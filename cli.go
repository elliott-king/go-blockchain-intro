package main

import (
    "flag"
    "os"
    "fmt"
    "strconv"
)

type CLI struct {
    bc *Blockchain
}

func (cli *CLI) printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  createblockchain -address ADDRESS - create blockchain and send genesis block reward to ADDRESS")
    fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
    if len(os.Args) < 2 {
        fmt.Println("Not enough args.")
        cli.printUsage()
        os.Exit(1)
    }
}

func (cli *CLI) Run() {
    cli.validateArgs()

    createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
    createBlockchainAddress := createBlockchainCmd.String("address", "", "Address for genesis block")

    switch os.Args[1] {
        case "createblockchain":
            _ = createBlockchainCmd.Parse(os.Args[2:])
        case "printchain":
            _ = printChainCmd.Parse(os.Args[2:])
        default:
            fmt.Println("No choice specified.")
            cli.printUsage()
            os.Exit(1)
    }

    if createBlockchainCmd.Parsed() {
        if *createBlockchainAddress == "" {
            fmt.Println("No address specified.")
            createBlockchainCmd.Usage()
            os.Exit(1)
        }
        cli.createBlockchain(*createBlockchainAddress)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }
}

func (cli *CLI) createBlockchain(address string) {
    bc := NewBlockchain(address)
    defer bc.db.Close()
    fmt.Println("Cli has created a blockchain")
}

//func (cli *CLI) addBlock(address string) {
//    cli.bc.AddBlock(address)
//    fmt.Println("Success")
//}

func(cli *CLI) printChain() {
    bc := NewBlockchain("")
    defer bc.db.Close()
    bci := cli.bc.Iterator()

    for {
        block := bci.Next()

        fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
        // fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)

        pow := NewProofOfWork(block)
        fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

        if len(block.PrevBlockHash) == 0 {
            break
        }
    }
}
