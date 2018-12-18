package tbsdk

// TmallItemSchemaUpdateRequest 天猫根据规则编辑商品
type TmallItemSchemaUpdateRequest struct {
    
    /* category_id optional商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* product_id optional商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* xml_data required根据tmall.item.update.schema.get生成的商品编辑规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaUpdateRequest) GetAPIName() string {
	return "tmall.item.schema.update"
}

// TmallItemSchemaUpdateResponse 天猫根据规则编辑商品
type TmallItemSchemaUpdateResponse struct {
    
    /* gmt_modified Basic商品更新操作成功时间 */
    gmt_modified Date `json:"gmt_modified";xml:"gmt_modified"`
    
    /* update_item_result Basic返回商品发布结果 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}
