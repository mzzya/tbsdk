package tbsdk

// TaobaoItemIncrementUpdateSchemaGetRequest 获取增量更新规则，目前开发的字段有title, subTitle, description, wl_description
type TaobaoItemIncrementUpdateSchemaGetRequest struct {
    
    /* category_id optional宝贝类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required宝贝id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* update_fields optional修改字段 */
    update_fields string `json:"update_fields";xml:"update_fields"`
    
}

func (req *TaobaoItemIncrementUpdateSchemaGetRequest) GetAPIName() string {
	return "taobao.item.increment.update.schema.get"
}

// TaobaoItemIncrementUpdateSchemaGetResponse 获取增量更新规则，目前开发的字段有title, subTitle, description, wl_description
type TaobaoItemIncrementUpdateSchemaGetResponse struct {
    
    /* update_rules Basic返回的结果集 */
    update_rules string `json:"update_rules";xml:"update_rules"`
    
}
