package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
)

var op ebiten.DrawImageOptions

// Render renders object.
func Reder(screen *ebiten.Image, img *ebiten.Image, position vector.Vector) {
	op.GeoM.Reset()
	width, height := img.Size()
	op.GeoM.Translate(position.X-(float64(width)*0.5), position.Y-(float64(height)*0.5))
	err := screen.DrawImage(img, &op)
	if err != nil {
		log.Println(err)
	}
}
