package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hueypark/mars/game"
	"github.com/jakecoffman/cp"
)

var (
	ball *game.Actor
	dot  *ebiten.Image
)

func init() {
	dot, _ = ebiten.NewImage(10, 10, ebiten.FilterDefault)
	_ = dot.Fill(color.White)
}

func main() {
	ball = game.NewActor(cp.Vector{})

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		ball.Body.ApplyForceAtLocalPoint(cp.Vector{0, -100}, cp.Vector{0, 0})
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		ball.Body.ApplyForceAtLocalPoint(cp.Vector{0, 100}, cp.Vector{0, 0})
	}

	game.Space.Step(1.0 / float64(ebiten.MaxTPS()))

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := screen.Fill(color.Black)
	if err != nil {
		return err
	}

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(200.0/255.0, 200.0/255.0, 200.0/255.0, 1)

	game.Space.EachBody(func(body *cp.Body) {
		op.GeoM.Reset()
		op.GeoM.Translate(body.Position().X, body.Position().Y)
		err := screen.DrawImage(dot, op)
		if err != nil {
			log.Println(err)
		}
	})

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
