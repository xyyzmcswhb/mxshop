package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type RegisterConsul struct {
	Host string
	Port int
}

type RegisterClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	Deregister(serviceId string) error
}

func NewRegisterClient(host string, port int) RegisterClient {
	return &RegisterConsul{
		Host: host,
		Port: port,
	}
}

func (r *RegisterConsul) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	//获取consul实例
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "300s", //5s检查一次
		DeregisterCriticalServiceAfter: "10s",  //10s有效期
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	//client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *RegisterConsul) Deregister(serviceId string) error {
	cfg := api.DefaultConfig()
	//获取consul实例
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
