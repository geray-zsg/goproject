package main

import (
	"fmt"
)

// 定义函数
/*
func 函数名(形参列表) (返回值类型列表) {  // 如果返回值类型只有一个的话，（）可以省略
	执行语句..
	return + 返回值列表
}
*/

// 自定义函数：功能：两个数相加 （单个返回值类型）
func cal(num1 int, num2 int) int {
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

// 定义函数求两个数的和，并求出差
func cal1(num1 int, num2 int) (int, int) {

	var sum int = 0
	sum += num1
	sum += num2

	var cha int = 0
	cha = num1 - num2

	return sum, cha

}

func main() {
	// var num1 int = 10
	// var num2 int = 20
	// var sum int = 0
	// // sum += num1
	// // sum += num2
	// sum = num1 + num2
	// fmt.Println(sum)

	// 调用函数，定义sum变量接收
	sum1 := cal(10, 20)
	fmt.Println(sum1)

	// 调用函数，定义sum2接收
	sum2 := cal(80, 20)
	fmt.Println(sum2)

	// 调用test.go中的函数
	// sum3 := test.sum(1, 3)
	// fmt.Println(sum3)

	// 接收和、差
	sum3, cha := cal1(10, 40)
	fmt.Println(sum3, cha)

	// 如果只想接收一个，可以使用_忽略
	sum4, _ := cal1(30, 10)
	fmt.Println(sum4)

}
