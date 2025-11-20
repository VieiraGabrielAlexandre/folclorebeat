package powerups

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/combat"
)

type Type int

const (
	TypeWolfOrb Type = iota
)

type PowerUp struct {
	X, Y      float64
	BaseY     float64
	Type      Type
	Collected bool
}

func NewWolfOrb(x, y float64) *PowerUp {
	return &PowerUp{
		X:     x,
		Y:     y,
		BaseY: y,
		Type:  TypeWolfOrb,
	}
}

func (p *PowerUp) Update(frame int) {
	// animação boba – orb flutuando
	offset := math.Sin(float64(frame) * 0.1)
	p.Y = p.BaseY + offset*4
}

func (p *PowerUp) Draw(screen *ebiten.Image) {
	if p.Collected {
		return
	}

	img := ebiten.NewImage(12, 12)
	// orb azul/branco
	img.Fill(color.RGBA{180, 220, 255, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(img, op)
}

func (p *PowerUp) Hitbox() combat.Rect {
	return combat.Rect{
		X: p.X,
		Y: p.Y,
		W: 12,
		H: 12,
	}
}
