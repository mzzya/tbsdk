package tbsdk

// CainiaoSmartdeliverySellerStatusIQueryRequest 查询智能发货引擎商家状态信息
type CainiaoSmartdeliverySellerStatusIQueryRequest struct {
    
}

func (req *CainiaoSmartdeliverySellerStatusIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.seller.status.i.query"
}

// CainiaoSmartdeliverySellerStatusIQueryResponse 查询智能发货引擎商家状态信息
type CainiaoSmartdeliverySellerStatusIQueryResponse struct {
    
    /* seller_status Object商家状态 */
    seller_status SellerStatus `json:"seller_status";xml:"seller_status"`
    
}
