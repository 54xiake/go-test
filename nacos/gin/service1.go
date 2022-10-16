package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var namingClient naming_client.INamingClient

func init() {
	namingClient = InitNacos()
}

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "8081",
		})
	})
	r.GET("/up", func(c *gin.Context) {
		Register()
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "up",
		})
	})
	r.GET("/down", func(c *gin.Context) {
		Deregister()
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "down",
		})
	})
	go Register()
	defer Deregister()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080

}

func InitNacos() naming_client.INamingClient {
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
	return namingClient
}

func Register() error {
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8801,
		ServiceName: "order-service",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	if success == false {
		return errors.New("注册失败")
	}
	return nil
}

func Deregister() error {
	success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8801,
		ServiceName: "order-service",
		Ephemeral:   true,
		Cluster:     "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	if success == false {
		return errors.New("踢除失败")
	}
	return nil
}
