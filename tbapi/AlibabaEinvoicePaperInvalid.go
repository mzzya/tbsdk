package tbsdk

// AlibabaEinvoicePaperInvalidRequest 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaEinvoicePaperInvalidRequest struct {
    
    /* invalid_operator required作废操作人 */
    invalid_operator string `json:"invalid_operator";xml:"invalid_operator"`
    
    /* invalid_type required作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废 */
    invalid_type int64 `json:"invalid_type";xml:"invalid_type"`
    
    /* invoice_code optional发票代码，空白作废时必填 */
    invoice_code string `json:"invoice_code";xml:"invoice_code"`
    
    /* invoice_no optional发票号码，空白作废时必填 */
    invoice_no string `json:"invoice_no";xml:"invoice_no"`
    
    /* payee_register_no required销售方纳税人识别号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* serial_no required开票流水号 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoicePaperInvalidRequest) GetAPIName() string {
	return "alibaba.einvoice.paper.invalid"
}

// AlibabaEinvoicePaperInvalidResponse 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaEinvoicePaperInvalidResponse struct {
    
    /* is_success Basic接口调用是否成功，操作结果tmc异步返回alibaba_invoice_PaperOpsReturn */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
