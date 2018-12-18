package tbsdk

// TaobaoQimenOrderstatusUpdateRequest 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
type TaobaoQimenOrderstatusUpdateRequest struct {
    
    /* action_time optional事件发生时间 */
    action_time string `json:"action_time";xml:"action_time"`
    
    /* allocation_code required星盘派单号 */
    allocation_code string `json:"allocation_code";xml:"allocation_code"`
    
    /* operator required操作人 */
    operator string `json:"operator";xml:"operator"`
    
    /* order_codes required淘系子订单号 */
    order_codes int64 `json:"order_codes";xml:"order_codes"`
    
    /* status required订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY */
    status string `json:"status";xml:"status"`
    
    /* store_id required目标门店的商户中心门店编码 */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* _type required业务类型，（枚举值：FAHUO、ZITI） */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoQimenOrderstatusUpdateRequest) GetAPIName() string {
	return "taobao.qimen.orderstatus.update"
}

// TaobaoQimenOrderstatusUpdateResponse 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
type TaobaoQimenOrderstatusUpdateResponse struct {
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
    /* result_code BasicresultCode */
    result_code string `json:"result_code";xml:"result_code"`
    
}
