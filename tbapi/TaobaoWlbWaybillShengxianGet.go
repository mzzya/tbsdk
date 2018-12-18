package tbsdk

// TaobaoWlbWaybillShengxianGetRequest 商家通过交易订单号获取电子面单接口
type TaobaoWlbWaybillShengxianGetRequest struct {
    
    /* biz_code required物流服务方代码，生鲜配送：YXSR */
    biz_code string `json:"biz_code";xml:"biz_code"`
    
    /* delivery_type required物流服务类型。冷链：602，常温：502 */
    delivery_type string `json:"delivery_type";xml:"delivery_type"`
    
    /* feature optional预留扩展字段 */
    feature string `json:"feature";xml:"feature"`
    
    /* order_channels_type required订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688"; */
    order_channels_type string `json:"order_channels_type";xml:"order_channels_type"`
    
    /* sender_address_id optional商家淘宝地址信息ID */
    sender_address_id string `json:"sender_address_id";xml:"sender_address_id"`
    
    /* service_code optional仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询) */
    service_code string `json:"service_code";xml:"service_code"`
    
    /* trade_id required交易号，传入交易号不能存在拆单场景。 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoWlbWaybillShengxianGetRequest) GetAPIName() string {
	return "taobao.wlb.waybill.shengxian.get"
}

// TaobaoWlbWaybillShengxianGetResponse 商家通过交易订单号获取电子面单接口
type TaobaoWlbWaybillShengxianGetResponse struct {
    
    /* fresh_waybill Object成功后返回的生鲜电子面单信息 */
    fresh_waybill FreshWaybill `json:"fresh_waybill";xml:"fresh_waybill"`
    
    /* is_success Basic生成是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
