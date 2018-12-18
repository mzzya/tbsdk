package tbsdk

// TaobaoItemSchemaAddRequest isv将宝贝信息，组装成特定格式的xml数据作为参数，进行发布商品
type TaobaoItemSchemaAddRequest struct {
    
    /* category_id required当前发布的叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* xml_data required将需要发布的商品数据组装成的xml格式数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TaobaoItemSchemaAddRequest) GetAPIName() string {
	return "taobao.item.schema.add"
}

// TaobaoItemSchemaAddResponse isv将宝贝信息，组装成特定格式的xml数据作为参数，进行发布商品
type TaobaoItemSchemaAddResponse struct {
    
    /* add_result Basic返回的结果 */
    add_result string `json:"add_result";xml:"add_result"`
    
}
