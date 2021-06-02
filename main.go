package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/teachain/chainrouter/stub"
	"github.com/teachain/chainrouter/manager"
	"log"
	"os"
	"os/exec"
)

func main() {
	app := "/Users/daminyang/Downloads/hello/check"
	driver, err := manager.NewStubDriver(app)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(driver.Greet())
	driver.Exited()
}

func test() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugins/simstub"),
		Logger:          logger,
	})
	defer client.Exited()

	//获取客户端
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	//请求对应的插件
	// Request the plugin
	raw, err := rpcClient.Dispense("stub")
	if err != nil {
		log.Fatal(err)
	}

	//
	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(stub.Stub)
	fmt.Println(greeter.Greet())

	{
		// We're a host! Start by launching the plugin process.
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: handshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command("./plugins/simstub2"),
			Logger:          logger,
		})
		defer client.Exited()

		//获取客户端
		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
		}
		//请求对应的插件
		// Request the plugin
		raw, err := rpcClient.Dispense("stub")
		if err != nil {
			log.Fatal(err)
		}
		//
		// We should have a Greeter now! This feels like a normal interface
		// implementation but is in fact over an RPC connection.
		greeter := raw.(stub.Stub)
		fmt.Println("from 2", greeter.Greet())
		//记得关闭
		rpcClient.Close()

	}
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"stub": &stub.StubPlugin{},
}
