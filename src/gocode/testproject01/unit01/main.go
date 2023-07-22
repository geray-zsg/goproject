package main

import "fmt"

// 定一个一个函数，交换两个数
func test(num1 int, num2 int) (int, int) {
	num1 += num2
	num2 = num1 - num2
	num1 = num1 - num2
	return num1, num2
}

// args...int 可以传入任意多个数量的int类型的数据，处理可变参数的时候，将可变参数当做切片来处理（可以理解为数组）
func test1(nums ...int) int {
	var sum int = 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		sum += nums[i]
	}
	return sum
}

func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数：num1 = %v， num2 = %v \n", num1, num2)
	num3, num4 := test(num1, num2)
	fmt.Printf("交换前的两个数：num1 = %v， num2 = %v \n", num3, num4)

	sum := test1(10, 20, 30)
	fmt.Println(sum)
}
