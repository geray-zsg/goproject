#!/bin/bash

IMAGES=(
registry.cn-beijing.aliyuncs.com/kse/ks-apiserver:v3.3.1-cmft-rc.1
registry.cn-beijing.aliyuncs.com/kse/ks-controller-manager:v3.3.1-cmft-rc.1
registry.cn-beijing.aliyuncs.com/kse/kubeeye-controller:v1.0.0
registry.cn-beijing.aliyuncs.com/kse/kubeeye-apiserver:v1.0.0
registry.cn-beijing.aliyuncs.com/kse/kubeeye-job:v1.0.0
registry.cn-beijing.aliyuncs.com/kse/conformance:v3.3.1-cmft-rc.1
registry.cn-beijing.aliyuncs.com/kse/fluid-apiserver:v0.1.0
registry.cn-beijing.aliyuncs.com/kse/kruise-apiserver:v0.1.0
registry.cn-beijing.aliyuncs.com/kse/ks-installer:v3.3.1-cmft-rc.1
)



for image in ${IMAGES[@]}; do
	docker pull $image

done

