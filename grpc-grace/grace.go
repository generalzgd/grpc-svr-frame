/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: grace.go
 * @time: 2019-05-20 11:28
 */

package grpc_grace

import (
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/astaxie/beego"
	"github.com/facebookgo/grace/gracenet"
	"google.golang.org/grpc"
)

var (
	envKey = "LISTEN_FDS"
	ppid   = os.Getppid()
)

type GraceGrpc struct {
	server   *grpc.Server
	grace    *gracenet.Net
	listener net.Listener
	errors   chan error
}

func New(s *grpc.Server, netType, addr string) (*GraceGrpc, error) {
	g := GraceGrpc{
		server: s,
		grace:  &gracenet.Net{},
		errors: make(chan error),
	}

	lis, err := g.grace.Listen(netType, addr)
	if err != nil {
		return nil, err
	}

	g.listener = lis
	return &g, nil
}

func (p *GraceGrpc) startServe() {
	if err := p.server.Serve(p.listener); err != nil {
		p.errors <- err
	}
}

func (p *GraceGrpc) handleSignal() <-chan struct{} {
	end := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 10)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
		for {
			sig := <-ch
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				//beego.Debug("handleSignal:", sig)
				signal.Stop(ch)
				p.server.GracefulStop()
				close(end)
				return
			case syscall.SIGUSR2:
				//beego.Debug("handleSignal:", sig)
				//p.server.GracefulStop()
				if _, err := p.grace.StartProcess(); err != nil {
					p.errors <- err
				}
			}
		}
	}()
	return end
}

func (p *GraceGrpc) Serve() error {
	if p.listener == nil {
		return errors.New("struct not initialized")
	}

	inherit := os.Getenv(envKey) != ""
	pid := os.Getpid()
	addr := p.listener.Addr().String()

	if inherit {
		if ppid == 1 {
			beego.Info("listening on init address", addr)
		} else {
			beego.Info("graceful hand off with new pid replace old", addr, pid, ppid)
		}
	} else {
		beego.Info("start serving pid", addr, pid)
	}

	go p.startServe()

	if inherit && ppid != 1 {
		if err := syscall.Kill(ppid, syscall.SIGTERM); err != nil {
			return err
		}
	}

	end := p.handleSignal()
	select {
	case err := <-p.errors:
		return err
	case <-end:
		beego.Info("exiting pid", pid)
	}
	return nil
}
