package tbsdk

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
	return cli.DoPostObj(req, session, nil)
}

func (cli *Client) DoPostObj(req BaseRequest, session string, v interface{}) ([]byte, error) {
	tr := &http.Transport{
		DisableCompression: true,
	}
	httpCli := http.Client{Transport: tr, Timeout: cli.Timeout}

	var param = make(map[string]string, 10)
	param["method"] = req.GetAPIName()
	param["app_key"] = cli.appKey
	// if session != "" {
	param["session"] = session
	// }
	param["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	param["format"] = cli.Formart
	param["v"] = "2.0"
	// partner_id	String	否	合作伙伴身份标识。
	// target_app_key	String	否	被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效。
	if cli.Formart == Formart_Json {
		param["simplify"] = "true"
	}
	param["sign_method"] = cli.SignMethod

	var reqParam = req.GetParams()
	for k, v := range reqParam {
		param[k] = GetValueStr(v)
	}

	param["sign"] = SignStringMap(param, cli.appSecret, cli.SignMethod)

	var postData = GetParamStr(param)
	// log.Printf("postData\t%s\n\n", postData)
	httpReq, err := http.NewRequest("POST", cli.APIAddr, strings.NewReader(postData))
	if err != nil {
		return nil, err
	}
	err = httpReq.ParseForm()
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpCli.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteResult, err := ioutil.ReadAll(resp.Body)
	if v == nil || err != nil {
		return byteResult, err
	}
	log.Printf("byteResult:%s\n", byteResult)

	firstStrIndex := 0
	for _, bt := range byteResult {
		if bt != ' ' && firstStrIndex == 0 {
			if bt == '{' {
				err = json.Unmarshal(byteResult, v)
			} else {
				err = xml.Unmarshal(byteResult, v)
			}
			return nil, err
		}
	}
	return byteResult, err
}

func GetValueStr(v interface{}) string {
	if t, ok := v.(time.Time); ok {
		return t.Format("2006-01-02 15:04:05")
	}
	return fmt.Sprint(v)
}

func GetParamStr(params map[string]string) string {
	var sb strings.Builder
	for k, v := range params {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		sb.WriteString("&")
	}
	return sb.String()[0 : sb.Len()-1]
}

func SignStringMap(params map[string]string, appSecret string, signMethod string) string {
	var keys = make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// log.Printf("SignStringKeys\t%+v\n\n", keys)
	var sb strings.Builder
	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString(params[key])
	}
	// log.Printf("SignString\t%s\n\n", sb.String())
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

// func GetResonseString(bytes []byte, err error) (string, error) {
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(bytes), err
// }

// func GetResonseObject(v interface{}, bytes []byte, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	firstStrIndex := 0
// 	for _, bt := range bytes {
// 		if bt != ' ' && firstStrIndex == 0 {
// 			if bt == '{' {
// 				err = json.Unmarshal(bytes, v)
// 			} else {
// 				err = xml.Unmarshal(bytes, v)
// 			}
// 			return err
// 		}
// 	}
// 	return nil
// }
