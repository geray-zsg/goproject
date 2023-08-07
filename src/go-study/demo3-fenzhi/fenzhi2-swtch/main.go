/**
 * @Author: Geray
 * @Date: 2023/8/7 20:37:26
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/7 20:37:26
 * Description: switch分支
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import "fmt"

func main() {
	// 实现功能：根据学生分数，判定学生成绩等级
	var score int = 89

	switch score / 10 {
	case 10:
		fmt.Println("满分")
	case 9:
		fmt.Println("优秀")
	case 8:
		fmt.Println("良好")
	case 7, 6: // Go中的switch可以多个判断，作用为or
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}
