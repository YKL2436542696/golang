package main

import (
	"context"
	"flag"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

//1. Resource方法指定了本次操作的资源类型；
//2. List方法向kubernetes发起请求；
//3. FromUnstructured将Unstructured数据结构转成PodList，其原理前面已经分析过；

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

	dynamicClient, err := dynamic.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	// dynamicClient的唯一关联方法所需的入参
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}

	// 使用dynamicClient的查询列表方法，查询指定namespace下的所有pod
	// 注意此方法返回的数据结构类型是UnstructuredList
	unstructObj, err := dynamicClient.
		Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		panic(err.Error())
	}

	// 实例化一个PodList数据结构，用于接收从unstructObj转换后的结果
	podList := &apiv1.PodList{}

	//转换
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)

	if err != nil {
		panic(err.Error())
	}

	// 表头
	fmt.Printf("namespace\t status\t\t name\n")

	for _, d := range podList.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}

}
