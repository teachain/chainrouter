package manager

import "github.com/teachain/chainrouter/stub"

type Resource struct {
	local        bool
	resourceInfo *stub.ResourceInfo
}

func (this *Resource) IsLocal() bool {
	return this.local
}
func (this *Resource) GetResourceInfo() *stub.ResourceInfo {
	return this.resourceInfo
}
func (this *Resource) SetResourceInfo(resourceInfo *stub.ResourceInfo) {
	this.resourceInfo = resourceInfo
}
