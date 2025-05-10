package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main() {
	// 1. 创建一个新的Consul客户端
	//_ = Register("192.168.1.7", 8022, "user-web", []string{"joyshop", "test", "walker"}, "user-web")
	//AllService()
	FilterService()
}

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.1.7:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	registration := new(api.AgentServiceRegistration)
	registration.ID = id
	registration.Name = name
	registration.Address = address
	registration.Port = port
	registration.Tags = tags
	// 生成对应的检查对象
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/health", address, port)
	check.Interval = "5s"
	check.Timeout = "5s"
	check.DeregisterCriticalServiceAfter = "10s"
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.1.7:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for _, service := range services {
		fmt.Println(service.Service)
	}

}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.1.7:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)

	if err != nil {
		panic(err)
	}
	for _, service := range services {
		fmt.Println(service.Service)
	}
}
