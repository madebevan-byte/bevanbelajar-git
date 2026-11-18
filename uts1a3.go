package main

import "fmt"

func main() {
	a := 22
	b := 4
	c := 2.5

	p := a / b
	q := a % b
	r := float64(a / b)
	s := float64(a) / float64(b)
	u := int(c*float64(b)) + a%3
	t := (a%2 == 1) && (float64(b) > c)

	fmt.Println("p =", p)
	fmt.Println("q =", q)
	fmt.Println("r =", r)
	fmt.Println("s =", s)
	fmt.Println("u =", u)
	fmt.Println("t =", t)
}
