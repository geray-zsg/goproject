/**
 * @Author: Geray
 * @Date: 2023/8/18 12:52:26
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/18 12:52:26
 * Description: 实现URL编解码
   容器平台提供给租户门户侧url地址为：租户门户地址+容器平台侧kubesphere-config中配置的cas地址的编码
   		租户门户地址：https://it.cmft/gateway/cas?service=
		ks配置文件配置地址：redirectURL: http://ks.cloud.cmft:30880/oauth/redirect/cas（编码后进行拼接：http%3A%2F%2Fks.cloud.cmft%3A30880%2Foauth%2Fredirect%2Fcas）

		结果：https://it.cmft/gateway/cas?service=http%3A%2F%2Fks.cloud.cmft%3A30880%2Foauth%2Fredirect%2Fcas

 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
*/
package main

import (
	"fmt"
	"net/url"
)

// 编码
func encodeURL(urlStr string) string {
	encodedURL := url.QueryEscape(urlStr)
	fmt.Println("编码结果", encodedURL)
	return encodedURL
}

// 解码
func getDecodeURL(urlStr string) {
	decodeURL, err := url.QueryUnescape(urlStr)
	if err != nil {
		fmt.Println("解码失败：", err)
		return
	}
	fmt.Println("解码结果：", decodeURL)
}

func main() {

	// 编码
	// encodeURLStr := "http://ks.cloud.cmft:30880/oauth/redirect/cas"
	// encodeURLStr := "http://10.127.128.10:30880/oauth/redirect/cas"
	encodeURLStr := "http://10.27.33.9:30880/oauth/redirect/cas"
	encodedURL := encodeURL(encodeURLStr)
	fmt.Println("编码后的拼接结果：https://it.cmft/gateway/cas?service=" + encodedURL)

	// URL
	decodeURLStr := "https://it.cmft/gateway/cas?service=http%3A%2F%2F10.127.128.10%3A30880%2Foauth%2Fredirect%2Fcas"
	// 解码
	getDecodeURL(decodeURLStr)

}
