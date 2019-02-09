package p1

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

    length := len(encoded_arr)*2 + 1
    hex_values := make([]uint8, length)
    for i, value := range encoded_arr {
        hex_values[i*2] = value/16
        hex_values[i*2+1] = value%16
    }

    hex_values[len(hex_values)-1] = 16
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

 func findMatch(match int, nibbles []uint8, encodedKey []uint8) int {

     for match < len(encodedKey) && match < len(nibbles) && encodedKey[match] == nibbles[match] {
         match++
     }
     return match
 }

 func (mpt *MerklePatriciaTrie) breakLeafNoMatch(currNode Node, nibbles []uint8, encodedKey []uint8, newValue string) string {
     delete(mpt.db, currNode.hash_node())
     newBranch := createNode(1, [17]string{}, []uint8{}, "")
     encodedKey = append(encodedKey, 16)
     newLeaf1 := createNode(2, [17]string{}, Compact_encode(encodedKey[1:]), newValue) //should we append 16
     newLeaf2 := createNode(2, [17]string{}, Compact_encode(nibbles[1:]), currNode.flag_value.value)
     newBranch.branch_value[encodedKey[0]] = newLeaf1.hash_node()
     newBranch.branch_value[nibbles[0]] = newLeaf2.hash_node()
     mpt.db[newLeaf1.hash_node()] = newLeaf1
     mpt.db[newLeaf2.hash_node()] = newLeaf2
     mpt.db[newBranch.hash_node()] = newBranch

     return newBranch.hash_node()
 }


