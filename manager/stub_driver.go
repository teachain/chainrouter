package manager

import (
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/teachain/chainrouter/common"
	"github.com/teachain/chainrouter/stub"
	"github.com/teachain/chainrouter/utils"
	"os"
	"os/exec"
)

type StubDriver struct {
	stub.Stub
	client *plugin.Client
}

//must be called
func (this *StubDriver) Exited() {
	this.client.Exited()
}
func NewStubDriver(stubPath string) (*StubDriver, error) {
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
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	appPath := stubPath + common.FILE_SEPARATOR + common.STUB_APP

	appConfigPath := stubPath + common.FILE_SEPARATOR + common.STUB_CONFIG_FILE

	if !utils.FileExists(appPath) {
		return nil, errors.New(common.STUB_APP + " not exists in " + stubPath)
	}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(appPath, "--config", appConfigPath),
		Logger:          logger,
	})
	//获取客户端
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}
	//请求对应的插件
	// Request the plugin
	raw, err := rpcClient.Dispense("stub")
	if err != nil {
		return nil, err
	}
	//
	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	driver, ok := raw.(stub.Stub)
	if ok {
		stubDriver := &StubDriver{
			client: client,
			Stub:   driver,
		}
		return stubDriver, nil
	}
	return nil, errors.New("stub not implement stub.Stub")
}
