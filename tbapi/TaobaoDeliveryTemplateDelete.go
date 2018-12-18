package tbsdk

// TaobaoDeliveryTemplateDeleteRequest 根据用户指定的模板ID删除指定的模板
type TaobaoDeliveryTemplateDeleteRequest struct {
    
    /* template_id required运费模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TaobaoDeliveryTemplateDeleteRequest) GetAPIName() string {
	return "taobao.delivery.template.delete"
}

// TaobaoDeliveryTemplateDeleteResponse 根据用户指定的模板ID删除指定的模板
type TaobaoDeliveryTemplateDeleteResponse struct {
    
    /* complete Basic表示删除成功还是失败 */
    complete bool `json:"complete";xml:"complete"`
    
}
