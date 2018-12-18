package tbsdk

// TaobaoOpenuidGetRequest 获取授权账号对应的OpenUid
type TaobaoOpenuidGetRequest struct {
    
}

func (req *TaobaoOpenuidGetRequest) GetAPIName() string {
	return "taobao.openuid.get"
}

// TaobaoOpenuidGetResponse 获取授权账号对应的OpenUid
type TaobaoOpenuidGetResponse struct {
    
    /* open_uid BasicOpenUID */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
}
