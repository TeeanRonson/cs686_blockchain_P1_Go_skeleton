package p2

import (
    "fmt"
    "reflect"
    "testing"
)

func TestDecodeEncode(t *testing.T) {

    message := "TestDecodeFromJson error"
    test1 := "{\"height\":1,\"timeStamp\":1234567890,\"hash\":\"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\",\"parentHash\":\"genesis\",\"size\":1174,\"mpt\":{\"charles\":\"ge\",\"hello\":\"world\"}}"
    block := DecodeFromJson(test1)

    fmt.Println(block.EncodeToJson())

    if !reflect.DeepEqual(block.EncodeToJson(), test1) {
        t.Errorf("Incorrect output match: %s", message)
    }
}

//func TestBlock_CreateGenesisBlock(t *testing.T) {
//
//    message := "TestCreateGenesisBlockError"
//    genesisJson := "{\"height\":0,\"timeStamp\":1551210463,\"hash\":\"GenesisBlock\",\"parentHash\":\"\",\"size\":0,\"mpt\":{}}"
//
//    var genesis Block
//    genesis.CreateGenesisBlock()
//
//    fmt.Println("\nHere", genesis.EncodeToJson())
//
//    if !reflect.DeepEqual(genesisJson, genesis.EncodeToJson()) {
//        t.Errorf("Incorrect genesis match %s", message)
//    }
//
//}
