package tbsdk

// TaobaoQimenTradeUserAddRequest 添加奇门订单链路用户
type TaobaoQimenTradeUserAddRequest struct {
    
    /* memo optional商家备注 */
    memo string `json:"memo";xml:"memo"`
    
}

func (req *TaobaoQimenTradeUserAddRequest) GetAPIName() string {
	return "taobao.qimen.trade.user.add"
}

// TaobaoQimenTradeUserAddResponse 添加奇门订单链路用户
type TaobaoQimenTradeUserAddResponse struct {
    
    /* appkey Basicappkey */
    appkey string `json:"appkey";xml:"appkey"`
    
    /* gmt_create Basic创建时间 */
    gmt_create Date `json:"gmt_create";xml:"gmt_create"`
    
    /* memo Basic卖家备注 */
    memo string `json:"memo";xml:"memo"`
    
}
