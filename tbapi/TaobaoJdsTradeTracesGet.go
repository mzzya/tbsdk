package tbsdk

// TaobaoJdsTradeTracesGetRequest 获取聚石塔数据共享的交易全链路信息
type TaobaoJdsTradeTracesGetRequest struct {
    
    /* return_user_status optional是否返回用户的状态(是否存在) */
    return_user_status bool `json:"return_user_status";xml:"return_user_status"`
    
    /* tid required淘宝的订单tid */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoJdsTradeTracesGetRequest) GetAPIName() string {
	return "taobao.jds.trade.traces.get"
}

// TaobaoJdsTradeTracesGetResponse 获取聚石塔数据共享的交易全链路信息
type TaobaoJdsTradeTracesGetResponse struct {
    
    /* traces Object Array跟踪信息列表 */
    traces TradeTrace `json:"traces";xml:"traces"`
    
    /* user_status Basic在订单全链路系统中的状态(比如是否存在) */
    user_status string `json:"user_status";xml:"user_status"`
    
}
