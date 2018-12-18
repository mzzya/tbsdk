package tbsdk

// AlipayXiaodaiUserPermitRequest 阿里金融为用户开通消息通道接口
type AlipayXiaodaiUserPermitRequest struct {
    
    /* user_id required用户数字ID */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *AlipayXiaodaiUserPermitRequest) GetAPIName() string {
	return "alipay.xiaodai.user.permit"
}

// AlipayXiaodaiUserPermitResponse 阿里金融为用户开通消息通道接口
type AlipayXiaodaiUserPermitResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
