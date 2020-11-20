package main

import "fmt"

func main(){
	fmt.Println("有朋自远方来，不亦乐乎")

	var a float32
	a = 5
	var d float32//该类型为浮点数,不同类型的值不能计算，必须改为同类型
	d = 2 * 3.14 * a
	fmt.Println("圆形周长为：",d)
	fmt.Printf("d的数据类型为：%T\n",d)

	//var age int = 18
	age := 18//省略var的形式
	fmt.Println("我的年龄是：",age)
	fmt.Printf("age的数据类型为：%T",age)

	//变量的多个变量
	var a1, b1, c1 int
	a1 = 3
	b1 = 4
	c1 = 5
	fmt.Println(a1,b1,c1)

	a1 = 6
	b1 = 7
	c1 = 8
	fmt.Println(a1,b1,c1)

	var st1,st2,st3 string = "赵","长","宇"
	fmt.Println(st1,st2,st3)

	var name, age1, address = "赵长宇",18,"江西南昌"
	fmt.Println(name,age1,address)
	fmt.Printf("name的内容为：%s,数据类型为：%T",name,name)//%s是格式化获取字符串变量的值
	fmt.Printf("age1的数值为：%d，数据类型为%T",age1,age1)//%d获取整数变量的值

	name1,name2 := 1,3
	fmt.Println(name1,name2)
	fmt.Println(A(2,3))
}
func A (i int ,j int)(int){
	return i*j
}