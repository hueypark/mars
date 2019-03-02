package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

var (
	actorImage *ebiten.Image
	nodeImage  *ebiten.Image
)

func init() {
	actorImage, _ = ebiten.NewImage(10, 10, ebiten.FilterDefault)
	if err := actorImage.Fill(colornames.Gray); err != nil {
		log.Fatal(err)
	}

	nodeImage, _ = ebiten.NewImage(30, 30, ebiten.FilterDefault)
	if err := nodeImage.Fill(color.White); err != nil {
		log.Fatal(err)
	}
}
