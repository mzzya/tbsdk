package tbsdk

// TaobaoWangwangAbstractGetwordlistRequest 获取关键词列表，只支持json返回
type TaobaoWangwangAbstractGetwordlistRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
}

func (req *TaobaoWangwangAbstractGetwordlistRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.getwordlist"
}

// TaobaoWangwangAbstractGetwordlistResponse 获取关键词列表，只支持json返回
type TaobaoWangwangAbstractGetwordlistResponse struct {
    
    /* error_msg Basic例如单词长度太长等，ret_code=-1才有 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
    /* word_lists Object Array关键词列表，ret_code=0才有 */
    word_lists WordList `json:"word_lists";xml:"word_lists"`
    
}
