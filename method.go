package main

import (
	"fmt"
	"strconv"
	"sort"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height*r.width
}
func (r rect) perim() int{
	return 2*r.height+2*r.width
}

func main()  {
	r:=rect{10,5}
	fmt.Printf("Area:%d\n",r.area())
	fmt.Printf("Perim:%d\n",r.perim())
	x:=func (ii int) int {
		var i int =10
		return i*ii
	} (2)
	fmt.Println(x)
	y,_:=strconv.Atoi("123")
	yy,_:=strconv.ParseFloat("12.12",64)
	fmt.Println(y+23,yy)
	var arr2 [5]int
	for i,v:=range arr2{
		fmt.Println(i,v)
	}
	arr3 :=[...]int{1,2,3,4,5}
	arr4:=&arr3
	arr4[1]=6
	fmt.Println(arr3[1:4],*arr4)
	//sli := []int{1,2,3,5,4}
	sli1:=make([]int,0)
	t:=[]int{5,21,31,11,41,7,6,8,9}
	//copy(t,sli)
	sli1=append(sli1, 1,2,3,4,5)
	fmt.Println(t,sli1)
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	var ar = []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(ar[5:7],len(ar[5:7]),cap(ar[5:7]))
	fmt.Println(items)
	for item:=range items{
		items[item]*=2
	}
	fmt.Println(items)
	sort.Ints(t)
	index:=sort.SearchInts(t,41)
	fmt.Println(t,index,)
}
