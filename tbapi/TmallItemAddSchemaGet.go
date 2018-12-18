package tbsdk

// TmallItemAddSchemaGetRequest 通过类目以及productId获取商品发布规则；
type TmallItemAddSchemaGetRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* isv_init optional正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可 */
    isv_init bool `json:"isv_init";xml:"isv_init"`
    
    /* product_id required商品发布的目标product_id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* _type optional发布商品类型，一口价填“b”，拍卖填"a" */
    _type string `json:"type";xml:"type"`
    
}

func (req *TmallItemAddSchemaGetRequest) GetAPIName() string {
	return "tmall.item.add.schema.get"
}

// TmallItemAddSchemaGetResponse 通过类目以及productId获取商品发布规则；
type TmallItemAddSchemaGetResponse struct {
    
    /* add_item_result Basic返回发布商品的规则文档 */
    add_item_result string `json:"add_item_result";xml:"add_item_result"`
    
}
