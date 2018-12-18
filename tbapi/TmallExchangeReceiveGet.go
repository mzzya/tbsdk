package tbsdk

// TmallExchangeReceiveGetRequest 卖家查询换货列表
type TmallExchangeReceiveGetRequest struct {
    
    /* biz_order_id optional正向订单号 */
    biz_order_id int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    /* buyer_id optional买家id */
    buyer_id int64 `json:"buyer_id";xml:"buyer_id"`
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* dispute_status_array optional换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14) */
    dispute_status_array int64 `json:"dispute_status_array";xml:"dispute_status_array"`
    
    /* end_created_time optional查询申请时间段的结束时间点 */
    end_created_time Date `json:"end_created_time";xml:"end_created_time"`
    
    /* end_gmt_modifed_time optional查询修改时间段的结束时间点 */
    end_gmt_modifed_time Date `json:"end_gmt_modifed_time";xml:"end_gmt_modifed_time"`
    
    /* fields required返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick */
    fields string `json:"fields";xml:"fields"`
    
    /* logistic_no optional快递单号 */
    logistic_no string `json:"logistic_no";xml:"logistic_no"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* refund_id_array optional退款单号ID列表，最多只能输入20个id */
    refund_id_array int64 `json:"refund_id_array";xml:"refund_id_array"`
    
    /* start_created_time optional查询申请时间段的开始时间点 */
    start_created_time Date `json:"start_created_time";xml:"start_created_time"`
    
    /* start_gmt_modified_time optional查询修改时间段的开始时间点 */
    start_gmt_modified_time Date `json:"start_gmt_modified_time";xml:"start_gmt_modified_time"`
    
}

func (req *TmallExchangeReceiveGetRequest) GetAPIName() string {
	return "tmall.exchange.receive.get"
}

// TmallExchangeReceiveGetResponse 卖家查询换货列表
type TmallExchangeReceiveGetResponse struct {
    
    /* error_codes Basic错误码 */
    error_codes string `json:"error_codes";xml:"error_codes"`
    
    /* error_msg Basic错误信息 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* exception Basic所抛出异常 */
    exception map[string]interface{} `json:"exception";xml:"exception"`
    
    /* has_next Basic是否还有下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* page_results Basic当前页的换货单数量 */
    page_results int64 `json:"page_results";xml:"page_results"`
    
    /* results Object Array返回结果 */
    results Exchange `json:"results";xml:"results"`
    
    /* total_results Basic所有符合查询条件的换货单的数量 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
