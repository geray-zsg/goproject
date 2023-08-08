/**
 * @Author: Geray
 * @Date: 2023/8/7 20:44:04
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/7 20:44:04
 * Description: for循环
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import "fmt"

func main() {
	// Go的for循环有三种形式，1、使用分好，和其他语言中的一样
	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			// fmt.Println("我要跳出本次循环了，所以你看不到i=", i)
			fmt.Println("我要跳出本次循环了，所以你看不到偶数")
			continue // 跳出本次循环，继续下次循环，continue用在类似的while循环要注意，容易死循环
		}
		fmt.Println("i=", i)
	}

	// 2、和while一样
	var j int = 0
	for j < 10 {
		if j > 5 {
			fmt.Println("我跳出循环了，bye！")
			break // break跳出循环
		}
		fmt.Println("j=", j)
		j++
	}
	fmt.Println("我已跳出for之外，不在循环之中")

	// 3、死循环
	// for {
	// }

}
