package tbsdk

// TaobaoItemSellerGetRequest 获取单个商品的全部信息
type TaobaoItemSellerGetRequest struct {
    
    /* fields required需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iid required商品数字ID */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemSellerGetRequest) GetAPIName() string {
	return "taobao.item.seller.get"
}

// TaobaoItemSellerGetResponse 获取单个商品的全部信息
type TaobaoItemSellerGetResponse struct {
    
    /* item Object商品详细信息 */
    item Item `json:"item";xml:"item"`
    
}
