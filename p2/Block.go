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
    Header Header
    Mpt p1.MerklePatriciaTrie
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

/**
convert block to JsonString
 */
func (b *Block) toString() string {

    out, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    fmt.Println("toString func in Block class")
    return string(out)
}

/**
Convert the MPT struct into bytes
 */
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
    fmt.Println("Stopped")

    // HERE ARE YOUR BYTES!!!!
    return network.Bytes()
}
/**
This function takes arguments(such as height, parentHash, and value of MPT type) and forms a block.
This is a method of the block struct.
 */
func (b *Block) NewBlock(height int32, parentHash string, value p1.MerklePatriciaTrie) {

    var header Header
    mptAsBytes := getBytes(value)
    fmt.Println("bytes length:", len(mptAsBytes))

    header.Height = height
    header.Timestamp = int64(time.Now().Unix())
    header.ParentHash = parentHash
    header.Size = int32(len(mptAsBytes))

    hashString := string(header.Height) + string(header.Timestamp) + header.ParentHash + value.Root + string(header.Size)
    sum := sha3.Sum256([]byte(hashString))
    header.Hash = hex.EncodeToString(sum[:])

    b.Header = header
    b.Mpt = value
}

func (b *Block) CreateGenesisBlock() {

    header := Header{0, int64(time.Now().Unix()), "GenesisBlock", "", 0}
    b.Mpt = p1.GetMPTrie()
    b.Header = header
}

/**
Reconstruct MPT from Map input
 */
func NewTrie(values map[string]string) p1.MerklePatriciaTrie {

    mpt := p1.GetMPTrie()
    for key, value := range values {
        mpt.Insert(key, value)
    }
    return mpt
}

func convertToBlockJson(jsonString string) (BlockJson, error) {

   blockJson := BlockJson{}
   if err := json.Unmarshal([]byte(jsonString), &blockJson); err != nil {
       panic(err)
       return blockJson, err
   }
   return blockJson, nil
}

/**
This function takes a string that represents the JSON value of a block as an input, and decodes the input string back to a block instance.
Note that you have to reconstruct an MPT from the JSON string, and use that MPT as the block's value.
 */
func DecodeFromJson(jsonString string) (Block, error) {

  var header Header
  newBlock := Block{}
  blockJson, err := convertToBlockJson(jsonString)
  if err != nil {
      return newBlock, err
  }

  mpt := NewTrie(blockJson.MPT)
  header.Height = blockJson.Height
  header.Timestamp = blockJson.Timestamp
  header.Hash = blockJson.Hash
  header.ParentHash = blockJson.ParentHash
  header.Size = blockJson.Size

  newBlock.Header = header
  newBlock.Mpt = mpt
  return newBlock, nil
}

func DecodeFromJson2(blockJson BlockJson) Block {

    var header Header
    newBlock := Block{}
    mpt := NewTrie(blockJson.MPT)
    header.Height = blockJson.Height
    header.Timestamp = blockJson.Timestamp
    header.Hash = blockJson.Hash
    header.ParentHash = blockJson.ParentHash
    header.Size = blockJson.Size

    newBlock.Header = header
    newBlock.Mpt = mpt
    return newBlock
}

/**
This function encodes a block instance into a JSON format string.
Note that the block's value is an MPT, and you have to record all of the (key, value)
pairs that have been inserted into the MPT in your JSON string.
There's an example with details on Piazza.
 */
func (b *Block) EncodeToJson() (string, error) {

    toJson := BlockJson{
        b.Header.Height,
        b.Header.Timestamp,
        b.Header.Hash,
        b.Header.ParentHash,
        b.Header.Size,
        b.Mpt.Inputs,
    }

    jsonFormatted, err := json.Marshal(toJson)
    if err != nil {
        fmt.Println("Error in EncodeToJson")
        return string(jsonFormatted), err
    }
    return string(jsonFormatted), nil
}














