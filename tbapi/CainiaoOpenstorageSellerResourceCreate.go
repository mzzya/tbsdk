package tbsdk

// CainiaoOpenstorageSellerResourceCreateRequest 商家资源创建接口(云打印开源存储)
type CainiaoOpenstorageSellerResourceCreateRequest struct {
    
    /* param_create_seller_resource_request optional商家创建资源参数 */
    param_create_seller_resource_request CreateSellerResourceRequest `json:"param_create_seller_resource_request";xml:"param_create_seller_resource_request"`
    
}

func (req *CainiaoOpenstorageSellerResourceCreateRequest) GetAPIName() string {
	return "cainiao.openstorage.seller.resource.create"
}

// CainiaoOpenstorageSellerResourceCreateResponse 商家资源创建接口(云打印开源存储)
type CainiaoOpenstorageSellerResourceCreateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
