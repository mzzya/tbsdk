package tbsdk

// TmallItemUpdateSchemaGetRequest Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetRequest struct {
    
    /* category_id optional商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* product_id optional商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallItemUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.item.update.schema.get"
}

// TmallItemUpdateSchemaGetResponse Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetResponse struct {
    
    /* update_item_result Basic返回发布商品的规则文档 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}
