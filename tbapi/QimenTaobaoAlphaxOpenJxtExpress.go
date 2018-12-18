package tbsdk

// QimenTaobaoAlphaxOpenJxtExpressRequest isv 接入指定快递api
type QimenTaobaoAlphaxOpenJxtExpressRequest struct {
    
    /* cap required快递公司名称 */
    cap string `json:"cap";xml:"cap"`
    
    /* order_id required主订单id */
    order_id string `json:"order_id";xml:"order_id"`
    
    /* seller_id required卖家主账号 */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtExpressRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.express"
}

// QimenTaobaoAlphaxOpenJxtExpressResponse isv 接入指定快递api
type QimenTaobaoAlphaxOpenJxtExpressResponse struct {
    
    /* result Object返回值 */
    result ResultDO `json:"result";xml:"result"`
    
}
