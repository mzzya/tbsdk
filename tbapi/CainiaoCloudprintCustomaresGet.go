package tbsdk

// CainiaoCloudprintCustomaresGetRequest 供isv使用，获取商家的自定义区的模板信息
type CainiaoCloudprintCustomaresGetRequest struct {
    
    /* template_id required用户使用的标准模板id */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *CainiaoCloudprintCustomaresGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.customares.get"
}

// CainiaoCloudprintCustomaresGetResponse 供isv使用，获取商家的自定义区的模板信息
type CainiaoCloudprintCustomaresGetResponse struct {
    
    /* result Object结果 */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
