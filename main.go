package main

import "FinUploader/helpers"

//go:generate abigen --abi ./abi/abi.json --pkg transactionLedger --type TransactionLedger --out ./abi/SwiftTransactionLedger.go

func main() {
	config := helpers.Config{}
	/*cfg*/ _, err := config.Setup()
	if err != nil {
		return
	}

	//cfg.AddTransaction()
	//cfg.GetTransaction()
	//cfg.GetTransaction()
}
