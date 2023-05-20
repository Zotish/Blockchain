package str

type Block struct {
	BlockHeight  uint64
	TimeStamp    string
	PreviousHash string
	CurrentHash  string
	Transaction  *TransactionDetails
}
type TransactionDetails struct {
	Sender   string
	Receiver string
	Amount   float64
	Fees     float64
	Status   bool
}

var Chain []*Block

type Blockchain struct {
	Coins   *Coin
	Wallets *Wallet
	St      *StakeAndUnStake
}

type Coin struct {
	CoinName      string
	Symbole       string
	InitialSupply uint64
}

var User []*Wallet

type Wallet struct {
	WalletName string
	PrivateKey string
	Publickey  string
	Address    string
	Balance    float64
	Password   string
}

var ValidatorList []*StakeAndUnStake

type StakeAndUnStake struct {
	ValidatorAddress     string
	UserAddress          string
	StakingAmount        float64
	ValidatorPrivateKey  string
	ValidatorID          uint64
	BlockCreationTime    string
	DurationToCompleteTx string
	GasPrice             uint64
	GasLim               float64
	ChainID              uint64
}
