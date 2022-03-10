package main

import "fmt"

func main() {

	a := 8 //2^3

	fmt.Println(a << 3) // 2^(3 * 3) = 2^6 // --> 64
	fmt.Println(a >> 3) // 2^(3 / 3) = 2^0 // --> 1

	/*
	 * i = 8
	 * j = 2
	 *
	 * i << j == i^j
	 * or (2^3)^3 = 2^(3 * 3)
	 */

}
