package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y     float64
	VX, VY   float64
	OnGround bool
	State    PlayerState
	Facing   int
	IsWolf   bool
}

func NewPlayer() *Player {
	return &Player{
		X:        100,
		Y:        200,
		Facing:   1,
		State:    StateIdle,
		OnGround: true,
	}
}

func (p *Player) Update() {
	p.handleMovement()
	p.handleAttacks()
	p.applyPhysics()
}

func (p *Player) Draw(screen *ebiten.Image) {
	// temporário: desenha um retângulo como placeholder
	rect := ebiten.NewImage(24, 48)
	if p.IsWolf {
		rect.Fill(color.RGBA{150, 50, 50, 255}) // cor diferente pro lobisomem
	} else {
		rect.Fill(color.RGBA{50, 200, 255, 255})
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(rect, op)
}
