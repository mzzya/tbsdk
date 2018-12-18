package tbsdk

// TaobaoItemUpdateSchemaGetRequest 在新的编辑模式下，isv需要先获取一份编辑商品的规则和数据，然后根据规则传入数据。该接口提供编辑规则查询功能
type TaobaoItemUpdateSchemaGetRequest struct {
    
    /* category_id optional商品类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoItemUpdateSchemaGetRequest) GetAPIName() string {
	return "taobao.item.update.schema.get"
}

// TaobaoItemUpdateSchemaGetResponse 在新的编辑模式下，isv需要先获取一份编辑商品的规则和数据，然后根据规则传入数据。该接口提供编辑规则查询功能
type TaobaoItemUpdateSchemaGetResponse struct {
    
    /* update_rules Basic返回的结果集 */
    update_rules string `json:"update_rules";xml:"update_rules"`
    
}
