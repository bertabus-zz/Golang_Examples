// Package go_sandbox_test provides ...
package go_sandbox_test

import "fmt"
import "math"

func ExampleHelloworld() {
	fmt.Println("hello world")
}

func ExampleArange(start, stop, step float64) []float64 {
	fmt.Println(arange(0., 10., 0.5))
}
