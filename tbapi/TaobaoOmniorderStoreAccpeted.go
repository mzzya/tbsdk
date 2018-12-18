package tbsdk

// TaobaoOmniorderStoreAccpetedRequest ISV Pos端门店接单，通知星盘
type TaobaoOmniorderStoreAccpetedRequest struct {
    
    /* report_timestamp requiredISV系统上报时间 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* sub_order_list required子订单列表 */
    sub_order_list StoreAcceptedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreAccpetedRequest) GetAPIName() string {
	return "taobao.omniorder.store.accpeted"
}

// TaobaoOmniorderStoreAccpetedResponse ISV Pos端门店接单，通知星盘
type TaobaoOmniorderStoreAccpetedResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic错误内容 */
    message string `json:"message";xml:"message"`
    
}
