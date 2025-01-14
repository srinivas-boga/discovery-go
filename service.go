package discovery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServiceInstance(serviceDiscoveryClient *ServiceDiscoveryClient) gin.HandlerFunc {

	return func(c *gin.Context) {

		serviceName := c.Param("service_name")
		serviceInstance := serviceDiscoveryClient.GetServiceInstance(serviceName)
		c.JSON(http.StatusOK, serviceInstance)
	}
}

func RegisterService(serviceDiscoveryClient *ServiceDiscoveryClient) gin.HandlerFunc {

	return func(c *gin.Context) {

		var serviceInstance ServiceInstance
		serviceName := c.Param("service_name")
		err := c.BindJSON(&serviceInstance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		serviceDiscoveryClient.RegisterService(serviceName, serviceInstance)
		c.JSON(http.StatusOK, gin.H{"message": "Service instance registered successfully"})
	}
}
