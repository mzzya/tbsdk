package tbsdk

// TmallItemSizemappingTemplatesListRequest 获取所有尺码表模板列表。
type TmallItemSizemappingTemplatesListRequest struct {
    
}

func (req *TmallItemSizemappingTemplatesListRequest) GetAPIName() string {
	return "tmall.item.sizemapping.templates.list"
}

// TmallItemSizemappingTemplatesListResponse 获取所有尺码表模板列表。
type TmallItemSizemappingTemplatesListResponse struct {
    
    /* size_mapping_templates Object Array尺码表模板列表 */
    size_mapping_templates SizeMappingTemplate `json:"size_mapping_templates";xml:"size_mapping_templates"`
    
}
