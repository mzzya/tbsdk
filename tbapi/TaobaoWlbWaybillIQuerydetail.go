package tbsdk

// TaobaoWlbWaybillIQuerydetailRequest 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailRequest struct {
    
    /* waybill_detail_query_request required面单查询请求 */
    waybill_detail_query_request WaybillDetailQueryRequest `json:"waybill_detail_query_request";xml:"waybill_detail_query_request"`
    
}

func (req *TaobaoWlbWaybillIQuerydetailRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.querydetail"
}

// TaobaoWlbWaybillIQuerydetailResponse 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailResponse struct {
    
    /* error_codes Basic Array面单查询错误编码 */
    error_codes string `json:"error_codes";xml:"error_codes"`
    
    /* inexistent_waybill_codes Basic Array不存在的面单号 */
    inexistent_waybill_codes string `json:"inexistent_waybill_codes";xml:"inexistent_waybill_codes"`
    
    /* query_success Basic查询是否成功 */
    query_success bool `json:"query_success";xml:"query_success"`
    
    /* waybill_details Object Array面单详情 */
    waybill_details WaybillDetailQueryInfo `json:"waybill_details";xml:"waybill_details"`
    
}
