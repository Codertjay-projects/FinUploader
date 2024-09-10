package helpers

import (
	transactionLedger "FinUploader/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
	"strconv"
)

type Config struct {
	ETHClient                *ethclient.Client
	TransactionLedger        *transactionLedger.TransactionLedger
	Authentication           *bind.TransactOpts
	HTTPRPCURL               string
	TransactionLedgerAddress common.Address
}

func (cfg *Config) Setup() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg.SetupRPCURL()

	ethClient, err := cfg.SetupETHClient(cfg.HTTPRPCURL)
	if err != nil {
		return nil, err
	}
	cfg.ETHClient = ethClient

	cfg.SetupAuthentication()
	cfg.SetupTransactionLedgerAddress()

	return cfg, nil
}

func (cfg *Config) SetupRPCURL() *Config {
	cfg.HTTPRPCURL = "https://polygon-mainnet.alchemyapi.io/v2/ydQb_t8kmgv_Q8lBoXYxD9wRJOP0SYFA"
	return cfg
}

func (cfg *Config) SetupETHClient(rpcURL string) (*ethclient.Client, error) {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(client), nil
}

func (cfg *Config) SetupTransactionLedgerAddress() *Config {
	var err error
	cfg.TransactionLedgerAddress = common.HexToAddress(os.Getenv("ARBITRAGE_ADDRESS"))
	cfg.TransactionLedger, err = transactionLedger.NewTransactionLedger(cfg.TransactionLedgerAddress, cfg.ETHClient)
	if err != nil {
		log.Fatal("Error creating TransactionLedger contract:", err)
		return nil
	}
	return cfg
}

func (cfg *Config) SetupAuthentication() *Config {
	chainIDStr := os.Getenv("CHAIN_ID")
	if chainIDStr == "" {
		log.Fatalf("CHAIN_ID environment variable not set")
	}
	chainIDInt, err := strconv.Atoi(chainIDStr)
	if err != nil {
		log.Fatalf("Failed to convert CHAIN_ID to integer: %v", err)
	}
	chainID := big.NewInt(int64(chainIDInt))

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY")[2:])
	if err != nil {
		log.Fatalf("Failed to convert private key to ECDSA: %v", err)
	}
	cfg.Authentication, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Set gas price to 50 Gwei (50 * 10^9 Wei)
	cfg.Authentication.GasPrice = big.NewInt(500 * 1000000000)

	return cfg
}

func (cfg *Config) AddTransaction(_transactionReference string, _dateCurrencyAmount string, _orderingCustomer string, _beneficiary string) error {
	_, err := cfg.TransactionLedger.AddTransaction(cfg.Authentication, _transactionReference, _dateCurrencyAmount, _orderingCustomer, _beneficiary)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) GetTransaction(index int64) (string, string, string, string, error) {
	transactionReference, dateCurrencyAmount, orderingCustomer, beneficiary, err := cfg.TransactionLedger.GetTransaction(nil, big.NewInt(index))
	if err != nil {
		return "", "", "", "", err
	}
	return transactionReference, dateCurrencyAmount, orderingCustomer, beneficiary, nil
}

func (cfg *Config) RecordFinFile(name, cid string) error {
	_, err := cfg.TransactionLedger.RecordFinFile(cfg.Authentication, name, cid)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) GetFinFile(name, cid string) (string, error) {
	finFile, err := cfg.TransactionLedger.GetFinFile(nil)
	if err != nil {
		return "", err
	}
	return finFile, nil
}
