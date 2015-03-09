package main

import (
	"fmt"
	"github.com/seefan/gossdb"
)


func main() {
	pool, err := gossdb.NewPool(&gossdb.Config{
		Host:             "222.73.225.235",
		Port:             1818,
		MinPoolSize:      5,
		MaxPoolSize:      50,
		AcquireIncrement: 5,
	})
	if err != nil {
		panic(err)
		return
	}
	c, err := pool.NewClient()
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()

	c.Set("test","hello world.")
	re, err := c.Get("test")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(re))
	}


	//设置10 秒过期
	c.Set("test1",1225,10)
	//取出数据，并指定类型为 int
	re, err = c.Get("test1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(re.Int(), "is get")
	}
}
