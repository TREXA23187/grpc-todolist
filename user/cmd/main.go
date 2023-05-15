package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"grpc-todolist/user/config"
	"grpc-todolist/user/discovery"
	"grpc-todolist/user/internal/handler"
	"grpc-todolist/user/internal/repository"
	"grpc-todolist/user/internal/service"
	"net"
)

func main() {
	config.InitConfig()
	repository.InitDB()

	// etcd
	etcdAddr := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewRegister(etcdAddr, logrus.New())

	grpcAddress := viper.GetString("server.grpcAddress")
	userNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()

	service.RegisterUserServiceServer(server, handler.NewUserService())
	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err = etcdRegister.Register(userNode, 10); err != nil {
		panic(err)
	}
	if err = server.Serve(listen); err != nil {
		panic(err)
	}

}
