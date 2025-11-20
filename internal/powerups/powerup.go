package powerups

// PowerUp represents a collectible that grants an effect.
type PowerUp struct {
	Name string
}

// Apply would apply the power-up effect to a target.
func (p *PowerUp) Apply() {}
