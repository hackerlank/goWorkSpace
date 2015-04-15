package dp

import (
	"os"
	"net"
	"fmt"
	"time"
	"go_spider-master/core/common/zk_util"
	"go_spider-master/core/common/config"
	"github.com/bitly/simplejson"
)


type DPStat struct {
	maxid		int
}


type DP struct {
	id			int
	hostname    string
	ipv4        string
	start_ts    int64
}

func (this *DP) init() {
	this.hostname, _ = os.Hostname()
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		//fmt.Println(addr)
		this.ipv4 += fmt.Sprintf("%v", addr) + ", "
	}
	this.ipv4 = string(this.ipv4[0:(len(this.ipv4)-2)])
	this.start_ts = time.Now().Unix()
}

func (this *DP) Run() {
	//init DP struct
	this.init()

	zkUtil := zk_util.Zk_util{}
	zkUtil.Connect()

	DPStat := DPStat{}

	// 注册zookeeper
	maxDpId := 0
	dp_stat := zkUtil.GetDPStat()
	if nil == dp_stat {
		zkUtil.CreateDPStat()
	} else {
		maxDpId, _ = config.GetJsonValueInt(dp_stat, "maxid")
	}
	newDpId := maxDpId+1
	this.id = newDpId
	DPStat.maxid = newDpId

	zkUtil.CreateDP(newDpId)
	zkUtil.CreateDPHeartbeat(newDpId)

	dpJson := this.dp2Json()
	zkUtil.SetDP(newDpId, dpJson)
	zkUtil.UpDPHeartBeat(newDpId, this.start_ts)

	dpStatJson := DPStat.dpStat2Json()
	zkUtil.SetDPStat(dpStatJson)

	// inform scheduler of this dp node with rpc


	fmt.Println(string(dp_stat))

	//主工作循环
	for {
		//update heartbeat time
		curtime := time.Now().Unix()
		zkUtil.UpDPHeartBeat(newDpId, curtime)

		fmt.Println("looping...")

		time.Sleep(time.Second)
	}
}


func (this *DP) dp2Json() []byte {
	js, _ := simplejson.NewJson([]byte("{}"))
	js.Set(config.ZK_CRAWLER_HOSTNAME_KEY, this.hostname)
	js.Set(config.ZK_CRAWLER_IP_KEY, this.ipv4)
	js.Set(config.ZK_CRAWLER_START_KEY, this.start_ts)
	jsonStr, _ := js.MarshalJSON()
	return jsonStr
}

func (this *DPStat) dpStat2Json() []byte {
	js, _ := simplejson.NewJson([]byte("{}"))
	js.Set("maxid", this.maxid)
	jsonStr, _ := js.MarshalJSON()
	return jsonStr
}
