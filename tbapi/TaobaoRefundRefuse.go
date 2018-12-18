package tbsdk

// TaobaoRefundRefuseRequest 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：
1. 传入的refund_id和相应的tid, oid必须匹配
2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
3. 只有卖家才能执行拒绝退款操作
4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
type TaobaoRefundRefuseRequest struct {
    
    /* oid optional退款记录对应的交易子订单号 */
    oid int64 `json:"oid";xml:"oid"`
    
    /* refund_id required退款单号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase optional可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version optional退款版本号，天猫退款为必填项。 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* refuse_message required拒绝退款时的说明信息，长度2-200 */
    refuse_message string `json:"refuse_message";xml:"refuse_message"`
    
    /* refuse_proof optional拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。 */
    refuse_proof byte[] `json:"refuse_proof";xml:"refuse_proof"`
    
    /* refuse_reason_id optional拒绝原因编号，会提供用户拒绝原因列表供选择 */
    refuse_reason_id int64 `json:"refuse_reason_id";xml:"refuse_reason_id"`
    
    /* tid optional退款记录对应的交易订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoRefundRefuseRequest) GetAPIName() string {
	return "taobao.refund.refuse"
}

// TaobaoRefundRefuseResponse 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：
1. 传入的refund_id和相应的tid, oid必须匹配
2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
3. 只有卖家才能执行拒绝退款操作
4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
type TaobaoRefundRefuseResponse struct {
    
    /* is_success Basic拒绝退款操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* refund Object拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段 */
    refund Refund `json:"refund";xml:"refund"`
    
}
