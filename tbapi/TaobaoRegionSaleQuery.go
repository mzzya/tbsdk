package tbsdk

// TaobaoRegionSaleQueryRequest 查询商品销售区域
type TaobaoRegionSaleQueryRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sale_region_level required1:国家;2:省;3: 市;4:区县 */
    sale_region_level int64 `json:"sale_region_level";xml:"sale_region_level"`
    
    /* sku_id required无sku传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionSaleQueryRequest) GetAPIName() string {
	return "taobao.region.sale.query"
}

// TaobaoRegionSaleQueryResponse 查询商品销售区域
type TaobaoRegionSaleQueryResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}
