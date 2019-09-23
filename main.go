/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: main.go
 * @time: 2019-05-17 19:08
 */

package main

import (
	"context"

	`github.com/astaxie/beego/logs`
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"

	grpc_grace `github.com/generalzgd/grpc-svr-frame/grpc-grace`
	"github.com/generalzgd/grpc-svr-frame/proto"

	// "github.com/bsm/grpclb"
)

type Svr struct {
}

func (p *Svr) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	// panic("implement me")
	return &proto.HelloResp{}, nil
}

// use "kill -31 pid" to restart svr graceful
func main() {
	// exampleGrpc()

	exampleNewResolver()
}

func exampleNewResolver() {
	const target = "helloworld"

	// balance := grpc.RoundRobin(grpclb.NewResolver(&grpclb.Options{Address:"127.0.0.1:8383"}))

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		logs.Error("did not connect:", err)
		return
	}
	defer func() {
		conn.Close()
	}()

	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "world"})
	if err != nil {
		logs.Error("could not greet:", err)
		return
	}

	logs.Info("Greeting:", r.Message)
}

func exampleGrpc() {
	logs.Info("start grace svr")

	addr := ":10011"
	s := grpc.NewServer()
	proto.RegisterDeployServer(s, &Svr{})
	reflection.Register(s)
	gr, err := grpc_grace.New(s, "tcp", addr)
	if err != nil {
		logs.Error("failed to new grace grpc.", err)
	}
	if err := gr.Serve(); err != nil {
		logs.Error("failed to serve:", addr)
	}

	logs.GetBeeLogger().Flush()
}
