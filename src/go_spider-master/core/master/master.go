package master

import (
	"go_spider-master/core/common/zk_util"
	"fmt"
	"time"
	"os"
	"net"
	"github.com/bitly/simplejson"
)

type Master struct {
	hostname		string
	ipv4			string
	start_ts		int64
	heartbeat_ts	int64
}

func (this *Master) init() {
	this.hostname, _ = os.Hostname()

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr)
		this.ipv4 += fmt.Sprintf("%v", addr) + ", "
	}
	this.ipv4 = string(this.ipv4[0:(len(this.ipv4)-2)])

	this.start_ts = time.Now().Unix()

	this.heartbeat_ts = 0
}


func (this *Master) Run() {
	//init Master struct
	this.init()

	zkUtil := zk_util.Zk_util{};
	zkUtil.Connect();


	// 注册zookeeper
	master_data := zkUtil.GetMaster()
	if nil == master_data {
		zkUtil.CreateMaster()
		masterJson := this.master2Json()
		zkUtil.
	} else {
		// if exists, check last heartbeat

		// if heartbeat is overdue

		// else if a master is already running
	}


	fmt.Println(string(master_data))

	//检测各个服务器节点心跳信息
	for {
		fmt.Println("looping...")
		time.Sleep(time.Second)
	}
}


func (this *Master) master2Json() []byte {
	js, _ := simplejson.NewJson([]byte("{}"))
	js.Set("hostname", this.hostname)
	js.Set("ipv4", this.ipv4)
	js.Set("start_ts", this.start_ts)
	js.Set("heartbeat_ts", this.heartbeat_ts)
	jsonStr, _ := js.MarshalJSON()
	return jsonStr
}
