package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {

	var kubeconfig *string

	// home目录
	if home := homedir.HomeDir(); home != "" {
		// 如果输入了kubeconfig参数，该参数的值就是kubeconfig文件的绝对路径，
		// 如果没有输入kubeconfig参数，就用默认路径~/.kube/config
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		// 如果取不到当前用户的home目录，就没办法设置kubeconfig的默认目录了，只能从入参中取
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	//新建discoveryClient的实例
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	//获取所有分组和资源数据
	APIGroup, APIResourceListSlice, err := discoveryClient.ServerGroupsAndResources()

	if err != nil {
		panic(err.Error())
	}

	//先开看Group信息
	fmt.Printf("APIGroup :\n\n %v\n\n\n\n", APIGroup)

	// APIResourceListSlice是个切片，里面的每个元素代表一个GroupVersion及其资源
	for _, singleAPIResourceList := range APIResourceListSlice {

		// GroupVersion是个字符串，例如”apps/v1"
		groupVersionStr := singleAPIResourceList.GroupVersion

		// parseGroupVersion方法将字符串转成数据结构
		gv, err := schema.ParseGroupVersion(groupVersionStr)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("*****************************************************************")
		fmt.Printf("GV string [%v]\nGV struct [%#v]\nresources :\n\n", groupVersionStr, gv)

		// APIResource字段是个切片，里面是当前GroupVersion下的所有资源
		for _, singleAPIResource := range singleAPIResourceList.APIResources {
			fmt.Printf("%v\n", singleAPIResource.Name)
		}

	}

}
