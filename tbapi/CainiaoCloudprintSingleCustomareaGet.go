package tbsdk

// CainiaoCloudprintSingleCustomareaGetRequest 商家所有快递公司模板只有一个自定义区
type CainiaoCloudprintSingleCustomareaGetRequest struct {
    
    /* seller_id optional这是商家用户id */
    seller_id int64 `json:"seller_id";xml:"seller_id"`
    
}

func (req *CainiaoCloudprintSingleCustomareaGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.single.customarea.get"
}

// CainiaoCloudprintSingleCustomareaGetResponse 商家所有快递公司模板只有一个自定义区
type CainiaoCloudprintSingleCustomareaGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
