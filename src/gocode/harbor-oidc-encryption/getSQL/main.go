package main

import (
	"encoding/csv"
	"fmt"
	"os"
	// "goproject/src/gocode/harborDB-setSQL/harbor"
)

func generateSQL(fileReader string, fileWriter string, dex_url string) {
	// 打开CSV文件
	filereader, err := os.Open(fileReader)
	if err != nil {
		panic(err)
	}
	defer filereader.Close()

	// 创建一个写入文件
	// filewriter, err := os.Create(fileWriter)
	// if err != nil {
	// 	panic(err)
	// }
	// defer filewriter.Close()
	// 创建一个bufio.Writer，将数据写入文件中
	// writer := bufio.NewWriter(filewriter)

	// 创建CSV Reader
	reader := csv.NewReader(filereader)

	// 读取CSV数据
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// fmt.Println(records)

	// 获取电话号码并连接到字符串中
	fmt.Println("循环csv文件内容，逐条数据加密")
	for _, record := range records {

		if record[5] == "phone" || record[5] == "" || record[6] == "" {
			fmt.Println("用户user_id:" + record[0] + "username:" + record[1] + "没有手机号或密码，跳过...")
			continue
		}

		// result := fmt.Sprintf("用户名：%v,手机号：%v, 密码：%v \n", record[1], record[5], record[6])
		result := fmt.Sprintf("harbor-tools user --update --username=%v --userid=%v --password=%v", record[1], record[5], record[6])
		result = "docker run --network=harbor_harbor --rm -it repos.cloud.cmft/kubesphere/harbor-tools sh -c '" + result + "'"
		fmt.Println(result)
	}

}

func main() {

	// 定义dex访问地址
	dex_url := "https://10.27.33.121:32000"

	// 定义数组类型的集群名称
	// harbor_name := [11]string{"host", "zzappzf", "csmgmt", "cstest", "zzapp", "csapp", "app1", "test1", "zzmgmtkf", "zztestkf", "mgmt1"}
	harbor_name := [1]string{"poc"}

	// 循环数组分别生成每个集群的sql
	var i int
	for i = 0; i <= len(harbor_name)-1; i++ {
		// str := fmt.Sprintf("%v ----> harbor集群的对应的k8s集群名称是：%v \n", i, harbor_name[i])
		// fmt.Println(str)

		writeFile := "output-" + harbor_name[i] + ".txt"
		readerFile := "../csvFile/harbor-" + harbor_name[i] + ".csv"

		// fmt.Println(writeFile)
		// fmt.Println(readerFile)

		// 调用方法传入文件到生成sql的函数
		fmt.Println("调用generateSQL函数加密并生成SQL，文件名为：", readerFile)
		generateSQL(readerFile, writeFile, dex_url)
	}

}
