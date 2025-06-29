package main

import (
	"log"

	"chess/chess"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := chess.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(chess.ScreenWidth, chess.ScreenHeight)
	ebiten.SetWindowTitle("CHESS")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
