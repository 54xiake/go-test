package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc"
)

type nacosRF struct {
}

func InitNacos() *nacosRF {
	//return &nacosRF{}
	d := &nacosRF{}
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./log",
		CacheDir:            "./cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			//Scheme:      "http",
		},
	}

	var err error
	d.client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	NacosInstance = d
	return d
}

func (this *nacosRF) Register(serverName string, port int) {
	ok, err := InitNacos().client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          utils.GetIp(),
		Port:        uint64(port),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{},
		ClusterName: "DEMO_SERVER",
		ServiceName: serverName,
		GroupName:   "DEMO_SERVER_GROUP",
		Ephemeral:   true,
	})
	if err != nil {
		panic(err)
	}
	if !ok {
		logger.Info("注册本服务发生错误")
		panic(errors.New("注册本服务发生错误"))
	}
}

//FindInstance 获取某个服务器连接信息
func (this *nacosRF) FindInstance(serverName string) (*ServerInfo, error) {
	ins, err := InitNacos().client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		Clusters:    []string{"DEMO_SERVER"},
		ServiceName: serverName,
		GroupName:   "DEMO_SERVER_GROUP",
	})
	if err != nil {
		return nil, err
	}
	return &ServerInfo{
		Ip:         ins.Ip,
		Port:       ins.Port,
		ServerName: ins.ServiceName,
	}, nil
}

func GetDemoOrderService(serviceName string) pb.RpcDemoServiceClient {
	info, err := nacosRF.NacosInstance.FindInstance(serviceName)
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

func main() {

	nacosRF.NacosInstance.Register("demo", 8300)

	demoClient := service.GetDemoOrderService("demo")
	result, err := demoClient.GetUserInfo(context.Background(), &pb.GetUserInfoRequest{
		Id: 10,
	})
	//
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
