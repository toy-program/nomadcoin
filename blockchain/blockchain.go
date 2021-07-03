package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

func (b *Block) GetDetail() string {
	return fmt.Sprintf("Data: %s\nHash: %s\nPrev Hash: %s\n", b.Data, b.Hash, b.PrevHash)
}

type Blockchain struct {
	Blocks []*Block
}

var b *Blockchain
var once sync.Once

func GetBlockchain() *Blockchain {
	if b == nil {
		once.Do(func() {
			b = &Blockchain{}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}

func (b *Blockchain) AddBlock(data string) {
	b.Blocks = append(b.Blocks, createBlock(data))
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().Blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().Blocks[totalBlocks-1].Hash
}

func (b *Block) getHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *Block {
	newBlock := Block{
		Data:     data,
		Hash:     "",
		PrevHash: getLastHash(),
	}

	newBlock.getHash()
	return &newBlock
}

func (b *Blockchain) AllBlocks() []*Block {
	return b.Blocks
}
