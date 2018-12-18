package tbsdk

// TmallProductUpdateSchemaGetRequest 获取用户更新产品的规则
type TmallProductUpdateSchemaGetRequest struct {
    
    /* product_id required产品编号 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallProductUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.product.update.schema.get"
}

// TmallProductUpdateSchemaGetResponse 获取用户更新产品的规则
type TmallProductUpdateSchemaGetResponse struct {
    
    /* update_product_schema Basic参数产品ID对产品的更新规则 */
    update_product_schema string `json:"update_product_schema";xml:"update_product_schema"`
    
}
