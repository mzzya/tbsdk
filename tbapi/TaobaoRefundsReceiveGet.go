package tbsdk

// TaobaoRefundsReceiveGetRequest 查询卖家收到的退款列表
type TaobaoRefundsReceiveGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_modified optional查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* fields required需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数; 默认值:1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_modified optional查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
    /* status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功) */
    status string `json:"status";xml:"status"`
    
    /* _type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a> */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoRefundsReceiveGetRequest) GetAPIName() string {
	return "taobao.refunds.receive.get"
}

// TaobaoRefundsReceiveGetResponse 查询卖家收到的退款列表
type TaobaoRefundsReceiveGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* refunds Object Array搜索到的退款信息列表 */
    refunds Refund `json:"refunds";xml:"refunds"`
    
    /* total_results Basic搜索到的退款信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
