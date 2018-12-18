package tbsdk

// AlibabaProviderEinvoiceCreateRequest 电子发票同步开票
type AlibabaProviderEinvoiceCreateRequest struct {
    
    /* biz_sign optional业务签名，用ca私钥做的业务签名 */
    biz_sign string `json:"biz_sign";xml:"biz_sign"`
    
    /* business_type required默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1 */
    business_type int64 `json:"business_type";xml:"business_type"`
    
    /* erp_tid optionalERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式 */
    erp_tid string `json:"erp_tid";xml:"erp_tid"`
    
    /* invoice_amount required开票金额(对应新版的价税合计,价税合计=合计金额+合计税额) */
    invoice_amount string `json:"invoice_amount";xml:"invoice_amount"`
    
    /* invoice_items optional电子发票明细 */
    invoice_items InvoiceItem `json:"invoice_items";xml:"invoice_items"`
    
    /* invoice_memo optional发票备注，有些省市会把此信息打印到PDF中 */
    invoice_memo string `json:"invoice_memo";xml:"invoice_memo"`
    
    /* invoice_time required开票日期, 格式"YYYY-MM-DD HH:SS:MM" */
    invoice_time Date `json:"invoice_time";xml:"invoice_time"`
    
    /* invoice_type required发票(开票)类型，蓝票blue,红票red，默认blue */
    invoice_type string `json:"invoice_type";xml:"invoice_type"`
    
    /* normal_invoice_code optional原发票代码(开红票时传入) */
    normal_invoice_code string `json:"normal_invoice_code";xml:"normal_invoice_code"`
    
    /* normal_invoice_no optional原发票号码(开红票时传入) */
    normal_invoice_no string `json:"normal_invoice_no";xml:"normal_invoice_no"`
    
    /* payee_address required开票方地址(新版中为必传) */
    payee_address string `json:"payee_address";xml:"payee_address"`
    
    /* payee_bankaccount optional开票方银行及 帐号 */
    payee_bankaccount string `json:"payee_bankaccount";xml:"payee_bankaccount"`
    
    /* payee_checker optional复核人 */
    payee_checker string `json:"payee_checker";xml:"payee_checker"`
    
    /* payee_name required开票方名称，公司名(如:XX商城) */
    payee_name string `json:"payee_name";xml:"payee_name"`
    
    /* payee_operator required开票人 */
    payee_operator string `json:"payee_operator";xml:"payee_operator"`
    
    /* payee_phone required开票方 电话(新版中为必传) */
    payee_phone string `json:"payee_phone";xml:"payee_phone"`
    
    /* payee_receiver optional收款人 */
    payee_receiver string `json:"payee_receiver";xml:"payee_receiver"`
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* payer_address optional消费者地址，开具增值税专用发票时必填，其他发票可以为空 */
    payer_address string `json:"payer_address";xml:"payer_address"`
    
    /* payer_bankaccount optional付款方开票开户银行及账号 */
    payer_bankaccount string `json:"payer_bankaccount";xml:"payer_bankaccount"`
    
    /* payer_email optional消费者电子邮箱 */
    payer_email string `json:"payer_email";xml:"payer_email"`
    
    /* payer_name required付款方名称, 对应发票台头 */
    payer_name string `json:"payer_name";xml:"payer_name"`
    
    /* payer_phone optional消费者联系电话，开具增值税专用发票时必填，其他发票可以为空。 */
    payer_phone string `json:"payer_phone";xml:"payer_phone"`
    
    /* payer_register_no optional付款方税务登记证号。 */
    payer_register_no string `json:"payer_register_no";xml:"payer_register_no"`
    
    /* platform_code required电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码) */
    platform_code string `json:"platform_code";xml:"platform_code"`
    
    /* platform_tid required电商平台对应的订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
    /* serial_no required开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
    /* sum_price required合计金额(新版中为必传) */
    sum_price string `json:"sum_price";xml:"sum_price"`
    
    /* sum_tax required合计税额(新版中为必传) */
    sum_tax string `json:"sum_tax";xml:"sum_tax"`
    
}

func (req *AlibabaProviderEinvoiceCreateRequest) GetAPIName() string {
	return "alibaba.provider.einvoice.create"
}

// AlibabaProviderEinvoiceCreateResponse 电子发票同步开票
type AlibabaProviderEinvoiceCreateResponse struct {
    
    /* anti_fake_code Basic防伪码 */
    anti_fake_code string `json:"anti_fake_code";xml:"anti_fake_code"`
    
    /* ciphertext Basic发票密文，密码区的字符串 */
    ciphertext string `json:"ciphertext";xml:"ciphertext"`
    
    /* device_no Basic税控设备编号(新版电子发票有) */
    device_no string `json:"device_no";xml:"device_no"`
    
    /* erp_tid Basicerp自定义单据号 */
    erp_tid string `json:"erp_tid";xml:"erp_tid"`
    
    /* file_data_type Basic文件类型(pdf,jpg,png) */
    file_data_type string `json:"file_data_type";xml:"file_data_type"`
    
    /* invoice_amount Basic实际开票金额 */
    invoice_amount string `json:"invoice_amount";xml:"invoice_amount"`
    
    /* invoice_code Basic发票代码 */
    invoice_code string `json:"invoice_code";xml:"invoice_code"`
    
    /* invoice_date Basic开票日期 */
    invoice_date string `json:"invoice_date";xml:"invoice_date"`
    
    /* invoice_file_data Basic发票文件PDF内容，PDF的byte[]用base64编码后的字段串。 */
    invoice_file_data string `json:"invoice_file_data";xml:"invoice_file_data"`
    
    /* invoice_no Basic发票号码 */
    invoice_no string `json:"invoice_no";xml:"invoice_no"`
    
    /* platform_code Basic电商平台身份标识码，如taobao,tmall */
    platform_code string `json:"platform_code";xml:"platform_code"`
    
    /* platform_tid Basic电商平台订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
    /* qr_code Basic二维码 */
    qr_code string `json:"qr_code";xml:"qr_code"`
    
    /* serial_no Basic电子发票流水号，标记业务上的唯一请求 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}
