package tbsdk

// TaobaoWlbSubscriptionQueryRequest 查询商家定购的所有服务,可通过入参状态来筛选
type TaobaoWlbSubscriptionQueryRequest struct {
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* status optional状态 
AUDITING 1-待审核; 
CANCEL 2-撤销 ;
CHECKED 3-审核通过 ;
FAILED 4-审核未通过 ;
SYNCHRONIZING 5-同步中;
只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。 */
    status string `json:"status";xml:"status"`
    
}

func (req *TaobaoWlbSubscriptionQueryRequest) GetAPIName() string {
	return "taobao.wlb.subscription.query"
}

// TaobaoWlbSubscriptionQueryResponse 查询商家定购的所有服务,可通过入参状态来筛选
type TaobaoWlbSubscriptionQueryResponse struct {
    
    /* seller_subscription_list Object Array卖家定购的服务列表 */
    seller_subscription_list WlbSellerSubscription `json:"seller_subscription_list";xml:"seller_subscription_list"`
    
    /* total_count Basic结果总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}
