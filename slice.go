package main

import "fmt"
func main() {
	var array [5]int=[5]int{1,2,3,4,5}
	array2:=[5]int{0,2,3,3,1}
	fmt.Println(array,array2)
	fmt.Printf("%d\n",array2[2])

	for i:=0;i<len(array);i++{
		fmt.Printf("索引：%d，值：%d \n",i,array[i])
	}
	fmt.Println()
	for i,v:=range array2 {
		fmt.Printf("索引：%d，值：%d \n",i,v)
	}
}
