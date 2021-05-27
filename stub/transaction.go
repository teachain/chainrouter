package stub

type Transaction struct {
	TxBytes             []byte
	ReceiptBytes        []byte
	AccountIdentity     string
	Resource            string
	TransactionRequest  TransactionRequest
	TransactionResponse TransactionResponse
	TransactionByProxy  bool
}
