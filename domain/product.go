package domain

import (
	"context"
	"fmt"
	"server-product/rpc"
)

type Product struct {
	rpc.UnimplementedProductServer
}

func (Product) Ping(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	fmt.Println("ping函数处理", req.Ping)
	return &rpc.Response{Pong: "21321312"}, nil
}
