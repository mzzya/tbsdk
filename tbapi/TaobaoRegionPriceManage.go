package tbsdk

// TaobaoRegionPriceManageRequest 编辑区域价格
type TaobaoRegionPriceManageRequest struct {
    
    /* is_full optionaltrue:全量, false:增量 */
    is_full bool `json:"is_full";xml:"is_full"`
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* regional_price_dtos required列表 */
    regional_price_dtos RegionalPriceDto `json:"regional_price_dtos";xml:"regional_price_dtos"`
    
    /* sku_id optional无sku传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceManageRequest) GetAPIName() string {
	return "taobao.region.price.manage"
}

// TaobaoRegionPriceManageResponse 编辑区域价格
type TaobaoRegionPriceManageResponse struct {
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
