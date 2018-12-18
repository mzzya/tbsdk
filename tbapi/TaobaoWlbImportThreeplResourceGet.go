package tbsdk

// TaobaoWlbImportThreeplResourceGetRequest 获取3pl直邮的发货可用资源
type TaobaoWlbImportThreeplResourceGetRequest struct {
    
    /* from_id required发货地区域id */
    from_id int64 `json:"from_id";xml:"from_id"`
    
    /* to_address required收件人地址 */
    to_address AddressDto `json:"to_address";xml:"to_address"`
    
    /* _type requiredONLINE或者OFFLINE */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbImportThreeplResourceGetRequest) GetAPIName() string {
	return "taobao.wlb.import.threepl.resource.get"
}

// TaobaoWlbImportThreeplResourceGetResponse 获取3pl直邮的发货可用资源
type TaobaoWlbImportThreeplResourceGetResponse struct {
    
    /* result Objectresult */
    result TopResult `json:"result";xml:"result"`
    
}
