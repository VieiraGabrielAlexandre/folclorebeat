package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Stage struct{}

func NewStage() *Stage {
	return &Stage{}
}

func (s *Stage) Draw(screen *ebiten.Image) {
	bg := ebiten.NewImage(480, 270)
	bg.Fill(color.RGBA{20, 20, 40, 255})
	screen.DrawImage(bg, nil)
}
