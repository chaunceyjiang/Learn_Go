package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Now:%s\n", time.Now())                                                                      //Now:2017-12-20 18:19:48.9307575 +0800 CST m=+0.003006001
	fmt.Printf("Day:%02d Month:%02d Year:%4d \n ", time.Now().Day(), time.Now().Month(), time.Now().Year()) //Day:20 Month:12 Year:2017
	t := time.Now().UTC()
	fmt.Println(t)                    // 2017-12-20 10:19:48.9407857 +0000 UTC
	fmt.Println(t.Format(time.ANSIC)) //Wed Dec 20 10:19:48 2017
}
