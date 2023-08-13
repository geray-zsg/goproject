/**
 * @Author: Geray
 * @Date: 2023/8/13 11:17:24
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/13 11:17:24
 * Description: 结构体，绑定方法
	方法使用的是值传递（拷贝副本修改）
	要实现修改原来的值，需要使用指针(*)，传递地址(&)，修改原来的数据
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
*/
package main

import "fmt"

// 定义一个Person结构体
type Person struct {
	Name string
	Age  int
}

// 结构体Person绑定方法
func (p Person) testName() {
	// testName未使用指针所以修改的是副本不影响原来的数据
	// 修改副本name
	p.Name = "露露"
	fmt.Println(p.Name)
}

func (p *Person) testAge() { // 这里定义了指针其他地方都可以简写（底层编译器会自动补全 * 或者 &）
	// 这里传递结构体的指针，调用的时候需要传递地址，修改原来的值，原来的值将被修改
	// (*p).Age = 20      // 这里也可以简写
	p.Age = 23
	fmt.Println(p.Age) // 这里是简写：(*p).Age
}

func main() {
	// 创建结构体对象
	var p Person

	// 由于testName未使用指针，方法使用值传递，所以不影响原数据
	p.Name = "莉莉"
	p.testName()
	fmt.Println(p.Name)

	// 由于testAge使用指针传递地址修改原数据，所以原数据被修改（两个输出值相同）
	p.Age = 28
	// (&p).testAge()	// 这里也可以简写
	p.testAge()
	fmt.Printf("p的地址: %p, age值为： %v", &p, p.Age)

}
