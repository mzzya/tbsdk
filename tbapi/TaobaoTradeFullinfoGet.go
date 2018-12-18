package tbsdk

// TaobaoTradeFullinfoGetRequest 获取单笔交易的详细信息
<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值 
<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>4. 获取红包优惠金额可以使用字段 coupon_fee
<br/>5. 请按需获取字段，减少TOP系统的压力
type TaobaoTradeFullinfoGetRequest struct {
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。 */
    fields string `json:"fields";xml:"fields"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeFullinfoGetRequest) GetAPIName() string {
	return "taobao.trade.fullinfo.get"
}

// TaobaoTradeFullinfoGetResponse 获取单笔交易的详细信息
<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值 
<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>4. 获取红包优惠金额可以使用字段 coupon_fee
<br/>5. 请按需获取字段，减少TOP系统的压力
type TaobaoTradeFullinfoGetResponse struct {
    
    /* trade Object交易主订单信息 */
    trade Trade `json:"trade";xml:"trade"`
    
}
