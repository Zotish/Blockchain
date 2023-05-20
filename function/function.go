package function

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"main/str"
	"math"
	"strconv"
	"time"
)

func CalculateHash(Block *str.Block) string {
	AddValue := Block.PreviousHash + Block.TimeStamp + Block.Transaction.Sender + Block.Transaction.Receiver + strconv.FormatFloat(Block.Transaction.Amount, 'f', -1, 64) + strconv.FormatFloat(Block.Transaction.Fees, 'f', -1, 64) + strconv.FormatBool(Block.Transaction.Status) + strconv.FormatUint(Block.BlockHeight, 32)
	var strBt [32]byte
	copy(strBt[:], []byte(AddValue))
	data2 := [][]byte{strBt[:]}
	Concates := bytes.Join(data2, []byte{})
	ConVToSHA := sha256.New()
	ConVToSHA.Write([]byte(Concates))
	var hashed [32]byte
	copy(hashed[:], ConVToSHA.Sum(nil))
	return hex.EncodeToString(hashed[:])
}
func CreateGenblock() *str.Block {
	GenBlock := &str.Block{
		BlockHeight:  0,
		PreviousHash: "",
		TimeStamp:    time.Now().Format("2006 01 01 12:00 "),
		Transaction: &str.TransactionDetails{
			Sender:   "",
			Receiver: "",
			Amount:   0,
			Fees:     0,
			Status:   true,
		},
	}
	GenBlock.CurrentHash = CalculateHash(GenBlock)
	str.Chain = append(str.Chain, GenBlock)
	return GenBlock
}
func CreateNewBlock(prevBlock *str.Block, tx *str.TransactionDetails) *str.Block {
	newBlock := &str.Block{
		BlockHeight:  prevBlock.BlockHeight + 1,
		PreviousHash: prevBlock.CurrentHash,
		Transaction: &str.TransactionDetails{
			Sender:   tx.Sender,
			Receiver: tx.Receiver,
			Amount:   tx.Amount,
			Fees:     tx.Fees,
			Status:   tx.Status,
		},
	}
	if !IsValid(newBlock, prevBlock) {
		newBlock.TimeStamp = time.Now().Format("2006 01 01 12:00")
		newBlock.CurrentHash = CalculateHash(newBlock)
	}
	CalculateBlockReward(int64(newBlock.BlockHeight))
	return newBlock
}
func IsValid(block *str.Block, prevBlock *str.Block) bool {
	if prevBlock.BlockHeight+1 != block.BlockHeight {
		return false
	}
	if prevBlock.CurrentHash != block.PreviousHash {
		return false
	}
	if prevBlock.TimeStamp != block.TimeStamp {
		return false
	}
	if CalculateHash(block) != block.CurrentHash {
		return false
	}
	return true
}
func CreateValidator(validatorAddress string, stakingAmount float64, validatorPrivateKey string, validatorID uint64, blockTime string, gasPrice uint64, gasLim float64, chainID uint64) error {
	// validate the input parameters
	if validatorAddress == "" {
		return errors.New("validator address cannot be empty")
	}

	if stakingAmount >= 100000 {
		return errors.New("staking amount must be greater than 100000")
	}
	if validatorPrivateKey == "" {
		return errors.New("validator private key cannot be empty")
	}
	if validatorID == 0 {
		return errors.New("validator ID cannot be 0")
	}
	if blockTime == "" {
		return errors.New("block time cannot be empty")
	}
	if gasPrice == 0 {
		return errors.New("gas price cannot be 0")
	}
	if gasLim <= 0 {
		return errors.New("gas limit must be greater than 0")
	}
	if chainID == 0 {
		return errors.New("chain ID cannot be 0")
	}

	// create a new stake and unstake object
	stakeAndUnStake := &str.StakeAndUnStake{
		ValidatorAddress:     validatorAddress,
		StakingAmount:        stakingAmount,
		ValidatorPrivateKey:  validatorPrivateKey,
		ValidatorID:          validatorID,
		DurationToCompleteTx: blockTime,
		GasPrice:             gasPrice,
		GasLim:               gasLim,
		ChainID:              chainID,
	}

	// add the stake and unstake object to the list of validators
	str.ValidatorList = append(str.ValidatorList, stakeAndUnStake)
	if len(str.ValidatorList) > 200 {
		fmt.Println("Noval")
	}

	return nil
}
func GetValidator() []*str.StakeAndUnStake {
	return str.ValidatorList
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
	str.User = append(str.User, GetWalletInfo)
	return GetWalletInfo
}
func CalculateBlockReward(blockHeight int64) float64 {
	var reward float64

	// Define the initial reward
	initialReward := 5.0

	// Halve the reward every 100,000 blocks
	halvingInterval := int64(1000000)
	halvings := blockHeight / halvingInterval

	// Calculate the reward based on the number of halvings
	reward = initialReward / math.Pow(2, float64(halvings))

	// Return the reward
	return reward
}
func GetValidatorID(ID uint64) uint64 {
	var id str.StakeAndUnStake
	for range str.ValidatorList {
		if id.ValidatorID != ID {
			fmt.Println("Id not found")
		}
	}
	return id.ValidatorID
}
func Stake(amount float64) bool {
	var UAddress str.Wallet
	var VaAddress str.StakeAndUnStake
	if GetBalanceByAddress(UAddress.Address) > amount {
		Getbal := GetBalanceByAddress(UAddress.Address)
		Getbal -= amount
		Deductbal := GetBalanceByAddress(VaAddress.ValidatorAddress)
		Deductbal += amount
		return true

	}
	return false
}
func GetBalanceByAddress(address string) float64 {
	var Address str.Wallet
	return Address.Balance
}

func IsUserStakes(Address string) bool {
	var StakeRecords []str.StakeAndUnStake
	for _, stake := range StakeRecords {
		if stake.UserAddress == Address {
			return true
		}
	}

	return false
}

// this function is still incomplete. i will do it as soon as possible when i will complete the react and implement these
/*func Unstake(amount float64) bool {
	var UnStake str.StakeAndUnStake
	var UserAdress str.Wallet
	if IsUserStakes(UnStake.ValidatorAddress) {
		Unbal := GetBalanceByAddress(UnStake.ValidatorAddress)
		Unbal -= amount
		userBal := GetBalanceByAddress(UserAddress.Address)
		unstakeTimestamp := time.Now().Add(21 * 24 * time.Hour)
		if uint64(unstakeTimestamp.Unix()) == 0 {
			userBal += amount
		} else {
			errors.New("")
		}

	}
	return true

}
*/
