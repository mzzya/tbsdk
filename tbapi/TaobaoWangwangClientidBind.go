package tbsdk

// TaobaoWangwangClientidBindRequest 绑定nick和客户端id
type TaobaoWangwangClientidBindRequest struct {
    
    /* app_name required应用名 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* client_id required客户端Id */
    client_id string `json:"client_id";xml:"client_id"`
    
}

func (req *TaobaoWangwangClientidBindRequest) GetAPIName() string {
	return "taobao.wangwang.clientid.bind"
}

// TaobaoWangwangClientidBindResponse 绑定nick和客户端id
type TaobaoWangwangClientidBindResponse struct {
    
    /* result Basic0:表示成功
其它为错误 */
    result int64 `json:"result";xml:"result"`
    
}
