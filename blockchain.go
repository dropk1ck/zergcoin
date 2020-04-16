package main

import (
    "fmt"
    "reflect"
    //"github.com/davecgh/go-spew/spew"
)

// Blockchain object
type Blockchain struct {
    blocks []*Block
}

func CreateGenesisBlock() *Block {
    genesisBlock := NewBlock([]byte("Genesis Block"), []byte{})
    genesisBlock.CalcPOW()
    return genesisBlock
}

func (bc *Blockchain) AddBlock(newBlock *Block) bool {
    lastBlock := bc.GetLastBlock()
    if newBlock.ValidateBlock() {
        if reflect.DeepEqual(lastBlock.Hash, newBlock.PrevBlockHash) {
            bc.blocks = append(bc.blocks, newBlock)
            return true
        } else {
            fmt.Println("BAD BLOCK: PrevBlockHash was incorrect")
        }
    } else {
        fmt.Println("BAD BLOCK: block did not validate")
    }
    return false
}

func (bc *Blockchain) PrintBlockchain() {
    for idx, block := range bc.blocks {
        fmt.Printf("Block %d\n", idx)
        fmt.Printf("Data: %v\n", string(block.Data))
        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Printf("Difficulty: %d\n\n", block.Target)
    }
}

// this function verifies the blockchain by:
//  - verifying that each block is valid
//  - checking the next block to ensure it includes the previous block's hash
func (bc *Blockchain) Verify() bool {
    return false
}

func (bc *Blockchain) GetLastBlock() *Block {
    return bc.blocks[len(bc.blocks)-1]
}

func NewBlockchain() *Blockchain {
    fmt.Println("Creating the genesis block...")
    genesisBlock := CreateGenesisBlock()

    // the genesis block gets accepted by default 
    return &Blockchain{[]*Block{genesisBlock}}
}
