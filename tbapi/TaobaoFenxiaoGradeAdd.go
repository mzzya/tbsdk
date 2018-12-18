package tbsdk

// TaobaoFenxiaoGradeAddRequest 新建等级
type TaobaoFenxiaoGradeAddRequest struct {
    
    /* name required等级名称，等级名称不可重复 */
    name string `json:"name";xml:"name"`
    
}

func (req *TaobaoFenxiaoGradeAddRequest) GetAPIName() string {
	return "taobao.fenxiao.grade.add"
}

// TaobaoFenxiaoGradeAddResponse 新建等级
type TaobaoFenxiaoGradeAddResponse struct {
    
    /* grade_id Basic等级ID */
    grade_id int64 `json:"grade_id";xml:"grade_id"`
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
