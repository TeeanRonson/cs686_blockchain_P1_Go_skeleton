package p2

import (
    "bytes"
    "encoding/gob"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p1"
    "golang.org/x/crypto/sha3"
    "log"
    "time"
)

type Block struct {
    header Header
    mpt p1.MerklePatriciaTrie
}

type Header struct {
    Height int32                `json:"height"`
    Timestamp int64             `json:"timeStamp"`
    Hash string                 `json:"hash"`
    ParentHash string           `json:"parentHash"`
    Size int32                  `json:"size"`
}

type BlockJson struct {
    Height     int32             `json:"height"`
    Timestamp  int64             `json:"timeStamp"`
    Hash       string            `json:"hash"`
    ParentHash string            `json:"parentHash"`
    Size       int32             `json:"size"`
    MPT        map[string]string `json:"mpt"`
}

func (b *Block) toString() string {

    out, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }

    fmt.Println("Inside", out)
    return string(out)
}

func getBytes(value p1.MerklePatriciaTrie) []byte {

    var network bytes.Buffer        // Stand-in for a network connection
    enc := gob.NewEncoder(&network) // Will write to network.
    //dec := gob.NewDecoder(&network) // Will read from network.
    // Encode (send) the value.
    err := enc.Encode(value)
    if err != nil {
        fmt.Println("error")
        log.Fatal("encode error:", err)
    }

    // HERE ARE YOUR BYTES!!!!
    return network.Bytes()
}
/**
This function takes arguments(such as height, parentHash, and value of MPT type) and forms a block.
This is a method of the block struct.
 */
func (b *Block) NewBlock(height int32, parentHash string, value p1.MerklePatriciaTrie) {

    var header Header
    //mptAsBytes := getBytes(value)
    mptAsBytes := value.GetRoot()
    fmt.Println("bytes length", len(mptAsBytes))

    header.Height = height
    header.Timestamp = int64(time.Now().Unix())
    header.ParentHash = parentHash
    header.Size = int32(len(mptAsBytes))

    hashString := string(header.Height) + string(header.Timestamp) + header.ParentHash + value.GetRoot() + string(header.Size)
    sum := sha3.Sum256([]byte(hashString))
    header.Hash = hex.EncodeToString(sum[:])

    b.header = header
    b.mpt = value
}

/**
Reconstruct MPT from Map input
 */
func NewTrie(values map[string]string) p1.MerklePatriciaTrie {
    //db := make(map[string]p1.Node)
    //root := "root"
    mpt := p1.GetMPTrie()

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

    //Empty block
    var header Header
    newBlock := Block{}
    //Empty BlockJson
    blockJson := BlockJson{}
    if err := json.Unmarshal([]byte(jsonString), &blockJson); err != nil {
        panic(err)
    }
    mpt := NewTrie(blockJson.MPT)
    header.Height = blockJson.Height
    header.Timestamp = blockJson.Timestamp
    header.Hash = blockJson.Hash
    header.ParentHash = blockJson.ParentHash
    header.Size = blockJson.Size

    newBlock.header = header
    newBlock.mpt = mpt
    return newBlock
}

/**
This function encodes a block instance into a JSON format string.
Note that the block's value is an MPT, and you have to record all of the (key, value)
pairs that have been inserted into the MPT in your JSON string.
There's an example with details on Piazza.
 */
func (b *Block) EncodeToJson() string {
    toJson := BlockJson{
        b.header.Height,
        b.header.Timestamp,
        b.header.Hash,
        b.header.ParentHash,
        b.header.Size,
        b.mpt.GetInputs(),
    }

    jsonFormatted, err := json.Marshal(toJson)
    if err != nil {
        fmt.Println("Error in EncodeToJson")
    }
    return string(jsonFormatted)
}














