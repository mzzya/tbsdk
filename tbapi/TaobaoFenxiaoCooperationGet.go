package tbsdk

// TaobaoFenxiaoCooperationGetRequest 获取供应商的合作关系信息
type TaobaoFenxiaoCooperationGetRequest struct {
    
    /* end_date optional合作结束时间yyyy-MM-dd HH:mm:ss */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_date optional合作开始时间yyyy-MM-dd HH:mm:ss */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* status optional合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止) */
    status string `json:"status";xml:"status"`
    
    /* trade_type optional分销方式：AGENT(代销) 、DEALER（经销） */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoCooperationGetRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.get"
}

// TaobaoFenxiaoCooperationGetResponse 获取供应商的合作关系信息
type TaobaoFenxiaoCooperationGetResponse struct {
    
    /* cooperations Object Array合作分销关系 */
    cooperations Cooperation `json:"cooperations";xml:"cooperations"`
    
    /* total_results Basic结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
