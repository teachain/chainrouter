package stub

type StubRPCServer struct {
	Impl Stub
}

func (this *StubRPCServer) Greet(args interface{}, resp *string) error {
	*resp = this.Impl.Greet()
	return nil
}

func (this *StubRPCServer) GetResources(args interface{}, resp *[]*ResourceInfo) error {
	*resp = this.Impl.GetResources()
	return nil
}
