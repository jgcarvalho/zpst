package dock

import (
	"math"
	"math/rand"
	"time"
	//"fmt"
)

func distance(a [3]float64, b [3]float64) float64 {
	return math.Sqrt(math.Pow(a[0]-b[0], 2) + math.Pow(a[1]-b[1], 2) + math.Pow(a[2]-b[2], 2))
}

func surfPoint(center [3]float64, radius float64) [3]float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	s_sin, s_cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	t_sin, t_cos := math.Sincos(rand.Float64() * 2 * math.Pi)

	dx := radius * s_cos * t_sin
	dy := radius * s_sin * t_sin
	dz := radius * t_cos

	return [3]float64{center[0] - dx, center[1] - dy, center[2] - dz}

}

func (p *Protein) AnchorPointsAtom(natom int, npoints int, radius float64) [][3]float64 {
	coordRef, _ := p.SelectAtomCoord(natom)
	var coordNearAtoms [][3]float64

	for _, r := range p.Residues {
		for _, a := range r.Atoms {
			if distance(coordRef, a.XYZ) < (radius*2.0+1.0) && a.N != natom {
				coordNearAtoms = append(coordNearAtoms, a.XYZ)
			}
		}
	}

	anchorPoints := make([][3]float64, npoints)
	for i := 0; i < npoints; {
		point := surfPoint(coordRef, radius)
		status := true
		for _, near := range coordNearAtoms {
			if distance(point, near) < (radius + 0.5) {
				status = false
			}
		}
		if status {
			anchorPoints[i] = point
			i++
		}

	}
	return anchorPoints
}

func (p *Protein) BestAnchorPointAtom(natom int, npoints int, radius float64) [3]float64 {
	anchorPoints := p.AnchorPointsAtom(natom, npoints, radius)

	var bestAnchor [3]float64
	rms_min := math.MaxFloat64

	for _, best := range anchorPoints {
		sum := 0.0
		for _, others := range anchorPoints {
			sum += math.Pow(distance(best, others), 2)
		}
		rms := math.Sqrt(sum / float64(len(anchorPoints)))
		if rms < rms_min {
			rms_min = rms
			bestAnchor = best
		}

	}
	return bestAnchor
}
