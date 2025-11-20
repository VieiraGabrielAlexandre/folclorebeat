package bosses

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/combat"
	"folclorebeat/internal/player"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Saci struct {
	X, Y          float64
	Width, Height float64
	HP, MaxHP     int
	Alive         bool

	teleportCooldown int
	frame            int
	baseY            float64
}

func NewSaci(x, y float64) *Saci {
	return &Saci{
		X:                x,
		Y:                y,
		baseY:            y,
		Width:            24,
		Height:           48,
		HP:               12,
		MaxHP:            12,
		Alive:            true,
		teleportCooldown: 90,
	}
}

func (s *Saci) Update(p *player.Player) {
	if !s.Alive {
		return
	}

	s.frame++
	if s.teleportCooldown > 0 {
		s.teleportCooldown--
	}

	offset := math.Sin(float64(s.frame) * 0.1)
	s.Y = s.baseY + offset*4

	if s.teleportCooldown <= 0 {
		dir := 1.0
		if rand.Intn(2) == 0 {
			dir = -1.0
		}
		dist := float64(40 + rand.Intn(80))
		newX := p.X + dir*dist

		if newX < 20 {
			newX = 20
		}
		if newX > 480-40 {
			newX = 480 - 40
		}
		s.X = newX

		s.teleportCooldown = 90
	}
}

func (s *Saci) Hitbox() combat.Rect {
	return combat.Rect{
		X: s.X,
		Y: s.Y,
		W: s.Width,
		H: s.Height,
	}
}

func (s *Saci) TakeDamage(d int) {
	if !s.Alive {
		return
	}
	s.HP -= d
	if s.HP <= 0 {
		s.HP = 0
		s.Alive = false
	}
}

func (s *Saci) Draw(screen *ebiten.Image) {
	if !s.Alive {
		return
	}

	img := ebiten.NewImage(int(s.Width), int(s.Height))
	img.Fill(color.RGBA{180, 30, 180, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(img, op)
}

func (s *Saci) IsAlive() bool {
	return s.Alive
}
