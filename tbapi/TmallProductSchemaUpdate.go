package tbsdk

// TmallProductSchemaUpdateRequest 产品更新接口
type TmallProductSchemaUpdateRequest struct {
    
    /* product_id required产品编号 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* xml_data required根据tmall.product.update.schema.get生成的产品更新规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallProductSchemaUpdateRequest) GetAPIName() string {
	return "tmall.product.schema.update"
}

// TmallProductSchemaUpdateResponse 产品更新接口
type TmallProductSchemaUpdateResponse struct {
    
    /* update_product_result Basic产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间 */
    update_product_result string `json:"update_product_result";xml:"update_product_result"`
    
}
