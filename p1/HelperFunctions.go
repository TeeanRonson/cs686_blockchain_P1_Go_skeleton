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
    //fmt.Println("ascii: ", ascii)
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

    flag := Flag_value{Compact_encode(encodedKey), newValue}
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
    return  true
}


