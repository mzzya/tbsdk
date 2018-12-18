package tbsdk

// QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest 在自动化任务核对地址卡片中，透出修改地址选择器，让用户自己选择要修改的地址，然后同步到商家后台ERP系统
type QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest struct {
    
    /* bizOrderId required交易订单ID */
    bizOrderId string `json:"bizOrderId";xml:"bizOrderId"`
    
    /* buyerNick required买家账号名 */
    buyerNick string `json:"buyerNick";xml:"buyerNick"`
    
    /* modifiedAddress required要修改的地址信息 */
    modifiedAddress ModifiedAddress `json:"modifiedAddress";xml:"modifiedAddress"`
    
    /* originalAddress optional订单原始收货地址信息 */
    originalAddress OriginalAddress `json:"originalAddress";xml:"originalAddress"`
    
    /* sellerNick required店铺主账号 */
    sellerNick string `json:"sellerNick";xml:"sellerNick"`
    
}

func (req *QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest) GetAPIName() string {
	return "qimen.taobao.qianniu.cloudkefu.address.self.modify"
}

// QimenTaobaoQianniuCloudkefuAddressSelfModifyResponse 在自动化任务核对地址卡片中，透出修改地址选择器，让用户自己选择要修改的地址，然后同步到商家后台ERP系统
type QimenTaobaoQianniuCloudkefuAddressSelfModifyResponse struct {
    
    /* result Object修改地址返回结果 */
    result ResultDO `json:"result";xml:"result"`
    
}
