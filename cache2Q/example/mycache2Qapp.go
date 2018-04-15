package main

import (
	"cache2Q"
	"fmt"
)
func main()  {
	c:=cache2Q.NewCache("第二次缓存",10,20)
	c.Add("1","dsdssss")
	c.SetAddedItemFunc(func(item *cache2Q.CacheItem) {
		fmt.Println("添加成功")
	})
	l:=c.GetFIFO()
	l.SetFifoTransferToLRU(func(item *cache2Q.CacheItem) {
		fmt.Println("开始传输")
	})
	l.SetAboutToDeleteItemFunc(func(item *cache2Q.CacheItem) {
		fmt.Println("开始删除。。。。。")
	})
	fmt.Println("-------------------")
	item,err:=c.Value("1")
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(item.Key(),item.Value())
	}
	item,err=c.Value("1")
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(item.Key(),item.Value())
	}
	item,err=c.Value("1")
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(item.Key(),item.Value())
	}
	item,err=c.Value("1")
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(item.Key(),item.Value())
	}
	for i:=0;i<24;i++{
		c.Add(i,i)
	}
}
