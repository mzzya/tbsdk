package tbsdk

// CainiaoBimTradeorderConsignRequest 驱动保税交易订单发货
type CainiaoBimTradeorderConsignRequest struct {
    
    /* res_id optional选择的线路ID非必填字段 */
    res_id string `json:"res_id";xml:"res_id"`
    
    /* store_code optional仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* trade_id required交易单号 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
}

func (req *CainiaoBimTradeorderConsignRequest) GetAPIName() string {
	return "cainiao.bim.tradeorder.consign"
}

// CainiaoBimTradeorderConsignResponse 驱动保税交易订单发货
type CainiaoBimTradeorderConsignResponse struct {
    
    /* lg_order_code Basic菜鸟物流订单号,格式为LP开头 */
    lg_order_code string `json:"lg_order_code";xml:"lg_order_code"`
    
    /* store_order_code Basic菜鸟仓库作业单据号 */
    store_order_code string `json:"store_order_code";xml:"store_order_code"`
    
}
