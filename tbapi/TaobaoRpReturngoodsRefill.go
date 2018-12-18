package tbsdk

// TaobaoRpReturngoodsRefillRequest 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillRequest struct {
    
    /* logistics_company_code required物流公司编号 */
    logistics_company_code string `json:"logistics_company_code";xml:"logistics_company_code"`
    
    /* logistics_waybill_no required物流公司运单号 */
    logistics_waybill_no string `json:"logistics_waybill_no";xml:"logistics_waybill_no"`
    
    /* refund_id required退款单编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required退款阶段，可选值：售中：onsale，售后：aftersale */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRpReturngoodsRefillRequest) GetAPIName() string {
	return "taobao.rp.returngoods.refill"
}

// TaobaoRpReturngoodsRefillResponse 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillResponse struct {
    
    /* is_success Basic验货操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
