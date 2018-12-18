package tbsdk

// TaobaoFenxiaoGradesGetRequest 根据供应商ID，查询他的分销商等级信息
type TaobaoFenxiaoGradesGetRequest struct {
    
}

func (req *TaobaoFenxiaoGradesGetRequest) GetAPIName() string {
	return "taobao.fenxiao.grades.get"
}

// TaobaoFenxiaoGradesGetResponse 根据供应商ID，查询他的分销商等级信息
type TaobaoFenxiaoGradesGetResponse struct {
    
    /* fenxiao_grades Object Array分销商等级信息 */
    fenxiao_grades FenxiaoGrade `json:"fenxiao_grades";xml:"fenxiao_grades"`
    
}
