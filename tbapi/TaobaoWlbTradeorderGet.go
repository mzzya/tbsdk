package tbsdk

// TaobaoWlbTradeorderGetRequest 根据交易类型和交易id查询物流宝订单详情
type TaobaoWlbTradeorderGetRequest struct {
    
    /* sub_trade_id optional子交易号 */
    sub_trade_id string `json:"sub_trade_id";xml:"sub_trade_id"`
    
    /* trade_id required指定交易类型的交易号 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
    /* trade_type required交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易 */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoWlbTradeorderGetRequest) GetAPIName() string {
	return "taobao.wlb.tradeorder.get"
}

// TaobaoWlbTradeorderGetResponse 根据交易类型和交易id查询物流宝订单详情
type TaobaoWlbTradeorderGetResponse struct {
    
    /* wlb_order_list Object Array物流宝订单对象 */
    wlb_order_list WlbOrder `json:"wlb_order_list";xml:"wlb_order_list"`
    
}
