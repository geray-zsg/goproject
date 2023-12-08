#!/bin/bash  
  
resources=("Group"  
            "GlobalRole"  
            "GlobalRoleBinding"  
            "workspaces.tenant.kubesphere.io"  
            "workspacetemplate"  
            "federatedworkspaces.types.kubefed.io"  
            "WorkspaceRole"  
            "WorkspaceRoleBinding"  
            "federatedusers.types.kubefed.io"  
            "users.iam.kubesphere.io")  
  
if [ -n "$1" ] && [ "$1" = "checkDelete" ]; then  
    for resource in "${resources[@]}"  
    do  
        result=$(kubectl get $resource -o yaml | grep deletionGracePeriodSeconds)  
  
        if [ -n "$result" ]; then  
            echo "---------> $resource  下存在 正在被标记删除的资源。"  
            echo "$result"
        else
            echo "---------> $resource 无被标记回收的资源"
        fi  
    done  
else  
    for resource in "${resources[@]}"  
    do  
        echo "------------> 资源名称名称：$resource"  
        kubectl get $resource  
    done  
fi