package tbsdk

// TmallItemOuteridUpdateRequest 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
type TmallItemOuteridUpdateRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* outer_id optional商品维度商家编码，如果不修改可以不传；清空请设置空串 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* sku_outers optional商品SKU更新OuterId时候用的数据 */
    sku_outers UpdateSkuOuterId `json:"sku_outers";xml:"sku_outers"`
    
}

func (req *TmallItemOuteridUpdateRequest) GetAPIName() string {
	return "tmall.item.outerid.update"
}

// TmallItemOuteridUpdateResponse 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
type TmallItemOuteridUpdateResponse struct {
    
    /* outerid_update_result Basic商家编码更新结果 */
    outerid_update_result string `json:"outerid_update_result";xml:"outerid_update_result"`
    
}
