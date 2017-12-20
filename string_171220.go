package main

import (
	"fmt"
	//	"math/rand"
	//	"time"
	"strconv"
	"strings"
)

/*
func main() {
	fmt.Printf("hello2\n")
	var x float32 =2.3
	fmt.Printf("%f\n",x)
	fmt.Printf("%t",true)
}
*/
/*
func main()  {
	for i:=0;i<10;i++{
		a:=rand.Int()
		fmt.Printf("%d / ",a)
	}
	for i:=0;i<5;i++{
		r:=rand.Intn(3)
		fmt.Printf("%d / ",r)
	}
	fmt.Println()
	times:=int64(time.Now().Nanosecond())
	rand.Seed(times)
	for i:=0;i<10;i++{
		fmt.Printf("%.5f\n",100*rand.Float32())
	}
}
*/
/*
func main()  {
	fmt.Println(`这里有一个换行
        --------
换行结束~~`)
}
*/

func main() {
	s := `This Is An Example Of A String`
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n", s, "Th")
	fmt.Printf("%t\n", strings.HasSuffix(s, "ing"))
	fmt.Printf("%t\n", strings.HasPrefix(s, "TH"))
	fmt.Printf("%t\n", strings.Contains(s, "example"))
	var str string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is :%d\n", strings.Index(str, "Marc"))
	fmt.Printf("The position of the last instance of \"Hi\" is:%d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("toLower: %s\n", strings.ToLower(s))
	fmt.Printf("Itoa:%s\n", strconv.Itoa(13+22))
	fmt.Printf("Ftoa:%s\n", strconv.FormatFloat(12.123456789, 'f', 5, 64))
}
