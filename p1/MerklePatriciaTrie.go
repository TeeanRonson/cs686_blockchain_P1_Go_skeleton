package p1

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"reflect"
)

type Flag_value struct {
	encoded_prefix []uint8
	value string

}

type Node struct {
	node_type int // 0: Null, 1: Branch, 2: Ext or Leaf
	branch_value [17]string
	flag_value Flag_value
}

type MerklePatriciaTrie struct {
	db map[string]Node
	root string //root hash
}

/**
Traverses the MPT to find the value associated with the key
 */
func getHelper(key string, mpt *MerklePatriciaTrie) string {

	return ""
}

/*
Takes a key as the argument, traverses down the MPT to find the value
iF the key doesnt exist, return an empty string
 */
func (mpt *MerklePatriciaTrie) Get(key string) string {

	//Create a path array
	//Convert the key string into Hexcode to finds its path through the MPT
	//Add each item of the Hexcode into the Path array
	//pass MPT Tree and Key into the Helper method
	path := make([]uint8, 0)
	keyPath := EncodeToHex(key)
	fmt.Println(keyPath)
	for _, value := range keyPath {
		path = append(path, value)
	}
	fmt.Println("path:", path)
	return ""
}

/*
Takes a pair of <key, value> as arguments. It will traverse down the MPT and find the right place to insert the value
 */
func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {

	//flag := Flag_value{EncodeToHex(key), new_value}
	//node := Node{2, nil, flag}
}

func (mpt *MerklePatriciaTrie) InsertHelp(node Node) {
	if len(mpt.root) == 0 {

	}
}

/*
Function takes a key as the argument, traverses down the MPT and finds the Key.
If the key exists, delete the corresponding value and re-balance the Trie, if necessary.
if the key doesn't exist, return 'path_not_found'
 */
func (mpt *MerklePatriciaTrie) Delete(key string) {
	// TODO
}


/**
Encodes the incoming Key(string) into Hex values

Example: do --> 646f in String form
			--> [54 52 54 102] in ASCII form
			--> for each value returned in the string we want to convert
				that into values from 1-16: where f = 16
			--> Algo:
				for each value in 6 4 6 f
					if value == number
						add to list
					else
						convert letter to value
 */
func EncodeToHex(key string) []uint8 {

	//for each item in the key
		//find its ASCII value
		//Find the hex values associated with that ASCII
		//Add it into the []uint8
	result := make([]uint8, 0)
	ascii := []byte(key)
	fmt.Println(ascii)
	for _, value := range ascii {
		result = append(result, value/16)
		result = append(result, value % 16)
	}



	//src := []uint8(key)
	//dst := make([]uint8, hex.EncodedLen(len(src)))
	//hex.Encode(dst, src)
	//fmt.Printf("Hex value: %s\n", dst)
	return result
}

func convertToChar1(letter string) int8 {

	switch letter {
	case "a":
		return 10 //index
	case "b":
		return 11
	case "c":
		return 12
	case "d":
		return 13
	case "e":
		return 14
	case "f":
		return 15
	}
	return -1
}

func convertToChar2(letter string) int8 {

	return -1
}
/*
Function takes an array of HEX value as the input, mark the Node type(Branch, Leaf, Extension),
makes sure the length is even, and converts it into an array of ASCII numbers as the output.
 */
func Compact_encode(hex_array []uint8) []uint8 {
	fmt.Println("Hex Array Original", hex_array)
	//0 --> extension 1 --> leaf node
	//If the last value is 16, it is a leaf node
	term := 0
	if hex_array[len(hex_array)-1] == 16 {
		term = 1
	}
	//Remove the last two values i.e. 16
	if term == 1 {
		hex_array = hex_array[0: len(hex_array) - 1]
	}
	fmt.Println("Hex Array Modified", hex_array)
	//create a new flags slice
	flags := make([]uint8, 0)
	oddlen := len(hex_array) % 2
	flags = append(flags, uint8(2*term+oddlen))
	fmt.Println("Flags", flags)
	//If the length is odd
	if oddlen == 1 {
		hex_array = append(flags, hex_array...)
	} else {
		flags = append(append(flags, 0))
		hex_array = append(flags, hex_array...)
	}
	fmt.Println("Hex Array with Odd check", hex_array)
	//Convert result to 4 item length array
	result := make([]uint8, 0)
	for i:= 0; i < len(hex_array); i += 2 {
		result = append(result, 16*hex_array[i]+hex_array[i+1])
	}
	fmt.Println(result)
	return result
}

// If Leaf, ignore 16 at the end
/*
Reverse the compact_encode function
 */
func compact_decode(encoded_arr []uint8) []uint8 {
	return []uint8{}
}

func Test_compact_encode() {
	fmt.Println(reflect.DeepEqual(compact_decode(Compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	//fmt.Println(reflect.DeepEqual(compact_decode(Compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	//fmt.Println(reflect.DeepEqual(compact_decode(Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	//fmt.Println(reflect.DeepEqual(compact_decode(Compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
}

/*
 */
func (node *Node) hash_node() string {
	var str string
	switch node.node_type {
	case 0:
		str = ""
	case 1:
		str = "branch_"
		for _, v := range node.branch_value {
			str += v
		}
	case 2:
		str = node.flag_value.value
	}

	sum := sha3.Sum256([]byte(str))
	return "HashStart_" + hex.EncodeToString(sum[:]) + "_HashEnd"
}