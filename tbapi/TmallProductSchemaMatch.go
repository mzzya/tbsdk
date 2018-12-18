package tbsdk

// TmallProductSchemaMatchRequest 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallProductSchemaMatchRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* propvalues required根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。 */
    propvalues string `json:"propvalues";xml:"propvalues"`
    
}

func (req *TmallProductSchemaMatchRequest) GetAPIName() string {
	return "tmall.product.schema.match"
}

// TmallProductSchemaMatchResponse 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallProductSchemaMatchResponse struct {
    
    /* match_result Basic返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。 */
    match_result string `json:"match_result";xml:"match_result"`
    
}
