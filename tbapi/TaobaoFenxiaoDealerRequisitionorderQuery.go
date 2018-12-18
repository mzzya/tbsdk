package tbsdk

// TaobaoFenxiaoDealerRequisitionorderQueryRequest 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
type TaobaoFenxiaoDealerRequisitionorderQueryRequest struct {
    
    /* dealer_order_ids required经销采购单编号。
多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。 */
    dealer_order_ids int64 `json:"dealer_order_ids";xml:"dealer_order_ids"`
    
    /* fields optional多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.query"
}

// TaobaoFenxiaoDealerRequisitionorderQueryResponse 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
type TaobaoFenxiaoDealerRequisitionorderQueryResponse struct {
    
    /* dealer_orders Object Array经销采购单结果列表 */
    dealer_orders DealerOrder `json:"dealer_orders";xml:"dealer_orders"`
    
}
