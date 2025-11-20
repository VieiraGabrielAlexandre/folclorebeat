package bosses

import (
	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/combat"
	"folclorebeat/internal/player"
)

type Boss interface {
	Update(p *player.Player)
	Draw(screen *ebiten.Image)
	Hitbox() combat.Rect
	TakeDamage(d int)
	IsAlive() bool
}
