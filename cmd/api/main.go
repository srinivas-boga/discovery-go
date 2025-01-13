package main

import (
	"github.com/gin-gonic/gin"

	"discovery"
)

func main() {
	r := gin.Default()

	serviceDiscoveryClient := discovery.NewServiceDiscoveryClient()

	r.GET("/service/:service_name", discovery.GetServiceInstance(serviceDiscoveryClient))
	r.POST("/service/:service_name", discovery.RegisterService(serviceDiscoveryClient))
	r.Run(":8080")
}
