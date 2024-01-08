#!/bin/bash

##########################
##  $1 指定存放文件名   ##
##########################


# 目标镜像仓库url
dsc_url="$3"

# 源集群的镜像URL
src_url="repos.cloud.cmft"

# 存放位置
save_path="/root/jszj/images/$2"
echo $save_path

IMAGES=`cat /root/jszj/images/image-file/${2}`

# 拉取镜像
pull_images() {
    for image in ${IMAGES[@]}; do
        echo "----------> 拉取镜像: ${src_url}${image}"
        docker pull ${src_url}/${image}
    done
}

# 保存镜像啊ing
save_images() {
    for image in ${IMAGES[@]}; do
        echo "----------> 保存镜像: ${src_url}/${image} 到 ${save_path}/${SAVE_IMAGE_NAME} "
        SAVE_IMAGE_NAME=$(echo $image | cut -d "/" -f 2 | tr ':' '_')
        #echo "SAVE_IMAGE_NAME: $SAVE_IMAGE_NAME"
        docker save -o ${save_path}/${SAVE_IMAGE_NAME} ${src_url}/${image}
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
    done
}

# 推送镜像
push_images() { 
    for image in ${IMAGES[@]}; do
        echo "----------> push镜像: ${src_url}/${image} 到 $dsc_url"
        docker push ${dsc_url}/${image}
    done
}

Help() {
    echo "提供参数
        pull       镜像拉取
        push       镜像推送
        tag        镜像tag
        save       镜像保存
        load       加载镜像

        例如：
        ./images.sh pull [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]
        ./images.sh save [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]
        ./images.sh tag  [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]  [repos.cloud.test]
        ./images.sh load [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]
        ./images.sh push [chaos-mesh|kubeeye|fluidcloudnative|gateway|openkruise|sonobuoy]  [repos.cloud.test]
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
    *)
        Help
        ;;
esac