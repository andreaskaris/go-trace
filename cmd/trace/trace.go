package main

import (
	"flag"
	"log"

	"github.com/andreaskaris/go-trace/trace"
)

var pid = flag.Int("pid", 0, "pid to trace")

func main() {
	flag.Parse()
	if *pid == 0 {
		log.Fatalf("Cannot trace pid: %d", *pid)
	}
	err, _ := trace.Strace(*pid)
	//err := trace.StraceAll(*pid)
	if err != nil {
		log.Fatal(err)
	}
}
