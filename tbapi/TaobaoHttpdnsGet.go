package tbsdk

// TaobaoHttpdnsGetRequest 获取TOP DNS配置
type TaobaoHttpdnsGetRequest struct {
    
}

func (req *TaobaoHttpdnsGetRequest) GetAPIName() string {
	return "taobao.httpdns.get"
}

// TaobaoHttpdnsGetResponse 获取TOP DNS配置
type TaobaoHttpdnsGetResponse struct {
    
    /* result BasicHTTP DNS配置信息 */
    result string `json:"result";xml:"result"`
    
}
