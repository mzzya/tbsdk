package tbsdk

// CainiaoSmartdeliveryIGetRequest 获取智选cp和电子面单信息
type CainiaoSmartdeliveryIGetRequest struct {
    
    /* smart_delivery_batch_request required<a href="http://open.taobao.com/docs/doc.htm?treeId=319&articleId=106295&docType=1">智能发货引擎</a>批量请求参数 */
    smart_delivery_batch_request SmartDeliveryBatchRequest `json:"smart_delivery_batch_request";xml:"smart_delivery_batch_request"`
    
}

func (req *CainiaoSmartdeliveryIGetRequest) GetAPIName() string {
	return "cainiao.smartdelivery.i.get"
}

// CainiaoSmartdeliveryIGetResponse 获取智选cp和电子面单信息
type CainiaoSmartdeliveryIGetResponse struct {
    
    /* smart_delivery_response_wrapper_list Object Array<a href="http://open.taobao.com/docs/doc.htm?treeId=319&articleId=106295&docType=1">智能发货引擎</a>结果包装类 */
    smart_delivery_response_wrapper_list SmartDeliveryResponseWrapper `json:"smart_delivery_response_wrapper_list";xml:"smart_delivery_response_wrapper_list"`
    
}
