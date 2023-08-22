package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pinyinToHan := make(map[string]string)
	filePath := "../userlist" // 请替换为您的文件路径

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
			pinyin := parts[0]
			han := extractHanFromDescription(parts[1])
			pinyinToHan[pinyin] = han
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	for pinyin, han := range pinyinToHan {
		fmt.Printf("%s: %s\n", pinyin, han)
	}
}

func extractHanFromDescription(description string) string {
	var han string
	for _, c := range description {
		if c >= '\u4e00' && c <= '\u9fff' {
			han += string(c)
		}
	}
	return han
}
