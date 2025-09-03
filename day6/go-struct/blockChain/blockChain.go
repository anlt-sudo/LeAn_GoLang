package blockChain

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	Data     string
	PrevHash string
	Hash     string
}

type BlockChain struct {
	Blocks []*Block
}

func calHash(data string, prevHash string) string {
	info := joinDataPrevHash(data, prevHash)
	hashInfo := []byte(info)
	hash := sha256.Sum256(hashInfo)
	return fmt.Sprintf("%x", hash[:])
}

func joinDataPrevHash(data string, prevHash string) string {
	info := []string{data, prevHash}
	return strings.Join(info, "")
}

func NewBlock(data string, prevHash string) *Block {
	block := &Block{
		Data:     data,
		PrevHash: prevHash,
	}
	block.Hash = calHash(block.Data, block.PrevHash)
	return block
}


func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash := prevBlock.Hash
	newBlock := NewBlock(data, prevHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}


func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}


func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

func (bc *BlockChain) GetBlocks() []*Block {
	return bc.Blocks
}

