package tbsdk

// TaobaoRegionPriceQueryRequest 区域价格查询
type TaobaoRegionPriceQueryRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* regional_price_dtos optional不传则返回所有设置的区域价格 */
    regional_price_dtos RegionalPriceDto `json:"regional_price_dtos";xml:"regional_price_dtos"`
    
    /* sku_id optional无sku可传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceQueryRequest) GetAPIName() string {
	return "taobao.region.price.query"
}

// TaobaoRegionPriceQueryResponse 区域价格查询
type TaobaoRegionPriceQueryResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}
