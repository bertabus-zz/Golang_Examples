// Package go_sandbox_test provides ...
package go_sandbox_test

import "fmt"
import "math"
import "github.com/bertabus/Golang_Examples"

func ExampleHelloworld() {
	fmt.Println("hello world")
}

func ExampleArange(start, stop, step float64) []float64 {
	fmt.Println(Arange(0., 10., 0.5))
}
