package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 获取 kubeconfig 文件路径
	/*
		这三行代码的作用是用于获取 Kubernetes 集群配置文件（kubeconfig 文件）的路径，并将其作为一个可选参数传递给你的程序。让我来详细解释一下：
			home := homeDir()：homeDir() 函数用于获取用户的主目录路径。这个函数会根据操作系统返回当前用户的主目录路径（例如，在 Linux 系统中是 /home/username，在 Windows 系统中是 C:\Users\username）。这样做是为了在程序中找到默认的 kubeconfig 文件路径。
			kubeconfig := flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")：这一行创建了一个名为 kubeconfig 的命令行标志。如果用户在运行程序时提供了 --kubeconfig 参数，则程序将使用用户提供的路径作为 kubeconfig 文件的路径。否则，它将使用默认路径，即用户的主目录下 .kube/config 文件。
			flag.Parse()：这个函数解析命令行参数，将用户提供的参数传递给程序。在这里，它用于解析和处理 --kubeconfig 参数。
		也就是可以使用--kubeconfig参数来指定k8s的文件
	*/
	home := homeDir()
	kubeconfig := flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// 生成 client-go 的 Config 对象
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %s\n", err.Error())
		os.Exit(1)
	}

	// 创建 Kubernetes 的客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating client: %s\n", err.Error())
		os.Exit(1)
	}

	// 获取 Pod 的状态
	podName := "your-pod-name" // 替换为你要获取状态的 Pod 的名称
	namespace := "default"     // 替换为 Pod 所在的命名空间
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting pod: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Pod %s status: %s\n", podName, pod.Status.Phase)

	// 删除 Pod
	deletePolicy := metav1.DeletePropagationForeground // 设置删除策略
	err = clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		fmt.Printf("Error deleting pod: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Pod %s deleted\n", podName)

	// 等待一段时间，然后再次获取 Pod 的状态，确认删除后的状态
	time.Sleep(10 * time.Second) // 等待 10 秒钟
	deletedPod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting deleted pod: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Pod %s status after deletion: %s\n", podName, deletedPod.Status.Phase)
}

// 获取 home 目录
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
