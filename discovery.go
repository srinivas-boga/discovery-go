package discovery

import (
	"sync"
)

type ServiceInstance struct {
	IpAddr string
	Port   int
}

type ServiceDiscoveryClient struct {
	Instances map[string]ServiceInstance
	mutex     sync.RWMutex
}

func NewServiceDiscoveryClient() *ServiceDiscoveryClient {
	return &ServiceDiscoveryClient{
		Instances: make(map[string]ServiceInstance),
	}
}

func (s *ServiceDiscoveryClient) RegisterService(serviceName string, serviceInstance ServiceInstance) {
	s.mutex.Lock()
	s.Instances[serviceName] = serviceInstance
	s.mutex.Unlock()
}

func (s *ServiceDiscoveryClient) GetServiceInstance(serviceName string) ServiceInstance {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.Instances[serviceName]
}
