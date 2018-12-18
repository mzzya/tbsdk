package tbsdk

// TaobaoTradePostageUpdateRequest 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
type TaobaoTradePostageUpdateRequest struct {
    
    /* post_fee required邮费价格(邮费单位是元） */
    post_fee string `json:"post_fee";xml:"post_fee"`
    
    /* tid required主订单编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradePostageUpdateRequest) GetAPIName() string {
	return "taobao.trade.postage.update"
}

// TaobaoTradePostageUpdateResponse 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
type TaobaoTradePostageUpdateResponse struct {
    
    /* trade Object返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment */
    trade Trade `json:"trade";xml:"trade"`
    
}
