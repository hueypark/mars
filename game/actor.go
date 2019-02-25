package game

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
)

var image *ebiten.Image

func init() {
	image, _ = ebiten.NewImage(10, 10, ebiten.FilterDefault)
	err := image.Fill(color.White)
	if err != nil {
		log.Fatal(err)
	}
}

// Actor represents basic game actor.
type Actor struct {
	Body            *cp.Body
	desiredPosition cp.Vector
}

// NewActor creates new actor.
func NewActor(position cp.Vector) *Actor {
	actor := &Actor{}

	actor.Body = cp.NewBody(1.0, cp.INFINITY)
	actor.Body.UserData = actor
	actor.Body.SetPosition(position)

	shape := cp.NewCircle(actor.Body, 10.0, cp.Vector{})
	shape.SetElasticity(0)
	shape.SetFriction(0)

	Space.AddBody(actor.Body)
	Space.AddShape(shape)

	return actor
}

// Tick ticks actor.
func (actor *Actor) Tick() {
	sub := actor.desiredPosition.Sub(actor.Position())
	dir := sub.Normalize()
	if math.IsNaN(dir.X) || math.IsNaN(dir.Y) {
		return
	}

	actor.Body.ApplyForceAtLocalPoint(dir.Mult(actor.DesiredSpeed(sub.Length())), cp.Vector{})

	actor.Body.SetVelocityVector(actor.Body.Velocity().Mult(0.99))
}

// Position returns actor position.
func (actor *Actor) Position() cp.Vector {
	return actor.Body.Position()
}

// SetDesiredPosition returns desired position.
func (actor *Actor) SetDesiredPosition(position cp.Vector) {
	actor.desiredPosition = position
}

// DesiredSpeed represents desired speed for desired position.
func (actor *Actor) DesiredSpeed(remainLength float64) float64 {
	return math.Min(remainLength, 100)
}

// Image returns image.
func (actor *Actor) Image() *ebiten.Image {
	return image
}