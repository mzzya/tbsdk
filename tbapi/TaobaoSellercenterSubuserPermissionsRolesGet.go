package tbsdk

// TaobaoSellercenterSubuserPermissionsRolesGetRequest 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetRequest struct {
    
    /* nick required子账号昵称(子账号标识) */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetAPIName() string {
	return "taobao.sellercenter.subuser.permissions.roles.get"
}

// TaobaoSellercenterSubuserPermissionsRolesGetResponse 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetResponse struct {
    
    /* subuser_permission Object子账号被所拥有的权限 */
    subuser_permission SubUserPermission `json:"subuser_permission";xml:"subuser_permission"`
    
}
