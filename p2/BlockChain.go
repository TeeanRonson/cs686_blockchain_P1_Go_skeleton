package p2

import (
    "encoding/json"
    "fmt"
    "github.com/pkg/errors"
    "log"
    "reflect"
)

/**
Chain = map which maps a block height to a list of blocks. The value is a list so that it can handle the forks.
Length = Length equals to the highest block height.
 */
type BlockChain struct {
    Length int32
    Chain map[int32][]Block
}

func NewBlockChain() BlockChain {
    chain := make(map[int32][]Block, 0)
    //var genesis Block
    //genesis.CreateGenesisBlock()
    //chain[0] = []Block{genesis}
    return BlockChain{0, chain}
}
/**
This function takes a height as the argument,
returns the list of blocks stored in that height or None if the height doesn't exist.
 */
func (bc *BlockChain) Get(height int32) []Block {

    currChain := bc.Chain[height]
    if currChain != nil {
        return currChain
    }
    return nil
}

/**
This function takes a block as the argument,
use its height to find the corresponding list in blockchain's Chain map.
If the list already contains that block's hash,
ignore it because we don't store duplicate blocks; if not, insert the block into the list.
 */
func (bc *BlockChain) Insert(block Block) {

    currChain := bc.Get(block.Header.Height)
    fmt.Println("\nRound-----")
    fmt.Println(len(bc.Get(0)))
    fmt.Println(len(bc.Get(1)))
    fmt.Println(len(bc.Get(2)))
    fmt.Println(len(bc.Get(3)))

    if currChain == nil {
        fmt.Println("No []Block at that height: append to Block height:", block.Header.Height)
        newChain := make([]Block, 0)
        newChain = append(newChain, block)
        //add a new chain at the height
        bc.Chain[block.Header.Height] = newChain
    } else {
        fmt.Println("\nWe should see this twice------")
        fmt.Println("Before length", len(currChain))
        for _, currBlock := range currChain {
            if reflect.DeepEqual(block.Header.Hash, currBlock.Header.Hash) {
                fmt.Println("MATCH")
                return
            }
        }
        //currChain = append(currChain, block)
        //fmt.Println("After length", len(currChain))
        //bc.Chain[block.Header.Height] = currChain
        bc.Chain[block.Header.Height] = append(bc.Chain[block.Header.Height], block)
    }
    if bc.Length < block.Header.Height {
        bc.Length = block.Header.Height
    }
    fmt.Println("Top Length", bc.Length)
    fmt.Println("End Round------")
}

/**
This function iterates over all the blocks,
generate blocks' JsonString by the function you implemented previously,
and return the list of those JsonStrings
 */
func (bc *BlockChain) BlockChainEncodeToJson() (string, error) {

    jsonList := make([]BlockJson, 0)
    for _, chain := range bc.Chain {
        for _, block := range chain {
            jsonList = append(jsonList, blockToBlockJson(block))
        }
    }

    result, err := json.MarshalIndent(jsonList, "", "")
    if err != nil {
        fmt.Println("Cannot Marshal Indent jsonList")
        log.Fatal(err)
    }
    return string(result), nil
}
/**
This function is called upon a blockchain instance.
It takes a blockchain JSON string as input,
decodes the JSON string back to a list of block JSON strings,
decodes each block JSON string back to a block instance, and inserts every block into the blockchain.
 */
func BlockChainDecodeFromJson(jsonString string) (BlockChain, error) {

    newBlockChain := NewBlockChain()
    blockJsonList := make([]BlockJson, 0)

    if err := json.Unmarshal([]byte(jsonString), &blockJsonList); err != nil {
        panic(err)
        return newBlockChain, errors.New("Blockchain DecodeFromJson error")
    }

    for _, item := range blockJsonList {
        createBlock := blockJsonToBlock(item)
        newBlockChain.Insert(createBlock)
    }
    fmt.Println(len(newBlockChain.Get(0)))
    fmt.Println(len(newBlockChain.Get(1)))
    fmt.Println(len(newBlockChain.Get(2)))
    fmt.Println(len(newBlockChain.Get(3)))
    return newBlockChain, nil
}