package manager

import (
	"github.com/teachain/chainrouter/stub"
	"sync"
)

type Chain struct {
	zoneName   string
	name       string
	stubType   string
	checkSum   string
	resources  map[string]*Resource
	properties map[string]string
	readLock   sync.RWMutex
	local      bool
	driver     *StubDriver
}
func (this *Chain) SetStubDriver(driver *StubDriver) {
	this.driver = driver
}

func (this *Chain) GetResource(path *stub.Path) *Resource {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	return this.resources[path.GetResource()]
}
func (this *Chain) GetResources() map[string]*Resource {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	return this.resources
}

//是否是本地链（就是本路由安装的链）
func (this *Chain) IsLocal() bool {
	return this.local
}

func (this *Chain) GetChainInfo() *ChainInfo {
	chainInfo := &ChainInfo{}
	chainInfo.SetZone(this.zoneName)
	chainInfo.SetName(this.name)
	chainInfo.SetStubType(this.stubType)
	chainInfo.SetProperties(this.properties)
	chainInfo.SetChecksum(this.checkSum)
	resourceInfos := this.getAllResourcesInfo(true)
	chainInfo.SetResources(resourceInfos)
	return chainInfo
}
func (this *Chain) getAllResourcesInfo(ignoreRemote bool) []*stub.ResourceInfo {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	result := make([]*stub.ResourceInfo, 0)
	for k, _ := range this.resources {
		resource := this.resources[k]
		if ignoreRemote && resource.IsLocal() {
			continue
		}
		result = append(result, resource.GetResourceInfo())
	}
	return result
}

//getAllResourcesInfo
