package main

import "os"

func HelpAndExit(code int) {
	println(`
USAGE
    qnctl -c config.json -b <bucket> <operation> <args>
OR
    export QINIU_AK=... QINIU_SK=... QINIU_BUCKET=...
    qnctl <operation> <args>

CONFIG EXAMPLE
    {
        "access_key": "DliZmM1OTVjZTVkNzkxMGQxOGE4NzJiNmM1ZmFmZ",
        "secret_key": "TIwY2Y4ZjRmOTJkNzRhOTc0YmE4NDkyM2FiZmVhZ"
    }

OPERATION
    add  <key|path> <file>
    rm   <key>
    stat <key>
    ls   <path>
	`)

	os.Exit(code)
}
