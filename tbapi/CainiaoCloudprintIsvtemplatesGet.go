package tbsdk

// CainiaoCloudprintIsvtemplatesGetRequest 获取商家使用的标准模板
type CainiaoCloudprintIsvtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintIsvtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.isvtemplates.get"
}

// CainiaoCloudprintIsvtemplatesGetResponse 获取商家使用的标准模板
type CainiaoCloudprintIsvtemplatesGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
