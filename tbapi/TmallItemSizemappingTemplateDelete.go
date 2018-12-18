package tbsdk

// TmallItemSizemappingTemplateDeleteRequest 删除天猫商品尺码表模板
type TmallItemSizemappingTemplateDeleteRequest struct {
    
    /* template_id required尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TmallItemSizemappingTemplateDeleteRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.delete"
}

// TmallItemSizemappingTemplateDeleteResponse 删除天猫商品尺码表模板
type TmallItemSizemappingTemplateDeleteResponse struct {
    
    /* template_id Basic尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}
