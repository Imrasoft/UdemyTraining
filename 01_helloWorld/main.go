package main

import "fmt"

func main() {
	fmt.Println("42")
	fmt.Printf("\n Decimal is %d \n Binary is %b \n Hexadecimal is %x", 42,42,42)

	for i:=60; i<120; i++ {
		fmt.Printf("\n %d \t %b \t %x \t %q \t %s \n", i,i,i,i,i)
	}
}