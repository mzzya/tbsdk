package tbsdk

// TaobaoFenxiaoRequisitionsGetRequest 合作申请查询
type TaobaoFenxiaoRequisitionsGetRequest struct {
    
    /* apply_end optional申请结束时间yyyy-MM-dd */
    apply_end Date `json:"apply_end";xml:"apply_end"`
    
    /* apply_start optional申请开始时间yyyy-MM-dd */
    apply_start Date `json:"apply_start";xml:"apply_start"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* status optional申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期） */
    status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoFenxiaoRequisitionsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.requisitions.get"
}

// TaobaoFenxiaoRequisitionsGetResponse 合作申请查询
type TaobaoFenxiaoRequisitionsGetResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* requisitions Object Array合作申请 */
    requisitions Requisition `json:"requisitions";xml:"requisitions"`
    
    /* total_results Basic结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
