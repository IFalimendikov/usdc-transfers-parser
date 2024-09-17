## 💰USDC Transfers Parser💰

### Overview

This is a Go-based application that tracks USDC transfers on the Ethereum blockchain. It fetches transfer events for a specific block and stores them in a SQLite database.

### Features

- Fetches USDC transfer events for a specified block number.
- Stores transfer details (block number, sender, recipient, value) in a SQLite database.
- Uses Infura for connecting to the Ethereum network.

### Prerequisites

- Go
- SQLite
- Infura account (for Ethereum network access)

### Installation 

1. Ensure you have Go installed on your system.
2. Clone this repository or download the source code, import packages.
3. Set up your Infura Key in the `.env.example`.
4. Rename `example.env` to `.env`.

### Example

```bash
# Run the program
$ go run cmd/main.go 20770696

# Output
Processed transfer: Block 20770696, From 0x8c1D39c2E24f2E5dF126851DF3E891a59221967C, To 0x931250786dFd106B1E63C7Fd8f0d854876a45200, Value 13272290000
Processed transfer: Block 20770696, From 0x36D85C5d10b858B9331E45cE9957D2a8257e0F7D, To 0xff8Ba4D1fC3762f6154cc942CCF30049A2A0cEC6, Value 1221000000
Processed transfer: Block 20770696, From 0xA9D1e08C7793af67e9d92fe308d5697FB81d3E43, To 0x1dD2a22dB3E0c8Dce032dDcc5983f927c46c03e0, Value 72024886
```



