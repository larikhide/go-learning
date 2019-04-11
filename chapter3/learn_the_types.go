package main

import "fmt"

func main() {
	// uint8 = byte, uint16, uint32 - "unsigned integer" can be only positive or zero values
	// int8, int16, int32 = rine - "signed integer"
	// uint, int, uintprt - is machine depended types
	// ATTENTION! if you want to use integer just use the "int" type
	// "NaN" - is not a number such as "0/0" and +∞ or -∞
	// float32 and float64 - is single and double precision real numbers
	// complex64 and complex128 - is a complex types =)
	// ATTENTION! if you want to use floating point numbers just use the "float64" type
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.0+1.0) /* just for show */
	// "+" is addition, "-" substraction, "*" multiplication, "/" division, "%" remainder of the division
	fmt.Println(len("Hello World")) /* lenght of the string. If you want to use control sequences of characters as a "/n" - new lune, "/t" - tabulation */
	// a space is also considired a character, therefore the lenght of a string is 11.
	// String are indexed starting at 0, not at 0
	fmt.Println('H') /* так тоже можно, но только 1 букву.Странно */
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	// start to show boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	// finish to show boolean operators
	fmt.Println("32132 * 42452 = ", 32132*42452)
}
