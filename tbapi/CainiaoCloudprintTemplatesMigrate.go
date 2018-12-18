package tbsdk

// CainiaoCloudprintTemplatesMigrateRequest 云打印模板迁移接口
type CainiaoCloudprintTemplatesMigrateRequest struct {
    
    /* custom_area_content optional自定义区内容 */
    custom_area_content string `json:"custom_area_content";xml:"custom_area_content"`
    
    /* custom_area_name optional自定义区名称 */
    custom_area_name string `json:"custom_area_name";xml:"custom_area_name"`
    
    /* tempalte_id optional标准电子面单模板的id */
    tempalte_id int64 `json:"tempalte_id";xml:"tempalte_id"`
    
}

func (req *CainiaoCloudprintTemplatesMigrateRequest) GetAPIName() string {
	return "cainiao.cloudprint.templates.migrate"
}

// CainiaoCloudprintTemplatesMigrateResponse 云打印模板迁移接口
type CainiaoCloudprintTemplatesMigrateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
