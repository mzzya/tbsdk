package tbsdk

// AlibabaEinvoiceCreatereqRequest ERP发起开票请求
type AlibabaEinvoiceCreatereqRequest struct {
    
    /* apply_id optional开票申请ID，接收了开票申请消息后，需要把apply_id带上 */
    apply_id string `json:"apply_id";xml:"apply_id"`
    
    /* business_type required默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1; */
    business_type int64 `json:"business_type";xml:"business_type"`
    
    /* erp_tid optionalERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式 */
    erp_tid string `json:"erp_tid";xml:"erp_tid"`
    
    /* invoice_amount required开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span> */
    invoice_amount string `json:"invoice_amount";xml:"invoice_amount"`
    
    /* invoice_items required电子发票明细 */
    invoice_items InvoiceItem `json:"invoice_items";xml:"invoice_items"`
    
    /* invoice_kind optional发票种类，0=电子发票,1=纸质发票,2=专票。注意：未订购纸票服务的税号无法开具纸票 */
    invoice_kind int64 `json:"invoice_kind";xml:"invoice_kind"`
    
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
    
    /* out_shop_name optional外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致 */
    out_shop_name string `json:"out_shop_name";xml:"out_shop_name"`
    
    /* payee_address required开票方地址(新版中为必传) */
    payee_address string `json:"payee_address";xml:"payee_address"`
    
    /* payee_bankaccount optional开票方银行及 帐号 */
    payee_bankaccount string `json:"payee_bankaccount";xml:"payee_bankaccount"`
    
    /* payee_checker optional复核人 */
    payee_checker string `json:"payee_checker";xml:"payee_checker"`
    
    /* payee_name required开票方名称，公司名(如:XX商城) */
    payee_name string `json:"payee_name";xml:"payee_name"`
    
    /* payee_operator optional开票人 */
    payee_operator string `json:"payee_operator";xml:"payee_operator"`
    
    /* payee_phone optional收款方电话 */
    payee_phone string `json:"payee_phone";xml:"payee_phone"`
    
    /* payee_receiver optional收款人 */
    payee_receiver string `json:"payee_receiver";xml:"payee_receiver"`
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* payer_address optional消费者地址 */
    payer_address string `json:"payer_address";xml:"payer_address"`
    
    /* payer_bankaccount optional付款方开票开户银行及账号 */
    payer_bankaccount string `json:"payer_bankaccount";xml:"payer_bankaccount"`
    
    /* payer_email optional消费者电子邮箱 */
    payer_email string `json:"payer_email";xml:"payer_email"`
    
    /* payer_name required付款方名称, 对应发票台头 */
    payer_name string `json:"payer_name";xml:"payer_name"`
    
    /* payer_phone optional消费者联系电话 */
    payer_phone string `json:"payer_phone";xml:"payer_phone"`
    
    /* payer_register_no optional付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空 */
    payer_register_no string `json:"payer_register_no";xml:"payer_register_no"`
    
    /* platform_code required电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码) */
    platform_code string `json:"platform_code";xml:"platform_code"`
    
    /* platform_tid required电商平台对应的主订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
    /* red_notice_no optional红字通知单号，专票冲红时需要，商家跟税局申请 */
    red_notice_no string `json:"red_notice_no";xml:"red_notice_no"`
    
    /* serial_no required开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
    /* sum_price required合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span> */
    sum_price string `json:"sum_price";xml:"sum_price"`
    
    /* sum_tax required合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span> */
    sum_tax string `json:"sum_tax";xml:"sum_tax"`
    
}

func (req *AlibabaEinvoiceCreatereqRequest) GetAPIName() string {
	return "alibaba.einvoice.createreq"
}

// AlibabaEinvoiceCreatereqResponse ERP发起开票请求
type AlibabaEinvoiceCreatereqResponse struct {
    
    /* is_success Basic开票信息是否成功接受 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
