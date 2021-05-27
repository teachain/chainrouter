package stub

import "math/big"

type BlockHeader struct {
	number          *big.Int
	PrevHash        string
	Hash            string
	StateRoot       string
	TransactionRoot string
	ReceiptRoot     string
}
