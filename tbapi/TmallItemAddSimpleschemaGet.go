package tbsdk

// TmallItemAddSimpleschemaGetRequest 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetRequest struct {
    
}

func (req *TmallItemAddSimpleschemaGetRequest) GetAPIName() string {
	return "tmall.item.add.simpleschema.get"
}

// TmallItemAddSimpleschemaGetResponse 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetResponse struct {
    
    /* result Basic返回发布商品的规则文档 */
    result string `json:"result";xml:"result"`
    
}
