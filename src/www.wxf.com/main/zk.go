package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	//"os"
	"strings"
	"time"
)

func must(err error) {
	if nil != err {
		panic(err)
	}
}

func connect() *zk.Conn {
	zksStr := "127.0.0.1"//os.Getenv("ZOOKEEPER_SERVERS")
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	must(err)
	fmt.Println("\n")
	fmt.Println(conn)
	fmt.Println("\n\n")
	return conn
}


func main() {
	conn := connect()
	defer conn.Close()

	testPath := "/test3"
	ephemeralPath := "/ephemeralTestPath"

	flags := int32(0)
	ephemeralFlags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	exists, stat, err := conn.Exists(testPath)
	must(err)
	fmt.Printf("exists: %+v %+v\n\n", exists, stat)

	path, err := conn.Create(testPath, []byte("somedata"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n\n", path)

	epath, eerr := conn.Create(ephemeralPath, []byte("someephemeraldata"), ephemeralFlags, acl)
	must(eerr)
	fmt.Printf("ephemeral test path: %+v\n\n", epath)

	data, stat, err := conn.Get(testPath)
	must(err)
	fmt.Printf("get: %+v %+v\n\n", string(data), stat)

	stat, err = conn.Set(testPath, []byte("somenewdata"), stat.Version)
	must(err)
	fmt.Printf("set: %+v\n\n", stat)

//	err = conn.Delete(testPath, -1)
//	must(err)
//	fmt.Printf("delete: %+v ok\n", testPath)

	exists, stat, err = conn.Exists(testPath)
	must(err)
	fmt.Printf("exists: %+v %+v\n\n", exists, stat)
}
