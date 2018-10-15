package main

import (
    "fmt"
    "bytes"
    "crypto/sha256"
    "encoding/gob"
    "log"
)

const subsidy = 10 // Arbitrary

// Inputs reference outputs of a previous transaction
type TXInput struct {
    Txid        []byte
    Vout        int
    ScriptSig   string
}

// Outputs are where coins are stored
type TXOutput struct {
    Value               int
    ScriptPubKey        string
}


type Transaction struct {
    ID          []byte
    Vin         []TXInput
    Vout        []TXOutput
}

// TX for mining, requires no prev outputs
func NewCoinbaseTX(to, data string) *Transaction {
    if data == "" {
        data = fmt.Sprintf("Reward to '%s'", to)
    }

    txin := TXInput{[]byte{}, -1, data}
    txout := TXOutput{subsidy, to}
    tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
    tx.SetId()

    return &tx
}

// will soon be tx.Hash()
func (t *Transaction) SetId() {
    var encoded bytes.Buffer
    var hash [32]byte

    enc := gob.NewEncoder(&encoded)
    err := enc.Encode(t)
    if err != nil {
        log.Panic(err)
    }

    hash = sha256.Sum256(encoded.Bytes())
    t.ID = hash[:]
}
