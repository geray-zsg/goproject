package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	etcdBackupStatus = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "node_loadapp_etcd_backup_status",
			Help: "Etcd backup status (1: Success, 0: Failure)",
		},
	)
	mutex = &sync.Mutex{}
)

func main() {
	prometheus.MustRegister(etcdBackupStatus)

	go func() {
		for {
			updateMetrics()
			time.Sleep(30 * time.Second) // Update metrics every 30 seconds
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Etcd Backup Metrics Exporter</title></head>
			<body>
			<h1>Etcd Backup Metrics Exporter</h1>
			<p><a href="/metrics">Metrics</a></p>
			</body>
			</html>`))
	})

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func updateMetrics() {
	mutex.Lock()
	defer mutex.Unlock()

	// Read the contents of the file specified by the environment variable
	filePath := os.Getenv("ETCD_BACKUP_FILE_PATH")
	if filePath == "" {
		fmt.Println("ETCD_BACKUP_FILE_PATH environment variable not set")
		return
	}

	// Read the contents of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Extract the metric value from the file content
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "node_loadapp_etcd_backup_status") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				val, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					fmt.Println("Error parsing metric value:", err)
					return
				}
				etcdBackupStatus.Set(val)
				return
			}
		}
	}
	fmt.Println("Metric not found in file")
	etcdBackupStatus.Set(0) // Set default value if metric not found or error occurs
}
