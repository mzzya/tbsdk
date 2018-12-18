package tbsdk

// TmallItemSchemaIncrementUpdateRequest 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
type TmallItemSchemaIncrementUpdateRequest struct {
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* xml_data required根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaIncrementUpdateRequest) GetAPIName() string {
	return "tmall.item.schema.increment.update"
}

// TmallItemSchemaIncrementUpdateResponse 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
type TmallItemSchemaIncrementUpdateResponse struct {
    
    /* gmt_modified Basic商品更新操作成功时间 */
    gmt_modified Date `json:"gmt_modified";xml:"gmt_modified"`
    
    /* update_item_result Basic返回商品发布结果 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}
