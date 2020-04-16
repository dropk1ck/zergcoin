package main

import (
    "bytes"
    "crypto/sha256"
    "fmt"
    "math"
    "math/big"
    "strconv"
    "time"
)

// Block object
type Block struct {
    Timestamp       int64       // time when block was created
    Data            []byte
    PrevBlockHash   []byte
    Target          int64
    Nonce           int64
    Hash            []byte
}

func (b *Block) GetHeader() []byte {
    return bytes.Join(
        [][]byte{
            b.Data,
            b.PrevBlockHash,
            []byte(strconv.FormatInt(b.Timestamp, 10)),
            []byte(strconv.FormatInt(b.Target, 10)),
            []byte(strconv.FormatInt(b.Nonce, 10)),
        },
        []byte{},
    )
}

func (b *Block) CalcPOW() bool {
    var hashInt big.Int
    var nonce int64
    var hash [32]byte

    fmt.Println("Mining block...")
    target := big.NewInt(1)
    target.Lsh(target, uint(256-b.Target))
    nonce = 0
    for nonce = 0; nonce < math.MaxInt64; nonce++ {
        b.Nonce = nonce
        data := b.GetHeader()
        hash = sha256.Sum256(data)
        hashInt.SetBytes(hash[:])
        if hashInt.Cmp(target) == -1 {
            b.Hash = hash[:]
            fmt.Printf("Found hash for new block: %x\n", b.Hash)
            return true
        }
    }
    fmt.Println("No hash found? Impossible....")
    b.Nonce = 0
    return false
}

func (b *Block) ValidateBlock() bool {
    var hashInt big.Int

    data := b.GetHeader()
    hash := sha256.Sum256(data)
    hashInt.SetBytes(hash[:])
    target := big.NewInt(1)
    target.Lsh(target, uint(256-b.Target))
    isValid := hashInt.Cmp(target) == -1
    return isValid
}

func NewBlock(data []byte, prevBlockHash []byte) *Block {
    block := new(Block)
    block.Timestamp = time.Now().Unix()
    block.Data = data
    block.PrevBlockHash = prevBlockHash
    block.Target = targetBits
    return block
}
