package tbsdk

// TaobaoFenxiaoTradePrepayOfflineAddRequest 渠道分销供应商上传线下流水预存款（增加）
type TaobaoFenxiaoTradePrepayOfflineAddRequest struct {
    
    /* offline_add_prepay_param required增加流水 */
    offline_add_prepay_param TopOfflineAddPrepayDto `json:"offline_add_prepay_param";xml:"offline_add_prepay_param"`
    
}

func (req *TaobaoFenxiaoTradePrepayOfflineAddRequest) GetAPIName() string {
	return "taobao.fenxiao.trade.prepay.offline.add"
}

// TaobaoFenxiaoTradePrepayOfflineAddResponse 渠道分销供应商上传线下流水预存款（增加）
type TaobaoFenxiaoTradePrepayOfflineAddResponse struct {
    
    /* result Objectresult */
    result ResultTopDo `json:"result";xml:"result"`
    
}
