package tbsdk

// CainiaoOpenstorageResourcePublishRequest isv和商家资源发布接口(云打印开源存储)
type CainiaoOpenstorageResourcePublishRequest struct {
    
    /* param_publish_resource_request optional资源发布参数 */
    param_publish_resource_request PublishResourceRequest `json:"param_publish_resource_request";xml:"param_publish_resource_request"`
    
}

func (req *CainiaoOpenstorageResourcePublishRequest) GetAPIName() string {
	return "cainiao.openstorage.resource.publish"
}

// CainiaoOpenstorageResourcePublishResponse isv和商家资源发布接口(云打印开源存储)
type CainiaoOpenstorageResourcePublishResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
