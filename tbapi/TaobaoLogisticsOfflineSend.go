package tbsdk

// TaobaoLogisticsOfflineSendRequest 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOfflineSendRequest struct {
    
    /* cancel_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址 */
    cancel_id int64 `json:"cancel_id";xml:"cancel_id"`
    
    /* company_code required物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。 */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* feature optionalfeature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。 */
    feature string `json:"feature";xml:"feature"`
    
    /* is_split optional表明是否是拆单，默认值0，1表示拆单 */
    is_split int64 `json:"is_split";xml:"is_split"`
    
    /* out_sid required运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
    out_sid string `json:"out_sid";xml:"out_sid"`
    
    /* seller_ip optional商家的IP地址 */
    seller_ip string `json:"seller_ip";xml:"seller_ip"`
    
    /* sender_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址 */
    sender_id int64 `json:"sender_id";xml:"sender_id"`
    
    /* sub_tid optional需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。 */
    sub_tid int64 `json:"sub_tid";xml:"sub_tid"`
    
    /* tid required淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOfflineSendRequest) GetAPIName() string {
	return "taobao.logistics.offline.send"
}

// TaobaoLogisticsOfflineSendResponse 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOfflineSendResponse struct {
    
    /* shipping Object自己联系的调用结果 */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}
