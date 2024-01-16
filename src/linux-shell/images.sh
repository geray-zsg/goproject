#!/bin/bash

##########################
##  $1 指定存放文件名   ##
##########################


# 目标镜像仓库url
dsc_url="$3"

# 源集群的镜像URL
src_url="repos.cloud.cmft"

# 存放位置或密码
#save_path="/root/jszj/images/$2"

############对于字符串比较，你应该使用 = 而不是 ==。而且，最好使用双括号 [[ ... ]] 进行条件测试，因为它提供了更多的功能和安全性。
############另外，! 在条件测试中需要紧跟在 [[ ... ]] 或 [ ... ] 的后面，不能有空格。所以你的条件应该写成 [[ ! $2 = passwd_* ]] 或者 ! [ $2 = passwd_* ]。
if [[ ! $2 = passwd_* ]]; then
    if [ -n "$2" ];then
        save_path="./$2"
        echo $save_path

        #IMAGES=`cat /root/jszj/images/image-file/${2}`
        IMAGES=`cat ./image-file/${2}`
    fi
else
    # ${2#passwd_} 重开头删除执行的字符串，获取真正的密码passwd_除外的部分，密码中不能带有！
    harbor_passwd=${2#passwd_}
    #harbor_passwd=${password%bash}  # 从末尾删除 "bash"
    #harbor_passwd="${password}"     # 使用引号包裹密码
    echo "密码是： $harbor_passwd"
fi



# 拉取镜像
pull_images() {
    for image in ${IMAGES[@]}; do
        echo "----------> 拉取镜像: ${src_url}/${image}"
        docker pull ${src_url}/${image}

        # 获取上一个命令的退出状态码
        exit_code=$?

        # 检查退出状态码，0 表示成功，非0 表示失败
        if [ $exit_code -ne 0 ]; then
            echo "拉取镜像 ${src_url}/${image} 失败            <-----------------------"
        fi
    done
}

# 保存镜像啊ing
save_images() {
    for image in ${IMAGES[@]}; do
        echo "----------> 保存镜像: ${src_url}/${image} 到 ${save_path}/${SAVE_IMAGE_NAME} "
        SAVE_IMAGE_NAME=$(echo $image | cut -d "/" -f 2 | tr ':' '_')
        #echo "SAVE_IMAGE_NAME: $SAVE_IMAGE_NAME"
        docker save -o ${save_path}/${SAVE_IMAGE_NAME} ${src_url}/${image}

        # 获取上一个命令的退出状态码
        exit_code=$?

        # 检查退出状态码，0 表示成功，非0 表示失败
        if [ $exit_code -ne 0 ]; then
            echo "保存镜像 ${src_url}/${image} 失败            <-----------------------"
        fi
    done
}

# tag镜像
tag_images() {
    for image in ${IMAGES[@]}; do
        echo "----------> tag镜像: ${src_url}/${image} 到 $dsc_url"
        docker tag ${src_url}/${image} ${dsc_url}/${image}
    done
}

# 推送镜像
load_images() { 
    for image in ${IMAGES[@]}; do
        echo "----------> 保存镜像: ${src_url}/${image} 到 ${save_path}/${SAVE_IMAGE_NAME} "
        SAVE_IMAGE_NAME=$(echo $image | cut -d "/" -f 2 | tr ':' '_')
        #echo "SAVE_IMAGE_NAME: $SAVE_IMAGE_NAME"
        docker load -i ${save_path}/${SAVE_IMAGE_NAME}

        # 获取上一个命令的退出状态码
        exit_code=$?

        # 检查退出状态码，0 表示成功，非0 表示失败
        if [ $exit_code -ne 0 ]; then
            echo "加载镜像 ${src_url}/${image} 失败            <-----------------------"
        fi
    done
}

# 推送镜像
push_images() { 
    for image in ${IMAGES[@]}; do
        echo "----------> push镜像: ${src_url}/${image} 到 $dsc_url"
        docker push ${dsc_url}/${image}

                # 获取上一个命令的退出状态码
        exit_code=$?

        # 检查退出状态码，0 表示成功，非0 表示失败
        if [ $exit_code -ne 0 ]; then
            echo "推送镜像 ${dsc_url}/${image} 失败            <-----------------------"
        fi
    done
}

# 创建项目
create_project() {

    url="https://${dsc_url}"
    user="admin"
    passwd="$harbor_passwd"

    harbor_projects=(
        fluidcloudnative
        openkruise
        sonobuoy 
    )

    for project in "${harbor_projects[@]}"; do
        echo "创建项目 $project"
        #curl -u "${user}:${passwd}" -k -X POST -H "Content-Type: application/json" "${url}/api/v2.0/projects" -d "{ \"project_name\": \"${project}\", \"public\": true}"
    done

}


Help() {
    echo "提供参数
        pull                镜像拉取
        push                镜像推送
        tag                 镜像tag
        save                镜像保存
        load                加载镜像
        createProject       创建harbor仓库（密码格式：passwd_<仓库密码>）【密码不能包含！,shell中有特殊意义】

        例如：
        ./images.sh pull [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]
        ./images.sh save [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]
        ./images.sh tag  [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]  [repos.cloud.test]
        ./images.sh load [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]
        ./images.sh push [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]  [repos.cloud.test]
        ./images.sh createProject  passwd_<password> [url]
    "
}


case $1 in
    "pull")
        pull_images
        ;;
    "save")
        # 保存镜像啊ing
        save_images $2
        ;;
    "tag")
        tag_images
        ;;
    "load")
       load_images
        ;;
    "push")
       push_images
        ;;
    "createProject")
       create_project 
        ;;
    *)
        Help
        ;;
esac