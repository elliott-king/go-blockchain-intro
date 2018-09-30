package main

import (
    "fmt"
)

func main() {
    bc := NewBlockchain()

    bc.AddBlock("Send 1 BTC to send_bitcoin")
    bc.AddBlock("Send 2 more BTC to send_bitcoin")

    for _, block := range bc.blocks {
        fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Println()
    }
}