.PHONY: all clean install

GOPATH=${PWD}/Godeps

all: clean
	ln -sf ${PWD}/vendor Godeps/src && go build

install: all
	sudo cp qnctl /usr/bin/

clean:
	rm -rf qnctl Godeps/src
