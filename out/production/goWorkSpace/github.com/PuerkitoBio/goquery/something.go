package goquery

import (
	"fmt"
)

func selIndex() {
	sel := DocB().Find("#Main")
	j := sel.Index()
	fmt.Println("Index=%d", j)
}

func Main() {
	selIndex()
}
