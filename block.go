package main

import (
    "bytes"
    "crypto/sha256"
    "strconv"
    "time"
)

// Taken from:
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

type Block struct {
    Timestamp           int64
    Data                []byte
    PrevBlockHash       []byte
    Hash                []byte
}

// Concatenate & calc SHA-256 on block fields. Oversimplification.
func (b *Block) SetHash() {
    timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
    headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
    hash := sha256.Sum256(headers)

    b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
    block.SetHash()
    return block
}

