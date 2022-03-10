package main

import "fmt"

func main() {
	/*
	 * &  = AND
	 * |  = OR
	 * ^  = XOR
	 * &^ = NOR
	 */

	a := 10 //1010
	b := 3  //0011

	fmt.Println(a & b)  // 0010 -> 2
	fmt.Println(a | b)  // 1011 -> 11
	fmt.Println(a ^ b)  // 1001 -> 9
	fmt.Println(a &^ b) // 0100 -> 8

	/*
	 * The following code does not work because this operands are only for 'bit' operations
	 */

	/*
		var c bool = true
		var d bool = false

		fmt.Println(c & d)
		fmt.Println(c | d)
		fmt.Println(c ^ d)
		fmt.Println(c &^ d)
	*/

}
