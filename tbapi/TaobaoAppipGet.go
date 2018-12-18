package tbsdk

// TaobaoAppipGetRequest 获取ISV发起请求服务器IP
type TaobaoAppipGetRequest struct {
    
}

func (req *TaobaoAppipGetRequest) GetAPIName() string {
	return "taobao.appip.get"
}

// TaobaoAppipGetResponse 获取ISV发起请求服务器IP
type TaobaoAppipGetResponse struct {
    
    /* ip BasicISV发起请求服务器IP */
    ip string `json:"ip";xml:"ip"`
    
}
