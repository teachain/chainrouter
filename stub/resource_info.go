package stub

type ResourceInfo struct {
	name       string
	stubType   string
	properties map[string]interface{}
	checksum   string
}

func (this *ResourceInfo) SetName(name string) {
	this.name = name
}
func (this *ResourceInfo) GetName() string {
	return this.name
}
func (this *ResourceInfo) SetStubType(stubType string) {
	this.stubType = stubType
}
func (this *ResourceInfo) GetStubType() string {
	return this.stubType
}
