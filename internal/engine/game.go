package engine

import (
	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/bosses"
	_ "folclorebeat/internal/combat"
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

	Boss        *bosses.Saci
	BossSpawned bool
	BossXPGiven bool

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

	// IA dos inimigos comuns
	for _, e := range g.Enemies {
		e.Update(g.Player)
	}

	// IA do boss, se existir
	if g.BossSpawned && g.Boss != nil && g.Boss.Alive {
		g.Boss.Update(g.Player)
	}

	// Combate: ataque do player contra inimigos
	if atkRect, ok := g.Player.AttackHitbox(); ok {
		for _, e := range g.Enemies {
			if !e.Alive {
				continue
			}
			if atkRect.Intersects(e.Hitbox()) {
				e.TakeDamage(g.Player.AttackPower)
			}
		}

		// ataque também acerta o boss
		if g.BossSpawned && g.Boss != nil && g.Boss.Alive {
			if atkRect.Intersects(g.Boss.Hitbox()) {
				g.Boss.TakeDamage(g.Player.AttackPower)
			}
		}
	}

	// Drop de power-up quando inimigo morre
	for _, e := range g.Enemies {
		if e.Killed {
			orb := powerups.NewWolfOrb(e.X, e.Y-10)
			g.PowerUps = append(g.PowerUps, orb)
			e.Killed = false
		}
	}

	// Atualiza e verifica coleta de power-ups
	for _, p := range g.PowerUps {
		if p.Collected {
			continue
		}
		p.Update(g.frame)

		if p.Hitbox().Intersects(g.Player.Hitbox()) {
			p.Collected = true
			switch p.Type {
			case powerups.TypeWolfOrb:
				g.Player.GainXP(1)
			}
		}
	}

	// Spawn do boss Saci:
	// quando todos os inimigos comuns estiverem mortos e ainda não tiver boss
	if !g.BossSpawned {
		allDead := true
		for _, e := range g.Enemies {
			if e.Alive {
				allDead = false
				break
			}
		}
		if allDead {
			g.Boss = bosses.NewSaci(380, 200)
			g.BossSpawned = true
		}
	}

	// Recompensa de XP ao matar o boss
	if g.BossSpawned && g.Boss != nil && !g.Boss.Alive && !g.BossXPGiven {
		g.Player.GainXP(3) // boss dá mais XP
		g.BossXPGiven = true
	}

	// Dano de contato do boss no player
	if g.BossSpawned && g.Boss != nil && g.Boss.Alive {
		if g.Boss.Hitbox().Intersects(g.Player.Hitbox()) {
			g.Player.TakeDamage(1)
		}
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

	if g.BossSpawned && g.Boss != nil {
		g.Boss.Draw(screen)
	}

	drawHUD(screen, g.Player)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 480, 270
}
