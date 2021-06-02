package service

import (
	"errors"
	"fmt"
	"github.com/teachain/chainrouter/manager"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/configor"

	"github.com/teachain/chainrouter/common"
	"github.com/teachain/chainrouter/config"
	"github.com/teachain/chainrouter/utils"


)

//1、加载配置文件
func LoadAppConfig(filename string) (*config.AppConfig, error) {
	var config config.AppConfig
	err := configor.Load(&config, filename)
	if err != nil {
		return nil, err
	}
	return &config, err
}
func LoadStubConfig(filename string) (*config.StubConfig, error) {
	var config config.StubConfig
	err := configor.Load(&config, filename)
	if err != nil {
		return nil, err
	}
	return &config, err
}
func NewZoneMap(config *config.AppConfig) (map[string]*manager.Zone, error) {
	if config.Common == nil {
		return nil, errors.New("common must be set")
	}
	if config.Common.Zone == "" {
		return nil, errors.New("common manager must be set")
	}
	network := config.Common.Zone
	visible := config.Common.Visible
	stubsPath := config.Chains.Path

	fmt.Println(network)
	fmt.Println(visible)
	fmt.Println(stubsPath)

	return nil, nil
}

func GetChains(zone string, chainsDir map[string]string) map[string]*manager.Chain {
	for chainName, _ := range chainsDir {
		//链对应的目录路径
		stubPath := chainsDir[chainName]
		//配置文件stub.yaml的路径
		stubConfigFile := stubPath + common.FILE_SEPARATOR + common.STUB_CONFIG_FILE
		config, err := LoadStubConfig(stubConfigFile)
		if err != nil {
			fmt.Println(err)
		}
		typeStr := config.Common.Type
		name := config.Common.Name
		//目录名称和这里的配置要一致
		if k != name {
			fmt.Println("error stub name")
			continue
		}
		chainInfo:=manager.NewChainInfo()
		chainInfo.SetName(chainName)
		//todo perperties
		chainInfo.SetStubType(typeStr)




	}
	return nil
}

func GetStubsDir(stubsPath string) (map[string]string, error) {
	result := make(map[string]string)
	f, err := os.Stat(stubsPath)
	if err != nil {
		return nil, err
	}
	if !f.IsDir() {
		return nil, errors.New(stubsPath + " is not a valid directory")
	}
	current := stubsPath
	if string(stubsPath[len(stubsPath)-1]) == "/" {
		current = stubsPath[:len(stubsPath)-1]
	}
	stubsDirs, err := ioutil.ReadDir(stubsPath)
	if err != nil {
		return nil, err
	}
	for _, stub := range stubsDirs {
		if !stub.IsDir() {
			continue
		}
		if strings.HasPrefix(stub.Name(), ".") {
			continue
		}
		stubPath := current + string(filepath.Separator) + stub.Name()
		//每个目录下必须存在这样一个文件stub.yaml
		filename := stubPath + string(filepath.Separator) + common.STUB_CONFIG_FILE
		if !utils.FileExists(filename) {
			return nil, err
		}
		result[stub.Name()] = stubPath

	}
	return result, nil
}

func Start() {
	//zoneManager := &manager.ZoneManager{}
	//config, err := LoadAppConfig("")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//zones := NewZoneMap(config)
	//zoneManager.SetZones(zones)
}
