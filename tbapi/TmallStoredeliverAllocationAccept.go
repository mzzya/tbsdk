package tbsdk

// TmallStoredeliverAllocationAcceptRequest 商品通门店发货业务门店接单拒单接口
type TmallStoredeliverAllocationAcceptRequest struct {
    
    /* allocation_code required派单号 */
    allocation_code string `json:"allocation_code";xml:"allocation_code"`
    
    /* is_accept required是否接单 */
    is_accept bool `json:"is_accept";xml:"is_accept"`
    
    /* sub_order_code required子订单号必须和派单号匹配 */
    sub_order_code int64 `json:"sub_order_code";xml:"sub_order_code"`
    
}

func (req *TmallStoredeliverAllocationAcceptRequest) GetAPIName() string {
	return "tmall.storedeliver.allocation.accept"
}

// TmallStoredeliverAllocationAcceptResponse 商品通门店发货业务门店接单拒单接口
type TmallStoredeliverAllocationAcceptResponse struct {
    
    /* is_success Basic是否执行成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
