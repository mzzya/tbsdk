package tbsdk

// TmallItemSimpleschemaUpdateRequest 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* schema_xml_fields required编辑商品时提交的xml信息 */
    schema_xml_fields string `json:"schema_xml_fields";xml:"schema_xml_fields"`
    
}

func (req *TmallItemSimpleschemaUpdateRequest) GetAPIName() string {
	return "tmall.item.simpleschema.update"
}

// TmallItemSimpleschemaUpdateResponse 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateResponse struct {
    
    /* gmt_modified Basic编辑商品操作成功时间 */
    gmt_modified Date `json:"gmt_modified";xml:"gmt_modified"`
    
    /* sku_map_json Basicsku与outerId映射信息 */
    sku_map_json string `json:"sku_map_json";xml:"sku_map_json"`
    
    /* update_item_result Basic编辑商品的itemid */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}
