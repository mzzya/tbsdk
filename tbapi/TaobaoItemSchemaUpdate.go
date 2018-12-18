package tbsdk

// TaobaoItemSchemaUpdateRequest 新模式下的商品编辑接口，在调用该接口的时候，需要先调用taobao.item.update.rules.get接口获取编辑规则和数据集。然后按照约定的参数传入规则，调用该接口进行编辑商品
type TaobaoItemSchemaUpdateRequest struct {
    
    /* category_id optional如果重新选择类目后，传入重新选择的叶子类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* incremental optional是否增量更新。为true只改xml_data里传入的值。为false时，将根据xml_data的数据对原始宝贝数据进行全量覆盖,这个参数暂时不支持 */
    incremental bool `json:"incremental";xml:"incremental"`
    
    /* item_id required编辑商品的商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* xml_data required编辑商品时候的修改内容 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TaobaoItemSchemaUpdateRequest) GetAPIName() string {
	return "taobao.item.schema.update"
}

// TaobaoItemSchemaUpdateResponse 新模式下的商品编辑接口，在调用该接口的时候，需要先调用taobao.item.update.rules.get接口获取编辑规则和数据集。然后按照约定的参数传入规则，调用该接口进行编辑商品
type TaobaoItemSchemaUpdateResponse struct {
    
    /* update_result Basic编辑商品返回的结果信息，暂时只返回 itemId */
    update_result string `json:"update_result";xml:"update_result"`
    
}
