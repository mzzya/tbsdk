package tbsdk

// TaobaoOmniorderAllocatedinfoSyncRequest ISV分单完成，将分单结果同步给星盘
type TaobaoOmniorderAllocatedinfoSyncRequest struct {
    
    /* message optional分单结果消息, 如果status为AllocateFail, 则表示失败的理由. */
    message string `json:"message";xml:"message"`
    
    /* report_timestamp required1231243213213 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* status required分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail) */
    status string `json:"status";xml:"status"`
    
    /* sub_order_list required门店的分单列表 */
    sub_order_list StoreAllocatedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderAllocatedinfoSyncRequest) GetAPIName() string {
	return "taobao.omniorder.allocatedinfo.sync"
}

// TaobaoOmniorderAllocatedinfoSyncResponse ISV分单完成，将分单结果同步给星盘
type TaobaoOmniorderAllocatedinfoSyncResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic错误内容 */
    message string `json:"message";xml:"message"`
    
}
