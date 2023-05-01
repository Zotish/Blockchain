package function

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"main/str"
)

type Wallet struct {
	WalletName string
	PrivateKey string
	Publickey  string
	Address    string
	Balance    uint64
	Password   string
}

func GetCoinName(CN *str.Coin) string {
	return CN.CoinName
}
func GetSymbole(CN *str.Coin) string {
	return CN.Symbole
}
func GetInitialSupply(CN *str.Coin) uint64 {
	return CN.InitialSupply
}

func CreateWallet(Password string) *str.Wallet {
	privateKeySEc, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKeyHexFormate := hex.EncodeToString(privateKeySEc.D.Bytes())
	Pubkey := append(privateKeySEc.PublicKey.X.Bytes(), privateKeySEc.PublicKey.Y.Bytes()...)
	NewPublicKey := hex.EncodeToString(Pubkey)
	GenerateAddress := privateKeyHexFormate + NewPublicKey + Password
	Bytes := []byte(GenerateAddress)
	ShA := sha256.Sum256(Bytes)
	GenerateAddressToHex := hex.EncodeToString(ShA[:15])
	if len(GenerateAddressToHex) < 30 {
		GenerateAddressToHex = "0" + GenerateAddressToHex
	}

	GetWalletInfo := &str.Wallet{
		WalletName: "Wallet1",
		PrivateKey: privateKeyHexFormate,
		Publickey:  NewPublicKey,
		Address:    "0x" + GenerateAddressToHex,
		Balance:    0,
		Password:   Password,
	}
	return GetWalletInfo
}
func (privateKey *Wallet) GetPirvateKey(pass string) string {
	if privateKey.Password != pass {
		return "Incorrect Password"
	}
	return privateKey.PrivateKey
}
func (pubKey *Wallet) GetPublicKey() string {
	return pubKey.Publickey
}
func (address *Wallet) Receive() string {
	return address.Address
}
func (WalletNamep *Wallet) UpdateWallet(Name string) {
	WalletNamep.WalletName = Name
}
