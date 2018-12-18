package tbsdk

// TaobaoWlbOrderstatusGetRequest 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
type TaobaoWlbOrderstatusGetRequest struct {
    
    /* order_code required物流宝订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbOrderstatusGetRequest) GetAPIName() string {
	return "taobao.wlb.orderstatus.get"
}

// TaobaoWlbOrderstatusGetResponse 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
type TaobaoWlbOrderstatusGetResponse struct {
    
    /* wlb_order_status Object Array订单流转信息状态列表 */
    wlb_order_status WlbProcessStatus `json:"wlb_order_status";xml:"wlb_order_status"`
    
}
