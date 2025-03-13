package main

import (
	"fmt"
	"os"
	"math"
	"github.com/KostianDev/sdmt-lab1/solver"
	"github.com/KostianDev/sdmt-lab1/interactive"
	"github.com/KostianDev/sdmt-lab1/noninteractive"
)

func main() {
	var a, b, c float64
	var x1, x2 float64

	switch len(os.Args) {
	case 1:
		a, b, c = interactive.RunInteractiveMode()
	case 2:
		a, b, c = noninteractive.RunNonInteractiveMode(os.Args[1])
	default:
		fmt.Println("Error: invalid number of arguments")
		os.Exit(1)
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		os.Exit(1)
	}

	fmt.Printf("The equation is: (%g)x^2 + (%g)x + (%g) = 0\n", a, b, c)

	x1, x2 = solver.SolveQuadEquation(a, b, c)
	if math.IsNaN(x1) || math.IsNaN(x2) {
		fmt.Println("Error: Equation has no real roots")
		os.Exit(1)
	} else if x1 == x2 {
		fmt.Printf("There is only one root:\nx = %f\n", x1)
	} else {
		fmt.Printf("The roots are:\nx1 = %f, x2 = %f\n", x1, x2)
	}

	os.Exit(0)
}