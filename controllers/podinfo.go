package controllers

import (
	"context"
	"github.com/centric-lt/k8s-101/gen/podinfo"
	"github.com/centric-lt/k8s-101/internal/kubernetes"
	"log"
)

// podinfo service example implementation.
// The example methods log the requests and return zero values.
type podinfosrvc struct {
	kubeClient kubernetes.KubernetesClient
	logger     *log.Logger
}

// NewPodinfo returns the podinfo service implementation.
func NewPodinfo(logger *log.Logger) podinfo.Service {
	return &podinfosrvc{
		kubeClient: nil, //kubernetes.NewKubernetesClient(),
		logger:     logger,
	}
}

// Pod implements pod.
func (s *podinfosrvc) Get(ctx context.Context) (*podinfo.Podinforesult, error) {
	res := &podinfo.Podinforesult{}
	s.logger.Print("podinfo.pod")
	res.Hostname = "asdasd"
	res.IP = "10.10.0.55"
	return res, nil
}
