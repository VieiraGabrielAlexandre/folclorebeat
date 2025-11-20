package enemies

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/player"
)

type Kind int

const (
	KindZombie Kind = iota
	KindVampire
)

type Enemy struct {
	X, Y          float64
	VX            float64
	Width, Height float64
	HP, MaxHP     int
	Kind          Kind
	Alive         bool
}

// Fábricas

func NewZombie(x, y float64) *Enemy {
	return &Enemy{
		X:      x,
		Y:      y,
		VX:     0.4,
		Width:  20,
		Height: 40,
		HP:     3,
		MaxHP:  3,
		Kind:   KindZombie,
		Alive:  true,
	}
}

func NewVampire(x, y float64) *Enemy {
	return &Enemy{
		X:      x,
		Y:      y,
		VX:     0.8,
		Width:  20,
		Height: 40,
		HP:     4,
		MaxHP:  4,
		Kind:   KindVampire,
		Alive:  true,
	}
}

// Lógica de IA simples: anda na direção do player

func (e *Enemy) Update(p *player.Player) {
	if !e.Alive {
		return
	}

	dx := p.X - e.X
	if math.Abs(dx) > 2 {
		if dx > 0 {
			e.X += e.VX
		} else {
			e.X -= e.VX
		}
	}
}

// Desenho placeholder (cores diferentes por tipo)

func (e *Enemy) Draw(screen *ebiten.Image) {
	if !e.Alive {
		return
	}

	img := ebiten.NewImage(int(e.Width), int(e.Height))

	switch e.Kind {
	case KindZombie:
		img.Fill(color.RGBA{80, 200, 80, 255}) // verde
	case KindVampire:
		img.Fill(color.RGBA{200, 80, 80, 255}) // vermelho
	default:
		img.Fill(color.RGBA{150, 150, 150, 255})
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(img, op)
}
