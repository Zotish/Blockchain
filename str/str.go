package str

type Block struct {
	Index        uint64
	TimeStamp    string
	PreviousHash string
	CurrentHash  string
	Nonce        uint64
	Transaction  *TransactionDetails
}
type TransactionDetails struct {
}
type Blockchain struct {
	Blockchain []Blockchain
}
type Coin struct {
	CoinName      string
	Symbole       string
	InitialSupply uint64
}

var User []Wallet

type Wallet struct {
	WalletName string
	PrivateKey string
	Publickey  string
	Address    string
	Balance    uint64
	Password   string
}
