package tbsdk

// CainiaoCloudprintIsvResourcesGetRequest isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaoCloudprintIsvResourcesGetRequest struct {
    
    /* isv_resource_type requiredisv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区） */
    isv_resource_type string `json:"isv_resource_type";xml:"isv_resource_type"`
    
}

func (req *CainiaoCloudprintIsvResourcesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.isv.resources.get"
}

// CainiaoCloudprintIsvResourcesGetResponse isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaoCloudprintIsvResourcesGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
