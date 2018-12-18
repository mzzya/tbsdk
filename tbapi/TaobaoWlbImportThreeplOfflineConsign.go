package tbsdk

// TaobaoWlbImportThreeplOfflineConsignRequest 菜鸟认证直邮线下发货
type TaobaoWlbImportThreeplOfflineConsignRequest struct {
    
    /* from_id optional发件人地址库id */
    from_id int64 `json:"from_id";xml:"from_id"`
    
    /* res_code optional资源code */
    res_code string `json:"res_code";xml:"res_code"`
    
    /* res_id optional资源id */
    res_id int64 `json:"res_id";xml:"res_id"`
    
    /* trade_id optional交易单号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
    /* waybill_no optional运单号 */
    waybill_no string `json:"waybill_no";xml:"waybill_no"`
    
}

func (req *TaobaoWlbImportThreeplOfflineConsignRequest) GetAPIName() string {
	return "taobao.wlb.import.threepl.offline.consign"
}

// TaobaoWlbImportThreeplOfflineConsignResponse 菜鸟认证直邮线下发货
type TaobaoWlbImportThreeplOfflineConsignResponse struct {
    
    /* result Objectresult */
    result TopResult `json:"result";xml:"result"`
    
}
