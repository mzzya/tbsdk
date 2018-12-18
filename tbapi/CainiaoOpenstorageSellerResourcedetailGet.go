package tbsdk

// CainiaoOpenstorageSellerResourcedetailGetRequest 商家资源详情获取（云打印开源存储）
type CainiaoOpenstorageSellerResourcedetailGetRequest struct {
    
    /* seller_resource_id optional商家资源id */
    seller_resource_id int64 `json:"seller_resource_id";xml:"seller_resource_id"`
    
}

func (req *CainiaoOpenstorageSellerResourcedetailGetRequest) GetAPIName() string {
	return "cainiao.openstorage.seller.resourcedetail.get"
}

// CainiaoOpenstorageSellerResourcedetailGetResponse 商家资源详情获取（云打印开源存储）
type CainiaoOpenstorageSellerResourcedetailGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
