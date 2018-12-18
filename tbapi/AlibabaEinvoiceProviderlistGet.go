package tbsdk

// AlibabaEinvoiceProviderlistGetRequest 获取能提供电子发票开票服务的开票服务商
type AlibabaEinvoiceProviderlistGetRequest struct {
    
}

func (req *AlibabaEinvoiceProviderlistGetRequest) GetAPIName() string {
	return "alibaba.einvoice.providerlist.get"
}

// AlibabaEinvoiceProviderlistGetResponse 获取能提供电子发票开票服务的开票服务商
type AlibabaEinvoiceProviderlistGetResponse struct {
    
    /* result Object查询结果集 */
    result ResultSet `json:"result";xml:"result"`
    
}
