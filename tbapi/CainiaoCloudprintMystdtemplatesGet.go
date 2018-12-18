package tbsdk

// CainiaoCloudprintMystdtemplatesGetRequest 获取用户使用的菜鸟电子面单
type CainiaoCloudprintMystdtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintMystdtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.mystdtemplates.get"
}

// CainiaoCloudprintMystdtemplatesGetResponse 获取用户使用的菜鸟电子面单
type CainiaoCloudprintMystdtemplatesGetResponse struct {
    
    /* result Object返回结果 */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
