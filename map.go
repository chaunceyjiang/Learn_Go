package main

import (
	"container/list"
	"fmt"
)

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
	_,ok:=map2[2]
	if ok{
		fmt.Println("yes")
		delete(map2,2)
	} else {
		fmt.Println("no")
	}
	l:=list.New()
	for i:=0;i<5;i++{
		l.PushBack(i)
	}
	for e:=l.Front();e!=nil;e=e.Next(){
		fmt.Println(e.Value,"list")
	}
	fmt.Println(len(map2))
	fmt.Println(mf[2](3))
}
