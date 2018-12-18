package tbsdk

// TaobaoOmniorderDtdConsignRequest 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoOmniorderDtdConsignRequest struct {
    
    /* main_order_id required淘宝订单主订单号 */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
    /* store_id optional发货对应的商户中心门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderDtdConsignRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.consign"
}

// TaobaoOmniorderDtdConsignResponse 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoOmniorderDtdConsignResponse struct {
    
    /* message Basic错误信息 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}
