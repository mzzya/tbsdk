package tbsdk

// TaobaoFenxiaoProductcatDeleteRequest 删除产品线
type TaobaoFenxiaoProductcatDeleteRequest struct {
    
    /* product_line_id required产品线ID */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
}

func (req *TaobaoFenxiaoProductcatDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.delete"
}

// TaobaoFenxiaoProductcatDeleteResponse 删除产品线
type TaobaoFenxiaoProductcatDeleteResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
