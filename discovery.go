package discovery

import (
	"discovery/database"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceInstance struct {
	IpAddr string
	Port   int
}

type ServiceDiscoveryClient struct {
	db_connection *mongo.Client
	cache         *Cache
}

func NewServiceDiscoveryClient() *ServiceDiscoveryClient {
	return &ServiceDiscoveryClient{
		cache:         NewCache(10000),
		db_connection: database.GetMongoDBClient(),
	}
}

func (s *ServiceDiscoveryClient) RegisterService(serviceName string, serviceInstance ServiceInstance) {
	deletedItem, isDeleted := s.cache.Set(serviceName, serviceInstance)
	if isDeleted {
		// write to database
		database.GetMongoDBCollection(s.db_connection, "service_discovery", "service_instances").InsertOne(context.Background(), deletedItem)
	}
}

func (s *ServiceDiscoveryClient) GetServiceInstance(serviceName string) (ServiceInstance, error) {
	if cachedServiceInstance := s.cache.Get(serviceName); cachedServiceInstance != nil {
		return cachedServiceInstance.(cacheItem).value.(ServiceInstance), nil
	}

	var serviceInstance ServiceInstance
	err := database.GetMongoDBCollection(s.db_connection, "service_discovery", "service_instances").FindOne(context.Background(), serviceName).Decode(&serviceInstance)
	if err != nil {
		return ServiceInstance{}, err
	}
	s.cache.Set(serviceName, serviceInstance)
	return serviceInstance, nil
}
