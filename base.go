package main

import (
	"fmt"
)
func main() {
	var v int = 2
	a:=1
	t:=5+3i
	const (
		x=iota
		y=4
		z
		w=iota
	)
	s1:=[]string{"1","2","3"}
	var s2 []string=[]string{"one","two","three","four"}

	numbers:= make(map[string]int)
	numbers["one"]=1
	numbers["two"]=2

	fmt.Println(numbers["one"],)
	fmt.Println(v,a,t)
	fmt.Println(x,y,z,w)
	fmt.Println(s1,s2[1],len(s1),cap(s2),s2[1:3])


}
