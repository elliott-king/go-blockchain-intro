package main

import (
    "fmt"
    "github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// In essences, blockchain is just a  db w/ ordered structure (each block linked to prev)
type Blockchain struct {
    tip []byte
    db *bolt.DB
}

type BlockchainIterator struct {
    currentHash []byte
    db          *bolt.DB
}

func(bc *Blockchain) AddBlock(data string) {
    var lastHash []byte
    err := bc.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(blocksBucket))
        lastHash = b.Get([]byte("1"))
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }

    newBlock := NewBlock(data, lastHash)
    
    if err = bc.db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(blocksBucket))
        _ = b.Put(newBlock.Hash, newBlock.Serialize())
        _ = b.Put([]byte("1"), newBlock.Hash)
        bc.tip = newBlock.Hash
        return nil
    }); err != nil {
        fmt.Println(err)
    }
}

func NewGenesisBlock(coinbase *(Transaction) *Block {
    return NewBlock([]*Transaction{coinbase}, []byte{})
}

func NewBlockchain() *Blockchain {
    var tip []byte
    db, _ := bolt.Open(dbFile, 0600, nil)

    _ = db.Update(func(tx *bolt.Tx) error {
        cbtx := NewCoinbaseTX(address, genesisCoinbaseData) //TODO
        genesis := NewGenesisBlock(cbtx)

        b := tx.Bucket([]byte(blocksBucket))

        if b == nil {
            fmt.Println("No existing blockchain found. Creating a new one...")
            genesis := NewGenesisBlock()
            b, err := tx.CreateBucket([]byte(blocksBucket))
            if err != nil {
                fmt.Println(err)
            }
            _ = b.Put(genesis.Hash, genesis.Serialize())
            _ = b.Put([]byte("1"), genesis.Hash)
            tip = genesis.Hash
        } else {
            tip = b.Get([]byte("1"))
        }
        return nil
    })

    bc := Blockchain{tip, db}
    return &bc
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
    bci := BlockchainIterator{bc.tip, bc.db}
    return &bci
}

func (i *BlockchainIterator) Next() *Block {
    var block *Block

    _ = i.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(blocksBucket))
        encodedBlock := b.Get(i.currentHash)
        block = DeserializeBlock(encodedBlock)

        return nil
    })

    i.currentHash = block.PrevBlockHash
    return block
}
