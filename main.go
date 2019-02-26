package main

import (
    "fmt"
    "github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p2"
)

func main() {

    test1 := "{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}"
    block := p2.DecodeFromJson(test1)

    fmt.Println("block", block)
    block.EncodeToJson()


}
