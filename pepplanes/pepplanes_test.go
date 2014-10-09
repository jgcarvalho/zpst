package pepplanes

import (
	"bitbucket.org/zgcarvalho/zpst"
	"fmt"
	"testing"
)

var p1 = [3]zpst.Atom{
	zpst.Atom{Name: "C", XYZ: [3]float64{28.444, -9.870, 37.349}},
	zpst.Atom{Name: "O", XYZ: [3]float64{27.970, -8.753, 37.553}},
	zpst.Atom{Name: "N", XYZ: [3]float64{29.158, -10.144, 36.237}},
}
var p2 = [3]zpst.Atom{
	zpst.Atom{Name: "C", XYZ: [3]float64{28.139, -8.768, 34.422}},
	zpst.Atom{Name: "O", XYZ: [3]float64{27.279, -9.621, 34.245}},
	zpst.Atom{Name: "N", XYZ: [3]float64{28.046, -7.520, 33.970}},
}

func TestTrans(t *testing.T) {
	fmt.Println("translation")
	t1, t2 := trans(p1, p2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("\nrotation Z")
	t1, t2 = rotZ(t1, t2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("\nrotation Y")
	t1, t2 = rotY(t1, t2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("\nrotation X")
	t1, t2 = rotX(t1, t2)
	fmt.Println(t1)
	fmt.Println(t2)
}
