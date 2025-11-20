package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"folclorebeat/internal/engine"
)

func main() {
	game := engine.NewGame()

	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Folclore Beat 'em Up — Rise of the Lobisomem")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
