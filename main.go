package main

import (
	"github.com/go-micro-v4-demo/firstService/handler"
	pb "github.com/go-micro-v4-demo/firstService/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "gsnimi@sina.cn"
	version = "gsnimi@sina.cn"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterFirstServiceHandler(srv.Server(), new(handler.FirstService)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
