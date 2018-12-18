package tbsdk

// TaobaoSellercenterUserPermissionsGetRequest 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
type TaobaoSellercenterUserPermissionsGetRequest struct {
    
    /* nick required用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterUserPermissionsGetRequest) GetAPIName() string {
	return "taobao.sellercenter.user.permissions.get"
}

// TaobaoSellercenterUserPermissionsGetResponse 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
type TaobaoSellercenterUserPermissionsGetResponse struct {
    
    /* permissions Object Array权限列表 */
    permissions Permission `json:"permissions";xml:"permissions"`
    
}
