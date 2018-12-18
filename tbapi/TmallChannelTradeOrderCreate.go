package tbsdk

// TmallChannelTradeOrderCreateRequest 创建渠道分销单
type TmallChannelTradeOrderCreateRequest struct {
    
    /* param0 required入参 */
    param0 TopChannelPurchaseOrderCreateParam `json:"param0";xml:"param0"`
    
}

func (req *TmallChannelTradeOrderCreateRequest) GetAPIName() string {
	return "tmall.channel.trade.order.create"
}

// TmallChannelTradeOrderCreateResponse 创建渠道分销单
type TmallChannelTradeOrderCreateResponse struct {
    
    /* main_purchase_order_list Basic Array采购单号 */
    main_purchase_order_list map[string]interface{} `json:"main_purchase_order_list";xml:"main_purchase_order_list"`
    
}
