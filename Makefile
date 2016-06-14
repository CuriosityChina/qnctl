.PHONY: all clean install

all: clean
	ln -sf ${PWD}/vendor Godeps/src && GOPATH=${PWD}/Godeps go build

install: all
	sudo cp qnctl /usr/bin/

clean:
	rm -rf qnctl Godeps/src
