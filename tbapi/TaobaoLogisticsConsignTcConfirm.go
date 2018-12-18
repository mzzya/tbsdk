package tbsdk

// TaobaoLogisticsConsignTcConfirmRequest 下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
type TaobaoLogisticsConsignTcConfirmRequest struct {
    
    /* app_name requiredERP的名称 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* extend_fields optional扩展字段 K:V; */
    extend_fields map[string]interface{} `json:"extend_fields";xml:"extend_fields"`
    
    /* goods_list optional发货的包裹 */
    goods_list ConfirmConsignGoodsDto `json:"goods_list";xml:"goods_list"`
    
    /* out_trade_no required已发货的外部订单号 */
    out_trade_no string `json:"out_trade_no";xml:"out_trade_no"`
    
    /* provider_id required货主id */
    provider_id int64 `json:"provider_id";xml:"provider_id"`
    
    /* seller_id required卖家id */
    seller_id int64 `json:"seller_id";xml:"seller_id"`
    
    /* trade_id required待发货的淘宝主交易号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoLogisticsConsignTcConfirmRequest) GetAPIName() string {
	return "taobao.logistics.consign.tc.confirm"
}

// TaobaoLogisticsConsignTcConfirmResponse 下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
type TaobaoLogisticsConsignTcConfirmResponse struct {
    
    /* order_consign_code Basic菜鸟发货单据 */
    order_consign_code string `json:"order_consign_code";xml:"order_consign_code"`
    
    /* retry Basic是否重试 */
    retry bool `json:"retry";xml:"retry"`
    
    /* trace_id Basic单次调用流程唯一id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}
