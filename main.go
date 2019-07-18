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

	"github.com/astaxie/beego"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"

	"generalzgd/grpc-svr-frame/grpcgrace"
	"generalzgd/grpc-svr-frame/proto"

	//"github.com/bsm/grpclb"
)

type Svr struct {
}

func (p *Svr) SayHello(context.Context, *proto.HelloReq) (*proto.HelloResp, error) {
	//panic("implement me")
	return &proto.HelloResp{}, nil
}

// use "kill -31 pid" to restart svr graceful
func main() {
	//exampleGrpc()

	exampleNewResolver()
}

func exampleNewResolver() {
	const target = "helloworld"

	//balance := grpc.RoundRobin(grpclb.NewResolver(&grpclb.Options{Address:"127.0.0.1:8383"}))

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		beego.Error("did not connect:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "world"})
	if err != nil {
		beego.Error("could not greet:", err)
		return
	}

	beego.Info("Greeting:", r.Message)
}

func exampleGrpc() {
	beego.Info("start grace svr")

	addr := ":10011"
	s := grpc.NewServer()
	proto.RegisterDeployServer(s, &Svr{})
	reflection.Register(s)
	gr, err := grpcgrace.New(s, "tcp", addr)
	if err != nil {
		beego.Error("failed to new grace grpc.", err)
	}
	if err := gr.Serve(); err != nil {
		beego.Error("failed to serve:", addr)
	}

	beego.BeeLogger.Flush()
}
