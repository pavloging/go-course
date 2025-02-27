package main

type UserID int

func main() {
	var idx int = 1
	var uid UserID = 42

	// Особенность Go
	// Если мы создадим type UserID и AdminID и они оба будут int
	// То это будут абсолютно разные типы, не смотря на то, что они являются чилом - int

	myID := UserID(idx)
	println(uid, myID)
}
