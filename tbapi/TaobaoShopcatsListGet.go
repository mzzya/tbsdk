package tbsdk

// TaobaoShopcatsListGetRequest 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoShopcatsListGetRequest struct {
    
    /* fields optional需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoShopcatsListGetRequest) GetAPIName() string {
	return "taobao.shopcats.list.get"
}

// TaobaoShopcatsListGetResponse 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoShopcatsListGetResponse struct {
    
    /* shop_cats Object Array店铺类目列表信息 */
    shop_cats ShopCat `json:"shop_cats";xml:"shop_cats"`
    
}
