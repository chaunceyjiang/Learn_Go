package main

import "fmt"

type User struct {
	ID int
	Name string
}
func main()  {
	var a [5]int
	fmt.Println("emp:",a)
	a[4]=100
	fmt.Println("set",a)
	fmt.Println("get:",a[4])

	b:=[5]int{1,2,3,4,5}
	var c [4]int = [4]int{1,2,3}
	fmt.Println("dcl:",b,c)
	d:=[...]User{
		{0,"User0"},
		{8,"User8"},
	}
	fmt.Println(d,"长度",len(d))
	
}

