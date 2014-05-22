// Package go_sandbox_test provides ...
package go_sandbox_test

import "fmt"
import "math"

func ExampleHelloworld() {
	fmt.Println("hello world")
}

// from http://stackoverflow.com/questions/19906544/
func ExampleArange(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N)
	for x := range rnge {
		rnge[x] = start + step*float64(x)
	}
	return rnge
}
