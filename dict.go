package main

import "fmt"

func main()  {
	m:=make(map[string]int)
	m["k1"]=7
	m["k2"]=13
	fmt.Println("map:",m,len(m))
	_,ok:=m["k3"]
	fmt.Println("ok:",ok)
	m2:=map[string]int{}
	fmt.Println("len",len(m2))
	m2["k1"]=1
	m2["k2"]=2
	fmt.Println("len",len(m2))
	m3:=map[int]string{}
	m3[1]="Strings"
	fmt.Println("len",len(m3),m3)
}

