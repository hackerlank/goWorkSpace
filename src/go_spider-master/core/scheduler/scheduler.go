// Copyright 2014 Hu Cong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
package scheduler

import (
    "github.com/hu17889/go_spider/core/common/request"
    "os"
    "net"
    "fmt"
    "time"
    "go_spider-master/core/common/zk_util"
    "strconv"
    "go_spider-master/core/common/config"
    "github.com/bitly/simplejson"
)

type Scheduler interface {
    Push(requ *request.Request)
    Poll() *request.Request
    Count() int
}

type Schedulerr struct {
	hostname    string
	ipv4        string
	start_ts    int64
}

func (this *Schedulerr) init() {
    this.hostname, _ = os.Hostname()
    addrs, _ := net.InterfaceAddrs()
    for _, addr := range addrs {
        //fmt.Println(addr)
        this.ipv4 += fmt.Sprintf("%v", addr) + ", "
    }
    this.ipv4 = string(this.ipv4[0:(len(this.ipv4)-2)])
    this.start_ts = time.Now().Unix()
}

func (this *Schedulerr) Run() {
    //init Scheduler struct
    this.init()

    zkUtil := zk_util.Zk_util{}
    zkUtil.Connect()


    // 注册zookeeper
    scheduler_data := zkUtil.GetScheduler()
    zkUtil.CreateSchedulerHeartbeat()
    if nil == scheduler_data {
        zkUtil.CreateScheduler()
        schedulerJson := this.scheduler2Json()
        zkUtil.SetScheduler(schedulerJson)
        zkUtil.UpSchedulerHeartBeat(this.start_ts)
    } else {
        lastHeartbeat := zkUtil.GetSchedulerHeartbeat()
        lastHeartbeat64, _ := strconv.ParseInt(string(lastHeartbeat), 10, 64)

        // if heartbeat is overdue
        curtime := time.Now().Unix()
        if curtime - lastHeartbeat64 > config.ZK_CRAWLER_SCHEDULER_HEARTBEAT {
            fmt.Println("scheduler heartbeat overdue! Continue startup process...")
            //update heartbeat & start time
            this.start_ts = curtime
            schedulerJson := this.scheduler2Json()
            zkUtil.SetScheduler(schedulerJson)
            zkUtil.UpSchedulerHeartBeat(curtime)
        } else {							// else if a scheduler is already running
            fmt.Println("scheduler is already running! Now exiting...")
            os.Exit(1)
        }
    }


    fmt.Println(string(scheduler_data))

    //主工作循环
    for {
        //update heartbeat time
        curtime := time.Now().Unix()
        zkUtil.UpSchedulerHeartBeat(curtime)

        fmt.Println("looping...")

        time.Sleep(time.Second)
    }
}


func (this *Schedulerr) scheduler2Json() []byte {
    js, _ := simplejson.NewJson([]byte("{}"))
    js.Set(config.ZK_CRAWLER_HOSTNAME_KEY, this.hostname)
    js.Set(config.ZK_CRAWLER_IP_KEY, this.ipv4)
    js.Set(config.ZK_CRAWLER_START_KEY, this.start_ts)
    jsonStr, _ := js.MarshalJSON()
    return jsonStr
}
