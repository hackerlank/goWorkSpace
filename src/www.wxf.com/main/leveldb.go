package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func must(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	db, err := leveldb.OpenFile("E:/BigData/lvdb", nil)
	must(err)
	defer db.Close()


	err = db.Put([]byte("key"), []byte("valve"), nil)
	must(err)

	data, err := db.Get([]byte("key"), nil)
	must(err)
	fmt.Println(string(data))


}
