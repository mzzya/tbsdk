package tbsdk

// TaobaoFenxiaoCooperationProductcatAddRequest 追加授权产品线
type TaobaoFenxiaoCooperationProductcatAddRequest struct {
    
    /* cooperate_id required合作关系id */
    cooperate_id int64 `json:"cooperate_id";xml:"cooperate_id"`
    
    /* grade_id optional等级ID（为空则不修改） */
    grade_id int64 `json:"grade_id";xml:"grade_id"`
    
    /* product_line_list required产品线id列表，若有多个，以逗号分隔 */
    product_line_list string `json:"product_line_list";xml:"product_line_list"`
    
}

func (req *TaobaoFenxiaoCooperationProductcatAddRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.productcat.add"
}

// TaobaoFenxiaoCooperationProductcatAddResponse 追加授权产品线
type TaobaoFenxiaoCooperationProductcatAddResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
