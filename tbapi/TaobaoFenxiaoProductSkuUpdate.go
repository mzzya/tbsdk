package tbsdk

// TaobaoFenxiaoProductSkuUpdateRequest 产品SKU信息更新
type TaobaoFenxiaoProductSkuUpdateRequest struct {
    
    /* agent_cost_price optional代销采购价 */
    agent_cost_price string `json:"agent_cost_price";xml:"agent_cost_price"`
    
    /* dealer_cost_price optional经销采购价 */
    dealer_cost_price string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties requiredsku属性 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity optional产品SKU库存 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* sku_number optional商家编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
    /* standard_price optional采购基准价 */
    standard_price string `json:"standard_price";xml:"standard_price"`
    
}

func (req *TaobaoFenxiaoProductSkuUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.update"
}

// TaobaoFenxiaoProductSkuUpdateResponse 产品SKU信息更新
type TaobaoFenxiaoProductSkuUpdateResponse struct {
    
    /* created Basic操作时间 */
    created Date `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}
