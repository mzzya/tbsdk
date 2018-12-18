package tbsdk

// TaobaoWlbOrderdetailDateGetRequest 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
type TaobaoWlbOrderdetailDateGetRequest struct {
    
    /* end_time required创建时间结束 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* page_no optional分页下标 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required创建时间起始 */
    start_time Date `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoWlbOrderdetailDateGetRequest) GetAPIName() string {
	return "taobao.wlb.orderdetail.date.get"
}

// TaobaoWlbOrderdetailDateGetResponse 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
type TaobaoWlbOrderdetailDateGetResponse struct {
    
    /* order_detail_list Object Array物流宝订单，并且包含订单详情 */
    order_detail_list WlbOrderDetail `json:"order_detail_list";xml:"order_detail_list"`
    
    /* total_count Basic总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
