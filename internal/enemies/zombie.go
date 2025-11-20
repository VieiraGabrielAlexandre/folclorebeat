package enemies

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	X, Y float64
}

func NewZombie(x, y float64) *Enemy {
	return &Enemy{X: x, Y: y}
}

func NewVampire(x, y float64) *Enemy {
	return &Enemy{X: x, Y: y}
}

func (e *Enemy) Update(player interface{}) {}

func (e *Enemy) Draw(screen *ebiten.Image) {
	rect := ebiten.NewImage(20, 40)
	rect.Fill(color.RGBA{80, 200, 80, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(rect, op)
}
