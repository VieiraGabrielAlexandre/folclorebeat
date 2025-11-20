package engine

import (
	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/bosses"
	"folclorebeat/internal/enemies"
	"folclorebeat/internal/player"
	"folclorebeat/internal/powerups"
	"folclorebeat/internal/world"
)

type Game struct {
	Player   *player.Player
	Enemies  []*enemies.Enemy
	Stage    *world.Stage
	PowerUps []*powerups.PowerUp

	Boss            bosses.Boss
	BossStage       int // 0 = nenhum, 1 = Saci, 2 = Cuca
	SaciRewardGiven bool
	CucaRewardGiven bool

	frame int
}

func NewGame() *Game {
	return &Game{
		Player: player.NewPlayer(),
		Stage:  world.NewStage(),
		Enemies: []*enemies.Enemy{
			enemies.NewZombie(260, 200),
			enemies.NewZombie(320, 200),
			enemies.NewVampire(380, 200),
		},
		PowerUps: []*powerups.PowerUp{},
	}
}

func (g *Game) Update() error {
	g.frame++

	g.Player.Update()

	// IA inimigos comuns
	for _, e := range g.Enemies {
		e.Update(g.Player)
	}

	// IA boss atual
	if g.Boss != nil && g.Boss.IsAlive() {
		g.Boss.Update(g.Player)
	}

	// ataques do player
	if atkRect, ok := g.Player.AttackHitbox(); ok {
		for _, e := range g.Enemies {
			if !e.Alive {
				continue
			}
			if atkRect.Intersects(e.Hitbox()) {
				e.TakeDamage(g.Player.AttackPower)
			}
		}
		if g.Boss != nil && g.Boss.IsAlive() {
			if atkRect.Intersects(g.Boss.Hitbox()) {
				g.Boss.TakeDamage(g.Player.AttackPower)
			}
		}
	}

	// drop de orbs
	for _, e := range g.Enemies {
		if e.Killed {
			orb := powerups.NewWolfOrb(e.X, e.Y-10)
			g.PowerUps = append(g.PowerUps, orb)
			e.Killed = false
		}
	}

	// orbs: update + coleta
	for _, p := range g.PowerUps {
		if p.Collected {
			continue
		}
		p.Update(g.frame)

		if p.Hitbox().Intersects(g.Player.Hitbox()) {
			p.Collected = true
			g.Player.GainXP(1)
		}
	}

	// spawn do Saci
	if g.BossStage == 0 {
		allDead := true
		for _, e := range g.Enemies {
			if e.Alive {
				allDead = false
				break
			}
		}
		if allDead {
			g.Boss = bosses.NewSaci(380, 200)
			g.BossStage = 1
		}
	}

	// pós-Saci: dar XP e criar Cuca
	if g.BossStage == 1 && g.Boss != nil && !g.Boss.IsAlive() && !g.SaciRewardGiven {
		g.Player.GainXP(2)
		g.SaciRewardGiven = true

		g.Boss = bosses.NewCuca(360, 160)
		g.BossStage = 2
	}

	// pós-Cuca: recompensa final
	if g.BossStage == 2 && g.Boss != nil && !g.Boss.IsAlive() && !g.CucaRewardGiven {
		g.Player.GainXP(3)
		g.CucaRewardGiven = true
		// aqui depois podemos colocar “fase completa”
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Stage.Draw(screen)
	g.Player.Draw(screen)

	for _, e := range g.Enemies {
		e.Draw(screen)
	}

	for _, p := range g.PowerUps {
		p.Draw(screen)
	}

	if g.Boss != nil && g.Boss.IsAlive() {
		g.Boss.Draw(screen)
	}

	drawHUD(screen, g.Player)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 480, 270
}
