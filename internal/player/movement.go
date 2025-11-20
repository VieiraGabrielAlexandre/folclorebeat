package player

import "github.com/hajimehoshi/ebiten/v2"

func (p *Player) handleMovement() {
	moving := false

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.X += 2
		p.Facing = 1
		moving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.X -= 2
		p.Facing = -1
		moving = true
	}

	if moving && p.OnGround {
		p.State = StateWalk
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.OnGround {
		p.VY = -6
		p.OnGround = false
		p.State = StateJump
	}
}

func (p *Player) applyPhysics() {
	if !p.OnGround {
		p.VY += 0.25 // gravidade
		p.Y += p.VY

		if p.Y >= 200 {
			p.Y = 200
			p.OnGround = true
			p.VY = 0
			p.State = StateIdle
		}
	}
}
