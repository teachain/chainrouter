package stub

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)


//在这里定义接口
type Stub interface {

	Greet() string

	GetResources()[]*ResourceInfo
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
