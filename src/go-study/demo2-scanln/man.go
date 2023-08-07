/**
 * @Author: Geray
 * @Date: 2023/8/7 20:03:52
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/7 20:03:52
 * Description:
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import "fmt"

func main() {
	// 实现功能：键盘录入学生的姓名、年龄、成绩、是否为VIP
	// 方式1：scanln
	var name string
	fmt.Println("请录入学生的姓名：")
	fmt.Scanln(&name)

	var age int
	fmt.Println("请录入学生的年龄：")
	fmt.Scanln(&age)

	var score float32
	fmt.Println("请录入学生的成绩：")
	fmt.Scanln(&score)

	var isVIP bool
	fmt.Println("请录入学生是否为VIP：")
	fmt.Scanln(&isVIP)

	fmt.Printf("学生的姓名：%v,年龄：%v,成绩：%v,是否为VIP：%v", name, age, score, isVIP)

	// 方式2：Scanf
	fmt.Println("请依次录入学生的姓名、年龄、成绩、是否为VIP，使用空格分割！")
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVIP)
	fmt.Printf("学生的姓名：%v,年龄：%v,成绩：%v,是否为VIP：%v", name, age, score, isVIP)

}
