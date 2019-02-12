package p1

import "fmt"

/**
Retrieves a new MPT
 */
func GetMPTrie() MerklePatriciaTrie {

    db := make(map[string]Node)
    root := "root"

    mpt := MerklePatriciaTrie{db, root}
    return mpt
}

/**
Converts the input key string into Hex format
 */
func EncodeToHex(key string) []uint8 {

    result := make([]uint8, 0)
    ascii := []byte(key)
    //fmt.Println("ascii:", ascii)
    for _, value := range ascii {
        result = append(result, value/16)
        result = append(result, value%16)
    }
    return result
}
/**
Converts the input key []uint8 into Hex format
which is passed as ASCII values
 */
func ConvertToHex(encoded_arr []uint8) []uint8 {

    length := len(encoded_arr)*2
    hex_values := make([]uint8, length)
    for i, value := range encoded_arr {
        hex_values[i*2] = value/16
        hex_values[i*2+1] = value%16
    }

    //hex_values[len(hex_values)-1] = 16
    return hex_values
}

/**
Creates a new Node
 */
func createNode(nodeType int, branchValue [17]string, encodedKey []uint8, newValue string) Node {
    encode := make([]uint8, 0)
    if len(encodedKey) != 0 {
        encode = Compact_encode(encodedKey)
    }
    flag := Flag_value{encode, newValue}
    newNode := Node{nodeType, branchValue, flag}
    return newNode
}

/**
Checks if the node is a leaf or extension node
 */
func isLeaf(currNode Node) bool {
    if ConvertToHex(currNode.flag_value.encoded_prefix)[0] < 2 {
        return false
    }
    return true
}

/**
find the matched portion of nibbles & encodedKey
 */

 func findMatch(match int, nibbles []uint8, encodedKey []uint8) uint8 {

     for match < len(encodedKey) && match < len(nibbles) && encodedKey[match] == nibbles[match] {
         match++
     }
     return uint8(match)
 }

 /**
 Breaks a leaf during insertion when there are no matches between input key and nibbles at leaf
  */
 func (mpt *MerklePatriciaTrie) breakNodeNoMatch(currNode Node, nibbles []uint8, encodedKey []uint8, newValue string, isLeaf bool) string {
     delete(mpt.db, currNode.hash_node())
     if isLeaf {
         nibbles = append(nibbles, 16) //may be an extension or a leaf
         encodedKey = append(encodedKey, 16) //always a leaf
     }

     if len(encodedKey[1:]) == 0 {
         //add new value to newBranch[16] = newValue
     }
     newLeaf1 := createNode(2, [17]string{}, encodedKey[1:], newValue)
     newLeaf2 := createNode(2, [17]string{}, nibbles[1:], currNode.flag_value.value)
     newBranch := createNode(1, [17]string{}, []uint8{}, "")
     mpt.addLeavesToBranch(newLeaf1, &newBranch, encodedKey[0])
     mpt.addLeavesToBranch(newLeaf2, &newBranch, nibbles[0])
     mpt.addToMap(newLeaf1)
     mpt.addToMap(newLeaf2)
     mpt.addToMap(newBranch)
     return newBranch.hash_node()
 }

 /**
 Partial match in the leaf with excess hex values in both nibbles and incoming encodedKey
  */
 func (mpt *MerklePatriciaTrie) breakNodeDoubleExcess(currNode Node, match uint8, nibbles []uint8, encodedKey []uint8, newValue string, isLeaf bool) string {
     delete(mpt.db, currNode.hash_node()) //delete my old leaf self from the db
     nibbles = append(nibbles, 16)
     encodedKey = append(encodedKey, 16)

     if !isLeaf {
         nibbles = nibbles[:len(nibbles)-1]
     }
     newLeaf1 := createNode(2, [17]string{}, nibbles[match + 1:], currNode.flag_value.value)
     newLeaf2 := createNode(2, [17]string{}, encodedKey[match + 1:], newValue)
     newBranch := createNode(1, [17]string{}, []uint8{}, "") //create a branch node
     mpt.addLeavesToBranch(newLeaf1, &newBranch, nibbles[match])
     mpt.addLeavesToBranch(newLeaf2, &newBranch, encodedKey[match])
     extension := createNode(2, [17]string{}, nibbles[0:match], newBranch.hash_node()) //change myself to an extension node
     mpt.addToMap(newLeaf1)
     mpt.addToMap(newLeaf2)
     mpt.addToMap(newBranch)
     mpt.addToMap(extension)
     return extension.hash_node()
 }

 /**
 Breaks a leaf during insertion when there is a partial match between nibbles and encodedkey
 Extension node: stores the partial matched
 Branch node: extends from the extension
 leafNode: newLeaf node placed in the respective branch node array
  */
 func (mpt *MerklePatriciaTrie) breakLeafSingleExcess(currNode Node, match uint8,  nibbles []uint8, encodedKey []uint8, newValue string, excessPath bool) string{
     delete(mpt.db, currNode.hash_node())
     pathway := nibbles
     index := nibbles[match]
     if excessPath {
         pathway = encodedKey
         index = encodedKey[match]
     }
     pathway = append(pathway, 16)
     newLeaf := createNode(2, [17]string{}, pathway[match+1:], newValue)
     newBranch := createNode(1, [17]string{}, []uint8{}, "")
     newBranch.branch_value[16] = currNode.flag_value.value
     mpt.addLeavesToBranch(newLeaf, &newBranch, index)
     extension := createNode(2, [17]string{}, pathway[0:match], newBranch.hash_node())
     mpt.addToMap(newLeaf)
     mpt.addToMap(newBranch)
     mpt.addToMap(extension)
     return extension.hash_node()
 }

 func (mpt *MerklePatriciaTrie) breakExtSingleExcess(currNode Node, match uint8, nibbles []uint8, encodedKey []uint8, newValue string, excessPath bool) string {
     delete(mpt.db, currNode.hash_node())
     pathway := nibbles
     index := nibbles[match]
     if excessPath {
         pathway = encodedKey
         index = encodedKey[match]
     }
     lowerExt := createNode(2, [17]string{}, pathway[match+1:], currNode.flag_value.value)
     newBranch := createNode(1, [17]string{}, []uint8{}, "")
     newBranch.branch_value[16] = newValue
     mpt.addLeavesToBranch(lowerExt, &newBranch, index)
     upperExt := createNode(2, [17]string{}, pathway[0:match], newBranch.hash_node())
     mpt.addToMap(lowerExt)
     mpt.addToMap(newBranch)
     mpt.addToMap(upperExt)
     fmt.Println(1)
     fmt.Println(upperExt.hash_node())
     return upperExt.hash_node()


 }

 /**
 Adds a new Node to the DB
  */
 func (mpt *MerklePatriciaTrie) addToMap(newNode Node) {
     mpt.db[newNode.hash_node()] = newNode
 }

 /**
 Adds the newLeaf to the branch node at its respective position
  */
func (mpt *MerklePatriciaTrie) addLeavesToBranch(newLeaf Node, branch *Node, index uint8) {
    branch.branch_value[index] = newLeaf.hash_node()
}
