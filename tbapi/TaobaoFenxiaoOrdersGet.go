package tbsdk

// TaobaoFenxiaoOrdersGetRequest 分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
type TaobaoFenxiaoOrdersGetRequest struct {
    
    /* end_created optional结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
    end_created Date `json:"end_created";xml:"end_created"`
    
    /* fields optional多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。（大于0的整数。默认为1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。（每页条数不超过50条） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* purchase_order_id optional采购单编号或分销流水号，若其它参数没传，则此参数必传。 */
    purchase_order_id int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
    /* start_created optional起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
    start_created Date `json:"start_created";xml:"start_created"`
    
    /* status optional交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) */
    status string `json:"status";xml:"status"`
    
    /* tc_order_id optional采购单下游买家订单id */
    tc_order_id int64 `json:"tc_order_id";xml:"tc_order_id"`
    
    /* time_type optional可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询) */
    time_type string `json:"time_type";xml:"time_type"`
    
}

func (req *TaobaoFenxiaoOrdersGetRequest) GetAPIName() string {
	return "taobao.fenxiao.orders.get"
}

// TaobaoFenxiaoOrdersGetResponse 分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
type TaobaoFenxiaoOrdersGetResponse struct {
    
    /* purchase_orders Object Array采购单及子采购单信息。返回 PurchaseOrder 包含的字段信息。 */
    purchase_orders TopDpOrderDo `json:"purchase_orders";xml:"purchase_orders"`
    
    /* total_results Basic搜索到的采购单记录总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
