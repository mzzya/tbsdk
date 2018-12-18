package tbsdk

// TaobaoSellercenterRolemembersGetRequest 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolemembersGetRequest struct {
    
    /* role_id required角色id */
    role_id int64 `json:"role_id";xml:"role_id"`
    
}

func (req *TaobaoSellercenterRolemembersGetRequest) GetAPIName() string {
	return "taobao.sellercenter.rolemembers.get"
}

// TaobaoSellercenterRolemembersGetResponse 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolemembersGetResponse struct {
    
    /* subusers Object Array子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流 */
    subusers SubUserInfo `json:"subusers";xml:"subusers"`
    
}
