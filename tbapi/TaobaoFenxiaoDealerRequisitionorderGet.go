package tbsdk

// TaobaoFenxiaoDealerRequisitionorderGetRequest 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
type TaobaoFenxiaoDealerRequisitionorderGetRequest struct {
    
    /* end_date required采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* fields optional多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
    fields string `json:"fields";xml:"fields"`
    
    /* identity required查询者自己在所要查询的采购申请/经销采购单中的身份。
1：供应商。
2：分销商。
注：填写其他值当做错误处理。 */
    identity int64 `json:"identity";xml:"identity"`
    
    /* order_status optional采购申请/经销采购单状态。
0：全部状态。
1：分销商提交申请，待供应商审核。
2：供应商驳回申请，待分销商确认。
3：供应商修改后，待分销商确认。
4：分销商拒绝修改，待供应商再审核。
5：审核通过下单成功，待分销商付款。
7：付款成功，待供应商发货。
8：供应商发货，待分销商收货。
9：分销商收货，交易成功。
10：采购申请/经销采购单关闭。

注：无值按默认值0计，超出状态范围返回错误信息。 */
    order_status int64 `json:"order_status";xml:"order_status"`
    
    /* page_no optional页码（大于0的整数。无值或小于1的值按默认值1计） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_date required采购申请/经销采购单最早修改时间 */
    start_date Date `json:"start_date";xml:"start_date"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderGetRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.get"
}

// TaobaoFenxiaoDealerRequisitionorderGetResponse 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
type TaobaoFenxiaoDealerRequisitionorderGetResponse struct {
    
    /* dealer_orders Object Array采购申请/经销采购单结果列表 */
    dealer_orders DealerOrder `json:"dealer_orders";xml:"dealer_orders"`
    
    /* total_results Basic按查询条件查到的记录总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
