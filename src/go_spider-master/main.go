package main

import (
	"flag"
	"go_spider-master/core/master"
	"go_spider-master/core/scheduler"
	"go_spider-master/core/sink"
	"go_spider-master/core/dp"
	"github.com/takama/daemon"
	"log"
)

//var omitNewline = flag.Bool("\n", false, "test")
var module = flag.String("M", "master", "set module to startup")

func main() {
	service, err := daemon.New("name", "desc")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	status, err := service.Install()
	if err != nil {
		log.Fatal(status, "\nerror: ", err)
	}

	flag.Parse()


	var s string = ""
	for _,arg := range flag.Args() {
		s += " " + arg
	}
	println(s)

	println("main : " + *module)





	switch *module {
	case "master":
		master_obj := master.Master{}
		master_obj.Run()
	case "scheduler":
		scheduler_obj := scheduler.Schedulerr{}
		scheduler_obj.Run()
	case "dp":
		dp_obj := dp.DP{}
		dp_obj.Run()
	case "sink":
		sink_obj := sink.Sink{}
		sink_obj.Run()
	}

}
