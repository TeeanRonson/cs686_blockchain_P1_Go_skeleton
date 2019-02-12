package main

import (
	"fmt"
	"github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p1"
)

func main() {

	//fmt.Println("Testing the Insert method")
	mpt := p1.GetMPTrie()
	mpt.Insert("p", "apple")
	mpt.Insert("aaaaa", "banana")
	mpt.Insert("aaaap", "orange")
	mpt.Insert("aa", "new")
	mpt.Insert("aaaab", "candle")
	mpt.Insert("king", "king")
	mpt.Insert("abc", "alphabet")
	//result, err := mpt.Delete("p")
	//fmt.Println(result, err)
	//deleteBanana, err := mpt.Delete("aaaaa")
	//fmt.Println(deleteBanana, err)



	fmt.Println("\nGet test")
	apple := mpt.Get("p")
	banana := mpt.Get("aaaaa")
	orange := mpt.Get("aaaap")
	newWord := mpt.Get("aa")
	candle := mpt.Get("aaaab")
	king := mpt.Get("king")
	alphabet := mpt.Get("abc")
	fmt.Println("Apple:", apple)
	fmt.Println("Banana:", banana)
	fmt.Println("Orange:", orange)
	fmt.Println("New:", newWord)
	fmt.Println("Candle:", candle)
	fmt.Println("King:", king)
	fmt.Println("alphabet:", alphabet)


	//fmt.Println(p1.Compact_encode([]uint8{7, 0}))
	//fmt.Println(p1.Compact_encode([]uint8{7, 0, 16}))

	//fmt.Println(p1.EncodeToHex("p"))
	//fmt.Println(p1.Compact_encode(p1.EncodeToHex("p")))
	//fmt.Println(p1.Compact_decode(p1.Compact_encode(p1.EncodeToHex("p"))))

	//-------------------//-------------------//------------------//-------------------//-------------------

	//fmt.Println("Testing some values")

	//mpt.Test()
	//encodedKey := []uint8{1, 6, 1, 6}
	//nibbles := []uint8{1, 6, 1, 6, 1, 6, 1, 6, 1}
	//
	//match := 0
	//for match < len(encodedKey) && encodedKey[match] == nibbles[match] {
	//	match++
	//}
	//fmt.Println(encodedKey[0:match])
	//fmt.Println(nibbles[0:match])
	//fmt.Println(len(encodedKey))
	//fmt.Println(match)
	//fmt.Println(encodedKey[match:])
	//fmt.Println(nibbles[match:])

	//dog := p1.EncodeToHex("dog")
	//fmt.Println("dog:", dog)
	//
	//fmt.Println(p1.Compact_encode([]uint8{1,6,1}))
	//doge := p1.EncodeToHex("doge")
	//fmt.Println("doge:", doge)
	//
	//horse := p1.EncodeToHex("horse")
	//fmt.Println("horse:", horse)
	//
	//fmt.Println(p1.Compact_encode([]uint8{6, 1, 5, 7, 2, 7, 3, 6, 5}))

	//fmt.Println("\nCompact Encode Values")
	//fmt.Println("Final-----------:", p1.Compact_encode(do))
	//fmt.Println("Final-----------:",p1.Compact_encode(dog))
	//fmt.Println("Final-----------:",p1.Compact_encode(doge))
	//fmt.Println("Final-----------:",p1.Compact_encode(horse))

	//fmt.Println("Test Encode:", p1.Compact_encode([]uint8 {1, 6, 1}))
	//fmt.Println("Test Decode:", p1.CompactToHex([]uint8 {17, 97}))
	//fmt.Println("Test Decode:", p1.CompactToHex([]uint8 {15, 1, 12, 11, 8, 16}))


	//fmt.Println("\nTest Cases")
	//p1.Test_compact_encode()
	//
	//fmt.Println(p1.Compact_encode([]uint8{1, 2, 3, 4, 5}))
	//fmt.Println()
	//fmt.Println(p1.Compact_encode([]uint8{0, 1, 2, 3, 4, 5}))
	//fmt.Println()
	//fmt.Println(p1.Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16}))
	//fmt.Println()
	//fmt.Println(p1.Compact_decode([]uint8{32, 15, 28, 184}))


}
