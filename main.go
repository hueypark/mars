package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jakecoffman/cp"
)

var (
	space *cp.Space
	dot   *ebiten.Image
	ball  *cp.Shape
)

func init() {
	dot, _ = ebiten.NewImage(10, 10, ebiten.FilterDefault)
	_ = dot.Fill(color.White)
}

func main() {
	space = cp.NewSpace()
	space.Iterations = 1

	ball = makeBall(100, 100)
	space.AddBody(ball.Body())
	space.AddShape(ball)

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		ball.Body().ApplyForceAtLocalPoint(cp.Vector{0, -100}, cp.Vector{0, 0})
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		ball.Body().ApplyForceAtLocalPoint(cp.Vector{0, 100}, cp.Vector{0, 0})
	}

	space.Step(1.0 / float64(ebiten.MaxTPS()))

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := screen.Fill(color.Black)
	if err != nil {
		return err
	}

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(200.0/255.0, 200.0/255.0, 200.0/255.0, 1)

	space.EachBody(func(body *cp.Body) {
		op.GeoM.Reset()
		op.GeoM.Translate(body.Position().X, body.Position().Y)
		err := screen.DrawImage(dot, op)
		if err != nil {
			log.Println(err)
		}
	})

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func makeBall(x, y float64) *cp.Shape {
	body := cp.NewBody(1.0, cp.INFINITY)
	body.SetPosition(cp.Vector{x, y})

	shape := cp.NewCircle(body, 0.95, cp.Vector{})
	shape.SetElasticity(0)
	shape.SetFriction(0)

	return shape
}
