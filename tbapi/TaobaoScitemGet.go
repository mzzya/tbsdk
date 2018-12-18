package tbsdk

// TaobaoScitemGetRequest 根据id查询商品
type TaobaoScitemGetRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoScitemGetRequest) GetAPIName() string {
	return "taobao.scitem.get"
}

// TaobaoScitemGetResponse 根据id查询商品
type TaobaoScitemGetResponse struct {
    
    /* sc_item Object后端商品 */
    sc_item ScItem `json:"sc_item";xml:"sc_item"`
    
}
