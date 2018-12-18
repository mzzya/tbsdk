package tbsdk

// TmallItemSimpleschemaAddRequest 天猫简化版schema发布商品。
type TmallItemSimpleschemaAddRequest struct {
    
    /* schema_xml_fields required根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据 */
    schema_xml_fields string `json:"schema_xml_fields";xml:"schema_xml_fields"`
    
}

func (req *TmallItemSimpleschemaAddRequest) GetAPIName() string {
	return "tmall.item.simpleschema.add"
}

// TmallItemSimpleschemaAddResponse 天猫简化版schema发布商品。
type TmallItemSimpleschemaAddResponse struct {
    
    /* gmt_modified Basic商品最后发布时间。 */
    gmt_modified Date `json:"gmt_modified";xml:"gmt_modified"`
    
    /* result Basic发布成功后返回商品ID */
    result string `json:"result";xml:"result"`
    
}
