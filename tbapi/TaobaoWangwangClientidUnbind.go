package tbsdk

// TaobaoWangwangClientidUnbindRequest 解除账号和客户端Id的绑定
type TaobaoWangwangClientidUnbindRequest struct {
    
    /* app_name required应用名 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* client_id required客户端Id */
    client_id string `json:"client_id";xml:"client_id"`
    
}

func (req *TaobaoWangwangClientidUnbindRequest) GetAPIName() string {
	return "taobao.wangwang.clientid.unbind"
}

// TaobaoWangwangClientidUnbindResponse 解除账号和客户端Id的绑定
type TaobaoWangwangClientidUnbindResponse struct {
    
    /* result Basic返回0表示成功， 其他值为错误 */
    result int64 `json:"result";xml:"result"`
    
}
