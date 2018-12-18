package tbsdk

// TaobaoWlbOrderConsignRequest 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaoWlbOrderConsignRequest struct {
    
    /* wlb_order_code required物流宝订单编号 */
    wlb_order_code string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbOrderConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.consign"
}

// TaobaoWlbOrderConsignResponse 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaoWlbOrderConsignResponse struct {
    
    /* modify_time Basic修改时间 */
    modify_time Date `json:"modify_time";xml:"modify_time"`
    
}
