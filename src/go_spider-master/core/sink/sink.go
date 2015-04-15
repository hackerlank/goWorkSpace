package sink

import (
	"os"
	"net"
	"fmt"
	"time"
	"go_spider-master/core/common/zk_util"
	"strconv"
	"go_spider-master/core/common/config"
	"github.com/bitly/simplejson"
)

type Sink struct {
	hostname    string
	ipv4        string
	start_ts    int64
}

func (this *Sink) init() {
	this.hostname, _ = os.Hostname()
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		//fmt.Println(addr)
		this.ipv4 += fmt.Sprintf("%v", addr) + ", "
	}
	this.ipv4 = string(this.ipv4[0:(len(this.ipv4)-2)])
	this.start_ts = time.Now().Unix()
}

func (this *Sink) Run() {
	//init Sink struct
	this.init()

	zkUtil := zk_util.Zk_util{}
	zkUtil.Connect()


	// 注册zookeeper
	sink_data := zkUtil.GetSink()
	zkUtil.CreateSinkHeartbeat()
	if nil == sink_data {
		zkUtil.CreateSink()
		sinkJson := this.sink2Json()
		zkUtil.SetSink(sinkJson)
		zkUtil.UpSinkHeartBeat(this.start_ts)
	} else {
		lastHeartbeat := zkUtil.GetSinkHeartbeat()
		lastHeartbeat64, _ := strconv.ParseInt(string(lastHeartbeat), 10, 64)

		// if heartbeat is overdue
		curtime := time.Now().Unix()
		if curtime - lastHeartbeat64 > config.ZK_CRAWLER_SINK_HEARTBEAT {
			fmt.Println("sink heartbeat overdue! Continue startup process...")
			//update heartbeat & start time
			this.start_ts = curtime
			sinkJson := this.sink2Json()
			zkUtil.SetSink(sinkJson)
			zkUtil.UpSinkHeartBeat(curtime)
		} else {							// else if a sink is already running
			fmt.Println("sink is already running! Now exiting...")
			os.Exit(1)
		}
	}


	fmt.Println(string(sink_data))

	//主工作循环
	for {
		//update heartbeat time
		curtime := time.Now().Unix()
		zkUtil.UpSinkHeartBeat(curtime)

		fmt.Println("looping...")

		time.Sleep(time.Second)
	}
}


func (this *Sink) sink2Json() []byte {
	js, _ := simplejson.NewJson([]byte("{}"))
	js.Set(config.ZK_CRAWLER_HOSTNAME_KEY, this.hostname)
	js.Set(config.ZK_CRAWLER_IP_KEY, this.ipv4)
	js.Set(config.ZK_CRAWLER_START_KEY, this.start_ts)
	jsonStr, _ := js.MarshalJSON()
	return jsonStr
}
