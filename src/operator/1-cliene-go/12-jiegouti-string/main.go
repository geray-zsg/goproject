/**
 * @Author: Geray
 * @Date: 2023/8/13 11:49:38
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/13 11:49:38
 * Description:
		如果一个类型实现了String()方法，那么fmt.Println默认会调用这个变量的String()进行输出
		注：以后定义结构体，常定义String()作为数据结构体信息的方法，在fmt.Println()会自动调用
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
*/

package main

import "fmt"

// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 绑定方法
func (p *Person) String() string {
	// Sprintf()函数将返回strin类型值
	str := fmt.Sprintf("Name: %v, Age: %v", p.Name, p.Age)
	return str

}

func main() {
	// 创建结构体对象
	p := Person{
		Name: "莉莉",
		Age:  26,
	}

	// 传入地址，自动调用该结构体的String()方法返回数据
	fmt.Println(&p) // 这里不能简写，如果简写将返回创建的结构体对象：Name: 莉莉, Age: 26
	fmt.Println(p)  // 如果简写将返回创建的结构体对象：{莉莉 26}

}
