package tbsdk

// TmallProductSchemaAddRequest Schema体系发布一个产品
type TmallProductSchemaAddRequest struct {
    
    /* brand_id special品牌ID */
    brand_id int64 `json:"brand_id";xml:"brand_id"`
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* xml_data required根据tmall.product.add.schema.get生成的产品发布规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallProductSchemaAddRequest) GetAPIName() string {
	return "tmall.product.schema.add"
}

// TmallProductSchemaAddResponse Schema体系发布一个产品
type TmallProductSchemaAddResponse struct {
    
    /* add_product_result Basic新发产品结果 */
    add_product_result string `json:"add_product_result";xml:"add_product_result"`
    
}
