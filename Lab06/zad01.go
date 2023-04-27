package main

import (
	"flag"
	"fmt"
	"math"
)

type l struct {
	a, b, c float64
}

func r(v *l) (x [2]float64, s bool) {
	d := v.b*v.b - 4*v.a*v.c
	if d < 0 {
		return
	}
	s = true
	x[0] = (-v.b + math.Sqrt(d)) / (2 * v.a)
	x[1] = (-v.b - math.Sqrt(d)) / (2 * v.a)
	return
}
func main() {
	a := flag.Float64("a", 0, "")
	b := flag.Float64("b", 0, "")
	c := flag.Float64("c", 0, "")
	flag.Parse()
	w := l{*a, *b, *c}
	fmt.Print(r(&w))
}
