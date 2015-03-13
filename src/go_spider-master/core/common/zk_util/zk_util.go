package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"go_spider-master/core/common/config"
	"go_spider-master/core/common/util"
	"fmt"
	"time"
	"strings"
)

type zk_util struct {
	conn			*zk.Conn
	flag			int32
	ephemeralflag	int32
	acl				[]zk.ACL
}

func (this *zk_util) Connect() *zk_util {
	zkHosts, err := config.GetSpecConfig(config.CONF_FILE_PATH, "zookeeper", "hosts")
	util.Panicable(err)
	zks := strings.Split(zkHosts, ",")
	this.conn, _, err = zk.Connect(zks, time.Second)
	util.Panicable(err)

	this.flag = int32(0)
	this.ephemeralflag = int32(zk.FlagEphemeral)
	this.acl = zk.WorldACL(zk.PermAll)

	return this
}

func (this *zk_util) createPathIfNotExist(dest string, ephemeral bool) *zk_util {
	if dest == "/" {
		util.Printf("reach root node!")
		return nil
	}

	targetIndex := strings.LastIndex(dest, "/")
	parent := dest[0:targetIndex]
//	target := dest[targetIndex+1:len(dest)]

	parentExists, _, _ := this.conn.Exists(parent)
	if parentExists == false {
		util.Printf(fmt.Sprintf("parent: %v NOT exists, step into recursion", parent))
		this.createPathIfNotExist(parent[0:len(parent)], ephemeral)
	}
	destExists, _, _ := this.conn.Exists(dest)
	if destExists == false {
		util.Printf(fmt.Sprintf("cur: %v NOT exists, creating...", dest))

		var flag int32
		if ephemeral == false {flag = this.flag} else {flag = this.ephemeralflag}

		_, err := this.conn.Create(dest, []byte("happy"), flag, this.acl)
		if nil != err {
			panic(err)
		}
	}

	return this
}

func (this *zk_util) CreatePermanentPathIfNotExist(dest string) *zk_util {
	return this.createPathIfNotExist(dest, false)
}

func (this *zk_util) CreateEphemeralPathIfNotExist(dest string) *zk_util {
	return this.createPathIfNotExist(dest, true)
}

func (this *zk_util) SetPathData(dest string, value string) *zk_util {
	_, stat, err := this.conn.Get(dest)

//	destExists, _, _ := this.conn.Exists(dest)
	switch {
	case dest == "/":
		util.Printf("You cannot set value to root node!")
		return nil
	case err != nil:
		util.Printf(fmt.Sprintf("dest: %v MAY NOT exists", dest))
		return nil
	}

	_, err = this.conn.Set(dest, []byte(value), stat.Version)
	util.Panicable(err)

	return this
}

// data format: json
func (this *zk_util) GetPathData(dest string) []byte {
	destExists, _, _ := this.conn.Exists(dest)
	if destExists == false {
		util.Printf(fmt.Sprintf("dest: %v NOT exists", dest))
		return nil
	}

	dataRes, _, err := this.conn.Get(dest)
	util.Panicable(err)

	return dataRes
}

func (this *zk_util) DeletePath(dest string) *zk_util {
	err := this.conn.Delete(dest, -1)
	if err != nil {
		util.Printf(fmt.Sprintf("dest: %v MAY NOT exists", dest))
	}

	return this
}





func (this *zk_util) GetDP() *zk_util {
	exists, stat, err := this.conn.Exists(config.ZK_CRAWLER_DP_PATH)
	util.Panicable(err)
	fmt.Println(exists, stat.Version)
	if false == exists {
		panic("ZK_DP_PATH not exist!")
		return nil
	}
	return this
}



func main() {
	zkUtil := zk_util{}
	zkUtil.Connect()

	zkUtil.CreateEphemeralPathIfNotExist(config.ZK_CRAWLER_ROOT)

	zkUtil.SetPathData(config.ZK_CRAWLER_ROOT, "hahahahahahahahaabc")

	fmt.Println(string(zkUtil.GetPathData(config.ZK_CRAWLER_ROOT)))

	zkUtil.DeletePath(config.ZK_CRAWLER_ROOT)
}
