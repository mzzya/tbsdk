package tbsdk

// AlibabaEinvoiceClosereqRequest 关闭失败开票请求，避免造成重复开票
type AlibabaEinvoiceClosereqRequest struct {
    
    /* payee_register_no required税号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* serial_no required流水号 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoiceClosereqRequest) GetAPIName() string {
	return "alibaba.einvoice.closereq"
}

// AlibabaEinvoiceClosereqResponse 关闭失败开票请求，避免造成重复开票
type AlibabaEinvoiceClosereqResponse struct {
    
    /* result Basic关闭是否成功 */
    result bool `json:"result";xml:"result"`
    
}
