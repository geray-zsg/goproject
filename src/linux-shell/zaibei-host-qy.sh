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
    kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration mutating-webhook-ks-cfg || echo "Webhook not found" 
    echo "----------> ws和ns关系"
    kubectl  --kubeconfig=${KUBECONFIG_PATH} get ns --show-labels | grep workspace  || echo "Workspace not found"
}

# 远程备份并删除 webhook
Remote_backup_and_delete_webhook() {
    # webhook名称
    local RESOURCE_NAME="mutating-webhook-ks-cfg"

    echo "正在操作：mutating-webhook-ks-cfg"
    # kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME
    if kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME > /dev/null 2>&1 ; then
        echo "资源存在，正在备份..."
        kubectl --kubeconfig=${KUBECONFIG_PATH} get MutatingWebhookConfiguration $RESOURCE_NAME -o yaml > ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml

        echo "备份完成，正在删除资源..."
        kubectl --kubeconfig=${KUBECONFIG_PATH} delete -f ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml

        echo "资源已删除。"
    else
        echo "${CLUSTER_NAME} 集群不存在webhook ${RESOURCE_NAME}。"
    fi

}

# 恢复webhook
Restore_webhook() {
    # webhook名称
    local RESOURCE_NAME="mutating-webhook-ks-cfg"
    
    if [ -f "./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml" ]; then
        kubectl --kubeconfig=${KUBECONFIG_PATH}  create -f ./webhooks/${CLUSTER_NAME}_${RESOURCE_NAME}-bak.yaml
    else
        echo "原${CLUSTER_NAME} 集群不存在该webhook，无需恢复"
    fi
}

# 远程解绑 workspace 和 namespace 的关系
Remote_unbind_workspace_namespace() {

    # 备份ws和ns
    #kubectl --kubeconfig=${KUBECONFIG_PATH} get ns --show-labels > ./namespaces/ws_ns_${CLUSTER_NAME}
    # 如果没有workspace这个则不添加到备份清单中
    kubectl --kubeconfig=${KUBECONFIG_PATH} get namespace -o=jsonpath='{range .items[*]}{.metadata.name},{.metadata.labels.kubesphere\.io/workspace}{"\n"}{end}' | awk -F',' '$2 != "" {print}' > ./namespaces/ws_ns_${CLUSTER_NAME}


    # 在这里使用 kubectl 连接到对应的 member 集群执行解绑的命令
    # 以循环方式对所有 namespace 执行解绑操作
    #for i in `kubectl  --kubeconfig=${KUBECONFIG_PATH}  get ns | grep -v NAME | awk '{print $1}'`;do kubectl  --kubeconfig=${KUBECONFIG_PATH}  label ns $i kubesphere.io/workspace- && kubectl  --kubeconfig=${KUBECONFIG_PATH}  patch ns $i -p '{"metadata":{"ownerReferences":[]}}' --type=merge;done
    kubectl --kubeconfig=${KUBECONFIG_PATH} get namespace -o=jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}' | xargs -I{} kubectl --kubeconfig=${KUBECONFIG_PATH} patch namespace {} -p '{"metadata": {"labels": {"kubesphere.io/workspace": null}}}'

}


# 恢复ns和ws关系
Restore_ws_ns() {

    if [ -f "./namespaces/ws_ns_${CLUSTER_NAME}" ]; then
        #awk '/workspace/{split($NF,w,",");for(i=1;i<=length(w);i++){if(w[i]~"kubesphere.io/workspace"){cmd="kubectl --kubeconfig=${KUBECONFIG_PATH} label ns "$1" "w[i]" --overwrite";system(cmd)}}}' ./namespaces/ws_ns_${CLUSTER_NAME}
        cat ./namespaces/ws_ns_${CLUSTER_NAME} | while IFS=, read -r namespace label; do kubectl --kubeconfig=${KUBECONFIG_PATH} label namespace "$namespace" kubesphere.io/workspace="$label"; done


    else
        echo "./namespaces/ws_ns_${CLUSTER_NAME} 备份文件不存在"
    fi
}
###########################结束：member集群##############################

###########################开始：备host集群清理待回收资源##############################
resources=(
    "Group"  
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
        # kubectl get $resource -o yaml | grep deletionGracePeriodSeconds -q || echo "$resource not found"
        kubectl get $resource -o jsonpath='{range .items[*]}{.metadata.name}{","}{.metadata.deletionGracePeriodSeconds}{"\n"}{end}' | awk -F',' '$2 == 0 {print $1}'
    done  
}

# 清理host集群待清理的资源
HostDeleting() {
    # 如果存在这个
    for resource in "${resources[@]}"  
    do  
        result=$(kubectl get $resource -o jsonpath='{range .items[*]}{.metadata.name}{","}{.metadata.deletionGracePeriodSeconds}{"\n"}{end}' | awk -F',' '$2 == 0 {print $1}')  

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
        setWebhook          恢复集群webhook
        deleteWsNS          删除集群workspace和namespace关系
        setWsNs             恢复集群workspace和namespace关系
        checkHostDel        检查host是否存在别标记回收(deletionGracePeriodSeconds)的资源
        delHost             清理host是否存在别标记回收(deletionGracePeriodSeconds)的资源
        checkKubeFed        检查host集群kubefed服务
        stopKubeFed         停止host集群kubefed服务
        startKubeFed        启动host集群kubefed服务
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

if [ "$1" = "checkHostDel" ]; then
    # 检查host是否存在别标记回收的资源
    CheckHostDeleting
elif [ "$1" = "delHost" ]; then
    # 清理host是否存在别标记回收的资源
    HostDeleting
elif [ "$1" = "checkKubeFed" ]; then
    # 检查kubefed服务
    kubectl -n kube-federation-system get pods | grep kubefed-controller-manager || echo "kubefed-controller-manager Pod Not Found"
elif [ "$1" = "stopKubeFed" ]; then
    # 停止kubefed服务
    kubectl scale --replicas=0 -n kube-federation-system deploy kubefed-controller-manager
elif [ "$1" = "startKubeFed" ]; then
    # 停止kubefed服务
    kubectl scale --replicas=2 -n kube-federation-system deploy kubefed-controller-manager
                       
elif [ "$1" = "get" ] || [ "$1" = "deleteWebhook" ] || [ "$1" = "deleteWsNS" ] || [ "$1" = "setWebhook" ] || [ "$1" = "setWsNs" ]; then

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
        fi
    done
else
    Help
fi