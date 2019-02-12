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
     return 17
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

  func (mpt *MerklePatriciaTrie) mergeBranchAndLeaf(currNode Node, childNode Node) Node {
      currNibble := []uint8{findNonEmptySpace(currNode)}
      //childNode := mpt.db[currNode.branch_value[currNibble[0]]]
      childNibbles  := Compact_decode(childNode.flag_value.encoded_prefix)
      newLeafNibbles := append(mergeArrays(currNibble, childNibbles), 16)
      mergedLeaf := createNode(2, [17]string{}, newLeafNibbles, childNode.flag_value.value)
      delete(mpt.db, currNode.hash_node())
      delete(mpt.db, childNode.hash_node())
      mpt.addToMap(mergedLeaf)
      return mergedLeaf
  }

  func (mpt *MerklePatriciaTrie) checkBranch(branchCount int, currNode Node)  (string, Node, error) {

      if branchCount == 0 {
          return "", currNode, nil
      } else if branchCount == 1 {
          //if the position is at position 16
          //then we return "", currNode, nil
          currNibble := []uint8{findNonEmptySpace(currNode)}
          if currNibble[0] == 16 {
              return "", currNode, nil
          } else {
              childNode := mpt.db[currNode.branch_value[currNibble[0]]]
              if isLeaf(childNode) {
                  mergedLeaf := mpt.mergeBranchAndLeaf(currNode, childNode)
                  return mergedLeaf.hash_node(), mergedLeaf, nil
              } else {
                  //below me is an extension node
                  //merge myself with the extension node
                  //push me up
              }
          }
      }
      mpt.addToMap(currNode)
      return currNode.hash_node(), currNode, nil
  }
