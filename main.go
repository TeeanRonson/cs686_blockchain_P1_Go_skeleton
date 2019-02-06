package main

import (
	"fmt"
	"github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p1"
)

func main() {

	//fmt.Println("Testing the Get method")
	mpt := p1.GetMPTrie()
	do := "do"
	result, err := mpt.Get(do)
	fmt.Println(result, err)






	//-------------------//-------------------//------------------//-------------------//-------------------

	//fmt.Println("Testing some values")
	//do := p1.EncodeToHex("do")
	//fmt.Println("do:", do);
	//
	//dog := p1.EncodeToHex("dog")
	//fmt.Println("dog:", dog)
	//
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

}
