package world

// Collider represents a simple collision box in the world.
type Collider struct {
	X, Y float64
	W, H float64
}

func (c Collider) Intersects(o Collider) bool {
	return !(c.X+c.W < o.X || o.X+o.W < c.X || c.Y+c.H < o.Y || o.Y+o.H < c.Y)
}
