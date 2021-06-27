package main

import (
	"fmt"

	"github.com/toy-program/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println(block.GetDetail())
	}
}
