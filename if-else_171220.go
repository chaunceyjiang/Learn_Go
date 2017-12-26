package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var s string = "2"
	var a int
	fmt.Printf("请输入一个整数： ")
	fmt.Scanf("%d", &a)
	b, _ := strconv.Atoi(s)
	if a > b {
		fmt.Printf("Greater than\n")
	} else if a == b {
		fmt.Printf("Equal\n")
	} else {
		fmt.Printf("Less than\n")
	}
	s = func(bool1 bool) string {
		if (1 == 1) && bool1 {
			return "yes"
		} else {
			return "no"
		}
	}(func() bool {
		if rand.Intn(2) == 0 {
			return false
		} else {
			return true
		}
	}())
	fmt.Printf("%s\n", s)
}
