package main

import (
	"bitbucket.org/zgcarvalho/zpst"
	"bitbucket.org/zgcarvalho/zpst/pepplanes"
	"fmt"
)

func main() {
	prot, err := zpst.LoadFromPDBFile("/home/zeh/Downloads/1A04.pdb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prot)
	pepplanes.PepPlanes(prot, "/home/zeh/Downloads/1A04.pp")

}
