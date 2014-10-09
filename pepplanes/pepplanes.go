package pepplanes

import (
	"bitbucket.org/zgcarvalho/zpst"
	"fmt"
	"math"
)

func PepPlanes(p *zpst.Protein, fn string) {
	// var ref [3]float64
	// var nxt [3]float64
	for i := 1; i < len(p.Residues)-1; i++ {
		c1, _ := p.Residues[i-1].AtomName("C")
		o1, _ := p.Residues[i-1].AtomName("O")
		n1, _ := p.Residues[i].AtomName("N")
		c2, _ := p.Residues[i].AtomName("C")
		o2, _ := p.Residues[i].AtomName("O")
		n2, _ := p.Residues[i+1].AtomName("N")
		fmt.Println("Residuo ", i)
		pl1, pl2 := TransRot([3]zpst.Atom{c1, o1, n1}, [3]zpst.Atom{c2, o2, n2})
		fmt.Println(pl1)
		fmt.Println(pl2)
	}
}

func TransRot(ref [3]zpst.Atom, nxt [3]zpst.Atom) ([3]zpst.Atom, [3]zpst.Atom) {
	pl1, pl2 := trans(ref, nxt)
	pl1, pl2 = rotZ(pl1, pl2)
	pl1, pl2 = rotY(pl1, pl2)
	pl1, pl2 = rotX(pl1, pl2)
	return pl1, pl2
}

func trans(ref [3]zpst.Atom, nxt [3]zpst.Atom) ([3]zpst.Atom, [3]zpst.Atom) {
	move := ref[0].XYZ
	for i := range ref {
		ref[i].XYZ = [3]float64{ref[i].XYZ[0] - move[0], ref[i].XYZ[1] - move[1], ref[i].XYZ[2] - move[2]}
		nxt[i].XYZ = [3]float64{nxt[i].XYZ[0] - move[0], nxt[i].XYZ[1] - move[1], nxt[i].XYZ[2] - move[2]}
	}
	return ref, nxt
}

func dot(mrot [3][3]float64, coord [3]float64) (new_coord [3]float64) {
	new_coord[0] = mrot[0][0]*coord[0] + mrot[0][1]*coord[1] + mrot[0][2]*coord[2]
	new_coord[1] = mrot[1][0]*coord[0] + mrot[1][1]*coord[1] + mrot[1][2]*coord[2]
	new_coord[2] = mrot[2][0]*coord[0] + mrot[2][1]*coord[1] + mrot[2][2]*coord[2]
	return new_coord
}

func rotX(ref [3]zpst.Atom, nxt [3]zpst.Atom) ([3]zpst.Atom, [3]zpst.Atom) {
	ang := math.Atan2(ref[1].XYZ[1]-ref[0].XYZ[1], ref[1].XYZ[2]-ref[0].XYZ[2])
	sin_a, cos_a := math.Sincos(-ang)
	m_rot := [3][3]float64{[3]float64{1.0, 0.0, 0.0}, [3]float64{0.0, cos_a, sin_a}, [3]float64{0.0, -sin_a, cos_a}}
	for i := range ref {
		ref[i].XYZ = dot(m_rot, ref[i].XYZ)
		nxt[i].XYZ = dot(m_rot, nxt[i].XYZ)
	}
	return ref, nxt
}

func rotY(ref [3]zpst.Atom, nxt [3]zpst.Atom) ([3]zpst.Atom, [3]zpst.Atom) {
	ang := math.Atan2(ref[2].XYZ[2]-ref[0].XYZ[2], ref[2].XYZ[0]-ref[0].XYZ[0])
	sin_b, cos_b := math.Sincos(-ang)
	m_rot := [3][3]float64{[3]float64{cos_b, 0.0, -sin_b}, [3]float64{0.0, 1.0, 0.0}, [3]float64{sin_b, 0.0, cos_b}}
	for i := range ref {
		ref[i].XYZ = dot(m_rot, ref[i].XYZ)
		nxt[i].XYZ = dot(m_rot, nxt[i].XYZ)
	}
	return ref, nxt
}

func rotZ(ref [3]zpst.Atom, nxt [3]zpst.Atom) ([3]zpst.Atom, [3]zpst.Atom) {
	ang := math.Atan2(ref[2].XYZ[1]-ref[0].XYZ[1], ref[2].XYZ[0]-ref[0].XYZ[0])
	sin_c, cos_c := math.Sincos(ang)
	m_rot := [3][3]float64{[3]float64{cos_c, sin_c, 0.0}, [3]float64{-sin_c, cos_c, 0.0}, [3]float64{0.0, 0.0, 1.0}}
	for i := range ref {
		ref[i].XYZ = dot(m_rot, ref[i].XYZ)
		nxt[i].XYZ = dot(m_rot, nxt[i].XYZ)
	}
	return ref, nxt
}
