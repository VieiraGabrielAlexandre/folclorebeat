package engine

import (
	"github.com/hajimehoshi/ebiten/v2"

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
			enemies.NewZombie(300, 300),
			enemies.NewVampire(500, 280),
		},
	}
}

func (g *Game) Update() error {
	g.Player.Update()

	for _, e := range g.Enemies {
		e.Update(g.Player)
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

func (g *Game) Layout(w, h int) (int, int) {
	return 480, 270
}
