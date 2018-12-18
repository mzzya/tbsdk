package tbsdk

// CainiaoOpenstorageIsvResourceCreateRequest isv资源创建接口（云打印开源存储）
type CainiaoOpenstorageIsvResourceCreateRequest struct {
    
    /* param_create_isv_resource_request optionalisv创建资源的参数 */
    param_create_isv_resource_request CreateIsvResourceRequest `json:"param_create_isv_resource_request";xml:"param_create_isv_resource_request"`
    
}

func (req *CainiaoOpenstorageIsvResourceCreateRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resource.create"
}

// CainiaoOpenstorageIsvResourceCreateResponse isv资源创建接口（云打印开源存储）
type CainiaoOpenstorageIsvResourceCreateResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_message Basic错误消息 */
    err_message string `json:"err_message";xml:"err_message"`
    
}
