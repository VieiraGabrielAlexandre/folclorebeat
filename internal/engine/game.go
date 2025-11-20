package engine

import (
	"github.com/hajimehoshi/ebiten/v2"

	_ "folclorebeat/internal/combat"
	"folclorebeat/internal/enemies"
	"folclorebeat/internal/player"
	"folclorebeat/internal/world"
)

type Game struct {
	Player  *player.Player
	Enemies []*enemies.Enemy
	Stage   *world.Stage
}

func NewGame() *Game {
	return &Game{
		Player: player.NewPlayer(),
		Stage:  world.NewStage(),
		Enemies: []*enemies.Enemy{
			enemies.NewZombie(300, 200),
			enemies.NewVampire(350, 200),
		},
	}
}

func (g *Game) Update() error {
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

	// Recompensa de XP por inimigos mortos
	for _, e := range g.Enemies {
		if e.Killed {
			g.Player.GainXP(e.XPReward)
			e.Killed = false
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
}

func (g *Game) Layout() (int, int) {
	return 480, 270
}
