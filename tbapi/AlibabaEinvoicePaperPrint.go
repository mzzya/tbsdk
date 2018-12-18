package tbsdk

// AlibabaEinvoicePaperPrintRequest 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintRequest struct {
    
    /* dialog_setting_flag required打印框设置，0=不弹打印设置框，1=弹出打印设置框 */
    dialog_setting_flag int64 `json:"dialog_setting_flag";xml:"dialog_setting_flag"`
    
    /* force_print required是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印 */
    force_print bool `json:"force_print";xml:"force_print"`
    
    /* payee_register_no required销售方纳税人识别号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* print_flag required打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。 */
    print_flag int64 `json:"print_flag";xml:"print_flag"`
    
    /* serial_no required开票流水号 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoicePaperPrintRequest) GetAPIName() string {
	return "alibaba.einvoice.paper.print"
}

// AlibabaEinvoicePaperPrintResponse 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintResponse struct {
    
    /* is_success Basic调用结果，打印结果tmc消息alibaba_invoice_PaperOpsReturn异步通知 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
