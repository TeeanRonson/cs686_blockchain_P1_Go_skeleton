package p1

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


