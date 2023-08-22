package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请提供拼音作为参数")
		return
	}

	pinyinToHan := make(map[string]string)
	filePath := "userlist" // 请替换为您的文件路径

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ",", 2)
		if len(parts) == 2 {
			pinyinToHan[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	targetPinyin := os.Args[1]
	if han, found := pinyinToHan[targetPinyin]; found {
		hanParts := strings.SplitN(han, "_", 2)
		nameHan := hanParts[len(hanParts)-1]
		fmt.Println(extractHan(nameHan))
	} else {
		fmt.Println("未找到匹配的汉字")
	}
}

func extractHan(s string) string {
	var han string
	for _, c := range s {
		if c >= '\u4e00' && c <= '\u9fff' {
			han += string(c)
		}
	}
	return han
}
