package tbsdk

import "time"


/* TaobaoRegionPriceManageRequest 编辑区域价格 */
type TaobaoRegionPriceManageRequest struct {
    
    /* is_full optionaltrue:全量, false:增量 */
    is_full bool `json:"is_full";xml:"is_full"`
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* regional_price_dtos required列表 */
    regional_price_dtos RegionalPriceDto `json:"regional_price_dtos";xml:"regional_price_dtos"`
    
    /* sku_id optional无sku传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceManageRequest) GetAPIName() string {
	return "taobao.region.price.manage"
}

/* TaobaoRegionPriceManageResponse 编辑区域价格 */
type TaobaoRegionPriceManageResponse struct {
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoTmcQueueGetRequest 根据appkey和groupName获取消息队列积压情况 */
type TaobaoTmcQueueGetRequest struct {
    
    /* group_name requiredTMC组名 */
    group_name string `json:"group_name";xml:"group_name"`
    
}

func (req *TaobaoTmcQueueGetRequest) GetAPIName() string {
	return "taobao.tmc.queue.get"
}

/* TaobaoTmcQueueGetResponse 根据appkey和groupName获取消息队列积压情况 */
type TaobaoTmcQueueGetResponse struct {
    
    /* datas Object Array队列详细信息 */
    datas TmcQueueInfo `json:"datas";xml:"datas"`
    
}

/* TaobaoWlbWmsReturnOrderNotifyRequest 销售退货通知 */
type TaobaoWlbWmsReturnOrderNotifyRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* extend_fields optional扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-” */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* order_code optionalERP单据编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time optionalERP订单创建时间 */
    order_create_time time.Time `json:"order_create_time";xml:"order_create_time"`
    
    /* order_flag optional用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票 */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list optional商品信息列表 */
    order_item_list Orderitemlistwlbwmsreturnordernotify `json:"order_item_list";xml:"order_item_list"`
    
    /* order_source optional订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他 */
    order_source string `json:"order_source";xml:"order_source"`
    
    /* order_type optional订单类型 501 销退入库 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* owner_user_id optional货主ID */
    owner_user_id string `json:"owner_user_id";xml:"owner_user_id"`
    
    /* prev_order_code optional来源单据号，销售退货时填充原发货的LBX号 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收件人信息 */
    receiver_info Receiverinfowlbwmsreturnordernotify `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* return_reason optional销退时请提供退货的原因 */
    return_reason string `json:"return_reason";xml:"return_reason"`
    
    /* sender_info optional发件人信息 */
    sender_info Senderinfowlbwmsreturnordernotify `json:"sender_info";xml:"sender_info"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tms_order_code optional运单号 */
    tms_order_code string `json:"tms_order_code";xml:"tms_order_code"`
    
    /* tms_service_code optional快递公司编码 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* tms_service_name optional快递公司名称 */
    tms_service_name string `json:"tms_service_name";xml:"tms_service_name"`
    
}

func (req *TaobaoWlbWmsReturnOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.return.order.notify"
}

/* TaobaoWlbWmsReturnOrderNotifyResponse 销售退货通知 */
type TaobaoWlbWmsReturnOrderNotifyResponse struct {
    
    /* create_time Basic订单创建时间 */
    create_time time.Time `json:"create_time";xml:"create_time"`
    
    /* store_order_code BasicLBX929829111 */
    store_order_code string `json:"store_order_code";xml:"store_order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoRegionPriceQueryRequest 区域价格查询 */
type TaobaoRegionPriceQueryRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* regional_price_dtos optional不传则返回所有设置的区域价格 */
    regional_price_dtos RegionalPriceDto `json:"regional_price_dtos";xml:"regional_price_dtos"`
    
    /* sku_id optional无sku可传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceQueryRequest) GetAPIName() string {
	return "taobao.region.price.query"
}

/* TaobaoRegionPriceQueryResponse 区域价格查询 */
type TaobaoRegionPriceQueryResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}

/* TaobaoJdsTradeTracesGetRequest 获取聚石塔数据共享的交易全链路信息 */
type TaobaoJdsTradeTracesGetRequest struct {
    
    /* return_user_status optional是否返回用户的状态(是否存在) */
    return_user_status bool `json:"return_user_status";xml:"return_user_status"`
    
    /* tid required淘宝的订单tid */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoJdsTradeTracesGetRequest) GetAPIName() string {
	return "taobao.jds.trade.traces.get"
}

/* TaobaoJdsTradeTracesGetResponse 获取聚石塔数据共享的交易全链路信息 */
type TaobaoJdsTradeTracesGetResponse struct {
    
    /* traces Object Array跟踪信息列表 */
    traces TradeTrace `json:"traces";xml:"traces"`
    
    /* user_status Basic在订单全链路系统中的状态(比如是否存在) */
    user_status string `json:"user_status";xml:"user_status"`
    
}

/* TaobaoTmcUserTopicsGetRequest 获取用户开通的topic列表 */
type TaobaoTmcUserTopicsGetRequest struct {
    
    /* nick optional卖家nick */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoTmcUserTopicsGetRequest) GetAPIName() string {
	return "taobao.tmc.user.topics.get"
}

/* TaobaoTmcUserTopicsGetResponse 获取用户开通的topic列表 */
type TaobaoTmcUserTopicsGetResponse struct {
    
    /* result_code Basic错误码 */
    result_code string `json:"result_code";xml:"result_code"`
    
    /* result_message Basic错误信息 */
    result_message string `json:"result_message";xml:"result_message"`
    
    /* topics Basic Arraytopic列表 */
    topics string `json:"topics";xml:"topics"`
    
}

/* TaobaoWlbWmsSkuGetRequest 商品信息查询 */
type TaobaoWlbWmsSkuGetRequest struct {
    
    /* item_code optional菜鸟商品ID,与itemcode必须有一个值不为空 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* item_id optional商家商品编码,与itemid必须有一个值不为空 */
    item_id string `json:"item_id";xml:"item_id"`
    
    /* owner_user_id optional货主ID */
    owner_user_id string `json:"owner_user_id";xml:"owner_user_id"`
    
}

func (req *TaobaoWlbWmsSkuGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.sku.get"
}

/* TaobaoWlbWmsSkuGetResponse 商品信息查询 */
type TaobaoWlbWmsSkuGetResponse struct {
    
    /* advent_lifecycle Basic保质期预警天数 */
    advent_lifecycle int64 `json:"advent_lifecycle";xml:"advent_lifecycle"`
    
    /* approval_number Basic批准文号 */
    approval_number string `json:"approval_number";xml:"approval_number"`
    
    /* bar_code Basic条形码，多条码请用”;”分隔； */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* brand Basic品牌编码 */
    brand string `json:"brand";xml:"brand"`
    
    /* brand_name Basic品牌名称 */
    brand_name string `json:"brand_name";xml:"brand_name"`
    
    /* category Basic商品类别编码（外部系统类别） */
    category string `json:"category";xml:"category"`
    
    /* category_name Basic商品类别名称 */
    category_name string `json:"category_name";xml:"category_name"`
    
    /* color Basic颜色 */
    color string `json:"color";xml:"color"`
    
    /* cost_price Basic成本价，单位分 */
    cost_price int64 `json:"cost_price";xml:"cost_price"`
    
    /* extend_fields Basic拓展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-” */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* gross_weight Basic毛重，单位克 */
    gross_weight int64 `json:"gross_weight";xml:"gross_weight"`
    
    /* height Basic高度，单位毫米 */
    height int64 `json:"height";xml:"height"`
    
    /* iitem_code Basic商家商品编码,与itemid必须有一个值不为空 */
    iitem_code string `json:"iitem_code";xml:"iitem_code"`
    
    /* is_area_sale Basic是否区域销售 */
    is_area_sale bool `json:"is_area_sale";xml:"is_area_sale"`
    
    /* is_batch_mgt Basic是否启用批次管理 */
    is_batch_mgt bool `json:"is_batch_mgt";xml:"is_batch_mgt"`
    
    /* is_danger Basic是否危险品 */
    is_danger bool `json:"is_danger";xml:"is_danger"`
    
    /* is_hygroscopic Basic是否易碎品 */
    is_hygroscopic bool `json:"is_hygroscopic";xml:"is_hygroscopic"`
    
    /* is_shelflife Basic是否启用保质期管理 */
    is_shelflife bool `json:"is_shelflife";xml:"is_shelflife"`
    
    /* is_sn_mgt Basic是否启用序列号管理 */
    is_sn_mgt bool `json:"is_sn_mgt";xml:"is_sn_mgt"`
    
    /* item_id Basic菜鸟商品ID,与itemcode必须有一个值不为空 */
    item_id string `json:"item_id";xml:"item_id"`
    
    /* item_price Basic零售价，单位分 */
    item_price int64 `json:"item_price";xml:"item_price"`
    
    /* length Basic长度，单位毫米 */
    length int64 `json:"length";xml:"length"`
    
    /* lifecycle Basic保质期天数 */
    lifecycle int64 `json:"lifecycle";xml:"lifecycle"`
    
    /* lockup_lifecycle Basic保质期禁售天数 */
    lockup_lifecycle int64 `json:"lockup_lifecycle";xml:"lockup_lifecycle"`
    
    /* name Basic商品名称 */
    name string `json:"name";xml:"name"`
    
    /* net_weight Basic净重，单位克 */
    net_weight int64 `json:"net_weight";xml:"net_weight"`
    
    /* origin_address Basic场地 */
    origin_address int64 `json:"origin_address";xml:"origin_address"`
    
    /* pcs Basic箱规 */
    pcs int64 `json:"pcs";xml:"pcs"`
    
    /* reject_lifecycle Basic保质期禁收天数 */
    reject_lifecycle int64 `json:"reject_lifecycle";xml:"reject_lifecycle"`
    
    /* size Basic尺寸 */
    size string `json:"size";xml:"size"`
    
    /* specification Basic规格 */
    specification string `json:"specification";xml:"specification"`
    
    /* tag_price Basic吊牌价，单位分 */
    tag_price int64 `json:"tag_price";xml:"tag_price"`
    
    /* title Basic商品标题 */
    title string `json:"title";xml:"title"`
    
    /* _type Basic商品类别 NORMAL：普通商品、 COMBINE：组合商品、 DISTRIBUTION：分销商品 */
    _type string `json:"type";xml:"type"`
    
    /* use_yn Basic启用标识 */
    use_yn bool `json:"use_yn";xml:"use_yn"`
    
    /* volume Basic体积，单位立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* width Basic宽度，单位毫米 */
    width int64 `json:"width";xml:"width"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoQimenInventoryReportRequest WMS调用奇门的接口,将库存盘点情况回传ERP */
type TaobaoQimenInventoryReportRequest struct {
    
    /* adjustType required变动类型：CHECK=盘点 ADJUST=调整 */
    adjustType string `json:"adjustType";xml:"adjustType"`
    
    /* checkOrderCode required盘点单编码 */
    checkOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    /* checkOrderId optional仓储系统的盘点单编码 */
    checkOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    /* checkTime optional盘点时间(YYYY-MM-DD HH:MM:SS) */
    checkTime string `json:"checkTime";xml:"checkTime"`
    
    /* currentPage required当前页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品库存信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* orderType optionalorderType */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* outBizCode required外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理) */
    outBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* pageSize required每页记录的条数 */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* totalPage required总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
    /* warehouseCode required仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventoryReportRequest) GetAPIName() string {
	return "taobao.qimen.inventory.report"
}

/* TaobaoQimenInventoryReportResponse WMS调用奇门的接口,将库存盘点情况回传ERP */
type TaobaoQimenInventoryReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenInventorycheckQueryRequest ERP调用奇门的接口,查询库存盘点情况 */
type TaobaoQimenInventorycheckQueryRequest struct {
    
    /* checkOrderCode required盘点单编码 */
    checkOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    /* checkOrderId optional仓储系统的盘点单编码 */
    checkOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventorycheckQueryRequest) GetAPIName() string {
	return "taobao.qimen.inventorycheck.query"
}

/* TaobaoQimenInventorycheckQueryResponse ERP调用奇门的接口,查询库存盘点情况 */
type TaobaoQimenInventorycheckQueryResponse struct {
    
    /* checkOrderCode Basic盘点单编码 */
    checkOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    /* checkOrderId Basic仓储系统的盘点单编码 */
    checkOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    /* checkTime Basic盘点时间(YYYY-MM-DD HH:MM:SS) */
    checkTime string `json:"checkTime";xml:"checkTime"`
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品库存列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* ownerCode Basic货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* remark Basic备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
    /* warehouseCode Basic仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

/* TaobaoQimenItemsSynchronizeRequest ERP调用奇门的接口,批量同步商品信息给WMS */
type TaobaoQimenItemsSynchronizeRequest struct {
    
    /* actionType required操作类型(两种类型：add|update) */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional同步的商品信息 */
    items Item `json:"items";xml:"items"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemsSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.items.synchronize"
}

/* TaobaoQimenItemsSynchronizeResponse ERP调用奇门的接口,批量同步商品信息给WMS */
type TaobaoQimenItemsSynchronizeResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品同步批量接口中单商品信息 */
    items BatchItemSynItem `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenWavenumReportRequest WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求 */
type TaobaoQimenWavenumReportRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional发货单号 */
    orders Order `json:"orders";xml:"orders"`
    
    /* waveNum required波次号 */
    waveNum string `json:"waveNum";xml:"waveNum"`
    
}

func (req *TaobaoQimenWavenumReportRequest) GetAPIName() string {
	return "taobao.qimen.wavenum.report"
}

/* TaobaoQimenWavenumReportResponse WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求 */
type TaobaoQimenWavenumReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoLogisticsOrderCreateRequest 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。 */
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

/* TaobaoLogisticsOrderCreateResponse 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。 */
type TaobaoLogisticsOrderCreateResponse struct {
    
    /* trade_id Basic淘宝物流订单交易号，如返回-1则表示错误。如果在新建订单时传入trade_id,此处会返回此id，如果未传入trade_id，此处会返回淘宝物流分配的交易号码。 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}

/* TaobaoItempropvaluesGetRequest 获取标准类目属性值 */
type TaobaoItempropvaluesGetRequest struct {
    
    /* attr_keys optional属性的Key，支持多条，以“,”分隔 */
    attr_keys string `json:"attr_keys";xml:"attr_keys"`
    
    /* cid required叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID */
    cid int64 `json:"cid";xml:"cid"`
    
    /* datetime special假如传2005-01-01 00:00:00，则取所有的属性和子属性(状态为删除的属性值不返回prop_name) */
    datetime time.Time `json:"datetime";xml:"datetime"`
    
    /* fields required需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order */
    fields string `json:"fields";xml:"fields"`
    
    /* pvs special属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2) */
    pvs string `json:"pvs";xml:"pvs"`
    
    /* _type optional获取类目的类型：1代表集市、2代表天猫 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItempropvaluesGetRequest) GetAPIName() string {
	return "taobao.itempropvalues.get"
}

/* TaobaoItempropvaluesGetResponse 获取标准类目属性值 */
type TaobaoItempropvaluesGetResponse struct {
    
    /* last_modified Basic最近修改时间。格式:yyyy-MM-dd HH:mm:ss */
    last_modified string `json:"last_modified";xml:"last_modified"`
    
    /* prop_values Object Array属性值,根据fields传入的参数返回相应的结果 */
    prop_values PropValue `json:"prop_values";xml:"prop_values"`
    
}

/* Gfq1qp4287ShopexAreaInfoGetRequest 地区编码同步 */
type Gfq1qp4287ShopexAreaInfoGetRequest struct {
    
    /* qimen_body optional请求参数 */
    qimen_body string `json:"qimen_body";xml:"qimen_body"`
    
}

func (req *Gfq1qp4287ShopexAreaInfoGetRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.area.info.get"
}

/* Gfq1qp4287ShopexAreaInfoGetResponse 地区编码同步 */
type Gfq1qp4287ShopexAreaInfoGetResponse struct {
    
    /* data Basic返回内容 */
    data string `json:"data";xml:"data"`
    
    /* msg_id Basic请求任务id */
    msg_id string `json:"msg_id";xml:"msg_id"`
    
    /* res Basic描述 */
    res string `json:"res";xml:"res"`
    
    /* rsp Basic成功失败标识 */
    rsp string `json:"rsp";xml:"rsp"`
    
}

/* TaobaoTmcMessagesConfirmRequest 确认消费消息的状态 */
type TaobaoTmcMessagesConfirmRequest struct {
    
    /* f_message_ids optional处理失败的消息ID列表--已废弃，无需传此字段 */
    f_message_ids int64 `json:"f_message_ids";xml:"f_message_ids"`
    
    /* group_name optional分组名称，不传代表默认分组 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* s_message_ids required处理成功的消息ID列表 最大 200个ID */
    s_message_ids int64 `json:"s_message_ids";xml:"s_message_ids"`
    
}

func (req *TaobaoTmcMessagesConfirmRequest) GetAPIName() string {
	return "taobao.tmc.messages.confirm"
}

/* TaobaoTmcMessagesConfirmResponse 确认消费消息的状态 */
type TaobaoTmcMessagesConfirmResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoTmcMessagesConsumeRequest 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。 */
type TaobaoTmcMessagesConsumeRequest struct {
    
    /* group_name optional用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* quantity optional每次批量消费消息的条数 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoTmcMessagesConsumeRequest) GetAPIName() string {
	return "taobao.tmc.messages.consume"
}

/* TaobaoTmcMessagesConsumeResponse 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。 */
type TaobaoTmcMessagesConsumeResponse struct {
    
    /* messages Object Array消息列表 */
    messages TmcMessage `json:"messages";xml:"messages"`
    
}

/* TaobaoTmcMessageProduceRequest 发布单条消息 */
type TaobaoTmcMessageProduceRequest struct {
    
    /* content required消息内容的JSON表述，必须按照topic的定义来填充 */
    content string `json:"content";xml:"content"`
    
    /* ex_content optional消息的扩增属性，用json格式表示 */
    ex_content string `json:"ex_content";xml:"ex_content"`
    
    /* media_content optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。 */
    media_content []byte `json:"media_content";xml:"media_content"`
    
    /* media_content2 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content2 []byte `json:"media_content2";xml:"media_content2"`
    
    /* media_content3 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content3 []byte `json:"media_content3";xml:"media_content3"`
    
    /* media_content4 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content4 []byte `json:"media_content4";xml:"media_content4"`
    
    /* media_content5 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content5 []byte `json:"media_content5";xml:"media_content5"`
    
    /* target_appkey optional直发消息需要传入目标appkey */
    target_appkey string `json:"target_appkey";xml:"target_appkey"`
    
    /* target_group optional目标分组，一般为default */
    target_group string `json:"target_group";xml:"target_group"`
    
    /* topic required消息类型 */
    topic string `json:"topic";xml:"topic"`
    
}

func (req *TaobaoTmcMessageProduceRequest) GetAPIName() string {
	return "taobao.tmc.message.produce"
}

/* TaobaoTmcMessageProduceResponse 发布单条消息 */
type TaobaoTmcMessageProduceResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* msg_ids Basic Array消息ID */
    msg_ids string `json:"msg_ids";xml:"msg_ids"`
    
    /* total Basic投递目标数 */
    total int64 `json:"total";xml:"total"`
    
}

/* TaobaoItemsOnsaleGetRequest 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取 */
type TaobaoItemsOnsaleGetRequest struct {
    
    /* auction_type optional商品类型：a-拍卖,b-一口价 */
    auction_type string `json:"auction_type";xml:"auction_type"`
    
    /* cid optional商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到 */
    cid int64 `json:"cid";xml:"cid"`
    
    /* end_modified optional结束的修改时间 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* fields required需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。 */
    fields string `json:"fields";xml:"fields"`
    
    /* has_discount optional是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
    has_discount bool `json:"has_discount";xml:"has_discount"`
    
    /* has_showcase optional是否橱窗推荐。 可选值：true，false。默认不过滤该条件 */
    has_showcase bool `json:"has_showcase";xml:"has_showcase"`
    
    /* is_combine optional组合商品 */
    is_combine bool `json:"is_combine";xml:"is_combine"`
    
    /* is_cspu optional是否挂接了达尔文标准产品体系 */
    is_cspu bool `json:"is_cspu";xml:"is_cspu"`
    
    /* is_ex optional商品是否在外部网店显示 */
    is_ex bool `json:"is_ex";xml:"is_ex"`
    
    /* is_taobao optional商品是否在淘宝显示 */
    is_taobao bool `json:"is_taobao";xml:"is_taobao"`
    
    /* order_by optional排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* q optional搜索字段。搜索商品的title。 */
    q string `json:"q";xml:"q"`
    
    /* seller_cids optional卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
    seller_cids string `json:"seller_cids";xml:"seller_cids"`
    
    /* start_modified optional起始的修改时间 */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoItemsOnsaleGetRequest) GetAPIName() string {
	return "taobao.items.onsale.get"
}

/* TaobaoItemsOnsaleGetResponse 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取 */
type TaobaoItemsOnsaleGetResponse struct {
    
    /* items Object Array搜索到的商品列表，具体字段根据设定的fields决定，不包括desc字段 */
    items Item `json:"items";xml:"items"`
    
    /* total_results Basic搜索到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* CainiaoCrmOmsRuleSyncRequest 将商家ERP订单处理规则同步到菜鸟CRM系统 */
type CainiaoCrmOmsRuleSyncRequest struct {
    
    /* check_rule_msg optional人工审单规则描述 */
    check_rule_msg string `json:"check_rule_msg";xml:"check_rule_msg"`
    
    /* is_auto_check required是否系统智能处理订单（无人工介入） */
    is_auto_check bool `json:"is_auto_check";xml:"is_auto_check"`
    
    /* is_open_cnauto required是否开启菜鸟自动流转规则 */
    is_open_cnauto bool `json:"is_open_cnauto";xml:"is_open_cnauto"`
    
    /* is_sys_merge_order optional是否开启了订单合单 */
    is_sys_merge_order bool `json:"is_sys_merge_order";xml:"is_sys_merge_order"`
    
    /* merge_order_cycle optional订单合单周期（单位：分钟） */
    merge_order_cycle int64 `json:"merge_order_cycle";xml:"merge_order_cycle"`
    
    /* other_rule optional其他未定义订单处理规则，格式｛name;stauts;cycle;｝ */
    other_rule string `json:"other_rule";xml:"other_rule"`
    
    /* shop_code required店铺nick */
    shop_code string `json:"shop_code";xml:"shop_code"`
    
}

func (req *CainiaoCrmOmsRuleSyncRequest) GetAPIName() string {
	return "cainiao.crm.oms.rule.sync"
}

/* CainiaoCrmOmsRuleSyncResponse 将商家ERP订单处理规则同步到菜鸟CRM系统 */
type CainiaoCrmOmsRuleSyncResponse struct {
    
    /* wl_error_code BasicerrorCode */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg BasicerrorMsg */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basicsuccess */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoItemUpdateRequest 根据传入的num_iid更新对应的商品的数据 
传入的num_iid所对应的商品必须属于当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。 */
type TaobaoItemUpdateRequest struct {
    
    /* after_sale_id optional售后说明模板id */
    after_sale_id int64 `json:"after_sale_id";xml:"after_sale_id"`
    
    /* approve_status optional商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true */
    approve_status string `json:"approve_status";xml:"approve_status"`
    
    /* auction_point optional商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
    auction_point int64 `json:"auction_point";xml:"auction_point"`
    
    /* auto_fill optional代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充) */
    auto_fill string `json:"auto_fill";xml:"auto_fill"`
    
    /* auto_repost optional自动重发。可选值:true,false; */
    auto_repost bool `json:"auto_repost";xml:"auto_repost"`
    
    /* barcode optional商品条形码 */
    barcode string `json:"barcode";xml:"barcode"`
    
    /* change_prop optional基础色数据 */
    change_prop string `json:"change_prop";xml:"change_prop"`
    
    /* chaoshi_extends_info optional天猫超市扩展字段，天猫超市专用。 */
    chaoshi_extends_info string `json:"chaoshi_extends_info";xml:"chaoshi_extends_info"`
    
    /* cid optional叶子类目id */
    cid int64 `json:"cid";xml:"cid"`
    
    /* cod_postage_id optional货到付款运费模板ID该字段已经废弃，货到付款模板已经集成到运费模板中。 */
    cod_postage_id int64 `json:"cod_postage_id";xml:"cod_postage_id"`
    
    /* cpv_memo optional针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷 */
    cpv_memo string `json:"cpv_memo";xml:"cpv_memo"`
    
    /* delivery_time_delivery_time optional商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11 */
    delivery_time_delivery_time string `json:"delivery_time.delivery_time";xml:"delivery_time.delivery_time"`
    
    /* delivery_time_delivery_time_type optional发货时间类型：绝对发货时间或者相对发货时间 */
    delivery_time_delivery_time_type string `json:"delivery_time.delivery_time_type";xml:"delivery_time.delivery_time_type"`
    
    /* delivery_time_need_delivery_time optional设置是否使用发货时间，商品级别，sku级别 */
    delivery_time_need_delivery_time string `json:"delivery_time.need_delivery_time";xml:"delivery_time.need_delivery_time"`
    
    /* desc optional商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制 */
    desc string `json:"desc";xml:"desc"`
    
    /* desc_modules optional商品描述模块化，模块列表，由List转化成jsonArray存入，后端逻辑验证通过，拼装成模块内容+锚点导航后存入desc中。数据结构具体参见Item_Desc_Module */
    desc_modules string `json:"desc_modules";xml:"desc_modules"`
    
    /* empty_fields optional支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock） */
    empty_fields string `json:"empty_fields";xml:"empty_fields"`
    
    /* ems_fee optionalems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
    ems_fee Price `json:"ems_fee";xml:"ems_fee"`
    
    /* express_fee optional快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
    express_fee Price `json:"express_fee";xml:"express_fee"`
    
    /* features optional宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp */
    features string `json:"features";xml:"features"`
    
    /* food_security_contact optional厂家联系方式 */
    food_security_contact string `json:"food_security.contact";xml:"food_security.contact"`
    
    /* food_security_design_code optional产品标准号 */
    food_security_design_code string `json:"food_security.design_code";xml:"food_security.design_code"`
    
    /* food_security_factory optional厂名 */
    food_security_factory string `json:"food_security.factory";xml:"food_security.factory"`
    
    /* food_security_factory_site optional厂址 */
    food_security_factory_site string `json:"food_security.factory_site";xml:"food_security.factory_site"`
    
    /* food_security_food_additive optional食品添加剂 */
    food_security_food_additive string `json:"food_security.food_additive";xml:"food_security.food_additive"`
    
    /* food_security_health_product_no optional健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
    food_security_health_product_no string `json:"food_security.health_product_no";xml:"food_security.health_product_no"`
    
    /* food_security_mix optional配料表 */
    food_security_mix string `json:"food_security.mix";xml:"food_security.mix"`
    
    /* food_security_period optional保质期 */
    food_security_period string `json:"food_security.period";xml:"food_security.period"`
    
    /* food_security_plan_storage optional储藏方法 */
    food_security_plan_storage string `json:"food_security.plan_storage";xml:"food_security.plan_storage"`
    
    /* food_security_prd_license_no optional生产许可证号 */
    food_security_prd_license_no string `json:"food_security.prd_license_no";xml:"food_security.prd_license_no"`
    
    /* food_security_product_date_end optional生产结束日期,格式必须为yyyy-MM-dd */
    food_security_product_date_end string `json:"food_security.product_date_end";xml:"food_security.product_date_end"`
    
    /* food_security_product_date_start optional生产开始日期，格式必须为yyyy-MM-dd */
    food_security_product_date_start string `json:"food_security.product_date_start";xml:"food_security.product_date_start"`
    
    /* food_security_stock_date_end optional进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
    food_security_stock_date_end string `json:"food_security.stock_date_end";xml:"food_security.stock_date_end"`
    
    /* food_security_stock_date_start optional进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
    food_security_stock_date_start string `json:"food_security.stock_date_start";xml:"food_security.stock_date_start"`
    
    /* food_security_supplier optional供货商 */
    food_security_supplier string `json:"food_security.supplier";xml:"food_security.supplier"`
    
    /* freight_payer optional运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担); */
    freight_payer string `json:"freight_payer";xml:"freight_payer"`
    
    /* global_stock_country optional全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他） */
    global_stock_country string `json:"global_stock_country";xml:"global_stock_country"`
    
    /* global_stock_delivery_place optional全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台” */
    global_stock_delivery_place string `json:"global_stock_delivery_place";xml:"global_stock_delivery_place"`
    
    /* global_stock_tax_free_promise optional全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。 */
    global_stock_tax_free_promise bool `json:"global_stock_tax_free_promise";xml:"global_stock_tax_free_promise"`
    
    /* global_stock_type optional全球购商品采购地（库存类型）全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
    global_stock_type string `json:"global_stock_type";xml:"global_stock_type"`
    
    /* has_discount optional支持会员打折。可选值:true,false; */
    has_discount bool `json:"has_discount";xml:"has_discount"`
    
    /* has_invoice optional是否有发票。可选值:true,false (商城卖家此字段必须为true) */
    has_invoice bool `json:"has_invoice";xml:"has_invoice"`
    
    /* has_showcase optional橱窗推荐。可选值:true,false; */
    has_showcase bool `json:"has_showcase";xml:"has_showcase"`
    
    /* has_warranty optional是否有保修。可选值:true,false; */
    has_warranty bool `json:"has_warranty";xml:"has_warranty"`
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* image optional商品图片。类型:JPG,GIF;最大长度:3M */
    image []byte `json:"image";xml:"image"`
    
    /* increment optional加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
    increment Price `json:"increment";xml:"increment"`
    
    /* input_custom_cpv optional针对当前商品的自定义属性值 */
    input_custom_cpv string `json:"input_custom_cpv";xml:"input_custom_cpv"`
    
    /* input_pids optional用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
    input_pids string `json:"input_pids";xml:"input_pids"`
    
    /* input_str optional用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
    input_str string `json:"input_str";xml:"input_str"`
    
    /* interactive_id optional主图视频互动信息id，必须有主图视频id才能传互动信息id */
    interactive_id int64 `json:"interactive_id";xml:"interactive_id"`
    
    /* is_3D optional是否是3D */
    is_3D bool `json:"is_3D";xml:"is_3D"`
    
    /* is_ex optional是否在外店显示 */
    is_ex bool `json:"is_ex";xml:"is_ex"`
    
    /* is_lightning_consignment optional实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记 */
    is_lightning_consignment bool `json:"is_lightning_consignment";xml:"is_lightning_consignment"`
    
    /* is_offline optional是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。 */
    is_offline string `json:"is_offline";xml:"is_offline"`
    
    /* is_replace_sku optional是否替换sku */
    is_replace_sku bool `json:"is_replace_sku";xml:"is_replace_sku"`
    
    /* is_taobao optional是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
    is_taobao bool `json:"is_taobao";xml:"is_taobao"`
    
    /* is_xinpin optional商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。 */
    is_xinpin bool `json:"is_xinpin";xml:"is_xinpin"`
    
    /* item_size optional表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0 */
    item_size string `json:"item_size";xml:"item_size"`
    
    /* item_weight optional商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0 */
    item_weight string `json:"item_weight";xml:"item_weight"`
    
    /* lang optional商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN” */
    lang string `json:"lang";xml:"lang"`
    
    /* lease_extends_info optional淘宝租赁扩展信息 */
    lease_extends_info string `json:"lease_extends_info";xml:"lease_extends_info"`
    
    /* list_time optional上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。 */
    list_time time.Time `json:"list_time";xml:"list_time"`
    
    /* locality_life_choose_logis optional编辑电子凭证宝贝时候表示是否使用邮寄0: 代表不使用邮寄；1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
    locality_life_choose_logis string `json:"locality_life.choose_logis";xml:"locality_life.choose_logis"`
    
    /* locality_life_eticket optional电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4 */
    locality_life_eticket string `json:"locality_life.eticket";xml:"locality_life.eticket"`
    
    /* locality_life_expirydate optional本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15 */
    locality_life_expirydate string `json:"locality_life.expirydate";xml:"locality_life.expirydate"`
    
    /* locality_life_merchant optional码商信息，格式为 码商id:nick */
    locality_life_merchant string `json:"locality_life.merchant";xml:"locality_life.merchant"`
    
    /* locality_life_network_id optional网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID */
    locality_life_network_id string `json:"locality_life.network_id";xml:"locality_life.network_id"`
    
    /* locality_life_obs optional预约门店是否支持门店自提,1:是 */
    locality_life_obs string `json:"locality_life.obs";xml:"locality_life.obs"`
    
    /* locality_life_onsale_auto_refund_ratio optional电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
    locality_life_onsale_auto_refund_ratio int64 `json:"locality_life.onsale_auto_refund_ratio";xml:"locality_life.onsale_auto_refund_ratio"`
    
    /* locality_life_packageid optional新版电子凭证包id */
    locality_life_packageid string `json:"locality_life.packageid";xml:"locality_life.packageid"`
    
    /* locality_life_refund_ratio optional退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例 */
    locality_life_refund_ratio int64 `json:"locality_life.refund_ratio";xml:"locality_life.refund_ratio"`
    
    /* locality_life_refundmafee optional退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
    locality_life_refundmafee string `json:"locality_life.refundmafee";xml:"locality_life.refundmafee"`
    
    /* locality_life_verification optional核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款 */
    locality_life_verification string `json:"locality_life.verification";xml:"locality_life.verification"`
    
    /* locality_life_version optional电子凭证版本 新版电子凭证值:1 */
    locality_life_version string `json:"locality_life.version";xml:"locality_life.version"`
    
    /* location_city optional所在地城市。如杭州 */
    location_city string `json:"location.city";xml:"location.city"`
    
    /* location_state optional所在地省份。如浙江 */
    location_state string `json:"location.state";xml:"location.state"`
    
    /* ms_payment_price optional订金 */
    ms_payment_price string `json:"ms_payment.price";xml:"ms_payment.price"`
    
    /* ms_payment_reference_price optional参考价 */
    ms_payment_reference_price string `json:"ms_payment.reference_price";xml:"ms_payment.reference_price"`
    
    /* ms_payment_voucher_price optional尾款可抵扣金额 */
    ms_payment_voucher_price string `json:"ms_payment.voucher_price";xml:"ms_payment.voucher_price"`
    
    /* newprepay optional该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；<br>注意：使用该API修改商品其它属性如标题title时，如需保持商品不支持7天无理由退货状态，该字段需传入0 。 */
    newprepay string `json:"newprepay";xml:"newprepay"`
    
    /* num optional商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。 */
    num int64 `json:"num";xml:"num"`
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* o2o_bind_service optional汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。 */
    o2o_bind_service bool `json:"o2o_bind_service";xml:"o2o_bind_service"`
    
    /* outer_id optional商家编码 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* paimai_info_deposit optional拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。 */
    paimai_info_deposit int64 `json:"paimai_info.deposit";xml:"paimai_info.deposit"`
    
    /* paimai_info_interval optional降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。 */
    paimai_info_interval int64 `json:"paimai_info.interval";xml:"paimai_info.interval"`
    
    /* paimai_info_mode optional拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。 */
    paimai_info_mode int64 `json:"paimai_info.mode";xml:"paimai_info.mode"`
    
    /* paimai_info_reserve optional降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
    paimai_info_reserve Price `json:"paimai_info.reserve";xml:"paimai_info.reserve"`
    
    /* paimai_info_valid_hour optional自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
    paimai_info_valid_hour int64 `json:"paimai_info.valid_hour";xml:"paimai_info.valid_hour"`
    
    /* paimai_info_valid_minute optional自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
    paimai_info_valid_minute int64 `json:"paimai_info.valid_minute";xml:"paimai_info.valid_minute"`
    
    /* pic_path optional商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* post_fee optional平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写 */
    post_fee Price `json:"post_fee";xml:"post_fee"`
    
    /* postage_id optional宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.templates.get获得当前会话用户的所有邮费模板） */
    postage_id int64 `json:"postage_id";xml:"postage_id"`
    
    /* price optional商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。 */
    price Price `json:"price";xml:"price"`
    
    /* product_id optional商品所属的产品ID(B商家发布商品需要用) */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* property_alias optional属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。 */
    property_alias string `json:"property_alias";xml:"property_alias"`
    
    /* props optional商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。 */
    props string `json:"props";xml:"props"`
    
    /* qualification optional商品资质信息 */
    qualification string `json:"qualification";xml:"qualification"`
    
    /* scenic_ticket_book_cost optional景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
    scenic_ticket_book_cost string `json:"scenic_ticket_book_cost";xml:"scenic_ticket_book_cost"`
    
    /* scenic_ticket_pay_way optional景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
    scenic_ticket_pay_way int64 `json:"scenic_ticket_pay_way";xml:"scenic_ticket_pay_way"`
    
    /* sell_point optional商品卖点信息，最长150个字符。天猫和集市都可用 */
    sell_point string `json:"sell_point";xml:"sell_point"`
    
    /* sell_promise optional是否承诺退换货服务!虚拟商品无须设置此项! */
    sell_promise bool `json:"sell_promise";xml:"sell_promise"`
    
    /* seller_cids optional重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
    seller_cids string `json:"seller_cids";xml:"seller_cids"`
    
    /* shape optional宝贝形态:1代表电子券;0或其他代表实物 */
    shape string `json:"shape";xml:"shape"`
    
    /* sku_barcode optionalsku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔 */
    sku_barcode string `json:"sku_barcode";xml:"sku_barcode"`
    
    /* sku_delivery_times optional此参数暂时不起作用 */
    sku_delivery_times string `json:"sku_delivery_times";xml:"sku_delivery_times"`
    
    /* sku_hd_height optional家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。 可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50 */
    sku_hd_height string `json:"sku_hd_height";xml:"sku_hd_height"`
    
    /* sku_hd_lamp_quantity optional家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7 */
    sku_hd_lamp_quantity string `json:"sku_hd_lamp_quantity";xml:"sku_hd_lamp_quantity"`
    
    /* sku_hd_length optional家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：20,30,30 */
    sku_hd_length string `json:"sku_hd_length";xml:"sku_hd_length"`
    
    /* sku_outer_ids optionalSku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
    sku_outer_ids string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    /* sku_prices optional更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
    sku_prices string `json:"sku_prices";xml:"sku_prices"`
    
    /* sku_properties optional更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。 */
    sku_properties string `json:"sku_properties";xml:"sku_properties"`
    
    /* sku_quantities optional更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4 */
    sku_quantities string `json:"sku_quantities";xml:"sku_quantities"`
    
    /* sku_spec_ids optional此参数暂时不起作用 */
    sku_spec_ids string `json:"sku_spec_ids";xml:"sku_spec_ids"`
    
    /* spu_confirm optional手机类目spu 确认信息字段 */
    spu_confirm bool `json:"spu_confirm";xml:"spu_confirm"`
    
    /* stuff_status optional商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。 */
    stuff_status string `json:"stuff_status";xml:"stuff_status"`
    
    /* sub_stock optional商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存 */
    sub_stock int64 `json:"sub_stock";xml:"sub_stock"`
    
    /* title optional宝贝标题. 不能超过30字符,受违禁词控制 */
    title string `json:"title";xml:"title"`
    
    /* valid_thru optional有效期。可选值:7,14;单位:天; */
    valid_thru int64 `json:"valid_thru";xml:"valid_thru"`
    
    /* video_id optional主图视频id */
    video_id int64 `json:"video_id";xml:"video_id"`
    
    /* weight optional商品的重量(商超卖家专用字段) */
    weight int64 `json:"weight";xml:"weight"`
    
    /* wireless_desc optional无线的宝贝描述 */
    wireless_desc string `json:"wireless_desc";xml:"wireless_desc"`
    
}

func (req *TaobaoItemUpdateRequest) GetAPIName() string {
	return "taobao.item.update"
}

/* TaobaoItemUpdateResponse 根据传入的num_iid更新对应的商品的数据 
传入的num_iid所对应的商品必须属于当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。 */
type TaobaoItemUpdateResponse struct {
    
    /* item Object商品结构里的num_iid，modified */
    item Item `json:"item";xml:"item"`
    
}

/* TaobaoTmcUserCancelRequest 取消用户的消息服务 */
type TaobaoTmcUserCancelRequest struct {
    
    /* nick required用户昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* user_platform optional用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcUserCancelRequest) GetAPIName() string {
	return "taobao.tmc.user.cancel"
}

/* TaobaoTmcUserCancelResponse 取消用户的消息服务 */
type TaobaoTmcUserCancelResponse struct {
    
    /* is_success Basic是否成功,如果为false并且没有错误码，表示删除的用户不存在。 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoNextoneLogisticsWarehouseUpdateRequest 商家上传退货入仓状态给ag */
type TaobaoNextoneLogisticsWarehouseUpdateRequest struct {
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* warehouse_status required退货入仓状态 1.已入仓 */
    warehouse_status int64 `json:"warehouse_status";xml:"warehouse_status"`
    
}

func (req *TaobaoNextoneLogisticsWarehouseUpdateRequest) GetAPIName() string {
	return "taobao.nextone.logistics.warehouse.update"
}

/* TaobaoNextoneLogisticsWarehouseUpdateResponse 商家上传退货入仓状态给ag */
type TaobaoNextoneLogisticsWarehouseUpdateResponse struct {
    
    /* err_code BasicerrorCode */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_info BasicerrorInfo */
    err_info string `json:"err_info";xml:"err_info"`
    
    /* result_data BasicresultData */
    result_data map[string]interface{} `json:"result_data";xml:"result_data"`
    
    /* succeed Basicsuccess */
    succeed bool `json:"succeed";xml:"succeed"`
    
}

/* TaobaoItemImgUploadRequest 添加一张商品图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。 */
type TaobaoItemImgUploadRequest struct {
    
    /* id optional商品图片id(如果是更新图片，则需要传该参数) */
    id int64 `json:"id";xml:"id"`
    
    /* image optional商品图片内容类型:JPG,GIF;最大:3M 。支持的文件类型：gif,jpg,jpeg,png */
    image []byte `json:"image";xml:"image"`
    
    /* is_major optional是否将该图片设为主图,可选值:true,false;默认值:false(非主图) */
    is_major bool `json:"is_major";xml:"is_major"`
    
    /* is_rectangle optional是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图 */
    is_rectangle bool `json:"is_rectangle";xml:"is_rectangle"`
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* position optional图片序号 */
    position int64 `json:"position";xml:"position"`
    
}

func (req *TaobaoItemImgUploadRequest) GetAPIName() string {
	return "taobao.item.img.upload"
}

/* TaobaoItemImgUploadResponse 添加一张商品图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。 */
type TaobaoItemImgUploadResponse struct {
    
    /* item_img Object商品图片结构 */
    item_img ItemImg `json:"item_img";xml:"item_img"`
    
}

/* TaobaoTmcUserPermitRequest 为已授权的用户开通消息服务 */
type TaobaoTmcUserPermitRequest struct {
    
    /* topics optional消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。 */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcUserPermitRequest) GetAPIName() string {
	return "taobao.tmc.user.permit"
}

/* TaobaoTmcUserPermitResponse 为已授权的用户开通消息服务 */
type TaobaoTmcUserPermitResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TmallChannelProductsGetRequest 查询供应商的产品数据。 

* 入参传入pids将优先查询，即只按这个条件查询。 
*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
* 入参fields传入images将查询多图数据，不传只返回主图数据。 
* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
* 查询结果按照产品发布时间倒序，即时间近的数据在前。
* 传入channel 渠道，会只返回相应渠道的产品 */
type TmallChannelProductsGetRequest struct {
    
    /* top_query_product_d_o optionaltop_query_product_d_o */
    top_query_product_d_o TopQueryProductDo `json:"top_query_product_d_o";xml:"top_query_product_d_o"`
    
}

func (req *TmallChannelProductsGetRequest) GetAPIName() string {
	return "tmall.channel.products.get"
}

/* TmallChannelProductsGetResponse 查询供应商的产品数据。 

* 入参传入pids将优先查询，即只按这个条件查询。 
*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
* 入参fields传入images将查询多图数据，不传只返回主图数据。 
* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
* 查询结果按照产品发布时间倒序，即时间近的数据在前。
* 传入channel 渠道，会只返回相应渠道的产品 */
type TmallChannelProductsGetResponse struct {
    
    /* products Object Array产品对象记录集 */
    products TopProductDO `json:"products";xml:"products"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoItemAddRequest 此接口用于新增一个商品 
商品所属的卖家是当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。 */
type TaobaoItemAddRequest struct {
    
    /* after_sale_id optional售后说明模板id */
    after_sale_id int64 `json:"after_sale_id";xml:"after_sale_id"`
    
    /* approve_status optional商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale */
    approve_status string `json:"approve_status";xml:"approve_status"`
    
    /* auction_point optional商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
    auction_point int64 `json:"auction_point";xml:"auction_point"`
    
    /* auto_fill optional代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充) */
    auto_fill string `json:"auto_fill";xml:"auto_fill"`
    
    /* auto_repost optional自动重发。可选值:true,false;默认值:false(不重发) */
    auto_repost bool `json:"auto_repost";xml:"auto_repost"`
    
    /* barcode optional商品条形码 */
    barcode string `json:"barcode";xml:"barcode"`
    
    /* change_prop optional基础色数据 */
    change_prop string `json:"change_prop";xml:"change_prop"`
    
    /* chaoshi_extends_info optional天猫超市扩展字段，天猫超市专用。 */
    chaoshi_extends_info string `json:"chaoshi_extends_info";xml:"chaoshi_extends_info"`
    
    /* cid required叶子类目id */
    cid int64 `json:"cid";xml:"cid"`
    
    /* cod_postage_id optional此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃 */
    cod_postage_id int64 `json:"cod_postage_id";xml:"cod_postage_id"`
    
    /* cpv_memo optional针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷 */
    cpv_memo string `json:"cpv_memo";xml:"cpv_memo"`
    
    /* custom_made_type_id optional定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1 */
    custom_made_type_id string `json:"custom_made_type_id";xml:"custom_made_type_id"`
    
    /* delivery_time_delivery_time optional商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11 */
    delivery_time_delivery_time string `json:"delivery_time.delivery_time";xml:"delivery_time.delivery_time"`
    
    /* delivery_time_delivery_time_type optional发货时间类型：绝对发货时间或者相对发货时间 */
    delivery_time_delivery_time_type string `json:"delivery_time.delivery_time_type";xml:"delivery_time.delivery_time_type"`
    
    /* delivery_time_need_delivery_time optional设置是否使用发货时间，商品级别，sku级别 */
    delivery_time_need_delivery_time string `json:"delivery_time.need_delivery_time";xml:"delivery_time.need_delivery_time"`
    
    /* desc required宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制 */
    desc string `json:"desc";xml:"desc"`
    
    /* desc_modules optional商品描述模块化，模块列表，由List转化成jsonArray存入，后端逻辑验证通过，拼装成模块内容+锚点导航后存入desc中。数据结构具体参见Item_Desc_Module */
    desc_modules string `json:"desc_modules";xml:"desc_modules"`
    
    /* ems_fee optionalems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
    ems_fee Price `json:"ems_fee";xml:"ems_fee"`
    
    /* express_fee optional快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
    express_fee Price `json:"express_fee";xml:"express_fee"`
    
    /* features optional宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp */
    features string `json:"features";xml:"features"`
    
    /* food_security_contact optional厂家联系方式 */
    food_security_contact string `json:"food_security.contact";xml:"food_security.contact"`
    
    /* food_security_design_code optional产品标准号 */
    food_security_design_code string `json:"food_security.design_code";xml:"food_security.design_code"`
    
    /* food_security_factory optional厂名 */
    food_security_factory string `json:"food_security.factory";xml:"food_security.factory"`
    
    /* food_security_factory_site optional厂址 */
    food_security_factory_site string `json:"food_security.factory_site";xml:"food_security.factory_site"`
    
    /* food_security_food_additive optional食品添加剂 */
    food_security_food_additive string `json:"food_security.food_additive";xml:"food_security.food_additive"`
    
    /* food_security_health_product_no optional健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
    food_security_health_product_no string `json:"food_security.health_product_no";xml:"food_security.health_product_no"`
    
    /* food_security_mix optional配料表 */
    food_security_mix string `json:"food_security.mix";xml:"food_security.mix"`
    
    /* food_security_period optional保质期 */
    food_security_period string `json:"food_security.period";xml:"food_security.period"`
    
    /* food_security_plan_storage optional储藏方法 */
    food_security_plan_storage string `json:"food_security.plan_storage";xml:"food_security.plan_storage"`
    
    /* food_security_prd_license_no optional生产许可证号 */
    food_security_prd_license_no string `json:"food_security.prd_license_no";xml:"food_security.prd_license_no"`
    
    /* food_security_product_date_end optional生产结束日期,格式必须为yyyy-MM-dd */
    food_security_product_date_end string `json:"food_security.product_date_end";xml:"food_security.product_date_end"`
    
    /* food_security_product_date_start optional生产开始日期，格式必须为yyyy-MM-dd */
    food_security_product_date_start string `json:"food_security.product_date_start";xml:"food_security.product_date_start"`
    
    /* food_security_stock_date_end optional进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
    food_security_stock_date_end string `json:"food_security.stock_date_end";xml:"food_security.stock_date_end"`
    
    /* food_security_stock_date_start optional进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
    food_security_stock_date_start string `json:"food_security.stock_date_start";xml:"food_security.stock_date_start"`
    
    /* food_security_supplier optional供货商 */
    food_security_supplier string `json:"food_security.supplier";xml:"food_security.supplier"`
    
    /* freight_payer optional运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee） */
    freight_payer string `json:"freight_payer";xml:"freight_payer"`
    
    /* global_stock_country optional全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他） */
    global_stock_country string `json:"global_stock_country";xml:"global_stock_country"`
    
    /* global_stock_delivery_place optional全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台” */
    global_stock_delivery_place string `json:"global_stock_delivery_place";xml:"global_stock_delivery_place"`
    
    /* global_stock_tax_free_promise optional全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约 */
    global_stock_tax_free_promise bool `json:"global_stock_tax_free_promise";xml:"global_stock_tax_free_promise"`
    
    /* global_stock_type optional全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
    global_stock_type string `json:"global_stock_type";xml:"global_stock_type"`
    
    /* has_discount optional支持会员打折。可选值:true,false;默认值:false(不打折) */
    has_discount bool `json:"has_discount";xml:"has_discount"`
    
    /* has_invoice optional是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票) */
    has_invoice bool `json:"has_invoice";xml:"has_invoice"`
    
    /* has_showcase optional橱窗推荐。可选值:true,false;默认值:false(不推荐) */
    has_showcase bool `json:"has_showcase";xml:"has_showcase"`
    
    /* has_warranty optional是否有保修。可选值:true,false;默认值:false(不保修) */
    has_warranty bool `json:"has_warranty";xml:"has_warranty"`
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* image optional商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间） */
    image []byte `json:"image";xml:"image"`
    
    /* increment optional加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
    increment Price `json:"increment";xml:"increment"`
    
    /* input_custom_cpv optional针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上 */
    input_custom_cpv string `json:"input_custom_cpv";xml:"input_custom_cpv"`
    
    /* input_pids optional用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
    input_pids string `json:"input_pids";xml:"input_pids"`
    
    /* input_str optional用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
    input_str string `json:"input_str";xml:"input_str"`
    
    /* interactive_id optional主图视频互动信息id，必须填写主图视频id才能有互动信息id */
    interactive_id int64 `json:"interactive_id";xml:"interactive_id"`
    
    /* is_3D optional是否是3D */
    is_3D bool `json:"is_3D";xml:"is_3D"`
    
    /* is_ex optional是否在外店显示 */
    is_ex bool `json:"is_ex";xml:"is_ex"`
    
    /* is_lightning_consignment optional实物闪电发货 */
    is_lightning_consignment bool `json:"is_lightning_consignment";xml:"is_lightning_consignment"`
    
    /* is_offline optional是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。 */
    is_offline string `json:"is_offline";xml:"is_offline"`
    
    /* is_taobao optional是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
    is_taobao bool `json:"is_taobao";xml:"is_taobao"`
    
    /* is_xinpin optional商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。 */
    is_xinpin bool `json:"is_xinpin";xml:"is_xinpin"`
    
    /* item_size optional表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m） */
    item_size string `json:"item_size";xml:"item_size"`
    
    /* item_weight optional商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。 */
    item_weight string `json:"item_weight";xml:"item_weight"`
    
    /* lang optional商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体 */
    lang string `json:"lang";xml:"lang"`
    
    /* lease_extends_info optional租赁扩展信息 */
    lease_extends_info string `json:"lease_extends_info";xml:"lease_extends_info"`
    
    /* list_time optional定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss) */
    list_time time.Time `json:"list_time";xml:"list_time"`
    
    /* locality_life_choose_logis optional发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
    locality_life_choose_logis string `json:"locality_life.choose_logis";xml:"locality_life.choose_logis"`
    
    /* locality_life_eticket optional电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4 */
    locality_life_eticket string `json:"locality_life.eticket";xml:"locality_life.eticket"`
    
    /* locality_life_expirydate optional本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15 */
    locality_life_expirydate string `json:"locality_life.expirydate";xml:"locality_life.expirydate"`
    
    /* locality_life_merchant optional码商信息，格式为 码商id:nick */
    locality_life_merchant string `json:"locality_life.merchant";xml:"locality_life.merchant"`
    
    /* locality_life_network_id optional网点ID */
    locality_life_network_id string `json:"locality_life.network_id";xml:"locality_life.network_id"`
    
    /* locality_life_obs optional预约门店是否支持门店自提,1:是 */
    locality_life_obs string `json:"locality_life.obs";xml:"locality_life.obs"`
    
    /* locality_life_onsale_auto_refund_ratio optional电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
    locality_life_onsale_auto_refund_ratio int64 `json:"locality_life.onsale_auto_refund_ratio";xml:"locality_life.onsale_auto_refund_ratio"`
    
    /* locality_life_packageid optional新版电子凭证包 id */
    locality_life_packageid string `json:"locality_life.packageid";xml:"locality_life.packageid"`
    
    /* locality_life_refund_ratio optional退款比例，百分比%前的数字,1-100的正整数值 */
    locality_life_refund_ratio int64 `json:"locality_life.refund_ratio";xml:"locality_life.refund_ratio"`
    
    /* locality_life_refundmafee optional退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
    locality_life_refundmafee string `json:"locality_life.refundmafee";xml:"locality_life.refundmafee"`
    
    /* locality_life_verification optional核销打款 1代表核销打款 0代表非核销打款 */
    locality_life_verification string `json:"locality_life.verification";xml:"locality_life.verification"`
    
    /* locality_life_version optional新版电子凭证字段 */
    locality_life_version string `json:"locality_life.version";xml:"locality_life.version"`
    
    /* location_city required所在地城市。如杭州 。 */
    location_city string `json:"location.city";xml:"location.city"`
    
    /* location_state required所在地省份。如浙江 */
    location_state string `json:"location.state";xml:"location.state"`
    
    /* ms_payment_price optional订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm */
    ms_payment_price string `json:"ms_payment.price";xml:"ms_payment.price"`
    
    /* ms_payment_reference_price optional参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm */
    ms_payment_reference_price string `json:"ms_payment.reference_price";xml:"ms_payment.reference_price"`
    
    /* ms_payment_voucher_price optional尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm */
    ms_payment_voucher_price string `json:"ms_payment.voucher_price";xml:"ms_payment.voucher_price"`
    
    /* newprepay optional该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货； */
    newprepay string `json:"newprepay";xml:"newprepay"`
    
    /* num required商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。 */
    num int64 `json:"num";xml:"num"`
    
    /* o2o_bind_service optional汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。 */
    o2o_bind_service bool `json:"o2o_bind_service";xml:"o2o_bind_service"`
    
    /* outer_id optional商品外部编码，该字段的最大长度是64个字节 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* paimai_info_deposit optional拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。 */
    paimai_info_deposit int64 `json:"paimai_info.deposit";xml:"paimai_info.deposit"`
    
    /* paimai_info_interval optional降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。 */
    paimai_info_interval int64 `json:"paimai_info.interval";xml:"paimai_info.interval"`
    
    /* paimai_info_mode optional拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。 */
    paimai_info_mode int64 `json:"paimai_info.mode";xml:"paimai_info.mode"`
    
    /* paimai_info_reserve optional降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
    paimai_info_reserve Price `json:"paimai_info.reserve";xml:"paimai_info.reserve"`
    
    /* paimai_info_valid_hour optional自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
    paimai_info_valid_hour int64 `json:"paimai_info.valid_hour";xml:"paimai_info.valid_hour"`
    
    /* paimai_info_valid_minute optional自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
    paimai_info_valid_minute int64 `json:"paimai_info.valid_minute";xml:"paimai_info.valid_minute"`
    
    /* pic_path optional（推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* post_fee optional平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写 */
    post_fee Price `json:"post_fee";xml:"post_fee"`
    
    /* postage_id optional宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板） */
    postage_id int64 `json:"postage_id";xml:"postage_id"`
    
    /* price optional商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。 */
    price Price `json:"price";xml:"price"`
    
    /* product_id optional商品所属的产品ID(B商家发布商品需要用) */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* property_alias optional属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。 */
    property_alias string `json:"property_alias";xml:"property_alias"`
    
    /* props optional商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值 */
    props string `json:"props";xml:"props"`
    
    /* qualification optional商品资质信息 */
    qualification string `json:"qualification";xml:"qualification"`
    
    /* scenic_ticket_book_cost optional景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
    scenic_ticket_book_cost string `json:"scenic_ticket_book_cost";xml:"scenic_ticket_book_cost"`
    
    /* scenic_ticket_pay_way optional景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
    scenic_ticket_pay_way int64 `json:"scenic_ticket_pay_way";xml:"scenic_ticket_pay_way"`
    
    /* sell_point optional商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。 */
    sell_point string `json:"sell_point";xml:"sell_point"`
    
    /* sell_promise optional是否承诺退换货服务!虚拟商品无须设置此项! */
    sell_promise bool `json:"sell_promise";xml:"sell_promise"`
    
    /* seller_cids optional商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
    seller_cids string `json:"seller_cids";xml:"seller_cids"`
    
    /* shape optional宝贝形态:1代表电子券;0或其他代表实物 */
    shape string `json:"shape";xml:"shape"`
    
    /* sku_barcode optionalsku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔 */
    sku_barcode string `json:"sku_barcode";xml:"sku_barcode"`
    
    /* sku_delivery_times optional此参数暂时不起作用 */
    sku_delivery_times string `json:"sku_delivery_times";xml:"sku_delivery_times"`
    
    /* sku_hd_height optional家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。 可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50 */
    sku_hd_height string `json:"sku_hd_height";xml:"sku_hd_height"`
    
    /* sku_hd_lamp_quantity optional家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7 */
    sku_hd_lamp_quantity string `json:"sku_hd_lamp_quantity";xml:"sku_hd_lamp_quantity"`
    
    /* sku_hd_length optional家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：20,30,30 */
    sku_hd_length string `json:"sku_hd_length";xml:"sku_hd_length"`
    
    /* sku_outer_ids optionalSku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
    sku_outer_ids string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    /* sku_prices optionalSku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
    sku_prices string `json:"sku_prices";xml:"sku_prices"`
    
    /* sku_properties optional更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。 */
    sku_properties string `json:"sku_properties";xml:"sku_properties"`
    
    /* sku_quantities optionalSku的数量串，结构如：num1,num2,num3 如：2,3 */
    sku_quantities string `json:"sku_quantities";xml:"sku_quantities"`
    
    /* sku_spec_ids optional此参数暂时不起作用 */
    sku_spec_ids string `json:"sku_spec_ids";xml:"sku_spec_ids"`
    
    /* spu_confirm optional手机类目spu 优化，信息确认字段 */
    spu_confirm bool `json:"spu_confirm";xml:"spu_confirm"`
    
    /* stuff_status required新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性 */
    stuff_status string `json:"stuff_status";xml:"stuff_status"`
    
    /* sub_stock optional商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存 */
    sub_stock int64 `json:"sub_stock";xml:"sub_stock"`
    
    /* support_custom_made optional是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改 */
    support_custom_made bool `json:"support_custom_made";xml:"support_custom_made"`
    
    /* title required宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符； */
    title string `json:"title";xml:"title"`
    
    /* _type required发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。 */
    _type string `json:"type";xml:"type"`
    
    /* valid_thru optional有效期。可选值:7,14;单位:天;默认值:14 */
    valid_thru int64 `json:"valid_thru";xml:"valid_thru"`
    
    /* video_id optional主图视频id */
    video_id int64 `json:"video_id";xml:"video_id"`
    
    /* weight optional商品的重量(商超卖家专用字段) */
    weight int64 `json:"weight";xml:"weight"`
    
    /* wireless_desc optional无线的宝贝描述 */
    wireless_desc string `json:"wireless_desc";xml:"wireless_desc"`
    
}

func (req *TaobaoItemAddRequest) GetAPIName() string {
	return "taobao.item.add"
}

/* TaobaoItemAddResponse 此接口用于新增一个商品 
商品所属的卖家是当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。 */
type TaobaoItemAddResponse struct {
    
    /* item Object商品结构,仅有numIid和created返回 */
    item Item `json:"item";xml:"item"`
    
}

/* TaobaoItemPropimgDeleteRequest 删除propimg_id 所指定的商品属性图片 
传入的num_iid所对应的商品必须属于当前会话的用户 
propimg_id对应的属性图片需要属于num_iid对应的商品 */
type TaobaoItemPropimgDeleteRequest struct {
    
    /* id required商品属性图片ID */
    id int64 `json:"id";xml:"id"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemPropimgDeleteRequest) GetAPIName() string {
	return "taobao.item.propimg.delete"
}

/* TaobaoItemPropimgDeleteResponse 删除propimg_id 所指定的商品属性图片 
传入的num_iid所对应的商品必须属于当前会话的用户 
propimg_id对应的属性图片需要属于num_iid对应的商品 */
type TaobaoItemPropimgDeleteResponse struct {
    
    /* prop_img Object属性图片结构 */
    prop_img PropImg `json:"prop_img";xml:"prop_img"`
    
}

/* TaobaoItemImgDeleteRequest 删除商品图片 */
type TaobaoItemImgDeleteRequest struct {
    
    /* id required商品图片ID；如果是竖图，请将id的值设置为1 */
    id int64 `json:"id";xml:"id"`
    
    /* is_sixth_pic optional标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下 */
    is_sixth_pic bool `json:"is_sixth_pic";xml:"is_sixth_pic"`
    
    /* num_iid required商品数字ID */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemImgDeleteRequest) GetAPIName() string {
	return "taobao.item.img.delete"
}

/* TaobaoItemImgDeleteResponse 删除商品图片 */
type TaobaoItemImgDeleteResponse struct {
    
    /* item_img Object商品图片结构 */
    item_img ItemImg `json:"item_img";xml:"item_img"`
    
}

/* TaobaoItemSkuAddRequest 新增一个sku到num_iid指定的商品中 
传入的iid所对应的商品必须属于当前会话的用户 */
type TaobaoItemSkuAddRequest struct {
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* item_price optionalsku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功 */
    item_price Price `json:"item_price";xml:"item_price"`
    
    /* lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
    lang string `json:"lang";xml:"lang"`
    
    /* num_iid requiredSku所属商品数字id，可通过 taobao.item.get 获取。必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outer_id optionalSku的商家外部id */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* price requiredSku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
    price Price `json:"price";xml:"price"`
    
    /* properties requiredSku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity requiredSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuAddRequest) GetAPIName() string {
	return "taobao.item.sku.add"
}

/* TaobaoItemSkuAddResponse 新增一个sku到num_iid指定的商品中 
传入的iid所对应的商品必须属于当前会话的用户 */
type TaobaoItemSkuAddResponse struct {
    
    /* sku Objectsku */
    sku Sku `json:"sku";xml:"sku"`
    
}

/* TaobaoItemPropimgUploadRequest 添加一张商品属性图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 
商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。 */
type TaobaoItemPropimgUploadRequest struct {
    
    /* id optional属性图片ID。如果是新增不需要填写 */
    id int64 `json:"id";xml:"id"`
    
    /* image optional属性图片内容。类型:JPG,GIF;图片大小不超过:3M */
    image []byte `json:"image";xml:"image"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* position optional图片位置 */
    position int64 `json:"position";xml:"position"`
    
    /* properties required属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemPropimgUploadRequest) GetAPIName() string {
	return "taobao.item.propimg.upload"
}

/* TaobaoItemPropimgUploadResponse 添加一张商品属性图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 
商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。 */
type TaobaoItemPropimgUploadResponse struct {
    
    /* prop_img ObjectPropImg属性图片结构 */
    prop_img PropImg `json:"prop_img";xml:"prop_img"`
    
}

/* TaobaoItemSkuUpdateRequest *更新一个sku的数据 
*需要更新的sku通过属性properties进行匹配查找 
*商品的数量和价格必须大于等于0 
*sku记录会更新到指定的num_iid对应的商品中 
*num_iid对应的商品必须属于当前的会话用户 */
type TaobaoItemSkuUpdateRequest struct {
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* item_price optionalsku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功 */
    item_price Price `json:"item_price";xml:"item_price"`
    
    /* lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
    lang string `json:"lang";xml:"lang"`
    
    /* num_iid requiredSku所属商品数字id，可通过 taobao.item.get 获取 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outer_id optionalSku的商家外部id */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* price optionalSku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中） */
    price Price `json:"price";xml:"price"`
    
    /* properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。
如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号， */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity optionalSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuUpdateRequest) GetAPIName() string {
	return "taobao.item.sku.update"
}

/* TaobaoItemSkuUpdateResponse *更新一个sku的数据 
*需要更新的sku通过属性properties进行匹配查找 
*商品的数量和价格必须大于等于0 
*sku记录会更新到指定的num_iid对应的商品中 
*num_iid对应的商品必须属于当前的会话用户 */
type TaobaoItemSkuUpdateResponse struct {
    
    /* sku Object商品Sku */
    sku Sku `json:"sku";xml:"sku"`
    
}

/* TaobaoItemSkuGetRequest 获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家 */
type TaobaoItemSkuGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick optional卖家nick(num_iid和nick必传一个)，只传卖家nick时候，该api返回的结果不包含cspu（SKu上的产品规格信息）。 */
    nick string `json:"nick";xml:"nick"`
    
    /* num_iid optional商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* sku_id requiredSku的id。可以通过taobao.item.get得到 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoItemSkuGetRequest) GetAPIName() string {
	return "taobao.item.sku.get"
}

/* TaobaoItemSkuGetResponse 获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家 */
type TaobaoItemSkuGetResponse struct {
    
    /* sku ObjectSku */
    sku Sku `json:"sku";xml:"sku"`
    
}

/* TaobaoItemUpdateDelistingRequest * 单个商品下架
    * 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateDelistingRequest struct {
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemUpdateDelistingRequest) GetAPIName() string {
	return "taobao.item.update.delisting"
}

/* TaobaoItemUpdateDelistingResponse * 单个商品下架
    * 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateDelistingResponse struct {
    
    /* item Object返回商品更新信息：返回的结果是:num_iid和modified */
    item Item `json:"item";xml:"item"`
    
}

/* TaobaoFenxiaoProductToChannelImportRequest 支持供应商将已有产品导入到某个渠道销售 */
type TaobaoFenxiaoProductToChannelImportRequest struct {
    
    /* channel required要导入的渠道[21 零售PLUS]目前仅支持此渠道 */
    channel int64 `json:"channel";xml:"channel"`
    
    /* product_id required要导入的产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TaobaoFenxiaoProductToChannelImportRequest) GetAPIName() string {
	return "taobao.fenxiao.product.to.channel.import"
}

/* TaobaoFenxiaoProductToChannelImportResponse 支持供应商将已有产品导入到某个渠道销售 */
type TaobaoFenxiaoProductToChannelImportResponse struct {
    
}

/* TaobaoItemSkusGetRequest * 获取多个商品下的所有sku */
type TaobaoItemSkusGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iids requiredsku所属商品数字id，必选。num_iid个数不能超过40个 */
    num_iids string `json:"num_iids";xml:"num_iids"`
    
}

func (req *TaobaoItemSkusGetRequest) GetAPIName() string {
	return "taobao.item.skus.get"
}

/* TaobaoItemSkusGetResponse * 获取多个商品下的所有sku */
type TaobaoItemSkusGetResponse struct {
    
    /* skus Object ArraySku列表 */
    skus Sku `json:"skus";xml:"skus"`
    
}

/* TaobaoWlbItemBatchQueryRequest 根据用户id，item id list和store code来查询商品库存信息和批次信息 */
type TaobaoWlbItemBatchQueryRequest struct {
    
    /* item_ids required需要查询的商品ID列表，以字符串表示，ID间以;隔开 */
    item_ids string `json:"item_ids";xml:"item_ids"`
    
    /* page_no optional分页查询参数，指定查询页数，默认为1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* store_code optional仓库编号 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbItemBatchQueryRequest) GetAPIName() string {
	return "taobao.wlb.item.batch.query"
}

/* TaobaoWlbItemBatchQueryResponse 根据用户id，item id list和store code来查询商品库存信息和批次信息 */
type TaobaoWlbItemBatchQueryResponse struct {
    
    /* item_inventory_batch_list Object Array商品库存及批次信息查询结果 */
    item_inventory_batch_list WlbItemBatchInventory `json:"item_inventory_batch_list";xml:"item_inventory_batch_list"`
    
    /* total_count Basic返回结果记录的数量 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoStreetestSessionGetRequest 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测 */
type TaobaoStreetestSessionGetRequest struct {
    
}

func (req *TaobaoStreetestSessionGetRequest) GetAPIName() string {
	return "taobao.streetest.session.get"
}

/* TaobaoStreetestSessionGetResponse 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测 */
type TaobaoStreetestSessionGetResponse struct {
    
    /* stree_test_session_key Basic压测账号对应的sessionKey */
    stree_test_session_key string `json:"stree_test_session_key";xml:"stree_test_session_key"`
    
}

/* TaobaoItemUpdateListingRequest * 单个商品上架
* 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateListingRequest struct {
    
    /* num required需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num */
    num int64 `json:"num";xml:"num"`
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemUpdateListingRequest) GetAPIName() string {
	return "taobao.item.update.listing"
}

/* TaobaoItemUpdateListingResponse * 单个商品上架
* 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateListingResponse struct {
    
    /* item Object上架后返回的商品信息：返回的结果就是:num_iid和modified */
    item Item `json:"item";xml:"item"`
    
}

/* CainiaoBimTradeorderConsignRequest 驱动保税交易订单发货 */
type CainiaoBimTradeorderConsignRequest struct {
    
    /* res_id optional选择的线路ID非必填字段 */
    res_id string `json:"res_id";xml:"res_id"`
    
    /* store_code optional仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* trade_id required交易单号 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
}

func (req *CainiaoBimTradeorderConsignRequest) GetAPIName() string {
	return "cainiao.bim.tradeorder.consign"
}

/* CainiaoBimTradeorderConsignResponse 驱动保税交易订单发货 */
type CainiaoBimTradeorderConsignResponse struct {
    
    /* lg_order_code Basic菜鸟物流订单号,格式为LP开头 */
    lg_order_code string `json:"lg_order_code";xml:"lg_order_code"`
    
    /* store_order_code Basic菜鸟仓库作业单据号 */
    store_order_code string `json:"store_order_code";xml:"store_order_code"`
    
}

/* TaobaoRpRefundReviewRequest 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。 */
type TaobaoRpRefundReviewRequest struct {
    
    /* message required审核留言 */
    message string `json:"message";xml:"message"`
    
    /* operator required审核人姓名 */
    operator string `json:"operator";xml:"operator"`
    
    /* refund_id required退款单编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required退款阶段，可选值：售中：onsale，售后：aftersale */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version required退款最后更新时间，以时间戳的方式表示 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* result required审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核） */
    result bool `json:"result";xml:"result"`
    
}

func (req *TaobaoRpRefundReviewRequest) GetAPIName() string {
	return "taobao.rp.refund.review"
}

/* TaobaoRpRefundReviewResponse 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。 */
type TaobaoRpRefundReviewResponse struct {
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoWlbWaybillICancelRequest 面单号有误需要取消的时候，调用该接口取消获取的电子面单。 */
type TaobaoWlbWaybillICancelRequest struct {
    
    /* waybill_apply_cancel_request required取消接口入参 */
    waybill_apply_cancel_request WaybillApplyCancelRequest `json:"waybill_apply_cancel_request";xml:"waybill_apply_cancel_request"`
    
}

func (req *TaobaoWlbWaybillICancelRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.cancel"
}

/* TaobaoWlbWaybillICancelResponse 面单号有误需要取消的时候，调用该接口取消获取的电子面单。 */
type TaobaoWlbWaybillICancelResponse struct {
    
    /* cancel_result Basic调用取消是否成功 */
    cancel_result bool `json:"cancel_result";xml:"cancel_result"`
    
}

/* TaobaoWlbWaybillIQuerydetailRequest 查看面单号的当前状态，如签收、发货、失效等。 */
type TaobaoWlbWaybillIQuerydetailRequest struct {
    
    /* waybill_detail_query_request required面单查询请求 */
    waybill_detail_query_request WaybillDetailQueryRequest `json:"waybill_detail_query_request";xml:"waybill_detail_query_request"`
    
}

func (req *TaobaoWlbWaybillIQuerydetailRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.querydetail"
}

/* TaobaoWlbWaybillIQuerydetailResponse 查看面单号的当前状态，如签收、发货、失效等。 */
type TaobaoWlbWaybillIQuerydetailResponse struct {
    
    /* error_codes Basic Array面单查询错误编码 */
    error_codes string `json:"error_codes";xml:"error_codes"`
    
    /* inexistent_waybill_codes Basic Array不存在的面单号 */
    inexistent_waybill_codes string `json:"inexistent_waybill_codes";xml:"inexistent_waybill_codes"`
    
    /* query_success Basic查询是否成功 */
    query_success bool `json:"query_success";xml:"query_success"`
    
    /* waybill_details Object Array面单详情 */
    waybill_details WaybillDetailQueryInfo `json:"waybill_details";xml:"waybill_details"`
    
}

/* TaobaoWlbWaybillIPrintRequest 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。 */
type TaobaoWlbWaybillIPrintRequest struct {
    
    /* waybill_apply_print_check_request required打印请求 */
    waybill_apply_print_check_request WaybillApplyPrintCheckRequest `json:"waybill_apply_print_check_request";xml:"waybill_apply_print_check_request"`
    
}

func (req *TaobaoWlbWaybillIPrintRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.print"
}

/* TaobaoWlbWaybillIPrintResponse 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。 */
type TaobaoWlbWaybillIPrintResponse struct {
    
    /* waybill_apply_print_check_infos Object Array面单打印信息 */
    waybill_apply_print_check_infos WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_infos";xml:"waybill_apply_print_check_infos"`
    
}

/* TaobaoTmcGroupAddRequest 为已开通用户添加用户分组 */
type TaobaoTmcGroupAddRequest struct {
    
    /* group_name required分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* nicks required用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户 */
    nicks string `json:"nicks";xml:"nicks"`
    
    /* user_platform optional用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcGroupAddRequest) GetAPIName() string {
	return "taobao.tmc.group.add"
}

/* TaobaoTmcGroupAddResponse 为已开通用户添加用户分组 */
type TaobaoTmcGroupAddResponse struct {
    
    /* created Basic创建时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* group_name Basic分组名称 */
    group_name string `json:"group_name";xml:"group_name"`
    
}

/* TaobaoTradesSoldGetRequest 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
<br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
<br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/> <span style="color:red">注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）</span> */
type TaobaoTradesSoldGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_created optional查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss */
    end_created time.Time `json:"end_created";xml:"end_created"`
    
    /* ext_type optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
    ext_type string `json:"ext_type";xml:"ext_type"`
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0,处方药未审核 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数; 默认值:1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* rate_status optional评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评) */
    rate_status string `json:"rate_status";xml:"rate_status"`
    
    /* start_created optional查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss */
    start_created time.Time `json:"start_created";xml:"start_created"`
    
    /* status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 */
    status string `json:"status";xml:"status"`
    
    /* tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
    tag string `json:"tag";xml:"tag"`
    
    /* _type optional交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade,auto_delivery,ec,cod,step 这5 种的交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>可选值：fixed(一口价)auction(拍卖)guarantee_trade(一口价、拍卖)step(分阶段付款，万人团，阶梯团订单）independent_simple_trade(旺店入门版交易)independent_shop_trade(旺店标准版交易)auto_delivery(自动发货)ec(直冲)cod(货到付款)game_equipment(游戏装备)shopex_trade(ShopEX交易)netcn_trade(万网交易)external_trade(统一外部交易)instant_trade (即时到账)b2c_cod(大商家货到付款)hotel_trade(酒店类型交易)super_market_trade(商超交易)super_market_cod_trade(商超货到付款交易)taohua(淘花网交易类型）waimai(外卖交易类型）o2o_offlinetrade（O2O交易）nopaid（即时到帐/趣味猜交易类型）step (万人团) eticket(电子凭证) tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。pre_auth_type(预授权0元购) lazada（获取lazada订单类型） */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldGetRequest) GetAPIName() string {
	return "taobao.trades.sold.get"
}

/* TaobaoTradesSoldGetResponse 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
<br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
<br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/> <span style="color:red">注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）</span> */
type TaobaoTradesSoldGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的交易信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
    trades Trade `json:"trades";xml:"trades"`
    
}

/* TaobaoTmcGroupDeleteRequest 删除指定的分组或分组下的用户 */
type TaobaoTmcGroupDeleteRequest struct {
    
    /* group_name required分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* nicks optional用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组 */
    nicks string `json:"nicks";xml:"nicks"`
    
    /* user_platform optional用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcGroupDeleteRequest) GetAPIName() string {
	return "taobao.tmc.group.delete"
}

/* TaobaoTmcGroupDeleteResponse 删除指定的分组或分组下的用户 */
type TaobaoTmcGroupDeleteResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoTmcGroupsGetRequest 获取自定义用户分组列表 */
type TaobaoTmcGroupsGetRequest struct {
    
    /* group_names optional要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。 */
    group_names string `json:"group_names";xml:"group_names"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页返回多少个分组 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoTmcGroupsGetRequest) GetAPIName() string {
	return "taobao.tmc.groups.get"
}

/* TaobaoTmcGroupsGetResponse 获取自定义用户分组列表 */
type TaobaoTmcGroupsGetResponse struct {
    
    /* groups Object Arraydasdasd */
    groups TmcGroup `json:"groups";xml:"groups"`
    
    /* total_results Basic分组总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoRpReturngoodsRefuseRequest 卖家拒绝退货，目前仅支持天猫退货。 */
type TaobaoRpReturngoodsRefuseRequest struct {
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required退款服务状态，售后或者售中 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version required退款版本号 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* refuse_proof required拒绝退货凭证图片，必须图片格式，大小不能超过5M */
    refuse_proof []byte `json:"refuse_proof";xml:"refuse_proof"`
    
    /* refuse_reason_id optional拒绝原因编号，会提供拒绝原因列表供选择 */
    refuse_reason_id int64 `json:"refuse_reason_id";xml:"refuse_reason_id"`
    
}

func (req *TaobaoRpReturngoodsRefuseRequest) GetAPIName() string {
	return "taobao.rp.returngoods.refuse"
}

/* TaobaoRpReturngoodsRefuseResponse 卖家拒绝退货，目前仅支持天猫退货。 */
type TaobaoRpReturngoodsRefuseResponse struct {
    
    /* result Basicasdf */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoRpReturngoodsRefillRequest 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。 */
type TaobaoRpReturngoodsRefillRequest struct {
    
    /* logistics_company_code required物流公司编号 */
    logistics_company_code string `json:"logistics_company_code";xml:"logistics_company_code"`
    
    /* logistics_waybill_no required物流公司运单号 */
    logistics_waybill_no string `json:"logistics_waybill_no";xml:"logistics_waybill_no"`
    
    /* refund_id required退款单编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required退款阶段，可选值：售中：onsale，售后：aftersale */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRpReturngoodsRefillRequest) GetAPIName() string {
	return "taobao.rp.returngoods.refill"
}

/* TaobaoRpReturngoodsRefillResponse 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。 */
type TaobaoRpReturngoodsRefillResponse struct {
    
    /* is_success Basic验货操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoRefundsApplyGetRequest 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型 */
type TaobaoRefundsApplyGetRequest struct {
    
    /* fields required需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* seller_nick optional卖家昵称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
    /* status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。
WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) 
WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) 
WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) 
SELLER_REFUSE_BUYER(卖家拒绝退款) 
CLOSED(退款关闭) 
SUCCESS(退款成功) */
    status string `json:"status";xml:"status"`
    
    /* _type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。
fixed(一口价) 
auction(拍卖) 
guarantee_trade(一口价、拍卖) 
independent_simple_trade(旺店入门版交易) 
independent_shop_trade(旺店标准版交易) 
auto_delivery(自动发货) 
ec(直冲) 
cod(货到付款) 
fenxiao(分销) 
game_equipment(游戏装备) 
shopex_trade(ShopEX交易) 
netcn_trade(万网交易) 
external_trade(统一外部交易) */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoRefundsApplyGetRequest) GetAPIName() string {
	return "taobao.refunds.apply.get"
}

/* TaobaoRefundsApplyGetResponse 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型 */
type TaobaoRefundsApplyGetResponse struct {
    
    /* refunds Object Array搜索到的退款信息列表 */
    refunds Refund `json:"refunds";xml:"refunds"`
    
    /* total_results Basic搜索到的交易信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoTradeMemoUpdateRequest 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能 */
type TaobaoTradeMemoUpdateRequest struct {
    
    /* flag optional卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0 */
    flag int64 `json:"flag";xml:"flag"`
    
    /* memo optional卖家交易备注。最大长度: 1000个字节 */
    memo string `json:"memo";xml:"memo"`
    
    /* reset optional是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值 */
    reset bool `json:"reset";xml:"reset"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeMemoUpdateRequest) GetAPIName() string {
	return "taobao.trade.memo.update"
}

/* TaobaoTradeMemoUpdateResponse 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能 */
type TaobaoTradeMemoUpdateResponse struct {
    
    /* trade Object更新交易的备注信息后返回的Trade，其中可用字段为tid和modified */
    trade Trade `json:"trade";xml:"trade"`
    
}

/* TaobaoTradeMemoAddRequest 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update */
type TaobaoTradeMemoAddRequest struct {
    
    /* flag optional交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0 */
    flag int64 `json:"flag";xml:"flag"`
    
    /* memo required交易备注。最大长度: 1000个字节 */
    memo string `json:"memo";xml:"memo"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeMemoAddRequest) GetAPIName() string {
	return "taobao.trade.memo.add"
}

/* TaobaoTradeMemoAddResponse 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update */
type TaobaoTradeMemoAddResponse struct {
    
    /* trade Object对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created */
    trade Trade `json:"trade";xml:"trade"`
    
}

/* TaobaoTraderatesGetRequest 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。 */
type TaobaoTraderatesGetRequest struct {
    
    /* end_date optional评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* fields required需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iid optional商品的数字ID */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* page_no optional页码。取值范围:大于零的整数最大限制为200; 默认值:1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页获取条数。默认值40，最小值1，最大值150。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* peer_nick optional评价对方昵称 */
    peer_nick string `json:"peer_nick";xml:"peer_nick"`
    
    /* rate_type required评价类型。可选值:get(得到),give(给出) */
    rate_type string `json:"rate_type";xml:"rate_type"`
    
    /* result optional评价结果。可选值:good(好评),neutral(中评),bad(差评) */
    result string `json:"result";xml:"result"`
    
    /* role required评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。 */
    role string `json:"role";xml:"role"`
    
    /* start_date optional评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。 */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* tid optional交易订单id，可以是父订单id号，也可以是子订单id号 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTraderatesGetRequest) GetAPIName() string {
	return "taobao.traderates.get"
}

/* TaobaoTraderatesGetResponse 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。 */
type TaobaoTraderatesGetResponse struct {
    
    /* has_next Basic当使用use_has_next时返回信息，如果还有下一页则返回true */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trade_rates Object Array评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息 */
    trade_rates TradeRate `json:"trade_rates";xml:"trade_rates"`
    
}

/* TaobaoTradeFullinfoGetRequest 获取单笔交易的详细信息
<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值 
<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>4. 获取红包优惠金额可以使用字段 coupon_fee
<br/>5. 请按需获取字段，减少TOP系统的压力 */
type TaobaoTradeFullinfoGetRequest struct {
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。 */
    fields string `json:"fields";xml:"fields"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeFullinfoGetRequest) GetAPIName() string {
	return "taobao.trade.fullinfo.get"
}

/* TaobaoTradeFullinfoGetResponse 获取单笔交易的详细信息
<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值 
<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>4. 获取红包优惠金额可以使用字段 coupon_fee
<br/>5. 请按需获取字段，减少TOP系统的压力 */
type TaobaoTradeFullinfoGetResponse struct {
    
    /* trade Object交易主订单信息 */
    trade Trade `json:"trade";xml:"trade"`
    
}

/* AlibabaProviderEinvoiceQueryRequest 开票服务商提供的对已开电子发票的查询功能，此接口是可选实现。只有部分电子发票支持通过此接口查询 */
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

/* AlibabaProviderEinvoiceQueryResponse 开票服务商提供的对已开电子发票的查询功能，此接口是可选实现。只有部分电子发票支持通过此接口查询 */
type AlibabaProviderEinvoiceQueryResponse struct {
    
    /* einvoice_list Object Array电子发票列表 */
    einvoice_list ProviderEinvoice `json:"einvoice_list";xml:"einvoice_list"`
    
}

/* TaobaoRefundGetRequest 获取单笔退款详情 */
type TaobaoRefundGetRequest struct {
    
    /* fields required需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku */
    fields string `json:"fields";xml:"fields"`
    
    /* refund_id required退款单号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
}

func (req *TaobaoRefundGetRequest) GetAPIName() string {
	return "taobao.refund.get"
}

/* TaobaoRefundGetResponse 获取单笔退款详情 */
type TaobaoRefundGetResponse struct {
    
    /* refund Object退款详情 */
    refund Refund `json:"refund";xml:"refund"`
    
}

/* TaobaoRefundsReceiveGetRequest 查询卖家收到的退款列表 */
type TaobaoRefundsReceiveGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_modified optional查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* fields required需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数; 默认值:1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_modified optional查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
    /* status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功) */
    status string `json:"status";xml:"status"`
    
    /* _type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a> */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoRefundsReceiveGetRequest) GetAPIName() string {
	return "taobao.refunds.receive.get"
}

/* TaobaoRefundsReceiveGetResponse 查询卖家收到的退款列表 */
type TaobaoRefundsReceiveGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* refunds Object Array搜索到的退款信息列表 */
    refunds Refund `json:"refunds";xml:"refunds"`
    
    /* total_results Basic搜索到的退款信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoAreasGetRequest 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a> */
type TaobaoAreasGetRequest struct {
    
    /* fields required需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip. */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoAreasGetRequest) GetAPIName() string {
	return "taobao.areas.get"
}

/* TaobaoAreasGetResponse 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a> */
type TaobaoAreasGetResponse struct {
    
    /* areas Object Array地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。 */
    areas Area `json:"areas";xml:"areas"`
    
}

/* TaobaoFenxiaoProductSkusGetRequest 产品sku查询 */
type TaobaoFenxiaoProductSkusGetRequest struct {
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TaobaoFenxiaoProductSkusGetRequest) GetAPIName() string {
	return "taobao.fenxiao.product.skus.get"
}

/* TaobaoFenxiaoProductSkusGetResponse 产品sku查询 */
type TaobaoFenxiaoProductSkusGetResponse struct {
    
    /* skus Object Arraysku信息 */
    skus FenxiaoSku `json:"skus";xml:"skus"`
    
    /* total_results Basic记录数量 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* CainiaoBmsOrderCreateRequest 通过接口，在菜鸟商家工作台创建订单，并通过商家工作台进行发货管理。 */
type CainiaoBmsOrderCreateRequest struct {
    
    /* buyer_message optional买家留言 */
    buyer_message string `json:"buyer_message";xml:"buyer_message"`
    
    /* buyer_nick optional买家名称" */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* cod_fee optionalCOD服务费 */
    cod_fee int64 `json:"cod_fee";xml:"cod_fee"`
    
    /* created required创建时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* discount_fee optional优惠金额，单位分 */
    discount_fee int64 `json:"discount_fee";xml:"discount_fee"`
    
    /* invoice_amount optional发票金额，单位分 */
    invoice_amount int64 `json:"invoice_amount";xml:"invoice_amount"`
    
    /* invoice_title optional发票抬头 */
    invoice_title string `json:"invoice_title";xml:"invoice_title"`
    
    /* invoice_type optional发票类型(1 电子、2 纸质) */
    invoice_type int64 `json:"invoice_type";xml:"invoice_type"`
    
    /* is_cod optional是否COD */
    is_cod bool `json:"is_cod";xml:"is_cod"`
    
    /* is_invoice optional是否打印发票 */
    is_invoice bool `json:"is_invoice";xml:"is_invoice"`
    
    /* items optionaldemo */
    items Items `json:"items";xml:"items"`
    
    /* order_amount required订单总金额，单位分 */
    order_amount int64 `json:"order_amount";xml:"order_amount"`
    
    /* order_type optional订单类型，默认0，普通销售订单，3:B2B单 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* paied_amount required支付金额，单位分 */
    paied_amount int64 `json:"paied_amount";xml:"paied_amount"`
    
    /* pay_time required支付时间 */
    pay_time time.Time `json:"pay_time";xml:"pay_time"`
    
    /* post_fee optional快递费，单位分 */
    post_fee int64 `json:"post_fee";xml:"post_fee"`
    
    /* receiver_address required收货地址 */
    receiver_address string `json:"receiver_address";xml:"receiver_address"`
    
    /* receiver_city required收货市 */
    receiver_city string `json:"receiver_city";xml:"receiver_city"`
    
    /* receiver_country optional收货人国籍 */
    receiver_country string `json:"receiver_country";xml:"receiver_country"`
    
    /* receiver_district optional收货区 */
    receiver_district string `json:"receiver_district";xml:"receiver_district"`
    
    /* receiver_mobile optional收货人手机，手机与电话不可同时为空 */
    receiver_mobile string `json:"receiver_mobile";xml:"receiver_mobile"`
    
    /* receiver_name required收货人名称 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* receiver_phone optional收货人电话，手机与电话不可同时为空 */
    receiver_phone string `json:"receiver_phone";xml:"receiver_phone"`
    
    /* receiver_state required收货省 */
    receiver_state string `json:"receiver_state";xml:"receiver_state"`
    
    /* receiver_town optional街道 */
    receiver_town string `json:"receiver_town";xml:"receiver_town"`
    
    /* receiver_zip required收货邮编 */
    receiver_zip string `json:"receiver_zip";xml:"receiver_zip"`
    
    /* seller_memo optional卖家备注 */
    seller_memo string `json:"seller_memo";xml:"seller_memo"`
    
    /* shop_code required店铺信息 */
    shop_code string `json:"shop_code";xml:"shop_code"`
    
    /* trade_id required交易单号 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
    /* wait_pay_amount required应收金额，单位分 */
    wait_pay_amount int64 `json:"wait_pay_amount";xml:"wait_pay_amount"`
    
}

func (req *CainiaoBmsOrderCreateRequest) GetAPIName() string {
	return "cainiao.bms.order.create"
}

/* CainiaoBmsOrderCreateResponse 通过接口，在菜鸟商家工作台创建订单，并通过商家工作台进行发货管理。 */
type CainiaoBmsOrderCreateResponse struct {
    
}

/* TaobaoFenxiaoProductSkuUpdateRequest 产品SKU信息更新 */
type TaobaoFenxiaoProductSkuUpdateRequest struct {
    
    /* agent_cost_price optional代销采购价 */
    agent_cost_price string `json:"agent_cost_price";xml:"agent_cost_price"`
    
    /* dealer_cost_price optional经销采购价 */
    dealer_cost_price string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties requiredsku属性 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity optional产品SKU库存 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* sku_number optional商家编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
    /* standard_price optional采购基准价 */
    standard_price string `json:"standard_price";xml:"standard_price"`
    
}

func (req *TaobaoFenxiaoProductSkuUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.update"
}

/* TaobaoFenxiaoProductSkuUpdateResponse 产品SKU信息更新 */
type TaobaoFenxiaoProductSkuUpdateResponse struct {
    
    /* created Basic操作时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoProductSkuAddRequest 添加产品SKU信息 */
type TaobaoFenxiaoProductSkuAddRequest struct {
    
    /* agent_cost_price special代销采购价 */
    agent_cost_price string `json:"agent_cost_price";xml:"agent_cost_price"`
    
    /* dealer_cost_price special经销采购价 */
    dealer_cost_price string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties requiredsku属性 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity optionalsku产品库存，在0到1000000之间，如果不传，则库存为0 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* sku_number optional商家编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
    /* standard_price required采购基准价，最大值1000000000 */
    standard_price string `json:"standard_price";xml:"standard_price"`
    
}

func (req *TaobaoFenxiaoProductSkuAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.add"
}

/* TaobaoFenxiaoProductSkuAddResponse 添加产品SKU信息 */
type TaobaoFenxiaoProductSkuAddResponse struct {
    
    /* created Basic操作时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoProductSkuDeleteRequest 根据sku properties删除sku数据 */
type TaobaoFenxiaoProductSkuDeleteRequest struct {
    
    /* product_id required产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties requiredsku属性 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductSkuDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.delete"
}

/* TaobaoFenxiaoProductSkuDeleteResponse 根据sku properties删除sku数据 */
type TaobaoFenxiaoProductSkuDeleteResponse struct {
    
    /* created Basic操作时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}

/* QimenTaobaoAutoReturnorderConfirmRequest 菜鸟自动创建销退单的入库确认回传 */
type QimenTaobaoAutoReturnorderConfirmRequest struct {
    
    /* customerId requiredcustomerId */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional */
    request AutoReturnInOrderConfirmRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoAutoReturnorderConfirmRequest) GetAPIName() string {
	return "qimen.taobao.auto.returnorder.confirm"
}

/* QimenTaobaoAutoReturnorderConfirmResponse 菜鸟自动创建销退单的入库确认回传 */
type QimenTaobaoAutoReturnorderConfirmResponse struct {
    
    /* response Objectresponse */
    response Response `json:"response";xml:"response"`
    
}

/* TaobaoFenxiaoProductImageDeleteRequest 产品图片删除，只删除图片信息，不真正删除图片 */
type TaobaoFenxiaoProductImageDeleteRequest struct {
    
    /* position required图片位置 */
    position int64 `json:"position";xml:"position"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties optionalproperties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductImageDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.product.image.delete"
}

/* TaobaoFenxiaoProductImageDeleteResponse 产品图片删除，只删除图片信息，不真正删除图片 */
type TaobaoFenxiaoProductImageDeleteResponse struct {
    
    /* created Basic操作时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoWlbWaybillIProductRequest 商家可以查询物流商的产品类型和服务能力。 */
type TaobaoWlbWaybillIProductRequest struct {
    
    /* waybill_product_type_request required查询物流商电子面单产品类型入参 */
    waybill_product_type_request WaybillProductTypeRequest `json:"waybill_product_type_request";xml:"waybill_product_type_request"`
    
}

func (req *TaobaoWlbWaybillIProductRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.product"
}

/* TaobaoWlbWaybillIProductResponse 商家可以查询物流商的产品类型和服务能力。 */
type TaobaoWlbWaybillIProductResponse struct {
    
    /* product_types Object Array产品类型返回 */
    product_types WaybillProductType `json:"product_types";xml:"product_types"`
    
}

/* TaobaoFenxiaoProductImageUploadRequest 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片 */
type TaobaoFenxiaoProductImageUploadRequest struct {
    
    /* image special产品图片 */
    image []byte `json:"image";xml:"image"`
    
    /* pic_path special产品主图图片空间相对路径或绝对路径 */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* position required图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图 */
    position int64 `json:"position";xml:"position"`
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties optionalproperties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductImageUploadRequest) GetAPIName() string {
	return "taobao.fenxiao.product.image.upload"
}

/* TaobaoFenxiaoProductImageUploadResponse 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片 */
type TaobaoFenxiaoProductImageUploadResponse struct {
    
    /* created Basic操作时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* result Basic操作是否成功 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoShopGetRequest 获取卖家店铺的基本信息 */
type TaobaoShopGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick required卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoShopGetRequest) GetAPIName() string {
	return "taobao.shop.get"
}

/* TaobaoShopGetResponse 获取卖家店铺的基本信息 */
type TaobaoShopGetResponse struct {
    
    /* shop Object店铺信息 */
    shop Shop `json:"shop";xml:"shop"`
    
}

/* TaobaoJdsTradesStatisticsDiffRequest 订单全链路状态统计差异比较 */
type TaobaoJdsTradesStatisticsDiffRequest struct {
    
    /* date required查询的日期，格式如YYYYMMDD的日期对应的数字 */
    date int64 `json:"date";xml:"date"`
    
    /* page_no optional页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* post_status required需要比较的状态 */
    post_status string `json:"post_status";xml:"post_status"`
    
    /* pre_status required需要比较的状态，将会和post_status做比较 */
    pre_status string `json:"pre_status";xml:"pre_status"`
    
}

func (req *TaobaoJdsTradesStatisticsDiffRequest) GetAPIName() string {
	return "taobao.jds.trades.statistics.diff"
}

/* TaobaoJdsTradesStatisticsDiffResponse 订单全链路状态统计差异比较 */
type TaobaoJdsTradesStatisticsDiffResponse struct {
    
    /* tids Basic Arraypre_status比post_status多的tid列表 */
    tids int64 `json:"tids";xml:"tids"`
    
    /* total_results Basic总记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoShopcatsListGetRequest 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异） */
type TaobaoShopcatsListGetRequest struct {
    
    /* fields optional需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoShopcatsListGetRequest) GetAPIName() string {
	return "taobao.shopcats.list.get"
}

/* TaobaoShopcatsListGetResponse 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异） */
type TaobaoShopcatsListGetResponse struct {
    
    /* shop_cats Object Array店铺类目列表信息 */
    shop_cats ShopCat `json:"shop_cats";xml:"shop_cats"`
    
}

/* TaobaoJdsTradesStatisticsGetRequest 获取订单数量统计结果 */
type TaobaoJdsTradesStatisticsGetRequest struct {
    
    /* date required查询的日期，格式如YYYYMMDD的日期对应的数字 */
    date int64 `json:"date";xml:"date"`
    
}

func (req *TaobaoJdsTradesStatisticsGetRequest) GetAPIName() string {
	return "taobao.jds.trades.statistics.get"
}

/* TaobaoJdsTradesStatisticsGetResponse 获取订单数量统计结果 */
type TaobaoJdsTradesStatisticsGetResponse struct {
    
    /* status_infos Object Array订单状态计数值 */
    status_infos TradeStat `json:"status_infos";xml:"status_infos"`
    
}

/* TaobaoSellercatsListGetRequest 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家） */
type TaobaoSellercatsListGetRequest struct {
    
    /* fields optionalfields参数 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick required卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercatsListGetRequest) GetAPIName() string {
	return "taobao.sellercats.list.get"
}

/* TaobaoSellercatsListGetResponse 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家） */
type TaobaoSellercatsListGetResponse struct {
    
    /* seller_cats Object Array卖家自定义类目 */
    seller_cats SellerCat `json:"seller_cats";xml:"seller_cats"`
    
}

/* TmallItemHscodeDetailGetRequest 通过hscode获取计量单位和销售单位 */
type TmallItemHscodeDetailGetRequest struct {
    
    /* hscode requiredhscode */
    hscode string `json:"hscode";xml:"hscode"`
    
}

func (req *TmallItemHscodeDetailGetRequest) GetAPIName() string {
	return "tmall.item.hscode.detail.get"
}

/* TmallItemHscodeDetailGetResponse 通过hscode获取计量单位和销售单位 */
type TmallItemHscodeDetailGetResponse struct {
    
    /* results Basic Array返回的计量单位和销售单位 */
    results map[string]interface{} `json:"results";xml:"results"`
    
}

/* TaobaoJdsRefundTracesGetRequest 获取聚石塔数据共享的交易全链路的退款信息 */
type TaobaoJdsRefundTracesGetRequest struct {
    
    /* refund_id required淘宝的退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* return_user_status optional是否返回用户状态(是否存在) */
    return_user_status bool `json:"return_user_status";xml:"return_user_status"`
    
}

func (req *TaobaoJdsRefundTracesGetRequest) GetAPIName() string {
	return "taobao.jds.refund.traces.get"
}

/* TaobaoJdsRefundTracesGetResponse 获取聚石塔数据共享的交易全链路的退款信息 */
type TaobaoJdsRefundTracesGetResponse struct {
    
    /* traces Object Array退款跟踪列表 */
    traces RefundTrace `json:"traces";xml:"traces"`
    
    /* user_status Basic用户在全链路系统中的状态(比如是否存在) */
    user_status string `json:"user_status";xml:"user_status"`
    
}

/* TaobaoJdsHluserGetRequest 订单全链路用户信息获取 */
type TaobaoJdsHluserGetRequest struct {
    
}

func (req *TaobaoJdsHluserGetRequest) GetAPIName() string {
	return "taobao.jds.hluser.get"
}

/* TaobaoJdsHluserGetResponse 订单全链路用户信息获取 */
type TaobaoJdsHluserGetResponse struct {
    
    /* hl_user Object回流用户信息 */
    hl_user HlUserDO `json:"hl_user";xml:"hl_user"`
    
}

/* TaobaoItemSkuPriceUpdateRequest 更新商品SKU的价格 */
type TaobaoItemSkuPriceUpdateRequest struct {
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* item_price optionalsku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功 */
    item_price Price `json:"item_price";xml:"item_price"`
    
    /* lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
    lang string `json:"lang";xml:"lang"`
    
    /* num_iid requiredSku所属商品数字id，可通过 taobao.item.get 获取 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outer_id optionalSku的商家外部id */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* price optionalSku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中） */
    price Price `json:"price";xml:"price"`
    
    /* properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity optionalSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuPriceUpdateRequest) GetAPIName() string {
	return "taobao.item.sku.price.update"
}

/* TaobaoItemSkuPriceUpdateResponse 更新商品SKU的价格 */
type TaobaoItemSkuPriceUpdateResponse struct {
    
    /* sku Object商品SKU信息（只包含num_iid和modified） */
    sku Sku `json:"sku";xml:"sku"`
    
}

/* TaobaoJdsHluserUpdateRequest 订单全链路用户信息修改，比如是否开放买家端展示 */
type TaobaoJdsHluserUpdateRequest struct {
    
    /* open_for_buyer optional回流信息是否开通买家端展示,可选值open,close */
    open_for_buyer string `json:"open_for_buyer";xml:"open_for_buyer"`
    
    /* open_nodes optional为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS */
    open_nodes string `json:"open_nodes";xml:"open_nodes"`
    
}

func (req *TaobaoJdsHluserUpdateRequest) GetAPIName() string {
	return "taobao.jds.hluser.update"
}

/* TaobaoJdsHluserUpdateResponse 订单全链路用户信息修改，比如是否开放买家端展示 */
type TaobaoJdsHluserUpdateResponse struct {
    
    /* result Basic是否成功 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoItemBarcodeUpdateRequest 通过该接口，将商品以及SKU上得条形码信息补全 */
type TaobaoItemBarcodeUpdateRequest struct {
    
    /* isforce optional是否强制保存商品条码。
true：强制保存
false ：需要执行条码库校验 */
    isforce bool `json:"isforce";xml:"isforce"`
    
    /* item_barcode optional商品条形码，如果不用更新，可选择不填 */
    item_barcode string `json:"item_barcode";xml:"item_barcode"`
    
    /* item_id required被更新商品的ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sku_barcodes optionalSKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔 */
    sku_barcodes string `json:"sku_barcodes";xml:"sku_barcodes"`
    
    /* sku_ids optional被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置 */
    sku_ids string `json:"sku_ids";xml:"sku_ids"`
    
    /* src optional访问来源，这字段提供给千牛扫码枪用，
其他调用方，不需要填写 */
    src string `json:"src";xml:"src"`
    
}

func (req *TaobaoItemBarcodeUpdateRequest) GetAPIName() string {
	return "taobao.item.barcode.update"
}

/* TaobaoItemBarcodeUpdateResponse 通过该接口，将商品以及SKU上得条形码信息补全 */
type TaobaoItemBarcodeUpdateResponse struct {
    
    /* item Object商品结构里的num_iid，modified */
    item Item `json:"item";xml:"item"`
    
}

/* TaobaoWlbWaybillIGetRequest 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。 */
type TaobaoWlbWaybillIGetRequest struct {
    
    /* waybill_apply_new_request required面单申请 */
    waybill_apply_new_request WaybillApplyNewRequest `json:"waybill_apply_new_request";xml:"waybill_apply_new_request"`
    
}

func (req *TaobaoWlbWaybillIGetRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.get"
}

/* TaobaoWlbWaybillIGetResponse 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。 */
type TaobaoWlbWaybillIGetResponse struct {
    
    /* waybill_apply_new_cols Object Array面单申请接口返回信息 */
    waybill_apply_new_cols WaybillApplyNewInfo `json:"waybill_apply_new_cols";xml:"waybill_apply_new_cols"`
    
}

/* TaobaoWlbWaybillISearchRequest 获取发货地&CP开通状态&账户的使用情况 */
type TaobaoWlbWaybillISearchRequest struct {
    
    /* waybill_apply_request required查询网点信息 */
    waybill_apply_request WaybillApplyRequest `json:"waybill_apply_request";xml:"waybill_apply_request"`
    
}

func (req *TaobaoWlbWaybillISearchRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.search"
}

/* TaobaoWlbWaybillISearchResponse 获取发货地&CP开通状态&账户的使用情况 */
type TaobaoWlbWaybillISearchResponse struct {
    
    /* subscribtions Object Array订购关系 */
    subscribtions WaybillApplySubscriptionInfo `json:"subscribtions";xml:"subscribtions"`
    
}

/* TaobaoRegionWarehouseManageRequest 编辑仓库覆盖范围 */
type TaobaoRegionWarehouseManageRequest struct {
    
    /* regions required可映射三级地址,例: 广东省 */
    regions string `json:"regions";xml:"regions"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoRegionWarehouseManageRequest) GetAPIName() string {
	return "taobao.region.warehouse.manage"
}

/* TaobaoRegionWarehouseManageResponse 编辑仓库覆盖范围 */
type TaobaoRegionWarehouseManageResponse struct {
    
    /* result Object返回结果 */
    result BaseResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbWaybillIFullupdateRequest 商家更新电子面单号对应的订单信息。
<br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；
<br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。 */
type TaobaoWlbWaybillIFullupdateRequest struct {
    
    /* waybill_apply_full_update_request required更新面单信息请求 */
    waybill_apply_full_update_request WaybillApplyFullUpdateRequest `json:"waybill_apply_full_update_request";xml:"waybill_apply_full_update_request"`
    
}

func (req *TaobaoWlbWaybillIFullupdateRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.fullupdate"
}

/* TaobaoWlbWaybillIFullupdateResponse 商家更新电子面单号对应的订单信息。
<br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；
<br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。 */
type TaobaoWlbWaybillIFullupdateResponse struct {
    
    /* waybill_apply_update_info Object更新接口出参 */
    waybill_apply_update_info WaybillApplyUpdateInfo `json:"waybill_apply_update_info";xml:"waybill_apply_update_info"`
    
}

/* TaobaoRegionSaleQueryRequest 查询商品销售区域 */
type TaobaoRegionSaleQueryRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sale_region_level required1:国家;2:省;3: 市;4:区县 */
    sale_region_level int64 `json:"sale_region_level";xml:"sale_region_level"`
    
    /* sku_id required无sku传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionSaleQueryRequest) GetAPIName() string {
	return "taobao.region.sale.query"
}

/* TaobaoRegionSaleQueryResponse 查询商品销售区域 */
type TaobaoRegionSaleQueryResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}

/* TaobaoPictureIsreferencedGetRequest 查询图片是否被引用，被引用返回true，未被引用返回false */
type TaobaoPictureIsreferencedGetRequest struct {
    
    /* picture_id required图片id */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureIsreferencedGetRequest) GetAPIName() string {
	return "taobao.picture.isreferenced.get"
}

/* TaobaoPictureIsreferencedGetResponse 查询图片是否被引用，被引用返回true，未被引用返回false */
type TaobaoPictureIsreferencedGetResponse struct {
    
    /* is_referenced Basic图片是否被引用 */
    is_referenced bool `json:"is_referenced";xml:"is_referenced"`
    
}

/* TaobaoRegionWarehouseQueryRequest 查询仓库覆盖范围 */
type TaobaoRegionWarehouseQueryRequest struct {
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoRegionWarehouseQueryRequest) GetAPIName() string {
	return "taobao.region.warehouse.query"
}

/* TaobaoRegionWarehouseQueryResponse 查询仓库覆盖范围 */
type TaobaoRegionWarehouseQueryResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}

/* TaobaoDeliveryTemplateDeleteRequest 根据用户指定的模板ID删除指定的模板 */
type TaobaoDeliveryTemplateDeleteRequest struct {
    
    /* template_id required运费模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TaobaoDeliveryTemplateDeleteRequest) GetAPIName() string {
	return "taobao.delivery.template.delete"
}

/* TaobaoDeliveryTemplateDeleteResponse 根据用户指定的模板ID删除指定的模板 */
type TaobaoDeliveryTemplateDeleteResponse struct {
    
    /* complete Basic表示删除成功还是失败 */
    complete bool `json:"complete";xml:"complete"`
    
}

/* TaobaoDeliveryTemplatesGetRequest 根据用户ID获取用户下所有模板 */
type TaobaoDeliveryTemplatesGetRequest struct {
    
    /* fields required需返回的字段列表。 <br/> 
<B>
可选值:示例中的值;各字段之间用","隔开
</B>
<br/><br/> 
<font color=blue>
template_id：查询模板ID <br/> 
template_name:查询模板名称 <br/> 
assumer：查询服务承担放 <br/> 
valuation:查询计价规则 <br/> 
supports:查询增值服务列表 <br/> 
created:查询模板创建时间 <br/> 
modified:查询修改时间<br/> 
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> 
address:卖家地址信息
</font>
<br/><br/> 
<font color=bule>
query_cod:查询货到付款运费信息<br/> 
query_post:查询平邮运费信息 <br/> 
query_express:查询快递公司运费信息 <br/> 
query_ems:查询EMS运费信息<br/> 
query_bzsd:查询普通保障速递运费信息<br/> 
query_wlb:查询物流宝保障速递运费信息<br/> 
query_furniture:查询家装物流运费信息
<font color=blue>
<br/><br/>
<font color=red>不能有空格</font> */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoDeliveryTemplatesGetRequest) GetAPIName() string {
	return "taobao.delivery.templates.get"
}

/* TaobaoDeliveryTemplatesGetResponse 根据用户ID获取用户下所有模板 */
type TaobaoDeliveryTemplatesGetResponse struct {
    
    /* delivery_templates Object Array运费模板列表 */
    delivery_templates DeliveryTemplate `json:"delivery_templates";xml:"delivery_templates"`
    
    /* total_results Basic获得到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoDeliveryTemplateUpdateRequest 修改运费模板 */
type TaobaoDeliveryTemplateUpdateRequest struct {
    
    /* assumer optional可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费 */
    assumer int64 `json:"assumer";xml:"assumer"`
    
    /* name optional模板名称，长度不能大于50个字节 */
    name string `json:"name";xml:"name"`
    
    /* template_add_fees required增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_add_fees string `json:"template_add_fees";xml:"template_add_fees"`
    
    /* template_add_standards required增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_add_standards string `json:"template_add_standards";xml:"template_add_standards"`
    
    /* template_dests required邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font> */
    template_dests string `json:"template_dests";xml:"template_dests"`
    
    /* template_id required需要修改的模板对应的模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
    /* template_start_fees required首费：输入0.01-999.99（最多包含两位小数）
<br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_start_fees string `json:"template_start_fees";xml:"template_start_fees"`
    
    /* template_start_standards required首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_start_standards string `json:"template_start_standards";xml:"template_start_standards"`
    
    /* template_types required运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4 
如: post;express;ems;cod 
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. 
 </font>
<br/>
<font color=blue>
普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。
货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font> */
    template_types string `json:"template_types";xml:"template_types"`
    
}

func (req *TaobaoDeliveryTemplateUpdateRequest) GetAPIName() string {
	return "taobao.delivery.template.update"
}

/* TaobaoDeliveryTemplateUpdateResponse 修改运费模板 */
type TaobaoDeliveryTemplateUpdateResponse struct {
    
    /* complete Basic表示修改是否成功 */
    complete bool `json:"complete";xml:"complete"`
    
}

/* TaobaoDeliveryTemplateAddRequest 增加运费模板的外部接口 */
type TaobaoDeliveryTemplateAddRequest struct {
    
    /* assumer required可选值：0、1 ，说明如下<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费 */
    assumer int64 `json:"assumer";xml:"assumer"`
    
    /* consign_area_id optional卖家发货地址区域ID
<br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm 
</font>

<br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font> */
    consign_area_id int64 `json:"consign_area_id";xml:"consign_area_id"`
    
    /* name required运费模板的名称，长度不能超过50个字节 */
    name string `json:"name";xml:"name"`
    
    /* template_add_fees required增费：输入0.00-999.99（最多包含两位小数）

<br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font>


<br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_add_fees string `json:"template_add_fees";xml:"template_add_fees"`
    
    /* template_add_standards required增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
<br/>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_add_standards string `json:"template_add_standards";xml:"template_add_standards"`
    
    /* template_dests required邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/>
<br/><font color=red>每个运费方式设置涉及的地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
<font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font> */
    template_dests string `json:"template_dests";xml:"template_dests"`
    
    /* template_start_fees required首费：输入0.00-999.99（最多包含两位小数）
<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_start_fees string `json:"template_start_fees";xml:"template_start_fees"`
    
    /* template_start_standards required首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
<br/>
<font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
    template_start_standards string `json:"template_start_standards";xml:"template_start_standards"`
    
    /* template_types required运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4 
如: post;express;ems;cod 
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分
。template_dests(指定地区)
template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font>
<br>
<font color=blue>
注意：<br/>
1、post,ems,express三种运费方式必须填写一个<br/>
2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture
只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>
3、如果是货到付款用户当没有填写cod运送方式时，运费模板会默认继承express的费用为cod的费用<br/>
4、如果是保障速递户当没有填写bzsd运送方式时，运费模板会默认继承express的费用为bzsd的费用<br/>
5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的
</font> */
    template_types string `json:"template_types";xml:"template_types"`
    
    /* valuation required可选值：0、1、3，说明如下。<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费
 <br>3:表示按宝贝体积计算运费 */
    valuation int64 `json:"valuation";xml:"valuation"`
    
}

func (req *TaobaoDeliveryTemplateAddRequest) GetAPIName() string {
	return "taobao.delivery.template.add"
}

/* TaobaoDeliveryTemplateAddResponse 增加运费模板的外部接口 */
type TaobaoDeliveryTemplateAddResponse struct {
    
    /* delivery_template Object模板对象 */
    delivery_template DeliveryTemplate `json:"delivery_template";xml:"delivery_template"`
    
}

/* TaobaoPictureUpdateRequest 修改指定图片的图片名 */
type TaobaoPictureUpdateRequest struct {
    
    /* new_name required新的图片名，最大长度50字符，不能为空 */
    new_name string `json:"new_name";xml:"new_name"`
    
    /* picture_id required要更改名字的图片的id */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureUpdateRequest) GetAPIName() string {
	return "taobao.picture.update"
}

/* TaobaoPictureUpdateResponse 修改指定图片的图片名 */
type TaobaoPictureUpdateResponse struct {
    
    /* done Basic更新是否成功 */
    done bool `json:"done";xml:"done"`
    
}

/* TaobaoOmniorderStoreSdtstatusRequest 提供给商家查询运力单的状态。 */
type TaobaoOmniorderStoreSdtstatusRequest struct {
    
    /* package_id optional菜鸟裹裹的包裹ID */
    package_id int64 `json:"package_id";xml:"package_id"`
    
}

func (req *TaobaoOmniorderStoreSdtstatusRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtstatus"
}

/* TaobaoOmniorderStoreSdtstatusResponse 提供给商家查询运力单的状态。 */
type TaobaoOmniorderStoreSdtstatusResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoDeliveryTemplateGetRequest 获取用户指定运费模板信息 */
type TaobaoDeliveryTemplateGetRequest struct {
    
    /* fields required需返回的字段列表。 <br/> 
<B>
可选值:示例中的值;各字段之间用","隔开
</B>
<br/><br/> 
<font color=blue>
template_id：查询模板ID <br/> 
template_name:查询模板名称 <br/> 
assumer：查询服务承担放 <br/> 
valuation:查询计价规则 <br/> 
supports:查询增值服务列表 <br/> 
created:查询模板创建时间 <br/> 
modified:查询修改时间<br/> 
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> 
address:卖家地址信息
</font>
<br/><br/> 
<font color=bule>
query_cod:查询货到付款运费信息<br/> 
query_post:查询平邮运费信息 <br/> 
query_express:查询快递公司运费信息 <br/> 
query_ems:查询EMS运费信息<br/> 
query_bzsd:查询普通保障速递运费信息<br/> 
query_wlb:查询物流宝保障速递运费信息<br/> 
query_furniture:查询家装物流运费信息
<font color=blue>
<br/><br/>
<font color=red>不能有空格</font> */
    fields string `json:"fields";xml:"fields"`
    
    /* template_ids required需要查询的运费模板ID列表 */
    template_ids string `json:"template_ids";xml:"template_ids"`
    
    /* user_nick optional在没有登录的情况下根据淘宝用户昵称查询指定的模板 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoDeliveryTemplateGetRequest) GetAPIName() string {
	return "taobao.delivery.template.get"
}

/* TaobaoDeliveryTemplateGetResponse 获取用户指定运费模板信息 */
type TaobaoDeliveryTemplateGetResponse struct {
    
    /* delivery_templates Object Array运费模板列表 */
    delivery_templates DeliveryTemplate `json:"delivery_templates";xml:"delivery_templates"`
    
    /* total_results Basic获得到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoPictureUserinfoGetRequest 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量 */
type TaobaoPictureUserinfoGetRequest struct {
    
}

func (req *TaobaoPictureUserinfoGetRequest) GetAPIName() string {
	return "taobao.picture.userinfo.get"
}

/* TaobaoPictureUserinfoGetResponse 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量 */
type TaobaoPictureUserinfoGetResponse struct {
    
    /* user_info Object用户使用图片空间的信息 */
    user_info UserInfo `json:"user_info";xml:"user_info"`
    
}

/* TaobaoPictureReplaceRequest 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。 */
type TaobaoPictureReplaceRequest struct {
    
    /* image_data required图片二进制文件流,不能为空,允许png、jpg、gif图片格式 */
    image_data []byte `json:"image_data";xml:"image_data"`
    
    /* picture_id required要替换的图片的id，必须大于0 */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureReplaceRequest) GetAPIName() string {
	return "taobao.picture.replace"
}

/* TaobaoPictureReplaceResponse 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。 */
type TaobaoPictureReplaceResponse struct {
    
    /* done Basic图片替换是否成功 */
    done bool `json:"done";xml:"done"`
    
}

/* TaobaoQimenItemsMarkingRequest 调用该接口，对商品进行XXXX标的打标、去标的动作。 */
type TaobaoQimenItemsMarkingRequest struct {
    
    /* action_type required操作类型，string（50），ADD=打标，DELETE=去标，必填 */
    action_type string `json:"action_type";xml:"action_type"`
    
    /* item_ids required线上商品ID，long，必填 */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
    /* remark optional备注，string（500） */
    remark string `json:"remark";xml:"remark"`
    
    /* tag_type required打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填 */
    tag_type string `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoQimenItemsMarkingRequest) GetAPIName() string {
	return "taobao.qimen.items.marking"
}

/* TaobaoQimenItemsMarkingResponse 调用该接口，对商品进行XXXX标的打标、去标的动作。 */
type TaobaoQimenItemsMarkingResponse struct {
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenTagItemsQueryRequest 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。 */
type TaobaoQimenTagItemsQueryRequest struct {
    
    /* remark optional备注，string（500） */
    remark string `json:"remark";xml:"remark"`
    
    /* tag_type required打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填 */
    tag_type string `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoQimenTagItemsQueryRequest) GetAPIName() string {
	return "taobao.qimen.tag.items.query"
}

/* TaobaoQimenTagItemsQueryResponse 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。 */
type TaobaoQimenTagItemsQueryResponse struct {
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* item_ids Basic ArrayitemIds */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
    /* tag_type BasictagType */
    tag_type string `json:"tag_type";xml:"tag_type"`
    
}

/* TaobaoPictureCategoryUpdateRequest 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。 */
type TaobaoPictureCategoryUpdateRequest struct {
    
    /* category_id required要更新的图片分类的id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* category_name optional图片分类的新名字，最大长度20字符，不能为空 */
    category_name string `json:"category_name";xml:"category_name"`
    
    /* parent_id optional图片分类的新父分类id */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
}

func (req *TaobaoPictureCategoryUpdateRequest) GetAPIName() string {
	return "taobao.picture.category.update"
}

/* TaobaoPictureCategoryUpdateResponse 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。 */
type TaobaoPictureCategoryUpdateResponse struct {
    
    /* done Basic更新图片分类是否成功 */
    done bool `json:"done";xml:"done"`
    
}

/* TaobaoFenxiaoDealerRequisitionorderCreateRequest 创建经销采购申请 */
type TaobaoFenxiaoDealerRequisitionorderCreateRequest struct {
    
    /* address required收货人所在街道地址 */
    address string `json:"address";xml:"address"`
    
    /* buyer_name required买家姓名（自提方式填写提货人姓名） */
    buyer_name string `json:"buyer_name";xml:"buyer_name"`
    
    /* city optional收货人所在市 */
    city string `json:"city";xml:"city"`
    
    /* district optional收货人所在区 */
    district string `json:"district";xml:"district"`
    
    /* id_card_number optional身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份） */
    id_card_number string `json:"id_card_number";xml:"id_card_number"`
    
    /* logistics_type optional配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货 */
    logistics_type string `json:"logistics_type";xml:"logistics_type"`
    
    /* mobile optional买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话） */
    mobile string `json:"mobile";xml:"mobile"`
    
    /* order_detail required采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价） */
    order_detail string `json:"order_detail";xml:"order_detail"`
    
    /* phone optional买家联系电话（此字段和mobile字段至少填写一个） */
    phone string `json:"phone";xml:"phone"`
    
    /* post_code optional收货人所在地区邮政编码 */
    post_code string `json:"post_code";xml:"post_code"`
    
    /* province required收货人所在省份 */
    province string `json:"province";xml:"province"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.create"
}

/* TaobaoFenxiaoDealerRequisitionorderCreateResponse 创建经销采购申请 */
type TaobaoFenxiaoDealerRequisitionorderCreateResponse struct {
    
    /* dealer_order_id Basic经销采购申请编号 */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
}

/* TaobaoQimenItemsTagQueryRequest 调用该接口，查询某个/某批商品上的标 */
type TaobaoQimenItemsTagQueryRequest struct {
    
    /* item_ids required线上淘宝商品ID，long，必填 */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
}

func (req *TaobaoQimenItemsTagQueryRequest) GetAPIName() string {
	return "taobao.qimen.items.tag.query"
}

/* TaobaoQimenItemsTagQueryResponse 调用该接口，查询某个/某批商品上的标 */
type TaobaoQimenItemsTagQueryResponse struct {
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* item_tags Object ArrayitemTags */
    item_tags ItemTag `json:"item_tags";xml:"item_tags"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoInventoryWarehouseManageRequest 创建商家仓或者更新商家仓信息 */
type TaobaoInventoryWarehouseManageRequest struct {
    
    /* ware_house_dto required仓库信息 */
    ware_house_dto WareHouseDto `json:"ware_house_dto";xml:"ware_house_dto"`
    
}

func (req *TaobaoInventoryWarehouseManageRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.manage"
}

/* TaobaoInventoryWarehouseManageResponse 创建商家仓或者更新商家仓信息 */
type TaobaoInventoryWarehouseManageResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoInventoryWarehouseGetRequest 获取商家仓信息 */
type TaobaoInventoryWarehouseGetRequest struct {
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryWarehouseGetRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.get"
}

/* TaobaoInventoryWarehouseGetResponse 获取商家仓信息 */
type TaobaoInventoryWarehouseGetResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbOutInventoryChangeNotifyRequest 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。 */
type TaobaoWlbOutInventoryChangeNotifyRequest struct {
    
    /* change_count required库存变化数量 */
    change_count int64 `json:"change_count";xml:"change_count"`
    
    /* item_id required物流宝商品id或前台宝贝id（由type类型决定） */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* op_type requiredOUT--出库 IN--入库 */
    op_type string `json:"op_type";xml:"op_type"`
    
    /* order_source_code optional订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号 */
    order_source_code string `json:"order_source_code";xml:"order_source_code"`
    
    /* out_biz_code required库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。 */
    out_biz_code string `json:"out_biz_code";xml:"out_biz_code"`
    
    /* result_count required本次库存变化后库存余额 */
    result_count int64 `json:"result_count";xml:"result_count"`
    
    /* source required（1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购 */
    source string `json:"source";xml:"source"`
    
    /* store_code optional目前非必须，系统会选择默认值 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* _type requiredWLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbOutInventoryChangeNotifyRequest) GetAPIName() string {
	return "taobao.wlb.out.inventory.change.notify"
}

/* TaobaoWlbOutInventoryChangeNotifyResponse 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。 */
type TaobaoWlbOutInventoryChangeNotifyResponse struct {
    
    /* gmt_modified Basic库存变化通知成功时间 */
    gmt_modified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
}

/* TaobaoInventoryWarehouseQueryRequest 分页查询商家仓信息 */
type TaobaoInventoryWarehouseQueryRequest struct {
    
    /* page_no required页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size required页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoInventoryWarehouseQueryRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.query"
}

/* TaobaoInventoryWarehouseQueryResponse 分页查询商家仓信息 */
type TaobaoInventoryWarehouseQueryResponse struct {
    
    /* result Objectresult */
    result PaginationResult `json:"result";xml:"result"`
    
}

/* TaobaoTradePostageUpdateRequest 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span> */
type TaobaoTradePostageUpdateRequest struct {
    
    /* post_fee required邮费价格(邮费单位是元） */
    post_fee string `json:"post_fee";xml:"post_fee"`
    
    /* tid required主订单编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradePostageUpdateRequest) GetAPIName() string {
	return "taobao.trade.postage.update"
}

/* TaobaoTradePostageUpdateResponse 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span> */
type TaobaoTradePostageUpdateResponse struct {
    
    /* trade Object返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment */
    trade Trade `json:"trade";xml:"trade"`
    
}

/* TaobaoQimenStockoutConfirmRequest 货品出库后，WMS将状态回传给ERP */
type TaobaoQimenStockoutConfirmRequest struct {
    
    /* deliveryOrder optionaldeliveryOrder */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderLines optionalorderLines */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages optionalpackages */
    packages Package `json:"packages";xml:"packages"`
    
}

func (req *TaobaoQimenStockoutConfirmRequest) GetAPIName() string {
	return "taobao.qimen.stockout.confirm"
}

/* TaobaoQimenStockoutConfirmResponse 货品出库后，WMS将状态回传给ERP */
type TaobaoQimenStockoutConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoItemDeleteRequest 删除单条商品 */
type TaobaoItemDeleteRequest struct {
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemDeleteRequest) GetAPIName() string {
	return "taobao.item.delete"
}

/* TaobaoItemDeleteResponse 删除单条商品 */
type TaobaoItemDeleteResponse struct {
    
    /* item Object被删除商品的相关信息 */
    item Item `json:"item";xml:"item"`
    
}

/* TaobaoOmniorderStoreSwitchstatusGetRequest 查询门店发货、门店自提状态 */
type TaobaoOmniorderStoreSwitchstatusGetRequest struct {
    
    /* seller_id required卖家ID */
    seller_id int64 `json:"seller_id";xml:"seller_id"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreSwitchstatusGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.switchstatus.get"
}

/* TaobaoOmniorderStoreSwitchstatusGetResponse 查询门店发货、门店自提状态 */
type TaobaoOmniorderStoreSwitchstatusGetResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* CainiaoSmartdeliveryStrategyIUpdateRequest 更新智能发货引擎发货策略设置 */
type CainiaoSmartdeliveryStrategyIUpdateRequest struct {
    
    /* delivery_strategy_info required发货策略信息 */
    delivery_strategy_info DeliveryStrategyInfo `json:"delivery_strategy_info";xml:"delivery_strategy_info"`
    
}

func (req *CainiaoSmartdeliveryStrategyIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.i.update"
}

/* CainiaoSmartdeliveryStrategyIUpdateResponse 更新智能发货引擎发货策略设置 */
type CainiaoSmartdeliveryStrategyIUpdateResponse struct {
    
    /* successful Basic设置是否成功 */
    successful bool `json:"successful";xml:"successful"`
    
}

/* CainiaoSmartdeliveryPriceofferIQueryRequest 查询智能发货引擎商家价格信息 */
type CainiaoSmartdeliveryPriceofferIQueryRequest struct {
    
    /* query_cp_price_info_request required请求参数 */
    query_cp_price_info_request QueryCpPriceInfoRequest `json:"query_cp_price_info_request";xml:"query_cp_price_info_request"`
    
}

func (req *CainiaoSmartdeliveryPriceofferIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.priceoffer.i.query"
}

/* CainiaoSmartdeliveryPriceofferIQueryResponse 查询智能发货引擎商家价格信息 */
type CainiaoSmartdeliveryPriceofferIQueryResponse struct {
    
    /* cp_price_info_list Object Array返回结果列表 */
    cp_price_info_list CpPriceInfo `json:"cp_price_info_list";xml:"cp_price_info_list"`
    
}

/* TaobaoOpenuidChangeRequest 将当前应用所属的openUId 转换为对应targetAppkey的openUid
规则：
1.如果两个appkey是应用前后台关系，可以直接转换；
2.如果appkey和targetAppkey都属于同一个开发者，不允许互相转换。
3.如果appkey和targetAppkey不属于同一个开发者，不允许互相转换。 */
type TaobaoOpenuidChangeRequest struct {
    
    /* open_uid requiredopenUid */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
    /* target_app_key required转换到的appkey */
    target_app_key string `json:"target_app_key";xml:"target_app_key"`
    
}

func (req *TaobaoOpenuidChangeRequest) GetAPIName() string {
	return "taobao.openuid.change"
}

/* TaobaoOpenuidChangeResponse 将当前应用所属的openUId 转换为对应targetAppkey的openUid
规则：
1.如果两个appkey是应用前后台关系，可以直接转换；
2.如果appkey和targetAppkey都属于同一个开发者，不允许互相转换。
3.如果appkey和targetAppkey不属于同一个开发者，不允许互相转换。 */
type TaobaoOpenuidChangeResponse struct {
    
    /* new_open_uid Basic转换到新的openUId */
    new_open_uid string `json:"new_open_uid";xml:"new_open_uid"`
    
}

/* AlibabaElectronicInvoiceGetRequest 查询已回传淘宝的电子发票,根据主订单id查询 */
type AlibabaElectronicInvoiceGetRequest struct {
    
    /* tid required淘宝主订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *AlibabaElectronicInvoiceGetRequest) GetAPIName() string {
	return "alibaba.electronic.invoice.get"
}

/* AlibabaElectronicInvoiceGetResponse 查询已回传淘宝的电子发票,根据主订单id查询 */
type AlibabaElectronicInvoiceGetResponse struct {
    
    /* invoice_detail Object电子发票详细信息 */
    invoice_detail ElectronicInvoiceDetail `json:"invoice_detail";xml:"invoice_detail"`
    
}

/* TaobaoRefundMessageAddRequest 创建退款留言/凭证 */
type TaobaoRefundMessageAddRequest struct {
    
    /* content required留言内容。最大长度: 400个字节 */
    content string `json:"content";xml:"content"`
    
    /* image required图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K */
    image []byte `json:"image";xml:"image"`
    
    /* refund_id required退款编号。 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
}

func (req *TaobaoRefundMessageAddRequest) GetAPIName() string {
	return "taobao.refund.message.add"
}

/* TaobaoRefundMessageAddResponse 创建退款留言/凭证 */
type TaobaoRefundMessageAddResponse struct {
    
    /* refund_message Object退款信息。包含id和created */
    refund_message RefundMessage `json:"refund_message";xml:"refund_message"`
    
}

/* TaobaoRefundMessagesGetRequest 查询退款留言/凭证列表 */
type TaobaoRefundMessagesGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* refund_id required退款单号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase optional退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRefundMessagesGetRequest) GetAPIName() string {
	return "taobao.refund.messages.get"
}

/* TaobaoRefundMessagesGetResponse 查询退款留言/凭证列表 */
type TaobaoRefundMessagesGetResponse struct {
    
    /* refund_messages Object Array查询到的退款留言/凭证列表 */
    refund_messages RefundMessage `json:"refund_messages";xml:"refund_messages"`
    
    /* total_results Basic查询到的退款留言/凭证总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* CainiaoSmartdeliveryPriceofferIUpdateRequest 智能发货引擎更新价格信息模板 */
type CainiaoSmartdeliveryPriceofferIUpdateRequest struct {
    
    /* cp_price_info required物流公司价格信息 */
    cp_price_info CpPriceInfo `json:"cp_price_info";xml:"cp_price_info"`
    
}

func (req *CainiaoSmartdeliveryPriceofferIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.priceoffer.i.update"
}

/* CainiaoSmartdeliveryPriceofferIUpdateResponse 智能发货引擎更新价格信息模板 */
type CainiaoSmartdeliveryPriceofferIUpdateResponse struct {
    
    /* successful Basic设置是否成功 */
    successful bool `json:"successful";xml:"successful"`
    
}

/* CainiaoSmartdeliveryCocpsIQueryRequest 获取电子面单订购关系中智能发货引擎支持的合作物流公司 */
type CainiaoSmartdeliveryCocpsIQueryRequest struct {
    
    /* send_address optional发货地址 */
    send_address Address `json:"send_address";xml:"send_address"`
    
}

func (req *CainiaoSmartdeliveryCocpsIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.cocps.i.query"
}

/* CainiaoSmartdeliveryCocpsIQueryResponse 获取电子面单订购关系中智能发货引擎支持的合作物流公司 */
type CainiaoSmartdeliveryCocpsIQueryResponse struct {
    
    /* smart_delivery_collaborate_cps_info_list Object Array返回结果 */
    smart_delivery_collaborate_cps_info_list SmartDeliveryCollaborateCpsInfo `json:"smart_delivery_collaborate_cps_info_list";xml:"smart_delivery_collaborate_cps_info_list"`
    
}

/* TaobaoTraderateImprImprwordsGetRequest 根据淘宝后台类目的一级类目和叶子类目 */
type TaobaoTraderateImprImprwordsGetRequest struct {
    
    /* cat_leaf_id optional淘宝叶子类目id */
    cat_leaf_id int64 `json:"cat_leaf_id";xml:"cat_leaf_id"`
    
    /* cat_root_id required淘宝一级类目id */
    cat_root_id int64 `json:"cat_root_id";xml:"cat_root_id"`
    
}

func (req *TaobaoTraderateImprImprwordsGetRequest) GetAPIName() string {
	return "taobao.traderate.impr.imprwords.get"
}

/* TaobaoTraderateImprImprwordsGetResponse 根据淘宝后台类目的一级类目和叶子类目 */
type TaobaoTraderateImprImprwordsGetResponse struct {
    
    /* impr_words Basic Array返回类目下所有大家印象的标签 */
    impr_words string `json:"impr_words";xml:"impr_words"`
    
}

/* TaobaoItemcatsGetRequest 获取后台供卖家发布商品的标准商品类目。 */
type TaobaoItemcatsGetRequest struct {
    
    /* cids special商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个) */
    cids int64 `json:"cids";xml:"cids"`
    
    /* datetime special无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化 */
    datetime time.Time `json:"datetime";xml:"datetime"`
    
    /* fields optional需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 */
    fields string `json:"fields";xml:"fields"`
    
    /* parent_cid special父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个) */
    parent_cid int64 `json:"parent_cid";xml:"parent_cid"`
    
}

func (req *TaobaoItemcatsGetRequest) GetAPIName() string {
	return "taobao.itemcats.get"
}

/* TaobaoItemcatsGetResponse 获取后台供卖家发布商品的标准商品类目。 */
type TaobaoItemcatsGetResponse struct {
    
    /* item_cats Object Array增量类目信息,根据fields传入的参数返回相应的结果；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 */
    item_cats ItemCat `json:"item_cats";xml:"item_cats"`
    
    /* last_modified Basic最近修改时间(如果取增量，会返回该字段)。 */
    last_modified time.Time `json:"last_modified";xml:"last_modified"`
    
}

/* TaobaoItempropsGetRequest 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。 */
type TaobaoItempropsGetRequest struct {
    
    /* attr_keys optional属性的Key，支持多条，以“,”分隔 */
    attr_keys string `json:"attr_keys";xml:"attr_keys"`
    
    /* child_path optional类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid */
    child_path string `json:"child_path";xml:"child_path"`
    
    /* cid required叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID */
    cid int64 `json:"cid";xml:"cid"`
    
    /* datetime optional增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime) */
    datetime time.Time `json:"datetime";xml:"datetime"`
    
    /* fields optional需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values */
    fields string `json:"fields";xml:"fields"`
    
    /* is_color_prop optional是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
    is_color_prop bool `json:"is_color_prop";xml:"is_color_prop"`
    
    /* is_enum_prop optional是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。 */
    is_enum_prop bool `json:"is_enum_prop";xml:"is_enum_prop"`
    
    /* is_input_prop optional在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
    is_input_prop bool `json:"is_input_prop";xml:"is_input_prop"`
    
    /* is_item_prop optional是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否) */
    is_item_prop bool `json:"is_item_prop";xml:"is_item_prop"`
    
    /* is_key_prop optional是否关键属性。可选值:true(是),false(否) */
    is_key_prop bool `json:"is_key_prop";xml:"is_key_prop"`
    
    /* is_sale_prop optional是否销售属性。可选值:true(是),false(否) */
    is_sale_prop bool `json:"is_sale_prop";xml:"is_sale_prop"`
    
    /* parent_pid optional父属性ID */
    parent_pid int64 `json:"parent_pid";xml:"parent_pid"`
    
    /* pid optional属性id (取类目属性时，传pid，不用同时传PID和parent_pid) */
    pid int64 `json:"pid";xml:"pid"`
    
    /* _type optional获取类目的类型：1代表集市、2代表天猫 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItempropsGetRequest) GetAPIName() string {
	return "taobao.itemprops.get"
}

/* TaobaoItempropsGetResponse 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。 */
type TaobaoItempropsGetResponse struct {
    
    /* item_props Object Array类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果 */
    item_props ItemProp `json:"item_props";xml:"item_props"`
    
    /* last_modified Basic最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss */
    last_modified time.Time `json:"last_modified";xml:"last_modified"`
    
}

/* CainiaoSmartdeliveryStrategyIQueryRequest 查询智能发货引擎发货策略设置 */
type CainiaoSmartdeliveryStrategyIQueryRequest struct {
    
    /* send_address optional发货地址 */
    send_address Address `json:"send_address";xml:"send_address"`
    
}

func (req *CainiaoSmartdeliveryStrategyIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.i.query"
}

/* CainiaoSmartdeliveryStrategyIQueryResponse 查询智能发货引擎发货策略设置 */
type CainiaoSmartdeliveryStrategyIQueryResponse struct {
    
    /* delivery_strategy_info_list Object Array返回结果列表 */
    delivery_strategy_info_list DeliveryStrategyInfo `json:"delivery_strategy_info_list";xml:"delivery_strategy_info_list"`
    
}

/* TaobaoTraderateImprImprwordByfeedidGetRequest 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签 */
type TaobaoTraderateImprImprwordByfeedidGetRequest struct {
    
    /* child_trade_id required交易订单id号（如果包含子订单，请使用子订单id号） */
    child_trade_id int64 `json:"child_trade_id";xml:"child_trade_id"`
    
}

func (req *TaobaoTraderateImprImprwordByfeedidGetRequest) GetAPIName() string {
	return "taobao.traderate.impr.imprword.byfeedid.get"
}

/* TaobaoTraderateImprImprwordByfeedidGetResponse 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签 */
type TaobaoTraderateImprImprwordByfeedidGetResponse struct {
    
    /* impr_feed Object根据子订单和买家昵称找到的评价和印象词结果 */
    impr_feed ImprFeedIdDO `json:"impr_feed";xml:"impr_feed"`
    
}

/* TaobaoTimeGetRequest 获取淘宝系统当前时间 */
type TaobaoTimeGetRequest struct {
    
}

func (req *TaobaoTimeGetRequest) GetAPIName() string {
	return "taobao.time.get"
}

/* TaobaoTimeGetResponse 获取淘宝系统当前时间 */
type TaobaoTimeGetResponse struct {
    
    /* time Basic淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss */
    time time.Time `json:"time";xml:"time"`
    
}

/* TaobaoPictureCategoryGetRequest 获取图片分类信息 */
type TaobaoPictureCategoryGetRequest struct {
    
    /* modified_time optional图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。 */
    modified_time time.Time `json:"modified_time";xml:"modified_time"`
    
    /* parent_id optional取二级分类时设置为对应父分类id
取一级分类时父分类id设为0
取全部分类的时候不设或设为-1 */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
    /* picture_category_id optional图片分类ID */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_category_name optional图片分类名，不支持模糊查询 */
    picture_category_name string `json:"picture_category_name";xml:"picture_category_name"`
    
    /* _type optional1 */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoPictureCategoryGetRequest) GetAPIName() string {
	return "taobao.picture.category.get"
}

/* TaobaoPictureCategoryGetResponse 获取图片分类信息 */
type TaobaoPictureCategoryGetResponse struct {
    
    /* picture_categories Object Array图片分类 */
    picture_categories PictureCategory `json:"picture_categories";xml:"picture_categories"`
    
}

/* TaobaoOmniorderStoreAccpetedRequest ISV Pos端门店接单，通知星盘 */
type TaobaoOmniorderStoreAccpetedRequest struct {
    
    /* report_timestamp requiredISV系统上报时间 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* sub_order_list required子订单列表 */
    sub_order_list StoreAcceptedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreAccpetedRequest) GetAPIName() string {
	return "taobao.omniorder.store.accpeted"
}

/* TaobaoOmniorderStoreAccpetedResponse ISV Pos端门店接单，通知星盘 */
type TaobaoOmniorderStoreAccpetedResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic错误内容 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoPictureDeleteRequest 删除图片空间图片 */
type TaobaoPictureDeleteRequest struct {
    
    /* picture_ids required图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100 */
    picture_ids string `json:"picture_ids";xml:"picture_ids"`
    
}

func (req *TaobaoPictureDeleteRequest) GetAPIName() string {
	return "taobao.picture.delete"
}

/* TaobaoPictureDeleteResponse 删除图片空间图片 */
type TaobaoPictureDeleteResponse struct {
    
    /* success Basic是否删除 */
    success string `json:"success";xml:"success"`
    
}

/* TaobaoOmniorderStoreRefusedRequest ISV Pos端门店拒单，通知星盘 */
type TaobaoOmniorderStoreRefusedRequest struct {
    
    /* report_timestamp requiredISV的系统时间 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* sub_order_list required子订单列表 */
    sub_order_list SubOrder `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreRefusedRequest) GetAPIName() string {
	return "taobao.omniorder.store.refused"
}

/* TaobaoOmniorderStoreRefusedResponse ISV Pos端门店拒单，通知星盘 */
type TaobaoOmniorderStoreRefusedResponse struct {
    
    /* err_code Basic正常为0,其他表示异常 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoPictureGetRequest 获取图片信息 */
type TaobaoPictureGetRequest struct {
    
    /* client_type optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* deleted optional是否删除,unfroze代表没有删除 */
    deleted string `json:"deleted";xml:"deleted"`
    
    /* end_date optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* is_https optional是否获取https的链接 */
    is_https bool `json:"is_https";xml:"is_https"`
    
    /* modified_time optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
    modified_time time.Time `json:"modified_time";xml:"modified_time"`
    
    /* order_by optional图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数.每页返回最多返回100条,默认值40 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* picture_category_id optional图片分类ID */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional图片ID */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* start_date optional查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* title optional图片标题,最大长度50字符,中英文都算一字符 */
    title string `json:"title";xml:"title"`
    
    /* urls optional图片url查询接口 */
    urls string `json:"urls";xml:"urls"`
    
}

func (req *TaobaoPictureGetRequest) GetAPIName() string {
	return "taobao.picture.get"
}

/* TaobaoPictureGetResponse 获取图片信息 */
type TaobaoPictureGetResponse struct {
    
    /* pictures Object Array图片信息列表 */
    pictures Picture `json:"pictures";xml:"pictures"`
    
    /* totalResults Basic总的结果数 */
    totalResults int64 `json:"totalResults";xml:"totalResults"`
    
}

/* TaobaoOmniorderStoreConsignedRequest ISV Pos端门店发货，通知星盘 */
type TaobaoOmniorderStoreConsignedRequest struct {
    
    /* report_timestamp requiredISV系统上报时间 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* sub_order_list required子订单列表 */
    sub_order_list StoreConsignedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreConsignedRequest) GetAPIName() string {
	return "taobao.omniorder.store.consigned"
}

/* TaobaoOmniorderStoreConsignedResponse ISV Pos端门店发货，通知星盘 */
type TaobaoOmniorderStoreConsignedResponse struct {
    
    /* data Objectdata */
    data StoreConsignedResponse `json:"data";xml:"data"`
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic错误内容 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoPictureUploadRequest 图片空间上传接口 */
type TaobaoPictureUploadRequest struct {
    
    /* client_type optional图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* image_input_title required包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名 */
    image_input_title string `json:"image_input_title";xml:"image_input_title"`
    
    /* img required图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。 */
    img []byte `json:"img";xml:"img"`
    
    /* is_https optional是否获取https连接 */
    is_https bool `json:"is_https";xml:"is_https"`
    
    /* picture_category_id required图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类 */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional如果此参数大于0，而且在后台能查到对应的图片，则此次上传为原图替换 */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* title optional图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1 */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoPictureUploadRequest) GetAPIName() string {
	return "taobao.picture.upload"
}

/* TaobaoPictureUploadResponse 图片空间上传接口 */
type TaobaoPictureUploadResponse struct {
    
    /* picture Object当前上传的一张图片信息 */
    picture Picture `json:"picture";xml:"picture"`
    
}

/* TaobaoTopSecretRegisterRequest 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密 */
type TaobaoTopSecretRegisterRequest struct {
    
    /* user_id optional用户id，保证唯一 */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoTopSecretRegisterRequest) GetAPIName() string {
	return "taobao.top.secret.register"
}

/* TaobaoTopSecretRegisterResponse 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密 */
type TaobaoTopSecretRegisterResponse struct {
    
    /* result Basic返回操作是否成功 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoMediaCategoryUpdateRequest 更新媒体文件分类 */
type TaobaoMediaCategoryUpdateRequest struct {
    
    /* media_category_id required文件分类ID,不能为空 */
    media_category_id int64 `json:"media_category_id";xml:"media_category_id"`
    
    /* media_category_name required文件分类名，最大长度20字符，中英文都算一字符,不能为空 */
    media_category_name string `json:"media_category_name";xml:"media_category_name"`
    
}

func (req *TaobaoMediaCategoryUpdateRequest) GetAPIName() string {
	return "taobao.media.category.update"
}

/* TaobaoMediaCategoryUpdateResponse 更新媒体文件分类 */
type TaobaoMediaCategoryUpdateResponse struct {
    
    /* success Basic更新是否成功标志 */
    success bool `json:"success";xml:"success"`
    
}

/* TaobaoOmniorderAllocatedinfoSyncRequest ISV分单完成，将分单结果同步给星盘 */
type TaobaoOmniorderAllocatedinfoSyncRequest struct {
    
    /* message optional分单结果消息, 如果status为AllocateFail, 则表示失败的理由. */
    message string `json:"message";xml:"message"`
    
    /* report_timestamp required1231243213213 */
    report_timestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    /* status required分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail) */
    status string `json:"status";xml:"status"`
    
    /* sub_order_list required门店的分单列表 */
    sub_order_list StoreAllocatedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    /* tid required淘宝交易主订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_id optional跟踪Id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderAllocatedinfoSyncRequest) GetAPIName() string {
	return "taobao.omniorder.allocatedinfo.sync"
}

/* TaobaoOmniorderAllocatedinfoSyncResponse ISV分单完成，将分单结果同步给星盘 */
type TaobaoOmniorderAllocatedinfoSyncResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic错误内容 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenOrderexceptionReportRequest WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况 */
type TaobaoQimenOrderexceptionReportRequest struct {
    
    /* createTime optional奇门仓储字段 */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderCode optional奇门仓储字段 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional奇门仓储字段 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* expressCode optional奇门仓储字段 */
    expressCode string `json:"expressCode";xml:"expressCode"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* logisticsCode optional奇门仓储字段 */
    logisticsCode string `json:"logisticsCode";xml:"logisticsCode"`
    
    /* messageDesc optional奇门仓储字段 */
    messageDesc string `json:"messageDesc";xml:"messageDesc"`
    
    /* messageId optional奇门仓储字段 */
    messageId string `json:"messageId";xml:"messageId"`
    
    /* messageType optional奇门仓储字段 */
    messageType string `json:"messageType";xml:"messageType"`
    
    /* orderLines optional奇门仓储字段 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* orderType optional奇门仓储字段 */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* remark optional奇门仓储字段 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional奇门仓储字段 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderexceptionReportRequest) GetAPIName() string {
	return "taobao.qimen.orderexception.report"
}

/* TaobaoQimenOrderexceptionReportResponse WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况 */
type TaobaoQimenOrderexceptionReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoTradesSoldIncrementGetRequest 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type） */
type TaobaoTradesSoldIncrementGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_modified required查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* ext_type optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
    ext_type string `json:"ext_type";xml:"ext_type"`
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0为处方药未审核状态 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* rate_status optional评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评) */
    rate_status string `json:"rate_status";xml:"rate_status"`
    
    /* start_modified required查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
    /* status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 */
    status string `json:"status";xml:"status"`
    
    /* tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
    tag string `json:"tag";xml:"tag"`
    
    /* _type optional交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。 */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldIncrementGetRequest) GetAPIName() string {
	return "taobao.trades.sold.increment.get"
}

/* TaobaoTradesSoldIncrementGetResponse 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type） */
type TaobaoTradesSoldIncrementGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的交易信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
    trades Trade `json:"trades";xml:"trades"`
    
}

/* TaobaoFenxiaoRefundGetRequest 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息 */
type TaobaoFenxiaoRefundGetRequest struct {
    
    /* query_seller_refund optional是否查询下游买家的退款信息 */
    query_seller_refund bool `json:"query_seller_refund";xml:"query_seller_refund"`
    
    /* sub_order_id required要查询的退款子单的id */
    sub_order_id int64 `json:"sub_order_id";xml:"sub_order_id"`
    
}

func (req *TaobaoFenxiaoRefundGetRequest) GetAPIName() string {
	return "taobao.fenxiao.refund.get"
}

/* TaobaoFenxiaoRefundGetResponse 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息 */
type TaobaoFenxiaoRefundGetResponse struct {
    
    /* refund_detail Object退款详情 */
    refund_detail RefundDetail `json:"refund_detail";xml:"refund_detail"`
    
}

/* TaobaoQimenInventoryreserveCancelRequest 库存预占取消 */
type TaobaoQimenInventoryreserveCancelRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemInventories optional奇门仓储字段 */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* orderCode optional奇门仓储字段 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderSource optional奇门仓储字段 */
    orderSource string `json:"orderSource";xml:"orderSource"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenInventoryreserveCancelRequest) GetAPIName() string {
	return "taobao.qimen.inventoryreserve.cancel"
}

/* TaobaoQimenInventoryreserveCancelResponse 库存预占取消 */
type TaobaoQimenInventoryreserveCancelResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* isRetry Basic奇门仓储字段 */
    isRetry string `json:"isRetry";xml:"isRetry"`
    
    /* itemInventories Object Array奇门仓储字段 */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderCode Basic奇门仓储字段 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
}

/* TaobaoQimenInventoryreserveCreateRequest ERP调用奇门的接口,查询发货库存预占用信息 */
type TaobaoQimenInventoryreserveCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemInventories optional交易订单信息 */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* maxWarehouseNum optional最大仓库数(当一个仓库不满足库存时.是否允许占用多个仓库库存.如果允许.需要指定最大分仓数量.但不能拆分订单明细. 0：不限个数.只要满足发货.不限分占几个仓库 1：默认值.只能单仓发 >1:最大 占用数量) */
    maxWarehouseNum int64 `json:"maxWarehouseNum";xml:"maxWarehouseNum"`
    
    /* orderCode requiredERP订单编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderSource required订单来源(213 天猫、201 淘宝、214 京东、202 1688 阿里中文站、203 苏宁在线、204 亚马逊中国、205 当当、208 1号店、207 唯品会、209 国美在线、210 拍拍、206 易贝ebay、211 聚美优品、212 乐蜂 网、215 邮乐、216 凡客、217 优购、218 银泰、219 易讯、221 聚尚网、222 蘑菇街、223 POS门店、301 其他) */
    orderSource int64 `json:"orderSource";xml:"orderSource"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* receiverInfo optional收件者信息 */
    receiverInfo ReceiverInfo `json:"receiverInfo";xml:"receiverInfo"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventoryreserveCreateRequest) GetAPIName() string {
	return "taobao.qimen.inventoryreserve.create"
}

/* TaobaoQimenInventoryreserveCreateResponse ERP调用奇门的接口,查询发货库存预占用信息 */
type TaobaoQimenInventoryreserveCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* isRetry Basic是否需要重试(Y/N (默认为N)返回Y时建议系统自动重试) */
    isRetry string `json:"isRetry";xml:"isRetry"`
    
    /* itemInventories Object Array */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderCode BasicERP订单编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
}

/* TaobaoFenxiaoDistributorProductsGetRequest 分销商查询供应商产品信息 */
type TaobaoFenxiaoDistributorProductsGetRequest struct {
    
    /* download_status optionaldownload_status */
    download_status string `json:"download_status";xml:"download_status"`
    
    /* end_time optional结束时间 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* fields optional指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
    fields string `json:"fields";xml:"fields"`
    
    /* item_ids optional根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
    item_ids int64 `json:"item_ids";xml:"item_ids"`
    
    /* order_by optionalorder_by */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* pids optional产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005” */
    pids int64 `json:"pids";xml:"pids"`
    
    /* productcat_id optional产品线ID */
    productcat_id int64 `json:"productcat_id";xml:"productcat_id"`
    
    /* start_time optional开始修改时间 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
    /* supplier_nick optional供应商nick，分页查询时，必传 */
    supplier_nick string `json:"supplier_nick";xml:"supplier_nick"`
    
    /* time_type optionaltime_type */
    time_type string `json:"time_type";xml:"time_type"`
    
    /* trade_type optionaltrade_type */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoDistributorProductsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributor.products.get"
}

/* TaobaoFenxiaoDistributorProductsGetResponse 分销商查询供应商产品信息 */
type TaobaoFenxiaoDistributorProductsGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* products Object Array产品对象记录集。返回 FenxiaoProduct 包含的字段信息。 */
    products FenxiaoProduct `json:"products";xml:"products"`
    
}

/* TaobaoWlbOrderJzpartnerQueryRequest 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口 */
type TaobaoWlbOrderJzpartnerQueryRequest struct {
    
    /* service_type optionalserviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的） */
    service_type int64 `json:"service_type";xml:"service_type"`
    
    /* taobao_trade_id optional淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。 */
    taobao_trade_id int64 `json:"taobao_trade_id";xml:"taobao_trade_id"`
    
}

func (req *TaobaoWlbOrderJzpartnerQueryRequest) GetAPIName() string {
	return "taobao.wlb.order.jzpartner.query"
}

/* TaobaoWlbOrderJzpartnerQueryResponse 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口 */
type TaobaoWlbOrderJzpartnerQueryResponse struct {
    
    /* install_list Object Array安装服务商列表 */
    install_list PartnerNew `json:"install_list";xml:"install_list"`
    
    /* is_success Basic接口查询成功或者失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* result_info Basic查询返回信息，如果失败，存储错误信息 */
    result_info string `json:"result_info";xml:"result_info"`
    
    /* server_list Object Array物流配送服务商对象列表 */
    server_list PartnerNew `json:"server_list";xml:"server_list"`
    
}

/* TaobaoWlbOrderJzwithinsConsignRequest 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口 */
type TaobaoWlbOrderJzwithinsConsignRequest struct {
    
    /* ins_partner optional物流服务商信息 */
    ins_partner JzPartnerNew `json:"ins_partner";xml:"ins_partner"`
    
    /* jz_consign_args required家装物流发货参数 */
    jz_consign_args JzConsignArgsNew `json:"jz_consign_args";xml:"jz_consign_args"`
    
    /* tid required淘宝交易订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* tms_partner required物流服务商信息 */
    tms_partner JzPartnerNew `json:"tms_partner";xml:"tms_partner"`
    
}

func (req *TaobaoWlbOrderJzwithinsConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.jzwithins.consign"
}

/* TaobaoWlbOrderJzwithinsConsignResponse 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口 */
type TaobaoWlbOrderJzwithinsConsignResponse struct {
    
    /* is_success Basic发货成功或者失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* result_info Basic发货返回信息，如果发货错误则报出对应错误 */
    result_info string `json:"result_info";xml:"result_info"`
    
}

/* TaobaoTradeConfirmfeeGetRequest 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用 */
type TaobaoTradeConfirmfeeGetRequest struct {
    
    /* tid required交易主订单或子订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeConfirmfeeGetRequest) GetAPIName() string {
	return "taobao.trade.confirmfee.get"
}

/* TaobaoTradeConfirmfeeGetResponse 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用 */
type TaobaoTradeConfirmfeeGetResponse struct {
    
    /* trade_confirm_fee Object获取到的交易确认收货费用 */
    trade_confirm_fee TradeConfirmFee `json:"trade_confirm_fee";xml:"trade_confirm_fee"`
    
}

/* TaobaoItemJointImgRequest * 关联一张商品图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额 */
type TaobaoItemJointImgRequest struct {
    
    /* id optional商品图片id(如果是更新图片，则需要传该参数) */
    id int64 `json:"id";xml:"id"`
    
    /* is_major optional上传的图片是否关联为商品主图 */
    is_major bool `json:"is_major";xml:"is_major"`
    
    /* is_rectangle optional是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图 */
    is_rectangle bool `json:"is_rectangle";xml:"is_rectangle"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* pic_path required图片URL,图片空间图片的相对地址 */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* position optional图片序号 */
    position int64 `json:"position";xml:"position"`
    
}

func (req *TaobaoItemJointImgRequest) GetAPIName() string {
	return "taobao.item.joint.img"
}

/* TaobaoItemJointImgResponse * 关联一张商品图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额 */
type TaobaoItemJointImgResponse struct {
    
    /* item_img Object商品图片信息 */
    item_img ItemImg `json:"item_img";xml:"item_img"`
    
}

/* TaobaoItemJointPropimgRequest * 关联一张商品属性图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张） */
type TaobaoItemJointPropimgRequest struct {
    
    /* id optional属性图片ID。如果是新增不需要填写 */
    id int64 `json:"id";xml:"id"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* pic_path required图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded ) */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* position optional图片序号 */
    position int64 `json:"position";xml:"position"`
    
    /* properties required属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemJointPropimgRequest) GetAPIName() string {
	return "taobao.item.joint.propimg"
}

/* TaobaoItemJointPropimgResponse * 关联一张商品属性图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张） */
type TaobaoItemJointPropimgResponse struct {
    
    /* prop_img Object属性图片对象信息 */
    prop_img PropImg `json:"prop_img";xml:"prop_img"`
    
}

/* TaobaoItemsCustomGetRequest 跟据卖家设定的商品外部id获取商品 
这个商品对应卖家从传入的session中获取，需要session绑定 */
type TaobaoItemsCustomGetRequest struct {
    
    /* fields required需返回的字段列表，参考：Item商品结构体说明，其中barcode、sku.barcode等条形码字段暂不支持；多个字段之间用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* outer_id required商品的外部商品ID，支持批量，最多不超过40个。 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
}

func (req *TaobaoItemsCustomGetRequest) GetAPIName() string {
	return "taobao.items.custom.get"
}

/* TaobaoItemsCustomGetResponse 跟据卖家设定的商品外部id获取商品 
这个商品对应卖家从传入的session中获取，需要session绑定 */
type TaobaoItemsCustomGetResponse struct {
    
    /* items Object Array商品列表，具体返回字段以fields决定 */
    items Item `json:"items";xml:"items"`
    
}

/* TaobaoItemsInventoryGetRequest 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取 */
type TaobaoItemsInventoryGetRequest struct {
    
    /* auction_type optional商品类型：a-拍卖,b-一口价 */
    auction_type string `json:"auction_type";xml:"auction_type"`
    
    /* banner optional分类字段。可选值:<br>
regular_shelved(定时上架)<br>
never_on_shelf(从未上架)<br>
off_shelf(我下架的)<br>
<font color='red'>for_shelved(等待所有上架)<br>
sold_out(全部卖完)<br>
violation_off_shelf(违规下架的)<br>
默认查询for_shelved(等待所有上架)这个状态的商品<br></font>
注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的) */
    banner string `json:"banner";xml:"banner"`
    
    /* cid optional商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到 */
    cid int64 `json:"cid";xml:"cid"`
    
    /* end_modified optional商品结束修改时间 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* fields required需返回的字段列表。可选值为 Item 商品结构体中的以下字段： 
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。 */
    fields string `json:"fields";xml:"fields"`
    
    /* has_discount optional是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
    has_discount bool `json:"has_discount";xml:"has_discount"`
    
    /* is_ex optional商品是否在外部网店显示 */
    is_ex bool `json:"is_ex";xml:"is_ex"`
    
    /* is_taobao optional商品是否在淘宝显示 */
    is_taobao bool `json:"is_taobao";xml:"is_taobao"`
    
    /* order_by optional排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数;最大值：200；默认值：40。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* q optional搜索字段。搜索商品的title。 */
    q string `json:"q";xml:"q"`
    
    /* seller_cids optional卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
    seller_cids string `json:"seller_cids";xml:"seller_cids"`
    
    /* start_modified optional商品起始修改时间 */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoItemsInventoryGetRequest) GetAPIName() string {
	return "taobao.items.inventory.get"
}

/* TaobaoItemsInventoryGetResponse 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取 */
type TaobaoItemsInventoryGetResponse struct {
    
    /* items Object Array搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段 */
    items Item `json:"items";xml:"items"`
    
    /* total_results Basic搜索到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoItemcatsAuthorizeGetRequest 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表 */
type TaobaoItemcatsAuthorizeGetRequest struct {
    
    /* fields required需要返回的字段。目前支持有：
brand.vid, brand.name, 
item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,
xinpin_item_cat.cid, 
xinpin_item_cat.name, 
xinpin_item_cat.status,
xinpin_item_cat.sort_order,
xinpin_item_cat.parent_cid,
xinpin_item_cat.is_parent */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoItemcatsAuthorizeGetRequest) GetAPIName() string {
	return "taobao.itemcats.authorize.get"
}

/* TaobaoItemcatsAuthorizeGetResponse 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表 */
type TaobaoItemcatsAuthorizeGetResponse struct {
    
    /* seller_authorize Object里面有3个数组：
Brand[]品牌列表,
ItemCat[] 类目列表
XinpinItemCat[] 针对于C卖家新品类目列表 */
    seller_authorize SellerAuthorize `json:"seller_authorize";xml:"seller_authorize"`
    
}

/* Gfq1qp4287ShopexMatrixDatapushRequest 订单转发 */
type Gfq1qp4287ShopexMatrixDatapushRequest struct {
    
    /* p optional数据内容 */
    p string `json:"p";xml:"p"`
    
}

func (req *Gfq1qp4287ShopexMatrixDatapushRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.matrix.datapush"
}

/* Gfq1qp4287ShopexMatrixDatapushResponse 订单转发 */
type Gfq1qp4287ShopexMatrixDatapushResponse struct {
    
    /* data Basicdata */
    data string `json:"data";xml:"data"`
    
    /* res Basicres */
    res string `json:"res";xml:"res"`
    
}

/* TaobaoSkusCustomGetRequest 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku 
这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用) */
type TaobaoSkusCustomGetRequest struct {
    
    /* fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开 */
    fields string `json:"fields";xml:"fields"`
    
    /* outer_id requiredSku的外部商家ID */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
}

func (req *TaobaoSkusCustomGetRequest) GetAPIName() string {
	return "taobao.skus.custom.get"
}

/* TaobaoSkusCustomGetResponse 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku 
这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用) */
type TaobaoSkusCustomGetResponse struct {
    
    /* skus Object ArraySku对象，具体字段以fields决定 */
    skus Sku `json:"skus";xml:"skus"`
    
}

/* TaobaoSellercenterRolesGetRequest 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。 */
type TaobaoSellercenterRolesGetRequest struct {
    
    /* nick required卖家昵称(只允许查询自己的信息：当前登陆者) */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterRolesGetRequest) GetAPIName() string {
	return "taobao.sellercenter.roles.get"
}

/* TaobaoSellercenterRolesGetResponse 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。 */
type TaobaoSellercenterRolesGetResponse struct {
    
    /* roles Object Array卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点) */
    roles Role `json:"roles";xml:"roles"`
    
}

/* TaobaoSellercenterRoleInfoGetRequest 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息) */
type TaobaoSellercenterRoleInfoGetRequest struct {
    
    /* role_id required角色id */
    role_id int64 `json:"role_id";xml:"role_id"`
    
}

func (req *TaobaoSellercenterRoleInfoGetRequest) GetAPIName() string {
	return "taobao.sellercenter.role.info.get"
}

/* TaobaoSellercenterRoleInfoGetResponse 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息) */
type TaobaoSellercenterRoleInfoGetResponse struct {
    
    /* role Object角色具体信息 */
    role Role `json:"role";xml:"role"`
    
}

/* TaobaoSellercenterRoleAddRequest 给指定的卖家创建新的子账号角色<br/>
如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
---------|---------|code=phone 手机淘宝店铺<br/>
调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/> */
type TaobaoSellercenterRoleAddRequest struct {
    
    /* description optional角色描述 */
    description string `json:"description";xml:"description"`
    
    /* name required角色名 */
    name string `json:"name";xml:"name"`
    
    /* nick required表示卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* permission_codes optional需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。 */
    permission_codes string `json:"permission_codes";xml:"permission_codes"`
    
}

func (req *TaobaoSellercenterRoleAddRequest) GetAPIName() string {
	return "taobao.sellercenter.role.add"
}

/* TaobaoSellercenterRoleAddResponse 给指定的卖家创建新的子账号角色<br/>
如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
---------|---------|code=phone 手机淘宝店铺<br/>
调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/>
code=sell 宝贝管理<br/>
---------|code=sm 店铺管理<br/>
---------|---------|code=sm-design 如店铺装修<br/>
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/> */
type TaobaoSellercenterRoleAddResponse struct {
    
    /* role Object子账号角色 */
    role Role `json:"role";xml:"role"`
    
}

/* CainiaoSmartdeliveryCpIModifyRequest 商家修改智能发货引擎推荐的cp */
type CainiaoSmartdeliveryCpIModifyRequest struct {
    
    /* modify_smart_delivery_cp_request required修改智选CP请求 */
    modify_smart_delivery_cp_request ModifySmartDeliveryCpRequest `json:"modify_smart_delivery_cp_request";xml:"modify_smart_delivery_cp_request"`
    
}

func (req *CainiaoSmartdeliveryCpIModifyRequest) GetAPIName() string {
	return "cainiao.smartdelivery.cp.i.modify"
}

/* CainiaoSmartdeliveryCpIModifyResponse 商家修改智能发货引擎推荐的cp */
type CainiaoSmartdeliveryCpIModifyResponse struct {
    
    /* modify_smart_delivery_cp_response Object更新智能发货智选cp返回结果 */
    modify_smart_delivery_cp_response ModifySmartDeliveryCpResponse `json:"modify_smart_delivery_cp_response";xml:"modify_smart_delivery_cp_response"`
    
}

/* TaobaoTmcMessagesProduceRequest 批量发送消息 */
type TaobaoTmcMessagesProduceRequest struct {
    
    /* messages requiredtmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}] */
    messages TmcPublishMessage `json:"messages";xml:"messages"`
    
}

func (req *TaobaoTmcMessagesProduceRequest) GetAPIName() string {
	return "taobao.tmc.messages.produce"
}

/* TaobaoTmcMessagesProduceResponse 批量发送消息 */
type TaobaoTmcMessagesProduceResponse struct {
    
    /* is_all_success Basic是否全部成功 */
    is_all_success bool `json:"is_all_success";xml:"is_all_success"`
    
    /* results Object Array发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功 */
    results TmcProduceResult `json:"results";xml:"results"`
    
}

/* CainiaoSmartdeliveryIGetRequest 获取智选cp和电子面单信息 */
type CainiaoSmartdeliveryIGetRequest struct {
    
    /* smart_delivery_batch_request required<a href="http://open.taobao.com/docs/doc.htm?treeId=319&articleId=106295&docType=1">智能发货引擎</a>批量请求参数 */
    smart_delivery_batch_request SmartDeliveryBatchRequest `json:"smart_delivery_batch_request";xml:"smart_delivery_batch_request"`
    
}

func (req *CainiaoSmartdeliveryIGetRequest) GetAPIName() string {
	return "cainiao.smartdelivery.i.get"
}

/* CainiaoSmartdeliveryIGetResponse 获取智选cp和电子面单信息 */
type CainiaoSmartdeliveryIGetResponse struct {
    
    /* smart_delivery_response_wrapper_list Object Array<a href="http://open.taobao.com/docs/doc.htm?treeId=319&articleId=106295&docType=1">智能发货引擎</a>结果包装类 */
    smart_delivery_response_wrapper_list SmartDeliveryResponseWrapper `json:"smart_delivery_response_wrapper_list";xml:"smart_delivery_response_wrapper_list"`
    
}

/* TaobaoFenxiaoOrdersGetRequest 分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。 */
type TaobaoFenxiaoOrdersGetRequest struct {
    
    /* end_created optional结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
    end_created time.Time `json:"end_created";xml:"end_created"`
    
    /* fields optional多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。（大于0的整数。默认为1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。（每页条数不超过50条） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* purchase_order_id optional采购单编号或分销流水号，若其它参数没传，则此参数必传。 */
    purchase_order_id int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
    /* start_created optional起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
    start_created time.Time `json:"start_created";xml:"start_created"`
    
    /* status optional交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) */
    status string `json:"status";xml:"status"`
    
    /* tc_order_id optional采购单下游买家订单id */
    tc_order_id int64 `json:"tc_order_id";xml:"tc_order_id"`
    
    /* time_type optional可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询) */
    time_type string `json:"time_type";xml:"time_type"`
    
}

func (req *TaobaoFenxiaoOrdersGetRequest) GetAPIName() string {
	return "taobao.fenxiao.orders.get"
}

/* TaobaoFenxiaoOrdersGetResponse 分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。 */
type TaobaoFenxiaoOrdersGetResponse struct {
    
    /* purchase_orders Object Array采购单及子采购单信息。返回 PurchaseOrder 包含的字段信息。 */
    purchase_orders TopDpOrderDo `json:"purchase_orders";xml:"purchase_orders"`
    
    /* total_results Basic搜索到的采购单记录总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* AlibabaEinvoiceSerialnoGenerateRequest erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突 */
type AlibabaEinvoiceSerialnoGenerateRequest struct {
    
}

func (req *AlibabaEinvoiceSerialnoGenerateRequest) GetAPIName() string {
	return "alibaba.einvoice.serialno.generate"
}

/* AlibabaEinvoiceSerialnoGenerateResponse erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突 */
type AlibabaEinvoiceSerialnoGenerateResponse struct {
    
    /* serial_no Basicresult */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

/* AlibabaEinvoiceSerialnoBatchGenerateRequest 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。 */
type AlibabaEinvoiceSerialnoBatchGenerateRequest struct {
    
}

func (req *AlibabaEinvoiceSerialnoBatchGenerateRequest) GetAPIName() string {
	return "alibaba.einvoice.serialno.batch.generate"
}

/* AlibabaEinvoiceSerialnoBatchGenerateResponse 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。 */
type AlibabaEinvoiceSerialnoBatchGenerateResponse struct {
    
    /* serial_no_list Basic Arrayresult */
    serial_no_list string `json:"serial_no_list";xml:"serial_no_list"`
    
}

/* TaobaoSellercenterRolemembersGetRequest 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。 */
type TaobaoSellercenterRolemembersGetRequest struct {
    
    /* role_id required角色id */
    role_id int64 `json:"role_id";xml:"role_id"`
    
}

func (req *TaobaoSellercenterRolemembersGetRequest) GetAPIName() string {
	return "taobao.sellercenter.rolemembers.get"
}

/* TaobaoSellercenterRolemembersGetResponse 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。 */
type TaobaoSellercenterRolemembersGetResponse struct {
    
    /* subusers Object Array子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流 */
    subusers SubUserInfo `json:"subusers";xml:"subusers"`
    
}

/* TaobaoSellercenterSubuserPermissionsRolesGetRequest 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。 */
type TaobaoSellercenterSubuserPermissionsRolesGetRequest struct {
    
    /* nick required子账号昵称(子账号标识) */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetAPIName() string {
	return "taobao.sellercenter.subuser.permissions.roles.get"
}

/* TaobaoSellercenterSubuserPermissionsRolesGetResponse 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。 */
type TaobaoSellercenterSubuserPermissionsRolesGetResponse struct {
    
    /* subuser_permission Object子账号被所拥有的权限 */
    subuser_permission SubUserPermission `json:"subuser_permission";xml:"subuser_permission"`
    
}

/* TaobaoWangwangAbstractInitializeRequest 模糊查询服务初始化，只支持json返回 */
type TaobaoWangwangAbstractInitializeRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
}

func (req *TaobaoWangwangAbstractInitializeRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.initialize"
}

/* TaobaoWangwangAbstractInitializeResponse 模糊查询服务初始化，只支持json返回 */
type TaobaoWangwangAbstractInitializeResponse struct {
    
    /* error_msg Basic当ret_code=-1时这个变量才有 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1表示成功或失败 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
}

/* TaobaoWangwangAbstractGetwordlistRequest 获取关键词列表，只支持json返回 */
type TaobaoWangwangAbstractGetwordlistRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
}

func (req *TaobaoWangwangAbstractGetwordlistRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.getwordlist"
}

/* TaobaoWangwangAbstractGetwordlistResponse 获取关键词列表，只支持json返回 */
type TaobaoWangwangAbstractGetwordlistResponse struct {
    
    /* error_msg Basic例如单词长度太长等，ret_code=-1才有 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
    /* word_lists Object Array关键词列表，ret_code=0才有 */
    word_lists WordList `json:"word_lists";xml:"word_lists"`
    
}

/* TaobaoOpenuidGetBymixnickRequest 通过mixnick转换openuid */
type TaobaoOpenuidGetBymixnickRequest struct {
    
    /* mix_nick required无线类应用获取到的混淆的nick */
    mix_nick string `json:"mix_nick";xml:"mix_nick"`
    
}

func (req *TaobaoOpenuidGetBymixnickRequest) GetAPIName() string {
	return "taobao.openuid.get.bymixnick"
}

/* TaobaoOpenuidGetBymixnickResponse 通过mixnick转换openuid */
type TaobaoOpenuidGetBymixnickResponse struct {
    
    /* open_uid BasicOpenUID */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
}

/* TaobaoWangwangAbstractLogqueryRequest 模糊聊天记录查询 */
type TaobaoWangwangAbstractLogqueryRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
    /* count optional获取记录条数，默认值是1000 */
    count int64 `json:"count";xml:"count"`
    
    /* end_date requiredutc */
    end_date int64 `json:"end_date";xml:"end_date"`
    
    /* from_id required卖家id，有cntaobao前缀 */
    from_id string `json:"from_id";xml:"from_id"`
    
    /* next_key optional设置了这个值，那么聊天记录就从这个点开始查询 */
    next_key string `json:"next_key";xml:"next_key"`
    
    /* start_date requiredutc */
    start_date int64 `json:"start_date";xml:"start_date"`
    
    /* to_id required买家id，有cntaobao前缀 */
    to_id string `json:"to_id";xml:"to_id"`
    
}

func (req *TaobaoWangwangAbstractLogqueryRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.logquery"
}

/* TaobaoWangwangAbstractLogqueryResponse 模糊聊天记录查询 */
type TaobaoWangwangAbstractLogqueryResponse struct {
    
    /* error_msg Basic例如单词长度太长等。
当ret_code不为0时这个值存在。 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* from_id Basic卖家id */
    from_id string `json:"from_id";xml:"from_id"`
    
    /* is_sliced Basic0或1
168的时候用户名设置有误 */
    is_sliced int64 `json:"is_sliced";xml:"is_sliced"`
    
    /* next_key Basic当is_sliced为真时才有这项 */
    next_key string `json:"next_key";xml:"next_key"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息.
为-1时就只有error_msg字段是有效的。 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
    /* to_id Basic买家id */
    to_id string `json:"to_id";xml:"to_id"`
    
    /* url_lists Object Arrayurl列表 */
    url_lists UrlList `json:"url_lists";xml:"url_lists"`
    
    /* word_count_lists Object Array关键词统计列表 */
    word_count_lists WordCountList `json:"word_count_lists";xml:"word_count_lists"`
    
}

/* TaobaoWangwangAbstractAddwordRequest 增加关键词，只支持json返回 */
type TaobaoWangwangAbstractAddwordRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
    /* word required关键词，长度大于0 */
    word string `json:"word";xml:"word"`
    
}

func (req *TaobaoWangwangAbstractAddwordRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.addword"
}

/* TaobaoWangwangAbstractAddwordResponse 增加关键词，只支持json返回 */
type TaobaoWangwangAbstractAddwordResponse struct {
    
    /* error_msg Basic例如单词长度太长等，当ret_code=-1时才有这项 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
}

/* TaobaoOpenuidGetBytradeRequest 通过订单获取对应买家的openUID,需要卖家授权 */
type TaobaoOpenuidGetBytradeRequest struct {
    
    /* tid required订单ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoOpenuidGetBytradeRequest) GetAPIName() string {
	return "taobao.openuid.get.bytrade"
}

/* TaobaoOpenuidGetBytradeResponse 通过订单获取对应买家的openUID,需要卖家授权 */
type TaobaoOpenuidGetBytradeResponse struct {
    
    /* open_uid Basic当前交易tid对应买家的openuid */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
}

/* TaobaoWangwangAbstractDeletewordRequest 删除关键词，只支持json返回 */
type TaobaoWangwangAbstractDeletewordRequest struct {
    
    /* charset optional传入参数的字符集 */
    charset string `json:"charset";xml:"charset"`
    
    /* word required关键词，长度大于0 */
    word string `json:"word";xml:"word"`
    
}

func (req *TaobaoWangwangAbstractDeletewordRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.deleteword"
}

/* TaobaoWangwangAbstractDeletewordResponse 删除关键词，只支持json返回 */
type TaobaoWangwangAbstractDeletewordResponse struct {
    
    /* error_msg Basic例如单词长度太长等 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* ret_code Basic0或-1，表示错误或正确，错误时有错误信息 */
    ret_code int64 `json:"ret_code";xml:"ret_code"`
    
}

/* TaobaoOpenuidGetRequest 获取授权账号对应的OpenUid */
type TaobaoOpenuidGetRequest struct {
    
}

func (req *TaobaoOpenuidGetRequest) GetAPIName() string {
	return "taobao.openuid.get"
}

/* TaobaoOpenuidGetResponse 获取授权账号对应的OpenUid */
type TaobaoOpenuidGetResponse struct {
    
    /* open_uid BasicOpenUID */
    open_uid string `json:"open_uid";xml:"open_uid"`
    
}

/* TaobaoLogisticsOrderShengxianConfirmRequest 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
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

/* TaobaoLogisticsOrderShengxianConfirmResponse 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
type TaobaoLogisticsOrderShengxianConfirmResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* ship_fresh Object发货成功后，返回承运商的信息 */
    ship_fresh ShipFresh `json:"ship_fresh";xml:"ship_fresh"`
    
}

/* TaobaoLogisticsConsignTcConfirmRequest 下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。 */
type TaobaoLogisticsConsignTcConfirmRequest struct {
    
    /* app_name requiredERP的名称 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* extend_fields optional扩展字段 K:V; */
    extend_fields map[string]interface{} `json:"extend_fields";xml:"extend_fields"`
    
    /* goods_list optional发货的包裹 */
    goods_list ConfirmConsignGoodsDto `json:"goods_list";xml:"goods_list"`
    
    /* out_trade_no required已发货的外部订单号 */
    out_trade_no string `json:"out_trade_no";xml:"out_trade_no"`
    
    /* provider_id required货主id */
    provider_id int64 `json:"provider_id";xml:"provider_id"`
    
    /* seller_id required卖家id */
    seller_id int64 `json:"seller_id";xml:"seller_id"`
    
    /* trade_id required待发货的淘宝主交易号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoLogisticsConsignTcConfirmRequest) GetAPIName() string {
	return "taobao.logistics.consign.tc.confirm"
}

/* TaobaoLogisticsConsignTcConfirmResponse 下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。 */
type TaobaoLogisticsConsignTcConfirmResponse struct {
    
    /* order_consign_code Basic菜鸟发货单据 */
    order_consign_code string `json:"order_consign_code";xml:"order_consign_code"`
    
    /* retry Basic是否重试 */
    retry bool `json:"retry";xml:"retry"`
    
    /* trace_id Basic单次调用流程唯一id */
    trace_id string `json:"trace_id";xml:"trace_id"`
    
}

/* TaobaoLogisticsOrdersGetRequest 批量查询物流订单。 */
type TaobaoLogisticsOrdersGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_created optional创建时间结束 */
    end_created time.Time `json:"end_created";xml:"end_created"`
    
    /* fields required需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br>
tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。 */
    fields string `json:"fields";xml:"fields"`
    
    /* freight_payer optional谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
    freight_payer string `json:"freight_payer";xml:"freight_payer"`
    
    /* page_no optional页码.该字段没传 或 值<1 ,则默认page_no为1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数.该字段没传 或 值<1 ,则默认page_size为40 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* receiver_name optional收货人姓名 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* seller_confirm optional卖家是否发货.可选值:yes(是),no(否).如:yes */
    seller_confirm string `json:"seller_confirm";xml:"seller_confirm"`
    
    /* start_created optional创建时间开始 */
    start_created time.Time `json:"start_created";xml:"start_created"`
    
    /* status optional物流状态.查看数据结构 Shipping 中的status字段. */
    status string `json:"status";xml:"status"`
    
    /* tid optional交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* _type optional物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoLogisticsOrdersGetRequest) GetAPIName() string {
	return "taobao.logistics.orders.get"
}

/* TaobaoLogisticsOrdersGetResponse 批量查询物流订单。 */
type TaobaoLogisticsOrdersGetResponse struct {
    
    /* shippings Object Array获取的物流订单详情列表 
返回的Shipping包含的具体信息为入参fields请求的字段信息 */
    shippings Shipping `json:"shippings";xml:"shippings"`
    
    /* total_results Basic搜索到的物流订单列表总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoLogisticsOrdersDetailGetRequest 查询物流订单的详细信息，涉及用户隐私字段。 */
type TaobaoLogisticsOrdersDetailGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_created optional创建时间结束.格式:yyyy-MM-dd HH:mm:ss */
    end_created time.Time `json:"end_created";xml:"end_created"`
    
    /* fields required需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔. */
    fields string `json:"fields";xml:"fields"`
    
    /* freight_payer optional谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
    freight_payer string `json:"freight_payer";xml:"freight_payer"`
    
    /* page_no optional页码.该字段没传 或 值<1 ,则默认page_no为1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数.该字段没传 或 值<1 ，则默认page_size为40 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* receiver_name optional收货人姓名 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* seller_confirm optional卖家是否发货.可选值:yes(是),no(否).如:yes. */
    seller_confirm string `json:"seller_confirm";xml:"seller_confirm"`
    
    /* start_created optional创建时间开始.格式:yyyy-MM-dd HH:mm:ss */
    start_created time.Time `json:"start_created";xml:"start_created"`
    
    /* status optional物流状态.可查看数据结构 Shipping 中的status字段. */
    status string `json:"status";xml:"status"`
    
    /* tid optional交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息. */
    tid int64 `json:"tid";xml:"tid"`
    
    /* _type optional物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoLogisticsOrdersDetailGetRequest) GetAPIName() string {
	return "taobao.logistics.orders.detail.get"
}

/* TaobaoLogisticsOrdersDetailGetResponse 查询物流订单的详细信息，涉及用户隐私字段。 */
type TaobaoLogisticsOrdersDetailGetResponse struct {
    
    /* shippings Object Array获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息. */
    shippings Shipping `json:"shippings";xml:"shippings"`
    
    /* total_results Basic搜索到的物流订单列表总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoLogisticsCompaniesGetRequest 查询淘宝网合作的物流公司信息，用于发货接口。 */
type TaobaoLogisticsCompaniesGetRequest struct {
    
    /* fields required需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开.
如:id,code,name,reg_mail_no
<br><font color='red'>说明：</font>
<br>id：物流公司ID
<br>code：物流公司code
<br>name：物流公司名称
<br>reg_mail_no：物流公司对应的运单规则 */
    fields string `json:"fields";xml:"fields"`
    
    /* is_recommended optional是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司. */
    is_recommended bool `json:"is_recommended";xml:"is_recommended"`
    
    /* order_mode optional推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司. */
    order_mode string `json:"order_mode";xml:"order_mode"`
    
}

func (req *TaobaoLogisticsCompaniesGetRequest) GetAPIName() string {
	return "taobao.logistics.companies.get"
}

/* TaobaoLogisticsCompaniesGetResponse 查询淘宝网合作的物流公司信息，用于发货接口。 */
type TaobaoLogisticsCompaniesGetResponse struct {
    
    /* logistics_companies Object Array物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。 */
    logistics_companies LogisticsCompany `json:"logistics_companies";xml:"logistics_companies"`
    
}

/* TaobaoAppipGetRequest 获取ISV发起请求服务器IP */
type TaobaoAppipGetRequest struct {
    
}

func (req *TaobaoAppipGetRequest) GetAPIName() string {
	return "taobao.appip.get"
}

/* TaobaoAppipGetResponse 获取ISV发起请求服务器IP */
type TaobaoAppipGetResponse struct {
    
    /* ip BasicISV发起请求服务器IP */
    ip string `json:"ip";xml:"ip"`
    
}

/* QimenTaobaoQimenEntryorderCallbackRequest ERP通过该接口回传预约入库单对应的商家仓出库单状态 */
type QimenTaobaoQimenEntryorderCallbackRequest struct {
    
    /* appkey required应用在奇门申请的appkey */
    appkey string `json:"appkey";xml:"appkey"`
    
    /* entryorderlist required入库单信息列表 */
    entryorderlist Struct `json:"entryorderlist";xml:"entryorderlist"`
    
    /* userId required货主在奇门授权的ID */
    userId string `json:"userId";xml:"userId"`
    
}

func (req *QimenTaobaoQimenEntryorderCallbackRequest) GetAPIName() string {
	return "qimen.taobao.qimen.entryorder.callback"
}

/* QimenTaobaoQimenEntryorderCallbackResponse ERP通过该接口回传预约入库单对应的商家仓出库单状态 */
type QimenTaobaoQimenEntryorderCallbackResponse struct {
    
    /* code Basic错误码:CD001 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic回传结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic异常信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoFenxiaoLoginUserGetRequest 获取用户登录信息 */
type TaobaoFenxiaoLoginUserGetRequest struct {
    
}

func (req *TaobaoFenxiaoLoginUserGetRequest) GetAPIName() string {
	return "taobao.fenxiao.login.user.get"
}

/* TaobaoFenxiaoLoginUserGetResponse 获取用户登录信息 */
type TaobaoFenxiaoLoginUserGetResponse struct {
    
    /* login_user Object登录用户信息 */
    login_user LoginUser `json:"login_user";xml:"login_user"`
    
}

/* TaobaoRdcAligeniusOrdermsgUpdateRequest 用于订单消息处理状态回传 */
type TaobaoRdcAligeniusOrdermsgUpdateRequest struct {
    
    /* oid required子订单（消息中传的子订单） */
    oid int64 `json:"oid";xml:"oid"`
    
    /* status required处理状态，1=成功，2=处理失败 */
    status int64 `json:"status";xml:"status"`
    
    /* tid required主订单（消息中传的主订单） */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoRdcAligeniusOrdermsgUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.ordermsg.update"
}

/* TaobaoRdcAligeniusOrdermsgUpdateResponse 用于订单消息处理状态回传 */
type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TmallItemQuantityUpdateRequest 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。 */
type TmallItemQuantityUpdateRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* item_quantity optional商品库存数；增量编辑方式支持正数、负数 */
    item_quantity int64 `json:"item_quantity";xml:"item_quantity"`
    
    /* options optional商品库存更新时候的可选参数 */
    options UpdateItemQuantityOption `json:"options";xml:"options"`
    
    /* sku_quantities optional更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！ */
    sku_quantities UpdateSkuQuantity `json:"sku_quantities";xml:"sku_quantities"`
    
}

func (req *TmallItemQuantityUpdateRequest) GetAPIName() string {
	return "tmall.item.quantity.update"
}

/* TmallItemQuantityUpdateResponse 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。 */
type TmallItemQuantityUpdateResponse struct {
    
    /* quantity_update_result Basic库存更新结果，商品id */
    quantity_update_result string `json:"quantity_update_result";xml:"quantity_update_result"`
    
}

/* TaobaoPictureCategoryAddRequest 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符 */
type TaobaoPictureCategoryAddRequest struct {
    
    /* parent_id optional图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
    /* picture_category_name required图片分类名称，最大长度20字符，中文字符算2个字符，不能为空 */
    picture_category_name string `json:"picture_category_name";xml:"picture_category_name"`
    
}

func (req *TaobaoPictureCategoryAddRequest) GetAPIName() string {
	return "taobao.picture.category.add"
}

/* TaobaoPictureCategoryAddResponse 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符 */
type TaobaoPictureCategoryAddResponse struct {
    
    /* picture_category Object图片分类信息 */
    picture_category PictureCategory `json:"picture_category";xml:"picture_category"`
    
}

/* TaobaoRegionPriceCancleRequest 取消区域价格 */
type TaobaoRegionPriceCancleRequest struct {
    
    /* item_id required商品 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sku_id required无sku传0 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceCancleRequest) GetAPIName() string {
	return "taobao.region.price.cancle"
}

/* TaobaoRegionPriceCancleResponse 取消区域价格 */
type TaobaoRegionPriceCancleResponse struct {
    
    /* result Objectresult */
    result BaseResult `json:"result";xml:"result"`
    
}

/* TaobaoQimenItemlackReportRequest WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP */
type TaobaoQimenItemlackReportRequest struct {
    
    /* createTime required缺货回告创建时间(YYYY-MM-DD HH:mm:ss) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderCode requiredERP的发货单编码 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional仓库系统的发货单编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional缺货商品列表 */
    items Item `json:"items";xml:"items"`
    
    /* outBizCode required外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理) */
    outBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode required仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemlackReportRequest) GetAPIName() string {
	return "taobao.qimen.itemlack.report"
}

/* TaobaoQimenItemlackReportResponse WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP */
type TaobaoQimenItemlackReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenOrderprocessQueryRequest ERP调用订单流水查询接口 */
type TaobaoQimenOrderprocessQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required单据号 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统单据号 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderSourceCode optional交易单号 */
    orderSourceCode string `json:"orderSourceCode";xml:"orderSourceCode"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderprocessQueryRequest) GetAPIName() string {
	return "taobao.qimen.orderprocess.query"
}

/* TaobaoQimenOrderprocessQueryResponse ERP调用订单流水查询接口 */
type TaobaoQimenOrderprocessQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderProcess Object订单处理流程 */
    orderProcess OrderProcess `json:"orderProcess";xml:"orderProcess"`
    
}

/* TaobaoFenxiaoRefundQueryRequest 供应商按查询条件批量查询代销采购退款 */
type TaobaoFenxiaoRefundQueryRequest struct {
    
    /* end_date required代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* page_no optional页码（大于0的整数。无值或小于1的值按默认值1计） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* query_seller_refund optional是否查询下游买家的退款信息 */
    query_seller_refund bool `json:"query_seller_refund";xml:"query_seller_refund"`
    
    /* start_date required代销采购退款单最早修改时间 */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
}

func (req *TaobaoFenxiaoRefundQueryRequest) GetAPIName() string {
	return "taobao.fenxiao.refund.query"
}

/* TaobaoFenxiaoRefundQueryResponse 供应商按查询条件批量查询代销采购退款 */
type TaobaoFenxiaoRefundQueryResponse struct {
    
    /* refund_list Object Array代销采购退款列表 */
    refund_list RefundDetail `json:"refund_list";xml:"refund_list"`
    
    /* total_results Basic按查询条件查到的记录总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoRefundRefusereasonGetRequest 获取商家拒绝原因列表 */
type TaobaoRefundRefusereasonGetRequest struct {
    
    /* fields required返回参数 */
    fields string `json:"fields";xml:"fields"`
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase required售中或售后 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRefundRefusereasonGetRequest) GetAPIName() string {
	return "taobao.refund.refusereason.get"
}

/* TaobaoRefundRefusereasonGetResponse 获取商家拒绝原因列表 */
type TaobaoRefundRefusereasonGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* reasons Object Array卖家拒绝原因对象 */
    reasons Reason `json:"reasons";xml:"reasons"`
    
    /* total_results Basic原因个数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoQimenReturnorderQueryRequest ERP调用奇门的接口，查询退货入库信息 */
type TaobaoQimenReturnorderQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* returnOrderCode required退货单编码 */
    returnOrderCode string `json:"returnOrderCode";xml:"returnOrderCode"`
    
    /* returnOrderId optional仓库系统订单编码 */
    returnOrderId string `json:"returnOrderId";xml:"returnOrderId"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenReturnorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.query"
}

/* TaobaoQimenReturnorderQueryResponse ERP调用奇门的接口，查询退货入库信息 */
type TaobaoQimenReturnorderQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* returnOrder Object退货单信息 */
    returnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

/* TaobaoQimenOrderstatusBatchqueryRequest ERP调用奇门的接口,查询订单在仓库的状态 */
type TaobaoQimenOrderstatusBatchqueryRequest struct {
    
    /* currentPage required当前第几页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* endTime optional订单最后操作时间(查询截止时间点;格式:YYYY-MM-DD hh:mm:ss) */
    endTime string `json:"endTime";xml:"endTime"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK= 换货入库;CNJG= 仓内加工单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* pageSize required页面大小(建议不超过100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* startTime required订单最后操作时间(查询起始时间点;格式:YYYY-MM-DD hh:mm:ss) */
    startTime string `json:"startTime";xml:"startTime"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderstatusBatchqueryRequest) GetAPIName() string {
	return "taobao.qimen.orderstatus.batchquery"
}

/* TaobaoQimenOrderstatusBatchqueryResponse ERP调用奇门的接口,查询订单在仓库的状态 */
type TaobaoQimenOrderstatusBatchqueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orders Object Array单据信息 */
    orders Order `json:"orders";xml:"orders"`
    
    /* totalPage Basic总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

/* TaobaoQimenItemlackQueryRequest ERP调用奇门的接口,查询库存商品缺货情况 */
type TaobaoQimenItemlackQueryRequest struct {
    
    /* deliveryOrderCode required出库单号 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional仓储系统出库单号 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemlackQueryRequest) GetAPIName() string {
	return "taobao.qimen.itemlack.query"
}

/* TaobaoQimenItemlackQueryResponse ERP调用奇门的接口,查询库存商品缺货情况 */
type TaobaoQimenItemlackQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* createTime Basic缺货回告创建时间(YYYY-MM-DD HH:mm:ss) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderCode BasicERP的发货单编码 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId Basic仓库系统的发货单编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array缺货商品列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
    /* warehouseCode Basic仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

/* TaobaoQimenOrderPendingRequest ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等 */
type TaobaoQimenOrderPendingRequest struct {
    
    /* actionType required操作类型(pending=挂起;restore=恢复) */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required单据编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统单据编码 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* reason optional挂起/恢复原因 */
    reason string `json:"reason";xml:"reason"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderPendingRequest) GetAPIName() string {
	return "taobao.qimen.order.pending"
}

/* TaobaoQimenOrderPendingResponse ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等 */
type TaobaoQimenOrderPendingResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenDeliveryorderBatchcreateAnswerRequest WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回） */
type TaobaoQimenDeliveryorderBatchcreateAnswerRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional发货单列表 */
    orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchcreate.answer"
}

/* TaobaoQimenDeliveryorderBatchcreateAnswerResponse WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回） */
type TaobaoQimenDeliveryorderBatchcreateAnswerResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenCombineitemSynchronizeRequest ERP调用奇门的接口,将商品信息同步给WMS */
type TaobaoQimenCombineitemSynchronizeRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemCode required组合商品的ERP编码 */
    itemCode string `json:"itemCode";xml:"itemCode"`
    
    /* itemId optionaltemp */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* items optional组合商品接口中的单商品信息 */
    items Item `json:"items";xml:"items"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenCombineitemSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.synchronize"
}

/* TaobaoQimenCombineitemSynchronizeResponse ERP调用奇门的接口,将商品信息同步给WMS */
type TaobaoQimenCombineitemSynchronizeResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoSellercenterSubusersGetRequest 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号) */
type TaobaoSellercenterSubusersGetRequest struct {
    
    /* nick required表示卖家昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterSubusersGetRequest) GetAPIName() string {
	return "taobao.sellercenter.subusers.get"
}

/* TaobaoSellercenterSubusersGetResponse 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号) */
type TaobaoSellercenterSubusersGetResponse struct {
    
    /* subusers Object Array子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流 */
    subusers SubUserInfo `json:"subusers";xml:"subusers"`
    
}

/* TaobaoQimenStockQueryRequest ERP调用奇门的接口,查询商品的库存量 */
type TaobaoQimenStockQueryRequest struct {
    
    /* batchCode optional批次编码 */
    batchCode string `json:"batchCode";xml:"batchCode"`
    
    /* expireDate optional商品过期日期(YYYY-MM-DD) */
    expireDate string `json:"expireDate";xml:"expireDate"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* inventoryType optional库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存) */
    inventoryType string `json:"inventoryType";xml:"inventoryType"`
    
    /* itemCode required商品编码 */
    itemCode string `json:"itemCode";xml:"itemCode"`
    
    /* itemId optional仓储系统商品ID */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* productDate optional商品生产日期(YYYY-MM-DD) */
    productDate string `json:"productDate";xml:"productDate"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockQueryRequest) GetAPIName() string {
	return "taobao.qimen.stock.query"
}

/* TaobaoQimenStockQueryResponse ERP调用奇门的接口,查询商品的库存量 */
type TaobaoQimenStockQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品的库存信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* totalCount Basic总数 */
    totalCount int64 `json:"totalCount";xml:"totalCount"`
    
}

/* TaobaoSellercenterUserPermissionsGetRequest 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号) */
type TaobaoSellercenterUserPermissionsGetRequest struct {
    
    /* nick required用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterUserPermissionsGetRequest) GetAPIName() string {
	return "taobao.sellercenter.user.permissions.get"
}

/* TaobaoSellercenterUserPermissionsGetResponse 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号) */
type TaobaoSellercenterUserPermissionsGetResponse struct {
    
    /* permissions Object Array权限列表 */
    permissions Permission `json:"permissions";xml:"permissions"`
    
}

/* TaobaoQimenStoreprocessConfirmRequest WMS调用奇门的接口,回传仓内加工单创建情况 */
type TaobaoQimenStoreprocessConfirmRequest struct {
    
    /* actualQty optional实际作业总数量 */
    actualQty int64 `json:"actualQty";xml:"actualQty"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* materialitems optional加工商品列表 */
    materialitems MaterialItem `json:"materialitems";xml:"materialitems"`
    
    /* orderCompleteTime required加工单完成时间(YYYY-MM-DD HH:MM:SS) */
    orderCompleteTime string `json:"orderCompleteTime";xml:"orderCompleteTime"`
    
    /* orderType required单据类型(CNJG=仓内加工作业单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* outBizCode required外部业务编码(一个合作伙伴中要求唯一多次确认时;每次传入要求唯一(一般传入WMS损益单据编码)) */
    outBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* processOrderCode required加工单编码 */
    processOrderCode string `json:"processOrderCode";xml:"processOrderCode"`
    
    /* processOrderId optional仓储系统加工单ID */
    processOrderId string `json:"processOrderId";xml:"processOrderId"`
    
    /* productitems optional加工商品列表 */
    productitems ProductItem `json:"productitems";xml:"productitems"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenStoreprocessConfirmRequest) GetAPIName() string {
	return "taobao.qimen.storeprocess.confirm"
}

/* TaobaoQimenStoreprocessConfirmResponse WMS调用奇门的接口,回传仓内加工单创建情况 */
type TaobaoQimenStoreprocessConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoTradeOrderskuUpdateRequest 只能更新发货前子订单的销售属性 
只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 
必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先 */
type TaobaoTradeOrderskuUpdateRequest struct {
    
    /* oid required子订单编号（对于单笔订单的交易可以传交易编号）。 */
    oid int64 `json:"oid";xml:"oid"`
    
    /* sku_id special销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* sku_props special销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。 */
    sku_props string `json:"sku_props";xml:"sku_props"`
    
}

func (req *TaobaoTradeOrderskuUpdateRequest) GetAPIName() string {
	return "taobao.trade.ordersku.update"
}

/* TaobaoTradeOrderskuUpdateResponse 只能更新发货前子订单的销售属性 
只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 
必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先 */
type TaobaoTradeOrderskuUpdateResponse struct {
    
    /* order Object只返回oid和modified */
    order Order `json:"order";xml:"order"`
    
}

/* TaobaoQimenStoreprocessCreateRequest ERP调用奇门的接口,创建仓内加工单 */
type TaobaoQimenStoreprocessCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* materialitems optional加工商品列表 */
    materialitems MaterialItem `json:"materialitems";xml:"materialitems"`
    
    /* orderCreateTime required加工单创建时间(YYYY-MM-DD HH:MM:SS) */
    orderCreateTime string `json:"orderCreateTime";xml:"orderCreateTime"`
    
    /* orderType required单据类型(CNJG=仓内加工作业单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* planQty optional成品计划数量 */
    planQty int64 `json:"planQty";xml:"planQty"`
    
    /* planTime required计划加工时间(YYYY-MM-DD HH:MM:SS) */
    planTime string `json:"planTime";xml:"planTime"`
    
    /* processOrderCode required加工单编码 */
    processOrderCode string `json:"processOrderCode";xml:"processOrderCode"`
    
    /* productitems optional商品列表 */
    productitems ProductItem `json:"productitems";xml:"productitems"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* serviceType required加工类型(1:仓内组合加工 2:仓内组合拆分) */
    serviceType string `json:"serviceType";xml:"serviceType"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStoreprocessCreateRequest) GetAPIName() string {
	return "taobao.qimen.storeprocess.create"
}

/* TaobaoQimenStoreprocessCreateResponse ERP调用奇门的接口,创建仓内加工单 */
type TaobaoQimenStoreprocessCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* processOrderId Basic仓储系统处理单ID */
    processOrderId string `json:"processOrderId";xml:"processOrderId"`
    
}

/* TaobaoTradeShippingaddressUpdateRequest 只能更新一笔交易里面的买家收货地址 
只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 
更新后的发货地址可以通过taobao.trade.fullinfo.get查到 
参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节） */
type TaobaoTradeShippingaddressUpdateRequest struct {
    
    /* receiver_address optional收货地址。最大长度为228个字节。 */
    receiver_address string `json:"receiver_address";xml:"receiver_address"`
    
    /* receiver_city optional城市。最大长度为32个字节。如：杭州 */
    receiver_city string `json:"receiver_city";xml:"receiver_city"`
    
    /* receiver_district optional区/县。最大长度为32个字节。如：西湖区 */
    receiver_district string `json:"receiver_district";xml:"receiver_district"`
    
    /* receiver_mobile optional移动电话。最大长度为11个字节。 */
    receiver_mobile string `json:"receiver_mobile";xml:"receiver_mobile"`
    
    /* receiver_name optional收货人全名。最大长度为50个字节。 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* receiver_phone optional固定电话。最大长度为30个字节。 */
    receiver_phone string `json:"receiver_phone";xml:"receiver_phone"`
    
    /* receiver_state optional省份。最大长度为32个字节。如：浙江 */
    receiver_state string `json:"receiver_state";xml:"receiver_state"`
    
    /* receiver_zip optional邮政编码。必须由6个数字组成。 */
    receiver_zip string `json:"receiver_zip";xml:"receiver_zip"`
    
    /* tid required交易编号。 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeShippingaddressUpdateRequest) GetAPIName() string {
	return "taobao.trade.shippingaddress.update"
}

/* TaobaoTradeShippingaddressUpdateResponse 只能更新一笔交易里面的买家收货地址 
只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 
更新后的发货地址可以通过taobao.trade.fullinfo.get查到 
参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节） */
type TaobaoTradeShippingaddressUpdateResponse struct {
    
    /* trade Object交易结构 */
    trade Trade `json:"trade";xml:"trade"`
    
}

/* TaobaoQimenStockoutQueryRequest 查询出库单详情 */
type TaobaoQimenStockoutQueryRequest struct {
    
    /* deliveryOrderCode required出库单号 */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional仓储系统出库单号 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页 */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockoutQueryRequest) GetAPIName() string {
	return "taobao.qimen.stockout.query"
}

/* TaobaoQimenStockoutQueryResponse 查询出库单详情 */
type TaobaoQimenStockoutQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* deliveryOrder Object出库单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array单据信息列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages Object Array包裹信息 */
    packages Package `json:"packages";xml:"packages"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
}

/* TaobaoInventoryInitialItemRequest 商家仓商品初始化在各个仓中库存 */
type TaobaoInventoryInitialItemRequest struct {
    
    /* sc_item_id required后端商品id */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* store_inventorys required商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}] */
    store_inventorys string `json:"store_inventorys";xml:"store_inventorys"`
    
}

func (req *TaobaoInventoryInitialItemRequest) GetAPIName() string {
	return "taobao.inventory.initial.item"
}

/* TaobaoInventoryInitialItemResponse 商家仓商品初始化在各个仓中库存 */
type TaobaoInventoryInitialItemResponse struct {
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

/* TaobaoPicturePicturesCountRequest 图片总数查询 */
type TaobaoPicturePicturesCountRequest struct {
    
    /* client_type optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* deleted optional是否删除,undeleted代表没有删除,deleted表示删除 */
    deleted string `json:"deleted";xml:"deleted"`
    
    /* end_date optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* picture_category_id optional图片分类 */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional图片ID */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* start_date optional查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* start_modified_date optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
    start_modified_date time.Time `json:"start_modified_date";xml:"start_modified_date"`
    
    /* title optional文件名 */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoPicturePicturesCountRequest) GetAPIName() string {
	return "taobao.picture.pictures.count"
}

/* TaobaoPicturePicturesCountResponse 图片总数查询 */
type TaobaoPicturePicturesCountResponse struct {
    
    /* totals Basic查询的文件总数 */
    totals int64 `json:"totals";xml:"totals"`
    
}

/* TaobaoJushitaJdpUserDeleteRequest 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。 */
type TaobaoJushitaJdpUserDeleteRequest struct {
    
    /* nick special要删除用户的昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* user_id special需要删除的用户编号 */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoJushitaJdpUserDeleteRequest) GetAPIName() string {
	return "taobao.jushita.jdp.user.delete"
}

/* TaobaoJushitaJdpUserDeleteResponse 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。 */
type TaobaoJushitaJdpUserDeleteResponse struct {
    
    /* is_success Basic是否删除成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoPicturePicturesGetRequest 图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd */
type TaobaoPicturePicturesGetRequest struct {
    
    /* client_type optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 */
    client_type string `json:"client_type";xml:"client_type"`
    
    /* current_page optional页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1 */
    current_page int64 `json:"current_page";xml:"current_page"`
    
    /* deleted optional是否删除,deleted代表没有删除 */
    deleted string `json:"deleted";xml:"deleted"`
    
    /* end_date optional结束时间 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* is_https optional是否获取https的链接 */
    is_https bool `json:"is_https";xml:"is_https"`
    
    /* order_by optional图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_size optional每页条数.每页返回最多返回100条,默认值40 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* picture_category_id optional图片分类 */
    picture_category_id int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    /* picture_id optional图片ID */
    picture_id int64 `json:"picture_id";xml:"picture_id"`
    
    /* start_date optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* start_modified_date optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
    start_modified_date time.Time `json:"start_modified_date";xml:"start_modified_date"`
    
    /* title optional图片标题,最大长度50字符,中英文都算一字符 */
    title string `json:"title";xml:"title"`
    
    /* urls optional图片url查询接口 */
    urls string `json:"urls";xml:"urls"`
    
}

func (req *TaobaoPicturePicturesGetRequest) GetAPIName() string {
	return "taobao.picture.pictures.get"
}

/* TaobaoPicturePicturesGetResponse 图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd */
type TaobaoPicturePicturesGetResponse struct {
    
    /* pictures Object Array图片空间图片数据对象 */
    pictures Picture `json:"pictures";xml:"pictures"`
    
}

/* TaobaoJushitaJdpUserAddRequest 提供给接入数据推送的应用添加数据推送服务的用户 */
type TaobaoJushitaJdpUserAddRequest struct {
    
    /* history_days optional推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。 */
    history_days int64 `json:"history_days";xml:"history_days"`
    
    /* host_ip optional已废弃，新版订单同步服务不要使用。同步用户数据的机器IP,必须是界面配置的IP。 */
    host_ip string `json:"host_ip";xml:"host_ip"`
    
    /* host_name optional已废弃，新版订单同步服务不要使用。绑定类型，用户数据同步的机器名称。 */
    host_name string `json:"host_name";xml:"host_name"`
    
    /* rds_name requiredRDS实例名称 */
    rds_name string `json:"rds_name";xml:"rds_name"`
    
    /* topics optional已废弃，使用页面中应用的配置。用户同步的数据类型,多个用','号分割。可选值：trade,refund,item */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoJushitaJdpUserAddRequest) GetAPIName() string {
	return "taobao.jushita.jdp.user.add"
}

/* TaobaoJushitaJdpUserAddResponse 提供给接入数据推送的应用添加数据推送服务的用户 */
type TaobaoJushitaJdpUserAddResponse struct {
    
    /* is_success Basic是否添加成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* CainiaoCloudprintTemplatesMigrateRequest 云打印模板迁移接口 */
type CainiaoCloudprintTemplatesMigrateRequest struct {
    
    /* custom_area_content optional自定义区内容 */
    custom_area_content string `json:"custom_area_content";xml:"custom_area_content"`
    
    /* custom_area_name optional自定义区名称 */
    custom_area_name string `json:"custom_area_name";xml:"custom_area_name"`
    
    /* tempalte_id optional标准电子面单模板的id */
    tempalte_id int64 `json:"tempalte_id";xml:"tempalte_id"`
    
}

func (req *CainiaoCloudprintTemplatesMigrateRequest) GetAPIName() string {
	return "cainiao.cloudprint.templates.migrate"
}

/* CainiaoCloudprintTemplatesMigrateResponse 云打印模板迁移接口 */
type CainiaoCloudprintTemplatesMigrateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* AlipayXiaodaiUserPermitRequest 阿里金融为用户开通消息通道接口 */
type AlipayXiaodaiUserPermitRequest struct {
    
    /* user_id required用户数字ID */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *AlipayXiaodaiUserPermitRequest) GetAPIName() string {
	return "alipay.xiaodai.user.permit"
}

/* AlipayXiaodaiUserPermitResponse 阿里金融为用户开通消息通道接口 */
type AlipayXiaodaiUserPermitResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoAppstoreSubscribeGetRequest 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get) */
type TaobaoAppstoreSubscribeGetRequest struct {
    
    /* lease_id optional插件实例ID */
    lease_id int64 `json:"lease_id";xml:"lease_id"`
    
    /* nick required用户昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoAppstoreSubscribeGetRequest) GetAPIName() string {
	return "taobao.appstore.subscribe.get"
}

/* TaobaoAppstoreSubscribeGetResponse 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get) */
type TaobaoAppstoreSubscribeGetResponse struct {
    
    /* user_subscribe Object用户订购信息 */
    user_subscribe UserSubscribe `json:"user_subscribe";xml:"user_subscribe"`
    
}

/* TaobaoQimenReplenishplanCreateRequest ERP调用奇门的接口,通知WMS创建补货计划 */
type TaobaoQimenReplenishplanCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* gmtDeadTime required最晚入库时间(YYYY-MM-DD HH:MM:SS) */
    gmtDeadTime string `json:"gmtDeadTime";xml:"gmtDeadTime"`
    
    /* items optional单据信息 */
    items ReplenishplanCreateItem `json:"items";xml:"items"`
    
    /* orderCode required外部系统单号(ERP单号) */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* userId required商家ID */
    userId string `json:"userId";xml:"userId"`
    
    /* warehouseCode required仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenReplenishplanCreateRequest) GetAPIName() string {
	return "taobao.qimen.replenishplan.create"
}

/* TaobaoQimenReplenishplanCreateResponse ERP调用奇门的接口,通知WMS创建补货计划 */
type TaobaoQimenReplenishplanCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TmallDisputeReceiveGetRequest 展示商家所有退款信息 */
type TmallDisputeReceiveGetRequest struct {
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* end_modified optional查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* fields required需要返回的字段。目前支持有：refund_id, alipay_no, tid, buyer_nick, seller_nick, status, created, modified, order_status, refund_fee, good_status, show_return_logistic(展现买家退货的物流信息), show_exchange_logistic(展现换货的物流信息), time_out, oid, refund_version, title, num, dispute_request, reason, desc */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数; 默认值:1 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数; 默认值:20;最大值:100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* refund_id optional逆向纠纷单号id */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* start_modified optional查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
    /* status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);CLOSED(退款关闭); SUCCESS(退款成功);SELLER_REFUSE_BUYER(卖家拒绝退款);WAIT_BUYER_CONFIRM_REDO_SEND_GOODS(等待买家确认重新邮寄的货物);WAIT_SELLER_CONFIRM_RETURN_ADDRESS(等待卖家确认退货地址);WAIT_SELLER_CONSIGN_GOOGDS(卖家确认收货,等待卖家发货);EXCHANGE_TRANSFORM_TO_REFUND(换货关闭,转退货退款);EXCHANGE_WAIT_BUYER_CONFIRM_GOODS(卖家已发货,等待买家确认收货);POST_FEE_DISPUTE_WAIT_ACTIVATE(邮费单已创建,待激活) */
    status string `json:"status";xml:"status"`
    
    /* _type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，查看可选值 */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TmallDisputeReceiveGetRequest) GetAPIName() string {
	return "tmall.dispute.receive.get"
}

/* TmallDisputeReceiveGetResponse 展示商家所有退款信息 */
type TmallDisputeReceiveGetResponse struct {
    
    /* result Objectresult */
    result ResultSet `json:"result";xml:"result"`
    
}

/* TaobaoLogisticsPartnersGetRequest 查询物流公司信息（可以查询目的地可不可达情况） */
type TaobaoLogisticsPartnersGetRequest struct {
    
    /* goods_value optional货物价格.只有当选择货到付款此参数才会有效 */
    goods_value string `json:"goods_value";xml:"goods_value"`
    
    /* is_need_carriage optional是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。 */
    is_need_carriage bool `json:"is_need_carriage";xml:"is_need_carriage"`
    
    /* service_type optional服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。 */
    service_type string `json:"service_type";xml:"service_type"`
    
    /* source_id optional物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取 */
    source_id string `json:"source_id";xml:"source_id"`
    
    /* target_id optional物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取 */
    target_id string `json:"target_id";xml:"target_id"`
    
}

func (req *TaobaoLogisticsPartnersGetRequest) GetAPIName() string {
	return "taobao.logistics.partners.get"
}

/* TaobaoLogisticsPartnersGetResponse 查询物流公司信息（可以查询目的地可不可达情况） */
type TaobaoLogisticsPartnersGetResponse struct {
    
    /* logistics_partners Object Array查询揽送范围之内的物流公司信息 */
    logistics_partners LogisticsPartner `json:"logistics_partners";xml:"logistics_partners"`
    
}

/* TaobaoTopSdkFeedbackUploadRequest sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性 */
type TaobaoTopSdkFeedbackUploadRequest struct {
    
    /* content optional具体内容，json形式 */
    content string `json:"content";xml:"content"`
    
    /* _type required1、回传加密信息 */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoTopSdkFeedbackUploadRequest) GetAPIName() string {
	return "taobao.top.sdk.feedback.upload"
}

/* TaobaoTopSdkFeedbackUploadResponse sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性 */
type TaobaoTopSdkFeedbackUploadResponse struct {
    
    /* upload_interval Basic控制回传间隔（单位：秒） */
    upload_interval int64 `json:"upload_interval";xml:"upload_interval"`
    
}

/* TaobaoSubuserDutysGetRequest 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息） */
type TaobaoSubuserDutysGetRequest struct {
    
    /* user_nick required主账号用户名 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubuserDutysGetRequest) GetAPIName() string {
	return "taobao.subuser.dutys.get"
}

/* TaobaoSubuserDutysGetResponse 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息） */
type TaobaoSubuserDutysGetResponse struct {
    
    /* dutys Object Array职务信息 */
    dutys Duty `json:"dutys";xml:"dutys"`
    
}

/* TaobaoSubuserEmployeeAddRequest 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息） */
type TaobaoSubuserEmployeeAddRequest struct {
    
    /* department_id required当前员工所属部门ID */
    department_id int64 `json:"department_id";xml:"department_id"`
    
    /* duty_id optional当前员工担任职务ID */
    duty_id int64 `json:"duty_id";xml:"duty_id"`
    
    /* employee_name required员工姓名 */
    employee_name string `json:"employee_name";xml:"employee_name"`
    
    /* employee_nickname optional员工花名 */
    employee_nickname string `json:"employee_nickname";xml:"employee_nickname"`
    
    /* employee_num optional员工工号(可以用大小写英文字母和数字) */
    employee_num string `json:"employee_num";xml:"employee_num"`
    
    /* entry_date optional员工入职时间 */
    entry_date time.Time `json:"entry_date";xml:"entry_date"`
    
    /* id_card_num optional员工身份证号码 */
    id_card_num string `json:"id_card_num";xml:"id_card_num"`
    
    /* leader_id optional直接上级的员工ID */
    leader_id int64 `json:"leader_id";xml:"leader_id"`
    
    /* office_phone optional办公电话 */
    office_phone string `json:"office_phone";xml:"office_phone"`
    
    /* personal_email optional员工私人邮箱 */
    personal_email string `json:"personal_email";xml:"personal_email"`
    
    /* personal_mobile optional员工手机号码 */
    personal_mobile string `json:"personal_mobile";xml:"personal_mobile"`
    
    /* sex required员工性别 1：男; 2:女 */
    sex int64 `json:"sex";xml:"sex"`
    
    /* sub_id required子账号ID */
    sub_id int64 `json:"sub_id";xml:"sub_id"`
    
    /* work_location optional工作地点 */
    work_location string `json:"work_location";xml:"work_location"`
    
}

func (req *TaobaoSubuserEmployeeAddRequest) GetAPIName() string {
	return "taobao.subuser.employee.add"
}

/* TaobaoSubuserEmployeeAddResponse 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息） */
type TaobaoSubuserEmployeeAddResponse struct {
    
    /* is_success Basic操作是否成功 true:操作成功; false:操作失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoSubuserInfoUpdateRequest 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息） */
type TaobaoSubuserInfoUpdateRequest struct {
    
    /* is_disable_subaccount optional是否停用子账号 true:表示停用该子账号false:表示开启该子账号 */
    is_disable_subaccount bool `json:"is_disable_subaccount";xml:"is_disable_subaccount"`
    
    /* is_dispatch optional子账号是否参与分流 true:参与分流 false:不参与分流 */
    is_dispatch bool `json:"is_dispatch";xml:"is_dispatch"`
    
    /* sub_id required子账号ID */
    sub_id int64 `json:"sub_id";xml:"sub_id"`
    
}

func (req *TaobaoSubuserInfoUpdateRequest) GetAPIName() string {
	return "taobao.subuser.info.update"
}

/* TaobaoSubuserInfoUpdateResponse 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息） */
type TaobaoSubuserInfoUpdateResponse struct {
    
    /* is_success Basic操作是否成功 true:操作成功; false:操作失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* Gfq1qp4287ShopexMatrxiOfflineSendRequest 商派矩阵线下发货接口 */
type Gfq1qp4287ShopexMatrxiOfflineSendRequest struct {
    
    /* params optional参数 */
    params string `json:"params";xml:"params"`
    
}

func (req *Gfq1qp4287ShopexMatrxiOfflineSendRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.matrxi.offline.send"
}

/* Gfq1qp4287ShopexMatrxiOfflineSendResponse 商派矩阵线下发货接口 */
type Gfq1qp4287ShopexMatrxiOfflineSendResponse struct {
    
    /* msg_id Basic任务id */
    msg_id string `json:"msg_id";xml:"msg_id"`
    
}

/* TaobaoUopTobOrderCreateRequest ToB仓储发货 */
type TaobaoUopTobOrderCreateRequest struct {
    
    /* delivery_order optionalERP出库对象 */
    delivery_order DeliveryOrder `json:"delivery_order";xml:"delivery_order"`
    
}

func (req *TaobaoUopTobOrderCreateRequest) GetAPIName() string {
	return "taobao.uop.tob.order.create"
}

/* TaobaoUopTobOrderCreateResponse ToB仓储发货 */
type TaobaoUopTobOrderCreateResponse struct {
    
    /* delivery_orders Object Array订单 */
    delivery_orders Deliveryorder `json:"delivery_orders";xml:"delivery_orders"`
    
    /* flag Basicflag */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
}

/* QimenTaobaoUopTmsupdatemessagetoerpRequest ERP配消息回告（拒签、签收、揽收等） */
type QimenTaobaoUopTmsupdatemessagetoerpRequest struct {
    
    /* request optional */
    request Struct `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoUopTmsupdatemessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.uop.tmsupdatemessagetoerp"
}

/* QimenTaobaoUopTmsupdatemessagetoerpResponse ERP配消息回告（拒签、签收、揽收等） */
type QimenTaobaoUopTmsupdatemessagetoerpResponse struct {
    
    /* response Objectresponse */
    response Response `json:"response";xml:"response"`
    
}

/* QimenTaobaoIcpOrderStockoutordermessagetoerpRequest 出库单信息推送接口 */
type QimenTaobaoIcpOrderStockoutordermessagetoerpRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOutOrderlist optional出库单记录集 */
    entryOutOrderlist EntryOutOrderlist `json:"entryOutOrderlist";xml:"entryOutOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockoutordermessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockoutordermessagetoerp"
}

/* QimenTaobaoIcpOrderStockoutordermessagetoerpResponse 出库单信息推送接口 */
type QimenTaobaoIcpOrderStockoutordermessagetoerpResponse struct {
    
    /* result Object返回结果 */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoItemSkuDeleteRequest 删除一个sku的数据
需要删除的sku通过属性properties进行匹配查找 */
type TaobaoItemSkuDeleteRequest struct {
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* item_num optionalsku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败 */
    item_num int64 `json:"item_num";xml:"item_num"`
    
    /* item_price optionalsku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功 */
    item_price Price `json:"item_price";xml:"item_price"`
    
    /* lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
    lang string `json:"lang";xml:"lang"`
    
    /* num_iid requiredSku所属商品数字id，可通过 taobao.item.get 获取。必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemSkuDeleteRequest) GetAPIName() string {
	return "taobao.item.sku.delete"
}

/* TaobaoItemSkuDeleteResponse 删除一个sku的数据
需要删除的sku通过属性properties进行匹配查找 */
type TaobaoItemSkuDeleteResponse struct {
    
    /* sku ObjectSku结构 */
    sku Sku `json:"sku";xml:"sku"`
    
}

/* AlibabaEinvoiceProviderlistGetRequest 获取能提供电子发票开票服务的开票服务商 */
type AlibabaEinvoiceProviderlistGetRequest struct {
    
}

func (req *AlibabaEinvoiceProviderlistGetRequest) GetAPIName() string {
	return "alibaba.einvoice.providerlist.get"
}

/* AlibabaEinvoiceProviderlistGetResponse 获取能提供电子发票开票服务的开票服务商 */
type AlibabaEinvoiceProviderlistGetResponse struct {
    
    /* result Object查询结果集 */
    result ResultSet `json:"result";xml:"result"`
    
}

/* TaobaoTmcUserGetRequest 查询指定用户开通的消息通道和组 */
type TaobaoTmcUserGetRequest struct {
    
    /* fields required需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。 */
    fields string `json:"fields";xml:"fields"`
    
    /* nick required用户昵称 */
    nick string `json:"nick";xml:"nick"`
    
    /* user_platform optional用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 */
    user_platform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcUserGetRequest) GetAPIName() string {
	return "taobao.tmc.user.get"
}

/* TaobaoTmcUserGetResponse 查询指定用户开通的消息通道和组 */
type TaobaoTmcUserGetResponse struct {
    
    /* tmc_user Object开通的用户数据 */
    tmc_user TmcUser `json:"tmc_user";xml:"tmc_user"`
    
}

/* CainiaoWaybillprintClientupdateGetconfigRequest 获取客户端更新配置信息 */
type CainiaoWaybillprintClientupdateGetconfigRequest struct {
    
    /* lan_ip required服务发起机器局域网ip */
    lan_ip string `json:"lan_ip";xml:"lan_ip"`
    
    /* mac required服务发起机器的物理地址 */
    mac string `json:"mac";xml:"mac"`
    
    /* msgid required当前消息版本 */
    msgid int64 `json:"msgid";xml:"msgid"`
    
    /* version required当前打印客户端版本 */
    version string `json:"version";xml:"version"`
    
}

func (req *CainiaoWaybillprintClientupdateGetconfigRequest) GetAPIName() string {
	return "cainiao.waybillprint.clientupdate.getconfig"
}

/* CainiaoWaybillprintClientupdateGetconfigResponse 获取客户端更新配置信息 */
type CainiaoWaybillprintClientupdateGetconfigResponse struct {
    
    /* result Objectresult */
    result UpdateConfigTopResult `json:"result";xml:"result"`
    
}

/* TmallTraderateItemtagsGetRequest 通过商品ID获取标签详细信息 */
type TmallTraderateItemtagsGetRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TmallTraderateItemtagsGetRequest) GetAPIName() string {
	return "tmall.traderate.itemtags.get"
}

/* TmallTraderateItemtagsGetResponse 通过商品ID获取标签详细信息 */
type TmallTraderateItemtagsGetResponse struct {
    
    /* tags Object Array标签列表 */
    tags TmallRateTagDetail `json:"tags";xml:"tags"`
    
}

/* TmallTraderateFeedsGetRequest 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签 */
type TmallTraderateFeedsGetRequest struct {
    
    /* child_trade_id required交易子订单ID */
    child_trade_id int64 `json:"child_trade_id";xml:"child_trade_id"`
    
}

func (req *TmallTraderateFeedsGetRequest) GetAPIName() string {
	return "tmall.traderate.feeds.get"
}

/* TmallTraderateFeedsGetResponse 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签 */
type TmallTraderateFeedsGetResponse struct {
    
    /* tmall_rate_info Object返回评价信息 */
    tmall_rate_info TmallRateInfo `json:"tmall_rate_info";xml:"tmall_rate_info"`
    
}

/* TaobaoFilesGetRequest 获取业务方暂存给ISV的文件列表 */
type TaobaoFilesGetRequest struct {
    
    /* end_date required搜索结束时间 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* start_date required搜索开始时间 */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* status optional下载链接状态。1:未下载。2:已下载 */
    status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoFilesGetRequest) GetAPIName() string {
	return "taobao.files.get"
}

/* TaobaoFilesGetResponse 获取业务方暂存给ISV的文件列表 */
type TaobaoFilesGetResponse struct {
    
    /* results Object Arrayresults */
    results TopDownloadRecordDo `json:"results";xml:"results"`
    
}

/* TaobaoWlbWaybillShengxianGetRequest 商家通过交易订单号获取电子面单接口 */
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

/* TaobaoWlbWaybillShengxianGetResponse 商家通过交易订单号获取电子面单接口 */
type TaobaoWlbWaybillShengxianGetResponse struct {
    
    /* fresh_waybill Object成功后返回的生鲜电子面单信息 */
    fresh_waybill FreshWaybill `json:"fresh_waybill";xml:"fresh_waybill"`
    
    /* is_success Basic生成是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoProductQuantityUpdateRequest 修改产品库存信息，支持全量修改以及增量修改两种方式 */
type TaobaoFenxiaoProductQuantityUpdateRequest struct {
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties optionalsku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity required库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。
当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
    quantity string `json:"quantity";xml:"quantity"`
    
    /* _type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoFenxiaoProductQuantityUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.quantity.update"
}

/* TaobaoFenxiaoProductQuantityUpdateResponse 修改产品库存信息，支持全量修改以及增量修改两种方式 */
type TaobaoFenxiaoProductQuantityUpdateResponse struct {
    
    /* created Basic操作时间 */
    created time.Time `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}

/* CainiaoOpenstorageSellerResourceCreateRequest 商家资源创建接口(云打印开源存储) */
type CainiaoOpenstorageSellerResourceCreateRequest struct {
    
    /* param_create_seller_resource_request optional商家创建资源参数 */
    param_create_seller_resource_request CreateSellerResourceRequest `json:"param_create_seller_resource_request";xml:"param_create_seller_resource_request"`
    
}

func (req *CainiaoOpenstorageSellerResourceCreateRequest) GetAPIName() string {
	return "cainiao.openstorage.seller.resource.create"
}

/* CainiaoOpenstorageSellerResourceCreateResponse 商家资源创建接口(云打印开源存储) */
type CainiaoOpenstorageSellerResourceCreateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* CainiaoOpenstorageResourcePublishRequest isv和商家资源发布接口(云打印开源存储) */
type CainiaoOpenstorageResourcePublishRequest struct {
    
    /* param_publish_resource_request optional资源发布参数 */
    param_publish_resource_request PublishResourceRequest `json:"param_publish_resource_request";xml:"param_publish_resource_request"`
    
}

func (req *CainiaoOpenstorageResourcePublishRequest) GetAPIName() string {
	return "cainiao.openstorage.resource.publish"
}

/* CainiaoOpenstorageResourcePublishResponse isv和商家资源发布接口(云打印开源存储) */
type CainiaoOpenstorageResourcePublishResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* AlipayPointBudgetGetRequest 查询已采购的集分宝余额，操作流程如下：
1、申请支付宝增值包。
2、申请支付宝应用上线时选择集分宝API。
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。 */
type AlipayPointBudgetGetRequest struct {
    
    /* auth_token special支付宝给用户的授权。如果没有top的授权，这个字段是必填项 */
    auth_token string `json:"auth_token";xml:"auth_token"`
    
}

func (req *AlipayPointBudgetGetRequest) GetAPIName() string {
	return "alipay.point.budget.get"
}

/* AlipayPointBudgetGetResponse 查询已采购的集分宝余额，操作流程如下：
1、申请支付宝增值包。
2、申请支付宝应用上线时选择集分宝API。
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。 */
type AlipayPointBudgetGetResponse struct {
    
    /* budget_amount Basic还可以发放的集分宝个数 */
    budget_amount int64 `json:"budget_amount";xml:"budget_amount"`
    
}

/* CainiaoOpenstorageResourceUpdateRequest isv和商家的资源获取接口（云打印开源存储） */
type CainiaoOpenstorageResourceUpdateRequest struct {
    
    /* param_update_resource_request required入参 */
    param_update_resource_request UpdateResourceRequest `json:"param_update_resource_request";xml:"param_update_resource_request"`
    
}

func (req *CainiaoOpenstorageResourceUpdateRequest) GetAPIName() string {
	return "cainiao.openstorage.resource.update"
}

/* CainiaoOpenstorageResourceUpdateResponse isv和商家的资源获取接口（云打印开源存储） */
type CainiaoOpenstorageResourceUpdateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* CainiaoOpenstorageIsvResourcesGetRequest isv资源列表查询（云打印开源存储） */
type CainiaoOpenstorageIsvResourcesGetRequest struct {
    
    /* isv_resource_type requiredisv资源类型，分为：TEMPLATE（模板）和PRINT_ITEM（打印项） */
    isv_resource_type string `json:"isv_resource_type";xml:"isv_resource_type"`
    
}

func (req *CainiaoOpenstorageIsvResourcesGetRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resources.get"
}

/* CainiaoOpenstorageIsvResourcesGetResponse isv资源列表查询（云打印开源存储） */
type CainiaoOpenstorageIsvResourcesGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* AlipayPointOrderAddRequest 向用户发送集分宝，发放集分宝之前，操作流程如下：
1、申请支付宝增值包。
2、申请支付宝应用上线时选择集分宝API。
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
7、调用发放API（alipay.point.order.add）发放集分宝。 */
type AlipayPointOrderAddRequest struct {
    
    /* auth_token special支付宝用户给应用发放集分宝的授权。 */
    auth_token string `json:"auth_token";xml:"auth_token"`
    
    /* memo required向用户展示集分宝发放备注 */
    memo string `json:"memo";xml:"memo"`
    
    /* merchant_order_no requiredisv提供的发放号订单号，由数字和组成，最大长度为32为，需要保证每笔发放的唯一性，支付宝会对该参数做唯一性控制。如果使用同样的订单号，支付宝将返回订单号已经存在的错误 */
    merchant_order_no string `json:"merchant_order_no";xml:"merchant_order_no"`
    
    /* order_time required发放集分宝时间 */
    order_time time.Time `json:"order_time";xml:"order_time"`
    
    /* point_count required发放集分宝的数量 */
    point_count int64 `json:"point_count";xml:"point_count"`
    
    /* user_symbol required用户标识符，用于指定集分宝发放的用户，和user_symbol_type一起使用，确定一个唯一的支付宝用户 */
    user_symbol string `json:"user_symbol";xml:"user_symbol"`
    
    /* user_symbol_type required用户标识符类型，现在支持ALIPAY_USER_ID:表示支付宝用户ID,ALIPAY_LOGON_ID:表示支付宝登陆号 */
    user_symbol_type string `json:"user_symbol_type";xml:"user_symbol_type"`
    
}

func (req *AlipayPointOrderAddRequest) GetAPIName() string {
	return "alipay.point.order.add"
}

/* AlipayPointOrderAddResponse 向用户发送集分宝，发放集分宝之前，操作流程如下：
1、申请支付宝增值包。
2、申请支付宝应用上线时选择集分宝API。
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
7、调用发放API（alipay.point.order.add）发放集分宝。 */
type AlipayPointOrderAddResponse struct {
    
    /* alipay_order_no Basic支付宝集分宝发放流水号 */
    alipay_order_no string `json:"alipay_order_no";xml:"alipay_order_no"`
    
    /* result_code Basic充值结果：SUCCESS表示成功，其他表示失败 */
    result_code bool `json:"result_code";xml:"result_code"`
    
}

/* TaobaoFenxiaoProductUpdateRequest 更新分销平台产品数据，不传更新数据返回失败<br>
1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br> */
type TaobaoFenxiaoProductUpdateRequest struct {
    
    /* category_id optional所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* city optional所在地：市，例：“杭州” */
    city string `json:"city";xml:"city"`
    
    /* cost_price optional代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    cost_price string `json:"cost_price";xml:"cost_price"`
    
    /* dealer_cost_price optional经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    dealer_cost_price string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    /* desc optional产品描述，长度为5到25000字符。 */
    desc string `json:"desc";xml:"desc"`
    
    /* discount_id optional折扣ID */
    discount_id int64 `json:"discount_id";xml:"discount_id"`
    
    /* have_invoice optional是否有发票，可选值：false（否）、true（是），默认false。 */
    have_invoice string `json:"have_invoice";xml:"have_invoice"`
    
    /* have_quarantee optional是否有保修，可选值：false（否）、true（是），默认false。 */
    have_quarantee string `json:"have_quarantee";xml:"have_quarantee"`
    
    /* image optional主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数 */
    image []byte `json:"image";xml:"image"`
    
    /* input_properties optional自定义属性。格式为pid:value;pid:value */
    input_properties string `json:"input_properties";xml:"input_properties"`
    
    /* is_authz optional产品是否需要授权isAuthz:yes|no 
yes:需要授权 
no:不需要授权 */
    is_authz string `json:"is_authz";xml:"is_authz"`
    
    /* name optional产品名称，长度不超过60个字节。 */
    name string `json:"name";xml:"name"`
    
    /* outer_id optional商家编码，长度不能超过60个字节。 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* pic_path optional产品主图图片空间相对路径或绝对路径 */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* pid required产品ID */
    pid int64 `json:"pid";xml:"pid"`
    
    /* postage_ems optionalems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
    postage_ems string `json:"postage_ems";xml:"postage_ems"`
    
    /* postage_fast optional快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
    postage_fast string `json:"postage_fast";xml:"postage_fast"`
    
    /* postage_id optional运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。 */
    postage_id int64 `json:"postage_id";xml:"postage_id"`
    
    /* postage_ordinary optional平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
    postage_ordinary string `json:"postage_ordinary";xml:"postage_ordinary"`
    
    /* postage_type optional运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。 */
    postage_type string `json:"postage_type";xml:"postage_type"`
    
    /* properties optional产品属性 */
    properties string `json:"properties";xml:"properties"`
    
    /* property_alias optional属性别名 */
    property_alias string `json:"property_alias";xml:"property_alias"`
    
    /* prov optional所在地：省，例：“浙江” */
    prov string `json:"prov";xml:"prov"`
    
    /* quantity optional产品库存必须是1到999999。 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* retail_price_high optional最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。 */
    retail_price_high string `json:"retail_price_high";xml:"retail_price_high"`
    
    /* retail_price_low optional最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    retail_price_low string `json:"retail_price_low";xml:"retail_price_low"`
    
    /* sku_cost_prices optionalsku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。 */
    sku_cost_prices string `json:"sku_cost_prices";xml:"sku_cost_prices"`
    
    /* sku_dealer_cost_prices optionalsku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。 */
    sku_dealer_cost_prices string `json:"sku_dealer_cost_prices";xml:"sku_dealer_cost_prices"`
    
    /* sku_ids optionalsku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。 */
    sku_ids string `json:"sku_ids";xml:"sku_ids"`
    
    /* sku_outer_ids optionalsku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",," */
    sku_outer_ids string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    /* sku_properties optionalsku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)
通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。 */
    sku_properties string `json:"sku_properties";xml:"sku_properties"`
    
    /* sku_properties_del optional根据sku属性删除sku信息。需要按组删除属性。 */
    sku_properties_del string `json:"sku_properties_del";xml:"sku_properties_del"`
    
    /* sku_quantitys optionalsku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。 */
    sku_quantitys string `json:"sku_quantitys";xml:"sku_quantitys"`
    
    /* sku_standard_prices optionalsku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。 */
    sku_standard_prices string `json:"sku_standard_prices";xml:"sku_standard_prices"`
    
    /* standard_price optional采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    standard_price string `json:"standard_price";xml:"standard_price"`
    
    /* standard_retail_price optional零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    standard_retail_price string `json:"standard_retail_price";xml:"standard_retail_price"`
    
    /* status optional发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。 */
    status string `json:"status";xml:"status"`
    
}

func (req *TaobaoFenxiaoProductUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.update"
}

/* TaobaoFenxiaoProductUpdateResponse 更新分销平台产品数据，不传更新数据返回失败<br>
1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br> */
type TaobaoFenxiaoProductUpdateResponse struct {
    
    /* modified Basic更新时间，时间格式：yyyy-MM-dd HH:mm:ss */
    modified time.Time `json:"modified";xml:"modified"`
    
    /* pid Basic产品ID */
    pid int64 `json:"pid";xml:"pid"`
    
}

/* TaobaoSubuserFullinfoGetRequest 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息） */
type TaobaoSubuserFullinfoGetRequest struct {
    
    /* fields optional传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email） */
    fields string `json:"fields";xml:"fields"`
    
    /* sub_id special子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号） */
    sub_id int64 `json:"sub_id";xml:"sub_id"`
    
    /* sub_nick special子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号） */
    sub_nick string `json:"sub_nick";xml:"sub_nick"`
    
}

func (req *TaobaoSubuserFullinfoGetRequest) GetAPIName() string {
	return "taobao.subuser.fullinfo.get"
}

/* TaobaoSubuserFullinfoGetResponse 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息） */
type TaobaoSubuserFullinfoGetResponse struct {
    
    /* sub_fullinfo Object子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息 */
    sub_fullinfo SubUserFullInfo `json:"sub_fullinfo";xml:"sub_fullinfo"`
    
}

/* TaobaoFenxiaoProductcatsGetRequest 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参 */
type TaobaoFenxiaoProductcatsGetRequest struct {
    
    /* fields optional返回字段列表 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoFenxiaoProductcatsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.productcats.get"
}

/* TaobaoFenxiaoProductcatsGetResponse 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参 */
type TaobaoFenxiaoProductcatsGetResponse struct {
    
    /* productcats Object Array产品线列表。返回 ProductCat 包含的字段信息。 */
    productcats ProductCat `json:"productcats";xml:"productcats"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoSubusersGetRequest 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息） */
type TaobaoSubusersGetRequest struct {
    
    /* user_nick required主账号用户名 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubusersGetRequest) GetAPIName() string {
	return "taobao.subusers.get"
}

/* TaobaoSubusersGetResponse 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息） */
type TaobaoSubusersGetResponse struct {
    
    /* subaccounts Object Array子账号基本信息 */
    subaccounts SubAccountInfo `json:"subaccounts";xml:"subaccounts"`
    
}

/* TaobaoFenxiaoProductAddRequest 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。

    * 产品图片默认为空
    * 产品发布后默认为下架状态 */
type TaobaoFenxiaoProductAddRequest struct {
    
    /* category_id required所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* city required所在地：市，例：“杭州” */
    city string `json:"city";xml:"city"`
    
    /* cost_price optional代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    cost_price string `json:"cost_price";xml:"cost_price"`
    
    /* dealer_cost_price optional经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    dealer_cost_price string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    /* desc required产品描述，长度为5到25000字符。 */
    desc string `json:"desc";xml:"desc"`
    
    /* discount_id optional折扣ID */
    discount_id int64 `json:"discount_id";xml:"discount_id"`
    
    /* have_invoice optional是否有发票，可选值：false（否）、true（是），默认false。 */
    have_invoice string `json:"have_invoice";xml:"have_invoice"`
    
    /* have_quarantee optional是否有保修，可选值：false（否）、true（是），默认false。 */
    have_quarantee string `json:"have_quarantee";xml:"have_quarantee"`
    
    /* image optional产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片 */
    image []byte `json:"image";xml:"image"`
    
    /* input_properties optional自定义属性。格式为pid:value;pid:value */
    input_properties string `json:"input_properties";xml:"input_properties"`
    
    /* is_authz optional添加产品时，添加入参isAuthz:yes|no 
yes:需要授权 
no:不需要授权 
默认是需要授权 */
    is_authz string `json:"is_authz";xml:"is_authz"`
    
    /* item_id optional导入的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* name required产品名称，长度不超过60个字节。 */
    name string `json:"name";xml:"name"`
    
    /* outer_id optional商家编码，长度不能超过60个字节。 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* pic_path optional产品主图图片空间相对路径或绝对路径 */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* postage_ems optionalems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。 */
    postage_ems string `json:"postage_ems";xml:"postage_ems"`
    
    /* postage_fast optional快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。 */
    postage_fast string `json:"postage_fast";xml:"postage_fast"`
    
    /* postage_id optional运费模板ID，参考taobao.postages.get。 */
    postage_id int64 `json:"postage_id";xml:"postage_id"`
    
    /* postage_ordinary optional平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。 */
    postage_ordinary string `json:"postage_ordinary";xml:"postage_ordinary"`
    
    /* postage_type optional运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。 */
    postage_type string `json:"postage_type";xml:"postage_type"`
    
    /* productcat_id required产品线ID */
    productcat_id int64 `json:"productcat_id";xml:"productcat_id"`
    
    /* properties optional产品属性，格式为pid:vid;pid:vid */
    properties string `json:"properties";xml:"properties"`
    
    /* property_alias optional属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名） */
    property_alias string `json:"property_alias";xml:"property_alias"`
    
    /* prov required所在地：省，例：“浙江” */
    prov string `json:"prov";xml:"prov"`
    
    /* quantity required产品库存必须是1到999999。 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* retail_price_high required最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。 */
    retail_price_high string `json:"retail_price_high";xml:"retail_price_high"`
    
    /* retail_price_low required最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    retail_price_low string `json:"retail_price_low";xml:"retail_price_low"`
    
    /* sku_cost_prices optionalsku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
    sku_cost_prices string `json:"sku_cost_prices";xml:"sku_cost_prices"`
    
    /* sku_dealer_cost_prices optionalsku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。 */
    sku_dealer_cost_prices string `json:"sku_dealer_cost_prices";xml:"sku_dealer_cost_prices"`
    
    /* sku_outer_ids optionalsku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
    sku_outer_ids string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    /* sku_properties optionalsku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
    sku_properties string `json:"sku_properties";xml:"sku_properties"`
    
    /* sku_quantitys optionalsku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
    sku_quantitys string `json:"sku_quantitys";xml:"sku_quantitys"`
    
    /* sku_standard_prices optionalsku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
    sku_standard_prices string `json:"sku_standard_prices";xml:"sku_standard_prices"`
    
    /* spu_id optional产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传 */
    spu_id int64 `json:"spu_id";xml:"spu_id"`
    
    /* standard_price required采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    standard_price string `json:"standard_price";xml:"standard_price"`
    
    /* standard_retail_price optional零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
    standard_retail_price string `json:"standard_retail_price";xml:"standard_retail_price"`
    
    /* trade_type optional分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做） */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.add"
}

/* TaobaoFenxiaoProductAddResponse 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。

    * 产品图片默认为空
    * 产品发布后默认为下架状态 */
type TaobaoFenxiaoProductAddResponse struct {
    
    /* created Basic产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss */
    created time.Time `json:"created";xml:"created"`
    
    /* pid Basic产品ID */
    pid int64 `json:"pid";xml:"pid"`
    
}

/* TaobaoTmcMsgSendrecordRequest 查询单条消息发送记录，只返回返回条数和时间。 */
type TaobaoTmcMsgSendrecordRequest struct {
    
    /* data_id required消息主键ID */
    data_id string `json:"data_id";xml:"data_id"`
    
    /* group_name required消息分组名 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* topic_name requiredTOPIC名称 */
    topic_name string `json:"topic_name";xml:"topic_name"`
    
}

func (req *TaobaoTmcMsgSendrecordRequest) GetAPIName() string {
	return "taobao.tmc.msg.sendrecord"
}

/* TaobaoTmcMsgSendrecordResponse 查询单条消息发送记录，只返回返回条数和时间。 */
type TaobaoTmcMsgSendrecordResponse struct {
    
    /* tb_send_list Basic淘宝发送时间 */
    tb_send_list string `json:"tb_send_list";xml:"tb_send_list"`
    
    /* tb_send_times Basic淘宝发送测试 */
    tb_send_times int64 `json:"tb_send_times";xml:"tb_send_times"`
    
    /* tmc_send_list BasicTMC发送时间 */
    tmc_send_list string `json:"tmc_send_list";xml:"tmc_send_list"`
    
    /* tmc_send_times Basictmc发送次数 */
    tmc_send_times int64 `json:"tmc_send_times";xml:"tmc_send_times"`
    
}

/* TaobaoWangwangEserviceChatpeersGetRequest 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。 */
type TaobaoWangwangEserviceChatpeersGetRequest struct {
    
}

func (req *TaobaoWangwangEserviceChatpeersGetRequest) GetAPIName() string {
	return "taobao.wangwang.eservice.chatpeers.get"
}

/* TaobaoWangwangEserviceChatpeersGetResponse 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。 */
type TaobaoWangwangEserviceChatpeersGetResponse struct {
    
    /* chatpeers Object Array聊天对象ID列表 */
    chatpeers Chatpeer `json:"chatpeers";xml:"chatpeers"`
    
    /* count Basic成员数目。 */
    count int64 `json:"count";xml:"count"`
    
    /* ret Basic返回码： 
10000:成功； 

60000：时间非法，包括1)时间段超过7天,或2)起始时间距今超过30天，或3)时间格式不对； 

50000：聊天用户ID不是该店铺的帐号； 

40000：系统错误，包括必填参数为空。 */
    ret int64 `json:"ret";xml:"ret"`
    
}

/* TaobaoFenxiaoProductsGetRequest 查询供应商的产品数据。

    * 入参传入pids将优先查询，即只按这个条件查询。
    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
    * 入参fields传入images将查询多图数据，不传只返回主图数据。
    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
    * 查询结果按照产品发布时间倒序，即时间近的数据在前。 */
type TaobaoFenxiaoProductsGetRequest struct {
    
    /* end_modified optional结束修改时间 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* fields optional指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
    fields string `json:"fields";xml:"fields"`
    
    /* is_authz optional查询产品列表时，查询入参“是否需要授权”
yes:需要授权 
no:不需要授权 */
    is_authz string `json:"is_authz";xml:"is_authz"`
    
    /* item_ids optional可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
    item_ids string `json:"item_ids";xml:"item_ids"`
    
    /* outer_id optional商家编码 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* pids optional产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005” */
    pids string `json:"pids";xml:"pids"`
    
    /* productcat_id optional产品线ID */
    productcat_id int64 `json:"productcat_id";xml:"productcat_id"`
    
    /* sku_number optionalsku商家编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
    /* start_modified optional开始修改时间 */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoFenxiaoProductsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.products.get"
}

/* TaobaoFenxiaoProductsGetResponse 查询供应商的产品数据。

    * 入参传入pids将优先查询，即只按这个条件查询。
    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
    * 入参fields传入images将查询多图数据，不传只返回主图数据。
    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
    * 查询结果按照产品发布时间倒序，即时间近的数据在前。 */
type TaobaoFenxiaoProductsGetResponse struct {
    
    /* products Object Array产品对象记录集。返回 FenxiaoProduct 包含的字段信息。 */
    products FenxiaoProduct `json:"products";xml:"products"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoSubuserDepartmentsGetRequest 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。 */
type TaobaoSubuserDepartmentsGetRequest struct {
    
    /* user_nick required主账号用户名 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubuserDepartmentsGetRequest) GetAPIName() string {
	return "taobao.subuser.departments.get"
}

/* TaobaoSubuserDepartmentsGetResponse 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。 */
type TaobaoSubuserDepartmentsGetResponse struct {
    
    /* departments Object Array部门信息 */
    departments Department `json:"departments";xml:"departments"`
    
}

/* AlipayUserAccountreportGetRequest 获取支付宝余额明细记录，不包含用户通过银行卡快捷支付进行交易的交易明细 */
type AlipayUserAccountreportGetRequest struct {
    
    /* end_time required对账单结束时间，其中end_time - start_time <= 1天，对于对账记录比较多的情况，强烈建议按天查询，否则会出现超时的情况。 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* fields required需要返回的字段列表。create_time:创建时间,type：账务类型,business_type:子业务类型,balance:当时支付宝账户余额,in_amount:收入金额,out_amount:支出金额,alipay_order_no:支付宝订单号,merchant_order_no:商户订单号,self_user_id:自己的支付宝ID,opt_user_id:对方的支付宝ID,memo:账号备注 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional要获取的对账单页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每次查询获取对账记录数量 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required对账单开始时间。最近一个月内的日期。 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
    /* _type optional账务类型。多个类型是，用逗号分隔，不传则查询所有类型的。PAYMENT:在线支付，TRANSFER:转账，DEPOSIT:充值，WITHDRAW:提现，CHARGE:收费，PREAUTH:预授权，OTHER：其它。 */
    _type string `json:"type";xml:"type"`
    
}

func (req *AlipayUserAccountreportGetRequest) GetAPIName() string {
	return "alipay.user.accountreport.get"
}

/* AlipayUserAccountreportGetResponse 获取支付宝余额明细记录，不包含用户通过银行卡快捷支付进行交易的交易明细 */
type AlipayUserAccountreportGetResponse struct {
    
    /* alipay_records Object Array对账记录列表 */
    alipay_records AlipayRecord `json:"alipay_records";xml:"alipay_records"`
    
    /* total_pages Basic总页数 */
    total_pages int64 `json:"total_pages";xml:"total_pages"`
    
    /* total_results Basic总记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoFuwuSaleLinkGenRequest 服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br>
注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。 */
type TaobaoFuwuSaleLinkGenRequest struct {
    
    /* nick optional用户需要营销的目标人群中的用户nick */
    nick string `json:"nick";xml:"nick"`
    
    /* param_str required从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。 */
    param_str string `json:"param_str";xml:"param_str"`
    
}

func (req *TaobaoFuwuSaleLinkGenRequest) GetAPIName() string {
	return "taobao.fuwu.sale.link.gen"
}

/* TaobaoFuwuSaleLinkGenResponse 服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br>
注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。 */
type TaobaoFuwuSaleLinkGenResponse struct {
    
    /* url Basic通过营销链接接口生成的营销链接短地址 */
    url string `json:"url";xml:"url"`
    
}

/* TmallProductSchemaGetRequest 产品信息获取接口schema形式返回 */
type TmallProductSchemaGetRequest struct {
    
    /* product_id required产品编号 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallProductSchemaGetRequest) GetAPIName() string {
	return "tmall.product.schema.get"
}

/* TmallProductSchemaGetResponse 产品信息获取接口schema形式返回 */
type TmallProductSchemaGetResponse struct {
    
    /* get_product_result Basic产品信息数据。schema形式 */
    get_product_result string `json:"get_product_result";xml:"get_product_result"`
    
}

/* CainiaoCloudprintIsvResourcesGetRequest isv资源查询，包括isv模板、打印项、预设的自定义区等 */
type CainiaoCloudprintIsvResourcesGetRequest struct {
    
    /* isv_resource_type requiredisv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区） */
    isv_resource_type string `json:"isv_resource_type";xml:"isv_resource_type"`
    
}

func (req *CainiaoCloudprintIsvResourcesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.isv.resources.get"
}

/* CainiaoCloudprintIsvResourcesGetResponse isv资源查询，包括isv模板、打印项、预设的自定义区等 */
type CainiaoCloudprintIsvResourcesGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* CainiaoOpenstorageIsvResourceCreateRequest isv资源创建接口（云打印开源存储） */
type CainiaoOpenstorageIsvResourceCreateRequest struct {
    
    /* param_create_isv_resource_request optionalisv创建资源的参数 */
    param_create_isv_resource_request CreateIsvResourceRequest `json:"param_create_isv_resource_request";xml:"param_create_isv_resource_request"`
    
}

func (req *CainiaoOpenstorageIsvResourceCreateRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resource.create"
}

/* CainiaoOpenstorageIsvResourceCreateResponse isv资源创建接口（云打印开源存储） */
type CainiaoOpenstorageIsvResourceCreateResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_message Basic错误消息 */
    err_message string `json:"err_message";xml:"err_message"`
    
}

/* TaobaoTopSecretGetRequest top sdk通过api获取对应解密秘钥 */
type TaobaoTopSecretGetRequest struct {
    
    /* customer_user_id optional自定义用户id */
    customer_user_id int64 `json:"customer_user_id";xml:"customer_user_id"`
    
    /* random_num required伪随机数 */
    random_num string `json:"random_num";xml:"random_num"`
    
    /* secret_version optional秘钥版本号 */
    secret_version int64 `json:"secret_version";xml:"secret_version"`
    
}

func (req *TaobaoTopSecretGetRequest) GetAPIName() string {
	return "taobao.top.secret.get"
}

/* TaobaoTopSecretGetResponse top sdk通过api获取对应解密秘钥 */
type TaobaoTopSecretGetResponse struct {
    
    /* app_config Basicapp配置信息 */
    app_config string `json:"app_config";xml:"app_config"`
    
    /* interval Basic下次更新秘钥间隔，单位（秒） */
    interval int64 `json:"interval";xml:"interval"`
    
    /* max_interval Basic最长有效期，容灾使用，单位（秒） */
    max_interval int64 `json:"max_interval";xml:"max_interval"`
    
    /* secret Basic秘钥值 */
    secret string `json:"secret";xml:"secret"`
    
    /* secret_version Basic秘钥版本号 */
    secret_version int64 `json:"secret_version";xml:"secret_version"`
    
}

/* TaobaoLogisticsAddressReachablebatchGetRequest 批量判定服务是否可达 */
type TaobaoLogisticsAddressReachablebatchGetRequest struct {
    
    /* address_list optional筛单用户输入地址结构 */
    address_list AddressReachable `json:"address_list";xml:"address_list"`
    
}

func (req *TaobaoLogisticsAddressReachablebatchGetRequest) GetAPIName() string {
	return "taobao.logistics.address.reachablebatch.get"
}

/* TaobaoLogisticsAddressReachablebatchGetResponse 批量判定服务是否可达 */
type TaobaoLogisticsAddressReachablebatchGetResponse struct {
    
    /* reachable_results Object Array物流是否可达结果列表 */
    reachable_results AddressReachableTopResult `json:"reachable_results";xml:"reachable_results"`
    
}

/* TaobaoOmniorderStoreReallocateRequest 门店发货提供改派接口 */
type TaobaoOmniorderStoreReallocateRequest struct {
    
    /* main_order_id required主订单号 */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
    /* store_id optional门店Id */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* sub_order_ids required子订单号 */
    sub_order_ids int64 `json:"sub_order_ids";xml:"sub_order_ids"`
    
    /* warehouse_code optional电商仓code */
    warehouse_code string `json:"warehouse_code";xml:"warehouse_code"`
    
}

func (req *TaobaoOmniorderStoreReallocateRequest) GetAPIName() string {
	return "taobao.omniorder.store.reallocate"
}

/* TaobaoOmniorderStoreReallocateResponse 门店发货提供改派接口 */
type TaobaoOmniorderStoreReallocateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoQimenSingleitemQueryRequest 商品查询接口 */
type TaobaoQimenSingleitemQueryRequest struct {
    
    /* itemCode optional商品编码,S1234,string(50),必填, */
    itemCode string `json:"itemCode";xml:"itemCode"`
    
    /* itemId optional仓储系统商品编码,C123,string(50),必填, */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional货主编码,H123,string(50),必填, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenSingleitemQueryRequest) GetAPIName() string {
	return "taobao.qimen.singleitem.query"
}

/* TaobaoQimenSingleitemQueryResponse 商品查询接口 */
type TaobaoQimenSingleitemQueryResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* item Objectitem */
    item Item `json:"item";xml:"item"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenItemmappingCreateRequest 前后端商品映射 */
type TaobaoQimenItemmappingCreateRequest struct {
    
    /* actionType optional奇门仓储字段,C123,string(50),必填, */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* itemId optional奇门仓储字段,C123,string(50),必填, */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* itemSource optional奇门仓储字段,C123,string(50),必填, */
    itemSource string `json:"itemSource";xml:"itemSource"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),必填, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* shopItemId optional奇门仓储字段,C123,string(50),必填, */
    shopItemId string `json:"shopItemId";xml:"shopItemId"`
    
    /* shopNick optional奇门仓储字段,C123,string(50),必填, */
    shopNick string `json:"shopNick";xml:"shopNick"`
    
    /* skuId optional奇门仓储字段,C123,string(50),必填, */
    skuId string `json:"skuId";xml:"skuId"`
    
}

func (req *TaobaoQimenItemmappingCreateRequest) GetAPIName() string {
	return "taobao.qimen.itemmapping.create"
}

/* TaobaoQimenItemmappingCreateResponse 前后端商品映射 */
type TaobaoQimenItemmappingCreateResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenInventoryruleCreateRequest 渠道间库存规则设置 */
type TaobaoQimenInventoryruleCreateRequest struct {
    
    /* inventoryRules optionalinventoryRules */
    inventoryRules InventoryRule `json:"inventoryRules";xml:"inventoryRules"`
    
    /* ownerCode optional奇门仓储字段,C1223,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenInventoryruleCreateRequest) GetAPIName() string {
	return "taobao.qimen.inventoryrule.create"
}

/* TaobaoQimenInventoryruleCreateResponse 渠道间库存规则设置 */
type TaobaoQimenInventoryruleCreateResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenCombineitemDeleteRequest 组合货品删除 */
type TaobaoQimenCombineitemDeleteRequest struct {
    
    /* itemId optional奇门仓储字段,C123,string(50),, */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenCombineitemDeleteRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.delete"
}

/* TaobaoQimenCombineitemDeleteResponse 组合货品删除 */
type TaobaoQimenCombineitemDeleteResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TmallItemCalculateHscodeGetRequest 算法获取hscode */
type TmallItemCalculateHscodeGetRequest struct {
    
    /* item_id optional商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TmallItemCalculateHscodeGetRequest) GetAPIName() string {
	return "tmall.item.calculate.hscode.get"
}

/* TmallItemCalculateHscodeGetResponse 算法获取hscode */
type TmallItemCalculateHscodeGetResponse struct {
    
    /* results Basic Array算法返回预测的hscode数据 */
    results map[string]interface{} `json:"results";xml:"results"`
    
}

/* TaobaoQimenChannelinventoryQueryRequest 渠道库存查询 */
type TaobaoQimenChannelinventoryQueryRequest struct {
    
    /* channelCodes optional奇门仓储字段 */
    channelCodes string `json:"channelCodes";xml:"channelCodes"`
    
    /* itemCodes optional奇门仓储字段 */
    itemCodes string `json:"itemCodes";xml:"itemCodes"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCodes optional奇门仓储字段 */
    warehouseCodes string `json:"warehouseCodes";xml:"warehouseCodes"`
    
}

func (req *TaobaoQimenChannelinventoryQueryRequest) GetAPIName() string {
	return "taobao.qimen.channelinventory.query"
}

/* TaobaoQimenChannelinventoryQueryResponse 渠道库存查询 */
type TaobaoQimenChannelinventoryQueryResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* itemInventories Object ArrayitemInventories */
    itemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoFenxiaoOrderRemarkUpdateRequest 供应商修改采购单备注 */
type TaobaoFenxiaoOrderRemarkUpdateRequest struct {
    
    /* purchase_order_id required采购单编号 */
    purchase_order_id int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
    /* supplier_memo required备注内容(供应商操作) */
    supplier_memo string `json:"supplier_memo";xml:"supplier_memo"`
    
    /* supplier_memo_flag optional旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色 */
    supplier_memo_flag int64 `json:"supplier_memo_flag";xml:"supplier_memo_flag"`
    
}

func (req *TaobaoFenxiaoOrderRemarkUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.order.remark.update"
}

/* TaobaoFenxiaoOrderRemarkUpdateResponse 供应商修改采购单备注 */
type TaobaoFenxiaoOrderRemarkUpdateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoInventoryAdjustExternalRequest 商家非交易调整库存，调拨出库、盘点等时调用 */
type TaobaoInventoryAdjustExternalRequest struct {
    
    /* biz_type optional外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他 */
    biz_type string `json:"biz_type";xml:"biz_type"`
    
    /* biz_unique_code required商家外部定单号 */
    biz_unique_code string `json:"biz_unique_code";xml:"biz_unique_code"`
    
    /* items required商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}] */
    items string `json:"items";xml:"items"`
    
    /* occupy_operate_code optional库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致 */
    occupy_operate_code string `json:"occupy_operate_code";xml:"occupy_operate_code"`
    
    /* operate_time optional业务操作时间 */
    operate_time string `json:"operate_time";xml:"operate_time"`
    
    /* operate_type optionaltest */
    operate_type string `json:"operate_type";xml:"operate_type"`
    
    /* reduce_type optionaltest */
    reduce_type string `json:"reduce_type";xml:"reduce_type"`
    
    /* store_code required商家仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryAdjustExternalRequest) GetAPIName() string {
	return "taobao.inventory.adjust.external"
}

/* TaobaoInventoryAdjustExternalResponse 商家非交易调整库存，调拨出库、盘点等时调用 */
type TaobaoInventoryAdjustExternalResponse struct {
    
    /* operate_code Basic操作返回码 */
    operate_code string `json:"operate_code";xml:"operate_code"`
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

/* TaobaoInventoryQueryRequest 商家查询商品总体库存信息 */
type TaobaoInventoryQueryRequest struct {
    
    /* sc_item_codes optional后端商品的商家编码列表，控制到50个 */
    sc_item_codes string `json:"sc_item_codes";xml:"sc_item_codes"`
    
    /* sc_item_ids required后端商品ID 列表，控制到50个 */
    sc_item_ids string `json:"sc_item_ids";xml:"sc_item_ids"`
    
    /* seller_nick optional卖家昵称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
    /* store_codes optional仓库列表:GLY001^GLY002 */
    store_codes string `json:"store_codes";xml:"store_codes"`
    
}

func (req *TaobaoInventoryQueryRequest) GetAPIName() string {
	return "taobao.inventory.query"
}

/* TaobaoInventoryQueryResponse 商家查询商品总体库存信息 */
type TaobaoInventoryQueryResponse struct {
    
    /* item_inventorys Object Array商品总体库存信息 */
    item_inventorys InventorySum `json:"item_inventorys";xml:"item_inventorys"`
    
    /* tip_infos Object Array提示信息，提示不存在的后端商品 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

/* TaobaoInventoryInitialRequest 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账 */
type TaobaoInventoryInitialRequest struct {
    
    /* items required商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}] */
    items string `json:"items";xml:"items"`
    
    /* store_code required商家仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryInitialRequest) GetAPIName() string {
	return "taobao.inventory.initial"
}

/* TaobaoInventoryInitialResponse 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账 */
type TaobaoInventoryInitialResponse struct {
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

/* TmallItemAddSimpleschemaGetRequest 通过商家信息获取商品发布字段和规则。 */
type TmallItemAddSimpleschemaGetRequest struct {
    
}

func (req *TmallItemAddSimpleschemaGetRequest) GetAPIName() string {
	return "tmall.item.add.simpleschema.get"
}

/* TmallItemAddSimpleschemaGetResponse 通过商家信息获取商品发布字段和规则。 */
type TmallItemAddSimpleschemaGetResponse struct {
    
    /* result Basic返回发布商品的规则文档 */
    result string `json:"result";xml:"result"`
    
}

/* TaobaoScitemMapAddRequest 创建IC商品或分销商品与后端商品的映射关系 */
type TaobaoScitemMapAddRequest struct {
    
    /* item_id required前台ic商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* need_check optional默认值为false
true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联
false:不进行高级校验 */
    need_check bool `json:"need_check";xml:"need_check"`
    
    /* outer_code optionalsc_item_id和outer_code 其中一个不能为空 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* sc_item_id optionalsc_item_id和outer_code 其中一个不能为空 */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* sku_id optional前台ic商品skuid */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoScitemMapAddRequest) GetAPIName() string {
	return "taobao.scitem.map.add"
}

/* TaobaoScitemMapAddResponse 创建IC商品或分销商品与后端商品的映射关系 */
type TaobaoScitemMapAddResponse struct {
    
    /* outer_code Basic接口调用返回结果信息：商家编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
}

/* TaobaoInventoryAdjustTradeRequest 商家交易调整库存，淘宝交易、B2B经销等 */
type TaobaoInventoryAdjustTradeRequest struct {
    
    /* biz_unique_code required商家外部定单号 */
    biz_unique_code string `json:"biz_unique_code";xml:"biz_unique_code"`
    
    /* items required商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}] */
    items string `json:"items";xml:"items"`
    
    /* operate_time required业务操作时间 */
    operate_time time.Time `json:"operate_time";xml:"operate_time"`
    
    /* tb_order_type required订单类型：B2C、B2B */
    tb_order_type string `json:"tb_order_type";xml:"tb_order_type"`
    
}

func (req *TaobaoInventoryAdjustTradeRequest) GetAPIName() string {
	return "taobao.inventory.adjust.trade"
}

/* TaobaoInventoryAdjustTradeResponse 商家交易调整库存，淘宝交易、B2B经销等 */
type TaobaoInventoryAdjustTradeResponse struct {
    
    /* operate_code Basic操作返回码 */
    operate_code string `json:"operate_code";xml:"operate_code"`
    
    /* tip_infos Object Array提示信息 */
    tip_infos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

/* TaobaoFenxiaoProductcatUpdateRequest 修改产品线 */
type TaobaoFenxiaoProductcatUpdateRequest struct {
    
    /* agent_cost_percent optional代销默认采购价比例，注意：100.00%，则输入为10000 */
    agent_cost_percent int64 `json:"agent_cost_percent";xml:"agent_cost_percent"`
    
    /* dealer_cost_percent optional经销默认采购价比例，注意：100.00%，则输入为10000 */
    dealer_cost_percent int64 `json:"dealer_cost_percent";xml:"dealer_cost_percent"`
    
    /* name optional产品线名称 */
    name string `json:"name";xml:"name"`
    
    /* product_line_id required产品线ID */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
    /* retail_high_percent optional最高零售价比例，注意：100.00%，则输入为10000 */
    retail_high_percent int64 `json:"retail_high_percent";xml:"retail_high_percent"`
    
    /* retail_low_percent optional最低零售价比例，注意：100.00%，则输入为10000 */
    retail_low_percent int64 `json:"retail_low_percent";xml:"retail_low_percent"`
    
}

func (req *TaobaoFenxiaoProductcatUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.update"
}

/* TaobaoFenxiaoProductcatUpdateResponse 修改产品线 */
type TaobaoFenxiaoProductcatUpdateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoProductcatAddRequest 新增产品线 */
type TaobaoFenxiaoProductcatAddRequest struct {
    
    /* agent_cost_percent required代销默认采购价比例，注意：100.00%，则输入为10000 */
    agent_cost_percent int64 `json:"agent_cost_percent";xml:"agent_cost_percent"`
    
    /* dealer_cost_percent required经销默认采购价比例，注意：100.00%，则输入为10000 */
    dealer_cost_percent int64 `json:"dealer_cost_percent";xml:"dealer_cost_percent"`
    
    /* name required产品线名称 */
    name string `json:"name";xml:"name"`
    
    /* retail_high_percent required最高零售价比例，注意：100.00%，则输入为10000 */
    retail_high_percent int64 `json:"retail_high_percent";xml:"retail_high_percent"`
    
    /* retail_low_percent required最低零售价比例，注意：100.00%，则输入为10000 */
    retail_low_percent int64 `json:"retail_low_percent";xml:"retail_low_percent"`
    
}

func (req *TaobaoFenxiaoProductcatAddRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.add"
}

/* TaobaoFenxiaoProductcatAddResponse 新增产品线 */
type TaobaoFenxiaoProductcatAddResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* product_line_id Basic产品线ID */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
}

/* TaobaoFenxiaoGradeAddRequest 新建等级 */
type TaobaoFenxiaoGradeAddRequest struct {
    
    /* name required等级名称，等级名称不可重复 */
    name string `json:"name";xml:"name"`
    
}

func (req *TaobaoFenxiaoGradeAddRequest) GetAPIName() string {
	return "taobao.fenxiao.grade.add"
}

/* TaobaoFenxiaoGradeAddResponse 新建等级 */
type TaobaoFenxiaoGradeAddResponse struct {
    
    /* grade_id Basic等级ID */
    grade_id int64 `json:"grade_id";xml:"grade_id"`
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoProductcatDeleteRequest 删除产品线 */
type TaobaoFenxiaoProductcatDeleteRequest struct {
    
    /* product_line_id required产品线ID */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
}

func (req *TaobaoFenxiaoProductcatDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.delete"
}

/* TaobaoFenxiaoProductcatDeleteResponse 删除产品线 */
type TaobaoFenxiaoProductcatDeleteResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TmallItemSchemaIncrementUpdateRequest 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名） */
type TmallItemSchemaIncrementUpdateRequest struct {
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* xml_data required根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaIncrementUpdateRequest) GetAPIName() string {
	return "tmall.item.schema.increment.update"
}

/* TmallItemSchemaIncrementUpdateResponse 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名） */
type TmallItemSchemaIncrementUpdateResponse struct {
    
    /* gmt_modified Basic商品更新操作成功时间 */
    gmt_modified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    /* update_item_result Basic返回商品发布结果 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}

/* TaobaoTopIpoutGetRequest 获取开放平台出口IP段 */
type TaobaoTopIpoutGetRequest struct {
    
}

func (req *TaobaoTopIpoutGetRequest) GetAPIName() string {
	return "taobao.top.ipout.get"
}

/* TaobaoTopIpoutGetResponse 获取开放平台出口IP段 */
type TaobaoTopIpoutGetResponse struct {
    
    /* ip_list BasicTOP网关出口IP列表 */
    ip_list string `json:"ip_list";xml:"ip_list"`
    
}

/* TmallItemIncrementUpdateSchemaGetRequest 增量方式修改天猫商品的规则获取的API。
1.接口返回支持增量修改的字段以及相应字段的规则。
2.如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好次字段以提升API整体性能）；
3.ISV初次接入，开发阶段，此字段不填可以看到所有支持增量的字段；但是如果上线功能明确，请尽量遵守第2条
4.如果ISV对字段规则非常清晰，可以直接组装入参数据提交到tmall.item.schema.increment.update进行数据更新。但是最好不要写死，比如每天还是有对此接口功能的一次比对。
---（感谢爱慕旗舰店提供API命名） */
type TmallItemIncrementUpdateSchemaGetRequest struct {
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* xml_data optional如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好此字段以提升API整体性能） */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemIncrementUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.item.increment.update.schema.get"
}

/* TmallItemIncrementUpdateSchemaGetResponse 增量方式修改天猫商品的规则获取的API。
1.接口返回支持增量修改的字段以及相应字段的规则。
2.如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好次字段以提升API整体性能）；
3.ISV初次接入，开发阶段，此字段不填可以看到所有支持增量的字段；但是如果上线功能明确，请尽量遵守第2条
4.如果ISV对字段规则非常清晰，可以直接组装入参数据提交到tmall.item.schema.increment.update进行数据更新。但是最好不要写死，比如每天还是有对此接口功能的一次比对。
---（感谢爱慕旗舰店提供API命名） */
type TmallItemIncrementUpdateSchemaGetResponse struct {
    
    /* update_item_result Basic返回增量更新商品的规则文档 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}

/* TaobaoFenxiaoCooperationProductcatAddRequest 追加授权产品线 */
type TaobaoFenxiaoCooperationProductcatAddRequest struct {
    
    /* cooperate_id required合作关系id */
    cooperate_id int64 `json:"cooperate_id";xml:"cooperate_id"`
    
    /* grade_id optional等级ID（为空则不修改） */
    grade_id int64 `json:"grade_id";xml:"grade_id"`
    
    /* product_line_list required产品线id列表，若有多个，以逗号分隔 */
    product_line_list string `json:"product_line_list";xml:"product_line_list"`
    
}

func (req *TaobaoFenxiaoCooperationProductcatAddRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.productcat.add"
}

/* TaobaoFenxiaoCooperationProductcatAddResponse 追加授权产品线 */
type TaobaoFenxiaoCooperationProductcatAddResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TmallItemSimpleschemaAddRequest 天猫简化版schema发布商品。 */
type TmallItemSimpleschemaAddRequest struct {
    
    /* schema_xml_fields required根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据 */
    schema_xml_fields string `json:"schema_xml_fields";xml:"schema_xml_fields"`
    
}

func (req *TmallItemSimpleschemaAddRequest) GetAPIName() string {
	return "tmall.item.simpleschema.add"
}

/* TmallItemSimpleschemaAddResponse 天猫简化版schema发布商品。 */
type TmallItemSimpleschemaAddResponse struct {
    
    /* gmt_modified Basic商品最后发布时间。 */
    gmt_modified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    /* result Basic发布成功后返回商品ID */
    result string `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoCooperationAuditRequest 合作授权审批 */
type TaobaoFenxiaoCooperationAuditRequest struct {
    
    /* audit_result required1:审批通过，审批通过后要加入授权产品线列表；
2:审批拒绝 */
    audit_result int64 `json:"audit_result";xml:"audit_result"`
    
    /* product_line_list_agent optional当审批通过时需要指定授权产品线id列表(代销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
    product_line_list_agent string `json:"product_line_list_agent";xml:"product_line_list_agent"`
    
    /* product_line_list_dealer optional当审批通过时需要指定授权产品线id列表(经销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
    product_line_list_dealer string `json:"product_line_list_dealer";xml:"product_line_list_dealer"`
    
    /* remark required备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* requisition_id required合作申请Id */
    requisition_id int64 `json:"requisition_id";xml:"requisition_id"`
    
}

func (req *TaobaoFenxiaoCooperationAuditRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.audit"
}

/* TaobaoFenxiaoCooperationAuditResponse 合作授权审批 */
type TaobaoFenxiaoCooperationAuditResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoRequisitionsGetRequest 合作申请查询 */
type TaobaoFenxiaoRequisitionsGetRequest struct {
    
    /* apply_end optional申请结束时间yyyy-MM-dd */
    apply_end time.Time `json:"apply_end";xml:"apply_end"`
    
    /* apply_start optional申请开始时间yyyy-MM-dd */
    apply_start time.Time `json:"apply_start";xml:"apply_start"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* status optional申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期） */
    status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoFenxiaoRequisitionsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.requisitions.get"
}

/* TaobaoFenxiaoRequisitionsGetResponse 合作申请查询 */
type TaobaoFenxiaoRequisitionsGetResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* requisitions Object Array合作申请 */
    requisitions Requisition `json:"requisitions";xml:"requisitions"`
    
    /* total_results Basic结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoLogisticsAddressReachableRequest 根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通 */
type TaobaoLogisticsAddressReachableRequest struct {
    
    /* address optional详细地址 */
    address string `json:"address";xml:"address"`
    
    /* area_code optional标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105 */
    area_code string `json:"area_code";xml:"area_code"`
    
    /* partner_ids required物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953” */
    partner_ids string `json:"partner_ids";xml:"partner_ids"`
    
    /* service_type required服务对应的数字编码，如揽派范围对应88 */
    service_type int64 `json:"service_type";xml:"service_type"`
    
    /* source_area_code optional发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011 */
    source_area_code string `json:"source_area_code";xml:"source_area_code"`
    
}

func (req *TaobaoLogisticsAddressReachableRequest) GetAPIName() string {
	return "taobao.logistics.address.reachable"
}

/* TaobaoLogisticsAddressReachableResponse 根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通 */
type TaobaoLogisticsAddressReachableResponse struct {
    
    /* reachable_result_list Object Array地址可达返回结果，每个TP对应一个 */
    reachable_result_list AddressReachableResult `json:"reachable_result_list";xml:"reachable_result_list"`
    
}

/* TaobaoFenxiaoCooperationTerminateRequest 终止与分销商的合作关系 */
type TaobaoFenxiaoCooperationTerminateRequest struct {
    
    /* cooperate_id required合作编号 */
    cooperate_id int64 `json:"cooperate_id";xml:"cooperate_id"`
    
    /* end_remain_days required终止天数，可选1,2,3,5,7,15，在多少天数内终止 */
    end_remain_days int64 `json:"end_remain_days";xml:"end_remain_days"`
    
    /* end_remark required终止说明（5-2000字） */
    end_remark string `json:"end_remark";xml:"end_remark"`
    
}

func (req *TaobaoFenxiaoCooperationTerminateRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.terminate"
}

/* TaobaoFenxiaoCooperationTerminateResponse 终止与分销商的合作关系 */
type TaobaoFenxiaoCooperationTerminateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoProductGradepriceGetRequest 等级折扣查询 */
type TaobaoFenxiaoProductGradepriceGetRequest struct {
    
    /* product_id required产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* sku_id optionalskuId */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* trade_type optional经、代销模式（1：代销、2：经销） */
    trade_type int64 `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductGradepriceGetRequest) GetAPIName() string {
	return "taobao.fenxiao.product.gradeprice.get"
}

/* TaobaoFenxiaoProductGradepriceGetResponse 等级折扣查询 */
type TaobaoFenxiaoProductGradepriceGetResponse struct {
    
    /* grade_discounts Object Array等级折扣列表 */
    grade_discounts GradeDiscount `json:"grade_discounts";xml:"grade_discounts"`
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoOmniorderStorecollectConsumeRequest 全渠道门店自提核销订单 */
type TaobaoOmniorderStorecollectConsumeRequest struct {
    
    /* code required核销码 */
    code string `json:"code";xml:"code"`
    
    /* main_order_id required淘宝主订单ID */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
    /* operator optional核销操作人信息 */
    operator string `json:"operator";xml:"operator"`
    
}

func (req *TaobaoOmniorderStorecollectConsumeRequest) GetAPIName() string {
	return "taobao.omniorder.storecollect.consume"
}

/* TaobaoOmniorderStorecollectConsumeResponse 全渠道门店自提核销订单 */
type TaobaoOmniorderStorecollectConsumeResponse struct {
    
    /* err_code Basic0表示成功 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_msg Basic核销错误信息 */
    err_msg string `json:"err_msg";xml:"err_msg"`
    
}

/* TaobaoOmniorderStorecollectQueryRequest 全渠道门店自提根据核销码查询订单 */
type TaobaoOmniorderStorecollectQueryRequest struct {
    
    /* code required核销码 */
    code string `json:"code";xml:"code"`
    
}

func (req *TaobaoOmniorderStorecollectQueryRequest) GetAPIName() string {
	return "taobao.omniorder.storecollect.query"
}

/* TaobaoOmniorderStorecollectQueryResponse 全渠道门店自提根据核销码查询订单 */
type TaobaoOmniorderStorecollectQueryResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoQianniuKefuevalGetRequest 获取买家对客服的服务评价 */
type TaobaoQianniuKefuevalGetRequest struct {
    
    /* btime required开始时间，格式yyyyMMddHHmmss */
    btime string `json:"btime";xml:"btime"`
    
    /* etime required结束时间,格式yyyyMMddHHmmss */
    etime string `json:"etime";xml:"etime"`
    
    /* query_ids required客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick */
    query_ids string `json:"query_ids";xml:"query_ids"`
    
}

func (req *TaobaoQianniuKefuevalGetRequest) GetAPIName() string {
	return "taobao.qianniu.kefueval.get"
}

/* TaobaoQianniuKefuevalGetResponse 获取买家对客服的服务评价 */
type TaobaoQianniuKefuevalGetResponse struct {
    
    /* result_count BasicresultCount */
    result_count int64 `json:"result_count";xml:"result_count"`
    
    /* staff_eval_details Object ArraystaffEvalDetails */
    staff_eval_details EvalDetail `json:"staff_eval_details";xml:"staff_eval_details"`
    
}

/* CainiaoSmartdeliveryStrategyWarehouseIQueryRequest 智能发货引擎仓维度策略查询 */
type CainiaoSmartdeliveryStrategyWarehouseIQueryRequest struct {
    
    /* param_query_delivery_strategy_request optional查询请求参数 */
    param_query_delivery_strategy_request QueryDeliveryStrategyRequest `json:"param_query_delivery_strategy_request";xml:"param_query_delivery_strategy_request"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.query"
}

/* CainiaoSmartdeliveryStrategyWarehouseIQueryResponse 智能发货引擎仓维度策略查询 */
type CainiaoSmartdeliveryStrategyWarehouseIQueryResponse struct {
    
    /* delivery_strategy_info_list Object Array返回结果列表 */
    delivery_strategy_info_list DeliveryStrategyInfo `json:"delivery_strategy_info_list";xml:"delivery_strategy_info_list"`
    
}

/* CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest 智能发货引擎发货策略设置仓维度 */
type CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest struct {
    
    /* delivery_strategy_set_request required智能发货设置请求参数 */
    delivery_strategy_set_request DeliveryStrategySetRequest `json:"delivery_strategy_set_request";xml:"delivery_strategy_set_request"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

/* CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse 智能发货引擎发货策略设置仓维度 */
type CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse struct {
    
    /* warehouse_info Object仓信息 */
    warehouse_info WarehouseDto `json:"warehouse_info";xml:"warehouse_info"`
    
}

/* TaobaoFenxiaoProductImportFromAuctionRequest 供应商选择关联店铺的前台宝贝，导入生成产品 */
type TaobaoFenxiaoProductImportFromAuctionRequest struct {
    
    /* auction_id required店铺宝贝id */
    auction_id int64 `json:"auction_id";xml:"auction_id"`
    
    /* product_line_id required产品线id */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
    /* trade_type required导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销] */
    trade_type int64 `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductImportFromAuctionRequest) GetAPIName() string {
	return "taobao.fenxiao.product.import.from.auction"
}

/* TaobaoFenxiaoProductImportFromAuctionResponse 供应商选择关联店铺的前台宝贝，导入生成产品 */
type TaobaoFenxiaoProductImportFromAuctionResponse struct {
    
    /* opt_time Basic操作时间 */
    opt_time time.Time `json:"opt_time";xml:"opt_time"`
    
    /* pid Basic生成的产品id */
    pid int64 `json:"pid";xml:"pid"`
    
}

/* TaobaoScitemOutercodeGetRequest 根据outerCode查询商品 */
type TaobaoScitemOutercodeGetRequest struct {
    
    /* outer_code required商品编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
}

func (req *TaobaoScitemOutercodeGetRequest) GetAPIName() string {
	return "taobao.scitem.outercode.get"
}

/* TaobaoScitemOutercodeGetResponse 根据outerCode查询商品 */
type TaobaoScitemOutercodeGetResponse struct {
    
    /* sc_item Object后台商品 */
    sc_item ScItem `json:"sc_item";xml:"sc_item"`
    
}

/* TaobaoHttpdnsGetRequest 获取TOP DNS配置 */
type TaobaoHttpdnsGetRequest struct {
    
}

func (req *TaobaoHttpdnsGetRequest) GetAPIName() string {
	return "taobao.httpdns.get"
}

/* TaobaoHttpdnsGetResponse 获取TOP DNS配置 */
type TaobaoHttpdnsGetResponse struct {
    
    /* result BasicHTTP DNS配置信息 */
    result string `json:"result";xml:"result"`
    
}

/* TaobaoScitemGetRequest 根据id查询商品 */
type TaobaoScitemGetRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoScitemGetRequest) GetAPIName() string {
	return "taobao.scitem.get"
}

/* TaobaoScitemGetResponse 根据id查询商品 */
type TaobaoScitemGetResponse struct {
    
    /* sc_item Object后端商品 */
    sc_item ScItem `json:"sc_item";xml:"sc_item"`
    
}

/* TaobaoScitemQueryRequest 查询后端商品 */
type TaobaoScitemQueryRequest struct {
    
    /* bar_code optional条形码 */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* item_name optional商品名称 */
    item_name string `json:"item_name";xml:"item_name"`
    
    /* item_type optionalITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION */
    item_type int64 `json:"item_type";xml:"item_type"`
    
    /* outer_code optional商家给商品的一个编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* page_index optional当前页码数 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* wms_code optional仓库编码 */
    wms_code string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemQueryRequest) GetAPIName() string {
	return "taobao.scitem.query"
}

/* TaobaoScitemQueryResponse 查询后端商品 */
type TaobaoScitemQueryResponse struct {
    
    /* query_pagination Object分页 */
    query_pagination QueryPagination `json:"query_pagination";xml:"query_pagination"`
    
    /* sc_item_list Object ArrayList<ScItemDO> */
    sc_item_list ScItem `json:"sc_item_list";xml:"sc_item_list"`
    
    /* total_page Basic商品条数 */
    total_page int64 `json:"total_page";xml:"total_page"`
    
}

/* TaobaoScitemUpdateRequest 根据商品ID或商家编码修改后端商品 */
type TaobaoScitemUpdateRequest struct {
    
    /* bar_code optional条形码 */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* brand_id optional品牌id */
    brand_id int64 `json:"brand_id";xml:"brand_id"`
    
    /* brand_name optionalbrand_Name */
    brand_name string `json:"brand_name";xml:"brand_name"`
    
    /* height optional高 单位：mm */
    height int64 `json:"height";xml:"height"`
    
    /* is_area_sale optional1表示区域销售，0或是空是非区域销售 */
    is_area_sale int64 `json:"is_area_sale";xml:"is_area_sale"`
    
    /* is_costly optional是否是贵重品 0:不是 1：是 */
    is_costly int64 `json:"is_costly";xml:"is_costly"`
    
    /* is_dangerous optional是否危险 0：不是  0：是 */
    is_dangerous int64 `json:"is_dangerous";xml:"is_dangerous"`
    
    /* is_friable optional是否易碎 0：不是  1：是 */
    is_friable int64 `json:"is_friable";xml:"is_friable"`
    
    /* is_warranty optional是否保质期：0:不是 1：是 */
    is_warranty int64 `json:"is_warranty";xml:"is_warranty"`
    
    /* item_id optional后端商品ID，跟outer_code必须指定一个 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* item_name optional商品名称 */
    item_name string `json:"item_name";xml:"item_name"`
    
    /* item_type optional0.普通供应链商品 1.供应链组合主商品 */
    item_type int64 `json:"item_type";xml:"item_type"`
    
    /* length optional长度 单位：mm */
    length int64 `json:"length";xml:"length"`
    
    /* matter_status optional0:液体，1：粉体，2：固体 */
    matter_status int64 `json:"matter_status";xml:"matter_status"`
    
    /* outer_code special商家编码，跟item_id必须指定一个 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* price optionalprice */
    price int64 `json:"price";xml:"price"`
    
    /* remark optionalremark */
    remark string `json:"remark";xml:"remark"`
    
    /* remove_properties optional移除商品属性P列表,P由系统分配：p1；p2 */
    remove_properties string `json:"remove_properties";xml:"remove_properties"`
    
    /* spu_id optional淘宝SKU产品级编码CSPU ID */
    spu_id int64 `json:"spu_id";xml:"spu_id"`
    
    /* update_properties optional需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3 */
    update_properties string `json:"update_properties";xml:"update_properties"`
    
    /* volume optional体积：立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* weight optionalweight */
    weight int64 `json:"weight";xml:"weight"`
    
    /* width optional宽 单位：mm */
    width int64 `json:"width";xml:"width"`
    
    /* wms_code optional仓储商编码 */
    wms_code string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemUpdateRequest) GetAPIName() string {
	return "taobao.scitem.update"
}

/* TaobaoScitemUpdateResponse 根据商品ID或商家编码修改后端商品 */
type TaobaoScitemUpdateResponse struct {
    
    /* update_rows Basic更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据 */
    update_rows int64 `json:"update_rows";xml:"update_rows"`
    
}

/* TaobaoScitemAddRequest 发布后端商品 */
type TaobaoScitemAddRequest struct {
    
    /* bar_code optional条形码 */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* brand_id optional品牌id */
    brand_id int64 `json:"brand_id";xml:"brand_id"`
    
    /* brand_name optionalbrand_Name */
    brand_name string `json:"brand_name";xml:"brand_name"`
    
    /* height optional高 单位：mm */
    height int64 `json:"height";xml:"height"`
    
    /* is_area_sale optional1表示区域销售，0或是空是非区域销售 */
    is_area_sale int64 `json:"is_area_sale";xml:"is_area_sale"`
    
    /* is_costly optional是否是贵重品 0:不是 1：是 */
    is_costly int64 `json:"is_costly";xml:"is_costly"`
    
    /* is_dangerous optional是否危险 0：不是  1：是 */
    is_dangerous int64 `json:"is_dangerous";xml:"is_dangerous"`
    
    /* is_friable optional是否易碎 0：不是  1：是 */
    is_friable int64 `json:"is_friable";xml:"is_friable"`
    
    /* is_warranty optional是否保质期：0:不是 1：是 */
    is_warranty int64 `json:"is_warranty";xml:"is_warranty"`
    
    /* item_name required商品名称 */
    item_name string `json:"item_name";xml:"item_name"`
    
    /* item_type optional0.普通供应链商品 1.供应链组合主商品 */
    item_type int64 `json:"item_type";xml:"item_type"`
    
    /* length optional长度 单位：mm */
    length int64 `json:"length";xml:"length"`
    
    /* matter_status optional0:液体，1：粉体，2：固体 */
    matter_status int64 `json:"matter_status";xml:"matter_status"`
    
    /* outer_code required商家编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* price optional价格 单位：分 */
    price int64 `json:"price";xml:"price"`
    
    /* properties optional商品属性格式是  p1:v1,p2:v2,p3:v3 */
    properties string `json:"properties";xml:"properties"`
    
    /* remark optionalremark */
    remark string `json:"remark";xml:"remark"`
    
    /* spu_id optionalspuId或是cspuid */
    spu_id int64 `json:"spu_id";xml:"spu_id"`
    
    /* volume optional体积：立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* weight optional重量 单位：g */
    weight int64 `json:"weight";xml:"weight"`
    
    /* width optional宽 单位：mm */
    width int64 `json:"width";xml:"width"`
    
    /* wms_code optional仓储商编码 */
    wms_code string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemAddRequest) GetAPIName() string {
	return "taobao.scitem.add"
}

/* TaobaoScitemAddResponse 发布后端商品 */
type TaobaoScitemAddResponse struct {
    
    /* sc_item Object后台商品信息 */
    sc_item ScItem `json:"sc_item";xml:"sc_item"`
    
}

/* TaobaoScitemMapDeleteRequest 根据后端商品Id，失效指定用户的商品与后端商品的映射关系 */
type TaobaoScitemMapDeleteRequest struct {
    
    /* sc_item_id required后台商品ID */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* user_nick optional店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoScitemMapDeleteRequest) GetAPIName() string {
	return "taobao.scitem.map.delete"
}

/* TaobaoScitemMapDeleteResponse 根据后端商品Id，失效指定用户的商品与后端商品的映射关系 */
type TaobaoScitemMapDeleteResponse struct {
    
    /* module Basic失效条数 */
    module int64 `json:"module";xml:"module"`
    
}

/* TaobaoScitemMapQueryRequest 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku */
type TaobaoScitemMapQueryRequest struct {
    
    /* item_id requiredmap_type为1：前台ic商品id
map_type为2：分销productid */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sku_id optionalmap_type为1：前台ic商品skuid 
map_type为2：分销商品的skuid */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoScitemMapQueryRequest) GetAPIName() string {
	return "taobao.scitem.map.query"
}

/* TaobaoScitemMapQueryResponse 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku */
type TaobaoScitemMapQueryResponse struct {
    
    /* sc_item_maps Object Array后端商品映射列表 */
    sc_item_maps ScItemMap `json:"sc_item_maps";xml:"sc_item_maps"`
    
}

/* AlipayUserAccountFreezeGetRequest 查询支付宝账户冻结类型的冻结金额。 */
type AlipayUserAccountFreezeGetRequest struct {
    
    /* freeze_type optional冻结类型，多个用,分隔。不传返回所有类型的冻结金额。
DEPOSIT_FREEZE,充值冻结
WITHDRAW_FREEZE,提现冻结
PAYMENT_FREEZE,支付冻结
BAIL_FREEZE,保证金冻结
CHARGE_FREEZE,收费冻结
PRE_DEPOSIT_FREEZE,预存款冻结
LOAN_FREEZE,贷款冻结
OTHER_FREEZE,其它 */
    freeze_type string `json:"freeze_type";xml:"freeze_type"`
    
}

func (req *AlipayUserAccountFreezeGetRequest) GetAPIName() string {
	return "alipay.user.account.freeze.get"
}

/* AlipayUserAccountFreezeGetResponse 查询支付宝账户冻结类型的冻结金额。 */
type AlipayUserAccountFreezeGetResponse struct {
    
    /* freeze_items Object Array冻结金额列表 */
    freeze_items AccountFreeze `json:"freeze_items";xml:"freeze_items"`
    
    /* total_results Basic冻结总条数 */
    total_results string `json:"total_results";xml:"total_results"`
    
}

/* AlipayUserAccountGetRequest 查询支付宝账户余额 */
type AlipayUserAccountGetRequest struct {
    
}

func (req *AlipayUserAccountGetRequest) GetAPIName() string {
	return "alipay.user.account.get"
}

/* AlipayUserAccountGetResponse 查询支付宝账户余额 */
type AlipayUserAccountGetResponse struct {
    
    /* alipay_account Object支付宝用户账户信息 */
    alipay_account AlipayAccount `json:"alipay_account";xml:"alipay_account"`
    
}

/* TaobaoRdcAligeniusSendgoodsCancelRequest 提供商家在仅退款中发送取消发货状态 */
type TaobaoRdcAligeniusSendgoodsCancelRequest struct {
    
    /* param required请求参数 */
    param CancelGoodsDto `json:"param";xml:"param"`
    
}

func (req *TaobaoRdcAligeniusSendgoodsCancelRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.sendgoods.cancel"
}

/* TaobaoRdcAligeniusSendgoodsCancelResponse 提供商家在仅退款中发送取消发货状态 */
type TaobaoRdcAligeniusSendgoodsCancelResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoProductMapAddRequest 创建分销和供应链商品映射关系。 */
type TaobaoFenxiaoProductMapAddRequest struct {
    
    /* not_check_outer_code optional是否需要校验商家编码，true不校验，false校验。 */
    not_check_outer_code bool `json:"not_check_outer_code";xml:"not_check_outer_code"`
    
    /* product_id required分销产品id。 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* sc_item_id optional后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。 */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* sc_item_ids optional在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。 */
    sc_item_ids string `json:"sc_item_ids";xml:"sc_item_ids"`
    
    /* sku_ids optional分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。 */
    sku_ids string `json:"sku_ids";xml:"sku_ids"`
    
}

func (req *TaobaoFenxiaoProductMapAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.map.add"
}

/* TaobaoFenxiaoProductMapAddResponse 创建分销和供应链商品映射关系。 */
type TaobaoFenxiaoProductMapAddResponse struct {
    
    /* is_success Basic操作结果 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* AlipayUserTradeSearchRequest 查询支付宝账户交易记录 */
type AlipayUserTradeSearchRequest struct {
    
    /* alipay_order_no optional支付宝订单号，为空查询所有记录 */
    alipay_order_no string `json:"alipay_order_no";xml:"alipay_order_no"`
    
    /* end_time required结束时间。与开始时间间隔在七天之内 */
    end_time string `json:"end_time";xml:"end_time"`
    
    /* merchant_order_no optional商户订单号，为空查询所有记录 */
    merchant_order_no string `json:"merchant_order_no";xml:"merchant_order_no"`
    
    /* order_from optional订单来源，为空查询所有来源。淘宝(TAOBAO)，支付宝(ALIPAY)，其它(OTHER) */
    order_from string `json:"order_from";xml:"order_from"`
    
    /* order_status optional订单状态，为空查询所有状态订单 */
    order_status string `json:"order_status";xml:"order_status"`
    
    /* order_type optional订单类型，为空查询所有类型订单。 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* page_no required页码。取值范围:大于零的整数; 默认值1 */
    page_no string `json:"page_no";xml:"page_no"`
    
    /* page_size required每页获取条数。最大值500。 */
    page_size string `json:"page_size";xml:"page_size"`
    
    /* start_time required开始时间，时间必须是今天范围之内。格式为yyyy-MM-dd HH:mm:ss，精确到秒 */
    start_time string `json:"start_time";xml:"start_time"`
    
}

func (req *AlipayUserTradeSearchRequest) GetAPIName() string {
	return "alipay.user.trade.search"
}

/* AlipayUserTradeSearchResponse 查询支付宝账户交易记录 */
type AlipayUserTradeSearchResponse struct {
    
    /* total_pages Basic总页数 */
    total_pages string `json:"total_pages";xml:"total_pages"`
    
    /* total_results Basic总记录数 */
    total_results string `json:"total_results";xml:"total_results"`
    
    /* trade_records Object Array交易记录列表 */
    trade_records TradeRecord `json:"trade_records";xml:"trade_records"`
    
}

/* CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest 删除智能发货引擎仓策略 */
type CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest struct {
    
    /* warehouse_id optional仓id */
    warehouse_id int64 `json:"warehouse_id";xml:"warehouse_id"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

/* CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse 删除智能发货引擎仓策略 */
type CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse struct {
    
    /* is_delete_success Basicdata */
    is_delete_success bool `json:"is_delete_success";xml:"is_delete_success"`
    
}

/* TmallItemShiptimeUpdateRequest 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
1.
    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 0 ---更新SKU
    },

    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

2.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 0 --更新SKU
    },
    按照指定SKU删除指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

3.

    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 1 ---更新商品
    },

    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间

4.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 1 --更新商品
    },
    删除商品级的发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间 */
type TmallItemShiptimeUpdateRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* option required批量更新商品/SKU发货时间的备选项 */
    option UpdateItemShipTimeOption `json:"option";xml:"option"`
    
    /* ship_time optional被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。 */
    ship_time string `json:"ship_time";xml:"ship_time"`
    
    /* sku_ship_times optional被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。 */
    sku_ship_times UpdateSkuShipTime `json:"sku_ship_times";xml:"sku_ship_times"`
    
}

func (req *TmallItemShiptimeUpdateRequest) GetAPIName() string {
	return "tmall.item.shiptime.update"
}

/* TmallItemShiptimeUpdateResponse 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
1.
    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 0 ---更新SKU
    },

    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

2.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 0 --更新SKU
    },
    按照指定SKU删除指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

3.

    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 1 ---更新商品
    },

    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间

4.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 1 --更新商品
    },
    删除商品级的发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间 */
type TmallItemShiptimeUpdateResponse struct {
    
    /* shiptime_update_result Basic被更新商品ID */
    shiptime_update_result string `json:"shiptime_update_result";xml:"shiptime_update_result"`
    
}

/* TmallInventoryQueryForstoreRequest 商家查询后端商品仓库库存 */
type TmallInventoryQueryForstoreRequest struct {
    
    /* param_list required查询列表 */
    param_list InventoryQueryForStoreRequest `json:"param_list";xml:"param_list"`
    
}

func (req *TmallInventoryQueryForstoreRequest) GetAPIName() string {
	return "tmall.inventory.query.forstore"
}

/* TmallInventoryQueryForstoreResponse 商家查询后端商品仓库库存 */
type TmallInventoryQueryForstoreResponse struct {
    
    /* errorcode Basic错误code */
    errorcode string `json:"errorcode";xml:"errorcode"`
    
    /* errormessage Basic错误信息 */
    errormessage string `json:"errormessage";xml:"errormessage"`
    
    /* issuccess Basic整体成功或失败 */
    issuccess bool `json:"issuccess";xml:"issuccess"`
    
    /* result Object查询结果 */
    result InventoryQueryResult `json:"result";xml:"result"`
    
}

/* TaobaoTopAuthTokenRefreshRequest 根据refresh_token重新生成token */
type TaobaoTopAuthTokenRefreshRequest struct {
    
    /* refresh_token requiredgrantType==refresh_token 时需要 */
    refresh_token string `json:"refresh_token";xml:"refresh_token"`
    
}

func (req *TaobaoTopAuthTokenRefreshRequest) GetAPIName() string {
	return "taobao.top.auth.token.refresh"
}

/* TaobaoTopAuthTokenRefreshResponse 根据refresh_token重新生成token */
type TaobaoTopAuthTokenRefreshResponse struct {
    
    /* token_result Basic返回的是json信息 */
    token_result map[string]interface{} `json:"token_result";xml:"token_result"`
    
}

/* TaobaoTopAuthTokenCreateRequest 用户通过code换获取access_token，https only */
type TaobaoTopAuthTokenCreateRequest struct {
    
    /* code required授权code，grantType==authorization_code 时需要 */
    code string `json:"code";xml:"code"`
    
    /* uuid optional与生成code的uuid配对 */
    uuid string `json:"uuid";xml:"uuid"`
    
}

func (req *TaobaoTopAuthTokenCreateRequest) GetAPIName() string {
	return "taobao.top.auth.token.create"
}

/* TaobaoTopAuthTokenCreateResponse 用户通过code换获取access_token，https only */
type TaobaoTopAuthTokenCreateResponse struct {
    
    /* token_result Basic返回的是json信息，和之前调用https://oauth.taobao.com/tac/token https://oauth.alibaba.com/token 换token返回的字段信息一致 */
    token_result map[string]interface{} `json:"token_result";xml:"token_result"`
    
}

/* AlibabaEinvoiceClosereqRequest 关闭失败开票请求，避免造成重复开票 */
type AlibabaEinvoiceClosereqRequest struct {
    
    /* payee_register_no required税号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* serial_no required流水号 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoiceClosereqRequest) GetAPIName() string {
	return "alibaba.einvoice.closereq"
}

/* AlibabaEinvoiceClosereqResponse 关闭失败开票请求，避免造成重复开票 */
type AlibabaEinvoiceClosereqResponse struct {
    
    /* result Basic关闭是否成功 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoQimenOrderCallbackRequest 配送拦截 */
type TaobaoQimenOrderCallbackRequest struct {
    
    /* deliveryOrderCode optional奇门仓储字段,C123,string(50),, */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* expressCode optional运单号 */
    expressCode string `json:"expressCode";xml:"expressCode"`
    
    /* orderId optional奇门仓储字段,C123,string(50),, */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* ownerCode optional奇门仓储字段,C123,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* warehouseCode optional奇门仓储字段,C123,string(50),, */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderCallbackRequest) GetAPIName() string {
	return "taobao.qimen.order.callback"
}

/* TaobaoQimenOrderCallbackResponse 配送拦截 */
type TaobaoQimenOrderCallbackResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenWarehouseinfoQueryRequest 货主仓库资源查询 */
type TaobaoQimenWarehouseinfoQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenWarehouseinfoQueryRequest) GetAPIName() string {
	return "taobao.qimen.warehouseinfo.query"
}

/* TaobaoQimenWarehouseinfoQueryResponse 货主仓库资源查询 */
type TaobaoQimenWarehouseinfoQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* ownerCode Basic奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* ownerName Basic奇门仓储字段 */
    ownerName string `json:"ownerName";xml:"ownerName"`
    
    /* warehouseInfos Object Array奇门仓储字段 */
    warehouseInfos WarehouseInfo `json:"warehouseInfos";xml:"warehouseInfos"`
    
}

/* TaobaoQimenExpressinfoQueryRequest 配送公司信息查询 */
type TaobaoQimenExpressinfoQueryRequest struct {
    
    /* expressCode optional奇门仓储字段 */
    expressCode string `json:"expressCode";xml:"expressCode"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenExpressinfoQueryRequest) GetAPIName() string {
	return "taobao.qimen.expressinfo.query"
}

/* TaobaoQimenExpressinfoQueryResponse 配送公司信息查询 */
type TaobaoQimenExpressinfoQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* expressInfos Object Array奇门仓储字段 */
    expressInfos ExpressInfo `json:"expressInfos";xml:"expressInfos"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoWlbWmsSnInfoQueryRequest 查询仓库作业的各类单据记录的Sn信息 */
type TaobaoWlbWmsSnInfoQueryRequest struct {
    
    /* order_code required订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_code_type required订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ） */
    order_code_type int64 `json:"order_code_type";xml:"order_code_type"`
    
    /* page_index required页数，默认每页50条 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
}

func (req *TaobaoWlbWmsSnInfoQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.sn.info.query"
}

/* TaobaoWlbWmsSnInfoQueryResponse 查询仓库作业的各类单据记录的Sn信息 */
type TaobaoWlbWmsSnInfoQueryResponse struct {
    
    /* result Object接口返回 */
    result Result `json:"result";xml:"result"`
    
}

/* CainiaoCloudprintCustomareaUpdateRequest 自定义区内容更新 */
type CainiaoCloudprintCustomareaUpdateRequest struct {
    
    /* custom_area_content required自定义区内容（可修改） */
    custom_area_content string `json:"custom_area_content";xml:"custom_area_content"`
    
    /* custom_area_id required自定义区id（不可修改） */
    custom_area_id int64 `json:"custom_area_id";xml:"custom_area_id"`
    
    /* custom_area_name required自定义区名称（可修改） */
    custom_area_name string `json:"custom_area_name";xml:"custom_area_name"`
    
}

func (req *CainiaoCloudprintCustomareaUpdateRequest) GetAPIName() string {
	return "cainiao.cloudprint.customarea.update"
}

/* CainiaoCloudprintCustomareaUpdateResponse 自定义区内容更新 */
type CainiaoCloudprintCustomareaUpdateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* TaobaoQimenTransferorderQueryRequest 调拨单查询 */
type TaobaoQimenTransferorderQueryRequest struct {
    
    /* erpOrderCode optionalERP单号 */
    erpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    /* ownerCode required111 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* transferOrderCode required调拨单号 */
    transferOrderCode string `json:"transferOrderCode";xml:"transferOrderCode"`
    
}

func (req *TaobaoQimenTransferorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.query"
}

/* TaobaoQimenTransferorderQueryResponse 调拨单查询 */
type TaobaoQimenTransferorderQueryResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
    /* transferOrderDetail Object调拨单细节 */
    transferOrderDetail TransferOrderDetail `json:"transferOrderDetail";xml:"transferOrderDetail"`
    
}

/* TaobaoQimenTransferorderReportRequest 调拨单通知 */
type TaobaoQimenTransferorderReportRequest struct {
    
    /* confirmInTime optional确认入库时间 */
    confirmInTime string `json:"confirmInTime";xml:"confirmInTime"`
    
    /* confirmOutTime optional确认出库时间 */
    confirmOutTime string `json:"confirmOutTime";xml:"confirmOutTime"`
    
    /* createTime optional调拨单创建时间 */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* erpOrderCode optionalerpOrderCode */
    erpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    /* fromWarehouseCode optional调拨出库仓编码 */
    fromWarehouseCode string `json:"fromWarehouseCode";xml:"fromWarehouseCode"`
    
    /* items optional项目集 */
    items Items `json:"items";xml:"items"`
    
    /* orderStatus optionalorderStatus */
    orderStatus string `json:"orderStatus";xml:"orderStatus"`
    
    /* ownerCode optional111 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* toWarehouseCode optional调拨入库仓编码 */
    toWarehouseCode string `json:"toWarehouseCode";xml:"toWarehouseCode"`
    
    /* transferInOrderCode optional调拨入库单号 */
    transferInOrderCode string `json:"transferInOrderCode";xml:"transferInOrderCode"`
    
    /* transferOrderCode optional调拨单号,0,string(50),必填, */
    transferOrderCode string `json:"transferOrderCode";xml:"transferOrderCode"`
    
    /* transferOutOrderCode optional调拨出库单号 */
    transferOutOrderCode string `json:"transferOutOrderCode";xml:"transferOutOrderCode"`
    
}

func (req *TaobaoQimenTransferorderReportRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.report"
}

/* TaobaoQimenTransferorderReportResponse 调拨单通知 */
type TaobaoQimenTransferorderReportResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
}

/* TmallChannelProductsQueryRequest 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息 */
type TmallChannelProductsQueryRequest struct {
    
    /* page_num required页码数，从1开始 */
    page_num int64 `json:"page_num";xml:"page_num"`
    
    /* page_size required分页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* product_ids optional产品Id */
    product_ids int64 `json:"product_ids";xml:"product_ids"`
    
    /* product_line_id optional产品线Id */
    product_line_id int64 `json:"product_line_id";xml:"product_line_id"`
    
    /* product_number optional商家产品编码 */
    product_number string `json:"product_number";xml:"product_number"`
    
    /* sku_number optional商家SKU编码 */
    sku_number string `json:"sku_number";xml:"sku_number"`
    
}

func (req *TmallChannelProductsQueryRequest) GetAPIName() string {
	return "tmall.channel.products.query"
}

/* TmallChannelProductsQueryResponse 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息 */
type TmallChannelProductsQueryResponse struct {
    
    /* result Objectresult */
    result PageResultDto `json:"result";xml:"result"`
    
}

/* TaobaoQimenTransferorderCreateRequest 调拨单创建 */
type TaobaoQimenTransferorderCreateRequest struct {
    
    /* attributes optional扩展属性,HZ1234,string(500),, */
    attributes string `json:"attributes";xml:"attributes"`
    
    /* erpOrderCode optional外部ERP订单号,HZ1234,string(50),必填, */
    erpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    /* expectStartTime optional期望调拨开始时间,Item1234,string(50),, */
    expectStartTime string `json:"expectStartTime";xml:"expectStartTime"`
    
    /* fromStoreCode optional出库仓编码,Item1234,string(50),必填, */
    fromStoreCode string `json:"fromStoreCode";xml:"fromStoreCode"`
    
    /* ownerCode required111 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* toStoreCode optional入库仓编码,HZ1234,string(50),必填, */
    toStoreCode string `json:"toStoreCode";xml:"toStoreCode"`
    
    /* transferItems optional项目集 */
    transferItems TransferItems `json:"transferItems";xml:"transferItems"`
    
}

func (req *TaobaoQimenTransferorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.create"
}

/* TaobaoQimenTransferorderCreateResponse 调拨单创建 */
type TaobaoQimenTransferorderCreateResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
    /* transferExecuteInfo Object调拨单信息 */
    transferExecuteInfo TransferExecuteInfo `json:"transferExecuteInfo";xml:"transferExecuteInfo"`
    
}

/* CainiaoOpenstorageSellerResourcedetailGetRequest 商家资源详情获取（云打印开源存储） */
type CainiaoOpenstorageSellerResourcedetailGetRequest struct {
    
    /* seller_resource_id optional商家资源id */
    seller_resource_id int64 `json:"seller_resource_id";xml:"seller_resource_id"`
    
}

func (req *CainiaoOpenstorageSellerResourcedetailGetRequest) GetAPIName() string {
	return "cainiao.openstorage.seller.resourcedetail.get"
}

/* CainiaoOpenstorageSellerResourcedetailGetResponse 商家资源详情获取（云打印开源存储） */
type CainiaoOpenstorageSellerResourcedetailGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* CainiaoOpenstorageIsvResourcedetailGetRequest isv资源详情查询（云打印开源存储） */
type CainiaoOpenstorageIsvResourcedetailGetRequest struct {
    
    /* isv_resource_id optionalisv资源id */
    isv_resource_id int64 `json:"isv_resource_id";xml:"isv_resource_id"`
    
}

func (req *CainiaoOpenstorageIsvResourcedetailGetRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resourcedetail.get"
}

/* CainiaoOpenstorageIsvResourcedetailGetResponse isv资源详情查询（云打印开源存储） */
type CainiaoOpenstorageIsvResourcedetailGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* QimenTaobaoIcpOrderStockinordermessagetoerpRequest 供应链创建的入库单推送商家ERP */
type QimenTaobaoIcpOrderStockinordermessagetoerpRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOrderlist optional入库单记录集 */
    entryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockinordermessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockinordermessagetoerp"
}

/* QimenTaobaoIcpOrderStockinordermessagetoerpResponse 供应链创建的入库单推送商家ERP */
type QimenTaobaoIcpOrderStockinordermessagetoerpResponse struct {
    
    /* result Object返回结果 */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoLogisticsTraceSearchRequest 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。
此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。） */
type TaobaoLogisticsTraceSearchRequest struct {
    
    /* is_split optional表明是否是拆单，默认值0，1表示拆单 */
    is_split int64 `json:"is_split";xml:"is_split"`
    
    /* seller_nick required卖家昵称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
    /* sub_tid optional拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。 */
    sub_tid int64 `json:"sub_tid";xml:"sub_tid"`
    
    /* tid required淘宝交易号，请勿传非淘宝交易号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsTraceSearchRequest) GetAPIName() string {
	return "taobao.logistics.trace.search"
}

/* TaobaoLogisticsTraceSearchResponse 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。
此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。） */
type TaobaoLogisticsTraceSearchResponse struct {
    
    /* company_name Basic物流公司名称 */
    company_name string `json:"company_name";xml:"company_name"`
    
    /* out_sid Basic运单号 */
    out_sid string `json:"out_sid";xml:"out_sid"`
    
    /* status Basic订单的物流状态（仅支持线上发货online订单，线下发货offline发出后直接变为已签收）
* 等候发送给物流公司
*已提交给物流公司,等待物流公司接单
*已经确认消息接收，等待物流公司接单
*物流公司已接单
*物流公司不接单
*物流公司揽收失败
*物流公司揽收成功
*签收失败
*对方已签收
*对方拒绝签收 */
    status string `json:"status";xml:"status"`
    
    /* tid Basic交易号 */
    tid int64 `json:"tid";xml:"tid"`
    
    /* trace_list Object Array流转信息列表 */
    trace_list TransitStepInfo `json:"trace_list";xml:"trace_list"`
    
}

/* TaobaoRdcAligeniusDistributionLogisticsCancelRequest 截配状态回传接口 */
type TaobaoRdcAligeniusDistributionLogisticsCancelRequest struct {
    
    /* param0 optional参数 */
    param0 CancelDistributionDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusDistributionLogisticsCancelRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.distribution.logistics.cancel"
}

/* TaobaoRdcAligeniusDistributionLogisticsCancelResponse 截配状态回传接口 */
type TaobaoRdcAligeniusDistributionLogisticsCancelResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* QimenTaobaoIcpOrderStockinordecanceltoerpRequest 一盘货入库单取消消息通知 */
type QimenTaobaoIcpOrderStockinordecanceltoerpRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOrderlist optional订单列表 */
    entryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockinordecanceltoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockinordecanceltoerp"
}

/* QimenTaobaoIcpOrderStockinordecanceltoerpResponse 一盘货入库单取消消息通知 */
type QimenTaobaoIcpOrderStockinordecanceltoerpResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}

/* CainiaoEndpointLockerTopStationAddorupdateRequest 新增或者修改代收点相关信息 */
type CainiaoEndpointLockerTopStationAddorupdateRequest struct {
    
    /* station_info optional站点信息 */
    station_info StationInfo `json:"station_info";xml:"station_info"`
    
}

func (req *CainiaoEndpointLockerTopStationAddorupdateRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.station.addorupdate"
}

/* CainiaoEndpointLockerTopStationAddorupdateResponse 新增或者修改代收点相关信息 */
type CainiaoEndpointLockerTopStationAddorupdateResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TaobaoRdcAligeniusWarehouseReverseUpdateRequest 销退单状态回传 */
type TaobaoRdcAligeniusWarehouseReverseUpdateRequest struct {
    
    /* param0 required参数 */
    param0 UpdateReverseStatusDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseReverseUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.update"
}

/* TaobaoRdcAligeniusWarehouseReverseUpdateResponse 销退单状态回传 */
type TaobaoRdcAligeniusWarehouseReverseUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* AlibabaProviderEinvoiceCreateRequest 电子发票同步开票 */
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
    invoice_time time.Time `json:"invoice_time";xml:"invoice_time"`
    
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

/* AlibabaProviderEinvoiceCreateResponse 电子发票同步开票 */
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

/* TaobaoRdcAligeniusRefundsCheckRequest 根据退款信息，对退款单进行审核 */
type TaobaoRdcAligeniusRefundsCheckRequest struct {
    
    /* param required入参 */
    param RefundCheckDto `json:"param";xml:"param"`
    
}

func (req *TaobaoRdcAligeniusRefundsCheckRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.refunds.check"
}

/* TaobaoRdcAligeniusRefundsCheckResponse 根据退款信息，对退款单进行审核 */
type TaobaoRdcAligeniusRefundsCheckResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoRpReturngoodsAgreeRequest 卖家同意退货，支持淘宝和天猫的订单。 */
type TaobaoRpReturngoodsAgreeRequest struct {
    
    /* address optional卖家提供的退货地址，淘宝退款为必填项。 */
    address string `json:"address";xml:"address"`
    
    /* mobile optional卖家手机，淘宝退款为必填项。 */
    mobile string `json:"mobile";xml:"mobile"`
    
    /* name optional卖家姓名，淘宝退款为必填项。 */
    name string `json:"name";xml:"name"`
    
    /* post optional卖家提供的退货地址的邮编，淘宝退款为必填项。 */
    post string `json:"post";xml:"post"`
    
    /* post_fee_bear_role optional邮费承担方，买家承担值为1，卖家承担值为0 */
    post_fee_bear_role int64 `json:"post_fee_bear_role";xml:"post_fee_bear_role"`
    
    /* refund_id required退款编号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase optional售中：onsale，售后：aftersale，天猫退款为必填项。 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version optional退款版本号，天猫退款为必填项。 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* remark optional卖家退货留言，天猫退款为必填项。 */
    remark string `json:"remark";xml:"remark"`
    
    /* seller_address_id optional卖家收货地址编号，天猫淘宝退款都为必填项。 */
    seller_address_id int64 `json:"seller_address_id";xml:"seller_address_id"`
    
    /* tel optional卖家座机，淘宝退款为必填项。 */
    tel string `json:"tel";xml:"tel"`
    
}

func (req *TaobaoRpReturngoodsAgreeRequest) GetAPIName() string {
	return "taobao.rp.returngoods.agree"
}

/* TaobaoRpReturngoodsAgreeResponse 卖家同意退货，支持淘宝和天猫的订单。 */
type TaobaoRpReturngoodsAgreeResponse struct {
    
    /* is_success Basic操作成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoRdcAligeniusAutorefundsQueryRequest （此接口后期将不再维护，请勿使用）供第三方商家查询授权给自己的所有退款款订单的退款信息 */
type TaobaoRdcAligeniusAutorefundsQueryRequest struct {
    
    /* end_time required查询数据时间段结束时间 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* page_no required页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size required每页数据数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required查询数据时间段开始时间 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoRdcAligeniusAutorefundsQueryRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.autorefunds.query"
}

/* TaobaoRdcAligeniusAutorefundsQueryResponse （此接口后期将不再维护，请勿使用）供第三方商家查询授权给自己的所有退款款订单的退款信息 */
type TaobaoRdcAligeniusAutorefundsQueryResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoTrademonitorGetRequest 仅限供应商调用此接口查询经销商品监控信息 */
type TaobaoFenxiaoTrademonitorGetRequest struct {
    
    /* distributor_nick optional经销商的淘宝账号 */
    distributor_nick string `json:"distributor_nick";xml:"distributor_nick"`
    
    /* end_created optional结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
    end_created time.Time `json:"end_created";xml:"end_created"`
    
    /* fields optional多个字段用","分隔。 fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。例如：trade_monitors.item_title表示只返回item_title */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。（大于0的整数。小于1按1计） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。（每页条数不超过50条，超过50或小于0均按50计） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* product_id optional产品id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* start_created optional起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
    start_created time.Time `json:"start_created";xml:"start_created"`
    
}

func (req *TaobaoFenxiaoTrademonitorGetRequest) GetAPIName() string {
	return "taobao.fenxiao.trademonitor.get"
}

/* TaobaoFenxiaoTrademonitorGetResponse 仅限供应商调用此接口查询经销商品监控信息 */
type TaobaoFenxiaoTrademonitorGetResponse struct {
    
    /* total_results Basic搜索到的经销商品订单数量 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trade_monitors Object Array经销商品订单监控信息 */
    trade_monitors TradeMonitor `json:"trade_monitors";xml:"trade_monitors"`
    
}

/* TaobaoRpRefundsAgreeRequest 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。 */
type TaobaoRpRefundsAgreeRequest struct {
    
    /* code optional短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。 */
    code string `json:"code";xml:"code"`
    
    /* refund_infos required退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。 */
    refund_infos string `json:"refund_infos";xml:"refund_infos"`
    
}

func (req *TaobaoRpRefundsAgreeRequest) GetAPIName() string {
	return "taobao.rp.refunds.agree"
}

/* TaobaoRpRefundsAgreeResponse 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。 */
type TaobaoRpRefundsAgreeResponse struct {
    
    /* message Basic信息 */
    message string `json:"message";xml:"message"`
    
    /* msg_code Basic批量退款操作情况，可选值：OP_SUCC（全部成功），SOME_OP_SUCC（部分成功），OP_FAILURE_UE（全部失败） */
    msg_code string `json:"msg_code";xml:"msg_code"`
    
    /* results Object Array退款操作结果列表 */
    results RefundMappingResult `json:"results";xml:"results"`
    
    /* succ Basic操作成功 */
    succ bool `json:"succ";xml:"succ"`
    
}

/* TaobaoItemAnchorGetRequest 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点 */
type TaobaoItemAnchorGetRequest struct {
    
    /* cat_id required对应类目编号 */
    cat_id int64 `json:"cat_id";xml:"cat_id"`
    
    /* _type required宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1. */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItemAnchorGetRequest) GetAPIName() string {
	return "taobao.item.anchor.get"
}

/* TaobaoItemAnchorGetResponse 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点 */
type TaobaoItemAnchorGetResponse struct {
    
    /* anchor_modules Object Array宝贝描述规范化可使用打标模块的锚点信息 */
    anchor_modules IdsModule `json:"anchor_modules";xml:"anchor_modules"`
    
    /* total_results Basic返回的宝贝描述模板结果数目 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoRefundRefuseRequest 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：
1. 传入的refund_id和相应的tid, oid必须匹配
2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
3. 只有卖家才能执行拒绝退款操作
4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单 */
type TaobaoRefundRefuseRequest struct {
    
    /* oid optional退款记录对应的交易子订单号 */
    oid int64 `json:"oid";xml:"oid"`
    
    /* refund_id required退款单号 */
    refund_id int64 `json:"refund_id";xml:"refund_id"`
    
    /* refund_phase optional可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。 */
    refund_phase string `json:"refund_phase";xml:"refund_phase"`
    
    /* refund_version optional退款版本号，天猫退款为必填项。 */
    refund_version int64 `json:"refund_version";xml:"refund_version"`
    
    /* refuse_message required拒绝退款时的说明信息，长度2-200 */
    refuse_message string `json:"refuse_message";xml:"refuse_message"`
    
    /* refuse_proof optional拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。 */
    refuse_proof []byte `json:"refuse_proof";xml:"refuse_proof"`
    
    /* refuse_reason_id optional拒绝原因编号，会提供用户拒绝原因列表供选择 */
    refuse_reason_id int64 `json:"refuse_reason_id";xml:"refuse_reason_id"`
    
    /* tid optional退款记录对应的交易订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoRefundRefuseRequest) GetAPIName() string {
	return "taobao.refund.refuse"
}

/* TaobaoRefundRefuseResponse 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：
1. 传入的refund_id和相应的tid, oid必须匹配
2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
3. 只有卖家才能执行拒绝退款操作
4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单 */
type TaobaoRefundRefuseResponse struct {
    
    /* is_success Basic拒绝退款操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* refund Object拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段 */
    refund Refund `json:"refund";xml:"refund"`
    
}

/* TaobaoTradeAmountGetRequest 卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同 */
type TaobaoTradeAmountGetRequest struct {
    
    /* fields required订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段) */
    fields string `json:"fields";xml:"fields"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeAmountGetRequest) GetAPIName() string {
	return "taobao.trade.amount.get"
}

/* TaobaoTradeAmountGetResponse 卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同 */
type TaobaoTradeAmountGetResponse struct {
    
    /* trade_amount Object主订单的财务信息详情 */
    trade_amount TradeAmount `json:"trade_amount";xml:"trade_amount"`
    
}

/* TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest 供应商修改经销采购单备注 */
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
    
    /* dealer_order_id required经销采购单ID */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
    /* supplier_memo optional备注留言，可为空 */
    supplier_memo string `json:"supplier_memo";xml:"supplier_memo"`
    
    /* supplier_memo_flag required旗子的标记，必选。
1-5之间的数字。
非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色 */
    supplier_memo_flag int64 `json:"supplier_memo_flag";xml:"supplier_memo_flag"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}

/* TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse 供应商修改经销采购单备注 */
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse struct {
    
    /* is_success Basic操作是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoWlbWmsItemCombinationCreateRequest 创建组合商品与子商品关系 */
type TaobaoWlbWmsItemCombinationCreateRequest struct {
    
    /* item_id required组合商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* sub_item_list required子货品列表 */
    sub_item_list SubItemList `json:"sub_item_list";xml:"sub_item_list"`
    
}

func (req *TaobaoWlbWmsItemCombinationCreateRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.create"
}

/* TaobaoWlbWmsItemCombinationCreateResponse 创建组合商品与子商品关系 */
type TaobaoWlbWmsItemCombinationCreateResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoWlbWmsItemCombinationDeleteRequest 删除货品组合关系 */
type TaobaoWlbWmsItemCombinationDeleteRequest struct {
    
    /* item_id optional组合货品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbWmsItemCombinationDeleteRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.delete"
}

/* TaobaoWlbWmsItemCombinationDeleteResponse 删除货品组合关系 */
type TaobaoWlbWmsItemCombinationDeleteResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoWlbWmsItemCombinationGetRequest 查询组合商品的组合关系 */
type TaobaoWlbWmsItemCombinationGetRequest struct {
    
    /* itemid required货品Id */
    itemid int64 `json:"itemid";xml:"itemid"`
    
}

func (req *TaobaoWlbWmsItemCombinationGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.get"
}

/* TaobaoWlbWmsItemCombinationGetResponse 查询组合商品的组合关系 */
type TaobaoWlbWmsItemCombinationGetResponse struct {
    
    /* result Object接口返回结果 */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoProductGradepriceUpdateRequest 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格 */
type TaobaoFenxiaoProductGradepriceUpdateRequest struct {
    
    /* ids required会员等级的id或者分销商id，例如：”1001,2001,1002” */
    ids int64 `json:"ids";xml:"ids"`
    
    /* prices required优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25 */
    prices string `json:"prices";xml:"prices"`
    
    /* product_id required产品Id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* sku_id optionalskuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* target_type required选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR" */
    target_type string `json:"target_type";xml:"target_type"`
    
    /* trade_type optional交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销) */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductGradepriceUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.gradeprice.update"
}

/* TaobaoFenxiaoProductGradepriceUpdateResponse 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格 */
type TaobaoFenxiaoProductGradepriceUpdateResponse struct {
    
    /* is_success Basic返回操作结果：成功或失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoRdcAligeniusOrderReturngoodsNotifyRequest 退货单创建结果反馈 */
type TaobaoRdcAligeniusOrderReturngoodsNotifyRequest struct {
    
    /* parent_order_id optional主订单号 */
    parent_order_id int64 `json:"parent_order_id";xml:"parent_order_id"`
    
    /* refund_return_notes optional退货单创建结果列表 */
    refund_return_notes RefundReturnNotes `json:"refund_return_notes";xml:"refund_return_notes"`
    
}

func (req *TaobaoRdcAligeniusOrderReturngoodsNotifyRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.order.returngoods.notify"
}

/* TaobaoRdcAligeniusOrderReturngoodsNotifyResponse 退货单创建结果反馈 */
type TaobaoRdcAligeniusOrderReturngoodsNotifyResponse struct {
    
    /* result Basicsuccess */
    result bool `json:"result";xml:"result"`
    
    /* result_code BasicerrorCode */
    result_code string `json:"result_code";xml:"result_code"`
    
    /* result_info BasicerrorInfo */
    result_info string `json:"result_info";xml:"result_info"`
    
}

/* TaobaoOmniorderStoreSdtcancelRequest 通知速店通取消取号 */
type TaobaoOmniorderStoreSdtcancelRequest struct {
    
    /* package_id required取号返回的packageId */
    package_id int64 `json:"package_id";xml:"package_id"`
    
}

func (req *TaobaoOmniorderStoreSdtcancelRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtcancel"
}

/* TaobaoOmniorderStoreSdtcancelResponse 通知速店通取消取号 */
type TaobaoOmniorderStoreSdtcancelResponse struct {
    
    /* result Object返回结果 */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoOmniorderStoreSdtquerystationRequest 速店通查询站点信息 */
type TaobaoOmniorderStoreSdtquerystationRequest struct {
    
    /* param_long2 required取号时返回的packageId */
    param_long2 int64 `json:"param_long2";xml:"param_long2"`
    
}

func (req *TaobaoOmniorderStoreSdtquerystationRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtquerystation"
}

/* TaobaoOmniorderStoreSdtquerystationResponse 速店通查询站点信息 */
type TaobaoOmniorderStoreSdtquerystationResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest 在自动化任务核对地址卡片中，透出修改地址选择器，让用户自己选择要修改的地址，然后同步到商家后台ERP系统 */
type QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest struct {
    
    /* bizOrderId required交易订单ID */
    bizOrderId string `json:"bizOrderId";xml:"bizOrderId"`
    
    /* buyerNick required买家账号名 */
    buyerNick string `json:"buyerNick";xml:"buyerNick"`
    
    /* modifiedAddress required要修改的地址信息 */
    modifiedAddress ModifiedAddress `json:"modifiedAddress";xml:"modifiedAddress"`
    
    /* originalAddress optional订单原始收货地址信息 */
    originalAddress OriginalAddress `json:"originalAddress";xml:"originalAddress"`
    
    /* sellerNick required店铺主账号 */
    sellerNick string `json:"sellerNick";xml:"sellerNick"`
    
}

func (req *QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest) GetAPIName() string {
	return "qimen.taobao.qianniu.cloudkefu.address.self.modify"
}

/* QimenTaobaoQianniuCloudkefuAddressSelfModifyResponse 在自动化任务核对地址卡片中，透出修改地址选择器，让用户自己选择要修改的地址，然后同步到商家后台ERP系统 */
type QimenTaobaoQianniuCloudkefuAddressSelfModifyResponse struct {
    
    /* result Object修改地址返回结果 */
    result ResultDO `json:"result";xml:"result"`
    
}

/* TmallItemSizemappingTemplateCreateRequest 新增天猫商品尺码表模板

男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：
脚长（cm） 必选

内衣-文胸类目，尺码表维度为：
上胸围（cm） 必选
下胸围（cm） 必选
上下胸围差（cm） 必选
身高（cm）
体重（公斤）

内衣-内裤类目，尺码表维度为：
腰围（cm） 必选
臀围（cm） 必选
身高（cm）
体重（公斤）
裤长（cm）
裆部（cm）
脚口（cm）
腿围（cm）

内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
腰围（cm）
肩宽（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
领口（cm）

内衣-睡裤/保暖裤类目，尺码维度为：
身高（cm） 必选
腰围（cm） 必选
体重（公斤）
臀围（cm）
裆部（cm）
裤长（cm）
脚口（cm）
腿围（cm）
裤侧长（cm）

内衣-睡裙类目，尺码维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
裙长（cm）
腰围（cm）
袖长（cm）
肩宽（cm）
背宽（cm）
腿围（cm）
臀围（cm）
底摆（cm）

男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
肩宽（cm）
胸围（cm）
腰围（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
摆围（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
中腰（cm）
领深（cm）
领高（cm）
领宽（cm）
领围（cm）
圆摆后中长（cm）
平摆衣长（cm）
圆摆衣长（cm）

男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
腰围（cm）
臀围（cm）
裤长（cm）
裙长（cm）
裙摆长（cm）
腿围（cm）
膝围（cm）
小脚围（cm）
拉伸腰围（cm）
坐围（cm）
拉伸坐围（cm）
脚口（cm）
前浪（cm）
后浪（cm）
横档（cm）

如果上述维度满足，可以自定义最多5个维度。

模板格式为：
尺码值:维度名称:数值
如：M:身高（cm）:160,L:身高（cm）:170 */
type TmallItemSizemappingTemplateCreateRequest struct {
    
    /* template_content required尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。 */
    template_content string `json:"template_content";xml:"template_content"`
    
    /* template_name required尺码表模板名称 */
    template_name string `json:"template_name";xml:"template_name"`
    
}

func (req *TmallItemSizemappingTemplateCreateRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.create"
}

/* TmallItemSizemappingTemplateCreateResponse 新增天猫商品尺码表模板

男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：
脚长（cm） 必选

内衣-文胸类目，尺码表维度为：
上胸围（cm） 必选
下胸围（cm） 必选
上下胸围差（cm） 必选
身高（cm）
体重（公斤）

内衣-内裤类目，尺码表维度为：
腰围（cm） 必选
臀围（cm） 必选
身高（cm）
体重（公斤）
裤长（cm）
裆部（cm）
脚口（cm）
腿围（cm）

内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
腰围（cm）
肩宽（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
领口（cm）

内衣-睡裤/保暖裤类目，尺码维度为：
身高（cm） 必选
腰围（cm） 必选
体重（公斤）
臀围（cm）
裆部（cm）
裤长（cm）
脚口（cm）
腿围（cm）
裤侧长（cm）

内衣-睡裙类目，尺码维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
裙长（cm）
腰围（cm）
袖长（cm）
肩宽（cm）
背宽（cm）
腿围（cm）
臀围（cm）
底摆（cm）

男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
肩宽（cm）
胸围（cm）
腰围（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
摆围（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
中腰（cm）
领深（cm）
领高（cm）
领宽（cm）
领围（cm）
圆摆后中长（cm）
平摆衣长（cm）
圆摆衣长（cm）

男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
腰围（cm）
臀围（cm）
裤长（cm）
裙长（cm）
裙摆长（cm）
腿围（cm）
膝围（cm）
小脚围（cm）
拉伸腰围（cm）
坐围（cm）
拉伸坐围（cm）
脚口（cm）
前浪（cm）
后浪（cm）
横档（cm）

如果上述维度满足，可以自定义最多5个维度。

模板格式为：
尺码值:维度名称:数值
如：M:身高（cm）:160,L:身高（cm）:170 */
type TmallItemSizemappingTemplateCreateResponse struct {
    
    /* size_mapping_template Object尺码表模板 */
    size_mapping_template SizeMappingTemplateDo `json:"size_mapping_template";xml:"size_mapping_template"`
    
}

/* TaobaoOmniorderStoreGetconsignmailcodeRequest 用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号 */
type TaobaoOmniorderStoreGetconsignmailcodeRequest struct {
    
    /* channel required淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS) */
    channel string `json:"channel";xml:"channel"`
    
    /* receiver optional收件人信息 */
    receiver ReceiverDto `json:"receiver";xml:"receiver"`
    
    /* sdt_extend_info_d_t_o optional扩展信息 */
    sdt_extend_info_d_t_o SdtExtendInfoDto `json:"sdt_extend_info_d_t_o";xml:"sdt_extend_info_d_t_o"`
    
    /* sender_contact optional发件人联系电话，如空则表示使用门店信息中的电话号码 */
    sender_contact string `json:"sender_contact";xml:"sender_contact"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* trades required订单信息，目前一次请求只支持一个主订单 */
    trades TradeOrderInfoDto `json:"trades";xml:"trades"`
    
}

func (req *TaobaoOmniorderStoreGetconsignmailcodeRequest) GetAPIName() string {
	return "taobao.omniorder.store.getconsignmailcode"
}

/* TaobaoOmniorderStoreGetconsignmailcodeResponse 用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号 */
type TaobaoOmniorderStoreGetconsignmailcodeResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* AlibabaEinvoiceCreatereqRequest ERP发起开票请求 */
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
    invoice_time time.Time `json:"invoice_time";xml:"invoice_time"`
    
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

/* AlibabaEinvoiceCreatereqResponse ERP发起开票请求 */
type AlibabaEinvoiceCreatereqResponse struct {
    
    /* is_success Basic开票信息是否成功接受 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TmallItemSizemappingTemplateUpdateRequest 更新天猫商品尺码表模板 */
type TmallItemSizemappingTemplateUpdateRequest struct {
    
    /* template_content required尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。 */
    template_content string `json:"template_content";xml:"template_content"`
    
    /* template_id required尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
    /* template_name required尺码表模板名称 */
    template_name string `json:"template_name";xml:"template_name"`
    
}

func (req *TmallItemSizemappingTemplateUpdateRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.update"
}

/* TmallItemSizemappingTemplateUpdateResponse 更新天猫商品尺码表模板 */
type TmallItemSizemappingTemplateUpdateResponse struct {
    
    /* size_mapping_template Object尺码表模板 */
    size_mapping_template SizeMappingTemplateDo `json:"size_mapping_template";xml:"size_mapping_template"`
    
}

/* TaobaoTradesSoldIncrementvGetRequest 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。 */
type TaobaoTradesSoldIncrementvGetRequest struct {
    
    /* end_create required查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
    end_create time.Time `json:"end_create";xml:"end_create"`
    
    /* ext_type optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
    ext_type string `json:"ext_type";xml:"ext_type"`
    
    /* fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。 */
    fields string `json:"fields";xml:"fields"`
    
    /* page_no optional页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_create required查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
    start_create time.Time `json:"start_create";xml:"start_create"`
    
    /* status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 */
    status string `json:"status";xml:"status"`
    
    /* tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
    tag string `json:"tag";xml:"tag"`
    
    /* _type optional交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。 */
    _type string `json:"type";xml:"type"`
    
    /* use_has_next optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
    use_has_next bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldIncrementvGetRequest) GetAPIName() string {
	return "taobao.trades.sold.incrementv.get"
}

/* TaobaoTradesSoldIncrementvGetResponse 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。 */
type TaobaoTradesSoldIncrementvGetResponse struct {
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic搜索到的交易信息总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
    trades Trade `json:"trades";xml:"trades"`
    
}

/* TaobaoOmniorderStoreWarehouseRequest 为了能够支持逆向到门店对应的区域仓，商家需要配置门店与区域仓的映射关系，这个接口可以给商家提供门店与区域仓映射关系的增删改查功能。 */
type TaobaoOmniorderStoreWarehouseRequest struct {
    
    /* operation required操作，1表示增加或者更新，2表示删除，3表示查询 */
    operation int64 `json:"operation";xml:"operation"`
    
    /* store_id optional门店id */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* warehouse_address optional仓联系地址 */
    warehouse_address string `json:"warehouse_address";xml:"warehouse_address"`
    
    /* warehouse_code optional仓编码 */
    warehouse_code string `json:"warehouse_code";xml:"warehouse_code"`
    
    /* warehouse_contact optional仓联系人 */
    warehouse_contact string `json:"warehouse_contact";xml:"warehouse_contact"`
    
    /* warehouse_mobile optional仓联系电话 */
    warehouse_mobile string `json:"warehouse_mobile";xml:"warehouse_mobile"`
    
}

func (req *TaobaoOmniorderStoreWarehouseRequest) GetAPIName() string {
	return "taobao.omniorder.store.warehouse"
}

/* TaobaoOmniorderStoreWarehouseResponse 为了能够支持逆向到门店对应的区域仓，商家需要配置门店与区域仓的映射关系，这个接口可以给商家提供门店与区域仓映射关系的增删改查功能。 */
type TaobaoOmniorderStoreWarehouseResponse struct {
    
    /* data Basic成功增加或者更新一条门店与区域仓的关联 */
    data string `json:"data";xml:"data"`
    
    /* fail_code Basiccode */
    fail_code string `json:"fail_code";xml:"fail_code"`
    
    /* fail_message Basicmessage */
    fail_message string `json:"fail_message";xml:"fail_message"`
    
}

/* TmallItemSizemappingTemplateDeleteRequest 删除天猫商品尺码表模板 */
type TmallItemSizemappingTemplateDeleteRequest struct {
    
    /* template_id required尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TmallItemSizemappingTemplateDeleteRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.delete"
}

/* TmallItemSizemappingTemplateDeleteResponse 删除天猫商品尺码表模板 */
type TmallItemSizemappingTemplateDeleteResponse struct {
    
    /* template_id Basic尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

/* TmallItemSizemappingTemplateGetRequest 获取天猫商品尺码表模板 */
type TmallItemSizemappingTemplateGetRequest struct {
    
    /* template_id required尺码表模板ID */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TmallItemSizemappingTemplateGetRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.get"
}

/* TmallItemSizemappingTemplateGetResponse 获取天猫商品尺码表模板 */
type TmallItemSizemappingTemplateGetResponse struct {
    
    /* size_mapping_template Object尺码表模板 */
    size_mapping_template Model `json:"size_mapping_template";xml:"size_mapping_template"`
    
}

/* TaobaoOmniorderStoreSdtconsignRequest ISV取完单号后通知菜鸟裹裹发货 */
type TaobaoOmniorderStoreSdtconsignRequest struct {
    
    /* package_id required取号接口返回的包裹id */
    package_id string `json:"package_id";xml:"package_id"`
    
    /* tag_code optional发货标签号 */
    tag_code string `json:"tag_code";xml:"tag_code"`
    
}

func (req *TaobaoOmniorderStoreSdtconsignRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtconsign"
}

/* TaobaoOmniorderStoreSdtconsignResponse ISV取完单号后通知菜鸟裹裹发货 */
type TaobaoOmniorderStoreSdtconsignResponse struct {
    
    /* data Objectdata */
    data SdtConsignResponse `json:"data";xml:"data"`
    
    /* err_code Basic异常码 0 为正常，否则异常 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* message Basic异常信息 */
    message string `json:"message";xml:"message"`
    
}

/* AlibabaEinvoiceCreateResultsIncrementGetRequest 增量开票结果获取 */
type AlibabaEinvoiceCreateResultsIncrementGetRequest struct {
    
    /* end_modified required终止查询时间 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* page_no optional显示的页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional页面大小(不能超过200) */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* start_modified required起始查询时间 */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
    /* status optional开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败) */
    status string `json:"status";xml:"status"`
    
}

func (req *AlibabaEinvoiceCreateResultsIncrementGetRequest) GetAPIName() string {
	return "alibaba.einvoice.create.results.increment.get"
}

/* AlibabaEinvoiceCreateResultsIncrementGetResponse 增量开票结果获取 */
type AlibabaEinvoiceCreateResultsIncrementGetResponse struct {
    
    /* invoice_result_list Object Array开票结果返回列表 */
    invoice_result_list InvoiceResult `json:"invoice_result_list";xml:"invoice_result_list"`
    
    /* total_count Basic符合条件的开票总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* AlibabaEinvoiceCreateResultGetRequest ERP开票结果获取 */
type AlibabaEinvoiceCreateResultGetRequest struct {
    
    /* out_shop_name optional外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致 */
    out_shop_name string `json:"out_shop_name";xml:"out_shop_name"`
    
    /* payee_register_no required收款方税务登记证号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* platform_code optional电商平台代码。淘宝：taobao，天猫：tmall */
    platform_code string `json:"platform_code";xml:"platform_code"`
    
    /* platform_tid optional电商平台对应的订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
    /* serial_no optional流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoiceCreateResultGetRequest) GetAPIName() string {
	return "alibaba.einvoice.create.result.get"
}

/* AlibabaEinvoiceCreateResultGetResponse ERP开票结果获取 */
type AlibabaEinvoiceCreateResultGetResponse struct {
    
    /* invoice_result_list Object Array开票返回结果数据列表 */
    invoice_result_list InvoiceResult `json:"invoice_result_list";xml:"invoice_result_list"`
    
}

/* TaobaoKfcKeywordSearchRequest 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果 */
type TaobaoKfcKeywordSearchRequest struct {
    
    /* apply optional应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。

这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。


通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。 */
    apply string `json:"apply";xml:"apply"`
    
    /* content required需要过滤的文本信息 */
    content string `json:"content";xml:"content"`
    
    /* nick optional发布信息的淘宝会员名，可以不传 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoKfcKeywordSearchRequest) GetAPIName() string {
	return "taobao.kfc.keyword.search"
}

/* TaobaoKfcKeywordSearchResponse 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果 */
type TaobaoKfcKeywordSearchResponse struct {
    
    /* kfc_search_result ObjectKFC 关键词过滤匹配结果 */
    kfc_search_result KfcSearchResult `json:"kfc_search_result";xml:"kfc_search_result"`
    
}

/* TaobaoQimenCombineitemQueryRequest 组合货品关系查询 */
type TaobaoQimenCombineitemQueryRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemId optional奇门仓储字段 */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* ownerCode optional奇门仓储字段 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenCombineitemQueryRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.query"
}

/* TaobaoQimenCombineitemQueryResponse 组合货品关系查询 */
type TaobaoQimenCombineitemQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* combItems Object Array奇门仓储字段 */
    combItems CombItem `json:"combItems";xml:"combItems"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TmallItemSizemappingTemplatesListRequest 获取所有尺码表模板列表。 */
type TmallItemSizemappingTemplatesListRequest struct {
    
}

func (req *TmallItemSizemappingTemplatesListRequest) GetAPIName() string {
	return "tmall.item.sizemapping.templates.list"
}

/* TmallItemSizemappingTemplatesListResponse 获取所有尺码表模板列表。 */
type TmallItemSizemappingTemplatesListResponse struct {
    
    /* size_mapping_templates Object Array尺码表模板列表 */
    size_mapping_templates SizeMappingTemplate `json:"size_mapping_templates";xml:"size_mapping_templates"`
    
}

/* TaobaoTmcTopicGroupAddRequest 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由 */
type TaobaoTmcTopicGroupAddRequest struct {
    
    /* group_name required消息分组名，如果不存在，会自动创建 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* topics required消息topic名称，多个以逗号(,)分割 */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcTopicGroupAddRequest) GetAPIName() string {
	return "taobao.tmc.topic.group.add"
}

/* TaobaoTmcTopicGroupAddResponse 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由 */
type TaobaoTmcTopicGroupAddResponse struct {
    
    /* result Basictrue */
    result bool `json:"result";xml:"result"`
    
}

/* TmallChannelTradeOrderCreateRequest 创建渠道分销单 */
type TmallChannelTradeOrderCreateRequest struct {
    
    /* param0 required入参 */
    param0 TopChannelPurchaseOrderCreateParam `json:"param0";xml:"param0"`
    
}

func (req *TmallChannelTradeOrderCreateRequest) GetAPIName() string {
	return "tmall.channel.trade.order.create"
}

/* TmallChannelTradeOrderCreateResponse 创建渠道分销单 */
type TmallChannelTradeOrderCreateResponse struct {
    
    /* main_purchase_order_list Basic Array采购单号 */
    main_purchase_order_list map[string]interface{} `json:"main_purchase_order_list";xml:"main_purchase_order_list"`
    
}

/* TaobaoFenxiaoDistributorsGetRequest 查询和当前登录供应商有合作关系的分销商的信息 */
type TaobaoFenxiaoDistributorsGetRequest struct {
    
    /* nicks required分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。 */
    nicks string `json:"nicks";xml:"nicks"`
    
}

func (req *TaobaoFenxiaoDistributorsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributors.get"
}

/* TaobaoFenxiaoDistributorsGetResponse 查询和当前登录供应商有合作关系的分销商的信息 */
type TaobaoFenxiaoDistributorsGetResponse struct {
    
    /* distributors Object Array分销商详细信息 */
    distributors Distributor `json:"distributors";xml:"distributors"`
    
}

/* QimenTaobaoAutoEntryorderGiftitemcancelRequest 该接口一次只能回传一个主交易号的BMS增加的货品取消信息 */
type QimenTaobaoAutoEntryorderGiftitemcancelRequest struct {
    
    /* customerId requiredcustomerId */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求 */
    request ErpBmsOrderGiftCancelRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoAutoEntryorderGiftitemcancelRequest) GetAPIName() string {
	return "qimen.taobao.auto.entryorder.giftitemcancel"
}

/* QimenTaobaoAutoEntryorderGiftitemcancelResponse 该接口一次只能回传一个主交易号的BMS增加的货品取消信息 */
type QimenTaobaoAutoEntryorderGiftitemcancelResponse struct {
    
    /* response Objectresponse */
    response Response `json:"response";xml:"response"`
    
}

/* QimenTaobaoBmsErptradeInterceptRequest 调用ERP接口，拦截订单 */
type QimenTaobaoBmsErptradeInterceptRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求主体 */
    request BmsTaobaoOrderIntercepteRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsErptradeInterceptRequest) GetAPIName() string {
	return "qimen.taobao.bms.erptrade.intercept"
}

/* QimenTaobaoBmsErptradeInterceptResponse 调用ERP接口，拦截订单 */
type QimenTaobaoBmsErptradeInterceptResponse struct {
    
    /* response ObjectResponse */
    response Response `json:"response";xml:"response"`
    
}

/* TmallStoredeliverAllocationAcceptRequest 商品通门店发货业务门店接单拒单接口 */
type TmallStoredeliverAllocationAcceptRequest struct {
    
    /* allocation_code required派单号 */
    allocation_code string `json:"allocation_code";xml:"allocation_code"`
    
    /* is_accept required是否接单 */
    is_accept bool `json:"is_accept";xml:"is_accept"`
    
    /* sub_order_code required子订单号必须和派单号匹配 */
    sub_order_code int64 `json:"sub_order_code";xml:"sub_order_code"`
    
}

func (req *TmallStoredeliverAllocationAcceptRequest) GetAPIName() string {
	return "tmall.storedeliver.allocation.accept"
}

/* TmallStoredeliverAllocationAcceptResponse 商品通门店发货业务门店接单拒单接口 */
type TmallStoredeliverAllocationAcceptResponse struct {
    
    /* is_success Basic是否执行成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* QimenTaobaoBmsErptradeTransferconsignRequest BMS调用ERP订单菜鸟仓&商家仓互转接口 */
type QimenTaobaoBmsErptradeTransferconsignRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求体 */
    request BmsErptradeTransferConsignRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsErptradeTransferconsignRequest) GetAPIName() string {
	return "qimen.taobao.bms.erptrade.transferconsign"
}

/* QimenTaobaoBmsErptradeTransferconsignResponse BMS调用ERP订单菜鸟仓&商家仓互转接口 */
type QimenTaobaoBmsErptradeTransferconsignResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}

/* QimenTaobaoBmsTradeConsignRequest BMS通知ERP交易单整单出库接口 */
type QimenTaobaoBmsTradeConsignRequest struct {
    
    /* customerId required货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optional请求实体 */
    request BmsTradeConsignRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsTradeConsignRequest) GetAPIName() string {
	return "qimen.taobao.bms.trade.consign"
}

/* QimenTaobaoBmsTradeConsignResponse BMS通知ERP交易单整单出库接口 */
type QimenTaobaoBmsTradeConsignResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}

/* TaobaoLogisticsExpressModifyAppointRequest 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期 */
type TaobaoLogisticsExpressModifyAppointRequest struct {
    
    /* express_modify_appoint_top_request required改约请求对象 */
    express_modify_appoint_top_request ExpressModifyAppointTopRequestDto `json:"express_modify_appoint_top_request";xml:"express_modify_appoint_top_request"`
    
}

func (req *TaobaoLogisticsExpressModifyAppointRequest) GetAPIName() string {
	return "taobao.logistics.express.modify.appoint"
}

/* TaobaoLogisticsExpressModifyAppointResponse 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期 */
type TaobaoLogisticsExpressModifyAppointResponse struct {
    
    /* result Object调用结果 */
    result SingleResultDto `json:"result";xml:"result"`
    
}

/* TaobaoLogisticsConsignResendRequest 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：<br>
1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。 */
type TaobaoLogisticsConsignResendRequest struct {
    
    /* company_code required物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取。<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font> */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* feature optionalfeature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
    feature string `json:"feature";xml:"feature"`
    
    /* is_split optional表明是否是拆单，默认值0，1表示拆单 */
    is_split int64 `json:"is_split";xml:"is_split"`
    
    /* out_sid required运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
    out_sid string `json:"out_sid";xml:"out_sid"`
    
    /* seller_ip optional商家的IP地址 */
    seller_ip string `json:"seller_ip";xml:"seller_ip"`
    
    /* sub_tid optional拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
    sub_tid int64 `json:"sub_tid";xml:"sub_tid"`
    
    /* tid required淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsConsignResendRequest) GetAPIName() string {
	return "taobao.logistics.consign.resend"
}

/* TaobaoLogisticsConsignResendResponse 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：<br>
1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。 */
type TaobaoLogisticsConsignResendResponse struct {
    
    /* shipping Object返回发货是否成功is_success */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}

/* TaobaoRdcAligeniusOrderEventUpdateRequest 供erp回传订单变更状态事件 */
type TaobaoRdcAligeniusOrderEventUpdateRequest struct {
    
    /* param0 optional参数 */
    param0 OrderEventDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusOrderEventUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.order.event.update"
}

/* TaobaoRdcAligeniusOrderEventUpdateResponse 供erp回传订单变更状态事件 */
type TaobaoRdcAligeniusOrderEventUpdateResponse struct {
    
    /* fail_code Basic错误码 */
    fail_code string `json:"fail_code";xml:"fail_code"`
    
    /* fail_info Basic错误描述 */
    fail_info string `json:"fail_info";xml:"fail_info"`
    
    /* success_flag Basic是否成功 */
    success_flag bool `json:"success_flag";xml:"success_flag"`
    
}

/* TaobaoTmcTopicGroupDeleteRequest 删除根据topic名称路由消息到不同的分组关系 */
type TaobaoTmcTopicGroupDeleteRequest struct {
    
    /* group_id optional消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系 */
    group_id int64 `json:"group_id";xml:"group_id"`
    
    /* group_name required消息分组名 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* topics required消息topic名称，多个以逗号(,)分割 */
    topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcTopicGroupDeleteRequest) GetAPIName() string {
	return "taobao.tmc.topic.group.delete"
}

/* TaobaoTmcTopicGroupDeleteResponse 删除根据topic名称路由消息到不同的分组关系 */
type TaobaoTmcTopicGroupDeleteResponse struct {
    
    /* result Basictrue */
    result bool `json:"result";xml:"result"`
    
}

/* TmallItemSchemaUpdateRequest 天猫根据规则编辑商品 */
type TmallItemSchemaUpdateRequest struct {
    
    /* category_id optional商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* product_id optional商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* xml_data required根据tmall.item.update.schema.get生成的商品编辑规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaUpdateRequest) GetAPIName() string {
	return "tmall.item.schema.update"
}

/* TmallItemSchemaUpdateResponse 天猫根据规则编辑商品 */
type TmallItemSchemaUpdateResponse struct {
    
    /* gmt_modified Basic商品更新操作成功时间 */
    gmt_modified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    /* update_item_result Basic返回商品发布结果 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}

/* CainiaoSmartdeliveryIssubscribeIQueryRequest 查询商家时候订购智能发货引擎服务 */
type CainiaoSmartdeliveryIssubscribeIQueryRequest struct {
    
}

func (req *CainiaoSmartdeliveryIssubscribeIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.issubscribe.i.query"
}

/* CainiaoSmartdeliveryIssubscribeIQueryResponse 查询商家时候订购智能发货引擎服务 */
type CainiaoSmartdeliveryIssubscribeIQueryResponse struct {
    
    /* successful Basictrue:商家已订购智能发货引擎服务,false:商家还没有订购或订购已过期 */
    successful bool `json:"successful";xml:"successful"`
    
}

/* TmallItemUpdateSchemaGetRequest Schema方式编辑天猫商品时，编辑商品规则获取 */
type TmallItemUpdateSchemaGetRequest struct {
    
    /* category_id optional商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required需要编辑的商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* product_id optional商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallItemUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.item.update.schema.get"
}

/* TmallItemUpdateSchemaGetResponse Schema方式编辑天猫商品时，编辑商品规则获取 */
type TmallItemUpdateSchemaGetResponse struct {
    
    /* update_item_result Basic返回发布商品的规则文档 */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}

/* TmallProductUpdateSchemaGetRequest 获取用户更新产品的规则 */
type TmallProductUpdateSchemaGetRequest struct {
    
    /* product_id required产品编号 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallProductUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.product.update.schema.get"
}

/* TmallProductUpdateSchemaGetResponse 获取用户更新产品的规则 */
type TmallProductUpdateSchemaGetResponse struct {
    
    /* update_product_schema Basic参数产品ID对产品的更新规则 */
    update_product_schema string `json:"update_product_schema";xml:"update_product_schema"`
    
}

/* TmallProductSchemaUpdateRequest 产品更新接口 */
type TmallProductSchemaUpdateRequest struct {
    
    /* product_id required产品编号 */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* xml_data required根据tmall.product.update.schema.get生成的产品更新规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallProductSchemaUpdateRequest) GetAPIName() string {
	return "tmall.product.schema.update"
}

/* TmallProductSchemaUpdateResponse 产品更新接口 */
type TmallProductSchemaUpdateResponse struct {
    
    /* update_product_result Basic产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间 */
    update_product_result string `json:"update_product_result";xml:"update_product_result"`
    
}

/* TmallExchangeRefusereasonGetRequest 获取拒绝换货原因列表 */
type TmallExchangeRefusereasonGetRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* dispute_type optional换货申请类型：0-任意类型；1-售中；2-售后 */
    dispute_type int64 `json:"dispute_type";xml:"dispute_type"`
    
    /* fields required返回字段 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeRefusereasonGetRequest) GetAPIName() string {
	return "tmall.exchange.refusereason.get"
}

/* TmallExchangeRefusereasonGetResponse 获取拒绝换货原因列表 */
type TmallExchangeRefusereasonGetResponse struct {
    
    /* result Object返回结果 */
    result ResultSet `json:"result";xml:"result"`
    
}

/* TaobaoTaeBillGetRequest 查询单笔账单明细 */
type TaobaoTaeBillGetRequest struct {
    
    /* account_id required虚拟账户科目编号 */
    account_id int64 `json:"account_id";xml:"account_id"`
    
    /* bid optional账单编号 */
    bid int64 `json:"bid";xml:"bid"`
    
    /* fields required传入需要返回的字段 */
    fields string `json:"fields";xml:"fields"`
    
    /* id required账单编号 */
    id int64 `json:"id";xml:"id"`
    
}

func (req *TaobaoTaeBillGetRequest) GetAPIName() string {
	return "taobao.tae.bill.get"
}

/* TaobaoTaeBillGetResponse 查询单笔账单明细 */
type TaobaoTaeBillGetResponse struct {
    
    /* bill Object账单明细 */
    bill BillDto `json:"bill";xml:"bill"`
    
}

/* TaobaoOmniorderGuideDataGetRequest 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。 */
type TaobaoOmniorderGuideDataGetRequest struct {
    
    /* page_no required页码，从1开始 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size required每页数量，不能大于1000 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required拉取数据开始时间 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
    /* _type requireddetail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购脚骨明细;detail_sxg_order: 随心购订单明细 */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoOmniorderGuideDataGetRequest) GetAPIName() string {
	return "taobao.omniorder.guide.data.get"
}

/* TaobaoOmniorderGuideDataGetResponse 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。 */
type TaobaoOmniorderGuideDataGetResponse struct {
    
    /* data_list Basic Array拉取的数据数组，如果为空，表示数据拉取完毕。拉取的数据字段包括打点时间、商家id、商品id和门店id等，传入的类型不同，返回的字段有所不同，可以根据具体类型的返回结果具体处理 */
    data_list map[string]interface{} `json:"data_list";xml:"data_list"`
    
}

/* TaobaoTaeAccountsGetRequest tae查询费用科目信息 */
type TaobaoTaeAccountsGetRequest struct {
    
    /* aids optional需要获取的科目ID */
    aids int64 `json:"aids";xml:"aids"`
    
    /* fields required需要返回的字段 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoTaeAccountsGetRequest) GetAPIName() string {
	return "taobao.tae.accounts.get"
}

/* TaobaoTaeAccountsGetResponse tae查询费用科目信息 */
type TaobaoTaeAccountsGetResponse struct {
    
    /* accounts Object Array返回的科目信息 */
    accounts TopAccountDto `json:"accounts";xml:"accounts"`
    
    /* total_results Basic返回记录行数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoTaeBillsGetRequest tae查询账单明细 */
type TaobaoTaeBillsGetRequest struct {
    
    /* current_page optional页数,建议不要超过100页,越大性能越低,有可能会超时 */
    current_page int64 `json:"current_page";xml:"current_page"`
    
    /* fields required传入需要返回的字段,参见Bill结构体 */
    fields string `json:"fields";xml:"fields"`
    
    /* item_id optional科目编号 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* p_trade_id optional交易编号 */
    p_trade_id int64 `json:"p_trade_id";xml:"p_trade_id"`
    
    /* page_size optional每页大小,默认40条,可选范围 ：40~100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* query_date_type optional查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序 */
    query_date_type int64 `json:"query_date_type";xml:"query_date_type"`
    
    /* query_end_date required结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外) */
    query_end_date time.Time `json:"query_end_date";xml:"query_end_date"`
    
    /* query_start_date required开始时间 */
    query_start_date time.Time `json:"query_start_date";xml:"query_start_date"`
    
    /* trade_id optional子订单编号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoTaeBillsGetRequest) GetAPIName() string {
	return "taobao.tae.bills.get"
}

/* TaobaoTaeBillsGetResponse tae查询账单明细 */
type TaobaoTaeBillsGetResponse struct {
    
    /* bills Object Array账单列表 */
    bills BillDto `json:"bills";xml:"bills"`
    
    /* has_next Basic是否存在下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoOmniorderItemTagOperateRequest 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。 */
type TaobaoOmniorderItemTagOperateRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* omni_setting optional分单&接单设置 */
    omni_setting OmniSettingDto `json:"omni_setting";xml:"omni_setting"`
    
    /* status required操作状态， 填 1 代表打标，填 -1 代表去标 */
    status int64 `json:"status";xml:"status"`
    
    /* types required商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提 */
    types string `json:"types";xml:"types"`
    
}

func (req *TaobaoOmniorderItemTagOperateRequest) GetAPIName() string {
	return "taobao.omniorder.item.tag.operate"
}

/* TaobaoOmniorderItemTagOperateResponse 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。 */
type TaobaoOmniorderItemTagOperateResponse struct {
    
    /* code Basic0 正常，否则异常 */
    code string `json:"code";xml:"code"`
    
    /* message Basiccode 不为 0时有值，代表异常信息 */
    message string `json:"message";xml:"message"`
    
}

/* TmallExchangeAgreeRequest 卖家同意换货申请 */
type TmallExchangeAgreeRequest struct {
    
    /* address_id required收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id */
    address_id int64 `json:"address_id";xml:"address_id"`
    
    /* complete_address optional详细收货地址 */
    complete_address string `json:"complete_address";xml:"complete_address"`
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。当前支持的有 dispute_id, bizorder_id, modified, status */
    fields string `json:"fields";xml:"fields"`
    
    /* leave_message optional卖家留言 */
    leave_message string `json:"leave_message";xml:"leave_message"`
    
    /* leave_message_pics optional上传图片举证 */
    leave_message_pics []byte `json:"leave_message_pics";xml:"leave_message_pics"`
    
    /* mobile optional收货人手机号 */
    mobile string `json:"mobile";xml:"mobile"`
    
    /* post optional邮政编码 */
    post string `json:"post";xml:"post"`
    
}

func (req *TmallExchangeAgreeRequest) GetAPIName() string {
	return "tmall.exchange.agree"
}

/* TmallExchangeAgreeResponse 卖家同意换货申请 */
type TmallExchangeAgreeResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

/* TmallExchangeRefuseRequest 卖家拒绝换货申请 */
type TmallExchangeRefuseRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持dispute_id, bizorder_id, modified, status */
    fields string `json:"fields";xml:"fields"`
    
    /* leave_message optional拒绝换货申请时的留言 */
    leave_message string `json:"leave_message";xml:"leave_message"`
    
    /* leave_message_pics optional凭证图片 */
    leave_message_pics []byte `json:"leave_message_pics";xml:"leave_message_pics"`
    
    /* seller_refuse_reason_id required换货原因对应ID */
    seller_refuse_reason_id int64 `json:"seller_refuse_reason_id";xml:"seller_refuse_reason_id"`
    
}

func (req *TmallExchangeRefuseRequest) GetAPIName() string {
	return "tmall.exchange.refuse"
}

/* TmallExchangeRefuseResponse 卖家拒绝换货申请 */
type TmallExchangeRefuseResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

/* TmallExchangeMessageAddRequest 卖家创建换货留言 */
type TmallExchangeMessageAddRequest struct {
    
    /* content required留言内容 */
    content string `json:"content";xml:"content"`
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type */
    fields string `json:"fields";xml:"fields"`
    
    /* message_pics optional凭证图片列表 */
    message_pics []byte `json:"message_pics";xml:"message_pics"`
    
}

func (req *TmallExchangeMessageAddRequest) GetAPIName() string {
	return "tmall.exchange.message.add"
}

/* TmallExchangeMessageAddResponse 卖家创建换货留言 */
type TmallExchangeMessageAddResponse struct {
    
    /* result Object返回结果 */
    result ResultSet `json:"result";xml:"result"`
    
}

/* TmallExchangeMessagesGetRequest 查询换货订单留言列表 */
type TmallExchangeMessagesGetRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type */
    fields string `json:"fields";xml:"fields"`
    
    /* operator_roles optional留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6) */
    operator_roles int64 `json:"operator_roles";xml:"operator_roles"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TmallExchangeMessagesGetRequest) GetAPIName() string {
	return "tmall.exchange.messages.get"
}

/* TmallExchangeMessagesGetResponse 查询换货订单留言列表 */
type TmallExchangeMessagesGetResponse struct {
    
    /* result Object返回结果 */
    result RefundMessageResult `json:"result";xml:"result"`
    
}

/* TmallExchangeReturngoodsRefuseRequest 卖家拒绝买家换货申请 */
type TmallExchangeReturngoodsRefuseRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* leave_message optional留言说明 */
    leave_message string `json:"leave_message";xml:"leave_message"`
    
    /* leave_message_pics optional凭证图片 */
    leave_message_pics []byte `json:"leave_message_pics";xml:"leave_message_pics"`
    
    /* seller_refuse_reason_id required拒绝原因ID */
    seller_refuse_reason_id int64 `json:"seller_refuse_reason_id";xml:"seller_refuse_reason_id"`
    
}

func (req *TmallExchangeReturngoodsRefuseRequest) GetAPIName() string {
	return "tmall.exchange.returngoods.refuse"
}

/* TmallExchangeReturngoodsRefuseResponse 卖家拒绝买家换货申请 */
type TmallExchangeReturngoodsRefuseResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

/* TaobaoQimenWmsReturnapplyReportRequest 退货异常包裹单通知接口 */
type TaobaoQimenWmsReturnapplyReportRequest struct {
    
    /* request optional请求对象 */
    request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenWmsReturnapplyReportRequest) GetAPIName() string {
	return "taobao.qimen.wms.returnapply.report"
}

/* TaobaoQimenWmsReturnapplyReportResponse 退货异常包裹单通知接口 */
type TaobaoQimenWmsReturnapplyReportResponse struct {
    
    /* response Object响应对象 */
    response Response `json:"response";xml:"response"`
    
}

/* CainiaoSmartdeliverySellerStatusIQueryRequest 查询智能发货引擎商家状态信息 */
type CainiaoSmartdeliverySellerStatusIQueryRequest struct {
    
}

func (req *CainiaoSmartdeliverySellerStatusIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.seller.status.i.query"
}

/* CainiaoSmartdeliverySellerStatusIQueryResponse 查询智能发货引擎商家状态信息 */
type CainiaoSmartdeliverySellerStatusIQueryResponse struct {
    
    /* seller_status Object商家状态 */
    seller_status SellerStatus `json:"seller_status";xml:"seller_status"`
    
}

/* TmallExchangeConsigngoodsRequest 卖家发货 */
type TmallExchangeConsigngoodsRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段 */
    fields string `json:"fields";xml:"fields"`
    
    /* logistics_company_name required卖家发货的快递公司 */
    logistics_company_name string `json:"logistics_company_name";xml:"logistics_company_name"`
    
    /* logistics_no required卖家发货的物流单号 */
    logistics_no string `json:"logistics_no";xml:"logistics_no"`
    
    /* logistics_type required卖家发货的物流类型，100表示平邮，200表示快递 */
    logistics_type int64 `json:"logistics_type";xml:"logistics_type"`
    
}

func (req *TmallExchangeConsigngoodsRequest) GetAPIName() string {
	return "tmall.exchange.consigngoods"
}

/* TmallExchangeConsigngoodsResponse 卖家发货 */
type TmallExchangeConsigngoodsResponse struct {
    
    /* result Objectresult */
    result RefundBaseResponse `json:"result";xml:"result"`
    
}

/* TaobaoTmcAuthGetRequest TMC连接授权Token */
type TaobaoTmcAuthGetRequest struct {
    
    /* group optionaltmc组名 */
    group string `json:"group";xml:"group"`
    
}

func (req *TaobaoTmcAuthGetRequest) GetAPIName() string {
	return "taobao.tmc.auth.get"
}

/* TaobaoTmcAuthGetResponse TMC连接授权Token */
type TaobaoTmcAuthGetResponse struct {
    
    /* result Basicresult */
    result string `json:"result";xml:"result"`
    
}

/* TaobaoInventoryMerchantAdjustRequest 货品库存商家端调整 ，入库，出库，盘点 */
type TaobaoInventoryMerchantAdjustRequest struct {
    
    /* inventory_check required调整库存对象 */
    inventory_check InventoryCheckDto `json:"inventory_check";xml:"inventory_check"`
    
}

func (req *TaobaoInventoryMerchantAdjustRequest) GetAPIName() string {
	return "taobao.inventory.merchant.adjust"
}

/* TaobaoInventoryMerchantAdjustResponse 货品库存商家端调整 ，入库，出库，盘点 */
type TaobaoInventoryMerchantAdjustResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsSkuCreateRequest 商品同步 */
type TaobaoWlbWmsSkuCreateRequest struct {
    
    /* advent_lifecycle optional保质期预警天数 */
    advent_lifecycle int64 `json:"advent_lifecycle";xml:"advent_lifecycle"`
    
    /* approval_number optional批准文号 */
    approval_number string `json:"approval_number";xml:"approval_number"`
    
    /* bar_code optional条形码，多条码请用”；”分隔； */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* brand optional品牌编码 */
    brand string `json:"brand";xml:"brand"`
    
    /* brand_name optional品牌名称 */
    brand_name string `json:"brand_name";xml:"brand_name"`
    
    /* category optional商品类别编码（外部系统类别） */
    category string `json:"category";xml:"category"`
    
    /* category_name optional商品类别名称 */
    category_name string `json:"category_name";xml:"category_name"`
    
    /* color optional颜色 */
    color string `json:"color";xml:"color"`
    
    /* cost_price optional成本价，单位分 */
    cost_price int64 `json:"cost_price";xml:"cost_price"`
    
    /* extend_fields optional拓展属性 */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* gross_weight optional毛重，单位克 */
    gross_weight int64 `json:"gross_weight";xml:"gross_weight"`
    
    /* height optional高度，单位毫米 */
    height int64 `json:"height";xml:"height"`
    
    /* is_area_sale optional是否区域销售 */
    is_area_sale bool `json:"is_area_sale";xml:"is_area_sale"`
    
    /* is_batch_mgt optional是否启用批次管理 */
    is_batch_mgt bool `json:"is_batch_mgt";xml:"is_batch_mgt"`
    
    /* is_danger optional是否危险品 */
    is_danger bool `json:"is_danger";xml:"is_danger"`
    
    /* is_hygroscopic optional是否易碎品 */
    is_hygroscopic bool `json:"is_hygroscopic";xml:"is_hygroscopic"`
    
    /* is_shelflife optional是否启用保质期管理 */
    is_shelflife bool `json:"is_shelflife";xml:"is_shelflife"`
    
    /* is_sn_mgt optional是否启用序列号管理 */
    is_sn_mgt bool `json:"is_sn_mgt";xml:"is_sn_mgt"`
    
    /* item_code required商家商品编码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* item_id optional商家商品ID */
    item_id string `json:"item_id";xml:"item_id"`
    
    /* item_price optional零售价，单位分 */
    item_price int64 `json:"item_price";xml:"item_price"`
    
    /* length optional长度，单位毫米 */
    length int64 `json:"length";xml:"length"`
    
    /* lifecycle optional商品保质期天数 */
    lifecycle int64 `json:"lifecycle";xml:"lifecycle"`
    
    /* lockup_lifecycle optional保质期禁售天数 */
    lockup_lifecycle int64 `json:"lockup_lifecycle";xml:"lockup_lifecycle"`
    
    /* name required商品名称 */
    name string `json:"name";xml:"name"`
    
    /* net_weight optional净重，单位克 */
    net_weight int64 `json:"net_weight";xml:"net_weight"`
    
    /* origin_address optional产地 */
    origin_address int64 `json:"origin_address";xml:"origin_address"`
    
    /* pcs optional箱规 */
    pcs int64 `json:"pcs";xml:"pcs"`
    
    /* reject_lifecycle optional保质期禁收天数 */
    reject_lifecycle int64 `json:"reject_lifecycle";xml:"reject_lifecycle"`
    
    /* size optional尺码 */
    size string `json:"size";xml:"size"`
    
    /* specification optional规格 */
    specification string `json:"specification";xml:"specification"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tag_price optional吊牌价，单位分 */
    tag_price int64 `json:"tag_price";xml:"tag_price"`
    
    /* title optional商品标题 */
    title string `json:"title";xml:"title"`
    
    /* _type required商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他) */
    _type string `json:"type";xml:"type"`
    
    /* use_yn required启用标识 */
    use_yn bool `json:"use_yn";xml:"use_yn"`
    
    /* volume optional体积，单位立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* width optional宽度，单位毫米 */
    width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbWmsSkuCreateRequest) GetAPIName() string {
	return "taobao.wlb.wms.sku.create"
}

/* TaobaoWlbWmsSkuCreateResponse 商品同步 */
type TaobaoWlbWmsSkuCreateResponse struct {
    
    /* item_id Basic系统自动生成 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* wl_error_code Basic错误码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoWlbWmsSkuUpdateRequest 商品信息的更新 */
type TaobaoWlbWmsSkuUpdateRequest struct {
    
    /* advent_lifecycle optional保质期预警天数 */
    advent_lifecycle int64 `json:"advent_lifecycle";xml:"advent_lifecycle"`
    
    /* approval_number optional批准文号 */
    approval_number string `json:"approval_number";xml:"approval_number"`
    
    /* attribute optional商品属性编码 */
    attribute string `json:"attribute";xml:"attribute"`
    
    /* bar_code optional条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002 */
    bar_code string `json:"bar_code";xml:"bar_code"`
    
    /* brand optional品牌编码 */
    brand string `json:"brand";xml:"brand"`
    
    /* brand_name optional品牌名称 */
    brand_name string `json:"brand_name";xml:"brand_name"`
    
    /* category optional商品类别编码（外部系统类别） */
    category string `json:"category";xml:"category"`
    
    /* category_name optional商品类别名称 */
    category_name string `json:"category_name";xml:"category_name"`
    
    /* color optional颜色 */
    color string `json:"color";xml:"color"`
    
    /* cost_price optional成本价，单位分 */
    cost_price int64 `json:"cost_price";xml:"cost_price"`
    
    /* extend_fields optional拓展属性 */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* gross_weight optional毛重，单位克 */
    gross_weight int64 `json:"gross_weight";xml:"gross_weight"`
    
    /* height optional高度，单位毫米 */
    height int64 `json:"height";xml:"height"`
    
    /* is_area_sale optional是否区域销售属性 */
    is_area_sale bool `json:"is_area_sale";xml:"is_area_sale"`
    
    /* is_batch_mgt optional是否启用批次管理 */
    is_batch_mgt bool `json:"is_batch_mgt";xml:"is_batch_mgt"`
    
    /* is_danger optional是否危险品 */
    is_danger bool `json:"is_danger";xml:"is_danger"`
    
    /* is_hygroscopic optional是否易碎品 */
    is_hygroscopic bool `json:"is_hygroscopic";xml:"is_hygroscopic"`
    
    /* is_shelflife optional是否启用保质期管理 */
    is_shelflife bool `json:"is_shelflife";xml:"is_shelflife"`
    
    /* is_sn_mgt optional是否启用序列号管理 */
    is_sn_mgt bool `json:"is_sn_mgt";xml:"is_sn_mgt"`
    
    /* item_id required外部系统ID */
    item_id string `json:"item_id";xml:"item_id"`
    
    /* item_price optional零售价，单位分 */
    item_price int64 `json:"item_price";xml:"item_price"`
    
    /* length optional长度，单位毫米 */
    length int64 `json:"length";xml:"length"`
    
    /* lifecycle optional商品保质期天数 */
    lifecycle int64 `json:"lifecycle";xml:"lifecycle"`
    
    /* lockup_lifecycle optional保质期禁售天数 */
    lockup_lifecycle int64 `json:"lockup_lifecycle";xml:"lockup_lifecycle"`
    
    /* name optional商品名称 */
    name string `json:"name";xml:"name"`
    
    /* net_weight optional净重，单位克 */
    net_weight int64 `json:"net_weight";xml:"net_weight"`
    
    /* origin_address optional产地 */
    origin_address int64 `json:"origin_address";xml:"origin_address"`
    
    /* pcs optional箱规 */
    pcs int64 `json:"pcs";xml:"pcs"`
    
    /* reject_lifecycle optional保质期禁收天数 */
    reject_lifecycle int64 `json:"reject_lifecycle";xml:"reject_lifecycle"`
    
    /* size optional尺码 */
    size string `json:"size";xml:"size"`
    
    /* specification optional规格 */
    specification string `json:"specification";xml:"specification"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tag_price optional吊牌价，单位分 */
    tag_price int64 `json:"tag_price";xml:"tag_price"`
    
    /* title optional商品标题 */
    title string `json:"title";xml:"title"`
    
    /* _type optional商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品 */
    _type string `json:"type";xml:"type"`
    
    /* use_yn required启用标识 */
    use_yn bool `json:"use_yn";xml:"use_yn"`
    
    /* volume optional体积，单位立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* width optional宽度，单位毫米 */
    width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbWmsSkuUpdateRequest) GetAPIName() string {
	return "taobao.wlb.wms.sku.update"
}

/* TaobaoWlbWmsSkuUpdateResponse 商品信息的更新 */
type TaobaoWlbWmsSkuUpdateResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoWlbWmsStockOutOrderNotifyRequest 出库单通知 */
type TaobaoWlbWmsStockOutOrderNotifyRequest struct {
    
    /* car_no optional车牌号 */
    car_no string `json:"car_no";xml:"car_no"`
    
    /* carriers_name optional承运商名称 */
    carriers_name string `json:"carriers_name";xml:"carriers_name"`
    
    /* extend_fields optional拓展属性 */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* order_code requiredERP单据ID */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time required订单创建时间 */
    order_create_time time.Time `json:"order_create_time";xml:"order_create_time"`
    
    /* order_item_list optional订单商品信息列表 */
    order_item_list Orderitemlistwlbwmsstockoutordernotify `json:"order_item_list";xml:"order_item_list"`
    
    /* order_type required单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* outbound_type_desc optionalERP可选择性文本透传至WMS */
    outbound_type_desc string `json:"outbound_type_desc";xml:"outbound_type_desc"`
    
    /* pick_call optional取件人电话 */
    pick_call string `json:"pick_call";xml:"pick_call"`
    
    /* pick_id optional取件人身份证ID */
    pick_id string `json:"pick_id";xml:"pick_id"`
    
    /* pick_name optional取件人姓名 */
    pick_name string `json:"pick_name";xml:"pick_name"`
    
    /* prev_order_code optional前物流订单号，如退货入库单可能会用到 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收件人信息 */
    receiver_info Receiverwlbwmsstockoutordernotify `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* send_time optional要求出库日期 */
    send_time time.Time `json:"send_time";xml:"send_time"`
    
    /* sender_info optional发货方信息 */
    sender_info Senderwlbwmsstockoutordernotify `json:"sender_info";xml:"sender_info"`
    
    /* store_code required仓储编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* transport_mode optional出库方式 */
    transport_mode string `json:"transport_mode";xml:"transport_mode"`
    
}

func (req *TaobaoWlbWmsStockOutOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.out.order.notify"
}

/* TaobaoWlbWmsStockOutOrderNotifyResponse 出库单通知 */
type TaobaoWlbWmsStockOutOrderNotifyResponse struct {
    
    /* order_code Basic仓储订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* CainiaoWaybillIiGetRequest 菜鸟电子面单的云打印申请电子面单号的方法 */
type CainiaoWaybillIiGetRequest struct {
    
    /* param_waybill_cloud_print_apply_new_request required入参信息 */
    param_waybill_cloud_print_apply_new_request WaybillCloudPrintApplyNewRequest `json:"param_waybill_cloud_print_apply_new_request";xml:"param_waybill_cloud_print_apply_new_request"`
    
}

func (req *CainiaoWaybillIiGetRequest) GetAPIName() string {
	return "cainiao.waybill.ii.get"
}

/* CainiaoWaybillIiGetResponse 菜鸟电子面单的云打印申请电子面单号的方法 */
type CainiaoWaybillIiGetResponse struct {
    
    /* modules Object Array系统自动生成 */
    modules WaybillCloudPrintResponse `json:"modules";xml:"modules"`
    
}

/* TaobaoWlbWmsConsignOrderNotifyRequest 发货订单通知 */
type TaobaoWlbWmsConsignOrderNotifyRequest struct {
    
    /* alipay_no optional废弃，支付宝交易号 */
    alipay_no string `json:"alipay_no";xml:"alipay_no"`
    
    /* ar_amount optional订单应收金额，消费者还需要付的金额，单位分 */
    ar_amount int64 `json:"ar_amount";xml:"ar_amount"`
    
    /* car_no optional废弃，车牌号 */
    car_no string `json:"car_no";xml:"car_no"`
    
    /* carriers_name optional废弃，承运商名称 */
    carriers_name string `json:"carriers_name";xml:"carriers_name"`
    
    /* deliver_requirements optional配送要求 */
    deliver_requirements Deliverrequirementswlbwmsconsignordernotify `json:"deliver_requirements";xml:"deliver_requirements"`
    
    /* discount_amount optional订单优惠金额，整单优惠金额，单位分 */
    discount_amount int64 `json:"discount_amount";xml:"discount_amount"`
    
    /* extend_fields optional拓展属性 */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* got_amount optional订单已付金额，消费者已经支付的金额，单位分 */
    got_amount int64 `json:"got_amount";xml:"got_amount"`
    
    /* invoice_info_list optional发票信息列表 */
    invoice_info_list Invoicelistwlbwmsconsignordernotify `json:"invoice_info_list";xml:"invoice_info_list"`
    
    /* order_amount optional订单总金额,=总商品金额-订单优惠金额+快递费用，单位分 */
    order_amount int64 `json:"order_amount";xml:"order_amount"`
    
    /* order_code requiredERP订单号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time optional订单创建时间,ERP创建订单时间 */
    order_create_time time.Time `json:"order_create_time";xml:"order_create_time"`
    
    /* order_examination_time optional订单审核时间,ERP创建支付时间 */
    order_examination_time time.Time `json:"order_examination_time";xml:"order_examination_time"`
    
    /* order_flag optional订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票) */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list optional订单商品信息列表 */
    order_item_list Orderitemlistwlbwmsconsignordernotify `json:"order_item_list";xml:"order_item_list"`
    
    /* order_pay_time optional订单支付时间 */
    order_pay_time time.Time `json:"order_pay_time";xml:"order_pay_time"`
    
    /* order_priority optional废弃，订单优先级0 -10，根据订单作业优先级设置，数字越大，优先级越高 */
    order_priority int64 `json:"order_priority";xml:"order_priority"`
    
    /* order_shop_create_time optional下单时间，订单在交易平台创建时间 */
    order_shop_create_time time.Time `json:"order_shop_create_time";xml:"order_shop_create_time"`
    
    /* order_source optional订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他） */
    order_source int64 `json:"order_source";xml:"order_source"`
    
    /* order_sub_source optional废弃，交易内部来源，文本透传 WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算) 一笔订单可能同时有以上多个标记，则以逗号分隔 */
    order_sub_source string `json:"order_sub_source";xml:"order_sub_source"`
    
    /* order_type required单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* picker_call optional废弃，取件人电话 */
    picker_call string `json:"picker_call";xml:"picker_call"`
    
    /* picker_id optional废弃，取件人身份证 */
    picker_id string `json:"picker_id";xml:"picker_id"`
    
    /* picker_name optional废弃，取件人姓名 */
    picker_name string `json:"picker_name";xml:"picker_name"`
    
    /* postfee optional快递费用，单位分 */
    postfee int64 `json:"postfee";xml:"postfee"`
    
    /* prev_order_code optional前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收件人信息 */
    receiver_info Receiverwlbwmsconsignordernotify `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* sender_info optional发货方信息 */
    sender_info Senderwlbwmsconsignordernotify `json:"sender_info";xml:"sender_info"`
    
    /* service_fee optionalCOD服务费，单位分 */
    service_fee int64 `json:"service_fee";xml:"service_fee"`
    
    /* store_code optional仓库编码，此字段为空时，由菜鸟路由仓库发货 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tms_service_code optional快递公司编码，此字段为空时，由菜鸟选择快递配送 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* tms_service_name optional快递公司名称 */
    tms_service_name string `json:"tms_service_name";xml:"tms_service_name"`
    
    /* transport_mode optional废弃，出库方式（自提，快递，销毁） */
    transport_mode string `json:"transport_mode";xml:"transport_mode"`
    
}

func (req *TaobaoWlbWmsConsignOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.order.notify"
}

/* TaobaoWlbWmsConsignOrderNotifyResponse 发货订单通知 */
type TaobaoWlbWmsConsignOrderNotifyResponse struct {
    
    /* consign_order_list Object Array系统自动生成 */
    consign_order_list Consignorderlist `json:"consign_order_list";xml:"consign_order_list"`
    
    /* order_code Basic订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* CainiaoWaybillIiUpdateRequest 商家更新电子面单号对应的面单信息。 */
type CainiaoWaybillIiUpdateRequest struct {
    
    /* param_waybill_cloud_print_update_request required更新请求信息 */
    param_waybill_cloud_print_update_request WaybillCloudPrintUpdateRequest `json:"param_waybill_cloud_print_update_request";xml:"param_waybill_cloud_print_update_request"`
    
}

func (req *CainiaoWaybillIiUpdateRequest) GetAPIName() string {
	return "cainiao.waybill.ii.update"
}

/* CainiaoWaybillIiUpdateResponse 商家更新电子面单号对应的面单信息。 */
type CainiaoWaybillIiUpdateResponse struct {
    
    /* print_data Basic模板内容 */
    print_data string `json:"print_data";xml:"print_data"`
    
    /* waybill_code Basic面单号 */
    waybill_code string `json:"waybill_code";xml:"waybill_code"`
    
}

/* TaobaoWlbWmsOrderCancelNotifyRequest 单据取消接口 */
type TaobaoWlbWmsOrderCancelNotifyRequest struct {
    
    /* order_code required订单类型 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_type required单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbWmsOrderCancelNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.order.cancel.notify"
}

/* TaobaoWlbWmsOrderCancelNotifyResponse 单据取消接口 */
type TaobaoWlbWmsOrderCancelNotifyResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* CainiaoEndpointLockerTopOrderTrackingNewRequest 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。 */
type CainiaoEndpointLockerTopOrderTrackingNewRequest struct {
    
    /* track_info required回传信息 */
    track_info CollectTrackingInfo `json:"track_info";xml:"track_info"`
    
}

func (req *CainiaoEndpointLockerTopOrderTrackingNewRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.tracking.new"
}

/* CainiaoEndpointLockerTopOrderTrackingNewResponse 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。 */
type CainiaoEndpointLockerTopOrderTrackingNewResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TaobaoQimenEntryorderQueryRequest ERP调用接口，查询入库单信息; */
type TaobaoQimenEntryorderQueryRequest struct {
    
    /* entryOrderCode required入库单编码 */
    entryOrderCode string `json:"entryOrderCode";xml:"entryOrderCode"`
    
    /* entryOrderId optional仓储系统入库单ID */
    entryOrderId string `json:"entryOrderId";xml:"entryOrderId"`
    
    /* extOrderCode optionalextOrderCode */
    extOrderCode string `json:"extOrderCode";xml:"extOrderCode"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderType optionalorderType */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页(从1开始) */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenEntryorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.query"
}

/* TaobaoQimenEntryorderQueryResponse ERP调用接口，查询入库单信息; */
type TaobaoQimenEntryorderQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* entryOrder Object入库单信息 */
    entryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array入库单单据信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
}

/* TmallExchangeGetRequest 获取单笔换货详情 */
type TmallExchangeGetRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeGetRequest) GetAPIName() string {
	return "tmall.exchange.get"
}

/* TmallExchangeGetResponse 获取单笔换货详情 */
type TmallExchangeGetResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

/* TaobaoQimenDeliveryorderQueryRequest ERP调用奇门的发货单查询接口，查询发货单详情 */
type TaobaoQimenDeliveryorderQueryRequest struct {
    
    /* deliveryOrderCode optional奇门仓储字段,说明,string(50),, */
    deliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    /* deliveryOrderId optional奇门仓储字段,说明,string(50),, */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required发库单号 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统发库单号 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderSourceCode optional交易单号 */
    orderSourceCode string `json:"orderSourceCode";xml:"orderSourceCode"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* page required当前页 */
    page int64 `json:"page";xml:"page"`
    
    /* pageSize required每页orderLine条数(最多100条) */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* status optional奇门仓储字段,说明,string(50),, */
    status string `json:"status";xml:"status"`
    
    /* warehouseCode optional仓库编码 */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenDeliveryorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.query"
}

/* TaobaoQimenDeliveryorderQueryResponse ERP调用奇门的发货单查询接口，查询发货单详情 */
type TaobaoQimenDeliveryorderQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* deliveryOrder Object发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orderLines Object Array单据列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages Object Array包裹信息 */
    packages Package `json:"packages";xml:"packages"`
    
    /* totalLines BasicorderLines总条数 */
    totalLines int64 `json:"totalLines";xml:"totalLines"`
    
}

/* CainiaoEndpointLockerTopOrderNoticeRequest 合作公司对订单手动触发短信，有次数限制 */
type CainiaoEndpointLockerTopOrderNoticeRequest struct {
    
    /* mail_no required运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* order_code required合作公司唯一订单编号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* scene_code required场景编号：0：重发短信，1：催取短信 */
    scene_code int64 `json:"scene_code";xml:"scene_code"`
    
    /* station_id required站点ID */
    station_id string `json:"station_id";xml:"station_id"`
    
}

func (req *CainiaoEndpointLockerTopOrderNoticeRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.notice"
}

/* CainiaoEndpointLockerTopOrderNoticeResponse 合作公司对订单手动触发短信，有次数限制 */
type CainiaoEndpointLockerTopOrderNoticeResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}

/* CainiaoEndpointLockerTopOrderNoticesendQueryRequest 合作公司查询消息发送的接口，判断是否裹裹发送消息 */
type CainiaoEndpointLockerTopOrderNoticesendQueryRequest struct {
    
    /* getter_phone optional收件人手机号 */
    getter_phone string `json:"getter_phone";xml:"getter_phone"`
    
    /* mail_no required运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* station_id required站点id */
    station_id string `json:"station_id";xml:"station_id"`
    
}

func (req *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.noticesend.query"
}

/* CainiaoEndpointLockerTopOrderNoticesendQueryResponse 合作公司查询消息发送的接口，判断是否裹裹发送消息 */
type CainiaoEndpointLockerTopOrderNoticesendQueryResponse struct {
    
    /* result Object返回结果 */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TmallExchangeReturngoodsAgreeRequest 卖家确认收货 */
type TmallExchangeReturngoodsAgreeRequest struct {
    
    /* dispute_id required换货单号ID */
    dispute_id int64 `json:"dispute_id";xml:"dispute_id"`
    
    /* fields required返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态） */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeReturngoodsAgreeRequest) GetAPIName() string {
	return "tmall.exchange.returngoods.agree"
}

/* TmallExchangeReturngoodsAgreeResponse 卖家确认收货 */
type TmallExchangeReturngoodsAgreeResponse struct {
    
    /* result Object返回结果 */
    result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

/* CainiaoCntmsLogisticsOrderConsignRequest 商家包装打印面单结束后，通知菜鸟包裹要发货 */
type CainiaoCntmsLogisticsOrderConsignRequest struct {
    
    /* content optional配送发货信息 */
    content CnTmsLogisticsOrderConsignContent `json:"content";xml:"content"`
    
}

func (req *CainiaoCntmsLogisticsOrderConsignRequest) GetAPIName() string {
	return "cainiao.cntms.logistics.order.consign"
}

/* CainiaoCntmsLogisticsOrderConsignResponse 商家包装打印面单结束后，通知菜鸟包裹要发货 */
type CainiaoCntmsLogisticsOrderConsignResponse struct {
    
    /* logistics_order_code Basic物流单号 */
    logistics_order_code string `json:"logistics_order_code";xml:"logistics_order_code"`
    
}

/* TmallExchangeReceiveGetRequest 卖家查询换货列表 */
type TmallExchangeReceiveGetRequest struct {
    
    /* biz_order_id optional正向订单号 */
    biz_order_id int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    /* buyer_id optional买家id */
    buyer_id int64 `json:"buyer_id";xml:"buyer_id"`
    
    /* buyer_nick optional买家昵称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* dispute_status_array optional换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14) */
    dispute_status_array int64 `json:"dispute_status_array";xml:"dispute_status_array"`
    
    /* end_created_time optional查询申请时间段的结束时间点 */
    end_created_time time.Time `json:"end_created_time";xml:"end_created_time"`
    
    /* end_gmt_modifed_time optional查询修改时间段的结束时间点 */
    end_gmt_modifed_time time.Time `json:"end_gmt_modifed_time";xml:"end_gmt_modifed_time"`
    
    /* fields required返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick */
    fields string `json:"fields";xml:"fields"`
    
    /* logistic_no optional快递单号 */
    logistic_no string `json:"logistic_no";xml:"logistic_no"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* refund_id_array optional退款单号ID列表，最多只能输入20个id */
    refund_id_array int64 `json:"refund_id_array";xml:"refund_id_array"`
    
    /* start_created_time optional查询申请时间段的开始时间点 */
    start_created_time time.Time `json:"start_created_time";xml:"start_created_time"`
    
    /* start_gmt_modified_time optional查询修改时间段的开始时间点 */
    start_gmt_modified_time time.Time `json:"start_gmt_modified_time";xml:"start_gmt_modified_time"`
    
}

func (req *TmallExchangeReceiveGetRequest) GetAPIName() string {
	return "tmall.exchange.receive.get"
}

/* TmallExchangeReceiveGetResponse 卖家查询换货列表 */
type TmallExchangeReceiveGetResponse struct {
    
    /* error_codes Basic错误码 */
    error_codes string `json:"error_codes";xml:"error_codes"`
    
    /* error_msg Basic错误信息 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* exception Basic所抛出异常 */
    exception map[string]interface{} `json:"exception";xml:"exception"`
    
    /* has_next Basic是否还有下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* page_results Basic当前页的换货单数量 */
    page_results int64 `json:"page_results";xml:"page_results"`
    
    /* results Object Array返回结果 */
    results Exchange `json:"results";xml:"results"`
    
    /* total_results Basic所有符合查询条件的换货单的数量 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* CainiaoCntmsMailnoGetRequest 打印面单时，通过此接口获取面单号及打打印信息 */
type CainiaoCntmsMailnoGetRequest struct {
    
    /* content optional获取菜鸟配送电子面单请求参数 */
    content CnTmsMailnoGetContent `json:"content";xml:"content"`
    
}

func (req *CainiaoCntmsMailnoGetRequest) GetAPIName() string {
	return "cainiao.cntms.mailno.get"
}

/* CainiaoCntmsMailnoGetResponse 打印面单时，通过此接口获取面单号及打打印信息 */
type CainiaoCntmsMailnoGetResponse struct {
    
    /* allocator_code Basic揽货商（分拨中心）编码 */
    allocator_code string `json:"allocator_code";xml:"allocator_code"`
    
    /* allocator_name Basic揽货商（分拨中心）名称 */
    allocator_name string `json:"allocator_name";xml:"allocator_name"`
    
    /* logistics_code Basic物流商公司编码
（ERP在调用菜鸟发货接口时此字段赋值到tms_code, 调用淘宝“自己联系物流（线下物流）发货”时，做为company_code传入） */
    logistics_code string `json:"logistics_code";xml:"logistics_code"`
    
    /* logistics_name Basic物流商名称 */
    logistics_name string `json:"logistics_name";xml:"logistics_name"`
    
    /* mailno Basic面单号 */
    mailno string `json:"mailno";xml:"mailno"`
    
    /* package_center_code Basic集包地代码 */
    package_center_code string `json:"package_center_code";xml:"package_center_code"`
    
    /* package_center_name Basic集包地名称 */
    package_center_name string `json:"package_center_name";xml:"package_center_name"`
    
    /* sec_distribution Basic二级流向信息 */
    sec_distribution string `json:"sec_distribution";xml:"sec_distribution"`
    
    /* short_address Basic一级流向信息，大头笔 */
    short_address string `json:"short_address";xml:"short_address"`
    
    /* tms_code Basic二级配送公司编码 */
    tms_code string `json:"tms_code";xml:"tms_code"`
    
    /* tms_name Basic二级配送公司名称 */
    tms_name string `json:"tms_name";xml:"tms_name"`
    
}

/* TaobaoLocationRelationQueryRequest 地点关联关系查询 
门店和仓库关联关系查询 */
type TaobaoLocationRelationQueryRequest struct {
    
    /* location_relation required关系查询 */
    location_relation LocationRelationDto `json:"location_relation";xml:"location_relation"`
    
}

func (req *TaobaoLocationRelationQueryRequest) GetAPIName() string {
	return "taobao.location.relation.query"
}

/* TaobaoLocationRelationQueryResponse 地点关联关系查询 
门店和仓库关联关系查询 */
type TaobaoLocationRelationQueryResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TaobaoLocationRelationEditRequest 地点关联关系增量编辑 */
type TaobaoLocationRelationEditRequest struct {
    
    /* location_relation_list required关系对象列表 */
    location_relation_list LocationRelationDto `json:"location_relation_list";xml:"location_relation_list"`
    
}

func (req *TaobaoLocationRelationEditRequest) GetAPIName() string {
	return "taobao.location.relation.edit"
}

/* TaobaoLocationRelationEditResponse 地点关联关系增量编辑 */
type TaobaoLocationRelationEditResponse struct {
    
    /* error_msg Basic错误信息 */
    error_msg string `json:"error_msg";xml:"error_msg"`
    
    /* errorcode Basic错误码 */
    errorcode string `json:"errorcode";xml:"errorcode"`
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* CainiaoCloudprintCmdprintRenderRequest isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。 */
type CainiaoCloudprintCmdprintRenderRequest struct {
    
    /* params required参数对象 */
    params CmdRenderParams `json:"params";xml:"params"`
    
}

func (req *CainiaoCloudprintCmdprintRenderRequest) GetAPIName() string {
	return "cainiao.cloudprint.cmdprint.render"
}

/* CainiaoCloudprintCmdprintRenderResponse isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。 */
type CainiaoCloudprintCmdprintRenderResponse struct {
    
    /* cmd_content Basic指令集内容串 */
    cmd_content string `json:"cmd_content";xml:"cmd_content"`
    
    /* cmd_encoding Basic指令集编码方式：origin-原串 gzip-采用gzip压缩并base64编码 */
    cmd_encoding string `json:"cmd_encoding";xml:"cmd_encoding"`
    
    /* ret_code Basic0成功，非0失败 */
    ret_code string `json:"ret_code";xml:"ret_code"`
    
    /* ret_msg Basic错误信息 */
    ret_msg string `json:"ret_msg";xml:"ret_msg"`
    
}

/* TaobaoWlbWmsReturnBillGetRequest 通过订单号获取单个销退单收货信息。 */
type TaobaoWlbWmsReturnBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsReturnBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.return.bill.get"
}

/* TaobaoWlbWmsReturnBillGetResponse 通过订单号获取单个销退单收货信息。 */
type TaobaoWlbWmsReturnBillGetResponse struct {
    
    /* return_order_info Object回退订单信息 */
    return_order_info CainiaoReturnBillReturnorderinfo `json:"return_order_info";xml:"return_order_info"`
    
}

/* TaobaoWlbWmsStockInBillGetRequest 获取入库单信息 */
type TaobaoWlbWmsStockInBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsStockInBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.in.bill.get"
}

/* TaobaoWlbWmsStockInBillGetResponse 获取入库单信息 */
type TaobaoWlbWmsStockInBillGetResponse struct {
    
    /* stock_in_info Object入库单信息 */
    stock_in_info CainiaoStockInBillStockininfo `json:"stock_in_info";xml:"stock_in_info"`
    
}

/* TaobaoOcTradesBytagGetRequest 根据标签查询订单编号 */
type TaobaoOcTradesBytagGetRequest struct {
    
    /* page optional当前页 */
    page int64 `json:"page";xml:"page"`
    
    /* page_size optional分页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* tag_name required标签名称 */
    tag_name string `json:"tag_name";xml:"tag_name"`
    
    /* tag_type required标签类型，1官方，2自定义 */
    tag_type int64 `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoOcTradesBytagGetRequest) GetAPIName() string {
	return "taobao.oc.trades.bytag.get"
}

/* TaobaoOcTradesBytagGetResponse 根据标签查询订单编号 */
type TaobaoOcTradesBytagGetResponse struct {
    
    /* tids Basic Array打了该标签的订单编号列表 */
    tids int64 `json:"tids";xml:"tids"`
    
    /* totals Basic总数 */
    totals int64 `json:"totals";xml:"totals"`
    
}

/* TaobaoOcTradetagAttachRequest 对订单添加标签和更新标签。标签分为官方标签和自定义标签。
官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731
自定义标签有2个通用属性:
    `show_str:给消费者显示的字符串（如果可以显示的话）
    `pic_urls:图片url,地址必须是图片空间的url,最多5张 */
type TaobaoOcTradetagAttachRequest struct {
    
    /* tag_name required标签名称 */
    tag_name string `json:"tag_name";xml:"tag_name"`
    
    /* tag_type optional标签类型       1：官方标签      2：自定义标签 */
    tag_type int64 `json:"tag_type";xml:"tag_type"`
    
    /* tag_value required标签值，json格式 */
    tag_value string `json:"tag_value";xml:"tag_value"`
    
    /* tid required订单id */
    tid int64 `json:"tid";xml:"tid"`
    
    /* visible optional该标签在消费者端是否显示,0:不显示,1：显示 */
    visible int64 `json:"visible";xml:"visible"`
    
}

func (req *TaobaoOcTradetagAttachRequest) GetAPIName() string {
	return "taobao.oc.tradetag.attach"
}

/* TaobaoOcTradetagAttachResponse 对订单添加标签和更新标签。标签分为官方标签和自定义标签。
官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731
自定义标签有2个通用属性:
    `show_str:给消费者显示的字符串（如果可以显示的话）
    `pic_urls:图片url,地址必须是图片空间的url,最多5张 */
type TaobaoOcTradetagAttachResponse struct {
    
    /* result Basic操作成功或者操作失败 */
    result bool `json:"result";xml:"result"`
    
}

/* TaobaoWlbOrderdetailDateGetRequest 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情 */
type TaobaoWlbOrderdetailDateGetRequest struct {
    
    /* end_time required创建时间结束 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* page_no optional分页下标 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页大小 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required创建时间起始 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoWlbOrderdetailDateGetRequest) GetAPIName() string {
	return "taobao.wlb.orderdetail.date.get"
}

/* TaobaoWlbOrderdetailDateGetResponse 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情 */
type TaobaoWlbOrderdetailDateGetResponse struct {
    
    /* order_detail_list Object Array物流宝订单，并且包含订单详情 */
    order_detail_list WlbOrderDetail `json:"order_detail_list";xml:"order_detail_list"`
    
    /* total_count Basic总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoOcTradetagsGetRequest 根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/> */
type TaobaoOcTradetagsGetRequest struct {
    
    /* history optional是否查询历史标签 */
    history int64 `json:"history";xml:"history"`
    
    /* tag_names optional不填，则不做标签名称过滤 */
    tag_names string `json:"tag_names";xml:"tag_names"`
    
    /* tag_types optional不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签 */
    tag_types string `json:"tag_types";xml:"tag_types"`
    
    /* tid required交易主订单id */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoOcTradetagsGetRequest) GetAPIName() string {
	return "taobao.oc.tradetags.get"
}

/* TaobaoOcTradetagsGetResponse 根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/> */
type TaobaoOcTradetagsGetResponse struct {
    
    /* trade_tags Object Array返回结果 */
    trade_tags TradeTagRelationDo `json:"trade_tags";xml:"trade_tags"`
    
}

/* QimenTaobaoCrmOrderSyncRequest CRM对于会员数据分析需要基于会员的购买行为，购买什么商品，价格，是否退换，使用什么优惠以及会员基础信息等等。这些信息通过销售订单数据来获取。 */
type QimenTaobaoCrmOrderSyncRequest struct {
    
    /* customerid requiredcustomerid */
    customerid string `json:"customerid";xml:"customerid"`
    
    /* data optional订单数据 */
    data Data `json:"data";xml:"data"`
    
    /* data_type optional数据类型（zx_order直销订单;fx_order分销订单） */
    data_type string `json:"data_type";xml:"data_type"`
    
    /* from_node_id optional数据来源ID */
    from_node_id string `json:"from_node_id";xml:"from_node_id"`
    
    /* from_type optional数据来源类型（taobao淘宝;JD京东;yihaodian一号店;dangdang当当;suning苏宁易购;amazon亚马逊;yinati银泰;mogujie蘑菇街;alibaba阿里巴巴;vop唯品会;meilishuo美丽说;youzan有赞;weixin微信;other其他） */
    from_type string `json:"from_type";xml:"from_type"`
    
    /* msg_id optional全局唯一任务编号 */
    msg_id string `json:"msg_id";xml:"msg_id"`
    
}

func (req *QimenTaobaoCrmOrderSyncRequest) GetAPIName() string {
	return "qimen.taobao.crm.order.sync"
}

/* QimenTaobaoCrmOrderSyncResponse CRM对于会员数据分析需要基于会员的购买行为，购买什么商品，价格，是否退换，使用什么优惠以及会员基础信息等等。这些信息通过销售订单数据来获取。 */
type QimenTaobaoCrmOrderSyncResponse struct {
    
    /* code Basic0成功(其他失败) */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果 */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* CainiaoMerchantInventoryAdjustRequest 商家仓库存调整接口，目前仅支持全量更新 */
type CainiaoMerchantInventoryAdjustRequest struct {
    
    /* adjust_request required商家仓编辑库存 */
    adjust_request MerStoreInvAdjustDto `json:"adjust_request";xml:"adjust_request"`
    
    /* app_name required调用方应用名 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* operation optional操作 */
    operation string `json:"operation";xml:"operation"`
    
}

func (req *CainiaoMerchantInventoryAdjustRequest) GetAPIName() string {
	return "cainiao.merchant.inventory.adjust"
}

/* CainiaoMerchantInventoryAdjustResponse 商家仓库存调整接口，目前仅支持全量更新 */
type CainiaoMerchantInventoryAdjustResponse struct {
    
    /* result Objectresult */
    result SingleResultDto `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsCainiaoBillQueryRequest 查询单据列表 */
type TaobaoWlbWmsCainiaoBillQueryRequest struct {
    
    /* end_modified_time required起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。 */
    end_modified_time time.Time `json:"end_modified_time";xml:"end_modified_time"`
    
    /* order_type optional订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* page_no optional页码。（大于0的整数。默认为1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。（每页条数不超过50条。默认为50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_modified_time required结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。 */
    start_modified_time time.Time `json:"start_modified_time";xml:"start_modified_time"`
    
}

func (req *TaobaoWlbWmsCainiaoBillQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.cainiao.bill.query"
}

/* TaobaoWlbWmsCainiaoBillQueryResponse 查询单据列表 */
type TaobaoWlbWmsCainiaoBillQueryResponse struct {
    
    /* order_info_list Object Array订单列表信息 */
    order_info_list CainiaoBillQueryOrderinfolist `json:"order_info_list";xml:"order_info_list"`
    
    /* total_count Basic总条数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* CainiaoWaybillIiCancelRequest 面单号有误需要取消的时候，调用该接口取消获取的电子面单。 */
type CainiaoWaybillIiCancelRequest struct {
    
    /* cp_code required快递公司code */
    cp_code string `json:"cp_code";xml:"cp_code"`
    
    /* waybill_code required电子面单号 */
    waybill_code string `json:"waybill_code";xml:"waybill_code"`
    
}

func (req *CainiaoWaybillIiCancelRequest) GetAPIName() string {
	return "cainiao.waybill.ii.cancel"
}

/* CainiaoWaybillIiCancelResponse 面单号有误需要取消的时候，调用该接口取消获取的电子面单。 */
type CainiaoWaybillIiCancelResponse struct {
    
    /* cancel_result Basic调用取消是否成功 */
    cancel_result bool `json:"cancel_result";xml:"cancel_result"`
    
}

/* CainiaoWaybillIiProductRequest 商家可以查询物流商的产品类型和服务能力。 */
type CainiaoWaybillIiProductRequest struct {
    
    /* cp_code required快递公司code */
    cp_code string `json:"cp_code";xml:"cp_code"`
    
}

func (req *CainiaoWaybillIiProductRequest) GetAPIName() string {
	return "cainiao.waybill.ii.product"
}

/* CainiaoWaybillIiProductResponse 商家可以查询物流商的产品类型和服务能力。 */
type CainiaoWaybillIiProductResponse struct {
    
    /* product_types Object Array返回值 */
    product_types WaybillProductType `json:"product_types";xml:"product_types"`
    
}

/* TaobaoRdcAligeniusAccountValidateRequest 提供应对接AG的erp系统查询其旗下的商家是否为AG商家 */
type TaobaoRdcAligeniusAccountValidateRequest struct {
    
}

func (req *TaobaoRdcAligeniusAccountValidateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.account.validate"
}

/* TaobaoRdcAligeniusAccountValidateResponse 提供应对接AG的erp系统查询其旗下的商家是否为AG商家 */
type TaobaoRdcAligeniusAccountValidateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsInventoryProfitlossGetRequest 通过订单列表批量获取库存损益单信息 */
type TaobaoWlbWmsInventoryProfitlossGetRequest struct {
    
    /* cn_order_code required菜鸟订单编码 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
}

func (req *TaobaoWlbWmsInventoryProfitlossGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.profitloss.get"
}

/* TaobaoWlbWmsInventoryProfitlossGetResponse 通过订单列表批量获取库存损益单信息 */
type TaobaoWlbWmsInventoryProfitlossGetResponse struct {
    
    /* profit_loss_info Object损益信息 */
    profit_loss_info CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info";xml:"profit_loss_info"`
    
}

/* CainiaoCloudprintMystdtemplatesGetRequest 获取用户使用的菜鸟电子面单 */
type CainiaoCloudprintMystdtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintMystdtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.mystdtemplates.get"
}

/* CainiaoCloudprintMystdtemplatesGetResponse 获取用户使用的菜鸟电子面单 */
type CainiaoCloudprintMystdtemplatesGetResponse struct {
    
    /* result Object返回结果 */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsConsignBillGetRequest 获取销售订单发货信息 */
type TaobaoWlbWmsConsignBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码,cnOrderCode与orderCode必须有一个值不可为空 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsConsignBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.bill.get"
}

/* TaobaoWlbWmsConsignBillGetResponse 获取销售订单发货信息 */
type TaobaoWlbWmsConsignBillGetResponse struct {
    
    /* consign_send_info_list Object Array商品信息列表 */
    consign_send_info_list Consignsendinfolist `json:"consign_send_info_list";xml:"consign_send_info_list"`
    
}

/* TaobaoWlbWmsStockOutBillGetRequest 通过订单号获取单个出库单发货信息 */
type TaobaoWlbWmsStockOutBillGetRequest struct {
    
    /* cn_order_code special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    cn_order_code string `json:"cn_order_code";xml:"cn_order_code"`
    
    /* order_code specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsStockOutBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.out.bill.get"
}

/* TaobaoWlbWmsStockOutBillGetResponse 通过订单号获取单个出库单发货信息 */
type TaobaoWlbWmsStockOutBillGetResponse struct {
    
    /* stock_out_info Object出库信息 */
    stock_out_info CainiaoStockOutBillStockoutinfo `json:"stock_out_info";xml:"stock_out_info"`
    
}

/* CainiaoCloudprintStdtemplatesGetRequest 获取菜鸟标准电子面单模板 */
type CainiaoCloudprintStdtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintStdtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.stdtemplates.get"
}

/* CainiaoCloudprintStdtemplatesGetResponse 获取菜鸟标准电子面单模板 */
type CainiaoCloudprintStdtemplatesGetResponse struct {
    
    /* result Object结果集 */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* TaobaoWangwangClientidBindRequest 绑定nick和客户端id */
type TaobaoWangwangClientidBindRequest struct {
    
    /* app_name required应用名 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* client_id required客户端Id */
    client_id string `json:"client_id";xml:"client_id"`
    
}

func (req *TaobaoWangwangClientidBindRequest) GetAPIName() string {
	return "taobao.wangwang.clientid.bind"
}

/* TaobaoWangwangClientidBindResponse 绑定nick和客户端id */
type TaobaoWangwangClientidBindResponse struct {
    
    /* result Basic0:表示成功
其它为错误 */
    result int64 `json:"result";xml:"result"`
    
}

/* TaobaoWangwangClientidUnbindRequest 解除账号和客户端Id的绑定 */
type TaobaoWangwangClientidUnbindRequest struct {
    
    /* app_name required应用名 */
    app_name string `json:"app_name";xml:"app_name"`
    
    /* client_id required客户端Id */
    client_id string `json:"client_id";xml:"client_id"`
    
}

func (req *TaobaoWangwangClientidUnbindRequest) GetAPIName() string {
	return "taobao.wangwang.clientid.unbind"
}

/* TaobaoWangwangClientidUnbindResponse 解除账号和客户端Id的绑定 */
type TaobaoWangwangClientidUnbindResponse struct {
    
    /* result Basic返回0表示成功， 其他值为错误 */
    result int64 `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsStockInOrderNotifyRequest 入库通知单 */
type TaobaoWlbWmsStockInOrderNotifyRequest struct {
    
    /* expect_end_time optional预期送达结束时间 */
    expect_end_time string `json:"expect_end_time";xml:"expect_end_time"`
    
    /* expect_start_time optional预期送达开始时间 */
    expect_start_time string `json:"expect_start_time";xml:"expect_start_time"`
    
    /* extend_fields optional扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-” */
    extend_fields string `json:"extend_fields";xml:"extend_fields"`
    
    /* inbound_type_desc optional可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等 */
    inbound_type_desc string `json:"inbound_type_desc";xml:"inbound_type_desc"`
    
    /* order_code required入库单据编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_create_time required单据创建时间 */
    order_create_time time.Time `json:"order_create_time";xml:"order_create_time"`
    
    /* order_flag optional订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9） */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list required系统自动生成 */
    order_item_list Orderitemlistwlbwmsstockinordernotifywl `json:"order_item_list";xml:"order_item_list"`
    
    /* order_type required单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库 */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* prev_order_code optional来源单据号，如销售退货时填充原销售订单号 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional系统自动生成 */
    receiver_info Receiverinfowlbwmsstockinordernotifywl `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注，销退入库订单需要留言备注填充到此字段 */
    remark string `json:"remark";xml:"remark"`
    
    /* return_reason optional销退时请提供退货的原因 */
    return_reason string `json:"return_reason";xml:"return_reason"`
    
    /* sender_info optional系统自动生成 */
    sender_info Senderinfowlbwmsstockinordernotifywl `json:"sender_info";xml:"sender_info"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* supplier_code optional供应商编码，往来单位编码 */
    supplier_code string `json:"supplier_code";xml:"supplier_code"`
    
    /* supplier_name optional供应商名称 ，往来单位名称 */
    supplier_name string `json:"supplier_name";xml:"supplier_name"`
    
    /* tms_order_code optional运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号 */
    tms_order_code string `json:"tms_order_code";xml:"tms_order_code"`
    
    /* tms_service_code optional配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* tms_service_name optional快递公司名称 */
    tms_service_name string `json:"tms_service_name";xml:"tms_service_name"`
    
}

func (req *TaobaoWlbWmsStockInOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.in.order.notify"
}

/* TaobaoWlbWmsStockInOrderNotifyResponse 入库通知单 */
type TaobaoWlbWmsStockInOrderNotifyResponse struct {
    
    /* order_code Basic仓储订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误详细 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoWlbWmsInventoryCountRequest 损益单回传 */
type TaobaoWlbWmsInventoryCountRequest struct {
    
    /* content optional损益单回传信息 */
    content WlbWmsInventoryCount `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsInventoryCountRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.count"
}

/* TaobaoWlbWmsInventoryCountResponse 损益单回传 */
type TaobaoWlbWmsInventoryCountResponse struct {
    
    /* result Objectresult */
    result WlbWmsInventoryCountResp `json:"result";xml:"result"`
    
}

/* AlipaySystemOauthTokenRequest OAuth2.0 授权的第二步，换取访问令牌。 */
type AlipaySystemOauthTokenRequest struct {
    
    /* code optional授权码，用户对应用授权后得到。 */
    code string `json:"code";xml:"code"`
    
    /* grant_type required获取访问令牌的类型，authorization_code表示用授权码换，refresh_token表示用刷新令牌来换。 */
    grant_type string `json:"grant_type";xml:"grant_type"`
    
    /* refresh_token optional刷新令牌，上次换取访问令牌是得到。 */
    refresh_token string `json:"refresh_token";xml:"refresh_token"`
    
}

func (req *AlipaySystemOauthTokenRequest) GetAPIName() string {
	return "alipay.system.oauth.token"
}

/* AlipaySystemOauthTokenResponse OAuth2.0 授权的第二步，换取访问令牌。 */
type AlipaySystemOauthTokenResponse struct {
    
    /* access_token Basic访问令牌 */
    access_token string `json:"access_token";xml:"access_token"`
    
    /* alipay_user_id Basic支付宝用户的id。 */
    alipay_user_id string `json:"alipay_user_id";xml:"alipay_user_id"`
    
    /* expires_in Basic访问令牌的有效时间，单位是秒。 */
    expires_in string `json:"expires_in";xml:"expires_in"`
    
    /* re_expires_in Basic刷新令牌的有效时间，单位是秒。 */
    re_expires_in string `json:"re_expires_in";xml:"re_expires_in"`
    
    /* refresh_token Basic刷新令牌 */
    refresh_token string `json:"refresh_token";xml:"refresh_token"`
    
}

/* TaobaoOmniorderStoreDeliverconfigUpdateRequest 修改门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigUpdateRequest struct {
    
    /* store_deliver_config required卖家发货配置 */
    store_deliver_config StoreDeliverConfig `json:"store_deliver_config";xml:"store_deliver_config"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetAPIName() string {
	return "taobao.omniorder.store.deliverconfig.update"
}

/* TaobaoOmniorderStoreDeliverconfigUpdateResponse 修改门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* QimenTaobaoUopTobDeliveryorderConfirmRequest 2B订单出库确认回传 */
type QimenTaobaoUopTobDeliveryorderConfirmRequest struct {
    
    /* customerId requiredcustomerId */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* request optionaltob出库确认对象 */
    request ToBDeliveryOrderConfirmRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoUopTobDeliveryorderConfirmRequest) GetAPIName() string {
	return "qimen.taobao.uop.tob.deliveryorder.confirm"
}

/* QimenTaobaoUopTobDeliveryorderConfirmResponse 2B订单出库确认回传 */
type QimenTaobaoUopTobDeliveryorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果 */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoOmniorderStoreCollectconfigUpdateRequest 修改门店自提配置内容 */
type TaobaoOmniorderStoreCollectconfigUpdateRequest struct {
    
    /* store_collect_config required门店自提配置 */
    store_collect_config StoreCollectConfig `json:"store_collect_config";xml:"store_collect_config"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreCollectconfigUpdateRequest) GetAPIName() string {
	return "taobao.omniorder.store.collectconfig.update"
}

/* TaobaoOmniorderStoreCollectconfigUpdateResponse 修改门店自提配置内容 */
type TaobaoOmniorderStoreCollectconfigUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoOmniorderStoreDeliverconfigGetRequest 查询门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigGetRequest struct {
    
    /* activity optional是否是活动期 */
    activity bool `json:"activity";xml:"activity"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreDeliverconfigGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.deliverconfig.get"
}

/* TaobaoOmniorderStoreDeliverconfigGetResponse 查询门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigGetResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* CainiaoCloudprintCustomaresGetRequest 供isv使用，获取商家的自定义区的模板信息 */
type CainiaoCloudprintCustomaresGetRequest struct {
    
    /* template_id required用户使用的标准模板id */
    template_id int64 `json:"template_id";xml:"template_id"`
    
}

func (req *CainiaoCloudprintCustomaresGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.customares.get"
}

/* CainiaoCloudprintCustomaresGetResponse 供isv使用，获取商家的自定义区的模板信息 */
type CainiaoCloudprintCustomaresGetResponse struct {
    
    /* result Object结果 */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsInventoryLackUploadRequest 缺货通知 */
type TaobaoWlbWmsInventoryLackUploadRequest struct {
    
    /* content optional缺货通知信息 */
    content WlbWmsInventoryLackUpload `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsInventoryLackUploadRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.lack.upload"
}

/* TaobaoWlbWmsInventoryLackUploadResponse 缺货通知 */
type TaobaoWlbWmsInventoryLackUploadResponse struct {
    
    /* result Object缺货回告 */
    result WlbWmsInventoryLackUploadResp `json:"result";xml:"result"`
    
}

/* TaobaoOmniorderStoreCollectconfigGetRequest 查询门店自提配置内容 */
type TaobaoOmniorderStoreCollectconfigGetRequest struct {
    
    /* activity optional是否是活动期 */
    activity bool `json:"activity";xml:"activity"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreCollectconfigGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.collectconfig.get"
}

/* TaobaoOmniorderStoreCollectconfigGetResponse 查询门店自提配置内容 */
type TaobaoOmniorderStoreCollectconfigGetResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoUserSellerGetRequest 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。 */
type TaobaoUserSellerGetRequest struct {
    
    /* fields required需要返回的字段列表，可选值为返回示例值中的可以看到的字段 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoUserSellerGetRequest) GetAPIName() string {
	return "taobao.user.seller.get"
}

/* TaobaoUserSellerGetResponse 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。 */
type TaobaoUserSellerGetResponse struct {
    
    /* user Object用户信息 */
    user User `json:"user";xml:"user"`
    
}

/* TaobaoWlbWmsInventoryQueryRequest 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存 */
type TaobaoWlbWmsInventoryQueryRequest struct {
    
    /* batch_code optional库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。 */
    batch_code string `json:"batch_code";xml:"batch_code"`
    
    /* channel_code optional渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他） */
    channel_code string `json:"channel_code";xml:"channel_code"`
    
    /* due_date optional失效日期，type=2时字段传值有效。 */
    due_date time.Time `json:"due_date";xml:"due_date"`
    
    /* inventory_type optional库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 ) */
    inventory_type int64 `json:"inventory_type";xml:"inventory_type"`
    
    /* item_id optional菜鸟商品ID */
    item_id string `json:"item_id";xml:"item_id"`
    
    /* page_no optional分页的第几页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页多少条，最大50条 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* produce_date optional生产日期，type=2时字段传值有效。 */
    produce_date time.Time `json:"produce_date";xml:"produce_date"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* _type optional库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性） */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbWmsInventoryQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.query"
}

/* TaobaoWlbWmsInventoryQueryResponse 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存 */
type TaobaoWlbWmsInventoryQueryResponse struct {
    
    /* item_list Object Array商品详情列表 */
    item_list WmsInventoryQueryItemlist `json:"item_list";xml:"item_list"`
    
    /* total_count Basic总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
    /* wl_error_code Basic错误代码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}

/* AlibabaEinvoicePaperPrintRequest 打印一张已开具成功的纸票 */
type AlibabaEinvoicePaperPrintRequest struct {
    
    /* dialog_setting_flag required打印框设置，0=不弹打印设置框，1=弹出打印设置框 */
    dialog_setting_flag int64 `json:"dialog_setting_flag";xml:"dialog_setting_flag"`
    
    /* force_print required是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印 */
    force_print bool `json:"force_print";xml:"force_print"`
    
    /* payee_register_no required销售方纳税人识别号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* print_flag required打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。 */
    print_flag int64 `json:"print_flag";xml:"print_flag"`
    
    /* serial_no required开票流水号 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoicePaperPrintRequest) GetAPIName() string {
	return "alibaba.einvoice.paper.print"
}

/* AlibabaEinvoicePaperPrintResponse 打印一张已开具成功的纸票 */
type AlibabaEinvoicePaperPrintResponse struct {
    
    /* is_success Basic调用结果，打印结果tmc消息alibaba_invoice_PaperOpsReturn异步通知 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoItemSellerGetRequest 获取单个商品的全部信息 */
type TaobaoItemSellerGetRequest struct {
    
    /* fields required需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iid required商品数字ID */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemSellerGetRequest) GetAPIName() string {
	return "taobao.item.seller.get"
}

/* TaobaoItemSellerGetResponse 获取单个商品的全部信息 */
type TaobaoItemSellerGetResponse struct {
    
    /* item Object商品详细信息 */
    item Item `json:"item";xml:"item"`
    
}

/* AlibabaEinvoicePaperInvalidRequest 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票 */
type AlibabaEinvoicePaperInvalidRequest struct {
    
    /* invalid_operator required作废操作人 */
    invalid_operator string `json:"invalid_operator";xml:"invalid_operator"`
    
    /* invalid_type required作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废 */
    invalid_type int64 `json:"invalid_type";xml:"invalid_type"`
    
    /* invoice_code optional发票代码，空白作废时必填 */
    invoice_code string `json:"invoice_code";xml:"invoice_code"`
    
    /* invoice_no optional发票号码，空白作废时必填 */
    invoice_no string `json:"invoice_no";xml:"invoice_no"`
    
    /* payee_register_no required销售方纳税人识别号 */
    payee_register_no string `json:"payee_register_no";xml:"payee_register_no"`
    
    /* serial_no required开票流水号 */
    serial_no string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoicePaperInvalidRequest) GetAPIName() string {
	return "alibaba.einvoice.paper.invalid"
}

/* AlibabaEinvoicePaperInvalidResponse 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票 */
type AlibabaEinvoicePaperInvalidResponse struct {
    
    /* is_success Basic接口调用是否成功，操作结果tmc异步返回alibaba_invoice_PaperOpsReturn */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoItemsSellerListGetRequest 批量获取商品详细信息 */
type TaobaoItemsSellerListGetRequest struct {
    
    /* fields required需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。 */
    fields string `json:"fields";xml:"fields"`
    
    /* num_iids required商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。 */
    num_iids string `json:"num_iids";xml:"num_iids"`
    
}

func (req *TaobaoItemsSellerListGetRequest) GetAPIName() string {
	return "taobao.items.seller.list.get"
}

/* TaobaoItemsSellerListGetResponse 批量获取商品详细信息 */
type TaobaoItemsSellerListGetResponse struct {
    
    /* items Object Array商品详细信息列表 */
    items Item `json:"items";xml:"items"`
    
}

/* TaobaoQimenSnReportRequest WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP */
type TaobaoQimenSnReportRequest struct {
    
    /* currentPage required当前页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* deliveryOrder optional发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品列表 */
    items Item `json:"items";xml:"items"`
    
    /* pageSize required每页记录的条数 */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* totalPage required总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

func (req *TaobaoQimenSnReportRequest) GetAPIName() string {
	return "taobao.qimen.sn.report"
}

/* TaobaoQimenSnReportResponse WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP */
type TaobaoQimenSnReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoTaeBookBillGetRequest tae查询单笔虚拟账户明细 */
type TaobaoTaeBookBillGetRequest struct {
    
    /* account_id required虚拟账户科目编号 */
    account_id int64 `json:"account_id";xml:"account_id"`
    
    /* bid optional虚拟账户流水编号 */
    bid int64 `json:"bid";xml:"bid"`
    
    /* fields required需要返回的字段:参见BookBill结构体 */
    fields string `json:"fields";xml:"fields"`
    
    /* id required虚拟账户流水编号 */
    id int64 `json:"id";xml:"id"`
    
}

func (req *TaobaoTaeBookBillGetRequest) GetAPIName() string {
	return "taobao.tae.book.bill.get"
}

/* TaobaoTaeBookBillGetResponse tae查询单笔虚拟账户明细 */
type TaobaoTaeBookBillGetResponse struct {
    
    /* bookbill Object虚拟账户账单 */
    bookbill TopAcctCashJourDto `json:"bookbill";xml:"bookbill"`
    
}

/* TaobaoTaeBookBillsGetRequest tae查询虚拟账户明细数据 */
type TaobaoTaeBookBillsGetRequest struct {
    
    /* account_id optional虚拟账户科目编号 */
    account_id int64 `json:"account_id";xml:"account_id"`
    
    /* end_time required记账结束时间,end_time与start_time相差不能超过30天 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* fields required需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略 */
    fields string `json:"fields";xml:"fields"`
    
    /* journal_types optional明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除 */
    journal_types int64 `json:"journal_types";xml:"journal_types"`
    
    /* page_no optional页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页大小,建议40~100,不能超过100 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time required记账开始时间 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoTaeBookBillsGetRequest) GetAPIName() string {
	return "taobao.tae.book.bills.get"
}

/* TaobaoTaeBookBillsGetResponse tae查询虚拟账户明细数据 */
type TaobaoTaeBookBillsGetResponse struct {
    
    /* bills Object Array虚拟账户账单列表 */
    bills TopAcctCashJourDto `json:"bills";xml:"bills"`
    
    /* has_next Basic是否有下一页 */
    has_next bool `json:"has_next";xml:"has_next"`
    
    /* total_results Basic当前查询的结果数,非总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoWlbOrderJzQueryRequest 家装业务查询物流公司api */
type TaobaoWlbOrderJzQueryRequest struct {
    
    /* ins_jz_receiver_t_o optional家装安装服务收货人信息 */
    ins_jz_receiver_t_o JzReceiverTo `json:"ins_jz_receiver_t_o";xml:"ins_jz_receiver_t_o"`
    
    /* jz_receiver_to optional家装收货人信息 */
    jz_receiver_to JzReceiverTo `json:"jz_receiver_to";xml:"jz_receiver_to"`
    
    /* sender_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址 */
    sender_id int64 `json:"sender_id";xml:"sender_id"`
    
    /* tid optional交易id */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoWlbOrderJzQueryRequest) GetAPIName() string {
	return "taobao.wlb.order.jz.query"
}

/* TaobaoWlbOrderJzQueryResponse 家装业务查询物流公司api */
type TaobaoWlbOrderJzQueryResponse struct {
    
    /* result Object结果信息 */
    result JzTopDto `json:"result";xml:"result"`
    
    /* result_error_code Basic错误编码 */
    result_error_code string `json:"result_error_code";xml:"result_error_code"`
    
    /* result_error_msg Basic错误信息 */
    result_error_msg string `json:"result_error_msg";xml:"result_error_msg"`
    
    /* result_success Basic是否成功 */
    result_success bool `json:"result_success";xml:"result_success"`
    
}

/* TaobaoRdcAligeniusWarehouseResendUpdateRequest 补发单状态回传接口 */
type TaobaoRdcAligeniusWarehouseResendUpdateRequest struct {
    
    /* param0 required参数 */
    param0 UpdateResendStatusDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.resend.update"
}

/* TaobaoRdcAligeniusWarehouseResendUpdateResponse 补发单状态回传接口 */
type TaobaoRdcAligeniusWarehouseResendUpdateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoQimenOrderCancelRequest ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景 */
type TaobaoQimenOrderCancelRequest struct {
    
    /* cancelReason optional取消原因 */
    cancelReason string `json:"cancelReason";xml:"cancelReason"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderCode required单据编码 */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderId optional仓储系统单据编码 */
    orderId string `json:"orderId";xml:"orderId"`
    
    /* orderType optional单据类型(JYCK=一般交易出库单;HHCK= 换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK=其他入库;XTRK= 销退入库;THRK=退货入库;HHRK= 换货入库;CNJG= 仓内加工单;CGTH=采购退货出库单) */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderCancelRequest) GetAPIName() string {
	return "taobao.qimen.order.cancel"
}

/* TaobaoQimenOrderCancelResponse ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景 */
type TaobaoQimenOrderCancelResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenDeliveryorderConfirmRequest taobao.qimen.deliveryorder.confirm */
type TaobaoQimenDeliveryorderConfirmRequest struct {
    
    /* deliveryOrder optional发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderLines optional单据列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* packages optional包裹信息 */
    packages Package `json:"packages";xml:"packages"`
    
}

func (req *TaobaoQimenDeliveryorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.confirm"
}

/* TaobaoQimenDeliveryorderConfirmResponse taobao.qimen.deliveryorder.confirm */
type TaobaoQimenDeliveryorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenDeliveryorderCreateRequest taobao.qimen.deliveryorder.create */
type TaobaoQimenDeliveryorderCreateRequest struct {
    
    /* deliveryOrder optional发货单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional订单列表 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenDeliveryorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.create"
}

/* TaobaoQimenDeliveryorderCreateResponse taobao.qimen.deliveryorder.create */
type TaobaoQimenDeliveryorderCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* createTime Basic订单创建时间(YYYY-MM-DD HH:MM:SS) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderId Basic出库单仓储系统编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* deliveryOrders Object Array发货单信息 */
    deliveryOrders DeliveryOrder `json:"deliveryOrders";xml:"deliveryOrders"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* logisticsCode Basic物流公司编码(统仓统配使用) */
    logisticsCode string `json:"logisticsCode";xml:"logisticsCode"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* warehouseCode Basic仓库编码(统仓统配使用) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

/* TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest 补发单erp物流信息回传平台 */
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest struct {
    
    /* param0 required参数 */
    param0 SendResendLogisticsMsgDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.resend.logistics.msg.post"
}

/* TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse 补发单erp物流信息回传平台 */
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoQimenOrderprocessReportRequest WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。 */
type TaobaoQimenOrderprocessReportRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* order optional订单信息 */
    order Order `json:"order";xml:"order"`
    
    /* process optional订单处理信息 */
    process Process `json:"process";xml:"process"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenOrderprocessReportRequest) GetAPIName() string {
	return "taobao.qimen.orderprocess.report"
}

/* TaobaoQimenOrderprocessReportResponse WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。 */
type TaobaoQimenOrderprocessReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoFenxiaoDistributorItemsGetRequest 供应商查询分销商商品下载记录。 */
type TaobaoFenxiaoDistributorItemsGetRequest struct {
    
    /* distributor_id special分销商ID 。 */
    distributor_id int64 `json:"distributor_id";xml:"distributor_id"`
    
    /* end_modified optional设置结束时间,空为不设置。 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* product_id special产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* start_modified optional设置开始时间。空为不设置。 */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoFenxiaoDistributorItemsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributor.items.get"
}

/* TaobaoFenxiaoDistributorItemsGetResponse 供应商查询分销商商品下载记录。 */
type TaobaoFenxiaoDistributorItemsGetResponse struct {
    
    /* records Object Array下载记录对象 */
    records FenxiaoItemRecord `json:"records";xml:"records"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoFenxiaoOrderConfirmPaidRequest 供应商确认收款（非支付宝交易）。 */
type TaobaoFenxiaoOrderConfirmPaidRequest struct {
    
    /* confirm_remark optional确认支付信息（字数小于100） */
    confirm_remark string `json:"confirm_remark";xml:"confirm_remark"`
    
    /* purchase_order_id required采购单编号。 */
    purchase_order_id int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
}

func (req *TaobaoFenxiaoOrderConfirmPaidRequest) GetAPIName() string {
	return "taobao.fenxiao.order.confirm.paid"
}

/* TaobaoFenxiaoOrderConfirmPaidResponse 供应商确认收款（非支付宝交易）。 */
type TaobaoFenxiaoOrderConfirmPaidResponse struct {
    
    /* is_success Basic确认结果成功与否 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoCooperationGetRequest 获取供应商的合作关系信息 */
type TaobaoFenxiaoCooperationGetRequest struct {
    
    /* end_date optional合作结束时间yyyy-MM-dd HH:mm:ss */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* page_no optional页码（大于0的整数，默认1） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数（默认20，最大50） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_date optional合作开始时间yyyy-MM-dd HH:mm:ss */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* status optional合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止) */
    status string `json:"status";xml:"status"`
    
    /* trade_type optional分销方式：AGENT(代销) 、DEALER（经销） */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoCooperationGetRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.get"
}

/* TaobaoFenxiaoCooperationGetResponse 获取供应商的合作关系信息 */
type TaobaoFenxiaoCooperationGetResponse struct {
    
    /* cooperations Object Array合作分销关系 */
    cooperations Cooperation `json:"cooperations";xml:"cooperations"`
    
    /* total_results Basic结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoQianniuCloudkefuStatuslogGetRequest 查询客服账号的操作记录。如：登录，下线，挂起等。接口即将下线，请ISV切换到taobao.qianniu.cloudkefu.onlinestatuslog.get接口上 */
type TaobaoQianniuCloudkefuStatuslogGetRequest struct {
    
    /* account_ids required子帐号列表，最多10个 */
    account_ids int64 `json:"account_ids";xml:"account_ids"`
    
    /* domian optional域，淘宝:cntaobao，1688:cnalichn */
    domian string `json:"domian";xml:"domian"`
    
    /* end_time optional查询结束时间，默认当天24点 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* main_account_id required主帐号ID */
    main_account_id int64 `json:"main_account_id";xml:"main_account_id"`
    
    /* page_num optional页码，每页50个 */
    page_num int64 `json:"page_num";xml:"page_num"`
    
    /* page_size optional分页 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time optional查询开始时间，默认当天0点 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
    /* _type required记录类型，PC上下线：8；移动上下线：4；挂起类型：5 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoQianniuCloudkefuStatuslogGetRequest) GetAPIName() string {
	return "taobao.qianniu.cloudkefu.statuslog.get"
}

/* TaobaoQianniuCloudkefuStatuslogGetResponse 查询客服账号的操作记录。如：登录，下线，挂起等。接口即将下线，请ISV切换到taobao.qianniu.cloudkefu.onlinestatuslog.get接口上 */
type TaobaoQianniuCloudkefuStatuslogGetResponse struct {
    
    /* result Objectresult */
    result ResultTo `json:"result";xml:"result"`
    
}

/* TaobaoLogisticsOnlineConfirmRequest <br><font color='red'>仅使用taobao.logistics.online.send 发货时，未输入运单号的情况下，需要使用该接口确认发货。发货接口主要有taobao.logistics.offline.send 和taobao.logistics.online.send <br>
确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外 */
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

/* TaobaoLogisticsOnlineConfirmResponse <br><font color='red'>仅使用taobao.logistics.online.send 发货时，未输入运单号的情况下，需要使用该接口确认发货。发货接口主要有taobao.logistics.offline.send 和taobao.logistics.online.send <br>
确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外 */
type TaobaoLogisticsOnlineConfirmResponse struct {
    
    /* shipping Object只返回is_success：是否成功。 */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}

/* TaobaoWlbOrderJzConsignRequest 家装类订单使用该接口发货 */
type TaobaoWlbOrderJzConsignRequest struct {
    
    /* ins_receiver_to optional安装收货人信息,如果为空,则取默认收货人信息 */
    ins_receiver_to JzReceiverTo `json:"ins_receiver_to";xml:"ins_receiver_to"`
    
    /* ins_tp_dto optional安装公司信息,需要安装时,才填写 */
    ins_tp_dto Tpdto `json:"ins_tp_dto";xml:"ins_tp_dto"`
    
    /* jz_receiver_to optional家装收货人信息,如果为空,则取默认收货信息 */
    jz_receiver_to JzReceiverTo `json:"jz_receiver_to";xml:"jz_receiver_to"`
    
    /* jz_top_args optional发货参数 */
    jz_top_args JzTopArgs `json:"jz_top_args";xml:"jz_top_args"`
    
    /* lg_tp_dto required物流公司信息 */
    lg_tp_dto Tpdto `json:"lg_tp_dto";xml:"lg_tp_dto"`
    
    /* sender_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址 */
    sender_id int64 `json:"sender_id";xml:"sender_id"`
    
    /* tid required交易号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoWlbOrderJzConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.jz.consign"
}

/* TaobaoWlbOrderJzConsignResponse 家装类订单使用该接口发货 */
type TaobaoWlbOrderJzConsignResponse struct {
    
    /* result_error_code Basic错误码 */
    result_error_code string `json:"result_error_code";xml:"result_error_code"`
    
    /* result_error_msg Basic错误信息描述 */
    result_error_msg string `json:"result_error_msg";xml:"result_error_msg"`
    
    /* result_success Basic是否成功 */
    result_success bool `json:"result_success";xml:"result_success"`
    
}

/* TaobaoLogisticsOnlineCancelRequest 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。 */
type TaobaoLogisticsOnlineCancelRequest struct {
    
    /* tid required淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOnlineCancelRequest) GetAPIName() string {
	return "taobao.logistics.online.cancel"
}

/* TaobaoLogisticsOnlineCancelResponse 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。 */
type TaobaoLogisticsOnlineCancelResponse struct {
    
    /* is_success Basic成功与失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* modify_time Basic修改时间 */
    modify_time string `json:"modify_time";xml:"modify_time"`
    
    /* recreated_order_id Basic重新创建物流订单id */
    recreated_order_id string `json:"recreated_order_id";xml:"recreated_order_id"`
    
}

/* AlibabaEinvoiceApplyTestRequest 开票申请消息测试接口 */
type AlibabaEinvoiceApplyTestRequest struct {
    
    /* business_type required0=个人，1=企业 */
    business_type int64 `json:"business_type";xml:"business_type"`
    
    /* payer_name required买家抬头 */
    payer_name string `json:"payer_name";xml:"payer_name"`
    
    /* payer_register_no optional买家税号 */
    payer_register_no string `json:"payer_register_no";xml:"payer_register_no"`
    
    /* platform_tid required订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
}

func (req *AlibabaEinvoiceApplyTestRequest) GetAPIName() string {
	return "alibaba.einvoice.apply.test"
}

/* AlibabaEinvoiceApplyTestResponse 开票申请消息测试接口 */
type AlibabaEinvoiceApplyTestResponse struct {
    
    /* is_success Basic消息是否发送成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoLogisticsDummySendRequest 用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货 */
type TaobaoLogisticsDummySendRequest struct {
    
    /* feature optionalfeature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
    feature string `json:"feature";xml:"feature"`
    
    /* seller_ip optional商家的IP地址 */
    seller_ip string `json:"seller_ip";xml:"seller_ip"`
    
    /* tid required淘宝交易ID */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsDummySendRequest) GetAPIName() string {
	return "taobao.logistics.dummy.send"
}

/* TaobaoLogisticsDummySendResponse 用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货 */
type TaobaoLogisticsDummySendResponse struct {
    
    /* shipping Object返回发货是否成功is_success */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}

/* TaobaoLogisticsOfflineSendRequest 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
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

/* TaobaoLogisticsOfflineSendResponse 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
type TaobaoLogisticsOfflineSendResponse struct {
    
    /* shipping Object自己联系的调用结果 */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}

/* TaobaoQimenReturnorderCreateRequest ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作 */
type TaobaoQimenReturnorderCreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* itemList optionalitemList */
    itemList Item `json:"itemList";xml:"itemList"`
    
    /* orderLines optional订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* returnOrder optional退货单信息 */
    returnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

func (req *TaobaoQimenReturnorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.create"
}

/* TaobaoQimenReturnorderCreateResponse ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作 */
type TaobaoQimenReturnorderCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* returnOrderId Basic仓储系统退货单编码 */
    returnOrderId string `json:"returnOrderId";xml:"returnOrderId"`
    
}

/* TaobaoWlbStoresBaseinfoGetRequest 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库 */
type TaobaoWlbStoresBaseinfoGetRequest struct {
    
    /* _type optional0.商家仓库.1.菜鸟仓库.2全部被划分的仓库 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbStoresBaseinfoGetRequest) GetAPIName() string {
	return "taobao.wlb.stores.baseinfo.get"
}

/* TaobaoWlbStoresBaseinfoGetResponse 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库 */
type TaobaoWlbStoresBaseinfoGetResponse struct {
    
    /* store_info_list Object Array仓库列表 */
    store_info_list StoreInfo `json:"store_info_list";xml:"store_info_list"`
    
    /* total_count Basic返回的总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoQimenEntryorderCreateRequest ERP调用接口，创建入库单; */
type TaobaoQimenEntryorderCreateRequest struct {
    
    /* entryOrder optional入库单信息 */
    entryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品信息 */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional入库单详情 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenEntryorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.create"
}

/* TaobaoQimenEntryorderCreateResponse ERP调用接口，创建入库单; */
type TaobaoQimenEntryorderCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* entryOrderId Basic仓储系统入库单编码 */
    entryOrderId string `json:"entryOrderId";xml:"entryOrderId"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenReturnorderConfirmRequest taobao.qimen.returnorder.confirm */
type TaobaoQimenReturnorderConfirmRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orderLines optional订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    /* returnOrder optional退货单信息 */
    returnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

func (req *TaobaoQimenReturnorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.confirm"
}

/* TaobaoQimenReturnorderConfirmResponse taobao.qimen.returnorder.confirm */
type TaobaoQimenReturnorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenStockchangeReportRequest WMS调用奇门的接口,将库存异动信息信息回传给ERP */
type TaobaoQimenStockchangeReportRequest struct {
    
    /* currentPage optional奇门仓储字段,说明,string(50),, */
    currentPage string `json:"currentPage";xml:"currentPage"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optionalitem */
    items Item `json:"items";xml:"items"`
    
    /* orderCode optional奇门仓储字段,说明,string(50),, */
    orderCode string `json:"orderCode";xml:"orderCode"`
    
    /* orderType optional奇门仓储字段,说明,string(50),, */
    orderType string `json:"orderType";xml:"orderType"`
    
    /* ownerCode optional奇门仓储字段,说明,string(50),, */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* pageSize optional奇门仓储字段,说明,string(50),, */
    pageSize string `json:"pageSize";xml:"pageSize"`
    
    /* remark optional奇门仓储字段,说明,string(50),, */
    remark string `json:"remark";xml:"remark"`
    
    /* totalPage optional奇门仓储字段,说明,string(50),, */
    totalPage string `json:"totalPage";xml:"totalPage"`
    
    /* warehouseCode optional奇门仓储字段,说明,string(50),, */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockchangeReportRequest) GetAPIName() string {
	return "taobao.qimen.stockchange.report"
}

/* TaobaoQimenStockchangeReportResponse WMS调用奇门的接口,将库存异动信息信息回传给ERP */
type TaobaoQimenStockchangeReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenInventoryQueryRequest ERP调用奇门的接口,查询商品的库存量 */
type TaobaoQimenInventoryQueryRequest struct {
    
    /* criteriaList optional查询准则 */
    criteriaList Criteria `json:"criteriaList";xml:"criteriaList"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenInventoryQueryRequest) GetAPIName() string {
	return "taobao.qimen.inventory.query"
}

/* TaobaoQimenInventoryQueryResponse ERP调用奇门的接口,查询商品的库存量 */
type TaobaoQimenInventoryQueryResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* items Object Array商品的库存信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenSingleitemSynchronizeRequest ERP调用奇门的接口,同步商品信息给WMS */
type TaobaoQimenSingleitemSynchronizeRequest struct {
    
    /* actionType required操作类型(两种类型：add|update) */
    actionType string `json:"actionType";xml:"actionType"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* item optional商品信息 */
    item Item `json:"item";xml:"item"`
    
    /* ownerCode required货主编码 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* supplierCode optional供应商编码 */
    supplierCode string `json:"supplierCode";xml:"supplierCode"`
    
    /* supplierName optional供应商名称 */
    supplierName string `json:"supplierName";xml:"supplierName"`
    
    /* warehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER) */
    warehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenSingleitemSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.singleitem.synchronize"
}

/* TaobaoQimenSingleitemSynchronizeResponse ERP调用奇门的接口,同步商品信息给WMS */
type TaobaoQimenSingleitemSynchronizeResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* itemId Basic仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传) */
    itemId string `json:"itemId";xml:"itemId"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoItemTemplatesGetRequest 查询当前登录用户的店铺的宝贝详情页的模板名称 */
type TaobaoItemTemplatesGetRequest struct {
    
}

func (req *TaobaoItemTemplatesGetRequest) GetAPIName() string {
	return "taobao.item.templates.get"
}

/* TaobaoItemTemplatesGetResponse 查询当前登录用户的店铺的宝贝详情页的模板名称 */
type TaobaoItemTemplatesGetResponse struct {
    
    /* item_template_list Object Array返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店） */
    item_template_list ItemTemplate `json:"item_template_list";xml:"item_template_list"`
    
}

/* TaobaoQimenDeliveryorderBatchconfirmRequest taobao.qimen.deliveryorder.batchconfirm */
type TaobaoQimenDeliveryorderBatchconfirmRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional发货单列表 */
    orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchconfirmRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchconfirm"
}

/* TaobaoQimenDeliveryorderBatchconfirmResponse taobao.qimen.deliveryorder.batchconfirm */
type TaobaoQimenDeliveryorderBatchconfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenStockoutCreateRequest ERP调用奇门接口，创建出库单信息 */
type TaobaoQimenStockoutCreateRequest struct {
    
    /* deliveryOrder optional出库单信息 */
    deliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional单据信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenStockoutCreateRequest) GetAPIName() string {
	return "taobao.qimen.stockout.create"
}

/* TaobaoQimenStockoutCreateResponse ERP调用奇门接口，创建出库单信息 */
type TaobaoQimenStockoutCreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* createTime Basic订单创建时间(YYYY-MM-DD HH:MM:SS) */
    createTime string `json:"createTime";xml:"createTime"`
    
    /* deliveryOrderId Basic出库单仓储系统编码 */
    deliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenEntryorderConfirmRequest WMS调用接口，回传入库单信息; */
type TaobaoQimenEntryorderConfirmRequest struct {
    
    /* entryOrder optional入库单信息 */
    entryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品信息列表 */
    items Item `json:"items";xml:"items"`
    
    /* orderLines optional订单信息 */
    orderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenEntryorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.confirm"
}

/* TaobaoQimenEntryorderConfirmResponse WMS调用接口，回传入库单信息; */
type TaobaoQimenEntryorderConfirmResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* TaobaoQimenDeliveryorderBatchcreateRequest ERP调用接口，将发货信息批量推送给WMS */
type TaobaoQimenDeliveryorderBatchcreateRequest struct {
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* orders optional订单信息 */
    orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchcreateRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchcreate"
}

/* TaobaoQimenDeliveryorderBatchcreateResponse ERP调用接口，将发货信息批量推送给WMS */
type TaobaoQimenDeliveryorderBatchcreateResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
    /* orders Object Array订单详情 */
    orders Order `json:"orders";xml:"orders"`
    
}

/* TaobaoWlbImportThreeplResourceGetRequest 获取3pl直邮的发货可用资源 */
type TaobaoWlbImportThreeplResourceGetRequest struct {
    
    /* from_id required发货地区域id */
    from_id int64 `json:"from_id";xml:"from_id"`
    
    /* to_address required收件人地址 */
    to_address AddressDto `json:"to_address";xml:"to_address"`
    
    /* _type requiredONLINE或者OFFLINE */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbImportThreeplResourceGetRequest) GetAPIName() string {
	return "taobao.wlb.import.threepl.resource.get"
}

/* TaobaoWlbImportThreeplResourceGetResponse 获取3pl直邮的发货可用资源 */
type TaobaoWlbImportThreeplResourceGetResponse struct {
    
    /* result Objectresult */
    result TopResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbImportThreeplOfflineConsignRequest 菜鸟认证直邮线下发货 */
type TaobaoWlbImportThreeplOfflineConsignRequest struct {
    
    /* from_id optional发件人地址库id */
    from_id int64 `json:"from_id";xml:"from_id"`
    
    /* res_code optional资源code */
    res_code string `json:"res_code";xml:"res_code"`
    
    /* res_id optional资源id */
    res_id int64 `json:"res_id";xml:"res_id"`
    
    /* trade_id optional交易单号 */
    trade_id int64 `json:"trade_id";xml:"trade_id"`
    
    /* waybill_no optional运单号 */
    waybill_no string `json:"waybill_no";xml:"waybill_no"`
    
}

func (req *TaobaoWlbImportThreeplOfflineConsignRequest) GetAPIName() string {
	return "taobao.wlb.import.threepl.offline.consign"
}

/* TaobaoWlbImportThreeplOfflineConsignResponse 菜鸟认证直邮线下发货 */
type TaobaoWlbImportThreeplOfflineConsignResponse struct {
    
    /* result Objectresult */
    result TopResult `json:"result";xml:"result"`
    
}

/* TmallItemOuteridUpdateRequest 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名） */
type TmallItemOuteridUpdateRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* outer_id optional商品维度商家编码，如果不修改可以不传；清空请设置空串 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* sku_outers optional商品SKU更新OuterId时候用的数据 */
    sku_outers UpdateSkuOuterId `json:"sku_outers";xml:"sku_outers"`
    
}

func (req *TmallItemOuteridUpdateRequest) GetAPIName() string {
	return "tmall.item.outerid.update"
}

/* TmallItemOuteridUpdateResponse 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名） */
type TmallItemOuteridUpdateResponse struct {
    
    /* outerid_update_result Basic商家编码更新结果 */
    outerid_update_result string `json:"outerid_update_result";xml:"outerid_update_result"`
    
}

/* TaobaoQimenOrderSnReportRequest WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表 */
type TaobaoQimenOrderSnReportRequest struct {
    
    /* currentPage required当前页(从1开始) */
    currentPage int64 `json:"currentPage";xml:"currentPage"`
    
    /* extendProps optional扩展属性 */
    extendProps Map `json:"extendProps";xml:"extendProps"`
    
    /* items optional商品列表 */
    items Items `json:"items";xml:"items"`
    
    /* order optional单据列表 */
    order Order `json:"order";xml:"order"`
    
    /* pageSize required每页记录的条数 */
    pageSize int64 `json:"pageSize";xml:"pageSize"`
    
    /* totalPage required总页数 */
    totalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

func (req *TaobaoQimenOrderSnReportRequest) GetAPIName() string {
	return "taobao.qimen.order.sn.report"
}

/* TaobaoQimenOrderSnReportResponse WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表 */
type TaobaoQimenOrderSnReportResponse struct {
    
    /* code Basic响应码 */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息 */
    message string `json:"message";xml:"message"`
    
}

/* QimenTaobaoAlphaxOpenJxtExpressRequest isv 接入指定快递api */
type QimenTaobaoAlphaxOpenJxtExpressRequest struct {
    
    /* cap required快递公司名称 */
    cap string `json:"cap";xml:"cap"`
    
    /* order_id required主订单id */
    order_id string `json:"order_id";xml:"order_id"`
    
    /* seller_id required卖家主账号 */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtExpressRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.express"
}

/* QimenTaobaoAlphaxOpenJxtExpressResponse isv 接入指定快递api */
type QimenTaobaoAlphaxOpenJxtExpressResponse struct {
    
    /* result Object返回值 */
    result ResultDO `json:"result";xml:"result"`
    
}

/* QimenTaobaoAlphaxOpenJxtDelivergoodsRequest isv 接入催发货接口api */
type QimenTaobaoAlphaxOpenJxtDelivergoodsRequest struct {
    
    /* order_id required主订单id */
    order_id string `json:"order_id";xml:"order_id"`
    
    /* seller_id required卖家主账号 */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtDelivergoodsRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.delivergoods"
}

/* QimenTaobaoAlphaxOpenJxtDelivergoodsResponse isv 接入催发货接口api */
type QimenTaobaoAlphaxOpenJxtDelivergoodsResponse struct {
    
    /* result Object返回值 */
    result ResultDO `json:"result";xml:"result"`
    
}

/* TaobaoSkusQuantityUpdateRequest 提供按照全量/增量的方式批量修改SKU库存的功能 */
type TaobaoSkusQuantityUpdateRequest struct {
    
    /* num_iid required商品数字ID，必填参数 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outerid_quantities special特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。 */
    outerid_quantities string `json:"outerid_quantities";xml:"outerid_quantities"`
    
    /* skuid_quantities specialsku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。 */
    skuid_quantities string `json:"skuid_quantities";xml:"skuid_quantities"`
    
    /* _type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0. */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoSkusQuantityUpdateRequest) GetAPIName() string {
	return "taobao.skus.quantity.update"
}

/* TaobaoSkusQuantityUpdateResponse 提供按照全量/增量的方式批量修改SKU库存的功能 */
type TaobaoSkusQuantityUpdateResponse struct {
    
    /* item Objectiid、numIid、num和modified，skus中每个sku的skuId、quantity和modified */
    item Item `json:"item";xml:"item"`
    
}

/* CainiaoEndpointLockerTopWithholdQueryRequest 查询是否有代扣欠款，是否签署代扣协议。 */
type CainiaoEndpointLockerTopWithholdQueryRequest struct {
    
    /* company_code required柜子公司编码 */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* open_user_id required开放用户Id */
    open_user_id string `json:"open_user_id";xml:"open_user_id"`
    
    /* order_type required柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用） */
    order_type int64 `json:"order_type";xml:"order_type"`
    
}

func (req *CainiaoEndpointLockerTopWithholdQueryRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.withhold.query"
}

/* CainiaoEndpointLockerTopWithholdQueryResponse 查询是否有代扣欠款，是否签署代扣协议。 */
type CainiaoEndpointLockerTopWithholdQueryResponse struct {
    
    /* result Objectresponse */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TaobaoWangwangEserviceChatrelationGetRequest 获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B */
type TaobaoWangwangEserviceChatrelationGetRequest struct {
    
    /* chat_relation_request required请求参数 */
    chat_relation_request ChatRelationRequest `json:"chat_relation_request";xml:"chat_relation_request"`
    
}

func (req *TaobaoWangwangEserviceChatrelationGetRequest) GetAPIName() string {
	return "taobao.wangwang.eservice.chatrelation.get"
}

/* TaobaoWangwangEserviceChatrelationGetResponse 获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B */
type TaobaoWangwangEserviceChatrelationGetResponse struct {
    
    /* result Objectresult */
    result ChatRelationResult `json:"result";xml:"result"`
    
}

/* CainiaoEndpointLockerTopOrderWithholdRequest 提供代扣，允许有一笔欠款。 */
type CainiaoEndpointLockerTopOrderWithholdRequest struct {
    
    /* company_code required柜子公司编码 */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* extra optional扩展字段 */
    extra string `json:"extra";xml:"extra"`
    
    /* gui_id required柜子id */
    gui_id string `json:"gui_id";xml:"gui_id"`
    
    /* mail_no required运单号 */
    mail_no string `json:"mail_no";xml:"mail_no"`
    
    /* open_user_id required开放用户id */
    open_user_id string `json:"open_user_id";xml:"open_user_id"`
    
    /* order_code required柜子订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_type required订单类型(0-取件业务，1-寄件业务，2-派样业务) */
    order_type int64 `json:"order_type";xml:"order_type"`
    
    /* total_fee required代扣金额（全额），单位：分 */
    total_fee int64 `json:"total_fee";xml:"total_fee"`
    
}

func (req *CainiaoEndpointLockerTopOrderWithholdRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.withhold"
}

/* CainiaoEndpointLockerTopOrderWithholdResponse 提供代扣，允许有一笔欠款。 */
type CainiaoEndpointLockerTopOrderWithholdResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbItemAddRequest 添加物流宝商品，支持物流宝子商品和属性添加 */
type TaobaoWlbItemAddRequest struct {
    
    /* color optional商品颜色 */
    color string `json:"color";xml:"color"`
    
    /* goods_cat optional货类 */
    goods_cat string `json:"goods_cat";xml:"goods_cat"`
    
    /* height optional商品高度，单位毫米 */
    height int64 `json:"height";xml:"height"`
    
    /* is_dangerous optional是否危险品 */
    is_dangerous bool `json:"is_dangerous";xml:"is_dangerous"`
    
    /* is_friable optional是否易碎品 */
    is_friable bool `json:"is_friable";xml:"is_friable"`
    
    /* is_sku required是否sku */
    is_sku string `json:"is_sku";xml:"is_sku"`
    
    /* item_code required商品编码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* length optional商品长度，单位毫米 */
    length int64 `json:"length";xml:"length"`
    
    /* name required商品名称 */
    name string `json:"name";xml:"name"`
    
    /* package_material optional商品包装材料类型 */
    package_material string `json:"package_material";xml:"package_material"`
    
    /* price optional商品价格，单位：分 */
    price int64 `json:"price";xml:"price"`
    
    /* pricing_cat optional计价货类 */
    pricing_cat string `json:"pricing_cat";xml:"pricing_cat"`
    
    /* pro_name_list optional属性名列表,目前支持的属性：
毛重:GWeight	
净重:Nweight
皮重: Tweight
自定义属性：
paramkey1
paramkey2
paramkey3
paramkey4 */
    pro_name_list string `json:"pro_name_list";xml:"pro_name_list"`
    
    /* pro_value_list optional属性值列表：
10,8 */
    pro_value_list string `json:"pro_value_list";xml:"pro_value_list"`
    
    /* remark optional商品备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* support_batch optional是否支持批次 */
    support_batch bool `json:"support_batch";xml:"support_batch"`
    
    /* title optional商品标题 */
    title string `json:"title";xml:"title"`
    
    /* _type optionalNORMAL--普通商品
COMBINE--组合商品
DISTRIBUTION--分销 */
    _type string `json:"type";xml:"type"`
    
    /* volume optional商品体积，单位立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* weight optional商品重量，单位G */
    weight int64 `json:"weight";xml:"weight"`
    
    /* width optional商品宽度，单位毫米 */
    width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbItemAddRequest) GetAPIName() string {
	return "taobao.wlb.item.add"
}

/* TaobaoWlbItemAddResponse 添加物流宝商品，支持物流宝子商品和属性添加 */
type TaobaoWlbItemAddResponse struct {
    
    /* item_id Basic新增的商品 */
    item_id map[string]interface{} `json:"item_id";xml:"item_id"`
    
}

/* TaobaoWlbWlborderGetRequest 根据物流宝订单编号查询物流宝订单概要信息 */
type TaobaoWlbWlborderGetRequest struct {
    
    /* wlb_order_code required物流宝订单编码 */
    wlb_order_code string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbWlborderGetRequest) GetAPIName() string {
	return "taobao.wlb.wlborder.get"
}

/* TaobaoWlbWlborderGetResponse 根据物流宝订单编号查询物流宝订单概要信息 */
type TaobaoWlbWlborderGetResponse struct {
    
    /* wlb_order Object物流宝订单对象 */
    wlb_order WlbOrder `json:"wlb_order";xml:"wlb_order"`
    
}

/* CainiaoNbaddAppointdeliverGetconsigninfoRequest 获取支持定时派送服务发货信息 */
type CainiaoNbaddAppointdeliverGetconsigninfoRequest struct {
    
    /* trade_order_id required淘宝交易订单id */
    trade_order_id int64 `json:"trade_order_id";xml:"trade_order_id"`
    
}

func (req *CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetAPIName() string {
	return "cainiao.nbadd.appointdeliver.getconsigninfo"
}

/* CainiaoNbaddAppointdeliverGetconsigninfoResponse 获取支持定时派送服务发货信息 */
type CainiaoNbaddAppointdeliverGetconsigninfoResponse struct {
    
    /* is_success Basic调用是否成功，正常情况下都会成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* result Object发货信息 */
    result ConsignSupportInfoDTO `json:"result";xml:"result"`
    
    /* result_code Basic错误编码 */
    result_code string `json:"result_code";xml:"result_code"`
    
    /* result_desc Basic错误描述 */
    result_desc string `json:"result_desc";xml:"result_desc"`
    
}

/* TaobaoOmniorderDtdResendRequest 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次 */
type TaobaoOmniorderDtdResendRequest struct {
    
    /* main_order_id required淘宝主订单ID */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
}

func (req *TaobaoOmniorderDtdResendRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.resend"
}

/* TaobaoOmniorderDtdResendResponse 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次 */
type TaobaoOmniorderDtdResendResponse struct {
    
    /* message Basic错误信息 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}

/* TaobaoOmniorderDtdConsignRequest 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单 */
type TaobaoOmniorderDtdConsignRequest struct {
    
    /* main_order_id required淘宝订单主订单号 */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
    /* store_id optional发货对应的商户中心门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderDtdConsignRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.consign"
}

/* TaobaoOmniorderDtdConsignResponse 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单 */
type TaobaoOmniorderDtdConsignResponse struct {
    
    /* message Basic错误信息 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}

/* TaobaoOmniorderDtdConsumeRequest 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。 */
type TaobaoOmniorderDtdConsumeRequest struct {
    
    /* param_door2door_consume_request optional核销信息 */
    param_door2door_consume_request Door2doorConsumeRequest `json:"param_door2door_consume_request";xml:"param_door2door_consume_request"`
    
}

func (req *TaobaoOmniorderDtdConsumeRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.consume"
}

/* TaobaoOmniorderDtdConsumeResponse 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。 */
type TaobaoOmniorderDtdConsumeResponse struct {
    
    /* message Basic错误西溪 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}

/* TaobaoOmniorderDtdQueryRequest 门店自送根据核销码码查询订单信息 */
type TaobaoOmniorderDtdQueryRequest struct {
    
    /* code required核销码 */
    code string `json:"code";xml:"code"`
    
}

func (req *TaobaoOmniorderDtdQueryRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.query"
}

/* TaobaoOmniorderDtdQueryResponse 门店自送根据核销码码查询订单信息 */
type TaobaoOmniorderDtdQueryResponse struct {
    
    /* data Objectdata */
    data Door2doorQueryResult `json:"data";xml:"data"`
    
    /* message Basic错误信息 */
    message string `json:"message";xml:"message"`
    
    /* result_code Basic错误码，为0表示成功，非0表示失败 */
    result_code string `json:"result_code";xml:"result_code"`
    
}

/* TaobaoWlbOrderstatusGetRequest 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表 */
type TaobaoWlbOrderstatusGetRequest struct {
    
    /* order_code required物流宝订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbOrderstatusGetRequest) GetAPIName() string {
	return "taobao.wlb.orderstatus.get"
}

/* TaobaoWlbOrderstatusGetResponse 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表 */
type TaobaoWlbOrderstatusGetResponse struct {
    
    /* wlb_order_status Object Array订单流转信息状态列表 */
    wlb_order_status WlbProcessStatus `json:"wlb_order_status";xml:"wlb_order_status"`
    
}

/* TmallItemSimpleschemaUpdateRequest 国外大商家天猫简化编辑商品 */
type TmallItemSimpleschemaUpdateRequest struct {
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* schema_xml_fields required编辑商品时提交的xml信息 */
    schema_xml_fields string `json:"schema_xml_fields";xml:"schema_xml_fields"`
    
}

func (req *TmallItemSimpleschemaUpdateRequest) GetAPIName() string {
	return "tmall.item.simpleschema.update"
}

/* TmallItemSimpleschemaUpdateResponse 国外大商家天猫简化编辑商品 */
type TmallItemSimpleschemaUpdateResponse struct {
    
    /* gmt_modified Basic编辑商品操作成功时间 */
    gmt_modified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    /* sku_map_json Basicsku与outerId映射信息 */
    sku_map_json string `json:"sku_map_json";xml:"sku_map_json"`
    
    /* update_item_result Basic编辑商品的itemid */
    update_item_result string `json:"update_item_result";xml:"update_item_result"`
    
}

/* TaobaoWlbItemGetRequest 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。 */
type TaobaoWlbItemGetRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemGetRequest) GetAPIName() string {
	return "taobao.wlb.item.get"
}

/* TaobaoWlbItemGetResponse 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。 */
type TaobaoWlbItemGetResponse struct {
    
    /* item Object商品信息 */
    item WlbItem `json:"item";xml:"item"`
    
}

/* TaobaoWlbOrderCancelRequest 取消物流宝订单 */
type TaobaoWlbOrderCancelRequest struct {
    
    /* wlb_order_code required物流宝订单编号 */
    wlb_order_code string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbOrderCancelRequest) GetAPIName() string {
	return "taobao.wlb.order.cancel"
}

/* TaobaoWlbOrderCancelResponse 取消物流宝订单 */
type TaobaoWlbOrderCancelResponse struct {
    
    /* error_code_list Basic错误编码列表 */
    error_code_list string `json:"error_code_list";xml:"error_code_list"`
    
    /* modify_time Basic修改时间，只有在取消成功的情况下，才可以做 */
    modify_time time.Time `json:"modify_time";xml:"modify_time"`
    
}

/* TaobaoWlbInventoryDetailGetRequest 查询库存详情，通过商品ID获取发送请求的卖家的库存详情 */
type TaobaoWlbInventoryDetailGetRequest struct {
    
    /* inventory_type_list optional库存类型列表，值包括：
VENDIBLE--可销售库存
FREEZE--冻结库存
ONWAY--在途库存
DEFECT--残次品库存
ENGINE_DAMAGE--机损
BOX_DAMAGE--箱损
EXPIRATION--过保 */
    inventory_type_list string `json:"inventory_type_list";xml:"inventory_type_list"`
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbInventoryDetailGetRequest) GetAPIName() string {
	return "taobao.wlb.inventory.detail.get"
}

/* TaobaoWlbInventoryDetailGetResponse 查询库存详情，通过商品ID获取发送请求的卖家的库存详情 */
type TaobaoWlbInventoryDetailGetResponse struct {
    
    /* inventory_list Object Array库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性 */
    inventory_list WlbInventory `json:"inventory_list";xml:"inventory_list"`
    
    /* item_id Basic入参的item_id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

/* TaobaoWlbTradeorderGetRequest 根据交易类型和交易id查询物流宝订单详情 */
type TaobaoWlbTradeorderGetRequest struct {
    
    /* sub_trade_id optional子交易号 */
    sub_trade_id string `json:"sub_trade_id";xml:"sub_trade_id"`
    
    /* trade_id required指定交易类型的交易号 */
    trade_id string `json:"trade_id";xml:"trade_id"`
    
    /* trade_type required交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易 */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoWlbTradeorderGetRequest) GetAPIName() string {
	return "taobao.wlb.tradeorder.get"
}

/* TaobaoWlbTradeorderGetResponse 根据交易类型和交易id查询物流宝订单详情 */
type TaobaoWlbTradeorderGetResponse struct {
    
    /* wlb_order_list Object Array物流宝订单对象 */
    wlb_order_list WlbOrder `json:"wlb_order_list";xml:"wlb_order_list"`
    
}

/* TaobaoWlbItemQueryRequest 根据状态、卖家、SKU等信息查询商品列表 */
type TaobaoWlbItemQueryRequest struct {
    
    /* is_sku optional是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理 */
    is_sku string `json:"is_sku";xml:"is_sku"`
    
    /* item_code optional商家编码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* item_type optionalITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理 */
    item_type string `json:"item_type";xml:"item_type"`
    
    /* name optional商品名称 */
    name string `json:"name";xml:"name"`
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* parent_id optional父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品 */
    parent_id int64 `json:"parent_id";xml:"parent_id"`
    
    /* status optional只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理 */
    status string `json:"status";xml:"status"`
    
    /* title optional商品前台销售名字 */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoWlbItemQueryRequest) GetAPIName() string {
	return "taobao.wlb.item.query"
}

/* TaobaoWlbItemQueryResponse 根据状态、卖家、SKU等信息查询商品列表 */
type TaobaoWlbItemQueryResponse struct {
    
    /* item_list Object Array商品信息列表 */
    item_list WlbItem `json:"item_list";xml:"item_list"`
    
    /* total_count Basic结果总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoWlbOrderCreateRequest 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生 */
type TaobaoWlbOrderCreateRequest struct {
    
    /* alipay_no optional支付宝交易号 */
    alipay_no string `json:"alipay_no";xml:"alipay_no"`
    
    /* attributes optional该字段暂时保留 */
    attributes string `json:"attributes";xml:"attributes"`
    
    /* buyer_nick optional买家呢称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* expect_end_time optional期望结束时间，在入库单会使用到 */
    expect_end_time time.Time `json:"expect_end_time";xml:"expect_end_time"`
    
    /* expect_start_time optional计划开始送达时间  在入库单中可能会使用 */
    expect_start_time time.Time `json:"expect_start_time";xml:"expect_start_time"`
    
    /* invoince_info optional{"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]} */
    invoince_info string `json:"invoince_info";xml:"invoince_info"`
    
    /* is_finished required该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。 */
    is_finished bool `json:"is_finished";xml:"is_finished"`
    
    /* order_code optional物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_flag optional用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，
: 是否可改配送方式  默认可更改
11 CONSIGN 物流宝代理发货,自动更改发货状态
12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票 */
    order_flag string `json:"order_flag";xml:"order_flag"`
    
    /* order_item_list required订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick
":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。 */
    order_item_list string `json:"order_item_list";xml:"order_item_list"`
    
    /* order_sub_type required订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购 */
    order_sub_type string `json:"order_sub_type";xml:"order_sub_type"`
    
    /* order_type required订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* out_biz_code required外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用 */
    out_biz_code string `json:"out_biz_code";xml:"out_biz_code"`
    
    /* package_count optional包裹件数，入库单和出库单中会用到 */
    package_count int64 `json:"package_count";xml:"package_count"`
    
    /* payable_amount optional应收金额，cod订单必选 */
    payable_amount int64 `json:"payable_amount";xml:"payable_amount"`
    
    /* prev_order_code optional源订单编号 */
    prev_order_code string `json:"prev_order_code";xml:"prev_order_code"`
    
    /* receiver_info optional收货方信息，必须传， 手机和电话必选其一。
收货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话

如果某一个字段的数据为空时，必须传NA */
    receiver_info string `json:"receiver_info";xml:"receiver_info"`
    
    /* remark optional备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* schedule_end optional投递时间范围要求,格式'15:20'用分号隔开 */
    schedule_end string `json:"schedule_end";xml:"schedule_end"`
    
    /* schedule_start optional投递时间范围要求,格式'13:20'用分号隔开 */
    schedule_start string `json:"schedule_start";xml:"schedule_start"`
    
    /* schedule_type optional投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY */
    schedule_type string `json:"schedule_type";xml:"schedule_type"`
    
    /* sender_info optional发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
如果某一个字段的数据为空时，必须传NA */
    sender_info string `json:"sender_info";xml:"sender_info"`
    
    /* service_fee optionalcod服务费，只有cod订单的时候，才需要这个字段 */
    service_fee int64 `json:"service_fee";xml:"service_fee"`
    
    /* store_code required仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
    /* tms_info optional出库单中可能会用到
运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号

========================================
如果某一个字段的数据为空时，必须传NA */
    tms_info string `json:"tms_info";xml:"tms_info"`
    
    /* tms_order_code optional运单编号，退货单时可能会使用 */
    tms_order_code string `json:"tms_order_code";xml:"tms_order_code"`
    
    /* tms_service_code optional物流公司编码 */
    tms_service_code string `json:"tms_service_code";xml:"tms_service_code"`
    
    /* total_amount optional总金额 */
    total_amount int64 `json:"total_amount";xml:"total_amount"`
    
}

func (req *TaobaoWlbOrderCreateRequest) GetAPIName() string {
	return "taobao.wlb.order.create"
}

/* TaobaoWlbOrderCreateResponse 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生 */
type TaobaoWlbOrderCreateResponse struct {
    
    /* create_time Basic订单创建时间 */
    create_time time.Time `json:"create_time";xml:"create_time"`
    
    /* order_code Basic物流宝订单创建成功后，返回物流宝的订单编号；如果订单创建失败，该字段为空。 */
    order_code string `json:"order_code";xml:"order_code"`
    
}

/* QimenTaobaoIcpOrderStockoutordecanceltoerpRequest 出入库单取消推送接口 */
type QimenTaobaoIcpOrderStockoutordecanceltoerpRequest struct {
    
    /* customerId required出库单所属货主ID */
    customerId string `json:"customerId";xml:"customerId"`
    
    /* entryOrderlist optional出入库单信息（一单） */
    entryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockoutordecanceltoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockoutordecanceltoerp"
}

/* QimenTaobaoIcpOrderStockoutordecanceltoerpResponse 出入库单取消推送接口 */
type QimenTaobaoIcpOrderStockoutordecanceltoerpResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}

/* TmallItemDescModulesGetRequest 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。 */
type TmallItemDescModulesGetRequest struct {
    
    /* cat_id required淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122 */
    cat_id int64 `json:"cat_id";xml:"cat_id"`
    
    /* usr_id required商家主帐号id */
    usr_id string `json:"usr_id";xml:"usr_id"`
    
}

func (req *TmallItemDescModulesGetRequest) GetAPIName() string {
	return "tmall.item.desc.modules.get"
}

/* TmallItemDescModulesGetResponse 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。 */
type TmallItemDescModulesGetResponse struct {
    
    /* modular_desc_info Object返回描述模块信息 */
    modular_desc_info ModularDescInfo `json:"modular_desc_info";xml:"modular_desc_info"`
    
}

/* TaobaoLogisticsConsignOrderCreateandsendRequest 创建物流订单，并发货。 */
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

/* TaobaoLogisticsConsignOrderCreateandsendResponse 创建物流订单，并发货。 */
type TaobaoLogisticsConsignOrderCreateandsendResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* order_id Basic订单ID */
    order_id int64 `json:"order_id";xml:"order_id"`
    
    /* result_desc Basic结果描述 */
    result_desc string `json:"result_desc";xml:"result_desc"`
    
}

/* TaobaoInventoryIpcInventorydetailGetRequest 查询库存明细 */
type TaobaoInventoryIpcInventorydetailGetRequest struct {
    
    /* biz_order_id optional主订单号，可选 */
    biz_order_id int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    /* biz_sub_order_id optional子订单号，可选 */
    biz_sub_order_id int64 `json:"biz_sub_order_id";xml:"biz_sub_order_id"`
    
    /* page_index required当前页数 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
    /* page_size required一页显示多少条 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* sc_item_id required仓储商品id */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* status_query required1:查询预扣  4：查询占用 */
    status_query int64 `json:"status_query";xml:"status_query"`
    
}

func (req *TaobaoInventoryIpcInventorydetailGetRequest) GetAPIName() string {
	return "taobao.inventory.ipc.inventorydetail.get"
}

/* TaobaoInventoryIpcInventorydetailGetResponse 查询库存明细 */
type TaobaoInventoryIpcInventorydetailGetResponse struct {
    
    /* inventory_details Object Array库存明细列表 */
    inventory_details IpcInventoryDetailDo `json:"inventory_details";xml:"inventory_details"`
    
}

/* TaobaoWlbWmsConsignInventoryReleaseRequest ERP释放预占用库存 */
type TaobaoWlbWmsConsignInventoryReleaseRequest struct {
    
    /* content optional库存释放 */
    content Content `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsConsignInventoryReleaseRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.inventory.release"
}

/* TaobaoWlbWmsConsignInventoryReleaseResponse ERP释放预占用库存 */
type TaobaoWlbWmsConsignInventoryReleaseResponse struct {
    
    /* result Object接口返回model */
    result Result `json:"result";xml:"result"`
    
}

/* TaobaoWlbWmsConsignInventoryOccupancyRequest ERP发货库存预占用 */
type TaobaoWlbWmsConsignInventoryOccupancyRequest struct {
    
    /* content optional库存占用 */
    content Content `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsConsignInventoryOccupancyRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.inventory.occupancy"
}

/* TaobaoWlbWmsConsignInventoryOccupancyResponse ERP发货库存预占用 */
type TaobaoWlbWmsConsignInventoryOccupancyResponse struct {
    
    /* is_retry Basic返回失败时，是否需求重试，为true时，建议系统自动重试 */
    is_retry bool `json:"is_retry";xml:"is_retry"`
    
    /* item_inventory_list Object Array库存占用明细列表 */
    item_inventory_list Iteminventorylist `json:"item_inventory_list";xml:"item_inventory_list"`
    
    /* order_code BasicERP订单编码 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* wl_error_code Basic错误码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success string `json:"wl_success";xml:"wl_success"`
    
}

/* TaobaoLogisticsAddressSearchRequest 通过此接口查询卖家地址库， */
type TaobaoLogisticsAddressSearchRequest struct {
    
    /* rdef optional可选，参数列表如下：<br><font color='red'>no_def:查询非默认地址<br>get_def:查询默认取货地址，也即对应卖家后台地址库中发货地址（send_def暂不起作用）<br>cancel_def:查询默认退货地址<br>缺省此参数时，查询所有当前用户的地址库</font> */
    rdef string `json:"rdef";xml:"rdef"`
    
}

func (req *TaobaoLogisticsAddressSearchRequest) GetAPIName() string {
	return "taobao.logistics.address.search"
}

/* TaobaoLogisticsAddressSearchResponse 通过此接口查询卖家地址库， */
type TaobaoLogisticsAddressSearchResponse struct {
    
    /* addresses Object Array一组地址库数据， */
    addresses AddressResult `json:"addresses";xml:"addresses"`
    
}

/* TaobaoLogisticsAddressAddRequest 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库 */
type TaobaoLogisticsAddressAddRequest struct {
    
    /* addr required详细街道地址，不需要重复填写省/市/区 */
    addr string `json:"addr";xml:"addr"`
    
    /* cancel_def optional默认退货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
    cancel_def bool `json:"cancel_def";xml:"cancel_def"`
    
    /* city required所在市 */
    city string `json:"city";xml:"city"`
    
    /* contact_name required联系人姓名 <font color='red'>长度不可超过20个字节</font> */
    contact_name string `json:"contact_name";xml:"contact_name"`
    
    /* country optional区、县
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
    country string `json:"country";xml:"country"`
    
    /* get_def optional默认取货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
    get_def bool `json:"get_def";xml:"get_def"`
    
    /* memo optional备注,<br><font color='red'>备注不能超过256字节</font> */
    memo string `json:"memo";xml:"memo"`
    
    /* mobile_phone special手机号码，手机与电话必需有一个
<br><font color='red'>手机号码不能超过20位</font> */
    mobile_phone string `json:"mobile_phone";xml:"mobile_phone"`
    
    /* phone special电话号码,手机与电话必需有一个 */
    phone string `json:"phone";xml:"phone"`
    
    /* province required所在省 */
    province string `json:"province";xml:"province"`
    
    /* seller_company optional公司名称,<br><font color="red">公司名称长能不能超过20字节</font> */
    seller_company string `json:"seller_company";xml:"seller_company"`
    
    /* send_def optional默认发货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认发货地址，撤消原默认发货地址</font> */
    send_def bool `json:"send_def";xml:"send_def"`
    
    /* zip_code optional地区邮政编码
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
    zip_code string `json:"zip_code";xml:"zip_code"`
    
}

func (req *TaobaoLogisticsAddressAddRequest) GetAPIName() string {
	return "taobao.logistics.address.add"
}

/* TaobaoLogisticsAddressAddResponse 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库 */
type TaobaoLogisticsAddressAddResponse struct {
    
    /* address_result Object只返回修改日期modify_date */
    address_result AddressResult `json:"address_result";xml:"address_result"`
    
}

/* TaobaoLogisticsOnlineSendRequest 用户调用该接口可实现在线订单发货（支持货到付款）
调用该接口实现在线下单发货，有两种情况：<br>
<font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br>
如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font> */
type TaobaoLogisticsOnlineSendRequest struct {
    
    /* cancel_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址 */
    cancel_id int64 `json:"cancel_id";xml:"cancel_id"`
    
    /* company_code required物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。 */
    company_code string `json:"company_code";xml:"company_code"`
    
    /* feature optionalfeature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。 */
    feature string `json:"feature";xml:"feature"`
    
    /* is_split optional表明是否是拆单，默认值0，1表示拆单 */
    is_split int64 `json:"is_split";xml:"is_split"`
    
    /* out_sid optional运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
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

func (req *TaobaoLogisticsOnlineSendRequest) GetAPIName() string {
	return "taobao.logistics.online.send"
}

/* TaobaoLogisticsOnlineSendResponse 用户调用该接口可实现在线订单发货（支持货到付款）
调用该接口实现在线下单发货，有两种情况：<br>
<font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br>
如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font> */
type TaobaoLogisticsOnlineSendResponse struct {
    
    /* shipping Objectde */
    shipping Shipping `json:"shipping";xml:"shipping"`
    
}

/* TaobaoLogisticsAddressRemoveRequest 用此接口删除卖家地址库 */
type TaobaoLogisticsAddressRemoveRequest struct {
    
    /* contact_id required地址库ID */
    contact_id int64 `json:"contact_id";xml:"contact_id"`
    
}

func (req *TaobaoLogisticsAddressRemoveRequest) GetAPIName() string {
	return "taobao.logistics.address.remove"
}

/* TaobaoLogisticsAddressRemoveResponse 用此接口删除卖家地址库 */
type TaobaoLogisticsAddressRemoveResponse struct {
    
    /* address_result Object只返回修改日期modify_date */
    address_result AddressResult `json:"address_result";xml:"address_result"`
    
}

/* TaobaoLogisticsAddressModifyRequest 卖家地址库修改 */
type TaobaoLogisticsAddressModifyRequest struct {
    
    /* addr required详细街道地址，不需要重复填写省/市/区 */
    addr string `json:"addr";xml:"addr"`
    
    /* cancel_def optional默认退货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
    cancel_def bool `json:"cancel_def";xml:"cancel_def"`
    
    /* city required所在市 */
    city string `json:"city";xml:"city"`
    
    /* contact_id required地址库ID */
    contact_id int64 `json:"contact_id";xml:"contact_id"`
    
    /* contact_name required联系人姓名
<font color='red'>长度不可超过20个字节</font> */
    contact_name string `json:"contact_name";xml:"contact_name"`
    
    /* country optional区、县
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
    country string `json:"country";xml:"country"`
    
    /* get_def optional默认取货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
    get_def bool `json:"get_def";xml:"get_def"`
    
    /* memo optional备注,<br><font color='red'>备注不能超过256字节</font> */
    memo string `json:"memo";xml:"memo"`
    
    /* mobile_phone special手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font> */
    mobile_phone string `json:"mobile_phone";xml:"mobile_phone"`
    
    /* phone special电话号码,手机与电话必需有一个 */
    phone string `json:"phone";xml:"phone"`
    
    /* province required所在省 */
    province string `json:"province";xml:"province"`
    
    /* seller_company optional公司名称,
<br><font color='red'>公司名称长能不能超过20字节</font> */
    seller_company string `json:"seller_company";xml:"seller_company"`
    
    /* send_def optional默认发货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认发货地址，撤消原默认发货地址</font> */
    send_def bool `json:"send_def";xml:"send_def"`
    
    /* zip_code optional地区邮政编码
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
    zip_code string `json:"zip_code";xml:"zip_code"`
    
}

func (req *TaobaoLogisticsAddressModifyRequest) GetAPIName() string {
	return "taobao.logistics.address.modify"
}

/* TaobaoLogisticsAddressModifyResponse 卖家地址库修改 */
type TaobaoLogisticsAddressModifyResponse struct {
    
    /* address_result Object只返回修改时间modify_date */
    address_result AddressResult `json:"address_result";xml:"address_result"`
    
}

/* TaobaoWlbWmsUnknownPackageUploadRequest 有货无单销退入库单消息回传 */
type TaobaoWlbWmsUnknownPackageUploadRequest struct {
    
    /* content requiredWlbWmsUnknownPackageUpload */
    content WlbWmsUnknownPackageUpload `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsUnknownPackageUploadRequest) GetAPIName() string {
	return "taobao.wlb.wms.unknown.package.upload"
}

/* TaobaoWlbWmsUnknownPackageUploadResponse 有货无单销退入库单消息回传 */
type TaobaoWlbWmsUnknownPackageUploadResponse struct {
    
    /* response ObjectWlbWmsUnknownPackageUploadResp */
    response WlbWmsUnknownPackageUploadResp `json:"response";xml:"response"`
    
}

/* TaobaoTradeReceivetimeDelayRequest 延长交易收货时间 */
type TaobaoTradeReceivetimeDelayRequest struct {
    
    /* days required延长收货的天数，可选值为：3, 5, 7, 10。 */
    days int64 `json:"days";xml:"days"`
    
    /* tid required主订单号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeReceivetimeDelayRequest) GetAPIName() string {
	return "taobao.trade.receivetime.delay"
}

/* TaobaoTradeReceivetimeDelayResponse 延长交易收货时间 */
type TaobaoTradeReceivetimeDelayResponse struct {
    
    /* trade Object更新后的交易数据，只包括tid和modified两个字段。 */
    trade Trade `json:"trade";xml:"trade"`
    
}

/* TaobaoFenxiaoCooperationUpdateRequest 供应商更新合作的分销商等级 */
type TaobaoFenxiaoCooperationUpdateRequest struct {
    
    /* distributor_id required分销商ID */
    distributor_id int64 `json:"distributor_id";xml:"distributor_id"`
    
    /* grade_id required等级ID(0代表取消) */
    grade_id int64 `json:"grade_id";xml:"grade_id"`
    
    /* trade_type optional分销方式(新增)： AGENT(代销)、DEALER(经销)(默认为代销) */
    trade_type string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoCooperationUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.update"
}

/* TaobaoFenxiaoCooperationUpdateResponse 供应商更新合作的分销商等级 */
type TaobaoFenxiaoCooperationUpdateResponse struct {
    
    /* is_success Basic更新结果成功失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoFenxiaoGradesGetRequest 根据供应商ID，查询他的分销商等级信息 */
type TaobaoFenxiaoGradesGetRequest struct {
    
}

func (req *TaobaoFenxiaoGradesGetRequest) GetAPIName() string {
	return "taobao.fenxiao.grades.get"
}

/* TaobaoFenxiaoGradesGetResponse 根据供应商ID，查询他的分销商等级信息 */
type TaobaoFenxiaoGradesGetResponse struct {
    
    /* fenxiao_grades Object Array分销商等级信息 */
    fenxiao_grades FenxiaoGrade `json:"fenxiao_grades";xml:"fenxiao_grades"`
    
}

/* TaobaoFenxiaoDiscountsGetRequest 查询折扣信息 */
type TaobaoFenxiaoDiscountsGetRequest struct {
    
    /* discount_id optional折扣ID */
    discount_id int64 `json:"discount_id";xml:"discount_id"`
    
    /* ext_fields optional指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用） */
    ext_fields string `json:"ext_fields";xml:"ext_fields"`
    
}

func (req *TaobaoFenxiaoDiscountsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.discounts.get"
}

/* TaobaoFenxiaoDiscountsGetResponse 查询折扣信息 */
type TaobaoFenxiaoDiscountsGetResponse struct {
    
    /* discounts Object Array折扣数据结构 */
    discounts Discount `json:"discounts";xml:"discounts"`
    
    /* total_results Basic记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* QimenTaobaoAlphaxOpenJxtInvoiceRequest isv 开发票接口api */
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

/* QimenTaobaoAlphaxOpenJxtInvoiceResponse isv 开发票接口api */
type QimenTaobaoAlphaxOpenJxtInvoiceResponse struct {
    
    /* result Object返回值 */
    result Data `json:"result";xml:"result"`
    
}

/* TaobaoFuwuScoresGetRequest 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟 */
type TaobaoFuwuScoresGetRequest struct {
    
    /* current_page required当前页 */
    current_page int64 `json:"current_page";xml:"current_page"`
    
    /* date required评价日期，查询某一天的评价 */
    date time.Time `json:"date";xml:"date"`
    
    /* page_size optional每页获取条数。默认值40，最小值1，最大值100。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoFuwuScoresGetRequest) GetAPIName() string {
	return "taobao.fuwu.scores.get"
}

/* TaobaoFuwuScoresGetResponse 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟 */
type TaobaoFuwuScoresGetResponse struct {
    
    /* score_result Object Array评价流水记录 */
    score_result ScoreResult `json:"score_result";xml:"score_result"`
    
}

/* TaobaoItemSchemaIncrementUpdateRequest 根据schema规则增量修改宝贝信息 */
type TaobaoItemSchemaIncrementUpdateRequest struct {
    
    /* category_id optional商品类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* parameters required修改字段 */
    parameters string `json:"parameters";xml:"parameters"`
    
}

func (req *TaobaoItemSchemaIncrementUpdateRequest) GetAPIName() string {
	return "taobao.item.schema.increment.update"
}

/* TaobaoItemSchemaIncrementUpdateResponse 根据schema规则增量修改宝贝信息 */
type TaobaoItemSchemaIncrementUpdateResponse struct {
    
    /* item_id Basic商品id */
    item_id string `json:"item_id";xml:"item_id"`
    
}

/* TaobaoVasSubscSearchRequest 用于ISV查询自己名下的应用及收费项目的订购记录 */
type TaobaoVasSubscSearchRequest struct {
    
    /* article_code required应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
    article_code string `json:"article_code";xml:"article_code"`
    
    /* autosub optional是否自动续费，true=自动续费 false=非自动续费 空=全部 */
    autosub bool `json:"autosub";xml:"autosub"`
    
    /* end_deadline optional到期时间结束值 */
    end_deadline time.Time `json:"end_deadline";xml:"end_deadline"`
    
    /* expire_notice optional是否到期提醒，true=到期提醒 false=非到期提醒 空=全部 */
    expire_notice bool `json:"expire_notice";xml:"expire_notice"`
    
    /* item_code optional收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* nick optional淘宝会员名 */
    nick string `json:"nick";xml:"nick"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional一页包含的记录数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_deadline optional到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据） */
    start_deadline time.Time `json:"start_deadline";xml:"start_deadline"`
    
    /* status optional订购记录状态，1=有效 2=过期 空=全部 */
    status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoVasSubscSearchRequest) GetAPIName() string {
	return "taobao.vas.subsc.search"
}

/* TaobaoVasSubscSearchResponse 用于ISV查询自己名下的应用及收费项目的订购记录 */
type TaobaoVasSubscSearchResponse struct {
    
    /* article_subs Object Array订购关系对象 */
    article_subs ArticleSub `json:"article_subs";xml:"article_subs"`
    
    /* total_item Basic总记录数 */
    total_item int64 `json:"total_item";xml:"total_item"`
    
}

/* TaobaoItemQuantityUpdateRequest 提供按照全量或增量形式修改宝贝/SKU库存的功能 */
type TaobaoItemQuantityUpdateRequest struct {
    
    /* num_iid required商品数字ID，必填参数 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outer_id optionalSKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* quantity required库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
    /* sku_id optional要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存 */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
    /* _type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItemQuantityUpdateRequest) GetAPIName() string {
	return "taobao.item.quantity.update"
}

/* TaobaoItemQuantityUpdateResponse 提供按照全量或增量形式修改宝贝/SKU库存的功能 */
type TaobaoItemQuantityUpdateResponse struct {
    
    /* item Objectiid、numIid、num和modified，skus中每个sku的skuId、quantity和modified */
    item Item `json:"item";xml:"item"`
    
}

/* TaobaoVasOrderSearchRequest 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。
建议用于查询前一日的历史记录，不适合用作实时数据查询。
现在只能查询90天以内的数据
该接口限制每分钟所有appkey调用总和只能有800次。 */
type TaobaoVasOrderSearchRequest struct {
    
    /* article_code required应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
    article_code string `json:"article_code";xml:"article_code"`
    
    /* biz_order_id optional订单号 */
    biz_order_id int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    /* biz_type optional订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部 */
    biz_type int64 `json:"biz_type";xml:"biz_type"`
    
    /* end_created optional订单创建时间（订购时间）结束值 */
    end_created time.Time `json:"end_created";xml:"end_created"`
    
    /* item_code optional收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
    item_code string `json:"item_code";xml:"item_code"`
    
    /* nick optional淘宝会员名 */
    nick string `json:"nick";xml:"nick"`
    
    /* order_id optional子订单号 */
    order_id int64 `json:"order_id";xml:"order_id"`
    
    /* page_no optional页码 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional一页包含的记录数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_created optional订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据） */
    start_created time.Time `json:"start_created";xml:"start_created"`
    
}

func (req *TaobaoVasOrderSearchRequest) GetAPIName() string {
	return "taobao.vas.order.search"
}

/* TaobaoVasOrderSearchResponse 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。
建议用于查询前一日的历史记录，不适合用作实时数据查询。
现在只能查询90天以内的数据
该接口限制每分钟所有appkey调用总和只能有800次。 */
type TaobaoVasOrderSearchResponse struct {
    
    /* article_biz_orders Object Array商品订单对象 */
    article_biz_orders ArticleBizOrder `json:"article_biz_orders";xml:"article_biz_orders"`
    
    /* total_item Basic总记录数 */
    total_item int64 `json:"total_item";xml:"total_item"`
    
}

/* TaobaoVasSubscribeGetRequest 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况 */
type TaobaoVasSubscribeGetRequest struct {
    
    /* article_code required商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码 */
    article_code string `json:"article_code";xml:"article_code"`
    
    /* nick required淘宝会员名 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoVasSubscribeGetRequest) GetAPIName() string {
	return "taobao.vas.subscribe.get"
}

/* TaobaoVasSubscribeGetResponse 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况 */
type TaobaoVasSubscribeGetResponse struct {
    
    /* article_user_subscribes Object Array用户订购信息 */
    article_user_subscribes ArticleUserSubscribe `json:"article_user_subscribes";xml:"article_user_subscribes"`
    
}

/* CainiaoWaybillIiQueryByWaybillcodeRequest 通过面单号查看面单号的当前状态，如签收、发货、失效等。 */
type CainiaoWaybillIiQueryByWaybillcodeRequest struct {
    
    /* param_list optional系统自动生成 */
    param_list WaybillDetailQueryByWaybillCodeRequest `json:"param_list";xml:"param_list"`
    
}

func (req *CainiaoWaybillIiQueryByWaybillcodeRequest) GetAPIName() string {
	return "cainiao.waybill.ii.query.by.waybillcode"
}

/* CainiaoWaybillIiQueryByWaybillcodeResponse 通过面单号查看面单号的当前状态，如签收、发货、失效等。 */
type CainiaoWaybillIiQueryByWaybillcodeResponse struct {
    
    /* modules Object Array查询返回值 */
    modules WaybillCloudPrintWithResultDescResponse `json:"modules";xml:"modules"`
    
}

/* CainiaoWaybillIiQueryByTradecodeRequest 通过订单号查看面单的信息 */
type CainiaoWaybillIiQueryByTradecodeRequest struct {
    
    /* param_list optional订单号列表 */
    param_list WaybillDetailQueryByBizSubCodeRequest `json:"param_list";xml:"param_list"`
    
}

func (req *CainiaoWaybillIiQueryByTradecodeRequest) GetAPIName() string {
	return "cainiao.waybill.ii.query.by.tradecode"
}

/* CainiaoWaybillIiQueryByTradecodeResponse 通过订单号查看面单的信息 */
type CainiaoWaybillIiQueryByTradecodeResponse struct {
    
    /* modules Object Array查询返回值 */
    modules WaybillCloudPrintWithResultDescResponse `json:"modules";xml:"modules"`
    
}

/* TaobaoWlbOrderConsignRequest 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货 */
type TaobaoWlbOrderConsignRequest struct {
    
    /* wlb_order_code required物流宝订单编号 */
    wlb_order_code string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbOrderConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.consign"
}

/* TaobaoWlbOrderConsignResponse 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货 */
type TaobaoWlbOrderConsignResponse struct {
    
    /* modify_time Basic修改时间 */
    modify_time time.Time `json:"modify_time";xml:"modify_time"`
    
}

/* CainiaoCloudprintSingleCustomareaGetRequest 商家所有快递公司模板只有一个自定义区 */
type CainiaoCloudprintSingleCustomareaGetRequest struct {
    
    /* seller_id optional这是商家用户id */
    seller_id int64 `json:"seller_id";xml:"seller_id"`
    
}

func (req *CainiaoCloudprintSingleCustomareaGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.single.customarea.get"
}

/* CainiaoCloudprintSingleCustomareaGetResponse 商家所有快递公司模板只有一个自定义区 */
type CainiaoCloudprintSingleCustomareaGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* TaobaoQimenShopSynchronizeRequest 店铺同步接口描述 */
type TaobaoQimenShopSynchronizeRequest struct {
    
    /* request optional请求 */
    request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenShopSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.shop.synchronize"
}

/* TaobaoQimenShopSynchronizeResponse 店铺同步接口描述 */
type TaobaoQimenShopSynchronizeResponse struct {
    
    /* response ObjectResponse */
    response Response `json:"response";xml:"response"`
    
}

/* CainiaoCloudprintIsvtemplatesGetRequest 获取商家使用的标准模板 */
type CainiaoCloudprintIsvtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintIsvtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.isvtemplates.get"
}

/* CainiaoCloudprintIsvtemplatesGetResponse 获取商家使用的标准模板 */
type CainiaoCloudprintIsvtemplatesGetResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* AlibabaEinvoiceApplyGetRequest ERP获取开票申请数据 */
type AlibabaEinvoiceApplyGetRequest struct {
    
    /* apply_id optional开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息 */
    apply_id string `json:"apply_id";xml:"apply_id"`
    
    /* platform_tid required平台订单号 */
    platform_tid string `json:"platform_tid";xml:"platform_tid"`
    
}

func (req *AlibabaEinvoiceApplyGetRequest) GetAPIName() string {
	return "alibaba.einvoice.apply.get"
}

/* AlibabaEinvoiceApplyGetResponse ERP获取开票申请数据 */
type AlibabaEinvoiceApplyGetResponse struct {
    
    /* apply_list Object Array开票明细 */
    apply_list Apply `json:"apply_list";xml:"apply_list"`
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoWlbNotifyMessagePageGetRequest 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息 */
type TaobaoWlbNotifyMessagePageGetRequest struct {
    
    /* end_date optional记录截至时间 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* msg_code optional通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功 */
    msg_code string `json:"msg_code";xml:"msg_code"`
    
    /* page_no optional分页查询页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页查询的每页页数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_date optional记录开始时间 */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* status optional消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM */
    status string `json:"status";xml:"status"`
    
}

func (req *TaobaoWlbNotifyMessagePageGetRequest) GetAPIName() string {
	return "taobao.wlb.notify.message.page.get"
}

/* TaobaoWlbNotifyMessagePageGetResponse 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息 */
type TaobaoWlbNotifyMessagePageGetResponse struct {
    
    /* total_count Basic2000-01-01 00:00:00 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
    /* wlb_messages Object Array通道消息 */
    wlb_messages WlbMessage `json:"wlb_messages";xml:"wlb_messages"`
    
}

/* CainiaoCloudprintClientinfoPutRequest 云打印客户端监控信息收集 */
type CainiaoCloudprintClientinfoPutRequest struct {
    
    /* json_data required客户端上传json数据 */
    json_data string `json:"json_data";xml:"json_data"`
    
}

func (req *CainiaoCloudprintClientinfoPutRequest) GetAPIName() string {
	return "cainiao.cloudprint.clientinfo.put"
}

/* CainiaoCloudprintClientinfoPutResponse 云打印客户端监控信息收集 */
type CainiaoCloudprintClientinfoPutResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

/* TaobaoWlbNotifyMessageConfirmRequest 确认物流宝可执行消息 */
type TaobaoWlbNotifyMessageConfirmRequest struct {
    
    /* message_id required物流宝通知消息的id，通过taobao.wlb.notify.message.page.get接口得到的WlbMessage数据结构中的id字段 */
    message_id string `json:"message_id";xml:"message_id"`
    
}

func (req *TaobaoWlbNotifyMessageConfirmRequest) GetAPIName() string {
	return "taobao.wlb.notify.message.confirm"
}

/* TaobaoWlbNotifyMessageConfirmResponse 确认物流宝可执行消息 */
type TaobaoWlbNotifyMessageConfirmResponse struct {
    
    /* gmt_modified Basic物流宝消息确认时间 */
    gmt_modified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
}

/* TaobaoWlbItemCombinationGetRequest 根据商品id查询商品组合关系 */
type TaobaoWlbItemCombinationGetRequest struct {
    
    /* item_id required要查询的组合商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemCombinationGetRequest) GetAPIName() string {
	return "taobao.wlb.item.combination.get"
}

/* TaobaoWlbItemCombinationGetResponse 根据商品id查询商品组合关系 */
type TaobaoWlbItemCombinationGetResponse struct {
    
    /* item_id_list Basic Array组合子商品id列表 */
    item_id_list int64 `json:"item_id_list";xml:"item_id_list"`
    
}

/* TaobaoWlbItemMapGetRequest 根据物流宝商品ID查询商品映射关系 */
type TaobaoWlbItemMapGetRequest struct {
    
    /* item_id required要查询映射关系的物流宝商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemMapGetRequest) GetAPIName() string {
	return "taobao.wlb.item.map.get"
}

/* TaobaoWlbItemMapGetResponse 根据物流宝商品ID查询商品映射关系 */
type TaobaoWlbItemMapGetResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* out_entity_item_list Object Array外部商品实体 */
    out_entity_item_list OutEntityItem `json:"out_entity_item_list";xml:"out_entity_item_list"`
    
}

/* TaobaoWlbItemMapGetByExtentityRequest 根据外部实体类型和实体id查询映射的物流宝商品id */
type TaobaoWlbItemMapGetByExtentityRequest struct {
    
    /* ext_entity_id required外部实体类型对应的商品id */
    ext_entity_id int64 `json:"ext_entity_id";xml:"ext_entity_id"`
    
    /* ext_entity_type required外部实体类型： IC_ITEM--ic商品 IC_SKU--ic销售单元 */
    ext_entity_type string `json:"ext_entity_type";xml:"ext_entity_type"`
    
}

func (req *TaobaoWlbItemMapGetByExtentityRequest) GetAPIName() string {
	return "taobao.wlb.item.map.get.by.extentity"
}

/* TaobaoWlbItemMapGetByExtentityResponse 根据外部实体类型和实体id查询映射的物流宝商品id */
type TaobaoWlbItemMapGetByExtentityResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* item_id Basic物流宝商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

/* TaobaoWlbTmsorderQueryRequest 通过物流订单编号分页查询物流信息 */
type TaobaoWlbTmsorderQueryRequest struct {
    
    /* order_code required物流订单编号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoWlbTmsorderQueryRequest) GetAPIName() string {
	return "taobao.wlb.tmsorder.query"
}

/* TaobaoWlbTmsorderQueryResponse 通过物流订单编号分页查询物流信息 */
type TaobaoWlbTmsorderQueryResponse struct {
    
    /* tms_order_list Object Array物流订单运单信息列表 */
    tms_order_list WlbTmsOrder `json:"tms_order_list";xml:"tms_order_list"`
    
    /* total_count Basic结果总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoItemIncrementUpdateSchemaGetRequest 获取增量更新规则，目前开发的字段有title, subTitle, description, wl_description */
type TaobaoItemIncrementUpdateSchemaGetRequest struct {
    
    /* category_id optional宝贝类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required宝贝id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* update_fields optional修改字段 */
    update_fields string `json:"update_fields";xml:"update_fields"`
    
}

func (req *TaobaoItemIncrementUpdateSchemaGetRequest) GetAPIName() string {
	return "taobao.item.increment.update.schema.get"
}

/* TaobaoItemIncrementUpdateSchemaGetResponse 获取增量更新规则，目前开发的字段有title, subTitle, description, wl_description */
type TaobaoItemIncrementUpdateSchemaGetResponse struct {
    
    /* update_rules Basic返回的结果集 */
    update_rules string `json:"update_rules";xml:"update_rules"`
    
}

/* TaobaoWlbOrderPageGetRequest 分页查询物流宝订单 */
type TaobaoWlbOrderPageGetRequest struct {
    
    /* end_time optional查询截止时间 */
    end_time time.Time `json:"end_time";xml:"end_time"`
    
    /* order_code optional物流订单编号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* order_status optionalTMS拒签：-100 接收方拒签：-200 */
    order_status int64 `json:"order_status";xml:"order_status"`
    
    /* order_sub_type optional订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单 */
    order_sub_type string `json:"order_sub_type";xml:"order_sub_type"`
    
    /* order_type optional订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库 */
    order_type string `json:"order_type";xml:"order_type"`
    
    /* page_no optional分页的第几页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页多少条 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time optional查询开始时间 */
    start_time time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoWlbOrderPageGetRequest) GetAPIName() string {
	return "taobao.wlb.order.page.get"
}

/* TaobaoWlbOrderPageGetResponse 分页查询物流宝订单 */
type TaobaoWlbOrderPageGetResponse struct {
    
    /* order_list Object Array物流宝订单对象 */
    order_list WlbOrder `json:"order_list";xml:"order_list"`
    
    /* total_count Basic总条数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoWlbSubscriptionQueryRequest 查询商家定购的所有服务,可通过入参状态来筛选 */
type TaobaoWlbSubscriptionQueryRequest struct {
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* status optional状态 
AUDITING 1-待审核; 
CANCEL 2-撤销 ;
CHECKED 3-审核通过 ;
FAILED 4-审核未通过 ;
SYNCHRONIZING 5-同步中;
只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。 */
    status string `json:"status";xml:"status"`
    
}

func (req *TaobaoWlbSubscriptionQueryRequest) GetAPIName() string {
	return "taobao.wlb.subscription.query"
}

/* TaobaoWlbSubscriptionQueryResponse 查询商家定购的所有服务,可通过入参状态来筛选 */
type TaobaoWlbSubscriptionQueryResponse struct {
    
    /* seller_subscription_list Object Array卖家定购的服务列表 */
    seller_subscription_list WlbSellerSubscription `json:"seller_subscription_list";xml:"seller_subscription_list"`
    
    /* total_count Basic结果总数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoWlbItemUpdateRequest 修改物流宝商品信息 */
type TaobaoWlbItemUpdateRequest struct {
    
    /* color optional商品颜色 */
    color string `json:"color";xml:"color"`
    
    /* delete_property_key_list optional需要删除的商品属性key列表 */
    delete_property_key_list string `json:"delete_property_key_list";xml:"delete_property_key_list"`
    
    /* goods_cat optional商品货类 */
    goods_cat string `json:"goods_cat";xml:"goods_cat"`
    
    /* height optional商品高度，单位厘米 */
    height int64 `json:"height";xml:"height"`
    
    /* id required要修改的商品id */
    id int64 `json:"id";xml:"id"`
    
    /* is_dangerous optional是否危险品 */
    is_dangerous bool `json:"is_dangerous";xml:"is_dangerous"`
    
    /* is_friable optional是否易碎品 */
    is_friable bool `json:"is_friable";xml:"is_friable"`
    
    /* length optional商品长度，单位厘米 */
    length int64 `json:"length";xml:"length"`
    
    /* name optional要修改的商品名称 */
    name string `json:"name";xml:"name"`
    
    /* package_material optional商品包装材料类型 */
    package_material string `json:"package_material";xml:"package_material"`
    
    /* pricing_cat optional商品计价货类 */
    pricing_cat string `json:"pricing_cat";xml:"pricing_cat"`
    
    /* remark optional要修改的商品备注 */
    remark string `json:"remark";xml:"remark"`
    
    /* title optional要修改的商品标题 */
    title string `json:"title";xml:"title"`
    
    /* update_property_key_list optional需要修改的商品属性值的列表，如果属性不存在，则新增属性 */
    update_property_key_list string `json:"update_property_key_list";xml:"update_property_key_list"`
    
    /* update_property_value_list optional需要修改的属性值的列表 */
    update_property_value_list string `json:"update_property_value_list";xml:"update_property_value_list"`
    
    /* volume optional商品体积，单位立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* weight optional商品重量，单位G */
    weight int64 `json:"weight";xml:"weight"`
    
    /* width optional商品宽度，单位厘米 */
    width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbItemUpdateRequest) GetAPIName() string {
	return "taobao.wlb.item.update"
}

/* TaobaoWlbItemUpdateResponse 修改物流宝商品信息 */
type TaobaoWlbItemUpdateResponse struct {
    
    /* gmt_modified Basic修改时间 */
    gmt_modified bool `json:"gmt_modified";xml:"gmt_modified"`
    
}

/* TaobaoWlbInventorylogQueryRequest 通过商品ID等几个条件来分页查询库存变更记录 */
type TaobaoWlbInventorylogQueryRequest struct {
    
    /* gmt_end optional结束修改时间,小于等于该时间 */
    gmt_end time.Time `json:"gmt_end";xml:"gmt_end"`
    
    /* gmt_start optional起始修改时间,大于等于该时间 */
    gmt_start time.Time `json:"gmt_start";xml:"gmt_start"`
    
    /* item_id optional商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* op_type optional库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理 */
    op_type string `json:"op_type";xml:"op_type"`
    
    /* op_user_id optional可指定授权的用户来查询 */
    op_user_id int64 `json:"op_user_id";xml:"op_user_id"`
    
    /* order_code optional单号 */
    order_code string `json:"order_code";xml:"order_code"`
    
    /* page_no optional当前页 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional分页记录个数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* store_code optional仓库编码 */
    store_code string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbInventorylogQueryRequest) GetAPIName() string {
	return "taobao.wlb.inventorylog.query"
}

/* TaobaoWlbInventorylogQueryResponse 通过商品ID等几个条件来分页查询库存变更记录 */
type TaobaoWlbInventorylogQueryResponse struct {
    
    /* inventory_log_list Object Array库存变更记录 */
    inventory_log_list WlbItemInventoryLog `json:"inventory_log_list";xml:"inventory_log_list"`
    
    /* total_count Basic记录数 */
    total_count int64 `json:"total_count";xml:"total_count"`
    
}

/* TaobaoTopatsJushitaJdpDatadeleteRequest 异步删除rds库数据推送表的数据 */
type TaobaoTopatsJushitaJdpDatadeleteRequest struct {
    
    /* end_date required删除数据时间段的结束修改时间，格式为：yyyy-MM-dd HH:mm:ss，结束时间必须为前天的23:59:59秒以前，根据是业务的modified时间 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* rds_name optional数据库实例名，当删除用户的绑定关系已经不在推送里时，此参数必填 */
    rds_name string `json:"rds_name";xml:"rds_name"`
    
    /* start_date required删除数据时间段的起始修改时间，格式为：yyyy-MM-dd HH:mm:ss,根据是业务的modified时间 */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
    /* sync_types required推送的数据类型，可选值为：tb_trade(淘宝交易)，tb_item(淘宝商品)，tb_refund(淘宝退款)，fx_trade(分销订单)，同时删除多种类型以分号分隔，如："tb_trade;tb_item;tb_refund;fx_trade" */
    sync_types string `json:"sync_types";xml:"sync_types"`
    
    /* user_nick required用户昵称，必填 */
    user_nick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoTopatsJushitaJdpDatadeleteRequest) GetAPIName() string {
	return "taobao.topats.jushita.jdp.datadelete"
}

/* TaobaoTopatsJushitaJdpDatadeleteResponse 异步删除rds库数据推送表的数据 */
type TaobaoTopatsJushitaJdpDatadeleteResponse struct {
    
    /* result Objectresult */
    result ResultDo `json:"result";xml:"result"`
    
}

/* CainiaoWaybillIiSearchRequest 获取发货地&CP开通状态&账户的使用情况 */
type CainiaoWaybillIiSearchRequest struct {
    
    /* cp_code optional物流公司code */
    cp_code string `json:"cp_code";xml:"cp_code"`
    
}

func (req *CainiaoWaybillIiSearchRequest) GetAPIName() string {
	return "cainiao.waybill.ii.search"
}

/* CainiaoWaybillIiSearchResponse 获取发货地&CP开通状态&账户的使用情况 */
type CainiaoWaybillIiSearchResponse struct {
    
    /* waybill_apply_subscription_cols Object ArrayCP网点信息及对应的商家的发货信息 */
    waybill_apply_subscription_cols WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_cols";xml:"waybill_apply_subscription_cols"`
    
}

/* TaobaoJushitaJdpUsersGetRequest 获取开通的订单同步服务的用户，含有rds的路由关系 */
type TaobaoJushitaJdpUsersGetRequest struct {
    
    /* end_modified optional此参数一般不用传，用于查询最后更改时间在某个时间段内的用户 */
    end_modified time.Time `json:"end_modified";xml:"end_modified"`
    
    /* page_no optional当前页数 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页记录数，默认200 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_modified optional此参数一般不用传，用于查询最后更改时间在某个时间段内的用户 */
    start_modified time.Time `json:"start_modified";xml:"start_modified"`
    
    /* target_appkey optional普通isv不能传入此参数 */
    target_appkey string `json:"target_appkey";xml:"target_appkey"`
    
    /* user_id optional如果传了user_id表示单条查询 */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoJushitaJdpUsersGetRequest) GetAPIName() string {
	return "taobao.jushita.jdp.users.get"
}

/* TaobaoJushitaJdpUsersGetResponse 获取开通的订单同步服务的用户，含有rds的路由关系 */
type TaobaoJushitaJdpUsersGetResponse struct {
    
    /* total_results Basic总记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
    /* users Object Array用户列表 */
    users JdpUser `json:"users";xml:"users"`
    
}

/* TaobaoShopGetbytitleRequest 根据店铺名称获取店铺信息 */
type TaobaoShopGetbytitleRequest struct {
    
    /* fields optional代表需要获取的店铺信息：sid,cid,title,nick,desc,bulletin,pic_path,created,modified,shop_score */
    fields string `json:"fields";xml:"fields"`
    
    /* title required店铺名称，必须严格匹配（不支持模糊查询） */
    title string `json:"title";xml:"title"`
    
}

func (req *TaobaoShopGetbytitleRequest) GetAPIName() string {
	return "taobao.shop.getbytitle"
}

/* TaobaoShopGetbytitleResponse 根据店铺名称获取店铺信息 */
type TaobaoShopGetbytitleResponse struct {
    
    /* err_code Basic错误码 */
    err_code string `json:"err_code";xml:"err_code"`
    
    /* err_msg Basic错误信息 */
    err_msg string `json:"err_msg";xml:"err_msg"`
    
    /* is_error Basic有无错误 */
    is_error bool `json:"is_error";xml:"is_error"`
    
    /* result_shop Basictest */
    result_shop map[string]interface{} `json:"result_shop";xml:"result_shop"`
    
}

/* TmallProductSchemaAddRequest Schema体系发布一个产品 */
type TmallProductSchemaAddRequest struct {
    
    /* brand_id special品牌ID */
    brand_id int64 `json:"brand_id";xml:"brand_id"`
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* xml_data required根据tmall.product.add.schema.get生成的产品发布规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallProductSchemaAddRequest) GetAPIName() string {
	return "tmall.product.schema.add"
}

/* TmallProductSchemaAddResponse Schema体系发布一个产品 */
type TmallProductSchemaAddResponse struct {
    
    /* add_product_result Basic新发产品结果 */
    add_product_result string `json:"add_product_result";xml:"add_product_result"`
    
}

/* TmallProductAddSchemaGetRequest 获取用户发布产品的规则 */
type TmallProductAddSchemaGetRequest struct {
    
    /* brand_id special品牌ID */
    brand_id int64 `json:"brand_id";xml:"brand_id"`
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TmallProductAddSchemaGetRequest) GetAPIName() string {
	return "tmall.product.add.schema.get"
}

/* TmallProductAddSchemaGetResponse 获取用户发布产品的规则 */
type TmallProductAddSchemaGetResponse struct {
    
    /* add_product_rule Basic返回发布产品的规则文档 */
    add_product_rule string `json:"add_product_rule";xml:"add_product_rule"`
    
}

/* TaobaoQimenReturnapplyReportRequest 退货异常包裹单通知接口 */
type TaobaoQimenReturnapplyReportRequest struct {
    
    /* request optional */
    request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenReturnapplyReportRequest) GetAPIName() string {
	return "taobao.qimen.returnapply.report"
}

/* TaobaoQimenReturnapplyReportResponse 退货异常包裹单通知接口 */
type TaobaoQimenReturnapplyReportResponse struct {
    
    /* response Object */
    response Response `json:"response";xml:"response"`
    
}

/* QimenTaobaoAlphaxOpenGetExpressRequest 调用接口 从isv获取  商家配置物流公司名单 */
type QimenTaobaoAlphaxOpenGetExpressRequest struct {
    
    /* seller_id required卖家id */
    seller_id string `json:"seller_id";xml:"seller_id"`
    
    /* seller_nick required卖家名称 */
    seller_nick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenGetExpressRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.get.express"
}

/* QimenTaobaoAlphaxOpenGetExpressResponse 调用接口 从isv获取  商家配置物流公司名单 */
type QimenTaobaoAlphaxOpenGetExpressResponse struct {
    
    /* result Object返回值 */
    result ResultDO `json:"result";xml:"result"`
    
}

/* TaobaoFenxiaoDealerRequisitionorderCloseRequest 供应商或分销商关闭采购申请/经销采购单 */
type TaobaoFenxiaoDealerRequisitionorderCloseRequest struct {
    
    /* dealer_order_id required经销采购单编号 */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
    /* reason required关闭原因：
1：长时间无法联系到分销商，取消交易。
2：分销商错误提交申请，取消交易。
3：缺货，近段时间都无法发货。
4：分销商恶意提交申请单。
5：其他原因。 */
    reason int64 `json:"reason";xml:"reason"`
    
    /* reason_detail required关闭详细原因，字数5-200字 */
    reason_detail string `json:"reason_detail";xml:"reason_detail"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}

/* TaobaoFenxiaoDealerRequisitionorderCloseResponse 供应商或分销商关闭采购申请/经销采购单 */
type TaobaoFenxiaoDealerRequisitionorderCloseResponse struct {
    
    /* is_success Basic操作是否成功。true：成功；false：失败。 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TmallItemAddSchemaGetRequest 通过类目以及productId获取商品发布规则； */
type TmallItemAddSchemaGetRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* isv_init optional正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可 */
    isv_init bool `json:"isv_init";xml:"isv_init"`
    
    /* product_id required商品发布的目标product_id */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* _type optional发布商品类型，一口价填“b”，拍卖填"a" */
    _type string `json:"type";xml:"type"`
    
}

func (req *TmallItemAddSchemaGetRequest) GetAPIName() string {
	return "tmall.item.add.schema.get"
}

/* TmallItemAddSchemaGetResponse 通过类目以及productId获取商品发布规则； */
type TmallItemAddSchemaGetResponse struct {
    
    /* add_item_result Basic返回发布商品的规则文档 */
    add_item_result string `json:"add_item_result";xml:"add_item_result"`
    
}

/* TaobaoFenxiaoDealerRequisitionorderAgreeRequest 供应商或分销商通过采购申请/经销采购单审核 */
type TaobaoFenxiaoDealerRequisitionorderAgreeRequest struct {
    
    /* dealer_order_id required采购申请/经销采购单编号 */
    dealer_order_id int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}

/* TaobaoFenxiaoDealerRequisitionorderAgreeResponse 供应商或分销商通过采购申请/经销采购单审核 */
type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {
    
    /* is_success Basic操作是否成功。true：成功；false：失败。 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TmallProductSchemaMatchRequest 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）； */
type TmallProductSchemaMatchRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* propvalues required根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。 */
    propvalues string `json:"propvalues";xml:"propvalues"`
    
}

func (req *TmallProductSchemaMatchRequest) GetAPIName() string {
	return "tmall.product.schema.match"
}

/* TmallProductSchemaMatchResponse 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）； */
type TmallProductSchemaMatchResponse struct {
    
    /* match_result Basic返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。 */
    match_result string `json:"match_result";xml:"match_result"`
    
}

/* TmallProductMatchSchemaGetRequest ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则 */
type TmallProductMatchSchemaGetRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TmallProductMatchSchemaGetRequest) GetAPIName() string {
	return "tmall.product.match.schema.get"
}

/* TmallProductMatchSchemaGetResponse ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则 */
type TmallProductMatchSchemaGetResponse struct {
    
    /* match_result Basic返回匹配product的规则文档 */
    match_result string `json:"match_result";xml:"match_result"`
    
}

/* TmallItemSchemaAddRequest 天猫TopSchema发布商品。 */
type TmallItemSchemaAddRequest struct {
    
    /* category_id required商品发布的目标类目，必须是叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* product_id required发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* xml_data required根据tmall.item.add.schema.get生成的商品发布规则入参数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaAddRequest) GetAPIName() string {
	return "tmall.item.schema.add"
}

/* TmallItemSchemaAddResponse 天猫TopSchema发布商品。 */
type TmallItemSchemaAddResponse struct {
    
    /* add_item_result Basic返回商品发布结果 */
    add_item_result string `json:"add_item_result";xml:"add_item_result"`
    
    /* gmt_create Basic发布商品操作成功时间 */
    gmt_create time.Time `json:"gmt_create";xml:"gmt_create"`
    
}

/* TaobaoFenxiaoDealerRequisitionorderGetRequest 批量查询采购申请/经销采购单，目前支持供应商和分销商查询 */
type TaobaoFenxiaoDealerRequisitionorderGetRequest struct {
    
    /* end_date required采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天 */
    end_date time.Time `json:"end_date";xml:"end_date"`
    
    /* fields optional多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
    fields string `json:"fields";xml:"fields"`
    
    /* identity required查询者自己在所要查询的采购申请/经销采购单中的身份。
1：供应商。
2：分销商。
注：填写其他值当做错误处理。 */
    identity int64 `json:"identity";xml:"identity"`
    
    /* order_status optional采购申请/经销采购单状态。
0：全部状态。
1：分销商提交申请，待供应商审核。
2：供应商驳回申请，待分销商确认。
3：供应商修改后，待分销商确认。
4：分销商拒绝修改，待供应商再审核。
5：审核通过下单成功，待分销商付款。
7：付款成功，待供应商发货。
8：供应商发货，待分销商收货。
9：分销商收货，交易成功。
10：采购申请/经销采购单关闭。

注：无值按默认值0计，超出状态范围返回错误信息。 */
    order_status int64 `json:"order_status";xml:"order_status"`
    
    /* page_no optional页码（大于0的整数。无值或小于1的值按默认值1计） */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_date required采购申请/经销采购单最早修改时间 */
    start_date time.Time `json:"start_date";xml:"start_date"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderGetRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.get"
}

/* TaobaoFenxiaoDealerRequisitionorderGetResponse 批量查询采购申请/经销采购单，目前支持供应商和分销商查询 */
type TaobaoFenxiaoDealerRequisitionorderGetResponse struct {
    
    /* dealer_orders Object Array采购申请/经销采购单结果列表 */
    dealer_orders DealerOrder `json:"dealer_orders";xml:"dealer_orders"`
    
    /* total_results Basic按查询条件查到的记录总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

/* TaobaoQimenOrderstatusUpdateRequest 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。 */
type TaobaoQimenOrderstatusUpdateRequest struct {
    
    /* action_time optional事件发生时间 */
    action_time string `json:"action_time";xml:"action_time"`
    
    /* allocation_code required星盘派单号 */
    allocation_code string `json:"allocation_code";xml:"allocation_code"`
    
    /* operator required操作人 */
    operator string `json:"operator";xml:"operator"`
    
    /* order_codes required淘系子订单号 */
    order_codes int64 `json:"order_codes";xml:"order_codes"`
    
    /* status required订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY */
    status string `json:"status";xml:"status"`
    
    /* store_id required目标门店的商户中心门店编码 */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* _type required业务类型，（枚举值：FAHUO、ZITI） */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoQimenOrderstatusUpdateRequest) GetAPIName() string {
	return "taobao.qimen.orderstatus.update"
}

/* TaobaoQimenOrderstatusUpdateResponse 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。 */
type TaobaoQimenOrderstatusUpdateResponse struct {
    
    /* is_success Basicsuccess */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* message Basicmessage */
    message string `json:"message";xml:"message"`
    
    /* result_code BasicresultCode */
    result_code string `json:"result_code";xml:"result_code"`
    
}

/* TaobaoFenxiaoDealerRequisitionorderQueryRequest 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。 */
type TaobaoFenxiaoDealerRequisitionorderQueryRequest struct {
    
    /* dealer_order_ids required经销采购单编号。
多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。 */
    dealer_order_ids int64 `json:"dealer_order_ids";xml:"dealer_order_ids"`
    
    /* fields optional多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.query"
}

/* TaobaoFenxiaoDealerRequisitionorderQueryResponse 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。 */
type TaobaoFenxiaoDealerRequisitionorderQueryResponse struct {
    
    /* dealer_orders Object Array经销采购单结果列表 */
    dealer_orders DealerOrder `json:"dealer_orders";xml:"dealer_orders"`
    
}

/* TaobaoQimenEventProduceRequest 当订单被处理时，用于通知奇门系统。 */
type TaobaoQimenEventProduceRequest struct {
    
    /* create optional订单创建时间,数字 */
    create int64 `json:"create";xml:"create"`
    
    /* ext optionalJSON格式扩展字段 */
    ext string `json:"ext";xml:"ext"`
    
    /* nick optional外部商家名称。必须同时填写platform */
    nick string `json:"nick";xml:"nick"`
    
    /* platform optional商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他 */
    platform string `json:"platform";xml:"platform"`
    
    /* status required事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK */
    status string `json:"status";xml:"status"`
    
    /* tid required淘宝订单号 */
    tid string `json:"tid";xml:"tid"`
    
}

func (req *TaobaoQimenEventProduceRequest) GetAPIName() string {
	return "taobao.qimen.event.produce"
}

/* TaobaoQimenEventProduceResponse 当订单被处理时，用于通知奇门系统。 */
type TaobaoQimenEventProduceResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}

/* TaobaoQimenEventsProduceRequest 批量发送消息 */
type TaobaoQimenEventsProduceRequest struct {
    
    /* messages required奇门事件列表, 最多50条 */
    messages QimenEvent `json:"messages";xml:"messages"`
    
}

func (req *TaobaoQimenEventsProduceRequest) GetAPIName() string {
	return "taobao.qimen.events.produce"
}

/* TaobaoQimenEventsProduceResponse 批量发送消息 */
type TaobaoQimenEventsProduceResponse struct {
    
    /* is_all_success Basic是否全部成功 */
    is_all_success bool `json:"is_all_success";xml:"is_all_success"`
    
    /* results Object Array发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功 */
    results QimenResult `json:"results";xml:"results"`
    
}

/* TaobaoFenxiaoTradePrepayOfflineReduceRequest 渠道分销供应商上传线下流水预存款（减少） */
type TaobaoFenxiaoTradePrepayOfflineReduceRequest struct {
    
    /* offline_reduce_prepay_param required减少流水 */
    offline_reduce_prepay_param TopOfflineReducePrepayDto `json:"offline_reduce_prepay_param";xml:"offline_reduce_prepay_param"`
    
}

func (req *TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetAPIName() string {
	return "taobao.fenxiao.trade.prepay.offline.reduce"
}

/* TaobaoFenxiaoTradePrepayOfflineReduceResponse 渠道分销供应商上传线下流水预存款（减少） */
type TaobaoFenxiaoTradePrepayOfflineReduceResponse struct {
    
    /* result Objectresult */
    result ResultTopDo `json:"result";xml:"result"`
    
}

/* TaobaoQimenTradeUserAddRequest 添加奇门订单链路用户 */
type TaobaoQimenTradeUserAddRequest struct {
    
    /* memo optional商家备注 */
    memo string `json:"memo";xml:"memo"`
    
}

func (req *TaobaoQimenTradeUserAddRequest) GetAPIName() string {
	return "taobao.qimen.trade.user.add"
}

/* TaobaoQimenTradeUserAddResponse 添加奇门订单链路用户 */
type TaobaoQimenTradeUserAddResponse struct {
    
    /* appkey Basicappkey */
    appkey string `json:"appkey";xml:"appkey"`
    
    /* gmt_create Basic创建时间 */
    gmt_create time.Time `json:"gmt_create";xml:"gmt_create"`
    
    /* memo Basic卖家备注 */
    memo string `json:"memo";xml:"memo"`
    
}

/* TaobaoFenxiaoTradePrepayOfflineAddRequest 渠道分销供应商上传线下流水预存款（增加） */
type TaobaoFenxiaoTradePrepayOfflineAddRequest struct {
    
    /* offline_add_prepay_param required增加流水 */
    offline_add_prepay_param TopOfflineAddPrepayDto `json:"offline_add_prepay_param";xml:"offline_add_prepay_param"`
    
}

func (req *TaobaoFenxiaoTradePrepayOfflineAddRequest) GetAPIName() string {
	return "taobao.fenxiao.trade.prepay.offline.add"
}

/* TaobaoFenxiaoTradePrepayOfflineAddResponse 渠道分销供应商上传线下流水预存款（增加） */
type TaobaoFenxiaoTradePrepayOfflineAddResponse struct {
    
    /* result Objectresult */
    result ResultTopDo `json:"result";xml:"result"`
    
}

/* TaobaoQimenTradeUserDeleteRequest 删除奇门订单链路用户 */
type TaobaoQimenTradeUserDeleteRequest struct {
    
}

func (req *TaobaoQimenTradeUserDeleteRequest) GetAPIName() string {
	return "taobao.qimen.trade.user.delete"
}

/* TaobaoQimenTradeUserDeleteResponse 删除奇门订单链路用户 */
type TaobaoQimenTradeUserDeleteResponse struct {
    
    /* modal Basicmodal */
    modal bool `json:"modal";xml:"modal"`
    
}

/* TaobaoQimenTradeUsersGetRequest 获取已开通奇门订单服务的用户列表 */
type TaobaoQimenTradeUsersGetRequest struct {
    
    /* page_index required每页的数量 */
    page_index int64 `json:"page_index";xml:"page_index"`
    
    /* page_size required页数 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoQimenTradeUsersGetRequest) GetAPIName() string {
	return "taobao.qimen.trade.users.get"
}

/* TaobaoQimenTradeUsersGetResponse 获取已开通奇门订单服务的用户列表 */
type TaobaoQimenTradeUsersGetResponse struct {
    
    /* total_count BasictotalCount */
    total_count int64 `json:"total_count";xml:"total_count"`
    
    /* users Object Arraymodal */
    users QimenUser `json:"users";xml:"users"`
    
}

/* TaobaoItemSchemaUpdateRequest 新模式下的商品编辑接口，在调用该接口的时候，需要先调用taobao.item.update.rules.get接口获取编辑规则和数据集。然后按照约定的参数传入规则，调用该接口进行编辑商品 */
type TaobaoItemSchemaUpdateRequest struct {
    
    /* category_id optional如果重新选择类目后，传入重新选择的叶子类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* incremental optional是否增量更新。为true只改xml_data里传入的值。为false时，将根据xml_data的数据对原始宝贝数据进行全量覆盖,这个参数暂时不支持 */
    incremental bool `json:"incremental";xml:"incremental"`
    
    /* item_id required编辑商品的商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* xml_data required编辑商品时候的修改内容 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TaobaoItemSchemaUpdateRequest) GetAPIName() string {
	return "taobao.item.schema.update"
}

/* TaobaoItemSchemaUpdateResponse 新模式下的商品编辑接口，在调用该接口的时候，需要先调用taobao.item.update.rules.get接口获取编辑规则和数据集。然后按照约定的参数传入规则，调用该接口进行编辑商品 */
type TaobaoItemSchemaUpdateResponse struct {
    
    /* update_result Basic编辑商品返回的结果信息，暂时只返回 itemId */
    update_result string `json:"update_result";xml:"update_result"`
    
}

/* TaobaoItemAddSchemaGetRequest 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能 */
type TaobaoItemAddSchemaGetRequest struct {
    
    /* category_id required发布宝贝的叶子类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TaobaoItemAddSchemaGetRequest) GetAPIName() string {
	return "taobao.item.add.schema.get"
}

/* TaobaoItemAddSchemaGetResponse 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能 */
type TaobaoItemAddSchemaGetResponse struct {
    
    /* add_rules Basic返回结果的集合 */
    add_rules string `json:"add_rules";xml:"add_rules"`
    
}

/* TaobaoItemSchemaAddRequest isv将宝贝信息，组装成特定格式的xml数据作为参数，进行发布商品 */
type TaobaoItemSchemaAddRequest struct {
    
    /* category_id required当前发布的叶子类目 */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* xml_data required将需要发布的商品数据组装成的xml格式数据 */
    xml_data string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TaobaoItemSchemaAddRequest) GetAPIName() string {
	return "taobao.item.schema.add"
}

/* TaobaoItemSchemaAddResponse isv将宝贝信息，组装成特定格式的xml数据作为参数，进行发布商品 */
type TaobaoItemSchemaAddResponse struct {
    
    /* add_result Basic返回的结果 */
    add_result string `json:"add_result";xml:"add_result"`
    
}

/* TaobaoItemUpdateSchemaGetRequest 在新的编辑模式下，isv需要先获取一份编辑商品的规则和数据，然后根据规则传入数据。该接口提供编辑规则查询功能 */
type TaobaoItemUpdateSchemaGetRequest struct {
    
    /* category_id optional商品类目id */
    category_id int64 `json:"category_id";xml:"category_id"`
    
    /* item_id required商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoItemUpdateSchemaGetRequest) GetAPIName() string {
	return "taobao.item.update.schema.get"
}

/* TaobaoItemUpdateSchemaGetResponse 在新的编辑模式下，isv需要先获取一份编辑商品的规则和数据，然后根据规则传入数据。该接口提供编辑规则查询功能 */
type TaobaoItemUpdateSchemaGetResponse struct {
    
    /* update_rules Basic返回的结果集 */
    update_rules string `json:"update_rules";xml:"update_rules"`
    
}

