package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
	}
	//for i,j:=0,5;i<j;i++,j++{fmt.Printf("yes\n")}   error , j unresolved reference
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		fmt.Printf("yes\n")
	}
	for i := 1; i < 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%2d * %2d = %2d ", j, i, i*j)
		}
		fmt.Println()
	}
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is :%d\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is :%c \n", i, str[i])
	}
	str2 := "Chinese: 日本語"
	fmt.Printf("\nThe llength os str2 is :%d\n", len(str2))
	for i := 0; i < len(str2); i++ {
		fmt.Printf("Character on position %d is :%c\n", i, str2[i])
	}
	gotoi := 1
	goto LABLE1
HEAR:
	fmt.Printf("goto %d\n", gotoi)
	gotoi++
	if gotoi > 10 {
		goto LABLE1
	}
	goto HEAR
LABLE1:
	str3 := "G"
	for i := 0; i < 5; i++ {
		fmt.Println(str3)
		str3 += "G"
	}

	for i := 1; i < 10; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
	for pos, char := range str2 {
		fmt.Printf("character %c at byte position %d\n", char, pos)
	}
	for index, rune2 := range str2 {
		fmt.Printf("%-2d    %d    %#U    '%c'    % X\n", index, rune2, rune2, rune2, []byte(string(rune2)))
	}
	str4 := '中'
	fmt.Println(str4)
	fmt.Printf("%d %c\n", str4, str4)
	str5 := '日'
	fmt.Println(str5)
	str6 := "中"
	fmt.Println(str6)
}
