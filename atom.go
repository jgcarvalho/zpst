package zpst

import (
	"fmt"
)

type Atom struct {
	N       int
	Name    string
	Type    string
	XYZ     [3]float64
	Occ     float64
	Bfactor float64
	AltLoc  string
	Charge  string
}

func (a Atom) String() string {
	return fmt.Sprintf("<Atom> Number: %d, Name: %s, Type: %s, Coord (%.3f, %.3f, %.3f)", a.N, a.Name, a.Type, a.XYZ[0], a.XYZ[1], a.XYZ[2])
}
