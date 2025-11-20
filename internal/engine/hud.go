package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/player"
)

func drawHUD(screen *ebiten.Image, p *player.Player) {
	// HP bar
	drawBar(screen, 10, 10, 100, 8,
		float64(p.HP)/float64(p.MaxHP),
		color.RGBA{50, 200, 50, 255},
	)

	// XP bar
	ratio := 0.0
	if p.XPToNext > 0 {
		ratio = float64(p.XP) / float64(p.XPToNext)
	}
	drawBar(screen, 10, 24, 100, 6,
		ratio,
		color.RGBA{80, 160, 255, 255},
	)
}

func drawBar(screen *ebiten.Image, x, y, w, h int, ratio float64, fillColor color.Color) {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	// fundo
	bg := ebiten.NewImage(w, h)
	bg.Fill(color.RGBA{30, 30, 30, 200})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(bg, op)

	// preenchimento
	fw := int(float64(w) * ratio)
	if fw <= 0 {
		return
	}

	bar := ebiten.NewImage(fw, h)
	bar.Fill(fillColor)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(bar, op2)
}
