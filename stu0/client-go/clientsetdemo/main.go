package main

import (
	"context"
	"flag"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/utils/pointer"
	"path/filepath"
)

const (
	NAMESPACE       = "test-clientset"
	DEPLOYMENT_NAME = "client-test-deployment"
	SERVICE_NAME    = "client-test-service"
)

func main() {
	var kubeconfig *string

	// home目录
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// 获取用户输入的操作类型，默认是create，还可以输入clean，用于清理所有资源
	operate := flag.String("operate", "create", "operate type : create or clean")
	flag.Parse()

	// 从本机加载kubeconfig配置文件
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	// 实例化clientset对象
	clienset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("operation is %v\n", *operate)

	// 如果执行清理操作
	if "clean" == *operate {
		clean(clienset)
	} else {
		//创建namespace
		createNamespace(clienset)
		//创建deployment
		createDeployment(clienset)
		//创建service
		createService(clienset)
	}
}

//清理创建的所有资源
func clean(clientset *kubernetes.Clientset) {
	emptyDeleteOptions := metav1.DeleteOptions{}

	//删除service
	if err := clientset.CoreV1().Services(NAMESPACE).Delete(context.TODO(), SERVICE_NAME, emptyDeleteOptions); err != nil {
		panic(err.Error())
	}

	//删除deployment
	if err := clientset.AppsV1().Deployments(NAMESPACE).Delete(context.TODO(), DEPLOYMENT_NAME, emptyDeleteOptions); err != nil {
		panic(err.Error())
	}

	//删除namespace
	if err := clientset.CoreV1().Namespaces().Delete(context.TODO(), NAMESPACE, emptyDeleteOptions); err != nil {
		panic(err.Error())
	}
}

// 新建namespace
func createNamespace(clientset *kubernetes.Clientset) {
	namespaceClient := clientset.CoreV1().Namespaces()

	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NAMESPACE,
		},
	}

	result, err := namespaceClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create namespace %s \n", result.GetName())
}

// 新建service
func createService(clientset *kubernetes.Clientset) {
	//得到service的客户端
	serviceClient := clientset.CoreV1().Services(NAMESPACE)

	//实例化一个数据结构
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: SERVICE_NAME,
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{{
				Name:     "http",
				Port:     8080,
				NodePort: 30080,
			},
			},
			Selector: map[string]string{
				"app": "tomcat",
			},
			Type: apiv1.ServiceTypeNodePort,
		},
	}

	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Create service %s \n", result.GetName())
}

// 新建deployment
func createDeployment(clientset *kubernetes.Clientset) {
	//得到deployment的客户端
	deploymentClient := clientset.
		AppsV1().
		Deployments(NAMESPACE)

	//实例化一个数据结构
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: DEPLOYMENT_NAME,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "tomcat",
				},
			},

			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "tomcat",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "tomcat",
							Image:           "tomcat:8.0.18-jre8",
							ImagePullPolicy: "IfNotPresent",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolSCTP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create deployment %s \n", result.GetName())

}
