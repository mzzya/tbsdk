package tbsdk

// CainiaoCloudprintStdtemplatesGetRequest 获取菜鸟标准电子面单模板
type CainiaoCloudprintStdtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintStdtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.stdtemplates.get"
}

// CainiaoCloudprintStdtemplatesGetResponse 获取菜鸟标准电子面单模板
type CainiaoCloudprintStdtemplatesGetResponse struct {
    
    /* result Object结果集 */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
