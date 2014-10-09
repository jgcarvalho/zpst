package zpst

import (
	"errors"
	"fmt"
)

type Residue struct {
	N     int
	Npdb  int
	ICode string
	Code  string
	Chain string
	Atoms []Atom
	Phi   float64
	Psi   float64
	Omega float64
	Chi1  float64
	Chi2  float64
	Chi3  float64
	Chi4  float64
}

func (r *Residue) AtomName(name string) (Atom, error) {
	for _, a := range r.Atoms {
		if a.Name == name {
			return a, nil
		}
	}
	var a Atom
	return a, errors.New(fmt.Sprintf("Atom not found: %s", name))
}

func (r Residue) String() string {
	return fmt.Sprintf("<Residue> Number: %d, pdbNumber: %d%s, Code: %s, Chain: %s", r.N, r.Npdb, r.ICode, r.Code, r.Chain)
}
