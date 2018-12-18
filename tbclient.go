package tbsdk

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	appKey     string
	appSecret  string
	APIAddr    string
	Timeout    time.Duration
	Formart    string
	SignMethod string
}

type BaseRequest interface {
	GetAPIName() string
	GetParams() map[string]interface{}
}

const (
	APIAddr         = "http://gw.api.taobao.com/router/rest"
	APIAddrTest     = "http://gw.api.tbsandbox.com/router/rest"
	Formart_Json    = "json"
	Formart_XML     = "xml"
	SignMethod_HMAC = "hmac"
	SignMethod_MD5  = "md5"
)

func NewClientWithProxy(appkey, appsecret, apiAddr string) *Client {
	client := new(Client)
	client.appKey = appkey
	client.appSecret = appsecret
	client.APIAddr = apiAddr
	client.Timeout = 30 * time.Second
	client.Formart = Formart_Json
	client.SignMethod = SignMethod_MD5
	return client
}

func NewClient(appkey, appsecret string) *Client {
	return NewClientWithProxy(appkey, appsecret, APIAddr)
}

func (cli *Client) DoPost(req BaseRequest, session string) ([]byte, error) {
	tr := &http.Transport{
		DisableCompression: true,
	}
	httpCli := http.Client{Transport: tr, Timeout: cli.Timeout}

	httpReq, err := http.NewRequest("POST", cli.APIAddr, strings.NewReader("name=cjb"))
	if err != nil {
		return nil, err
	}
	var param = make(map[string]string, 10)
	param["method"] = req.GetAPIName()
	param["app_key"] = cli.appKey
	param["session"] = session
	param["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	param["format"] = cli.Formart
	param["v"] = "2.0"
	// partner_id	String	否	合作伙伴身份标识。
	// target_app_key	String	否	被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效。
	if cli.Formart == Formart_Json {
		param["simplify"] = "true"
	}

	param["sign_method"] = cli.SignMethod
	param["sign"] = ""
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("Cookie", "name=anny")

	resp, err := httpCli.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func SignStringMap(params map[string]string, appSecret string, signMethod string) string {
	var keys = make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString(params[key])
	}
	return SignString(sb.String(), appSecret, signMethod)
}
func SignString(params string, appSecret string, signMethod string) string {
	if signMethod == SignMethod_MD5 {
		h := md5.New()
		h.Write([]byte(appSecret))
		h.Write([]byte(params))
		h.Write([]byte(appSecret))
		return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	} else {
		h := hmac.New(md5.New, []byte(appSecret))
		h.Write([]byte(params))
		return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	}
}

func GetResonseString(bytes []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func GetResonseObject(bytes []byte, err error, v interface{}) error {
	if err != nil {
		return err
	}
	firstStrIndex := 0
	for _, bt := range bytes {
		if bt != ' ' && firstStrIndex == 0 {
			if bt == '{' {
				err = json.Unmarshal(bytes, v)
			} else {
				err = xml.Unmarshal(bytes, v)
			}
			return err
		}
	}
	return nil
}
