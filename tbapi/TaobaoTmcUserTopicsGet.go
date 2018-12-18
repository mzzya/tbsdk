package tbsdk

// TaobaoTmcUserTopicsGetRequest 获取用户开通的topic列表
type TaobaoTmcUserTopicsGetRequest struct {
    
    /* nick optional卖家nick */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoTmcUserTopicsGetRequest) GetAPIName() string {
	return "taobao.tmc.user.topics.get"
}

// TaobaoTmcUserTopicsGetResponse 获取用户开通的topic列表
type TaobaoTmcUserTopicsGetResponse struct {
    
    /* result_code Basic错误码 */
    result_code string `json:"result_code";xml:"result_code"`
    
    /* result_message Basic错误信息 */
    result_message string `json:"result_message";xml:"result_message"`
    
    /* topics Basic Arraytopic列表 */
    topics string `json:"topics";xml:"topics"`
    
}
