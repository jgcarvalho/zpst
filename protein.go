package zpst

import (
	"errors"
	"fmt"
)

type Protein struct {
	Name     string
	Chains   []string
	Residues []Residue
}

func (p *Protein) Residue(n int, chain string) (Residue, error) {
	for _, r := range p.Residues {
		if r.N == n && r.Chain == chain {
			return r, nil
		}
	}
	var r Residue
	return r, errors.New("Residue not found!")

}

func (p *Protein) checkChain(chain string) bool {
	for _, c := range p.Chains {
		if c == chain {
			return true
		}
	}
	return false
}

//Algo não esta funcionando como devia
func (p *Protein) SelectChainResidues(chain string) ([]*Residue, error) {
	if !p.checkChain(chain) {
		//var r []Residue
		return nil, errors.New(fmt.Sprintf("Chain not found: %q!\n", chain))
	}
	begin, end := len(p.Residues), 0
	for i, r := range p.Residues {
		if r.Chain == chain {
			if i < begin {
				begin = i
			}
			if i > end {
				end = i
			}
		}
	}

	var p_res []*Residue
	for _, res := range p.Residues[begin : end+1] {
		p_res = append(p_res, &res)
	}
	return p_res, nil
}

func (p *Protein) SelectAtomCoord(atomNumber int) ([3]float64, error) {
	for _, r := range p.Residues {
		for _, a := range r.Atoms {
			if a.N == atomNumber {
				return a.XYZ, nil
			}
		}
	}
	return [3]float64{}, errors.New(fmt.Sprintf("Atom not found: %q!\n", atomNumber))
}

func (p *Protein) Copy() Protein {
	var cp Protein
	cp.Name = p.Name
	copy(cp.Chains, p.Chains)
	cp.Residues = make([]Residue, len(p.Residues))
	for i, r := range p.Residues {
		cp.Residues[i].N = r.N
		cp.Residues[i].Npdb = r.Npdb
		cp.Residues[i].ICode = r.ICode
		cp.Residues[i].Code = r.Code
		cp.Residues[i].Chain = r.Chain
		cp.Residues[i].Phi = r.Phi
		cp.Residues[i].Psi = r.Psi
		cp.Residues[i].Omega = r.Omega
		cp.Residues[i].Chi1 = r.Chi1
		cp.Residues[i].Chi2 = r.Chi2
		cp.Residues[i].Chi3 = r.Chi3
		cp.Residues[i].Chi4 = r.Chi4

		cp.Residues[i].Atoms = make([]Atom, len(p.Residues[i].Atoms))
		for j, a := range r.Atoms {
			cp.Residues[i].Atoms[j].N = a.N
			cp.Residues[i].Atoms[j].Name = a.Name
			cp.Residues[i].Atoms[j].Type = a.Type
			cp.Residues[i].Atoms[j].XYZ = a.XYZ
			cp.Residues[i].Atoms[j].Occ = a.Occ
			cp.Residues[i].Atoms[j].Bfactor = a.Bfactor
			cp.Residues[i].Atoms[j].AltLoc = a.AltLoc
			cp.Residues[i].Atoms[j].Charge = a.Charge
		}
	}
	return cp

}

func (p Protein) String() string {
	return fmt.Sprintf("<Protein> Name: %s, Chains:%s, Total Residues:%d", p.Name, p.Chains, len(p.Residues))
}
