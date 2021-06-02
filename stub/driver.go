package stub

import "math/big"

type DriverCallback interface {
	onTransactionResponse(err error, transactionResponse TransactionResponse)
}

type GetBlockNumberCallback interface {
	OnResponse(err error, blockNumber *big.Int)
}

type GetBlockCallback interface {
	onResponse(err error, block Block)
}
type GetTransactionCallback interface {
	onResponse(err error, transaction Transaction)
}
type Driver interface {
	DecodeTransactionRequest(request Request) (bool, TransactionRequest)

	GetResources(connection Connection) []ResourceInfo

	AsyncCall(context TransactionContext, request TransactionRequest, byProxy bool, connection Connection, callback DriverCallback)

	AsyncSendTransaction(context TransactionContext, request TransactionRequest, byProxy bool, connection Connection, callback DriverCallback)

	AsyncGetBlockNumber(connection Connection, callback GetBlockNumberCallback)

	AsyncGetBlock(blockNumber *big.Int, onlyHeader bool, connection Connection, callback GetBlockCallback)
}
