package tbsdk

// TaobaoJdsHluserGetRequest 订单全链路用户信息获取
type TaobaoJdsHluserGetRequest struct {
    
}

func (req *TaobaoJdsHluserGetRequest) GetAPIName() string {
	return "taobao.jds.hluser.get"
}

// TaobaoJdsHluserGetResponse 订单全链路用户信息获取
type TaobaoJdsHluserGetResponse struct {
    
    /* hl_user Object回流用户信息 */
    hl_user HlUserDO `json:"hl_user";xml:"hl_user"`
    
}
