package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    //"github.com/davecgh/go-spew/spew"
)

const targetBits = 24

func main() {
    fmt.Print("*** ZERGCOIN ***\n\n")
    blockchain := NewBlockchain()

    reader := bufio.NewReader(os.Stdin)
    // loop and add blocks from stdin
    for {
        fmt.Print("Enter data to generate new block: ")
        data, _ := reader.ReadString('\n')
        data = strings.TrimSuffix(data, "\n")
        lastBlock := blockchain.GetLastBlock()
        newBlock := NewBlock([]byte(data), lastBlock.Hash)
        newBlock.CalcPOW()
        if blockchain.AddBlock(newBlock) {
            fmt.Println("Block successfully added")
            fmt.Println("Current blockchain:\n")
            blockchain.PrintBlockchain()
        }
    }
    //spew.Dump(blockchain)
}
