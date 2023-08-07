/**
 * @Author: Geray
 * @Date: 2023/8/7 19:27:13
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/7 19:27:13
 * Description:	学习指针
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import "fmt"

func main() {

	var age int = 18

	// &符号+变量：就可以获取这个变量内存地址
	fmt.Println(&age) //0xc000014088

	// 定义一个指针变量（指针：就可以理解为一个内存地址）
	// ptr指针变量的名字
	// ptr对应的就是类型就是：*int 是一个指针类型（可以理解为 指向 int类型的指针）
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)

	// 想通过ptr指针或地址获取那个数据
	fmt.Println("ptr指向的数值为：", *ptr)

	// 1.可以通过指针改变指向的值
	*ptr = 20
	fmt.Println("修改后的age值：", age)

	// 2.指针接收的一定是地址值
	// 3.指针的地址不可以不匹配（数据类型要匹配吗，int类型的数据不能使用其他数据类型来接收指针）
	// 4.基本数据类型（又叫值类型），都有对应的指针类型，形式为*数据类型，比如int对应的指针就是*int，float32对应的聚水*float32
}
