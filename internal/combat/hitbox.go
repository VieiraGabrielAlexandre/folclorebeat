package combat

// Hitbox represents an area that can hit others, causing damage.
type Hitbox struct {
	Active bool
	X, Y   float64
	W, H   float64
	Damage int
}
