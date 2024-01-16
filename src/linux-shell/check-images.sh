#!/bin/bash

##########################
##   $2 指定域名        ##
##########################


# 目标镜像仓库url
harbor_url=$2

IMAGES=(
    # chaosmesh
    kubesphere/chaos-mesh:v2.6.1-rc.1
    kubesphere/chaos-daemon:v2.6.1
    kubesphere/chaos-dashboard:v2.6.1
    kubesphere/chaos-coredns:v0.2.6
    kubesphere/chaos-dlv:v2.6.1
    kubesphere/chaos-kernel:v2.6.1
    # fluidcloudnative
   fluidcloudnative/alluxioruntime-controller:v0.9.1-e79e93b
    fluidcloudnative/csi-node-driver-registrar:v2.3.0
    fluidcloudnative/fluid-csi:v0.9.1-e79e93b
    fluidcloudnative/dataset-controller:v0.9.1-e79e93b
    kubesphere/fluid-apiserver:v0.1.0
    kubesphere/fluid-apiserver:v0.1.0
    fluidcloudnative/fluid-webhook:v0.9.1-e79e93b
    fluidcloudnative/application-controller:v0.9.1-e79e93b
    fluidcloudnative/thinruntime-controller:v0.9.1-e79e93b 
    # gateway
    kubespheredev/ks-apiserver:v3.3.1-cmft-rc.1
    kubespheredev/ks-controller-manager:v3.3.1-cmft-rc.1
    kubespheredev/ks-console:v3.3.1-cmft-rc.1
    kubespheredev/ks-installer:v3.3.1-cmft
    # kubeeye
    kubesphere/kubeeye-apiserver:v1.0.1-rc.1
    kubesphere/kube-rbac-proxy:v0.11.0
    kubesphere/kubeeye-controller:v1.0.0
    kubesphere/kubeeye-job:v1.0.0
    # openkruise
    kubesphere/kruise-apiserver:v0.1.0
    kubesphere/kruise-apiserver:v0.1.0
    openkruise/kruise-manager:v1.4.0
    # sonobuoy
    sonobuoy/glusterdynamic-provisioner:v1.0
    sonobuoy/httpd:2.4.38-alpine
    sonobuoy/nonroot:1.0
    sonobuoy/nfs:1.0
    sonobuoy/sample-apiserver:1.17
    sonobuoy/echoserver:2.2
    sonobuoy/etcd:3.4.13-0
    sonobuoy/httpd:2.4.39-alpine
    sonobuoy/nginx:1.14-alpine
    sonobuoy/regression-issue-74839-amd64:1.0
    sonobuoy/debian-iptables:v12.1.2
    sonobuoy/etcd:3.4.13-0
    sonobuoy/nautilus:1.0
    sonobuoy/prometheus-to-sd:v0.5.0
    sonobuoy/busybox:1.29
    sonobuoy/kitten:1.0
    sonobuoy/redis:5.0.5-alpine
    sonobuoy/agnhost:2.20
    sonobuoy/cuda-vector-add:2.0
    sonobuoy/ipc-utils:1.0
    sonobuoy/nonewprivs:1.0
    sonobuoy/jessie-dnsutils:1.0
    sonobuoy/resource-consumer:1.5
    sonobuoy/iscsi:2.0
    sonobuoy/gluster:1.0
    sonobuoy/apparmor-loader:1.0
    sonobuoy/metadata-concealment:1.2
    sonobuoy/cuda-vector-add:1.0
    sonobuoy/nfs-provisioner:v2.2.2
    sonobuoy/nginx:1.15-alpine
    sonobuoy/perl:5.26
    sonobuoy/conformance:v1.19.9
    sonobuoy/systemd-logs:v0.4
    sonobuoy/conformance:v3.3.1-cmft-rc.1
    sonobuoy/sonobuoy:test-1115
)

# 拉取镜像
pull_images() {
    for image in ${IMAGES[@]}; do
        echo "----------> 拉取镜像: ${harbor_url}/${image}"
        echo "docker pull ${harbor_url}/${image}"

        # 获取上一个命令的退出状态码
        exit_code=$?

        # 检查退出状态码，0 表示成功，非0 表示失败
        if [ $exit_code -ne 0 ]; then
            echo "拉取镜像 ${harbor_url}/${image} 失败            <-----------------------"
        fi
    done
}


Help() {
    echo "提供参数
        pull                镜像拉取

        例如：
        ./check-images.sh pull [url]
    "
}


case $1 in
    "pull")
        pull_images
        ;;
    *)
        Help
        ;;
esac