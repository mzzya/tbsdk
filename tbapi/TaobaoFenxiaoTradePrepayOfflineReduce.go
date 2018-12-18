package tbsdk

// TaobaoFenxiaoTradePrepayOfflineReduceRequest 渠道分销供应商上传线下流水预存款（减少）
type TaobaoFenxiaoTradePrepayOfflineReduceRequest struct {
    
    /* offline_reduce_prepay_param required减少流水 */
    offline_reduce_prepay_param TopOfflineReducePrepayDto `json:"offline_reduce_prepay_param";xml:"offline_reduce_prepay_param"`
    
}

func (req *TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetAPIName() string {
	return "taobao.fenxiao.trade.prepay.offline.reduce"
}

// TaobaoFenxiaoTradePrepayOfflineReduceResponse 渠道分销供应商上传线下流水预存款（减少）
type TaobaoFenxiaoTradePrepayOfflineReduceResponse struct {
    
    /* result Objectresult */
    result ResultTopDo `json:"result";xml:"result"`
    
}
