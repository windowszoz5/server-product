package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server-product/common"
	"server-product/compose"
	"server-product/domain"
	"server-product/rpc"
)

func main() {
	// 监听端口
	url := fmt.Sprintf("localhost:%d", 3366)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to list2213en: %v", err)
	}

	//初始ES
	compose.InitEs()

	// 实例化server
	s := grpc.NewServer(grpc.StatsHandler(&common.ServerStats{}))

	//注册服务
	rpc.RegisterProductServer(s, &domain.Product{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
