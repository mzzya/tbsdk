package tbsdk

// TmallItemSizemappingTemplateGetRequest 获取天猫商品尺码表模板
type TmallItemSizemappingTemplateGetRequest struct {
    
    /* template_id required尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TmallItemSizemappingTemplateGetRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.get"
}

// TmallItemSizemappingTemplateGetResponse 获取天猫商品尺码表模板
type TmallItemSizemappingTemplateGetResponse struct {
    
    /* size_mapping_template Object尺码表模板 */
    size_mapping_template Model `json:"size_mapping_template";xml:"size_mapping_template"`
    
}
