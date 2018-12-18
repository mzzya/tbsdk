package tbsdk

// TaobaoFenxiaoCooperationAuditRequest 合作授权审批
type TaobaoFenxiaoCooperationAuditRequest struct {
    
    /* audit_result required1:审批通过，审批通过后要加入授权产品线列表；
2:审批拒绝 */
    audit_result int64 `json:"audit_result";xml:"audit_result"`
    
    /* product_line_list_agent optional当审批通过时需要指定授权产品线id列表(代销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
    product_line_list_agent string `json:"product_line_list_agent";xml:"product_line_list_agent"`
    
    /* product_line_list_dealer optional当审批通过时需要指定授权产品线id列表(经销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
    product_line_list_dealer string `json:"product_line_list_dealer";xml:"product_line_list_dealer"`
    
    /* remark required备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* requisition_id required合作申请Id */
    requisition_id int64 `json:"requisition_id";xml:"requisition_id"`
    
}

func (req *TaobaoFenxiaoCooperationAuditRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.audit"
}

// TaobaoFenxiaoCooperationAuditResponse 合作授权审批
type TaobaoFenxiaoCooperationAuditResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
