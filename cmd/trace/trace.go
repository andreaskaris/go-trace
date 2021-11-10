package main

import (
	"flag"
	"log"

	"github.com/andreaskaris/go-trace/trace"
)

var pid = flag.Int("p", 0, "pid to trace")
var follow = flag.Bool("f", false, "follow children")

func main() {
	flag.Parse()
	if *pid == 0 {
		log.Fatalf("Cannot trace pid: %d", *pid)
	}
	err := trace.Strace(*pid, *follow)
	if err != nil {
		log.Fatal(err)
	}
}
