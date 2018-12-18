package tbsdk

// TaobaoRegionPriceCancleRequest 取消区域价格
type TaobaoRegionPriceCancleRequest struct {
    
    /* item_id required商品 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sku_id required无sku传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceCancleRequest) GetAPIName() string {
	return "taobao.region.price.cancle"
}

// TaobaoRegionPriceCancleResponse 取消区域价格
type TaobaoRegionPriceCancleResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}
