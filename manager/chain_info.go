package manager

import (
	"github.com/teachain/chainrouter/stub"
)

//链信息
type ChainInfo struct {
	zone       string
	name       string
	stubType   string
	resources  []*stub.ResourceInfo
	properties map[string]string
	checksum   string
}

func NewChainInfo()*ChainInfo{
	return &ChainInfo{}
}

func (this *ChainInfo) SetProperties(properties map[string]string) {
	this.properties = properties
}
func (this *ChainInfo) GetProperties() map[string]string {
	return this.properties
}
func (this *ChainInfo) SetResources(resources []*stub.ResourceInfo) {
	this.resources = resources
}
func (this *ChainInfo) GetResources() []*stub.ResourceInfo {
	return this.resources
}

func (this *ChainInfo) GetName() string {
	return this.name
}
func (this *ChainInfo) SetName(name string) {
	this.name = name
}
func (this *ChainInfo) SetStubType(stubType string) {
	this.stubType = stubType
}
func (this *ChainInfo) GetStubType() string {
	return this.stubType
}
func (this *ChainInfo) SetZone(zone string) {
	this.zone = zone
}
func (this *ChainInfo) GetZone() string {
	return this.zone
}
func (this *ChainInfo) SetChecksum(checksum string) {
	this.checksum = checksum
}
func (this *ChainInfo) GetChecksum() string {
	return this.checksum
}
