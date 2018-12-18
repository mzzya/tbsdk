package tbsdk

// AlibabaEinvoiceApplyGetRequest ERP获取开票申请数据
type AlibabaEinvoiceApplyGetRequest struct {
    
    /* apply_id optional开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息 */
    apply_id string `json:"apply_id";xml:"apply_id"`
    
    /* platform_tid required平台订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
}

func (req *AlibabaEinvoiceApplyGetRequest) GetAPIName() string {
	return "alibaba.einvoice.apply.get"
}

// AlibabaEinvoiceApplyGetResponse ERP获取开票申请数据
type AlibabaEinvoiceApplyGetResponse struct {
    
    /* apply_list Object Array开票明细 */
    apply_list Apply `json:"apply_list";xml:"apply_list"`
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
