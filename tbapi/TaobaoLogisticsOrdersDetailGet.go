package tbsdk

// TaobaoLogisticsOrdersDetailGetRequest 查询物流订单的详细信息，涉及用户隐私字段。
type TaobaoLogisticsOrdersDetailGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_created optional创建时间结束.格式:yyyy-MM-dd HH:mm:ss */
    end_created Date `json:"end_created";xml:"end_created"`
    
    /* fields required需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔. */
    fields string `json:"fields";xml:"fields"`
    
    /* freight_payer optional谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
    freight_payer string `json:"freight_payer";xml:"freight_payer"`
    
    /* page_no optional页码.该字段没传 或 值<1 ,则默认page_no为1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数.该字段没传 或 值<1 ，则默认page_size为40 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* receiver_name optional收货人姓名 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* seller_confirm optional卖家是否发货.可选值:yes(是),no(否).如:yes. */
    seller_confirm string `json:"seller_confirm";xml:"seller_confirm"`
    
    /* start_created optional创建时间开始.格式:yyyy-MM-dd HH:mm:ss */
    start_created Date `json:"start_created";xml:"start_created"`
    
    /* status optional物流状态.可查看数据结构 Shipping 中的status字段. */
    status string `json:"status";xml:"status"`
    
    /* tid optional交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息. */
    tid int64 `json:"tid";xml:"tid"`
    
    /* _type optional物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoLogisticsOrdersDetailGetRequest) GetAPIName() string {
	return "taobao.logistics.orders.detail.get"
}

// TaobaoLogisticsOrdersDetailGetResponse 查询物流订单的详细信息，涉及用户隐私字段。
type TaobaoLogisticsOrdersDetailGetResponse struct {
    
    /* shippings Object Array获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息. */
    shippings Shipping `json:"shippings";xml:"shippings"`
    
    /* total_results Basic搜索到的物流订单列表总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
