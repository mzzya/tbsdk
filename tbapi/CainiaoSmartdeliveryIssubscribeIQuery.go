package tbsdk

// CainiaoSmartdeliveryIssubscribeIQueryRequest 查询商家时候订购智能发货引擎服务
type CainiaoSmartdeliveryIssubscribeIQueryRequest struct {
    
}

func (req *CainiaoSmartdeliveryIssubscribeIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.issubscribe.i.query"
}

// CainiaoSmartdeliveryIssubscribeIQueryResponse 查询商家时候订购智能发货引擎服务
type CainiaoSmartdeliveryIssubscribeIQueryResponse struct {
    
    /* successful Basictrue:商家已订购智能发货引擎服务,false:商家还没有订购或订购已过期 */
    successful bool `json:"successful";xml:"successful"`
    
}
