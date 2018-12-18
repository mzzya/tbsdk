package tbsdk

// TmallItemSizemappingTemplateUpdateRequest 更新天猫商品尺码表模板
type TmallItemSizemappingTemplateUpdateRequest struct {
    
    /* template_content required尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。 */
    template_content string `json:"template_content";xml:"template_content"`
    
    /* template_id required尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
    /* template_name required尺码表模板名称 */
    template_name string `json:"template_name";xml:"template_name"`
    
}

func (req *TmallItemSizemappingTemplateUpdateRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.update"
}

// TmallItemSizemappingTemplateUpdateResponse 更新天猫商品尺码表模板
type TmallItemSizemappingTemplateUpdateResponse struct {
    
    /* size_mapping_template Object尺码表模板 */
    size_mapping_template SizeMappingTemplateDo `json:"size_mapping_template";xml:"size_mapping_template"`
    
}
