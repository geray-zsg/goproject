package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 准备请求参数
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "zhushougui")
	data.Set("password", "Zhu88jie!!")
	data.Set("client_id", "kubesphere")
	data.Set("client_secret", "kubesphere")

	// 发送POST请求
	resp, err := http.PostForm("http://10.27.33.9:31407/oauth/token", data)
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}

	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}

	// 输出响应结果
	fmt.Println("响应结果：", string(body))
}
