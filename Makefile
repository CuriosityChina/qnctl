.PHONY: all clean install

GOPATH=${PWD}/Godeps/_workspace

all:
	go build

install: all
	sudo cp qnctl /usr/bin/

clean:
	rm -rf qnctl
