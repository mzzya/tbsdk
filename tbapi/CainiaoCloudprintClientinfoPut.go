package tbsdk

// CainiaoCloudprintClientinfoPutRequest 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutRequest struct {
    
    /* json_data required客户端上传json数据 */
    json_data string `json:"json_data";xml:"json_data"`
    
}

func (req *CainiaoCloudprintClientinfoPutRequest) GetAPIName() string {
	return "cainiao.cloudprint.clientinfo.put"
}

// CainiaoCloudprintClientinfoPutResponse 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
