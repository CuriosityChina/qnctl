# qnctl ![build status](https://travis-ci.org/wizawu/qnctl.svg)

Built on the official library `qiniupkg.com/api.v7/kodo`, lightweight and friendly.

Supported operations:
+ add
+ rm
+ ls
+ stat
+ sync

### Installation

```shell
make
cp qnctl /usr/bin/      # optional
```

### Usage

#### Set Up Qiniu Keys

There are two ways to do that. Export `QINIU_AK`, `QINIU_SK` and `QINIU_BUCKET`

```shell
export QINIU_AK=abcd...
export QINIU_SK=dcba...
export QINIU_BUCKET=test
```

Or create a JSON file like

```shell
$ cat config.json
{
    "access_key": "abcd...",
    "secret_key": "dcba..."
}

$ qnctl -c config.json -b test ls images
```

#### add

If the first argument after `add` **does not** end with `/`, it would be the exact key. Otherwise, the key to save would be the concatenation of the first argument and the **basename** of the second argument.

```shell
$ qnctl add images/001.png path/to/001.png
$ qnctl add images/ path/to/001.png         # same as previous

# Save to another key
$ qnctl add images/002.png path/to/001.png

# Add multiple
$ qnctl add images/ path/to/*.png

# DON'T DO
$ qnctl add images path/to/001.png
$ qnctl add images path/to/*.png
```

#### rm

Delete one objects.

```shell
$ qnctl rm images/001.png
```

#### ls

List objects that match the **prefix**. Wildcard is not supported.

```
$ qnctl ls images/
    326842  2016-01-27 03:06:16  images/001.png
    140797  2016-01-27 03:06:16  images/002.png
       983  2016-01-27 03:06:16  images/003.png
```

#### stat

Return status of one object.

```
$ qnctl stat images/001.png
      Key: images/001.png
     Hash: FiBrvrIq6fUvc4p0BR9ZfNmILym1
     Size: 326842
Mime Type: image/png
 Put Time: 2016-01-27 03:06:16 +0800 CST
 End User: 
```

### sync

Refresh CDN cache.

```
$ qnctl sync http://your-domain.qiniudns.com/key1 http://your-domain.qiniudns.com/key2
```
