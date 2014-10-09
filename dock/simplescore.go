package dock

//import "fmt"

func (p *Protein) selectChainAtomCoord(c string) (bb_atoms [][3]float64, sc_atoms [][3]float64) {
	for _, r := range p.Residues {
		if r.Chain == c {
			for _, a := range r.Atoms {
				if a.Name == "CA" || a.Name == "N" || a.Name == "C" || a.Name == "O" {
					bb_atoms = append(bb_atoms, a.XYZ)
				} else {
					sc_atoms = append(sc_atoms, a.XYZ)
				}

			}
		}
	}
	return
}

func (p *Protein) ScoreDockSimple(c1 string, c2 string) float64 {
	bb_atoms_c1, sc_atoms_c1 := p.selectChainAtomCoord(c1)
	bb_atoms_c2, sc_atoms_c2 := p.selectChainAtomCoord(c2)
	atoms_c2 := bb_atoms_c2
	atoms_c2 = append(atoms_c2, sc_atoms_c2...)
	//atoms_c2 := append(bb_atoms_c2, sc_atoms_c2)

	score := 0.0
	for _, a := range bb_atoms_c1 {
		for _, b := range bb_atoms_c2 {
			if distance(a, b) < 3.5 {
				score += 5.0
			}
		}
	}

	for _, a := range bb_atoms_c1 {
		for _, b := range sc_atoms_c2 {
			if distance(a, b) < 3.5 {
				score += 3.0
			}
		}
	}

	for _, a := range sc_atoms_c1 {
		for _, b := range bb_atoms_c2 {
			if distance(a, b) < 3.5 {
				score += 3.0
			}
		}
	}

	for _, a := range sc_atoms_c1 {
		for _, b := range sc_atoms_c2 {
			if distance(a, b) < 3.5 {
				score += 1.0
			}
		}
	}

	return score
}
