package bosses

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/combat"
	"folclorebeat/internal/player"
)

type fireball struct {
	X, Y   float64
	VX, VY float64
	Alive  bool
}

type Cuca struct {
	X, Y          float64
	Width, Height float64
	HP, MaxHP     int
	Alive         bool

	frame         int
	baseY         float64
	shootCooldown int

	projectiles []*fireball
}

func NewCuca(x, y float64) *Cuca {
	return &Cuca{
		X:             x,
		Y:             y,
		baseY:         y,
		Width:         32,
		Height:        52,
		HP:            18,
		MaxHP:         18,
		Alive:         true,
		shootCooldown: 45,
		projectiles:   []*fireball{},
	}
}

func (c *Cuca) Update(p *player.Player) {
	if !c.Alive {
		return
	}

	c.frame++
	if c.shootCooldown > 0 {
		c.shootCooldown--
	}

	// flutuar
	offset := math.Sin(float64(c.frame) * 0.07)
	c.Y = c.baseY + offset*5

	// 🐊 andar pela fase seguindo o player (horizontal)
	moveSpeed := 0.6
	dx := (p.X + 12) - (c.X + c.Width/2)
	if math.Abs(dx) > 2 {
		if dx > 0 {
			c.X += moveSpeed
		} else {
			c.X -= moveSpeed
		}
	}

	// limita a área de movimento
	if c.X < 40 {
		c.X = 40
	}
	if c.X > 480-80 {
		c.X = 480 - 80
	}

	// 🔥 atira fireballs diagonais em direção ao player
	if c.shootCooldown <= 0 {
		// posição da Cuca (centro)
		cx := c.X + c.Width/2
		cy := c.Y + c.Height/2

		// posição do player (centro aproximado)
		px := p.X + 12
		py := p.Y + 24

		dirX := px - cx
		dirY := py - cy
		dist := math.Hypot(dirX, dirY)
		if dist == 0 {
			dist = 1
		}

		dirX /= dist
		dirY /= dist

		speed := 2.0 + rand.Float64()*0.8

		fb := &fireball{
			X:     cx,
			Y:     cy,
			VX:    dirX * speed,
			VY:    dirY * speed,
			Alive: true,
		}
		c.projectiles = append(c.projectiles, fb)
		c.shootCooldown = 45
	}

	// atualiza projéteis
	active := make([]*fireball, 0, len(c.projectiles))
	for _, fb := range c.projectiles {
		if !fb.Alive {
			continue
		}

		fb.X += fb.VX
		fb.Y += fb.VY

		// se sair da tela, morre
		if fb.X < -10 || fb.X > 480+10 || fb.Y < -10 || fb.Y > 270+10 {
			fb.Alive = false
			continue
		}

		// colisão com player
		fbRect := combat.Rect{X: fb.X, Y: fb.Y, W: 6, H: 6}
		if fbRect.Intersects(p.Hitbox()) {
			p.TakeDamage(1)
			fb.Alive = false
			continue
		}

		active = append(active, fb)
	}
	c.projectiles = active
}

func (c *Cuca) Hitbox() combat.Rect {
	return combat.Rect{
		X: c.X,
		Y: c.Y,
		W: c.Width,
		H: c.Height,
	}
}

func (c *Cuca) TakeDamage(d int) {
	if !c.Alive {
		return
	}
	c.HP -= d
	if c.HP <= 0 {
		c.HP = 0
		c.Alive = false
	}
}

func (c *Cuca) Draw(screen *ebiten.Image) {
	if !c.Alive {
		return
	}

	// corpo da Cuca
	img := ebiten.NewImage(int(c.Width), int(c.Height))
	img.Fill(color.RGBA{20, 120, 40, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.X, c.Y)
	screen.DrawImage(img, op)

	// fireballs
	for _, fb := range c.projectiles {
		if !fb.Alive {
			continue
		}
		b := ebiten.NewImage(6, 6)
		b.Fill(color.RGBA{255, 200, 60, 255})
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(fb.X, fb.Y)
		screen.DrawImage(b, op2)
	}
}

func (c *Cuca) IsAlive() bool {
	return c.Alive
}
