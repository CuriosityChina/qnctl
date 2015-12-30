package main

import "os"

func helpAndExit(code int) {
	println(`
USAGE
    QINIU_BUCKET=<bucket> qnctl -c config.json <operation> <args>
OR
    export QINIU_AK=... QINIU_SK=... QINIU_BUCKET=...
    qnctl <operation> <args>

QINIU KEYS
    Use environment variables QINIU_AK, QINIU_SK or config.json like
    {
        access_key: "DliZmM1OTVjZTVkNzkxMGQxOGE4NzJiNmM1ZmFmZ",
        secret_key: "TIwY2Y4ZjRmOTJkNzRhOTc0YmE4NDkyM2FiZmVhZ"
    }

OPERATION
    ls   <path>
    add  <path|key> <file>
    stat <key>
    rm   <key>
    `)

	os.Exit(code)
}
