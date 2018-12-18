package tbsdk

// AlibabaElectronicInvoiceGetRequest 查询已回传淘宝的电子发票,根据主订单id查询
type AlibabaElectronicInvoiceGetRequest struct {
    
    /* tid required淘宝主订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *AlibabaElectronicInvoiceGetRequest) GetAPIName() string {
	return "alibaba.electronic.invoice.get"
}

// AlibabaElectronicInvoiceGetResponse 查询已回传淘宝的电子发票,根据主订单id查询
type AlibabaElectronicInvoiceGetResponse struct {
    
    /* invoice_detail Object电子发票详细信息 */
    invoice_detail ElectronicInvoiceDetail `json:"invoice_detail";xml:"invoice_detail"`
    
}
