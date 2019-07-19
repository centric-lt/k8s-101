package kubernetes

import (
	"io/ioutil"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

type PodInfo struct {
	Ip       string
	Hostname string
}

type KubernetesClient interface {
	GetCurrentPodInfo() *PodInfo
}

type kubeClient struct {
	clientset *kubernetes.Clientset
	namespace string
}

func NewKubernetesClient() KubernetesClient {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ns, err := getCurrentNamesapce()
	if err != nil {
		panic(err)
	}
	return &kubeClient{
		clientset: clientset,
		namespace: ns,
	}
}

func getCurrentNamesapce() (string, error) {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (k *kubeClient) GetCurrentPodInfo() *PodInfo {
	podName := os.Getenv("HOSTNAME")
	pod, err := k.clientset.CoreV1().Pods(k.namespace).Get(podName, v1.GetOptions{})
	if err != nil {
		return nil
	}
	return &PodInfo{
		Ip:       pod.Status.HostIP,
		Hostname: pod.Name,
	}
}
