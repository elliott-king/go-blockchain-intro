package main

import (
    "fmt"
    "github.com/boltdb/bolt"
    "log"
)

const dbFile = "db/blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "Hello from sunny Louisville Kentucky"

// In essences, blockchain is just a  db w/ ordered structure (each block linked to prev)
type Blockchain struct {
    tip []byte
    db *bolt.DB
}

type BlockchainIterator struct {
    currentHash []byte
    db          *bolt.DB
}

func(bc *Blockchain) AddBlock(transactions []*Transaction) {
    var lastHash []byte
    err := bc.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(blocksBucket))
        lastHash = b.Get([]byte("1"))
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }

    newBlock := NewBlock(transactions, lastHash)

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

// Creates new blockchain and genesis block
func NewBlockchain(address string) *Blockchain {
    var tip []byte
    db, err := bolt.Open(dbFile, 0600, nil)
    if err != nil {
        log.Panic("Could not open db.\n", err)
    }
    fmt.Println("Got here first!")

    if err := db.Update(func(tx *bolt.Tx) error {
        fmt.Println("Got here!")
        cbtx := NewCoinbaseTX(address, genesisCoinbaseData) //TODO
        genesis := NewGenesisBlock(cbtx)

        b, _ := tx.CreateBucket([]byte(blocksBucket))
        _ = b.Put(genesis.Hash, genesis.Serialize())
        _ = b.Put([]byte("1"), genesis.Hash)

        tip = genesis.Hash
        return nil
    }); err != nil {
        //fmt.Println("Error
        log.Panic(err)
    }

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
