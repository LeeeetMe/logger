package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = 6
	d = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
