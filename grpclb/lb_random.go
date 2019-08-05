/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: lb_random.go
 * @time: 2019-07-30 16:08
 */

package grpclb

import (
	"context"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
)

func init() {
	balancer.Register(newBuilder())
}

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilderWithConfig("grpc-random", &randomPickerBuilder{}, base.Config{})
}

//
//type RandomBuilder struct {
//	name string
//}

//func (*RandomBuilder) Build(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
//	//panic("implement me")
//	return &RandomBalancer{
//		cc:cc,
//		subConns: make(map[resolver.Address]balancer.SubConn),
//		scStates: make(map[balancer.SubConn]connectivity.State),
//		csEvltr:  &balancer.ConnectivityStateEvaluator{},
//	}
//}

//func (*RandomBuilder) Name() string {
//	//panic("implement me")
//	return "grpclb-random"
//}

// 随机策略
/*type RandomBalancer struct {
	cc balancer.ClientConn

	csEvltr *balancer.ConnectivityStateEvaluator
	state   connectivity.State

	subConns map[resolver.Address]balancer.SubConn
	scStates map[balancer.SubConn]connectivity.State
	picker   balancer.Picker

	config base.Config
}

func (p *RandomBalancer) HandleSubConnStateChange(sc balancer.SubConn, state connectivity.State) {
	panic("implement me")
}

func (p *RandomBalancer) HandleResolvedAddrs([]resolver.Address, error) {
	panic("implement me")
}

func (p *RandomBalancer) UpdateResolverState(s resolver.State) {
	//panic("implement me")
	addrSet := make(map[resolver.Address]struct{}, len(s.Addresses))
	for _, a := range s.Addresses {
		addrSet[a] = struct{}{}
		if _, ok := p.subConns[a]; !ok {
			// a is a new address (not existing in b.subConns).
			sc, err := p.cc.NewSubConn([]resolver.Address{a}, balancer.NewSubConnOptions{HealthCheckEnabled: p.config.HealthCheck})
			if err != nil {
				grpclog.Warningf("base.baseBalancer: failed to create new SubConn: %v", err)
				continue
			}
			p.subConns[a] = sc
			p.scStates[sc] = connectivity.Idle
			sc.Connect()
		}
	}
	for a, sc := range p.subConns {
		if _, ok := addrSet[a]; !ok {
			p.cc.RemoveSubConn(sc)
			delete(p.subConns, a)
		}
	}
}

func (p *RandomBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	//panic("implement me")
	s := state.ConnectivityState
	grpclog.Infof("base.baseBalancer: handle SubConn state change: %p, %v", sc, s)
	oldS, ok := p.scStates[sc]
	if !ok {
		grpclog.Infof("base.baseBalancer: got state changes for an unknown SubConn: %p, %v", sc, s)
		return
	}
	p.scStates[sc] = s
	switch s {
	case connectivity.Idle:
		sc.Connect()
	case connectivity.Shutdown:
		// When an address was removed by resolver, b called RemoveSubConn but
		// kept the sc's state in scStates. Remove state for this sc here.
		delete(p.scStates, sc)
	}

	//oldAggrState := p.state
	p.state = p.csEvltr.RecordTransition(oldS, s)

	p.cc.UpdateBalancerState(p.state, p.picker)
}

func (p *RandomBalancer) Close() {
}*/

type randomPickerBuilder struct {
}

func (*randomPickerBuilder) Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker {
	//panic("implement me")
	if len(readySCs) < 1 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	picker := RandomPicker{
		subConn: readySCs,
	}
	return &picker
}

//
type RandomPicker struct {
	subConn map[resolver.Address]balancer.SubConn
	lock    sync.Mutex
}

func (*RandomPicker) Pick(ctx context.Context, opts balancer.PickOptions) (conn balancer.SubConn, done func(balancer.DoneInfo), err error) {
	//panic("implement me")

}
