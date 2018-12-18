package tbsdk

// CainiaoOpenstorageResourceUpdateRequest isv和商家的资源获取接口（云打印开源存储）
type CainiaoOpenstorageResourceUpdateRequest struct {
    
    /* param_update_resource_request required入参 */
    param_update_resource_request UpdateResourceRequest `json:"param_update_resource_request";xml:"param_update_resource_request"`
    
}

func (req *CainiaoOpenstorageResourceUpdateRequest) GetAPIName() string {
	return "cainiao.openstorage.resource.update"
}

// CainiaoOpenstorageResourceUpdateResponse isv和商家的资源获取接口（云打印开源存储）
type CainiaoOpenstorageResourceUpdateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
