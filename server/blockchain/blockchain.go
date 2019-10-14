package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block ...
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain ...
type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// NewBlock @parms: data string, prevBlockHash string @return: *Block
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.setHash()

	return block
}

// AddBlock @parms: data string @return: *Block
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// NewBlockchain @return: *Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock @return: *Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}
