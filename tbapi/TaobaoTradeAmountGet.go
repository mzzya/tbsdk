package tbsdk

// TaobaoTradeAmountGetRequest 卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
type TaobaoTradeAmountGetRequest struct {
    
    /* fields required订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段) */
    fields string `json:"fields";xml:"fields"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeAmountGetRequest) GetAPIName() string {
	return "taobao.trade.amount.get"
}

// TaobaoTradeAmountGetResponse 卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
type TaobaoTradeAmountGetResponse struct {
    
    /* trade_amount Object主订单的财务信息详情 */
    trade_amount TradeAmount `json:"trade_amount";xml:"trade_amount"`
    
}
