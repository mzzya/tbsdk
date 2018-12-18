package tbsdk

// TaobaoRpReturngoodsRefuseRequest 卖家拒绝退货，目前仅支持天猫退货。
type TaobaoRpReturngoodsRefuseRequest struct {
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required退款服务状态，售后或者售中 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version required退款版本号 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* refuse_proof required拒绝退货凭证图片，必须图片格式，大小不能超过5M */
    refuse_proof byte[] `json:"refuse_proof";xml:"refuse_proof"`
    
    /* refuse_reason_id optional拒绝原因编号，会提供拒绝原因列表供选择 */
    refuse_reason_id int64 `json:"refuse_reason_id";xml:"refuse_reason_id"`
    
}

func (req *TaobaoRpReturngoodsRefuseRequest) GetAPIName() string {
	return "taobao.rp.returngoods.refuse"
}

// TaobaoRpReturngoodsRefuseResponse 卖家拒绝退货，目前仅支持天猫退货。
type TaobaoRpReturngoodsRefuseResponse struct {
    
    /* result Basicasdf */
    result bool `json:"result";xml:"result"`
    
}
