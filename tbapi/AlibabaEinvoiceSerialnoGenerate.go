package tbsdk

// AlibabaEinvoiceSerialnoGenerateRequest erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaEinvoiceSerialnoGenerateRequest struct {
    
}

func (req *AlibabaEinvoiceSerialnoGenerateRequest) GetAPIName() string {
	return "alibaba.einvoice.serialno.generate"
}

// AlibabaEinvoiceSerialnoGenerateResponse erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaEinvoiceSerialnoGenerateResponse struct {
    
    /* serial_no Basicresult */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}
