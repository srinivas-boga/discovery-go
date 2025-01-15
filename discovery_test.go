package discovery

import (
	"testing"
)

// TestRegisterGetService tests the RegisterService and GetServiceInstance functions
func TestRegisterGetService(t *testing.T) {
	// Create a new service discovery client
	serviceDiscoveryClient := NewServiceDiscoveryClient()

	// Create a new service instance
	serviceInstance := ServiceInstance{
		IpAddr: "localhost",
		Port:   8080,
	}

	// Register the service instance
	serviceDiscoveryClient.RegisterService("test-service", serviceInstance)

	// Get the service instance
	registeredServiceInstance, err := serviceDiscoveryClient.GetServiceInstance("test-service")
	if err != nil {
		t.Errorf("Error getting service instance: %s", err.Error())
	}

	// Check if the service instance is correct
	if registeredServiceInstance.IpAddr != serviceInstance.IpAddr {
		t.Errorf("Expected host to be %s, got %s", serviceInstance.IpAddr, registeredServiceInstance.IpAddr)
	}

	if registeredServiceInstance.Port != serviceInstance.Port {
		t.Errorf("Expected port to be %d, got %d", serviceInstance.Port, registeredServiceInstance.Port)
	}
}
