package dbutils

import "fmt"

// 获取数据库连接，如果想要让这个函数被其他包使用，首字母必须大写
func GetConn() {
	fmt.Println("执行了dbutils包下的GetConn函数")
}
