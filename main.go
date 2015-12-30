package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func readConf(path string) map[string]string {
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

func getAuth() (string, string) {
	if len(os.Args) >= 3 && os.Args[1] == "-c" {
		conf := readConf(os.Args[2])
		return conf["access_key"], conf["secret_key"]
	} else {
		return os.Getenv("QINIU_AK"), os.Getenv("QINIU_SK")
	}
}

func main() {
	accessKey, secretKey := getAuth()
	if accessKey == "" || secretKey == "" {
		helpAndExit(1)
	} else {
		Operate(accessKey, secretKey)
	}
}
