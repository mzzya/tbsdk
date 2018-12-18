package tbsdk

// CainiaoOpenstorageIsvResourcesGetRequest isv资源列表查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcesGetRequest struct {
    
    /* isv_resource_type requiredisv资源类型，分为：TEMPLATE（模板）和PRINT_ITEM（打印项） */
    isv_resource_type string `json:"isv_resource_type";xml:"isv_resource_type"`
    
}

func (req *CainiaoOpenstorageIsvResourcesGetRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resources.get"
}

// CainiaoOpenstorageIsvResourcesGetResponse isv资源列表查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcesGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
