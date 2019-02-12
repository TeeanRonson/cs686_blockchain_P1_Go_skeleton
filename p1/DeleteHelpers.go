package p1

/**
Checks the number if items in the branch node
 */
func branchItems(currNode Node) int {

    count := 0
    for _, value := range currNode.branch_value {
        if value != "" {
            count++
        }
    }
    return count
}

/**
Finds the non empty space in branch node
 */
 func findNonEmptySpace(currNode Node) uint8 {

     for i, value := range currNode.branch_value {
         if value != "" {
             return uint8(i)
         }
     }
     return -1
 }

 /**
 Merge two []uint8 arrays
  */
  func mergeArrays(one []uint8, two[]uint8) []uint8 {
      var result []uint8
      for _, value := range one {
          result = append(result, value)
      }
      for _, value := range two {
          result = append(result, value)
      }
      return result
  }

  func (mpt *MerklePatriciaTrie) mergeBranchAndLeaf(currNode Node) Node {

      currNibble := []uint8{findNonEmptySpace(currNode)}
      childNode := mpt.db[currNode.branch_value[currNibble[0]]]
      childNibbles  := Compact_decode(childNode.flag_value.encoded_prefix)
      newLeafNibbles := append(mergeArrays(currNibble, childNibbles), 16)
      mergedLeaf := createNode(2, [17]string{}, newLeafNibbles, childNode.flag_value.value)
      //delete child and branch
      delete(mpt.db, currNode.hash_node())
      delete(mpt.db, childNode.hash_node())
      mpt.addToMap(mergedLeaf)
      //send up the merged leaf
      //get the parent to check if I am a leaf
      //if I am leaf, and he is an extension he has to merge with me
      //if I am a leaf and he is a branch, he places me in the right place
      return mergedLeaf
  }

  func (mpt *MerklePatriciaTrie) checkBranch(branchCount int, currNode Node, child Node)  (string, Node, error) {

      if branchCount == 0 {
          return "", currNode, nil
      } else if branchCount == 1 && isLeaf(child) {
          mergedLeaf := mpt.mergeBranchAndLeaf(currNode)
          return mergedLeaf.hash_node(), mergedLeaf, nil
      } else {
          mpt.addToMap(currNode)
          return currNode.hash_node(), currNode, nil
      }

  }
