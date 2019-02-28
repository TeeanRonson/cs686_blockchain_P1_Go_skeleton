package main

import (
    "fmt"
    "github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p2"
)

func main() {

    var genesis p2.Block
    var block1 p2.Block
    bc := p2.NewBlockChain()

    genesis.CreateGenesisBlock()
    bc.Insert(genesis)
    fmt.Println(bc)
    genesis.Mpt.Insert("Rong", "Liew")
    block1.NewBlock(genesis.Header.Height + 1, genesis.Header.ParentHash, genesis.Mpt)
    fmt.Println("Block 1:", block1)
    bc.Insert(block1)
    fmt.Println("BlockChain:", bc)



}
