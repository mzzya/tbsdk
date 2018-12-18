package tbsdk

// TaobaoSellercenterRolesGetRequest 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolesGetRequest struct {
    
    /* nick required卖家昵称(只允许查询自己的信息：当前登陆者) */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterRolesGetRequest) GetAPIName() string {
	return "taobao.sellercenter.roles.get"
}

// TaobaoSellercenterRolesGetResponse 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolesGetResponse struct {
    
    /* roles Object Array卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点) */
    roles Role `json:"roles";xml:"roles"`
    
}
