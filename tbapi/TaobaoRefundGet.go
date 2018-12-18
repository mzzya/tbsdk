package tbsdk

// TaobaoRefundGetRequest 获取单笔退款详情
type TaobaoRefundGetRequest struct {
    
    /* fields required需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku */
    fields string `json:"fields";xml:"fields"`
    
    /* refund_id required退款单号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
}

func (req *TaobaoRefundGetRequest) GetAPIName() string {
	return "taobao.refund.get"
}

// TaobaoRefundGetResponse 获取单笔退款详情
type TaobaoRefundGetResponse struct {
    
    /* refund Object退款详情 */
    refund Refund `json:"refund";xml:"refund"`
    
}
