#!/bin/bash

##########################
# 目录：
#   ./qy-zhunbei.sh     脚本
#   ./members           kubeconfig集群配置目录
#   ./webhooks          webhook备份目录
#   ./namespaces        ws和ns解绑备份目录
##########################

# 远程备份并删除 webhook
remote_backup_and_delete_webhook() {
    local KUBECONFIG_PATH=$1
    local CLUSTER_NAME=$2
    local RESOURCE_NAME="mutating-webhook-ks-cfg"

    echo "正在处理备份 webhook ..."
    if kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME > /dev/null 2>&1 ; then
        echo "资源存在，正在备份..."
        echo "命令是：kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME -o yaml > ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml"
        #kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME -o yaml > ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml

        echo "备份完成，正在删除资源..."
        echo "命令是：kubectl --kubeconfig=${KUBECONFIG_PATH} delete -f ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml"
        #kubectl --kubeconfig=${KUBECONFIG_PATH} delete -f ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml

        echo "资源已删除。"
    else
        echo "资源不存在。"
    fi

}

# 远程解绑 workspace 和 namespace 的关系
remote_unbind_workspace_namespace() {
    local KUBECONFIG_PATH=$1
    local CLUSTER_NAME=$2

    echo "正在处理 ${CLUSTER_NAME} 集群..."
    # 在这里使用 kubectl 连接到对应的 member 集群执行解绑的命令
    # 以循环方式对所有 namespace 执行解绑操作
}

# 定义集群名称数组
CLUSTER_NAMES=(
#    app1
#    csapp
#    csappzf
#    csmgmt   
#    cstest
#    host   
#    mgmt1
#    test1
#    zzapp
#    zzappzf
#    zzmgmtkf
#    zztestkf
    poc-test
)

# 执行操作
for CLUSTER_NAME in "${CLUSTER_NAMES[@]}"; do
    echo "----------------------------------------------> 正在操作的集群是：$CLUSTER_NAME"

    # 调用备份并删除 webhook 函数
    remote_backup_and_delete_webhook "./members/member_kubeconfig_${CLUSTER_NAME}.yaml" $CLUSTER_NAME

    # 调用解绑 workspace 和 namespace 的关系函数
    # remote_unbind_workspace_namespace "./members/member_kubeconfig_${CLUSTER_NAME}.yaml" $CLUSTER_NAME
done
