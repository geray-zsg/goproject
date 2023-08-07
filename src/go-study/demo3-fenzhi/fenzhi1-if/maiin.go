/**
* @Author: Geray
* @Date: 2023/8/7 20:22:26
* @LastEditors: Geray
* @LastEditTime: 2023/8/7 20:22:26
* Description: 单分支和双分支
* Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import (
	"fmt"
)

func main() {
	// 实现功能：如果口罩的库存小于30个，提示库存不足
	var count int = 20
	if count < 30 { //{}一定不能省略
		fmt.Println("对不起，库存不足")
	}
	// 在go语言中，if后面可以并列加入变量的定义
	if count1 := 20; count1 < 30 {
		fmt.Println("对不起，库存不足")
	}

	// 上面都是单分支，双分支if{代码1}else{代码2}
	// 实现功能：如果口罩的库存小于30个，提示库存不足，否则展示存储充足
	var count3 int = 20
	if count3 < 30 {
		fmt.Println("对不起，库存不足")
	} else {
		fmt.Println("库存充足")
	}

	// 多分支，if{代码1}else if{代码1} else if{代码1}else{代码1}
	// 实现功能：根据学生分数，判定学生成绩等级
	var score int = 99
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不合格")
	}
}
