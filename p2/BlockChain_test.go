package p2

import (
    "encoding/json"
    "fmt"
    "reflect"
    "testing"
)


func TestBlockChainBasic(t *testing.T) {
    jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
    bc, err := BlockChainDecodeFromJson(jsonBlockChain)

    if err != nil {
        fmt.Println("BlockChain Decode error:", err)
        t.Fail()
    }
    jsonNew, err := bc.BlockChainEncodeToJson()
    fmt.Println("JsonNew:", jsonNew)
    if err != nil {
        fmt.Println("BlockChain Encode error:", err)
        t.Fail()
    }

    var realValue []BlockJson
    var expectedValue []BlockJson
    err = json.Unmarshal([]byte(jsonNew), &realValue)
    if err != nil {
        fmt.Println("real value error:", err)
        t.Fail()
    }
    err = json.Unmarshal([]byte(jsonBlockChain), &expectedValue)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }

    if !reflect.DeepEqual(realValue, expectedValue) {
        fmt.Println("=========Real=========")
        fmt.Println(realValue)
        fmt.Println("=========Expected=========")
        fmt.Println(expectedValue)
        t.Fail()
    }
}