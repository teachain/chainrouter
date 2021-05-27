package stub

type Block struct {
	RawBytes           []byte
	BlockHeader        BlockHeader
	TransactionsHashes []string
}
