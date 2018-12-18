package tbsdk

// TaobaoFenxiaoCooperationTerminateRequest 终止与分销商的合作关系
type TaobaoFenxiaoCooperationTerminateRequest struct {
    
    /* cooperate_id required合作编号 */
    cooperate_id int64 `json:"cooperate_id";xml:"cooperate_id"`
    
    /* end_remain_days required终止天数，可选1,2,3,5,7,15，在多少天数内终止 */
    end_remain_days int64 `json:"end_remain_days";xml:"end_remain_days"`
    
    /* end_remark required终止说明（5-2000字） */
    end_remark string `json:"end_remark";xml:"end_remark"`
    
}

func (req *TaobaoFenxiaoCooperationTerminateRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.terminate"
}

// TaobaoFenxiaoCooperationTerminateResponse 终止与分销商的合作关系
type TaobaoFenxiaoCooperationTerminateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
