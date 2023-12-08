package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// prometheusURL := "http://prometheus.example.com/api/v1/query"
	prometheusURL := "http://10.27.0.30:30093//api/v1/query"

	// 构建 PromQL 查询
	queryCPU := "100 - avg by(instance) (irate(node_cpu_seconds_total{mode=\"idle\"}[5m])) * 100"
	queryMemory := "100 - (avg by (instance) (node_memory_MemFree_bytes + node_memory_Cached_bytes) / avg by (instance) (node_memory_MemTotal_bytes) * 100)"
	queryDisk := "100 - (avg by (instance) (node_filesystem_avail_bytes{mountpoint=\"/\"}) / avg by (instance) (node_filesystem_size_bytes{mountpoint=\"/\"}) * 100)"

	// 发送查询请求
	queryAndPrintMetric(prometheusURL, "CPU", queryCPU)
	queryAndPrintMetric(prometheusURL, "Memory", queryMemory)
	queryAndPrintMetric(prometheusURL, "Disk", queryDisk)
}

// 发送查询请求并打印结果
func queryAndPrintMetric(url, metricName, query string) {
	// 构建查询参数
	params := fmt.Sprintf("query=%s", query)
	resp, err := http.Get(fmt.Sprintf("%s?%s", url, params))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 打印结果
	fmt.Println("Metrics for", metricName, ":", result)
}
