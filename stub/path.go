package stub

type Path struct {
	zone     string //业务领域
	chain    string //链的名称
	resource string //合约的名称
}

func (this *Path) ToURI() string {
	if this.resource != "" {
		return this.zone + "/" + this.chain + "/" + this.resource
	}
	return this.zone + "/" + this.chain
}
func (this *Path) GetChain() string {
	return this.chain
}
func (this *Path) GetZone() string {
	return this.zone
}
func (this *Path) GetResource() string {
	return this.resource
}
func (this *Path) ToString() string {
	if this.resource != "" {
		return this.zone + "." + this.chain + "." + this.resource
	}
	return this.zone + "." + this.chain
}

func (this *Path) SetZone(zone string) {
	this.zone = zone
}
func (this *Path) SetChain(chain string) {
	this.chain = chain
}

func (this *Path) SetResource(resource string) {
	this.resource = resource
}
