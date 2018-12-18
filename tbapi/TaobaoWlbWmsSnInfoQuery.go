package tbsdk

// TaobaoWlbWmsSnInfoQueryRequest 查询仓库作业的各类单据记录的Sn信息
type TaobaoWlbWmsSnInfoQueryRequest struct {
    
    /* order_code required订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_code_type required订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ） */
    order_code_type int64 `json:"order_code_type";xml:"order_code_type"`
    
    /* page_index required页数，默认每页50条 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
}

func (req *TaobaoWlbWmsSnInfoQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.sn.info.query"
}

// TaobaoWlbWmsSnInfoQueryResponse 查询仓库作业的各类单据记录的Sn信息
type TaobaoWlbWmsSnInfoQueryResponse struct {
    
    /* result Object接口返回 */
    result Result `json:"result";xml:"result"`
    
}
