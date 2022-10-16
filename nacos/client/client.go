package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc"
)

func main() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "c9eda8a8-91fe-46f6-a167-6fd69d8d14d2", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// Create naming client for service discovery
	_, _ = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	// Create config client for dynamic configuration
	_, _ = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	// Another way of create naming client for service discovery (recommend)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	services, err := namingClient.GetService(vo.GetServiceParam{
		ServiceName: "order-service",
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
	})

	fmt.Println(services)

	instances, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: "order-service",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
	})

	fmt.Println(instances)

	instances, err = namingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: "order-service",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		HealthyOnly: true,
	})

	fmt.Println(instances)

	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "order-service",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
	})

	fmt.Println(instance)

	err = namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: "order-service",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			fmt.Println(services)
		},
	})
	fmt.Println(err)
}

type ServerInfo struct {
	Ip         string
	Port       uint64
	ServerName string
}

func FindInstance(serviceName string) (*ServerInfo, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "c9eda8a8-91fe-46f6-a167-6fd69d8d14d2", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
	})
	if err != nil {
		return nil, err
	}
	return &ServerInfo{
		Ip:         instance.Ip,
		Port:       instance.Port,
		ServerName: instance.ServiceName,
	}, nil
}

func GetDemoOrderService(serviceName string) pb.RpcDemoServiceClient {
	info, err := FindInstance(serviceName)
	if err != nil {
		return nil
	}
	c, err := grpc.Dial(fmt.Sprintf("%s:%d", info.Ip, info.Port), grpc.WithInsecure())
	if err != nil {
		return nil
	}
	p := pb.NewRpcDemoServiceClient(c)
	return p
}
