package main

import (
    "time"
)

// Taken from:
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

type Block struct {
    Timestamp           int64
    Data                []byte
    PrevBlockHash       []byte
    Hash                []byte
    Nonce               int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

