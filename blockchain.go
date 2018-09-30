package main

// In essences, blockchain is just a  db w/ ordered structure (each block linked to prev)
type Blockchain struct {
    blocks []*Block
}

func(bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks)-1]
    newBlock := NewBlock(data, prevBlock.Hash)
    bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}
