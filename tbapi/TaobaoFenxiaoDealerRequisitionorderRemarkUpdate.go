package tbsdk

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest 供应商修改经销采购单备注
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
    
    /* dealer_order_id required经销采购单ID */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
    /* supplier_memo optional备注留言，可为空 */
    supplier_memo string `json:"supplier_memo";xml:"supplier_memo"`
    
    /* supplier_memo_flag required旗子的标记，必选。
1-5之间的数字。
非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色 */
    supplier_memo_flag int64 `json:"supplier_memo_flag";xml:"supplier_memo_flag"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse 供应商修改经销采购单备注
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
