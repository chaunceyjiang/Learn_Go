package main

import "fmt"

type identifier struct {
	Name string
	Age float32
}
type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}
type Dog struct {
	types string
}
type Node struct {
	value int
	L *Node
	R *Node
}
func (self *Dog)PrintName()  {
	fmt.Println(self.types)
}
func (self *identifier)PrintName()  {
	fmt.Println(self.Name)
}
func pNode(root *Node)  {
	if root==nil{
		return
	}
	fmt.Printf("value: %d adderss:%p\n",root.value,root)
	pNode(root.L)
	pNode(root.R)
}
type Integer struct {n int}

func (self *Integer)Get() int {
	return self.n
}
func (self *Integer)Set(a int)  {
	self.n=a
}
func main() {
	var t3 identifier
	var t *identifier
	type A struct {a int}
	type B struct {a, b int}

	type C struct {A; B}
	var c C;
	c.A.a=1
	fmt.Println(c.A.a)
	t2:=new(Dog)
	(*t2).types="2"
	t3.Age=2
	t3.Name="3"
	t=&t3
	t3.PrintName()
	t.Name="333"
	t3.PrintName()
	//t2.types="斗牛犬"
	t2.PrintName()
	fmt.Printf("t2: %p %p",&t3,t)

	root:=new(Node)
	lists:=[]int{2,6,3,1,4,8,5}
	root.value=lists[0]
	var p *Node
	for _,i:=range lists[1:]{
		p=root
		for {
			if i > p.value {
				if p.R == nil {
					p.R = new(Node)
					p.R.value=i
					break
				}
				p = p.R

			} else {
				if p.L == nil {
					p.L = new(Node)
					p.L.value=i
					break
				}
				p = p.L
			}
		}
	}
	fmt.Println(root)
	pNode(root)
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
	n:=new(Integer)
	n.n=1
	n.Set(2)
	fmt.Println(n.Get())
}