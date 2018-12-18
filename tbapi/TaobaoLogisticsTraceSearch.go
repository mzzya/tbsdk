package tbsdk

// TaobaoLogisticsTraceSearchRequest 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。
此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
type TaobaoLogisticsTraceSearchRequest struct {
    
    /* is_split optional表明是否是拆单，默认值0，1表示拆单 */
    is_split int64 `json:"is_split";xml:"is_split"`
    
    /* seller_nick required卖家昵称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
    /* sub_tid optional拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。 */
    sub_tid int64 `json:"sub_tid";xml:"sub_tid"`
    
    /* tid required淘宝交易号，请勿传非淘宝交易号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsTraceSearchRequest) GetAPIName() string {
	return "taobao.logistics.trace.search"
}

// TaobaoLogisticsTraceSearchResponse 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。
此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
type TaobaoLogisticsTraceSearchResponse struct {
    
    /* company_name Basic物流公司名称 */
    company_name string `json:"company_name";xml:"company_name"`
    
    /* out_sid Basic运单号 */
    out_sid string `json:"out_sid";xml:"out_sid"`
    
    /* status Basic订单的物流状态（仅支持线上发货online订单，线下发货offline发出后直接变为已签收）
* 等候发送给物流公司
*已提交给物流公司,等待物流公司接单
*已经确认消息接收，等待物流公司接单
*物流公司已接单
*物流公司不接单
*物流公司揽收失败
*物流公司揽收成功
*签收失败
*对方已签收
*对方拒绝签收 */
    status string `json:"status";xml:"status"`
    
    /* tid Basic交易号 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_list Object Array流转信息列表 */
    trace_list TransitStepInfo `json:"trace_list";xml:"trace_list"`
    
}
