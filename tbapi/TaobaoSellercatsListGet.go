package tbsdk

// TaobaoSellercatsListGetRequest 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaoSellercatsListGetRequest struct {
    
    /* fields optionalfields参数 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick required卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercatsListGetRequest) GetAPIName() string {
	return "taobao.sellercats.list.get"
}

// TaobaoSellercatsListGetResponse 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaoSellercatsListGetResponse struct {
    
    /* seller_cats Object Array卖家自定义类目 */
    seller_cats SellerCat `json:"seller_cats";xml:"seller_cats"`
    
}
