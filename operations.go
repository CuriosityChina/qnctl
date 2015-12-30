package main

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"qiniupkg.com/api.v7/kodo"
	"strings"
	"time"
)

var client *kodo.Client

func Operate(accessKey, secretKey string) {
	kodo.SetMac(accessKey, secretKey)
	client = kodo.New(0, nil)

	var index = 0
	var nargs = len(os.Args)
	if nargs >= 4 && os.Args[1] == "-c" {
		index = 3
	} else if nargs >= 2 {
		index = 1
	} else {
		helpAndExit(1)
	}

	var op = os.Args[index]
	if op == "ls" && nargs == index+2 {
		Ls(os.Args[index+1])
	} else if op == "add" && nargs == index+3 {
		Add(os.Args[index+1], os.Args[index+2])
	} else if op == "stat" && nargs == index+2 {
		Stat(os.Args[index+1])
	} else if op == "rm" && nargs == index+2 {
		Rm(os.Args[index+1])
	} else {
		helpAndExit(1)
	}
}

func Ls(path string) {
	bucket, ctx := Bucket()
	list, _, err := bucket.List(ctx, path, "", 100)
	if err == nil {
		for i := range list {
			PrintListItem(&list[i])
		}
	} else {
		println(err.Error())
	}
}

func Add(key, file string) {
	paths, err := filepath.Glob(file)
	if err != nil {
		println(err.Error())
	} else {
		for i := range paths {
			fileInfo, _ := os.Stat(paths[i])
			if !fileInfo.IsDir() {
				AddOne(key, paths[i])
				println("")
			}
		}
	}
}

func AddOne(key, file string) {
	if strings.HasSuffix(key, "/") {
		key = key + filepath.Base(file)
	}
	bucket, ctx := Bucket()
	err := bucket.PutFile(ctx, nil, key, file, nil)
	if err != nil {
		println(err.Error())
	} else {
		Stat(key)
	}
}

func Stat(key string) {
	bucket, ctx := Bucket()
	entry, err := bucket.Stat(ctx, key)
	if err == nil {
		PrintEntry(&entry, key)
	} else {
		println(err.Error())
	}
}

func Rm(key string) {
	bucket, ctx := Bucket()
	err := bucket.Delete(ctx, key)
	if err != nil {
		println(err.Error())
	}
}

func PrintEntry(i *kodo.Entry, key string) {
	fmt.Printf("      Key: %s\n", key)
	fmt.Printf("     Hash: %s\n", i.Hash)
	fmt.Printf("     Size: %d\n", i.Fsize)
	fmt.Printf("Mime Type: %s\n", i.MimeType)
	fmt.Printf(" Put Time: %s\n", time.Unix(i.PutTime/1e7, 0).String())
	fmt.Printf(" End User: %s\n", i.EndUser)
}

func PrintListItem(i *kodo.ListItem) {
	ts := time.Unix(i.PutTime/1e7, 0).String()
	fmt.Printf("%10d  %.19s  %s\n", i.Fsize, ts, i.Key)
}

func Bucket() (kodo.Bucket, context.Context) {
	name := os.Getenv("QINIU_BUCKET")
	if name == "" {
		helpAndExit(1)
	}
	return client.Bucket(name), context.Background()
}
