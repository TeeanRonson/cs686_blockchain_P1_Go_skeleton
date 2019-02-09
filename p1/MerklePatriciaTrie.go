package p1

/**
If tree is empty
Insert a leaf node: if leaf node then we need to insert [6, 4, 6, 15, 16] into compact encode
if we see a leaf

When we create a new node, remove the old node from the DB

when comparing nibbles with paths
- common path
- rest of path
- rest of nibble

insert function : return the hashvalue of the current Node up


 */

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"reflect"
	"text/template/parse"
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


func (mpt *MerklePatriciaTrie) GetHelper2(node string, path []uint8, position int) string {

	//Do some hashing, find the path all the way down and return the value
	currNode := mpt.db[node]
	nodeType := currNode.node_type
	switch nodeType {
	case 0:
		return ""
	case 1:
		//fmt.Println("Branch")
		if len(path) == 0 {
			return currNode.branch_value[16]
		}
		nextHash := currNode.branch_value[path[0]]
		return mpt.GetHelper2(nextHash, path[1:], 0)
	case 2:
		isLeaf := isLeaf(currNode)
		//fmt.Println(isLeaf)
		nibbles := Compact_decode(currNode.flag_value.encoded_prefix)
		if reflect.DeepEqual(nibbles, path) && isLeaf {
			return currNode.flag_value.value
		} else if !reflect.DeepEqual(nibbles, path) && isLeaf {
			return ""
		} else {
			length := len(nibbles)
			nextHash := currNode.flag_value.value
			return mpt.GetHelper2(nextHash, path[length:], 0)
		}
	}
	return ""
}

/**
Traverses the MPT to find the value associated with the key
 */
func (mpt *MerklePatriciaTrie) GetHelper1(path []uint8) string {

	if path == nil || mpt.root == "" {
		fmt.Println("Nothing")
		return ""
	}

	value := mpt.GetHelper2(mpt.root, path, 0)

	return value
}

/*
Takes a key as the argument, traverses down the MPT to find the value
if the key doesnt exist, return an empty string
 */
func (mpt *MerklePatriciaTrie) Get(key string) string {
	//Create a path array
	//Convert the key string into Hexcode
	//Add each item of the Hexcode into the Path array
	//pass MPT Tree and Key into the Helper method
	path := EncodeToHex(key)
	fmt.Println("New Get\nPath:", path)
	return mpt.GetHelper1(path)
}

/**
Should we always append 16 to a leaf node?

 */
func (mpt *MerklePatriciaTrie) insertHelp(parent string, currHash string, encodedKey []uint8, newValue string) (newHash string) {

	currNode := mpt.db[currHash]
	nodeType := currNode.node_type
	switch nodeType {
	case 0: //null
		encodedKey = append(encodedKey, 16)
		leaf := createNode(2, [17]string{}, encodedKey, newValue)
		mpt.addToMap(leaf)
		return leaf.hash_node()
	case 1: //branch
		if len(encodedKey) == 0 {
			delete(mpt.db, currHash)
			currNode.branch_value[16] = newValue
			mpt.addToMap(currNode)
			return currNode.hash_node()
		}

		if currNode.branch_value[encodedKey[0]] == "" {
			delete(mpt.db, currHash)
			encodedKey = append(encodedKey, 16)
			newLeaf := createNode(2, [17]string{}, encodedKey[1:], newValue)
			mpt.addLeavesToBranch(newLeaf, &currNode, encodedKey[0])
			mpt.addToMap(currNode)
			return currNode.hash_node()
		} else {
			fmt.Println("Into Branch")
			nextHash := currNode.branch_value[encodedKey[0]]
			newHash := mpt.insertHelp(currHash, nextHash, encodedKey[1:], newValue)
			//fmt.Println(2)
			if newHash != nextHash {
				delete(mpt.db, currHash)
				//fmt.Println(3)
				currNode.branch_value[encodedKey[0]] = newHash
				//fmt.Println("Checking currNode:", currNode)
				mpt.addToMap(currNode)
				return currNode.hash_node()
			}
			return currHash
		}
	case 2: //leaf/extension
		nibbles := Compact_decode(currNode.flag_value.encoded_prefix)
		match := findMatch(0, nibbles, encodedKey)
		if isLeaf(currNode) {
			fmt.Println("InLeaf")
			if match == 0 { //no match
				fmt.Println("No Match")
				return mpt.breakNodeNoMatch(currNode, nibbles, encodedKey, newValue, true)
			} else if reflect.DeepEqual(encodedKey, nibbles) { //full match, replace value
				fmt.Println("Full Match")
				delete(mpt.db, currHash)
				currNode.flag_value.value = newValue
				mpt.addToMap(currNode)
				return currNode.hash_node()
			} else if match > 0 && int(match) <= len(nibbles) { //partial matches - 3 types
				fmt.Println("Partial Match")
				if len(encodedKey[match:]) != 0 && len(nibbles[match:]) != 0 { //excess path and excess nibbles
					fmt.Println("Excess Path & Excess Nibbles")
					return mpt.breakNodeDoubleExcess(currNode, match, nibbles, encodedKey, newValue, true)
				} else if len(encodedKey[match:]) != 0 && len(nibbles[match:]) == 0 { //partial match with excess path only
					return mpt.breakLeafSingleExcess(currNode, match, nibbles, encodedKey, newValue, true)
				} else if len(encodedKey[match:]) == 0 && len(nibbles[match:]) != 0 { //partial match with excess nibbles
					return mpt.breakLeafSingleExcess(currNode, match, nibbles, encodedKey, newValue, false)
				}
			}
		} else { //is extension
			fmt.Println("Entering Extension Node")
			match := findMatch(0, nibbles, encodedKey)
			fmt.Println("Number of matches:", match)
			if match == 0 { //no match
				return mpt.breakNodeNoMatch(currNode, nibbles, encodedKey, newValue, false)
			} else if reflect.DeepEqual(nibbles, encodedKey) { //exact match
				return mpt.insertHelp(currHash, currNode.flag_value.value, encodedKey[match:], newValue)
			} else if match > 0 && int(match) <= len(nibbles) {
				fmt.Println("Partial Match")
				if len(encodedKey[match:]) != 0 && len(nibbles[match:]) != 0 {
					fmt.Println("Case 1")
					return mpt.breakNodeDoubleExcess(currNode, match, nibbles, encodedKey, newValue, false)
				} else if len(encodedKey[match:]) != 0 && len(nibbles[match:]) == 0 {
					fmt.Println("Case 2")
					newHash = mpt.insertHelp(currHash, currNode.flag_value.value, encodedKey[match:], newValue)
					if newHash != currNode.flag_value.value {
						delete(mpt.db, currHash)
						currNode.flag_value.value = newHash
						mpt.addToMap(currNode)
						return currNode.hash_node()
					}
				} else if len(encodedKey[match:]) == 0 && len(nibbles[match:]) != 0 {
					fmt.Println("Case 3")
					return mpt.breakExtSingleExcess(currNode, match, nibbles, encodedKey, newValue, false)
				}
			}
		}
	}
	return currHash
}

