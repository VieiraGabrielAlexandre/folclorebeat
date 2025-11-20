package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/combat"
)

type Player struct {
	X, Y     float64
	VX, VY   float64
	OnGround bool
	State    PlayerState
	Facing   int
	IsWolf   bool

	// Combate / progressão
	HP, MaxHP      int
	Level          int
	XP             int
	XPToNext       int
	AttackPower    int
	AttackCooldown int // frames até poder atacar de novo
}

func NewPlayer() *Player {
	return &Player{
		X:              100,
		Y:              200,
		Facing:         1,
		State:          StateIdle,
		OnGround:       true,
		MaxHP:          10,
		HP:             10,
		Level:          1,
		XP:             0,
		XPToNext:       2, // mata 2 inimigos -> transforma
		AttackPower:    1,
		AttackCooldown: 0,
	}
}

func (p *Player) Update() {
	if p.AttackCooldown > 0 {
		p.AttackCooldown--
	}

	p.handleMovement()
	p.handleAttacks()
	p.applyPhysics()
}

func (p *Player) Draw(screen *ebiten.Image) {
	// temporário: desenha um retângulo como placeholder
	rect := ebiten.NewImage(24, 48)
	if p.IsWolf {
		rect.Fill(color.RGBA{150, 50, 50, 255}) // lobisomem = vermelho escuro
	} else {
		rect.Fill(color.RGBA{50, 200, 255, 255})
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(rect, op)
}

// Hitbox "corpo" do player (usado pra levar dano ou colisão futura)
func (p *Player) Hitbox() combat.Rect {
	return combat.Rect{
		X: p.X,
		Y: p.Y,
		W: 24,
		H: 48,
	}
}

// Hitbox do ataque atual (soco, chute, voadora)
func (p *Player) AttackHitbox() (combat.Rect, bool) {
	switch p.State {
	case StatePunch, StateKick, StateAirKick:
		// retângulo na frente do jogador
		w := 30.0
		h := 40.0
		x := p.X
		if p.Facing > 0 {
			x += 24 // à direita do player
		} else {
			x -= w // à esquerda do player
		}
		return combat.Rect{
			X: x,
			Y: p.Y,
			W: w,
			H: h,
		}, true
	default:
		return combat.Rect{}, false
	}
}

// XP / Level / Transformação

func (p *Player) GainXP(amount int) {
	p.XP += amount
	if p.XP >= p.XPToNext && !p.IsWolf {
		p.Level++
		p.XP = 0
		p.XPToNext += 2
		p.TransformToWolf()
	}
}

func (p *Player) TransformToWolf() {
	p.IsWolf = true
	p.State = StateWolf
	p.AttackPower = 3
}
