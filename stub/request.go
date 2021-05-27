package stub

type Request struct {
	Type         int
	Path         string
	Data         []byte
	ResourceInfo ResourceInfo
}
