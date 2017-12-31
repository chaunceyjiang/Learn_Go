package main

import "fmt"

func main() {
	//声明map
	var map1 map[string]int
	map2:=make(map[int]string)
	//初始化map
	map1=map[string]int{"1":1}
	map2=map[int]string{2:"2"}
	fmt.Println(map1,map2)
	//声明并初始化
	mf:=map[int]func(int) int{1: func(a int) int {
		return 1+a
	},
	2: func(a int) int {
		return 2+a
	},
	3: func(a int) int {
		return 3+a
	}}

	fmt.Println(mf[2](3))
}
