package tbsdk

// TaobaoQimenShopSynchronizeRequest 店铺同步接口描述
type TaobaoQimenShopSynchronizeRequest struct {
    
    /* request optional请求 */
    request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenShopSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.shop.synchronize"
}

// TaobaoQimenShopSynchronizeResponse 店铺同步接口描述
type TaobaoQimenShopSynchronizeResponse struct {
    
    /* response ObjectResponse */
    response Response `json:"response";xml:"response"`
    
}
