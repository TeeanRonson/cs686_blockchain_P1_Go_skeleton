package Blockchain

type Blockchain struct {
    Length int32
    Chain map[int32][]Block
}

/**
This function takes a height as the argument,
returns the list of blocks stored in that height or None if the height doesn't exist.
 */
func Get(height int32) []Block {

    return make([]Block, 0)
}

/**
This function takes a block as the argument,
use its height to find the corresponding list in blockchain's Chain map.
If the list has already contained that block's hash,
ignore it because we don't store duplicate blocks; if not, insert the block into the list.
 */
func Insert(block Block) {


}

/**
This function iterates over all the blocks,
generate blocks' JsonString by the function you implemented previously,
and return the list of those JsonStrings
 */
func (blockchain *Blockchain) EncodeToJson() string {

    return ""
}


/**
This function is called upon a blockchain instance.
It takes a blockchain JSON string as input,
decodes the JSON string back to a list of block JSON strings,
decodes each block JSON string back to a block instance, and inserts every block into the blockchain.
 */
func (blockchain *Blockchain) DecodeFromJson() string {

    return ""
}