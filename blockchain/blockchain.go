package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func (b *block) GetDetail() string {
	return fmt.Sprintf("Data: %s\nHash: %s\nPrev Hash: %s\n", b.data, b.hash, b.prevHash)
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].hash
}

func (b *block) getHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *block {
	newBlock := block{
		data:     data,
		hash:     "",
		prevHash: getLastHash(),
	}

	newBlock.getHash()
	return &newBlock
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
