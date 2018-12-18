package tbsdk

// AlibabaProviderEinvoiceQueryRequest 开票服务商提供的对已开电子发票的查询功能，此接口是可选实现。只有部分电子发票支持通过此接口查询
type AlibabaProviderEinvoiceQueryRequest struct {
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* platform_code optional电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码) */
    platform_code string `json:"platform_code";xml:"platform_code"`
    
    /* platform_tid optional电商平台对应的订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
    /* serial_no special开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。与erp_tid二选一 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaProviderEinvoiceQueryRequest) GetAPIName() string {
	return "alibaba.provider.einvoice.query"
}

// AlibabaProviderEinvoiceQueryResponse 开票服务商提供的对已开电子发票的查询功能，此接口是可选实现。只有部分电子发票支持通过此接口查询
type AlibabaProviderEinvoiceQueryResponse struct {
    
    /* einvoice_list Object Array电子发票列表 */
    einvoice_list ProviderEinvoice `json:"einvoice_list";xml:"einvoice_list"`
    
}
