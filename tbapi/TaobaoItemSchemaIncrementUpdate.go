package tbsdk

// TaobaoItemSchemaIncrementUpdateRequest 根据schema规则增量修改宝贝信息
type TaobaoItemSchemaIncrementUpdateRequest struct {
    
    /* category_id optional商品类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* parameters required修改字段 */
    parameters string `json:"parameters";xml:"parameters"`
    
}

func (req *TaobaoItemSchemaIncrementUpdateRequest) GetAPIName() string {
	return "taobao.item.schema.increment.update"
}

// TaobaoItemSchemaIncrementUpdateResponse 根据schema规则增量修改宝贝信息
type TaobaoItemSchemaIncrementUpdateResponse struct {
    
    /* item_id Basic商品id */
    item_id string `json:"item_id";xml:"item_id"`
    
}
