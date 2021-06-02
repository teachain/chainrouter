package stub

type StubFactory interface {
	NewDriver() Driver

	NewConnection(path string) Connection

	NewAccount(properties map[string]interface{}) Account

	GenerateAccount(path string, args []string)

	GenerateConnection(path string, args []string)
}
