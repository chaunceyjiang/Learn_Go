package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Print("First example with -1： ")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Printf("错误 : %s\n", err1)
	} else {
		fmt.Printf("%g %s\n", ret1, err1)
	}
	ret2, err2 := MySqrt(21.4)
	if err2 != nil {
		fmt.Printf("错误 : %s\n", err2)
	} else {
		fmt.Printf("%g %v\n", ret2, err2)
	}

	Greeting("yes", "1", "2", "3")

	arr := []int{7, 5, 6, 1, 4, 5}
	fmt.Println(Min(arr...))
	fmt.Println(arr[1:3])
	//Parm(arr[1:3]...)  error
	deferA()
	fmt.Println(deferB(), "deferB")
	fmt.Println(CallBack(3, Add))
	func(a int) { fmt.Println("yes", a) }(1)
}

func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("发生除0错误！")
	}
	return math.Sqrt(f), nil
}
func Greeting(prefix string, who ...string) {
	fmt.Println(prefix)
	fmt.Println(len(who), who)
}
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Parm(a, b int) {
	fmt.Println(a, b)
}
func deferA() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func deferB() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func Add(a, b int) int {
	return a + b
}

func CallBack(y int, f func(a, b int) int) int { //回调函数
	return y + f(1, 2)
}
