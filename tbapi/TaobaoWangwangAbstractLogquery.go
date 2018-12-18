package tbsdk

// TaobaoWangwangAbstractLogqueryRequest 模糊聊天记录查询
type TaobaoWangwangAbstractLogqueryRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
    /* count optional获取记录条数，默认值是1000 */
    count int64 `json:"count";xml:"count"`
    
    /* end_date requiredutc */
    end_date int64 `json:"end_date";xml:"end_date"`
    
    /* from_id required卖家id，有cntaobao前缀 */
    from_id string `json:"from_id";xml:"from_id"`
    
    /* next_key optional设置了这个值，那么聊天记录就从这个点开始查询 */
    next_key string `json:"next_key";xml:"next_key"`
    
    /* start_date requiredutc */
    start_date int64 `json:"start_date";xml:"start_date"`
    
    /* to_id required买家id，有cntaobao前缀 */
    to_id string `json:"to_id";xml:"to_id"`
    
}

func (req *TaobaoWangwangAbstractLogqueryRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.logquery"
}

// TaobaoWangwangAbstractLogqueryResponse 模糊聊天记录查询
type TaobaoWangwangAbstractLogqueryResponse struct {
    
    /* error_msg Basic例如单词长度太长等。
当ret_code不为0时这个值存在。 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* from_id Basic卖家id */
    from_id string `json:"from_id";xml:"from_id"`
    
    /* is_sliced Basic0或1
168的时候用户名设置有误 */
    is_sliced int64 `json:"is_sliced";xml:"is_sliced"`
    
    /* next_key Basic当is_sliced为真时才有这项 */
    next_key string `json:"next_key";xml:"next_key"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息.
为-1时就只有error_msg字段是有效的。 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
    /* to_id Basic买家id */
    to_id string `json:"to_id";xml:"to_id"`
    
    /* url_lists Object Arrayurl列表 */
    url_lists UrlList `json:"url_lists";xml:"url_lists"`
    
    /* word_count_lists Object Array关键词统计列表 */
    word_count_lists WordCountList `json:"word_count_lists";xml:"word_count_lists"`
    
}
