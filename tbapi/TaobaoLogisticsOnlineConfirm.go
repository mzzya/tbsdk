package tbsdk

// TaobaoLogisticsOnlineConfirmRequest <br><font color='red'>仅使用taobao.logistics.online.send 发货时，未输入运单号的情况下，需要使用该接口确认发货。发货接口主要有taobao.logistics.offline.send 和taobao.logistics.online.send <br>
确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外
type TaobaoLogisticsOnlineConfirmRequest struct {
    
    /* is_split optional表明是否是拆单，默认值0，1表示拆单 */
    is_split int64 `json:"is_split";xml:"is_split"`
    
    /* out_sid required运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
    out_sid string `json:"out_sid";xml:"out_sid"`
    
    /* seller_ip optional商家的IP地址 */
    seller_ip string `json:"seller_ip";xml:"seller_ip"`
    
    /* sub_tid optional拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集 */
    sub_tid int64 `json:"sub_tid";xml:"sub_tid"`
    
    /* tid required淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOnlineConfirmRequest) GetAPIName() string {
	return "taobao.logistics.online.confirm"
}

// TaobaoLogisticsOnlineConfirmResponse <br><font color='red'>仅使用taobao.logistics.online.send 发货时，未输入运单号的情况下，需要使用该接口确认发货。发货接口主要有taobao.logistics.offline.send 和taobao.logistics.online.send <br>
确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外
type TaobaoLogisticsOnlineConfirmResponse struct {
    
    /* shipping Object只返回is_success：是否成功。 */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}
