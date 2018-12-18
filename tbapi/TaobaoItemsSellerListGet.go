package tbsdk

// TaobaoItemsSellerListGetRequest 批量获取商品详细信息
type TaobaoItemsSellerListGetRequest struct {
    
    /* fields required需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iids required商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。 */
    num_iids string `json:"num_iids";xml:"num_iids"`
    
}

func (req *TaobaoItemsSellerListGetRequest) GetAPIName() string {
	return "taobao.items.seller.list.get"
}

// TaobaoItemsSellerListGetResponse 批量获取商品详细信息
type TaobaoItemsSellerListGetResponse struct {
    
    /* items Object Array商品详细信息列表 */
    items Item `json:"items";xml:"items"`
    
}
