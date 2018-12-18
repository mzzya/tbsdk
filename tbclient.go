package tbsdk

import (
	"time"
)

type Client struct {
	userHttps bool
	appKey    string
	appSecret string
	APIAddr   string
	Timeout   time.Duration
}

const (
	APIAddr     = "gw.api.taobao.com/router/rest"
	APIAddrTest = "gw.api.tbsandbox.com/router/rest"
)

func NewClientWithProxy(useHttps bool, appkey, appsecret, apiAddr string) *Client {
	client := new(Client)
	client.userHttps = useHttps
	client.appKey = appkey
	client.appSecret = appsecret
	client.APIAddr = apiAddr
	return client
}

func NewClient(useHttps bool, appkey, appsecret string) *Client {
	return NewClientWithProxy(useHttps, appkey, appsecret, APIAddr)
}

func (cli *Client) DoPost(v interface{}) map[string]interface{} {
	// httpClient := &http.Client{}
	// httpClient
	return nil
}
