package main // 1.package进行包的声明，建议：包的声明和这个包所在的文件夹同名
import (
	"fmt"
	test "goproject/src/gocode/testproject02/crm/main/dbutils"
)

// 1.main包是程序的入口包，一般main函数会放在这个包下

func main() {
	fmt.Println("这个是main函数的执行，程序入口...")
	test.GetConn()
}
