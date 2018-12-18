package tbsdk

// TaobaoTradesSoldIncrementGetRequest 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）
type TaobaoTradesSoldIncrementGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_modified required查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* ext_type optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
    ext_type string `json:"ext_type";xml:"ext_type"`
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0为处方药未审核状态 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* rate_status optional评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评) */
    rate_status string `json:"rate_status";xml:"rate_status"`
    
    /* start_modified required查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
    /* status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 */
    status string `json:"status";xml:"status"`
    
    /* tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
    tag string `json:"tag";xml:"tag"`
    
    /* _type optional交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。 */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldIncrementGetRequest) GetAPIName() string {
	return "taobao.trades.sold.increment.get"
}

// TaobaoTradesSoldIncrementGetResponse 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）
type TaobaoTradesSoldIncrementGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的交易信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
    trades Trade `json:"trades";xml:"trades"`
    
}
