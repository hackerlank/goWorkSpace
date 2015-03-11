package main

import (
	"fmt"
	"www.wxf.com/test"
	"net/http"
	"io/ioutil"
)

func main() {
	test.Test()
	test.Test2()
	fmt.Printf("testfmt")


	resp, _ := http.Get("http://www.google.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
