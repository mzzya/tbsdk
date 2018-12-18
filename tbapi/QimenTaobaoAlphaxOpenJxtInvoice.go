package tbsdk

// QimenTaobaoAlphaxOpenJxtInvoiceRequest isv 开发票接口api
type QimenTaobaoAlphaxOpenJxtInvoiceRequest struct {
    
    /* company_title optional公司抬头 */
    company_title string `json:"company_title";xml:"company_title"`
    
    /* extend_arg optionalapi新增字段，主要用于扩展参数，例如增值税扩展字段（registered_address 注册地址、registered_phone 注册电话、bank 开户行、账户 ） */
    extend_arg string `json:"extend_arg";xml:"extend_arg"`
    
    /* invoice_attr required发票属性(0:公司；1：个人) */
    invoice_attr int64 `json:"invoice_attr";xml:"invoice_attr"`
    
    /* invoice_kind required发票形态 （1:电子发票;   2：纸质发票) */
    invoice_kind int64 `json:"invoice_kind";xml:"invoice_kind"`
    
    /* invoice_type required发票类型（1:普通发票；2：增值税专用发票） */
    invoice_type int64 `json:"invoice_type";xml:"invoice_type"`
    
    /* order_id required订单id */
    order_id string `json:"order_id";xml:"order_id"`
    
    /* seller_id required卖家主账号id */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
    /* tax_no optional税号 */
    tax_no string `json:"tax_no";xml:"tax_no"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtInvoiceRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.invoice"
}

// QimenTaobaoAlphaxOpenJxtInvoiceResponse isv 开发票接口api
type QimenTaobaoAlphaxOpenJxtInvoiceResponse struct {
    
    /* result Object返回值 */
    result Data `json:"result";xml:"result"`
    
}
