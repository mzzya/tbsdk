package tbsdk

// TaobaoSubusersGetRequest 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaoSubusersGetRequest struct {
    
    /* user_nick required主账号用户名 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubusersGetRequest) GetAPIName() string {
	return "taobao.subusers.get"
}

// TaobaoSubusersGetResponse 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaoSubusersGetResponse struct {
    
    /* subaccounts Object Array子账号基本信息 */
    subaccounts SubAccountInfo `json:"subaccounts";xml:"subaccounts"`
    
}
