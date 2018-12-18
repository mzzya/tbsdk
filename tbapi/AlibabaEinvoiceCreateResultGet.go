package tbsdk

// AlibabaEinvoiceCreateResultGetRequest ERP开票结果获取
type AlibabaEinvoiceCreateResultGetRequest struct {
    
    /* out_shop_name optional外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致 */
    out_shop_name string `json:"out_shop_name";xml:"out_shop_name"`
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* platform_code optional电商平台代码。淘宝：taobao，天猫：tmall */
    platform_code string `json:"platform_code";xml:"platform_code"`
    
    /* platform_tid optional电商平台对应的订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
    /* serial_no optional流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoiceCreateResultGetRequest) GetAPIName() string {
	return "alibaba.einvoice.create.result.get"
}

// AlibabaEinvoiceCreateResultGetResponse ERP开票结果获取
type AlibabaEinvoiceCreateResultGetResponse struct {
    
    /* invoice_result_list Object Array开票返回结果数据列表 */
    invoice_result_list InvoiceResult `json:"invoice_result_list";xml:"invoice_result_list"`
    
}
