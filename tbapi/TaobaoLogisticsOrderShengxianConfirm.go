package tbsdk

// TaobaoLogisticsOrderShengxianConfirmRequest 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOrderShengxianConfirmRequest struct {
    
    /* cancel_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br> */
    cancel_id int64 `json:"cancel_id";xml:"cancel_id"`
    
    /* delivery_type required1：冷链。0：常温 */
    delivery_type int64 `json:"delivery_type";xml:"delivery_type"`
    
    /* logis_id optional物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准 */
    logis_id int64 `json:"logis_id";xml:"logis_id"`
    
    /* out_sid required运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入； */
    out_sid string `json:"out_sid";xml:"out_sid"`
    
    /* seller_ip optional商家的IP地址 */
    seller_ip string `json:"seller_ip";xml:"seller_ip"`
    
    /* sender_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货 */
    sender_id int64 `json:"sender_id";xml:"sender_id"`
    
    /* service_code optional仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货 */
    service_code string `json:"service_code";xml:"service_code"`
    
    /* tid optional淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOrderShengxianConfirmRequest) GetAPIName() string {
	return "taobao.logistics.order.shengxian.confirm"
}

// TaobaoLogisticsOrderShengxianConfirmResponse 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOrderShengxianConfirmResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* ship_fresh Object发货成功后，返回承运商的信息 */
    ship_fresh ShipFresh `json:"ship_fresh";xml:"ship_fresh"`
    
}
