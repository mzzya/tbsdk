package tbsdk

// AlibabaEinvoiceApplyTestRequest 开票申请消息测试接口
type AlibabaEinvoiceApplyTestRequest struct {
    
    /* business_type required0=个人，1=企业 */
    business_type int64 `json:"business_type";xml:"business_type"`
    
    /* payer_name required买家抬头 */
    payer_name string `json:"payer_name";xml:"payer_name"`
    
    /* payer_register_no optional买家税号 */
    payer_register_no string `json:"payer_register_no";xml:"payer_register_no"`
    
    /* platform_tid required订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
}

func (req *AlibabaEinvoiceApplyTestRequest) GetAPIName() string {
	return "alibaba.einvoice.apply.test"
}

// AlibabaEinvoiceApplyTestResponse 开票申请消息测试接口
type AlibabaEinvoiceApplyTestResponse struct {
    
    /* is_success Basic消息是否发送成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
