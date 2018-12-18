package tbsdk

// QimenTaobaoAlphaxOpenJxtDelivergoodsRequest isv 接入催发货接口api
type QimenTaobaoAlphaxOpenJxtDelivergoodsRequest struct {
    
    /* order_id required主订单id */
    order_id string `json:"order_id";xml:"order_id"`
    
    /* seller_id required卖家主账号 */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtDelivergoodsRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.delivergoods"
}

// QimenTaobaoAlphaxOpenJxtDelivergoodsResponse isv 接入催发货接口api
type QimenTaobaoAlphaxOpenJxtDelivergoodsResponse struct {
    
    /* result Object返回值 */
    result ResultDO `json:"result";xml:"result"`
    
}
