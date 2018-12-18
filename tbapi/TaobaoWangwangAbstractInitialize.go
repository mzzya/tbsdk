package tbsdk

// TaobaoWangwangAbstractInitializeRequest 模糊查询服务初始化，只支持json返回
type TaobaoWangwangAbstractInitializeRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
}

func (req *TaobaoWangwangAbstractInitializeRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.initialize"
}

// TaobaoWangwangAbstractInitializeResponse 模糊查询服务初始化，只支持json返回
type TaobaoWangwangAbstractInitializeResponse struct {
    
    /* error_msg Basic当ret_code=-1时这个变量才有 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1表示成功或失败 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
}
