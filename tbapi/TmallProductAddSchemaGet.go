package tbsdk

// TmallProductAddSchemaGetRequest 获取用户发布产品的规则
type TmallProductAddSchemaGetRequest struct {
    
    /* brand_id special品牌ID */
    brand_id int64 `json:"brand_id";xml:"brand_id"`
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TmallProductAddSchemaGetRequest) GetAPIName() string {
	return "tmall.product.add.schema.get"
}

// TmallProductAddSchemaGetResponse 获取用户发布产品的规则
type TmallProductAddSchemaGetResponse struct {
    
    /* add_product_rule Basic返回发布产品的规则文档 */
    add_product_rule string `json:"add_product_rule";xml:"add_product_rule"`
    
}
