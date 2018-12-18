package tbsdk

// TaobaoTradesSoldIncrementvGetRequest 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。
type TaobaoTradesSoldIncrementvGetRequest struct {
    
    /* end_create required查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
    end_create Date `json:"end_create";xml:"end_create"`
    
    /* ext_type optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
    ext_type string `json:"ext_type";xml:"ext_type"`
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_create required查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
    start_create Date `json:"start_create";xml:"start_create"`
    
    /* status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 */
    status string `json:"status";xml:"status"`
    
    /* tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
    tag string `json:"tag";xml:"tag"`
    
    /* _type optional交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。 */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldIncrementvGetRequest) GetAPIName() string {
	return "taobao.trades.sold.incrementv.get"
}

// TaobaoTradesSoldIncrementvGetResponse 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。
type TaobaoTradesSoldIncrementvGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的交易信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
    trades Trade `json:"trades";xml:"trades"`
    
}
