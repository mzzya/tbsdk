package tbsdk

// TaobaoItemAddSchemaGetRequest 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能
type TaobaoItemAddSchemaGetRequest struct {
    
    /* category_id required发布宝贝的叶子类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TaobaoItemAddSchemaGetRequest) GetAPIName() string {
	return "taobao.item.add.schema.get"
}

// TaobaoItemAddSchemaGetResponse 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能
type TaobaoItemAddSchemaGetResponse struct {
    
    /* add_rules Basic返回结果的集合 */
    add_rules string `json:"add_rules";xml:"add_rules"`
    
}
