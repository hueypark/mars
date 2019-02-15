package game

import "github.com/jakecoffman/cp"

type Actor struct {
	Body *cp.Body
}

func NewActor(position cp.Vector) *Actor {
	actor := &Actor{}

	actor.Body = cp.NewBody(1.0, cp.INFINITY)
	actor.Body.SetPosition(position)

	shape := cp.NewCircle(actor.Body, 10.0, cp.Vector{})
	shape.SetElasticity(0)
	shape.SetFriction(0)

	Space.AddBody(actor.Body)
	Space.AddShape(shape)

	return actor
}

func (actor *Actor) Position() cp.Vector {
	return actor.Body.Position()
}
