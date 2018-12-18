package tbsdk

// TaobaoRpRefundsAgreeRequest 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
type TaobaoRpRefundsAgreeRequest struct {
    
    /* code optional短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。 */
    code string `json:"code";xml:"code"`
    
    /* refund_infos required退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。 */
    refund_infos string `json:"refund_infos";xml:"refund_infos"`
    
}

func (req *TaobaoRpRefundsAgreeRequest) GetAPIName() string {
	return "taobao.rp.refunds.agree"
}

// TaobaoRpRefundsAgreeResponse 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
type TaobaoRpRefundsAgreeResponse struct {
    
    /* message Basic信息 */
    message string `json:"message";xml:"message"`
    
    /* msg_code Basic批量退款操作情况，可选值：OP_SUCC（全部成功），SOME_OP_SUCC（部分成功），OP_FAILURE_UE（全部失败） */
    msg_code string `json:"msg_code";xml:"msg_code"`
    
    /* results Object Array退款操作结果列表 */
    results RefundMappingResult `json:"results";xml:"results"`
    
    /* succ Basic操作成功 */
    succ bool `json:"succ";xml:"succ"`
    
}
