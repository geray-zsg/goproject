/**
 * @Author: Geray
 * @Date: 2023/8/10 18:49:18
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/10 18:49:18
 * Description:  Reflector是Client-go中用来监听指定资源的组件，当资源发生变化是，例如：增、删、改等操作，会以时间的形式存入本地队列，然后有对应的方法处理
	在Reflector中核心部署就是List-Watch，其他功能基本也是围绕这块来搞的。
	在实例化Reflector的过程中，有一个ListerWatcher的接口对象，这个结构对象有两个方法，分别是List和Watch，这两个方法就是实现了前面介绍的List-Watch功能。
Reflector核心逻辑，可以简单归纳为三部分：
	List：调用List方法获取资源全部列表数据，转换为资源对象列表，然后保存到本地缓存；
	定时同步：定时器定时触发同步机制，定时更新缓存数据，在Reflector的结构体对象中，是可以配置定时同步的周期时间的；
	Watch：监听资源的变化，并且调用对应的时间处理函数来进行处理。

	Reflector组件对数据更新同步，都是基于ResourceVersion来进行的，每个资源对象都具有这个属性，当数据变化的时候这个值也会递增。



 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("watch开启...")

}
