package tbsdk

// TaobaoJdsTradesStatisticsDiffRequest 订单全链路状态统计差异比较
type TaobaoJdsTradesStatisticsDiffRequest struct {
    
    /* date required查询的日期，格式如YYYYMMDD的日期对应的数字 */
    date int64 `json:"date";xml:"date"`
    
    /* page_no optional页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* post_status required需要比较的状态 */
    post_status string `json:"post_status";xml:"post_status"`
    
    /* pre_status required需要比较的状态，将会和post_status做比较 */
    pre_status string `json:"pre_status";xml:"pre_status"`
    
}

func (req *TaobaoJdsTradesStatisticsDiffRequest) GetAPIName() string {
	return "taobao.jds.trades.statistics.diff"
}

// TaobaoJdsTradesStatisticsDiffResponse 订单全链路状态统计差异比较
type TaobaoJdsTradesStatisticsDiffResponse struct {
    
    /* tids Basic Arraypre_status比post_status多的tid列表 */
    tids int64 `json:"tids";xml:"tids"`
    
    /* total_results Basic总记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
