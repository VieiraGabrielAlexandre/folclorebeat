package player

import "github.com/hajimehoshi/ebiten/v2"

func (p *Player) handleAttacks() {
	// não pode atacar enquanto está "em recarga"
	if p.AttackCooldown > 0 {
		return
	}

	// Soco (A)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.State = StatePunch
		p.AttackCooldown = 12 // ~curta duração
		return
	}

	// Chute (S) ou voadora se estiver no ar
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if !p.OnGround {
			p.State = StateAirKick
		} else {
			p.State = StateKick
		}
		p.AttackCooldown = 16
		return
	}

	// Transformação manual (debug) - ainda mantemos
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		p.TransformToWolf()
	}
}
