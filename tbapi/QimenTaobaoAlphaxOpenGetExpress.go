package tbsdk

// QimenTaobaoAlphaxOpenGetExpressRequest 调用接口 从isv获取  商家配置物流公司名单
type QimenTaobaoAlphaxOpenGetExpressRequest struct {
    
    /* seller_id required卖家id */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenGetExpressRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.get.express"
}

// QimenTaobaoAlphaxOpenGetExpressResponse 调用接口 从isv获取  商家配置物流公司名单
type QimenTaobaoAlphaxOpenGetExpressResponse struct {
    
    /* result Object返回值 */
    result ResultDO `json:"result";xml:"result"`
    
}
