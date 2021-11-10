build:
	go build -o bin/trace cmd/trace/trace.go

build-test:
	go build -o bin/testcmd cmd/testcmd/testcmd.go
