package tbsdk

// TaobaoLogisticsOnlineCancelRequest 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
type TaobaoLogisticsOnlineCancelRequest struct {
    
    /* tid required淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOnlineCancelRequest) GetAPIName() string {
	return "taobao.logistics.online.cancel"
}

// TaobaoLogisticsOnlineCancelResponse 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
type TaobaoLogisticsOnlineCancelResponse struct {
    
    /* is_success Basic成功与失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* modify_time Basic修改时间 */
    modify_time string `json:"modify_time";xml:"modify_time"`
    
    /* recreated_order_id Basic重新创建物流订单id */
    recreated_order_id string `json:"recreated_order_id";xml:"recreated_order_id"`
    
}
