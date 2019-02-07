package p1

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"reflect"
)

type Flag_value struct {
	encoded_prefix []uint8
	value string

}

/**
The variable "branch_value" is used only when it's a Branch Node
Variable "flag_value" is used only when it's an Ext or Leaf Node.
You don't have to use "branch_value" if it's an Ext or Leaf Node
 */
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
Retrieves a new MPT
 */
func GetMPTrie() *MerklePatriciaTrie {

	db := new (map[string]Node)
	root := "root"
	return &MerklePatriciaTrie{*db, root}
}

func (mpt *MerklePatriciaTrie) GetHelper2(node string, path []uint8, position int) (value string, newRoot string, err error) {

	//Do some hashing, find the path all the way down and return the value
	currNode := mpt.db[node]
	nodeType := currNode.node_type
	//Use switch statement to determine the type of node
	switch nodeType {
	case 0:
		return "", "", err
	case 1:
		//Branch Node
		//Remove the first item in the path = item
		//find the next node in the DB using the string in branch_value[item}
		//Recurse into the function
		if len(path) == 0 {
			return currNode.branch_value[16], "", err
		}
		nextHash := currNode.branch_value[path[0]]
		return mpt.GetHelper2(nextHash, path[1:], 0)
	case 2:
		//Extension/Leaf node
		//nibbles = Decode encoded_prefix at Node
		//find substring of nibblesAtNode with path
		//if nibblesAtNode == path
		//Node is a leaf node
		//return value at leaf node
		//Check for similar prefix
		//if nibblesAtNode matches path up to[len(nibblesAtNode)]
		//Node is an extension node
		//recurse down with the rest of the path
		isLeaf := false
		if ConvertToHex(currNode.flag_value.encoded_prefix)[0] < 2 {
			isLeaf = false
		} else {
			isLeaf = true
		}
		nibbles := Compact_decode(currNode.flag_value.encoded_prefix)

		if reflect.DeepEqual(nibbles, path) && isLeaf {
			return currNode.flag_value.value, "", err
		} else if !reflect.DeepEqual(nibbles, path) && isLeaf {
			return "", "", err
		} else {
			length := len(nibbles)
			nextHash := currNode.flag_value.value
			return mpt.GetHelper2(nextHash, path[length:], 0)
		}
	}
	return "", "", err
}

/**
Traverses the MPT to find the value associated with the key
 */
func (mpt *MerklePatriciaTrie) GetHelper1(path []uint8) (string, error) {

	if path == nil || mpt.root == "" {
		return "", errors.New("Path is Nil - Get")
	}

	value, newRoot, err := mpt.GetHelper2(mpt.root, path, 0)
	if err == nil {
		mpt.root = newRoot
	}
	//Find the Hash Value of the Node
	return value, errors.New("Path Not Found - Get")

}

/*
Takes a key as the argument, traverses down the MPT to find the value
if the key doesnt exist, return an empty string
 */
func (mpt *MerklePatriciaTrie) Get(key string) (string, error) {
	//Create a path array
	//Convert the key string into Hexcode
	//Add each item of the Hexcode into the Path array
	//pass MPT Tree and Key into the Helper method
	path := EncodeToHex(key)
	fmt.Println("Path:", path)
	return mpt.GetHelper1(path)
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
func (mpt *MerklePatriciaTrie) Delete(key string) error {
	return errors.New("Path Not Found - Delete")
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
		flags = append(flags, 0)
		hex_array = append(flags, hex_array...)
	}
	fmt.Println("Hex Array with Odd check", hex_array)
	//Convert result to 4 item length array
	result := make([]uint8, 0)
	for i:= 0; i < len(hex_array); i += 2 {
		result = append(result, 16*hex_array[i]+hex_array[i+1])
	}
	fmt.Println("Result Encode:", result)
	return result
}


// If Leaf, ignore 16 at the end
/*
Reverse the compact_encode function
 */
func Compact_decode(encoded_arr []uint8) []uint8 {

	unpack := ConvertToHex(encoded_arr)
	if unpack[0] < 2 { //Extension node, remove 16
		unpack = unpack[:len(unpack)-1]
	}
	fmt.Println("First Unpack: ", unpack)
	checkPrefix := 2 - unpack[0]&1
	fmt.Println(unpack[checkPrefix:])

	if unpack[len(unpack)-1] == 16 {
		return unpack[checkPrefix:len(unpack)-1]
	}
	return unpack[checkPrefix:]
}

func Test_compact_encode() {
	//fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	//fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	//fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
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