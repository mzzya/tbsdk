package tbsdk

// TaobaoFenxiaoDealerRequisitionorderCloseRequest 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseRequest struct {
    
    /* dealer_order_id required经销采购单编号 */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
    /* reason required关闭原因：
1：长时间无法联系到分销商，取消交易。
2：分销商错误提交申请，取消交易。
3：缺货，近段时间都无法发货。
4：分销商恶意提交申请单。
5：其他原因。 */
    reason int64 `json:"reason";xml:"reason"`
    
    /* reason_detail required关闭详细原因，字数5-200字 */
    reason_detail string `json:"reason_detail";xml:"reason_detail"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}

// TaobaoFenxiaoDealerRequisitionorderCloseResponse 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseResponse struct {
    
    /* is_success Basic操作是否成功。true：成功；false：失败。 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
