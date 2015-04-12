package main

import (
	"flag"
	"go_spider-master/core/master"
)

//var omitNewline = flag.Bool("\n", false, "test")
var module = flag.String("M", "master", "set module to startup")

func main() {
	flag.Parse()


	var s string = ""
	for _,arg := range flag.Args() {
		s += " " + arg
	}
	println(s)

	println(*module)





	switch *module {
	case "master":
		master_obj := master.Master{}
		master_obj.Run()
	case "scheduler":
	case "dp":
	case "sink":

	}

}
