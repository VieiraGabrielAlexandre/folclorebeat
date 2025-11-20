package player

// Animation handles current animation state for the player.
type Animation struct {
	Name string
}

// Set switches the current animation.
func (a *Animation) Set(name string) { a.Name = name }
