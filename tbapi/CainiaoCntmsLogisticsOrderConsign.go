package tbsdk

// CainiaoCntmsLogisticsOrderConsignRequest 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaoCntmsLogisticsOrderConsignRequest struct {
    
    /* content optional配送发货信息 */
    content CnTmsLogisticsOrderConsignContent `json:"content";xml:"content"`
    
}

func (req *CainiaoCntmsLogisticsOrderConsignRequest) GetAPIName() string {
	return "cainiao.cntms.logistics.order.consign"
}

// CainiaoCntmsLogisticsOrderConsignResponse 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaoCntmsLogisticsOrderConsignResponse struct {
    
    /* logistics_order_code Basic物流单号 */
    logistics_order_code string `json:"logistics_order_code";xml:"logistics_order_code"`
    
}
