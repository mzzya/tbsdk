package tbsdk

// TaobaoVasSubscribeGetRequest 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
type TaobaoVasSubscribeGetRequest struct {
    
    /* article_code required商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码 */
    article_code string `json:"article_code";xml:"article_code"`
    
    /* nick required淘宝会员名 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoVasSubscribeGetRequest) GetAPIName() string {
	return "taobao.vas.subscribe.get"
}

// TaobaoVasSubscribeGetResponse 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
type TaobaoVasSubscribeGetResponse struct {
    
    /* article_user_subscribes Object Array用户订购信息 */
    article_user_subscribes ArticleUserSubscribe `json:"article_user_subscribes";xml:"article_user_subscribes"`
    
}
