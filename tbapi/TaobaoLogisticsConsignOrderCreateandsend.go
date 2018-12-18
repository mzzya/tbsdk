package tbsdk

// TaobaoLogisticsConsignOrderCreateandsendRequest 创建物流订单，并发货。
type TaobaoLogisticsConsignOrderCreateandsendRequest struct {
    
    /* company_id required物流公司ID */
    company_id int64 `json:"company_id";xml:"company_id"`
    
    /* item_json_string required物品的json数据。 */
    item_json_string string `json:"item_json_string";xml:"item_json_string"`
    
    /* logis_type required物流订单物流类型，值固定选择：2 */
    logis_type int64 `json:"logis_type";xml:"logis_type"`
    
    /* mail_no optional运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* order_source required订单来源，值选择：30 */
    order_source int64 `json:"order_source";xml:"order_source"`
    
    /* order_type required订单类型，值固定选择：30 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* r_address required收件人街道地址 */
    r_address string `json:"r_address";xml:"r_address"`
    
    /* r_area_id required收件人区域ID */
    r_area_id int64 `json:"r_area_id";xml:"r_area_id"`
    
    /* r_city_name required市 */
    r_city_name string `json:"r_city_name";xml:"r_city_name"`
    
    /* r_dist_name optional区 */
    r_dist_name string `json:"r_dist_name";xml:"r_dist_name"`
    
    /* r_mobile_phone optional手机号码 */
    r_mobile_phone string `json:"r_mobile_phone";xml:"r_mobile_phone"`
    
    /* r_name required收件人名称 */
    r_name string `json:"r_name";xml:"r_name"`
    
    /* r_prov_name required省 */
    r_prov_name string `json:"r_prov_name";xml:"r_prov_name"`
    
    /* r_telephone optional电话号码 */
    r_telephone string `json:"r_telephone";xml:"r_telephone"`
    
    /* r_zip_code required收件人邮编 */
    r_zip_code string `json:"r_zip_code";xml:"r_zip_code"`
    
    /* s_address required发件人街道地址 */
    s_address string `json:"s_address";xml:"s_address"`
    
    /* s_area_id required发件人区域ID */
    s_area_id int64 `json:"s_area_id";xml:"s_area_id"`
    
    /* s_city_name required市 */
    s_city_name string `json:"s_city_name";xml:"s_city_name"`
    
    /* s_dist_name optional区 */
    s_dist_name string `json:"s_dist_name";xml:"s_dist_name"`
    
    /* s_mobile_phone optional手机号码 */
    s_mobile_phone string `json:"s_mobile_phone";xml:"s_mobile_phone"`
    
    /* s_name required发件人名称 */
    s_name string `json:"s_name";xml:"s_name"`
    
    /* s_prov_name required省 */
    s_prov_name string `json:"s_prov_name";xml:"s_prov_name"`
    
    /* s_telephone optional电话号码 */
    s_telephone string `json:"s_telephone";xml:"s_telephone"`
    
    /* s_zip_code required发件人出编 */
    s_zip_code string `json:"s_zip_code";xml:"s_zip_code"`
    
    /* shipping optional费用承担方式 1买家承担运费 2卖家承担运费 */
    shipping string `json:"shipping";xml:"shipping"`
    
    /* trade_id required交易流水号，淘外订单号或者商家内部交易流水号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
    /* user_id required用户ID */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoLogisticsConsignOrderCreateandsendRequest) GetAPIName() string {
	return "taobao.logistics.consign.order.createandsend"
}

// TaobaoLogisticsConsignOrderCreateandsendResponse 创建物流订单，并发货。
type TaobaoLogisticsConsignOrderCreateandsendResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* order_id Basic订单ID */
    order_id int64 `json:"order_id";xml:"order_id"`
    
    /* result_desc Basic结果描述 */
    result_desc string `json:"result_desc";xml:"result_desc"`
    
}
