package tbsdk

// TaobaoWlbTmsorderQueryRequest 通过物流订单编号分页查询物流信息
type TaobaoWlbTmsorderQueryRequest struct {
    
    /* order_code required物流订单编号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoWlbTmsorderQueryRequest) GetAPIName() string {
	return "taobao.wlb.tmsorder.query"
}

// TaobaoWlbTmsorderQueryResponse 通过物流订单编号分页查询物流信息
type TaobaoWlbTmsorderQueryResponse struct {
    
    /* tms_order_list Object Array物流订单运单信息列表 */
    tms_order_list WlbTmsOrder `json:"tms_order_list";xml:"tms_order_list"`
    
    /* total_count Basic结果总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
