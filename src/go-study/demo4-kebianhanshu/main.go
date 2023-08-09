/**
 * @Author: Geray
 * @Date: 2023/8/8 15:36:11
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/8 15:36:11
 * Description:
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import "fmt"

func change(s ...string) {

	s[1] = "Go"
	s = append(s, "Playground")
	fmt.Println(s)
}
func main() {
	welcome := []string{"Hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