/*
Takes a pair of <key, value> as arguments. It will traverse down the MPT and find the right place to insert the value
 */
func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {

	fmt.Println("\nNewInsertion")
	encodedKey := EncodeToHex(key)
	fmt.Println("Insert Path:", encodedKey, new_value)
	newHash := mpt.insertHelp("", mpt.root, encodedKey, new_value)
	if newHash != mpt.root {
		mpt.root = newHash
		fmt.Println("Newhash:", newHash)
		fmt.Println("DB final:", mpt.db)
	}
}

func (mpt *MerklePatriciaTrie) deleteHelp(parent string, currHash string, path []uint8) string {

	currNode := mpt.db[currHash]
	nodeType := currNode.node_type
	switch nodeType {

	}

	return currHash
}
/*
Function takes a key as the argument, traverses down the MPT and finds the Key.
If the key exists, delete the corresponding value and re-balance the Trie, if necessary.
if the key doesn't exist, return 'path_not_found'
 */
func (mpt *MerklePatriciaTrie) Delete(key string) error {

	fmt.Println("\nNewDeletion")
	encodedKey := EncodeToHex(key)
	fmt.Println("Delete Path:", encodedKey)
	newHash := mpt.deleteHelp("", mpt.root, encodedKey)
	if newHash != mpt.root {
		mpt.root = newHash
		fmt.Println("Newhash:", newHash)
		fmt.Println("DB final:", mpt.db)
	}

	return errors.New("Path Not Found - Delete")
}
/*
Function takes an array of HEX value as the input, mark the Node type(Branch, Leaf, Extension),
makes sure the length is even, and converts it into an array of ASCII numbers as the output.
//If the last value is 16, it is a leaf node
 */
func Compact_encode(hex_array []uint8) []uint8 {
	//fmt.Println("Encoding........")
	//fmt.Println(hex_array)
	term := 0
	if hex_array[len(hex_array)-1] == 16 {
		hex_array = hex_array[0: len(hex_array) - 1]
		term = 1
	}

	flags := make([]uint8, 0)
	oddlen := len(hex_array) % 2
	flags = append(flags, uint8(2*term+oddlen))
	//fmt.Println(flags)
	if oddlen == 1 {
		hex_array = append(flags, hex_array...)
	} else {
		flags = append(flags, 0)
		hex_array = append(flags, hex_array...)
	}
	//fmt.Println("hex:", hex_array)
	result := make([]uint8, 0)
	for i:= 0; i < len(hex_array); i += 2 {
		result = append(result, 16*hex_array[i]+hex_array[i+1])
	}
	return result
}


// If Leaf, ignore 16 at the end
/*
Reverse the compact_encode function
 */
func Compact_decode(encoded_arr []uint8) []uint8 {

	unpack := ConvertToHex(encoded_arr)
	//if unpack[0] < 2 { //Extension node, remove 16
	//	unpack = unpack[:len(unpack)-1]
	//}

	checkPrefix := 2 - unpack[0]&1
	//if unpack[len(unpack)-1] == 16 {
	//	return unpack[checkPrefix:len(unpack)-1]
	//}
	return unpack[checkPrefix:]
}

func Test_compact_encode() {
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
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