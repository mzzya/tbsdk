package tbsdk

// TaobaoFenxiaoProductSkuAddRequest 添加产品SKU信息
type TaobaoFenxiaoProductSkuAddRequest struct {
    
    /* agent_cost_price special代销采购价 */
    agent_cost_price string `json:"agent_cost_price";xml:"agent_cost_price"`
    
    /* dealer_cost_price special经销采购价 */
    dealer_cost_price string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties requiredsku属性 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity optionalsku产品库存，在0到1000000之间，如果不传，则库存为0 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* sku_number optional商家编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
    /* standard_price required采购基准价，最大值1000000000 */
    standard_price string `json:"standard_price";xml:"standard_price"`
    
}

func (req *TaobaoFenxiaoProductSkuAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.add"
}

// TaobaoFenxiaoProductSkuAddResponse 添加产品SKU信息
type TaobaoFenxiaoProductSkuAddResponse struct {
    
    /* created Basic操作时间 */
    created Date `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}
