package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"os"
	"path/filepath"
	"qiniupkg.com/api.v7/kodo"
	"strings"
	"time"
)

var client *kodo.Client
var bucketName string

func Operate(accessKey, secretKey, bucket string) {
	kodo.SetMac(accessKey, secretKey)
	client = kodo.New(0, nil)
	bucketName = bucket

	args := flag.Args()
	op := args[0]
	if op == "ls" && len(args) == 2 {
		Ls(args[1])
	} else if op == "add" && len(args) == 3 {
		Add(args[1], args[2])
	} else if op == "rm" && len(args) == 2 {
		Rm(args[1])
	} else if op == "stat" && len(args) == 2 {
		Stat(args[1])
	} else {
		HelpAndExit(1)
	}
}

func Ls(path string) {
	bucket, ctx := Bucket()
	list, _, _, err := bucket.List(ctx, path, "", "", 100000)
	for i := range list {
		PrintListItem(&list[i])
	}
	if err != nil && err != io.EOF {
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
	return client.Bucket(bucketName), context.Background()
}
