package tbsdk

// TaobaoLogisticsOrderCreateRequest 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
type TaobaoLogisticsOrderCreateRequest struct {
    
    /* goods_names required运送的货物名称列表，用|号隔开 */
    goods_names string `json:"goods_names";xml:"goods_names"`
    
    /* goods_quantities required运送货物的数量列表，用|号隔开 */
    goods_quantities string `json:"goods_quantities";xml:"goods_quantities"`
    
    /* is_consign optional创建订单同时是否进行发货，默认发货。 */
    is_consign bool `json:"is_consign";xml:"is_consign"`
    
    /* item_values required运送货物的单价列表(注意：单位为分），用|号隔开 */
    item_values string `json:"item_values";xml:"item_values"`
    
    /* logis_company_code special发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。 */
    logis_company_code string `json:"logis_company_code";xml:"logis_company_code"`
    
    /* logis_type optional发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。 */
    logis_type string `json:"logis_type";xml:"logis_type"`
    
    /* mail_no special发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* order_type optional物流发货类型：1-普通订单
不填为默认类型 1-普通订单 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* seller_wangwang_id required卖家旺旺号 */
    seller_wangwang_id string `json:"seller_wangwang_id";xml:"seller_wangwang_id"`
    
    /* shipping optional运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。 */
    shipping int64 `json:"shipping";xml:"shipping"`
    
    /* trade_id optional订单的交易号码 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoLogisticsOrderCreateRequest) GetAPIName() string {
	return "taobao.logistics.order.create"
}

// TaobaoLogisticsOrderCreateResponse 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
type TaobaoLogisticsOrderCreateResponse struct {
    
    /* trade_id Basic淘宝物流订单交易号，如返回-1则表示错误。如果在新建订单时传入trade_id,此处会返回此id，如果未传入trade_id，此处会返回淘宝物流分配的交易号码。 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}
