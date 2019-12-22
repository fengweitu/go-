package main

import "fmt"

type Man struct {
	age, height int
}

func main() {
	const NAME string = "STEVEN"
	//man := Man{age:30,height:177}
	//const person  = man

	const (
		x = "A"
		y
		a = 10
		b
		c
	)
	fmt.Println(x, y, a, b, c)

	const (
		L = iota
		M = iota
		N = iota
	)
	fmt.Println(L, M, N)
	const (
		L1 = iota
		M1
		N1
	)
	fmt.Println(L1, M1, N1)
	const (
		L2 = "steven"
		M2 = iota
		N2 = 123
	)
	fmt.Println(L2, M2, N2)

	const (
		i = 1 << iota //1*2^0
		j = 3 << iota //3*2^1
		k             // = 3<<iota
		l             // = 3<<iota
	)
	fmt.Println(i, j, k, l)

	const (
		a1 = '一'
		b1
		c1 = iota
		d1
	)
	fmt.Println(a1, b1, c1, d1)
	const name  = iota
	fmt.Println(name)

}
