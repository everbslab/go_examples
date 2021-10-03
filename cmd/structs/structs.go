package main

import "fmt"

type A struct{}

type B struct {
	A
}

type C struct {
	A
}

func (a A) Foo() string {
	return "a"
}

func (b B) Foo() string {
	return "b"
}

func main() {
	b := B{}
	c := C{}

	fmt.Println(b.Foo(), c.Foo()) // b, a
}
