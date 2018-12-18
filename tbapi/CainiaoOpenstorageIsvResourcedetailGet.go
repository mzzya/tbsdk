package tbsdk

// CainiaoOpenstorageIsvResourcedetailGetRequest isv资源详情查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcedetailGetRequest struct {
    
    /* isv_resource_id optionalisv资源id */
    isv_resource_id int64 `json:"isv_resource_id";xml:"isv_resource_id"`
    
}

func (req *CainiaoOpenstorageIsvResourcedetailGetRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resourcedetail.get"
}

// CainiaoOpenstorageIsvResourcedetailGetResponse isv资源详情查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcedetailGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
