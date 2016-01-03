.PHONY: all clean

GOPATH=${PWD}/Godeps/_workspace

all:
	go build

clean:
	rm -rf qnctl
