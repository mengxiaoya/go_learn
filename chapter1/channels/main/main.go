package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (r *Person) hello() string {
	return r.name
}

// 笔记：Go语言中*和&的使用
// 1. * 可以表示一个变量是指针类型，例如 hello方法的接受者就是一个Person指针类型
// 2. & 直接修饰在结构类型Person上并没有什么效果，看打印结果即可
//    但是 & 修饰在变量r上的意思就是取r变量的地址值，例如第22行代码，打印的是0xc000006030
func main() {
	fmt.Println(&Person{"liu", 10})
	var r = &Person{"liu", 10}
	fmt.Println(r)
	fmt.Println(*r)
	fmt.Println(&r)
	var r2 *Person = &Person{"liu2", 10} // 声明一个指向Person结构体类型的指针变量r2
	var r3 Person = Person{"liu3", 20}   // 声明一个普通的Person结构体类型的变量r3
	fmt.Println(r2.hello())
	fmt.Println(r2)
	fmt.Println(r3)

	// 结果依次如下
	//&{liu 10}
	//&{liu 10}
	//{liu 10}
	//0xc000006030
	//liu2
	//&{liu2 10}
	//{liu3 20}
}
