build:
	go build -o bin/trace cmd/trace/trace.go

build-test:
	go build -o bin/testcmd cmd/testcmd/testcmd.go

build-ctest:
	gcc -pthread -o bin/testccmd cmd/testccmd/testccmd.c
