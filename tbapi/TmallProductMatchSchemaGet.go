package tbsdk

// TmallProductMatchSchemaGetRequest ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallProductMatchSchemaGetRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TmallProductMatchSchemaGetRequest) GetAPIName() string {
	return "tmall.product.match.schema.get"
}

// TmallProductMatchSchemaGetResponse ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallProductMatchSchemaGetResponse struct {
    
    /* match_result Basic返回匹配product的规则文档 */
    match_result string `json:"match_result";xml:"match_result"`
    
}
