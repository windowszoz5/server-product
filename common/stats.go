package common

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"server-product/compose"
)

type ServerStats struct {
	OutPayload *stats.OutPayload
	InPayload  *stats.InPayload
	InHeader   *stats.InHeader
	LocalAddr  string
	TractId    string
}

func (h *ServerStats) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	//获取上下文数据
	if md, b := metadata.FromIncomingContext(ctx); b {
		if len(md.Get("X-Track-id")) > 0 {
			h.TractId = md.Get("X-Track-id")[0]
		}
	}

	return ctx
}

func (h *ServerStats) HandleRPC(ctx context.Context, s stats.RPCStats) {
	switch sd := s.(type) {
	case *stats.InHeader:
		h.InHeader = sd
		break
	//记录数据输出
	case *stats.OutPayload:
		h.OutPayload = sd
		break
	//记录数据输入
	case *stats.InPayload:
		h.InPayload = sd
		break
	case *stats.End:
		h.Kiblog()
	}
}

func (h *ServerStats) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	h.LocalAddr = info.LocalAddr.String()
	return ctx
}

func (h *ServerStats) HandleConn(ctx context.Context, rpcStats stats.ConnStats) {
	fmt.Println(rpcStats)
}

type KibnanaLog struct {
	Url       string `json:"url"`       //请求地址
	Body      string `json:"body"`      //post请求体
	Ip        string `json:"ip"`        //请求IP
	ReqHeader string `json:"reqHeader"` //请求头
	TractId   string `json:"tractId"`   //请求UID
	Method    string `json:"method"`    //请求方法
	Message   string `json:"message"`   //请求相应
}

func (h *ServerStats) Kiblog() {
	esData := KibnanaLog{
		Url:     h.InHeader.FullMethod,
		Body:    string(h.InPayload.Data),
		TractId: h.TractId,
		//ReqHeader: ,
		Ip:      h.LocalAddr,
		Method:  h.InHeader.FullMethod,
		Message: string(h.OutPayload.Data),
	}
	fmt.Println()
	_, err := compose.EsClient.Index().
		Index("master").
		Type("server-product").
		BodyJson(esData).
		Do()
	if err != nil {
		fmt.Println("写入es错误")
	}
}
