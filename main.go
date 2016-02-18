package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

func readConfig(path string) map[string]string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	var result map[string]string
	err = json.Unmarshal(content, &result)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return result
}

func main() {
	accessKey := os.Getenv("QINIU_AK")
	secretKey := os.Getenv("QINIU_SK")
	bucket := os.Getenv("QINIU_BUCKET")

	bucketArg := flag.String("b", "", "")
	configArg := flag.String("c", "", "")
	flag.Parse()

	if *bucketArg != "" {
		bucket = *bucketArg
	}
	if *configArg != "" {
		config := readConfig(*configArg)
		accessKey = config["access_key"]
		secretKey = config["secret_key"]
	}

	if accessKey == "" || secretKey == "" || bucket == "" {
		HelpAndExit(1)
	} else {
		Operate(accessKey, secretKey, bucket)
	}
}
