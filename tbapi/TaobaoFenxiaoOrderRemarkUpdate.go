package tbsdk

// TaobaoFenxiaoOrderRemarkUpdateRequest 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateRequest struct {
    
    /* purchase_order_id required采购单编号 */
    purchase_order_id int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
    /* supplier_memo required备注内容(供应商操作) */
    supplier_memo string `json:"supplier_memo";xml:"supplier_memo"`
    
    /* supplier_memo_flag optional旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色 */
    supplier_memo_flag int64 `json:"supplier_memo_flag";xml:"supplier_memo_flag"`
    
}

func (req *TaobaoFenxiaoOrderRemarkUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.order.remark.update"
}

// TaobaoFenxiaoOrderRemarkUpdateResponse 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
