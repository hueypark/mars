package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/mars/conn"
	"github.com/hueypark/mars/game"
	"github.com/hueypark/mars/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

var (
	ball *game.Actor
)

func main() {
	conn.SendLogin(0)

	ball = game.NewActor(cp.Vector{})

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		ball.SetDesiredPosition(cp.Vector{X: float64(x), Y: float64(y)})
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

	game.ForEachNode(func(node *game.Node) {
		renderer.Reder(screen, node.Image(), node.Position())

		node.ForEachEdge(func(edge *game.Edge) {
			if toNode := game.GetNode(edge.To); toNode != nil {
				ebitenutil.DrawLine(
					screen,
					node.Position().X,
					node.Position().Y,
					toNode.Position().X,
					toNode.Position().Y,
					colornames.White)
			}
		})
	})

	game.EachActor(func(actor *game.Actor) {
		actor.Tick()

		renderer.Reder(screen, actor.Image(), vector.Vector{X: actor.Position().X, Y: actor.Position().Y})
	})

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
