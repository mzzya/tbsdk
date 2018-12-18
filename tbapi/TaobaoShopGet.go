package tbsdk

// TaobaoShopGetRequest 获取卖家店铺的基本信息
type TaobaoShopGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick required卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoShopGetRequest) GetAPIName() string {
	return "taobao.shop.get"
}

// TaobaoShopGetResponse 获取卖家店铺的基本信息
type TaobaoShopGetResponse struct {
    
    /* shop Object店铺信息 */
    shop Shop `json:"shop";xml:"shop"`
    
}
