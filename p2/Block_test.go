package p2

import (
    "testing"
)

func TestDecodeFromJson(t *testing.T) {

    message := "TestDecodeFromJson error"
    test1 := "{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}"
    testResult := "{{1 1234567890 3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48 genesis 1174} " +
        "{map[HashStart_420baf620e3fcd9b3715b42b92506e9304d56e02d3a103499a3a292560cb66b2_HashEnd:{2 [                ] " +
        "{[32 101 108 108 111] world}} HashStart_fc0859986f03df2a1c373e98e983ba74c226b10fe709e728ef7f8e3395ce057e_HashEnd:" +
        "{2 [                ] {[32 104 97 114 108 101 115] ge}} HashStart_45b4d59f210cc39eb106b24ed127cd1c7435cf00eb28a3bcaf" +
        "820841a8c96327_HashEnd:{1 [   HashStart_fc0859986f03df2a1c373e98e983ba74c226b10fe709e728ef7f8e3395ce057e_HashEnd     " +
        "HashStart_420baf620e3fcd9b3715b42b92506e9304d56e02d3a103499a3a292560cb66b2_HashEnd        ] {[] }} HashStart_8c5f2eb9500" +
        "ff43ef88a772cbd51f19705b887d4e03757bf1601135da8534703_HashEnd:{2 [                ] {[22] HashStart_45b4d59f210cc39eb106b24" +
        "ed127cd1c7435cf00eb28a3bcaf820841a8c96327_HashEnd}}] map[hello:world charles:ge] HashStart_8c5f2eb9500ff43ef88a772cbd51f19705b" +
        "887d4e03757bf1601135da8534703_HashEnd}}"

    var block Block
    block = DecodeFromJson(test1)

    //fmt.Println("Empty", block.toString(), block)

    if block.toString() != testResult {
        t.Errorf("Incorrect output match %s", message)
    }
}
