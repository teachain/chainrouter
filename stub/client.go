package stub

import "net/rpc"

// Here is an implementation that talks over RPC
type StubRPC struct{ client *rpc.Client }

//本路由作为调用端，stub插件作为服务器端
func (this *StubRPC) Greet() string {
	var resp string
	err := this.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}
	return resp
}
