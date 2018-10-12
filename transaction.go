// Describes transaction upon the blockchain


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
