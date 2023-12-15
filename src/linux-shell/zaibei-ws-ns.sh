#!/bin/bash

set -e
##########################
# 目录：
#   ./qy-zhunbei.sh  [参数]   脚本
#   ./members           kubeconfig集群配置目录
#   ./webhooks          webhook备份目录
#   ./namespaces        ws和ns解绑备份目录
##########################


###########################开始：member集群##############################
# 获取webhook、ws和ns的关系
Get_webhook_ws_ns() {
    echo "----------> webhook"
    kubectl  --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration mutating-webhook-ks-cfg
    echo "----------> ws和ns关系"
    kubectl  --kubeconfig=${KUBECONFIG_PATH} get ns --show-labels | grep workspace 
}

# 远程备份并删除 webhook
Remote_backup_and_delete_webhook() {
    local RESOURCE_NAME="mutating-webhook-ks-cfg"

    echo "正在操作：mutating-webhook-ks-cfg"
    kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME
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

# 恢复webhook
Restore_webhook() {

    if [ -f "./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml" ]; then
        kubectl --kubeconfig=${KUBECONFIG_PATH}  create -f ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml
    else
        echo "${CLUSTER_NAME} 集群不存在"
    fi
}

# 远程解绑 workspace 和 namespace 的关系
Remote_unbind_workspace_namespace() {

    # 备份ws和ns
    kubectl --kubeconfig=${KUBECONFIG_PATH} get ns --show-labels > ./namespaces/ws_ns_${CLUSTER_NAME}

    # 在这里使用 kubectl 连接到对应的 member 集群执行解绑的命令
    # 以循环方式对所有 namespace 执行解绑操作
    for i in `kubectl  --kubeconfig=${KUBECONFIG_PATH}  get ns | grep -v NAME | awk '{print $1}'`;do kubectl  --kubeconfig=${KUBECONFIG_PATH}  label ns $i kubesphere.io/workspace- && kubectl  --kubeconfig=${KUBECONFIG_PATH}  patch ns $i -p '{"metadata":{"ownerReferences":[]}}' --type=merge;done
}


# 恢复ns和ws关系
Restore_ws_ns() {

    if [ -f "./namespaces/ws_ns_${CLUSTER_NAME}" ]; then
        awk '/workspace/{split($NF,w,",");for(i=1;i<=length(w);i++){if(w[i]~"kubesphere.io/workspace"){cmd="kubectl --kubeconfig=${KUBECONFIG_PATH} label ns "$1" "w[i]" --overwrite";system(cmd)}}}' ./namespaces/ws_ns_${CLUSTER_NAME}
    else
        echo "${CLUSTER_NAME} 集群不存在"
    fi
}
###########################结束：member集群##############################

###########################开始：备host集群清理待回收资源##############################
resources=("
    Group"  
    "GlobalRole"  
    "GlobalRoleBinding"  
    "workspaces.tenant.kubesphere.io"  
    "workspacetemplate"  
    "federatedworkspaces.types.kubefed.io"  
    "WorkspaceRole"  
    "WorkspaceRoleBinding"  
    "federatedusers.types.kubefed.io"  
    "users.iam.kubesphere.io"
)

# 检查host集群是否存在待清理的资源
CheckHostDeleting(){
    for resource in "${resources[@]}"  
    do  
        echo "######################| $resource"
        kubectl get $resource -o yaml | grep deletionGracePeriodSeconds
    done  
}

# 清理host集群待清理的资源
HostDeleting() {
    # 如果存在这个
    for resource in "${resources[@]}"  
    do  
        result=$(kubectl  get $resource -o yaml | grep deletionGracePeriodSeconds)  
  
        if [ -n "$result" ]; then  
            echo "---------> $resource  下存在 正在被标记删除的资源。"  
            #如果存在则删除
            kubectl patch  $resource $result --type json -p '[{"op": "replace", "path": "/metadata/finalizers", "value": []}]'
        else
            echo "---------> $resource 无被标记回收的资源"
        fi
    done  
}

###########################结束：备host集群清理待回收资源##############################

# 请给定环境变量
Help() {
    echo "提供参数
        get                 获取webhook、workspace和namespace关系
        deleteWebhook       删除集群webhook
        ddeleteWsNS         删除集群workspace和namespace关系
        setWebhook          恢复集群webhook
        setWsNs             恢复集群workspace和namespace关系
        checkHostDel        检查host是否存在别标记回收的资源
        delHost             清理host是否存在别标记回收的资源
    "
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

if [ -z "$1" ]; then
    Help
elif [ "$1" = "checkHostDel" ]; then
    # 检查host是否存在别标记回收的资源
    CheckHostDeleting
elif [ "$1" = "delHost" ]; then
    # 清理host是否存在别标记回收的资源
    HostDeleting
else

    # 检查目录是否存在  
    if [ ! -d "members" ]; then  
        mkdir members  
    fi  
    
    if [ ! -d "namespaces" ]; then  
        mkdir namespaces  
    fi  
    
    if [ ! -d "webhooks" ]; then  
        mkdir webhooks  
    fi

    # 执行操作
    for CLUSTER_NAME in "${CLUSTER_NAMES[@]}"; do
        # 定义kubeconfig路径
        KUBECONFIG_PATH=./members/member_kubeconfig_${CLUSTER_NAME}.yaml

        echo "----------------------------------------------> 正在操作的集群是：$CLUSTER_NAME"
        if [ "$1" = "get" ]; then
            # 获取webhook以及ws和ns绑定关系
            Get_webhook_ws_ns
        elif [ "$1" = "deleteWebhook" ]; then
            # 备份并删除 webhook 函数
            Remote_backup_and_delete_webhook
        elif [ "$1" = "deleteWsNS" ]; then

            echo "--------------------| 解绑ws和ns关系"
            # 调用解绑 workspace 和 namespace 的关系
            Remote_unbind_workspace_namespace
        elif [ "$1" = "setWebhook" ]; then
            # 恢复 webhook
            Restore_webhook
        elif [ "$1" = "setWsNs" ]; then
            # 恢复ws和关系
            Restore_ws_ns
        else
            Help
        fi
    done
fi