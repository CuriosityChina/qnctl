package kodo

import (
	"net/http"

	"qiniupkg.com/api.v7/auth/qbox"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/x/rpc.v7"
)

// ----------------------------------------------------------

type zoneConfig struct {
	IoHost  string
	UpHosts []string
}

var zones = []zoneConfig{
	// z0:
	{
		IoHost: "http://iovip.qbox.me",
		UpHosts: []string{
			"http://up.qiniu.com",
			"http://upload.qiniu.com",
			"-H up.qiniu.com http://183.136.139.16",
		},
	},
	// z1:
	{
		IoHost: "http://iovip-z1.qbox.me",
		UpHosts: []string{
			"http://up-z1.qiniu.com",
			"http://upload-z1.qiniu.com",
			"-H up-z1.qiniu.com http://106.38.227.27",
		},
	},
}

const (
	defaultRsHost  = "http://rs.qbox.me"
	defaultRsfHost = "http://rsf.qbox.me"
)

// ----------------------------------------------------------

type Config struct {
	AccessKey string
	SecretKey string
	RSHost    string
	RSFHost   string
	IoHost    string
	UpHosts   []string
	Transport http.RoundTripper
}

// ----------------------------------------------------------

type Client struct {
	rpc.Client
	mac *qbox.Mac
	Config
}

func New(zone int, cfg *Config) (p *Client) {

	p = new(Client)
	if cfg != nil {
		p.Config = *cfg
	}

	p.mac = qbox.NewMac(p.AccessKey, p.SecretKey)
	p.Client = rpc.Client{qbox.NewClient(p.mac, p.Transport)}

	if p.RSHost == "" {
		p.RSHost = defaultRsHost
	}
	if p.RSFHost == "" {
		p.RSFHost = defaultRsfHost
	}

	if zone < 0 || zone >= len(zones) {
		panic("invalid config: invalid zone")
	}
	if len(p.UpHosts) == 0 {
		p.UpHosts = zones[zone].UpHosts
	}
	if p.IoHost == "" {
		p.IoHost = zones[zone].IoHost
	}
	return
}

// ----------------------------------------------------------

// 设置全局默认的 ACCESS_KEY, SECRET_KEY 变量。
//
func SetMac(accessKey, secretKey string) {

	conf.ACCESS_KEY, conf.SECRET_KEY = accessKey, secretKey
}

// ----------------------------------------------------------

// 设置使用这个SDK的应用程序名。userApp 必须满足 [A-Za-z0-9_\ \-\.]*
//
func SetAppName(userApp string) error {

	return conf.SetAppName(userApp)
}

// ----------------------------------------------------------

