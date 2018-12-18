package tbsdk

// TaobaoJdsTradesStatisticsGetRequest 获取订单数量统计结果
type TaobaoJdsTradesStatisticsGetRequest struct {
    
    /* date required查询的日期，格式如YYYYMMDD的日期对应的数字 */
    date int64 `json:"date";xml:"date"`
    
}

func (req *TaobaoJdsTradesStatisticsGetRequest) GetAPIName() string {
	return "taobao.jds.trades.statistics.get"
}

// TaobaoJdsTradesStatisticsGetResponse 获取订单数量统计结果
type TaobaoJdsTradesStatisticsGetResponse struct {
    
    /* status_infos Object Array订单状态计数值 */
    status_infos TradeStat `json:"status_infos";xml:"status_infos"`
    
}
