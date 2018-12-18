package tbsdk

// TmallProductSchemaGetRequest 产品信息获取接口schema形式返回
type TmallProductSchemaGetRequest struct {
    
    /* product_id required产品编号 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallProductSchemaGetRequest) GetAPIName() string {
	return "tmall.product.schema.get"
}

// TmallProductSchemaGetResponse 产品信息获取接口schema形式返回
type TmallProductSchemaGetResponse struct {
    
    /* get_product_result Basic产品信息数据。schema形式 */
    get_product_result string `json:"get_product_result";xml:"get_product_result"`
    
}
