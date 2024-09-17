// main.go
package main

import (
	"usdc-transfers/internal/storage"
	"usdc-transfers/internal/services"
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
)

const (
	usdcAddress   = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	transferTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Infura URL from .env
	infuraURL := os.Getenv("INFURA_URL")
	if infuraURL == "" {
		log.Fatal("INFURA_URL not set in .env file")
	}

	// Check if block number is provided
	if len(os.Args) < 2 {
		log.Fatal("Please provide a block number as an argument")
	}

	// Parse block number from command-line argument
	blockNumber, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("Invalid block number: %v", err)
	}

	// Connect to Ethereum network
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	// Open SQLite database
	db, err := sql.Open("sqlite3", "usdc_transfers.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = storage.DropTable(db)
	if err != nil {
		log.Fatalf("Failed to drop database: %v", err)
	}

	// Create table if not exists
	err = storage.CreateTable(db)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Process the specified block
	services.ProcessBlock(client, db, blockNumber, usdcAddress, transferTopic)
}


