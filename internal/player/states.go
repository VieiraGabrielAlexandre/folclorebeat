package player

type PlayerState int

const (
	StateIdle PlayerState = iota
	StateWalk
	StateJump
	StatePunch
	StateKick
	StateAirKick // voadora
	StateWolf    // transformado em lobisomem
)
