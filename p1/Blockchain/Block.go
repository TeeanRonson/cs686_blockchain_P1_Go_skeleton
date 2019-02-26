package Blockchain

import (
    "bytes"
    "encoding/binary"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p1/MerklePatriciaTrie"
    "golang.org/x/crypto/sha3"
    "time"
)

type Block struct {
    header Header
    mpt MerklePatriciaTrie.MerklePatriciaTrie
}

type Header struct {
    Height int32
    Timestamp int64
    Hash string
    ParentHash string
    Size int32
}

type BlockJson struct {
    Height     int32             `json:"height"`
    Timestamp  int64             `json:"timeStamp"`
    Hash       string            `json:"hash"`
    ParentHash string            `json:"parentHash"`
    Size       int32             `json:"size"`
    MPT        map[string]string `json:"mpt"`
}


/**
This function takes arguments(such as height, parentHash, and value of MPT type) and forms a block.
This is a method of the block struct.
 */
func (b *Block) NewBlock(height int32, parentHash string, value MerklePatriciaTrie.MerklePatriciaTrie) *Block {

    var header Header
    buf := &bytes.Buffer{}
    err := binary.Write(buf, binary.BigEndian, value)
    if err != nil {
        panic(err)
    }
    //fmt.Println(buf.Bytes())

    header.Height = 1
    header.Size = int32(len(buf.Bytes()))
    header.Timestamp = int64(time.Now().Unix())
    header.ParentHash = parentHash

    hash_str := string(header.Height) + string(header.Timestamp) + header.ParentHash + value.GetRoot() + string(header.Size)
    sum := sha3.Sum256([]byte(hash_str))
    header.Hash = hex.EncodeToString(sum[:])

    return &Block{header, b.mpt}

}

/**
Reconstruct MPT from Map input
 */
func createTrie(values map[string]string) MerklePatriciaTrie.MerklePatriciaTrie {

    mpt := MerklePatriciaTrie.GetMPTrie()

    for key, value := range values {
        mpt.Insert(key, value)
    }
    return mpt
}

/**
This function takes a string that represents the JSON value of a block as an input, and decodes the input string back to a block instance.
Note that you have to reconstruct an MPT from the JSON string, and use that MPT as the block's value.
 */
func DecodeFromJson(jsonString string) Block {

    newBlock := Block{}
    blockJson := BlockJson{}
    if err := json.Unmarshal([]byte(jsonString), &blockJson); err != nil {
        panic(err)
    }
    mpt := createTrie(blockJson.MPT)
    newBlock.Initial(blockJson.Height, blockJson.ParentHash, mpt)
    return newBlock
}

/**
This function encodes a block instance into a JSON format string.
Note that the block's value is an MPT, and you have to record all of the (key, value)
pairs that have been inserted into the MPT in your JSON string.
There's an example with details on Piazza.
 */
func (b *Block) EncodeToJson() string {

    return ""
}














