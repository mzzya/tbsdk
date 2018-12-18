package tbsdk

// TmallItemSchemaAddRequest 天猫TopSchema发布商品。
type TmallItemSchemaAddRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* product_id required发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* xml_data required根据tmall.item.add.schema.get生成的商品发布规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaAddRequest) GetAPIName() string {
	return "tmall.item.schema.add"
}

// TmallItemSchemaAddResponse 天猫TopSchema发布商品。
type TmallItemSchemaAddResponse struct {
    
    /* add_item_result Basic返回商品发布结果 */
    add_item_result string `json:"add_item_result";xml:"add_item_result"`
    
    /* gmt_create Basic发布商品操作成功时间 */
    gmt_create Date `json:"gmt_create";xml:"gmt_create"`
    
}
