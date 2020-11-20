package main

import "fmt"

type Student struct {
	name string
	age int
	sex string
	class string
}
func main(){
	fmt.Println("众志成城，抗击疫情。致敬最美逆行者白衣天使们")
	fmt.Print("明天就是周末了，可以休息了\n")
	//ln表示换行，但也可以用转义字符\n换行
	fmt.Print("哈哈哈哈\t")
	//\t表示空格
	fmt.Print("我知道了")
	var a int = 3
	//定义为var，名字是a，最后用int表示类型,值为3
	var b int = 4
	var c int
	c = (a * b) / 2
	//等号表示赋予一个值
	fmt.Println("三角形面积是：",c)//字符串和定义值不能在同一个打印里,需要逗号隔开
	//fmt.Print(c)
	fmt.Printf("c的数据类型是：%T\n",c)
	//数据类型转换
	var num int8
	num = 100
	var num1 int16
	num1 = 200
	sum := int16(num) + num1//不同的类型不能相加减，并且像字符串类型数据改为数值类型会乱码
	fmt.Println(sum)
	//关于floot的小数点位数
	var pin = 3.1415926
	fmt.Printf("保留n位小数：%.nf\n",pin)
	fmt.Printf("保留两位小数：%.2f\n",pin)
	//查看字符串的长度
	var name1 = "asljasljd"
	name2 := len(name1)
	fmt.Println("字符串长度为：",len(name1))
	fmt.Println(name2)
	for i :=1; i <=5; i++ {
		//空格规律等于5-行
		//5-1=4个空格
		//5-2=3个空格
		for  num := 1; num <= 5-i; num++ {
			fmt.Print(" ")
		}
		for str := 1; str <= 2 * i -1; str++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//中间分界
	for i := 1; i <= 4; i++ {
		for  num := 1; num <= i; num++ {
			fmt.Print(" ")
		}
		for str := 1; str <= 9 - 2 *i; str++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//直角三角形
	for i :=1; i <=5; i++ {//i表示几行
		for str :=1;  str <= 2*i-1; str++ {
			fmt.Print("*")
		}
		fmt.Println()//每次走完一行回车进行下一行
	}
	//九九乘法表
	for i :=1; i <= 9; i++{//i表示几行
		for j :=1; j <= i; j++ {//一行里的列
			fmt.Printf("%d * %d = %d",j,i,i*j)
			fmt.Print(" | ")
		}
		fmt.Println()
	}
	//2-50的素数
	fmt.Print("2-50的素数为：")
	for a :=2; a <=50; a++{
		for b :=2; b <=(a/b); b++ {
			if a % b == 0 {
				break
			}
		}
		if b> (a/b) {
			fmt.Print(a)
		}
		fmt.Print(" ")
	}
	fmt.Println()
	//数组
	var agear = [3]int{10,12,13}
	fmt.Println(agear)
	agear[0] = 11
	fmt.Println(agear)
	//Student{"赵长宇",18,"男","c190603"}
	fmt.Println(Student{"赵长宇",18,"男","c190603"})
	fmt.Println()
}
