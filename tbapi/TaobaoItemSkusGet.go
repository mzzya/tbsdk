package tbsdk

// TaobaoItemSkusGetRequest * 获取多个商品下的所有sku
type TaobaoItemSkusGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iids requiredsku所属商品数字id，必选。num_iid个数不能超过40个 */
    num_iids string `json:"num_iids";xml:"num_iids"`
    
}

func (req *TaobaoItemSkusGetRequest) GetAPIName() string {
	return "taobao.item.skus.get"
}

// TaobaoItemSkusGetResponse * 获取多个商品下的所有sku
type TaobaoItemSkusGetResponse struct {
    
    /* skus Object ArraySku列表 */
    skus Sku `json:"skus";xml:"skus"`
    
}
