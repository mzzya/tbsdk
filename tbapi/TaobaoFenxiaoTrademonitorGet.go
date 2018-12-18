package tbsdk

// TaobaoFenxiaoTrademonitorGetRequest 仅限供应商调用此接口查询经销商品监控信息
type TaobaoFenxiaoTrademonitorGetRequest struct {
    
    /* distributor_nick optional经销商的淘宝账号 */
    distributor_nick string `json:"distributor_nick";xml:"distributor_nick"`
    
    /* end_created optional结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
    end_created Date `json:"end_created";xml:"end_created"`
    
    /* fields optional多个字段用","分隔。 fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。例如：trade_monitors.item_title表示只返回item_title */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。（大于0的整数。小于1按1计） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。（每页条数不超过50条，超过50或小于0均按50计） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* product_id optional产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* start_created optional起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
    start_created Date `json:"start_created";xml:"start_created"`
    
}

func (req *TaobaoFenxiaoTrademonitorGetRequest) GetAPIName() string {
	return "taobao.fenxiao.trademonitor.get"
}

// TaobaoFenxiaoTrademonitorGetResponse 仅限供应商调用此接口查询经销商品监控信息
type TaobaoFenxiaoTrademonitorGetResponse struct {
    
    /* total_results Basic搜索到的经销商品订单数量 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trade_monitors Object Array经销商品订单监控信息 */
    trade_monitors TradeMonitor `json:"trade_monitors";xml:"trade_monitors"`
    
}
