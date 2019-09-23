/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: balance.go
 * @time: 2019/9/11 13:47
 */
package grpc_consul

import (
	`errors`
	`fmt`
	`net`
	`strings`
	`sync`

	`github.com/astaxie/beego/logs`
	`github.com/hashicorp/consul/api`
	`github.com/hashicorp/consul/api/watch`
	`google.golang.org/grpc/resolver`
)

var (
	currentConsulAddr = "http://127.0.0.1:8500"
)

const (
	defaultPort = "0"
)

func init() {
	InitRegister(currentConsulAddr)
}

/*
* 默认为本地的consul agent代理，如有其它consul地址，请手动注册
 */
func InitRegister(consulAddr string) {
	currentConsulAddr = consulAddr
	resolver.Register(NewBuilder())
}

func NewBuilder() resolver.Builder {
	return &consulBuilder{
		data: map[string]resolver.Resolver{},
	}
}

type consulBuilder struct {
	lock   sync.Mutex
	data   map[string]resolver.Resolver
	client *api.Client
}

func parseTarget(target, defaultPort string) (host, port string, err error) {
	if len(target) < 1 {
		return "", "", errors.New("address emtpy")
	}
	if ip := net.ParseIP(target); ip != nil {
		return target, defaultPort, nil
	}
	if host, port, err := net.SplitHostPort(target); err == nil {
		if len(port) < 1 {
			return "", "", errors.New("port empty")
		}
		if len(host) < 1 {
			host = "localhost"
		}
		return host, port, nil
	}
	if host, port, err := net.SplitHostPort(target + ":" + defaultPort); err == nil {
		return host, port, nil
	}
	return "", "", fmt.Errorf("invalid target address %v, error info:%v", target, err)
}

func (p *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	logs.Debug("build:", target)

	host, _, err := parseTarget(target.Endpoint, defaultPort)
	if err != nil {
		return nil, err
	}
	host = strings.ToLower(host)

	p.lock.Lock()
	defer p.lock.Unlock()

	if p.client == nil {
		cfg := api.DefaultConfig()
		cfg.Address = currentConsulAddr
		cli, err := api.NewClient(cfg)
		if err != nil {
			return nil, err
		}
		p.client = cli
	}

	// if r, ok := p.data[host]; ok {
		// panic("resolve is already builded"+host)
		// return r, nil
	// }

	r := &consulResolver{
		address:              target.Endpoint, // Authority,
		name:                 host,            // 服务名
		cc:                   cc,
		disableServiceConfig: opts.DisableServiceConfig,
		client:               p.client,
	}
	// p.data[host] = r
	go r.watch()
	return r, nil
}

func (*consulBuilder) Scheme() string {
	return "consul"
}

// *************************************************************************************
type consulResolver struct {
	address              string
	method               string
	name                 string
	cc                   resolver.ClientConn
	disableServiceConfig bool
	client               *api.Client
	currentWatcher       *watch.Plan
}

func (p *consulResolver) watch() {
	params := map[string]interface{}{
		"type":    "service",
		"service": p.name,// + "_grpc",
		"tag":     "grpc",
	}

	plan, err := watch.Parse(params)
	if err != nil {
		return
	}
	p.currentWatcher = plan
	plan.Handler = func(idx uint64, val interface{}) {
		if val != nil {
			// logs.Info("consul resolver watch:", idx, val)
			if ents, ok := val.([]*api.ServiceEntry); ok {
				addrList := make([]resolver.Address, 0, len(ents))
				for _, it := range ents {
					addrList = append(addrList, resolver.Address{
						Addr:       fmt.Sprintf("%s:%d", it.Service.Address, it.Service.Port),
						Type:       resolver.Backend,
						ServerName: it.Service.Service,
					})
				}
				p.onServiceChange(addrList)
			}
		}
	}
	plan.Run(currentConsulAddr)
}

func (p *consulResolver) onServiceChange(list []resolver.Address) {
	logs.Info("service list: %v", list)
	p.cc.NewAddress(list)
	p.cc.NewServiceConfig(p.name)
}

func (p *consulResolver) ResolveNow(opt resolver.ResolveNowOption) {
	logs.Info("consulResolver::ResolveNow:", opt)
}

func (p *consulResolver) Close() {
	logs.Info("consulResolver::Close")
	p.cc.UpdateState(resolver.State{})
}

// 测试函数
// func testCallRpc() {
// 	address := "consul:///authsvr"
//
// 	opts := []grpc.DialOption{
// 		grpc.WithBlock(),
// 		grpc.WithBalancerName(roundrobin.Name),
// 		grpc.WithInsecure(),
// 	}
// 	ctx := context.Background()
// 	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
// 	conn, err := grpc.DialContext(ctx, address, opts...)
// 	if err != nil {
// 		return
// 	}
// 	defer conn.Close()
// 	client := ZQProto.NewAuthorizeClient(conn)
//
// 	req := &ZQProto.ImLoginRequest{
// 		Uid:      163,
// 		Token:    "123123.xfsdfs",
// 		Platform: 1,
// 	}
//
// 	// for {
// 		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
// 		rep, err := client.Login(ctx, req)
// 		if err != nil {
// 			logs.Error("rpc call error:%v", err)
// 		} else {
// 			logs.Info("rpc call ok:", rep.String())
// 		}
// 		// time.Sleep(time.Second * 2)
// 	// }
//
// }
