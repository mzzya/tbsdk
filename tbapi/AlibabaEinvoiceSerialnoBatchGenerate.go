package tbsdk

// AlibabaEinvoiceSerialnoBatchGenerateRequest 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。
type AlibabaEinvoiceSerialnoBatchGenerateRequest struct {
    
}

func (req *AlibabaEinvoiceSerialnoBatchGenerateRequest) GetAPIName() string {
	return "alibaba.einvoice.serialno.batch.generate"
}

// AlibabaEinvoiceSerialnoBatchGenerateResponse 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。
type AlibabaEinvoiceSerialnoBatchGenerateResponse struct {
    
    /* serial_no_list Basic Arrayresult */
    serial_no_list string `json:"serial_no_list";xml:"serial_no_list"`
    
}
