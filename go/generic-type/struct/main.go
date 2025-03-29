package main

import "fmt"

type adder[T int | float32 | float64 | string] struct {
	a T
	b T
}

func (a adder[T]) add() T {
	return a.a + a.b
}

func main() {
	i := adder[int]{a: 1, b: 2}
	fmt.Println(i, i.add())

	s := adder[string]{a: "hello", b: "world"}
	fmt.Println(s, s.add())
}
