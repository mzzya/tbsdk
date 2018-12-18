package tbsdk

// TaobaoTimeGetRequest 获取淘宝系统当前时间
type TaobaoTimeGetRequest struct {
    
}

func (req *TaobaoTimeGetRequest) GetAPIName() string {
	return "taobao.time.get"
}

// TaobaoTimeGetResponse 获取淘宝系统当前时间
type TaobaoTimeGetResponse struct {
    
    /* time Basic淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss */
    time Date `json:"time";xml:"time"`
    
}
