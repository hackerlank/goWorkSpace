package master

import (
	"go_spider-master/core/common/zk_util"
	"go_spider-master/core/common/config"
	"fmt"
	"time"
	"os"
	"net"
	"github.com/bitly/simplejson"
	"strconv"
)

type Master struct {
	hostname		string
	ipv4			string
	start_ts		int64
	//heartbeat_ts	int64
}

func (this *Master) init() {
	this.hostname, _ = os.Hostname()
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		//fmt.Println(addr)
		this.ipv4 += fmt.Sprintf("%v", addr) + ", "
	}
	this.ipv4 = string(this.ipv4[0:(len(this.ipv4)-2)])
	this.start_ts = time.Now().Unix()
	//this.heartbeat_ts = 0
}


func (this *Master) Run() {
	//init Master struct
	this.init()

	zkUtil := zk_util.Zk_util{}
	zkUtil.Connect()


	// 注册zookeeper
	master_data := zkUtil.GetMaster()
	zkUtil.CreateMasterHeartbeat()
	if nil == master_data {
		zkUtil.CreateMaster()
		masterJson := this.master2Json()
		zkUtil.SetMaster(masterJson)
		zkUtil.UpMasterHeartBeat(this.start_ts)
	} else {
		lastHeartbeat := zkUtil.GetMasterHeartbeat()
		lastHeartbeat64, _ := strconv.ParseInt(string(lastHeartbeat), 10, 64)

		// if heartbeat is overdue
		curtime := time.Now().Unix()
		if curtime - lastHeartbeat64 > config.ZK_CRAWLER_MASTER_HEARTBEAT {
			fmt.Println("master heartbeat overdue! Continue startup process...")
			//update heartbeat & start time
			this.start_ts = curtime
			//this.heartbeat_ts = curtime
			masterJson := this.master2Json()
			zkUtil.SetMaster(masterJson)
			zkUtil.UpMasterHeartBeat(curtime)
		} else {							// else if a master is already running
			fmt.Println("master is already running! Now exiting...")
			os.Exit(1)
		}
	}


	fmt.Println(string(master_data))

	//检测各个服务器节点心跳信息
	for {
		//update heartbeat time
		curtime := time.Now().Unix()
		zkUtil.UpMasterHeartBeat(curtime)

		fmt.Println("looping...")
		// 1. scheduler
		lastSchedulerHeartbeat := zkUtil.GetSchedulerHeartbeat()
		lastSchedulerHeartbeat64, _ := strconv.ParseInt(string(lastSchedulerHeartbeat), 10, 64)
		if curtime - lastSchedulerHeartbeat64 > config.ZK_CRAWLER_SCHEDULER_HEARTBEAT {
			fmt.Println("Warning! Scheduler node's heartbeat overdued!")
		}
		// 2. dp
		// get dp node list from scheduler dp cache with rpc
		dpNodeIdList := []int{1,2,3,4,5}
		for dpNodeId := range dpNodeIdList {
			lastDpHeartbeat := zkUtil.GetDPHeartbeat(dpNodeId)
			lastDpHeartbeat64, _ := strconv.ParseInt(string(lastDpHeartbeat), 10, 64)
			if curtime - lastDpHeartbeat64 > config.ZK_CRAWLER_DP_HEARTBEAT {
				fmt.Println("Warning! DP id:" + strconv.Itoa(dpNodeId) + "'s heartbeat overdued!")
			}
		}

		// 3. sink
		lastSinkHeartbeat := zkUtil.GetSinkHeartbeat()
		lastSinkHeartbeat64, _ := strconv.ParseInt(string(lastSinkHeartbeat), 10, 64)
		if curtime - lastSinkHeartbeat64 > config.ZK_CRAWLER_SINK_HEARTBEAT {
			fmt.Println("Warning! Sink node's heartbeat overdued!")
		}


		time.Sleep(time.Second)
	}
}


func (this *Master) master2Json() []byte {
	js, _ := simplejson.NewJson([]byte("{}"))
	js.Set(config.ZK_CRAWLER_HOSTNAME_KEY, this.hostname)
	js.Set(config.ZK_CRAWLER_IP_KEY, this.ipv4)
	js.Set(config.ZK_CRAWLER_START_KEY, this.start_ts)
	//js.Set(config.ZK_CRAWLER_HEARTBEAT_KEY, this.heartbeat_ts)
	jsonStr, _ := js.MarshalJSON()
	return jsonStr
}
