package main

import (
    "fmt"
    "time"
    "bytes"
    "encoding/gob"
)

// Taken from:
// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

type Block struct {
    Timestamp           int64
    Transactions        []*Transaction
    PrevBlockHash       []byte
    Hash                []byte
    Nonce               int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte)
*Block {
    block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

func (b *Block) Serialize() []byte {
    var result bytes.Buffer
    encoder := gob.NewEncoder(&result)

    if err:= encoder.Encode(b); err != nil {
        fmt.Println(err)
    }

    return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
    var block Block

    decoder := gob.NewDecoder(bytes.NewReader(d))
    if err := decoder.Decode(&block); err != nil {
        fmt.Println(err)
    }
    return &block
}
