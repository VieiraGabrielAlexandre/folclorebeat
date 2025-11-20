package player

import "github.com/hajimehoshi/ebiten/v2"

func (p *Player) handleAttacks() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.State = StatePunch
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if !p.OnGround {
			p.State = StateAirKick // voadora
		} else {
			p.State = StateKick
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyF) {
		p.TransformToWolf()
	}
}

func (p *Player) TransformToWolf() {
	p.IsWolf = true
	p.State = StateWolf
}
