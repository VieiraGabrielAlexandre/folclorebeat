package engine

import (
	"github.com/hajimehoshi/ebiten/v2"

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

	frame int
}

func NewGame() *Game {
	return &Game{
		Player: player.NewPlayer(),
		Stage:  world.NewStage(),
		Enemies: []*enemies.Enemy{
			enemies.NewZombie(300, 200),
			enemies.NewVampire(350, 200),
		},
		PowerUps: []*powerups.PowerUp{},
	}
}

func (g *Game) Update() error {
	g.frame++

	g.Player.Update()

	// IA dos inimigos
	for _, e := range g.Enemies {
		e.Update(g.Player)
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
			// orbe de lobisomem => dá XP
			switch p.Type {
			case powerups.TypeWolfOrb:
				g.Player.GainXP(1)
			}
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

	drawHUD(screen, g.Player)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 480, 270
}
