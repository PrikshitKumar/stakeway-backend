package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Struct for deposit data from JSON
type DepositData struct {
	PubKey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Amount                int64  `json:"amount"`
	Signature             string `json:"signature"`
	DepositMessageRoot    string `json:"deposit_message_root"`
	DepositDataRoot       string `json:"deposit_data_root"`
	ForkVersion           string `json:"fork_version"`
	NetworkName           string `json:"network_name"`
}

var (
	holeskyRPC      = "https://ethereum-holesky.publicnode.com"                         // Public RPC for Holesky testnet
	depositContract = common.HexToAddress("0x4242424242424242424242424242424242424242") // Staking contract address
	depositDataFile = "deposit_data.json"                                               // Provided deposit data file
	gasLimit        = uint64(500000)                                                    // Adjusted for staking transaction
	depositAmount   = 32                                                                // Staking Amount
)

// Load deposit data from JSON
func loadDepositData(filename string) (*DepositData, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var depositData []DepositData
	if err := json.Unmarshal(file, &depositData); err != nil {
		return nil, err
	}

	return &depositData[0], nil
}

func sendDepositTransaction() {
	// Load deposit data
	depositData, err := loadDepositData(depositDataFile)
	if err != nil {
		log.Fatalf("Failed to load deposit data: %v", err)
	}

	// Connect to Holesky testnet RPC
	client, err := rpc.Dial(holeskyRPC)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	defer client.Close()

	ethClient := ethclient.NewClient(client)

	// Load private key (Ensure you securely store this, never hardcode in production!)
	privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	senderAddress := crypto.PubkeyToAddress(*publicKey)
	log.Println("Sending from address:", senderAddress.Hex())

	// Get nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	// Construct transaction
	// Convert ETH amount to wei
	amount := new(big.Int).Mul(big.NewInt(int64(depositAmount)), big.NewInt(1e18))

	var depositDataRootArray [32]byte
	copy(depositDataRootArray[:], common.Hex2BytesFixed(depositData.DepositDataRoot, 32))

	// Encode transaction data
	parsedABI, _ := abi.JSON(strings.NewReader(`[{"name": "deposit", "type": "function", "inputs": [{"name": "pubkey", "type": "bytes"},{"name": "withdrawal_credentials", "type": "bytes"},{"name": "signature", "type": "bytes"},{"name": "deposit_data_root", "type": "bytes32"}]}]`))
	txData, err := parsedABI.Pack("deposit",
		common.Hex2Bytes(depositData.PubKey),
		common.Hex2Bytes(depositData.WithdrawalCredentials),
		common.Hex2Bytes(depositData.Signature),
		depositDataRootArray,
	)
	if err != nil {
		log.Fatalf("Failed to encode transaction data: %v", err)
	}

	// Create transaction
	tx := types.NewTransaction(nonce, depositContract, amount, gasLimit, gasPrice, txData)

	// Sign the transaction
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	log.Printf("Transaction sent! Tx Hash: %s", signedTx.Hash().Hex())

	// Wait for transaction receipt
	revertReason, success := waitForTransaction(ethClient, signedTx.Hash(), senderAddress)
	if !success {
		log.Println("Transaction failed. Revert reason:", revertReason)
	} else {
		log.Println("Transaction successful.")
	}
}

func waitForTransaction(client *ethclient.Client, txHash common.Hash, senderAddress common.Address) (string, bool) {
	ctx := context.Background()
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			if receipt.Status == 1 {
				return "", true
			}
			return fetchRevertReason(client, txHash, senderAddress), false
		}
		time.Sleep(3 * time.Second)
	}
}

func fetchRevertReason(client *ethclient.Client, txHash common.Hash, senderAddress common.Address) string {
	ctx := context.Background()
	tx, _, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		return "Failed to fetch transaction."
	}

	msg := ethereum.CallMsg{
		From:     senderAddress,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	revertData, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return fmt.Sprintf("Failed to fetch revert reason: %v", err)
	}

	if len(revertData) > 0 {
		return fmt.Sprintf("Revert reason: %s", string(revertData))
	}
	return "Transaction failed without a revert reason."
}

func main() {
	sendDepositTransaction()
}
