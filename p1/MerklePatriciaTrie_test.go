package p1

import "testing"

func TestMerklePatriciaTrie_Get1(t *testing.T) {

    mpt := GetMPTrie()
    mpt.Insert("p", "apple")
    mpt.Insert("aaaaa", "banana")
    mpt.Insert("aaaap", "orange")
    mpt.Insert("aa", "new")
    mpt.Insert("aaaab", "candle")
    mpt.Insert("king", "king")
    mpt.Insert("abc", "alphabet")

    apple := mpt.Get("p")
    banana := mpt.Get("aaaaa")
    orange := mpt.Get("aaaap")
    newWord := mpt.Get("aa")
    candle := mpt.Get("aaaab")
    king := mpt.Get("king")
    alphabet := mpt.Get("abc")
    if  apple != "apple" || banana != "banana" || orange != "orange" || newWord != "new" || candle != "candle" || king != "king" || alphabet != "alphabet"{
        t.Errorf("Result is %s", apple)
        t.Errorf("Result is %s", banana)
        t.Errorf("Result is %s", orange)
        t.Errorf("Result is %s", newWord)
        t.Errorf("Result is %s", candle)
        t.Errorf("Result is %s", king)
        t.Errorf("Result is %s", alphabet)
    }


}
