package tbsdk

// TaobaoSellercenterRoleAddRequest 给指定的卖家创建新的子账号角色<br/>
如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
---------|---------|code=phone 手机淘宝店铺<br/>
调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
type TaobaoSellercenterRoleAddRequest struct {
    
    /* description optional角色描述 */
    description string `json:"description";xml:"description"`
    
    /* name required角色名 */
    name string `json:"name";xml:"name"`
    
    /* nick required表示卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* permission_codes optional需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。 */
    permission_codes string `json:"permission_codes";xml:"permission_codes"`
    
}

func (req *TaobaoSellercenterRoleAddRequest) GetAPIName() string {
	return "taobao.sellercenter.role.add"
}

// TaobaoSellercenterRoleAddResponse 给指定的卖家创建新的子账号角色<br/>
如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
---------|---------|code=phone 手机淘宝店铺<br/>
调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
type TaobaoSellercenterRoleAddResponse struct {
    
    /* role Object子账号角色 */
    role Role `json:"role";xml:"role"`
    
}
