package tbsdk

// TaobaoTmcUserPermitRequest 为已授权的用户开通消息服务
type TaobaoTmcUserPermitRequest struct {
    
    /* topics optional消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。 */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcUserPermitRequest) GetAPIName() string {
	return "taobao.tmc.user.permit"
}

// TaobaoTmcUserPermitResponse 为已授权的用户开通消息服务
type TaobaoTmcUserPermitResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
