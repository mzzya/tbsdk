package tbsdk

// TaobaoWangwangAbstractDeletewordRequest 删除关键词，只支持json返回
type TaobaoWangwangAbstractDeletewordRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
    /* word required关键词，长度大于0 */
    word string `json:"word";xml:"word"`
    
}

func (req *TaobaoWangwangAbstractDeletewordRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.deleteword"
}

// TaobaoWangwangAbstractDeletewordResponse 删除关键词，只支持json返回
type TaobaoWangwangAbstractDeletewordResponse struct {
    
    /* error_msg Basic例如单词长度太长等 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
}
