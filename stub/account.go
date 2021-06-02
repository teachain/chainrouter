package stub

type Account interface {
	GetName() string
	GetType() string
	GetIdentity() string
	GetKeyID() int
	IsDefault() bool
}
