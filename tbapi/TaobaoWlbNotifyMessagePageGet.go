package tbsdk

// TaobaoWlbNotifyMessagePageGetRequest 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaoWlbNotifyMessagePageGetRequest struct {
    
    /* end_date optional记录截至时间 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* msg_code optional通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功 */
    msg_code string `json:"msg_code";xml:"msg_code"`
    
    /* page_no optional分页查询页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页查询的每页页数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_date optional记录开始时间 */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* status optional消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM */
    status string `json:"status";xml:"status"`
    
}

func (req *TaobaoWlbNotifyMessagePageGetRequest) GetAPIName() string {
	return "taobao.wlb.notify.message.page.get"
}

// TaobaoWlbNotifyMessagePageGetResponse 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaoWlbNotifyMessagePageGetResponse struct {
    
    /* total_count Basic2000-01-01 00:00:00 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
    /* wlb_messages Object Array通道消息 */
    wlb_messages WlbMessage `json:"wlb_messages";xml:"wlb_messages"`
    
}
