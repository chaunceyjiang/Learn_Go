package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	var a = 7
	switch a {
	case 10:
		fmt.Printf("10")
		fmt.Printf("20")
		fallthrough
	case 7:
		fmt.Printf("7")
		fmt.Printf("14")
	default:
		fmt.Printf("default\n")
	}

	switch {
	case a >= 10:
		fmt.Printf("\n-10-")
		fmt.Printf("\n-20-")
	case a == 7:
		fmt.Printf("7\n")
	default:
		fmt.Printf("default")
	}
}
