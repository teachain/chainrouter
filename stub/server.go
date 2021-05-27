package stub

type StubRPCServer struct {
	Impl Stub
}
//要实现
func (this *StubRPCServer) Greet(args interface{}, resp *string) error {
	*resp = this.Impl.Greet()
	return nil
}
