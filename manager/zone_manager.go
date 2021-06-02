package manager

import (
	"github.com/teachain/chainrouter/stub"
	"sync"
)

type ZoneManager struct {
	zones       map[string]*Zone
	stubManager StubManager
	readLock    sync.RWMutex
}

func (this *ZoneManager) SetZones(zones map[string]*Zone) {
	this.zones = zones
}

func (this *ZoneManager) GetChain(path *stub.Path) *Chain {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	return this.getChain(path)
}
func (this *ZoneManager) getChain(path *stub.Path) *Chain {
	zone := this.getZoneByPath(path)
	if zone != nil {
		return zone.GetChain(path)
	}
	return nil
}

//调用前先加锁
func (this *ZoneManager) getZoneByPath(path *stub.Path) *Zone {
	return this.getZone(path.GetZone())
}

//调用前先加锁
func (this *ZoneManager) getZone(name string) *Zone {
	return this.zones[name]
}
func (this *ZoneManager) GetResource(path *stub.Path, create bool) *Resource {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	zone := this.getZoneByPath(path)
	if zone == nil {
		return nil
	}
	chain := zone.GetChain(path)
	if chain == nil {
		return nil
	}
	resource := chain.GetResource(path)
	if resource != nil {
		return resource
	}
	if create {
		//todo
		//资源不存在就创建默认的
	}
	return nil

}

//获取所有的资源
func (this *ZoneManager) GetChainResources(path *stub.Path) map[string]*Resource {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	chain := this.getChain(path)
	if chain == nil {
		return nil
	}
	return chain.GetResources()
}

//获取所有的资源
func (this *ZoneManager) GetAllResources(ignoreRemote bool) map[string]*Resource {
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	result := make(map[string]*Resource)
	for k, _ := range this.zones {
		zone := this.zones[k]
		zoneName := toPureName(k)
		chains := zone.GetChains()
		for j, _ := range chains {
			chain := chains[j]
			stubName := toPureName(j)
			resources := chain.GetResources()
			for c, _ := range resources {
				resource := resources[c]
				if ignoreRemote && !resource.IsLocal() {
					continue
				}
				resourceName := toPureName(c)
				name := zoneName + "." + stubName + "." + resourceName
				result[name] = resource
			}
		}
	}
	return result
}
func (this *ZoneManager) getAllChainsInfo(ignoreRemote bool) map[string]*ChainInfo {
	result := make(map[string]*ChainInfo)
	this.readLock.RLock()
	defer this.readLock.RUnlock()
	for k, _ := range this.zones {
		zone := this.zones[k]
		chains := zone.GetChains()
		for j, _ := range chains {
			chain := chains[j]
			if ignoreRemote && !chain.IsLocal() {
				continue
			}
			chainName := toPureName(j)
			path := &stub.Path{}
			path.SetZone(k)
			path.SetChain(chainName)
			result[path.ToString()] = chain.GetChainInfo()
		}
	}
	return result
}
