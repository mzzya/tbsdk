package tbsdk

// TaobaoFenxiaoProductMapAddRequest 创建分销和供应链商品映射关系。
type TaobaoFenxiaoProductMapAddRequest struct {
    
    /* not_check_outer_code optional是否需要校验商家编码，true不校验，false校验。 */
    not_check_outer_code bool `json:"not_check_outer_code";xml:"not_check_outer_code"`
    
    /* product_id required分销产品id。 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* sc_item_id optional后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。 */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* sc_item_ids optional在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。 */
    sc_item_ids string `json:"sc_item_ids";xml:"sc_item_ids"`
    
    /* sku_ids optional分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。 */
    sku_ids string `json:"sku_ids";xml:"sku_ids"`
    
}

func (req *TaobaoFenxiaoProductMapAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.map.add"
}

// TaobaoFenxiaoProductMapAddResponse 创建分销和供应链商品映射关系。
type TaobaoFenxiaoProductMapAddResponse struct {
    
    /* is_success Basic操作结果 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
