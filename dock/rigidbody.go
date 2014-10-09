package dock

import "math"

func vector(begin [3]float64, end [3]float64) [3]float64 {
	return [3]float64{end[0] - begin[0], end[1] - begin[1], end[2] - begin[2]}
}

func (p *Protein) RigidTrans(chain string, from [3]float64, to [3]float64) Protein {
	p_out := p.Copy()

	if chain == "" || chain == "All" {
		move := vector(from, to)
		for i, _ := range p_out.Residues {
			for j, a := range p_out.Residues[i].Atoms {
				p_out.Residues[i].Atoms[j].XYZ[0] = a.XYZ[0] + move[0]
				p_out.Residues[i].Atoms[j].XYZ[1] = a.XYZ[1] + move[1]
				p_out.Residues[i].Atoms[j].XYZ[2] = a.XYZ[2] + move[2]

			}
		}
	} else {
		move := vector(from, to)
		for i, r := range p_out.Residues {
			if r.Chain == chain {
				for j, a := range p_out.Residues[i].Atoms {
					p_out.Residues[i].Atoms[j].XYZ[0] = a.XYZ[0] + move[0]
					p_out.Residues[i].Atoms[j].XYZ[1] = a.XYZ[1] + move[1]
					p_out.Residues[i].Atoms[j].XYZ[2] = a.XYZ[2] + move[2]

				}
			}
		}
	}
	return p_out
}

func dot(mrot [3][3]float64, coord [3]float64) (new_coord [3]float64) {
	new_coord[0] = mrot[0][0]*coord[0] + mrot[0][1]*coord[1] + mrot[0][2]*coord[2]
	new_coord[1] = mrot[1][0]*coord[0] + mrot[1][1]*coord[1] + mrot[1][2]*coord[2]
	new_coord[2] = mrot[2][0]*coord[0] + mrot[2][1]*coord[1] + mrot[2][2]*coord[2]
	return new_coord
}

func rotX(coord [3]float64, a float64) [3]float64 {
	sin_a, cos_a := math.Sincos(a)
	m_rot := [3][3]float64{[3]float64{1.0, 0.0, 0.0}, [3]float64{0.0, cos_a, sin_a}, [3]float64{0.0, -sin_a, cos_a}}
	return dot(m_rot, coord)
}

func rotY(coord [3]float64, b float64) [3]float64 {
	sin_b, cos_b := math.Sincos(b)
	m_rot := [3][3]float64{[3]float64{cos_b, 0.0, -sin_b}, [3]float64{0.0, 1.0, 0.0}, [3]float64{sin_b, 0.0, cos_b}}
	return dot(m_rot, coord)
}

func rotZ(coord [3]float64, c float64) [3]float64 {
	sin_c, cos_c := math.Sincos(c)
	m_rot := [3][3]float64{[3]float64{cos_c, sin_c, 0.0}, [3]float64{-sin_c, cos_c, 0.0}, [3]float64{0.0, 0.0, 1.0}}
	return dot(m_rot, coord)
}

//func (p *Protein) RigidSpin(chain string, center [3]float64) Protein {
func (p *Protein) RigidSpin(chain string, center [3]float64, angles [3]float64) Protein {
	orig := [3]float64{0.0, 0.0, 0.0}
	p_out := p.RigidTrans("All", center, orig)

	if chain == "" || chain == "All" {
		for i, _ := range p_out.Residues {
			for j, _ := range p_out.Residues[i].Atoms {
				p_out.Residues[i].Atoms[j].XYZ = rotX(p_out.Residues[i].Atoms[j].XYZ, angles[0])
				p_out.Residues[i].Atoms[j].XYZ = rotY(p_out.Residues[i].Atoms[j].XYZ, angles[1])
				p_out.Residues[i].Atoms[j].XYZ = rotZ(p_out.Residues[i].Atoms[j].XYZ, angles[2])
			}
		}
	} else {
		for i, r := range p_out.Residues {
			if r.Chain == chain {
				for j, _ := range p_out.Residues[i].Atoms {
					p_out.Residues[i].Atoms[j].XYZ = rotX(p_out.Residues[i].Atoms[j].XYZ, angles[0])
					p_out.Residues[i].Atoms[j].XYZ = rotY(p_out.Residues[i].Atoms[j].XYZ, angles[1])
					p_out.Residues[i].Atoms[j].XYZ = rotZ(p_out.Residues[i].Atoms[j].XYZ, angles[2])
				}
			}
		}
	}
	return p_out.RigidTrans("All", orig, center)
}
