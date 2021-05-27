package stub

import "math/big"

type TransactionResponse struct {
	ErrorCode   int
	Message     string
	Hash        string
	ExtraHashes []string
	BlockNumber *big.Int
	Result      []string
}
