package main

import (
	"AA_L42/pkg/mis"
	"fmt"
)

func main() {
	g := mis.Graph_NewRandGraph(1000, 0.1)

	ind := mis.MaximalIndependentSet(g)

	fmt.Println(ind)
	fmt.Println(len(ind))
}
