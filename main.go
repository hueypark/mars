package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	err := ebiten.Run(update, 800, 600, 1, "Hello world!")
	if err != nil {
		log.Fatalln(err)
	}
}

func update(screen *ebiten.Image) error {
	return ebitenutil.DebugPrint(screen, "Hello world!")
}
