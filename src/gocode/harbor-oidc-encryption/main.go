package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"goproject/src/gocode/harbor-oidc-encryption/harbor"
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
	filewriter, err := os.Create(fileWriter)
	if err != nil {
		panic(err)
	}
	defer filewriter.Close()
	// 创建一个bufio.Writer，将数据写入文件中
	writer := bufio.NewWriter(filewriter)

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
		// 如果手机号码为空则退出本次循环
		if record[2] == "" || record[2] == "phone" {
			fmt.Println("user_id: " + record[0] + "  没有手机号信息，跳过...")
			continue
		}
		// 对手机号码进行拼接加密，由于文件中手机号在第三列所以这里使用索引为2
		encryption_phone, _ := harbor.Marshal(&harbor.IDTokenSubject{ConnId: "cas", UserId: record[2]})
		/*
			这里如果要测试自己的账号单条数据不用csv文件则注释上一条和上面的判断语句
			因为要循环文件不行啊，后面在调整代码吧，先使用csv文件：顺序以此如下
			user_id	username phone email
		*/
		// encryption_phone, _ := harbor.Marshal(&harbor.IDTokenSubject{ConnId: "cas", UserId: "17339872165"})
		fmt.Println(record[2])
		fmt.Println(encryption_phone)

		// 加密后拼接url
		// encryption_phone_url := encryption_phone + "https://10.27.33.121:32000"
		encryption_phone_url := encryption_phone + dex_url
		fmt.Println(encryption_phone_url)

		// 根据user_id生成对应的SQL，插入到oidc_user表中的数据
		// INSERT INTO oidc_user (user_id,secret,subiss) VALUES ('3','','CgsxNzMzOTg3MjE2NRIDY2Fzhttps://10.27.33.121:32000');
		sql := fmt.Sprintf("INSERT INTO oidc_user (user_id,secret,subiss) VALUES ('%s','','%s');", record[0], encryption_phone_url)
		fmt.Println(sql)

		// 将数据写入文件中
		_, err := writer.WriteString(sql + "\n")
		if err != nil {
			panic(err)
		}
	}
	// 循环结束后刷新bufio.NewWriter，确保所有的数据都被写入到文件中
	err = writer.Flush()
	if err != nil {
		// 处理刷新bufio.NewWriter失败的情况
		panic(err)
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
		readerFile := "csvFile/harbor-" + harbor_name[i] + ".csv"

		// fmt.Println(writeFile)
		// fmt.Println(readerFile)

		// 调用方法传入文件到生成sql的函数
		fmt.Println("调用generateSQL函数加密并生成SQL，文件名为：", readerFile)
		generateSQL(readerFile, writeFile, dex_url)
	}

}
