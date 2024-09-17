package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
)

func ProcessBlock(client *ethclient.Client, db *sql.DB, blockNumber uint64, address, topic string) {
	// Get logs for the ERC20 contract in the specified block
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: []common.Address{common.HexToAddress(address)},
		Topics:    [][]common.Hash{{common.HexToHash(topic)}},
	})
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	stmt, err := db.Prepare(`
		INSERT INTO usdc_transfers (block_number, sender, recipient, value)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		log.Printf("Failed to prepare insert statement: %v", err)
		return
	}
	defer stmt.Close()

	// Process each log
	for _, vLog := range logs {
		sender := common.HexToAddress(vLog.Topics[1].Hex()).Hex()
		recipient := common.HexToAddress(vLog.Topics[2].Hex()).Hex()
		value := new(big.Int).SetBytes(vLog.Data).String()

		// Execute the prepared statement
		_, err := stmt.Exec(blockNumber, sender, recipient, value)
		if err != nil {
			log.Printf("Failed to insert transfer: %v", err)
			continue
		}

		fmt.Printf("Processed transfer: Block %d, From %s, To %s, Value %s\n", blockNumber, sender, recipient, value)
	}
}
