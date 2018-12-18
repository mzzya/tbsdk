package tbsdk

// TaobaoFenxiaoCooperationUpdateRequest 供应商更新合作的分销商等级
type TaobaoFenxiaoCooperationUpdateRequest struct {
    
    /* distributor_id required分销商ID */
    distributor_id int64 `json:"distributor_id";xml:"distributor_id"`
    
    /* grade_id required等级ID(0代表取消) */
    grade_id int64 `json:"grade_id";xml:"grade_id"`
    
    /* trade_type optional分销方式(新增)： AGENT(代销)、DEALER(经销)(默认为代销) */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoCooperationUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.update"
}

// TaobaoFenxiaoCooperationUpdateResponse 供应商更新合作的分销商等级
type TaobaoFenxiaoCooperationUpdateResponse struct {
    
    /* is_success Basic更新结果成功失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
