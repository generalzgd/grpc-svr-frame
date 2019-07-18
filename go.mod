module grpc-svr-frame

go 1.12

require (
	github.com/astaxie/beego v1.11.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9 // indirect
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/funny/link v0.0.0-20190321084249-bd07e4e9e63f // indirect
	github.com/funny/slab v0.0.0-20180511031532-b1fad5e5d478 // indirect
	github.com/generalzgd/securegotcp v0.0.0-20190423103017-2d895b9280f8 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/net v0.0.0-20190514140710-3ec191127204 // indirect
	golang.org/x/sys v0.0.0-20190516110030-61b9204099cb // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190516172635-bb713bdc0e52 // indirect
	google.golang.org/grpc v1.21.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f

replace golang.org/x/net => github.com/golang/net v0.0.0-20190514140710-3ec191127204

replace (
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace google.golang.org/appengine => github.com/golang/appengine v1.1.0

replace golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190517181255-950ef44c6e07

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190517183331-d88f79806bbd

replace google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190516172635-bb713bdc0e52

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190516110030-61b9204099cb
