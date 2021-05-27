package stub

import (
	"net/rpc"
	"github.com/hashicorp/go-plugin"
)

type Stub interface {
	//全部接口
	Greet() string
}

type StubPlugin struct {
	Impl Stub
}

func (this *StubPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &StubRPCServer{Impl: this.Impl}, nil
}

func (this *StubPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &StubRPC{client: c}, nil
}
