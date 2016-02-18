package main

import "os"

func HelpAndExit(code int) {
	println(`
Usage:
    qnctl -c config.json -b <bucket> <operation> <args>
    # or
    export QINIU_AK=... QINIU_SK=... QINIU_BUCKET=...
    qnctl <operation> <args>

Config example:
    {
        "access_key": "DliZmM1OTVjZTVkNzkxMGQxOGE4NzJiNmM1ZmFmZ",
        "secret_key": "TIwY2Y4ZjRmOTJkNzRhOTc0YmE4NDkyM2FiZmVhZ"
    }

Operations:
    add     <key|path>  <file>
    rm      <key>
    stat    <key>
    ls      <path>
	`)

	os.Exit(code)
}
