package lib

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	consulapi "github.com/hashicorp/consul/api"
	"os"
)

func init() {
	consulRegister()
}

func consulRegister() {
	_ = web.LoadAppConfig("ini", "conf/app.conf")
	ConsulAddress, _ := web.AppConfig.String("ConsulAddress")
	ServiceType, _ := web.AppConfig.String("ServiceType")
	Port, _ := web.AppConfig.Int("Port")
	ServiceAdvHost, _ := web.AppConfig.String("ServiceAdvHost")

	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = ConsulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
		os.Exit(1)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.Name = ServiceType
	registration.Port = Port
	registration.Tags = []string{ServiceType}
	registration.Address = ServiceAdvHost

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "300s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}
