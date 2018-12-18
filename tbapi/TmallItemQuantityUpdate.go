package tbsdk

// TmallItemQuantityUpdateRequest 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
type TmallItemQuantityUpdateRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* item_quantity optional商品库存数；增量编辑方式支持正数、负数 */
    item_quantity int64 `json:"item_quantity";xml:"item_quantity"`
    
    /* options optional商品库存更新时候的可选参数 */
    options UpdateItemQuantityOption `json:"options";xml:"options"`
    
    /* sku_quantities optional更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！ */
    sku_quantities UpdateSkuQuantity `json:"sku_quantities";xml:"sku_quantities"`
    
}

func (req *TmallItemQuantityUpdateRequest) GetAPIName() string {
	return "tmall.item.quantity.update"
}

// TmallItemQuantityUpdateResponse 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
type TmallItemQuantityUpdateResponse struct {
    
    /* quantity_update_result Basic库存更新结果，商品id */
    quantity_update_result string `json:"quantity_update_result";xml:"quantity_update_result"`
    
}
