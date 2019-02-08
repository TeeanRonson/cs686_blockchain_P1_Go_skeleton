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
		isLeaf := isLeaf(currNode)
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

func (mpt *MerklePatriciaTrie) InsertHelp(parent string, curr string, encodedKey []uint8, newValue string) (newHash string){

	currNode := mpt.db[curr]
	nodeType := currNode.node_type
	switch nodeType {
	case 0:
		leaf := createNode(2, [17]string{}, encodedKey, newValue)
		hash := leaf.hash_node()
		mpt.db[hash] = leaf
		return hash
	//case1 - nodeType is branch node
		//if no more items inside encoded_key
		// - place item at branch[16]
		//Hash the branch node
		//Store the hash value in the DB[hash][newBranch node]
		//continue to hash nodes up to the root node***

	//there are still items inside encoded_key
		//check if branch[encoded_key[0]] == null
			//if null
			//create a new leaf node and add encoded_key[1:] as the nibble
			//add the value into the leaf node
			//hash the leaf node
			//store the hashed string at branch[encoded_key[0]]
			//store the hash value and the leaf node in the DB
			//continue to hash nodes up to the root node***

			//if not null
			//traverse down to next level
			//send in: string value at branch_node[x], encoded_key[1:], value
	case 1:
		if len(encodedKey) == 0 {
			currNode.branch_value[16] = newValue
			mpt.db[currNode.hash_node()] = currNode
		}

		if currNode.branch_value[encodedKey[0]] == "" {
			position := encodedKey[0]
			leaf := createNode(2, [17]string{}, encodedKey[1:], newValue)
			hash := leaf.hash_node()
			currNode.branch_value[position] = hash
			mpt.db[hash] = leaf
		} else {
			position := encodedKey[0]
			nextNode := currNode.branch_value[position]
			mpt.InsertHelp(curr, nextNode, encodedKey[1:], newValue)
		}
	case 2:
		//if there is nothing left in encodedKey, then we should have a repeat key
		//replace the value of the key
		//would we still have values in the leaf?
		if encodedKey == nil && isLeaf(currNode) {
			currNode.flag_value.value = newValue
			return
		}

		if isLeaf(currNode) {
			nibbles := Compact_decode(currNode.flag_value.encoded_prefix)
			match := 0
			for encodedKey[match] == nibbles[match] {
				match++
			}
			if match == 0 {
				branch := createNode(1, [17]string{}, []uint8{}, "")
				oldLeaf := createNode(2, [17]string{}, nibbles[1:], currNode.flag_value.value)
				newLeaf := createNode(2, [17]string{}, encodedKey[1:], newValue)
				branch.branch_value[nibbles[0]] = oldLeaf.hash_node()
				branch.branch_value[encodedKey[0]] = newLeaf.hash_node()
				mpt.db[oldLeaf.hash_node()] = oldLeaf
				mpt.db[newLeaf.hash_node()] = newLeaf
			} else {

			}
		}
	}

	//All of the nibbles match, but some extra path at the end
	//All of path matches, but with some extra nibbles


	//case2 - nodeType is an extension/leaf node
		//decode nibbles to check if node is extension or leaf
		//if leaf
			//1. find matching pattern between encoded_key and leaf nibbles
			//if no matching patterns
				//create a branch node
				//create a new leaf node for the incoming encoded_key
				//position the new and old leaf node into the branch node
				//hash the leaf nodes, store them into the DB
				//continue to hash nodes up to root node
			//if there are matching patterns
				//create an extension node and add matching patterns as the nibble except the last value
				//create a branch node that extends from the extension node
				//create a new leaf node for the incoming encoded_key
				//position the new and old leaf node into the branch node
				//hash the leaf nodes, store them into the DB
				//continue to hash nodes up to root node
		//if extension
			//2. find matching pattern between encoded_key and extension nibbles
			//if no matching patterns
				//

			//if there are matching patterns
				//change nibbles in extension node up to the matched portion
				//hash the extension node and add it to the DB
				//create a new branch node
				//hash the branch node and add it to the DB
				//return the 'hash string', remainder of the encoded key, value


}

/*
Takes a pair of <key, value> as arguments. It will traverse down the MPT and find the right place to insert the value
 */
func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	//Encode the key
	//pass an empty parent, rootString, key and value into the InsertHelp
	fmt.Println("\nNewInsertion")
	encodedKey := EncodeToHex(key)
	fmt.Println(encodedKey, new_value)
	newHash := mpt.InsertHelp("", mpt.root, encodedKey, new_value)
	if newHash != mpt.root {
		delete(mpt.db, mpt.root)
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
//If the last value is 16, it is a leaf node
 */
func Compact_encode(hex_array []uint8) []uint8 {
	term := 0
	if hex_array[len(hex_array)-1] == 16 {
		term = 1
	}
	if term == 1 {
		hex_array = hex_array[0: len(hex_array) - 1]
	}
	flags := make([]uint8, 0)
	oddlen := len(hex_array) % 2
	flags = append(flags, uint8(2*term+oddlen))
	if oddlen == 1 {
		hex_array = append(flags, hex_array...)
	} else {
		flags = append(flags, 0)
		hex_array = append(flags, hex_array...)
	}
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
	if unpack[0] < 2 { //Extension node, remove 16
		unpack = unpack[:len(unpack)-1]
	}
	checkPrefix := 2 - unpack[0]&1
	fmt.Println(unpack[checkPrefix:])

	if unpack[len(unpack)-1] == 16 {
		return unpack[checkPrefix:len(unpack)-1]
	}
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