package stub

// Callback for asyncSend
type ConnectionCallback interface {
	onResponse(response Response)
}
type ConnectionEventHandler interface {
	onResourcesChange(resourceInfos []ResourceInfo)
}
type Connection interface {
	AsyncSend(request Request, callback ConnectionCallback)
	SetConnectionEventHandler(eventHandler ConnectionEventHandler)
	GetProperties() map[string]string
}
