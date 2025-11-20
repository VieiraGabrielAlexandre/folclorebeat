package enemies

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/combat"
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

	// recompensa / controle de morte
	XPReward int
	Killed   bool
}

// Fábricas

func NewZombie(x, y float64) *Enemy {
	return &Enemy{
		X:        x,
		Y:        y,
		VX:       0.4,
		Width:    20,
		Height:   40,
		HP:       3,
		MaxHP:    3,
		Kind:     KindZombie,
		Alive:    true,
		XPReward: 1,
	}
}

func NewVampire(x, y float64) *Enemy {
	return &Enemy{
		X:        x,
		Y:        y,
		VX:       0.8,
		Width:    20,
		Height:   40,
		HP:       4,
		MaxHP:    4,
		Kind:     KindVampire,
		Alive:    true,
		XPReward: 1,
	}
}

// AI simples: anda na direção do player

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

// Hitbox do inimigo

func (e *Enemy) Hitbox() combat.Rect {
	return combat.Rect{
		X: e.X,
		Y: e.Y,
		W: e.Width,
		H: e.Height,
	}
}

// Sofrer dano

func (e *Enemy) TakeDamage(d int) {
	if !e.Alive {
		return
	}

	e.HP -= d
	if e.HP <= 0 {
		e.HP = 0
		e.Alive = false
		e.Killed = true // importante pro drop do power-up
	}
}

// Desenho placeholder

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
