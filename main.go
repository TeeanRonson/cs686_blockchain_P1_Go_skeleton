package main

import (
	"fmt"
	"github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p1"
)

func main() {

	fmt.Println("Testing the Get method")
	//prefix := make([]uint8, 0)
	//flagValue := p1.Flag_value{prefix, "Hello"}
	//arr := make([]string, 0)
	//nod := p1.Node{0, arr, flagValue}

	//db := make(map[string]p1.Node)
	//mpt := p1.MerklePatriciaTrie{db, "rootHash"}
	//mpt.Get("dog")

	//var hex_array = []uint8(value)

	//fmt.Println("\nRun Compact Encode")
	//fmt.Println("Result:", p1.Compact_encode(toInsert))

	//-------------------//-------------------//-------------------//-------------------//-------------------

	fmt.Println("Testing some values")
	do := p1.EncodeToHex("do")
	fmt.Println("Outside:", do);


	//fmt.Println("-------------------")
	//dog := p1.EncodeToHex("dog")
	//puppy := p1.EncodeToHex("puppy")
	//horse := p1.EncodeToHex("horse")
	//fmt.Printf("dog: %s\n", dog)
	//
	//fmt.Println()
	//fmt.Println("\nCompact Encode Values")
	//fmt.Println(p1.Compact_encode(dog))

	//fmt.Println(p1.Compact_encode([]uint8 {6, 1}))
	//fmt.Println(p1.Compact_encode([]uint8 {6, 4, 6, 15, 6, 7}))
	//fmt.Println(p1.Compact_encode(puppy))
	//fmt.Println(p1.Compact_encode(horse))

	//fmt.Println("\nTest Cases")
	//a := p1.EncodeToHex("do")
	//fmt.Printf("a: %x\n", a)
	//
	//fmt.Println("Output:", p1.Compact_encode(a))

	//p1.Test_compact_encode()












}
