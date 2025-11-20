package combat

// Hurtbox represents an area that can take damage from hitboxes.
type Hurtbox struct {
	X, Y   float64
	W, H   float64
	Health int
}

// ApplyDamage reduces health by the specified amount, not going below zero.
func (h *Hurtbox) ApplyDamage(d Damage) {
	if d.Amount <= 0 || h.Health <= 0 {
		return
	}
	if h.Health-d.Amount < 0 {
		h.Health = 0
		return
	}
	h.Health -= d.Amount
}
