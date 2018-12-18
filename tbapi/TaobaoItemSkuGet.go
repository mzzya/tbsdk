package tbsdk

// TaobaoItemSkuGetRequest 获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家
type TaobaoItemSkuGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick optional卖家nick(num_iid和nick必传一个)，只传卖家nick时候，该api返回的结果不包含cspu（SKu上的产品规格信息）。 */
    nick string `json:"nick";xml:"nick"`
    
    /* num_iid optional商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* sku_id requiredSku的id。可以通过taobao.item.get得到 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoItemSkuGetRequest) GetAPIName() string {
	return "taobao.item.sku.get"
}

// TaobaoItemSkuGetResponse 获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家
type TaobaoItemSkuGetResponse struct {
    
    /* sku ObjectSku */
    sku Sku `json:"sku";xml:"sku"`
    
}
