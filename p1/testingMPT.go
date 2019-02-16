package p1

import ("fmt"
)

func TestInsertAndGet() {

    mpt := GetMPTrie()
    mpt.Insert("p", "apple")
    mpt.Insert("aaaaa", "banana")
    mpt.Insert("aaaap", "orange")
    mpt.Insert("aa", "new")
    mpt.Insert("aaaab", "candle")
    mpt.Insert("king", "king")
    mpt.Insert("abc", "alphabet")


    fmt.Println("\nGet test")
    apple := mpt.Get("p")
    banana := mpt.Get("aaaaa")
    orange := mpt.Get("aaaap")
    newWord := mpt.Get("aa")
    candle := mpt.Get("aaaab")
    king := mpt.Get("king")
    alphabet := mpt.Get("abc")
    fmt.Println("Apple:", apple)
    fmt.Println("Banana:", banana)
    fmt.Println("Orange:", orange)
    fmt.Println("New:", newWord)
    fmt.Println("Candle:", candle)
    fmt.Println("King:", king)
    fmt.Println("alphabet:", alphabet)

    fmt.Println("------------------------------\n")

    //result, err := mpt.Delete("p")
    //fmt.Println(result, err)
    //deleteBanana, err := mpt.Delete("aaaaa")
    //fmt.Println(deleteBanana, err)
    //
    //fmt.Println("\nPost Get test")
    //appleDelete := mpt.Get("p")
    //bananaDelete := mpt.Get("aaaaa")
    //orangeDelete := mpt.Get("aaaap")
    ////newWord := mpt.Get("aa")
    ////candle := mpt.Get("aaaab")
    ////king := mpt.Get("king")
    ////alphabet := mpt.Get("abc")
    //fmt.Println("Apple:", appleDelete)
    //fmt.Println("Banana:", bananaDelete)
    //fmt.Println("Orange:", orangeDelete)
    ////fmt.Println("New:", newWord)
    ////fmt.Println("Candle:", candle)
    ////fmt.Println("King:", king)
    ////fmt.Println("alphabet:", alphabet)
}

func TestCharles1InsertAndDelete() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aa", "banana")
    mpt.Insert("ap", "orange")
    mpt.Insert("b", "new")

    fmt.Println("/nGet Values")
    apple := mpt.Get("p")
    fmt.Println("Get:", apple)
    banana := mpt.Get("aa")
    fmt.Println("Get:", banana)

    fmt.Println("\nDeleting values")
    deleteIncorrect, err1 := mpt.Delete("c")
    fmt.Println(deleteIncorrect, err1)
    deleteNew, err2 := mpt.Delete("b")
    fmt.Println(deleteNew, err2)
    deleteOrange, err3 := mpt.Delete("ap")
    fmt.Println(deleteOrange, err3)
    deleteBanana, err4 := mpt.Delete("aa")
    fmt.Println(deleteBanana, err4)
    deleteApple, err5 := mpt.Delete("p")
    fmt.Println(deleteApple, err5)

    fmt.Println("\nGet Values")
    orange := mpt.Get("ap")
    fmt.Println("Get:", orange)
    newWord := mpt.Get("b")
    fmt.Println("Get:", newWord)

    //fmt.Println(mpt.db[mpt.root])
    //fmt.Println(mpt.db["HashStart_42a990655bffe188c9823a2f914641a32dcbb1b28e8586bd29af291db7dcd4e8_HashEnd"])
    //fmt.Println(mpt.db["HashStart_2fdf6310583baee09f440c41749fd03f2542d1bcb9cf24b78045caf56d77758c_HashEnd"])
    //fmt.Println(mpt.db["HashStart_23ca1c3a6072294f27e66941c8cd3531b5d5ed16d7bf05883b7e30fbf32cb59b_HashEnd"])
    //fmt.Println(mpt.db["HashStart_afb91e31b95ddfc4cc5b179ee86e4ed9d5d5681b0feeb15b21f9564c03749d01_HashEnd"])
    //fmt.Println(mpt.db["HashStart_3c255775632b05b1194107f9ac8b40f9d498720c70536a3f90be2686b31d1b67_HashEnd"])
}

func TestCharles2InsertAndDelete() {
    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aa", "banana")
    mpt.Insert("ap", "orange")
    mpt.Insert("ba", "new")

    fmt.Println("\nDeleting values")
    deleteIncorrect, err1 := mpt.Delete("c")
    fmt.Println(deleteIncorrect, err1)
    deleteNew, err2 := mpt.Delete("ba")
    fmt.Println(deleteNew, err2)

    banana := mpt.Get("aa")
    fmt.Println("Get:", banana)

    aaa := mpt.Get("aaa")
    fmt.Println(aaa)
}

func TestCharles3InsertAndDelete() {
    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aaa", "banana")
    mpt.Insert("aap", "orange")
    mpt.Insert("b", "new")

    fmt.Println("\nDeleting values")
    //deleteNew, err1 := mpt.Delete("b")
    //fmt.Println(deleteNew, err1)

    banana := mpt.Get("aaa")
    fmt.Println("Banana:", banana)

    getNew := mpt.Get("b")
    fmt.Println(getNew)
}

func TestCharles4InsertAndDelete() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aa", "apple")
    mpt.Insert("ap", "banana")
    mpt.Insert("bc", "new")

    //fmt.Println("\nDeleting values")
    //deleteC, err1 := mpt.Delete("c")
    //fmt.Println(deleteC, err1)

    deleteNew, err2 := mpt.Delete("bc")
    fmt.Println(deleteNew, err2)

    banana := mpt.Get("ap")
    fmt.Println("Banana:", banana)
}

func TestCharles5InsertAndDelete() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aaaa", "apple")
    mpt.Insert("aaap", "banana")
    mpt.Insert("a", "new")

    fmt.Println("\nDeleting values")
    deleteC, err1 := mpt.Delete("c")
    fmt.Println(deleteC, err1)

    deleteNew, err2 := mpt.Delete("a")
    fmt.Println(deleteNew, err2)

    apple := mpt.Get("aaaa")
    fmt.Println("Banana:", apple)
}


func TestCharles6InsertAndDelete() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aaa", "apple")
    mpt.Insert("aap", "banana")
    mpt.Insert("bc", "new")

    fmt.Println("\nDeleting values")
    deleteC, err1 := mpt.Delete("c")
    fmt.Println(deleteC, err1)

}

func TestLeaf1() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("b", "banana")
    mpt.Insert("a", "new")

    a2 := mpt.Get("a")
    fmt.Println("Get a = new:", a2)

    fmt.Println("\nDeleting values")
    deleteC, err1 := mpt.Delete("c")
    fmt.Println(deleteC, err1)

    deleteNew, err2 := mpt.Delete("a")
    fmt.Println(deleteNew, err2)

}

func TestLeaf2() {

mpt := GetMPTrie()
fmt.Println("Inserting values")
mpt.Insert("a", "apple")
mpt.Insert("b", "banana")
mpt.Insert("ab", "new")

fmt.Println("\nDeleting values")
//deleteC, err1 := mpt.Delete("c")
//fmt.Println(deleteC, err1)

getNew := mpt.Get("ab")
fmt.Println("New:", getNew)

deleteNew, err2 := mpt.Delete("ab")
fmt.Println(deleteNew, err2)

}
