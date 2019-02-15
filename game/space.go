package game

import "github.com/jakecoffman/cp"

var Space *cp.Space

func init() {
	Space = cp.NewSpace()
	Space.Iterations = 1
}
