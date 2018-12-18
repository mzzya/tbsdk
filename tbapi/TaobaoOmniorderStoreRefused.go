package tbsdk

// TaobaoOmniorderStoreRefusedRequest ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedRequest struct {
    
    /* report_timestamp requiredISV的系统时间 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* sub_order_list required子订单列表 */
    sub_order_list SubOrder `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreRefusedRequest) GetAPIName() string {
	return "taobao.omniorder.store.refused"
}

// TaobaoOmniorderStoreRefusedResponse ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedResponse struct {
    
    /* err_code Basic正常为0,其他表示异常 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}
