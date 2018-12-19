package tbsdk

import "time"

type ErrorResponse struct{
    SubMsg string `json:"sub_msg",xml:"sub_msg"`
    Code int `json:"code",xml:"code"`
    SubCode string `json:"sub_code",xml:"sub_code"`
    Msg string `json:"msg",xml:"msg"`
    RequestId string `json:"request_id",xml:"request_id"`
}


// TaobaoRegionPriceManageRequest 编辑区域价格
type TaobaoRegionPriceManageRequest struct {
    
    // IsFull optionaltrue:全量, false:增量
    IsFull bool `json:"is_full";xml:"is_full"`
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // RegionalPriceDtos required列表
    RegionalPriceDtos RegionalPriceDto `json:"regional_price_dtos";xml:"regional_price_dtos"`
    
    // SkuId optional无sku传0
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceManageRequest) GetAPIName() string {
	return "taobao.region.price.manage"
}

func (req *TaobaoRegionPriceManageRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["IsFull"] = req.IsFull
    
    params["ItemId"] = req.ItemId
    
    params["RegionalPriceDtos"] = req.RegionalPriceDtos
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoRegionPriceManageResponse 编辑区域价格
type TaobaoRegionPriceManageResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basicsuccess
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoTmcQueueGetRequest 根据appkey和groupName获取消息队列积压情况
type TaobaoTmcQueueGetRequest struct {
    
    // GroupName requiredTMC组名
    GroupName string `json:"group_name";xml:"group_name"`
    
}

func (req *TaobaoTmcQueueGetRequest) GetAPIName() string {
	return "taobao.tmc.queue.get"
}

func (req *TaobaoTmcQueueGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["GroupName"] = req.GroupName
    
    return params
}

// TaobaoTmcQueueGetResponse 根据appkey和groupName获取消息队列积压情况
type TaobaoTmcQueueGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Datas Object Array队列详细信息
    Datas TmcQueueInfo `json:"datas";xml:"datas"`
    
}

// TaobaoWlbWmsReturnOrderNotifyRequest 销售退货通知
type TaobaoWlbWmsReturnOrderNotifyRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // ExtendFields optional扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // OrderCode optionalERP单据编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderCreateTime optionalERP订单创建时间
    OrderCreateTime time.Time `json:"order_create_time";xml:"order_create_time"`
    
    // OrderFlag optional用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    OrderFlag string `json:"order_flag";xml:"order_flag"`
    
    // OrderItemList optional商品信息列表
    OrderItemList Orderitemlistwlbwmsreturnordernotify `json:"order_item_list";xml:"order_item_list"`
    
    // OrderSource optional订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
    OrderSource string `json:"order_source";xml:"order_source"`
    
    // OrderType optional订单类型 501 销退入库
    OrderType string `json:"order_type";xml:"order_type"`
    
    // OwnerUserId optional货主ID
    OwnerUserId string `json:"owner_user_id";xml:"owner_user_id"`
    
    // PrevOrderCode optional来源单据号，销售退货时填充原发货的LBX号
    PrevOrderCode string `json:"prev_order_code";xml:"prev_order_code"`
    
    // ReceiverInfo optional收件人信息
    ReceiverInfo Receiverinfowlbwmsreturnordernotify `json:"receiver_info";xml:"receiver_info"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // ReturnReason optional销退时请提供退货的原因
    ReturnReason string `json:"return_reason";xml:"return_reason"`
    
    // SenderInfo optional发件人信息
    SenderInfo Senderinfowlbwmsreturnordernotify `json:"sender_info";xml:"sender_info"`
    
    // StoreCode optional仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TmsOrderCode optional运单号
    TmsOrderCode string `json:"tms_order_code";xml:"tms_order_code"`
    
    // TmsServiceCode optional快递公司编码
    TmsServiceCode string `json:"tms_service_code";xml:"tms_service_code"`
    
    // TmsServiceName optional快递公司名称
    TmsServiceName string `json:"tms_service_name";xml:"tms_service_name"`
    
}

func (req *TaobaoWlbWmsReturnOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.return.order.notify"
}

func (req *TaobaoWlbWmsReturnOrderNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 18)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["ExtendFields"] = req.ExtendFields
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderCreateTime"] = req.OrderCreateTime
    
    params["OrderFlag"] = req.OrderFlag
    
    params["OrderItemList"] = req.OrderItemList
    
    params["OrderSource"] = req.OrderSource
    
    params["OrderType"] = req.OrderType
    
    params["OwnerUserId"] = req.OwnerUserId
    
    params["PrevOrderCode"] = req.PrevOrderCode
    
    params["ReceiverInfo"] = req.ReceiverInfo
    
    params["Remark"] = req.Remark
    
    params["ReturnReason"] = req.ReturnReason
    
    params["SenderInfo"] = req.SenderInfo
    
    params["StoreCode"] = req.StoreCode
    
    params["TmsOrderCode"] = req.TmsOrderCode
    
    params["TmsServiceCode"] = req.TmsServiceCode
    
    params["TmsServiceName"] = req.TmsServiceName
    
    return params
}

// TaobaoWlbWmsReturnOrderNotifyResponse 销售退货通知
type TaobaoWlbWmsReturnOrderNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CreateTime Basic订单创建时间
    CreateTime time.Time `json:"create_time";xml:"create_time"`
    
    // StoreOrderCode BasicLBX929829111
    StoreOrderCode string `json:"store_order_code";xml:"store_order_code"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoRegionPriceQueryRequest 区域价格查询
type TaobaoRegionPriceQueryRequest struct {
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // RegionalPriceDtos optional不传则返回所有设置的区域价格
    RegionalPriceDtos RegionalPriceDto `json:"regional_price_dtos";xml:"regional_price_dtos"`
    
    // SkuId optional无sku可传0
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceQueryRequest) GetAPIName() string {
	return "taobao.region.price.query"
}

func (req *TaobaoRegionPriceQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ItemId"] = req.ItemId
    
    params["RegionalPriceDtos"] = req.RegionalPriceDtos
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoRegionPriceQueryResponse 区域价格查询
type TaobaoRegionPriceQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result BaseResult `json:"result";xml:"result"`
    
}

// TaobaoJdsTradeTracesGetRequest 获取聚石塔数据共享的交易全链路信息
type TaobaoJdsTradeTracesGetRequest struct {
    
    // ReturnUserStatus optional是否返回用户的状态(是否存在)
    ReturnUserStatus bool `json:"return_user_status";xml:"return_user_status"`
    
    // Tid required淘宝的订单tid
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoJdsTradeTracesGetRequest) GetAPIName() string {
	return "taobao.jds.trade.traces.get"
}

func (req *TaobaoJdsTradeTracesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ReturnUserStatus"] = req.ReturnUserStatus
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoJdsTradeTracesGetResponse 获取聚石塔数据共享的交易全链路信息
type TaobaoJdsTradeTracesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Traces Object Array跟踪信息列表
    Traces TradeTrace `json:"traces";xml:"traces"`
    
    // UserStatus Basic在订单全链路系统中的状态(比如是否存在)
    UserStatus string `json:"user_status";xml:"user_status"`
    
}

// TaobaoTmcUserTopicsGetRequest 获取用户开通的topic列表
type TaobaoTmcUserTopicsGetRequest struct {
    
    // Nick optional卖家nick
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoTmcUserTopicsGetRequest) GetAPIName() string {
	return "taobao.tmc.user.topics.get"
}

func (req *TaobaoTmcUserTopicsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoTmcUserTopicsGetResponse 获取用户开通的topic列表
type TaobaoTmcUserTopicsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ResultCode Basic错误码
    ResultCode string `json:"result_code";xml:"result_code"`
    
    // ResultMessage Basic错误信息
    ResultMessage string `json:"result_message";xml:"result_message"`
    
    // Topics Basic Arraytopic列表
    Topics string `json:"topics";xml:"topics"`
    
}

// TaobaoWlbWmsSkuGetRequest 商品信息查询
type TaobaoWlbWmsSkuGetRequest struct {
    
    // ItemCode optional菜鸟商品ID,与itemcode必须有一个值不为空
    ItemCode string `json:"item_code";xml:"item_code"`
    
    // ItemId optional商家商品编码,与itemid必须有一个值不为空
    ItemId string `json:"item_id";xml:"item_id"`
    
    // OwnerUserId optional货主ID
    OwnerUserId string `json:"owner_user_id";xml:"owner_user_id"`
    
}

func (req *TaobaoWlbWmsSkuGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.sku.get"
}

func (req *TaobaoWlbWmsSkuGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ItemCode"] = req.ItemCode
    
    params["ItemId"] = req.ItemId
    
    params["OwnerUserId"] = req.OwnerUserId
    
    return params
}

// TaobaoWlbWmsSkuGetResponse 商品信息查询
type TaobaoWlbWmsSkuGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AdventLifecycle Basic保质期预警天数
    AdventLifecycle int64 `json:"advent_lifecycle";xml:"advent_lifecycle"`
    
    // ApprovalNumber Basic批准文号
    ApprovalNumber string `json:"approval_number";xml:"approval_number"`
    
    // BarCode Basic条形码，多条码请用”;”分隔；
    BarCode string `json:"bar_code";xml:"bar_code"`
    
    // Brand Basic品牌编码
    Brand string `json:"brand";xml:"brand"`
    
    // BrandName Basic品牌名称
    BrandName string `json:"brand_name";xml:"brand_name"`
    
    // Category Basic商品类别编码（外部系统类别）
    Category string `json:"category";xml:"category"`
    
    // CategoryName Basic商品类别名称
    CategoryName string `json:"category_name";xml:"category_name"`
    
    // Color Basic颜色
    Color string `json:"color";xml:"color"`
    
    // CostPrice Basic成本价，单位分
    CostPrice int64 `json:"cost_price";xml:"cost_price"`
    
    // ExtendFields Basic拓展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // GrossWeight Basic毛重，单位克
    GrossWeight int64 `json:"gross_weight";xml:"gross_weight"`
    
    // Height Basic高度，单位毫米
    Height int64 `json:"height";xml:"height"`
    
    // IitemCode Basic商家商品编码,与itemid必须有一个值不为空
    IitemCode string `json:"iitem_code";xml:"iitem_code"`
    
    // IsAreaSale Basic是否区域销售
    IsAreaSale bool `json:"is_area_sale";xml:"is_area_sale"`
    
    // IsBatchMgt Basic是否启用批次管理
    IsBatchMgt bool `json:"is_batch_mgt";xml:"is_batch_mgt"`
    
    // IsDanger Basic是否危险品
    IsDanger bool `json:"is_danger";xml:"is_danger"`
    
    // IsHygroscopic Basic是否易碎品
    IsHygroscopic bool `json:"is_hygroscopic";xml:"is_hygroscopic"`
    
    // IsShelflife Basic是否启用保质期管理
    IsShelflife bool `json:"is_shelflife";xml:"is_shelflife"`
    
    // IsSnMgt Basic是否启用序列号管理
    IsSnMgt bool `json:"is_sn_mgt";xml:"is_sn_mgt"`
    
    // ItemId Basic菜鸟商品ID,与itemcode必须有一个值不为空
    ItemId string `json:"item_id";xml:"item_id"`
    
    // ItemPrice Basic零售价，单位分
    ItemPrice int64 `json:"item_price";xml:"item_price"`
    
    // Length Basic长度，单位毫米
    Length int64 `json:"length";xml:"length"`
    
    // Lifecycle Basic保质期天数
    Lifecycle int64 `json:"lifecycle";xml:"lifecycle"`
    
    // LockupLifecycle Basic保质期禁售天数
    LockupLifecycle int64 `json:"lockup_lifecycle";xml:"lockup_lifecycle"`
    
    // Name Basic商品名称
    Name string `json:"name";xml:"name"`
    
    // NetWeight Basic净重，单位克
    NetWeight int64 `json:"net_weight";xml:"net_weight"`
    
    // OriginAddress Basic场地
    OriginAddress int64 `json:"origin_address";xml:"origin_address"`
    
    // Pcs Basic箱规
    Pcs int64 `json:"pcs";xml:"pcs"`
    
    // RejectLifecycle Basic保质期禁收天数
    RejectLifecycle int64 `json:"reject_lifecycle";xml:"reject_lifecycle"`
    
    // Size Basic尺寸
    Size string `json:"size";xml:"size"`
    
    // Specification Basic规格
    Specification string `json:"specification";xml:"specification"`
    
    // TagPrice Basic吊牌价，单位分
    TagPrice int64 `json:"tag_price";xml:"tag_price"`
    
    // Title Basic商品标题
    Title string `json:"title";xml:"title"`
    
    // Type Basic商品类别 NORMAL：普通商品、 COMBINE：组合商品、 DISTRIBUTION：分销商品
    Type string `json:"type";xml:"type"`
    
    // UseYn Basic启用标识
    UseYn bool `json:"use_yn";xml:"use_yn"`
    
    // Volume Basic体积，单位立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Width Basic宽度，单位毫米
    Width int64 `json:"width";xml:"width"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoQimenInventoryReportRequest WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportRequest struct {
    
    // AdjustType required变动类型：CHECK=盘点 ADJUST=调整
    AdjustType string `json:"adjustType";xml:"adjustType"`
    
    // CheckOrderCode required盘点单编码
    CheckOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    // CheckOrderId optional仓储系统的盘点单编码
    CheckOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    // CheckTime optional盘点时间(YYYY-MM-DD HH:MM:SS)
    CheckTime string `json:"checkTime";xml:"checkTime"`
    
    // CurrentPage required当前页(从1开始)
    CurrentPage int64 `json:"currentPage";xml:"currentPage"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional商品库存信息列表
    Items Item `json:"items";xml:"items"`
    
    // OrderType optionalorderType
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OutBizCode required外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
    OutBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    // OwnerCode required货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // PageSize required每页记录的条数
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // TotalPage required总页数
    TotalPage int64 `json:"totalPage";xml:"totalPage"`
    
    // WarehouseCode required仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventoryReportRequest) GetAPIName() string {
	return "taobao.qimen.inventory.report"
}

func (req *TaobaoQimenInventoryReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 14)
    
    params["AdjustType"] = req.AdjustType
    
    params["CheckOrderCode"] = req.CheckOrderCode
    
    params["CheckOrderId"] = req.CheckOrderId
    
    params["CheckTime"] = req.CheckTime
    
    params["CurrentPage"] = req.CurrentPage
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OrderType"] = req.OrderType
    
    params["OutBizCode"] = req.OutBizCode
    
    params["OwnerCode"] = req.OwnerCode
    
    params["PageSize"] = req.PageSize
    
    params["Remark"] = req.Remark
    
    params["TotalPage"] = req.TotalPage
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenInventoryReportResponse WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenInventorycheckQueryRequest ERP调用奇门的接口,查询库存盘点情况
type TaobaoQimenInventorycheckQueryRequest struct {
    
    // CheckOrderCode required盘点单编码
    CheckOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    // CheckOrderId optional仓储系统的盘点单编码
    CheckOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页(从1开始)
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页orderLine条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventorycheckQueryRequest) GetAPIName() string {
	return "taobao.qimen.inventorycheck.query"
}

func (req *TaobaoQimenInventorycheckQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["CheckOrderCode"] = req.CheckOrderCode
    
    params["CheckOrderId"] = req.CheckOrderId
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenInventorycheckQueryResponse ERP调用奇门的接口,查询库存盘点情况
type TaobaoQimenInventorycheckQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CheckOrderCode Basic盘点单编码
    CheckOrderCode string `json:"checkOrderCode";xml:"checkOrderCode"`
    
    // CheckOrderId Basic仓储系统的盘点单编码
    CheckOrderId string `json:"checkOrderId";xml:"checkOrderId"`
    
    // CheckTime Basic盘点时间(YYYY-MM-DD HH:MM:SS)
    CheckTime string `json:"checkTime";xml:"checkTime"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Items Object Array商品库存列表
    Items Item `json:"items";xml:"items"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OwnerCode Basic货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Remark Basic备注
    Remark string `json:"remark";xml:"remark"`
    
    // TotalLines BasicorderLines总条数
    TotalLines int64 `json:"totalLines";xml:"totalLines"`
    
    // WarehouseCode Basic仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

// TaobaoQimenItemsSynchronizeRequest ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoQimenItemsSynchronizeRequest struct {
    
    // ActionType required操作类型(两种类型：add|update)
    ActionType string `json:"actionType";xml:"actionType"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional同步的商品信息
    Items Item `json:"items";xml:"items"`
    
    // OwnerCode required货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // WarehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemsSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.items.synchronize"
}

func (req *TaobaoQimenItemsSynchronizeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["ActionType"] = req.ActionType
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OwnerCode"] = req.OwnerCode
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenItemsSynchronizeResponse ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoQimenItemsSynchronizeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Items Object Array商品同步批量接口中单商品信息
    Items BatchItemSynItem `json:"items";xml:"items"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenWavenumReportRequest WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoQimenWavenumReportRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Orders optional发货单号
    Orders Order `json:"orders";xml:"orders"`
    
    // WaveNum required波次号
    WaveNum string `json:"waveNum";xml:"waveNum"`
    
}

func (req *TaobaoQimenWavenumReportRequest) GetAPIName() string {
	return "taobao.qimen.wavenum.report"
}

func (req *TaobaoQimenWavenumReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Orders"] = req.Orders
    
    params["WaveNum"] = req.WaveNum
    
    return params
}

// TaobaoQimenWavenumReportResponse WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoQimenWavenumReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoLogisticsOrderCreateRequest 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
type TaobaoLogisticsOrderCreateRequest struct {
    
    // GoodsNames required运送的货物名称列表，用|号隔开
    GoodsNames string `json:"goods_names";xml:"goods_names"`
    
    // GoodsQuantities required运送货物的数量列表，用|号隔开
    GoodsQuantities string `json:"goods_quantities";xml:"goods_quantities"`
    
    // IsConsign optional创建订单同时是否进行发货，默认发货。
    IsConsign bool `json:"is_consign";xml:"is_consign"`
    
    // ItemValues required运送货物的单价列表(注意：单位为分），用|号隔开
    ItemValues string `json:"item_values";xml:"item_values"`
    
    // LogisCompanyCode special发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
    LogisCompanyCode string `json:"logis_company_code";xml:"logis_company_code"`
    
    // LogisType optional发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。
    LogisType string `json:"logis_type";xml:"logis_type"`
    
    // MailNo special发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
    MailNo string `json:"mail_no";xml:"mail_no"`
    
    // OrderType optional物流发货类型：1-普通订单
    // 不填为默认类型 1-普通订单
    OrderType int64 `json:"order_type";xml:"order_type"`
    
    // SellerWangwangId required卖家旺旺号
    SellerWangwangId string `json:"seller_wangwang_id";xml:"seller_wangwang_id"`
    
    // Shipping optional运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
    Shipping int64 `json:"shipping";xml:"shipping"`
    
    // TradeId optional订单的交易号码
    TradeId int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoLogisticsOrderCreateRequest) GetAPIName() string {
	return "taobao.logistics.order.create"
}

func (req *TaobaoLogisticsOrderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 11)
    
    params["GoodsNames"] = req.GoodsNames
    
    params["GoodsQuantities"] = req.GoodsQuantities
    
    params["IsConsign"] = req.IsConsign
    
    params["ItemValues"] = req.ItemValues
    
    params["LogisCompanyCode"] = req.LogisCompanyCode
    
    params["LogisType"] = req.LogisType
    
    params["MailNo"] = req.MailNo
    
    params["OrderType"] = req.OrderType
    
    params["SellerWangwangId"] = req.SellerWangwangId
    
    params["Shipping"] = req.Shipping
    
    params["TradeId"] = req.TradeId
    
    return params
}

// TaobaoLogisticsOrderCreateResponse 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
type TaobaoLogisticsOrderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TradeId Basic淘宝物流订单交易号，如返回-1则表示错误。如果在新建订单时传入trade_id,此处会返回此id，如果未传入trade_id，此处会返回淘宝物流分配的交易号码。
    TradeId int64 `json:"trade_id";xml:"trade_id"`
    
}

// TaobaoItempropvaluesGetRequest 获取标准类目属性值
type TaobaoItempropvaluesGetRequest struct {
    
    // AttrKeys optional属性的Key，支持多条，以“,”分隔
    AttrKeys string `json:"attr_keys";xml:"attr_keys"`
    
    // Cid required叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
    Cid int64 `json:"cid";xml:"cid"`
    
    // Datetime special假如传2005-01-01 00:00:00，则取所有的属性和子属性(状态为删除的属性值不返回prop_name)
    Datetime time.Time `json:"datetime";xml:"datetime"`
    
    // Fields required需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
    Fields string `json:"fields";xml:"fields"`
    
    // Pvs special属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
    Pvs string `json:"pvs";xml:"pvs"`
    
    // Type optional获取类目的类型：1代表集市、2代表天猫
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItempropvaluesGetRequest) GetAPIName() string {
	return "taobao.itempropvalues.get"
}

func (req *TaobaoItempropvaluesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["AttrKeys"] = req.AttrKeys
    
    params["Cid"] = req.Cid
    
    params["Datetime"] = req.Datetime
    
    params["Fields"] = req.Fields
    
    params["Pvs"] = req.Pvs
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoItempropvaluesGetResponse 获取标准类目属性值
type TaobaoItempropvaluesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // LastModified Basic最近修改时间。格式:yyyy-MM-dd HH:mm:ss
    LastModified string `json:"last_modified";xml:"last_modified"`
    
    // PropValues Object Array属性值,根据fields传入的参数返回相应的结果
    PropValues PropValue `json:"prop_values";xml:"prop_values"`
    
}

// Gfq1qp4287ShopexAreaInfoGetRequest 地区编码同步
type Gfq1qp4287ShopexAreaInfoGetRequest struct {
    
    // QimenBody optional请求参数
    QimenBody string `json:"qimen_body";xml:"qimen_body"`
    
}

func (req *Gfq1qp4287ShopexAreaInfoGetRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.area.info.get"
}

func (req *Gfq1qp4287ShopexAreaInfoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["QimenBody"] = req.QimenBody
    
    return params
}

// Gfq1qp4287ShopexAreaInfoGetResponse 地区编码同步
type Gfq1qp4287ShopexAreaInfoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Data Basic返回内容
    Data string `json:"data";xml:"data"`
    
    // MsgId Basic请求任务id
    MsgId string `json:"msg_id";xml:"msg_id"`
    
    // Res Basic描述
    Res string `json:"res";xml:"res"`
    
    // Rsp Basic成功失败标识
    Rsp string `json:"rsp";xml:"rsp"`
    
}

// TaobaoTmcMessagesConfirmRequest 确认消费消息的状态
type TaobaoTmcMessagesConfirmRequest struct {
    
    // FMessageIds optional处理失败的消息ID列表--已废弃，无需传此字段
    FMessageIds int64 `json:"f_message_ids";xml:"f_message_ids"`
    
    // GroupName optional分组名称，不传代表默认分组
    GroupName string `json:"group_name";xml:"group_name"`
    
    // SMessageIds required处理成功的消息ID列表 最大 200个ID
    SMessageIds int64 `json:"s_message_ids";xml:"s_message_ids"`
    
}

func (req *TaobaoTmcMessagesConfirmRequest) GetAPIName() string {
	return "taobao.tmc.messages.confirm"
}

func (req *TaobaoTmcMessagesConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["FMessageIds"] = req.FMessageIds
    
    params["GroupName"] = req.GroupName
    
    params["SMessageIds"] = req.SMessageIds
    
    return params
}

// TaobaoTmcMessagesConfirmResponse 确认消费消息的状态
type TaobaoTmcMessagesConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoTmcMessagesConsumeRequest 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaoTmcMessagesConsumeRequest struct {
    
    // GroupName optional用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
    GroupName string `json:"group_name";xml:"group_name"`
    
    // Quantity optional每次批量消费消息的条数
    Quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoTmcMessagesConsumeRequest) GetAPIName() string {
	return "taobao.tmc.messages.consume"
}

func (req *TaobaoTmcMessagesConsumeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["GroupName"] = req.GroupName
    
    params["Quantity"] = req.Quantity
    
    return params
}

// TaobaoTmcMessagesConsumeResponse 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaoTmcMessagesConsumeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Messages Object Array消息列表
    Messages TmcMessage `json:"messages";xml:"messages"`
    
}

// TaobaoTmcMessageProduceRequest 发布单条消息
type TaobaoTmcMessageProduceRequest struct {
    
    // Content required消息内容的JSON表述，必须按照topic的定义来填充
    Content string `json:"content";xml:"content"`
    
    // ExContent optional消息的扩增属性，用json格式表示
    ExContent string `json:"ex_content";xml:"ex_content"`
    
    // MediaContent optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。
    MediaContent []byte `json:"media_content";xml:"media_content"`
    
    // MediaContent2 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    MediaContent2 []byte `json:"media_content2";xml:"media_content2"`
    
    // MediaContent3 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    MediaContent3 []byte `json:"media_content3";xml:"media_content3"`
    
    // MediaContent4 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    MediaContent4 []byte `json:"media_content4";xml:"media_content4"`
    
    // MediaContent5 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    MediaContent5 []byte `json:"media_content5";xml:"media_content5"`
    
    // TargetAppkey optional直发消息需要传入目标appkey
    TargetAppkey string `json:"target_appkey";xml:"target_appkey"`
    
    // TargetGroup optional目标分组，一般为default
    TargetGroup string `json:"target_group";xml:"target_group"`
    
    // Topic required消息类型
    Topic string `json:"topic";xml:"topic"`
    
}

func (req *TaobaoTmcMessageProduceRequest) GetAPIName() string {
	return "taobao.tmc.message.produce"
}

func (req *TaobaoTmcMessageProduceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["Content"] = req.Content
    
    params["ExContent"] = req.ExContent
    
    params["MediaContent"] = req.MediaContent
    
    params["MediaContent2"] = req.MediaContent2
    
    params["MediaContent3"] = req.MediaContent3
    
    params["MediaContent4"] = req.MediaContent4
    
    params["MediaContent5"] = req.MediaContent5
    
    params["TargetAppkey"] = req.TargetAppkey
    
    params["TargetGroup"] = req.TargetGroup
    
    params["Topic"] = req.Topic
    
    return params
}

// TaobaoTmcMessageProduceResponse 发布单条消息
type TaobaoTmcMessageProduceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // MsgIds Basic Array消息ID
    MsgIds string `json:"msg_ids";xml:"msg_ids"`
    
    // Total Basic投递目标数
    Total int64 `json:"total";xml:"total"`
    
}

// TaobaoItemsOnsaleGetRequest 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤 
//只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取
type TaobaoItemsOnsaleGetRequest struct {
    
    // AuctionType optional商品类型：a-拍卖,b-一口价
    AuctionType string `json:"auction_type";xml:"auction_type"`
    
    // Cid optional商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
    Cid int64 `json:"cid";xml:"cid"`
    
    // EndModified optional结束的修改时间
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // Fields required需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。
    Fields string `json:"fields";xml:"fields"`
    
    // HasDiscount optional是否参与会员折扣。可选值：true，false。默认不过滤该条件
    HasDiscount bool `json:"has_discount";xml:"has_discount"`
    
    // HasShowcase optional是否橱窗推荐。 可选值：true，false。默认不过滤该条件
    HasShowcase bool `json:"has_showcase";xml:"has_showcase"`
    
    // IsCombine optional组合商品
    IsCombine bool `json:"is_combine";xml:"is_combine"`
    
    // IsCspu optional是否挂接了达尔文标准产品体系
    IsCspu bool `json:"is_cspu";xml:"is_cspu"`
    
    // IsEx optional商品是否在外部网店显示
    IsEx bool `json:"is_ex";xml:"is_ex"`
    
    // IsTaobao optional商品是否在淘宝显示
    IsTaobao bool `json:"is_taobao";xml:"is_taobao"`
    
    // OrderBy optional排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
    OrderBy string `json:"order_by";xml:"order_by"`
    
    // PageNo optional页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // Q optional搜索字段。搜索商品的title。
    Q string `json:"q";xml:"q"`
    
    // SellerCids optional卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
    SellerCids string `json:"seller_cids";xml:"seller_cids"`
    
    // StartModified optional起始的修改时间
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoItemsOnsaleGetRequest) GetAPIName() string {
	return "taobao.items.onsale.get"
}

func (req *TaobaoItemsOnsaleGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 16)
    
    params["AuctionType"] = req.AuctionType
    
    params["Cid"] = req.Cid
    
    params["EndModified"] = req.EndModified
    
    params["Fields"] = req.Fields
    
    params["HasDiscount"] = req.HasDiscount
    
    params["HasShowcase"] = req.HasShowcase
    
    params["IsCombine"] = req.IsCombine
    
    params["IsCspu"] = req.IsCspu
    
    params["IsEx"] = req.IsEx
    
    params["IsTaobao"] = req.IsTaobao
    
    params["OrderBy"] = req.OrderBy
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["Q"] = req.Q
    
    params["SellerCids"] = req.SellerCids
    
    params["StartModified"] = req.StartModified
    
    return params
}

// TaobaoItemsOnsaleGetResponse 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤 
//只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取
type TaobaoItemsOnsaleGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Items Object Array搜索到的商品列表，具体字段根据设定的fields决定，不包括desc字段
    Items Item `json:"items";xml:"items"`
    
    // TotalResults Basic搜索到符合条件的结果总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// CainiaoCrmOmsRuleSyncRequest 将商家ERP订单处理规则同步到菜鸟CRM系统
type CainiaoCrmOmsRuleSyncRequest struct {
    
    // CheckRuleMsg optional人工审单规则描述
    CheckRuleMsg string `json:"check_rule_msg";xml:"check_rule_msg"`
    
    // IsAutoCheck required是否系统智能处理订单（无人工介入）
    IsAutoCheck bool `json:"is_auto_check";xml:"is_auto_check"`
    
    // IsOpenCnauto required是否开启菜鸟自动流转规则
    IsOpenCnauto bool `json:"is_open_cnauto";xml:"is_open_cnauto"`
    
    // IsSysMergeOrder optional是否开启了订单合单
    IsSysMergeOrder bool `json:"is_sys_merge_order";xml:"is_sys_merge_order"`
    
    // MergeOrderCycle optional订单合单周期（单位：分钟）
    MergeOrderCycle int64 `json:"merge_order_cycle";xml:"merge_order_cycle"`
    
    // OtherRule optional其他未定义订单处理规则，格式｛name;stauts;cycle;｝
    OtherRule string `json:"other_rule";xml:"other_rule"`
    
    // ShopCode required店铺nick
    ShopCode string `json:"shop_code";xml:"shop_code"`
    
}

func (req *CainiaoCrmOmsRuleSyncRequest) GetAPIName() string {
	return "cainiao.crm.oms.rule.sync"
}

func (req *CainiaoCrmOmsRuleSyncRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["CheckRuleMsg"] = req.CheckRuleMsg
    
    params["IsAutoCheck"] = req.IsAutoCheck
    
    params["IsOpenCnauto"] = req.IsOpenCnauto
    
    params["IsSysMergeOrder"] = req.IsSysMergeOrder
    
    params["MergeOrderCycle"] = req.MergeOrderCycle
    
    params["OtherRule"] = req.OtherRule
    
    params["ShopCode"] = req.ShopCode
    
    return params
}

// CainiaoCrmOmsRuleSyncResponse 将商家ERP订单处理规则同步到菜鸟CRM系统
type CainiaoCrmOmsRuleSyncResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlErrorCode BasicerrorCode
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg BasicerrorMsg
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basicsuccess
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoItemUpdateRequest 根据传入的num_iid更新对应的商品的数据 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败） 
//商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
//当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。
type TaobaoItemUpdateRequest struct {
    
    // AfterSaleId optional售后说明模板id
    AfterSaleId int64 `json:"after_sale_id";xml:"after_sale_id"`
    
    // ApproveStatus optional商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true
    ApproveStatus string `json:"approve_status";xml:"approve_status"`
    
    // AuctionPoint optional商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
    AuctionPoint int64 `json:"auction_point";xml:"auction_point"`
    
    // AutoFill optional代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
    AutoFill string `json:"auto_fill";xml:"auto_fill"`
    
    // AutoRepost optional自动重发。可选值:true,false;
    AutoRepost bool `json:"auto_repost";xml:"auto_repost"`
    
    // Barcode optional商品条形码
    Barcode string `json:"barcode";xml:"barcode"`
    
    // ChangeProp optional基础色数据
    ChangeProp string `json:"change_prop";xml:"change_prop"`
    
    // ChaoshiExtendsInfo optional天猫超市扩展字段，天猫超市专用。
    ChaoshiExtendsInfo string `json:"chaoshi_extends_info";xml:"chaoshi_extends_info"`
    
    // Cid optional叶子类目id
    Cid int64 `json:"cid";xml:"cid"`
    
    // CodPostageId optional货到付款运费模板ID该字段已经废弃，货到付款模板已经集成到运费模板中。
    CodPostageId int64 `json:"cod_postage_id";xml:"cod_postage_id"`
    
    // CpvMemo optional针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
    CpvMemo string `json:"cpv_memo";xml:"cpv_memo"`
    
    // DeliveryTimeDeliveryTime optional商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
    DeliveryTimeDeliveryTime string `json:"delivery_time.delivery_time";xml:"delivery_time.delivery_time"`
    
    // DeliveryTimeDeliveryTimeType optional发货时间类型：绝对发货时间或者相对发货时间
    DeliveryTimeDeliveryTimeType string `json:"delivery_time.delivery_time_type";xml:"delivery_time.delivery_time_type"`
    
    // DeliveryTimeNeedDeliveryTime optional设置是否使用发货时间，商品级别，sku级别
    DeliveryTimeNeedDeliveryTime string `json:"delivery_time.need_delivery_time";xml:"delivery_time.need_delivery_time"`
    
    // Desc optional商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制
    Desc string `json:"desc";xml:"desc"`
    
    // DescModules optional商品描述模块化，模块列表，由List转化成jsonArray存入，后端逻辑验证通过，拼装成模块内容+锚点导航后存入desc中。数据结构具体参见Item_Desc_Module
    DescModules string `json:"desc_modules";xml:"desc_modules"`
    
    // EmptyFields optional支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock）
    EmptyFields string `json:"empty_fields";xml:"empty_fields"`
    
    // EmsFee optionalems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
    EmsFee float32 `json:"ems_fee";xml:"ems_fee"`
    
    // ExpressFee optional快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
    ExpressFee float32 `json:"express_fee";xml:"express_fee"`
    
    // Features optional宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp
    Features string `json:"features";xml:"features"`
    
    // FoodSecurityContact optional厂家联系方式
    FoodSecurityContact string `json:"food_security.contact";xml:"food_security.contact"`
    
    // FoodSecurityDesignCode optional产品标准号
    FoodSecurityDesignCode string `json:"food_security.design_code";xml:"food_security.design_code"`
    
    // FoodSecurityFactory optional厂名
    FoodSecurityFactory string `json:"food_security.factory";xml:"food_security.factory"`
    
    // FoodSecurityFactorySite optional厂址
    FoodSecurityFactorySite string `json:"food_security.factory_site";xml:"food_security.factory_site"`
    
    // FoodSecurityFoodAdditive optional食品添加剂
    FoodSecurityFoodAdditive string `json:"food_security.food_additive";xml:"food_security.food_additive"`
    
    // FoodSecurityHealthProductNo optional健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
    FoodSecurityHealthProductNo string `json:"food_security.health_product_no";xml:"food_security.health_product_no"`
    
    // FoodSecurityMix optional配料表
    FoodSecurityMix string `json:"food_security.mix";xml:"food_security.mix"`
    
    // FoodSecurityPeriod optional保质期
    FoodSecurityPeriod string `json:"food_security.period";xml:"food_security.period"`
    
    // FoodSecurityPlanStorage optional储藏方法
    FoodSecurityPlanStorage string `json:"food_security.plan_storage";xml:"food_security.plan_storage"`
    
    // FoodSecurityPrdLicenseNo optional生产许可证号
    FoodSecurityPrdLicenseNo string `json:"food_security.prd_license_no";xml:"food_security.prd_license_no"`
    
    // FoodSecurityProductDateEnd optional生产结束日期,格式必须为yyyy-MM-dd
    FoodSecurityProductDateEnd string `json:"food_security.product_date_end";xml:"food_security.product_date_end"`
    
    // FoodSecurityProductDateStart optional生产开始日期，格式必须为yyyy-MM-dd
    FoodSecurityProductDateStart string `json:"food_security.product_date_start";xml:"food_security.product_date_start"`
    
    // FoodSecurityStockDateEnd optional进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
    FoodSecurityStockDateEnd string `json:"food_security.stock_date_end";xml:"food_security.stock_date_end"`
    
    // FoodSecurityStockDateStart optional进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
    FoodSecurityStockDateStart string `json:"food_security.stock_date_start";xml:"food_security.stock_date_start"`
    
    // FoodSecuritySupplier optional供货商
    FoodSecuritySupplier string `json:"food_security.supplier";xml:"food_security.supplier"`
    
    // FreightPayer optional运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);
    FreightPayer string `json:"freight_payer";xml:"freight_payer"`
    
    // GlobalStockCountry optional全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）
    GlobalStockCountry string `json:"global_stock_country";xml:"global_stock_country"`
    
    // GlobalStockDeliveryPlace optional全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”
    GlobalStockDeliveryPlace string `json:"global_stock_delivery_place";xml:"global_stock_delivery_place"`
    
    // GlobalStockTaxFreePromise optional全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。
    GlobalStockTaxFreePromise bool `json:"global_stock_tax_free_promise";xml:"global_stock_tax_free_promise"`
    
    // GlobalStockType optional全球购商品采购地（库存类型）全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
    GlobalStockType string `json:"global_stock_type";xml:"global_stock_type"`
    
    // HasDiscount optional支持会员打折。可选值:true,false;
    HasDiscount bool `json:"has_discount";xml:"has_discount"`
    
    // HasInvoice optional是否有发票。可选值:true,false (商城卖家此字段必须为true)
    HasInvoice bool `json:"has_invoice";xml:"has_invoice"`
    
    // HasShowcase optional橱窗推荐。可选值:true,false;
    HasShowcase bool `json:"has_showcase";xml:"has_showcase"`
    
    // HasWarranty optional是否有保修。可选值:true,false;
    HasWarranty bool `json:"has_warranty";xml:"has_warranty"`
    
    // Ignorewarning optional忽略警告提示.
    Ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    // Image optional商品图片。类型:JPG,GIF;最大长度:3M
    Image []byte `json:"image";xml:"image"`
    
    // Increment optional加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
    Increment float32 `json:"increment";xml:"increment"`
    
    // InputCustomCpv optional针对当前商品的自定义属性值
    InputCustomCpv string `json:"input_custom_cpv";xml:"input_custom_cpv"`
    
    // InputPids optional用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
    InputPids string `json:"input_pids";xml:"input_pids"`
    
    // InputStr optional用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
    InputStr string `json:"input_str";xml:"input_str"`
    
    // InteractiveId optional主图视频互动信息id，必须有主图视频id才能传互动信息id
    InteractiveId int64 `json:"interactive_id";xml:"interactive_id"`
    
    // Is3D optional是否是3D
    Is3D bool `json:"is_3D";xml:"is_3D"`
    
    // IsEx optional是否在外店显示
    IsEx bool `json:"is_ex";xml:"is_ex"`
    
    // IsLightningConsignment optional实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记
    IsLightningConsignment bool `json:"is_lightning_consignment";xml:"is_lightning_consignment"`
    
    // IsOffline optional是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
    IsOffline string `json:"is_offline";xml:"is_offline"`
    
    // IsReplaceSku optional是否替换sku
    IsReplaceSku bool `json:"is_replace_sku";xml:"is_replace_sku"`
    
    // IsTaobao optional是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
    IsTaobao bool `json:"is_taobao";xml:"is_taobao"`
    
    // IsXinpin optional商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。
    IsXinpin bool `json:"is_xinpin";xml:"is_xinpin"`
    
    // ItemSize optional表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0
    ItemSize string `json:"item_size";xml:"item_size"`
    
    // ItemWeight optional商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0
    ItemWeight string `json:"item_weight";xml:"item_weight"`
    
    // Lang optional商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN”
    Lang string `json:"lang";xml:"lang"`
    
    // LeaseExtendsInfo optional淘宝租赁扩展信息
    LeaseExtendsInfo string `json:"lease_extends_info";xml:"lease_extends_info"`
    
    // ListTime optional上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。
    ListTime time.Time `json:"list_time";xml:"list_time"`
    
    // LocalityLifeChooseLogis optional编辑电子凭证宝贝时候表示是否使用邮寄0: 代表不使用邮寄；1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
    LocalityLifeChooseLogis string `json:"locality_life.choose_logis";xml:"locality_life.choose_logis"`
    
    // LocalityLifeEticket optional电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
    LocalityLifeEticket string `json:"locality_life.eticket";xml:"locality_life.eticket"`
    
    // LocalityLifeExpirydate optional本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
    LocalityLifeExpirydate string `json:"locality_life.expirydate";xml:"locality_life.expirydate"`
    
    // LocalityLifeMerchant optional码商信息，格式为 码商id:nick
    LocalityLifeMerchant string `json:"locality_life.merchant";xml:"locality_life.merchant"`
    
    // LocalityLifeNetworkId optional网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID
    LocalityLifeNetworkId string `json:"locality_life.network_id";xml:"locality_life.network_id"`
    
    // LocalityLifeObs optional预约门店是否支持门店自提,1:是
    LocalityLifeObs string `json:"locality_life.obs";xml:"locality_life.obs"`
    
    // LocalityLifeOnsaleAutoRefundRatio optional电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
    LocalityLifeOnsaleAutoRefundRatio int64 `json:"locality_life.onsale_auto_refund_ratio";xml:"locality_life.onsale_auto_refund_ratio"`
    
    // LocalityLifePackageid optional新版电子凭证包id
    LocalityLifePackageid string `json:"locality_life.packageid";xml:"locality_life.packageid"`
    
    // LocalityLifeRefundRatio optional退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例
    LocalityLifeRefundRatio int64 `json:"locality_life.refund_ratio";xml:"locality_life.refund_ratio"`
    
    // LocalityLifeRefundmafee optional退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
    LocalityLifeRefundmafee string `json:"locality_life.refundmafee";xml:"locality_life.refundmafee"`
    
    // LocalityLifeVerification optional核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款
    LocalityLifeVerification string `json:"locality_life.verification";xml:"locality_life.verification"`
    
    // LocalityLifeVersion optional电子凭证版本 新版电子凭证值:1
    LocalityLifeVersion string `json:"locality_life.version";xml:"locality_life.version"`
    
    // LocationCity optional所在地城市。如杭州
    LocationCity string `json:"location.city";xml:"location.city"`
    
    // LocationState optional所在地省份。如浙江
    LocationState string `json:"location.state";xml:"location.state"`
    
    // MsPaymentPrice optional订金
    MsPaymentPrice string `json:"ms_payment.price";xml:"ms_payment.price"`
    
    // MsPaymentReferencePrice optional参考价
    MsPaymentReferencePrice string `json:"ms_payment.reference_price";xml:"ms_payment.reference_price"`
    
    // MsPaymentVoucherPrice optional尾款可抵扣金额
    MsPaymentVoucherPrice string `json:"ms_payment.voucher_price";xml:"ms_payment.voucher_price"`
    
    // Newprepay optional该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；<br>注意：使用该API修改商品其它属性如标题title时，如需保持商品不支持7天无理由退货状态，该字段需传入0 。
    Newprepay string `json:"newprepay";xml:"newprepay"`
    
    // Num optional商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
    Num int64 `json:"num";xml:"num"`
    
    // NumIid required商品数字ID，该参数必须
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // O2oBindService optional汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
    O2oBindService bool `json:"o2o_bind_service";xml:"o2o_bind_service"`
    
    // OuterId optional商家编码
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // PaimaiInfoDeposit optional拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。
    PaimaiInfoDeposit int64 `json:"paimai_info.deposit";xml:"paimai_info.deposit"`
    
    // PaimaiInfoInterval optional降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
    PaimaiInfoInterval int64 `json:"paimai_info.interval";xml:"paimai_info.interval"`
    
    // PaimaiInfoMode optional拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
    PaimaiInfoMode int64 `json:"paimai_info.mode";xml:"paimai_info.mode"`
    
    // PaimaiInfoReserve optional降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
    PaimaiInfoReserve float32 `json:"paimai_info.reserve";xml:"paimai_info.reserve"`
    
    // PaimaiInfoValidHour optional自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    PaimaiInfoValidHour int64 `json:"paimai_info.valid_hour";xml:"paimai_info.valid_hour"`
    
    // PaimaiInfoValidMinute optional自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    PaimaiInfoValidMinute int64 `json:"paimai_info.valid_minute";xml:"paimai_info.valid_minute"`
    
    // PicPath optional商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // PostFee optional平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写
    PostFee float32 `json:"post_fee";xml:"post_fee"`
    
    // PostageId optional宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.templates.get获得当前会话用户的所有邮费模板）
    PostageId int64 `json:"postage_id";xml:"postage_id"`
    
    // Price optional商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。
    Price float32 `json:"price";xml:"price"`
    
    // ProductId optional商品所属的产品ID(B商家发布商品需要用)
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // PropertyAlias optional属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。
    PropertyAlias string `json:"property_alias";xml:"property_alias"`
    
    // Props optional商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。
    Props string `json:"props";xml:"props"`
    
    // Qualification optional商品资质信息
    Qualification string `json:"qualification";xml:"qualification"`
    
    // ScenicTicketBookCost optional景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
    ScenicTicketBookCost string `json:"scenic_ticket_book_cost";xml:"scenic_ticket_book_cost"`
    
    // ScenicTicketPayWay optional景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
    ScenicTicketPayWay int64 `json:"scenic_ticket_pay_way";xml:"scenic_ticket_pay_way"`
    
    // SellPoint optional商品卖点信息，最长150个字符。天猫和集市都可用
    SellPoint string `json:"sell_point";xml:"sell_point"`
    
    // SellPromise optional是否承诺退换货服务!虚拟商品无须设置此项!
    SellPromise bool `json:"sell_promise";xml:"sell_promise"`
    
    // SellerCids optional重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
    SellerCids string `json:"seller_cids";xml:"seller_cids"`
    
    // Shape optional宝贝形态:1代表电子券;0或其他代表实物
    Shape string `json:"shape";xml:"shape"`
    
    // SkuBarcode optionalsku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
    SkuBarcode string `json:"sku_barcode";xml:"sku_barcode"`
    
    // SkuDeliveryTimes optional此参数暂时不起作用
    SkuDeliveryTimes string `json:"sku_delivery_times";xml:"sku_delivery_times"`
    
    // SkuHdHeight optional家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。 可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50
    SkuHdHeight string `json:"sku_hd_height";xml:"sku_hd_height"`
    
    // SkuHdLampQuantity optional家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7
    SkuHdLampQuantity string `json:"sku_hd_lamp_quantity";xml:"sku_hd_lamp_quantity"`
    
    // SkuHdLength optional家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：20,30,30
    SkuHdLength string `json:"sku_hd_length";xml:"sku_hd_length"`
    
    // SkuOuterIds optionalSku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
    SkuOuterIds string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    // SkuPrices optional更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
    SkuPrices string `json:"sku_prices";xml:"sku_prices"`
    
    // SkuProperties optional更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
    SkuProperties string `json:"sku_properties";xml:"sku_properties"`
    
    // SkuQuantities optional更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4
    SkuQuantities string `json:"sku_quantities";xml:"sku_quantities"`
    
    // SkuSpecIds optional此参数暂时不起作用
    SkuSpecIds string `json:"sku_spec_ids";xml:"sku_spec_ids"`
    
    // SpuConfirm optional手机类目spu 确认信息字段
    SpuConfirm bool `json:"spu_confirm";xml:"spu_confirm"`
    
    // StuffStatus optional商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。
    StuffStatus string `json:"stuff_status";xml:"stuff_status"`
    
    // SubStock optional商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存
    SubStock int64 `json:"sub_stock";xml:"sub_stock"`
    
    // Title optional宝贝标题. 不能超过30字符,受违禁词控制
    Title string `json:"title";xml:"title"`
    
    // ValidThru optional有效期。可选值:7,14;单位:天;
    ValidThru int64 `json:"valid_thru";xml:"valid_thru"`
    
    // VideoId optional主图视频id
    VideoId int64 `json:"video_id";xml:"video_id"`
    
    // Weight optional商品的重量(商超卖家专用字段)
    Weight int64 `json:"weight";xml:"weight"`
    
    // WirelessDesc optional无线的宝贝描述
    WirelessDesc string `json:"wireless_desc";xml:"wireless_desc"`
    
}

func (req *TaobaoItemUpdateRequest) GetAPIName() string {
	return "taobao.item.update"
}

func (req *TaobaoItemUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 123)
    
    params["AfterSaleId"] = req.AfterSaleId
    
    params["ApproveStatus"] = req.ApproveStatus
    
    params["AuctionPoint"] = req.AuctionPoint
    
    params["AutoFill"] = req.AutoFill
    
    params["AutoRepost"] = req.AutoRepost
    
    params["Barcode"] = req.Barcode
    
    params["ChangeProp"] = req.ChangeProp
    
    params["ChaoshiExtendsInfo"] = req.ChaoshiExtendsInfo
    
    params["Cid"] = req.Cid
    
    params["CodPostageId"] = req.CodPostageId
    
    params["CpvMemo"] = req.CpvMemo
    
    params["DeliveryTimeDeliveryTime"] = req.DeliveryTimeDeliveryTime
    
    params["DeliveryTimeDeliveryTimeType"] = req.DeliveryTimeDeliveryTimeType
    
    params["DeliveryTimeNeedDeliveryTime"] = req.DeliveryTimeNeedDeliveryTime
    
    params["Desc"] = req.Desc
    
    params["DescModules"] = req.DescModules
    
    params["EmptyFields"] = req.EmptyFields
    
    params["EmsFee"] = req.EmsFee
    
    params["ExpressFee"] = req.ExpressFee
    
    params["Features"] = req.Features
    
    params["FoodSecurityContact"] = req.FoodSecurityContact
    
    params["FoodSecurityDesignCode"] = req.FoodSecurityDesignCode
    
    params["FoodSecurityFactory"] = req.FoodSecurityFactory
    
    params["FoodSecurityFactorySite"] = req.FoodSecurityFactorySite
    
    params["FoodSecurityFoodAdditive"] = req.FoodSecurityFoodAdditive
    
    params["FoodSecurityHealthProductNo"] = req.FoodSecurityHealthProductNo
    
    params["FoodSecurityMix"] = req.FoodSecurityMix
    
    params["FoodSecurityPeriod"] = req.FoodSecurityPeriod
    
    params["FoodSecurityPlanStorage"] = req.FoodSecurityPlanStorage
    
    params["FoodSecurityPrdLicenseNo"] = req.FoodSecurityPrdLicenseNo
    
    params["FoodSecurityProductDateEnd"] = req.FoodSecurityProductDateEnd
    
    params["FoodSecurityProductDateStart"] = req.FoodSecurityProductDateStart
    
    params["FoodSecurityStockDateEnd"] = req.FoodSecurityStockDateEnd
    
    params["FoodSecurityStockDateStart"] = req.FoodSecurityStockDateStart
    
    params["FoodSecuritySupplier"] = req.FoodSecuritySupplier
    
    params["FreightPayer"] = req.FreightPayer
    
    params["GlobalStockCountry"] = req.GlobalStockCountry
    
    params["GlobalStockDeliveryPlace"] = req.GlobalStockDeliveryPlace
    
    params["GlobalStockTaxFreePromise"] = req.GlobalStockTaxFreePromise
    
    params["GlobalStockType"] = req.GlobalStockType
    
    params["HasDiscount"] = req.HasDiscount
    
    params["HasInvoice"] = req.HasInvoice
    
    params["HasShowcase"] = req.HasShowcase
    
    params["HasWarranty"] = req.HasWarranty
    
    params["Ignorewarning"] = req.Ignorewarning
    
    params["Image"] = req.Image
    
    params["Increment"] = req.Increment
    
    params["InputCustomCpv"] = req.InputCustomCpv
    
    params["InputPids"] = req.InputPids
    
    params["InputStr"] = req.InputStr
    
    params["InteractiveId"] = req.InteractiveId
    
    params["Is3D"] = req.Is3D
    
    params["IsEx"] = req.IsEx
    
    params["IsLightningConsignment"] = req.IsLightningConsignment
    
    params["IsOffline"] = req.IsOffline
    
    params["IsReplaceSku"] = req.IsReplaceSku
    
    params["IsTaobao"] = req.IsTaobao
    
    params["IsXinpin"] = req.IsXinpin
    
    params["ItemSize"] = req.ItemSize
    
    params["ItemWeight"] = req.ItemWeight
    
    params["Lang"] = req.Lang
    
    params["LeaseExtendsInfo"] = req.LeaseExtendsInfo
    
    params["ListTime"] = req.ListTime
    
    params["LocalityLifeChooseLogis"] = req.LocalityLifeChooseLogis
    
    params["LocalityLifeEticket"] = req.LocalityLifeEticket
    
    params["LocalityLifeExpirydate"] = req.LocalityLifeExpirydate
    
    params["LocalityLifeMerchant"] = req.LocalityLifeMerchant
    
    params["LocalityLifeNetworkId"] = req.LocalityLifeNetworkId
    
    params["LocalityLifeObs"] = req.LocalityLifeObs
    
    params["LocalityLifeOnsaleAutoRefundRatio"] = req.LocalityLifeOnsaleAutoRefundRatio
    
    params["LocalityLifePackageid"] = req.LocalityLifePackageid
    
    params["LocalityLifeRefundRatio"] = req.LocalityLifeRefundRatio
    
    params["LocalityLifeRefundmafee"] = req.LocalityLifeRefundmafee
    
    params["LocalityLifeVerification"] = req.LocalityLifeVerification
    
    params["LocalityLifeVersion"] = req.LocalityLifeVersion
    
    params["LocationCity"] = req.LocationCity
    
    params["LocationState"] = req.LocationState
    
    params["MsPaymentPrice"] = req.MsPaymentPrice
    
    params["MsPaymentReferencePrice"] = req.MsPaymentReferencePrice
    
    params["MsPaymentVoucherPrice"] = req.MsPaymentVoucherPrice
    
    params["Newprepay"] = req.Newprepay
    
    params["Num"] = req.Num
    
    params["NumIid"] = req.NumIid
    
    params["O2oBindService"] = req.O2oBindService
    
    params["OuterId"] = req.OuterId
    
    params["PaimaiInfoDeposit"] = req.PaimaiInfoDeposit
    
    params["PaimaiInfoInterval"] = req.PaimaiInfoInterval
    
    params["PaimaiInfoMode"] = req.PaimaiInfoMode
    
    params["PaimaiInfoReserve"] = req.PaimaiInfoReserve
    
    params["PaimaiInfoValidHour"] = req.PaimaiInfoValidHour
    
    params["PaimaiInfoValidMinute"] = req.PaimaiInfoValidMinute
    
    params["PicPath"] = req.PicPath
    
    params["PostFee"] = req.PostFee
    
    params["PostageId"] = req.PostageId
    
    params["Price"] = req.Price
    
    params["ProductId"] = req.ProductId
    
    params["PropertyAlias"] = req.PropertyAlias
    
    params["Props"] = req.Props
    
    params["Qualification"] = req.Qualification
    
    params["ScenicTicketBookCost"] = req.ScenicTicketBookCost
    
    params["ScenicTicketPayWay"] = req.ScenicTicketPayWay
    
    params["SellPoint"] = req.SellPoint
    
    params["SellPromise"] = req.SellPromise
    
    params["SellerCids"] = req.SellerCids
    
    params["Shape"] = req.Shape
    
    params["SkuBarcode"] = req.SkuBarcode
    
    params["SkuDeliveryTimes"] = req.SkuDeliveryTimes
    
    params["SkuHdHeight"] = req.SkuHdHeight
    
    params["SkuHdLampQuantity"] = req.SkuHdLampQuantity
    
    params["SkuHdLength"] = req.SkuHdLength
    
    params["SkuOuterIds"] = req.SkuOuterIds
    
    params["SkuPrices"] = req.SkuPrices
    
    params["SkuProperties"] = req.SkuProperties
    
    params["SkuQuantities"] = req.SkuQuantities
    
    params["SkuSpecIds"] = req.SkuSpecIds
    
    params["SpuConfirm"] = req.SpuConfirm
    
    params["StuffStatus"] = req.StuffStatus
    
    params["SubStock"] = req.SubStock
    
    params["Title"] = req.Title
    
    params["ValidThru"] = req.ValidThru
    
    params["VideoId"] = req.VideoId
    
    params["Weight"] = req.Weight
    
    params["WirelessDesc"] = req.WirelessDesc
    
    return params
}

// TaobaoItemUpdateResponse 根据传入的num_iid更新对应的商品的数据 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败） 
//商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
//当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。
type TaobaoItemUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object商品结构里的num_iid，modified
    Item Item `json:"item";xml:"item"`
    
}

// TaobaoTmcUserCancelRequest 取消用户的消息服务
type TaobaoTmcUserCancelRequest struct {
    
    // Nick required用户昵称
    Nick string `json:"nick";xml:"nick"`
    
    // UserPlatform optional用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    UserPlatform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcUserCancelRequest) GetAPIName() string {
	return "taobao.tmc.user.cancel"
}

func (req *TaobaoTmcUserCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Nick"] = req.Nick
    
    params["UserPlatform"] = req.UserPlatform
    
    return params
}

// TaobaoTmcUserCancelResponse 取消用户的消息服务
type TaobaoTmcUserCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功,如果为false并且没有错误码，表示删除的用户不存在。
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoNextoneLogisticsWarehouseUpdateRequest 商家上传退货入仓状态给ag
type TaobaoNextoneLogisticsWarehouseUpdateRequest struct {
    
    // RefundId required退款编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // WarehouseStatus required退货入仓状态 1.已入仓
    WarehouseStatus int64 `json:"warehouse_status";xml:"warehouse_status"`
    
}

func (req *TaobaoNextoneLogisticsWarehouseUpdateRequest) GetAPIName() string {
	return "taobao.nextone.logistics.warehouse.update"
}

func (req *TaobaoNextoneLogisticsWarehouseUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["RefundId"] = req.RefundId
    
    params["WarehouseStatus"] = req.WarehouseStatus
    
    return params
}

// TaobaoNextoneLogisticsWarehouseUpdateResponse 商家上传退货入仓状态给ag
type TaobaoNextoneLogisticsWarehouseUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode BasicerrorCode
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // ErrInfo BasicerrorInfo
    ErrInfo string `json:"err_info";xml:"err_info"`
    
    // ResultData BasicresultData
    ResultData map[string]interface{} `json:"result_data";xml:"result_data"`
    
    // Succeed Basicsuccess
    Succeed bool `json:"succeed";xml:"succeed"`
    
}

// TaobaoItemImgUploadRequest 添加一张商品图片到num_iid指定的商品中 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
//使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
//商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
type TaobaoItemImgUploadRequest struct {
    
    // Id optional商品图片id(如果是更新图片，则需要传该参数)
    Id int64 `json:"id";xml:"id"`
    
    // Image optional商品图片内容类型:JPG,GIF;最大:3M 。支持的文件类型：gif,jpg,jpeg,png
    Image []byte `json:"image";xml:"image"`
    
    // IsMajor optional是否将该图片设为主图,可选值:true,false;默认值:false(非主图)
    IsMajor bool `json:"is_major";xml:"is_major"`
    
    // IsRectangle optional是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
    IsRectangle bool `json:"is_rectangle";xml:"is_rectangle"`
    
    // NumIid required商品数字ID，该参数必须
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // Position optional图片序号
    Position int64 `json:"position";xml:"position"`
    
}

func (req *TaobaoItemImgUploadRequest) GetAPIName() string {
	return "taobao.item.img.upload"
}

func (req *TaobaoItemImgUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Id"] = req.Id
    
    params["Image"] = req.Image
    
    params["IsMajor"] = req.IsMajor
    
    params["IsRectangle"] = req.IsRectangle
    
    params["NumIid"] = req.NumIid
    
    params["Position"] = req.Position
    
    return params
}

// TaobaoItemImgUploadResponse 添加一张商品图片到num_iid指定的商品中 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
//使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
//商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
type TaobaoItemImgUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemImg Object商品图片结构
    ItemImg ItemImg `json:"item_img";xml:"item_img"`
    
}

// TaobaoTmcUserPermitRequest 为已授权的用户开通消息服务
type TaobaoTmcUserPermitRequest struct {
    
    // Topics optional消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
    Topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcUserPermitRequest) GetAPIName() string {
	return "taobao.tmc.user.permit"
}

func (req *TaobaoTmcUserPermitRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Topics"] = req.Topics
    
    return params
}

// TaobaoTmcUserPermitResponse 为已授权的用户开通消息服务
type TaobaoTmcUserPermitResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TmallChannelProductsGetRequest 查询供应商的产品数据。 
//
//* 入参传入pids将优先查询，即只按这个条件查询。 
//*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
//* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
//* 入参fields传入images将查询多图数据，不传只返回主图数据。 
//* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
//* 查询结果按照产品发布时间倒序，即时间近的数据在前。
//* 传入channel 渠道，会只返回相应渠道的产品
type TmallChannelProductsGetRequest struct {
    
    // TopQueryProductDO optionaltop_query_product_d_o
    TopQueryProductDO TopQueryProductDo `json:"top_query_product_d_o";xml:"top_query_product_d_o"`
    
}

func (req *TmallChannelProductsGetRequest) GetAPIName() string {
	return "tmall.channel.products.get"
}

func (req *TmallChannelProductsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TopQueryProductDO"] = req.TopQueryProductDO
    
    return params
}

// TmallChannelProductsGetResponse 查询供应商的产品数据。 
//
//* 入参传入pids将优先查询，即只按这个条件查询。 
//*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
//* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
//* 入参fields传入images将查询多图数据，不传只返回主图数据。 
//* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
//* 查询结果按照产品发布时间倒序，即时间近的数据在前。
//* 传入channel 渠道，会只返回相应渠道的产品
type TmallChannelProductsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Products Object Array产品对象记录集
    Products TopProductDO `json:"products";xml:"products"`
    
    // TotalResults Basic查询结果记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoItemAddRequest 此接口用于新增一个商品 
//商品所属的卖家是当前会话的用户 
//商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
//商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
//商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
//当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
type TaobaoItemAddRequest struct {
    
    // AfterSaleId optional售后说明模板id
    AfterSaleId int64 `json:"after_sale_id";xml:"after_sale_id"`
    
    // ApproveStatus optional商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale
    ApproveStatus string `json:"approve_status";xml:"approve_status"`
    
    // AuctionPoint optional商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
    AuctionPoint int64 `json:"auction_point";xml:"auction_point"`
    
    // AutoFill optional代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
    AutoFill string `json:"auto_fill";xml:"auto_fill"`
    
    // AutoRepost optional自动重发。可选值:true,false;默认值:false(不重发)
    AutoRepost bool `json:"auto_repost";xml:"auto_repost"`
    
    // Barcode optional商品条形码
    Barcode string `json:"barcode";xml:"barcode"`
    
    // ChangeProp optional基础色数据
    ChangeProp string `json:"change_prop";xml:"change_prop"`
    
    // ChaoshiExtendsInfo optional天猫超市扩展字段，天猫超市专用。
    ChaoshiExtendsInfo string `json:"chaoshi_extends_info";xml:"chaoshi_extends_info"`
    
    // Cid required叶子类目id
    Cid int64 `json:"cid";xml:"cid"`
    
    // CodPostageId optional此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃
    CodPostageId int64 `json:"cod_postage_id";xml:"cod_postage_id"`
    
    // CpvMemo optional针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
    CpvMemo string `json:"cpv_memo";xml:"cpv_memo"`
    
    // CustomMadeTypeId optional定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1
    CustomMadeTypeId string `json:"custom_made_type_id";xml:"custom_made_type_id"`
    
    // DeliveryTimeDeliveryTime optional商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
    DeliveryTimeDeliveryTime string `json:"delivery_time.delivery_time";xml:"delivery_time.delivery_time"`
    
    // DeliveryTimeDeliveryTimeType optional发货时间类型：绝对发货时间或者相对发货时间
    DeliveryTimeDeliveryTimeType string `json:"delivery_time.delivery_time_type";xml:"delivery_time.delivery_time_type"`
    
    // DeliveryTimeNeedDeliveryTime optional设置是否使用发货时间，商品级别，sku级别
    DeliveryTimeNeedDeliveryTime string `json:"delivery_time.need_delivery_time";xml:"delivery_time.need_delivery_time"`
    
    // Desc required宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制
    Desc string `json:"desc";xml:"desc"`
    
    // DescModules optional商品描述模块化，模块列表，由List转化成jsonArray存入，后端逻辑验证通过，拼装成模块内容+锚点导航后存入desc中。数据结构具体参见Item_Desc_Module
    DescModules string `json:"desc_modules";xml:"desc_modules"`
    
    // EmsFee optionalems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
    EmsFee float32 `json:"ems_fee";xml:"ems_fee"`
    
    // ExpressFee optional快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
    ExpressFee float32 `json:"express_fee";xml:"express_fee"`
    
    // Features optional宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp
    Features string `json:"features";xml:"features"`
    
    // FoodSecurityContact optional厂家联系方式
    FoodSecurityContact string `json:"food_security.contact";xml:"food_security.contact"`
    
    // FoodSecurityDesignCode optional产品标准号
    FoodSecurityDesignCode string `json:"food_security.design_code";xml:"food_security.design_code"`
    
    // FoodSecurityFactory optional厂名
    FoodSecurityFactory string `json:"food_security.factory";xml:"food_security.factory"`
    
    // FoodSecurityFactorySite optional厂址
    FoodSecurityFactorySite string `json:"food_security.factory_site";xml:"food_security.factory_site"`
    
    // FoodSecurityFoodAdditive optional食品添加剂
    FoodSecurityFoodAdditive string `json:"food_security.food_additive";xml:"food_security.food_additive"`
    
    // FoodSecurityHealthProductNo optional健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
    FoodSecurityHealthProductNo string `json:"food_security.health_product_no";xml:"food_security.health_product_no"`
    
    // FoodSecurityMix optional配料表
    FoodSecurityMix string `json:"food_security.mix";xml:"food_security.mix"`
    
    // FoodSecurityPeriod optional保质期
    FoodSecurityPeriod string `json:"food_security.period";xml:"food_security.period"`
    
    // FoodSecurityPlanStorage optional储藏方法
    FoodSecurityPlanStorage string `json:"food_security.plan_storage";xml:"food_security.plan_storage"`
    
    // FoodSecurityPrdLicenseNo optional生产许可证号
    FoodSecurityPrdLicenseNo string `json:"food_security.prd_license_no";xml:"food_security.prd_license_no"`
    
    // FoodSecurityProductDateEnd optional生产结束日期,格式必须为yyyy-MM-dd
    FoodSecurityProductDateEnd string `json:"food_security.product_date_end";xml:"food_security.product_date_end"`
    
    // FoodSecurityProductDateStart optional生产开始日期，格式必须为yyyy-MM-dd
    FoodSecurityProductDateStart string `json:"food_security.product_date_start";xml:"food_security.product_date_start"`
    
    // FoodSecurityStockDateEnd optional进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
    FoodSecurityStockDateEnd string `json:"food_security.stock_date_end";xml:"food_security.stock_date_end"`
    
    // FoodSecurityStockDateStart optional进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
    FoodSecurityStockDateStart string `json:"food_security.stock_date_start";xml:"food_security.stock_date_start"`
    
    // FoodSecuritySupplier optional供货商
    FoodSecuritySupplier string `json:"food_security.supplier";xml:"food_security.supplier"`
    
    // FreightPayer optional运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee）
    FreightPayer string `json:"freight_payer";xml:"freight_payer"`
    
    // GlobalStockCountry optional全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）
    GlobalStockCountry string `json:"global_stock_country";xml:"global_stock_country"`
    
    // GlobalStockDeliveryPlace optional全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台”
    GlobalStockDeliveryPlace string `json:"global_stock_delivery_place";xml:"global_stock_delivery_place"`
    
    // GlobalStockTaxFreePromise optional全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约
    GlobalStockTaxFreePromise bool `json:"global_stock_tax_free_promise";xml:"global_stock_tax_free_promise"`
    
    // GlobalStockType optional全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
    GlobalStockType string `json:"global_stock_type";xml:"global_stock_type"`
    
    // HasDiscount optional支持会员打折。可选值:true,false;默认值:false(不打折)
    HasDiscount bool `json:"has_discount";xml:"has_discount"`
    
    // HasInvoice optional是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)
    HasInvoice bool `json:"has_invoice";xml:"has_invoice"`
    
    // HasShowcase optional橱窗推荐。可选值:true,false;默认值:false(不推荐)
    HasShowcase bool `json:"has_showcase";xml:"has_showcase"`
    
    // HasWarranty optional是否有保修。可选值:true,false;默认值:false(不保修)
    HasWarranty bool `json:"has_warranty";xml:"has_warranty"`
    
    // Ignorewarning optional忽略警告提示.
    Ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    // Image optional商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间）
    Image []byte `json:"image";xml:"image"`
    
    // Increment optional加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
    Increment float32 `json:"increment";xml:"increment"`
    
    // InputCustomCpv optional针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上
    InputCustomCpv string `json:"input_custom_cpv";xml:"input_custom_cpv"`
    
    // InputPids optional用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
    InputPids string `json:"input_pids";xml:"input_pids"`
    
    // InputStr optional用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
    InputStr string `json:"input_str";xml:"input_str"`
    
    // InteractiveId optional主图视频互动信息id，必须填写主图视频id才能有互动信息id
    InteractiveId int64 `json:"interactive_id";xml:"interactive_id"`
    
    // Is3D optional是否是3D
    Is3D bool `json:"is_3D";xml:"is_3D"`
    
    // IsEx optional是否在外店显示
    IsEx bool `json:"is_ex";xml:"is_ex"`
    
    // IsLightningConsignment optional实物闪电发货
    IsLightningConsignment bool `json:"is_lightning_consignment";xml:"is_lightning_consignment"`
    
    // IsOffline optional是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
    IsOffline string `json:"is_offline";xml:"is_offline"`
    
    // IsTaobao optional是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
    IsTaobao bool `json:"is_taobao";xml:"is_taobao"`
    
    // IsXinpin optional商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。
    IsXinpin bool `json:"is_xinpin";xml:"is_xinpin"`
    
    // ItemSize optional表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）
    ItemSize string `json:"item_size";xml:"item_size"`
    
    // ItemWeight optional商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。
    ItemWeight string `json:"item_weight";xml:"item_weight"`
    
    // Lang optional商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体
    Lang string `json:"lang";xml:"lang"`
    
    // LeaseExtendsInfo optional租赁扩展信息
    LeaseExtendsInfo string `json:"lease_extends_info";xml:"lease_extends_info"`
    
    // ListTime optional定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss)
    ListTime time.Time `json:"list_time";xml:"list_time"`
    
    // LocalityLifeChooseLogis optional发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
    LocalityLifeChooseLogis string `json:"locality_life.choose_logis";xml:"locality_life.choose_logis"`
    
    // LocalityLifeEticket optional电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
    LocalityLifeEticket string `json:"locality_life.eticket";xml:"locality_life.eticket"`
    
    // LocalityLifeExpirydate optional本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
    LocalityLifeExpirydate string `json:"locality_life.expirydate";xml:"locality_life.expirydate"`
    
    // LocalityLifeMerchant optional码商信息，格式为 码商id:nick
    LocalityLifeMerchant string `json:"locality_life.merchant";xml:"locality_life.merchant"`
    
    // LocalityLifeNetworkId optional网点ID
    LocalityLifeNetworkId string `json:"locality_life.network_id";xml:"locality_life.network_id"`
    
    // LocalityLifeObs optional预约门店是否支持门店自提,1:是
    LocalityLifeObs string `json:"locality_life.obs";xml:"locality_life.obs"`
    
    // LocalityLifeOnsaleAutoRefundRatio optional电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
    LocalityLifeOnsaleAutoRefundRatio int64 `json:"locality_life.onsale_auto_refund_ratio";xml:"locality_life.onsale_auto_refund_ratio"`
    
    // LocalityLifePackageid optional新版电子凭证包 id
    LocalityLifePackageid string `json:"locality_life.packageid";xml:"locality_life.packageid"`
    
    // LocalityLifeRefundRatio optional退款比例，百分比%前的数字,1-100的正整数值
    LocalityLifeRefundRatio int64 `json:"locality_life.refund_ratio";xml:"locality_life.refund_ratio"`
    
    // LocalityLifeRefundmafee optional退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
    LocalityLifeRefundmafee string `json:"locality_life.refundmafee";xml:"locality_life.refundmafee"`
    
    // LocalityLifeVerification optional核销打款 1代表核销打款 0代表非核销打款
    LocalityLifeVerification string `json:"locality_life.verification";xml:"locality_life.verification"`
    
    // LocalityLifeVersion optional新版电子凭证字段
    LocalityLifeVersion string `json:"locality_life.version";xml:"locality_life.version"`
    
    // LocationCity required所在地城市。如杭州 。
    LocationCity string `json:"location.city";xml:"location.city"`
    
    // LocationState required所在地省份。如浙江
    LocationState string `json:"location.state";xml:"location.state"`
    
    // MsPaymentPrice optional订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
    MsPaymentPrice string `json:"ms_payment.price";xml:"ms_payment.price"`
    
    // MsPaymentReferencePrice optional参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
    MsPaymentReferencePrice string `json:"ms_payment.reference_price";xml:"ms_payment.reference_price"`
    
    // MsPaymentVoucherPrice optional尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
    MsPaymentVoucherPrice string `json:"ms_payment.voucher_price";xml:"ms_payment.voucher_price"`
    
    // Newprepay optional该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；
    Newprepay string `json:"newprepay";xml:"newprepay"`
    
    // Num required商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
    Num int64 `json:"num";xml:"num"`
    
    // O2oBindService optional汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
    O2oBindService bool `json:"o2o_bind_service";xml:"o2o_bind_service"`
    
    // OuterId optional商品外部编码，该字段的最大长度是64个字节
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // PaimaiInfoDeposit optional拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。
    PaimaiInfoDeposit int64 `json:"paimai_info.deposit";xml:"paimai_info.deposit"`
    
    // PaimaiInfoInterval optional降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
    PaimaiInfoInterval int64 `json:"paimai_info.interval";xml:"paimai_info.interval"`
    
    // PaimaiInfoMode optional拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
    PaimaiInfoMode int64 `json:"paimai_info.mode";xml:"paimai_info.mode"`
    
    // PaimaiInfoReserve optional降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
    PaimaiInfoReserve float32 `json:"paimai_info.reserve";xml:"paimai_info.reserve"`
    
    // PaimaiInfoValidHour optional自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    PaimaiInfoValidHour int64 `json:"paimai_info.valid_hour";xml:"paimai_info.valid_hour"`
    
    // PaimaiInfoValidMinute optional自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    PaimaiInfoValidMinute int64 `json:"paimai_info.valid_minute";xml:"paimai_info.valid_minute"`
    
    // PicPath optional（推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // PostFee optional平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写
    PostFee float32 `json:"post_fee";xml:"post_fee"`
    
    // PostageId optional宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板）
    PostageId int64 `json:"postage_id";xml:"postage_id"`
    
    // Price optional商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。
    Price float32 `json:"price";xml:"price"`
    
    // ProductId optional商品所属的产品ID(B商家发布商品需要用)
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // PropertyAlias optional属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。
    PropertyAlias string `json:"property_alias";xml:"property_alias"`
    
    // Props optional商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值
    Props string `json:"props";xml:"props"`
    
    // Qualification optional商品资质信息
    Qualification string `json:"qualification";xml:"qualification"`
    
    // ScenicTicketBookCost optional景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
    ScenicTicketBookCost string `json:"scenic_ticket_book_cost";xml:"scenic_ticket_book_cost"`
    
    // ScenicTicketPayWay optional景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
    ScenicTicketPayWay int64 `json:"scenic_ticket_pay_way";xml:"scenic_ticket_pay_way"`
    
    // SellPoint optional商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。
    SellPoint string `json:"sell_point";xml:"sell_point"`
    
    // SellPromise optional是否承诺退换货服务!虚拟商品无须设置此项!
    SellPromise bool `json:"sell_promise";xml:"sell_promise"`
    
    // SellerCids optional商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
    SellerCids string `json:"seller_cids";xml:"seller_cids"`
    
    // Shape optional宝贝形态:1代表电子券;0或其他代表实物
    Shape string `json:"shape";xml:"shape"`
    
    // SkuBarcode optionalsku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
    SkuBarcode string `json:"sku_barcode";xml:"sku_barcode"`
    
    // SkuDeliveryTimes optional此参数暂时不起作用
    SkuDeliveryTimes string `json:"sku_delivery_times";xml:"sku_delivery_times"`
    
    // SkuHdHeight optional家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。 可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50
    SkuHdHeight string `json:"sku_hd_height";xml:"sku_hd_height"`
    
    // SkuHdLampQuantity optional家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7
    SkuHdLampQuantity string `json:"sku_hd_lamp_quantity";xml:"sku_hd_lamp_quantity"`
    
    // SkuHdLength optional家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：20,30,30
    SkuHdLength string `json:"sku_hd_length";xml:"sku_hd_length"`
    
    // SkuOuterIds optionalSku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
    SkuOuterIds string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    // SkuPrices optionalSku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
    SkuPrices string `json:"sku_prices";xml:"sku_prices"`
    
    // SkuProperties optional更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
    SkuProperties string `json:"sku_properties";xml:"sku_properties"`
    
    // SkuQuantities optionalSku的数量串，结构如：num1,num2,num3 如：2,3
    SkuQuantities string `json:"sku_quantities";xml:"sku_quantities"`
    
    // SkuSpecIds optional此参数暂时不起作用
    SkuSpecIds string `json:"sku_spec_ids";xml:"sku_spec_ids"`
    
    // SpuConfirm optional手机类目spu 优化，信息确认字段
    SpuConfirm bool `json:"spu_confirm";xml:"spu_confirm"`
    
    // StuffStatus required新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性
    StuffStatus string `json:"stuff_status";xml:"stuff_status"`
    
    // SubStock optional商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存
    SubStock int64 `json:"sub_stock";xml:"sub_stock"`
    
    // SupportCustomMade optional是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改
    SupportCustomMade bool `json:"support_custom_made";xml:"support_custom_made"`
    
    // Title required宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；
    Title string `json:"title";xml:"title"`
    
    // Type required发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。
    Type string `json:"type";xml:"type"`
    
    // ValidThru optional有效期。可选值:7,14;单位:天;默认值:14
    ValidThru int64 `json:"valid_thru";xml:"valid_thru"`
    
    // VideoId optional主图视频id
    VideoId int64 `json:"video_id";xml:"video_id"`
    
    // Weight optional商品的重量(商超卖家专用字段)
    Weight int64 `json:"weight";xml:"weight"`
    
    // WirelessDesc optional无线的宝贝描述
    WirelessDesc string `json:"wireless_desc";xml:"wireless_desc"`
    
}

func (req *TaobaoItemAddRequest) GetAPIName() string {
	return "taobao.item.add"
}

func (req *TaobaoItemAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 123)
    
    params["AfterSaleId"] = req.AfterSaleId
    
    params["ApproveStatus"] = req.ApproveStatus
    
    params["AuctionPoint"] = req.AuctionPoint
    
    params["AutoFill"] = req.AutoFill
    
    params["AutoRepost"] = req.AutoRepost
    
    params["Barcode"] = req.Barcode
    
    params["ChangeProp"] = req.ChangeProp
    
    params["ChaoshiExtendsInfo"] = req.ChaoshiExtendsInfo
    
    params["Cid"] = req.Cid
    
    params["CodPostageId"] = req.CodPostageId
    
    params["CpvMemo"] = req.CpvMemo
    
    params["CustomMadeTypeId"] = req.CustomMadeTypeId
    
    params["DeliveryTimeDeliveryTime"] = req.DeliveryTimeDeliveryTime
    
    params["DeliveryTimeDeliveryTimeType"] = req.DeliveryTimeDeliveryTimeType
    
    params["DeliveryTimeNeedDeliveryTime"] = req.DeliveryTimeNeedDeliveryTime
    
    params["Desc"] = req.Desc
    
    params["DescModules"] = req.DescModules
    
    params["EmsFee"] = req.EmsFee
    
    params["ExpressFee"] = req.ExpressFee
    
    params["Features"] = req.Features
    
    params["FoodSecurityContact"] = req.FoodSecurityContact
    
    params["FoodSecurityDesignCode"] = req.FoodSecurityDesignCode
    
    params["FoodSecurityFactory"] = req.FoodSecurityFactory
    
    params["FoodSecurityFactorySite"] = req.FoodSecurityFactorySite
    
    params["FoodSecurityFoodAdditive"] = req.FoodSecurityFoodAdditive
    
    params["FoodSecurityHealthProductNo"] = req.FoodSecurityHealthProductNo
    
    params["FoodSecurityMix"] = req.FoodSecurityMix
    
    params["FoodSecurityPeriod"] = req.FoodSecurityPeriod
    
    params["FoodSecurityPlanStorage"] = req.FoodSecurityPlanStorage
    
    params["FoodSecurityPrdLicenseNo"] = req.FoodSecurityPrdLicenseNo
    
    params["FoodSecurityProductDateEnd"] = req.FoodSecurityProductDateEnd
    
    params["FoodSecurityProductDateStart"] = req.FoodSecurityProductDateStart
    
    params["FoodSecurityStockDateEnd"] = req.FoodSecurityStockDateEnd
    
    params["FoodSecurityStockDateStart"] = req.FoodSecurityStockDateStart
    
    params["FoodSecuritySupplier"] = req.FoodSecuritySupplier
    
    params["FreightPayer"] = req.FreightPayer
    
    params["GlobalStockCountry"] = req.GlobalStockCountry
    
    params["GlobalStockDeliveryPlace"] = req.GlobalStockDeliveryPlace
    
    params["GlobalStockTaxFreePromise"] = req.GlobalStockTaxFreePromise
    
    params["GlobalStockType"] = req.GlobalStockType
    
    params["HasDiscount"] = req.HasDiscount
    
    params["HasInvoice"] = req.HasInvoice
    
    params["HasShowcase"] = req.HasShowcase
    
    params["HasWarranty"] = req.HasWarranty
    
    params["Ignorewarning"] = req.Ignorewarning
    
    params["Image"] = req.Image
    
    params["Increment"] = req.Increment
    
    params["InputCustomCpv"] = req.InputCustomCpv
    
    params["InputPids"] = req.InputPids
    
    params["InputStr"] = req.InputStr
    
    params["InteractiveId"] = req.InteractiveId
    
    params["Is3D"] = req.Is3D
    
    params["IsEx"] = req.IsEx
    
    params["IsLightningConsignment"] = req.IsLightningConsignment
    
    params["IsOffline"] = req.IsOffline
    
    params["IsTaobao"] = req.IsTaobao
    
    params["IsXinpin"] = req.IsXinpin
    
    params["ItemSize"] = req.ItemSize
    
    params["ItemWeight"] = req.ItemWeight
    
    params["Lang"] = req.Lang
    
    params["LeaseExtendsInfo"] = req.LeaseExtendsInfo
    
    params["ListTime"] = req.ListTime
    
    params["LocalityLifeChooseLogis"] = req.LocalityLifeChooseLogis
    
    params["LocalityLifeEticket"] = req.LocalityLifeEticket
    
    params["LocalityLifeExpirydate"] = req.LocalityLifeExpirydate
    
    params["LocalityLifeMerchant"] = req.LocalityLifeMerchant
    
    params["LocalityLifeNetworkId"] = req.LocalityLifeNetworkId
    
    params["LocalityLifeObs"] = req.LocalityLifeObs
    
    params["LocalityLifeOnsaleAutoRefundRatio"] = req.LocalityLifeOnsaleAutoRefundRatio
    
    params["LocalityLifePackageid"] = req.LocalityLifePackageid
    
    params["LocalityLifeRefundRatio"] = req.LocalityLifeRefundRatio
    
    params["LocalityLifeRefundmafee"] = req.LocalityLifeRefundmafee
    
    params["LocalityLifeVerification"] = req.LocalityLifeVerification
    
    params["LocalityLifeVersion"] = req.LocalityLifeVersion
    
    params["LocationCity"] = req.LocationCity
    
    params["LocationState"] = req.LocationState
    
    params["MsPaymentPrice"] = req.MsPaymentPrice
    
    params["MsPaymentReferencePrice"] = req.MsPaymentReferencePrice
    
    params["MsPaymentVoucherPrice"] = req.MsPaymentVoucherPrice
    
    params["Newprepay"] = req.Newprepay
    
    params["Num"] = req.Num
    
    params["O2oBindService"] = req.O2oBindService
    
    params["OuterId"] = req.OuterId
    
    params["PaimaiInfoDeposit"] = req.PaimaiInfoDeposit
    
    params["PaimaiInfoInterval"] = req.PaimaiInfoInterval
    
    params["PaimaiInfoMode"] = req.PaimaiInfoMode
    
    params["PaimaiInfoReserve"] = req.PaimaiInfoReserve
    
    params["PaimaiInfoValidHour"] = req.PaimaiInfoValidHour
    
    params["PaimaiInfoValidMinute"] = req.PaimaiInfoValidMinute
    
    params["PicPath"] = req.PicPath
    
    params["PostFee"] = req.PostFee
    
    params["PostageId"] = req.PostageId
    
    params["Price"] = req.Price
    
    params["ProductId"] = req.ProductId
    
    params["PropertyAlias"] = req.PropertyAlias
    
    params["Props"] = req.Props
    
    params["Qualification"] = req.Qualification
    
    params["ScenicTicketBookCost"] = req.ScenicTicketBookCost
    
    params["ScenicTicketPayWay"] = req.ScenicTicketPayWay
    
    params["SellPoint"] = req.SellPoint
    
    params["SellPromise"] = req.SellPromise
    
    params["SellerCids"] = req.SellerCids
    
    params["Shape"] = req.Shape
    
    params["SkuBarcode"] = req.SkuBarcode
    
    params["SkuDeliveryTimes"] = req.SkuDeliveryTimes
    
    params["SkuHdHeight"] = req.SkuHdHeight
    
    params["SkuHdLampQuantity"] = req.SkuHdLampQuantity
    
    params["SkuHdLength"] = req.SkuHdLength
    
    params["SkuOuterIds"] = req.SkuOuterIds
    
    params["SkuPrices"] = req.SkuPrices
    
    params["SkuProperties"] = req.SkuProperties
    
    params["SkuQuantities"] = req.SkuQuantities
    
    params["SkuSpecIds"] = req.SkuSpecIds
    
    params["SpuConfirm"] = req.SpuConfirm
    
    params["StuffStatus"] = req.StuffStatus
    
    params["SubStock"] = req.SubStock
    
    params["SupportCustomMade"] = req.SupportCustomMade
    
    params["Title"] = req.Title
    
    params["Type"] = req.Type
    
    params["ValidThru"] = req.ValidThru
    
    params["VideoId"] = req.VideoId
    
    params["Weight"] = req.Weight
    
    params["WirelessDesc"] = req.WirelessDesc
    
    return params
}

// TaobaoItemAddResponse 此接口用于新增一个商品 
//商品所属的卖家是当前会话的用户 
//商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
//商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
//商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
//当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
type TaobaoItemAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object商品结构,仅有numIid和created返回
    Item Item `json:"item";xml:"item"`
    
}

// TaobaoItemPropimgDeleteRequest 删除propimg_id 所指定的商品属性图片 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//propimg_id对应的属性图片需要属于num_iid对应的商品
type TaobaoItemPropimgDeleteRequest struct {
    
    // Id required商品属性图片ID
    Id int64 `json:"id";xml:"id"`
    
    // NumIid required商品数字ID，必选
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemPropimgDeleteRequest) GetAPIName() string {
	return "taobao.item.propimg.delete"
}

func (req *TaobaoItemPropimgDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Id"] = req.Id
    
    params["NumIid"] = req.NumIid
    
    return params
}

// TaobaoItemPropimgDeleteResponse 删除propimg_id 所指定的商品属性图片 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//propimg_id对应的属性图片需要属于num_iid对应的商品
type TaobaoItemPropimgDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PropImg Object属性图片结构
    PropImg PropImg `json:"prop_img";xml:"prop_img"`
    
}

// TaobaoItemImgDeleteRequest 删除商品图片
type TaobaoItemImgDeleteRequest struct {
    
    // Id required商品图片ID；如果是竖图，请将id的值设置为1
    Id int64 `json:"id";xml:"id"`
    
    // IsSixthPic optional标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
    IsSixthPic bool `json:"is_sixth_pic";xml:"is_sixth_pic"`
    
    // NumIid required商品数字ID
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemImgDeleteRequest) GetAPIName() string {
	return "taobao.item.img.delete"
}

func (req *TaobaoItemImgDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Id"] = req.Id
    
    params["IsSixthPic"] = req.IsSixthPic
    
    params["NumIid"] = req.NumIid
    
    return params
}

// TaobaoItemImgDeleteResponse 删除商品图片
type TaobaoItemImgDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemImg Object商品图片结构
    ItemImg ItemImg `json:"item_img";xml:"item_img"`
    
}

// TaobaoItemSkuAddRequest 新增一个sku到num_iid指定的商品中 
//传入的iid所对应的商品必须属于当前会话的用户
type TaobaoItemSkuAddRequest struct {
    
    // Ignorewarning optional忽略警告提示.
    Ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    // ItemPrice optionalsku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
    ItemPrice float32 `json:"item_price";xml:"item_price"`
    
    // Lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    Lang string `json:"lang";xml:"lang"`
    
    // NumIid requiredSku所属商品数字id，可通过 taobao.item.get 获取。必选
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // OuterId optionalSku的商家外部id
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // Price requiredSku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
    Price float32 `json:"price";xml:"price"`
    
    // Properties requiredSku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
    Properties string `json:"properties";xml:"properties"`
    
    // Quantity requiredSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
    Quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuAddRequest) GetAPIName() string {
	return "taobao.item.sku.add"
}

func (req *TaobaoItemSkuAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["Ignorewarning"] = req.Ignorewarning
    
    params["ItemPrice"] = req.ItemPrice
    
    params["Lang"] = req.Lang
    
    params["NumIid"] = req.NumIid
    
    params["OuterId"] = req.OuterId
    
    params["Price"] = req.Price
    
    params["Properties"] = req.Properties
    
    params["Quantity"] = req.Quantity
    
    return params
}

// TaobaoItemSkuAddResponse 新增一个sku到num_iid指定的商品中 
//传入的iid所对应的商品必须属于当前会话的用户
type TaobaoItemSkuAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Sku Objectsku
    Sku Sku `json:"sku";xml:"sku"`
    
}

// TaobaoItemPropimgUploadRequest 添加一张商品属性图片到num_iid指定的商品中 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
//商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 
//商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoItemPropimgUploadRequest struct {
    
    // Id optional属性图片ID。如果是新增不需要填写
    Id int64 `json:"id";xml:"id"`
    
    // Image optional属性图片内容。类型:JPG,GIF;图片大小不超过:3M
    Image []byte `json:"image";xml:"image"`
    
    // NumIid required商品数字ID，必选
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // Position optional图片位置
    Position int64 `json:"position";xml:"position"`
    
    // Properties required属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
    Properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemPropimgUploadRequest) GetAPIName() string {
	return "taobao.item.propimg.upload"
}

func (req *TaobaoItemPropimgUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["Id"] = req.Id
    
    params["Image"] = req.Image
    
    params["NumIid"] = req.NumIid
    
    params["Position"] = req.Position
    
    params["Properties"] = req.Properties
    
    return params
}

// TaobaoItemPropimgUploadResponse 添加一张商品属性图片到num_iid指定的商品中 
//传入的num_iid所对应的商品必须属于当前会话的用户 
//图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
//商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 
//商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoItemPropimgUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PropImg ObjectPropImg属性图片结构
    PropImg PropImg `json:"prop_img";xml:"prop_img"`
    
}

// TaobaoItemSkuUpdateRequest *更新一个sku的数据 
//*需要更新的sku通过属性properties进行匹配查找 
//*商品的数量和价格必须大于等于0 
//*sku记录会更新到指定的num_iid对应的商品中 
//*num_iid对应的商品必须属于当前的会话用户
type TaobaoItemSkuUpdateRequest struct {
    
    // Ignorewarning optional忽略警告提示.
    Ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    // ItemPrice optionalsku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
    ItemPrice float32 `json:"item_price";xml:"item_price"`
    
    // Lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    Lang string `json:"lang";xml:"lang"`
    
    // NumIid requiredSku所属商品数字id，可通过 taobao.item.get 获取
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // OuterId optionalSku的商家外部id
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // Price optionalSku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
    Price float32 `json:"price";xml:"price"`
    
    // Properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。
    // 如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号，
    Properties string `json:"properties";xml:"properties"`
    
    // Quantity optionalSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
    Quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuUpdateRequest) GetAPIName() string {
	return "taobao.item.sku.update"
}

func (req *TaobaoItemSkuUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["Ignorewarning"] = req.Ignorewarning
    
    params["ItemPrice"] = req.ItemPrice
    
    params["Lang"] = req.Lang
    
    params["NumIid"] = req.NumIid
    
    params["OuterId"] = req.OuterId
    
    params["Price"] = req.Price
    
    params["Properties"] = req.Properties
    
    params["Quantity"] = req.Quantity
    
    return params
}

// TaobaoItemSkuUpdateResponse *更新一个sku的数据 
//*需要更新的sku通过属性properties进行匹配查找 
//*商品的数量和价格必须大于等于0 
//*sku记录会更新到指定的num_iid对应的商品中 
//*num_iid对应的商品必须属于当前的会话用户
type TaobaoItemSkuUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Sku Object商品Sku
    Sku Sku `json:"sku";xml:"sku"`
    
}

// TaobaoItemSkuGetRequest 获取sku_id所对应的sku数据 
//sku_id对应的sku要属于传入的nick对应的卖家
type TaobaoItemSkuGetRequest struct {
    
    // Fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
    Fields string `json:"fields";xml:"fields"`
    
    // Nick optional卖家nick(num_iid和nick必传一个)，只传卖家nick时候，该api返回的结果不包含cspu（SKu上的产品规格信息）。
    Nick string `json:"nick";xml:"nick"`
    
    // NumIid optional商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // SkuId requiredSku的id。可以通过taobao.item.get得到
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoItemSkuGetRequest) GetAPIName() string {
	return "taobao.item.sku.get"
}

func (req *TaobaoItemSkuGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Fields"] = req.Fields
    
    params["Nick"] = req.Nick
    
    params["NumIid"] = req.NumIid
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoItemSkuGetResponse 获取sku_id所对应的sku数据 
//sku_id对应的sku要属于传入的nick对应的卖家
type TaobaoItemSkuGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Sku ObjectSku
    Sku Sku `json:"sku";xml:"sku"`
    
}

// TaobaoItemUpdateDelistingRequest * 单个商品下架
//    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingRequest struct {
    
    // NumIid required商品数字ID，该参数必须
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemUpdateDelistingRequest) GetAPIName() string {
	return "taobao.item.update.delisting"
}

func (req *TaobaoItemUpdateDelistingRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["NumIid"] = req.NumIid
    
    return params
}

// TaobaoItemUpdateDelistingResponse * 单个商品下架
//    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object返回商品更新信息：返回的结果是:num_iid和modified
    Item Item `json:"item";xml:"item"`
    
}

// TaobaoFenxiaoProductToChannelImportRequest 支持供应商将已有产品导入到某个渠道销售
type TaobaoFenxiaoProductToChannelImportRequest struct {
    
    // Channel required要导入的渠道[21 零售PLUS]目前仅支持此渠道
    Channel int64 `json:"channel";xml:"channel"`
    
    // ProductId required要导入的产品id
    ProductId int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TaobaoFenxiaoProductToChannelImportRequest) GetAPIName() string {
	return "taobao.fenxiao.product.to.channel.import"
}

func (req *TaobaoFenxiaoProductToChannelImportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Channel"] = req.Channel
    
    params["ProductId"] = req.ProductId
    
    return params
}

// TaobaoFenxiaoProductToChannelImportResponse 支持供应商将已有产品导入到某个渠道销售
type TaobaoFenxiaoProductToChannelImportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
}

// TaobaoItemSkusGetRequest * 获取多个商品下的所有sku
type TaobaoItemSkusGetRequest struct {
    
    // Fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
    Fields string `json:"fields";xml:"fields"`
    
    // NumIids requiredsku所属商品数字id，必选。num_iid个数不能超过40个
    NumIids string `json:"num_iids";xml:"num_iids"`
    
}

func (req *TaobaoItemSkusGetRequest) GetAPIName() string {
	return "taobao.item.skus.get"
}

func (req *TaobaoItemSkusGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["NumIids"] = req.NumIids
    
    return params
}

// TaobaoItemSkusGetResponse * 获取多个商品下的所有sku
type TaobaoItemSkusGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Skus Object ArraySku列表
    Skus Sku `json:"skus";xml:"skus"`
    
}

// TaobaoWlbItemBatchQueryRequest 根据用户id，item id list和store code来查询商品库存信息和批次信息
type TaobaoWlbItemBatchQueryRequest struct {
    
    // ItemIds required需要查询的商品ID列表，以字符串表示，ID间以;隔开
    ItemIds string `json:"item_ids";xml:"item_ids"`
    
    // PageNo optional分页查询参数，指定查询页数，默认为1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StoreCode optional仓库编号
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbItemBatchQueryRequest) GetAPIName() string {
	return "taobao.wlb.item.batch.query"
}

func (req *TaobaoWlbItemBatchQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ItemIds"] = req.ItemIds
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoWlbItemBatchQueryResponse 根据用户id，item id list和store code来查询商品库存信息和批次信息
type TaobaoWlbItemBatchQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemInventoryBatchList Object Array商品库存及批次信息查询结果
    ItemInventoryBatchList WlbItemBatchInventory `json:"item_inventory_batch_list";xml:"item_inventory_batch_list"`
    
    // TotalCount Basic返回结果记录的数量
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoStreetestSessionGetRequest 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaoStreetestSessionGetRequest struct {
    
}

func (req *TaobaoStreetestSessionGetRequest) GetAPIName() string {
	return "taobao.streetest.session.get"
}

func (req *TaobaoStreetestSessionGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoStreetestSessionGetResponse 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaoStreetestSessionGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // StreeTestSessionKey Basic压测账号对应的sessionKey
    StreeTestSessionKey string `json:"stree_test_session_key";xml:"stree_test_session_key"`
    
}

// TaobaoItemUpdateListingRequest * 单个商品上架
//* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingRequest struct {
    
    // Num required需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
    Num int64 `json:"num";xml:"num"`
    
    // NumIid required商品数字ID，该参数必须
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemUpdateListingRequest) GetAPIName() string {
	return "taobao.item.update.listing"
}

func (req *TaobaoItemUpdateListingRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Num"] = req.Num
    
    params["NumIid"] = req.NumIid
    
    return params
}

// TaobaoItemUpdateListingResponse * 单个商品上架
//* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object上架后返回的商品信息：返回的结果就是:num_iid和modified
    Item Item `json:"item";xml:"item"`
    
}

// CainiaoBimTradeorderConsignRequest 驱动保税交易订单发货
type CainiaoBimTradeorderConsignRequest struct {
    
    // ResId optional选择的线路ID非必填字段
    ResId string `json:"res_id";xml:"res_id"`
    
    // StoreCode optional仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TradeId required交易单号
    TradeId string `json:"trade_id";xml:"trade_id"`
    
}

func (req *CainiaoBimTradeorderConsignRequest) GetAPIName() string {
	return "cainiao.bim.tradeorder.consign"
}

func (req *CainiaoBimTradeorderConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ResId"] = req.ResId
    
    params["StoreCode"] = req.StoreCode
    
    params["TradeId"] = req.TradeId
    
    return params
}

// CainiaoBimTradeorderConsignResponse 驱动保税交易订单发货
type CainiaoBimTradeorderConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // LgOrderCode Basic菜鸟物流订单号,格式为LP开头
    LgOrderCode string `json:"lg_order_code";xml:"lg_order_code"`
    
    // StoreOrderCode Basic菜鸟仓库作业单据号
    StoreOrderCode string `json:"store_order_code";xml:"store_order_code"`
    
}

// TaobaoRpRefundReviewRequest 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
type TaobaoRpRefundReviewRequest struct {
    
    // Message required审核留言
    Message string `json:"message";xml:"message"`
    
    // Operator required审核人姓名
    Operator string `json:"operator";xml:"operator"`
    
    // RefundId required退款单编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase required退款阶段，可选值：售中：onsale，售后：aftersale
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
    // RefundVersion required退款最后更新时间，以时间戳的方式表示
    RefundVersion int64 `json:"refund_version";xml:"refund_version"`
    
    // Result required审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）
    Result bool `json:"result";xml:"result"`
    
}

func (req *TaobaoRpRefundReviewRequest) GetAPIName() string {
	return "taobao.rp.refund.review"
}

func (req *TaobaoRpRefundReviewRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Message"] = req.Message
    
    params["Operator"] = req.Operator
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    params["RefundVersion"] = req.RefundVersion
    
    params["Result"] = req.Result
    
    return params
}

// TaobaoRpRefundReviewResponse 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
type TaobaoRpRefundReviewResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basicsuccess
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoWlbWaybillICancelRequest 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaoWlbWaybillICancelRequest struct {
    
    // WaybillApplyCancelRequest required取消接口入参
    WaybillApplyCancelRequest WaybillApplyCancelRequest `json:"waybill_apply_cancel_request";xml:"waybill_apply_cancel_request"`
    
}

func (req *TaobaoWlbWaybillICancelRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.cancel"
}

func (req *TaobaoWlbWaybillICancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillApplyCancelRequest"] = req.WaybillApplyCancelRequest
    
    return params
}

// TaobaoWlbWaybillICancelResponse 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaoWlbWaybillICancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CancelResult Basic调用取消是否成功
    CancelResult bool `json:"cancel_result";xml:"cancel_result"`
    
}

// TaobaoWlbWaybillIQuerydetailRequest 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailRequest struct {
    
    // WaybillDetailQueryRequest required面单查询请求
    WaybillDetailQueryRequest WaybillDetailQueryRequest `json:"waybill_detail_query_request";xml:"waybill_detail_query_request"`
    
}

func (req *TaobaoWlbWaybillIQuerydetailRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.querydetail"
}

func (req *TaobaoWlbWaybillIQuerydetailRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillDetailQueryRequest"] = req.WaybillDetailQueryRequest
    
    return params
}

// TaobaoWlbWaybillIQuerydetailResponse 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorCodes Basic Array面单查询错误编码
    ErrorCodes string `json:"error_codes";xml:"error_codes"`
    
    // InexistentWaybillCodes Basic Array不存在的面单号
    InexistentWaybillCodes string `json:"inexistent_waybill_codes";xml:"inexistent_waybill_codes"`
    
    // QuerySuccess Basic查询是否成功
    QuerySuccess bool `json:"query_success";xml:"query_success"`
    
    // WaybillDetails Object Array面单详情
    WaybillDetails WaybillDetailQueryInfo `json:"waybill_details";xml:"waybill_details"`
    
}

// TaobaoWlbWaybillIPrintRequest 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaoWlbWaybillIPrintRequest struct {
    
    // WaybillApplyPrintCheckRequest required打印请求
    WaybillApplyPrintCheckRequest WaybillApplyPrintCheckRequest `json:"waybill_apply_print_check_request";xml:"waybill_apply_print_check_request"`
    
}

func (req *TaobaoWlbWaybillIPrintRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.print"
}

func (req *TaobaoWlbWaybillIPrintRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillApplyPrintCheckRequest"] = req.WaybillApplyPrintCheckRequest
    
    return params
}

// TaobaoWlbWaybillIPrintResponse 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaoWlbWaybillIPrintResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WaybillApplyPrintCheckInfos Object Array面单打印信息
    WaybillApplyPrintCheckInfos WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_infos";xml:"waybill_apply_print_check_infos"`
    
}

// TaobaoTmcGroupAddRequest 为已开通用户添加用户分组
type TaobaoTmcGroupAddRequest struct {
    
    // GroupName required分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。
    GroupName string `json:"group_name";xml:"group_name"`
    
    // Nicks required用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户
    Nicks string `json:"nicks";xml:"nicks"`
    
    // UserPlatform optional用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    UserPlatform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcGroupAddRequest) GetAPIName() string {
	return "taobao.tmc.group.add"
}

func (req *TaobaoTmcGroupAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["GroupName"] = req.GroupName
    
    params["Nicks"] = req.Nicks
    
    params["UserPlatform"] = req.UserPlatform
    
    return params
}

// TaobaoTmcGroupAddResponse 为已开通用户添加用户分组
type TaobaoTmcGroupAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic创建时间
    Created time.Time `json:"created";xml:"created"`
    
    // GroupName Basic分组名称
    GroupName string `json:"group_name";xml:"group_name"`
    
}

// TaobaoTradesSoldGetRequest 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
//<br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
//<br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
//<br/> <span style="color:red">注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）</span>
type TaobaoTradesSoldGetRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // EndCreated optional查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
    EndCreated time.Time `json:"end_created";xml:"end_created"`
    
    // ExtType optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
    ExtType string `json:"ext_type";xml:"ext_type"`
    
    // Fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0,处方药未审核
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。取值范围:大于零的整数; 默认值:1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // RateStatus optional评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评)
    RateStatus string `json:"rate_status";xml:"rate_status"`
    
    // StartCreated optional查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
    StartCreated time.Time `json:"start_created";xml:"start_created"`
    
    // Status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
    Status string `json:"status";xml:"status"`
    
    // Tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
    Tag string `json:"tag";xml:"tag"`
    
    // Type optional交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade,auto_delivery,ec,cod,step 这5 种的交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>可选值：fixed(一口价)auction(拍卖)guarantee_trade(一口价、拍卖)step(分阶段付款，万人团，阶梯团订单）independent_simple_trade(旺店入门版交易)independent_shop_trade(旺店标准版交易)auto_delivery(自动发货)ec(直冲)cod(货到付款)game_equipment(游戏装备)shopex_trade(ShopEX交易)netcn_trade(万网交易)external_trade(统一外部交易)instant_trade (即时到账)b2c_cod(大商家货到付款)hotel_trade(酒店类型交易)super_market_trade(商超交易)super_market_cod_trade(商超货到付款交易)taohua(淘花网交易类型）waimai(外卖交易类型）o2o_offlinetrade（O2O交易）nopaid（即时到帐/趣味猜交易类型）step (万人团) eticket(电子凭证) tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。pre_auth_type(预授权0元购) lazada（获取lazada订单类型）
    Type string `json:"type";xml:"type"`
    
    // UseHasNext optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。
    UseHasNext bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldGetRequest) GetAPIName() string {
	return "taobao.trades.sold.get"
}

func (req *TaobaoTradesSoldGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["EndCreated"] = req.EndCreated
    
    params["ExtType"] = req.ExtType
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["RateStatus"] = req.RateStatus
    
    params["StartCreated"] = req.StartCreated
    
    params["Status"] = req.Status
    
    params["Tag"] = req.Tag
    
    params["Type"] = req.Type
    
    params["UseHasNext"] = req.UseHasNext
    
    return params
}

// TaobaoTradesSoldGetResponse 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
//<br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
//<br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
//<br/> <span style="color:red">注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）</span>
type TaobaoTradesSoldGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // TotalResults Basic搜索到的交易信息总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
    // Trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    Trades Trade `json:"trades";xml:"trades"`
    
}

// TaobaoTmcGroupDeleteRequest 删除指定的分组或分组下的用户
type TaobaoTmcGroupDeleteRequest struct {
    
    // GroupName required分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
    GroupName string `json:"group_name";xml:"group_name"`
    
    // Nicks optional用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
    Nicks string `json:"nicks";xml:"nicks"`
    
    // UserPlatform optional用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    UserPlatform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcGroupDeleteRequest) GetAPIName() string {
	return "taobao.tmc.group.delete"
}

func (req *TaobaoTmcGroupDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["GroupName"] = req.GroupName
    
    params["Nicks"] = req.Nicks
    
    params["UserPlatform"] = req.UserPlatform
    
    return params
}

// TaobaoTmcGroupDeleteResponse 删除指定的分组或分组下的用户
type TaobaoTmcGroupDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoTmcGroupsGetRequest 获取自定义用户分组列表
type TaobaoTmcGroupsGetRequest struct {
    
    // GroupNames optional要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
    GroupNames string `json:"group_names";xml:"group_names"`
    
    // PageNo optional页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页返回多少个分组
    PageSize int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoTmcGroupsGetRequest) GetAPIName() string {
	return "taobao.tmc.groups.get"
}

func (req *TaobaoTmcGroupsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["GroupNames"] = req.GroupNames
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    return params
}

// TaobaoTmcGroupsGetResponse 获取自定义用户分组列表
type TaobaoTmcGroupsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Groups Object Arraydasdasd
    Groups TmcGroup `json:"groups";xml:"groups"`
    
    // TotalResults Basic分组总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoRpReturngoodsRefuseRequest 卖家拒绝退货，目前仅支持天猫退货。
type TaobaoRpReturngoodsRefuseRequest struct {
    
    // RefundId required退款编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase required退款服务状态，售后或者售中
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
    // RefundVersion required退款版本号
    RefundVersion int64 `json:"refund_version";xml:"refund_version"`
    
    // RefuseProof required拒绝退货凭证图片，必须图片格式，大小不能超过5M
    RefuseProof []byte `json:"refuse_proof";xml:"refuse_proof"`
    
    // RefuseReasonId optional拒绝原因编号，会提供拒绝原因列表供选择
    RefuseReasonId int64 `json:"refuse_reason_id";xml:"refuse_reason_id"`
    
}

func (req *TaobaoRpReturngoodsRefuseRequest) GetAPIName() string {
	return "taobao.rp.returngoods.refuse"
}

func (req *TaobaoRpReturngoodsRefuseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    params["RefundVersion"] = req.RefundVersion
    
    params["RefuseProof"] = req.RefuseProof
    
    params["RefuseReasonId"] = req.RefuseReasonId
    
    return params
}

// TaobaoRpReturngoodsRefuseResponse 卖家拒绝退货，目前仅支持天猫退货。
type TaobaoRpReturngoodsRefuseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basicasdf
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoRpReturngoodsRefillRequest 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillRequest struct {
    
    // LogisticsCompanyCode required物流公司编号
    LogisticsCompanyCode string `json:"logistics_company_code";xml:"logistics_company_code"`
    
    // LogisticsWaybillNo required物流公司运单号
    LogisticsWaybillNo string `json:"logistics_waybill_no";xml:"logistics_waybill_no"`
    
    // RefundId required退款单编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase required退款阶段，可选值：售中：onsale，售后：aftersale
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRpReturngoodsRefillRequest) GetAPIName() string {
	return "taobao.rp.returngoods.refill"
}

func (req *TaobaoRpReturngoodsRefillRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["LogisticsCompanyCode"] = req.LogisticsCompanyCode
    
    params["LogisticsWaybillNo"] = req.LogisticsWaybillNo
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    return params
}

// TaobaoRpReturngoodsRefillResponse 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic验货操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoRefundsApplyGetRequest 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
type TaobaoRefundsApplyGetRequest struct {
    
    // Fields required需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // SellerNick optional卖家昵称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
    // Status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。
    // WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) 
    // WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) 
    // WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) 
    // SELLER_REFUSE_BUYER(卖家拒绝退款) 
    // CLOSED(退款关闭) 
    // SUCCESS(退款成功)
    Status string `json:"status";xml:"status"`
    
    // Type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。
    // fixed(一口价) 
    // auction(拍卖) 
    // guarantee_trade(一口价、拍卖) 
    // independent_simple_trade(旺店入门版交易) 
    // independent_shop_trade(旺店标准版交易) 
    // auto_delivery(自动发货) 
    // ec(直冲) 
    // cod(货到付款) 
    // fenxiao(分销) 
    // game_equipment(游戏装备) 
    // shopex_trade(ShopEX交易) 
    // netcn_trade(万网交易) 
    // external_trade(统一外部交易)
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoRefundsApplyGetRequest) GetAPIName() string {
	return "taobao.refunds.apply.get"
}

func (req *TaobaoRefundsApplyGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["SellerNick"] = req.SellerNick
    
    params["Status"] = req.Status
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoRefundsApplyGetResponse 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
type TaobaoRefundsApplyGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Refunds Object Array搜索到的退款信息列表
    Refunds Refund `json:"refunds";xml:"refunds"`
    
    // TotalResults Basic搜索到的交易信息总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoTradeMemoUpdateRequest 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
type TaobaoTradeMemoUpdateRequest struct {
    
    // Flag optional卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    Flag int64 `json:"flag";xml:"flag"`
    
    // Memo optional卖家交易备注。最大长度: 1000个字节
    Memo string `json:"memo";xml:"memo"`
    
    // Reset optional是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
    Reset bool `json:"reset";xml:"reset"`
    
    // Tid required交易编号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeMemoUpdateRequest) GetAPIName() string {
	return "taobao.trade.memo.update"
}

func (req *TaobaoTradeMemoUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Flag"] = req.Flag
    
    params["Memo"] = req.Memo
    
    params["Reset"] = req.Reset
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeMemoUpdateResponse 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
type TaobaoTradeMemoUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Trade Object更新交易的备注信息后返回的Trade，其中可用字段为tid和modified
    Trade Trade `json:"trade";xml:"trade"`
    
}

// TaobaoTradeMemoAddRequest 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
type TaobaoTradeMemoAddRequest struct {
    
    // Flag optional交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    Flag int64 `json:"flag";xml:"flag"`
    
    // Memo required交易备注。最大长度: 1000个字节
    Memo string `json:"memo";xml:"memo"`
    
    // Tid required交易编号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeMemoAddRequest) GetAPIName() string {
	return "taobao.trade.memo.add"
}

func (req *TaobaoTradeMemoAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Flag"] = req.Flag
    
    params["Memo"] = req.Memo
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeMemoAddResponse 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
type TaobaoTradeMemoAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Trade Object对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created
    Trade Trade `json:"trade";xml:"trade"`
    
}

// TaobaoTraderatesGetRequest 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
type TaobaoTraderatesGetRequest struct {
    
    // EndDate optional评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // Fields required需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
    Fields string `json:"fields";xml:"fields"`
    
    // NumIid optional商品的数字ID
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // PageNo optional页码。取值范围:大于零的整数最大限制为200; 默认值:1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页获取条数。默认值40，最小值1，最大值150。
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // PeerNick optional评价对方昵称
    PeerNick string `json:"peer_nick";xml:"peer_nick"`
    
    // RateType required评价类型。可选值:get(得到),give(给出)
    RateType string `json:"rate_type";xml:"rate_type"`
    
    // Result optional评价结果。可选值:good(好评),neutral(中评),bad(差评)
    Result string `json:"result";xml:"result"`
    
    // Role required评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
    Role string `json:"role";xml:"role"`
    
    // StartDate optional评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // Tid optional交易订单id，可以是父订单id号，也可以是子订单id号
    Tid int64 `json:"tid";xml:"tid"`
    
    // UseHasNext optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
    UseHasNext bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTraderatesGetRequest) GetAPIName() string {
	return "taobao.traderates.get"
}

func (req *TaobaoTraderatesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["EndDate"] = req.EndDate
    
    params["Fields"] = req.Fields
    
    params["NumIid"] = req.NumIid
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["PeerNick"] = req.PeerNick
    
    params["RateType"] = req.RateType
    
    params["Result"] = req.Result
    
    params["Role"] = req.Role
    
    params["StartDate"] = req.StartDate
    
    params["Tid"] = req.Tid
    
    params["UseHasNext"] = req.UseHasNext
    
    return params
}

// TaobaoTraderatesGetResponse 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
type TaobaoTraderatesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic当使用use_has_next时返回信息，如果还有下一页则返回true
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // TotalResults Basic搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
    // TradeRates Object Array评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息
    TradeRates TradeRate `json:"trade_rates";xml:"trade_rates"`
    
}

// TaobaoTradeFullinfoGetRequest 获取单笔交易的详细信息
//<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值 
//<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
//<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
//<br/>4. 获取红包优惠金额可以使用字段 coupon_fee
//<br/>5. 请按需获取字段，减少TOP系统的压力
type TaobaoTradeFullinfoGetRequest struct {
    
    // Fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
    Fields string `json:"fields";xml:"fields"`
    
    // Tid required交易编号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeFullinfoGetRequest) GetAPIName() string {
	return "taobao.trade.fullinfo.get"
}

func (req *TaobaoTradeFullinfoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeFullinfoGetResponse 获取单笔交易的详细信息
//<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值 
//<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息 
//<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
//<br/>4. 获取红包优惠金额可以使用字段 coupon_fee
//<br/>5. 请按需获取字段，减少TOP系统的压力
type TaobaoTradeFullinfoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Trade Object交易主订单信息
    Trade Trade `json:"trade";xml:"trade"`
    
}

// AlibabaProviderEinvoiceQueryRequest 开票服务商提供的对已开电子发票的查询功能，此接口是可选实现。只有部分电子发票支持通过此接口查询
type AlibabaProviderEinvoiceQueryRequest struct {
    
    // PayeeRegisterNo required收款方税务登记证号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // PlatformCode optional电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    PlatformCode string `json:"platform_code";xml:"platform_code"`
    
    // PlatformTid optional电商平台对应的订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
    // SerialNo special开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。与erp_tid二选一
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaProviderEinvoiceQueryRequest) GetAPIName() string {
	return "alibaba.provider.einvoice.query"
}

func (req *AlibabaProviderEinvoiceQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["PlatformCode"] = req.PlatformCode
    
    params["PlatformTid"] = req.PlatformTid
    
    params["SerialNo"] = req.SerialNo
    
    return params
}

// AlibabaProviderEinvoiceQueryResponse 开票服务商提供的对已开电子发票的查询功能，此接口是可选实现。只有部分电子发票支持通过此接口查询
type AlibabaProviderEinvoiceQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // EinvoiceList Object Array电子发票列表
    EinvoiceList ProviderEinvoice `json:"einvoice_list";xml:"einvoice_list"`
    
}

// TaobaoRefundGetRequest 获取单笔退款详情
type TaobaoRefundGetRequest struct {
    
    // Fields required需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
    Fields string `json:"fields";xml:"fields"`
    
    // RefundId required退款单号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
}

func (req *TaobaoRefundGetRequest) GetAPIName() string {
	return "taobao.refund.get"
}

func (req *TaobaoRefundGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["RefundId"] = req.RefundId
    
    return params
}

// TaobaoRefundGetResponse 获取单笔退款详情
type TaobaoRefundGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Refund Object退款详情
    Refund Refund `json:"refund";xml:"refund"`
    
}

// TaobaoRefundsReceiveGetRequest 查询卖家收到的退款列表
type TaobaoRefundsReceiveGetRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // EndModified optional查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // Fields required需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。取值范围:大于零的整数; 默认值:1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartModified optional查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
    // Status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
    Status string `json:"status";xml:"status"`
    
    // Type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>
    Type string `json:"type";xml:"type"`
    
    // UseHasNext optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
    UseHasNext bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoRefundsReceiveGetRequest) GetAPIName() string {
	return "taobao.refunds.receive.get"
}

func (req *TaobaoRefundsReceiveGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["EndModified"] = req.EndModified
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartModified"] = req.StartModified
    
    params["Status"] = req.Status
    
    params["Type"] = req.Type
    
    params["UseHasNext"] = req.UseHasNext
    
    return params
}

// TaobaoRefundsReceiveGetResponse 查询卖家收到的退款列表
type TaobaoRefundsReceiveGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // Refunds Object Array搜索到的退款信息列表
    Refunds Refund `json:"refunds";xml:"refunds"`
    
    // TotalResults Basic搜索到的退款信息总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoAreasGetRequest 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
//<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
type TaobaoAreasGetRequest struct {
    
    // Fields required需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoAreasGetRequest) GetAPIName() string {
	return "taobao.areas.get"
}

func (req *TaobaoAreasGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoAreasGetResponse 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
//<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
type TaobaoAreasGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Areas Object Array地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。
    Areas Area `json:"areas";xml:"areas"`
    
}

// TaobaoFenxiaoProductSkusGetRequest 产品sku查询
type TaobaoFenxiaoProductSkusGetRequest struct {
    
    // ProductId required产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TaobaoFenxiaoProductSkusGetRequest) GetAPIName() string {
	return "taobao.fenxiao.product.skus.get"
}

func (req *TaobaoFenxiaoProductSkusGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ProductId"] = req.ProductId
    
    return params
}

// TaobaoFenxiaoProductSkusGetResponse 产品sku查询
type TaobaoFenxiaoProductSkusGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Skus Object Arraysku信息
    Skus FenxiaoSku `json:"skus";xml:"skus"`
    
    // TotalResults Basic记录数量
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// CainiaoBmsOrderCreateRequest 通过接口，在菜鸟商家工作台创建订单，并通过商家工作台进行发货管理。
type CainiaoBmsOrderCreateRequest struct {
    
    // BuyerMessage optional买家留言
    BuyerMessage string `json:"buyer_message";xml:"buyer_message"`
    
    // BuyerNick optional买家名称"
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // CodFee optionalCOD服务费
    CodFee int64 `json:"cod_fee";xml:"cod_fee"`
    
    // Created required创建时间
    Created time.Time `json:"created";xml:"created"`
    
    // DiscountFee optional优惠金额，单位分
    DiscountFee int64 `json:"discount_fee";xml:"discount_fee"`
    
    // InvoiceAmount optional发票金额，单位分
    InvoiceAmount int64 `json:"invoice_amount";xml:"invoice_amount"`
    
    // InvoiceTitle optional发票抬头
    InvoiceTitle string `json:"invoice_title";xml:"invoice_title"`
    
    // InvoiceType optional发票类型(1 电子、2 纸质)
    InvoiceType int64 `json:"invoice_type";xml:"invoice_type"`
    
    // IsCod optional是否COD
    IsCod bool `json:"is_cod";xml:"is_cod"`
    
    // IsInvoice optional是否打印发票
    IsInvoice bool `json:"is_invoice";xml:"is_invoice"`
    
    // Items optionaldemo
    Items Items `json:"items";xml:"items"`
    
    // OrderAmount required订单总金额，单位分
    OrderAmount int64 `json:"order_amount";xml:"order_amount"`
    
    // OrderType optional订单类型，默认0，普通销售订单，3:B2B单
    OrderType string `json:"order_type";xml:"order_type"`
    
    // PaiedAmount required支付金额，单位分
    PaiedAmount int64 `json:"paied_amount";xml:"paied_amount"`
    
    // PayTime required支付时间
    PayTime time.Time `json:"pay_time";xml:"pay_time"`
    
    // PostFee optional快递费，单位分
    PostFee int64 `json:"post_fee";xml:"post_fee"`
    
    // ReceiverAddress required收货地址
    ReceiverAddress string `json:"receiver_address";xml:"receiver_address"`
    
    // ReceiverCity required收货市
    ReceiverCity string `json:"receiver_city";xml:"receiver_city"`
    
    // ReceiverCountry optional收货人国籍
    ReceiverCountry string `json:"receiver_country";xml:"receiver_country"`
    
    // ReceiverDistrict optional收货区
    ReceiverDistrict string `json:"receiver_district";xml:"receiver_district"`
    
    // ReceiverMobile optional收货人手机，手机与电话不可同时为空
    ReceiverMobile string `json:"receiver_mobile";xml:"receiver_mobile"`
    
    // ReceiverName required收货人名称
    ReceiverName string `json:"receiver_name";xml:"receiver_name"`
    
    // ReceiverPhone optional收货人电话，手机与电话不可同时为空
    ReceiverPhone string `json:"receiver_phone";xml:"receiver_phone"`
    
    // ReceiverState required收货省
    ReceiverState string `json:"receiver_state";xml:"receiver_state"`
    
    // ReceiverTown optional街道
    ReceiverTown string `json:"receiver_town";xml:"receiver_town"`
    
    // ReceiverZip required收货邮编
    ReceiverZip string `json:"receiver_zip";xml:"receiver_zip"`
    
    // SellerMemo optional卖家备注
    SellerMemo string `json:"seller_memo";xml:"seller_memo"`
    
    // ShopCode required店铺信息
    ShopCode string `json:"shop_code";xml:"shop_code"`
    
    // TradeId required交易单号
    TradeId string `json:"trade_id";xml:"trade_id"`
    
    // WaitPayAmount required应收金额，单位分
    WaitPayAmount int64 `json:"wait_pay_amount";xml:"wait_pay_amount"`
    
}

func (req *CainiaoBmsOrderCreateRequest) GetAPIName() string {
	return "cainiao.bms.order.create"
}

func (req *CainiaoBmsOrderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 30)
    
    params["BuyerMessage"] = req.BuyerMessage
    
    params["BuyerNick"] = req.BuyerNick
    
    params["CodFee"] = req.CodFee
    
    params["Created"] = req.Created
    
    params["DiscountFee"] = req.DiscountFee
    
    params["InvoiceAmount"] = req.InvoiceAmount
    
    params["InvoiceTitle"] = req.InvoiceTitle
    
    params["InvoiceType"] = req.InvoiceType
    
    params["IsCod"] = req.IsCod
    
    params["IsInvoice"] = req.IsInvoice
    
    params["Items"] = req.Items
    
    params["OrderAmount"] = req.OrderAmount
    
    params["OrderType"] = req.OrderType
    
    params["PaiedAmount"] = req.PaiedAmount
    
    params["PayTime"] = req.PayTime
    
    params["PostFee"] = req.PostFee
    
    params["ReceiverAddress"] = req.ReceiverAddress
    
    params["ReceiverCity"] = req.ReceiverCity
    
    params["ReceiverCountry"] = req.ReceiverCountry
    
    params["ReceiverDistrict"] = req.ReceiverDistrict
    
    params["ReceiverMobile"] = req.ReceiverMobile
    
    params["ReceiverName"] = req.ReceiverName
    
    params["ReceiverPhone"] = req.ReceiverPhone
    
    params["ReceiverState"] = req.ReceiverState
    
    params["ReceiverTown"] = req.ReceiverTown
    
    params["ReceiverZip"] = req.ReceiverZip
    
    params["SellerMemo"] = req.SellerMemo
    
    params["ShopCode"] = req.ShopCode
    
    params["TradeId"] = req.TradeId
    
    params["WaitPayAmount"] = req.WaitPayAmount
    
    return params
}

// CainiaoBmsOrderCreateResponse 通过接口，在菜鸟商家工作台创建订单，并通过商家工作台进行发货管理。
type CainiaoBmsOrderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
}

// TaobaoFenxiaoProductSkuUpdateRequest 产品SKU信息更新
type TaobaoFenxiaoProductSkuUpdateRequest struct {
    
    // AgentCostPrice optional代销采购价
    AgentCostPrice string `json:"agent_cost_price";xml:"agent_cost_price"`
    
    // DealerCostPrice optional经销采购价
    DealerCostPrice string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    // ProductId required产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Properties requiredsku属性
    Properties string `json:"properties";xml:"properties"`
    
    // Quantity optional产品SKU库存
    Quantity int64 `json:"quantity";xml:"quantity"`
    
    // SkuNumber optional商家编码
    SkuNumber string `json:"sku_number";xml:"sku_number"`
    
    // StandardPrice optional采购基准价
    StandardPrice string `json:"standard_price";xml:"standard_price"`
    
}

func (req *TaobaoFenxiaoProductSkuUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.update"
}

func (req *TaobaoFenxiaoProductSkuUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["AgentCostPrice"] = req.AgentCostPrice
    
    params["DealerCostPrice"] = req.DealerCostPrice
    
    params["ProductId"] = req.ProductId
    
    params["Properties"] = req.Properties
    
    params["Quantity"] = req.Quantity
    
    params["SkuNumber"] = req.SkuNumber
    
    params["StandardPrice"] = req.StandardPrice
    
    return params
}

// TaobaoFenxiaoProductSkuUpdateResponse 产品SKU信息更新
type TaobaoFenxiaoProductSkuUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic操作时间
    Created time.Time `json:"created";xml:"created"`
    
    // Result Basic操作结果
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoProductSkuAddRequest 添加产品SKU信息
type TaobaoFenxiaoProductSkuAddRequest struct {
    
    // AgentCostPrice special代销采购价
    AgentCostPrice string `json:"agent_cost_price";xml:"agent_cost_price"`
    
    // DealerCostPrice special经销采购价
    DealerCostPrice string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    // ProductId required产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Properties requiredsku属性
    Properties string `json:"properties";xml:"properties"`
    
    // Quantity optionalsku产品库存，在0到1000000之间，如果不传，则库存为0
    Quantity int64 `json:"quantity";xml:"quantity"`
    
    // SkuNumber optional商家编码
    SkuNumber string `json:"sku_number";xml:"sku_number"`
    
    // StandardPrice required采购基准价，最大值1000000000
    StandardPrice string `json:"standard_price";xml:"standard_price"`
    
}

func (req *TaobaoFenxiaoProductSkuAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.add"
}

func (req *TaobaoFenxiaoProductSkuAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["AgentCostPrice"] = req.AgentCostPrice
    
    params["DealerCostPrice"] = req.DealerCostPrice
    
    params["ProductId"] = req.ProductId
    
    params["Properties"] = req.Properties
    
    params["Quantity"] = req.Quantity
    
    params["SkuNumber"] = req.SkuNumber
    
    params["StandardPrice"] = req.StandardPrice
    
    return params
}

// TaobaoFenxiaoProductSkuAddResponse 添加产品SKU信息
type TaobaoFenxiaoProductSkuAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic操作时间
    Created time.Time `json:"created";xml:"created"`
    
    // Result Basic操作结果
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoProductSkuDeleteRequest 根据sku properties删除sku数据
type TaobaoFenxiaoProductSkuDeleteRequest struct {
    
    // ProductId required产品id
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Properties requiredsku属性
    Properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductSkuDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.product.sku.delete"
}

func (req *TaobaoFenxiaoProductSkuDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ProductId"] = req.ProductId
    
    params["Properties"] = req.Properties
    
    return params
}

// TaobaoFenxiaoProductSkuDeleteResponse 根据sku properties删除sku数据
type TaobaoFenxiaoProductSkuDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic操作时间
    Created time.Time `json:"created";xml:"created"`
    
    // Result Basic操作结果
    Result bool `json:"result";xml:"result"`
    
}

// QimenTaobaoAutoReturnorderConfirmRequest 菜鸟自动创建销退单的入库确认回传
type QimenTaobaoAutoReturnorderConfirmRequest struct {
    
    // CustomerId requiredcustomerId
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // Request optional
    Request AutoReturnInOrderConfirmRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoAutoReturnorderConfirmRequest) GetAPIName() string {
	return "qimen.taobao.auto.returnorder.confirm"
}

func (req *QimenTaobaoAutoReturnorderConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoAutoReturnorderConfirmResponse 菜鸟自动创建销退单的入库确认回传
type QimenTaobaoAutoReturnorderConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Objectresponse
    Response Response `json:"response";xml:"response"`
    
}

// TaobaoFenxiaoProductImageDeleteRequest 产品图片删除，只删除图片信息，不真正删除图片
type TaobaoFenxiaoProductImageDeleteRequest struct {
    
    // Position required图片位置
    Position int64 `json:"position";xml:"position"`
    
    // ProductId required产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Properties optionalproperties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    Properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductImageDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.product.image.delete"
}

func (req *TaobaoFenxiaoProductImageDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Position"] = req.Position
    
    params["ProductId"] = req.ProductId
    
    params["Properties"] = req.Properties
    
    return params
}

// TaobaoFenxiaoProductImageDeleteResponse 产品图片删除，只删除图片信息，不真正删除图片
type TaobaoFenxiaoProductImageDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic操作时间
    Created time.Time `json:"created";xml:"created"`
    
    // Result Basic操作结果
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoWlbWaybillIProductRequest 商家可以查询物流商的产品类型和服务能力。
type TaobaoWlbWaybillIProductRequest struct {
    
    // WaybillProductTypeRequest required查询物流商电子面单产品类型入参
    WaybillProductTypeRequest WaybillProductTypeRequest `json:"waybill_product_type_request";xml:"waybill_product_type_request"`
    
}

func (req *TaobaoWlbWaybillIProductRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.product"
}

func (req *TaobaoWlbWaybillIProductRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillProductTypeRequest"] = req.WaybillProductTypeRequest
    
    return params
}

// TaobaoWlbWaybillIProductResponse 商家可以查询物流商的产品类型和服务能力。
type TaobaoWlbWaybillIProductResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ProductTypes Object Array产品类型返回
    ProductTypes WaybillProductType `json:"product_types";xml:"product_types"`
    
}

// TaobaoFenxiaoProductImageUploadRequest 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaoFenxiaoProductImageUploadRequest struct {
    
    // Image special产品图片
    Image []byte `json:"image";xml:"image"`
    
    // PicPath special产品主图图片空间相对路径或绝对路径
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // Position required图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
    Position int64 `json:"position";xml:"position"`
    
    // ProductId required产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Properties optionalproperties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    Properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoFenxiaoProductImageUploadRequest) GetAPIName() string {
	return "taobao.fenxiao.product.image.upload"
}

func (req *TaobaoFenxiaoProductImageUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["Image"] = req.Image
    
    params["PicPath"] = req.PicPath
    
    params["Position"] = req.Position
    
    params["ProductId"] = req.ProductId
    
    params["Properties"] = req.Properties
    
    return params
}

// TaobaoFenxiaoProductImageUploadResponse 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaoFenxiaoProductImageUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic操作时间
    Created time.Time `json:"created";xml:"created"`
    
    // Result Basic操作是否成功
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoShopGetRequest 获取卖家店铺的基本信息
type TaobaoShopGetRequest struct {
    
    // Fields required需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
    Fields string `json:"fields";xml:"fields"`
    
    // Nick required卖家昵称
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoShopGetRequest) GetAPIName() string {
	return "taobao.shop.get"
}

func (req *TaobaoShopGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoShopGetResponse 获取卖家店铺的基本信息
type TaobaoShopGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shop Object店铺信息
    Shop Shop `json:"shop";xml:"shop"`
    
}

// TaobaoJdsTradesStatisticsDiffRequest 订单全链路状态统计差异比较
type TaobaoJdsTradesStatisticsDiffRequest struct {
    
    // Date required查询的日期，格式如YYYYMMDD的日期对应的数字
    Date int64 `json:"date";xml:"date"`
    
    // PageNo optional页数
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PostStatus required需要比较的状态
    PostStatus string `json:"post_status";xml:"post_status"`
    
    // PreStatus required需要比较的状态，将会和post_status做比较
    PreStatus string `json:"pre_status";xml:"pre_status"`
    
}

func (req *TaobaoJdsTradesStatisticsDiffRequest) GetAPIName() string {
	return "taobao.jds.trades.statistics.diff"
}

func (req *TaobaoJdsTradesStatisticsDiffRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Date"] = req.Date
    
    params["PageNo"] = req.PageNo
    
    params["PostStatus"] = req.PostStatus
    
    params["PreStatus"] = req.PreStatus
    
    return params
}

// TaobaoJdsTradesStatisticsDiffResponse 订单全链路状态统计差异比较
type TaobaoJdsTradesStatisticsDiffResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Tids Basic Arraypre_status比post_status多的tid列表
    Tids int64 `json:"tids";xml:"tids"`
    
    // TotalResults Basic总记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoShopcatsListGetRequest 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoShopcatsListGetRequest struct {
    
    // Fields optional需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoShopcatsListGetRequest) GetAPIName() string {
	return "taobao.shopcats.list.get"
}

func (req *TaobaoShopcatsListGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoShopcatsListGetResponse 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoShopcatsListGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ShopCats Object Array店铺类目列表信息
    ShopCats ShopCat `json:"shop_cats";xml:"shop_cats"`
    
}

// TaobaoJdsTradesStatisticsGetRequest 获取订单数量统计结果
type TaobaoJdsTradesStatisticsGetRequest struct {
    
    // Date required查询的日期，格式如YYYYMMDD的日期对应的数字
    Date int64 `json:"date";xml:"date"`
    
}

func (req *TaobaoJdsTradesStatisticsGetRequest) GetAPIName() string {
	return "taobao.jds.trades.statistics.get"
}

func (req *TaobaoJdsTradesStatisticsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Date"] = req.Date
    
    return params
}

// TaobaoJdsTradesStatisticsGetResponse 获取订单数量统计结果
type TaobaoJdsTradesStatisticsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // StatusInfos Object Array订单状态计数值
    StatusInfos TradeStat `json:"status_infos";xml:"status_infos"`
    
}

// TaobaoSellercatsListGetRequest 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaoSellercatsListGetRequest struct {
    
    // Fields optionalfields参数
    Fields string `json:"fields";xml:"fields"`
    
    // Nick required卖家昵称
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercatsListGetRequest) GetAPIName() string {
	return "taobao.sellercats.list.get"
}

func (req *TaobaoSellercatsListGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoSellercatsListGetResponse 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaoSellercatsListGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SellerCats Object Array卖家自定义类目
    SellerCats SellerCat `json:"seller_cats";xml:"seller_cats"`
    
}

// TmallItemHscodeDetailGetRequest 通过hscode获取计量单位和销售单位
type TmallItemHscodeDetailGetRequest struct {
    
    // Hscode requiredhscode
    Hscode string `json:"hscode";xml:"hscode"`
    
}

func (req *TmallItemHscodeDetailGetRequest) GetAPIName() string {
	return "tmall.item.hscode.detail.get"
}

func (req *TmallItemHscodeDetailGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Hscode"] = req.Hscode
    
    return params
}

// TmallItemHscodeDetailGetResponse 通过hscode获取计量单位和销售单位
type TmallItemHscodeDetailGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Results Basic Array返回的计量单位和销售单位
    Results map[string]interface{} `json:"results";xml:"results"`
    
}

// TaobaoJdsRefundTracesGetRequest 获取聚石塔数据共享的交易全链路的退款信息
type TaobaoJdsRefundTracesGetRequest struct {
    
    // RefundId required淘宝的退款编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // ReturnUserStatus optional是否返回用户状态(是否存在)
    ReturnUserStatus bool `json:"return_user_status";xml:"return_user_status"`
    
}

func (req *TaobaoJdsRefundTracesGetRequest) GetAPIName() string {
	return "taobao.jds.refund.traces.get"
}

func (req *TaobaoJdsRefundTracesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["RefundId"] = req.RefundId
    
    params["ReturnUserStatus"] = req.ReturnUserStatus
    
    return params
}

// TaobaoJdsRefundTracesGetResponse 获取聚石塔数据共享的交易全链路的退款信息
type TaobaoJdsRefundTracesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Traces Object Array退款跟踪列表
    Traces RefundTrace `json:"traces";xml:"traces"`
    
    // UserStatus Basic用户在全链路系统中的状态(比如是否存在)
    UserStatus string `json:"user_status";xml:"user_status"`
    
}

// TaobaoJdsHluserGetRequest 订单全链路用户信息获取
type TaobaoJdsHluserGetRequest struct {
    
}

func (req *TaobaoJdsHluserGetRequest) GetAPIName() string {
	return "taobao.jds.hluser.get"
}

func (req *TaobaoJdsHluserGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoJdsHluserGetResponse 订单全链路用户信息获取
type TaobaoJdsHluserGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HlUser Object回流用户信息
    HlUser HlUserDO `json:"hl_user";xml:"hl_user"`
    
}

// TaobaoItemSkuPriceUpdateRequest 更新商品SKU的价格
type TaobaoItemSkuPriceUpdateRequest struct {
    
    // Ignorewarning optional忽略警告提示.
    Ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    // ItemPrice optionalsku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
    ItemPrice float32 `json:"item_price";xml:"item_price"`
    
    // Lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    Lang string `json:"lang";xml:"lang"`
    
    // NumIid requiredSku所属商品数字id，可通过 taobao.item.get 获取
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // OuterId optionalSku的商家外部id
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // Price optionalSku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
    Price float32 `json:"price";xml:"price"`
    
    // Properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
    Properties string `json:"properties";xml:"properties"`
    
    // Quantity optionalSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
    Quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuPriceUpdateRequest) GetAPIName() string {
	return "taobao.item.sku.price.update"
}

func (req *TaobaoItemSkuPriceUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["Ignorewarning"] = req.Ignorewarning
    
    params["ItemPrice"] = req.ItemPrice
    
    params["Lang"] = req.Lang
    
    params["NumIid"] = req.NumIid
    
    params["OuterId"] = req.OuterId
    
    params["Price"] = req.Price
    
    params["Properties"] = req.Properties
    
    params["Quantity"] = req.Quantity
    
    return params
}

// TaobaoItemSkuPriceUpdateResponse 更新商品SKU的价格
type TaobaoItemSkuPriceUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Sku Object商品SKU信息（只包含num_iid和modified）
    Sku Sku `json:"sku";xml:"sku"`
    
}

// TaobaoJdsHluserUpdateRequest 订单全链路用户信息修改，比如是否开放买家端展示
type TaobaoJdsHluserUpdateRequest struct {
    
    // OpenForBuyer optional回流信息是否开通买家端展示,可选值open,close
    OpenForBuyer string `json:"open_for_buyer";xml:"open_for_buyer"`
    
    // OpenNodes optional为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
    OpenNodes string `json:"open_nodes";xml:"open_nodes"`
    
}

func (req *TaobaoJdsHluserUpdateRequest) GetAPIName() string {
	return "taobao.jds.hluser.update"
}

func (req *TaobaoJdsHluserUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["OpenForBuyer"] = req.OpenForBuyer
    
    params["OpenNodes"] = req.OpenNodes
    
    return params
}

// TaobaoJdsHluserUpdateResponse 订单全链路用户信息修改，比如是否开放买家端展示
type TaobaoJdsHluserUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic是否成功
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoItemBarcodeUpdateRequest 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoItemBarcodeUpdateRequest struct {
    
    // Isforce optional是否强制保存商品条码。
    // true：强制保存
    // false ：需要执行条码库校验
    Isforce bool `json:"isforce";xml:"isforce"`
    
    // ItemBarcode optional商品条形码，如果不用更新，可选择不填
    ItemBarcode string `json:"item_barcode";xml:"item_barcode"`
    
    // ItemId required被更新商品的ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // SkuBarcodes optionalSKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
    SkuBarcodes string `json:"sku_barcodes";xml:"sku_barcodes"`
    
    // SkuIds optional被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
    SkuIds string `json:"sku_ids";xml:"sku_ids"`
    
    // Src optional访问来源，这字段提供给千牛扫码枪用，
    // 其他调用方，不需要填写
    Src string `json:"src";xml:"src"`
    
}

func (req *TaobaoItemBarcodeUpdateRequest) GetAPIName() string {
	return "taobao.item.barcode.update"
}

func (req *TaobaoItemBarcodeUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Isforce"] = req.Isforce
    
    params["ItemBarcode"] = req.ItemBarcode
    
    params["ItemId"] = req.ItemId
    
    params["SkuBarcodes"] = req.SkuBarcodes
    
    params["SkuIds"] = req.SkuIds
    
    params["Src"] = req.Src
    
    return params
}

// TaobaoItemBarcodeUpdateResponse 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoItemBarcodeUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object商品结构里的num_iid，modified
    Item Item `json:"item";xml:"item"`
    
}

// TaobaoWlbWaybillIGetRequest 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaoWlbWaybillIGetRequest struct {
    
    // WaybillApplyNewRequest required面单申请
    WaybillApplyNewRequest WaybillApplyNewRequest `json:"waybill_apply_new_request";xml:"waybill_apply_new_request"`
    
}

func (req *TaobaoWlbWaybillIGetRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.get"
}

func (req *TaobaoWlbWaybillIGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillApplyNewRequest"] = req.WaybillApplyNewRequest
    
    return params
}

// TaobaoWlbWaybillIGetResponse 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaoWlbWaybillIGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WaybillApplyNewCols Object Array面单申请接口返回信息
    WaybillApplyNewCols WaybillApplyNewInfo `json:"waybill_apply_new_cols";xml:"waybill_apply_new_cols"`
    
}

// TaobaoWlbWaybillISearchRequest 获取发货地&CP开通状态&账户的使用情况
type TaobaoWlbWaybillISearchRequest struct {
    
    // WaybillApplyRequest required查询网点信息
    WaybillApplyRequest WaybillApplyRequest `json:"waybill_apply_request";xml:"waybill_apply_request"`
    
}

func (req *TaobaoWlbWaybillISearchRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.search"
}

func (req *TaobaoWlbWaybillISearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillApplyRequest"] = req.WaybillApplyRequest
    
    return params
}

// TaobaoWlbWaybillISearchResponse 获取发货地&CP开通状态&账户的使用情况
type TaobaoWlbWaybillISearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Subscribtions Object Array订购关系
    Subscribtions WaybillApplySubscriptionInfo `json:"subscribtions";xml:"subscribtions"`
    
}

// TaobaoRegionWarehouseManageRequest 编辑仓库覆盖范围
type TaobaoRegionWarehouseManageRequest struct {
    
    // Regions required可映射三级地址,例: 广东省
    Regions string `json:"regions";xml:"regions"`
    
    // StoreCode required仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoRegionWarehouseManageRequest) GetAPIName() string {
	return "taobao.region.warehouse.manage"
}

func (req *TaobaoRegionWarehouseManageRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Regions"] = req.Regions
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoRegionWarehouseManageResponse 编辑仓库覆盖范围
type TaobaoRegionWarehouseManageResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result BaseResult `json:"result";xml:"result"`
    
}

// TaobaoWlbWaybillIFullupdateRequest 商家更新电子面单号对应的订单信息。
//<br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；
//<br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
type TaobaoWlbWaybillIFullupdateRequest struct {
    
    // WaybillApplyFullUpdateRequest required更新面单信息请求
    WaybillApplyFullUpdateRequest WaybillApplyFullUpdateRequest `json:"waybill_apply_full_update_request";xml:"waybill_apply_full_update_request"`
    
}

func (req *TaobaoWlbWaybillIFullupdateRequest) GetAPIName() string {
	return "taobao.wlb.waybill.i.fullupdate"
}

func (req *TaobaoWlbWaybillIFullupdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WaybillApplyFullUpdateRequest"] = req.WaybillApplyFullUpdateRequest
    
    return params
}

// TaobaoWlbWaybillIFullupdateResponse 商家更新电子面单号对应的订单信息。
//<br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；
//<br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
type TaobaoWlbWaybillIFullupdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WaybillApplyUpdateInfo Object更新接口出参
    WaybillApplyUpdateInfo WaybillApplyUpdateInfo `json:"waybill_apply_update_info";xml:"waybill_apply_update_info"`
    
}

// TaobaoRegionSaleQueryRequest 查询商品销售区域
type TaobaoRegionSaleQueryRequest struct {
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // SaleRegionLevel required1:国家;2:省;3: 市;4:区县
    SaleRegionLevel int64 `json:"sale_region_level";xml:"sale_region_level"`
    
    // SkuId required无sku传0
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionSaleQueryRequest) GetAPIName() string {
	return "taobao.region.sale.query"
}

func (req *TaobaoRegionSaleQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ItemId"] = req.ItemId
    
    params["SaleRegionLevel"] = req.SaleRegionLevel
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoRegionSaleQueryResponse 查询商品销售区域
type TaobaoRegionSaleQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result BaseResult `json:"result";xml:"result"`
    
}

// TaobaoPictureIsreferencedGetRequest 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaoPictureIsreferencedGetRequest struct {
    
    // PictureId required图片id
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureIsreferencedGetRequest) GetAPIName() string {
	return "taobao.picture.isreferenced.get"
}

func (req *TaobaoPictureIsreferencedGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["PictureId"] = req.PictureId
    
    return params
}

// TaobaoPictureIsreferencedGetResponse 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaoPictureIsreferencedGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsReferenced Basic图片是否被引用
    IsReferenced bool `json:"is_referenced";xml:"is_referenced"`
    
}

// TaobaoRegionWarehouseQueryRequest 查询仓库覆盖范围
type TaobaoRegionWarehouseQueryRequest struct {
    
    // StoreCode required仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoRegionWarehouseQueryRequest) GetAPIName() string {
	return "taobao.region.warehouse.query"
}

func (req *TaobaoRegionWarehouseQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoRegionWarehouseQueryResponse 查询仓库覆盖范围
type TaobaoRegionWarehouseQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result BaseResult `json:"result";xml:"result"`
    
}

// TaobaoDeliveryTemplateDeleteRequest 根据用户指定的模板ID删除指定的模板
type TaobaoDeliveryTemplateDeleteRequest struct {
    
    // TemplateId required运费模板ID
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TaobaoDeliveryTemplateDeleteRequest) GetAPIName() string {
	return "taobao.delivery.template.delete"
}

func (req *TaobaoDeliveryTemplateDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TemplateId"] = req.TemplateId
    
    return params
}

// TaobaoDeliveryTemplateDeleteResponse 根据用户指定的模板ID删除指定的模板
type TaobaoDeliveryTemplateDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Complete Basic表示删除成功还是失败
    Complete bool `json:"complete";xml:"complete"`
    
}

// TaobaoDeliveryTemplatesGetRequest 根据用户ID获取用户下所有模板
type TaobaoDeliveryTemplatesGetRequest struct {
    
    // Fields required需返回的字段列表。 <br/> 
    // <B>
    // 可选值:示例中的值;各字段之间用","隔开
    // </B>
    // <br/><br/> 
    // <font color=blue>
    // template_id：查询模板ID <br/> 
    // template_name:查询模板名称 <br/> 
    // assumer：查询服务承担放 <br/> 
    // valuation:查询计价规则 <br/> 
    // supports:查询增值服务列表 <br/> 
    // created:查询模板创建时间 <br/> 
    // modified:查询修改时间<br/> 
    // consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> 
    // address:卖家地址信息
    // </font>
    // <br/><br/> 
    // <font color=bule>
    // query_cod:查询货到付款运费信息<br/> 
    // query_post:查询平邮运费信息 <br/> 
    // query_express:查询快递公司运费信息 <br/> 
    // query_ems:查询EMS运费信息<br/> 
    // query_bzsd:查询普通保障速递运费信息<br/> 
    // query_wlb:查询物流宝保障速递运费信息<br/> 
    // query_furniture:查询家装物流运费信息
    // <font color=blue>
    // <br/><br/>
    // <font color=red>不能有空格</font>
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoDeliveryTemplatesGetRequest) GetAPIName() string {
	return "taobao.delivery.templates.get"
}

func (req *TaobaoDeliveryTemplatesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoDeliveryTemplatesGetResponse 根据用户ID获取用户下所有模板
type TaobaoDeliveryTemplatesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DeliveryTemplates Object Array运费模板列表
    DeliveryTemplates DeliveryTemplate `json:"delivery_templates";xml:"delivery_templates"`
    
    // TotalResults Basic获得到符合条件的结果总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoDeliveryTemplateUpdateRequest 修改运费模板
type TaobaoDeliveryTemplateUpdateRequest struct {
    
    // Assumer optional可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
    Assumer int64 `json:"assumer";xml:"assumer"`
    
    // Name optional模板名称，长度不能大于50个字节
    Name string `json:"name";xml:"name"`
    
    // TemplateAddFees required增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateAddFees string `json:"template_add_fees";xml:"template_add_fees"`
    
    // TemplateAddStandards required增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font>
    // <br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateAddStandards string `json:"template_add_standards";xml:"template_add_standards"`
    
    // TemplateDests required邮费子项涉及的地区.结构: value1;value2;value3,value4
    // <br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
    // 如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
    // <br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
    // <br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
    // <br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
    // template_start_standards(首费标准)、template_start_fees(首费)、
    // template_add_standards(增费标准)、
    // template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
    TemplateDests string `json:"template_dests";xml:"template_dests"`
    
    // TemplateId required需要修改的模板对应的模板ID
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
    // TemplateStartFees required首费：输入0.01-999.99（最多包含两位小数）
    // <br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateStartFees string `json:"template_start_fees";xml:"template_start_fees"`
    
    // TemplateStartStandards required首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br>
    // <br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateStartStandards string `json:"template_start_standards";xml:"template_start_standards"`
    
    // TemplateTypes required运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4 
    // 如: post;express;ems;cod 
    // <br/><font color=red>
    // 注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. 
    //  </font>
    // <br/>
    // <font color=blue>
    // 普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。
    // 货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font>
    TemplateTypes string `json:"template_types";xml:"template_types"`
    
}

func (req *TaobaoDeliveryTemplateUpdateRequest) GetAPIName() string {
	return "taobao.delivery.template.update"
}

func (req *TaobaoDeliveryTemplateUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["Assumer"] = req.Assumer
    
    params["Name"] = req.Name
    
    params["TemplateAddFees"] = req.TemplateAddFees
    
    params["TemplateAddStandards"] = req.TemplateAddStandards
    
    params["TemplateDests"] = req.TemplateDests
    
    params["TemplateId"] = req.TemplateId
    
    params["TemplateStartFees"] = req.TemplateStartFees
    
    params["TemplateStartStandards"] = req.TemplateStartStandards
    
    params["TemplateTypes"] = req.TemplateTypes
    
    return params
}

// TaobaoDeliveryTemplateUpdateResponse 修改运费模板
type TaobaoDeliveryTemplateUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Complete Basic表示修改是否成功
    Complete bool `json:"complete";xml:"complete"`
    
}

// TaobaoDeliveryTemplateAddRequest 增加运费模板的外部接口
type TaobaoDeliveryTemplateAddRequest struct {
    
    // Assumer required可选值：0、1 ，说明如下<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
    Assumer int64 `json:"assumer";xml:"assumer"`
    
    // ConsignAreaId optional卖家发货地址区域ID
    // <br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm 
    // </font>
    // 
    // <br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font>
    ConsignAreaId int64 `json:"consign_area_id";xml:"consign_area_id"`
    
    // Name required运费模板的名称，长度不能超过50个字节
    Name string `json:"name";xml:"name"`
    
    // TemplateAddFees required增费：输入0.00-999.99（最多包含两位小数）
    // 
    // <br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font>
    // 
    // 
    // <br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateAddFees string `json:"template_add_fees";xml:"template_add_fees"`
    
    // TemplateAddStandards required增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
    // <br/>
    // <br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateAddStandards string `json:"template_add_standards";xml:"template_add_standards"`
    
    // TemplateDests required邮费子项涉及的地区.结构: value1;value2;value3,value4
    // <br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
    // 如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
    // <br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/>
    // <br/><font color=red>每个运费方式设置涉及的地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
    // <br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
    // template_start_standards(首费标准)、template_start_fees(首费)、
    // template_add_standards(增费标准)、
    // template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
    // <font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font>
    TemplateDests string `json:"template_dests";xml:"template_dests"`
    
    // TemplateStartFees required首费：输入0.00-999.99（最多包含两位小数）
    // <br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateStartFees string `json:"template_start_fees";xml:"template_start_fees"`
    
    // TemplateStartStandards required首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
    // <br/>
    // <font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    TemplateStartStandards string `json:"template_start_standards";xml:"template_start_standards"`
    
    // TemplateTypes required运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4 
    // 如: post;express;ems;cod 
    // <br/><font color=red>
    // 注意:在添加多个运费方式时,字符串中使用 ";" 分号区分
    // 。template_dests(指定地区)
    // template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font>
    // <br>
    // <font color=blue>
    // 注意：<br/>
    // 1、post,ems,express三种运费方式必须填写一个<br/>
    // 2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture
    // 只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>
    // 3、如果是货到付款用户当没有填写cod运送方式时，运费模板会默认继承express的费用为cod的费用<br/>
    // 4、如果是保障速递户当没有填写bzsd运送方式时，运费模板会默认继承express的费用为bzsd的费用<br/>
    // 5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的
    // </font>
    TemplateTypes string `json:"template_types";xml:"template_types"`
    
    // Valuation required可选值：0、1、3，说明如下。<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费
    //  <br>3:表示按宝贝体积计算运费
    Valuation int64 `json:"valuation";xml:"valuation"`
    
}

func (req *TaobaoDeliveryTemplateAddRequest) GetAPIName() string {
	return "taobao.delivery.template.add"
}

func (req *TaobaoDeliveryTemplateAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["Assumer"] = req.Assumer
    
    params["ConsignAreaId"] = req.ConsignAreaId
    
    params["Name"] = req.Name
    
    params["TemplateAddFees"] = req.TemplateAddFees
    
    params["TemplateAddStandards"] = req.TemplateAddStandards
    
    params["TemplateDests"] = req.TemplateDests
    
    params["TemplateStartFees"] = req.TemplateStartFees
    
    params["TemplateStartStandards"] = req.TemplateStartStandards
    
    params["TemplateTypes"] = req.TemplateTypes
    
    params["Valuation"] = req.Valuation
    
    return params
}

// TaobaoDeliveryTemplateAddResponse 增加运费模板的外部接口
type TaobaoDeliveryTemplateAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DeliveryTemplate Object模板对象
    DeliveryTemplate DeliveryTemplate `json:"delivery_template";xml:"delivery_template"`
    
}

// TaobaoPictureUpdateRequest 修改指定图片的图片名
type TaobaoPictureUpdateRequest struct {
    
    // NewName required新的图片名，最大长度50字符，不能为空
    NewName string `json:"new_name";xml:"new_name"`
    
    // PictureId required要更改名字的图片的id
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureUpdateRequest) GetAPIName() string {
	return "taobao.picture.update"
}

func (req *TaobaoPictureUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["NewName"] = req.NewName
    
    params["PictureId"] = req.PictureId
    
    return params
}

// TaobaoPictureUpdateResponse 修改指定图片的图片名
type TaobaoPictureUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Done Basic更新是否成功
    Done bool `json:"done";xml:"done"`
    
}

// TaobaoOmniorderStoreSdtstatusRequest 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusRequest struct {
    
    // PackageId optional菜鸟裹裹的包裹ID
    PackageId int64 `json:"package_id";xml:"package_id"`
    
}

func (req *TaobaoOmniorderStoreSdtstatusRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtstatus"
}

func (req *TaobaoOmniorderStoreSdtstatusRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["PackageId"] = req.PackageId
    
    return params
}

// TaobaoOmniorderStoreSdtstatusResponse 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoDeliveryTemplateGetRequest 获取用户指定运费模板信息
type TaobaoDeliveryTemplateGetRequest struct {
    
    // Fields required需返回的字段列表。 <br/> 
    // <B>
    // 可选值:示例中的值;各字段之间用","隔开
    // </B>
    // <br/><br/> 
    // <font color=blue>
    // template_id：查询模板ID <br/> 
    // template_name:查询模板名称 <br/> 
    // assumer：查询服务承担放 <br/> 
    // valuation:查询计价规则 <br/> 
    // supports:查询增值服务列表 <br/> 
    // created:查询模板创建时间 <br/> 
    // modified:查询修改时间<br/> 
    // consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> 
    // address:卖家地址信息
    // </font>
    // <br/><br/> 
    // <font color=bule>
    // query_cod:查询货到付款运费信息<br/> 
    // query_post:查询平邮运费信息 <br/> 
    // query_express:查询快递公司运费信息 <br/> 
    // query_ems:查询EMS运费信息<br/> 
    // query_bzsd:查询普通保障速递运费信息<br/> 
    // query_wlb:查询物流宝保障速递运费信息<br/> 
    // query_furniture:查询家装物流运费信息
    // <font color=blue>
    // <br/><br/>
    // <font color=red>不能有空格</font>
    Fields string `json:"fields";xml:"fields"`
    
    // TemplateIds required需要查询的运费模板ID列表
    TemplateIds string `json:"template_ids";xml:"template_ids"`
    
    // UserNick optional在没有登录的情况下根据淘宝用户昵称查询指定的模板
    UserNick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoDeliveryTemplateGetRequest) GetAPIName() string {
	return "taobao.delivery.template.get"
}

func (req *TaobaoDeliveryTemplateGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Fields"] = req.Fields
    
    params["TemplateIds"] = req.TemplateIds
    
    params["UserNick"] = req.UserNick
    
    return params
}

// TaobaoDeliveryTemplateGetResponse 获取用户指定运费模板信息
type TaobaoDeliveryTemplateGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DeliveryTemplates Object Array运费模板列表
    DeliveryTemplates DeliveryTemplate `json:"delivery_templates";xml:"delivery_templates"`
    
    // TotalResults Basic获得到符合条件的结果总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoPictureUserinfoGetRequest 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
type TaobaoPictureUserinfoGetRequest struct {
    
}

func (req *TaobaoPictureUserinfoGetRequest) GetAPIName() string {
	return "taobao.picture.userinfo.get"
}

func (req *TaobaoPictureUserinfoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoPictureUserinfoGetResponse 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
type TaobaoPictureUserinfoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UserInfo Object用户使用图片空间的信息
    UserInfo UserInfo `json:"user_info";xml:"user_info"`
    
}

// TaobaoPictureReplaceRequest 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
type TaobaoPictureReplaceRequest struct {
    
    // ImageData required图片二进制文件流,不能为空,允许png、jpg、gif图片格式
    ImageData []byte `json:"image_data";xml:"image_data"`
    
    // PictureId required要替换的图片的id，必须大于0
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
}

func (req *TaobaoPictureReplaceRequest) GetAPIName() string {
	return "taobao.picture.replace"
}

func (req *TaobaoPictureReplaceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ImageData"] = req.ImageData
    
    params["PictureId"] = req.PictureId
    
    return params
}

// TaobaoPictureReplaceResponse 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
type TaobaoPictureReplaceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Done Basic图片替换是否成功
    Done bool `json:"done";xml:"done"`
    
}

// TaobaoQimenItemsMarkingRequest 调用该接口，对商品进行XXXX标的打标、去标的动作。
type TaobaoQimenItemsMarkingRequest struct {
    
    // ActionType required操作类型，string（50），ADD=打标，DELETE=去标，必填
    ActionType string `json:"action_type";xml:"action_type"`
    
    // ItemIds required线上商品ID，long，必填
    ItemIds int64 `json:"item_ids";xml:"item_ids"`
    
    // Remark optional备注，string（500）
    Remark string `json:"remark";xml:"remark"`
    
    // TagType required打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
    TagType string `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoQimenItemsMarkingRequest) GetAPIName() string {
	return "taobao.qimen.items.marking"
}

func (req *TaobaoQimenItemsMarkingRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ActionType"] = req.ActionType
    
    params["ItemIds"] = req.ItemIds
    
    params["Remark"] = req.Remark
    
    params["TagType"] = req.TagType
    
    return params
}

// TaobaoQimenItemsMarkingResponse 调用该接口，对商品进行XXXX标的打标、去标的动作。
type TaobaoQimenItemsMarkingResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Flag Basicflag
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basicmessage
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenTagItemsQueryRequest 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoQimenTagItemsQueryRequest struct {
    
    // Remark optional备注，string（500）
    Remark string `json:"remark";xml:"remark"`
    
    // TagType required打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
    TagType string `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoQimenTagItemsQueryRequest) GetAPIName() string {
	return "taobao.qimen.tag.items.query"
}

func (req *TaobaoQimenTagItemsQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Remark"] = req.Remark
    
    params["TagType"] = req.TagType
    
    return params
}

// TaobaoQimenTagItemsQueryResponse 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoQimenTagItemsQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Flag Basicflag
    Flag string `json:"flag";xml:"flag"`
    
    // ItemIds Basic ArrayitemIds
    ItemIds int64 `json:"item_ids";xml:"item_ids"`
    
    // Message Basicmessage
    Message string `json:"message";xml:"message"`
    
    // TagType BasictagType
    TagType string `json:"tag_type";xml:"tag_type"`
    
}

// TaobaoPictureCategoryUpdateRequest 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
type TaobaoPictureCategoryUpdateRequest struct {
    
    // CategoryId required要更新的图片分类的id
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // CategoryName optional图片分类的新名字，最大长度20字符，不能为空
    CategoryName string `json:"category_name";xml:"category_name"`
    
    // ParentId optional图片分类的新父分类id
    ParentId int64 `json:"parent_id";xml:"parent_id"`
    
}

func (req *TaobaoPictureCategoryUpdateRequest) GetAPIName() string {
	return "taobao.picture.category.update"
}

func (req *TaobaoPictureCategoryUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CategoryId"] = req.CategoryId
    
    params["CategoryName"] = req.CategoryName
    
    params["ParentId"] = req.ParentId
    
    return params
}

// TaobaoPictureCategoryUpdateResponse 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
type TaobaoPictureCategoryUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Done Basic更新图片分类是否成功
    Done bool `json:"done";xml:"done"`
    
}

// TaobaoFenxiaoDealerRequisitionorderCreateRequest 创建经销采购申请
type TaobaoFenxiaoDealerRequisitionorderCreateRequest struct {
    
    // Address required收货人所在街道地址
    Address string `json:"address";xml:"address"`
    
    // BuyerName required买家姓名（自提方式填写提货人姓名）
    BuyerName string `json:"buyer_name";xml:"buyer_name"`
    
    // City optional收货人所在市
    City string `json:"city";xml:"city"`
    
    // District optional收货人所在区
    District string `json:"district";xml:"district"`
    
    // IdCardNumber optional身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
    IdCardNumber string `json:"id_card_number";xml:"id_card_number"`
    
    // LogisticsType optional配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
    LogisticsType string `json:"logistics_type";xml:"logistics_type"`
    
    // Mobile optional买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
    Mobile string `json:"mobile";xml:"mobile"`
    
    // OrderDetail required采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
    OrderDetail string `json:"order_detail";xml:"order_detail"`
    
    // Phone optional买家联系电话（此字段和mobile字段至少填写一个）
    Phone string `json:"phone";xml:"phone"`
    
    // PostCode optional收货人所在地区邮政编码
    PostCode string `json:"post_code";xml:"post_code"`
    
    // Province required收货人所在省份
    Province string `json:"province";xml:"province"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.create"
}

func (req *TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 11)
    
    params["Address"] = req.Address
    
    params["BuyerName"] = req.BuyerName
    
    params["City"] = req.City
    
    params["District"] = req.District
    
    params["IdCardNumber"] = req.IdCardNumber
    
    params["LogisticsType"] = req.LogisticsType
    
    params["Mobile"] = req.Mobile
    
    params["OrderDetail"] = req.OrderDetail
    
    params["Phone"] = req.Phone
    
    params["PostCode"] = req.PostCode
    
    params["Province"] = req.Province
    
    return params
}

// TaobaoFenxiaoDealerRequisitionorderCreateResponse 创建经销采购申请
type TaobaoFenxiaoDealerRequisitionorderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DealerOrderId Basic经销采购申请编号
    DealerOrderId int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
}

// TaobaoQimenItemsTagQueryRequest 调用该接口，查询某个/某批商品上的标
type TaobaoQimenItemsTagQueryRequest struct {
    
    // ItemIds required线上淘宝商品ID，long，必填
    ItemIds int64 `json:"item_ids";xml:"item_ids"`
    
}

func (req *TaobaoQimenItemsTagQueryRequest) GetAPIName() string {
	return "taobao.qimen.items.tag.query"
}

func (req *TaobaoQimenItemsTagQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemIds"] = req.ItemIds
    
    return params
}

// TaobaoQimenItemsTagQueryResponse 调用该接口，查询某个/某批商品上的标
type TaobaoQimenItemsTagQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Flag Basicflag
    Flag string `json:"flag";xml:"flag"`
    
    // ItemTags Object ArrayitemTags
    ItemTags ItemTag `json:"item_tags";xml:"item_tags"`
    
    // Message Basicmessage
    Message string `json:"message";xml:"message"`
    
}

// TaobaoInventoryWarehouseManageRequest 创建商家仓或者更新商家仓信息
type TaobaoInventoryWarehouseManageRequest struct {
    
    // WareHouseDto required仓库信息
    WareHouseDto WareHouseDto `json:"ware_house_dto";xml:"ware_house_dto"`
    
}

func (req *TaobaoInventoryWarehouseManageRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.manage"
}

func (req *TaobaoInventoryWarehouseManageRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WareHouseDto"] = req.WareHouseDto
    
    return params
}

// TaobaoInventoryWarehouseManageResponse 创建商家仓或者更新商家仓信息
type TaobaoInventoryWarehouseManageResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoInventoryWarehouseGetRequest 获取商家仓信息
type TaobaoInventoryWarehouseGetRequest struct {
    
    // StoreCode required仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryWarehouseGetRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.get"
}

func (req *TaobaoInventoryWarehouseGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoInventoryWarehouseGetResponse 获取商家仓信息
type TaobaoInventoryWarehouseGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result BaseResult `json:"result";xml:"result"`
    
}

// TaobaoWlbOutInventoryChangeNotifyRequest 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaoWlbOutInventoryChangeNotifyRequest struct {
    
    // ChangeCount required库存变化数量
    ChangeCount int64 `json:"change_count";xml:"change_count"`
    
    // ItemId required物流宝商品id或前台宝贝id（由type类型决定）
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // OpType requiredOUT--出库 IN--入库
    OpType string `json:"op_type";xml:"op_type"`
    
    // OrderSourceCode optional订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号
    OrderSourceCode string `json:"order_source_code";xml:"order_source_code"`
    
    // OutBizCode required库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。
    OutBizCode string `json:"out_biz_code";xml:"out_biz_code"`
    
    // ResultCount required本次库存变化后库存余额
    ResultCount int64 `json:"result_count";xml:"result_count"`
    
    // Source required（1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购
    Source string `json:"source";xml:"source"`
    
    // StoreCode optional目前非必须，系统会选择默认值
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // Type requiredWLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbOutInventoryChangeNotifyRequest) GetAPIName() string {
	return "taobao.wlb.out.inventory.change.notify"
}

func (req *TaobaoWlbOutInventoryChangeNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["ChangeCount"] = req.ChangeCount
    
    params["ItemId"] = req.ItemId
    
    params["OpType"] = req.OpType
    
    params["OrderSourceCode"] = req.OrderSourceCode
    
    params["OutBizCode"] = req.OutBizCode
    
    params["ResultCount"] = req.ResultCount
    
    params["Source"] = req.Source
    
    params["StoreCode"] = req.StoreCode
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoWlbOutInventoryChangeNotifyResponse 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaoWlbOutInventoryChangeNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic库存变化通知成功时间
    GmtModified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
}

// TaobaoInventoryWarehouseQueryRequest 分页查询商家仓信息
type TaobaoInventoryWarehouseQueryRequest struct {
    
    // PageNo required页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize required页大小
    PageSize int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoInventoryWarehouseQueryRequest) GetAPIName() string {
	return "taobao.inventory.warehouse.query"
}

func (req *TaobaoInventoryWarehouseQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    return params
}

// TaobaoInventoryWarehouseQueryResponse 分页查询商家仓信息
type TaobaoInventoryWarehouseQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result PaginationResult `json:"result";xml:"result"`
    
}

// TaobaoTradePostageUpdateRequest 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
//<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
type TaobaoTradePostageUpdateRequest struct {
    
    // PostFee required邮费价格(邮费单位是元）
    PostFee string `json:"post_fee";xml:"post_fee"`
    
    // Tid required主订单编号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradePostageUpdateRequest) GetAPIName() string {
	return "taobao.trade.postage.update"
}

func (req *TaobaoTradePostageUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["PostFee"] = req.PostFee
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradePostageUpdateResponse 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
//<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
type TaobaoTradePostageUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Trade Object返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment
    Trade Trade `json:"trade";xml:"trade"`
    
}

// TaobaoQimenStockoutConfirmRequest 货品出库后，WMS将状态回传给ERP
type TaobaoQimenStockoutConfirmRequest struct {
    
    // DeliveryOrder optionaldeliveryOrder
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderLines optionalorderLines
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // Packages optionalpackages
    Packages Package `json:"packages";xml:"packages"`
    
}

func (req *TaobaoQimenStockoutConfirmRequest) GetAPIName() string {
	return "taobao.qimen.stockout.confirm"
}

func (req *TaobaoQimenStockoutConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["DeliveryOrder"] = req.DeliveryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderLines"] = req.OrderLines
    
    params["Packages"] = req.Packages
    
    return params
}

// TaobaoQimenStockoutConfirmResponse 货品出库后，WMS将状态回传给ERP
type TaobaoQimenStockoutConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoItemDeleteRequest 删除单条商品
type TaobaoItemDeleteRequest struct {
    
    // NumIid required商品数字ID，该参数必须
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemDeleteRequest) GetAPIName() string {
	return "taobao.item.delete"
}

func (req *TaobaoItemDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["NumIid"] = req.NumIid
    
    return params
}

// TaobaoItemDeleteResponse 删除单条商品
type TaobaoItemDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object被删除商品的相关信息
    Item Item `json:"item";xml:"item"`
    
}

// TaobaoOmniorderStoreSwitchstatusGetRequest 查询门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusGetRequest struct {
    
    // SellerId required卖家ID
    SellerId int64 `json:"seller_id";xml:"seller_id"`
    
    // StoreId required门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreSwitchstatusGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.switchstatus.get"
}

func (req *TaobaoOmniorderStoreSwitchstatusGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["SellerId"] = req.SellerId
    
    params["StoreId"] = req.StoreId
    
    return params
}

// TaobaoOmniorderStoreSwitchstatusGetResponse 查询门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// CainiaoSmartdeliveryStrategyIUpdateRequest 更新智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIUpdateRequest struct {
    
    // DeliveryStrategyInfo required发货策略信息
    DeliveryStrategyInfo DeliveryStrategyInfo `json:"delivery_strategy_info";xml:"delivery_strategy_info"`
    
}

func (req *CainiaoSmartdeliveryStrategyIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.i.update"
}

func (req *CainiaoSmartdeliveryStrategyIUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["DeliveryStrategyInfo"] = req.DeliveryStrategyInfo
    
    return params
}

// CainiaoSmartdeliveryStrategyIUpdateResponse 更新智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Successful Basic设置是否成功
    Successful bool `json:"successful";xml:"successful"`
    
}

// CainiaoSmartdeliveryPriceofferIQueryRequest 查询智能发货引擎商家价格信息
type CainiaoSmartdeliveryPriceofferIQueryRequest struct {
    
    // QueryCpPriceInfoRequest required请求参数
    QueryCpPriceInfoRequest QueryCpPriceInfoRequest `json:"query_cp_price_info_request";xml:"query_cp_price_info_request"`
    
}

func (req *CainiaoSmartdeliveryPriceofferIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.priceoffer.i.query"
}

func (req *CainiaoSmartdeliveryPriceofferIQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["QueryCpPriceInfoRequest"] = req.QueryCpPriceInfoRequest
    
    return params
}

// CainiaoSmartdeliveryPriceofferIQueryResponse 查询智能发货引擎商家价格信息
type CainiaoSmartdeliveryPriceofferIQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CpPriceInfoList Object Array返回结果列表
    CpPriceInfoList CpPriceInfo `json:"cp_price_info_list";xml:"cp_price_info_list"`
    
}

// TaobaoOpenuidChangeRequest 将当前应用所属的openUId 转换为对应targetAppkey的openUid
//规则：
//1.如果两个appkey是应用前后台关系，可以直接转换；
//2.如果appkey和targetAppkey都属于同一个开发者，不允许互相转换。
//3.如果appkey和targetAppkey不属于同一个开发者，不允许互相转换。
type TaobaoOpenuidChangeRequest struct {
    
    // OpenUid requiredopenUid
    OpenUid string `json:"open_uid";xml:"open_uid"`
    
    // TargetAppKey required转换到的appkey
    TargetAppKey string `json:"target_app_key";xml:"target_app_key"`
    
}

func (req *TaobaoOpenuidChangeRequest) GetAPIName() string {
	return "taobao.openuid.change"
}

func (req *TaobaoOpenuidChangeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["OpenUid"] = req.OpenUid
    
    params["TargetAppKey"] = req.TargetAppKey
    
    return params
}

// TaobaoOpenuidChangeResponse 将当前应用所属的openUId 转换为对应targetAppkey的openUid
//规则：
//1.如果两个appkey是应用前后台关系，可以直接转换；
//2.如果appkey和targetAppkey都属于同一个开发者，不允许互相转换。
//3.如果appkey和targetAppkey不属于同一个开发者，不允许互相转换。
type TaobaoOpenuidChangeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // NewOpenUid Basic转换到新的openUId
    NewOpenUid string `json:"new_open_uid";xml:"new_open_uid"`
    
}

// AlibabaElectronicInvoiceGetRequest 查询已回传淘宝的电子发票,根据主订单id查询
type AlibabaElectronicInvoiceGetRequest struct {
    
    // Tid required淘宝主订单号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *AlibabaElectronicInvoiceGetRequest) GetAPIName() string {
	return "alibaba.electronic.invoice.get"
}

func (req *AlibabaElectronicInvoiceGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Tid"] = req.Tid
    
    return params
}

// AlibabaElectronicInvoiceGetResponse 查询已回传淘宝的电子发票,根据主订单id查询
type AlibabaElectronicInvoiceGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InvoiceDetail Object电子发票详细信息
    InvoiceDetail ElectronicInvoiceDetail `json:"invoice_detail";xml:"invoice_detail"`
    
}

// TaobaoRefundMessageAddRequest 创建退款留言/凭证
type TaobaoRefundMessageAddRequest struct {
    
    // Content required留言内容。最大长度: 400个字节
    Content string `json:"content";xml:"content"`
    
    // Image required图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
    Image []byte `json:"image";xml:"image"`
    
    // RefundId required退款编号。
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
}

func (req *TaobaoRefundMessageAddRequest) GetAPIName() string {
	return "taobao.refund.message.add"
}

func (req *TaobaoRefundMessageAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Content"] = req.Content
    
    params["Image"] = req.Image
    
    params["RefundId"] = req.RefundId
    
    return params
}

// TaobaoRefundMessageAddResponse 创建退款留言/凭证
type TaobaoRefundMessageAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // RefundMessage Object退款信息。包含id和created
    RefundMessage RefundMessage `json:"refund_message";xml:"refund_message"`
    
}

// TaobaoRefundMessagesGetRequest 查询退款留言/凭证列表
type TaobaoRefundMessagesGetRequest struct {
    
    // Fields required需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // RefundId required退款单号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase optional退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRefundMessagesGetRequest) GetAPIName() string {
	return "taobao.refund.messages.get"
}

func (req *TaobaoRefundMessagesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    return params
}

// TaobaoRefundMessagesGetResponse 查询退款留言/凭证列表
type TaobaoRefundMessagesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // RefundMessages Object Array查询到的退款留言/凭证列表
    RefundMessages RefundMessage `json:"refund_messages";xml:"refund_messages"`
    
    // TotalResults Basic查询到的退款留言/凭证总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// CainiaoSmartdeliveryPriceofferIUpdateRequest 智能发货引擎更新价格信息模板
type CainiaoSmartdeliveryPriceofferIUpdateRequest struct {
    
    // CpPriceInfo required物流公司价格信息
    CpPriceInfo CpPriceInfo `json:"cp_price_info";xml:"cp_price_info"`
    
}

func (req *CainiaoSmartdeliveryPriceofferIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.priceoffer.i.update"
}

func (req *CainiaoSmartdeliveryPriceofferIUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["CpPriceInfo"] = req.CpPriceInfo
    
    return params
}

// CainiaoSmartdeliveryPriceofferIUpdateResponse 智能发货引擎更新价格信息模板
type CainiaoSmartdeliveryPriceofferIUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Successful Basic设置是否成功
    Successful bool `json:"successful";xml:"successful"`
    
}

// CainiaoSmartdeliveryCocpsIQueryRequest 获取电子面单订购关系中智能发货引擎支持的合作物流公司
type CainiaoSmartdeliveryCocpsIQueryRequest struct {
    
    // SendAddress optional发货地址
    SendAddress Address `json:"send_address";xml:"send_address"`
    
}

func (req *CainiaoSmartdeliveryCocpsIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.cocps.i.query"
}

func (req *CainiaoSmartdeliveryCocpsIQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["SendAddress"] = req.SendAddress
    
    return params
}

// CainiaoSmartdeliveryCocpsIQueryResponse 获取电子面单订购关系中智能发货引擎支持的合作物流公司
type CainiaoSmartdeliveryCocpsIQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SmartDeliveryCollaborateCpsInfoList Object Array返回结果
    SmartDeliveryCollaborateCpsInfoList SmartDeliveryCollaborateCpsInfo `json:"smart_delivery_collaborate_cps_info_list";xml:"smart_delivery_collaborate_cps_info_list"`
    
}

// TaobaoTraderateImprImprwordsGetRequest 根据淘宝后台类目的一级类目和叶子类目
type TaobaoTraderateImprImprwordsGetRequest struct {
    
    // CatLeafId optional淘宝叶子类目id
    CatLeafId int64 `json:"cat_leaf_id";xml:"cat_leaf_id"`
    
    // CatRootId required淘宝一级类目id
    CatRootId int64 `json:"cat_root_id";xml:"cat_root_id"`
    
}

func (req *TaobaoTraderateImprImprwordsGetRequest) GetAPIName() string {
	return "taobao.traderate.impr.imprwords.get"
}

func (req *TaobaoTraderateImprImprwordsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CatLeafId"] = req.CatLeafId
    
    params["CatRootId"] = req.CatRootId
    
    return params
}

// TaobaoTraderateImprImprwordsGetResponse 根据淘宝后台类目的一级类目和叶子类目
type TaobaoTraderateImprImprwordsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ImprWords Basic Array返回类目下所有大家印象的标签
    ImprWords string `json:"impr_words";xml:"impr_words"`
    
}

// TaobaoItemcatsGetRequest 获取后台供卖家发布商品的标准商品类目。
type TaobaoItemcatsGetRequest struct {
    
    // Cids special商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
    Cids int64 `json:"cids";xml:"cids"`
    
    // Datetime special无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化
    Datetime time.Time `json:"datetime";xml:"datetime"`
    
    // Fields optional需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
    Fields string `json:"fields";xml:"fields"`
    
    // ParentCid special父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
    ParentCid int64 `json:"parent_cid";xml:"parent_cid"`
    
}

func (req *TaobaoItemcatsGetRequest) GetAPIName() string {
	return "taobao.itemcats.get"
}

func (req *TaobaoItemcatsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Cids"] = req.Cids
    
    params["Datetime"] = req.Datetime
    
    params["Fields"] = req.Fields
    
    params["ParentCid"] = req.ParentCid
    
    return params
}

// TaobaoItemcatsGetResponse 获取后台供卖家发布商品的标准商品类目。
type TaobaoItemcatsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemCats Object Array增量类目信息,根据fields传入的参数返回相应的结果；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
    ItemCats ItemCat `json:"item_cats";xml:"item_cats"`
    
    // LastModified Basic最近修改时间(如果取增量，会返回该字段)。
    LastModified time.Time `json:"last_modified";xml:"last_modified"`
    
}

// TaobaoItempropsGetRequest 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
type TaobaoItempropsGetRequest struct {
    
    // AttrKeys optional属性的Key，支持多条，以“,”分隔
    AttrKeys string `json:"attr_keys";xml:"attr_keys"`
    
    // ChildPath optional类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
    ChildPath string `json:"child_path";xml:"child_path"`
    
    // Cid required叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
    Cid int64 `json:"cid";xml:"cid"`
    
    // Datetime optional增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime)
    Datetime time.Time `json:"datetime";xml:"datetime"`
    
    // Fields optional需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
    Fields string `json:"fields";xml:"fields"`
    
    // IsColorProp optional是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
    IsColorProp bool `json:"is_color_prop";xml:"is_color_prop"`
    
    // IsEnumProp optional是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
    IsEnumProp bool `json:"is_enum_prop";xml:"is_enum_prop"`
    
    // IsInputProp optional在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
    IsInputProp bool `json:"is_input_prop";xml:"is_input_prop"`
    
    // IsItemProp optional是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
    IsItemProp bool `json:"is_item_prop";xml:"is_item_prop"`
    
    // IsKeyProp optional是否关键属性。可选值:true(是),false(否)
    IsKeyProp bool `json:"is_key_prop";xml:"is_key_prop"`
    
    // IsSaleProp optional是否销售属性。可选值:true(是),false(否)
    IsSaleProp bool `json:"is_sale_prop";xml:"is_sale_prop"`
    
    // ParentPid optional父属性ID
    ParentPid int64 `json:"parent_pid";xml:"parent_pid"`
    
    // Pid optional属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
    Pid int64 `json:"pid";xml:"pid"`
    
    // Type optional获取类目的类型：1代表集市、2代表天猫
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItempropsGetRequest) GetAPIName() string {
	return "taobao.itemprops.get"
}

func (req *TaobaoItempropsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 14)
    
    params["AttrKeys"] = req.AttrKeys
    
    params["ChildPath"] = req.ChildPath
    
    params["Cid"] = req.Cid
    
    params["Datetime"] = req.Datetime
    
    params["Fields"] = req.Fields
    
    params["IsColorProp"] = req.IsColorProp
    
    params["IsEnumProp"] = req.IsEnumProp
    
    params["IsInputProp"] = req.IsInputProp
    
    params["IsItemProp"] = req.IsItemProp
    
    params["IsKeyProp"] = req.IsKeyProp
    
    params["IsSaleProp"] = req.IsSaleProp
    
    params["ParentPid"] = req.ParentPid
    
    params["Pid"] = req.Pid
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoItempropsGetResponse 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
type TaobaoItempropsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemProps Object Array类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
    ItemProps ItemProp `json:"item_props";xml:"item_props"`
    
    // LastModified Basic最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss
    LastModified time.Time `json:"last_modified";xml:"last_modified"`
    
}

// CainiaoSmartdeliveryStrategyIQueryRequest 查询智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIQueryRequest struct {
    
    // SendAddress optional发货地址
    SendAddress Address `json:"send_address";xml:"send_address"`
    
}

func (req *CainiaoSmartdeliveryStrategyIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.i.query"
}

func (req *CainiaoSmartdeliveryStrategyIQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["SendAddress"] = req.SendAddress
    
    return params
}

// CainiaoSmartdeliveryStrategyIQueryResponse 查询智能发货引擎发货策略设置
type CainiaoSmartdeliveryStrategyIQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DeliveryStrategyInfoList Object Array返回结果列表
    DeliveryStrategyInfoList DeliveryStrategyInfo `json:"delivery_strategy_info_list";xml:"delivery_strategy_info_list"`
    
}

// TaobaoTraderateImprImprwordByfeedidGetRequest 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签
type TaobaoTraderateImprImprwordByfeedidGetRequest struct {
    
    // ChildTradeId required交易订单id号（如果包含子订单，请使用子订单id号）
    ChildTradeId int64 `json:"child_trade_id";xml:"child_trade_id"`
    
}

func (req *TaobaoTraderateImprImprwordByfeedidGetRequest) GetAPIName() string {
	return "taobao.traderate.impr.imprword.byfeedid.get"
}

func (req *TaobaoTraderateImprImprwordByfeedidGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ChildTradeId"] = req.ChildTradeId
    
    return params
}

// TaobaoTraderateImprImprwordByfeedidGetResponse 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签
type TaobaoTraderateImprImprwordByfeedidGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ImprFeed Object根据子订单和买家昵称找到的评价和印象词结果
    ImprFeed ImprFeedIdDO `json:"impr_feed";xml:"impr_feed"`
    
}

// TaobaoTimeGetRequest 获取淘宝系统当前时间
type TaobaoTimeGetRequest struct {
    
}

func (req *TaobaoTimeGetRequest) GetAPIName() string {
	return "taobao.time.get"
}

func (req *TaobaoTimeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoTimeGetResponse 获取淘宝系统当前时间
type TaobaoTimeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Time Basic淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
    Time time.Time `json:"time";xml:"time"`
    
}

// TaobaoPictureCategoryGetRequest 获取图片分类信息
type TaobaoPictureCategoryGetRequest struct {
    
    // ModifiedTime optional图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。
    ModifiedTime time.Time `json:"modified_time";xml:"modified_time"`
    
    // ParentId optional取二级分类时设置为对应父分类id
    // 取一级分类时父分类id设为0
    // 取全部分类的时候不设或设为-1
    ParentId int64 `json:"parent_id";xml:"parent_id"`
    
    // PictureCategoryId optional图片分类ID
    PictureCategoryId int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    // PictureCategoryName optional图片分类名，不支持模糊查询
    PictureCategoryName string `json:"picture_category_name";xml:"picture_category_name"`
    
    // Type optional1
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoPictureCategoryGetRequest) GetAPIName() string {
	return "taobao.picture.category.get"
}

func (req *TaobaoPictureCategoryGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["ModifiedTime"] = req.ModifiedTime
    
    params["ParentId"] = req.ParentId
    
    params["PictureCategoryId"] = req.PictureCategoryId
    
    params["PictureCategoryName"] = req.PictureCategoryName
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoPictureCategoryGetResponse 获取图片分类信息
type TaobaoPictureCategoryGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PictureCategories Object Array图片分类
    PictureCategories PictureCategory `json:"picture_categories";xml:"picture_categories"`
    
}

// TaobaoOmniorderStoreAccpetedRequest ISV Pos端门店接单，通知星盘
type TaobaoOmniorderStoreAccpetedRequest struct {
    
    // ReportTimestamp requiredISV系统上报时间
    ReportTimestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    // SubOrderList required子订单列表
    SubOrderList StoreAcceptedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    // Tid required淘宝交易主订单ID
    Tid int64 `json:"tid";xml:"tid"`
    
    // TraceId optional跟踪Id
    TraceId string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreAccpetedRequest) GetAPIName() string {
	return "taobao.omniorder.store.accpeted"
}

func (req *TaobaoOmniorderStoreAccpetedRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ReportTimestamp"] = req.ReportTimestamp
    
    params["SubOrderList"] = req.SubOrderList
    
    params["Tid"] = req.Tid
    
    params["TraceId"] = req.TraceId
    
    return params
}

// TaobaoOmniorderStoreAccpetedResponse ISV Pos端门店接单，通知星盘
type TaobaoOmniorderStoreAccpetedResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode Basic错误码
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // Message Basic错误内容
    Message string `json:"message";xml:"message"`
    
}

// TaobaoPictureDeleteRequest 删除图片空间图片
type TaobaoPictureDeleteRequest struct {
    
    // PictureIds required图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
    PictureIds string `json:"picture_ids";xml:"picture_ids"`
    
}

func (req *TaobaoPictureDeleteRequest) GetAPIName() string {
	return "taobao.picture.delete"
}

func (req *TaobaoPictureDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["PictureIds"] = req.PictureIds
    
    return params
}

// TaobaoPictureDeleteResponse 删除图片空间图片
type TaobaoPictureDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Success Basic是否删除
    Success string `json:"success";xml:"success"`
    
}

// TaobaoOmniorderStoreRefusedRequest ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedRequest struct {
    
    // ReportTimestamp requiredISV的系统时间
    ReportTimestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    // SubOrderList required子订单列表
    SubOrderList SubOrder `json:"sub_order_list";xml:"sub_order_list"`
    
    // Tid required淘宝交易主订单ID
    Tid int64 `json:"tid";xml:"tid"`
    
    // TraceId optional跟踪Id
    TraceId string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreRefusedRequest) GetAPIName() string {
	return "taobao.omniorder.store.refused"
}

func (req *TaobaoOmniorderStoreRefusedRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ReportTimestamp"] = req.ReportTimestamp
    
    params["SubOrderList"] = req.SubOrderList
    
    params["Tid"] = req.Tid
    
    params["TraceId"] = req.TraceId
    
    return params
}

// TaobaoOmniorderStoreRefusedResponse ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode Basic正常为0,其他表示异常
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // Message Basicmessage
    Message string `json:"message";xml:"message"`
    
}

// TaobaoPictureGetRequest 获取图片信息
type TaobaoPictureGetRequest struct {
    
    // ClientType optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    ClientType string `json:"client_type";xml:"client_type"`
    
    // Deleted optional是否删除,unfroze代表没有删除
    Deleted string `json:"deleted";xml:"deleted"`
    
    // EndDate optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // IsHttps optional是否获取https的链接
    IsHttps bool `json:"is_https";xml:"is_https"`
    
    // ModifiedTime optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    ModifiedTime time.Time `json:"modified_time";xml:"modified_time"`
    
    // OrderBy optional图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
    OrderBy string `json:"order_by";xml:"order_by"`
    
    // PageNo optional页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数.每页返回最多返回100条,默认值40
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // PictureCategoryId optional图片分类ID
    PictureCategoryId int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    // PictureId optional图片ID
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
    // StartDate optional查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // Title optional图片标题,最大长度50字符,中英文都算一字符
    Title string `json:"title";xml:"title"`
    
    // Urls optional图片url查询接口
    Urls string `json:"urls";xml:"urls"`
    
}

func (req *TaobaoPictureGetRequest) GetAPIName() string {
	return "taobao.picture.get"
}

func (req *TaobaoPictureGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 13)
    
    params["ClientType"] = req.ClientType
    
    params["Deleted"] = req.Deleted
    
    params["EndDate"] = req.EndDate
    
    params["IsHttps"] = req.IsHttps
    
    params["ModifiedTime"] = req.ModifiedTime
    
    params["OrderBy"] = req.OrderBy
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["PictureCategoryId"] = req.PictureCategoryId
    
    params["PictureId"] = req.PictureId
    
    params["StartDate"] = req.StartDate
    
    params["Title"] = req.Title
    
    params["Urls"] = req.Urls
    
    return params
}

// TaobaoPictureGetResponse 获取图片信息
type TaobaoPictureGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Pictures Object Array图片信息列表
    Pictures Picture `json:"pictures";xml:"pictures"`
    
    // TotalResults Basic总的结果数
    TotalResults int64 `json:"totalResults";xml:"totalResults"`
    
}

// TaobaoOmniorderStoreConsignedRequest ISV Pos端门店发货，通知星盘
type TaobaoOmniorderStoreConsignedRequest struct {
    
    // ReportTimestamp requiredISV系统上报时间
    ReportTimestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    // SubOrderList required子订单列表
    SubOrderList StoreConsignedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    // Tid required淘宝交易主订单ID
    Tid int64 `json:"tid";xml:"tid"`
    
    // TraceId optional跟踪Id
    TraceId string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderStoreConsignedRequest) GetAPIName() string {
	return "taobao.omniorder.store.consigned"
}

func (req *TaobaoOmniorderStoreConsignedRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ReportTimestamp"] = req.ReportTimestamp
    
    params["SubOrderList"] = req.SubOrderList
    
    params["Tid"] = req.Tid
    
    params["TraceId"] = req.TraceId
    
    return params
}

// TaobaoOmniorderStoreConsignedResponse ISV Pos端门店发货，通知星盘
type TaobaoOmniorderStoreConsignedResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Data Objectdata
    Data StoreConsignedResponse `json:"data";xml:"data"`
    
    // ErrCode Basic错误码
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // Message Basic错误内容
    Message string `json:"message";xml:"message"`
    
}

// TaobaoPictureUploadRequest 图片空间上传接口
type TaobaoPictureUploadRequest struct {
    
    // ClientType optional图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布client:computer电脑版本宝贝使用，client:phone手机版本宝贝使用。注意：当client:phone时，图片限制为宽度在480-620之间，长度不能超过960，否则会报错。
    ClientType string `json:"client_type";xml:"client_type"`
    
    // ImageInputTitle required包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名
    ImageInputTitle string `json:"image_input_title";xml:"image_input_title"`
    
    // Img required图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。
    Img []byte `json:"img";xml:"img"`
    
    // IsHttps optional是否获取https连接
    IsHttps bool `json:"is_https";xml:"is_https"`
    
    // PictureCategoryId required图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类
    PictureCategoryId int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    // PictureId optional如果此参数大于0，而且在后台能查到对应的图片，则此次上传为原图替换
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
    // Title optional图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1
    Title string `json:"title";xml:"title"`
    
}

func (req *TaobaoPictureUploadRequest) GetAPIName() string {
	return "taobao.picture.upload"
}

func (req *TaobaoPictureUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["ClientType"] = req.ClientType
    
    params["ImageInputTitle"] = req.ImageInputTitle
    
    params["Img"] = req.Img
    
    params["IsHttps"] = req.IsHttps
    
    params["PictureCategoryId"] = req.PictureCategoryId
    
    params["PictureId"] = req.PictureId
    
    params["Title"] = req.Title
    
    return params
}

// TaobaoPictureUploadResponse 图片空间上传接口
type TaobaoPictureUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Picture Object当前上传的一张图片信息
    Picture Picture `json:"picture";xml:"picture"`
    
}

// TaobaoTopSecretRegisterRequest 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TaobaoTopSecretRegisterRequest struct {
    
    // UserId optional用户id，保证唯一
    UserId int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoTopSecretRegisterRequest) GetAPIName() string {
	return "taobao.top.secret.register"
}

func (req *TaobaoTopSecretRegisterRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["UserId"] = req.UserId
    
    return params
}

// TaobaoTopSecretRegisterResponse 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TaobaoTopSecretRegisterResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic返回操作是否成功
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoMediaCategoryUpdateRequest 更新媒体文件分类
type TaobaoMediaCategoryUpdateRequest struct {
    
    // MediaCategoryId required文件分类ID,不能为空
    MediaCategoryId int64 `json:"media_category_id";xml:"media_category_id"`
    
    // MediaCategoryName required文件分类名，最大长度20字符，中英文都算一字符,不能为空
    MediaCategoryName string `json:"media_category_name";xml:"media_category_name"`
    
}

func (req *TaobaoMediaCategoryUpdateRequest) GetAPIName() string {
	return "taobao.media.category.update"
}

func (req *TaobaoMediaCategoryUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["MediaCategoryId"] = req.MediaCategoryId
    
    params["MediaCategoryName"] = req.MediaCategoryName
    
    return params
}

// TaobaoMediaCategoryUpdateResponse 更新媒体文件分类
type TaobaoMediaCategoryUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Success Basic更新是否成功标志
    Success bool `json:"success";xml:"success"`
    
}

// TaobaoOmniorderAllocatedinfoSyncRequest ISV分单完成，将分单结果同步给星盘
type TaobaoOmniorderAllocatedinfoSyncRequest struct {
    
    // Message optional分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
    Message string `json:"message";xml:"message"`
    
    // ReportTimestamp required1231243213213
    ReportTimestamp int64 `json:"report_timestamp";xml:"report_timestamp"`
    
    // Status required分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
    Status string `json:"status";xml:"status"`
    
    // SubOrderList required门店的分单列表
    SubOrderList StoreAllocatedResult `json:"sub_order_list";xml:"sub_order_list"`
    
    // Tid required淘宝交易主订单ID
    Tid int64 `json:"tid";xml:"tid"`
    
    // TraceId optional跟踪Id
    TraceId string `json:"trace_id";xml:"trace_id"`
    
}

func (req *TaobaoOmniorderAllocatedinfoSyncRequest) GetAPIName() string {
	return "taobao.omniorder.allocatedinfo.sync"
}

func (req *TaobaoOmniorderAllocatedinfoSyncRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Message"] = req.Message
    
    params["ReportTimestamp"] = req.ReportTimestamp
    
    params["Status"] = req.Status
    
    params["SubOrderList"] = req.SubOrderList
    
    params["Tid"] = req.Tid
    
    params["TraceId"] = req.TraceId
    
    return params
}

// TaobaoOmniorderAllocatedinfoSyncResponse ISV分单完成，将分单结果同步给星盘
type TaobaoOmniorderAllocatedinfoSyncResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode Basic错误码
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // Message Basic错误内容
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenOrderexceptionReportRequest WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
type TaobaoQimenOrderexceptionReportRequest struct {
    
    // CreateTime optional奇门仓储字段
    CreateTime string `json:"createTime";xml:"createTime"`
    
    // DeliveryOrderCode optional奇门仓储字段
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // DeliveryOrderId optional奇门仓储字段
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // ExpressCode optional奇门仓储字段
    ExpressCode string `json:"expressCode";xml:"expressCode"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // LogisticsCode optional奇门仓储字段
    LogisticsCode string `json:"logisticsCode";xml:"logisticsCode"`
    
    // MessageDesc optional奇门仓储字段
    MessageDesc string `json:"messageDesc";xml:"messageDesc"`
    
    // MessageId optional奇门仓储字段
    MessageId string `json:"messageId";xml:"messageId"`
    
    // MessageType optional奇门仓储字段
    MessageType string `json:"messageType";xml:"messageType"`
    
    // OrderLines optional奇门仓储字段
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // OrderType optional奇门仓储字段
    OrderType string `json:"orderType";xml:"orderType"`
    
    // Remark optional奇门仓储字段
    Remark string `json:"remark";xml:"remark"`
    
    // WarehouseCode optional奇门仓储字段
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderexceptionReportRequest) GetAPIName() string {
	return "taobao.qimen.orderexception.report"
}

func (req *TaobaoQimenOrderexceptionReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 13)
    
    params["CreateTime"] = req.CreateTime
    
    params["DeliveryOrderCode"] = req.DeliveryOrderCode
    
    params["DeliveryOrderId"] = req.DeliveryOrderId
    
    params["ExpressCode"] = req.ExpressCode
    
    params["ExtendProps"] = req.ExtendProps
    
    params["LogisticsCode"] = req.LogisticsCode
    
    params["MessageDesc"] = req.MessageDesc
    
    params["MessageId"] = req.MessageId
    
    params["MessageType"] = req.MessageType
    
    params["OrderLines"] = req.OrderLines
    
    params["OrderType"] = req.OrderType
    
    params["Remark"] = req.Remark
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenOrderexceptionReportResponse WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
type TaobaoQimenOrderexceptionReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoTradesSoldIncrementGetRequest 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
//<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
//<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
//<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
//<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
//<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）
type TaobaoTradesSoldIncrementGetRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // EndModified required查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // ExtType optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
    ExtType string `json:"ext_type";xml:"ext_type"`
    
    // Fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0为处方药未审核状态
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span>
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // RateStatus optional评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评)
    RateStatus string `json:"rate_status";xml:"rate_status"`
    
    // StartModified required查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
    // Status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
    Status string `json:"status";xml:"status"`
    
    // Tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
    Tag string `json:"tag";xml:"tag"`
    
    // Type optional交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。
    Type string `json:"type";xml:"type"`
    
    // UseHasNext optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。
    UseHasNext bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldIncrementGetRequest) GetAPIName() string {
	return "taobao.trades.sold.increment.get"
}

func (req *TaobaoTradesSoldIncrementGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["EndModified"] = req.EndModified
    
    params["ExtType"] = req.ExtType
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["RateStatus"] = req.RateStatus
    
    params["StartModified"] = req.StartModified
    
    params["Status"] = req.Status
    
    params["Tag"] = req.Tag
    
    params["Type"] = req.Type
    
    params["UseHasNext"] = req.UseHasNext
    
    return params
}

// TaobaoTradesSoldIncrementGetResponse 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
//<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
//<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
//<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
//<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
//<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）
type TaobaoTradesSoldIncrementGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // TotalResults Basic搜索到的交易信息总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
    // Trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    Trades Trade `json:"trades";xml:"trades"`
    
}

// TaobaoFenxiaoRefundGetRequest 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaoFenxiaoRefundGetRequest struct {
    
    // QuerySellerRefund optional是否查询下游买家的退款信息
    QuerySellerRefund bool `json:"query_seller_refund";xml:"query_seller_refund"`
    
    // SubOrderId required要查询的退款子单的id
    SubOrderId int64 `json:"sub_order_id";xml:"sub_order_id"`
    
}

func (req *TaobaoFenxiaoRefundGetRequest) GetAPIName() string {
	return "taobao.fenxiao.refund.get"
}

func (req *TaobaoFenxiaoRefundGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["QuerySellerRefund"] = req.QuerySellerRefund
    
    params["SubOrderId"] = req.SubOrderId
    
    return params
}

// TaobaoFenxiaoRefundGetResponse 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaoFenxiaoRefundGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // RefundDetail Object退款详情
    RefundDetail RefundDetail `json:"refund_detail";xml:"refund_detail"`
    
}

// TaobaoQimenInventoryreserveCancelRequest 库存预占取消
type TaobaoQimenInventoryreserveCancelRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // ItemInventories optional奇门仓储字段
    ItemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    // OrderCode optional奇门仓储字段
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderSource optional奇门仓储字段
    OrderSource string `json:"orderSource";xml:"orderSource"`
    
    // OwnerCode optional奇门仓储字段
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenInventoryreserveCancelRequest) GetAPIName() string {
	return "taobao.qimen.inventoryreserve.cancel"
}

func (req *TaobaoQimenInventoryreserveCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["ItemInventories"] = req.ItemInventories
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderSource"] = req.OrderSource
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenInventoryreserveCancelResponse 库存预占取消
type TaobaoQimenInventoryreserveCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // IsRetry Basic奇门仓储字段
    IsRetry string `json:"isRetry";xml:"isRetry"`
    
    // ItemInventories Object Array奇门仓储字段
    ItemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderCode Basic奇门仓储字段
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
}

// TaobaoQimenInventoryreserveCreateRequest ERP调用奇门的接口,查询发货库存预占用信息
type TaobaoQimenInventoryreserveCreateRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // ItemInventories optional交易订单信息
    ItemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    // MaxWarehouseNum optional最大仓库数(当一个仓库不满足库存时.是否允许占用多个仓库库存.如果允许.需要指定最大分仓数量.但不能拆分订单明细. 0：不限个数.只要满足发货.不限分占几个仓库 1：默认值.只能单仓发 >1:最大 占用数量)
    MaxWarehouseNum int64 `json:"maxWarehouseNum";xml:"maxWarehouseNum"`
    
    // OrderCode requiredERP订单编码
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderSource required订单来源(213 天猫、201 淘宝、214 京东、202 1688 阿里中文站、203 苏宁在线、204 亚马逊中国、205 当当、208 1号店、207 唯品会、209 国美在线、210 拍拍、206 易贝ebay、211 聚美优品、212 乐蜂 网、215 邮乐、216 凡客、217 优购、218 银泰、219 易讯、221 聚尚网、222 蘑菇街、223 POS门店、301 其他)
    OrderSource int64 `json:"orderSource";xml:"orderSource"`
    
    // OwnerCode required货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // ReceiverInfo optional收件者信息
    ReceiverInfo ReceiverInfo `json:"receiverInfo";xml:"receiverInfo"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenInventoryreserveCreateRequest) GetAPIName() string {
	return "taobao.qimen.inventoryreserve.create"
}

func (req *TaobaoQimenInventoryreserveCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["ItemInventories"] = req.ItemInventories
    
    params["MaxWarehouseNum"] = req.MaxWarehouseNum
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderSource"] = req.OrderSource
    
    params["OwnerCode"] = req.OwnerCode
    
    params["ReceiverInfo"] = req.ReceiverInfo
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenInventoryreserveCreateResponse ERP调用奇门的接口,查询发货库存预占用信息
type TaobaoQimenInventoryreserveCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // IsRetry Basic是否需要重试(Y/N (默认为N)返回Y时建议系统自动重试)
    IsRetry string `json:"isRetry";xml:"isRetry"`
    
    // ItemInventories Object Array
    ItemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderCode BasicERP订单编码
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
}

// TaobaoFenxiaoDistributorProductsGetRequest 分销商查询供应商产品信息
type TaobaoFenxiaoDistributorProductsGetRequest struct {
    
    // DownloadStatus optionaldownload_status
    DownloadStatus string `json:"download_status";xml:"download_status"`
    
    // EndTime optional结束时间
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // Fields optional指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
    Fields string `json:"fields";xml:"fields"`
    
    // ItemIds optional根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
    ItemIds int64 `json:"item_ids";xml:"item_ids"`
    
    // OrderBy optionalorder_by
    OrderBy string `json:"order_by";xml:"order_by"`
    
    // PageNo optional页码（大于0的整数，默认1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页记录数（默认20，最大50）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // Pids optional产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
    Pids int64 `json:"pids";xml:"pids"`
    
    // ProductcatId optional产品线ID
    ProductcatId int64 `json:"productcat_id";xml:"productcat_id"`
    
    // StartTime optional开始修改时间
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
    // SupplierNick optional供应商nick，分页查询时，必传
    SupplierNick string `json:"supplier_nick";xml:"supplier_nick"`
    
    // TimeType optionaltime_type
    TimeType string `json:"time_type";xml:"time_type"`
    
    // TradeType optionaltrade_type
    TradeType string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoDistributorProductsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributor.products.get"
}

func (req *TaobaoFenxiaoDistributorProductsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 13)
    
    params["DownloadStatus"] = req.DownloadStatus
    
    params["EndTime"] = req.EndTime
    
    params["Fields"] = req.Fields
    
    params["ItemIds"] = req.ItemIds
    
    params["OrderBy"] = req.OrderBy
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["Pids"] = req.Pids
    
    params["ProductcatId"] = req.ProductcatId
    
    params["StartTime"] = req.StartTime
    
    params["SupplierNick"] = req.SupplierNick
    
    params["TimeType"] = req.TimeType
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoDistributorProductsGetResponse 分销商查询供应商产品信息
type TaobaoFenxiaoDistributorProductsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // Products Object Array产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
    Products FenxiaoProduct `json:"products";xml:"products"`
    
}

// TaobaoWlbOrderJzpartnerQueryRequest 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
type TaobaoWlbOrderJzpartnerQueryRequest struct {
    
    // ServiceType optionalserviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的）
    ServiceType int64 `json:"service_type";xml:"service_type"`
    
    // TaobaoTradeId optional淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。
    TaobaoTradeId int64 `json:"taobao_trade_id";xml:"taobao_trade_id"`
    
}

func (req *TaobaoWlbOrderJzpartnerQueryRequest) GetAPIName() string {
	return "taobao.wlb.order.jzpartner.query"
}

func (req *TaobaoWlbOrderJzpartnerQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ServiceType"] = req.ServiceType
    
    params["TaobaoTradeId"] = req.TaobaoTradeId
    
    return params
}

// TaobaoWlbOrderJzpartnerQueryResponse 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
type TaobaoWlbOrderJzpartnerQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InstallList Object Array安装服务商列表
    InstallList PartnerNew `json:"install_list";xml:"install_list"`
    
    // IsSuccess Basic接口查询成功或者失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // ResultInfo Basic查询返回信息，如果失败，存储错误信息
    ResultInfo string `json:"result_info";xml:"result_info"`
    
    // ServerList Object Array物流配送服务商对象列表
    ServerList PartnerNew `json:"server_list";xml:"server_list"`
    
}

// TaobaoWlbOrderJzwithinsConsignRequest 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
type TaobaoWlbOrderJzwithinsConsignRequest struct {
    
    // InsPartner optional物流服务商信息
    InsPartner JzPartnerNew `json:"ins_partner";xml:"ins_partner"`
    
    // JzConsignArgs required家装物流发货参数
    JzConsignArgs JzConsignArgsNew `json:"jz_consign_args";xml:"jz_consign_args"`
    
    // Tid required淘宝交易订单号
    Tid int64 `json:"tid";xml:"tid"`
    
    // TmsPartner required物流服务商信息
    TmsPartner JzPartnerNew `json:"tms_partner";xml:"tms_partner"`
    
}

func (req *TaobaoWlbOrderJzwithinsConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.jzwithins.consign"
}

func (req *TaobaoWlbOrderJzwithinsConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["InsPartner"] = req.InsPartner
    
    params["JzConsignArgs"] = req.JzConsignArgs
    
    params["Tid"] = req.Tid
    
    params["TmsPartner"] = req.TmsPartner
    
    return params
}

// TaobaoWlbOrderJzwithinsConsignResponse 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
type TaobaoWlbOrderJzwithinsConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic发货成功或者失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // ResultInfo Basic发货返回信息，如果发货错误则报出对应错误
    ResultInfo string `json:"result_info";xml:"result_info"`
    
}

// TaobaoTradeConfirmfeeGetRequest 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaoTradeConfirmfeeGetRequest struct {
    
    // Tid required交易主订单或子订单ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeConfirmfeeGetRequest) GetAPIName() string {
	return "taobao.trade.confirmfee.get"
}

func (req *TaobaoTradeConfirmfeeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeConfirmfeeGetResponse 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaoTradeConfirmfeeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TradeConfirmFee Object获取到的交易确认收货费用
    TradeConfirmFee TradeConfirmFee `json:"trade_confirm_fee";xml:"trade_confirm_fee"`
    
}

// TaobaoItemJointImgRequest * 关联一张商品图片到num_iid指定的商品中
//    * 传入的num_iid所对应的商品必须属于当前会话的用户
//    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
//    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
type TaobaoItemJointImgRequest struct {
    
    // Id optional商品图片id(如果是更新图片，则需要传该参数)
    Id int64 `json:"id";xml:"id"`
    
    // IsMajor optional上传的图片是否关联为商品主图
    IsMajor bool `json:"is_major";xml:"is_major"`
    
    // IsRectangle optional是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
    IsRectangle bool `json:"is_rectangle";xml:"is_rectangle"`
    
    // NumIid required商品数字ID，必选
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // PicPath required图片URL,图片空间图片的相对地址
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // Position optional图片序号
    Position int64 `json:"position";xml:"position"`
    
}

func (req *TaobaoItemJointImgRequest) GetAPIName() string {
	return "taobao.item.joint.img"
}

func (req *TaobaoItemJointImgRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Id"] = req.Id
    
    params["IsMajor"] = req.IsMajor
    
    params["IsRectangle"] = req.IsRectangle
    
    params["NumIid"] = req.NumIid
    
    params["PicPath"] = req.PicPath
    
    params["Position"] = req.Position
    
    return params
}

// TaobaoItemJointImgResponse * 关联一张商品图片到num_iid指定的商品中
//    * 传入的num_iid所对应的商品必须属于当前会话的用户
//    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
//    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
type TaobaoItemJointImgResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemImg Object商品图片信息
    ItemImg ItemImg `json:"item_img";xml:"item_img"`
    
}

// TaobaoItemJointPropimgRequest * 关联一张商品属性图片到num_iid指定的商品中
//    * 传入的num_iid所对应的商品必须属于当前会话的用户
//    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的
//    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
//    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
type TaobaoItemJointPropimgRequest struct {
    
    // Id optional属性图片ID。如果是新增不需要填写
    Id int64 `json:"id";xml:"id"`
    
    // NumIid required商品数字ID，必选
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // PicPath required图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // Position optional图片序号
    Position int64 `json:"position";xml:"position"`
    
    // Properties required属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。
    Properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemJointPropimgRequest) GetAPIName() string {
	return "taobao.item.joint.propimg"
}

func (req *TaobaoItemJointPropimgRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["Id"] = req.Id
    
    params["NumIid"] = req.NumIid
    
    params["PicPath"] = req.PicPath
    
    params["Position"] = req.Position
    
    params["Properties"] = req.Properties
    
    return params
}

// TaobaoItemJointPropimgResponse * 关联一张商品属性图片到num_iid指定的商品中
//    * 传入的num_iid所对应的商品必须属于当前会话的用户
//    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的
//    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
//    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
type TaobaoItemJointPropimgResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PropImg Object属性图片对象信息
    PropImg PropImg `json:"prop_img";xml:"prop_img"`
    
}

// TaobaoItemsCustomGetRequest 跟据卖家设定的商品外部id获取商品 
//这个商品对应卖家从传入的session中获取，需要session绑定
type TaobaoItemsCustomGetRequest struct {
    
    // Fields required需返回的字段列表，参考：Item商品结构体说明，其中barcode、sku.barcode等条形码字段暂不支持；多个字段之间用“,”分隔。
    Fields string `json:"fields";xml:"fields"`
    
    // OuterId required商品的外部商品ID，支持批量，最多不超过40个。
    OuterId string `json:"outer_id";xml:"outer_id"`
    
}

func (req *TaobaoItemsCustomGetRequest) GetAPIName() string {
	return "taobao.items.custom.get"
}

func (req *TaobaoItemsCustomGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["OuterId"] = req.OuterId
    
    return params
}

// TaobaoItemsCustomGetResponse 跟据卖家设定的商品外部id获取商品 
//这个商品对应卖家从传入的session中获取，需要session绑定
type TaobaoItemsCustomGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Items Object Array商品列表，具体返回字段以fields决定
    Items Item `json:"items";xml:"items"`
    
}

// TaobaoItemsInventoryGetRequest 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
//只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取
type TaobaoItemsInventoryGetRequest struct {
    
    // AuctionType optional商品类型：a-拍卖,b-一口价
    AuctionType string `json:"auction_type";xml:"auction_type"`
    
    // Banner optional分类字段。可选值:<br>
    // regular_shelved(定时上架)<br>
    // never_on_shelf(从未上架)<br>
    // off_shelf(我下架的)<br>
    // <font color='red'>for_shelved(等待所有上架)<br>
    // sold_out(全部卖完)<br>
    // violation_off_shelf(违规下架的)<br>
    // 默认查询for_shelved(等待所有上架)这个状态的商品<br></font>
    // 注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)
    Banner string `json:"banner";xml:"banner"`
    
    // Cid optional商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
    Cid int64 `json:"cid";xml:"cid"`
    
    // EndModified optional商品结束修改时间
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // Fields required需返回的字段列表。可选值为 Item 商品结构体中的以下字段： 
    // approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 
    // 不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。
    Fields string `json:"fields";xml:"fields"`
    
    // HasDiscount optional是否参与会员折扣。可选值：true，false。默认不过滤该条件
    HasDiscount bool `json:"has_discount";xml:"has_discount"`
    
    // IsEx optional商品是否在外部网店显示
    IsEx bool `json:"is_ex";xml:"is_ex"`
    
    // IsTaobao optional商品是否在淘宝显示
    IsTaobao bool `json:"is_taobao";xml:"is_taobao"`
    
    // OrderBy optional排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
    OrderBy string `json:"order_by";xml:"order_by"`
    
    // PageNo optional页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围:大于零的整数;最大值：200；默认值：40。
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // Q optional搜索字段。搜索商品的title。
    Q string `json:"q";xml:"q"`
    
    // SellerCids optional卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
    SellerCids string `json:"seller_cids";xml:"seller_cids"`
    
    // StartModified optional商品起始修改时间
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoItemsInventoryGetRequest) GetAPIName() string {
	return "taobao.items.inventory.get"
}

func (req *TaobaoItemsInventoryGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 14)
    
    params["AuctionType"] = req.AuctionType
    
    params["Banner"] = req.Banner
    
    params["Cid"] = req.Cid
    
    params["EndModified"] = req.EndModified
    
    params["Fields"] = req.Fields
    
    params["HasDiscount"] = req.HasDiscount
    
    params["IsEx"] = req.IsEx
    
    params["IsTaobao"] = req.IsTaobao
    
    params["OrderBy"] = req.OrderBy
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["Q"] = req.Q
    
    params["SellerCids"] = req.SellerCids
    
    params["StartModified"] = req.StartModified
    
    return params
}

// TaobaoItemsInventoryGetResponse 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
//只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取
type TaobaoItemsInventoryGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Items Object Array搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段
    Items Item `json:"items";xml:"items"`
    
    // TotalResults Basic搜索到符合条件的结果总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoItemcatsAuthorizeGetRequest 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
type TaobaoItemcatsAuthorizeGetRequest struct {
    
    // Fields required需要返回的字段。目前支持有：
    // brand.vid, brand.name, 
    // item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,
    // xinpin_item_cat.cid, 
    // xinpin_item_cat.name, 
    // xinpin_item_cat.status,
    // xinpin_item_cat.sort_order,
    // xinpin_item_cat.parent_cid,
    // xinpin_item_cat.is_parent
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoItemcatsAuthorizeGetRequest) GetAPIName() string {
	return "taobao.itemcats.authorize.get"
}

func (req *TaobaoItemcatsAuthorizeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoItemcatsAuthorizeGetResponse 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
type TaobaoItemcatsAuthorizeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SellerAuthorize Object里面有3个数组：
    // Brand[]品牌列表,
    // ItemCat[] 类目列表
    // XinpinItemCat[] 针对于C卖家新品类目列表
    SellerAuthorize SellerAuthorize `json:"seller_authorize";xml:"seller_authorize"`
    
}

// Gfq1qp4287ShopexMatrixDatapushRequest 订单转发
type Gfq1qp4287ShopexMatrixDatapushRequest struct {
    
    // P optional数据内容
    P string `json:"p";xml:"p"`
    
}

func (req *Gfq1qp4287ShopexMatrixDatapushRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.matrix.datapush"
}

func (req *Gfq1qp4287ShopexMatrixDatapushRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["P"] = req.P
    
    return params
}

// Gfq1qp4287ShopexMatrixDatapushResponse 订单转发
type Gfq1qp4287ShopexMatrixDatapushResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Data Basicdata
    Data string `json:"data";xml:"data"`
    
    // Res Basicres
    Res string `json:"res";xml:"res"`
    
}

// TaobaoSkusCustomGetRequest 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku 
//这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoSkusCustomGetRequest struct {
    
    // Fields required需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
    Fields string `json:"fields";xml:"fields"`
    
    // OuterId requiredSku的外部商家ID
    OuterId string `json:"outer_id";xml:"outer_id"`
    
}

func (req *TaobaoSkusCustomGetRequest) GetAPIName() string {
	return "taobao.skus.custom.get"
}

func (req *TaobaoSkusCustomGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["OuterId"] = req.OuterId
    
    return params
}

// TaobaoSkusCustomGetResponse 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku 
//这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoSkusCustomGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Skus Object ArraySku对象，具体字段以fields决定
    Skus Sku `json:"skus";xml:"skus"`
    
}

// TaobaoSellercenterRolesGetRequest 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolesGetRequest struct {
    
    // Nick required卖家昵称(只允许查询自己的信息：当前登陆者)
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterRolesGetRequest) GetAPIName() string {
	return "taobao.sellercenter.roles.get"
}

func (req *TaobaoSellercenterRolesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoSellercenterRolesGetResponse 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Roles Object Array卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点)
    Roles Role `json:"roles";xml:"roles"`
    
}

// TaobaoSellercenterRoleInfoGetRequest 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息)
type TaobaoSellercenterRoleInfoGetRequest struct {
    
    // RoleId required角色id
    RoleId int64 `json:"role_id";xml:"role_id"`
    
}

func (req *TaobaoSellercenterRoleInfoGetRequest) GetAPIName() string {
	return "taobao.sellercenter.role.info.get"
}

func (req *TaobaoSellercenterRoleInfoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["RoleId"] = req.RoleId
    
    return params
}

// TaobaoSellercenterRoleInfoGetResponse 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息)
type TaobaoSellercenterRoleInfoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Role Object角色具体信息
    Role Role `json:"role";xml:"role"`
    
}

// TaobaoSellercenterRoleAddRequest 给指定的卖家创建新的子账号角色<br/>
//如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/>
//code=sell 宝贝管理<br/>
//---------|code=sm 店铺管理<br/>
//---------|---------|code=sm-design 如店铺装修<br/>
//---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
//---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
//---------|---------|code=phone 手机淘宝店铺<br/>
//调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/>
//code=sell 宝贝管理<br/>
//---------|code=sm 店铺管理<br/>
//---------|---------|code=sm-design 如店铺装修<br/>
//---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
//---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
type TaobaoSellercenterRoleAddRequest struct {
    
    // Description optional角色描述
    Description string `json:"description";xml:"description"`
    
    // Name required角色名
    Name string `json:"name";xml:"name"`
    
    // Nick required表示卖家昵称
    Nick string `json:"nick";xml:"nick"`
    
    // PermissionCodes optional需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。
    PermissionCodes string `json:"permission_codes";xml:"permission_codes"`
    
}

func (req *TaobaoSellercenterRoleAddRequest) GetAPIName() string {
	return "taobao.sellercenter.role.add"
}

func (req *TaobaoSellercenterRoleAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Description"] = req.Description
    
    params["Name"] = req.Name
    
    params["Nick"] = req.Nick
    
    params["PermissionCodes"] = req.PermissionCodes
    
    return params
}

// TaobaoSellercenterRoleAddResponse 给指定的卖家创建新的子账号角色<br/>
//如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/>
//code=sell 宝贝管理<br/>
//---------|code=sm 店铺管理<br/>
//---------|---------|code=sm-design 如店铺装修<br/>
//---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
//---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
//---------|---------|code=phone 手机淘宝店铺<br/>
//调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/>
//code=sell 宝贝管理<br/>
//---------|code=sm 店铺管理<br/>
//---------|---------|code=sm-design 如店铺装修<br/>
//---------|---------|---------|code=sm-tbd-visit内店装修入口<br/>
//---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
type TaobaoSellercenterRoleAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Role Object子账号角色
    Role Role `json:"role";xml:"role"`
    
}

// CainiaoSmartdeliveryCpIModifyRequest 商家修改智能发货引擎推荐的cp
type CainiaoSmartdeliveryCpIModifyRequest struct {
    
    // ModifySmartDeliveryCpRequest required修改智选CP请求
    ModifySmartDeliveryCpRequest ModifySmartDeliveryCpRequest `json:"modify_smart_delivery_cp_request";xml:"modify_smart_delivery_cp_request"`
    
}

func (req *CainiaoSmartdeliveryCpIModifyRequest) GetAPIName() string {
	return "cainiao.smartdelivery.cp.i.modify"
}

func (req *CainiaoSmartdeliveryCpIModifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ModifySmartDeliveryCpRequest"] = req.ModifySmartDeliveryCpRequest
    
    return params
}

// CainiaoSmartdeliveryCpIModifyResponse 商家修改智能发货引擎推荐的cp
type CainiaoSmartdeliveryCpIModifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ModifySmartDeliveryCpResponse Object更新智能发货智选cp返回结果
    ModifySmartDeliveryCpResponse ModifySmartDeliveryCpResponse `json:"modify_smart_delivery_cp_response";xml:"modify_smart_delivery_cp_response"`
    
}

// TaobaoTmcMessagesProduceRequest 批量发送消息
type TaobaoTmcMessagesProduceRequest struct {
    
    // Messages requiredtmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}]
    Messages TmcPublishMessage `json:"messages";xml:"messages"`
    
}

func (req *TaobaoTmcMessagesProduceRequest) GetAPIName() string {
	return "taobao.tmc.messages.produce"
}

func (req *TaobaoTmcMessagesProduceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Messages"] = req.Messages
    
    return params
}

// TaobaoTmcMessagesProduceResponse 批量发送消息
type TaobaoTmcMessagesProduceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsAllSuccess Basic是否全部成功
    IsAllSuccess bool `json:"is_all_success";xml:"is_all_success"`
    
    // Results Object Array发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
    Results TmcProduceResult `json:"results";xml:"results"`
    
}

// CainiaoSmartdeliveryIGetRequest 获取智选cp和电子面单信息
type CainiaoSmartdeliveryIGetRequest struct {
    
    // SmartDeliveryBatchRequest required<a href="http://open.taobao.com/docs/doc.htm?treeId=319&articleId=106295&docType=1">智能发货引擎</a>批量请求参数
    SmartDeliveryBatchRequest SmartDeliveryBatchRequest `json:"smart_delivery_batch_request";xml:"smart_delivery_batch_request"`
    
}

func (req *CainiaoSmartdeliveryIGetRequest) GetAPIName() string {
	return "cainiao.smartdelivery.i.get"
}

func (req *CainiaoSmartdeliveryIGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["SmartDeliveryBatchRequest"] = req.SmartDeliveryBatchRequest
    
    return params
}

// CainiaoSmartdeliveryIGetResponse 获取智选cp和电子面单信息
type CainiaoSmartdeliveryIGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SmartDeliveryResponseWrapperList Object Array<a href="http://open.taobao.com/docs/doc.htm?treeId=319&articleId=106295&docType=1">智能发货引擎</a>结果包装类
    SmartDeliveryResponseWrapperList SmartDeliveryResponseWrapper `json:"smart_delivery_response_wrapper_list";xml:"smart_delivery_response_wrapper_list"`
    
}

// TaobaoFenxiaoOrdersGetRequest 分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
type TaobaoFenxiaoOrdersGetRequest struct {
    
    // EndCreated optional结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
    EndCreated time.Time `json:"end_created";xml:"end_created"`
    
    // Fields optional多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。（大于0的整数。默认为1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。（每页条数不超过50条）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // PurchaseOrderId optional采购单编号或分销流水号，若其它参数没传，则此参数必传。
    PurchaseOrderId int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
    // StartCreated optional起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
    StartCreated time.Time `json:"start_created";xml:"start_created"`
    
    // Status optional交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭)
    Status string `json:"status";xml:"status"`
    
    // TcOrderId optional采购单下游买家订单id
    TcOrderId int64 `json:"tc_order_id";xml:"tc_order_id"`
    
    // TimeType optional可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询)
    TimeType string `json:"time_type";xml:"time_type"`
    
}

func (req *TaobaoFenxiaoOrdersGetRequest) GetAPIName() string {
	return "taobao.fenxiao.orders.get"
}

func (req *TaobaoFenxiaoOrdersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["EndCreated"] = req.EndCreated
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["PurchaseOrderId"] = req.PurchaseOrderId
    
    params["StartCreated"] = req.StartCreated
    
    params["Status"] = req.Status
    
    params["TcOrderId"] = req.TcOrderId
    
    params["TimeType"] = req.TimeType
    
    return params
}

// TaobaoFenxiaoOrdersGetResponse 分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
type TaobaoFenxiaoOrdersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PurchaseOrders Object Array采购单及子采购单信息。返回 PurchaseOrder 包含的字段信息。
    PurchaseOrders TopDpOrderDo `json:"purchase_orders";xml:"purchase_orders"`
    
    // TotalResults Basic搜索到的采购单记录总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// AlibabaEinvoiceSerialnoGenerateRequest erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaEinvoiceSerialnoGenerateRequest struct {
    
}

func (req *AlibabaEinvoiceSerialnoGenerateRequest) GetAPIName() string {
	return "alibaba.einvoice.serialno.generate"
}

func (req *AlibabaEinvoiceSerialnoGenerateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// AlibabaEinvoiceSerialnoGenerateResponse erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaEinvoiceSerialnoGenerateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SerialNo Basicresult
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

// AlibabaEinvoiceSerialnoBatchGenerateRequest 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
//优先使用alibaba.einvoice.serial.generate。
type AlibabaEinvoiceSerialnoBatchGenerateRequest struct {
    
}

func (req *AlibabaEinvoiceSerialnoBatchGenerateRequest) GetAPIName() string {
	return "alibaba.einvoice.serialno.batch.generate"
}

func (req *AlibabaEinvoiceSerialnoBatchGenerateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// AlibabaEinvoiceSerialnoBatchGenerateResponse 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
//优先使用alibaba.einvoice.serial.generate。
type AlibabaEinvoiceSerialnoBatchGenerateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SerialNoList Basic Arrayresult
    SerialNoList string `json:"serial_no_list";xml:"serial_no_list"`
    
}

// TaobaoSellercenterRolemembersGetRequest 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolemembersGetRequest struct {
    
    // RoleId required角色id
    RoleId int64 `json:"role_id";xml:"role_id"`
    
}

func (req *TaobaoSellercenterRolemembersGetRequest) GetAPIName() string {
	return "taobao.sellercenter.rolemembers.get"
}

func (req *TaobaoSellercenterRolemembersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["RoleId"] = req.RoleId
    
    return params
}

// TaobaoSellercenterRolemembersGetResponse 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolemembersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Subusers Object Array子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流
    Subusers SubUserInfo `json:"subusers";xml:"subusers"`
    
}

// TaobaoSellercenterSubuserPermissionsRolesGetRequest 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetRequest struct {
    
    // Nick required子账号昵称(子账号标识)
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetAPIName() string {
	return "taobao.sellercenter.subuser.permissions.roles.get"
}

func (req *TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoSellercenterSubuserPermissionsRolesGetResponse 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SubuserPermission Object子账号被所拥有的权限
    SubuserPermission SubUserPermission `json:"subuser_permission";xml:"subuser_permission"`
    
}

// TaobaoWangwangAbstractInitializeRequest 模糊查询服务初始化，只支持json返回
type TaobaoWangwangAbstractInitializeRequest struct {
    
    // Charset optional传入参数的字符集
    Charset string `json:"charset";xml:"charset"`
    
}

func (req *TaobaoWangwangAbstractInitializeRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.initialize"
}

func (req *TaobaoWangwangAbstractInitializeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Charset"] = req.Charset
    
    return params
}

// TaobaoWangwangAbstractInitializeResponse 模糊查询服务初始化，只支持json返回
type TaobaoWangwangAbstractInitializeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorMsg Basic当ret_code=-1时这个变量才有
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // RetCode Basic0或-1表示成功或失败
    RetCode int64 `json:"ret_code";xml:"ret_code"`
    
}

// TaobaoWangwangAbstractGetwordlistRequest 获取关键词列表，只支持json返回
type TaobaoWangwangAbstractGetwordlistRequest struct {
    
    // Charset optional传入参数的字符集
    Charset string `json:"charset";xml:"charset"`
    
}

func (req *TaobaoWangwangAbstractGetwordlistRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.getwordlist"
}

func (req *TaobaoWangwangAbstractGetwordlistRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Charset"] = req.Charset
    
    return params
}

// TaobaoWangwangAbstractGetwordlistResponse 获取关键词列表，只支持json返回
type TaobaoWangwangAbstractGetwordlistResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorMsg Basic例如单词长度太长等，ret_code=-1才有
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // RetCode Basic0或-1，表示错误或正确，错误时有错误信息
    RetCode int64 `json:"ret_code";xml:"ret_code"`
    
    // WordLists Object Array关键词列表，ret_code=0才有
    WordLists WordList `json:"word_lists";xml:"word_lists"`
    
}

// TaobaoOpenuidGetBymixnickRequest 通过mixnick转换openuid
type TaobaoOpenuidGetBymixnickRequest struct {
    
    // MixNick required无线类应用获取到的混淆的nick
    MixNick string `json:"mix_nick";xml:"mix_nick"`
    
}

func (req *TaobaoOpenuidGetBymixnickRequest) GetAPIName() string {
	return "taobao.openuid.get.bymixnick"
}

func (req *TaobaoOpenuidGetBymixnickRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["MixNick"] = req.MixNick
    
    return params
}

// TaobaoOpenuidGetBymixnickResponse 通过mixnick转换openuid
type TaobaoOpenuidGetBymixnickResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OpenUid BasicOpenUID
    OpenUid string `json:"open_uid";xml:"open_uid"`
    
}

// TaobaoWangwangAbstractLogqueryRequest 模糊聊天记录查询
type TaobaoWangwangAbstractLogqueryRequest struct {
    
    // Charset optional传入参数的字符集
    Charset string `json:"charset";xml:"charset"`
    
    // Count optional获取记录条数，默认值是1000
    Count int64 `json:"count";xml:"count"`
    
    // EndDate requiredutc
    EndDate int64 `json:"end_date";xml:"end_date"`
    
    // FromId required卖家id，有cntaobao前缀
    FromId string `json:"from_id";xml:"from_id"`
    
    // NextKey optional设置了这个值，那么聊天记录就从这个点开始查询
    NextKey string `json:"next_key";xml:"next_key"`
    
    // StartDate requiredutc
    StartDate int64 `json:"start_date";xml:"start_date"`
    
    // ToId required买家id，有cntaobao前缀
    ToId string `json:"to_id";xml:"to_id"`
    
}

func (req *TaobaoWangwangAbstractLogqueryRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.logquery"
}

func (req *TaobaoWangwangAbstractLogqueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["Charset"] = req.Charset
    
    params["Count"] = req.Count
    
    params["EndDate"] = req.EndDate
    
    params["FromId"] = req.FromId
    
    params["NextKey"] = req.NextKey
    
    params["StartDate"] = req.StartDate
    
    params["ToId"] = req.ToId
    
    return params
}

// TaobaoWangwangAbstractLogqueryResponse 模糊聊天记录查询
type TaobaoWangwangAbstractLogqueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorMsg Basic例如单词长度太长等。
    // 当ret_code不为0时这个值存在。
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // FromId Basic卖家id
    FromId string `json:"from_id";xml:"from_id"`
    
    // IsSliced Basic0或1
    // 168的时候用户名设置有误
    IsSliced int64 `json:"is_sliced";xml:"is_sliced"`
    
    // NextKey Basic当is_sliced为真时才有这项
    NextKey string `json:"next_key";xml:"next_key"`
    
    // RetCode Basic0或-1，表示错误或正确，错误时有错误信息.
    // 为-1时就只有error_msg字段是有效的。
    RetCode int64 `json:"ret_code";xml:"ret_code"`
    
    // ToId Basic买家id
    ToId string `json:"to_id";xml:"to_id"`
    
    // UrlLists Object Arrayurl列表
    UrlLists UrlList `json:"url_lists";xml:"url_lists"`
    
    // WordCountLists Object Array关键词统计列表
    WordCountLists WordCountList `json:"word_count_lists";xml:"word_count_lists"`
    
}

// TaobaoWangwangAbstractAddwordRequest 增加关键词，只支持json返回
type TaobaoWangwangAbstractAddwordRequest struct {
    
    // Charset optional传入参数的字符集
    Charset string `json:"charset";xml:"charset"`
    
    // Word required关键词，长度大于0
    Word string `json:"word";xml:"word"`
    
}

func (req *TaobaoWangwangAbstractAddwordRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.addword"
}

func (req *TaobaoWangwangAbstractAddwordRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Charset"] = req.Charset
    
    params["Word"] = req.Word
    
    return params
}

// TaobaoWangwangAbstractAddwordResponse 增加关键词，只支持json返回
type TaobaoWangwangAbstractAddwordResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorMsg Basic例如单词长度太长等，当ret_code=-1时才有这项
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // RetCode Basic0或-1，表示错误或正确，错误时有错误信息
    RetCode int64 `json:"ret_code";xml:"ret_code"`
    
}

// TaobaoOpenuidGetBytradeRequest 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoOpenuidGetBytradeRequest struct {
    
    // Tid required订单ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoOpenuidGetBytradeRequest) GetAPIName() string {
	return "taobao.openuid.get.bytrade"
}

func (req *TaobaoOpenuidGetBytradeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoOpenuidGetBytradeResponse 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoOpenuidGetBytradeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OpenUid Basic当前交易tid对应买家的openuid
    OpenUid string `json:"open_uid";xml:"open_uid"`
    
}

// TaobaoWangwangAbstractDeletewordRequest 删除关键词，只支持json返回
type TaobaoWangwangAbstractDeletewordRequest struct {
    
    // Charset optional传入参数的字符集
    Charset string `json:"charset";xml:"charset"`
    
    // Word required关键词，长度大于0
    Word string `json:"word";xml:"word"`
    
}

func (req *TaobaoWangwangAbstractDeletewordRequest) GetAPIName() string {
	return "taobao.wangwang.abstract.deleteword"
}

func (req *TaobaoWangwangAbstractDeletewordRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Charset"] = req.Charset
    
    params["Word"] = req.Word
    
    return params
}

// TaobaoWangwangAbstractDeletewordResponse 删除关键词，只支持json返回
type TaobaoWangwangAbstractDeletewordResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorMsg Basic例如单词长度太长等
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // RetCode Basic0或-1，表示错误或正确，错误时有错误信息
    RetCode int64 `json:"ret_code";xml:"ret_code"`
    
}

// TaobaoOpenuidGetRequest 获取授权账号对应的OpenUid
type TaobaoOpenuidGetRequest struct {
    
}

func (req *TaobaoOpenuidGetRequest) GetAPIName() string {
	return "taobao.openuid.get"
}

func (req *TaobaoOpenuidGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoOpenuidGetResponse 获取授权账号对应的OpenUid
type TaobaoOpenuidGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OpenUid BasicOpenUID
    OpenUid string `json:"open_uid";xml:"open_uid"`
    
}

// TaobaoLogisticsOrderShengxianConfirmRequest 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOrderShengxianConfirmRequest struct {
    
    // CancelId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br>
    CancelId int64 `json:"cancel_id";xml:"cancel_id"`
    
    // DeliveryType required1：冷链。0：常温
    DeliveryType int64 `json:"delivery_type";xml:"delivery_type"`
    
    // LogisId optional物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准
    LogisId int64 `json:"logis_id";xml:"logis_id"`
    
    // OutSid required运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入；
    OutSid string `json:"out_sid";xml:"out_sid"`
    
    // SellerIp optional商家的IP地址
    SellerIp string `json:"seller_ip";xml:"seller_ip"`
    
    // SenderId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货
    SenderId int64 `json:"sender_id";xml:"sender_id"`
    
    // ServiceCode optional仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货
    ServiceCode string `json:"service_code";xml:"service_code"`
    
    // Tid optional淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOrderShengxianConfirmRequest) GetAPIName() string {
	return "taobao.logistics.order.shengxian.confirm"
}

func (req *TaobaoLogisticsOrderShengxianConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["CancelId"] = req.CancelId
    
    params["DeliveryType"] = req.DeliveryType
    
    params["LogisId"] = req.LogisId
    
    params["OutSid"] = req.OutSid
    
    params["SellerIp"] = req.SellerIp
    
    params["SenderId"] = req.SenderId
    
    params["ServiceCode"] = req.ServiceCode
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsOrderShengxianConfirmResponse 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOrderShengxianConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // ShipFresh Object发货成功后，返回承运商的信息
    ShipFresh ShipFresh `json:"ship_fresh";xml:"ship_fresh"`
    
}

// TaobaoLogisticsConsignTcConfirmRequest 下述业务场景可以使用此接口通知相关的交易订单发货：
//1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
//2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
type TaobaoLogisticsConsignTcConfirmRequest struct {
    
    // AppName requiredERP的名称
    AppName string `json:"app_name";xml:"app_name"`
    
    // ExtendFields optional扩展字段 K:V;
    ExtendFields map[string]interface{} `json:"extend_fields";xml:"extend_fields"`
    
    // GoodsList optional发货的包裹
    GoodsList ConfirmConsignGoodsDto `json:"goods_list";xml:"goods_list"`
    
    // OutTradeNo required已发货的外部订单号
    OutTradeNo string `json:"out_trade_no";xml:"out_trade_no"`
    
    // ProviderId required货主id
    ProviderId int64 `json:"provider_id";xml:"provider_id"`
    
    // SellerId required卖家id
    SellerId int64 `json:"seller_id";xml:"seller_id"`
    
    // TradeId required待发货的淘宝主交易号
    TradeId int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoLogisticsConsignTcConfirmRequest) GetAPIName() string {
	return "taobao.logistics.consign.tc.confirm"
}

func (req *TaobaoLogisticsConsignTcConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["AppName"] = req.AppName
    
    params["ExtendFields"] = req.ExtendFields
    
    params["GoodsList"] = req.GoodsList
    
    params["OutTradeNo"] = req.OutTradeNo
    
    params["ProviderId"] = req.ProviderId
    
    params["SellerId"] = req.SellerId
    
    params["TradeId"] = req.TradeId
    
    return params
}

// TaobaoLogisticsConsignTcConfirmResponse 下述业务场景可以使用此接口通知相关的交易订单发货：
//1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
//2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
type TaobaoLogisticsConsignTcConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OrderConsignCode Basic菜鸟发货单据
    OrderConsignCode string `json:"order_consign_code";xml:"order_consign_code"`
    
    // Retry Basic是否重试
    Retry bool `json:"retry";xml:"retry"`
    
    // TraceId Basic单次调用流程唯一id
    TraceId string `json:"trace_id";xml:"trace_id"`
    
}

// TaobaoLogisticsOrdersGetRequest 批量查询物流订单。
type TaobaoLogisticsOrdersGetRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // EndCreated optional创建时间结束
    EndCreated time.Time `json:"end_created";xml:"end_created"`
    
    // Fields required需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br>
    // tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。
    Fields string `json:"fields";xml:"fields"`
    
    // FreightPayer optional谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
    FreightPayer string `json:"freight_payer";xml:"freight_payer"`
    
    // PageNo optional页码.该字段没传 或 值<1 ,则默认page_no为1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数.该字段没传 或 值<1 ,则默认page_size为40
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ReceiverName optional收货人姓名
    ReceiverName string `json:"receiver_name";xml:"receiver_name"`
    
    // SellerConfirm optional卖家是否发货.可选值:yes(是),no(否).如:yes
    SellerConfirm string `json:"seller_confirm";xml:"seller_confirm"`
    
    // StartCreated optional创建时间开始
    StartCreated time.Time `json:"start_created";xml:"start_created"`
    
    // Status optional物流状态.查看数据结构 Shipping 中的status字段.
    Status string `json:"status";xml:"status"`
    
    // Tid optional交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
    Tid int64 `json:"tid";xml:"tid"`
    
    // Type optional物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoLogisticsOrdersGetRequest) GetAPIName() string {
	return "taobao.logistics.orders.get"
}

func (req *TaobaoLogisticsOrdersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["EndCreated"] = req.EndCreated
    
    params["Fields"] = req.Fields
    
    params["FreightPayer"] = req.FreightPayer
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["ReceiverName"] = req.ReceiverName
    
    params["SellerConfirm"] = req.SellerConfirm
    
    params["StartCreated"] = req.StartCreated
    
    params["Status"] = req.Status
    
    params["Tid"] = req.Tid
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoLogisticsOrdersGetResponse 批量查询物流订单。
type TaobaoLogisticsOrdersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shippings Object Array获取的物流订单详情列表 
    // 返回的Shipping包含的具体信息为入参fields请求的字段信息
    Shippings Shipping `json:"shippings";xml:"shippings"`
    
    // TotalResults Basic搜索到的物流订单列表总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoLogisticsOrdersDetailGetRequest 查询物流订单的详细信息，涉及用户隐私字段。
type TaobaoLogisticsOrdersDetailGetRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // EndCreated optional创建时间结束.格式:yyyy-MM-dd HH:mm:ss
    EndCreated time.Time `json:"end_created";xml:"end_created"`
    
    // Fields required需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔.
    Fields string `json:"fields";xml:"fields"`
    
    // FreightPayer optional谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
    FreightPayer string `json:"freight_payer";xml:"freight_payer"`
    
    // PageNo optional页码.该字段没传 或 值<1 ,则默认page_no为1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数.该字段没传 或 值<1 ，则默认page_size为40
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ReceiverName optional收货人姓名
    ReceiverName string `json:"receiver_name";xml:"receiver_name"`
    
    // SellerConfirm optional卖家是否发货.可选值:yes(是),no(否).如:yes.
    SellerConfirm string `json:"seller_confirm";xml:"seller_confirm"`
    
    // StartCreated optional创建时间开始.格式:yyyy-MM-dd HH:mm:ss
    StartCreated time.Time `json:"start_created";xml:"start_created"`
    
    // Status optional物流状态.可查看数据结构 Shipping 中的status字段.
    Status string `json:"status";xml:"status"`
    
    // Tid optional交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息.
    Tid int64 `json:"tid";xml:"tid"`
    
    // Type optional物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoLogisticsOrdersDetailGetRequest) GetAPIName() string {
	return "taobao.logistics.orders.detail.get"
}

func (req *TaobaoLogisticsOrdersDetailGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["EndCreated"] = req.EndCreated
    
    params["Fields"] = req.Fields
    
    params["FreightPayer"] = req.FreightPayer
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["ReceiverName"] = req.ReceiverName
    
    params["SellerConfirm"] = req.SellerConfirm
    
    params["StartCreated"] = req.StartCreated
    
    params["Status"] = req.Status
    
    params["Tid"] = req.Tid
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoLogisticsOrdersDetailGetResponse 查询物流订单的详细信息，涉及用户隐私字段。
type TaobaoLogisticsOrdersDetailGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shippings Object Array获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息.
    Shippings Shipping `json:"shippings";xml:"shippings"`
    
    // TotalResults Basic搜索到的物流订单列表总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoLogisticsCompaniesGetRequest 查询淘宝网合作的物流公司信息，用于发货接口。
type TaobaoLogisticsCompaniesGetRequest struct {
    
    // Fields required需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开.
    // 如:id,code,name,reg_mail_no
    // <br><font color='red'>说明：</font>
    // <br>id：物流公司ID
    // <br>code：物流公司code
    // <br>name：物流公司名称
    // <br>reg_mail_no：物流公司对应的运单规则
    Fields string `json:"fields";xml:"fields"`
    
    // IsRecommended optional是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司.
    IsRecommended bool `json:"is_recommended";xml:"is_recommended"`
    
    // OrderMode optional推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司.
    OrderMode string `json:"order_mode";xml:"order_mode"`
    
}

func (req *TaobaoLogisticsCompaniesGetRequest) GetAPIName() string {
	return "taobao.logistics.companies.get"
}

func (req *TaobaoLogisticsCompaniesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Fields"] = req.Fields
    
    params["IsRecommended"] = req.IsRecommended
    
    params["OrderMode"] = req.OrderMode
    
    return params
}

// TaobaoLogisticsCompaniesGetResponse 查询淘宝网合作的物流公司信息，用于发货接口。
type TaobaoLogisticsCompaniesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // LogisticsCompanies Object Array物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。
    LogisticsCompanies LogisticsCompany `json:"logistics_companies";xml:"logistics_companies"`
    
}

// TaobaoAppipGetRequest 获取ISV发起请求服务器IP
type TaobaoAppipGetRequest struct {
    
}

func (req *TaobaoAppipGetRequest) GetAPIName() string {
	return "taobao.appip.get"
}

func (req *TaobaoAppipGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoAppipGetResponse 获取ISV发起请求服务器IP
type TaobaoAppipGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Ip BasicISV发起请求服务器IP
    Ip string `json:"ip";xml:"ip"`
    
}

// QimenTaobaoQimenEntryorderCallbackRequest ERP通过该接口回传预约入库单对应的商家仓出库单状态
type QimenTaobaoQimenEntryorderCallbackRequest struct {
    
    // Appkey required应用在奇门申请的appkey
    Appkey string `json:"appkey";xml:"appkey"`
    
    // Entryorderlist required入库单信息列表
    Entryorderlist Struct `json:"entryorderlist";xml:"entryorderlist"`
    
    // UserId required货主在奇门授权的ID
    UserId string `json:"userId";xml:"userId"`
    
}

func (req *QimenTaobaoQimenEntryorderCallbackRequest) GetAPIName() string {
	return "qimen.taobao.qimen.entryorder.callback"
}

func (req *QimenTaobaoQimenEntryorderCallbackRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Appkey"] = req.Appkey
    
    params["Entryorderlist"] = req.Entryorderlist
    
    params["UserId"] = req.UserId
    
    return params
}

// QimenTaobaoQimenEntryorderCallbackResponse ERP通过该接口回传预约入库单对应的商家仓出库单状态
type QimenTaobaoQimenEntryorderCallbackResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic错误码:CD001
    Code string `json:"code";xml:"code"`
    
    // Flag Basic回传结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic异常信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoFenxiaoLoginUserGetRequest 获取用户登录信息
type TaobaoFenxiaoLoginUserGetRequest struct {
    
}

func (req *TaobaoFenxiaoLoginUserGetRequest) GetAPIName() string {
	return "taobao.fenxiao.login.user.get"
}

func (req *TaobaoFenxiaoLoginUserGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoFenxiaoLoginUserGetResponse 获取用户登录信息
type TaobaoFenxiaoLoginUserGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // LoginUser Object登录用户信息
    LoginUser LoginUser `json:"login_user";xml:"login_user"`
    
}

// TaobaoRdcAligeniusOrdermsgUpdateRequest 用于订单消息处理状态回传
type TaobaoRdcAligeniusOrdermsgUpdateRequest struct {
    
    // Oid required子订单（消息中传的子订单）
    Oid int64 `json:"oid";xml:"oid"`
    
    // Status required处理状态，1=成功，2=处理失败
    Status int64 `json:"status";xml:"status"`
    
    // Tid required主订单（消息中传的主订单）
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoRdcAligeniusOrdermsgUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.ordermsg.update"
}

func (req *TaobaoRdcAligeniusOrdermsgUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Oid"] = req.Oid
    
    params["Status"] = req.Status
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoRdcAligeniusOrdermsgUpdateResponse 用于订单消息处理状态回传
type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TmallItemQuantityUpdateRequest 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
type TmallItemQuantityUpdateRequest struct {
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // ItemQuantity optional商品库存数；增量编辑方式支持正数、负数
    ItemQuantity int64 `json:"item_quantity";xml:"item_quantity"`
    
    // Options optional商品库存更新时候的可选参数
    Options UpdateItemQuantityOption `json:"options";xml:"options"`
    
    // SkuQuantities optional更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
    SkuQuantities UpdateSkuQuantity `json:"sku_quantities";xml:"sku_quantities"`
    
}

func (req *TmallItemQuantityUpdateRequest) GetAPIName() string {
	return "tmall.item.quantity.update"
}

func (req *TmallItemQuantityUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ItemId"] = req.ItemId
    
    params["ItemQuantity"] = req.ItemQuantity
    
    params["Options"] = req.Options
    
    params["SkuQuantities"] = req.SkuQuantities
    
    return params
}

// TmallItemQuantityUpdateResponse 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
type TmallItemQuantityUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // QuantityUpdateResult Basic库存更新结果，商品id
    QuantityUpdateResult string `json:"quantity_update_result";xml:"quantity_update_result"`
    
}

// TaobaoPictureCategoryAddRequest 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
type TaobaoPictureCategoryAddRequest struct {
    
    // ParentId optional图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
    ParentId int64 `json:"parent_id";xml:"parent_id"`
    
    // PictureCategoryName required图片分类名称，最大长度20字符，中文字符算2个字符，不能为空
    PictureCategoryName string `json:"picture_category_name";xml:"picture_category_name"`
    
}

func (req *TaobaoPictureCategoryAddRequest) GetAPIName() string {
	return "taobao.picture.category.add"
}

func (req *TaobaoPictureCategoryAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ParentId"] = req.ParentId
    
    params["PictureCategoryName"] = req.PictureCategoryName
    
    return params
}

// TaobaoPictureCategoryAddResponse 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
type TaobaoPictureCategoryAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PictureCategory Object图片分类信息
    PictureCategory PictureCategory `json:"picture_category";xml:"picture_category"`
    
}

// TaobaoRegionPriceCancleRequest 取消区域价格
type TaobaoRegionPriceCancleRequest struct {
    
    // ItemId required商品
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // SkuId required无sku传0
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoRegionPriceCancleRequest) GetAPIName() string {
	return "taobao.region.price.cancle"
}

func (req *TaobaoRegionPriceCancleRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoRegionPriceCancleResponse 取消区域价格
type TaobaoRegionPriceCancleResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result BaseResult `json:"result";xml:"result"`
    
}

// TaobaoQimenItemlackReportRequest WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportRequest struct {
    
    // CreateTime required缺货回告创建时间(YYYY-MM-DD HH:mm:ss)
    CreateTime string `json:"createTime";xml:"createTime"`
    
    // DeliveryOrderCode requiredERP的发货单编码
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // DeliveryOrderId optional仓库系统的发货单编码
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional缺货商品列表
    Items Item `json:"items";xml:"items"`
    
    // OutBizCode required外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
    OutBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // WarehouseCode required仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemlackReportRequest) GetAPIName() string {
	return "taobao.qimen.itemlack.report"
}

func (req *TaobaoQimenItemlackReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["CreateTime"] = req.CreateTime
    
    params["DeliveryOrderCode"] = req.DeliveryOrderCode
    
    params["DeliveryOrderId"] = req.DeliveryOrderId
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OutBizCode"] = req.OutBizCode
    
    params["Remark"] = req.Remark
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenItemlackReportResponse WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenOrderprocessQueryRequest ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderCode required单据号
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderId optional仓储系统单据号
    OrderId string `json:"orderId";xml:"orderId"`
    
    // OrderSourceCode optional交易单号
    OrderSourceCode string `json:"orderSourceCode";xml:"orderSourceCode"`
    
    // OrderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderprocessQueryRequest) GetAPIName() string {
	return "taobao.qimen.orderprocess.query"
}

func (req *TaobaoQimenOrderprocessQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderId"] = req.OrderId
    
    params["OrderSourceCode"] = req.OrderSourceCode
    
    params["OrderType"] = req.OrderType
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Remark"] = req.Remark
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenOrderprocessQueryResponse ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderProcess Object订单处理流程
    OrderProcess OrderProcess `json:"orderProcess";xml:"orderProcess"`
    
}

// TaobaoFenxiaoRefundQueryRequest 供应商按查询条件批量查询代销采购退款
type TaobaoFenxiaoRefundQueryRequest struct {
    
    // EndDate required代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // PageNo optional页码（大于0的整数。无值或小于1的值按默认值1计）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // QuerySellerRefund optional是否查询下游买家的退款信息
    QuerySellerRefund bool `json:"query_seller_refund";xml:"query_seller_refund"`
    
    // StartDate required代销采购退款单最早修改时间
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
}

func (req *TaobaoFenxiaoRefundQueryRequest) GetAPIName() string {
	return "taobao.fenxiao.refund.query"
}

func (req *TaobaoFenxiaoRefundQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["EndDate"] = req.EndDate
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["QuerySellerRefund"] = req.QuerySellerRefund
    
    params["StartDate"] = req.StartDate
    
    return params
}

// TaobaoFenxiaoRefundQueryResponse 供应商按查询条件批量查询代销采购退款
type TaobaoFenxiaoRefundQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // RefundList Object Array代销采购退款列表
    RefundList RefundDetail `json:"refund_list";xml:"refund_list"`
    
    // TotalResults Basic按查询条件查到的记录总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoRefundRefusereasonGetRequest 获取商家拒绝原因列表
type TaobaoRefundRefusereasonGetRequest struct {
    
    // Fields required返回参数
    Fields string `json:"fields";xml:"fields"`
    
    // RefundId required退款编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase required售中或售后
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
}

func (req *TaobaoRefundRefusereasonGetRequest) GetAPIName() string {
	return "taobao.refund.refusereason.get"
}

func (req *TaobaoRefundRefusereasonGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Fields"] = req.Fields
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    return params
}

// TaobaoRefundRefusereasonGetResponse 获取商家拒绝原因列表
type TaobaoRefundRefusereasonGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // Reasons Object Array卖家拒绝原因对象
    Reasons Reason `json:"reasons";xml:"reasons"`
    
    // TotalResults Basic原因个数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoQimenReturnorderQueryRequest ERP调用奇门的接口，查询退货入库信息
type TaobaoQimenReturnorderQueryRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页(从1开始)
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页orderLine条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // ReturnOrderCode required退货单编码
    ReturnOrderCode string `json:"returnOrderCode";xml:"returnOrderCode"`
    
    // ReturnOrderId optional仓库系统订单编码
    ReturnOrderId string `json:"returnOrderId";xml:"returnOrderId"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenReturnorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.query"
}

func (req *TaobaoQimenReturnorderQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["ReturnOrderCode"] = req.ReturnOrderCode
    
    params["ReturnOrderId"] = req.ReturnOrderId
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenReturnorderQueryResponse ERP调用奇门的接口，查询退货入库信息
type TaobaoQimenReturnorderQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderLines Object Array订单信息
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // ReturnOrder Object退货单信息
    ReturnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

// TaobaoQimenOrderstatusBatchqueryRequest ERP调用奇门的接口,查询订单在仓库的状态
type TaobaoQimenOrderstatusBatchqueryRequest struct {
    
    // CurrentPage required当前第几页(从1开始)
    CurrentPage int64 `json:"currentPage";xml:"currentPage"`
    
    // EndTime optional订单最后操作时间(查询截止时间点;格式:YYYY-MM-DD hh:mm:ss)
    EndTime string `json:"endTime";xml:"endTime"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK= 换货入库;CNJG= 仓内加工单)
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OwnerCode required货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // PageSize required页面大小(建议不超过100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // StartTime required订单最后操作时间(查询起始时间点;格式:YYYY-MM-DD hh:mm:ss)
    StartTime string `json:"startTime";xml:"startTime"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderstatusBatchqueryRequest) GetAPIName() string {
	return "taobao.qimen.orderstatus.batchquery"
}

func (req *TaobaoQimenOrderstatusBatchqueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["CurrentPage"] = req.CurrentPage
    
    params["EndTime"] = req.EndTime
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderType"] = req.OrderType
    
    params["OwnerCode"] = req.OwnerCode
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenOrderstatusBatchqueryResponse ERP调用奇门的接口,查询订单在仓库的状态
type TaobaoQimenOrderstatusBatchqueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // Orders Object Array单据信息
    Orders Order `json:"orders";xml:"orders"`
    
    // TotalPage Basic总页数
    TotalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

// TaobaoQimenItemlackQueryRequest ERP调用奇门的接口,查询库存商品缺货情况
type TaobaoQimenItemlackQueryRequest struct {
    
    // DeliveryOrderCode required出库单号
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // DeliveryOrderId optional仓储系统出库单号
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页(从1开始)
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页orderLine条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenItemlackQueryRequest) GetAPIName() string {
	return "taobao.qimen.itemlack.query"
}

func (req *TaobaoQimenItemlackQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["DeliveryOrderCode"] = req.DeliveryOrderCode
    
    params["DeliveryOrderId"] = req.DeliveryOrderId
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenItemlackQueryResponse ERP调用奇门的接口,查询库存商品缺货情况
type TaobaoQimenItemlackQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // CreateTime Basic缺货回告创建时间(YYYY-MM-DD HH:mm:ss)
    CreateTime string `json:"createTime";xml:"createTime"`
    
    // DeliveryOrderCode BasicERP的发货单编码
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // DeliveryOrderId Basic仓库系统的发货单编码
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Items Object Array缺货商品列表
    Items Item `json:"items";xml:"items"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // TotalLines BasicorderLines总条数
    TotalLines int64 `json:"totalLines";xml:"totalLines"`
    
    // WarehouseCode Basic仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

// TaobaoQimenOrderPendingRequest ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
type TaobaoQimenOrderPendingRequest struct {
    
    // ActionType required操作类型(pending=挂起;restore=恢复)
    ActionType string `json:"actionType";xml:"actionType"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderCode required单据编码
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderId optional仓储系统单据编码
    OrderId string `json:"orderId";xml:"orderId"`
    
    // OrderType optional单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Reason optional挂起/恢复原因
    Reason string `json:"reason";xml:"reason"`
    
    // WarehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderPendingRequest) GetAPIName() string {
	return "taobao.qimen.order.pending"
}

func (req *TaobaoQimenOrderPendingRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["ActionType"] = req.ActionType
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderId"] = req.OrderId
    
    params["OrderType"] = req.OrderType
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Reason"] = req.Reason
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenOrderPendingResponse ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
type TaobaoQimenOrderPendingResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenDeliveryorderBatchcreateAnswerRequest WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoQimenDeliveryorderBatchcreateAnswerRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Orders optional发货单列表
    Orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchcreate.answer"
}

func (req *TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Orders"] = req.Orders
    
    return params
}

// TaobaoQimenDeliveryorderBatchcreateAnswerResponse WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoQimenDeliveryorderBatchcreateAnswerResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenCombineitemSynchronizeRequest ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoQimenCombineitemSynchronizeRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // ItemCode required组合商品的ERP编码
    ItemCode string `json:"itemCode";xml:"itemCode"`
    
    // ItemId optionaltemp
    ItemId string `json:"itemId";xml:"itemId"`
    
    // Items optional组合商品接口中的单商品信息
    Items Item `json:"items";xml:"items"`
    
    // OwnerCode required货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenCombineitemSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.synchronize"
}

func (req *TaobaoQimenCombineitemSynchronizeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["ItemCode"] = req.ItemCode
    
    params["ItemId"] = req.ItemId
    
    params["Items"] = req.Items
    
    params["OwnerCode"] = req.OwnerCode
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenCombineitemSynchronizeResponse ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoQimenCombineitemSynchronizeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoSellercenterSubusersGetRequest 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaoSellercenterSubusersGetRequest struct {
    
    // Nick required表示卖家昵称
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterSubusersGetRequest) GetAPIName() string {
	return "taobao.sellercenter.subusers.get"
}

func (req *TaobaoSellercenterSubusersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoSellercenterSubusersGetResponse 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaoSellercenterSubusersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Subusers Object Array子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流
    Subusers SubUserInfo `json:"subusers";xml:"subusers"`
    
}

// TaobaoQimenStockQueryRequest ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenStockQueryRequest struct {
    
    // BatchCode optional批次编码
    BatchCode string `json:"batchCode";xml:"batchCode"`
    
    // ExpireDate optional商品过期日期(YYYY-MM-DD)
    ExpireDate string `json:"expireDate";xml:"expireDate"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // InventoryType optional库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
    InventoryType string `json:"inventoryType";xml:"inventoryType"`
    
    // ItemCode required商品编码
    ItemCode string `json:"itemCode";xml:"itemCode"`
    
    // ItemId optional仓储系统商品ID
    ItemId string `json:"itemId";xml:"itemId"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页(从1开始)
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // ProductDate optional商品生产日期(YYYY-MM-DD)
    ProductDate string `json:"productDate";xml:"productDate"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockQueryRequest) GetAPIName() string {
	return "taobao.qimen.stock.query"
}

func (req *TaobaoQimenStockQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["BatchCode"] = req.BatchCode
    
    params["ExpireDate"] = req.ExpireDate
    
    params["ExtendProps"] = req.ExtendProps
    
    params["InventoryType"] = req.InventoryType
    
    params["ItemCode"] = req.ItemCode
    
    params["ItemId"] = req.ItemId
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["ProductDate"] = req.ProductDate
    
    params["Remark"] = req.Remark
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenStockQueryResponse ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenStockQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Items Object Array商品的库存信息列表
    Items Item `json:"items";xml:"items"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // TotalCount Basic总数
    TotalCount int64 `json:"totalCount";xml:"totalCount"`
    
}

// TaobaoSellercenterUserPermissionsGetRequest 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
type TaobaoSellercenterUserPermissionsGetRequest struct {
    
    // Nick required用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoSellercenterUserPermissionsGetRequest) GetAPIName() string {
	return "taobao.sellercenter.user.permissions.get"
}

func (req *TaobaoSellercenterUserPermissionsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoSellercenterUserPermissionsGetResponse 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
type TaobaoSellercenterUserPermissionsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Permissions Object Array权限列表
    Permissions Permission `json:"permissions";xml:"permissions"`
    
}

// TaobaoQimenStoreprocessConfirmRequest WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoQimenStoreprocessConfirmRequest struct {
    
    // ActualQty optional实际作业总数量
    ActualQty int64 `json:"actualQty";xml:"actualQty"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Materialitems optional加工商品列表
    Materialitems MaterialItem `json:"materialitems";xml:"materialitems"`
    
    // OrderCompleteTime required加工单完成时间(YYYY-MM-DD HH:MM:SS)
    OrderCompleteTime string `json:"orderCompleteTime";xml:"orderCompleteTime"`
    
    // OrderType required单据类型(CNJG=仓内加工作业单)
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OutBizCode required外部业务编码(一个合作伙伴中要求唯一多次确认时;每次传入要求唯一(一般传入WMS损益单据编码))
    OutBizCode string `json:"outBizCode";xml:"outBizCode"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // ProcessOrderCode required加工单编码
    ProcessOrderCode string `json:"processOrderCode";xml:"processOrderCode"`
    
    // ProcessOrderId optional仓储系统加工单ID
    ProcessOrderId string `json:"processOrderId";xml:"processOrderId"`
    
    // Productitems optional加工商品列表
    Productitems ProductItem `json:"productitems";xml:"productitems"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenStoreprocessConfirmRequest) GetAPIName() string {
	return "taobao.qimen.storeprocess.confirm"
}

func (req *TaobaoQimenStoreprocessConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 11)
    
    params["ActualQty"] = req.ActualQty
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Materialitems"] = req.Materialitems
    
    params["OrderCompleteTime"] = req.OrderCompleteTime
    
    params["OrderType"] = req.OrderType
    
    params["OutBizCode"] = req.OutBizCode
    
    params["OwnerCode"] = req.OwnerCode
    
    params["ProcessOrderCode"] = req.ProcessOrderCode
    
    params["ProcessOrderId"] = req.ProcessOrderId
    
    params["Productitems"] = req.Productitems
    
    params["Remark"] = req.Remark
    
    return params
}

// TaobaoQimenStoreprocessConfirmResponse WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoQimenStoreprocessConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoTradeOrderskuUpdateRequest 只能更新发货前子订单的销售属性 
//只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 
//必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
type TaobaoTradeOrderskuUpdateRequest struct {
    
    // Oid required子订单编号（对于单笔订单的交易可以传交易编号）。
    Oid int64 `json:"oid";xml:"oid"`
    
    // SkuId special销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
    // SkuProps special销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
    SkuProps string `json:"sku_props";xml:"sku_props"`
    
}

func (req *TaobaoTradeOrderskuUpdateRequest) GetAPIName() string {
	return "taobao.trade.ordersku.update"
}

func (req *TaobaoTradeOrderskuUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Oid"] = req.Oid
    
    params["SkuId"] = req.SkuId
    
    params["SkuProps"] = req.SkuProps
    
    return params
}

// TaobaoTradeOrderskuUpdateResponse 只能更新发货前子订单的销售属性 
//只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 
//必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
type TaobaoTradeOrderskuUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Order Object只返回oid和modified
    Order Order `json:"order";xml:"order"`
    
}

// TaobaoQimenStoreprocessCreateRequest ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Materialitems optional加工商品列表
    Materialitems MaterialItem `json:"materialitems";xml:"materialitems"`
    
    // OrderCreateTime required加工单创建时间(YYYY-MM-DD HH:MM:SS)
    OrderCreateTime string `json:"orderCreateTime";xml:"orderCreateTime"`
    
    // OrderType required单据类型(CNJG=仓内加工作业单)
    OrderType string `json:"orderType";xml:"orderType"`
    
    // PlanQty optional成品计划数量
    PlanQty int64 `json:"planQty";xml:"planQty"`
    
    // PlanTime required计划加工时间(YYYY-MM-DD HH:MM:SS)
    PlanTime string `json:"planTime";xml:"planTime"`
    
    // ProcessOrderCode required加工单编码
    ProcessOrderCode string `json:"processOrderCode";xml:"processOrderCode"`
    
    // Productitems optional商品列表
    Productitems ProductItem `json:"productitems";xml:"productitems"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // ServiceType required加工类型(1:仓内组合加工 2:仓内组合拆分)
    ServiceType string `json:"serviceType";xml:"serviceType"`
    
    // WarehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStoreprocessCreateRequest) GetAPIName() string {
	return "taobao.qimen.storeprocess.create"
}

func (req *TaobaoQimenStoreprocessCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 11)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Materialitems"] = req.Materialitems
    
    params["OrderCreateTime"] = req.OrderCreateTime
    
    params["OrderType"] = req.OrderType
    
    params["PlanQty"] = req.PlanQty
    
    params["PlanTime"] = req.PlanTime
    
    params["ProcessOrderCode"] = req.ProcessOrderCode
    
    params["Productitems"] = req.Productitems
    
    params["Remark"] = req.Remark
    
    params["ServiceType"] = req.ServiceType
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenStoreprocessCreateResponse ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // ProcessOrderId Basic仓储系统处理单ID
    ProcessOrderId string `json:"processOrderId";xml:"processOrderId"`
    
}

// TaobaoTradeShippingaddressUpdateRequest 只能更新一笔交易里面的买家收货地址 
//只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 
//更新后的发货地址可以通过taobao.trade.fullinfo.get查到 
//参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaoTradeShippingaddressUpdateRequest struct {
    
    // ReceiverAddress optional收货地址。最大长度为228个字节。
    ReceiverAddress string `json:"receiver_address";xml:"receiver_address"`
    
    // ReceiverCity optional城市。最大长度为32个字节。如：杭州
    ReceiverCity string `json:"receiver_city";xml:"receiver_city"`
    
    // ReceiverDistrict optional区/县。最大长度为32个字节。如：西湖区
    ReceiverDistrict string `json:"receiver_district";xml:"receiver_district"`
    
    // ReceiverMobile optional移动电话。最大长度为11个字节。
    ReceiverMobile string `json:"receiver_mobile";xml:"receiver_mobile"`
    
    // ReceiverName optional收货人全名。最大长度为50个字节。
    ReceiverName string `json:"receiver_name";xml:"receiver_name"`
    
    // ReceiverPhone optional固定电话。最大长度为30个字节。
    ReceiverPhone string `json:"receiver_phone";xml:"receiver_phone"`
    
    // ReceiverState optional省份。最大长度为32个字节。如：浙江
    ReceiverState string `json:"receiver_state";xml:"receiver_state"`
    
    // ReceiverZip optional邮政编码。必须由6个数字组成。
    ReceiverZip string `json:"receiver_zip";xml:"receiver_zip"`
    
    // Tid required交易编号。
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeShippingaddressUpdateRequest) GetAPIName() string {
	return "taobao.trade.shippingaddress.update"
}

func (req *TaobaoTradeShippingaddressUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["ReceiverAddress"] = req.ReceiverAddress
    
    params["ReceiverCity"] = req.ReceiverCity
    
    params["ReceiverDistrict"] = req.ReceiverDistrict
    
    params["ReceiverMobile"] = req.ReceiverMobile
    
    params["ReceiverName"] = req.ReceiverName
    
    params["ReceiverPhone"] = req.ReceiverPhone
    
    params["ReceiverState"] = req.ReceiverState
    
    params["ReceiverZip"] = req.ReceiverZip
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeShippingaddressUpdateResponse 只能更新一笔交易里面的买家收货地址 
//只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 
//更新后的发货地址可以通过taobao.trade.fullinfo.get查到 
//参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaoTradeShippingaddressUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Trade Object交易结构
    Trade Trade `json:"trade";xml:"trade"`
    
}

// TaobaoQimenStockoutQueryRequest 查询出库单详情
type TaobaoQimenStockoutQueryRequest struct {
    
    // DeliveryOrderCode required出库单号
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // DeliveryOrderId optional仓储系统出库单号
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页orderLine条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockoutQueryRequest) GetAPIName() string {
	return "taobao.qimen.stockout.query"
}

func (req *TaobaoQimenStockoutQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["DeliveryOrderCode"] = req.DeliveryOrderCode
    
    params["DeliveryOrderId"] = req.DeliveryOrderId
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenStockoutQueryResponse 查询出库单详情
type TaobaoQimenStockoutQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // DeliveryOrder Object出库单信息
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderLines Object Array单据信息列表
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // Packages Object Array包裹信息
    Packages Package `json:"packages";xml:"packages"`
    
    // TotalLines BasicorderLines总条数
    TotalLines int64 `json:"totalLines";xml:"totalLines"`
    
}

// TaobaoInventoryInitialItemRequest 商家仓商品初始化在各个仓中库存
type TaobaoInventoryInitialItemRequest struct {
    
    // ScItemId required后端商品id
    ScItemId int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    // StoreInventorys required商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
    StoreInventorys string `json:"store_inventorys";xml:"store_inventorys"`
    
}

func (req *TaobaoInventoryInitialItemRequest) GetAPIName() string {
	return "taobao.inventory.initial.item"
}

func (req *TaobaoInventoryInitialItemRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ScItemId"] = req.ScItemId
    
    params["StoreInventorys"] = req.StoreInventorys
    
    return params
}

// TaobaoInventoryInitialItemResponse 商家仓商品初始化在各个仓中库存
type TaobaoInventoryInitialItemResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TipInfos Object Array提示信息
    TipInfos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

// TaobaoPicturePicturesCountRequest 图片总数查询
type TaobaoPicturePicturesCountRequest struct {
    
    // ClientType optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    ClientType string `json:"client_type";xml:"client_type"`
    
    // Deleted optional是否删除,undeleted代表没有删除,deleted表示删除
    Deleted string `json:"deleted";xml:"deleted"`
    
    // EndDate optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // PictureCategoryId optional图片分类
    PictureCategoryId int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    // PictureId optional图片ID
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
    // StartDate optional查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // StartModifiedDate optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    StartModifiedDate time.Time `json:"start_modified_date";xml:"start_modified_date"`
    
    // Title optional文件名
    Title string `json:"title";xml:"title"`
    
}

func (req *TaobaoPicturePicturesCountRequest) GetAPIName() string {
	return "taobao.picture.pictures.count"
}

func (req *TaobaoPicturePicturesCountRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["ClientType"] = req.ClientType
    
    params["Deleted"] = req.Deleted
    
    params["EndDate"] = req.EndDate
    
    params["PictureCategoryId"] = req.PictureCategoryId
    
    params["PictureId"] = req.PictureId
    
    params["StartDate"] = req.StartDate
    
    params["StartModifiedDate"] = req.StartModifiedDate
    
    params["Title"] = req.Title
    
    return params
}

// TaobaoPicturePicturesCountResponse 图片总数查询
type TaobaoPicturePicturesCountResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Totals Basic查询的文件总数
    Totals int64 `json:"totals";xml:"totals"`
    
}

// TaobaoJushitaJdpUserDeleteRequest 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteRequest struct {
    
    // Nick special要删除用户的昵称
    Nick string `json:"nick";xml:"nick"`
    
    // UserId special需要删除的用户编号
    UserId int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoJushitaJdpUserDeleteRequest) GetAPIName() string {
	return "taobao.jushita.jdp.user.delete"
}

func (req *TaobaoJushitaJdpUserDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Nick"] = req.Nick
    
    params["UserId"] = req.UserId
    
    return params
}

// TaobaoJushitaJdpUserDeleteResponse 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否删除成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoPicturePicturesGetRequest 图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
type TaobaoPicturePicturesGetRequest struct {
    
    // ClientType optional图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    ClientType string `json:"client_type";xml:"client_type"`
    
    // CurrentPage optional页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
    CurrentPage int64 `json:"current_page";xml:"current_page"`
    
    // Deleted optional是否删除,deleted代表没有删除
    Deleted string `json:"deleted";xml:"deleted"`
    
    // EndDate optional结束时间
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // IsHttps optional是否获取https的链接
    IsHttps bool `json:"is_https";xml:"is_https"`
    
    // OrderBy optional图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
    OrderBy string `json:"order_by";xml:"order_by"`
    
    // PageSize optional每页条数.每页返回最多返回100条,默认值40
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // PictureCategoryId optional图片分类
    PictureCategoryId int64 `json:"picture_category_id";xml:"picture_category_id"`
    
    // PictureId optional图片ID
    PictureId int64 `json:"picture_id";xml:"picture_id"`
    
    // StartDate optional查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // StartModifiedDate optional图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    StartModifiedDate time.Time `json:"start_modified_date";xml:"start_modified_date"`
    
    // Title optional图片标题,最大长度50字符,中英文都算一字符
    Title string `json:"title";xml:"title"`
    
    // Urls optional图片url查询接口
    Urls string `json:"urls";xml:"urls"`
    
}

func (req *TaobaoPicturePicturesGetRequest) GetAPIName() string {
	return "taobao.picture.pictures.get"
}

func (req *TaobaoPicturePicturesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 13)
    
    params["ClientType"] = req.ClientType
    
    params["CurrentPage"] = req.CurrentPage
    
    params["Deleted"] = req.Deleted
    
    params["EndDate"] = req.EndDate
    
    params["IsHttps"] = req.IsHttps
    
    params["OrderBy"] = req.OrderBy
    
    params["PageSize"] = req.PageSize
    
    params["PictureCategoryId"] = req.PictureCategoryId
    
    params["PictureId"] = req.PictureId
    
    params["StartDate"] = req.StartDate
    
    params["StartModifiedDate"] = req.StartModifiedDate
    
    params["Title"] = req.Title
    
    params["Urls"] = req.Urls
    
    return params
}

// TaobaoPicturePicturesGetResponse 图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
type TaobaoPicturePicturesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Pictures Object Array图片空间图片数据对象
    Pictures Picture `json:"pictures";xml:"pictures"`
    
}

// TaobaoJushitaJdpUserAddRequest 提供给接入数据推送的应用添加数据推送服务的用户
type TaobaoJushitaJdpUserAddRequest struct {
    
    // HistoryDays optional推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。
    HistoryDays int64 `json:"history_days";xml:"history_days"`
    
    // HostIp optional已废弃，新版订单同步服务不要使用。同步用户数据的机器IP,必须是界面配置的IP。
    HostIp string `json:"host_ip";xml:"host_ip"`
    
    // HostName optional已废弃，新版订单同步服务不要使用。绑定类型，用户数据同步的机器名称。
    HostName string `json:"host_name";xml:"host_name"`
    
    // RdsName requiredRDS实例名称
    RdsName string `json:"rds_name";xml:"rds_name"`
    
    // Topics optional已废弃，使用页面中应用的配置。用户同步的数据类型,多个用','号分割。可选值：trade,refund,item
    Topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoJushitaJdpUserAddRequest) GetAPIName() string {
	return "taobao.jushita.jdp.user.add"
}

func (req *TaobaoJushitaJdpUserAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["HistoryDays"] = req.HistoryDays
    
    params["HostIp"] = req.HostIp
    
    params["HostName"] = req.HostName
    
    params["RdsName"] = req.RdsName
    
    params["Topics"] = req.Topics
    
    return params
}

// TaobaoJushitaJdpUserAddResponse 提供给接入数据推送的应用添加数据推送服务的用户
type TaobaoJushitaJdpUserAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否添加成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// CainiaoCloudprintTemplatesMigrateRequest 云打印模板迁移接口
type CainiaoCloudprintTemplatesMigrateRequest struct {
    
    // CustomAreaContent optional自定义区内容
    CustomAreaContent string `json:"custom_area_content";xml:"custom_area_content"`
    
    // CustomAreaName optional自定义区名称
    CustomAreaName string `json:"custom_area_name";xml:"custom_area_name"`
    
    // TempalteId optional标准电子面单模板的id
    TempalteId int64 `json:"tempalte_id";xml:"tempalte_id"`
    
}

func (req *CainiaoCloudprintTemplatesMigrateRequest) GetAPIName() string {
	return "cainiao.cloudprint.templates.migrate"
}

func (req *CainiaoCloudprintTemplatesMigrateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CustomAreaContent"] = req.CustomAreaContent
    
    params["CustomAreaName"] = req.CustomAreaName
    
    params["TempalteId"] = req.TempalteId
    
    return params
}

// CainiaoCloudprintTemplatesMigrateResponse 云打印模板迁移接口
type CainiaoCloudprintTemplatesMigrateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// AlipayXiaodaiUserPermitRequest 阿里金融为用户开通消息通道接口
type AlipayXiaodaiUserPermitRequest struct {
    
    // UserId required用户数字ID
    UserId int64 `json:"user_id";xml:"user_id"`
    
}

func (req *AlipayXiaodaiUserPermitRequest) GetAPIName() string {
	return "alipay.xiaodai.user.permit"
}

func (req *AlipayXiaodaiUserPermitRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["UserId"] = req.UserId
    
    return params
}

// AlipayXiaodaiUserPermitResponse 阿里金融为用户开通消息通道接口
type AlipayXiaodaiUserPermitResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoAppstoreSubscribeGetRequest 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
type TaobaoAppstoreSubscribeGetRequest struct {
    
    // LeaseId optional插件实例ID
    LeaseId int64 `json:"lease_id";xml:"lease_id"`
    
    // Nick required用户昵称
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoAppstoreSubscribeGetRequest) GetAPIName() string {
	return "taobao.appstore.subscribe.get"
}

func (req *TaobaoAppstoreSubscribeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["LeaseId"] = req.LeaseId
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoAppstoreSubscribeGetResponse 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
type TaobaoAppstoreSubscribeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UserSubscribe Object用户订购信息
    UserSubscribe UserSubscribe `json:"user_subscribe";xml:"user_subscribe"`
    
}

// TaobaoQimenReplenishplanCreateRequest ERP调用奇门的接口,通知WMS创建补货计划
type TaobaoQimenReplenishplanCreateRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // GmtDeadTime required最晚入库时间(YYYY-MM-DD HH:MM:SS)
    GmtDeadTime string `json:"gmtDeadTime";xml:"gmtDeadTime"`
    
    // Items optional单据信息
    Items ReplenishplanCreateItem `json:"items";xml:"items"`
    
    // OrderCode required外部系统单号(ERP单号)
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // UserId required商家ID
    UserId string `json:"userId";xml:"userId"`
    
    // WarehouseCode required仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenReplenishplanCreateRequest) GetAPIName() string {
	return "taobao.qimen.replenishplan.create"
}

func (req *TaobaoQimenReplenishplanCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["GmtDeadTime"] = req.GmtDeadTime
    
    params["Items"] = req.Items
    
    params["OrderCode"] = req.OrderCode
    
    params["UserId"] = req.UserId
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenReplenishplanCreateResponse ERP调用奇门的接口,通知WMS创建补货计划
type TaobaoQimenReplenishplanCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TmallDisputeReceiveGetRequest 展示商家所有退款信息
type TmallDisputeReceiveGetRequest struct {
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // EndModified optional查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // Fields required需要返回的字段。目前支持有：refund_id, alipay_no, tid, buyer_nick, seller_nick, status, created, modified, order_status, refund_fee, good_status, show_return_logistic(展现买家退货的物流信息), show_exchange_logistic(展现换货的物流信息), time_out, oid, refund_version, title, num, dispute_request, reason, desc
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。取值范围:大于零的整数; 默认值:1
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围:大于零的整数; 默认值:20;最大值:100
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // RefundId optional逆向纠纷单号id
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // StartModified optional查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
    // Status optional退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);CLOSED(退款关闭); SUCCESS(退款成功);SELLER_REFUSE_BUYER(卖家拒绝退款);WAIT_BUYER_CONFIRM_REDO_SEND_GOODS(等待买家确认重新邮寄的货物);WAIT_SELLER_CONFIRM_RETURN_ADDRESS(等待卖家确认退货地址);WAIT_SELLER_CONSIGN_GOOGDS(卖家确认收货,等待卖家发货);EXCHANGE_TRANSFORM_TO_REFUND(换货关闭,转退货退款);EXCHANGE_WAIT_BUYER_CONFIRM_GOODS(卖家已发货,等待买家确认收货);POST_FEE_DISPUTE_WAIT_ACTIVATE(邮费单已创建,待激活)
    Status string `json:"status";xml:"status"`
    
    // Type optional交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，查看可选值
    Type string `json:"type";xml:"type"`
    
    // UseHasNext optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
    UseHasNext bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TmallDisputeReceiveGetRequest) GetAPIName() string {
	return "tmall.dispute.receive.get"
}

func (req *TmallDisputeReceiveGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["BuyerNick"] = req.BuyerNick
    
    params["EndModified"] = req.EndModified
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["RefundId"] = req.RefundId
    
    params["StartModified"] = req.StartModified
    
    params["Status"] = req.Status
    
    params["Type"] = req.Type
    
    params["UseHasNext"] = req.UseHasNext
    
    return params
}

// TmallDisputeReceiveGetResponse 展示商家所有退款信息
type TmallDisputeReceiveGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result ResultSet `json:"result";xml:"result"`
    
}

// TaobaoLogisticsPartnersGetRequest 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaoLogisticsPartnersGetRequest struct {
    
    // GoodsValue optional货物价格.只有当选择货到付款此参数才会有效
    GoodsValue string `json:"goods_value";xml:"goods_value"`
    
    // IsNeedCarriage optional是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。
    IsNeedCarriage bool `json:"is_need_carriage";xml:"is_need_carriage"`
    
    // ServiceType optional服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。
    ServiceType string `json:"service_type";xml:"service_type"`
    
    // SourceId optional物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
    SourceId string `json:"source_id";xml:"source_id"`
    
    // TargetId optional物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
    TargetId string `json:"target_id";xml:"target_id"`
    
}

func (req *TaobaoLogisticsPartnersGetRequest) GetAPIName() string {
	return "taobao.logistics.partners.get"
}

func (req *TaobaoLogisticsPartnersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["GoodsValue"] = req.GoodsValue
    
    params["IsNeedCarriage"] = req.IsNeedCarriage
    
    params["ServiceType"] = req.ServiceType
    
    params["SourceId"] = req.SourceId
    
    params["TargetId"] = req.TargetId
    
    return params
}

// TaobaoLogisticsPartnersGetResponse 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaoLogisticsPartnersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // LogisticsPartners Object Array查询揽送范围之内的物流公司信息
    LogisticsPartners LogisticsPartner `json:"logistics_partners";xml:"logistics_partners"`
    
}

// TaobaoTopSdkFeedbackUploadRequest sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaoTopSdkFeedbackUploadRequest struct {
    
    // Content optional具体内容，json形式
    Content string `json:"content";xml:"content"`
    
    // Type required1、回传加密信息
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoTopSdkFeedbackUploadRequest) GetAPIName() string {
	return "taobao.top.sdk.feedback.upload"
}

func (req *TaobaoTopSdkFeedbackUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Content"] = req.Content
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoTopSdkFeedbackUploadResponse sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaoTopSdkFeedbackUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UploadInterval Basic控制回传间隔（单位：秒）
    UploadInterval int64 `json:"upload_interval";xml:"upload_interval"`
    
}

// TaobaoSubuserDutysGetRequest 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaoSubuserDutysGetRequest struct {
    
    // UserNick required主账号用户名
    UserNick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubuserDutysGetRequest) GetAPIName() string {
	return "taobao.subuser.dutys.get"
}

func (req *TaobaoSubuserDutysGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["UserNick"] = req.UserNick
    
    return params
}

// TaobaoSubuserDutysGetResponse 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaoSubuserDutysGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Dutys Object Array职务信息
    Dutys Duty `json:"dutys";xml:"dutys"`
    
}

// TaobaoSubuserEmployeeAddRequest 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息）
type TaobaoSubuserEmployeeAddRequest struct {
    
    // DepartmentId required当前员工所属部门ID
    DepartmentId int64 `json:"department_id";xml:"department_id"`
    
    // DutyId optional当前员工担任职务ID
    DutyId int64 `json:"duty_id";xml:"duty_id"`
    
    // EmployeeName required员工姓名
    EmployeeName string `json:"employee_name";xml:"employee_name"`
    
    // EmployeeNickname optional员工花名
    EmployeeNickname string `json:"employee_nickname";xml:"employee_nickname"`
    
    // EmployeeNum optional员工工号(可以用大小写英文字母和数字)
    EmployeeNum string `json:"employee_num";xml:"employee_num"`
    
    // EntryDate optional员工入职时间
    EntryDate time.Time `json:"entry_date";xml:"entry_date"`
    
    // IdCardNum optional员工身份证号码
    IdCardNum string `json:"id_card_num";xml:"id_card_num"`
    
    // LeaderId optional直接上级的员工ID
    LeaderId int64 `json:"leader_id";xml:"leader_id"`
    
    // OfficePhone optional办公电话
    OfficePhone string `json:"office_phone";xml:"office_phone"`
    
    // PersonalEmail optional员工私人邮箱
    PersonalEmail string `json:"personal_email";xml:"personal_email"`
    
    // PersonalMobile optional员工手机号码
    PersonalMobile string `json:"personal_mobile";xml:"personal_mobile"`
    
    // Sex required员工性别 1：男; 2:女
    Sex int64 `json:"sex";xml:"sex"`
    
    // SubId required子账号ID
    SubId int64 `json:"sub_id";xml:"sub_id"`
    
    // WorkLocation optional工作地点
    WorkLocation string `json:"work_location";xml:"work_location"`
    
}

func (req *TaobaoSubuserEmployeeAddRequest) GetAPIName() string {
	return "taobao.subuser.employee.add"
}

func (req *TaobaoSubuserEmployeeAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 14)
    
    params["DepartmentId"] = req.DepartmentId
    
    params["DutyId"] = req.DutyId
    
    params["EmployeeName"] = req.EmployeeName
    
    params["EmployeeNickname"] = req.EmployeeNickname
    
    params["EmployeeNum"] = req.EmployeeNum
    
    params["EntryDate"] = req.EntryDate
    
    params["IdCardNum"] = req.IdCardNum
    
    params["LeaderId"] = req.LeaderId
    
    params["OfficePhone"] = req.OfficePhone
    
    params["PersonalEmail"] = req.PersonalEmail
    
    params["PersonalMobile"] = req.PersonalMobile
    
    params["Sex"] = req.Sex
    
    params["SubId"] = req.SubId
    
    params["WorkLocation"] = req.WorkLocation
    
    return params
}

// TaobaoSubuserEmployeeAddResponse 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息）
type TaobaoSubuserEmployeeAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功 true:操作成功; false:操作失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoSubuserInfoUpdateRequest 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
type TaobaoSubuserInfoUpdateRequest struct {
    
    // IsDisableSubaccount optional是否停用子账号 true:表示停用该子账号false:表示开启该子账号
    IsDisableSubaccount bool `json:"is_disable_subaccount";xml:"is_disable_subaccount"`
    
    // IsDispatch optional子账号是否参与分流 true:参与分流 false:不参与分流
    IsDispatch bool `json:"is_dispatch";xml:"is_dispatch"`
    
    // SubId required子账号ID
    SubId int64 `json:"sub_id";xml:"sub_id"`
    
}

func (req *TaobaoSubuserInfoUpdateRequest) GetAPIName() string {
	return "taobao.subuser.info.update"
}

func (req *TaobaoSubuserInfoUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["IsDisableSubaccount"] = req.IsDisableSubaccount
    
    params["IsDispatch"] = req.IsDispatch
    
    params["SubId"] = req.SubId
    
    return params
}

// TaobaoSubuserInfoUpdateResponse 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
type TaobaoSubuserInfoUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功 true:操作成功; false:操作失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// Gfq1qp4287ShopexMatrxiOfflineSendRequest 商派矩阵线下发货接口
type Gfq1qp4287ShopexMatrxiOfflineSendRequest struct {
    
    // Params optional参数
    Params string `json:"params";xml:"params"`
    
}

func (req *Gfq1qp4287ShopexMatrxiOfflineSendRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.matrxi.offline.send"
}

func (req *Gfq1qp4287ShopexMatrxiOfflineSendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Params"] = req.Params
    
    return params
}

// Gfq1qp4287ShopexMatrxiOfflineSendResponse 商派矩阵线下发货接口
type Gfq1qp4287ShopexMatrxiOfflineSendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // MsgId Basic任务id
    MsgId string `json:"msg_id";xml:"msg_id"`
    
}

// TaobaoUopTobOrderCreateRequest ToB仓储发货
type TaobaoUopTobOrderCreateRequest struct {
    
    // DeliveryOrder optionalERP出库对象
    DeliveryOrder DeliveryOrder `json:"delivery_order";xml:"delivery_order"`
    
}

func (req *TaobaoUopTobOrderCreateRequest) GetAPIName() string {
	return "taobao.uop.tob.order.create"
}

func (req *TaobaoUopTobOrderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["DeliveryOrder"] = req.DeliveryOrder
    
    return params
}

// TaobaoUopTobOrderCreateResponse ToB仓储发货
type TaobaoUopTobOrderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DeliveryOrders Object Array订单
    DeliveryOrders Deliveryorder `json:"delivery_orders";xml:"delivery_orders"`
    
    // Flag Basicflag
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basicmessage
    Message string `json:"message";xml:"message"`
    
}

// QimenTaobaoUopTmsupdatemessagetoerpRequest ERP配消息回告（拒签、签收、揽收等）
type QimenTaobaoUopTmsupdatemessagetoerpRequest struct {
    
    // Request optional
    Request Struct `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoUopTmsupdatemessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.uop.tmsupdatemessagetoerp"
}

func (req *QimenTaobaoUopTmsupdatemessagetoerpRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoUopTmsupdatemessagetoerpResponse ERP配消息回告（拒签、签收、揽收等）
type QimenTaobaoUopTmsupdatemessagetoerpResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Objectresponse
    Response Response `json:"response";xml:"response"`
    
}

// QimenTaobaoIcpOrderStockoutordermessagetoerpRequest 出库单信息推送接口
type QimenTaobaoIcpOrderStockoutordermessagetoerpRequest struct {
    
    // CustomerId required货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // EntryOutOrderlist optional出库单记录集
    EntryOutOrderlist EntryOutOrderlist `json:"entryOutOrderlist";xml:"entryOutOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockoutordermessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockoutordermessagetoerp"
}

func (req *QimenTaobaoIcpOrderStockoutordermessagetoerpRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["EntryOutOrderlist"] = req.EntryOutOrderlist
    
    return params
}

// QimenTaobaoIcpOrderStockoutordermessagetoerpResponse 出库单信息推送接口
type QimenTaobaoIcpOrderStockoutordermessagetoerpResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoItemSkuDeleteRequest 删除一个sku的数据
//需要删除的sku通过属性properties进行匹配查找
type TaobaoItemSkuDeleteRequest struct {
    
    // Ignorewarning optional忽略警告提示.
    Ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    // ItemNum optionalsku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败
    ItemNum int64 `json:"item_num";xml:"item_num"`
    
    // ItemPrice optionalsku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功
    ItemPrice float32 `json:"item_price";xml:"item_price"`
    
    // Lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    Lang string `json:"lang";xml:"lang"`
    
    // NumIid requiredSku所属商品数字id，可通过 taobao.item.get 获取。必选
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // Properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
    Properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemSkuDeleteRequest) GetAPIName() string {
	return "taobao.item.sku.delete"
}

func (req *TaobaoItemSkuDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Ignorewarning"] = req.Ignorewarning
    
    params["ItemNum"] = req.ItemNum
    
    params["ItemPrice"] = req.ItemPrice
    
    params["Lang"] = req.Lang
    
    params["NumIid"] = req.NumIid
    
    params["Properties"] = req.Properties
    
    return params
}

// TaobaoItemSkuDeleteResponse 删除一个sku的数据
//需要删除的sku通过属性properties进行匹配查找
type TaobaoItemSkuDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Sku ObjectSku结构
    Sku Sku `json:"sku";xml:"sku"`
    
}

// AlibabaEinvoiceProviderlistGetRequest 获取能提供电子发票开票服务的开票服务商
type AlibabaEinvoiceProviderlistGetRequest struct {
    
}

func (req *AlibabaEinvoiceProviderlistGetRequest) GetAPIName() string {
	return "alibaba.einvoice.providerlist.get"
}

func (req *AlibabaEinvoiceProviderlistGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// AlibabaEinvoiceProviderlistGetResponse 获取能提供电子发票开票服务的开票服务商
type AlibabaEinvoiceProviderlistGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object查询结果集
    Result ResultSet `json:"result";xml:"result"`
    
}

// TaobaoTmcUserGetRequest 查询指定用户开通的消息通道和组
type TaobaoTmcUserGetRequest struct {
    
    // Fields required需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
    Fields string `json:"fields";xml:"fields"`
    
    // Nick required用户昵称
    Nick string `json:"nick";xml:"nick"`
    
    // UserPlatform optional用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    UserPlatform string `json:"user_platform";xml:"user_platform"`
    
}

func (req *TaobaoTmcUserGetRequest) GetAPIName() string {
	return "taobao.tmc.user.get"
}

func (req *TaobaoTmcUserGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Fields"] = req.Fields
    
    params["Nick"] = req.Nick
    
    params["UserPlatform"] = req.UserPlatform
    
    return params
}

// TaobaoTmcUserGetResponse 查询指定用户开通的消息通道和组
type TaobaoTmcUserGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TmcUser Object开通的用户数据
    TmcUser TmcUser `json:"tmc_user";xml:"tmc_user"`
    
}

// CainiaoWaybillprintClientupdateGetconfigRequest 获取客户端更新配置信息
type CainiaoWaybillprintClientupdateGetconfigRequest struct {
    
    // LanIp required服务发起机器局域网ip
    LanIp string `json:"lan_ip";xml:"lan_ip"`
    
    // Mac required服务发起机器的物理地址
    Mac string `json:"mac";xml:"mac"`
    
    // Msgid required当前消息版本
    Msgid int64 `json:"msgid";xml:"msgid"`
    
    // Version required当前打印客户端版本
    Version string `json:"version";xml:"version"`
    
}

func (req *CainiaoWaybillprintClientupdateGetconfigRequest) GetAPIName() string {
	return "cainiao.waybillprint.clientupdate.getconfig"
}

func (req *CainiaoWaybillprintClientupdateGetconfigRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["LanIp"] = req.LanIp
    
    params["Mac"] = req.Mac
    
    params["Msgid"] = req.Msgid
    
    params["Version"] = req.Version
    
    return params
}

// CainiaoWaybillprintClientupdateGetconfigResponse 获取客户端更新配置信息
type CainiaoWaybillprintClientupdateGetconfigResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result UpdateConfigTopResult `json:"result";xml:"result"`
    
}

// TmallTraderateItemtagsGetRequest 通过商品ID获取标签详细信息
type TmallTraderateItemtagsGetRequest struct {
    
    // ItemId required商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TmallTraderateItemtagsGetRequest) GetAPIName() string {
	return "tmall.traderate.itemtags.get"
}

func (req *TmallTraderateItemtagsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TmallTraderateItemtagsGetResponse 通过商品ID获取标签详细信息
type TmallTraderateItemtagsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Tags Object Array标签列表
    Tags TmallRateTagDetail `json:"tags";xml:"tags"`
    
}

// TmallTraderateFeedsGetRequest 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
type TmallTraderateFeedsGetRequest struct {
    
    // ChildTradeId required交易子订单ID
    ChildTradeId int64 `json:"child_trade_id";xml:"child_trade_id"`
    
}

func (req *TmallTraderateFeedsGetRequest) GetAPIName() string {
	return "tmall.traderate.feeds.get"
}

func (req *TmallTraderateFeedsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ChildTradeId"] = req.ChildTradeId
    
    return params
}

// TmallTraderateFeedsGetResponse 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
type TmallTraderateFeedsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TmallRateInfo Object返回评价信息
    TmallRateInfo TmallRateInfo `json:"tmall_rate_info";xml:"tmall_rate_info"`
    
}

// TaobaoFilesGetRequest 获取业务方暂存给ISV的文件列表
type TaobaoFilesGetRequest struct {
    
    // EndDate required搜索结束时间
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // StartDate required搜索开始时间
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // Status optional下载链接状态。1:未下载。2:已下载
    Status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoFilesGetRequest) GetAPIName() string {
	return "taobao.files.get"
}

func (req *TaobaoFilesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["EndDate"] = req.EndDate
    
    params["StartDate"] = req.StartDate
    
    params["Status"] = req.Status
    
    return params
}

// TaobaoFilesGetResponse 获取业务方暂存给ISV的文件列表
type TaobaoFilesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Results Object Arrayresults
    Results TopDownloadRecordDo `json:"results";xml:"results"`
    
}

// TaobaoWlbWaybillShengxianGetRequest 商家通过交易订单号获取电子面单接口
type TaobaoWlbWaybillShengxianGetRequest struct {
    
    // BizCode required物流服务方代码，生鲜配送：YXSR
    BizCode string `json:"biz_code";xml:"biz_code"`
    
    // DeliveryType required物流服务类型。冷链：602，常温：502
    DeliveryType string `json:"delivery_type";xml:"delivery_type"`
    
    // Feature optional预留扩展字段
    Feature string `json:"feature";xml:"feature"`
    
    // OrderChannelsType required订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688";
    OrderChannelsType string `json:"order_channels_type";xml:"order_channels_type"`
    
    // SenderAddressId optional商家淘宝地址信息ID
    SenderAddressId string `json:"sender_address_id";xml:"sender_address_id"`
    
    // ServiceCode optional仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
    ServiceCode string `json:"service_code";xml:"service_code"`
    
    // TradeId required交易号，传入交易号不能存在拆单场景。
    TradeId string `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoWlbWaybillShengxianGetRequest) GetAPIName() string {
	return "taobao.wlb.waybill.shengxian.get"
}

func (req *TaobaoWlbWaybillShengxianGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["BizCode"] = req.BizCode
    
    params["DeliveryType"] = req.DeliveryType
    
    params["Feature"] = req.Feature
    
    params["OrderChannelsType"] = req.OrderChannelsType
    
    params["SenderAddressId"] = req.SenderAddressId
    
    params["ServiceCode"] = req.ServiceCode
    
    params["TradeId"] = req.TradeId
    
    return params
}

// TaobaoWlbWaybillShengxianGetResponse 商家通过交易订单号获取电子面单接口
type TaobaoWlbWaybillShengxianGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // FreshWaybill Object成功后返回的生鲜电子面单信息
    FreshWaybill FreshWaybill `json:"fresh_waybill";xml:"fresh_waybill"`
    
    // IsSuccess Basic生成是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoProductQuantityUpdateRequest 修改产品库存信息，支持全量修改以及增量修改两种方式
type TaobaoFenxiaoProductQuantityUpdateRequest struct {
    
    // ProductId required产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Properties optionalsku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。
    Properties string `json:"properties";xml:"properties"`
    
    // Quantity required库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。
    // 当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
    Quantity string `json:"quantity";xml:"quantity"`
    
    // Type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoFenxiaoProductQuantityUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.quantity.update"
}

func (req *TaobaoFenxiaoProductQuantityUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ProductId"] = req.ProductId
    
    params["Properties"] = req.Properties
    
    params["Quantity"] = req.Quantity
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoFenxiaoProductQuantityUpdateResponse 修改产品库存信息，支持全量修改以及增量修改两种方式
type TaobaoFenxiaoProductQuantityUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic操作时间
    Created time.Time `json:"created";xml:"created"`
    
    // Result Basic操作结果
    Result bool `json:"result";xml:"result"`
    
}

// CainiaoOpenstorageSellerResourceCreateRequest 商家资源创建接口(云打印开源存储)
type CainiaoOpenstorageSellerResourceCreateRequest struct {
    
    // ParamCreateSellerResourceRequest optional商家创建资源参数
    ParamCreateSellerResourceRequest CreateSellerResourceRequest `json:"param_create_seller_resource_request";xml:"param_create_seller_resource_request"`
    
}

func (req *CainiaoOpenstorageSellerResourceCreateRequest) GetAPIName() string {
	return "cainiao.openstorage.seller.resource.create"
}

func (req *CainiaoOpenstorageSellerResourceCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamCreateSellerResourceRequest"] = req.ParamCreateSellerResourceRequest
    
    return params
}

// CainiaoOpenstorageSellerResourceCreateResponse 商家资源创建接口(云打印开源存储)
type CainiaoOpenstorageSellerResourceCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// CainiaoOpenstorageResourcePublishRequest isv和商家资源发布接口(云打印开源存储)
type CainiaoOpenstorageResourcePublishRequest struct {
    
    // ParamPublishResourceRequest optional资源发布参数
    ParamPublishResourceRequest PublishResourceRequest `json:"param_publish_resource_request";xml:"param_publish_resource_request"`
    
}

func (req *CainiaoOpenstorageResourcePublishRequest) GetAPIName() string {
	return "cainiao.openstorage.resource.publish"
}

func (req *CainiaoOpenstorageResourcePublishRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamPublishResourceRequest"] = req.ParamPublishResourceRequest
    
    return params
}

// CainiaoOpenstorageResourcePublishResponse isv和商家资源发布接口(云打印开源存储)
type CainiaoOpenstorageResourcePublishResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// AlipayPointBudgetGetRequest 查询已采购的集分宝余额，操作流程如下：
//1、申请支付宝增值包。
//2、申请支付宝应用上线时选择集分宝API。
//3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
//4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
//5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
//6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
type AlipayPointBudgetGetRequest struct {
    
    // AuthToken special支付宝给用户的授权。如果没有top的授权，这个字段是必填项
    AuthToken string `json:"auth_token";xml:"auth_token"`
    
}

func (req *AlipayPointBudgetGetRequest) GetAPIName() string {
	return "alipay.point.budget.get"
}

func (req *AlipayPointBudgetGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["AuthToken"] = req.AuthToken
    
    return params
}

// AlipayPointBudgetGetResponse 查询已采购的集分宝余额，操作流程如下：
//1、申请支付宝增值包。
//2、申请支付宝应用上线时选择集分宝API。
//3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
//4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
//5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
//6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
type AlipayPointBudgetGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // BudgetAmount Basic还可以发放的集分宝个数
    BudgetAmount int64 `json:"budget_amount";xml:"budget_amount"`
    
}

// CainiaoOpenstorageResourceUpdateRequest isv和商家的资源获取接口（云打印开源存储）
type CainiaoOpenstorageResourceUpdateRequest struct {
    
    // ParamUpdateResourceRequest required入参
    ParamUpdateResourceRequest UpdateResourceRequest `json:"param_update_resource_request";xml:"param_update_resource_request"`
    
}

func (req *CainiaoOpenstorageResourceUpdateRequest) GetAPIName() string {
	return "cainiao.openstorage.resource.update"
}

func (req *CainiaoOpenstorageResourceUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamUpdateResourceRequest"] = req.ParamUpdateResourceRequest
    
    return params
}

// CainiaoOpenstorageResourceUpdateResponse isv和商家的资源获取接口（云打印开源存储）
type CainiaoOpenstorageResourceUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// CainiaoOpenstorageIsvResourcesGetRequest isv资源列表查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcesGetRequest struct {
    
    // IsvResourceType requiredisv资源类型，分为：TEMPLATE（模板）和PRINT_ITEM（打印项）
    IsvResourceType string `json:"isv_resource_type";xml:"isv_resource_type"`
    
}

func (req *CainiaoOpenstorageIsvResourcesGetRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resources.get"
}

func (req *CainiaoOpenstorageIsvResourcesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["IsvResourceType"] = req.IsvResourceType
    
    return params
}

// CainiaoOpenstorageIsvResourcesGetResponse isv资源列表查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// AlipayPointOrderAddRequest 向用户发送集分宝，发放集分宝之前，操作流程如下：
//1、申请支付宝增值包。
//2、申请支付宝应用上线时选择集分宝API。
//3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
//4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
//5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
//6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
//7、调用发放API（alipay.point.order.add）发放集分宝。
type AlipayPointOrderAddRequest struct {
    
    // AuthToken special支付宝用户给应用发放集分宝的授权。
    AuthToken string `json:"auth_token";xml:"auth_token"`
    
    // Memo required向用户展示集分宝发放备注
    Memo string `json:"memo";xml:"memo"`
    
    // MerchantOrderNo requiredisv提供的发放号订单号，由数字和组成，最大长度为32为，需要保证每笔发放的唯一性，支付宝会对该参数做唯一性控制。如果使用同样的订单号，支付宝将返回订单号已经存在的错误
    MerchantOrderNo string `json:"merchant_order_no";xml:"merchant_order_no"`
    
    // OrderTime required发放集分宝时间
    OrderTime time.Time `json:"order_time";xml:"order_time"`
    
    // PointCount required发放集分宝的数量
    PointCount int64 `json:"point_count";xml:"point_count"`
    
    // UserSymbol required用户标识符，用于指定集分宝发放的用户，和user_symbol_type一起使用，确定一个唯一的支付宝用户
    UserSymbol string `json:"user_symbol";xml:"user_symbol"`
    
    // UserSymbolType required用户标识符类型，现在支持ALIPAY_USER_ID:表示支付宝用户ID,ALIPAY_LOGON_ID:表示支付宝登陆号
    UserSymbolType string `json:"user_symbol_type";xml:"user_symbol_type"`
    
}

func (req *AlipayPointOrderAddRequest) GetAPIName() string {
	return "alipay.point.order.add"
}

func (req *AlipayPointOrderAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["AuthToken"] = req.AuthToken
    
    params["Memo"] = req.Memo
    
    params["MerchantOrderNo"] = req.MerchantOrderNo
    
    params["OrderTime"] = req.OrderTime
    
    params["PointCount"] = req.PointCount
    
    params["UserSymbol"] = req.UserSymbol
    
    params["UserSymbolType"] = req.UserSymbolType
    
    return params
}

// AlipayPointOrderAddResponse 向用户发送集分宝，发放集分宝之前，操作流程如下：
//1、申请支付宝增值包。
//2、申请支付宝应用上线时选择集分宝API。
//3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
//4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
//5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
//6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
//7、调用发放API（alipay.point.order.add）发放集分宝。
type AlipayPointOrderAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AlipayOrderNo Basic支付宝集分宝发放流水号
    AlipayOrderNo string `json:"alipay_order_no";xml:"alipay_order_no"`
    
    // ResultCode Basic充值结果：SUCCESS表示成功，其他表示失败
    ResultCode bool `json:"result_code";xml:"result_code"`
    
}

// TaobaoFenxiaoProductUpdateRequest 更新分销平台产品数据，不传更新数据返回失败<br>
//1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
type TaobaoFenxiaoProductUpdateRequest struct {
    
    // CategoryId optional所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // City optional所在地：市，例：“杭州”
    City string `json:"city";xml:"city"`
    
    // CostPrice optional代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    CostPrice string `json:"cost_price";xml:"cost_price"`
    
    // DealerCostPrice optional经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    DealerCostPrice string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    // Desc optional产品描述，长度为5到25000字符。
    Desc string `json:"desc";xml:"desc"`
    
    // DiscountId optional折扣ID
    DiscountId int64 `json:"discount_id";xml:"discount_id"`
    
    // HaveInvoice optional是否有发票，可选值：false（否）、true（是），默认false。
    HaveInvoice string `json:"have_invoice";xml:"have_invoice"`
    
    // HaveQuarantee optional是否有保修，可选值：false（否）、true（是），默认false。
    HaveQuarantee string `json:"have_quarantee";xml:"have_quarantee"`
    
    // Image optional主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数
    Image []byte `json:"image";xml:"image"`
    
    // InputProperties optional自定义属性。格式为pid:value;pid:value
    InputProperties string `json:"input_properties";xml:"input_properties"`
    
    // IsAuthz optional产品是否需要授权isAuthz:yes|no 
    // yes:需要授权 
    // no:不需要授权
    IsAuthz string `json:"is_authz";xml:"is_authz"`
    
    // Name optional产品名称，长度不超过60个字节。
    Name string `json:"name";xml:"name"`
    
    // OuterId optional商家编码，长度不能超过60个字节。
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // PicPath optional产品主图图片空间相对路径或绝对路径
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // Pid required产品ID
    Pid int64 `json:"pid";xml:"pid"`
    
    // PostageEms optionalems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
    PostageEms string `json:"postage_ems";xml:"postage_ems"`
    
    // PostageFast optional快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
    PostageFast string `json:"postage_fast";xml:"postage_fast"`
    
    // PostageId optional运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。
    PostageId int64 `json:"postage_id";xml:"postage_id"`
    
    // PostageOrdinary optional平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
    PostageOrdinary string `json:"postage_ordinary";xml:"postage_ordinary"`
    
    // PostageType optional运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。
    PostageType string `json:"postage_type";xml:"postage_type"`
    
    // Properties optional产品属性
    Properties string `json:"properties";xml:"properties"`
    
    // PropertyAlias optional属性别名
    PropertyAlias string `json:"property_alias";xml:"property_alias"`
    
    // Prov optional所在地：省，例：“浙江”
    Prov string `json:"prov";xml:"prov"`
    
    // Quantity optional产品库存必须是1到999999。
    Quantity int64 `json:"quantity";xml:"quantity"`
    
    // RetailPriceHigh optional最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
    RetailPriceHigh string `json:"retail_price_high";xml:"retail_price_high"`
    
    // RetailPriceLow optional最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    RetailPriceLow string `json:"retail_price_low";xml:"retail_price_low"`
    
    // SkuCostPrices optionalsku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
    SkuCostPrices string `json:"sku_cost_prices";xml:"sku_cost_prices"`
    
    // SkuDealerCostPrices optionalsku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
    SkuDealerCostPrices string `json:"sku_dealer_cost_prices";xml:"sku_dealer_cost_prices"`
    
    // SkuIds optionalsku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。
    SkuIds string `json:"sku_ids";xml:"sku_ids"`
    
    // SkuOuterIds optionalsku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",,"
    SkuOuterIds string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    // SkuProperties optionalsku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)
    // 通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。
    SkuProperties string `json:"sku_properties";xml:"sku_properties"`
    
    // SkuPropertiesDel optional根据sku属性删除sku信息。需要按组删除属性。
    SkuPropertiesDel string `json:"sku_properties_del";xml:"sku_properties_del"`
    
    // SkuQuantitys optionalsku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。
    SkuQuantitys string `json:"sku_quantitys";xml:"sku_quantitys"`
    
    // SkuStandardPrices optionalsku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
    SkuStandardPrices string `json:"sku_standard_prices";xml:"sku_standard_prices"`
    
    // StandardPrice optional采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    StandardPrice string `json:"standard_price";xml:"standard_price"`
    
    // StandardRetailPrice optional零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    StandardRetailPrice string `json:"standard_retail_price";xml:"standard_retail_price"`
    
    // Status optional发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。
    Status string `json:"status";xml:"status"`
    
}

func (req *TaobaoFenxiaoProductUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.update"
}

func (req *TaobaoFenxiaoProductUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 37)
    
    params["CategoryId"] = req.CategoryId
    
    params["City"] = req.City
    
    params["CostPrice"] = req.CostPrice
    
    params["DealerCostPrice"] = req.DealerCostPrice
    
    params["Desc"] = req.Desc
    
    params["DiscountId"] = req.DiscountId
    
    params["HaveInvoice"] = req.HaveInvoice
    
    params["HaveQuarantee"] = req.HaveQuarantee
    
    params["Image"] = req.Image
    
    params["InputProperties"] = req.InputProperties
    
    params["IsAuthz"] = req.IsAuthz
    
    params["Name"] = req.Name
    
    params["OuterId"] = req.OuterId
    
    params["PicPath"] = req.PicPath
    
    params["Pid"] = req.Pid
    
    params["PostageEms"] = req.PostageEms
    
    params["PostageFast"] = req.PostageFast
    
    params["PostageId"] = req.PostageId
    
    params["PostageOrdinary"] = req.PostageOrdinary
    
    params["PostageType"] = req.PostageType
    
    params["Properties"] = req.Properties
    
    params["PropertyAlias"] = req.PropertyAlias
    
    params["Prov"] = req.Prov
    
    params["Quantity"] = req.Quantity
    
    params["RetailPriceHigh"] = req.RetailPriceHigh
    
    params["RetailPriceLow"] = req.RetailPriceLow
    
    params["SkuCostPrices"] = req.SkuCostPrices
    
    params["SkuDealerCostPrices"] = req.SkuDealerCostPrices
    
    params["SkuIds"] = req.SkuIds
    
    params["SkuOuterIds"] = req.SkuOuterIds
    
    params["SkuProperties"] = req.SkuProperties
    
    params["SkuPropertiesDel"] = req.SkuPropertiesDel
    
    params["SkuQuantitys"] = req.SkuQuantitys
    
    params["SkuStandardPrices"] = req.SkuStandardPrices
    
    params["StandardPrice"] = req.StandardPrice
    
    params["StandardRetailPrice"] = req.StandardRetailPrice
    
    params["Status"] = req.Status
    
    return params
}

// TaobaoFenxiaoProductUpdateResponse 更新分销平台产品数据，不传更新数据返回失败<br>
//1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
type TaobaoFenxiaoProductUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Modified Basic更新时间，时间格式：yyyy-MM-dd HH:mm:ss
    Modified time.Time `json:"modified";xml:"modified"`
    
    // Pid Basic产品ID
    Pid int64 `json:"pid";xml:"pid"`
    
}

// TaobaoSubuserFullinfoGetRequest 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
type TaobaoSubuserFullinfoGetRequest struct {
    
    // Fields optional传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email）
    Fields string `json:"fields";xml:"fields"`
    
    // SubId special子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
    SubId int64 `json:"sub_id";xml:"sub_id"`
    
    // SubNick special子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）
    SubNick string `json:"sub_nick";xml:"sub_nick"`
    
}

func (req *TaobaoSubuserFullinfoGetRequest) GetAPIName() string {
	return "taobao.subuser.fullinfo.get"
}

func (req *TaobaoSubuserFullinfoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Fields"] = req.Fields
    
    params["SubId"] = req.SubId
    
    params["SubNick"] = req.SubNick
    
    return params
}

// TaobaoSubuserFullinfoGetResponse 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
type TaobaoSubuserFullinfoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SubFullinfo Object子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息
    SubFullinfo SubUserFullInfo `json:"sub_fullinfo";xml:"sub_fullinfo"`
    
}

// TaobaoFenxiaoProductcatsGetRequest 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaoFenxiaoProductcatsGetRequest struct {
    
    // Fields optional返回字段列表
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoFenxiaoProductcatsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.productcats.get"
}

func (req *TaobaoFenxiaoProductcatsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoFenxiaoProductcatsGetResponse 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaoFenxiaoProductcatsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Productcats Object Array产品线列表。返回 ProductCat 包含的字段信息。
    Productcats ProductCat `json:"productcats";xml:"productcats"`
    
    // TotalResults Basic查询结果记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoSubusersGetRequest 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaoSubusersGetRequest struct {
    
    // UserNick required主账号用户名
    UserNick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubusersGetRequest) GetAPIName() string {
	return "taobao.subusers.get"
}

func (req *TaobaoSubusersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["UserNick"] = req.UserNick
    
    return params
}

// TaobaoSubusersGetResponse 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaoSubusersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Subaccounts Object Array子账号基本信息
    Subaccounts SubAccountInfo `json:"subaccounts";xml:"subaccounts"`
    
}

// TaobaoFenxiaoProductAddRequest 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。
//
//    * 产品图片默认为空
//    * 产品发布后默认为下架状态
type TaobaoFenxiaoProductAddRequest struct {
    
    // CategoryId required所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // City required所在地：市，例：“杭州”
    City string `json:"city";xml:"city"`
    
    // CostPrice optional代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    CostPrice string `json:"cost_price";xml:"cost_price"`
    
    // DealerCostPrice optional经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    DealerCostPrice string `json:"dealer_cost_price";xml:"dealer_cost_price"`
    
    // Desc required产品描述，长度为5到25000字符。
    Desc string `json:"desc";xml:"desc"`
    
    // DiscountId optional折扣ID
    DiscountId int64 `json:"discount_id";xml:"discount_id"`
    
    // HaveInvoice optional是否有发票，可选值：false（否）、true（是），默认false。
    HaveInvoice string `json:"have_invoice";xml:"have_invoice"`
    
    // HaveQuarantee optional是否有保修，可选值：false（否）、true（是），默认false。
    HaveQuarantee string `json:"have_quarantee";xml:"have_quarantee"`
    
    // Image optional产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
    Image []byte `json:"image";xml:"image"`
    
    // InputProperties optional自定义属性。格式为pid:value;pid:value
    InputProperties string `json:"input_properties";xml:"input_properties"`
    
    // IsAuthz optional添加产品时，添加入参isAuthz:yes|no 
    // yes:需要授权 
    // no:不需要授权 
    // 默认是需要授权
    IsAuthz string `json:"is_authz";xml:"is_authz"`
    
    // ItemId optional导入的商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // Name required产品名称，长度不超过60个字节。
    Name string `json:"name";xml:"name"`
    
    // OuterId optional商家编码，长度不能超过60个字节。
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // PicPath optional产品主图图片空间相对路径或绝对路径
    PicPath string `json:"pic_path";xml:"pic_path"`
    
    // PostageEms optionalems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
    PostageEms string `json:"postage_ems";xml:"postage_ems"`
    
    // PostageFast optional快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
    PostageFast string `json:"postage_fast";xml:"postage_fast"`
    
    // PostageId optional运费模板ID，参考taobao.postages.get。
    PostageId int64 `json:"postage_id";xml:"postage_id"`
    
    // PostageOrdinary optional平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
    PostageOrdinary string `json:"postage_ordinary";xml:"postage_ordinary"`
    
    // PostageType optional运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
    PostageType string `json:"postage_type";xml:"postage_type"`
    
    // ProductcatId required产品线ID
    ProductcatId int64 `json:"productcat_id";xml:"productcat_id"`
    
    // Properties optional产品属性，格式为pid:vid;pid:vid
    Properties string `json:"properties";xml:"properties"`
    
    // PropertyAlias optional属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
    PropertyAlias string `json:"property_alias";xml:"property_alias"`
    
    // Prov required所在地：省，例：“浙江”
    Prov string `json:"prov";xml:"prov"`
    
    // Quantity required产品库存必须是1到999999。
    Quantity int64 `json:"quantity";xml:"quantity"`
    
    // RetailPriceHigh required最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
    RetailPriceHigh string `json:"retail_price_high";xml:"retail_price_high"`
    
    // RetailPriceLow required最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    RetailPriceLow string `json:"retail_price_low";xml:"retail_price_low"`
    
    // SkuCostPrices optionalsku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    SkuCostPrices string `json:"sku_cost_prices";xml:"sku_cost_prices"`
    
    // SkuDealerCostPrices optionalsku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
    SkuDealerCostPrices string `json:"sku_dealer_cost_prices";xml:"sku_dealer_cost_prices"`
    
    // SkuOuterIds optionalsku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    SkuOuterIds string `json:"sku_outer_ids";xml:"sku_outer_ids"`
    
    // SkuProperties optionalsku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    SkuProperties string `json:"sku_properties";xml:"sku_properties"`
    
    // SkuQuantitys optionalsku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    SkuQuantitys string `json:"sku_quantitys";xml:"sku_quantitys"`
    
    // SkuStandardPrices optionalsku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    SkuStandardPrices string `json:"sku_standard_prices";xml:"sku_standard_prices"`
    
    // SpuId optional产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
    SpuId int64 `json:"spu_id";xml:"spu_id"`
    
    // StandardPrice required采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    StandardPrice string `json:"standard_price";xml:"standard_price"`
    
    // StandardRetailPrice optional零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    StandardRetailPrice string `json:"standard_retail_price";xml:"standard_retail_price"`
    
    // TradeType optional分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
    TradeType string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.add"
}

func (req *TaobaoFenxiaoProductAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 37)
    
    params["CategoryId"] = req.CategoryId
    
    params["City"] = req.City
    
    params["CostPrice"] = req.CostPrice
    
    params["DealerCostPrice"] = req.DealerCostPrice
    
    params["Desc"] = req.Desc
    
    params["DiscountId"] = req.DiscountId
    
    params["HaveInvoice"] = req.HaveInvoice
    
    params["HaveQuarantee"] = req.HaveQuarantee
    
    params["Image"] = req.Image
    
    params["InputProperties"] = req.InputProperties
    
    params["IsAuthz"] = req.IsAuthz
    
    params["ItemId"] = req.ItemId
    
    params["Name"] = req.Name
    
    params["OuterId"] = req.OuterId
    
    params["PicPath"] = req.PicPath
    
    params["PostageEms"] = req.PostageEms
    
    params["PostageFast"] = req.PostageFast
    
    params["PostageId"] = req.PostageId
    
    params["PostageOrdinary"] = req.PostageOrdinary
    
    params["PostageType"] = req.PostageType
    
    params["ProductcatId"] = req.ProductcatId
    
    params["Properties"] = req.Properties
    
    params["PropertyAlias"] = req.PropertyAlias
    
    params["Prov"] = req.Prov
    
    params["Quantity"] = req.Quantity
    
    params["RetailPriceHigh"] = req.RetailPriceHigh
    
    params["RetailPriceLow"] = req.RetailPriceLow
    
    params["SkuCostPrices"] = req.SkuCostPrices
    
    params["SkuDealerCostPrices"] = req.SkuDealerCostPrices
    
    params["SkuOuterIds"] = req.SkuOuterIds
    
    params["SkuProperties"] = req.SkuProperties
    
    params["SkuQuantitys"] = req.SkuQuantitys
    
    params["SkuStandardPrices"] = req.SkuStandardPrices
    
    params["SpuId"] = req.SpuId
    
    params["StandardPrice"] = req.StandardPrice
    
    params["StandardRetailPrice"] = req.StandardRetailPrice
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoProductAddResponse 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。
//
//    * 产品图片默认为空
//    * 产品发布后默认为下架状态
type TaobaoFenxiaoProductAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Created Basic产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss
    Created time.Time `json:"created";xml:"created"`
    
    // Pid Basic产品ID
    Pid int64 `json:"pid";xml:"pid"`
    
}

// TaobaoTmcMsgSendrecordRequest 查询单条消息发送记录，只返回返回条数和时间。
type TaobaoTmcMsgSendrecordRequest struct {
    
    // DataId required消息主键ID
    DataId string `json:"data_id";xml:"data_id"`
    
    // GroupName required消息分组名
    GroupName string `json:"group_name";xml:"group_name"`
    
    // TopicName requiredTOPIC名称
    TopicName string `json:"topic_name";xml:"topic_name"`
    
}

func (req *TaobaoTmcMsgSendrecordRequest) GetAPIName() string {
	return "taobao.tmc.msg.sendrecord"
}

func (req *TaobaoTmcMsgSendrecordRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["DataId"] = req.DataId
    
    params["GroupName"] = req.GroupName
    
    params["TopicName"] = req.TopicName
    
    return params
}

// TaobaoTmcMsgSendrecordResponse 查询单条消息发送记录，只返回返回条数和时间。
type TaobaoTmcMsgSendrecordResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TbSendList Basic淘宝发送时间
    TbSendList string `json:"tb_send_list";xml:"tb_send_list"`
    
    // TbSendTimes Basic淘宝发送测试
    TbSendTimes int64 `json:"tb_send_times";xml:"tb_send_times"`
    
    // TmcSendList BasicTMC发送时间
    TmcSendList string `json:"tmc_send_list";xml:"tmc_send_list"`
    
    // TmcSendTimes Basictmc发送次数
    TmcSendTimes int64 `json:"tmc_send_times";xml:"tmc_send_times"`
    
}

// TaobaoWangwangEserviceChatpeersGetRequest 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。
type TaobaoWangwangEserviceChatpeersGetRequest struct {
    
}

func (req *TaobaoWangwangEserviceChatpeersGetRequest) GetAPIName() string {
	return "taobao.wangwang.eservice.chatpeers.get"
}

func (req *TaobaoWangwangEserviceChatpeersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoWangwangEserviceChatpeersGetResponse 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。
type TaobaoWangwangEserviceChatpeersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Chatpeers Object Array聊天对象ID列表
    Chatpeers Chatpeer `json:"chatpeers";xml:"chatpeers"`
    
    // Count Basic成员数目。
    Count int64 `json:"count";xml:"count"`
    
    // Ret Basic返回码： 
    // 10000:成功； 
    // 
    // 60000：时间非法，包括1)时间段超过7天,或2)起始时间距今超过30天，或3)时间格式不对； 
    // 
    // 50000：聊天用户ID不是该店铺的帐号； 
    // 
    // 40000：系统错误，包括必填参数为空。
    Ret int64 `json:"ret";xml:"ret"`
    
}

// TaobaoFenxiaoProductsGetRequest 查询供应商的产品数据。
//
//    * 入参传入pids将优先查询，即只按这个条件查询。
//    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
//    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
//    * 入参fields传入images将查询多图数据，不传只返回主图数据。
//    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
//    * 查询结果按照产品发布时间倒序，即时间近的数据在前。
type TaobaoFenxiaoProductsGetRequest struct {
    
    // EndModified optional结束修改时间
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // Fields optional指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
    Fields string `json:"fields";xml:"fields"`
    
    // IsAuthz optional查询产品列表时，查询入参“是否需要授权”
    // yes:需要授权 
    // no:不需要授权
    IsAuthz string `json:"is_authz";xml:"is_authz"`
    
    // ItemIds optional可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
    ItemIds string `json:"item_ids";xml:"item_ids"`
    
    // OuterId optional商家编码
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // PageNo optional页码（大于0的整数，默认1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页记录数（默认20，最大50）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // Pids optional产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
    Pids string `json:"pids";xml:"pids"`
    
    // ProductcatId optional产品线ID
    ProductcatId int64 `json:"productcat_id";xml:"productcat_id"`
    
    // SkuNumber optionalsku商家编码
    SkuNumber string `json:"sku_number";xml:"sku_number"`
    
    // StartModified optional开始修改时间
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoFenxiaoProductsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.products.get"
}

func (req *TaobaoFenxiaoProductsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 11)
    
    params["EndModified"] = req.EndModified
    
    params["Fields"] = req.Fields
    
    params["IsAuthz"] = req.IsAuthz
    
    params["ItemIds"] = req.ItemIds
    
    params["OuterId"] = req.OuterId
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["Pids"] = req.Pids
    
    params["ProductcatId"] = req.ProductcatId
    
    params["SkuNumber"] = req.SkuNumber
    
    params["StartModified"] = req.StartModified
    
    return params
}

// TaobaoFenxiaoProductsGetResponse 查询供应商的产品数据。
//
//    * 入参传入pids将优先查询，即只按这个条件查询。
//    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
//    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
//    * 入参fields传入images将查询多图数据，不传只返回主图数据。
//    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
//    * 查询结果按照产品发布时间倒序，即时间近的数据在前。
type TaobaoFenxiaoProductsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Products Object Array产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
    Products FenxiaoProduct `json:"products";xml:"products"`
    
    // TotalResults Basic查询结果记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoSubuserDepartmentsGetRequest 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
type TaobaoSubuserDepartmentsGetRequest struct {
    
    // UserNick required主账号用户名
    UserNick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoSubuserDepartmentsGetRequest) GetAPIName() string {
	return "taobao.subuser.departments.get"
}

func (req *TaobaoSubuserDepartmentsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["UserNick"] = req.UserNick
    
    return params
}

// TaobaoSubuserDepartmentsGetResponse 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
type TaobaoSubuserDepartmentsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Departments Object Array部门信息
    Departments Department `json:"departments";xml:"departments"`
    
}

// AlipayUserAccountreportGetRequest 获取支付宝余额明细记录，不包含用户通过银行卡快捷支付进行交易的交易明细
type AlipayUserAccountreportGetRequest struct {
    
    // EndTime required对账单结束时间，其中end_time - start_time <= 1天，对于对账记录比较多的情况，强烈建议按天查询，否则会出现超时的情况。
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // Fields required需要返回的字段列表。create_time:创建时间,type：账务类型,business_type:子业务类型,balance:当时支付宝账户余额,in_amount:收入金额,out_amount:支出金额,alipay_order_no:支付宝订单号,merchant_order_no:商户订单号,self_user_id:自己的支付宝ID,opt_user_id:对方的支付宝ID,memo:账号备注
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional要获取的对账单页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每次查询获取对账记录数量
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime required对账单开始时间。最近一个月内的日期。
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
    // Type optional账务类型。多个类型是，用逗号分隔，不传则查询所有类型的。PAYMENT:在线支付，TRANSFER:转账，DEPOSIT:充值，WITHDRAW:提现，CHARGE:收费，PREAUTH:预授权，OTHER：其它。
    Type string `json:"type";xml:"type"`
    
}

func (req *AlipayUserAccountreportGetRequest) GetAPIName() string {
	return "alipay.user.accountreport.get"
}

func (req *AlipayUserAccountreportGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["EndTime"] = req.EndTime
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    params["Type"] = req.Type
    
    return params
}

// AlipayUserAccountreportGetResponse 获取支付宝余额明细记录，不包含用户通过银行卡快捷支付进行交易的交易明细
type AlipayUserAccountreportGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AlipayRecords Object Array对账记录列表
    AlipayRecords AlipayRecord `json:"alipay_records";xml:"alipay_records"`
    
    // TotalPages Basic总页数
    TotalPages int64 `json:"total_pages";xml:"total_pages"`
    
    // TotalResults Basic总记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoFuwuSaleLinkGenRequest 服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br>
//注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。
type TaobaoFuwuSaleLinkGenRequest struct {
    
    // Nick optional用户需要营销的目标人群中的用户nick
    Nick string `json:"nick";xml:"nick"`
    
    // ParamStr required从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。
    ParamStr string `json:"param_str";xml:"param_str"`
    
}

func (req *TaobaoFuwuSaleLinkGenRequest) GetAPIName() string {
	return "taobao.fuwu.sale.link.gen"
}

func (req *TaobaoFuwuSaleLinkGenRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Nick"] = req.Nick
    
    params["ParamStr"] = req.ParamStr
    
    return params
}

// TaobaoFuwuSaleLinkGenResponse 服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br>
//注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。
type TaobaoFuwuSaleLinkGenResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Url Basic通过营销链接接口生成的营销链接短地址
    Url string `json:"url";xml:"url"`
    
}

// TmallProductSchemaGetRequest 产品信息获取接口schema形式返回
type TmallProductSchemaGetRequest struct {
    
    // ProductId required产品编号
    ProductId int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallProductSchemaGetRequest) GetAPIName() string {
	return "tmall.product.schema.get"
}

func (req *TmallProductSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ProductId"] = req.ProductId
    
    return params
}

// TmallProductSchemaGetResponse 产品信息获取接口schema形式返回
type TmallProductSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GetProductResult Basic产品信息数据。schema形式
    GetProductResult string `json:"get_product_result";xml:"get_product_result"`
    
}

// CainiaoCloudprintIsvResourcesGetRequest isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaoCloudprintIsvResourcesGetRequest struct {
    
    // IsvResourceType requiredisv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
    IsvResourceType string `json:"isv_resource_type";xml:"isv_resource_type"`
    
}

func (req *CainiaoCloudprintIsvResourcesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.isv.resources.get"
}

func (req *CainiaoCloudprintIsvResourcesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["IsvResourceType"] = req.IsvResourceType
    
    return params
}

// CainiaoCloudprintIsvResourcesGetResponse isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaoCloudprintIsvResourcesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// CainiaoOpenstorageIsvResourceCreateRequest isv资源创建接口（云打印开源存储）
type CainiaoOpenstorageIsvResourceCreateRequest struct {
    
    // ParamCreateIsvResourceRequest optionalisv创建资源的参数
    ParamCreateIsvResourceRequest CreateIsvResourceRequest `json:"param_create_isv_resource_request";xml:"param_create_isv_resource_request"`
    
}

func (req *CainiaoOpenstorageIsvResourceCreateRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resource.create"
}

func (req *CainiaoOpenstorageIsvResourceCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamCreateIsvResourceRequest"] = req.ParamCreateIsvResourceRequest
    
    return params
}

// CainiaoOpenstorageIsvResourceCreateResponse isv资源创建接口（云打印开源存储）
type CainiaoOpenstorageIsvResourceCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode Basic错误码
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // ErrMessage Basic错误消息
    ErrMessage string `json:"err_message";xml:"err_message"`
    
}

// TaobaoTopSecretGetRequest top sdk通过api获取对应解密秘钥
type TaobaoTopSecretGetRequest struct {
    
    // CustomerUserId optional自定义用户id
    CustomerUserId int64 `json:"customer_user_id";xml:"customer_user_id"`
    
    // RandomNum required伪随机数
    RandomNum string `json:"random_num";xml:"random_num"`
    
    // SecretVersion optional秘钥版本号
    SecretVersion int64 `json:"secret_version";xml:"secret_version"`
    
}

func (req *TaobaoTopSecretGetRequest) GetAPIName() string {
	return "taobao.top.secret.get"
}

func (req *TaobaoTopSecretGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CustomerUserId"] = req.CustomerUserId
    
    params["RandomNum"] = req.RandomNum
    
    params["SecretVersion"] = req.SecretVersion
    
    return params
}

// TaobaoTopSecretGetResponse top sdk通过api获取对应解密秘钥
type TaobaoTopSecretGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AppConfig Basicapp配置信息
    AppConfig string `json:"app_config";xml:"app_config"`
    
    // Interval Basic下次更新秘钥间隔，单位（秒）
    Interval int64 `json:"interval";xml:"interval"`
    
    // MaxInterval Basic最长有效期，容灾使用，单位（秒）
    MaxInterval int64 `json:"max_interval";xml:"max_interval"`
    
    // Secret Basic秘钥值
    Secret string `json:"secret";xml:"secret"`
    
    // SecretVersion Basic秘钥版本号
    SecretVersion int64 `json:"secret_version";xml:"secret_version"`
    
}

// TaobaoLogisticsAddressReachablebatchGetRequest 批量判定服务是否可达
type TaobaoLogisticsAddressReachablebatchGetRequest struct {
    
    // AddressList optional筛单用户输入地址结构
    AddressList AddressReachable `json:"address_list";xml:"address_list"`
    
}

func (req *TaobaoLogisticsAddressReachablebatchGetRequest) GetAPIName() string {
	return "taobao.logistics.address.reachablebatch.get"
}

func (req *TaobaoLogisticsAddressReachablebatchGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["AddressList"] = req.AddressList
    
    return params
}

// TaobaoLogisticsAddressReachablebatchGetResponse 批量判定服务是否可达
type TaobaoLogisticsAddressReachablebatchGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ReachableResults Object Array物流是否可达结果列表
    ReachableResults AddressReachableTopResult `json:"reachable_results";xml:"reachable_results"`
    
}

// TaobaoOmniorderStoreReallocateRequest 门店发货提供改派接口
type TaobaoOmniorderStoreReallocateRequest struct {
    
    // MainOrderId required主订单号
    MainOrderId int64 `json:"main_order_id";xml:"main_order_id"`
    
    // StoreId optional门店Id
    StoreId int64 `json:"store_id";xml:"store_id"`
    
    // SubOrderIds required子订单号
    SubOrderIds int64 `json:"sub_order_ids";xml:"sub_order_ids"`
    
    // WarehouseCode optional电商仓code
    WarehouseCode string `json:"warehouse_code";xml:"warehouse_code"`
    
}

func (req *TaobaoOmniorderStoreReallocateRequest) GetAPIName() string {
	return "taobao.omniorder.store.reallocate"
}

func (req *TaobaoOmniorderStoreReallocateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["MainOrderId"] = req.MainOrderId
    
    params["StoreId"] = req.StoreId
    
    params["SubOrderIds"] = req.SubOrderIds
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoOmniorderStoreReallocateResponse 门店发货提供改派接口
type TaobaoOmniorderStoreReallocateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoQimenSingleitemQueryRequest 商品查询接口
type TaobaoQimenSingleitemQueryRequest struct {
    
    // ItemCode optional商品编码,S1234,string(50),必填,
    ItemCode string `json:"itemCode";xml:"itemCode"`
    
    // ItemId optional仓储系统商品编码,C123,string(50),必填,
    ItemId string `json:"itemId";xml:"itemId"`
    
    // OwnerCode optional货主编码,H123,string(50),必填,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenSingleitemQueryRequest) GetAPIName() string {
	return "taobao.qimen.singleitem.query"
}

func (req *TaobaoQimenSingleitemQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ItemCode"] = req.ItemCode
    
    params["ItemId"] = req.ItemId
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenSingleitemQueryResponse 商品查询接口
type TaobaoQimenSingleitemQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Item Objectitem
    Item Item `json:"item";xml:"item"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenItemmappingCreateRequest 前后端商品映射
type TaobaoQimenItemmappingCreateRequest struct {
    
    // ActionType optional奇门仓储字段,C123,string(50),必填,
    ActionType string `json:"actionType";xml:"actionType"`
    
    // ItemId optional奇门仓储字段,C123,string(50),必填,
    ItemId string `json:"itemId";xml:"itemId"`
    
    // ItemSource optional奇门仓储字段,C123,string(50),必填,
    ItemSource string `json:"itemSource";xml:"itemSource"`
    
    // OwnerCode optional奇门仓储字段,C123,string(50),必填,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // ShopItemId optional奇门仓储字段,C123,string(50),必填,
    ShopItemId string `json:"shopItemId";xml:"shopItemId"`
    
    // ShopNick optional奇门仓储字段,C123,string(50),必填,
    ShopNick string `json:"shopNick";xml:"shopNick"`
    
    // SkuId optional奇门仓储字段,C123,string(50),必填,
    SkuId string `json:"skuId";xml:"skuId"`
    
}

func (req *TaobaoQimenItemmappingCreateRequest) GetAPIName() string {
	return "taobao.qimen.itemmapping.create"
}

func (req *TaobaoQimenItemmappingCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["ActionType"] = req.ActionType
    
    params["ItemId"] = req.ItemId
    
    params["ItemSource"] = req.ItemSource
    
    params["OwnerCode"] = req.OwnerCode
    
    params["ShopItemId"] = req.ShopItemId
    
    params["ShopNick"] = req.ShopNick
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoQimenItemmappingCreateResponse 前后端商品映射
type TaobaoQimenItemmappingCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenInventoryruleCreateRequest 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateRequest struct {
    
    // InventoryRules optionalinventoryRules
    InventoryRules InventoryRule `json:"inventoryRules";xml:"inventoryRules"`
    
    // OwnerCode optional奇门仓储字段,C1223,string(50),,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenInventoryruleCreateRequest) GetAPIName() string {
	return "taobao.qimen.inventoryrule.create"
}

func (req *TaobaoQimenInventoryruleCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["InventoryRules"] = req.InventoryRules
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenInventoryruleCreateResponse 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenCombineitemDeleteRequest 组合货品删除
type TaobaoQimenCombineitemDeleteRequest struct {
    
    // ItemId optional奇门仓储字段,C123,string(50),,
    ItemId string `json:"itemId";xml:"itemId"`
    
    // OwnerCode optional奇门仓储字段,C123,string(50),,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenCombineitemDeleteRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.delete"
}

func (req *TaobaoQimenCombineitemDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenCombineitemDeleteResponse 组合货品删除
type TaobaoQimenCombineitemDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TmallItemCalculateHscodeGetRequest 算法获取hscode
type TmallItemCalculateHscodeGetRequest struct {
    
    // ItemId optional商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TmallItemCalculateHscodeGetRequest) GetAPIName() string {
	return "tmall.item.calculate.hscode.get"
}

func (req *TmallItemCalculateHscodeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TmallItemCalculateHscodeGetResponse 算法获取hscode
type TmallItemCalculateHscodeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Results Basic Array算法返回预测的hscode数据
    Results map[string]interface{} `json:"results";xml:"results"`
    
}

// TaobaoQimenChannelinventoryQueryRequest 渠道库存查询
type TaobaoQimenChannelinventoryQueryRequest struct {
    
    // ChannelCodes optional奇门仓储字段
    ChannelCodes string `json:"channelCodes";xml:"channelCodes"`
    
    // ItemCodes optional奇门仓储字段
    ItemCodes string `json:"itemCodes";xml:"itemCodes"`
    
    // OwnerCode optional奇门仓储字段,C123,string(50),,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // WarehouseCodes optional奇门仓储字段
    WarehouseCodes string `json:"warehouseCodes";xml:"warehouseCodes"`
    
}

func (req *TaobaoQimenChannelinventoryQueryRequest) GetAPIName() string {
	return "taobao.qimen.channelinventory.query"
}

func (req *TaobaoQimenChannelinventoryQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ChannelCodes"] = req.ChannelCodes
    
    params["ItemCodes"] = req.ItemCodes
    
    params["OwnerCode"] = req.OwnerCode
    
    params["WarehouseCodes"] = req.WarehouseCodes
    
    return params
}

// TaobaoQimenChannelinventoryQueryResponse 渠道库存查询
type TaobaoQimenChannelinventoryQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // ItemInventories Object ArrayitemInventories
    ItemInventories ItemInventory `json:"itemInventories";xml:"itemInventories"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TaobaoFenxiaoOrderRemarkUpdateRequest 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateRequest struct {
    
    // PurchaseOrderId required采购单编号
    PurchaseOrderId int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
    // SupplierMemo required备注内容(供应商操作)
    SupplierMemo string `json:"supplier_memo";xml:"supplier_memo"`
    
    // SupplierMemoFlag optional旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。
    // 1:红色
    // 2:黄色
    // 3:绿色
    // 4:蓝色
    // 5:粉红色
    SupplierMemoFlag int64 `json:"supplier_memo_flag";xml:"supplier_memo_flag"`
    
}

func (req *TaobaoFenxiaoOrderRemarkUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.order.remark.update"
}

func (req *TaobaoFenxiaoOrderRemarkUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["PurchaseOrderId"] = req.PurchaseOrderId
    
    params["SupplierMemo"] = req.SupplierMemo
    
    params["SupplierMemoFlag"] = req.SupplierMemoFlag
    
    return params
}

// TaobaoFenxiaoOrderRemarkUpdateResponse 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoInventoryAdjustExternalRequest 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoInventoryAdjustExternalRequest struct {
    
    // BizType optional外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
    BizType string `json:"biz_type";xml:"biz_type"`
    
    // BizUniqueCode required商家外部定单号
    BizUniqueCode string `json:"biz_unique_code";xml:"biz_unique_code"`
    
    // Items required商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]
    Items string `json:"items";xml:"items"`
    
    // OccupyOperateCode optional库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
    OccupyOperateCode string `json:"occupy_operate_code";xml:"occupy_operate_code"`
    
    // OperateTime optional业务操作时间
    OperateTime string `json:"operate_time";xml:"operate_time"`
    
    // OperateType optionaltest
    OperateType string `json:"operate_type";xml:"operate_type"`
    
    // ReduceType optionaltest
    ReduceType string `json:"reduce_type";xml:"reduce_type"`
    
    // StoreCode required商家仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryAdjustExternalRequest) GetAPIName() string {
	return "taobao.inventory.adjust.external"
}

func (req *TaobaoInventoryAdjustExternalRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["BizType"] = req.BizType
    
    params["BizUniqueCode"] = req.BizUniqueCode
    
    params["Items"] = req.Items
    
    params["OccupyOperateCode"] = req.OccupyOperateCode
    
    params["OperateTime"] = req.OperateTime
    
    params["OperateType"] = req.OperateType
    
    params["ReduceType"] = req.ReduceType
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoInventoryAdjustExternalResponse 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoInventoryAdjustExternalResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OperateCode Basic操作返回码
    OperateCode string `json:"operate_code";xml:"operate_code"`
    
    // TipInfos Object Array提示信息
    TipInfos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

// TaobaoInventoryQueryRequest 商家查询商品总体库存信息
type TaobaoInventoryQueryRequest struct {
    
    // ScItemCodes optional后端商品的商家编码列表，控制到50个
    ScItemCodes string `json:"sc_item_codes";xml:"sc_item_codes"`
    
    // ScItemIds required后端商品ID 列表，控制到50个
    ScItemIds string `json:"sc_item_ids";xml:"sc_item_ids"`
    
    // SellerNick optional卖家昵称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
    // StoreCodes optional仓库列表:GLY001^GLY002
    StoreCodes string `json:"store_codes";xml:"store_codes"`
    
}

func (req *TaobaoInventoryQueryRequest) GetAPIName() string {
	return "taobao.inventory.query"
}

func (req *TaobaoInventoryQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ScItemCodes"] = req.ScItemCodes
    
    params["ScItemIds"] = req.ScItemIds
    
    params["SellerNick"] = req.SellerNick
    
    params["StoreCodes"] = req.StoreCodes
    
    return params
}

// TaobaoInventoryQueryResponse 商家查询商品总体库存信息
type TaobaoInventoryQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemInventorys Object Array商品总体库存信息
    ItemInventorys InventorySum `json:"item_inventorys";xml:"item_inventorys"`
    
    // TipInfos Object Array提示信息，提示不存在的后端商品
    TipInfos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

// TaobaoInventoryInitialRequest 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账
type TaobaoInventoryInitialRequest struct {
    
    // Items required商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]
    Items string `json:"items";xml:"items"`
    
    // StoreCode required商家仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoInventoryInitialRequest) GetAPIName() string {
	return "taobao.inventory.initial"
}

func (req *TaobaoInventoryInitialRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Items"] = req.Items
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoInventoryInitialResponse 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账
type TaobaoInventoryInitialResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TipInfos Object Array提示信息
    TipInfos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

// TmallItemAddSimpleschemaGetRequest 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetRequest struct {
    
}

func (req *TmallItemAddSimpleschemaGetRequest) GetAPIName() string {
	return "tmall.item.add.simpleschema.get"
}

func (req *TmallItemAddSimpleschemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TmallItemAddSimpleschemaGetResponse 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic返回发布商品的规则文档
    Result string `json:"result";xml:"result"`
    
}

// TaobaoScitemMapAddRequest 创建IC商品或分销商品与后端商品的映射关系
type TaobaoScitemMapAddRequest struct {
    
    // ItemId required前台ic商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // NeedCheck optional默认值为false
    // true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联
    // false:不进行高级校验
    NeedCheck bool `json:"need_check";xml:"need_check"`
    
    // OuterCode optionalsc_item_id和outer_code 其中一个不能为空
    OuterCode string `json:"outer_code";xml:"outer_code"`
    
    // ScItemId optionalsc_item_id和outer_code 其中一个不能为空
    ScItemId int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    // SkuId optional前台ic商品skuid
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoScitemMapAddRequest) GetAPIName() string {
	return "taobao.scitem.map.add"
}

func (req *TaobaoScitemMapAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["ItemId"] = req.ItemId
    
    params["NeedCheck"] = req.NeedCheck
    
    params["OuterCode"] = req.OuterCode
    
    params["ScItemId"] = req.ScItemId
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoScitemMapAddResponse 创建IC商品或分销商品与后端商品的映射关系
type TaobaoScitemMapAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OuterCode Basic接口调用返回结果信息：商家编码
    OuterCode string `json:"outer_code";xml:"outer_code"`
    
}

// TaobaoInventoryAdjustTradeRequest 商家交易调整库存，淘宝交易、B2B经销等
type TaobaoInventoryAdjustTradeRequest struct {
    
    // BizUniqueCode required商家外部定单号
    BizUniqueCode string `json:"biz_unique_code";xml:"biz_unique_code"`
    
    // Items required商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}]
    Items string `json:"items";xml:"items"`
    
    // OperateTime required业务操作时间
    OperateTime time.Time `json:"operate_time";xml:"operate_time"`
    
    // TbOrderType required订单类型：B2C、B2B
    TbOrderType string `json:"tb_order_type";xml:"tb_order_type"`
    
}

func (req *TaobaoInventoryAdjustTradeRequest) GetAPIName() string {
	return "taobao.inventory.adjust.trade"
}

func (req *TaobaoInventoryAdjustTradeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["BizUniqueCode"] = req.BizUniqueCode
    
    params["Items"] = req.Items
    
    params["OperateTime"] = req.OperateTime
    
    params["TbOrderType"] = req.TbOrderType
    
    return params
}

// TaobaoInventoryAdjustTradeResponse 商家交易调整库存，淘宝交易、B2B经销等
type TaobaoInventoryAdjustTradeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OperateCode Basic操作返回码
    OperateCode string `json:"operate_code";xml:"operate_code"`
    
    // TipInfos Object Array提示信息
    TipInfos TipInfo `json:"tip_infos";xml:"tip_infos"`
    
}

// TaobaoFenxiaoProductcatUpdateRequest 修改产品线
type TaobaoFenxiaoProductcatUpdateRequest struct {
    
    // AgentCostPercent optional代销默认采购价比例，注意：100.00%，则输入为10000
    AgentCostPercent int64 `json:"agent_cost_percent";xml:"agent_cost_percent"`
    
    // DealerCostPercent optional经销默认采购价比例，注意：100.00%，则输入为10000
    DealerCostPercent int64 `json:"dealer_cost_percent";xml:"dealer_cost_percent"`
    
    // Name optional产品线名称
    Name string `json:"name";xml:"name"`
    
    // ProductLineId required产品线ID
    ProductLineId int64 `json:"product_line_id";xml:"product_line_id"`
    
    // RetailHighPercent optional最高零售价比例，注意：100.00%，则输入为10000
    RetailHighPercent int64 `json:"retail_high_percent";xml:"retail_high_percent"`
    
    // RetailLowPercent optional最低零售价比例，注意：100.00%，则输入为10000
    RetailLowPercent int64 `json:"retail_low_percent";xml:"retail_low_percent"`
    
}

func (req *TaobaoFenxiaoProductcatUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.update"
}

func (req *TaobaoFenxiaoProductcatUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["AgentCostPercent"] = req.AgentCostPercent
    
    params["DealerCostPercent"] = req.DealerCostPercent
    
    params["Name"] = req.Name
    
    params["ProductLineId"] = req.ProductLineId
    
    params["RetailHighPercent"] = req.RetailHighPercent
    
    params["RetailLowPercent"] = req.RetailLowPercent
    
    return params
}

// TaobaoFenxiaoProductcatUpdateResponse 修改产品线
type TaobaoFenxiaoProductcatUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoProductcatAddRequest 新增产品线
type TaobaoFenxiaoProductcatAddRequest struct {
    
    // AgentCostPercent required代销默认采购价比例，注意：100.00%，则输入为10000
    AgentCostPercent int64 `json:"agent_cost_percent";xml:"agent_cost_percent"`
    
    // DealerCostPercent required经销默认采购价比例，注意：100.00%，则输入为10000
    DealerCostPercent int64 `json:"dealer_cost_percent";xml:"dealer_cost_percent"`
    
    // Name required产品线名称
    Name string `json:"name";xml:"name"`
    
    // RetailHighPercent required最高零售价比例，注意：100.00%，则输入为10000
    RetailHighPercent int64 `json:"retail_high_percent";xml:"retail_high_percent"`
    
    // RetailLowPercent required最低零售价比例，注意：100.00%，则输入为10000
    RetailLowPercent int64 `json:"retail_low_percent";xml:"retail_low_percent"`
    
}

func (req *TaobaoFenxiaoProductcatAddRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.add"
}

func (req *TaobaoFenxiaoProductcatAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["AgentCostPercent"] = req.AgentCostPercent
    
    params["DealerCostPercent"] = req.DealerCostPercent
    
    params["Name"] = req.Name
    
    params["RetailHighPercent"] = req.RetailHighPercent
    
    params["RetailLowPercent"] = req.RetailLowPercent
    
    return params
}

// TaobaoFenxiaoProductcatAddResponse 新增产品线
type TaobaoFenxiaoProductcatAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // ProductLineId Basic产品线ID
    ProductLineId int64 `json:"product_line_id";xml:"product_line_id"`
    
}

// TaobaoFenxiaoGradeAddRequest 新建等级
type TaobaoFenxiaoGradeAddRequest struct {
    
    // Name required等级名称，等级名称不可重复
    Name string `json:"name";xml:"name"`
    
}

func (req *TaobaoFenxiaoGradeAddRequest) GetAPIName() string {
	return "taobao.fenxiao.grade.add"
}

func (req *TaobaoFenxiaoGradeAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Name"] = req.Name
    
    return params
}

// TaobaoFenxiaoGradeAddResponse 新建等级
type TaobaoFenxiaoGradeAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GradeId Basic等级ID
    GradeId int64 `json:"grade_id";xml:"grade_id"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoProductcatDeleteRequest 删除产品线
type TaobaoFenxiaoProductcatDeleteRequest struct {
    
    // ProductLineId required产品线ID
    ProductLineId int64 `json:"product_line_id";xml:"product_line_id"`
    
}

func (req *TaobaoFenxiaoProductcatDeleteRequest) GetAPIName() string {
	return "taobao.fenxiao.productcat.delete"
}

func (req *TaobaoFenxiaoProductcatDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ProductLineId"] = req.ProductLineId
    
    return params
}

// TaobaoFenxiaoProductcatDeleteResponse 删除产品线
type TaobaoFenxiaoProductcatDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TmallItemSchemaIncrementUpdateRequest 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
type TmallItemSchemaIncrementUpdateRequest struct {
    
    // ItemId required需要编辑的商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // XmlData required根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaIncrementUpdateRequest) GetAPIName() string {
	return "tmall.item.schema.increment.update"
}

func (req *TmallItemSchemaIncrementUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TmallItemSchemaIncrementUpdateResponse 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
type TmallItemSchemaIncrementUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic商品更新操作成功时间
    GmtModified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    // UpdateItemResult Basic返回商品发布结果
    UpdateItemResult string `json:"update_item_result";xml:"update_item_result"`
    
}

// TaobaoTopIpoutGetRequest 获取开放平台出口IP段
type TaobaoTopIpoutGetRequest struct {
    
}

func (req *TaobaoTopIpoutGetRequest) GetAPIName() string {
	return "taobao.top.ipout.get"
}

func (req *TaobaoTopIpoutGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoTopIpoutGetResponse 获取开放平台出口IP段
type TaobaoTopIpoutGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IpList BasicTOP网关出口IP列表
    IpList string `json:"ip_list";xml:"ip_list"`
    
}

// TmallItemIncrementUpdateSchemaGetRequest 增量方式修改天猫商品的规则获取的API。
//1.接口返回支持增量修改的字段以及相应字段的规则。
//2.如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好次字段以提升API整体性能）；
//3.ISV初次接入，开发阶段，此字段不填可以看到所有支持增量的字段；但是如果上线功能明确，请尽量遵守第2条
//4.如果ISV对字段规则非常清晰，可以直接组装入参数据提交到tmall.item.schema.increment.update进行数据更新。但是最好不要写死，比如每天还是有对此接口功能的一次比对。
//---（感谢爱慕旗舰店提供API命名）
type TmallItemIncrementUpdateSchemaGetRequest struct {
    
    // ItemId required需要编辑的商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // XmlData optional如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好此字段以提升API整体性能）
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemIncrementUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.item.increment.update.schema.get"
}

func (req *TmallItemIncrementUpdateSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TmallItemIncrementUpdateSchemaGetResponse 增量方式修改天猫商品的规则获取的API。
//1.接口返回支持增量修改的字段以及相应字段的规则。
//2.如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好次字段以提升API整体性能）；
//3.ISV初次接入，开发阶段，此字段不填可以看到所有支持增量的字段；但是如果上线功能明确，请尽量遵守第2条
//4.如果ISV对字段规则非常清晰，可以直接组装入参数据提交到tmall.item.schema.increment.update进行数据更新。但是最好不要写死，比如每天还是有对此接口功能的一次比对。
//---（感谢爱慕旗舰店提供API命名）
type TmallItemIncrementUpdateSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateItemResult Basic返回增量更新商品的规则文档
    UpdateItemResult string `json:"update_item_result";xml:"update_item_result"`
    
}

// TaobaoFenxiaoCooperationProductcatAddRequest 追加授权产品线
type TaobaoFenxiaoCooperationProductcatAddRequest struct {
    
    // CooperateId required合作关系id
    CooperateId int64 `json:"cooperate_id";xml:"cooperate_id"`
    
    // GradeId optional等级ID（为空则不修改）
    GradeId int64 `json:"grade_id";xml:"grade_id"`
    
    // ProductLineList required产品线id列表，若有多个，以逗号分隔
    ProductLineList string `json:"product_line_list";xml:"product_line_list"`
    
}

func (req *TaobaoFenxiaoCooperationProductcatAddRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.productcat.add"
}

func (req *TaobaoFenxiaoCooperationProductcatAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CooperateId"] = req.CooperateId
    
    params["GradeId"] = req.GradeId
    
    params["ProductLineList"] = req.ProductLineList
    
    return params
}

// TaobaoFenxiaoCooperationProductcatAddResponse 追加授权产品线
type TaobaoFenxiaoCooperationProductcatAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TmallItemSimpleschemaAddRequest 天猫简化版schema发布商品。
type TmallItemSimpleschemaAddRequest struct {
    
    // SchemaXmlFields required根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
    SchemaXmlFields string `json:"schema_xml_fields";xml:"schema_xml_fields"`
    
}

func (req *TmallItemSimpleschemaAddRequest) GetAPIName() string {
	return "tmall.item.simpleschema.add"
}

func (req *TmallItemSimpleschemaAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["SchemaXmlFields"] = req.SchemaXmlFields
    
    return params
}

// TmallItemSimpleschemaAddResponse 天猫简化版schema发布商品。
type TmallItemSimpleschemaAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic商品最后发布时间。
    GmtModified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    // Result Basic发布成功后返回商品ID
    Result string `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoCooperationAuditRequest 合作授权审批
type TaobaoFenxiaoCooperationAuditRequest struct {
    
    // AuditResult required1:审批通过，审批通过后要加入授权产品线列表；
    // 2:审批拒绝
    AuditResult int64 `json:"audit_result";xml:"audit_result"`
    
    // ProductLineListAgent optional当审批通过时需要指定授权产品线id列表(代销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。
    ProductLineListAgent string `json:"product_line_list_agent";xml:"product_line_list_agent"`
    
    // ProductLineListDealer optional当审批通过时需要指定授权产品线id列表(经销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。
    ProductLineListDealer string `json:"product_line_list_dealer";xml:"product_line_list_dealer"`
    
    // Remark required备注
    Remark string `json:"remark";xml:"remark"`
    
    // RequisitionId required合作申请Id
    RequisitionId int64 `json:"requisition_id";xml:"requisition_id"`
    
}

func (req *TaobaoFenxiaoCooperationAuditRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.audit"
}

func (req *TaobaoFenxiaoCooperationAuditRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["AuditResult"] = req.AuditResult
    
    params["ProductLineListAgent"] = req.ProductLineListAgent
    
    params["ProductLineListDealer"] = req.ProductLineListDealer
    
    params["Remark"] = req.Remark
    
    params["RequisitionId"] = req.RequisitionId
    
    return params
}

// TaobaoFenxiaoCooperationAuditResponse 合作授权审批
type TaobaoFenxiaoCooperationAuditResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoRequisitionsGetRequest 合作申请查询
type TaobaoFenxiaoRequisitionsGetRequest struct {
    
    // ApplyEnd optional申请结束时间yyyy-MM-dd
    ApplyEnd time.Time `json:"apply_end";xml:"apply_end"`
    
    // ApplyStart optional申请开始时间yyyy-MM-dd
    ApplyStart time.Time `json:"apply_start";xml:"apply_start"`
    
    // PageNo optional页码（大于0的整数，默认1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页记录数（默认20，最大50）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // Status optional申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
    Status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoFenxiaoRequisitionsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.requisitions.get"
}

func (req *TaobaoFenxiaoRequisitionsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["ApplyEnd"] = req.ApplyEnd
    
    params["ApplyStart"] = req.ApplyStart
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["Status"] = req.Status
    
    return params
}

// TaobaoFenxiaoRequisitionsGetResponse 合作申请查询
type TaobaoFenxiaoRequisitionsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // Requisitions Object Array合作申请
    Requisitions Requisition `json:"requisitions";xml:"requisitions"`
    
    // TotalResults Basic结果记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoLogisticsAddressReachableRequest 根据输入的目标地址，判断服务是否可达。
//现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
type TaobaoLogisticsAddressReachableRequest struct {
    
    // Address optional详细地址
    Address string `json:"address";xml:"address"`
    
    // AreaCode optional标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
    AreaCode string `json:"area_code";xml:"area_code"`
    
    // PartnerIds required物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
    PartnerIds string `json:"partner_ids";xml:"partner_ids"`
    
    // ServiceType required服务对应的数字编码，如揽派范围对应88
    ServiceType int64 `json:"service_type";xml:"service_type"`
    
    // SourceAreaCode optional发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
    SourceAreaCode string `json:"source_area_code";xml:"source_area_code"`
    
}

func (req *TaobaoLogisticsAddressReachableRequest) GetAPIName() string {
	return "taobao.logistics.address.reachable"
}

func (req *TaobaoLogisticsAddressReachableRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["Address"] = req.Address
    
    params["AreaCode"] = req.AreaCode
    
    params["PartnerIds"] = req.PartnerIds
    
    params["ServiceType"] = req.ServiceType
    
    params["SourceAreaCode"] = req.SourceAreaCode
    
    return params
}

// TaobaoLogisticsAddressReachableResponse 根据输入的目标地址，判断服务是否可达。
//现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
type TaobaoLogisticsAddressReachableResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ReachableResultList Object Array地址可达返回结果，每个TP对应一个
    ReachableResultList AddressReachableResult `json:"reachable_result_list";xml:"reachable_result_list"`
    
}

// TaobaoFenxiaoCooperationTerminateRequest 终止与分销商的合作关系
type TaobaoFenxiaoCooperationTerminateRequest struct {
    
    // CooperateId required合作编号
    CooperateId int64 `json:"cooperate_id";xml:"cooperate_id"`
    
    // EndRemainDays required终止天数，可选1,2,3,5,7,15，在多少天数内终止
    EndRemainDays int64 `json:"end_remain_days";xml:"end_remain_days"`
    
    // EndRemark required终止说明（5-2000字）
    EndRemark string `json:"end_remark";xml:"end_remark"`
    
}

func (req *TaobaoFenxiaoCooperationTerminateRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.terminate"
}

func (req *TaobaoFenxiaoCooperationTerminateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CooperateId"] = req.CooperateId
    
    params["EndRemainDays"] = req.EndRemainDays
    
    params["EndRemark"] = req.EndRemark
    
    return params
}

// TaobaoFenxiaoCooperationTerminateResponse 终止与分销商的合作关系
type TaobaoFenxiaoCooperationTerminateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoProductGradepriceGetRequest 等级折扣查询
type TaobaoFenxiaoProductGradepriceGetRequest struct {
    
    // ProductId required产品id
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // SkuId optionalskuId
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
    // TradeType optional经、代销模式（1：代销、2：经销）
    TradeType int64 `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductGradepriceGetRequest) GetAPIName() string {
	return "taobao.fenxiao.product.gradeprice.get"
}

func (req *TaobaoFenxiaoProductGradepriceGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ProductId"] = req.ProductId
    
    params["SkuId"] = req.SkuId
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoProductGradepriceGetResponse 等级折扣查询
type TaobaoFenxiaoProductGradepriceGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GradeDiscounts Object Array等级折扣列表
    GradeDiscounts GradeDiscount `json:"grade_discounts";xml:"grade_discounts"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoOmniorderStorecollectConsumeRequest 全渠道门店自提核销订单
type TaobaoOmniorderStorecollectConsumeRequest struct {
    
    // Code required核销码
    Code string `json:"code";xml:"code"`
    
    // MainOrderId required淘宝主订单ID
    MainOrderId int64 `json:"main_order_id";xml:"main_order_id"`
    
    // Operator optional核销操作人信息
    Operator string `json:"operator";xml:"operator"`
    
}

func (req *TaobaoOmniorderStorecollectConsumeRequest) GetAPIName() string {
	return "taobao.omniorder.storecollect.consume"
}

func (req *TaobaoOmniorderStorecollectConsumeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Code"] = req.Code
    
    params["MainOrderId"] = req.MainOrderId
    
    params["Operator"] = req.Operator
    
    return params
}

// TaobaoOmniorderStorecollectConsumeResponse 全渠道门店自提核销订单
type TaobaoOmniorderStorecollectConsumeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode Basic0表示成功
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // ErrMsg Basic核销错误信息
    ErrMsg string `json:"err_msg";xml:"err_msg"`
    
}

// TaobaoOmniorderStorecollectQueryRequest 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryRequest struct {
    
    // Code required核销码
    Code string `json:"code";xml:"code"`
    
}

func (req *TaobaoOmniorderStorecollectQueryRequest) GetAPIName() string {
	return "taobao.omniorder.storecollect.query"
}

func (req *TaobaoOmniorderStorecollectQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Code"] = req.Code
    
    return params
}

// TaobaoOmniorderStorecollectQueryResponse 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoQianniuKefuevalGetRequest 获取买家对客服的服务评价
type TaobaoQianniuKefuevalGetRequest struct {
    
    // Btime required开始时间，格式yyyyMMddHHmmss
    Btime string `json:"btime";xml:"btime"`
    
    // Etime required结束时间,格式yyyyMMddHHmmss
    Etime string `json:"etime";xml:"etime"`
    
    // QueryIds required客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
    QueryIds string `json:"query_ids";xml:"query_ids"`
    
}

func (req *TaobaoQianniuKefuevalGetRequest) GetAPIName() string {
	return "taobao.qianniu.kefueval.get"
}

func (req *TaobaoQianniuKefuevalGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Btime"] = req.Btime
    
    params["Etime"] = req.Etime
    
    params["QueryIds"] = req.QueryIds
    
    return params
}

// TaobaoQianniuKefuevalGetResponse 获取买家对客服的服务评价
type TaobaoQianniuKefuevalGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ResultCount BasicresultCount
    ResultCount int64 `json:"result_count";xml:"result_count"`
    
    // StaffEvalDetails Object ArraystaffEvalDetails
    StaffEvalDetails EvalDetail `json:"staff_eval_details";xml:"staff_eval_details"`
    
}

// CainiaoSmartdeliveryStrategyWarehouseIQueryRequest 智能发货引擎仓维度策略查询
type CainiaoSmartdeliveryStrategyWarehouseIQueryRequest struct {
    
    // ParamQueryDeliveryStrategyRequest optional查询请求参数
    ParamQueryDeliveryStrategyRequest QueryDeliveryStrategyRequest `json:"param_query_delivery_strategy_request";xml:"param_query_delivery_strategy_request"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.query"
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamQueryDeliveryStrategyRequest"] = req.ParamQueryDeliveryStrategyRequest
    
    return params
}

// CainiaoSmartdeliveryStrategyWarehouseIQueryResponse 智能发货引擎仓维度策略查询
type CainiaoSmartdeliveryStrategyWarehouseIQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DeliveryStrategyInfoList Object Array返回结果列表
    DeliveryStrategyInfoList DeliveryStrategyInfo `json:"delivery_strategy_info_list";xml:"delivery_strategy_info_list"`
    
}

// CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest 智能发货引擎发货策略设置仓维度
type CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest struct {
    
    // DeliveryStrategySetRequest required智能发货设置请求参数
    DeliveryStrategySetRequest DeliveryStrategySetRequest `json:"delivery_strategy_set_request";xml:"delivery_strategy_set_request"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["DeliveryStrategySetRequest"] = req.DeliveryStrategySetRequest
    
    return params
}

// CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse 智能发货引擎发货策略设置仓维度
type CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WarehouseInfo Object仓信息
    WarehouseInfo WarehouseDto `json:"warehouse_info";xml:"warehouse_info"`
    
}

// TaobaoFenxiaoProductImportFromAuctionRequest 供应商选择关联店铺的前台宝贝，导入生成产品
type TaobaoFenxiaoProductImportFromAuctionRequest struct {
    
    // AuctionId required店铺宝贝id
    AuctionId int64 `json:"auction_id";xml:"auction_id"`
    
    // ProductLineId required产品线id
    ProductLineId int64 `json:"product_line_id";xml:"product_line_id"`
    
    // TradeType required导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
    TradeType int64 `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductImportFromAuctionRequest) GetAPIName() string {
	return "taobao.fenxiao.product.import.from.auction"
}

func (req *TaobaoFenxiaoProductImportFromAuctionRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["AuctionId"] = req.AuctionId
    
    params["ProductLineId"] = req.ProductLineId
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoProductImportFromAuctionResponse 供应商选择关联店铺的前台宝贝，导入生成产品
type TaobaoFenxiaoProductImportFromAuctionResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OptTime Basic操作时间
    OptTime time.Time `json:"opt_time";xml:"opt_time"`
    
    // Pid Basic生成的产品id
    Pid int64 `json:"pid";xml:"pid"`
    
}

// TaobaoScitemOutercodeGetRequest 根据outerCode查询商品
type TaobaoScitemOutercodeGetRequest struct {
    
    // OuterCode required商品编码
    OuterCode string `json:"outer_code";xml:"outer_code"`
    
}

func (req *TaobaoScitemOutercodeGetRequest) GetAPIName() string {
	return "taobao.scitem.outercode.get"
}

func (req *TaobaoScitemOutercodeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["OuterCode"] = req.OuterCode
    
    return params
}

// TaobaoScitemOutercodeGetResponse 根据outerCode查询商品
type TaobaoScitemOutercodeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ScItem Object后台商品
    ScItem ScItem `json:"sc_item";xml:"sc_item"`
    
}

// TaobaoHttpdnsGetRequest 获取TOP DNS配置
type TaobaoHttpdnsGetRequest struct {
    
}

func (req *TaobaoHttpdnsGetRequest) GetAPIName() string {
	return "taobao.httpdns.get"
}

func (req *TaobaoHttpdnsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoHttpdnsGetResponse 获取TOP DNS配置
type TaobaoHttpdnsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result BasicHTTP DNS配置信息
    Result string `json:"result";xml:"result"`
    
}

// TaobaoScitemGetRequest 根据id查询商品
type TaobaoScitemGetRequest struct {
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoScitemGetRequest) GetAPIName() string {
	return "taobao.scitem.get"
}

func (req *TaobaoScitemGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TaobaoScitemGetResponse 根据id查询商品
type TaobaoScitemGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ScItem Object后端商品
    ScItem ScItem `json:"sc_item";xml:"sc_item"`
    
}

// TaobaoScitemQueryRequest 查询后端商品
type TaobaoScitemQueryRequest struct {
    
    // BarCode optional条形码
    BarCode string `json:"bar_code";xml:"bar_code"`
    
    // ItemName optional商品名称
    ItemName string `json:"item_name";xml:"item_name"`
    
    // ItemType optionalITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
    ItemType int64 `json:"item_type";xml:"item_type"`
    
    // OuterCode optional商家给商品的一个编码
    OuterCode string `json:"outer_code";xml:"outer_code"`
    
    // PageIndex optional当前页码数
    PageIndex int64 `json:"page_index";xml:"page_index"`
    
    // PageSize optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // WmsCode optional仓库编码
    WmsCode string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemQueryRequest) GetAPIName() string {
	return "taobao.scitem.query"
}

func (req *TaobaoScitemQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["BarCode"] = req.BarCode
    
    params["ItemName"] = req.ItemName
    
    params["ItemType"] = req.ItemType
    
    params["OuterCode"] = req.OuterCode
    
    params["PageIndex"] = req.PageIndex
    
    params["PageSize"] = req.PageSize
    
    params["WmsCode"] = req.WmsCode
    
    return params
}

// TaobaoScitemQueryResponse 查询后端商品
type TaobaoScitemQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // QueryPagination Object分页
    QueryPagination QueryPagination `json:"query_pagination";xml:"query_pagination"`
    
    // ScItemList Object ArrayList<ScItemDO>
    ScItemList ScItem `json:"sc_item_list";xml:"sc_item_list"`
    
    // TotalPage Basic商品条数
    TotalPage int64 `json:"total_page";xml:"total_page"`
    
}

// TaobaoScitemUpdateRequest 根据商品ID或商家编码修改后端商品
type TaobaoScitemUpdateRequest struct {
    
    // BarCode optional条形码
    BarCode string `json:"bar_code";xml:"bar_code"`
    
    // BrandId optional品牌id
    BrandId int64 `json:"brand_id";xml:"brand_id"`
    
    // BrandName optionalbrand_Name
    BrandName string `json:"brand_name";xml:"brand_name"`
    
    // Height optional高 单位：mm
    Height int64 `json:"height";xml:"height"`
    
    // IsAreaSale optional1表示区域销售，0或是空是非区域销售
    IsAreaSale int64 `json:"is_area_sale";xml:"is_area_sale"`
    
    // IsCostly optional是否是贵重品 0:不是 1：是
    IsCostly int64 `json:"is_costly";xml:"is_costly"`
    
    // IsDangerous optional是否危险 0：不是  0：是
    IsDangerous int64 `json:"is_dangerous";xml:"is_dangerous"`
    
    // IsFriable optional是否易碎 0：不是  1：是
    IsFriable int64 `json:"is_friable";xml:"is_friable"`
    
    // IsWarranty optional是否保质期：0:不是 1：是
    IsWarranty int64 `json:"is_warranty";xml:"is_warranty"`
    
    // ItemId optional后端商品ID，跟outer_code必须指定一个
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // ItemName optional商品名称
    ItemName string `json:"item_name";xml:"item_name"`
    
    // ItemType optional0.普通供应链商品 1.供应链组合主商品
    ItemType int64 `json:"item_type";xml:"item_type"`
    
    // Length optional长度 单位：mm
    Length int64 `json:"length";xml:"length"`
    
    // MatterStatus optional0:液体，1：粉体，2：固体
    MatterStatus int64 `json:"matter_status";xml:"matter_status"`
    
    // OuterCode special商家编码，跟item_id必须指定一个
    OuterCode string `json:"outer_code";xml:"outer_code"`
    
    // Price optionalprice
    Price int64 `json:"price";xml:"price"`
    
    // Remark optionalremark
    Remark string `json:"remark";xml:"remark"`
    
    // RemoveProperties optional移除商品属性P列表,P由系统分配：p1；p2
    RemoveProperties string `json:"remove_properties";xml:"remove_properties"`
    
    // SpuId optional淘宝SKU产品级编码CSPU ID
    SpuId int64 `json:"spu_id";xml:"spu_id"`
    
    // UpdateProperties optional需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
    UpdateProperties string `json:"update_properties";xml:"update_properties"`
    
    // Volume optional体积：立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Weight optionalweight
    Weight int64 `json:"weight";xml:"weight"`
    
    // Width optional宽 单位：mm
    Width int64 `json:"width";xml:"width"`
    
    // WmsCode optional仓储商编码
    WmsCode string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemUpdateRequest) GetAPIName() string {
	return "taobao.scitem.update"
}

func (req *TaobaoScitemUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 24)
    
    params["BarCode"] = req.BarCode
    
    params["BrandId"] = req.BrandId
    
    params["BrandName"] = req.BrandName
    
    params["Height"] = req.Height
    
    params["IsAreaSale"] = req.IsAreaSale
    
    params["IsCostly"] = req.IsCostly
    
    params["IsDangerous"] = req.IsDangerous
    
    params["IsFriable"] = req.IsFriable
    
    params["IsWarranty"] = req.IsWarranty
    
    params["ItemId"] = req.ItemId
    
    params["ItemName"] = req.ItemName
    
    params["ItemType"] = req.ItemType
    
    params["Length"] = req.Length
    
    params["MatterStatus"] = req.MatterStatus
    
    params["OuterCode"] = req.OuterCode
    
    params["Price"] = req.Price
    
    params["Remark"] = req.Remark
    
    params["RemoveProperties"] = req.RemoveProperties
    
    params["SpuId"] = req.SpuId
    
    params["UpdateProperties"] = req.UpdateProperties
    
    params["Volume"] = req.Volume
    
    params["Weight"] = req.Weight
    
    params["Width"] = req.Width
    
    params["WmsCode"] = req.WmsCode
    
    return params
}

// TaobaoScitemUpdateResponse 根据商品ID或商家编码修改后端商品
type TaobaoScitemUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateRows Basic更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据
    UpdateRows int64 `json:"update_rows";xml:"update_rows"`
    
}

// TaobaoScitemAddRequest 发布后端商品
type TaobaoScitemAddRequest struct {
    
    // BarCode optional条形码
    BarCode string `json:"bar_code";xml:"bar_code"`
    
    // BrandId optional品牌id
    BrandId int64 `json:"brand_id";xml:"brand_id"`
    
    // BrandName optionalbrand_Name
    BrandName string `json:"brand_name";xml:"brand_name"`
    
    // Height optional高 单位：mm
    Height int64 `json:"height";xml:"height"`
    
    // IsAreaSale optional1表示区域销售，0或是空是非区域销售
    IsAreaSale int64 `json:"is_area_sale";xml:"is_area_sale"`
    
    // IsCostly optional是否是贵重品 0:不是 1：是
    IsCostly int64 `json:"is_costly";xml:"is_costly"`
    
    // IsDangerous optional是否危险 0：不是  1：是
    IsDangerous int64 `json:"is_dangerous";xml:"is_dangerous"`
    
    // IsFriable optional是否易碎 0：不是  1：是
    IsFriable int64 `json:"is_friable";xml:"is_friable"`
    
    // IsWarranty optional是否保质期：0:不是 1：是
    IsWarranty int64 `json:"is_warranty";xml:"is_warranty"`
    
    // ItemName required商品名称
    ItemName string `json:"item_name";xml:"item_name"`
    
    // ItemType optional0.普通供应链商品 1.供应链组合主商品
    ItemType int64 `json:"item_type";xml:"item_type"`
    
    // Length optional长度 单位：mm
    Length int64 `json:"length";xml:"length"`
    
    // MatterStatus optional0:液体，1：粉体，2：固体
    MatterStatus int64 `json:"matter_status";xml:"matter_status"`
    
    // OuterCode required商家编码
    OuterCode string `json:"outer_code";xml:"outer_code"`
    
    // Price optional价格 单位：分
    Price int64 `json:"price";xml:"price"`
    
    // Properties optional商品属性格式是  p1:v1,p2:v2,p3:v3
    Properties string `json:"properties";xml:"properties"`
    
    // Remark optionalremark
    Remark string `json:"remark";xml:"remark"`
    
    // SpuId optionalspuId或是cspuid
    SpuId int64 `json:"spu_id";xml:"spu_id"`
    
    // Volume optional体积：立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Weight optional重量 单位：g
    Weight int64 `json:"weight";xml:"weight"`
    
    // Width optional宽 单位：mm
    Width int64 `json:"width";xml:"width"`
    
    // WmsCode optional仓储商编码
    WmsCode string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemAddRequest) GetAPIName() string {
	return "taobao.scitem.add"
}

func (req *TaobaoScitemAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 22)
    
    params["BarCode"] = req.BarCode
    
    params["BrandId"] = req.BrandId
    
    params["BrandName"] = req.BrandName
    
    params["Height"] = req.Height
    
    params["IsAreaSale"] = req.IsAreaSale
    
    params["IsCostly"] = req.IsCostly
    
    params["IsDangerous"] = req.IsDangerous
    
    params["IsFriable"] = req.IsFriable
    
    params["IsWarranty"] = req.IsWarranty
    
    params["ItemName"] = req.ItemName
    
    params["ItemType"] = req.ItemType
    
    params["Length"] = req.Length
    
    params["MatterStatus"] = req.MatterStatus
    
    params["OuterCode"] = req.OuterCode
    
    params["Price"] = req.Price
    
    params["Properties"] = req.Properties
    
    params["Remark"] = req.Remark
    
    params["SpuId"] = req.SpuId
    
    params["Volume"] = req.Volume
    
    params["Weight"] = req.Weight
    
    params["Width"] = req.Width
    
    params["WmsCode"] = req.WmsCode
    
    return params
}

// TaobaoScitemAddResponse 发布后端商品
type TaobaoScitemAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ScItem Object后台商品信息
    ScItem ScItem `json:"sc_item";xml:"sc_item"`
    
}

// TaobaoScitemMapDeleteRequest 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
type TaobaoScitemMapDeleteRequest struct {
    
    // ScItemId required后台商品ID
    ScItemId int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    // UserNick optional店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
    UserNick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoScitemMapDeleteRequest) GetAPIName() string {
	return "taobao.scitem.map.delete"
}

func (req *TaobaoScitemMapDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ScItemId"] = req.ScItemId
    
    params["UserNick"] = req.UserNick
    
    return params
}

// TaobaoScitemMapDeleteResponse 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
type TaobaoScitemMapDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Module Basic失效条数
    Module int64 `json:"module";xml:"module"`
    
}

// TaobaoScitemMapQueryRequest 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryRequest struct {
    
    // ItemId requiredmap_type为1：前台ic商品id
    // map_type为2：分销productid
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // SkuId optionalmap_type为1：前台ic商品skuid 
    // map_type为2：分销商品的skuid
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoScitemMapQueryRequest) GetAPIName() string {
	return "taobao.scitem.map.query"
}

func (req *TaobaoScitemMapQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["SkuId"] = req.SkuId
    
    return params
}

// TaobaoScitemMapQueryResponse 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ScItemMaps Object Array后端商品映射列表
    ScItemMaps ScItemMap `json:"sc_item_maps";xml:"sc_item_maps"`
    
}

// AlipayUserAccountFreezeGetRequest 查询支付宝账户冻结类型的冻结金额。
type AlipayUserAccountFreezeGetRequest struct {
    
    // FreezeType optional冻结类型，多个用,分隔。不传返回所有类型的冻结金额。
    // DEPOSIT_FREEZE,充值冻结
    // WITHDRAW_FREEZE,提现冻结
    // PAYMENT_FREEZE,支付冻结
    // BAIL_FREEZE,保证金冻结
    // CHARGE_FREEZE,收费冻结
    // PRE_DEPOSIT_FREEZE,预存款冻结
    // LOAN_FREEZE,贷款冻结
    // OTHER_FREEZE,其它
    FreezeType string `json:"freeze_type";xml:"freeze_type"`
    
}

func (req *AlipayUserAccountFreezeGetRequest) GetAPIName() string {
	return "alipay.user.account.freeze.get"
}

func (req *AlipayUserAccountFreezeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["FreezeType"] = req.FreezeType
    
    return params
}

// AlipayUserAccountFreezeGetResponse 查询支付宝账户冻结类型的冻结金额。
type AlipayUserAccountFreezeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // FreezeItems Object Array冻结金额列表
    FreezeItems AccountFreeze `json:"freeze_items";xml:"freeze_items"`
    
    // TotalResults Basic冻结总条数
    TotalResults string `json:"total_results";xml:"total_results"`
    
}

// AlipayUserAccountGetRequest 查询支付宝账户余额
type AlipayUserAccountGetRequest struct {
    
}

func (req *AlipayUserAccountGetRequest) GetAPIName() string {
	return "alipay.user.account.get"
}

func (req *AlipayUserAccountGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// AlipayUserAccountGetResponse 查询支付宝账户余额
type AlipayUserAccountGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AlipayAccount Object支付宝用户账户信息
    AlipayAccount AlipayAccount `json:"alipay_account";xml:"alipay_account"`
    
}

// TaobaoRdcAligeniusSendgoodsCancelRequest 提供商家在仅退款中发送取消发货状态
type TaobaoRdcAligeniusSendgoodsCancelRequest struct {
    
    // Param required请求参数
    Param CancelGoodsDto `json:"param";xml:"param"`
    
}

func (req *TaobaoRdcAligeniusSendgoodsCancelRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.sendgoods.cancel"
}

func (req *TaobaoRdcAligeniusSendgoodsCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param"] = req.Param
    
    return params
}

// TaobaoRdcAligeniusSendgoodsCancelResponse 提供商家在仅退款中发送取消发货状态
type TaobaoRdcAligeniusSendgoodsCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoProductMapAddRequest 创建分销和供应链商品映射关系。
type TaobaoFenxiaoProductMapAddRequest struct {
    
    // NotCheckOuterCode optional是否需要校验商家编码，true不校验，false校验。
    NotCheckOuterCode bool `json:"not_check_outer_code";xml:"not_check_outer_code"`
    
    // ProductId required分销产品id。
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // ScItemId optional后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。
    ScItemId int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    // ScItemIds optional在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。
    ScItemIds string `json:"sc_item_ids";xml:"sc_item_ids"`
    
    // SkuIds optional分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。
    SkuIds string `json:"sku_ids";xml:"sku_ids"`
    
}

func (req *TaobaoFenxiaoProductMapAddRequest) GetAPIName() string {
	return "taobao.fenxiao.product.map.add"
}

func (req *TaobaoFenxiaoProductMapAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["NotCheckOuterCode"] = req.NotCheckOuterCode
    
    params["ProductId"] = req.ProductId
    
    params["ScItemId"] = req.ScItemId
    
    params["ScItemIds"] = req.ScItemIds
    
    params["SkuIds"] = req.SkuIds
    
    return params
}

// TaobaoFenxiaoProductMapAddResponse 创建分销和供应链商品映射关系。
type TaobaoFenxiaoProductMapAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作结果
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// AlipayUserTradeSearchRequest 查询支付宝账户交易记录
type AlipayUserTradeSearchRequest struct {
    
    // AlipayOrderNo optional支付宝订单号，为空查询所有记录
    AlipayOrderNo string `json:"alipay_order_no";xml:"alipay_order_no"`
    
    // EndTime required结束时间。与开始时间间隔在七天之内
    EndTime string `json:"end_time";xml:"end_time"`
    
    // MerchantOrderNo optional商户订单号，为空查询所有记录
    MerchantOrderNo string `json:"merchant_order_no";xml:"merchant_order_no"`
    
    // OrderFrom optional订单来源，为空查询所有来源。淘宝(TAOBAO)，支付宝(ALIPAY)，其它(OTHER)
    OrderFrom string `json:"order_from";xml:"order_from"`
    
    // OrderStatus optional订单状态，为空查询所有状态订单
    OrderStatus string `json:"order_status";xml:"order_status"`
    
    // OrderType optional订单类型，为空查询所有类型订单。
    OrderType string `json:"order_type";xml:"order_type"`
    
    // PageNo required页码。取值范围:大于零的整数; 默认值1
    PageNo string `json:"page_no";xml:"page_no"`
    
    // PageSize required每页获取条数。最大值500。
    PageSize string `json:"page_size";xml:"page_size"`
    
    // StartTime required开始时间，时间必须是今天范围之内。格式为yyyy-MM-dd HH:mm:ss，精确到秒
    StartTime string `json:"start_time";xml:"start_time"`
    
}

func (req *AlipayUserTradeSearchRequest) GetAPIName() string {
	return "alipay.user.trade.search"
}

func (req *AlipayUserTradeSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["AlipayOrderNo"] = req.AlipayOrderNo
    
    params["EndTime"] = req.EndTime
    
    params["MerchantOrderNo"] = req.MerchantOrderNo
    
    params["OrderFrom"] = req.OrderFrom
    
    params["OrderStatus"] = req.OrderStatus
    
    params["OrderType"] = req.OrderType
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    return params
}

// AlipayUserTradeSearchResponse 查询支付宝账户交易记录
type AlipayUserTradeSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TotalPages Basic总页数
    TotalPages string `json:"total_pages";xml:"total_pages"`
    
    // TotalResults Basic总记录数
    TotalResults string `json:"total_results";xml:"total_results"`
    
    // TradeRecords Object Array交易记录列表
    TradeRecords TradeRecord `json:"trade_records";xml:"trade_records"`
    
}

// CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest 删除智能发货引擎仓策略
type CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest struct {
    
    // WarehouseId optional仓id
    WarehouseId int64 `json:"warehouse_id";xml:"warehouse_id"`
    
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetAPIName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

func (req *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WarehouseId"] = req.WarehouseId
    
    return params
}

// CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse 删除智能发货引擎仓策略
type CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsDeleteSuccess Basicdata
    IsDeleteSuccess bool `json:"is_delete_success";xml:"is_delete_success"`
    
}

// TmallItemShiptimeUpdateRequest 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
//1.
//    {
//        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
//        "updateType": 0 ---更新SKU
//    },
//
//    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空
//
//2.
//    {
//        "shipTimeType": 0, -- 删除发货时间
//        "updateType": 0 --更新SKU
//    },
//    按照指定SKU删除指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空
//
//3.
//
//    {
//        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
//        "updateType": 1 ---更新商品
//    },
//
//    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
//
//4.
//    {
//        "shipTimeType": 0, -- 删除发货时间
//        "updateType": 1 --更新商品
//    },
//    删除商品级的发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
type TmallItemShiptimeUpdateRequest struct {
    
    // ItemId required商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // Option required批量更新商品/SKU发货时间的备选项
    Option UpdateItemShipTimeOption `json:"option";xml:"option"`
    
    // ShipTime optional被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。
    ShipTime string `json:"ship_time";xml:"ship_time"`
    
    // SkuShipTimes optional被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。
    SkuShipTimes UpdateSkuShipTime `json:"sku_ship_times";xml:"sku_ship_times"`
    
}

func (req *TmallItemShiptimeUpdateRequest) GetAPIName() string {
	return "tmall.item.shiptime.update"
}

func (req *TmallItemShiptimeUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ItemId"] = req.ItemId
    
    params["Option"] = req.Option
    
    params["ShipTime"] = req.ShipTime
    
    params["SkuShipTimes"] = req.SkuShipTimes
    
    return params
}

// TmallItemShiptimeUpdateResponse 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
//1.
//    {
//        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
//        "updateType": 0 ---更新SKU
//    },
//
//    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空
//
//2.
//    {
//        "shipTimeType": 0, -- 删除发货时间
//        "updateType": 0 --更新SKU
//    },
//    按照指定SKU删除指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空
//
//3.
//
//    {
//        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
//        "updateType": 1 ---更新商品
//    },
//
//    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
//
//4.
//    {
//        "shipTimeType": 0, -- 删除发货时间
//        "updateType": 1 --更新商品
//    },
//    删除商品级的发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
type TmallItemShiptimeUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ShiptimeUpdateResult Basic被更新商品ID
    ShiptimeUpdateResult string `json:"shiptime_update_result";xml:"shiptime_update_result"`
    
}

// TmallInventoryQueryForstoreRequest 商家查询后端商品仓库库存
type TmallInventoryQueryForstoreRequest struct {
    
    // ParamList required查询列表
    ParamList InventoryQueryForStoreRequest `json:"param_list";xml:"param_list"`
    
}

func (req *TmallInventoryQueryForstoreRequest) GetAPIName() string {
	return "tmall.inventory.query.forstore"
}

func (req *TmallInventoryQueryForstoreRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamList"] = req.ParamList
    
    return params
}

// TmallInventoryQueryForstoreResponse 商家查询后端商品仓库库存
type TmallInventoryQueryForstoreResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Errorcode Basic错误code
    Errorcode string `json:"errorcode";xml:"errorcode"`
    
    // Errormessage Basic错误信息
    Errormessage string `json:"errormessage";xml:"errormessage"`
    
    // Issuccess Basic整体成功或失败
    Issuccess bool `json:"issuccess";xml:"issuccess"`
    
    // Result Object查询结果
    Result InventoryQueryResult `json:"result";xml:"result"`
    
}

// TaobaoTopAuthTokenRefreshRequest 根据refresh_token重新生成token
type TaobaoTopAuthTokenRefreshRequest struct {
    
    // RefreshToken requiredgrantType==refresh_token 时需要
    RefreshToken string `json:"refresh_token";xml:"refresh_token"`
    
}

func (req *TaobaoTopAuthTokenRefreshRequest) GetAPIName() string {
	return "taobao.top.auth.token.refresh"
}

func (req *TaobaoTopAuthTokenRefreshRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["RefreshToken"] = req.RefreshToken
    
    return params
}

// TaobaoTopAuthTokenRefreshResponse 根据refresh_token重新生成token
type TaobaoTopAuthTokenRefreshResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TokenResult Basic返回的是json信息
    TokenResult map[string]interface{} `json:"token_result";xml:"token_result"`
    
}

// TaobaoTopAuthTokenCreateRequest 用户通过code换获取access_token，https only
type TaobaoTopAuthTokenCreateRequest struct {
    
    // Code required授权code，grantType==authorization_code 时需要
    Code string `json:"code";xml:"code"`
    
    // Uuid optional与生成code的uuid配对
    Uuid string `json:"uuid";xml:"uuid"`
    
}

func (req *TaobaoTopAuthTokenCreateRequest) GetAPIName() string {
	return "taobao.top.auth.token.create"
}

func (req *TaobaoTopAuthTokenCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Code"] = req.Code
    
    params["Uuid"] = req.Uuid
    
    return params
}

// TaobaoTopAuthTokenCreateResponse 用户通过code换获取access_token，https only
type TaobaoTopAuthTokenCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TokenResult Basic返回的是json信息，和之前调用https://oauth.taobao.com/tac/token https://oauth.alibaba.com/token 换token返回的字段信息一致
    TokenResult map[string]interface{} `json:"token_result";xml:"token_result"`
    
}

// AlibabaEinvoiceClosereqRequest 关闭失败开票请求，避免造成重复开票
type AlibabaEinvoiceClosereqRequest struct {
    
    // PayeeRegisterNo required税号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // SerialNo required流水号
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoiceClosereqRequest) GetAPIName() string {
	return "alibaba.einvoice.closereq"
}

func (req *AlibabaEinvoiceClosereqRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["SerialNo"] = req.SerialNo
    
    return params
}

// AlibabaEinvoiceClosereqResponse 关闭失败开票请求，避免造成重复开票
type AlibabaEinvoiceClosereqResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic关闭是否成功
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoQimenOrderCallbackRequest 配送拦截
type TaobaoQimenOrderCallbackRequest struct {
    
    // DeliveryOrderCode optional奇门仓储字段,C123,string(50),,
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // ExpressCode optional运单号
    ExpressCode string `json:"expressCode";xml:"expressCode"`
    
    // OrderId optional奇门仓储字段,C123,string(50),,
    OrderId string `json:"orderId";xml:"orderId"`
    
    // OwnerCode optional奇门仓储字段,C123,string(50),,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // WarehouseCode optional奇门仓储字段,C123,string(50),,
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderCallbackRequest) GetAPIName() string {
	return "taobao.qimen.order.callback"
}

func (req *TaobaoQimenOrderCallbackRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["DeliveryOrderCode"] = req.DeliveryOrderCode
    
    params["ExpressCode"] = req.ExpressCode
    
    params["OrderId"] = req.OrderId
    
    params["OwnerCode"] = req.OwnerCode
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenOrderCallbackResponse 配送拦截
type TaobaoQimenOrderCallbackResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenWarehouseinfoQueryRequest 货主仓库资源查询
type TaobaoQimenWarehouseinfoQueryRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OwnerCode optional奇门仓储字段
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenWarehouseinfoQueryRequest) GetAPIName() string {
	return "taobao.qimen.warehouseinfo.query"
}

func (req *TaobaoQimenWarehouseinfoQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenWarehouseinfoQueryResponse 货主仓库资源查询
type TaobaoQimenWarehouseinfoQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OwnerCode Basic奇门仓储字段
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // OwnerName Basic奇门仓储字段
    OwnerName string `json:"ownerName";xml:"ownerName"`
    
    // WarehouseInfos Object Array奇门仓储字段
    WarehouseInfos WarehouseInfo `json:"warehouseInfos";xml:"warehouseInfos"`
    
}

// TaobaoQimenExpressinfoQueryRequest 配送公司信息查询
type TaobaoQimenExpressinfoQueryRequest struct {
    
    // ExpressCode optional奇门仓储字段
    ExpressCode string `json:"expressCode";xml:"expressCode"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OwnerCode optional奇门仓储字段
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenExpressinfoQueryRequest) GetAPIName() string {
	return "taobao.qimen.expressinfo.query"
}

func (req *TaobaoQimenExpressinfoQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ExpressCode"] = req.ExpressCode
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenExpressinfoQueryResponse 配送公司信息查询
type TaobaoQimenExpressinfoQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // ExpressInfos Object Array奇门仓储字段
    ExpressInfos ExpressInfo `json:"expressInfos";xml:"expressInfos"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoWlbWmsSnInfoQueryRequest 查询仓库作业的各类单据记录的Sn信息
type TaobaoWlbWmsSnInfoQueryRequest struct {
    
    // OrderCode required订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderCodeType required订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
    OrderCodeType int64 `json:"order_code_type";xml:"order_code_type"`
    
    // PageIndex required页数，默认每页50条
    PageIndex int64 `json:"page_index";xml:"page_index"`
    
}

func (req *TaobaoWlbWmsSnInfoQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.sn.info.query"
}

func (req *TaobaoWlbWmsSnInfoQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderCodeType"] = req.OrderCodeType
    
    params["PageIndex"] = req.PageIndex
    
    return params
}

// TaobaoWlbWmsSnInfoQueryResponse 查询仓库作业的各类单据记录的Sn信息
type TaobaoWlbWmsSnInfoQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object接口返回
    Result Result `json:"result";xml:"result"`
    
}

// CainiaoCloudprintCustomareaUpdateRequest 自定义区内容更新
type CainiaoCloudprintCustomareaUpdateRequest struct {
    
    // CustomAreaContent required自定义区内容（可修改）
    CustomAreaContent string `json:"custom_area_content";xml:"custom_area_content"`
    
    // CustomAreaId required自定义区id（不可修改）
    CustomAreaId int64 `json:"custom_area_id";xml:"custom_area_id"`
    
    // CustomAreaName required自定义区名称（可修改）
    CustomAreaName string `json:"custom_area_name";xml:"custom_area_name"`
    
}

func (req *CainiaoCloudprintCustomareaUpdateRequest) GetAPIName() string {
	return "cainiao.cloudprint.customarea.update"
}

func (req *CainiaoCloudprintCustomareaUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CustomAreaContent"] = req.CustomAreaContent
    
    params["CustomAreaId"] = req.CustomAreaId
    
    params["CustomAreaName"] = req.CustomAreaName
    
    return params
}

// CainiaoCloudprintCustomareaUpdateResponse 自定义区内容更新
type CainiaoCloudprintCustomareaUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// TaobaoQimenTransferorderQueryRequest 调拨单查询
type TaobaoQimenTransferorderQueryRequest struct {
    
    // ErpOrderCode optionalERP单号
    ErpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    // OwnerCode required111
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // TransferOrderCode required调拨单号
    TransferOrderCode string `json:"transferOrderCode";xml:"transferOrderCode"`
    
}

func (req *TaobaoQimenTransferorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.query"
}

func (req *TaobaoQimenTransferorderQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ErpOrderCode"] = req.ErpOrderCode
    
    params["OwnerCode"] = req.OwnerCode
    
    params["TransferOrderCode"] = req.TransferOrderCode
    
    return params
}

// TaobaoQimenTransferorderQueryResponse 调拨单查询
type TaobaoQimenTransferorderQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
    // TransferOrderDetail Object调拨单细节
    TransferOrderDetail TransferOrderDetail `json:"transferOrderDetail";xml:"transferOrderDetail"`
    
}

// TaobaoQimenTransferorderReportRequest 调拨单通知
type TaobaoQimenTransferorderReportRequest struct {
    
    // ConfirmInTime optional确认入库时间
    ConfirmInTime string `json:"confirmInTime";xml:"confirmInTime"`
    
    // ConfirmOutTime optional确认出库时间
    ConfirmOutTime string `json:"confirmOutTime";xml:"confirmOutTime"`
    
    // CreateTime optional调拨单创建时间
    CreateTime string `json:"createTime";xml:"createTime"`
    
    // ErpOrderCode optionalerpOrderCode
    ErpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    // FromWarehouseCode optional调拨出库仓编码
    FromWarehouseCode string `json:"fromWarehouseCode";xml:"fromWarehouseCode"`
    
    // Items optional项目集
    Items Items `json:"items";xml:"items"`
    
    // OrderStatus optionalorderStatus
    OrderStatus string `json:"orderStatus";xml:"orderStatus"`
    
    // OwnerCode optional111
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // ToWarehouseCode optional调拨入库仓编码
    ToWarehouseCode string `json:"toWarehouseCode";xml:"toWarehouseCode"`
    
    // TransferInOrderCode optional调拨入库单号
    TransferInOrderCode string `json:"transferInOrderCode";xml:"transferInOrderCode"`
    
    // TransferOrderCode optional调拨单号,0,string(50),必填,
    TransferOrderCode string `json:"transferOrderCode";xml:"transferOrderCode"`
    
    // TransferOutOrderCode optional调拨出库单号
    TransferOutOrderCode string `json:"transferOutOrderCode";xml:"transferOutOrderCode"`
    
}

func (req *TaobaoQimenTransferorderReportRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.report"
}

func (req *TaobaoQimenTransferorderReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["ConfirmInTime"] = req.ConfirmInTime
    
    params["ConfirmOutTime"] = req.ConfirmOutTime
    
    params["CreateTime"] = req.CreateTime
    
    params["ErpOrderCode"] = req.ErpOrderCode
    
    params["FromWarehouseCode"] = req.FromWarehouseCode
    
    params["Items"] = req.Items
    
    params["OrderStatus"] = req.OrderStatus
    
    params["OwnerCode"] = req.OwnerCode
    
    params["ToWarehouseCode"] = req.ToWarehouseCode
    
    params["TransferInOrderCode"] = req.TransferInOrderCode
    
    params["TransferOrderCode"] = req.TransferOrderCode
    
    params["TransferOutOrderCode"] = req.TransferOutOrderCode
    
    return params
}

// TaobaoQimenTransferorderReportResponse 调拨单通知
type TaobaoQimenTransferorderReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
}

// TmallChannelProductsQueryRequest 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
type TmallChannelProductsQueryRequest struct {
    
    // PageNum required页码数，从1开始
    PageNum int64 `json:"page_num";xml:"page_num"`
    
    // PageSize required分页大小
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ProductIds optional产品Id
    ProductIds int64 `json:"product_ids";xml:"product_ids"`
    
    // ProductLineId optional产品线Id
    ProductLineId int64 `json:"product_line_id";xml:"product_line_id"`
    
    // ProductNumber optional商家产品编码
    ProductNumber string `json:"product_number";xml:"product_number"`
    
    // SkuNumber optional商家SKU编码
    SkuNumber string `json:"sku_number";xml:"sku_number"`
    
}

func (req *TmallChannelProductsQueryRequest) GetAPIName() string {
	return "tmall.channel.products.query"
}

func (req *TmallChannelProductsQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["PageNum"] = req.PageNum
    
    params["PageSize"] = req.PageSize
    
    params["ProductIds"] = req.ProductIds
    
    params["ProductLineId"] = req.ProductLineId
    
    params["ProductNumber"] = req.ProductNumber
    
    params["SkuNumber"] = req.SkuNumber
    
    return params
}

// TmallChannelProductsQueryResponse 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
type TmallChannelProductsQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result PageResultDto `json:"result";xml:"result"`
    
}

// TaobaoQimenTransferorderCreateRequest 调拨单创建
type TaobaoQimenTransferorderCreateRequest struct {
    
    // Attributes optional扩展属性,HZ1234,string(500),,
    Attributes string `json:"attributes";xml:"attributes"`
    
    // ErpOrderCode optional外部ERP订单号,HZ1234,string(50),必填,
    ErpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    // ExpectStartTime optional期望调拨开始时间,Item1234,string(50),,
    ExpectStartTime string `json:"expectStartTime";xml:"expectStartTime"`
    
    // FromStoreCode optional出库仓编码,Item1234,string(50),必填,
    FromStoreCode string `json:"fromStoreCode";xml:"fromStoreCode"`
    
    // OwnerCode required111
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // ToStoreCode optional入库仓编码,HZ1234,string(50),必填,
    ToStoreCode string `json:"toStoreCode";xml:"toStoreCode"`
    
    // TransferItems optional项目集
    TransferItems TransferItems `json:"transferItems";xml:"transferItems"`
    
}

func (req *TaobaoQimenTransferorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.create"
}

func (req *TaobaoQimenTransferorderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["Attributes"] = req.Attributes
    
    params["ErpOrderCode"] = req.ErpOrderCode
    
    params["ExpectStartTime"] = req.ExpectStartTime
    
    params["FromStoreCode"] = req.FromStoreCode
    
    params["OwnerCode"] = req.OwnerCode
    
    params["ToStoreCode"] = req.ToStoreCode
    
    params["TransferItems"] = req.TransferItems
    
    return params
}

// TaobaoQimenTransferorderCreateResponse 调拨单创建
type TaobaoQimenTransferorderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码,0,string(50),,
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure,success,string(10),必填,
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息,invalid appkey,string(100),,
    Message string `json:"message";xml:"message"`
    
    // TransferExecuteInfo Object调拨单信息
    TransferExecuteInfo TransferExecuteInfo `json:"transferExecuteInfo";xml:"transferExecuteInfo"`
    
}

// CainiaoOpenstorageSellerResourcedetailGetRequest 商家资源详情获取（云打印开源存储）
type CainiaoOpenstorageSellerResourcedetailGetRequest struct {
    
    // SellerResourceId optional商家资源id
    SellerResourceId int64 `json:"seller_resource_id";xml:"seller_resource_id"`
    
}

func (req *CainiaoOpenstorageSellerResourcedetailGetRequest) GetAPIName() string {
	return "cainiao.openstorage.seller.resourcedetail.get"
}

func (req *CainiaoOpenstorageSellerResourcedetailGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["SellerResourceId"] = req.SellerResourceId
    
    return params
}

// CainiaoOpenstorageSellerResourcedetailGetResponse 商家资源详情获取（云打印开源存储）
type CainiaoOpenstorageSellerResourcedetailGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// CainiaoOpenstorageIsvResourcedetailGetRequest isv资源详情查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcedetailGetRequest struct {
    
    // IsvResourceId optionalisv资源id
    IsvResourceId int64 `json:"isv_resource_id";xml:"isv_resource_id"`
    
}

func (req *CainiaoOpenstorageIsvResourcedetailGetRequest) GetAPIName() string {
	return "cainiao.openstorage.isv.resourcedetail.get"
}

func (req *CainiaoOpenstorageIsvResourcedetailGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["IsvResourceId"] = req.IsvResourceId
    
    return params
}

// CainiaoOpenstorageIsvResourcedetailGetResponse isv资源详情查询（云打印开源存储）
type CainiaoOpenstorageIsvResourcedetailGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// QimenTaobaoIcpOrderStockinordermessagetoerpRequest 供应链创建的入库单推送商家ERP
type QimenTaobaoIcpOrderStockinordermessagetoerpRequest struct {
    
    // CustomerId required货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // EntryOrderlist optional入库单记录集
    EntryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockinordermessagetoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockinordermessagetoerp"
}

func (req *QimenTaobaoIcpOrderStockinordermessagetoerpRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["EntryOrderlist"] = req.EntryOrderlist
    
    return params
}

// QimenTaobaoIcpOrderStockinordermessagetoerpResponse 供应链创建的入库单推送商家ERP
type QimenTaobaoIcpOrderStockinordermessagetoerpResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoLogisticsTraceSearchRequest 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。
//此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
type TaobaoLogisticsTraceSearchRequest struct {
    
    // IsSplit optional表明是否是拆单，默认值0，1表示拆单
    IsSplit int64 `json:"is_split";xml:"is_split"`
    
    // SellerNick required卖家昵称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
    // SubTid optional拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。
    SubTid int64 `json:"sub_tid";xml:"sub_tid"`
    
    // Tid required淘宝交易号，请勿传非淘宝交易号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsTraceSearchRequest) GetAPIName() string {
	return "taobao.logistics.trace.search"
}

func (req *TaobaoLogisticsTraceSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["IsSplit"] = req.IsSplit
    
    params["SellerNick"] = req.SellerNick
    
    params["SubTid"] = req.SubTid
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsTraceSearchResponse 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。
//此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
type TaobaoLogisticsTraceSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CompanyName Basic物流公司名称
    CompanyName string `json:"company_name";xml:"company_name"`
    
    // OutSid Basic运单号
    OutSid string `json:"out_sid";xml:"out_sid"`
    
    // Status Basic订单的物流状态（仅支持线上发货online订单，线下发货offline发出后直接变为已签收）
    // * 等候发送给物流公司
    // *已提交给物流公司,等待物流公司接单
    // *已经确认消息接收，等待物流公司接单
    // *物流公司已接单
    // *物流公司不接单
    // *物流公司揽收失败
    // *物流公司揽收成功
    // *签收失败
    // *对方已签收
    // *对方拒绝签收
    Status string `json:"status";xml:"status"`
    
    // Tid Basic交易号
    Tid int64 `json:"tid";xml:"tid"`
    
    // TraceList Object Array流转信息列表
    TraceList TransitStepInfo `json:"trace_list";xml:"trace_list"`
    
}

// TaobaoRdcAligeniusDistributionLogisticsCancelRequest 截配状态回传接口
type TaobaoRdcAligeniusDistributionLogisticsCancelRequest struct {
    
    // Param0 optional参数
    Param0 CancelDistributionDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusDistributionLogisticsCancelRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.distribution.logistics.cancel"
}

func (req *TaobaoRdcAligeniusDistributionLogisticsCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param0"] = req.Param0
    
    return params
}

// TaobaoRdcAligeniusDistributionLogisticsCancelResponse 截配状态回传接口
type TaobaoRdcAligeniusDistributionLogisticsCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// QimenTaobaoIcpOrderStockinordecanceltoerpRequest 一盘货入库单取消消息通知
type QimenTaobaoIcpOrderStockinordecanceltoerpRequest struct {
    
    // CustomerId required货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // EntryOrderlist optional订单列表
    EntryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockinordecanceltoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockinordecanceltoerp"
}

func (req *QimenTaobaoIcpOrderStockinordecanceltoerpRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["EntryOrderlist"] = req.EntryOrderlist
    
    return params
}

// QimenTaobaoIcpOrderStockinordecanceltoerpResponse 一盘货入库单取消消息通知
type QimenTaobaoIcpOrderStockinordecanceltoerpResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Object
    Response Response `json:"response";xml:"response"`
    
}

// CainiaoEndpointLockerTopStationAddorupdateRequest 新增或者修改代收点相关信息
type CainiaoEndpointLockerTopStationAddorupdateRequest struct {
    
    // StationInfo optional站点信息
    StationInfo StationInfo `json:"station_info";xml:"station_info"`
    
}

func (req *CainiaoEndpointLockerTopStationAddorupdateRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.station.addorupdate"
}

func (req *CainiaoEndpointLockerTopStationAddorupdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["StationInfo"] = req.StationInfo
    
    return params
}

// CainiaoEndpointLockerTopStationAddorupdateResponse 新增或者修改代收点相关信息
type CainiaoEndpointLockerTopStationAddorupdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResult `json:"result";xml:"result"`
    
}

// TaobaoRdcAligeniusWarehouseReverseUpdateRequest 销退单状态回传
type TaobaoRdcAligeniusWarehouseReverseUpdateRequest struct {
    
    // Param0 required参数
    Param0 UpdateReverseStatusDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseReverseUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.update"
}

func (req *TaobaoRdcAligeniusWarehouseReverseUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param0"] = req.Param0
    
    return params
}

// TaobaoRdcAligeniusWarehouseReverseUpdateResponse 销退单状态回传
type TaobaoRdcAligeniusWarehouseReverseUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// AlibabaProviderEinvoiceCreateRequest 电子发票同步开票
type AlibabaProviderEinvoiceCreateRequest struct {
    
    // BizSign optional业务签名，用ca私钥做的业务签名
    BizSign string `json:"biz_sign";xml:"biz_sign"`
    
    // BusinessType required默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1
    BusinessType int64 `json:"business_type";xml:"business_type"`
    
    // ErpTid optionalERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
    ErpTid string `json:"erp_tid";xml:"erp_tid"`
    
    // InvoiceAmount required开票金额(对应新版的价税合计,价税合计=合计金额+合计税额)
    InvoiceAmount string `json:"invoice_amount";xml:"invoice_amount"`
    
    // InvoiceItems optional电子发票明细
    InvoiceItems InvoiceItem `json:"invoice_items";xml:"invoice_items"`
    
    // InvoiceMemo optional发票备注，有些省市会把此信息打印到PDF中
    InvoiceMemo string `json:"invoice_memo";xml:"invoice_memo"`
    
    // InvoiceTime required开票日期, 格式"YYYY-MM-DD HH:SS:MM"
    InvoiceTime time.Time `json:"invoice_time";xml:"invoice_time"`
    
    // InvoiceType required发票(开票)类型，蓝票blue,红票red，默认blue
    InvoiceType string `json:"invoice_type";xml:"invoice_type"`
    
    // NormalInvoiceCode optional原发票代码(开红票时传入)
    NormalInvoiceCode string `json:"normal_invoice_code";xml:"normal_invoice_code"`
    
    // NormalInvoiceNo optional原发票号码(开红票时传入)
    NormalInvoiceNo string `json:"normal_invoice_no";xml:"normal_invoice_no"`
    
    // PayeeAddress required开票方地址(新版中为必传)
    PayeeAddress string `json:"payee_address";xml:"payee_address"`
    
    // PayeeBankaccount optional开票方银行及 帐号
    PayeeBankaccount string `json:"payee_bankaccount";xml:"payee_bankaccount"`
    
    // PayeeChecker optional复核人
    PayeeChecker string `json:"payee_checker";xml:"payee_checker"`
    
    // PayeeName required开票方名称，公司名(如:XX商城)
    PayeeName string `json:"payee_name";xml:"payee_name"`
    
    // PayeeOperator required开票人
    PayeeOperator string `json:"payee_operator";xml:"payee_operator"`
    
    // PayeePhone required开票方 电话(新版中为必传)
    PayeePhone string `json:"payee_phone";xml:"payee_phone"`
    
    // PayeeReceiver optional收款人
    PayeeReceiver string `json:"payee_receiver";xml:"payee_receiver"`
    
    // PayeeRegisterNo required收款方税务登记证号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // PayerAddress optional消费者地址，开具增值税专用发票时必填，其他发票可以为空
    PayerAddress string `json:"payer_address";xml:"payer_address"`
    
    // PayerBankaccount optional付款方开票开户银行及账号
    PayerBankaccount string `json:"payer_bankaccount";xml:"payer_bankaccount"`
    
    // PayerEmail optional消费者电子邮箱
    PayerEmail string `json:"payer_email";xml:"payer_email"`
    
    // PayerName required付款方名称, 对应发票台头
    PayerName string `json:"payer_name";xml:"payer_name"`
    
    // PayerPhone optional消费者联系电话，开具增值税专用发票时必填，其他发票可以为空。
    PayerPhone string `json:"payer_phone";xml:"payer_phone"`
    
    // PayerRegisterNo optional付款方税务登记证号。
    PayerRegisterNo string `json:"payer_register_no";xml:"payer_register_no"`
    
    // PlatformCode required电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    PlatformCode string `json:"platform_code";xml:"platform_code"`
    
    // PlatformTid required电商平台对应的订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
    // SerialNo required开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
    // SumPrice required合计金额(新版中为必传)
    SumPrice string `json:"sum_price";xml:"sum_price"`
    
    // SumTax required合计税额(新版中为必传)
    SumTax string `json:"sum_tax";xml:"sum_tax"`
    
}

func (req *AlibabaProviderEinvoiceCreateRequest) GetAPIName() string {
	return "alibaba.provider.einvoice.create"
}

func (req *AlibabaProviderEinvoiceCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 29)
    
    params["BizSign"] = req.BizSign
    
    params["BusinessType"] = req.BusinessType
    
    params["ErpTid"] = req.ErpTid
    
    params["InvoiceAmount"] = req.InvoiceAmount
    
    params["InvoiceItems"] = req.InvoiceItems
    
    params["InvoiceMemo"] = req.InvoiceMemo
    
    params["InvoiceTime"] = req.InvoiceTime
    
    params["InvoiceType"] = req.InvoiceType
    
    params["NormalInvoiceCode"] = req.NormalInvoiceCode
    
    params["NormalInvoiceNo"] = req.NormalInvoiceNo
    
    params["PayeeAddress"] = req.PayeeAddress
    
    params["PayeeBankaccount"] = req.PayeeBankaccount
    
    params["PayeeChecker"] = req.PayeeChecker
    
    params["PayeeName"] = req.PayeeName
    
    params["PayeeOperator"] = req.PayeeOperator
    
    params["PayeePhone"] = req.PayeePhone
    
    params["PayeeReceiver"] = req.PayeeReceiver
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["PayerAddress"] = req.PayerAddress
    
    params["PayerBankaccount"] = req.PayerBankaccount
    
    params["PayerEmail"] = req.PayerEmail
    
    params["PayerName"] = req.PayerName
    
    params["PayerPhone"] = req.PayerPhone
    
    params["PayerRegisterNo"] = req.PayerRegisterNo
    
    params["PlatformCode"] = req.PlatformCode
    
    params["PlatformTid"] = req.PlatformTid
    
    params["SerialNo"] = req.SerialNo
    
    params["SumPrice"] = req.SumPrice
    
    params["SumTax"] = req.SumTax
    
    return params
}

// AlibabaProviderEinvoiceCreateResponse 电子发票同步开票
type AlibabaProviderEinvoiceCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AntiFakeCode Basic防伪码
    AntiFakeCode string `json:"anti_fake_code";xml:"anti_fake_code"`
    
    // Ciphertext Basic发票密文，密码区的字符串
    Ciphertext string `json:"ciphertext";xml:"ciphertext"`
    
    // DeviceNo Basic税控设备编号(新版电子发票有)
    DeviceNo string `json:"device_no";xml:"device_no"`
    
    // ErpTid Basicerp自定义单据号
    ErpTid string `json:"erp_tid";xml:"erp_tid"`
    
    // FileDataType Basic文件类型(pdf,jpg,png)
    FileDataType string `json:"file_data_type";xml:"file_data_type"`
    
    // InvoiceAmount Basic实际开票金额
    InvoiceAmount string `json:"invoice_amount";xml:"invoice_amount"`
    
    // InvoiceCode Basic发票代码
    InvoiceCode string `json:"invoice_code";xml:"invoice_code"`
    
    // InvoiceDate Basic开票日期
    InvoiceDate string `json:"invoice_date";xml:"invoice_date"`
    
    // InvoiceFileData Basic发票文件PDF内容，PDF的byte[]用base64编码后的字段串。
    InvoiceFileData string `json:"invoice_file_data";xml:"invoice_file_data"`
    
    // InvoiceNo Basic发票号码
    InvoiceNo string `json:"invoice_no";xml:"invoice_no"`
    
    // PlatformCode Basic电商平台身份标识码，如taobao,tmall
    PlatformCode string `json:"platform_code";xml:"platform_code"`
    
    // PlatformTid Basic电商平台订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
    // QrCode Basic二维码
    QrCode string `json:"qr_code";xml:"qr_code"`
    
    // SerialNo Basic电子发票流水号，标记业务上的唯一请求
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

// TaobaoRdcAligeniusRefundsCheckRequest 根据退款信息，对退款单进行审核
type TaobaoRdcAligeniusRefundsCheckRequest struct {
    
    // Param required入参
    Param RefundCheckDto `json:"param";xml:"param"`
    
}

func (req *TaobaoRdcAligeniusRefundsCheckRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.refunds.check"
}

func (req *TaobaoRdcAligeniusRefundsCheckRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param"] = req.Param
    
    return params
}

// TaobaoRdcAligeniusRefundsCheckResponse 根据退款信息，对退款单进行审核
type TaobaoRdcAligeniusRefundsCheckResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoRpReturngoodsAgreeRequest 卖家同意退货，支持淘宝和天猫的订单。
type TaobaoRpReturngoodsAgreeRequest struct {
    
    // Address optional卖家提供的退货地址，淘宝退款为必填项。
    Address string `json:"address";xml:"address"`
    
    // Mobile optional卖家手机，淘宝退款为必填项。
    Mobile string `json:"mobile";xml:"mobile"`
    
    // Name optional卖家姓名，淘宝退款为必填项。
    Name string `json:"name";xml:"name"`
    
    // Post optional卖家提供的退货地址的邮编，淘宝退款为必填项。
    Post string `json:"post";xml:"post"`
    
    // PostFeeBearRole optional邮费承担方，买家承担值为1，卖家承担值为0
    PostFeeBearRole int64 `json:"post_fee_bear_role";xml:"post_fee_bear_role"`
    
    // RefundId required退款编号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase optional售中：onsale，售后：aftersale，天猫退款为必填项。
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
    // RefundVersion optional退款版本号，天猫退款为必填项。
    RefundVersion int64 `json:"refund_version";xml:"refund_version"`
    
    // Remark optional卖家退货留言，天猫退款为必填项。
    Remark string `json:"remark";xml:"remark"`
    
    // SellerAddressId optional卖家收货地址编号，天猫淘宝退款都为必填项。
    SellerAddressId int64 `json:"seller_address_id";xml:"seller_address_id"`
    
    // Tel optional卖家座机，淘宝退款为必填项。
    Tel string `json:"tel";xml:"tel"`
    
}

func (req *TaobaoRpReturngoodsAgreeRequest) GetAPIName() string {
	return "taobao.rp.returngoods.agree"
}

func (req *TaobaoRpReturngoodsAgreeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 11)
    
    params["Address"] = req.Address
    
    params["Mobile"] = req.Mobile
    
    params["Name"] = req.Name
    
    params["Post"] = req.Post
    
    params["PostFeeBearRole"] = req.PostFeeBearRole
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    params["RefundVersion"] = req.RefundVersion
    
    params["Remark"] = req.Remark
    
    params["SellerAddressId"] = req.SellerAddressId
    
    params["Tel"] = req.Tel
    
    return params
}

// TaobaoRpReturngoodsAgreeResponse 卖家同意退货，支持淘宝和天猫的订单。
type TaobaoRpReturngoodsAgreeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoRdcAligeniusAutorefundsQueryRequest （此接口后期将不再维护，请勿使用）供第三方商家查询授权给自己的所有退款款订单的退款信息
type TaobaoRdcAligeniusAutorefundsQueryRequest struct {
    
    // EndTime required查询数据时间段结束时间
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // PageNo required页数
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize required每页数据数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime required查询数据时间段开始时间
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoRdcAligeniusAutorefundsQueryRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.autorefunds.query"
}

func (req *TaobaoRdcAligeniusAutorefundsQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["EndTime"] = req.EndTime
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    return params
}

// TaobaoRdcAligeniusAutorefundsQueryResponse （此接口后期将不再维护，请勿使用）供第三方商家查询授权给自己的所有退款款订单的退款信息
type TaobaoRdcAligeniusAutorefundsQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoTrademonitorGetRequest 仅限供应商调用此接口查询经销商品监控信息
type TaobaoFenxiaoTrademonitorGetRequest struct {
    
    // DistributorNick optional经销商的淘宝账号
    DistributorNick string `json:"distributor_nick";xml:"distributor_nick"`
    
    // EndCreated optional结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。
    EndCreated time.Time `json:"end_created";xml:"end_created"`
    
    // Fields optional多个字段用","分隔。 fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。例如：trade_monitors.item_title表示只返回item_title
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。（大于0的整数。小于1按1计）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。（每页条数不超过50条，超过50或小于0均按50计）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ProductId optional产品id
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // StartCreated optional起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。
    StartCreated time.Time `json:"start_created";xml:"start_created"`
    
}

func (req *TaobaoFenxiaoTrademonitorGetRequest) GetAPIName() string {
	return "taobao.fenxiao.trademonitor.get"
}

func (req *TaobaoFenxiaoTrademonitorGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["DistributorNick"] = req.DistributorNick
    
    params["EndCreated"] = req.EndCreated
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["ProductId"] = req.ProductId
    
    params["StartCreated"] = req.StartCreated
    
    return params
}

// TaobaoFenxiaoTrademonitorGetResponse 仅限供应商调用此接口查询经销商品监控信息
type TaobaoFenxiaoTrademonitorGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TotalResults Basic搜索到的经销商品订单数量
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
    // TradeMonitors Object Array经销商品订单监控信息
    TradeMonitors TradeMonitor `json:"trade_monitors";xml:"trade_monitors"`
    
}

// TaobaoRpRefundsAgreeRequest 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
type TaobaoRpRefundsAgreeRequest struct {
    
    // Code optional短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。
    Code string `json:"code";xml:"code"`
    
    // RefundInfos required退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。
    RefundInfos string `json:"refund_infos";xml:"refund_infos"`
    
}

func (req *TaobaoRpRefundsAgreeRequest) GetAPIName() string {
	return "taobao.rp.refunds.agree"
}

func (req *TaobaoRpRefundsAgreeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Code"] = req.Code
    
    params["RefundInfos"] = req.RefundInfos
    
    return params
}

// TaobaoRpRefundsAgreeResponse 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
type TaobaoRpRefundsAgreeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Message Basic信息
    Message string `json:"message";xml:"message"`
    
    // MsgCode Basic批量退款操作情况，可选值：OP_SUCC（全部成功），SOME_OP_SUCC（部分成功），OP_FAILURE_UE（全部失败）
    MsgCode string `json:"msg_code";xml:"msg_code"`
    
    // Results Object Array退款操作结果列表
    Results RefundMappingResult `json:"results";xml:"results"`
    
    // Succ Basic操作成功
    Succ bool `json:"succ";xml:"succ"`
    
}

// TaobaoItemAnchorGetRequest 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
type TaobaoItemAnchorGetRequest struct {
    
    // CatId required对应类目编号
    CatId int64 `json:"cat_id";xml:"cat_id"`
    
    // Type required宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItemAnchorGetRequest) GetAPIName() string {
	return "taobao.item.anchor.get"
}

func (req *TaobaoItemAnchorGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CatId"] = req.CatId
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoItemAnchorGetResponse 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
type TaobaoItemAnchorGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AnchorModules Object Array宝贝描述规范化可使用打标模块的锚点信息
    AnchorModules IdsModule `json:"anchor_modules";xml:"anchor_modules"`
    
    // TotalResults Basic返回的宝贝描述模板结果数目
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoRefundRefuseRequest 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：
//1. 传入的refund_id和相应的tid, oid必须匹配
//2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
//3. 只有卖家才能执行拒绝退款操作
//4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
type TaobaoRefundRefuseRequest struct {
    
    // Oid optional退款记录对应的交易子订单号
    Oid int64 `json:"oid";xml:"oid"`
    
    // RefundId required退款单号
    RefundId int64 `json:"refund_id";xml:"refund_id"`
    
    // RefundPhase optional可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。
    RefundPhase string `json:"refund_phase";xml:"refund_phase"`
    
    // RefundVersion optional退款版本号，天猫退款为必填项。
    RefundVersion int64 `json:"refund_version";xml:"refund_version"`
    
    // RefuseMessage required拒绝退款时的说明信息，长度2-200
    RefuseMessage string `json:"refuse_message";xml:"refuse_message"`
    
    // RefuseProof optional拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。
    RefuseProof []byte `json:"refuse_proof";xml:"refuse_proof"`
    
    // RefuseReasonId optional拒绝原因编号，会提供用户拒绝原因列表供选择
    RefuseReasonId int64 `json:"refuse_reason_id";xml:"refuse_reason_id"`
    
    // Tid optional退款记录对应的交易订单号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoRefundRefuseRequest) GetAPIName() string {
	return "taobao.refund.refuse"
}

func (req *TaobaoRefundRefuseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["Oid"] = req.Oid
    
    params["RefundId"] = req.RefundId
    
    params["RefundPhase"] = req.RefundPhase
    
    params["RefundVersion"] = req.RefundVersion
    
    params["RefuseMessage"] = req.RefuseMessage
    
    params["RefuseProof"] = req.RefuseProof
    
    params["RefuseReasonId"] = req.RefuseReasonId
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoRefundRefuseResponse 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：
//1. 传入的refund_id和相应的tid, oid必须匹配
//2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
//3. 只有卖家才能执行拒绝退款操作
//4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
type TaobaoRefundRefuseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic拒绝退款操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // Refund Object拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段
    Refund Refund `json:"refund";xml:"refund"`
    
}

// TaobaoTradeAmountGetRequest 卖家查询该笔交易的资金帐务相关的数据；
//1. 只供卖家使用，买家不可使用
//2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
type TaobaoTradeAmountGetRequest struct {
    
    // Fields required订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)
    Fields string `json:"fields";xml:"fields"`
    
    // Tid required交易编号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeAmountGetRequest) GetAPIName() string {
	return "taobao.trade.amount.get"
}

func (req *TaobaoTradeAmountGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeAmountGetResponse 卖家查询该笔交易的资金帐务相关的数据；
//1. 只供卖家使用，买家不可使用
//2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
type TaobaoTradeAmountGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TradeAmount Object主订单的财务信息详情
    TradeAmount TradeAmount `json:"trade_amount";xml:"trade_amount"`
    
}

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest 供应商修改经销采购单备注
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
    
    // DealerOrderId required经销采购单ID
    DealerOrderId int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
    // SupplierMemo optional备注留言，可为空
    SupplierMemo string `json:"supplier_memo";xml:"supplier_memo"`
    
    // SupplierMemoFlag required旗子的标记，必选。
    // 1-5之间的数字。
    // 非1-5之间，都采用1作为默认。
    // 1:红色
    // 2:黄色
    // 3:绿色
    // 4:蓝色
    // 5:粉红色
    SupplierMemoFlag int64 `json:"supplier_memo_flag";xml:"supplier_memo_flag"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}

func (req *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["DealerOrderId"] = req.DealerOrderId
    
    params["SupplierMemo"] = req.SupplierMemo
    
    params["SupplierMemoFlag"] = req.SupplierMemoFlag
    
    return params
}

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse 供应商修改经销采购单备注
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoWlbWmsItemCombinationCreateRequest 创建组合商品与子商品关系
type TaobaoWlbWmsItemCombinationCreateRequest struct {
    
    // ItemId required组合商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // SubItemList required子货品列表
    SubItemList SubItemList `json:"sub_item_list";xml:"sub_item_list"`
    
}

func (req *TaobaoWlbWmsItemCombinationCreateRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.create"
}

func (req *TaobaoWlbWmsItemCombinationCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["SubItemList"] = req.SubItemList
    
    return params
}

// TaobaoWlbWmsItemCombinationCreateResponse 创建组合商品与子商品关系
type TaobaoWlbWmsItemCombinationCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoWlbWmsItemCombinationDeleteRequest 删除货品组合关系
type TaobaoWlbWmsItemCombinationDeleteRequest struct {
    
    // ItemId optional组合货品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbWmsItemCombinationDeleteRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.delete"
}

func (req *TaobaoWlbWmsItemCombinationDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TaobaoWlbWmsItemCombinationDeleteResponse 删除货品组合关系
type TaobaoWlbWmsItemCombinationDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoWlbWmsItemCombinationGetRequest 查询组合商品的组合关系
type TaobaoWlbWmsItemCombinationGetRequest struct {
    
    // Itemid required货品Id
    Itemid int64 `json:"itemid";xml:"itemid"`
    
}

func (req *TaobaoWlbWmsItemCombinationGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.item.combination.get"
}

func (req *TaobaoWlbWmsItemCombinationGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Itemid"] = req.Itemid
    
    return params
}

// TaobaoWlbWmsItemCombinationGetResponse 查询组合商品的组合关系
type TaobaoWlbWmsItemCombinationGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object接口返回结果
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoProductGradepriceUpdateRequest 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
type TaobaoFenxiaoProductGradepriceUpdateRequest struct {
    
    // Ids required会员等级的id或者分销商id，例如：”1001,2001,1002”
    Ids int64 `json:"ids";xml:"ids"`
    
    // Prices required优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25
    Prices string `json:"prices";xml:"prices"`
    
    // ProductId required产品Id
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // SkuId optionalskuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
    // TargetType required选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR"
    TargetType string `json:"target_type";xml:"target_type"`
    
    // TradeType optional交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销)
    TradeType string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoProductGradepriceUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.gradeprice.update"
}

func (req *TaobaoFenxiaoProductGradepriceUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Ids"] = req.Ids
    
    params["Prices"] = req.Prices
    
    params["ProductId"] = req.ProductId
    
    params["SkuId"] = req.SkuId
    
    params["TargetType"] = req.TargetType
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoProductGradepriceUpdateResponse 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
type TaobaoFenxiaoProductGradepriceUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic返回操作结果：成功或失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoRdcAligeniusOrderReturngoodsNotifyRequest 退货单创建结果反馈
type TaobaoRdcAligeniusOrderReturngoodsNotifyRequest struct {
    
    // ParentOrderId optional主订单号
    ParentOrderId int64 `json:"parent_order_id";xml:"parent_order_id"`
    
    // RefundReturnNotes optional退货单创建结果列表
    RefundReturnNotes RefundReturnNotes `json:"refund_return_notes";xml:"refund_return_notes"`
    
}

func (req *TaobaoRdcAligeniusOrderReturngoodsNotifyRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.order.returngoods.notify"
}

func (req *TaobaoRdcAligeniusOrderReturngoodsNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ParentOrderId"] = req.ParentOrderId
    
    params["RefundReturnNotes"] = req.RefundReturnNotes
    
    return params
}

// TaobaoRdcAligeniusOrderReturngoodsNotifyResponse 退货单创建结果反馈
type TaobaoRdcAligeniusOrderReturngoodsNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basicsuccess
    Result bool `json:"result";xml:"result"`
    
    // ResultCode BasicerrorCode
    ResultCode string `json:"result_code";xml:"result_code"`
    
    // ResultInfo BasicerrorInfo
    ResultInfo string `json:"result_info";xml:"result_info"`
    
}

// TaobaoOmniorderStoreSdtcancelRequest 通知速店通取消取号
type TaobaoOmniorderStoreSdtcancelRequest struct {
    
    // PackageId required取号返回的packageId
    PackageId int64 `json:"package_id";xml:"package_id"`
    
}

func (req *TaobaoOmniorderStoreSdtcancelRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtcancel"
}

func (req *TaobaoOmniorderStoreSdtcancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["PackageId"] = req.PackageId
    
    return params
}

// TaobaoOmniorderStoreSdtcancelResponse 通知速店通取消取号
type TaobaoOmniorderStoreSdtcancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoOmniorderStoreSdtquerystationRequest 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationRequest struct {
    
    // ParamLong2 required取号时返回的packageId
    ParamLong2 int64 `json:"param_long2";xml:"param_long2"`
    
}

func (req *TaobaoOmniorderStoreSdtquerystationRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtquerystation"
}

func (req *TaobaoOmniorderStoreSdtquerystationRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamLong2"] = req.ParamLong2
    
    return params
}

// TaobaoOmniorderStoreSdtquerystationResponse 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest 在自动化任务核对地址卡片中，透出修改地址选择器，让用户自己选择要修改的地址，然后同步到商家后台ERP系统
type QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest struct {
    
    // BizOrderId required交易订单ID
    BizOrderId string `json:"bizOrderId";xml:"bizOrderId"`
    
    // BuyerNick required买家账号名
    BuyerNick string `json:"buyerNick";xml:"buyerNick"`
    
    // ModifiedAddress required要修改的地址信息
    ModifiedAddress ModifiedAddress `json:"modifiedAddress";xml:"modifiedAddress"`
    
    // OriginalAddress optional订单原始收货地址信息
    OriginalAddress OriginalAddress `json:"originalAddress";xml:"originalAddress"`
    
    // SellerNick required店铺主账号
    SellerNick string `json:"sellerNick";xml:"sellerNick"`
    
}

func (req *QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest) GetAPIName() string {
	return "qimen.taobao.qianniu.cloudkefu.address.self.modify"
}

func (req *QimenTaobaoQianniuCloudkefuAddressSelfModifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["BizOrderId"] = req.BizOrderId
    
    params["BuyerNick"] = req.BuyerNick
    
    params["ModifiedAddress"] = req.ModifiedAddress
    
    params["OriginalAddress"] = req.OriginalAddress
    
    params["SellerNick"] = req.SellerNick
    
    return params
}

// QimenTaobaoQianniuCloudkefuAddressSelfModifyResponse 在自动化任务核对地址卡片中，透出修改地址选择器，让用户自己选择要修改的地址，然后同步到商家后台ERP系统
type QimenTaobaoQianniuCloudkefuAddressSelfModifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object修改地址返回结果
    Result ResultDO `json:"result";xml:"result"`
    
}

// TmallItemSizemappingTemplateCreateRequest 新增天猫商品尺码表模板
//
//男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：
//脚长（cm） 必选
//
//内衣-文胸类目，尺码表维度为：
//上胸围（cm） 必选
//下胸围（cm） 必选
//上下胸围差（cm） 必选
//身高（cm）
//体重（公斤）
//
//内衣-内裤类目，尺码表维度为：
//腰围（cm） 必选
//臀围（cm） 必选
//身高（cm）
//体重（公斤）
//裤长（cm）
//裆部（cm）
//脚口（cm）
//腿围（cm）
//
//内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：
//身高（cm） 必选
//胸围（cm） 必选
//体重（公斤）
//腰围（cm）
//肩宽（cm）
//袖长（cm）
//衣长（cm）
//背宽（cm）
//前长（cm）
//下摆围（cm）
//袖口（cm）
//袖肥（cm）
//领口（cm）
//
//内衣-睡裤/保暖裤类目，尺码维度为：
//身高（cm） 必选
//腰围（cm） 必选
//体重（公斤）
//臀围（cm）
//裆部（cm）
//裤长（cm）
//脚口（cm）
//腿围（cm）
//裤侧长（cm）
//
//内衣-睡裙类目，尺码维度为：
//身高（cm） 必选
//胸围（cm） 必选
//体重（公斤）
//裙长（cm）
//腰围（cm）
//袖长（cm）
//肩宽（cm）
//背宽（cm）
//腿围（cm）
//臀围（cm）
//底摆（cm）
//
//男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：
//身高（cm）
//体重（公斤）
//肩宽（cm）
//胸围（cm）
//腰围（cm）
//袖长（cm）
//衣长（cm）
//背宽（cm）
//前长（cm）
//摆围（cm）
//下摆围（cm）
//袖口（cm）
//袖肥（cm）
//中腰（cm）
//领深（cm）
//领高（cm）
//领宽（cm）
//领围（cm）
//圆摆后中长（cm）
//平摆衣长（cm）
//圆摆衣长（cm）
//
//男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：
//身高（cm）
//体重（公斤）
//腰围（cm）
//臀围（cm）
//裤长（cm）
//裙长（cm）
//裙摆长（cm）
//腿围（cm）
//膝围（cm）
//小脚围（cm）
//拉伸腰围（cm）
//坐围（cm）
//拉伸坐围（cm）
//脚口（cm）
//前浪（cm）
//后浪（cm）
//横档（cm）
//
//如果上述维度满足，可以自定义最多5个维度。
//
//模板格式为：
//尺码值:维度名称:数值
//如：M:身高（cm）:160,L:身高（cm）:170
type TmallItemSizemappingTemplateCreateRequest struct {
    
    // TemplateContent required尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
    TemplateContent string `json:"template_content";xml:"template_content"`
    
    // TemplateName required尺码表模板名称
    TemplateName string `json:"template_name";xml:"template_name"`
    
}

func (req *TmallItemSizemappingTemplateCreateRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.create"
}

func (req *TmallItemSizemappingTemplateCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["TemplateContent"] = req.TemplateContent
    
    params["TemplateName"] = req.TemplateName
    
    return params
}

// TmallItemSizemappingTemplateCreateResponse 新增天猫商品尺码表模板
//
//男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：
//脚长（cm） 必选
//
//内衣-文胸类目，尺码表维度为：
//上胸围（cm） 必选
//下胸围（cm） 必选
//上下胸围差（cm） 必选
//身高（cm）
//体重（公斤）
//
//内衣-内裤类目，尺码表维度为：
//腰围（cm） 必选
//臀围（cm） 必选
//身高（cm）
//体重（公斤）
//裤长（cm）
//裆部（cm）
//脚口（cm）
//腿围（cm）
//
//内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：
//身高（cm） 必选
//胸围（cm） 必选
//体重（公斤）
//腰围（cm）
//肩宽（cm）
//袖长（cm）
//衣长（cm）
//背宽（cm）
//前长（cm）
//下摆围（cm）
//袖口（cm）
//袖肥（cm）
//领口（cm）
//
//内衣-睡裤/保暖裤类目，尺码维度为：
//身高（cm） 必选
//腰围（cm） 必选
//体重（公斤）
//臀围（cm）
//裆部（cm）
//裤长（cm）
//脚口（cm）
//腿围（cm）
//裤侧长（cm）
//
//内衣-睡裙类目，尺码维度为：
//身高（cm） 必选
//胸围（cm） 必选
//体重（公斤）
//裙长（cm）
//腰围（cm）
//袖长（cm）
//肩宽（cm）
//背宽（cm）
//腿围（cm）
//臀围（cm）
//底摆（cm）
//
//男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：
//身高（cm）
//体重（公斤）
//肩宽（cm）
//胸围（cm）
//腰围（cm）
//袖长（cm）
//衣长（cm）
//背宽（cm）
//前长（cm）
//摆围（cm）
//下摆围（cm）
//袖口（cm）
//袖肥（cm）
//中腰（cm）
//领深（cm）
//领高（cm）
//领宽（cm）
//领围（cm）
//圆摆后中长（cm）
//平摆衣长（cm）
//圆摆衣长（cm）
//
//男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：
//身高（cm）
//体重（公斤）
//腰围（cm）
//臀围（cm）
//裤长（cm）
//裙长（cm）
//裙摆长（cm）
//腿围（cm）
//膝围（cm）
//小脚围（cm）
//拉伸腰围（cm）
//坐围（cm）
//拉伸坐围（cm）
//脚口（cm）
//前浪（cm）
//后浪（cm）
//横档（cm）
//
//如果上述维度满足，可以自定义最多5个维度。
//
//模板格式为：
//尺码值:维度名称:数值
//如：M:身高（cm）:160,L:身高（cm）:170
type TmallItemSizemappingTemplateCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SizeMappingTemplate Object尺码表模板
    SizeMappingTemplate SizeMappingTemplateDo `json:"size_mapping_template";xml:"size_mapping_template"`
    
}

// TaobaoOmniorderStoreGetconsignmailcodeRequest 用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
type TaobaoOmniorderStoreGetconsignmailcodeRequest struct {
    
    // Channel required淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS)
    Channel string `json:"channel";xml:"channel"`
    
    // Receiver optional收件人信息
    Receiver ReceiverDto `json:"receiver";xml:"receiver"`
    
    // SdtExtendInfoDTO optional扩展信息
    SdtExtendInfoDTO SdtExtendInfoDto `json:"sdt_extend_info_d_t_o";xml:"sdt_extend_info_d_t_o"`
    
    // SenderContact optional发件人联系电话，如空则表示使用门店信息中的电话号码
    SenderContact string `json:"sender_contact";xml:"sender_contact"`
    
    // StoreId required门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
    // Trades required订单信息，目前一次请求只支持一个主订单
    Trades TradeOrderInfoDto `json:"trades";xml:"trades"`
    
}

func (req *TaobaoOmniorderStoreGetconsignmailcodeRequest) GetAPIName() string {
	return "taobao.omniorder.store.getconsignmailcode"
}

func (req *TaobaoOmniorderStoreGetconsignmailcodeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Channel"] = req.Channel
    
    params["Receiver"] = req.Receiver
    
    params["SdtExtendInfoDTO"] = req.SdtExtendInfoDTO
    
    params["SenderContact"] = req.SenderContact
    
    params["StoreId"] = req.StoreId
    
    params["Trades"] = req.Trades
    
    return params
}

// TaobaoOmniorderStoreGetconsignmailcodeResponse 用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
type TaobaoOmniorderStoreGetconsignmailcodeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// AlibabaEinvoiceCreatereqRequest ERP发起开票请求
type AlibabaEinvoiceCreatereqRequest struct {
    
    // ApplyId optional开票申请ID，接收了开票申请消息后，需要把apply_id带上
    ApplyId string `json:"apply_id";xml:"apply_id"`
    
    // BusinessType required默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
    BusinessType int64 `json:"business_type";xml:"business_type"`
    
    // ErpTid optionalERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
    ErpTid string `json:"erp_tid";xml:"erp_tid"`
    
    // InvoiceAmount required开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    InvoiceAmount string `json:"invoice_amount";xml:"invoice_amount"`
    
    // InvoiceItems required电子发票明细
    InvoiceItems InvoiceItem `json:"invoice_items";xml:"invoice_items"`
    
    // InvoiceKind optional发票种类，0=电子发票,1=纸质发票,2=专票。注意：未订购纸票服务的税号无法开具纸票
    InvoiceKind int64 `json:"invoice_kind";xml:"invoice_kind"`
    
    // InvoiceMemo optional发票备注，有些省市会把此信息打印到PDF中
    InvoiceMemo string `json:"invoice_memo";xml:"invoice_memo"`
    
    // InvoiceTime required开票日期, 格式"YYYY-MM-DD HH:SS:MM"
    InvoiceTime time.Time `json:"invoice_time";xml:"invoice_time"`
    
    // InvoiceType required发票(开票)类型，蓝票blue,红票red，默认blue
    InvoiceType string `json:"invoice_type";xml:"invoice_type"`
    
    // NormalInvoiceCode optional原发票代码(开红票时传入)
    NormalInvoiceCode string `json:"normal_invoice_code";xml:"normal_invoice_code"`
    
    // NormalInvoiceNo optional原发票号码(开红票时传入)
    NormalInvoiceNo string `json:"normal_invoice_no";xml:"normal_invoice_no"`
    
    // OutShopName optional外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
    OutShopName string `json:"out_shop_name";xml:"out_shop_name"`
    
    // PayeeAddress required开票方地址(新版中为必传)
    PayeeAddress string `json:"payee_address";xml:"payee_address"`
    
    // PayeeBankaccount optional开票方银行及 帐号
    PayeeBankaccount string `json:"payee_bankaccount";xml:"payee_bankaccount"`
    
    // PayeeChecker optional复核人
    PayeeChecker string `json:"payee_checker";xml:"payee_checker"`
    
    // PayeeName required开票方名称，公司名(如:XX商城)
    PayeeName string `json:"payee_name";xml:"payee_name"`
    
    // PayeeOperator optional开票人
    PayeeOperator string `json:"payee_operator";xml:"payee_operator"`
    
    // PayeePhone optional收款方电话
    PayeePhone string `json:"payee_phone";xml:"payee_phone"`
    
    // PayeeReceiver optional收款人
    PayeeReceiver string `json:"payee_receiver";xml:"payee_receiver"`
    
    // PayeeRegisterNo required收款方税务登记证号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // PayerAddress optional消费者地址
    PayerAddress string `json:"payer_address";xml:"payer_address"`
    
    // PayerBankaccount optional付款方开票开户银行及账号
    PayerBankaccount string `json:"payer_bankaccount";xml:"payer_bankaccount"`
    
    // PayerEmail optional消费者电子邮箱
    PayerEmail string `json:"payer_email";xml:"payer_email"`
    
    // PayerName required付款方名称, 对应发票台头
    PayerName string `json:"payer_name";xml:"payer_name"`
    
    // PayerPhone optional消费者联系电话
    PayerPhone string `json:"payer_phone";xml:"payer_phone"`
    
    // PayerRegisterNo optional付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
    PayerRegisterNo string `json:"payer_register_no";xml:"payer_register_no"`
    
    // PlatformCode required电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    PlatformCode string `json:"platform_code";xml:"platform_code"`
    
    // PlatformTid required电商平台对应的主订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
    // RedNoticeNo optional红字通知单号，专票冲红时需要，商家跟税局申请
    RedNoticeNo string `json:"red_notice_no";xml:"red_notice_no"`
    
    // SerialNo required开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
    // SumPrice required合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    SumPrice string `json:"sum_price";xml:"sum_price"`
    
    // SumTax required合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    SumTax string `json:"sum_tax";xml:"sum_tax"`
    
}

func (req *AlibabaEinvoiceCreatereqRequest) GetAPIName() string {
	return "alibaba.einvoice.createreq"
}

func (req *AlibabaEinvoiceCreatereqRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 32)
    
    params["ApplyId"] = req.ApplyId
    
    params["BusinessType"] = req.BusinessType
    
    params["ErpTid"] = req.ErpTid
    
    params["InvoiceAmount"] = req.InvoiceAmount
    
    params["InvoiceItems"] = req.InvoiceItems
    
    params["InvoiceKind"] = req.InvoiceKind
    
    params["InvoiceMemo"] = req.InvoiceMemo
    
    params["InvoiceTime"] = req.InvoiceTime
    
    params["InvoiceType"] = req.InvoiceType
    
    params["NormalInvoiceCode"] = req.NormalInvoiceCode
    
    params["NormalInvoiceNo"] = req.NormalInvoiceNo
    
    params["OutShopName"] = req.OutShopName
    
    params["PayeeAddress"] = req.PayeeAddress
    
    params["PayeeBankaccount"] = req.PayeeBankaccount
    
    params["PayeeChecker"] = req.PayeeChecker
    
    params["PayeeName"] = req.PayeeName
    
    params["PayeeOperator"] = req.PayeeOperator
    
    params["PayeePhone"] = req.PayeePhone
    
    params["PayeeReceiver"] = req.PayeeReceiver
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["PayerAddress"] = req.PayerAddress
    
    params["PayerBankaccount"] = req.PayerBankaccount
    
    params["PayerEmail"] = req.PayerEmail
    
    params["PayerName"] = req.PayerName
    
    params["PayerPhone"] = req.PayerPhone
    
    params["PayerRegisterNo"] = req.PayerRegisterNo
    
    params["PlatformCode"] = req.PlatformCode
    
    params["PlatformTid"] = req.PlatformTid
    
    params["RedNoticeNo"] = req.RedNoticeNo
    
    params["SerialNo"] = req.SerialNo
    
    params["SumPrice"] = req.SumPrice
    
    params["SumTax"] = req.SumTax
    
    return params
}

// AlibabaEinvoiceCreatereqResponse ERP发起开票请求
type AlibabaEinvoiceCreatereqResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic开票信息是否成功接受
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TmallItemSizemappingTemplateUpdateRequest 更新天猫商品尺码表模板
type TmallItemSizemappingTemplateUpdateRequest struct {
    
    // TemplateContent required尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
    TemplateContent string `json:"template_content";xml:"template_content"`
    
    // TemplateId required尺码表模板ID
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
    // TemplateName required尺码表模板名称
    TemplateName string `json:"template_name";xml:"template_name"`
    
}

func (req *TmallItemSizemappingTemplateUpdateRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.update"
}

func (req *TmallItemSizemappingTemplateUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["TemplateContent"] = req.TemplateContent
    
    params["TemplateId"] = req.TemplateId
    
    params["TemplateName"] = req.TemplateName
    
    return params
}

// TmallItemSizemappingTemplateUpdateResponse 更新天猫商品尺码表模板
type TmallItemSizemappingTemplateUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SizeMappingTemplate Object尺码表模板
    SizeMappingTemplate SizeMappingTemplateDo `json:"size_mapping_template";xml:"size_mapping_template"`
    
}

// TaobaoTradesSoldIncrementvGetRequest 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
//<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
//<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
//<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
//<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。
type TaobaoTradesSoldIncrementvGetRequest struct {
    
    // EndCreate required查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。
    EndCreate time.Time `json:"end_create";xml:"end_create"`
    
    // ExtType optional可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
    ExtType string `json:"ext_type";xml:"ext_type"`
    
    // Fields required需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
    Fields string `json:"fields";xml:"fields"`
    
    // PageNo optional页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span>
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartCreate required查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
    StartCreate time.Time `json:"start_create";xml:"start_create"`
    
    // Status optional交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
    Status string `json:"status";xml:"status"`
    
    // Tag optional卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
    Tag string `json:"tag";xml:"tag"`
    
    // Type optional交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。
    Type string `json:"type";xml:"type"`
    
    // UseHasNext optional是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。
    UseHasNext bool `json:"use_has_next";xml:"use_has_next"`
    
}

func (req *TaobaoTradesSoldIncrementvGetRequest) GetAPIName() string {
	return "taobao.trades.sold.incrementv.get"
}

func (req *TaobaoTradesSoldIncrementvGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["EndCreate"] = req.EndCreate
    
    params["ExtType"] = req.ExtType
    
    params["Fields"] = req.Fields
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartCreate"] = req.StartCreate
    
    params["Status"] = req.Status
    
    params["Tag"] = req.Tag
    
    params["Type"] = req.Type
    
    params["UseHasNext"] = req.UseHasNext
    
    return params
}

// TaobaoTradesSoldIncrementvGetResponse 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息） 
//<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。 
//<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。 
//<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。 
//<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。
type TaobaoTradesSoldIncrementvGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // TotalResults Basic搜索到的交易信息总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
    // Trades Object Array搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    Trades Trade `json:"trades";xml:"trades"`
    
}

// TaobaoOmniorderStoreWarehouseRequest 为了能够支持逆向到门店对应的区域仓，商家需要配置门店与区域仓的映射关系，这个接口可以给商家提供门店与区域仓映射关系的增删改查功能。
type TaobaoOmniorderStoreWarehouseRequest struct {
    
    // Operation required操作，1表示增加或者更新，2表示删除，3表示查询
    Operation int64 `json:"operation";xml:"operation"`
    
    // StoreId optional门店id
    StoreId int64 `json:"store_id";xml:"store_id"`
    
    // WarehouseAddress optional仓联系地址
    WarehouseAddress string `json:"warehouse_address";xml:"warehouse_address"`
    
    // WarehouseCode optional仓编码
    WarehouseCode string `json:"warehouse_code";xml:"warehouse_code"`
    
    // WarehouseContact optional仓联系人
    WarehouseContact string `json:"warehouse_contact";xml:"warehouse_contact"`
    
    // WarehouseMobile optional仓联系电话
    WarehouseMobile string `json:"warehouse_mobile";xml:"warehouse_mobile"`
    
}

func (req *TaobaoOmniorderStoreWarehouseRequest) GetAPIName() string {
	return "taobao.omniorder.store.warehouse"
}

func (req *TaobaoOmniorderStoreWarehouseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Operation"] = req.Operation
    
    params["StoreId"] = req.StoreId
    
    params["WarehouseAddress"] = req.WarehouseAddress
    
    params["WarehouseCode"] = req.WarehouseCode
    
    params["WarehouseContact"] = req.WarehouseContact
    
    params["WarehouseMobile"] = req.WarehouseMobile
    
    return params
}

// TaobaoOmniorderStoreWarehouseResponse 为了能够支持逆向到门店对应的区域仓，商家需要配置门店与区域仓的映射关系，这个接口可以给商家提供门店与区域仓映射关系的增删改查功能。
type TaobaoOmniorderStoreWarehouseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Data Basic成功增加或者更新一条门店与区域仓的关联
    Data string `json:"data";xml:"data"`
    
    // FailCode Basiccode
    FailCode string `json:"fail_code";xml:"fail_code"`
    
    // FailMessage Basicmessage
    FailMessage string `json:"fail_message";xml:"fail_message"`
    
}

// TmallItemSizemappingTemplateDeleteRequest 删除天猫商品尺码表模板
type TmallItemSizemappingTemplateDeleteRequest struct {
    
    // TemplateId required尺码表模板ID
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TmallItemSizemappingTemplateDeleteRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.delete"
}

func (req *TmallItemSizemappingTemplateDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TemplateId"] = req.TemplateId
    
    return params
}

// TmallItemSizemappingTemplateDeleteResponse 删除天猫商品尺码表模板
type TmallItemSizemappingTemplateDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TemplateId Basic尺码表模板ID
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
}

// TmallItemSizemappingTemplateGetRequest 获取天猫商品尺码表模板
type TmallItemSizemappingTemplateGetRequest struct {
    
    // TemplateId required尺码表模板ID
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
}

func (req *TmallItemSizemappingTemplateGetRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.get"
}

func (req *TmallItemSizemappingTemplateGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TemplateId"] = req.TemplateId
    
    return params
}

// TmallItemSizemappingTemplateGetResponse 获取天猫商品尺码表模板
type TmallItemSizemappingTemplateGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SizeMappingTemplate Object尺码表模板
    SizeMappingTemplate Model `json:"size_mapping_template";xml:"size_mapping_template"`
    
}

// TaobaoOmniorderStoreSdtconsignRequest ISV取完单号后通知菜鸟裹裹发货
type TaobaoOmniorderStoreSdtconsignRequest struct {
    
    // PackageId required取号接口返回的包裹id
    PackageId string `json:"package_id";xml:"package_id"`
    
    // TagCode optional发货标签号
    TagCode string `json:"tag_code";xml:"tag_code"`
    
}

func (req *TaobaoOmniorderStoreSdtconsignRequest) GetAPIName() string {
	return "taobao.omniorder.store.sdtconsign"
}

func (req *TaobaoOmniorderStoreSdtconsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["PackageId"] = req.PackageId
    
    params["TagCode"] = req.TagCode
    
    return params
}

// TaobaoOmniorderStoreSdtconsignResponse ISV取完单号后通知菜鸟裹裹发货
type TaobaoOmniorderStoreSdtconsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Data Objectdata
    Data SdtConsignResponse `json:"data";xml:"data"`
    
    // ErrCode Basic异常码 0 为正常，否则异常
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // Message Basic异常信息
    Message string `json:"message";xml:"message"`
    
}

// AlibabaEinvoiceCreateResultsIncrementGetRequest 增量开票结果获取
type AlibabaEinvoiceCreateResultsIncrementGetRequest struct {
    
    // EndModified required终止查询时间
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // PageNo optional显示的页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional页面大小(不能超过200)
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // PayeeRegisterNo required收款方税务登记证号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // StartModified required起始查询时间
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
    // Status optional开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
    Status string `json:"status";xml:"status"`
    
}

func (req *AlibabaEinvoiceCreateResultsIncrementGetRequest) GetAPIName() string {
	return "alibaba.einvoice.create.results.increment.get"
}

func (req *AlibabaEinvoiceCreateResultsIncrementGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["EndModified"] = req.EndModified
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["StartModified"] = req.StartModified
    
    params["Status"] = req.Status
    
    return params
}

// AlibabaEinvoiceCreateResultsIncrementGetResponse 增量开票结果获取
type AlibabaEinvoiceCreateResultsIncrementGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InvoiceResultList Object Array开票结果返回列表
    InvoiceResultList InvoiceResult `json:"invoice_result_list";xml:"invoice_result_list"`
    
    // TotalCount Basic符合条件的开票总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// AlibabaEinvoiceCreateResultGetRequest ERP开票结果获取
type AlibabaEinvoiceCreateResultGetRequest struct {
    
    // OutShopName optional外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
    OutShopName string `json:"out_shop_name";xml:"out_shop_name"`
    
    // PayeeRegisterNo required收款方税务登记证号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // PlatformCode optional电商平台代码。淘宝：taobao，天猫：tmall
    PlatformCode string `json:"platform_code";xml:"platform_code"`
    
    // PlatformTid optional电商平台对应的订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
    // SerialNo optional流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoiceCreateResultGetRequest) GetAPIName() string {
	return "alibaba.einvoice.create.result.get"
}

func (req *AlibabaEinvoiceCreateResultGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["OutShopName"] = req.OutShopName
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["PlatformCode"] = req.PlatformCode
    
    params["PlatformTid"] = req.PlatformTid
    
    params["SerialNo"] = req.SerialNo
    
    return params
}

// AlibabaEinvoiceCreateResultGetResponse ERP开票结果获取
type AlibabaEinvoiceCreateResultGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InvoiceResultList Object Array开票返回结果数据列表
    InvoiceResultList InvoiceResult `json:"invoice_result_list";xml:"invoice_result_list"`
    
}

// TaobaoKfcKeywordSearchRequest 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
type TaobaoKfcKeywordSearchRequest struct {
    
    // Apply optional应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。
    // 
    // 这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。
    // 
    // 
    // 通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。
    Apply string `json:"apply";xml:"apply"`
    
    // Content required需要过滤的文本信息
    Content string `json:"content";xml:"content"`
    
    // Nick optional发布信息的淘宝会员名，可以不传
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoKfcKeywordSearchRequest) GetAPIName() string {
	return "taobao.kfc.keyword.search"
}

func (req *TaobaoKfcKeywordSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Apply"] = req.Apply
    
    params["Content"] = req.Content
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoKfcKeywordSearchResponse 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
type TaobaoKfcKeywordSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // KfcSearchResult ObjectKFC 关键词过滤匹配结果
    KfcSearchResult KfcSearchResult `json:"kfc_search_result";xml:"kfc_search_result"`
    
}

// TaobaoQimenCombineitemQueryRequest 组合货品关系查询
type TaobaoQimenCombineitemQueryRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // ItemId optional奇门仓储字段
    ItemId string `json:"itemId";xml:"itemId"`
    
    // OwnerCode optional奇门仓储字段
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
}

func (req *TaobaoQimenCombineitemQueryRequest) GetAPIName() string {
	return "taobao.qimen.combineitem.query"
}

func (req *TaobaoQimenCombineitemQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["ItemId"] = req.ItemId
    
    params["OwnerCode"] = req.OwnerCode
    
    return params
}

// TaobaoQimenCombineitemQueryResponse 组合货品关系查询
type TaobaoQimenCombineitemQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // CombItems Object Array奇门仓储字段
    CombItems CombItem `json:"combItems";xml:"combItems"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TmallItemSizemappingTemplatesListRequest 获取所有尺码表模板列表。
type TmallItemSizemappingTemplatesListRequest struct {
    
}

func (req *TmallItemSizemappingTemplatesListRequest) GetAPIName() string {
	return "tmall.item.sizemapping.templates.list"
}

func (req *TmallItemSizemappingTemplatesListRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TmallItemSizemappingTemplatesListResponse 获取所有尺码表模板列表。
type TmallItemSizemappingTemplatesListResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SizeMappingTemplates Object Array尺码表模板列表
    SizeMappingTemplates SizeMappingTemplate `json:"size_mapping_templates";xml:"size_mapping_templates"`
    
}

// TaobaoTmcTopicGroupAddRequest 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
//如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
type TaobaoTmcTopicGroupAddRequest struct {
    
    // GroupName required消息分组名，如果不存在，会自动创建
    GroupName string `json:"group_name";xml:"group_name"`
    
    // Topics required消息topic名称，多个以逗号(,)分割
    Topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcTopicGroupAddRequest) GetAPIName() string {
	return "taobao.tmc.topic.group.add"
}

func (req *TaobaoTmcTopicGroupAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["GroupName"] = req.GroupName
    
    params["Topics"] = req.Topics
    
    return params
}

// TaobaoTmcTopicGroupAddResponse 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
//如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
type TaobaoTmcTopicGroupAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basictrue
    Result bool `json:"result";xml:"result"`
    
}

// TmallChannelTradeOrderCreateRequest 创建渠道分销单
type TmallChannelTradeOrderCreateRequest struct {
    
    // Param0 required入参
    Param0 TopChannelPurchaseOrderCreateParam `json:"param0";xml:"param0"`
    
}

func (req *TmallChannelTradeOrderCreateRequest) GetAPIName() string {
	return "tmall.channel.trade.order.create"
}

func (req *TmallChannelTradeOrderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param0"] = req.Param0
    
    return params
}

// TmallChannelTradeOrderCreateResponse 创建渠道分销单
type TmallChannelTradeOrderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // MainPurchaseOrderList Basic Array采购单号
    MainPurchaseOrderList map[string]interface{} `json:"main_purchase_order_list";xml:"main_purchase_order_list"`
    
}

// TaobaoFenxiaoDistributorsGetRequest 查询和当前登录供应商有合作关系的分销商的信息
type TaobaoFenxiaoDistributorsGetRequest struct {
    
    // Nicks required分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
    Nicks string `json:"nicks";xml:"nicks"`
    
}

func (req *TaobaoFenxiaoDistributorsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributors.get"
}

func (req *TaobaoFenxiaoDistributorsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Nicks"] = req.Nicks
    
    return params
}

// TaobaoFenxiaoDistributorsGetResponse 查询和当前登录供应商有合作关系的分销商的信息
type TaobaoFenxiaoDistributorsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Distributors Object Array分销商详细信息
    Distributors Distributor `json:"distributors";xml:"distributors"`
    
}

// QimenTaobaoAutoEntryorderGiftitemcancelRequest 该接口一次只能回传一个主交易号的BMS增加的货品取消信息
type QimenTaobaoAutoEntryorderGiftitemcancelRequest struct {
    
    // CustomerId requiredcustomerId
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // Request optional请求
    Request ErpBmsOrderGiftCancelRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoAutoEntryorderGiftitemcancelRequest) GetAPIName() string {
	return "qimen.taobao.auto.entryorder.giftitemcancel"
}

func (req *QimenTaobaoAutoEntryorderGiftitemcancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoAutoEntryorderGiftitemcancelResponse 该接口一次只能回传一个主交易号的BMS增加的货品取消信息
type QimenTaobaoAutoEntryorderGiftitemcancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Objectresponse
    Response Response `json:"response";xml:"response"`
    
}

// QimenTaobaoBmsErptradeInterceptRequest 调用ERP接口，拦截订单
type QimenTaobaoBmsErptradeInterceptRequest struct {
    
    // CustomerId required货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // Request optional请求主体
    Request BmsTaobaoOrderIntercepteRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsErptradeInterceptRequest) GetAPIName() string {
	return "qimen.taobao.bms.erptrade.intercept"
}

func (req *QimenTaobaoBmsErptradeInterceptRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoBmsErptradeInterceptResponse 调用ERP接口，拦截订单
type QimenTaobaoBmsErptradeInterceptResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response ObjectResponse
    Response Response `json:"response";xml:"response"`
    
}

// TmallStoredeliverAllocationAcceptRequest 商品通门店发货业务门店接单拒单接口
type TmallStoredeliverAllocationAcceptRequest struct {
    
    // AllocationCode required派单号
    AllocationCode string `json:"allocation_code";xml:"allocation_code"`
    
    // IsAccept required是否接单
    IsAccept bool `json:"is_accept";xml:"is_accept"`
    
    // SubOrderCode required子订单号必须和派单号匹配
    SubOrderCode int64 `json:"sub_order_code";xml:"sub_order_code"`
    
}

func (req *TmallStoredeliverAllocationAcceptRequest) GetAPIName() string {
	return "tmall.storedeliver.allocation.accept"
}

func (req *TmallStoredeliverAllocationAcceptRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["AllocationCode"] = req.AllocationCode
    
    params["IsAccept"] = req.IsAccept
    
    params["SubOrderCode"] = req.SubOrderCode
    
    return params
}

// TmallStoredeliverAllocationAcceptResponse 商品通门店发货业务门店接单拒单接口
type TmallStoredeliverAllocationAcceptResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否执行成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// QimenTaobaoBmsErptradeTransferconsignRequest BMS调用ERP订单菜鸟仓&商家仓互转接口
type QimenTaobaoBmsErptradeTransferconsignRequest struct {
    
    // CustomerId required货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // Request optional请求体
    Request BmsErptradeTransferConsignRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsErptradeTransferconsignRequest) GetAPIName() string {
	return "qimen.taobao.bms.erptrade.transferconsign"
}

func (req *QimenTaobaoBmsErptradeTransferconsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoBmsErptradeTransferconsignResponse BMS调用ERP订单菜鸟仓&商家仓互转接口
type QimenTaobaoBmsErptradeTransferconsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Object
    Response Response `json:"response";xml:"response"`
    
}

// QimenTaobaoBmsTradeConsignRequest BMS通知ERP交易单整单出库接口
type QimenTaobaoBmsTradeConsignRequest struct {
    
    // CustomerId required货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // Request optional请求实体
    Request BmsTradeConsignRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoBmsTradeConsignRequest) GetAPIName() string {
	return "qimen.taobao.bms.trade.consign"
}

func (req *QimenTaobaoBmsTradeConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoBmsTradeConsignResponse BMS通知ERP交易单整单出库接口
type QimenTaobaoBmsTradeConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Object
    Response Response `json:"response";xml:"response"`
    
}

// TaobaoLogisticsExpressModifyAppointRequest 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaoLogisticsExpressModifyAppointRequest struct {
    
    // ExpressModifyAppointTopRequest required改约请求对象
    ExpressModifyAppointTopRequest ExpressModifyAppointTopRequestDto `json:"express_modify_appoint_top_request";xml:"express_modify_appoint_top_request"`
    
}

func (req *TaobaoLogisticsExpressModifyAppointRequest) GetAPIName() string {
	return "taobao.logistics.express.modify.appoint"
}

func (req *TaobaoLogisticsExpressModifyAppointRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ExpressModifyAppointTopRequest"] = req.ExpressModifyAppointTopRequest
    
    return params
}

// TaobaoLogisticsExpressModifyAppointResponse 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaoLogisticsExpressModifyAppointResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object调用结果
    Result SingleResultDto `json:"result";xml:"result"`
    
}

// TaobaoLogisticsConsignResendRequest 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：<br>
//1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
//2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。
type TaobaoLogisticsConsignResendRequest struct {
    
    // CompanyCode required物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取。<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font>
    CompanyCode string `json:"company_code";xml:"company_code"`
    
    // Feature optionalfeature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。
    Feature string `json:"feature";xml:"feature"`
    
    // IsSplit optional表明是否是拆单，默认值0，1表示拆单
    IsSplit int64 `json:"is_split";xml:"is_split"`
    
    // OutSid required运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    OutSid string `json:"out_sid";xml:"out_sid"`
    
    // SellerIp optional商家的IP地址
    SellerIp string `json:"seller_ip";xml:"seller_ip"`
    
    // SubTid optional拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！
    SubTid int64 `json:"sub_tid";xml:"sub_tid"`
    
    // Tid required淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsConsignResendRequest) GetAPIName() string {
	return "taobao.logistics.consign.resend"
}

func (req *TaobaoLogisticsConsignResendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["CompanyCode"] = req.CompanyCode
    
    params["Feature"] = req.Feature
    
    params["IsSplit"] = req.IsSplit
    
    params["OutSid"] = req.OutSid
    
    params["SellerIp"] = req.SellerIp
    
    params["SubTid"] = req.SubTid
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsConsignResendResponse 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：<br>
//1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
//2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。
type TaobaoLogisticsConsignResendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shipping Object返回发货是否成功is_success
    Shipping Shipping `json:"shipping";xml:"shipping"`
    
}

// TaobaoRdcAligeniusOrderEventUpdateRequest 供erp回传订单变更状态事件
type TaobaoRdcAligeniusOrderEventUpdateRequest struct {
    
    // Param0 optional参数
    Param0 OrderEventDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusOrderEventUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.order.event.update"
}

func (req *TaobaoRdcAligeniusOrderEventUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param0"] = req.Param0
    
    return params
}

// TaobaoRdcAligeniusOrderEventUpdateResponse 供erp回传订单变更状态事件
type TaobaoRdcAligeniusOrderEventUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // FailCode Basic错误码
    FailCode string `json:"fail_code";xml:"fail_code"`
    
    // FailInfo Basic错误描述
    FailInfo string `json:"fail_info";xml:"fail_info"`
    
    // SuccessFlag Basic是否成功
    SuccessFlag bool `json:"success_flag";xml:"success_flag"`
    
}

// TaobaoTmcTopicGroupDeleteRequest 删除根据topic名称路由消息到不同的分组关系
type TaobaoTmcTopicGroupDeleteRequest struct {
    
    // GroupId optional消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
    GroupId int64 `json:"group_id";xml:"group_id"`
    
    // GroupName required消息分组名
    GroupName string `json:"group_name";xml:"group_name"`
    
    // Topics required消息topic名称，多个以逗号(,)分割
    Topics string `json:"topics";xml:"topics"`
    
}

func (req *TaobaoTmcTopicGroupDeleteRequest) GetAPIName() string {
	return "taobao.tmc.topic.group.delete"
}

func (req *TaobaoTmcTopicGroupDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["GroupId"] = req.GroupId
    
    params["GroupName"] = req.GroupName
    
    params["Topics"] = req.Topics
    
    return params
}

// TaobaoTmcTopicGroupDeleteResponse 删除根据topic名称路由消息到不同的分组关系
type TaobaoTmcTopicGroupDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basictrue
    Result bool `json:"result";xml:"result"`
    
}

// TmallItemSchemaUpdateRequest 天猫根据规则编辑商品
type TmallItemSchemaUpdateRequest struct {
    
    // CategoryId optional商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // ItemId required需要编辑的商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // ProductId optional商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // XmlData required根据tmall.item.update.schema.get生成的商品编辑规则入参数据
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaUpdateRequest) GetAPIName() string {
	return "tmall.item.schema.update"
}

func (req *TmallItemSchemaUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["CategoryId"] = req.CategoryId
    
    params["ItemId"] = req.ItemId
    
    params["ProductId"] = req.ProductId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TmallItemSchemaUpdateResponse 天猫根据规则编辑商品
type TmallItemSchemaUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic商品更新操作成功时间
    GmtModified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    // UpdateItemResult Basic返回商品发布结果
    UpdateItemResult string `json:"update_item_result";xml:"update_item_result"`
    
}

// CainiaoSmartdeliveryIssubscribeIQueryRequest 查询商家时候订购智能发货引擎服务
type CainiaoSmartdeliveryIssubscribeIQueryRequest struct {
    
}

func (req *CainiaoSmartdeliveryIssubscribeIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.issubscribe.i.query"
}

func (req *CainiaoSmartdeliveryIssubscribeIQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// CainiaoSmartdeliveryIssubscribeIQueryResponse 查询商家时候订购智能发货引擎服务
type CainiaoSmartdeliveryIssubscribeIQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Successful Basictrue:商家已订购智能发货引擎服务,false:商家还没有订购或订购已过期
    Successful bool `json:"successful";xml:"successful"`
    
}

// TmallItemUpdateSchemaGetRequest Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetRequest struct {
    
    // CategoryId optional商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // ItemId required需要编辑的商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // ProductId optional商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
    ProductId int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallItemUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.item.update.schema.get"
}

func (req *TmallItemUpdateSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CategoryId"] = req.CategoryId
    
    params["ItemId"] = req.ItemId
    
    params["ProductId"] = req.ProductId
    
    return params
}

// TmallItemUpdateSchemaGetResponse Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateItemResult Basic返回发布商品的规则文档
    UpdateItemResult string `json:"update_item_result";xml:"update_item_result"`
    
}

// TmallProductUpdateSchemaGetRequest 获取用户更新产品的规则
type TmallProductUpdateSchemaGetRequest struct {
    
    // ProductId required产品编号
    ProductId int64 `json:"product_id";xml:"product_id"`
    
}

func (req *TmallProductUpdateSchemaGetRequest) GetAPIName() string {
	return "tmall.product.update.schema.get"
}

func (req *TmallProductUpdateSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ProductId"] = req.ProductId
    
    return params
}

// TmallProductUpdateSchemaGetResponse 获取用户更新产品的规则
type TmallProductUpdateSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateProductSchema Basic参数产品ID对产品的更新规则
    UpdateProductSchema string `json:"update_product_schema";xml:"update_product_schema"`
    
}

// TmallProductSchemaUpdateRequest 产品更新接口
type TmallProductSchemaUpdateRequest struct {
    
    // ProductId required产品编号
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // XmlData required根据tmall.product.update.schema.get生成的产品更新规则入参数据
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallProductSchemaUpdateRequest) GetAPIName() string {
	return "tmall.product.schema.update"
}

func (req *TmallProductSchemaUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ProductId"] = req.ProductId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TmallProductSchemaUpdateResponse 产品更新接口
type TmallProductSchemaUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateProductResult Basic产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间
    UpdateProductResult string `json:"update_product_result";xml:"update_product_result"`
    
}

// TmallExchangeRefusereasonGetRequest 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // DisputeType optional换货申请类型：0-任意类型；1-售中；2-售后
    DisputeType int64 `json:"dispute_type";xml:"dispute_type"`
    
    // Fields required返回字段
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeRefusereasonGetRequest) GetAPIName() string {
	return "tmall.exchange.refusereason.get"
}

func (req *TmallExchangeRefusereasonGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["DisputeId"] = req.DisputeId
    
    params["DisputeType"] = req.DisputeType
    
    params["Fields"] = req.Fields
    
    return params
}

// TmallExchangeRefusereasonGetResponse 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ResultSet `json:"result";xml:"result"`
    
}

// TaobaoTaeBillGetRequest 查询单笔账单明细
type TaobaoTaeBillGetRequest struct {
    
    // AccountId required虚拟账户科目编号
    AccountId int64 `json:"account_id";xml:"account_id"`
    
    // Bid optional账单编号
    Bid int64 `json:"bid";xml:"bid"`
    
    // Fields required传入需要返回的字段
    Fields string `json:"fields";xml:"fields"`
    
    // Id required账单编号
    Id int64 `json:"id";xml:"id"`
    
}

func (req *TaobaoTaeBillGetRequest) GetAPIName() string {
	return "taobao.tae.bill.get"
}

func (req *TaobaoTaeBillGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["AccountId"] = req.AccountId
    
    params["Bid"] = req.Bid
    
    params["Fields"] = req.Fields
    
    params["Id"] = req.Id
    
    return params
}

// TaobaoTaeBillGetResponse 查询单笔账单明细
type TaobaoTaeBillGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Bill Object账单明细
    Bill BillDto `json:"bill";xml:"bill"`
    
}

// TaobaoOmniorderGuideDataGetRequest 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
type TaobaoOmniorderGuideDataGetRequest struct {
    
    // PageNo required页码，从1开始
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize required每页数量，不能大于1000
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime required拉取数据开始时间
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
    // Type requireddetail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购脚骨明细;detail_sxg_order: 随心购订单明细
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoOmniorderGuideDataGetRequest) GetAPIName() string {
	return "taobao.omniorder.guide.data.get"
}

func (req *TaobaoOmniorderGuideDataGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoOmniorderGuideDataGetResponse 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
type TaobaoOmniorderGuideDataGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DataList Basic Array拉取的数据数组，如果为空，表示数据拉取完毕。拉取的数据字段包括打点时间、商家id、商品id和门店id等，传入的类型不同，返回的字段有所不同，可以根据具体类型的返回结果具体处理
    DataList map[string]interface{} `json:"data_list";xml:"data_list"`
    
}

// TaobaoTaeAccountsGetRequest tae查询费用科目信息
type TaobaoTaeAccountsGetRequest struct {
    
    // Aids optional需要获取的科目ID
    Aids int64 `json:"aids";xml:"aids"`
    
    // Fields required需要返回的字段
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoTaeAccountsGetRequest) GetAPIName() string {
	return "taobao.tae.accounts.get"
}

func (req *TaobaoTaeAccountsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Aids"] = req.Aids
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoTaeAccountsGetResponse tae查询费用科目信息
type TaobaoTaeAccountsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Accounts Object Array返回的科目信息
    Accounts TopAccountDto `json:"accounts";xml:"accounts"`
    
    // TotalResults Basic返回记录行数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoTaeBillsGetRequest tae查询账单明细
type TaobaoTaeBillsGetRequest struct {
    
    // CurrentPage optional页数,建议不要超过100页,越大性能越低,有可能会超时
    CurrentPage int64 `json:"current_page";xml:"current_page"`
    
    // Fields required传入需要返回的字段,参见Bill结构体
    Fields string `json:"fields";xml:"fields"`
    
    // ItemId optional科目编号
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // PTradeId optional交易编号
    PTradeId int64 `json:"p_trade_id";xml:"p_trade_id"`
    
    // PageSize optional每页大小,默认40条,可选范围 ：40~100
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // QueryDateType optional查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
    QueryDateType int64 `json:"query_date_type";xml:"query_date_type"`
    
    // QueryEndDate required结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
    QueryEndDate time.Time `json:"query_end_date";xml:"query_end_date"`
    
    // QueryStartDate required开始时间
    QueryStartDate time.Time `json:"query_start_date";xml:"query_start_date"`
    
    // TradeId optional子订单编号
    TradeId int64 `json:"trade_id";xml:"trade_id"`
    
}

func (req *TaobaoTaeBillsGetRequest) GetAPIName() string {
	return "taobao.tae.bills.get"
}

func (req *TaobaoTaeBillsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["CurrentPage"] = req.CurrentPage
    
    params["Fields"] = req.Fields
    
    params["ItemId"] = req.ItemId
    
    params["PTradeId"] = req.PTradeId
    
    params["PageSize"] = req.PageSize
    
    params["QueryDateType"] = req.QueryDateType
    
    params["QueryEndDate"] = req.QueryEndDate
    
    params["QueryStartDate"] = req.QueryStartDate
    
    params["TradeId"] = req.TradeId
    
    return params
}

// TaobaoTaeBillsGetResponse tae查询账单明细
type TaobaoTaeBillsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Bills Object Array账单列表
    Bills BillDto `json:"bills";xml:"bills"`
    
    // HasNext Basic是否存在下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // TotalResults Basic当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoOmniorderItemTagOperateRequest 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
type TaobaoOmniorderItemTagOperateRequest struct {
    
    // ItemId required商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // OmniSetting optional分单&接单设置
    OmniSetting OmniSettingDto `json:"omni_setting";xml:"omni_setting"`
    
    // Status required操作状态， 填 1 代表打标，填 -1 代表去标
    Status int64 `json:"status";xml:"status"`
    
    // Types required商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
    Types string `json:"types";xml:"types"`
    
}

func (req *TaobaoOmniorderItemTagOperateRequest) GetAPIName() string {
	return "taobao.omniorder.item.tag.operate"
}

func (req *TaobaoOmniorderItemTagOperateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ItemId"] = req.ItemId
    
    params["OmniSetting"] = req.OmniSetting
    
    params["Status"] = req.Status
    
    params["Types"] = req.Types
    
    return params
}

// TaobaoOmniorderItemTagOperateResponse 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
type TaobaoOmniorderItemTagOperateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic0 正常，否则异常
    Code string `json:"code";xml:"code"`
    
    // Message Basiccode 不为 0时有值，代表异常信息
    Message string `json:"message";xml:"message"`
    
}

// TmallExchangeAgreeRequest 卖家同意换货申请
type TmallExchangeAgreeRequest struct {
    
    // AddressId required收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
    AddressId int64 `json:"address_id";xml:"address_id"`
    
    // CompleteAddress optional详细收货地址
    CompleteAddress string `json:"complete_address";xml:"complete_address"`
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
    Fields string `json:"fields";xml:"fields"`
    
    // LeaveMessage optional卖家留言
    LeaveMessage string `json:"leave_message";xml:"leave_message"`
    
    // LeaveMessagePics optional上传图片举证
    LeaveMessagePics []byte `json:"leave_message_pics";xml:"leave_message_pics"`
    
    // Mobile optional收货人手机号
    Mobile string `json:"mobile";xml:"mobile"`
    
    // Post optional邮政编码
    Post string `json:"post";xml:"post"`
    
}

func (req *TmallExchangeAgreeRequest) GetAPIName() string {
	return "tmall.exchange.agree"
}

func (req *TmallExchangeAgreeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["AddressId"] = req.AddressId
    
    params["CompleteAddress"] = req.CompleteAddress
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    params["LeaveMessage"] = req.LeaveMessage
    
    params["LeaveMessagePics"] = req.LeaveMessagePics
    
    params["Mobile"] = req.Mobile
    
    params["Post"] = req.Post
    
    return params
}

// TmallExchangeAgreeResponse 卖家同意换货申请
type TmallExchangeAgreeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

// TmallExchangeRefuseRequest 卖家拒绝换货申请
type TmallExchangeRefuseRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回字段。目前支持dispute_id, bizorder_id, modified, status
    Fields string `json:"fields";xml:"fields"`
    
    // LeaveMessage optional拒绝换货申请时的留言
    LeaveMessage string `json:"leave_message";xml:"leave_message"`
    
    // LeaveMessagePics optional凭证图片
    LeaveMessagePics []byte `json:"leave_message_pics";xml:"leave_message_pics"`
    
    // SellerRefuseReasonId required换货原因对应ID
    SellerRefuseReasonId int64 `json:"seller_refuse_reason_id";xml:"seller_refuse_reason_id"`
    
}

func (req *TmallExchangeRefuseRequest) GetAPIName() string {
	return "tmall.exchange.refuse"
}

func (req *TmallExchangeRefuseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    params["LeaveMessage"] = req.LeaveMessage
    
    params["LeaveMessagePics"] = req.LeaveMessagePics
    
    params["SellerRefuseReasonId"] = req.SellerRefuseReasonId
    
    return params
}

// TmallExchangeRefuseResponse 卖家拒绝换货申请
type TmallExchangeRefuseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

// TmallExchangeMessageAddRequest 卖家创建换货留言
type TmallExchangeMessageAddRequest struct {
    
    // Content required留言内容
    Content string `json:"content";xml:"content"`
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
    Fields string `json:"fields";xml:"fields"`
    
    // MessagePics optional凭证图片列表
    MessagePics []byte `json:"message_pics";xml:"message_pics"`
    
}

func (req *TmallExchangeMessageAddRequest) GetAPIName() string {
	return "tmall.exchange.message.add"
}

func (req *TmallExchangeMessageAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Content"] = req.Content
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    params["MessagePics"] = req.MessagePics
    
    return params
}

// TmallExchangeMessageAddResponse 卖家创建换货留言
type TmallExchangeMessageAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ResultSet `json:"result";xml:"result"`
    
}

// TmallExchangeMessagesGetRequest 查询换货订单留言列表
type TmallExchangeMessagesGetRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
    Fields string `json:"fields";xml:"fields"`
    
    // OperatorRoles optional留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6)
    OperatorRoles int64 `json:"operator_roles";xml:"operator_roles"`
    
    // PageNo optional页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TmallExchangeMessagesGetRequest) GetAPIName() string {
	return "tmall.exchange.messages.get"
}

func (req *TmallExchangeMessagesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    params["OperatorRoles"] = req.OperatorRoles
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    return params
}

// TmallExchangeMessagesGetResponse 查询换货订单留言列表
type TmallExchangeMessagesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result RefundMessageResult `json:"result";xml:"result"`
    
}

// TmallExchangeReturngoodsRefuseRequest 卖家拒绝买家换货申请
type TmallExchangeReturngoodsRefuseRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // LeaveMessage optional留言说明
    LeaveMessage string `json:"leave_message";xml:"leave_message"`
    
    // LeaveMessagePics optional凭证图片
    LeaveMessagePics []byte `json:"leave_message_pics";xml:"leave_message_pics"`
    
    // SellerRefuseReasonId required拒绝原因ID
    SellerRefuseReasonId int64 `json:"seller_refuse_reason_id";xml:"seller_refuse_reason_id"`
    
}

func (req *TmallExchangeReturngoodsRefuseRequest) GetAPIName() string {
	return "tmall.exchange.returngoods.refuse"
}

func (req *TmallExchangeReturngoodsRefuseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["DisputeId"] = req.DisputeId
    
    params["LeaveMessage"] = req.LeaveMessage
    
    params["LeaveMessagePics"] = req.LeaveMessagePics
    
    params["SellerRefuseReasonId"] = req.SellerRefuseReasonId
    
    return params
}

// TmallExchangeReturngoodsRefuseResponse 卖家拒绝买家换货申请
type TmallExchangeReturngoodsRefuseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

// TaobaoQimenWmsReturnapplyReportRequest 退货异常包裹单通知接口
type TaobaoQimenWmsReturnapplyReportRequest struct {
    
    // Request optional请求对象
    Request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenWmsReturnapplyReportRequest) GetAPIName() string {
	return "taobao.qimen.wms.returnapply.report"
}

func (req *TaobaoQimenWmsReturnapplyReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Request"] = req.Request
    
    return params
}

// TaobaoQimenWmsReturnapplyReportResponse 退货异常包裹单通知接口
type TaobaoQimenWmsReturnapplyReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Object响应对象
    Response Response `json:"response";xml:"response"`
    
}

// CainiaoSmartdeliverySellerStatusIQueryRequest 查询智能发货引擎商家状态信息
type CainiaoSmartdeliverySellerStatusIQueryRequest struct {
    
}

func (req *CainiaoSmartdeliverySellerStatusIQueryRequest) GetAPIName() string {
	return "cainiao.smartdelivery.seller.status.i.query"
}

func (req *CainiaoSmartdeliverySellerStatusIQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// CainiaoSmartdeliverySellerStatusIQueryResponse 查询智能发货引擎商家状态信息
type CainiaoSmartdeliverySellerStatusIQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SellerStatus Object商家状态
    SellerStatus SellerStatus `json:"seller_status";xml:"seller_status"`
    
}

// TmallExchangeConsigngoodsRequest 卖家发货
type TmallExchangeConsigngoodsRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回字段
    Fields string `json:"fields";xml:"fields"`
    
    // LogisticsCompanyName required卖家发货的快递公司
    LogisticsCompanyName string `json:"logistics_company_name";xml:"logistics_company_name"`
    
    // LogisticsNo required卖家发货的物流单号
    LogisticsNo string `json:"logistics_no";xml:"logistics_no"`
    
    // LogisticsType required卖家发货的物流类型，100表示平邮，200表示快递
    LogisticsType int64 `json:"logistics_type";xml:"logistics_type"`
    
}

func (req *TmallExchangeConsigngoodsRequest) GetAPIName() string {
	return "tmall.exchange.consigngoods"
}

func (req *TmallExchangeConsigngoodsRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    params["LogisticsCompanyName"] = req.LogisticsCompanyName
    
    params["LogisticsNo"] = req.LogisticsNo
    
    params["LogisticsType"] = req.LogisticsType
    
    return params
}

// TmallExchangeConsigngoodsResponse 卖家发货
type TmallExchangeConsigngoodsResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result RefundBaseResponse `json:"result";xml:"result"`
    
}

// TaobaoTmcAuthGetRequest TMC连接授权Token
type TaobaoTmcAuthGetRequest struct {
    
    // Group optionaltmc组名
    Group string `json:"group";xml:"group"`
    
}

func (req *TaobaoTmcAuthGetRequest) GetAPIName() string {
	return "taobao.tmc.auth.get"
}

func (req *TaobaoTmcAuthGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Group"] = req.Group
    
    return params
}

// TaobaoTmcAuthGetResponse TMC连接授权Token
type TaobaoTmcAuthGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basicresult
    Result string `json:"result";xml:"result"`
    
}

// TaobaoInventoryMerchantAdjustRequest 货品库存商家端调整 ，入库，出库，盘点
type TaobaoInventoryMerchantAdjustRequest struct {
    
    // InventoryCheck required调整库存对象
    InventoryCheck InventoryCheckDto `json:"inventory_check";xml:"inventory_check"`
    
}

func (req *TaobaoInventoryMerchantAdjustRequest) GetAPIName() string {
	return "taobao.inventory.merchant.adjust"
}

func (req *TaobaoInventoryMerchantAdjustRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["InventoryCheck"] = req.InventoryCheck
    
    return params
}

// TaobaoInventoryMerchantAdjustResponse 货品库存商家端调整 ，入库，出库，盘点
type TaobaoInventoryMerchantAdjustResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResult `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsSkuCreateRequest 商品同步
type TaobaoWlbWmsSkuCreateRequest struct {
    
    // AdventLifecycle optional保质期预警天数
    AdventLifecycle int64 `json:"advent_lifecycle";xml:"advent_lifecycle"`
    
    // ApprovalNumber optional批准文号
    ApprovalNumber string `json:"approval_number";xml:"approval_number"`
    
    // BarCode optional条形码，多条码请用”；”分隔；
    BarCode string `json:"bar_code";xml:"bar_code"`
    
    // Brand optional品牌编码
    Brand string `json:"brand";xml:"brand"`
    
    // BrandName optional品牌名称
    BrandName string `json:"brand_name";xml:"brand_name"`
    
    // Category optional商品类别编码（外部系统类别）
    Category string `json:"category";xml:"category"`
    
    // CategoryName optional商品类别名称
    CategoryName string `json:"category_name";xml:"category_name"`
    
    // Color optional颜色
    Color string `json:"color";xml:"color"`
    
    // CostPrice optional成本价，单位分
    CostPrice int64 `json:"cost_price";xml:"cost_price"`
    
    // ExtendFields optional拓展属性
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // GrossWeight optional毛重，单位克
    GrossWeight int64 `json:"gross_weight";xml:"gross_weight"`
    
    // Height optional高度，单位毫米
    Height int64 `json:"height";xml:"height"`
    
    // IsAreaSale optional是否区域销售
    IsAreaSale bool `json:"is_area_sale";xml:"is_area_sale"`
    
    // IsBatchMgt optional是否启用批次管理
    IsBatchMgt bool `json:"is_batch_mgt";xml:"is_batch_mgt"`
    
    // IsDanger optional是否危险品
    IsDanger bool `json:"is_danger";xml:"is_danger"`
    
    // IsHygroscopic optional是否易碎品
    IsHygroscopic bool `json:"is_hygroscopic";xml:"is_hygroscopic"`
    
    // IsShelflife optional是否启用保质期管理
    IsShelflife bool `json:"is_shelflife";xml:"is_shelflife"`
    
    // IsSnMgt optional是否启用序列号管理
    IsSnMgt bool `json:"is_sn_mgt";xml:"is_sn_mgt"`
    
    // ItemCode required商家商品编码
    ItemCode string `json:"item_code";xml:"item_code"`
    
    // ItemId optional商家商品ID
    ItemId string `json:"item_id";xml:"item_id"`
    
    // ItemPrice optional零售价，单位分
    ItemPrice int64 `json:"item_price";xml:"item_price"`
    
    // Length optional长度，单位毫米
    Length int64 `json:"length";xml:"length"`
    
    // Lifecycle optional商品保质期天数
    Lifecycle int64 `json:"lifecycle";xml:"lifecycle"`
    
    // LockupLifecycle optional保质期禁售天数
    LockupLifecycle int64 `json:"lockup_lifecycle";xml:"lockup_lifecycle"`
    
    // Name required商品名称
    Name string `json:"name";xml:"name"`
    
    // NetWeight optional净重，单位克
    NetWeight int64 `json:"net_weight";xml:"net_weight"`
    
    // OriginAddress optional产地
    OriginAddress int64 `json:"origin_address";xml:"origin_address"`
    
    // Pcs optional箱规
    Pcs int64 `json:"pcs";xml:"pcs"`
    
    // RejectLifecycle optional保质期禁收天数
    RejectLifecycle int64 `json:"reject_lifecycle";xml:"reject_lifecycle"`
    
    // Size optional尺码
    Size string `json:"size";xml:"size"`
    
    // Specification optional规格
    Specification string `json:"specification";xml:"specification"`
    
    // StoreCode optional仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TagPrice optional吊牌价，单位分
    TagPrice int64 `json:"tag_price";xml:"tag_price"`
    
    // Title optional商品标题
    Title string `json:"title";xml:"title"`
    
    // Type required商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
    Type string `json:"type";xml:"type"`
    
    // UseYn required启用标识
    UseYn bool `json:"use_yn";xml:"use_yn"`
    
    // Volume optional体积，单位立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Width optional宽度，单位毫米
    Width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbWmsSkuCreateRequest) GetAPIName() string {
	return "taobao.wlb.wms.sku.create"
}

func (req *TaobaoWlbWmsSkuCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 38)
    
    params["AdventLifecycle"] = req.AdventLifecycle
    
    params["ApprovalNumber"] = req.ApprovalNumber
    
    params["BarCode"] = req.BarCode
    
    params["Brand"] = req.Brand
    
    params["BrandName"] = req.BrandName
    
    params["Category"] = req.Category
    
    params["CategoryName"] = req.CategoryName
    
    params["Color"] = req.Color
    
    params["CostPrice"] = req.CostPrice
    
    params["ExtendFields"] = req.ExtendFields
    
    params["GrossWeight"] = req.GrossWeight
    
    params["Height"] = req.Height
    
    params["IsAreaSale"] = req.IsAreaSale
    
    params["IsBatchMgt"] = req.IsBatchMgt
    
    params["IsDanger"] = req.IsDanger
    
    params["IsHygroscopic"] = req.IsHygroscopic
    
    params["IsShelflife"] = req.IsShelflife
    
    params["IsSnMgt"] = req.IsSnMgt
    
    params["ItemCode"] = req.ItemCode
    
    params["ItemId"] = req.ItemId
    
    params["ItemPrice"] = req.ItemPrice
    
    params["Length"] = req.Length
    
    params["Lifecycle"] = req.Lifecycle
    
    params["LockupLifecycle"] = req.LockupLifecycle
    
    params["Name"] = req.Name
    
    params["NetWeight"] = req.NetWeight
    
    params["OriginAddress"] = req.OriginAddress
    
    params["Pcs"] = req.Pcs
    
    params["RejectLifecycle"] = req.RejectLifecycle
    
    params["Size"] = req.Size
    
    params["Specification"] = req.Specification
    
    params["StoreCode"] = req.StoreCode
    
    params["TagPrice"] = req.TagPrice
    
    params["Title"] = req.Title
    
    params["Type"] = req.Type
    
    params["UseYn"] = req.UseYn
    
    params["Volume"] = req.Volume
    
    params["Width"] = req.Width
    
    return params
}

// TaobaoWlbWmsSkuCreateResponse 商品同步
type TaobaoWlbWmsSkuCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemId Basic系统自动生成
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // WlErrorCode Basic错误码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoWlbWmsSkuUpdateRequest 商品信息的更新
type TaobaoWlbWmsSkuUpdateRequest struct {
    
    // AdventLifecycle optional保质期预警天数
    AdventLifecycle int64 `json:"advent_lifecycle";xml:"advent_lifecycle"`
    
    // ApprovalNumber optional批准文号
    ApprovalNumber string `json:"approval_number";xml:"approval_number"`
    
    // Attribute optional商品属性编码
    Attribute string `json:"attribute";xml:"attribute"`
    
    // BarCode optional条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
    BarCode string `json:"bar_code";xml:"bar_code"`
    
    // Brand optional品牌编码
    Brand string `json:"brand";xml:"brand"`
    
    // BrandName optional品牌名称
    BrandName string `json:"brand_name";xml:"brand_name"`
    
    // Category optional商品类别编码（外部系统类别）
    Category string `json:"category";xml:"category"`
    
    // CategoryName optional商品类别名称
    CategoryName string `json:"category_name";xml:"category_name"`
    
    // Color optional颜色
    Color string `json:"color";xml:"color"`
    
    // CostPrice optional成本价，单位分
    CostPrice int64 `json:"cost_price";xml:"cost_price"`
    
    // ExtendFields optional拓展属性
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // GrossWeight optional毛重，单位克
    GrossWeight int64 `json:"gross_weight";xml:"gross_weight"`
    
    // Height optional高度，单位毫米
    Height int64 `json:"height";xml:"height"`
    
    // IsAreaSale optional是否区域销售属性
    IsAreaSale bool `json:"is_area_sale";xml:"is_area_sale"`
    
    // IsBatchMgt optional是否启用批次管理
    IsBatchMgt bool `json:"is_batch_mgt";xml:"is_batch_mgt"`
    
    // IsDanger optional是否危险品
    IsDanger bool `json:"is_danger";xml:"is_danger"`
    
    // IsHygroscopic optional是否易碎品
    IsHygroscopic bool `json:"is_hygroscopic";xml:"is_hygroscopic"`
    
    // IsShelflife optional是否启用保质期管理
    IsShelflife bool `json:"is_shelflife";xml:"is_shelflife"`
    
    // IsSnMgt optional是否启用序列号管理
    IsSnMgt bool `json:"is_sn_mgt";xml:"is_sn_mgt"`
    
    // ItemId required外部系统ID
    ItemId string `json:"item_id";xml:"item_id"`
    
    // ItemPrice optional零售价，单位分
    ItemPrice int64 `json:"item_price";xml:"item_price"`
    
    // Length optional长度，单位毫米
    Length int64 `json:"length";xml:"length"`
    
    // Lifecycle optional商品保质期天数
    Lifecycle int64 `json:"lifecycle";xml:"lifecycle"`
    
    // LockupLifecycle optional保质期禁售天数
    LockupLifecycle int64 `json:"lockup_lifecycle";xml:"lockup_lifecycle"`
    
    // Name optional商品名称
    Name string `json:"name";xml:"name"`
    
    // NetWeight optional净重，单位克
    NetWeight int64 `json:"net_weight";xml:"net_weight"`
    
    // OriginAddress optional产地
    OriginAddress int64 `json:"origin_address";xml:"origin_address"`
    
    // Pcs optional箱规
    Pcs int64 `json:"pcs";xml:"pcs"`
    
    // RejectLifecycle optional保质期禁收天数
    RejectLifecycle int64 `json:"reject_lifecycle";xml:"reject_lifecycle"`
    
    // Size optional尺码
    Size string `json:"size";xml:"size"`
    
    // Specification optional规格
    Specification string `json:"specification";xml:"specification"`
    
    // StoreCode optional仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TagPrice optional吊牌价，单位分
    TagPrice int64 `json:"tag_price";xml:"tag_price"`
    
    // Title optional商品标题
    Title string `json:"title";xml:"title"`
    
    // Type optional商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
    Type string `json:"type";xml:"type"`
    
    // UseYn required启用标识
    UseYn bool `json:"use_yn";xml:"use_yn"`
    
    // Volume optional体积，单位立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Width optional宽度，单位毫米
    Width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbWmsSkuUpdateRequest) GetAPIName() string {
	return "taobao.wlb.wms.sku.update"
}

func (req *TaobaoWlbWmsSkuUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 38)
    
    params["AdventLifecycle"] = req.AdventLifecycle
    
    params["ApprovalNumber"] = req.ApprovalNumber
    
    params["Attribute"] = req.Attribute
    
    params["BarCode"] = req.BarCode
    
    params["Brand"] = req.Brand
    
    params["BrandName"] = req.BrandName
    
    params["Category"] = req.Category
    
    params["CategoryName"] = req.CategoryName
    
    params["Color"] = req.Color
    
    params["CostPrice"] = req.CostPrice
    
    params["ExtendFields"] = req.ExtendFields
    
    params["GrossWeight"] = req.GrossWeight
    
    params["Height"] = req.Height
    
    params["IsAreaSale"] = req.IsAreaSale
    
    params["IsBatchMgt"] = req.IsBatchMgt
    
    params["IsDanger"] = req.IsDanger
    
    params["IsHygroscopic"] = req.IsHygroscopic
    
    params["IsShelflife"] = req.IsShelflife
    
    params["IsSnMgt"] = req.IsSnMgt
    
    params["ItemId"] = req.ItemId
    
    params["ItemPrice"] = req.ItemPrice
    
    params["Length"] = req.Length
    
    params["Lifecycle"] = req.Lifecycle
    
    params["LockupLifecycle"] = req.LockupLifecycle
    
    params["Name"] = req.Name
    
    params["NetWeight"] = req.NetWeight
    
    params["OriginAddress"] = req.OriginAddress
    
    params["Pcs"] = req.Pcs
    
    params["RejectLifecycle"] = req.RejectLifecycle
    
    params["Size"] = req.Size
    
    params["Specification"] = req.Specification
    
    params["StoreCode"] = req.StoreCode
    
    params["TagPrice"] = req.TagPrice
    
    params["Title"] = req.Title
    
    params["Type"] = req.Type
    
    params["UseYn"] = req.UseYn
    
    params["Volume"] = req.Volume
    
    params["Width"] = req.Width
    
    return params
}

// TaobaoWlbWmsSkuUpdateResponse 商品信息的更新
type TaobaoWlbWmsSkuUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoWlbWmsStockOutOrderNotifyRequest 出库单通知
type TaobaoWlbWmsStockOutOrderNotifyRequest struct {
    
    // CarNo optional车牌号
    CarNo string `json:"car_no";xml:"car_no"`
    
    // CarriersName optional承运商名称
    CarriersName string `json:"carriers_name";xml:"carriers_name"`
    
    // ExtendFields optional拓展属性
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // OrderCode requiredERP单据ID
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderCreateTime required订单创建时间
    OrderCreateTime time.Time `json:"order_create_time";xml:"order_create_time"`
    
    // OrderItemList optional订单商品信息列表
    OrderItemList Orderitemlistwlbwmsstockoutordernotify `json:"order_item_list";xml:"order_item_list"`
    
    // OrderType required单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
    OrderType int64 `json:"order_type";xml:"order_type"`
    
    // OutboundTypeDesc optionalERP可选择性文本透传至WMS
    OutboundTypeDesc string `json:"outbound_type_desc";xml:"outbound_type_desc"`
    
    // PickCall optional取件人电话
    PickCall string `json:"pick_call";xml:"pick_call"`
    
    // PickId optional取件人身份证ID
    PickId string `json:"pick_id";xml:"pick_id"`
    
    // PickName optional取件人姓名
    PickName string `json:"pick_name";xml:"pick_name"`
    
    // PrevOrderCode optional前物流订单号，如退货入库单可能会用到
    PrevOrderCode string `json:"prev_order_code";xml:"prev_order_code"`
    
    // ReceiverInfo optional收件人信息
    ReceiverInfo Receiverwlbwmsstockoutordernotify `json:"receiver_info";xml:"receiver_info"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // SendTime optional要求出库日期
    SendTime time.Time `json:"send_time";xml:"send_time"`
    
    // SenderInfo optional发货方信息
    SenderInfo Senderwlbwmsstockoutordernotify `json:"sender_info";xml:"sender_info"`
    
    // StoreCode required仓储编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TransportMode optional出库方式
    TransportMode string `json:"transport_mode";xml:"transport_mode"`
    
}

func (req *TaobaoWlbWmsStockOutOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.out.order.notify"
}

func (req *TaobaoWlbWmsStockOutOrderNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 18)
    
    params["CarNo"] = req.CarNo
    
    params["CarriersName"] = req.CarriersName
    
    params["ExtendFields"] = req.ExtendFields
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderCreateTime"] = req.OrderCreateTime
    
    params["OrderItemList"] = req.OrderItemList
    
    params["OrderType"] = req.OrderType
    
    params["OutboundTypeDesc"] = req.OutboundTypeDesc
    
    params["PickCall"] = req.PickCall
    
    params["PickId"] = req.PickId
    
    params["PickName"] = req.PickName
    
    params["PrevOrderCode"] = req.PrevOrderCode
    
    params["ReceiverInfo"] = req.ReceiverInfo
    
    params["Remark"] = req.Remark
    
    params["SendTime"] = req.SendTime
    
    params["SenderInfo"] = req.SenderInfo
    
    params["StoreCode"] = req.StoreCode
    
    params["TransportMode"] = req.TransportMode
    
    return params
}

// TaobaoWlbWmsStockOutOrderNotifyResponse 出库单通知
type TaobaoWlbWmsStockOutOrderNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OrderCode Basic仓储订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误详细
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// CainiaoWaybillIiGetRequest 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaoWaybillIiGetRequest struct {
    
    // ParamWaybillCloudPrintApplyNewRequest required入参信息
    ParamWaybillCloudPrintApplyNewRequest WaybillCloudPrintApplyNewRequest `json:"param_waybill_cloud_print_apply_new_request";xml:"param_waybill_cloud_print_apply_new_request"`
    
}

func (req *CainiaoWaybillIiGetRequest) GetAPIName() string {
	return "cainiao.waybill.ii.get"
}

func (req *CainiaoWaybillIiGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamWaybillCloudPrintApplyNewRequest"] = req.ParamWaybillCloudPrintApplyNewRequest
    
    return params
}

// CainiaoWaybillIiGetResponse 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaoWaybillIiGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Modules Object Array系统自动生成
    Modules WaybillCloudPrintResponse `json:"modules";xml:"modules"`
    
}

// TaobaoWlbWmsConsignOrderNotifyRequest 发货订单通知
type TaobaoWlbWmsConsignOrderNotifyRequest struct {
    
    // AlipayNo optional废弃，支付宝交易号
    AlipayNo string `json:"alipay_no";xml:"alipay_no"`
    
    // ArAmount optional订单应收金额，消费者还需要付的金额，单位分
    ArAmount int64 `json:"ar_amount";xml:"ar_amount"`
    
    // CarNo optional废弃，车牌号
    CarNo string `json:"car_no";xml:"car_no"`
    
    // CarriersName optional废弃，承运商名称
    CarriersName string `json:"carriers_name";xml:"carriers_name"`
    
    // DeliverRequirements optional配送要求
    DeliverRequirements Deliverrequirementswlbwmsconsignordernotify `json:"deliver_requirements";xml:"deliver_requirements"`
    
    // DiscountAmount optional订单优惠金额，整单优惠金额，单位分
    DiscountAmount int64 `json:"discount_amount";xml:"discount_amount"`
    
    // ExtendFields optional拓展属性
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // GotAmount optional订单已付金额，消费者已经支付的金额，单位分
    GotAmount int64 `json:"got_amount";xml:"got_amount"`
    
    // InvoiceInfoList optional发票信息列表
    InvoiceInfoList Invoicelistwlbwmsconsignordernotify `json:"invoice_info_list";xml:"invoice_info_list"`
    
    // OrderAmount optional订单总金额,=总商品金额-订单优惠金额+快递费用，单位分
    OrderAmount int64 `json:"order_amount";xml:"order_amount"`
    
    // OrderCode requiredERP订单号
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderCreateTime optional订单创建时间,ERP创建订单时间
    OrderCreateTime time.Time `json:"order_create_time";xml:"order_create_time"`
    
    // OrderExaminationTime optional订单审核时间,ERP创建支付时间
    OrderExaminationTime time.Time `json:"order_examination_time";xml:"order_examination_time"`
    
    // OrderFlag optional订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票)
    OrderFlag string `json:"order_flag";xml:"order_flag"`
    
    // OrderItemList optional订单商品信息列表
    OrderItemList Orderitemlistwlbwmsconsignordernotify `json:"order_item_list";xml:"order_item_list"`
    
    // OrderPayTime optional订单支付时间
    OrderPayTime time.Time `json:"order_pay_time";xml:"order_pay_time"`
    
    // OrderPriority optional废弃，订单优先级0 -10，根据订单作业优先级设置，数字越大，优先级越高
    OrderPriority int64 `json:"order_priority";xml:"order_priority"`
    
    // OrderShopCreateTime optional下单时间，订单在交易平台创建时间
    OrderShopCreateTime time.Time `json:"order_shop_create_time";xml:"order_shop_create_time"`
    
    // OrderSource optional订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
    OrderSource int64 `json:"order_source";xml:"order_source"`
    
    // OrderSubSource optional废弃，交易内部来源，文本透传 WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算) 一笔订单可能同时有以上多个标记，则以逗号分隔
    OrderSubSource string `json:"order_sub_source";xml:"order_sub_source"`
    
    // OrderType required单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单
    OrderType int64 `json:"order_type";xml:"order_type"`
    
    // PickerCall optional废弃，取件人电话
    PickerCall string `json:"picker_call";xml:"picker_call"`
    
    // PickerId optional废弃，取件人身份证
    PickerId string `json:"picker_id";xml:"picker_id"`
    
    // PickerName optional废弃，取件人姓名
    PickerName string `json:"picker_name";xml:"picker_name"`
    
    // Postfee optional快递费用，单位分
    Postfee int64 `json:"postfee";xml:"postfee"`
    
    // PrevOrderCode optional前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容
    PrevOrderCode string `json:"prev_order_code";xml:"prev_order_code"`
    
    // ReceiverInfo optional收件人信息
    ReceiverInfo Receiverwlbwmsconsignordernotify `json:"receiver_info";xml:"receiver_info"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // SenderInfo optional发货方信息
    SenderInfo Senderwlbwmsconsignordernotify `json:"sender_info";xml:"sender_info"`
    
    // ServiceFee optionalCOD服务费，单位分
    ServiceFee int64 `json:"service_fee";xml:"service_fee"`
    
    // StoreCode optional仓库编码，此字段为空时，由菜鸟路由仓库发货
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TmsServiceCode optional快递公司编码，此字段为空时，由菜鸟选择快递配送
    TmsServiceCode string `json:"tms_service_code";xml:"tms_service_code"`
    
    // TmsServiceName optional快递公司名称
    TmsServiceName string `json:"tms_service_name";xml:"tms_service_name"`
    
    // TransportMode optional废弃，出库方式（自提，快递，销毁）
    TransportMode string `json:"transport_mode";xml:"transport_mode"`
    
}

func (req *TaobaoWlbWmsConsignOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.order.notify"
}

func (req *TaobaoWlbWmsConsignOrderNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 34)
    
    params["AlipayNo"] = req.AlipayNo
    
    params["ArAmount"] = req.ArAmount
    
    params["CarNo"] = req.CarNo
    
    params["CarriersName"] = req.CarriersName
    
    params["DeliverRequirements"] = req.DeliverRequirements
    
    params["DiscountAmount"] = req.DiscountAmount
    
    params["ExtendFields"] = req.ExtendFields
    
    params["GotAmount"] = req.GotAmount
    
    params["InvoiceInfoList"] = req.InvoiceInfoList
    
    params["OrderAmount"] = req.OrderAmount
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderCreateTime"] = req.OrderCreateTime
    
    params["OrderExaminationTime"] = req.OrderExaminationTime
    
    params["OrderFlag"] = req.OrderFlag
    
    params["OrderItemList"] = req.OrderItemList
    
    params["OrderPayTime"] = req.OrderPayTime
    
    params["OrderPriority"] = req.OrderPriority
    
    params["OrderShopCreateTime"] = req.OrderShopCreateTime
    
    params["OrderSource"] = req.OrderSource
    
    params["OrderSubSource"] = req.OrderSubSource
    
    params["OrderType"] = req.OrderType
    
    params["PickerCall"] = req.PickerCall
    
    params["PickerId"] = req.PickerId
    
    params["PickerName"] = req.PickerName
    
    params["Postfee"] = req.Postfee
    
    params["PrevOrderCode"] = req.PrevOrderCode
    
    params["ReceiverInfo"] = req.ReceiverInfo
    
    params["Remark"] = req.Remark
    
    params["SenderInfo"] = req.SenderInfo
    
    params["ServiceFee"] = req.ServiceFee
    
    params["StoreCode"] = req.StoreCode
    
    params["TmsServiceCode"] = req.TmsServiceCode
    
    params["TmsServiceName"] = req.TmsServiceName
    
    params["TransportMode"] = req.TransportMode
    
    return params
}

// TaobaoWlbWmsConsignOrderNotifyResponse 发货订单通知
type TaobaoWlbWmsConsignOrderNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ConsignOrderList Object Array系统自动生成
    ConsignOrderList Consignorderlist `json:"consign_order_list";xml:"consign_order_list"`
    
    // OrderCode Basic订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误详细
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// CainiaoWaybillIiUpdateRequest 商家更新电子面单号对应的面单信息。
type CainiaoWaybillIiUpdateRequest struct {
    
    // ParamWaybillCloudPrintUpdateRequest required更新请求信息
    ParamWaybillCloudPrintUpdateRequest WaybillCloudPrintUpdateRequest `json:"param_waybill_cloud_print_update_request";xml:"param_waybill_cloud_print_update_request"`
    
}

func (req *CainiaoWaybillIiUpdateRequest) GetAPIName() string {
	return "cainiao.waybill.ii.update"
}

func (req *CainiaoWaybillIiUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamWaybillCloudPrintUpdateRequest"] = req.ParamWaybillCloudPrintUpdateRequest
    
    return params
}

// CainiaoWaybillIiUpdateResponse 商家更新电子面单号对应的面单信息。
type CainiaoWaybillIiUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // PrintData Basic模板内容
    PrintData string `json:"print_data";xml:"print_data"`
    
    // WaybillCode Basic面单号
    WaybillCode string `json:"waybill_code";xml:"waybill_code"`
    
}

// TaobaoWlbWmsOrderCancelNotifyRequest 单据取消接口
type TaobaoWlbWmsOrderCancelNotifyRequest struct {
    
    // OrderCode required订单类型
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderType required单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
    OrderType string `json:"order_type";xml:"order_type"`
    
    // StoreCode required仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbWmsOrderCancelNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.order.cancel.notify"
}

func (req *TaobaoWlbWmsOrderCancelNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderType"] = req.OrderType
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoWlbWmsOrderCancelNotifyResponse 单据取消接口
type TaobaoWlbWmsOrderCancelNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误详细
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// CainiaoEndpointLockerTopOrderTrackingNewRequest 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
type CainiaoEndpointLockerTopOrderTrackingNewRequest struct {
    
    // TrackInfo required回传信息
    TrackInfo CollectTrackingInfo `json:"track_info";xml:"track_info"`
    
}

func (req *CainiaoEndpointLockerTopOrderTrackingNewRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.tracking.new"
}

func (req *CainiaoEndpointLockerTopOrderTrackingNewRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TrackInfo"] = req.TrackInfo
    
    return params
}

// CainiaoEndpointLockerTopOrderTrackingNewResponse 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
type CainiaoEndpointLockerTopOrderTrackingNewResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResult `json:"result";xml:"result"`
    
}

// TaobaoQimenEntryorderQueryRequest ERP调用接口，查询入库单信息;
type TaobaoQimenEntryorderQueryRequest struct {
    
    // EntryOrderCode required入库单编码
    EntryOrderCode string `json:"entryOrderCode";xml:"entryOrderCode"`
    
    // EntryOrderId optional仓储系统入库单ID
    EntryOrderId string `json:"entryOrderId";xml:"entryOrderId"`
    
    // ExtOrderCode optionalextOrderCode
    ExtOrderCode string `json:"extOrderCode";xml:"extOrderCode"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderType optionalorderType
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页(从1开始)
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页orderLine条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenEntryorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.query"
}

func (req *TaobaoQimenEntryorderQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["EntryOrderCode"] = req.EntryOrderCode
    
    params["EntryOrderId"] = req.EntryOrderId
    
    params["ExtOrderCode"] = req.ExtOrderCode
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderType"] = req.OrderType
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["Remark"] = req.Remark
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenEntryorderQueryResponse ERP调用接口，查询入库单信息;
type TaobaoQimenEntryorderQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // EntryOrder Object入库单信息
    EntryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderLines Object Array入库单单据信息
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // TotalLines BasicorderLines总条数
    TotalLines int64 `json:"totalLines";xml:"totalLines"`
    
}

// TmallExchangeGetRequest 获取单笔换货详情
type TmallExchangeGetRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeGetRequest) GetAPIName() string {
	return "tmall.exchange.get"
}

func (req *TmallExchangeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    return params
}

// TmallExchangeGetResponse 获取单笔换货详情
type TmallExchangeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

// TaobaoQimenDeliveryorderQueryRequest ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryRequest struct {
    
    // DeliveryOrderCode optional奇门仓储字段,说明,string(50),,
    DeliveryOrderCode string `json:"deliveryOrderCode";xml:"deliveryOrderCode"`
    
    // DeliveryOrderId optional奇门仓储字段,说明,string(50),,
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderCode required发库单号
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderId optional仓储系统发库单号
    OrderId string `json:"orderId";xml:"orderId"`
    
    // OrderSourceCode optional交易单号
    OrderSourceCode string `json:"orderSourceCode";xml:"orderSourceCode"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Page required当前页
    Page int64 `json:"page";xml:"page"`
    
    // PageSize required每页orderLine条数(最多100条)
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // Status optional奇门仓储字段,说明,string(50),,
    Status string `json:"status";xml:"status"`
    
    // WarehouseCode optional仓库编码
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenDeliveryorderQueryRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.query"
}

func (req *TaobaoQimenDeliveryorderQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 12)
    
    params["DeliveryOrderCode"] = req.DeliveryOrderCode
    
    params["DeliveryOrderId"] = req.DeliveryOrderId
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderId"] = req.OrderId
    
    params["OrderSourceCode"] = req.OrderSourceCode
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["Remark"] = req.Remark
    
    params["Status"] = req.Status
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenDeliveryorderQueryResponse ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // DeliveryOrder Object发货单信息
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // OrderLines Object Array单据列表
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // Packages Object Array包裹信息
    Packages Package `json:"packages";xml:"packages"`
    
    // TotalLines BasicorderLines总条数
    TotalLines int64 `json:"totalLines";xml:"totalLines"`
    
}

// CainiaoEndpointLockerTopOrderNoticeRequest 合作公司对订单手动触发短信，有次数限制
type CainiaoEndpointLockerTopOrderNoticeRequest struct {
    
    // MailNo required运单号
    MailNo string `json:"mail_no";xml:"mail_no"`
    
    // OrderCode required合作公司唯一订单编号
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // SceneCode required场景编号：0：重发短信，1：催取短信
    SceneCode int64 `json:"scene_code";xml:"scene_code"`
    
    // StationId required站点ID
    StationId string `json:"station_id";xml:"station_id"`
    
}

func (req *CainiaoEndpointLockerTopOrderNoticeRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.notice"
}

func (req *CainiaoEndpointLockerTopOrderNoticeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["MailNo"] = req.MailNo
    
    params["OrderCode"] = req.OrderCode
    
    params["SceneCode"] = req.SceneCode
    
    params["StationId"] = req.StationId
    
    return params
}

// CainiaoEndpointLockerTopOrderNoticeResponse 合作公司对订单手动触发短信，有次数限制
type CainiaoEndpointLockerTopOrderNoticeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResult `json:"result";xml:"result"`
    
}

// CainiaoEndpointLockerTopOrderNoticesendQueryRequest 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoEndpointLockerTopOrderNoticesendQueryRequest struct {
    
    // GetterPhone optional收件人手机号
    GetterPhone string `json:"getter_phone";xml:"getter_phone"`
    
    // MailNo required运单号
    MailNo string `json:"mail_no";xml:"mail_no"`
    
    // StationId required站点id
    StationId string `json:"station_id";xml:"station_id"`
    
}

func (req *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.noticesend.query"
}

func (req *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["GetterPhone"] = req.GetterPhone
    
    params["MailNo"] = req.MailNo
    
    params["StationId"] = req.StationId
    
    return params
}

// CainiaoEndpointLockerTopOrderNoticesendQueryResponse 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoEndpointLockerTopOrderNoticesendQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result SingleResult `json:"result";xml:"result"`
    
}

// TmallExchangeReturngoodsAgreeRequest 卖家确认收货
type TmallExchangeReturngoodsAgreeRequest struct {
    
    // DisputeId required换货单号ID
    DisputeId int64 `json:"dispute_id";xml:"dispute_id"`
    
    // Fields required返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TmallExchangeReturngoodsAgreeRequest) GetAPIName() string {
	return "tmall.exchange.returngoods.agree"
}

func (req *TmallExchangeReturngoodsAgreeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["DisputeId"] = req.DisputeId
    
    params["Fields"] = req.Fields
    
    return params
}

// TmallExchangeReturngoodsAgreeResponse 卖家确认收货
type TmallExchangeReturngoodsAgreeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result ExchangeBaseResponse `json:"result";xml:"result"`
    
}

// CainiaoCntmsLogisticsOrderConsignRequest 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaoCntmsLogisticsOrderConsignRequest struct {
    
    // Content optional配送发货信息
    Content CnTmsLogisticsOrderConsignContent `json:"content";xml:"content"`
    
}

func (req *CainiaoCntmsLogisticsOrderConsignRequest) GetAPIName() string {
	return "cainiao.cntms.logistics.order.consign"
}

func (req *CainiaoCntmsLogisticsOrderConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// CainiaoCntmsLogisticsOrderConsignResponse 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaoCntmsLogisticsOrderConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // LogisticsOrderCode Basic物流单号
    LogisticsOrderCode string `json:"logistics_order_code";xml:"logistics_order_code"`
    
}

// TmallExchangeReceiveGetRequest 卖家查询换货列表
type TmallExchangeReceiveGetRequest struct {
    
    // BizOrderId optional正向订单号
    BizOrderId int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    // BuyerId optional买家id
    BuyerId int64 `json:"buyer_id";xml:"buyer_id"`
    
    // BuyerNick optional买家昵称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // DisputeStatusArray optional换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
    DisputeStatusArray int64 `json:"dispute_status_array";xml:"dispute_status_array"`
    
    // EndCreatedTime optional查询申请时间段的结束时间点
    EndCreatedTime time.Time `json:"end_created_time";xml:"end_created_time"`
    
    // EndGmtModifedTime optional查询修改时间段的结束时间点
    EndGmtModifedTime time.Time `json:"end_gmt_modifed_time";xml:"end_gmt_modifed_time"`
    
    // Fields required返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
    Fields string `json:"fields";xml:"fields"`
    
    // LogisticNo optional快递单号
    LogisticNo string `json:"logistic_no";xml:"logistic_no"`
    
    // PageNo optional页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // RefundIdArray optional退款单号ID列表，最多只能输入20个id
    RefundIdArray int64 `json:"refund_id_array";xml:"refund_id_array"`
    
    // StartCreatedTime optional查询申请时间段的开始时间点
    StartCreatedTime time.Time `json:"start_created_time";xml:"start_created_time"`
    
    // StartGmtModifiedTime optional查询修改时间段的开始时间点
    StartGmtModifiedTime time.Time `json:"start_gmt_modified_time";xml:"start_gmt_modified_time"`
    
}

func (req *TmallExchangeReceiveGetRequest) GetAPIName() string {
	return "tmall.exchange.receive.get"
}

func (req *TmallExchangeReceiveGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 13)
    
    params["BizOrderId"] = req.BizOrderId
    
    params["BuyerId"] = req.BuyerId
    
    params["BuyerNick"] = req.BuyerNick
    
    params["DisputeStatusArray"] = req.DisputeStatusArray
    
    params["EndCreatedTime"] = req.EndCreatedTime
    
    params["EndGmtModifedTime"] = req.EndGmtModifedTime
    
    params["Fields"] = req.Fields
    
    params["LogisticNo"] = req.LogisticNo
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["RefundIdArray"] = req.RefundIdArray
    
    params["StartCreatedTime"] = req.StartCreatedTime
    
    params["StartGmtModifiedTime"] = req.StartGmtModifiedTime
    
    return params
}

// TmallExchangeReceiveGetResponse 卖家查询换货列表
type TmallExchangeReceiveGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorCodes Basic错误码
    ErrorCodes string `json:"error_codes";xml:"error_codes"`
    
    // ErrorMsg Basic错误信息
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // Exception Basic所抛出异常
    Exception map[string]interface{} `json:"exception";xml:"exception"`
    
    // HasNext Basic是否还有下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // PageResults Basic当前页的换货单数量
    PageResults int64 `json:"page_results";xml:"page_results"`
    
    // Results Object Array返回结果
    Results Exchange `json:"results";xml:"results"`
    
    // TotalResults Basic所有符合查询条件的换货单的数量
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// CainiaoCntmsMailnoGetRequest 打印面单时，通过此接口获取面单号及打打印信息
type CainiaoCntmsMailnoGetRequest struct {
    
    // Content optional获取菜鸟配送电子面单请求参数
    Content CnTmsMailnoGetContent `json:"content";xml:"content"`
    
}

func (req *CainiaoCntmsMailnoGetRequest) GetAPIName() string {
	return "cainiao.cntms.mailno.get"
}

func (req *CainiaoCntmsMailnoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// CainiaoCntmsMailnoGetResponse 打印面单时，通过此接口获取面单号及打打印信息
type CainiaoCntmsMailnoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AllocatorCode Basic揽货商（分拨中心）编码
    AllocatorCode string `json:"allocator_code";xml:"allocator_code"`
    
    // AllocatorName Basic揽货商（分拨中心）名称
    AllocatorName string `json:"allocator_name";xml:"allocator_name"`
    
    // LogisticsCode Basic物流商公司编码
    // （ERP在调用菜鸟发货接口时此字段赋值到tms_code, 调用淘宝“自己联系物流（线下物流）发货”时，做为company_code传入）
    LogisticsCode string `json:"logistics_code";xml:"logistics_code"`
    
    // LogisticsName Basic物流商名称
    LogisticsName string `json:"logistics_name";xml:"logistics_name"`
    
    // Mailno Basic面单号
    Mailno string `json:"mailno";xml:"mailno"`
    
    // PackageCenterCode Basic集包地代码
    PackageCenterCode string `json:"package_center_code";xml:"package_center_code"`
    
    // PackageCenterName Basic集包地名称
    PackageCenterName string `json:"package_center_name";xml:"package_center_name"`
    
    // SecDistribution Basic二级流向信息
    SecDistribution string `json:"sec_distribution";xml:"sec_distribution"`
    
    // ShortAddress Basic一级流向信息，大头笔
    ShortAddress string `json:"short_address";xml:"short_address"`
    
    // TmsCode Basic二级配送公司编码
    TmsCode string `json:"tms_code";xml:"tms_code"`
    
    // TmsName Basic二级配送公司名称
    TmsName string `json:"tms_name";xml:"tms_name"`
    
}

// TaobaoLocationRelationQueryRequest 地点关联关系查询 
//门店和仓库关联关系查询
type TaobaoLocationRelationQueryRequest struct {
    
    // LocationRelation required关系查询
    LocationRelation LocationRelationDto `json:"location_relation";xml:"location_relation"`
    
}

func (req *TaobaoLocationRelationQueryRequest) GetAPIName() string {
	return "taobao.location.relation.query"
}

func (req *TaobaoLocationRelationQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["LocationRelation"] = req.LocationRelation
    
    return params
}

// TaobaoLocationRelationQueryResponse 地点关联关系查询 
//门店和仓库关联关系查询
type TaobaoLocationRelationQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResult `json:"result";xml:"result"`
    
}

// TaobaoLocationRelationEditRequest 地点关联关系增量编辑
type TaobaoLocationRelationEditRequest struct {
    
    // LocationRelationList required关系对象列表
    LocationRelationList LocationRelationDto `json:"location_relation_list";xml:"location_relation_list"`
    
}

func (req *TaobaoLocationRelationEditRequest) GetAPIName() string {
	return "taobao.location.relation.edit"
}

func (req *TaobaoLocationRelationEditRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["LocationRelationList"] = req.LocationRelationList
    
    return params
}

// TaobaoLocationRelationEditResponse 地点关联关系增量编辑
type TaobaoLocationRelationEditResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorMsg Basic错误信息
    ErrorMsg string `json:"error_msg";xml:"error_msg"`
    
    // Errorcode Basic错误码
    Errorcode string `json:"errorcode";xml:"errorcode"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// CainiaoCloudprintCmdprintRenderRequest isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
type CainiaoCloudprintCmdprintRenderRequest struct {
    
    // Params required参数对象
    Params CmdRenderParams `json:"params";xml:"params"`
    
}

func (req *CainiaoCloudprintCmdprintRenderRequest) GetAPIName() string {
	return "cainiao.cloudprint.cmdprint.render"
}

func (req *CainiaoCloudprintCmdprintRenderRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Params"] = req.Params
    
    return params
}

// CainiaoCloudprintCmdprintRenderResponse isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
type CainiaoCloudprintCmdprintRenderResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CmdContent Basic指令集内容串
    CmdContent string `json:"cmd_content";xml:"cmd_content"`
    
    // CmdEncoding Basic指令集编码方式：origin-原串 gzip-采用gzip压缩并base64编码
    CmdEncoding string `json:"cmd_encoding";xml:"cmd_encoding"`
    
    // RetCode Basic0成功，非0失败
    RetCode string `json:"ret_code";xml:"ret_code"`
    
    // RetMsg Basic错误信息
    RetMsg string `json:"ret_msg";xml:"ret_msg"`
    
}

// TaobaoWlbWmsReturnBillGetRequest 通过订单号获取单个销退单收货信息。
type TaobaoWlbWmsReturnBillGetRequest struct {
    
    // CnOrderCode special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    CnOrderCode string `json:"cn_order_code";xml:"cn_order_code"`
    
    // OrderCode specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    OrderCode string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsReturnBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.return.bill.get"
}

func (req *TaobaoWlbWmsReturnBillGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CnOrderCode"] = req.CnOrderCode
    
    params["OrderCode"] = req.OrderCode
    
    return params
}

// TaobaoWlbWmsReturnBillGetResponse 通过订单号获取单个销退单收货信息。
type TaobaoWlbWmsReturnBillGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ReturnOrderInfo Object回退订单信息
    ReturnOrderInfo CainiaoReturnBillReturnorderinfo `json:"return_order_info";xml:"return_order_info"`
    
}

// TaobaoWlbWmsStockInBillGetRequest 获取入库单信息
type TaobaoWlbWmsStockInBillGetRequest struct {
    
    // CnOrderCode special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    CnOrderCode string `json:"cn_order_code";xml:"cn_order_code"`
    
    // OrderCode specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    OrderCode string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsStockInBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.in.bill.get"
}

func (req *TaobaoWlbWmsStockInBillGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CnOrderCode"] = req.CnOrderCode
    
    params["OrderCode"] = req.OrderCode
    
    return params
}

// TaobaoWlbWmsStockInBillGetResponse 获取入库单信息
type TaobaoWlbWmsStockInBillGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // StockInInfo Object入库单信息
    StockInInfo CainiaoStockInBillStockininfo `json:"stock_in_info";xml:"stock_in_info"`
    
}

// TaobaoOcTradesBytagGetRequest 根据标签查询订单编号
type TaobaoOcTradesBytagGetRequest struct {
    
    // Page optional当前页
    Page int64 `json:"page";xml:"page"`
    
    // PageSize optional分页大小
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // TagName required标签名称
    TagName string `json:"tag_name";xml:"tag_name"`
    
    // TagType required标签类型，1官方，2自定义
    TagType int64 `json:"tag_type";xml:"tag_type"`
    
}

func (req *TaobaoOcTradesBytagGetRequest) GetAPIName() string {
	return "taobao.oc.trades.bytag.get"
}

func (req *TaobaoOcTradesBytagGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Page"] = req.Page
    
    params["PageSize"] = req.PageSize
    
    params["TagName"] = req.TagName
    
    params["TagType"] = req.TagType
    
    return params
}

// TaobaoOcTradesBytagGetResponse 根据标签查询订单编号
type TaobaoOcTradesBytagGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Tids Basic Array打了该标签的订单编号列表
    Tids int64 `json:"tids";xml:"tids"`
    
    // Totals Basic总数
    Totals int64 `json:"totals";xml:"totals"`
    
}

// TaobaoOcTradetagAttachRequest 对订单添加标签和更新标签。标签分为官方标签和自定义标签。
//官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731
//自定义标签有2个通用属性:
//    `show_str:给消费者显示的字符串（如果可以显示的话）
//    `pic_urls:图片url,地址必须是图片空间的url,最多5张
type TaobaoOcTradetagAttachRequest struct {
    
    // TagName required标签名称
    TagName string `json:"tag_name";xml:"tag_name"`
    
    // TagType optional标签类型       1：官方标签      2：自定义标签
    TagType int64 `json:"tag_type";xml:"tag_type"`
    
    // TagValue required标签值，json格式
    TagValue string `json:"tag_value";xml:"tag_value"`
    
    // Tid required订单id
    Tid int64 `json:"tid";xml:"tid"`
    
    // Visible optional该标签在消费者端是否显示,0:不显示,1：显示
    Visible int64 `json:"visible";xml:"visible"`
    
}

func (req *TaobaoOcTradetagAttachRequest) GetAPIName() string {
	return "taobao.oc.tradetag.attach"
}

func (req *TaobaoOcTradetagAttachRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["TagName"] = req.TagName
    
    params["TagType"] = req.TagType
    
    params["TagValue"] = req.TagValue
    
    params["Tid"] = req.Tid
    
    params["Visible"] = req.Visible
    
    return params
}

// TaobaoOcTradetagAttachResponse 对订单添加标签和更新标签。标签分为官方标签和自定义标签。
//官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731
//自定义标签有2个通用属性:
//    `show_str:给消费者显示的字符串（如果可以显示的话）
//    `pic_urls:图片url,地址必须是图片空间的url,最多5张
type TaobaoOcTradetagAttachResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic操作成功或者操作失败
    Result bool `json:"result";xml:"result"`
    
}

// TaobaoWlbOrderdetailDateGetRequest 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
type TaobaoWlbOrderdetailDateGetRequest struct {
    
    // EndTime required创建时间结束
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // PageNo optional分页下标
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页大小
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime required创建时间起始
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoWlbOrderdetailDateGetRequest) GetAPIName() string {
	return "taobao.wlb.orderdetail.date.get"
}

func (req *TaobaoWlbOrderdetailDateGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["EndTime"] = req.EndTime
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    return params
}

// TaobaoWlbOrderdetailDateGetResponse 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
type TaobaoWlbOrderdetailDateGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OrderDetailList Object Array物流宝订单，并且包含订单详情
    OrderDetailList WlbOrderDetail `json:"order_detail_list";xml:"order_detail_list"`
    
    // TotalCount Basic总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoOcTradetagsGetRequest 根据订单查询订单标签。<br/>
//返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
//官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
//主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
type TaobaoOcTradetagsGetRequest struct {
    
    // History optional是否查询历史标签
    History int64 `json:"history";xml:"history"`
    
    // TagNames optional不填，则不做标签名称过滤
    TagNames string `json:"tag_names";xml:"tag_names"`
    
    // TagTypes optional不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签
    TagTypes string `json:"tag_types";xml:"tag_types"`
    
    // Tid required交易主订单id
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoOcTradetagsGetRequest) GetAPIName() string {
	return "taobao.oc.tradetags.get"
}

func (req *TaobaoOcTradetagsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["History"] = req.History
    
    params["TagNames"] = req.TagNames
    
    params["TagTypes"] = req.TagTypes
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoOcTradetagsGetResponse 根据订单查询订单标签。<br/>
//返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
//官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
//主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
type TaobaoOcTradetagsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TradeTags Object Array返回结果
    TradeTags TradeTagRelationDo `json:"trade_tags";xml:"trade_tags"`
    
}

// QimenTaobaoCrmOrderSyncRequest CRM对于会员数据分析需要基于会员的购买行为，购买什么商品，价格，是否退换，使用什么优惠以及会员基础信息等等。这些信息通过销售订单数据来获取。
type QimenTaobaoCrmOrderSyncRequest struct {
    
    // Customerid requiredcustomerid
    Customerid string `json:"customerid";xml:"customerid"`
    
    // Data optional订单数据
    Data Data `json:"data";xml:"data"`
    
    // DataType optional数据类型（zx_order直销订单;fx_order分销订单）
    DataType string `json:"data_type";xml:"data_type"`
    
    // FromNodeId optional数据来源ID
    FromNodeId string `json:"from_node_id";xml:"from_node_id"`
    
    // FromType optional数据来源类型（taobao淘宝;JD京东;yihaodian一号店;dangdang当当;suning苏宁易购;amazon亚马逊;yinati银泰;mogujie蘑菇街;alibaba阿里巴巴;vop唯品会;meilishuo美丽说;youzan有赞;weixin微信;other其他）
    FromType string `json:"from_type";xml:"from_type"`
    
    // MsgId optional全局唯一任务编号
    MsgId string `json:"msg_id";xml:"msg_id"`
    
}

func (req *QimenTaobaoCrmOrderSyncRequest) GetAPIName() string {
	return "qimen.taobao.crm.order.sync"
}

func (req *QimenTaobaoCrmOrderSyncRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Customerid"] = req.Customerid
    
    params["Data"] = req.Data
    
    params["DataType"] = req.DataType
    
    params["FromNodeId"] = req.FromNodeId
    
    params["FromType"] = req.FromType
    
    params["MsgId"] = req.MsgId
    
    return params
}

// QimenTaobaoCrmOrderSyncResponse CRM对于会员数据分析需要基于会员的购买行为，购买什么商品，价格，是否退换，使用什么优惠以及会员基础信息等等。这些信息通过销售订单数据来获取。
type QimenTaobaoCrmOrderSyncResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic0成功(其他失败)
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// CainiaoMerchantInventoryAdjustRequest 商家仓库存调整接口，目前仅支持全量更新
type CainiaoMerchantInventoryAdjustRequest struct {
    
    // AdjustRequest required商家仓编辑库存
    AdjustRequest MerStoreInvAdjustDto `json:"adjust_request";xml:"adjust_request"`
    
    // AppName required调用方应用名
    AppName string `json:"app_name";xml:"app_name"`
    
    // Operation optional操作
    Operation string `json:"operation";xml:"operation"`
    
}

func (req *CainiaoMerchantInventoryAdjustRequest) GetAPIName() string {
	return "cainiao.merchant.inventory.adjust"
}

func (req *CainiaoMerchantInventoryAdjustRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["AdjustRequest"] = req.AdjustRequest
    
    params["AppName"] = req.AppName
    
    params["Operation"] = req.Operation
    
    return params
}

// CainiaoMerchantInventoryAdjustResponse 商家仓库存调整接口，目前仅支持全量更新
type CainiaoMerchantInventoryAdjustResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResultDto `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsCainiaoBillQueryRequest 查询单据列表
type TaobaoWlbWmsCainiaoBillQueryRequest struct {
    
    // EndModifiedTime required起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
    EndModifiedTime time.Time `json:"end_modified_time";xml:"end_modified_time"`
    
    // OrderType optional订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
    OrderType string `json:"order_type";xml:"order_type"`
    
    // PageNo optional页码。（大于0的整数。默认为1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数。（每页条数不超过50条。默认为50）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartModifiedTime required结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
    StartModifiedTime time.Time `json:"start_modified_time";xml:"start_modified_time"`
    
}

func (req *TaobaoWlbWmsCainiaoBillQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.cainiao.bill.query"
}

func (req *TaobaoWlbWmsCainiaoBillQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["EndModifiedTime"] = req.EndModifiedTime
    
    params["OrderType"] = req.OrderType
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartModifiedTime"] = req.StartModifiedTime
    
    return params
}

// TaobaoWlbWmsCainiaoBillQueryResponse 查询单据列表
type TaobaoWlbWmsCainiaoBillQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OrderInfoList Object Array订单列表信息
    OrderInfoList CainiaoBillQueryOrderinfolist `json:"order_info_list";xml:"order_info_list"`
    
    // TotalCount Basic总条数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// CainiaoWaybillIiCancelRequest 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaoWaybillIiCancelRequest struct {
    
    // CpCode required快递公司code
    CpCode string `json:"cp_code";xml:"cp_code"`
    
    // WaybillCode required电子面单号
    WaybillCode string `json:"waybill_code";xml:"waybill_code"`
    
}

func (req *CainiaoWaybillIiCancelRequest) GetAPIName() string {
	return "cainiao.waybill.ii.cancel"
}

func (req *CainiaoWaybillIiCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CpCode"] = req.CpCode
    
    params["WaybillCode"] = req.WaybillCode
    
    return params
}

// CainiaoWaybillIiCancelResponse 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaoWaybillIiCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CancelResult Basic调用取消是否成功
    CancelResult bool `json:"cancel_result";xml:"cancel_result"`
    
}

// CainiaoWaybillIiProductRequest 商家可以查询物流商的产品类型和服务能力。
type CainiaoWaybillIiProductRequest struct {
    
    // CpCode required快递公司code
    CpCode string `json:"cp_code";xml:"cp_code"`
    
}

func (req *CainiaoWaybillIiProductRequest) GetAPIName() string {
	return "cainiao.waybill.ii.product"
}

func (req *CainiaoWaybillIiProductRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["CpCode"] = req.CpCode
    
    return params
}

// CainiaoWaybillIiProductResponse 商家可以查询物流商的产品类型和服务能力。
type CainiaoWaybillIiProductResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ProductTypes Object Array返回值
    ProductTypes WaybillProductType `json:"product_types";xml:"product_types"`
    
}

// TaobaoRdcAligeniusAccountValidateRequest 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaoRdcAligeniusAccountValidateRequest struct {
    
}

func (req *TaobaoRdcAligeniusAccountValidateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.account.validate"
}

func (req *TaobaoRdcAligeniusAccountValidateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoRdcAligeniusAccountValidateResponse 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaoRdcAligeniusAccountValidateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsInventoryProfitlossGetRequest 通过订单列表批量获取库存损益单信息
type TaobaoWlbWmsInventoryProfitlossGetRequest struct {
    
    // CnOrderCode required菜鸟订单编码
    CnOrderCode string `json:"cn_order_code";xml:"cn_order_code"`
    
}

func (req *TaobaoWlbWmsInventoryProfitlossGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.profitloss.get"
}

func (req *TaobaoWlbWmsInventoryProfitlossGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["CnOrderCode"] = req.CnOrderCode
    
    return params
}

// TaobaoWlbWmsInventoryProfitlossGetResponse 通过订单列表批量获取库存损益单信息
type TaobaoWlbWmsInventoryProfitlossGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ProfitLossInfo Object损益信息
    ProfitLossInfo CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info";xml:"profit_loss_info"`
    
}

// CainiaoCloudprintMystdtemplatesGetRequest 获取用户使用的菜鸟电子面单
type CainiaoCloudprintMystdtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintMystdtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.mystdtemplates.get"
}

func (req *CainiaoCloudprintMystdtemplatesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// CainiaoCloudprintMystdtemplatesGetResponse 获取用户使用的菜鸟电子面单
type CainiaoCloudprintMystdtemplatesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回结果
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsConsignBillGetRequest 获取销售订单发货信息
type TaobaoWlbWmsConsignBillGetRequest struct {
    
    // CnOrderCode special菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
    CnOrderCode string `json:"cn_order_code";xml:"cn_order_code"`
    
    // OrderCode specialERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
    OrderCode string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsConsignBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.bill.get"
}

func (req *TaobaoWlbWmsConsignBillGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CnOrderCode"] = req.CnOrderCode
    
    params["OrderCode"] = req.OrderCode
    
    return params
}

// TaobaoWlbWmsConsignBillGetResponse 获取销售订单发货信息
type TaobaoWlbWmsConsignBillGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ConsignSendInfoList Object Array商品信息列表
    ConsignSendInfoList Consignsendinfolist `json:"consign_send_info_list";xml:"consign_send_info_list"`
    
}

// TaobaoWlbWmsStockOutBillGetRequest 通过订单号获取单个出库单发货信息
type TaobaoWlbWmsStockOutBillGetRequest struct {
    
    // CnOrderCode special菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    CnOrderCode string `json:"cn_order_code";xml:"cn_order_code"`
    
    // OrderCode specialERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    OrderCode string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbWmsStockOutBillGetRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.out.bill.get"
}

func (req *TaobaoWlbWmsStockOutBillGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CnOrderCode"] = req.CnOrderCode
    
    params["OrderCode"] = req.OrderCode
    
    return params
}

// TaobaoWlbWmsStockOutBillGetResponse 通过订单号获取单个出库单发货信息
type TaobaoWlbWmsStockOutBillGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // StockOutInfo Object出库信息
    StockOutInfo CainiaoStockOutBillStockoutinfo `json:"stock_out_info";xml:"stock_out_info"`
    
}

// CainiaoCloudprintStdtemplatesGetRequest 获取菜鸟标准电子面单模板
type CainiaoCloudprintStdtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintStdtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.stdtemplates.get"
}

func (req *CainiaoCloudprintStdtemplatesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// CainiaoCloudprintStdtemplatesGetResponse 获取菜鸟标准电子面单模板
type CainiaoCloudprintStdtemplatesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object结果集
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// TaobaoWangwangClientidBindRequest 绑定nick和客户端id
type TaobaoWangwangClientidBindRequest struct {
    
    // AppName required应用名
    AppName string `json:"app_name";xml:"app_name"`
    
    // ClientId required客户端Id
    ClientId string `json:"client_id";xml:"client_id"`
    
}

func (req *TaobaoWangwangClientidBindRequest) GetAPIName() string {
	return "taobao.wangwang.clientid.bind"
}

func (req *TaobaoWangwangClientidBindRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["AppName"] = req.AppName
    
    params["ClientId"] = req.ClientId
    
    return params
}

// TaobaoWangwangClientidBindResponse 绑定nick和客户端id
type TaobaoWangwangClientidBindResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic0:表示成功
    // 其它为错误
    Result int64 `json:"result";xml:"result"`
    
}

// TaobaoWangwangClientidUnbindRequest 解除账号和客户端Id的绑定
type TaobaoWangwangClientidUnbindRequest struct {
    
    // AppName required应用名
    AppName string `json:"app_name";xml:"app_name"`
    
    // ClientId required客户端Id
    ClientId string `json:"client_id";xml:"client_id"`
    
}

func (req *TaobaoWangwangClientidUnbindRequest) GetAPIName() string {
	return "taobao.wangwang.clientid.unbind"
}

func (req *TaobaoWangwangClientidUnbindRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["AppName"] = req.AppName
    
    params["ClientId"] = req.ClientId
    
    return params
}

// TaobaoWangwangClientidUnbindResponse 解除账号和客户端Id的绑定
type TaobaoWangwangClientidUnbindResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Basic返回0表示成功， 其他值为错误
    Result int64 `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsStockInOrderNotifyRequest 入库通知单
type TaobaoWlbWmsStockInOrderNotifyRequest struct {
    
    // ExpectEndTime optional预期送达结束时间
    ExpectEndTime string `json:"expect_end_time";xml:"expect_end_time"`
    
    // ExpectStartTime optional预期送达开始时间
    ExpectStartTime string `json:"expect_start_time";xml:"expect_start_time"`
    
    // ExtendFields optional扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    ExtendFields string `json:"extend_fields";xml:"extend_fields"`
    
    // InboundTypeDesc optional可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
    InboundTypeDesc string `json:"inbound_type_desc";xml:"inbound_type_desc"`
    
    // OrderCode required入库单据编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderCreateTime required单据创建时间
    OrderCreateTime time.Time `json:"order_create_time";xml:"order_create_time"`
    
    // OrderFlag optional订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
    OrderFlag string `json:"order_flag";xml:"order_flag"`
    
    // OrderItemList required系统自动生成
    OrderItemList Orderitemlistwlbwmsstockinordernotifywl `json:"order_item_list";xml:"order_item_list"`
    
    // OrderType required单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
    OrderType int64 `json:"order_type";xml:"order_type"`
    
    // PrevOrderCode optional来源单据号，如销售退货时填充原销售订单号
    PrevOrderCode string `json:"prev_order_code";xml:"prev_order_code"`
    
    // ReceiverInfo optional系统自动生成
    ReceiverInfo Receiverinfowlbwmsstockinordernotifywl `json:"receiver_info";xml:"receiver_info"`
    
    // Remark optional备注，销退入库订单需要留言备注填充到此字段
    Remark string `json:"remark";xml:"remark"`
    
    // ReturnReason optional销退时请提供退货的原因
    ReturnReason string `json:"return_reason";xml:"return_reason"`
    
    // SenderInfo optional系统自动生成
    SenderInfo Senderinfowlbwmsstockinordernotifywl `json:"sender_info";xml:"sender_info"`
    
    // StoreCode required仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // SupplierCode optional供应商编码，往来单位编码
    SupplierCode string `json:"supplier_code";xml:"supplier_code"`
    
    // SupplierName optional供应商名称 ，往来单位名称
    SupplierName string `json:"supplier_name";xml:"supplier_name"`
    
    // TmsOrderCode optional运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
    TmsOrderCode string `json:"tms_order_code";xml:"tms_order_code"`
    
    // TmsServiceCode optional配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
    TmsServiceCode string `json:"tms_service_code";xml:"tms_service_code"`
    
    // TmsServiceName optional快递公司名称
    TmsServiceName string `json:"tms_service_name";xml:"tms_service_name"`
    
}

func (req *TaobaoWlbWmsStockInOrderNotifyRequest) GetAPIName() string {
	return "taobao.wlb.wms.stock.in.order.notify"
}

func (req *TaobaoWlbWmsStockInOrderNotifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 20)
    
    params["ExpectEndTime"] = req.ExpectEndTime
    
    params["ExpectStartTime"] = req.ExpectStartTime
    
    params["ExtendFields"] = req.ExtendFields
    
    params["InboundTypeDesc"] = req.InboundTypeDesc
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderCreateTime"] = req.OrderCreateTime
    
    params["OrderFlag"] = req.OrderFlag
    
    params["OrderItemList"] = req.OrderItemList
    
    params["OrderType"] = req.OrderType
    
    params["PrevOrderCode"] = req.PrevOrderCode
    
    params["ReceiverInfo"] = req.ReceiverInfo
    
    params["Remark"] = req.Remark
    
    params["ReturnReason"] = req.ReturnReason
    
    params["SenderInfo"] = req.SenderInfo
    
    params["StoreCode"] = req.StoreCode
    
    params["SupplierCode"] = req.SupplierCode
    
    params["SupplierName"] = req.SupplierName
    
    params["TmsOrderCode"] = req.TmsOrderCode
    
    params["TmsServiceCode"] = req.TmsServiceCode
    
    params["TmsServiceName"] = req.TmsServiceName
    
    return params
}

// TaobaoWlbWmsStockInOrderNotifyResponse 入库通知单
type TaobaoWlbWmsStockInOrderNotifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OrderCode Basic仓储订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // WlErrorCode Basic错误编码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误详细
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoWlbWmsInventoryCountRequest 损益单回传
type TaobaoWlbWmsInventoryCountRequest struct {
    
    // Content optional损益单回传信息
    Content WlbWmsInventoryCount `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsInventoryCountRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.count"
}

func (req *TaobaoWlbWmsInventoryCountRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// TaobaoWlbWmsInventoryCountResponse 损益单回传
type TaobaoWlbWmsInventoryCountResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result WlbWmsInventoryCountResp `json:"result";xml:"result"`
    
}

// AlipaySystemOauthTokenRequest OAuth2.0 授权的第二步，换取访问令牌。
type AlipaySystemOauthTokenRequest struct {
    
    // Code optional授权码，用户对应用授权后得到。
    Code string `json:"code";xml:"code"`
    
    // GrantType required获取访问令牌的类型，authorization_code表示用授权码换，refresh_token表示用刷新令牌来换。
    GrantType string `json:"grant_type";xml:"grant_type"`
    
    // RefreshToken optional刷新令牌，上次换取访问令牌是得到。
    RefreshToken string `json:"refresh_token";xml:"refresh_token"`
    
}

func (req *AlipaySystemOauthTokenRequest) GetAPIName() string {
	return "alipay.system.oauth.token"
}

func (req *AlipaySystemOauthTokenRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Code"] = req.Code
    
    params["GrantType"] = req.GrantType
    
    params["RefreshToken"] = req.RefreshToken
    
    return params
}

// AlipaySystemOauthTokenResponse OAuth2.0 授权的第二步，换取访问令牌。
type AlipaySystemOauthTokenResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AccessToken Basic访问令牌
    AccessToken string `json:"access_token";xml:"access_token"`
    
    // AlipayUserId Basic支付宝用户的id。
    AlipayUserId string `json:"alipay_user_id";xml:"alipay_user_id"`
    
    // ExpiresIn Basic访问令牌的有效时间，单位是秒。
    ExpiresIn string `json:"expires_in";xml:"expires_in"`
    
    // ReExpiresIn Basic刷新令牌的有效时间，单位是秒。
    ReExpiresIn string `json:"re_expires_in";xml:"re_expires_in"`
    
    // RefreshToken Basic刷新令牌
    RefreshToken string `json:"refresh_token";xml:"refresh_token"`
    
}

// TaobaoOmniorderStoreDeliverconfigUpdateRequest 修改门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigUpdateRequest struct {
    
    // StoreDeliverConfig required卖家发货配置
    StoreDeliverConfig StoreDeliverConfig `json:"store_deliver_config";xml:"store_deliver_config"`
    
    // StoreId required门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetAPIName() string {
	return "taobao.omniorder.store.deliverconfig.update"
}

func (req *TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["StoreDeliverConfig"] = req.StoreDeliverConfig
    
    params["StoreId"] = req.StoreId
    
    return params
}

// TaobaoOmniorderStoreDeliverconfigUpdateResponse 修改门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// QimenTaobaoUopTobDeliveryorderConfirmRequest 2B订单出库确认回传
type QimenTaobaoUopTobDeliveryorderConfirmRequest struct {
    
    // CustomerId requiredcustomerId
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // Request optionaltob出库确认对象
    Request ToBDeliveryOrderConfirmRequest `json:"request";xml:"request"`
    
}

func (req *QimenTaobaoUopTobDeliveryorderConfirmRequest) GetAPIName() string {
	return "qimen.taobao.uop.tob.deliveryorder.confirm"
}

func (req *QimenTaobaoUopTobDeliveryorderConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["Request"] = req.Request
    
    return params
}

// QimenTaobaoUopTobDeliveryorderConfirmResponse 2B订单出库确认回传
type QimenTaobaoUopTobDeliveryorderConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoOmniorderStoreCollectconfigUpdateRequest 修改门店自提配置内容
type TaobaoOmniorderStoreCollectconfigUpdateRequest struct {
    
    // StoreCollectConfig required门店自提配置
    StoreCollectConfig StoreCollectConfig `json:"store_collect_config";xml:"store_collect_config"`
    
    // StoreId required门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreCollectconfigUpdateRequest) GetAPIName() string {
	return "taobao.omniorder.store.collectconfig.update"
}

func (req *TaobaoOmniorderStoreCollectconfigUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["StoreCollectConfig"] = req.StoreCollectConfig
    
    params["StoreId"] = req.StoreId
    
    return params
}

// TaobaoOmniorderStoreCollectconfigUpdateResponse 修改门店自提配置内容
type TaobaoOmniorderStoreCollectconfigUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoOmniorderStoreDeliverconfigGetRequest 查询门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigGetRequest struct {
    
    // Activity optional是否是活动期
    Activity bool `json:"activity";xml:"activity"`
    
    // StoreId required门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreDeliverconfigGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.deliverconfig.get"
}

func (req *TaobaoOmniorderStoreDeliverconfigGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Activity"] = req.Activity
    
    params["StoreId"] = req.StoreId
    
    return params
}

// TaobaoOmniorderStoreDeliverconfigGetResponse 查询门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// CainiaoCloudprintCustomaresGetRequest 供isv使用，获取商家的自定义区的模板信息
type CainiaoCloudprintCustomaresGetRequest struct {
    
    // TemplateId required用户使用的标准模板id
    TemplateId int64 `json:"template_id";xml:"template_id"`
    
}

func (req *CainiaoCloudprintCustomaresGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.customares.get"
}

func (req *CainiaoCloudprintCustomaresGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TemplateId"] = req.TemplateId
    
    return params
}

// CainiaoCloudprintCustomaresGetResponse 供isv使用，获取商家的自定义区的模板信息
type CainiaoCloudprintCustomaresGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object结果
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsInventoryLackUploadRequest 缺货通知
type TaobaoWlbWmsInventoryLackUploadRequest struct {
    
    // Content optional缺货通知信息
    Content WlbWmsInventoryLackUpload `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsInventoryLackUploadRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.lack.upload"
}

func (req *TaobaoWlbWmsInventoryLackUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// TaobaoWlbWmsInventoryLackUploadResponse 缺货通知
type TaobaoWlbWmsInventoryLackUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object缺货回告
    Result WlbWmsInventoryLackUploadResp `json:"result";xml:"result"`
    
}

// TaobaoOmniorderStoreCollectconfigGetRequest 查询门店自提配置内容
type TaobaoOmniorderStoreCollectconfigGetRequest struct {
    
    // Activity optional是否是活动期
    Activity bool `json:"activity";xml:"activity"`
    
    // StoreId required门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderStoreCollectconfigGetRequest) GetAPIName() string {
	return "taobao.omniorder.store.collectconfig.get"
}

func (req *TaobaoOmniorderStoreCollectconfigGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Activity"] = req.Activity
    
    params["StoreId"] = req.StoreId
    
    return params
}

// TaobaoOmniorderStoreCollectconfigGetResponse 查询门店自提配置内容
type TaobaoOmniorderStoreCollectconfigGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoUserSellerGetRequest 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaoUserSellerGetRequest struct {
    
    // Fields required需要返回的字段列表，可选值为返回示例值中的可以看到的字段
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoUserSellerGetRequest) GetAPIName() string {
	return "taobao.user.seller.get"
}

func (req *TaobaoUserSellerGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoUserSellerGetResponse 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaoUserSellerGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // User Object用户信息
    User User `json:"user";xml:"user"`
    
}

// TaobaoWlbWmsInventoryQueryRequest 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
type TaobaoWlbWmsInventoryQueryRequest struct {
    
    // BatchCode optional库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
    BatchCode string `json:"batch_code";xml:"batch_code"`
    
    // ChannelCode optional渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
    ChannelCode string `json:"channel_code";xml:"channel_code"`
    
    // DueDate optional失效日期，type=2时字段传值有效。
    DueDate time.Time `json:"due_date";xml:"due_date"`
    
    // InventoryType optional库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
    InventoryType int64 `json:"inventory_type";xml:"inventory_type"`
    
    // ItemId optional菜鸟商品ID
    ItemId string `json:"item_id";xml:"item_id"`
    
    // PageNo optional分页的第几页
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页多少条，最大50条
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ProduceDate optional生产日期，type=2时字段传值有效。
    ProduceDate time.Time `json:"produce_date";xml:"produce_date"`
    
    // StoreCode optional仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // Type optional库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbWmsInventoryQueryRequest) GetAPIName() string {
	return "taobao.wlb.wms.inventory.query"
}

func (req *TaobaoWlbWmsInventoryQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["BatchCode"] = req.BatchCode
    
    params["ChannelCode"] = req.ChannelCode
    
    params["DueDate"] = req.DueDate
    
    params["InventoryType"] = req.InventoryType
    
    params["ItemId"] = req.ItemId
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["ProduceDate"] = req.ProduceDate
    
    params["StoreCode"] = req.StoreCode
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoWlbWmsInventoryQueryResponse 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
type TaobaoWlbWmsInventoryQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemList Object Array商品详情列表
    ItemList WmsInventoryQueryItemlist `json:"item_list";xml:"item_list"`
    
    // TotalCount Basic总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
    // WlErrorCode Basic错误代码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess bool `json:"wl_success";xml:"wl_success"`
    
}

// AlibabaEinvoicePaperPrintRequest 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintRequest struct {
    
    // DialogSettingFlag required打印框设置，0=不弹打印设置框，1=弹出打印设置框
    DialogSettingFlag int64 `json:"dialog_setting_flag";xml:"dialog_setting_flag"`
    
    // ForcePrint required是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
    ForcePrint bool `json:"force_print";xml:"force_print"`
    
    // PayeeRegisterNo required销售方纳税人识别号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // PrintFlag required打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
    PrintFlag int64 `json:"print_flag";xml:"print_flag"`
    
    // SerialNo required开票流水号
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoicePaperPrintRequest) GetAPIName() string {
	return "alibaba.einvoice.paper.print"
}

func (req *AlibabaEinvoicePaperPrintRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["DialogSettingFlag"] = req.DialogSettingFlag
    
    params["ForcePrint"] = req.ForcePrint
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["PrintFlag"] = req.PrintFlag
    
    params["SerialNo"] = req.SerialNo
    
    return params
}

// AlibabaEinvoicePaperPrintResponse 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic调用结果，打印结果tmc消息alibaba_invoice_PaperOpsReturn异步通知
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoItemSellerGetRequest 获取单个商品的全部信息
type TaobaoItemSellerGetRequest struct {
    
    // Fields required需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。
    Fields string `json:"fields";xml:"fields"`
    
    // NumIid required商品数字ID
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemSellerGetRequest) GetAPIName() string {
	return "taobao.item.seller.get"
}

func (req *TaobaoItemSellerGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["NumIid"] = req.NumIid
    
    return params
}

// TaobaoItemSellerGetResponse 获取单个商品的全部信息
type TaobaoItemSellerGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object商品详细信息
    Item Item `json:"item";xml:"item"`
    
}

// AlibabaEinvoicePaperInvalidRequest 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaEinvoicePaperInvalidRequest struct {
    
    // InvalidOperator required作废操作人
    InvalidOperator string `json:"invalid_operator";xml:"invalid_operator"`
    
    // InvalidType required作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
    InvalidType int64 `json:"invalid_type";xml:"invalid_type"`
    
    // InvoiceCode optional发票代码，空白作废时必填
    InvoiceCode string `json:"invoice_code";xml:"invoice_code"`
    
    // InvoiceNo optional发票号码，空白作废时必填
    InvoiceNo string `json:"invoice_no";xml:"invoice_no"`
    
    // PayeeRegisterNo required销售方纳税人识别号
    PayeeRegisterNo string `json:"payee_register_no";xml:"payee_register_no"`
    
    // SerialNo required开票流水号
    SerialNo string `json:"serial_no";xml:"serial_no"`
    
}

func (req *AlibabaEinvoicePaperInvalidRequest) GetAPIName() string {
	return "alibaba.einvoice.paper.invalid"
}

func (req *AlibabaEinvoicePaperInvalidRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["InvalidOperator"] = req.InvalidOperator
    
    params["InvalidType"] = req.InvalidType
    
    params["InvoiceCode"] = req.InvoiceCode
    
    params["InvoiceNo"] = req.InvoiceNo
    
    params["PayeeRegisterNo"] = req.PayeeRegisterNo
    
    params["SerialNo"] = req.SerialNo
    
    return params
}

// AlibabaEinvoicePaperInvalidResponse 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaEinvoicePaperInvalidResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic接口调用是否成功，操作结果tmc异步返回alibaba_invoice_PaperOpsReturn
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoItemsSellerListGetRequest 批量获取商品详细信息
type TaobaoItemsSellerListGetRequest struct {
    
    // Fields required需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。
    Fields string `json:"fields";xml:"fields"`
    
    // NumIids required商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。
    NumIids string `json:"num_iids";xml:"num_iids"`
    
}

func (req *TaobaoItemsSellerListGetRequest) GetAPIName() string {
	return "taobao.items.seller.list.get"
}

func (req *TaobaoItemsSellerListGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["NumIids"] = req.NumIids
    
    return params
}

// TaobaoItemsSellerListGetResponse 批量获取商品详细信息
type TaobaoItemsSellerListGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Items Object Array商品详细信息列表
    Items Item `json:"items";xml:"items"`
    
}

// TaobaoQimenSnReportRequest WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
type TaobaoQimenSnReportRequest struct {
    
    // CurrentPage required当前页(从1开始)
    CurrentPage int64 `json:"currentPage";xml:"currentPage"`
    
    // DeliveryOrder optional发货单信息
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional商品列表
    Items Item `json:"items";xml:"items"`
    
    // PageSize required每页记录的条数
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // TotalPage required总页数
    TotalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

func (req *TaobaoQimenSnReportRequest) GetAPIName() string {
	return "taobao.qimen.sn.report"
}

func (req *TaobaoQimenSnReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["CurrentPage"] = req.CurrentPage
    
    params["DeliveryOrder"] = req.DeliveryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["PageSize"] = req.PageSize
    
    params["TotalPage"] = req.TotalPage
    
    return params
}

// TaobaoQimenSnReportResponse WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
type TaobaoQimenSnReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoTaeBookBillGetRequest tae查询单笔虚拟账户明细
type TaobaoTaeBookBillGetRequest struct {
    
    // AccountId required虚拟账户科目编号
    AccountId int64 `json:"account_id";xml:"account_id"`
    
    // Bid optional虚拟账户流水编号
    Bid int64 `json:"bid";xml:"bid"`
    
    // Fields required需要返回的字段:参见BookBill结构体
    Fields string `json:"fields";xml:"fields"`
    
    // Id required虚拟账户流水编号
    Id int64 `json:"id";xml:"id"`
    
}

func (req *TaobaoTaeBookBillGetRequest) GetAPIName() string {
	return "taobao.tae.book.bill.get"
}

func (req *TaobaoTaeBookBillGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["AccountId"] = req.AccountId
    
    params["Bid"] = req.Bid
    
    params["Fields"] = req.Fields
    
    params["Id"] = req.Id
    
    return params
}

// TaobaoTaeBookBillGetResponse tae查询单笔虚拟账户明细
type TaobaoTaeBookBillGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Bookbill Object虚拟账户账单
    Bookbill TopAcctCashJourDto `json:"bookbill";xml:"bookbill"`
    
}

// TaobaoTaeBookBillsGetRequest tae查询虚拟账户明细数据
type TaobaoTaeBookBillsGetRequest struct {
    
    // AccountId optional虚拟账户科目编号
    AccountId int64 `json:"account_id";xml:"account_id"`
    
    // EndTime required记账结束时间,end_time与start_time相差不能超过30天
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // Fields required需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
    Fields string `json:"fields";xml:"fields"`
    
    // JournalTypes optional明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
    JournalTypes int64 `json:"journal_types";xml:"journal_types"`
    
    // PageNo optional页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页大小,建议40~100,不能超过100
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime required记账开始时间
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoTaeBookBillsGetRequest) GetAPIName() string {
	return "taobao.tae.book.bills.get"
}

func (req *TaobaoTaeBookBillsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["AccountId"] = req.AccountId
    
    params["EndTime"] = req.EndTime
    
    params["Fields"] = req.Fields
    
    params["JournalTypes"] = req.JournalTypes
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    return params
}

// TaobaoTaeBookBillsGetResponse tae查询虚拟账户明细数据
type TaobaoTaeBookBillsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Bills Object Array虚拟账户账单列表
    Bills TopAcctCashJourDto `json:"bills";xml:"bills"`
    
    // HasNext Basic是否有下一页
    HasNext bool `json:"has_next";xml:"has_next"`
    
    // TotalResults Basic当前查询的结果数,非总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoWlbOrderJzQueryRequest 家装业务查询物流公司api
type TaobaoWlbOrderJzQueryRequest struct {
    
    // InsJzReceiverTO optional家装安装服务收货人信息
    InsJzReceiverTO JzReceiverTo `json:"ins_jz_receiver_t_o";xml:"ins_jz_receiver_t_o"`
    
    // JzReceiverTo optional家装收货人信息
    JzReceiverTo JzReceiverTo `json:"jz_receiver_to";xml:"jz_receiver_to"`
    
    // SenderId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    SenderId int64 `json:"sender_id";xml:"sender_id"`
    
    // Tid optional交易id
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoWlbOrderJzQueryRequest) GetAPIName() string {
	return "taobao.wlb.order.jz.query"
}

func (req *TaobaoWlbOrderJzQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["InsJzReceiverTO"] = req.InsJzReceiverTO
    
    params["JzReceiverTo"] = req.JzReceiverTo
    
    params["SenderId"] = req.SenderId
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoWlbOrderJzQueryResponse 家装业务查询物流公司api
type TaobaoWlbOrderJzQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object结果信息
    Result JzTopDto `json:"result";xml:"result"`
    
    // ResultErrorCode Basic错误编码
    ResultErrorCode string `json:"result_error_code";xml:"result_error_code"`
    
    // ResultErrorMsg Basic错误信息
    ResultErrorMsg string `json:"result_error_msg";xml:"result_error_msg"`
    
    // ResultSuccess Basic是否成功
    ResultSuccess bool `json:"result_success";xml:"result_success"`
    
}

// TaobaoRdcAligeniusWarehouseResendUpdateRequest 补发单状态回传接口
type TaobaoRdcAligeniusWarehouseResendUpdateRequest struct {
    
    // Param0 required参数
    Param0 UpdateResendStatusDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.resend.update"
}

func (req *TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param0"] = req.Param0
    
    return params
}

// TaobaoRdcAligeniusWarehouseResendUpdateResponse 补发单状态回传接口
type TaobaoRdcAligeniusWarehouseResendUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoQimenOrderCancelRequest ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
type TaobaoQimenOrderCancelRequest struct {
    
    // CancelReason optional取消原因
    CancelReason string `json:"cancelReason";xml:"cancelReason"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderCode required单据编码
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderId optional仓储系统单据编码
    OrderId string `json:"orderId";xml:"orderId"`
    
    // OrderType optional单据类型(JYCK=一般交易出库单;HHCK= 换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK=其他入库;XTRK= 销退入库;THRK=退货入库;HHRK= 换货入库;CNJG= 仓内加工单;CGTH=采购退货出库单)
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OwnerCode optional货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // WarehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenOrderCancelRequest) GetAPIName() string {
	return "taobao.qimen.order.cancel"
}

func (req *TaobaoQimenOrderCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["CancelReason"] = req.CancelReason
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderId"] = req.OrderId
    
    params["OrderType"] = req.OrderType
    
    params["OwnerCode"] = req.OwnerCode
    
    params["Remark"] = req.Remark
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenOrderCancelResponse ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
type TaobaoQimenOrderCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenDeliveryorderConfirmRequest taobao.qimen.deliveryorder.confirm
type TaobaoQimenDeliveryorderConfirmRequest struct {
    
    // DeliveryOrder optional发货单信息
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderLines optional单据列表
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // Packages optional包裹信息
    Packages Package `json:"packages";xml:"packages"`
    
}

func (req *TaobaoQimenDeliveryorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.confirm"
}

func (req *TaobaoQimenDeliveryorderConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["DeliveryOrder"] = req.DeliveryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderLines"] = req.OrderLines
    
    params["Packages"] = req.Packages
    
    return params
}

// TaobaoQimenDeliveryorderConfirmResponse taobao.qimen.deliveryorder.confirm
type TaobaoQimenDeliveryorderConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenDeliveryorderCreateRequest taobao.qimen.deliveryorder.create
type TaobaoQimenDeliveryorderCreateRequest struct {
    
    // DeliveryOrder optional发货单信息
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional
    Items Item `json:"items";xml:"items"`
    
    // OrderLines optional订单列表
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenDeliveryorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.create"
}

func (req *TaobaoQimenDeliveryorderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["DeliveryOrder"] = req.DeliveryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OrderLines"] = req.OrderLines
    
    return params
}

// TaobaoQimenDeliveryorderCreateResponse taobao.qimen.deliveryorder.create
type TaobaoQimenDeliveryorderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // CreateTime Basic订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime string `json:"createTime";xml:"createTime"`
    
    // DeliveryOrderId Basic出库单仓储系统编码
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // DeliveryOrders Object Array发货单信息
    DeliveryOrders DeliveryOrder `json:"deliveryOrders";xml:"deliveryOrders"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // LogisticsCode Basic物流公司编码(统仓统配使用)
    LogisticsCode string `json:"logisticsCode";xml:"logisticsCode"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // WarehouseCode Basic仓库编码(统仓统配使用)
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest 补发单erp物流信息回传平台
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest struct {
    
    // Param0 required参数
    Param0 SendResendLogisticsMsgDto `json:"param0";xml:"param0"`
    
}

func (req *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest) GetAPIName() string {
	return "taobao.rdc.aligenius.warehouse.resend.logistics.msg.post"
}

func (req *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Param0"] = req.Param0
    
    return params
}

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse 补发单erp物流信息回传平台
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoQimenOrderprocessReportRequest WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
type TaobaoQimenOrderprocessReportRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Order optional订单信息
    Order Order `json:"order";xml:"order"`
    
    // Process optional订单处理信息
    Process Process `json:"process";xml:"process"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenOrderprocessReportRequest) GetAPIName() string {
	return "taobao.qimen.orderprocess.report"
}

func (req *TaobaoQimenOrderprocessReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Order"] = req.Order
    
    params["Process"] = req.Process
    
    params["Remark"] = req.Remark
    
    return params
}

// TaobaoQimenOrderprocessReportResponse WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
type TaobaoQimenOrderprocessReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoFenxiaoDistributorItemsGetRequest 供应商查询分销商商品下载记录。
type TaobaoFenxiaoDistributorItemsGetRequest struct {
    
    // DistributorId special分销商ID 。
    DistributorId int64 `json:"distributor_id";xml:"distributor_id"`
    
    // EndModified optional设置结束时间,空为不设置。
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // PageNo optional页码（大于0的整数，默认1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页记录数（默认20，最大50）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ProductId special产品ID
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // StartModified optional设置开始时间。空为不设置。
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoFenxiaoDistributorItemsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributor.items.get"
}

func (req *TaobaoFenxiaoDistributorItemsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["DistributorId"] = req.DistributorId
    
    params["EndModified"] = req.EndModified
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["ProductId"] = req.ProductId
    
    params["StartModified"] = req.StartModified
    
    return params
}

// TaobaoFenxiaoDistributorItemsGetResponse 供应商查询分销商商品下载记录。
type TaobaoFenxiaoDistributorItemsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Records Object Array下载记录对象
    Records FenxiaoItemRecord `json:"records";xml:"records"`
    
    // TotalResults Basic查询结果记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoFenxiaoOrderConfirmPaidRequest 供应商确认收款（非支付宝交易）。
type TaobaoFenxiaoOrderConfirmPaidRequest struct {
    
    // ConfirmRemark optional确认支付信息（字数小于100）
    ConfirmRemark string `json:"confirm_remark";xml:"confirm_remark"`
    
    // PurchaseOrderId required采购单编号。
    PurchaseOrderId int64 `json:"purchase_order_id";xml:"purchase_order_id"`
    
}

func (req *TaobaoFenxiaoOrderConfirmPaidRequest) GetAPIName() string {
	return "taobao.fenxiao.order.confirm.paid"
}

func (req *TaobaoFenxiaoOrderConfirmPaidRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ConfirmRemark"] = req.ConfirmRemark
    
    params["PurchaseOrderId"] = req.PurchaseOrderId
    
    return params
}

// TaobaoFenxiaoOrderConfirmPaidResponse 供应商确认收款（非支付宝交易）。
type TaobaoFenxiaoOrderConfirmPaidResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic确认结果成功与否
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoCooperationGetRequest 获取供应商的合作关系信息
type TaobaoFenxiaoCooperationGetRequest struct {
    
    // EndDate optional合作结束时间yyyy-MM-dd HH:mm:ss
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // PageNo optional页码（大于0的整数，默认1）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页记录数（默认20，最大50）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartDate optional合作开始时间yyyy-MM-dd HH:mm:ss
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // Status optional合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
    Status string `json:"status";xml:"status"`
    
    // TradeType optional分销方式：AGENT(代销) 、DEALER（经销）
    TradeType string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoCooperationGetRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.get"
}

func (req *TaobaoFenxiaoCooperationGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["EndDate"] = req.EndDate
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartDate"] = req.StartDate
    
    params["Status"] = req.Status
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoCooperationGetResponse 获取供应商的合作关系信息
type TaobaoFenxiaoCooperationGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Cooperations Object Array合作分销关系
    Cooperations Cooperation `json:"cooperations";xml:"cooperations"`
    
    // TotalResults Basic结果记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoQianniuCloudkefuStatuslogGetRequest 查询客服账号的操作记录。如：登录，下线，挂起等。接口即将下线，请ISV切换到taobao.qianniu.cloudkefu.onlinestatuslog.get接口上
type TaobaoQianniuCloudkefuStatuslogGetRequest struct {
    
    // AccountIds required子帐号列表，最多10个
    AccountIds int64 `json:"account_ids";xml:"account_ids"`
    
    // Domian optional域，淘宝:cntaobao，1688:cnalichn
    Domian string `json:"domian";xml:"domian"`
    
    // EndTime optional查询结束时间，默认当天24点
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // MainAccountId required主帐号ID
    MainAccountId int64 `json:"main_account_id";xml:"main_account_id"`
    
    // PageNum optional页码，每页50个
    PageNum int64 `json:"page_num";xml:"page_num"`
    
    // PageSize optional分页
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime optional查询开始时间，默认当天0点
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
    // Type required记录类型，PC上下线：8；移动上下线：4；挂起类型：5
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoQianniuCloudkefuStatuslogGetRequest) GetAPIName() string {
	return "taobao.qianniu.cloudkefu.statuslog.get"
}

func (req *TaobaoQianniuCloudkefuStatuslogGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["AccountIds"] = req.AccountIds
    
    params["Domian"] = req.Domian
    
    params["EndTime"] = req.EndTime
    
    params["MainAccountId"] = req.MainAccountId
    
    params["PageNum"] = req.PageNum
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoQianniuCloudkefuStatuslogGetResponse 查询客服账号的操作记录。如：登录，下线，挂起等。接口即将下线，请ISV切换到taobao.qianniu.cloudkefu.onlinestatuslog.get接口上
type TaobaoQianniuCloudkefuStatuslogGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result ResultTo `json:"result";xml:"result"`
    
}

// TaobaoLogisticsOnlineConfirmRequest <br><font color='red'>仅使用taobao.logistics.online.send 发货时，未输入运单号的情况下，需要使用该接口确认发货。发货接口主要有taobao.logistics.offline.send 和taobao.logistics.online.send <br>
//确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外
type TaobaoLogisticsOnlineConfirmRequest struct {
    
    // IsSplit optional表明是否是拆单，默认值0，1表示拆单
    IsSplit int64 `json:"is_split";xml:"is_split"`
    
    // OutSid required运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    OutSid string `json:"out_sid";xml:"out_sid"`
    
    // SellerIp optional商家的IP地址
    SellerIp string `json:"seller_ip";xml:"seller_ip"`
    
    // SubTid optional拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
    SubTid int64 `json:"sub_tid";xml:"sub_tid"`
    
    // Tid required淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOnlineConfirmRequest) GetAPIName() string {
	return "taobao.logistics.online.confirm"
}

func (req *TaobaoLogisticsOnlineConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["IsSplit"] = req.IsSplit
    
    params["OutSid"] = req.OutSid
    
    params["SellerIp"] = req.SellerIp
    
    params["SubTid"] = req.SubTid
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsOnlineConfirmResponse <br><font color='red'>仅使用taobao.logistics.online.send 发货时，未输入运单号的情况下，需要使用该接口确认发货。发货接口主要有taobao.logistics.offline.send 和taobao.logistics.online.send <br>
//确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外
type TaobaoLogisticsOnlineConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shipping Object只返回is_success：是否成功。
    Shipping Shipping `json:"shipping";xml:"shipping"`
    
}

// TaobaoWlbOrderJzConsignRequest 家装类订单使用该接口发货
type TaobaoWlbOrderJzConsignRequest struct {
    
    // InsReceiverTo optional安装收货人信息,如果为空,则取默认收货人信息
    InsReceiverTo JzReceiverTo `json:"ins_receiver_to";xml:"ins_receiver_to"`
    
    // InsTpDto optional安装公司信息,需要安装时,才填写
    InsTpDto Tpdto `json:"ins_tp_dto";xml:"ins_tp_dto"`
    
    // JzReceiverTo optional家装收货人信息,如果为空,则取默认收货信息
    JzReceiverTo JzReceiverTo `json:"jz_receiver_to";xml:"jz_receiver_to"`
    
    // JzTopArgs optional发货参数
    JzTopArgs JzTopArgs `json:"jz_top_args";xml:"jz_top_args"`
    
    // LgTpDto required物流公司信息
    LgTpDto Tpdto `json:"lg_tp_dto";xml:"lg_tp_dto"`
    
    // SenderId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    SenderId int64 `json:"sender_id";xml:"sender_id"`
    
    // Tid required交易号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoWlbOrderJzConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.jz.consign"
}

func (req *TaobaoWlbOrderJzConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["InsReceiverTo"] = req.InsReceiverTo
    
    params["InsTpDto"] = req.InsTpDto
    
    params["JzReceiverTo"] = req.JzReceiverTo
    
    params["JzTopArgs"] = req.JzTopArgs
    
    params["LgTpDto"] = req.LgTpDto
    
    params["SenderId"] = req.SenderId
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoWlbOrderJzConsignResponse 家装类订单使用该接口发货
type TaobaoWlbOrderJzConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ResultErrorCode Basic错误码
    ResultErrorCode string `json:"result_error_code";xml:"result_error_code"`
    
    // ResultErrorMsg Basic错误信息描述
    ResultErrorMsg string `json:"result_error_msg";xml:"result_error_msg"`
    
    // ResultSuccess Basic是否成功
    ResultSuccess bool `json:"result_success";xml:"result_success"`
    
}

// TaobaoLogisticsOnlineCancelRequest 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
type TaobaoLogisticsOnlineCancelRequest struct {
    
    // Tid required淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOnlineCancelRequest) GetAPIName() string {
	return "taobao.logistics.online.cancel"
}

func (req *TaobaoLogisticsOnlineCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsOnlineCancelResponse 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
type TaobaoLogisticsOnlineCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic成功与失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // ModifyTime Basic修改时间
    ModifyTime string `json:"modify_time";xml:"modify_time"`
    
    // RecreatedOrderId Basic重新创建物流订单id
    RecreatedOrderId string `json:"recreated_order_id";xml:"recreated_order_id"`
    
}

// AlibabaEinvoiceApplyTestRequest 开票申请消息测试接口
type AlibabaEinvoiceApplyTestRequest struct {
    
    // BusinessType required0=个人，1=企业
    BusinessType int64 `json:"business_type";xml:"business_type"`
    
    // PayerName required买家抬头
    PayerName string `json:"payer_name";xml:"payer_name"`
    
    // PayerRegisterNo optional买家税号
    PayerRegisterNo string `json:"payer_register_no";xml:"payer_register_no"`
    
    // PlatformTid required订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
}

func (req *AlibabaEinvoiceApplyTestRequest) GetAPIName() string {
	return "alibaba.einvoice.apply.test"
}

func (req *AlibabaEinvoiceApplyTestRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["BusinessType"] = req.BusinessType
    
    params["PayerName"] = req.PayerName
    
    params["PayerRegisterNo"] = req.PayerRegisterNo
    
    params["PlatformTid"] = req.PlatformTid
    
    return params
}

// AlibabaEinvoiceApplyTestResponse 开票申请消息测试接口
type AlibabaEinvoiceApplyTestResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic消息是否发送成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoLogisticsDummySendRequest 用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货
type TaobaoLogisticsDummySendRequest struct {
    
    // Feature optionalfeature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。
    Feature string `json:"feature";xml:"feature"`
    
    // SellerIp optional商家的IP地址
    SellerIp string `json:"seller_ip";xml:"seller_ip"`
    
    // Tid required淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsDummySendRequest) GetAPIName() string {
	return "taobao.logistics.dummy.send"
}

func (req *TaobaoLogisticsDummySendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["Feature"] = req.Feature
    
    params["SellerIp"] = req.SellerIp
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsDummySendResponse 用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货
type TaobaoLogisticsDummySendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shipping Object返回发货是否成功is_success
    Shipping Shipping `json:"shipping";xml:"shipping"`
    
}

// TaobaoLogisticsOfflineSendRequest 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOfflineSendRequest struct {
    
    // CancelId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
    CancelId int64 `json:"cancel_id";xml:"cancel_id"`
    
    // CompanyCode required物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
    CompanyCode string `json:"company_code";xml:"company_code"`
    
    // Feature optionalfeature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
    Feature string `json:"feature";xml:"feature"`
    
    // IsSplit optional表明是否是拆单，默认值0，1表示拆单
    IsSplit int64 `json:"is_split";xml:"is_split"`
    
    // OutSid required运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    OutSid string `json:"out_sid";xml:"out_sid"`
    
    // SellerIp optional商家的IP地址
    SellerIp string `json:"seller_ip";xml:"seller_ip"`
    
    // SenderId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    SenderId int64 `json:"sender_id";xml:"sender_id"`
    
    // SubTid optional需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
    SubTid int64 `json:"sub_tid";xml:"sub_tid"`
    
    // Tid required淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOfflineSendRequest) GetAPIName() string {
	return "taobao.logistics.offline.send"
}

func (req *TaobaoLogisticsOfflineSendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["CancelId"] = req.CancelId
    
    params["CompanyCode"] = req.CompanyCode
    
    params["Feature"] = req.Feature
    
    params["IsSplit"] = req.IsSplit
    
    params["OutSid"] = req.OutSid
    
    params["SellerIp"] = req.SellerIp
    
    params["SenderId"] = req.SenderId
    
    params["SubTid"] = req.SubTid
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsOfflineSendResponse 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOfflineSendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shipping Object自己联系的调用结果
    Shipping Shipping `json:"shipping";xml:"shipping"`
    
}

// TaobaoQimenReturnorderCreateRequest ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
type TaobaoQimenReturnorderCreateRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // ItemList optionalitemList
    ItemList Item `json:"itemList";xml:"itemList"`
    
    // OrderLines optional订单信息
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // ReturnOrder optional退货单信息
    ReturnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

func (req *TaobaoQimenReturnorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.create"
}

func (req *TaobaoQimenReturnorderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["ItemList"] = req.ItemList
    
    params["OrderLines"] = req.OrderLines
    
    params["ReturnOrder"] = req.ReturnOrder
    
    return params
}

// TaobaoQimenReturnorderCreateResponse ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
type TaobaoQimenReturnorderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // ReturnOrderId Basic仓储系统退货单编码
    ReturnOrderId string `json:"returnOrderId";xml:"returnOrderId"`
    
}

// TaobaoWlbStoresBaseinfoGetRequest 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaoWlbStoresBaseinfoGetRequest struct {
    
    // Type optional0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbStoresBaseinfoGetRequest) GetAPIName() string {
	return "taobao.wlb.stores.baseinfo.get"
}

func (req *TaobaoWlbStoresBaseinfoGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoWlbStoresBaseinfoGetResponse 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaoWlbStoresBaseinfoGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // StoreInfoList Object Array仓库列表
    StoreInfoList StoreInfo `json:"store_info_list";xml:"store_info_list"`
    
    // TotalCount Basic返回的总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoQimenEntryorderCreateRequest ERP调用接口，创建入库单;
type TaobaoQimenEntryorderCreateRequest struct {
    
    // EntryOrder optional入库单信息
    EntryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional商品信息
    Items Item `json:"items";xml:"items"`
    
    // OrderLines optional入库单详情
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenEntryorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.create"
}

func (req *TaobaoQimenEntryorderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["EntryOrder"] = req.EntryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OrderLines"] = req.OrderLines
    
    return params
}

// TaobaoQimenEntryorderCreateResponse ERP调用接口，创建入库单;
type TaobaoQimenEntryorderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // EntryOrderId Basic仓储系统入库单编码
    EntryOrderId string `json:"entryOrderId";xml:"entryOrderId"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenReturnorderConfirmRequest taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // OrderLines optional订单信息
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
    // ReturnOrder optional退货单信息
    ReturnOrder ReturnOrder `json:"returnOrder";xml:"returnOrder"`
    
}

func (req *TaobaoQimenReturnorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.returnorder.confirm"
}

func (req *TaobaoQimenReturnorderConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["OrderLines"] = req.OrderLines
    
    params["ReturnOrder"] = req.ReturnOrder
    
    return params
}

// TaobaoQimenReturnorderConfirmResponse taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenStockchangeReportRequest WMS调用奇门的接口,将库存异动信息信息回传给ERP
type TaobaoQimenStockchangeReportRequest struct {
    
    // CurrentPage optional奇门仓储字段,说明,string(50),,
    CurrentPage string `json:"currentPage";xml:"currentPage"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optionalitem
    Items Item `json:"items";xml:"items"`
    
    // OrderCode optional奇门仓储字段,说明,string(50),,
    OrderCode string `json:"orderCode";xml:"orderCode"`
    
    // OrderType optional奇门仓储字段,说明,string(50),,
    OrderType string `json:"orderType";xml:"orderType"`
    
    // OwnerCode optional奇门仓储字段,说明,string(50),,
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // PageSize optional奇门仓储字段,说明,string(50),,
    PageSize string `json:"pageSize";xml:"pageSize"`
    
    // Remark optional奇门仓储字段,说明,string(50),,
    Remark string `json:"remark";xml:"remark"`
    
    // TotalPage optional奇门仓储字段,说明,string(50),,
    TotalPage string `json:"totalPage";xml:"totalPage"`
    
    // WarehouseCode optional奇门仓储字段,说明,string(50),,
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenStockchangeReportRequest) GetAPIName() string {
	return "taobao.qimen.stockchange.report"
}

func (req *TaobaoQimenStockchangeReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["CurrentPage"] = req.CurrentPage
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderType"] = req.OrderType
    
    params["OwnerCode"] = req.OwnerCode
    
    params["PageSize"] = req.PageSize
    
    params["Remark"] = req.Remark
    
    params["TotalPage"] = req.TotalPage
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenStockchangeReportResponse WMS调用奇门的接口,将库存异动信息信息回传给ERP
type TaobaoQimenStockchangeReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenInventoryQueryRequest ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenInventoryQueryRequest struct {
    
    // CriteriaList optional查询准则
    CriteriaList Criteria `json:"criteriaList";xml:"criteriaList"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
}

func (req *TaobaoQimenInventoryQueryRequest) GetAPIName() string {
	return "taobao.qimen.inventory.query"
}

func (req *TaobaoQimenInventoryQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CriteriaList"] = req.CriteriaList
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Remark"] = req.Remark
    
    return params
}

// TaobaoQimenInventoryQueryResponse ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenInventoryQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Items Object Array商品的库存信息列表
    Items Item `json:"items";xml:"items"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenSingleitemSynchronizeRequest ERP调用奇门的接口,同步商品信息给WMS
type TaobaoQimenSingleitemSynchronizeRequest struct {
    
    // ActionType required操作类型(两种类型：add|update)
    ActionType string `json:"actionType";xml:"actionType"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Item optional商品信息
    Item Item `json:"item";xml:"item"`
    
    // OwnerCode required货主编码
    OwnerCode string `json:"ownerCode";xml:"ownerCode"`
    
    // SupplierCode optional供应商编码
    SupplierCode string `json:"supplierCode";xml:"supplierCode"`
    
    // SupplierName optional供应商名称
    SupplierName string `json:"supplierName";xml:"supplierName"`
    
    // WarehouseCode required仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode string `json:"warehouseCode";xml:"warehouseCode"`
    
}

func (req *TaobaoQimenSingleitemSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.singleitem.synchronize"
}

func (req *TaobaoQimenSingleitemSynchronizeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["ActionType"] = req.ActionType
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Item"] = req.Item
    
    params["OwnerCode"] = req.OwnerCode
    
    params["SupplierCode"] = req.SupplierCode
    
    params["SupplierName"] = req.SupplierName
    
    params["WarehouseCode"] = req.WarehouseCode
    
    return params
}

// TaobaoQimenSingleitemSynchronizeResponse ERP调用奇门的接口,同步商品信息给WMS
type TaobaoQimenSingleitemSynchronizeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // ItemId Basic仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传)
    ItemId string `json:"itemId";xml:"itemId"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoItemTemplatesGetRequest 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoItemTemplatesGetRequest struct {
    
}

func (req *TaobaoItemTemplatesGetRequest) GetAPIName() string {
	return "taobao.item.templates.get"
}

func (req *TaobaoItemTemplatesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoItemTemplatesGetResponse 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoItemTemplatesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemTemplateList Object Array返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店）
    ItemTemplateList ItemTemplate `json:"item_template_list";xml:"item_template_list"`
    
}

// TaobaoQimenDeliveryorderBatchconfirmRequest taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Orders optional发货单列表
    Orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchconfirmRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchconfirm"
}

func (req *TaobaoQimenDeliveryorderBatchconfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Orders"] = req.Orders
    
    return params
}

// TaobaoQimenDeliveryorderBatchconfirmResponse taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenStockoutCreateRequest ERP调用奇门接口，创建出库单信息
type TaobaoQimenStockoutCreateRequest struct {
    
    // DeliveryOrder optional出库单信息
    DeliveryOrder DeliveryOrder `json:"deliveryOrder";xml:"deliveryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional
    Items Item `json:"items";xml:"items"`
    
    // OrderLines optional单据信息
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenStockoutCreateRequest) GetAPIName() string {
	return "taobao.qimen.stockout.create"
}

func (req *TaobaoQimenStockoutCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["DeliveryOrder"] = req.DeliveryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OrderLines"] = req.OrderLines
    
    return params
}

// TaobaoQimenStockoutCreateResponse ERP调用奇门接口，创建出库单信息
type TaobaoQimenStockoutCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // CreateTime Basic订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime string `json:"createTime";xml:"createTime"`
    
    // DeliveryOrderId Basic出库单仓储系统编码
    DeliveryOrderId string `json:"deliveryOrderId";xml:"deliveryOrderId"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenEntryorderConfirmRequest WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmRequest struct {
    
    // EntryOrder optional入库单信息
    EntryOrder EntryOrder `json:"entryOrder";xml:"entryOrder"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional商品信息列表
    Items Item `json:"items";xml:"items"`
    
    // OrderLines optional订单信息
    OrderLines OrderLine `json:"orderLines";xml:"orderLines"`
    
}

func (req *TaobaoQimenEntryorderConfirmRequest) GetAPIName() string {
	return "taobao.qimen.entryorder.confirm"
}

func (req *TaobaoQimenEntryorderConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["EntryOrder"] = req.EntryOrder
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["OrderLines"] = req.OrderLines
    
    return params
}

// TaobaoQimenEntryorderConfirmResponse WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// TaobaoQimenDeliveryorderBatchcreateRequest ERP调用接口，将发货信息批量推送给WMS
type TaobaoQimenDeliveryorderBatchcreateRequest struct {
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Orders optional订单信息
    Orders Order `json:"orders";xml:"orders"`
    
}

func (req *TaobaoQimenDeliveryorderBatchcreateRequest) GetAPIName() string {
	return "taobao.qimen.deliveryorder.batchcreate"
}

func (req *TaobaoQimenDeliveryorderBatchcreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Orders"] = req.Orders
    
    return params
}

// TaobaoQimenDeliveryorderBatchcreateResponse ERP调用接口，将发货信息批量推送给WMS
type TaobaoQimenDeliveryorderBatchcreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
    // Orders Object Array订单详情
    Orders Order `json:"orders";xml:"orders"`
    
}

// TaobaoWlbImportThreeplResourceGetRequest 获取3pl直邮的发货可用资源
type TaobaoWlbImportThreeplResourceGetRequest struct {
    
    // FromId required发货地区域id
    FromId int64 `json:"from_id";xml:"from_id"`
    
    // ToAddress required收件人地址
    ToAddress AddressDto `json:"to_address";xml:"to_address"`
    
    // Type requiredONLINE或者OFFLINE
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoWlbImportThreeplResourceGetRequest) GetAPIName() string {
	return "taobao.wlb.import.threepl.resource.get"
}

func (req *TaobaoWlbImportThreeplResourceGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["FromId"] = req.FromId
    
    params["ToAddress"] = req.ToAddress
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoWlbImportThreeplResourceGetResponse 获取3pl直邮的发货可用资源
type TaobaoWlbImportThreeplResourceGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result TopResult `json:"result";xml:"result"`
    
}

// TaobaoWlbImportThreeplOfflineConsignRequest 菜鸟认证直邮线下发货
type TaobaoWlbImportThreeplOfflineConsignRequest struct {
    
    // FromId optional发件人地址库id
    FromId int64 `json:"from_id";xml:"from_id"`
    
    // ResCode optional资源code
    ResCode string `json:"res_code";xml:"res_code"`
    
    // ResId optional资源id
    ResId int64 `json:"res_id";xml:"res_id"`
    
    // TradeId optional交易单号
    TradeId int64 `json:"trade_id";xml:"trade_id"`
    
    // WaybillNo optional运单号
    WaybillNo string `json:"waybill_no";xml:"waybill_no"`
    
}

func (req *TaobaoWlbImportThreeplOfflineConsignRequest) GetAPIName() string {
	return "taobao.wlb.import.threepl.offline.consign"
}

func (req *TaobaoWlbImportThreeplOfflineConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["FromId"] = req.FromId
    
    params["ResCode"] = req.ResCode
    
    params["ResId"] = req.ResId
    
    params["TradeId"] = req.TradeId
    
    params["WaybillNo"] = req.WaybillNo
    
    return params
}

// TaobaoWlbImportThreeplOfflineConsignResponse 菜鸟认证直邮线下发货
type TaobaoWlbImportThreeplOfflineConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result TopResult `json:"result";xml:"result"`
    
}

// TmallItemOuteridUpdateRequest 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
type TmallItemOuteridUpdateRequest struct {
    
    // ItemId required商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // OuterId optional商品维度商家编码，如果不修改可以不传；清空请设置空串
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // SkuOuters optional商品SKU更新OuterId时候用的数据
    SkuOuters UpdateSkuOuterId `json:"sku_outers";xml:"sku_outers"`
    
}

func (req *TmallItemOuteridUpdateRequest) GetAPIName() string {
	return "tmall.item.outerid.update"
}

func (req *TmallItemOuteridUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["ItemId"] = req.ItemId
    
    params["OuterId"] = req.OuterId
    
    params["SkuOuters"] = req.SkuOuters
    
    return params
}

// TmallItemOuteridUpdateResponse 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
type TmallItemOuteridUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OuteridUpdateResult Basic商家编码更新结果
    OuteridUpdateResult string `json:"outerid_update_result";xml:"outerid_update_result"`
    
}

// TaobaoQimenOrderSnReportRequest WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
type TaobaoQimenOrderSnReportRequest struct {
    
    // CurrentPage required当前页(从1开始)
    CurrentPage int64 `json:"currentPage";xml:"currentPage"`
    
    // ExtendProps optional扩展属性
    ExtendProps map[string]interface{} `json:"extendProps";xml:"extendProps"`
    
    // Items optional商品列表
    Items Items `json:"items";xml:"items"`
    
    // Order optional单据列表
    Order Order `json:"order";xml:"order"`
    
    // PageSize required每页记录的条数
    PageSize int64 `json:"pageSize";xml:"pageSize"`
    
    // TotalPage required总页数
    TotalPage int64 `json:"totalPage";xml:"totalPage"`
    
}

func (req *TaobaoQimenOrderSnReportRequest) GetAPIName() string {
	return "taobao.qimen.order.sn.report"
}

func (req *TaobaoQimenOrderSnReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["CurrentPage"] = req.CurrentPage
    
    params["ExtendProps"] = req.ExtendProps
    
    params["Items"] = req.Items
    
    params["Order"] = req.Order
    
    params["PageSize"] = req.PageSize
    
    params["TotalPage"] = req.TotalPage
    
    return params
}

// TaobaoQimenOrderSnReportResponse WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
type TaobaoQimenOrderSnReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Code Basic响应码
    Code string `json:"code";xml:"code"`
    
    // Flag Basic响应结果:success|failure
    Flag string `json:"flag";xml:"flag"`
    
    // Message Basic响应信息
    Message string `json:"message";xml:"message"`
    
}

// QimenTaobaoAlphaxOpenJxtExpressRequest isv 接入指定快递api
type QimenTaobaoAlphaxOpenJxtExpressRequest struct {
    
    // Cap required快递公司名称
    Cap string `json:"cap";xml:"cap"`
    
    // OrderId required主订单id
    OrderId string `json:"order_id";xml:"order_id"`
    
    // SellerId required卖家主账号
    SellerId string `json:"seller_id";xml:"seller_id"`
    
    // SellerNick required卖家名称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtExpressRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.express"
}

func (req *QimenTaobaoAlphaxOpenJxtExpressRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["Cap"] = req.Cap
    
    params["OrderId"] = req.OrderId
    
    params["SellerId"] = req.SellerId
    
    params["SellerNick"] = req.SellerNick
    
    return params
}

// QimenTaobaoAlphaxOpenJxtExpressResponse isv 接入指定快递api
type QimenTaobaoAlphaxOpenJxtExpressResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回值
    Result ResultDO `json:"result";xml:"result"`
    
}

// QimenTaobaoAlphaxOpenJxtDelivergoodsRequest isv 接入催发货接口api
type QimenTaobaoAlphaxOpenJxtDelivergoodsRequest struct {
    
    // OrderId required主订单id
    OrderId string `json:"order_id";xml:"order_id"`
    
    // SellerId required卖家主账号
    SellerId string `json:"seller_id";xml:"seller_id"`
    
    // SellerNick required卖家名称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtDelivergoodsRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.delivergoods"
}

func (req *QimenTaobaoAlphaxOpenJxtDelivergoodsRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["OrderId"] = req.OrderId
    
    params["SellerId"] = req.SellerId
    
    params["SellerNick"] = req.SellerNick
    
    return params
}

// QimenTaobaoAlphaxOpenJxtDelivergoodsResponse isv 接入催发货接口api
type QimenTaobaoAlphaxOpenJxtDelivergoodsResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回值
    Result ResultDO `json:"result";xml:"result"`
    
}

// TaobaoSkusQuantityUpdateRequest 提供按照全量/增量的方式批量修改SKU库存的功能
type TaobaoSkusQuantityUpdateRequest struct {
    
    // NumIid required商品数字ID，必填参数
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // OuteridQuantities special特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。
    OuteridQuantities string `json:"outerid_quantities";xml:"outerid_quantities"`
    
    // SkuidQuantities specialsku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。
    SkuidQuantities string `json:"skuid_quantities";xml:"skuid_quantities"`
    
    // Type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0.
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoSkusQuantityUpdateRequest) GetAPIName() string {
	return "taobao.skus.quantity.update"
}

func (req *TaobaoSkusQuantityUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["NumIid"] = req.NumIid
    
    params["OuteridQuantities"] = req.OuteridQuantities
    
    params["SkuidQuantities"] = req.SkuidQuantities
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoSkusQuantityUpdateResponse 提供按照全量/增量的方式批量修改SKU库存的功能
type TaobaoSkusQuantityUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Objectiid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    Item Item `json:"item";xml:"item"`
    
}

// CainiaoEndpointLockerTopWithholdQueryRequest 查询是否有代扣欠款，是否签署代扣协议。
type CainiaoEndpointLockerTopWithholdQueryRequest struct {
    
    // CompanyCode required柜子公司编码
    CompanyCode string `json:"company_code";xml:"company_code"`
    
    // OpenUserId required开放用户Id
    OpenUserId string `json:"open_user_id";xml:"open_user_id"`
    
    // OrderType required柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用）
    OrderType int64 `json:"order_type";xml:"order_type"`
    
}

func (req *CainiaoEndpointLockerTopWithholdQueryRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.withhold.query"
}

func (req *CainiaoEndpointLockerTopWithholdQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CompanyCode"] = req.CompanyCode
    
    params["OpenUserId"] = req.OpenUserId
    
    params["OrderType"] = req.OrderType
    
    return params
}

// CainiaoEndpointLockerTopWithholdQueryResponse 查询是否有代扣欠款，是否签署代扣协议。
type CainiaoEndpointLockerTopWithholdQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresponse
    Result SingleResult `json:"result";xml:"result"`
    
}

// TaobaoWangwangEserviceChatrelationGetRequest 获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
//A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
//2016-09-01， B
//2016-09-02， B
type TaobaoWangwangEserviceChatrelationGetRequest struct {
    
    // ChatRelationRequest required请求参数
    ChatRelationRequest ChatRelationRequest `json:"chat_relation_request";xml:"chat_relation_request"`
    
}

func (req *TaobaoWangwangEserviceChatrelationGetRequest) GetAPIName() string {
	return "taobao.wangwang.eservice.chatrelation.get"
}

func (req *TaobaoWangwangEserviceChatrelationGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ChatRelationRequest"] = req.ChatRelationRequest
    
    return params
}

// TaobaoWangwangEserviceChatrelationGetResponse 获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
//A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
//2016-09-01， B
//2016-09-02， B
type TaobaoWangwangEserviceChatrelationGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result ChatRelationResult `json:"result";xml:"result"`
    
}

// CainiaoEndpointLockerTopOrderWithholdRequest 提供代扣，允许有一笔欠款。
type CainiaoEndpointLockerTopOrderWithholdRequest struct {
    
    // CompanyCode required柜子公司编码
    CompanyCode string `json:"company_code";xml:"company_code"`
    
    // Extra optional扩展字段
    Extra string `json:"extra";xml:"extra"`
    
    // GuiId required柜子id
    GuiId string `json:"gui_id";xml:"gui_id"`
    
    // MailNo required运单号
    MailNo string `json:"mail_no";xml:"mail_no"`
    
    // OpenUserId required开放用户id
    OpenUserId string `json:"open_user_id";xml:"open_user_id"`
    
    // OrderCode required柜子订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderType required订单类型(0-取件业务，1-寄件业务，2-派样业务)
    OrderType int64 `json:"order_type";xml:"order_type"`
    
    // TotalFee required代扣金额（全额），单位：分
    TotalFee int64 `json:"total_fee";xml:"total_fee"`
    
}

func (req *CainiaoEndpointLockerTopOrderWithholdRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.order.withhold"
}

func (req *CainiaoEndpointLockerTopOrderWithholdRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["CompanyCode"] = req.CompanyCode
    
    params["Extra"] = req.Extra
    
    params["GuiId"] = req.GuiId
    
    params["MailNo"] = req.MailNo
    
    params["OpenUserId"] = req.OpenUserId
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderType"] = req.OrderType
    
    params["TotalFee"] = req.TotalFee
    
    return params
}

// CainiaoEndpointLockerTopOrderWithholdResponse 提供代扣，允许有一笔欠款。
type CainiaoEndpointLockerTopOrderWithholdResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result SingleResult `json:"result";xml:"result"`
    
}

// TaobaoWlbItemAddRequest 添加物流宝商品，支持物流宝子商品和属性添加
type TaobaoWlbItemAddRequest struct {
    
    // Color optional商品颜色
    Color string `json:"color";xml:"color"`
    
    // GoodsCat optional货类
    GoodsCat string `json:"goods_cat";xml:"goods_cat"`
    
    // Height optional商品高度，单位毫米
    Height int64 `json:"height";xml:"height"`
    
    // IsDangerous optional是否危险品
    IsDangerous bool `json:"is_dangerous";xml:"is_dangerous"`
    
    // IsFriable optional是否易碎品
    IsFriable bool `json:"is_friable";xml:"is_friable"`
    
    // IsSku required是否sku
    IsSku string `json:"is_sku";xml:"is_sku"`
    
    // ItemCode required商品编码
    ItemCode string `json:"item_code";xml:"item_code"`
    
    // Length optional商品长度，单位毫米
    Length int64 `json:"length";xml:"length"`
    
    // Name required商品名称
    Name string `json:"name";xml:"name"`
    
    // PackageMaterial optional商品包装材料类型
    PackageMaterial string `json:"package_material";xml:"package_material"`
    
    // Price optional商品价格，单位：分
    Price int64 `json:"price";xml:"price"`
    
    // PricingCat optional计价货类
    PricingCat string `json:"pricing_cat";xml:"pricing_cat"`
    
    // ProNameList optional属性名列表,目前支持的属性：
    // 毛重:GWeight	
    // 净重:Nweight
    // 皮重: Tweight
    // 自定义属性：
    // paramkey1
    // paramkey2
    // paramkey3
    // paramkey4
    ProNameList string `json:"pro_name_list";xml:"pro_name_list"`
    
    // ProValueList optional属性值列表：
    // 10,8
    ProValueList string `json:"pro_value_list";xml:"pro_value_list"`
    
    // Remark optional商品备注
    Remark string `json:"remark";xml:"remark"`
    
    // SupportBatch optional是否支持批次
    SupportBatch bool `json:"support_batch";xml:"support_batch"`
    
    // Title optional商品标题
    Title string `json:"title";xml:"title"`
    
    // Type optionalNORMAL--普通商品
    // COMBINE--组合商品
    // DISTRIBUTION--分销
    Type string `json:"type";xml:"type"`
    
    // Volume optional商品体积，单位立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Weight optional商品重量，单位G
    Weight int64 `json:"weight";xml:"weight"`
    
    // Width optional商品宽度，单位毫米
    Width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbItemAddRequest) GetAPIName() string {
	return "taobao.wlb.item.add"
}

func (req *TaobaoWlbItemAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 21)
    
    params["Color"] = req.Color
    
    params["GoodsCat"] = req.GoodsCat
    
    params["Height"] = req.Height
    
    params["IsDangerous"] = req.IsDangerous
    
    params["IsFriable"] = req.IsFriable
    
    params["IsSku"] = req.IsSku
    
    params["ItemCode"] = req.ItemCode
    
    params["Length"] = req.Length
    
    params["Name"] = req.Name
    
    params["PackageMaterial"] = req.PackageMaterial
    
    params["Price"] = req.Price
    
    params["PricingCat"] = req.PricingCat
    
    params["ProNameList"] = req.ProNameList
    
    params["ProValueList"] = req.ProValueList
    
    params["Remark"] = req.Remark
    
    params["SupportBatch"] = req.SupportBatch
    
    params["Title"] = req.Title
    
    params["Type"] = req.Type
    
    params["Volume"] = req.Volume
    
    params["Weight"] = req.Weight
    
    params["Width"] = req.Width
    
    return params
}

// TaobaoWlbItemAddResponse 添加物流宝商品，支持物流宝子商品和属性添加
type TaobaoWlbItemAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemId Basic新增的商品
    ItemId map[string]interface{} `json:"item_id";xml:"item_id"`
    
}

// TaobaoWlbWlborderGetRequest 根据物流宝订单编号查询物流宝订单概要信息
type TaobaoWlbWlborderGetRequest struct {
    
    // WlbOrderCode required物流宝订单编码
    WlbOrderCode string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbWlborderGetRequest) GetAPIName() string {
	return "taobao.wlb.wlborder.get"
}

func (req *TaobaoWlbWlborderGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WlbOrderCode"] = req.WlbOrderCode
    
    return params
}

// TaobaoWlbWlborderGetResponse 根据物流宝订单编号查询物流宝订单概要信息
type TaobaoWlbWlborderGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlbOrder Object物流宝订单对象
    WlbOrder WlbOrder `json:"wlb_order";xml:"wlb_order"`
    
}

// CainiaoNbaddAppointdeliverGetconsigninfoRequest 获取支持定时派送服务发货信息
type CainiaoNbaddAppointdeliverGetconsigninfoRequest struct {
    
    // TradeOrderId required淘宝交易订单id
    TradeOrderId int64 `json:"trade_order_id";xml:"trade_order_id"`
    
}

func (req *CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetAPIName() string {
	return "cainiao.nbadd.appointdeliver.getconsigninfo"
}

func (req *CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["TradeOrderId"] = req.TradeOrderId
    
    return params
}

// CainiaoNbaddAppointdeliverGetconsigninfoResponse 获取支持定时派送服务发货信息
type CainiaoNbaddAppointdeliverGetconsigninfoResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic调用是否成功，正常情况下都会成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // Result Object发货信息
    Result ConsignSupportInfoDTO `json:"result";xml:"result"`
    
    // ResultCode Basic错误编码
    ResultCode string `json:"result_code";xml:"result_code"`
    
    // ResultDesc Basic错误描述
    ResultDesc string `json:"result_desc";xml:"result_desc"`
    
}

// TaobaoOmniorderDtdResendRequest 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoOmniorderDtdResendRequest struct {
    
    // MainOrderId required淘宝主订单ID
    MainOrderId int64 `json:"main_order_id";xml:"main_order_id"`
    
}

func (req *TaobaoOmniorderDtdResendRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.resend"
}

func (req *TaobaoOmniorderDtdResendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["MainOrderId"] = req.MainOrderId
    
    return params
}

// TaobaoOmniorderDtdResendResponse 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoOmniorderDtdResendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Message Basic错误信息
    Message string `json:"message";xml:"message"`
    
    // ResultCode Basic错误码，为0表示成功，非0表示失败
    ResultCode string `json:"result_code";xml:"result_code"`
    
}

// TaobaoOmniorderDtdConsignRequest 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoOmniorderDtdConsignRequest struct {
    
    // MainOrderId required淘宝订单主订单号
    MainOrderId int64 `json:"main_order_id";xml:"main_order_id"`
    
    // StoreId optional发货对应的商户中心门店ID
    StoreId int64 `json:"store_id";xml:"store_id"`
    
}

func (req *TaobaoOmniorderDtdConsignRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.consign"
}

func (req *TaobaoOmniorderDtdConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["MainOrderId"] = req.MainOrderId
    
    params["StoreId"] = req.StoreId
    
    return params
}

// TaobaoOmniorderDtdConsignResponse 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoOmniorderDtdConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Message Basic错误信息
    Message string `json:"message";xml:"message"`
    
    // ResultCode Basic错误码，为0表示成功，非0表示失败
    ResultCode string `json:"result_code";xml:"result_code"`
    
}

// TaobaoOmniorderDtdConsumeRequest 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoOmniorderDtdConsumeRequest struct {
    
    // ParamDoor2doorConsumeRequest optional核销信息
    ParamDoor2doorConsumeRequest Door2doorConsumeRequest `json:"param_door2door_consume_request";xml:"param_door2door_consume_request"`
    
}

func (req *TaobaoOmniorderDtdConsumeRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.consume"
}

func (req *TaobaoOmniorderDtdConsumeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamDoor2doorConsumeRequest"] = req.ParamDoor2doorConsumeRequest
    
    return params
}

// TaobaoOmniorderDtdConsumeResponse 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoOmniorderDtdConsumeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Message Basic错误西溪
    Message string `json:"message";xml:"message"`
    
    // ResultCode Basic错误码，为0表示成功，非0表示失败
    ResultCode string `json:"result_code";xml:"result_code"`
    
}

// TaobaoOmniorderDtdQueryRequest 门店自送根据核销码码查询订单信息
type TaobaoOmniorderDtdQueryRequest struct {
    
    // Code required核销码
    Code string `json:"code";xml:"code"`
    
}

func (req *TaobaoOmniorderDtdQueryRequest) GetAPIName() string {
	return "taobao.omniorder.dtd.query"
}

func (req *TaobaoOmniorderDtdQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Code"] = req.Code
    
    return params
}

// TaobaoOmniorderDtdQueryResponse 门店自送根据核销码码查询订单信息
type TaobaoOmniorderDtdQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Data Objectdata
    Data Door2doorQueryResult `json:"data";xml:"data"`
    
    // Message Basic错误信息
    Message string `json:"message";xml:"message"`
    
    // ResultCode Basic错误码，为0表示成功，非0表示失败
    ResultCode string `json:"result_code";xml:"result_code"`
    
}

// TaobaoWlbOrderstatusGetRequest 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
type TaobaoWlbOrderstatusGetRequest struct {
    
    // OrderCode required物流宝订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
}

func (req *TaobaoWlbOrderstatusGetRequest) GetAPIName() string {
	return "taobao.wlb.orderstatus.get"
}

func (req *TaobaoWlbOrderstatusGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["OrderCode"] = req.OrderCode
    
    return params
}

// TaobaoWlbOrderstatusGetResponse 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
type TaobaoWlbOrderstatusGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlbOrderStatus Object Array订单流转信息状态列表
    WlbOrderStatus WlbProcessStatus `json:"wlb_order_status";xml:"wlb_order_status"`
    
}

// TmallItemSimpleschemaUpdateRequest 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateRequest struct {
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // SchemaXmlFields required编辑商品时提交的xml信息
    SchemaXmlFields string `json:"schema_xml_fields";xml:"schema_xml_fields"`
    
}

func (req *TmallItemSimpleschemaUpdateRequest) GetAPIName() string {
	return "tmall.item.simpleschema.update"
}

func (req *TmallItemSimpleschemaUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ItemId"] = req.ItemId
    
    params["SchemaXmlFields"] = req.SchemaXmlFields
    
    return params
}

// TmallItemSimpleschemaUpdateResponse 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic编辑商品操作成功时间
    GmtModified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
    // SkuMapJson Basicsku与outerId映射信息
    SkuMapJson string `json:"sku_map_json";xml:"sku_map_json"`
    
    // UpdateItemResult Basic编辑商品的itemid
    UpdateItemResult string `json:"update_item_result";xml:"update_item_result"`
    
}

// TaobaoWlbItemGetRequest 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaoWlbItemGetRequest struct {
    
    // ItemId required商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemGetRequest) GetAPIName() string {
	return "taobao.wlb.item.get"
}

func (req *TaobaoWlbItemGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TaobaoWlbItemGetResponse 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaoWlbItemGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Object商品信息
    Item WlbItem `json:"item";xml:"item"`
    
}

// TaobaoWlbOrderCancelRequest 取消物流宝订单
type TaobaoWlbOrderCancelRequest struct {
    
    // WlbOrderCode required物流宝订单编号
    WlbOrderCode string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbOrderCancelRequest) GetAPIName() string {
	return "taobao.wlb.order.cancel"
}

func (req *TaobaoWlbOrderCancelRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WlbOrderCode"] = req.WlbOrderCode
    
    return params
}

// TaobaoWlbOrderCancelResponse 取消物流宝订单
type TaobaoWlbOrderCancelResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrorCodeList Basic错误编码列表
    ErrorCodeList string `json:"error_code_list";xml:"error_code_list"`
    
    // ModifyTime Basic修改时间，只有在取消成功的情况下，才可以做
    ModifyTime time.Time `json:"modify_time";xml:"modify_time"`
    
}

// TaobaoWlbInventoryDetailGetRequest 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
type TaobaoWlbInventoryDetailGetRequest struct {
    
    // InventoryTypeList optional库存类型列表，值包括：
    // VENDIBLE--可销售库存
    // FREEZE--冻结库存
    // ONWAY--在途库存
    // DEFECT--残次品库存
    // ENGINE_DAMAGE--机损
    // BOX_DAMAGE--箱损
    // EXPIRATION--过保
    InventoryTypeList string `json:"inventory_type_list";xml:"inventory_type_list"`
    
    // ItemId required商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // StoreCode optional仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbInventoryDetailGetRequest) GetAPIName() string {
	return "taobao.wlb.inventory.detail.get"
}

func (req *TaobaoWlbInventoryDetailGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["InventoryTypeList"] = req.InventoryTypeList
    
    params["ItemId"] = req.ItemId
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoWlbInventoryDetailGetResponse 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
type TaobaoWlbInventoryDetailGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InventoryList Object Array库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性
    InventoryList WlbInventory `json:"inventory_list";xml:"inventory_list"`
    
    // ItemId Basic入参的item_id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

// TaobaoWlbTradeorderGetRequest 根据交易类型和交易id查询物流宝订单详情
type TaobaoWlbTradeorderGetRequest struct {
    
    // SubTradeId optional子交易号
    SubTradeId string `json:"sub_trade_id";xml:"sub_trade_id"`
    
    // TradeId required指定交易类型的交易号
    TradeId string `json:"trade_id";xml:"trade_id"`
    
    // TradeType required交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
    TradeType string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoWlbTradeorderGetRequest) GetAPIName() string {
	return "taobao.wlb.tradeorder.get"
}

func (req *TaobaoWlbTradeorderGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["SubTradeId"] = req.SubTradeId
    
    params["TradeId"] = req.TradeId
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoWlbTradeorderGetResponse 根据交易类型和交易id查询物流宝订单详情
type TaobaoWlbTradeorderGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WlbOrderList Object Array物流宝订单对象
    WlbOrderList WlbOrder `json:"wlb_order_list";xml:"wlb_order_list"`
    
}

// TaobaoWlbItemQueryRequest 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryRequest struct {
    
    // IsSku optional是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理
    IsSku string `json:"is_sku";xml:"is_sku"`
    
    // ItemCode optional商家编码
    ItemCode string `json:"item_code";xml:"item_code"`
    
    // ItemType optionalITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理
    ItemType string `json:"item_type";xml:"item_type"`
    
    // Name optional商品名称
    Name string `json:"name";xml:"name"`
    
    // PageNo optional当前页
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ParentId optional父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品
    ParentId int64 `json:"parent_id";xml:"parent_id"`
    
    // Status optional只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理
    Status string `json:"status";xml:"status"`
    
    // Title optional商品前台销售名字
    Title string `json:"title";xml:"title"`
    
}

func (req *TaobaoWlbItemQueryRequest) GetAPIName() string {
	return "taobao.wlb.item.query"
}

func (req *TaobaoWlbItemQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["IsSku"] = req.IsSku
    
    params["ItemCode"] = req.ItemCode
    
    params["ItemType"] = req.ItemType
    
    params["Name"] = req.Name
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["ParentId"] = req.ParentId
    
    params["Status"] = req.Status
    
    params["Title"] = req.Title
    
    return params
}

// TaobaoWlbItemQueryResponse 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemList Object Array商品信息列表
    ItemList WlbItem `json:"item_list";xml:"item_list"`
    
    // TotalCount Basic结果总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoWlbOrderCreateRequest 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
type TaobaoWlbOrderCreateRequest struct {
    
    // AlipayNo optional支付宝交易号
    AlipayNo string `json:"alipay_no";xml:"alipay_no"`
    
    // Attributes optional该字段暂时保留
    Attributes string `json:"attributes";xml:"attributes"`
    
    // BuyerNick optional买家呢称
    BuyerNick string `json:"buyer_nick";xml:"buyer_nick"`
    
    // ExpectEndTime optional期望结束时间，在入库单会使用到
    ExpectEndTime time.Time `json:"expect_end_time";xml:"expect_end_time"`
    
    // ExpectStartTime optional计划开始送达时间  在入库单中可能会使用
    ExpectStartTime time.Time `json:"expect_start_time";xml:"expect_start_time"`
    
    // InvoinceInfo optional{"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]}
    InvoinceInfo string `json:"invoince_info";xml:"invoince_info"`
    
    // IsFinished required该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。
    IsFinished bool `json:"is_finished";xml:"is_finished"`
    
    // OrderCode optional物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderFlag optional用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，
    // : 是否可改配送方式  默认可更改
    // 11 CONSIGN 物流宝代理发货,自动更改发货状态
    // 12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    OrderFlag string `json:"order_flag";xml:"order_flag"`
    
    // OrderItemList required订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick
    // ":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。
    OrderItemList string `json:"order_item_list";xml:"order_item_list"`
    
    // OrderSubType required订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购
    OrderSubType string `json:"order_sub_type";xml:"order_sub_type"`
    
    // OrderType required订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库
    OrderType string `json:"order_type";xml:"order_type"`
    
    // OutBizCode required外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用
    OutBizCode string `json:"out_biz_code";xml:"out_biz_code"`
    
    // PackageCount optional包裹件数，入库单和出库单中会用到
    PackageCount int64 `json:"package_count";xml:"package_count"`
    
    // PayableAmount optional应收金额，cod订单必选
    PayableAmount int64 `json:"payable_amount";xml:"payable_amount"`
    
    // PrevOrderCode optional源订单编号
    PrevOrderCode string `json:"prev_order_code";xml:"prev_order_code"`
    
    // ReceiverInfo optional收货方信息，必须传， 手机和电话必选其一。
    // 收货方信息：
    // 邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
    // 
    // 如果某一个字段的数据为空时，必须传NA
    ReceiverInfo string `json:"receiver_info";xml:"receiver_info"`
    
    // Remark optional备注
    Remark string `json:"remark";xml:"remark"`
    
    // ScheduleEnd optional投递时间范围要求,格式'15:20'用分号隔开
    ScheduleEnd string `json:"schedule_end";xml:"schedule_end"`
    
    // ScheduleStart optional投递时间范围要求,格式'13:20'用分号隔开
    ScheduleStart string `json:"schedule_start";xml:"schedule_start"`
    
    // ScheduleType optional投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY
    ScheduleType string `json:"schedule_type";xml:"schedule_type"`
    
    // SenderInfo optional发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：
    // 邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
    // 如果某一个字段的数据为空时，必须传NA
    SenderInfo string `json:"sender_info";xml:"sender_info"`
    
    // ServiceFee optionalcod服务费，只有cod订单的时候，才需要这个字段
    ServiceFee int64 `json:"service_fee";xml:"service_fee"`
    
    // StoreCode required仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
    // TmsInfo optional出库单中可能会用到
    // 运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号
    // 
    // ========================================
    // 如果某一个字段的数据为空时，必须传NA
    TmsInfo string `json:"tms_info";xml:"tms_info"`
    
    // TmsOrderCode optional运单编号，退货单时可能会使用
    TmsOrderCode string `json:"tms_order_code";xml:"tms_order_code"`
    
    // TmsServiceCode optional物流公司编码
    TmsServiceCode string `json:"tms_service_code";xml:"tms_service_code"`
    
    // TotalAmount optional总金额
    TotalAmount int64 `json:"total_amount";xml:"total_amount"`
    
}

func (req *TaobaoWlbOrderCreateRequest) GetAPIName() string {
	return "taobao.wlb.order.create"
}

func (req *TaobaoWlbOrderCreateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 28)
    
    params["AlipayNo"] = req.AlipayNo
    
    params["Attributes"] = req.Attributes
    
    params["BuyerNick"] = req.BuyerNick
    
    params["ExpectEndTime"] = req.ExpectEndTime
    
    params["ExpectStartTime"] = req.ExpectStartTime
    
    params["InvoinceInfo"] = req.InvoinceInfo
    
    params["IsFinished"] = req.IsFinished
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderFlag"] = req.OrderFlag
    
    params["OrderItemList"] = req.OrderItemList
    
    params["OrderSubType"] = req.OrderSubType
    
    params["OrderType"] = req.OrderType
    
    params["OutBizCode"] = req.OutBizCode
    
    params["PackageCount"] = req.PackageCount
    
    params["PayableAmount"] = req.PayableAmount
    
    params["PrevOrderCode"] = req.PrevOrderCode
    
    params["ReceiverInfo"] = req.ReceiverInfo
    
    params["Remark"] = req.Remark
    
    params["ScheduleEnd"] = req.ScheduleEnd
    
    params["ScheduleStart"] = req.ScheduleStart
    
    params["ScheduleType"] = req.ScheduleType
    
    params["SenderInfo"] = req.SenderInfo
    
    params["ServiceFee"] = req.ServiceFee
    
    params["StoreCode"] = req.StoreCode
    
    params["TmsInfo"] = req.TmsInfo
    
    params["TmsOrderCode"] = req.TmsOrderCode
    
    params["TmsServiceCode"] = req.TmsServiceCode
    
    params["TotalAmount"] = req.TotalAmount
    
    return params
}

// TaobaoWlbOrderCreateResponse 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
type TaobaoWlbOrderCreateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // CreateTime Basic订单创建时间
    CreateTime time.Time `json:"create_time";xml:"create_time"`
    
    // OrderCode Basic物流宝订单创建成功后，返回物流宝的订单编号；如果订单创建失败，该字段为空。
    OrderCode string `json:"order_code";xml:"order_code"`
    
}

// QimenTaobaoIcpOrderStockoutordecanceltoerpRequest 出入库单取消推送接口
type QimenTaobaoIcpOrderStockoutordecanceltoerpRequest struct {
    
    // CustomerId required出库单所属货主ID
    CustomerId string `json:"customerId";xml:"customerId"`
    
    // EntryOrderlist optional出入库单信息（一单）
    EntryOrderlist EntryOrderlist `json:"entryOrderlist";xml:"entryOrderlist"`
    
}

func (req *QimenTaobaoIcpOrderStockoutordecanceltoerpRequest) GetAPIName() string {
	return "qimen.taobao.icp.order.stockoutordecanceltoerp"
}

func (req *QimenTaobaoIcpOrderStockoutordecanceltoerpRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CustomerId"] = req.CustomerId
    
    params["EntryOrderlist"] = req.EntryOrderlist
    
    return params
}

// QimenTaobaoIcpOrderStockoutordecanceltoerpResponse 出入库单取消推送接口
type QimenTaobaoIcpOrderStockoutordecanceltoerpResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Object
    Response Response `json:"response";xml:"response"`
    
}

// TmallItemDescModulesGetRequest 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallItemDescModulesGetRequest struct {
    
    // CatId required淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122
    CatId int64 `json:"cat_id";xml:"cat_id"`
    
    // UsrId required商家主帐号id
    UsrId string `json:"usr_id";xml:"usr_id"`
    
}

func (req *TmallItemDescModulesGetRequest) GetAPIName() string {
	return "tmall.item.desc.modules.get"
}

func (req *TmallItemDescModulesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CatId"] = req.CatId
    
    params["UsrId"] = req.UsrId
    
    return params
}

// TmallItemDescModulesGetResponse 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallItemDescModulesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ModularDescInfo Object返回描述模块信息
    ModularDescInfo ModularDescInfo `json:"modular_desc_info";xml:"modular_desc_info"`
    
}

// TaobaoLogisticsConsignOrderCreateandsendRequest 创建物流订单，并发货。
type TaobaoLogisticsConsignOrderCreateandsendRequest struct {
    
    // CompanyId required物流公司ID
    CompanyId int64 `json:"company_id";xml:"company_id"`
    
    // ItemJsonString required物品的json数据。
    ItemJsonString string `json:"item_json_string";xml:"item_json_string"`
    
    // LogisType required物流订单物流类型，值固定选择：2
    LogisType int64 `json:"logis_type";xml:"logis_type"`
    
    // MailNo optional运单号
    MailNo string `json:"mail_no";xml:"mail_no"`
    
    // OrderSource required订单来源，值选择：30
    OrderSource int64 `json:"order_source";xml:"order_source"`
    
    // OrderType required订单类型，值固定选择：30
    OrderType int64 `json:"order_type";xml:"order_type"`
    
    // RAddress required收件人街道地址
    RAddress string `json:"r_address";xml:"r_address"`
    
    // RAreaId required收件人区域ID
    RAreaId int64 `json:"r_area_id";xml:"r_area_id"`
    
    // RCityName required市
    RCityName string `json:"r_city_name";xml:"r_city_name"`
    
    // RDistName optional区
    RDistName string `json:"r_dist_name";xml:"r_dist_name"`
    
    // RMobilePhone optional手机号码
    RMobilePhone string `json:"r_mobile_phone";xml:"r_mobile_phone"`
    
    // RName required收件人名称
    RName string `json:"r_name";xml:"r_name"`
    
    // RProvName required省
    RProvName string `json:"r_prov_name";xml:"r_prov_name"`
    
    // RTelephone optional电话号码
    RTelephone string `json:"r_telephone";xml:"r_telephone"`
    
    // RZipCode required收件人邮编
    RZipCode string `json:"r_zip_code";xml:"r_zip_code"`
    
    // SAddress required发件人街道地址
    SAddress string `json:"s_address";xml:"s_address"`
    
    // SAreaId required发件人区域ID
    SAreaId int64 `json:"s_area_id";xml:"s_area_id"`
    
    // SCityName required市
    SCityName string `json:"s_city_name";xml:"s_city_name"`
    
    // SDistName optional区
    SDistName string `json:"s_dist_name";xml:"s_dist_name"`
    
    // SMobilePhone optional手机号码
    SMobilePhone string `json:"s_mobile_phone";xml:"s_mobile_phone"`
    
    // SName required发件人名称
    SName string `json:"s_name";xml:"s_name"`
    
    // SProvName required省
    SProvName string `json:"s_prov_name";xml:"s_prov_name"`
    
    // STelephone optional电话号码
    STelephone string `json:"s_telephone";xml:"s_telephone"`
    
    // SZipCode required发件人出编
    SZipCode string `json:"s_zip_code";xml:"s_zip_code"`
    
    // Shipping optional费用承担方式 1买家承担运费 2卖家承担运费
    Shipping string `json:"shipping";xml:"shipping"`
    
    // TradeId required交易流水号，淘外订单号或者商家内部交易流水号
    TradeId int64 `json:"trade_id";xml:"trade_id"`
    
    // UserId required用户ID
    UserId int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoLogisticsConsignOrderCreateandsendRequest) GetAPIName() string {
	return "taobao.logistics.consign.order.createandsend"
}

func (req *TaobaoLogisticsConsignOrderCreateandsendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 27)
    
    params["CompanyId"] = req.CompanyId
    
    params["ItemJsonString"] = req.ItemJsonString
    
    params["LogisType"] = req.LogisType
    
    params["MailNo"] = req.MailNo
    
    params["OrderSource"] = req.OrderSource
    
    params["OrderType"] = req.OrderType
    
    params["RAddress"] = req.RAddress
    
    params["RAreaId"] = req.RAreaId
    
    params["RCityName"] = req.RCityName
    
    params["RDistName"] = req.RDistName
    
    params["RMobilePhone"] = req.RMobilePhone
    
    params["RName"] = req.RName
    
    params["RProvName"] = req.RProvName
    
    params["RTelephone"] = req.RTelephone
    
    params["RZipCode"] = req.RZipCode
    
    params["SAddress"] = req.SAddress
    
    params["SAreaId"] = req.SAreaId
    
    params["SCityName"] = req.SCityName
    
    params["SDistName"] = req.SDistName
    
    params["SMobilePhone"] = req.SMobilePhone
    
    params["SName"] = req.SName
    
    params["SProvName"] = req.SProvName
    
    params["STelephone"] = req.STelephone
    
    params["SZipCode"] = req.SZipCode
    
    params["Shipping"] = req.Shipping
    
    params["TradeId"] = req.TradeId
    
    params["UserId"] = req.UserId
    
    return params
}

// TaobaoLogisticsConsignOrderCreateandsendResponse 创建物流订单，并发货。
type TaobaoLogisticsConsignOrderCreateandsendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // OrderId Basic订单ID
    OrderId int64 `json:"order_id";xml:"order_id"`
    
    // ResultDesc Basic结果描述
    ResultDesc string `json:"result_desc";xml:"result_desc"`
    
}

// TaobaoInventoryIpcInventorydetailGetRequest 查询库存明细
type TaobaoInventoryIpcInventorydetailGetRequest struct {
    
    // BizOrderId optional主订单号，可选
    BizOrderId int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    // BizSubOrderId optional子订单号，可选
    BizSubOrderId int64 `json:"biz_sub_order_id";xml:"biz_sub_order_id"`
    
    // PageIndex required当前页数
    PageIndex int64 `json:"page_index";xml:"page_index"`
    
    // PageSize required一页显示多少条
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // ScItemId required仓储商品id
    ScItemId int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    // StatusQuery required1:查询预扣  4：查询占用
    StatusQuery int64 `json:"status_query";xml:"status_query"`
    
}

func (req *TaobaoInventoryIpcInventorydetailGetRequest) GetAPIName() string {
	return "taobao.inventory.ipc.inventorydetail.get"
}

func (req *TaobaoInventoryIpcInventorydetailGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["BizOrderId"] = req.BizOrderId
    
    params["BizSubOrderId"] = req.BizSubOrderId
    
    params["PageIndex"] = req.PageIndex
    
    params["PageSize"] = req.PageSize
    
    params["ScItemId"] = req.ScItemId
    
    params["StatusQuery"] = req.StatusQuery
    
    return params
}

// TaobaoInventoryIpcInventorydetailGetResponse 查询库存明细
type TaobaoInventoryIpcInventorydetailGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InventoryDetails Object Array库存明细列表
    InventoryDetails IpcInventoryDetailDo `json:"inventory_details";xml:"inventory_details"`
    
}

// TaobaoWlbWmsConsignInventoryReleaseRequest ERP释放预占用库存
type TaobaoWlbWmsConsignInventoryReleaseRequest struct {
    
    // Content optional库存释放
    Content Content `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsConsignInventoryReleaseRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.inventory.release"
}

func (req *TaobaoWlbWmsConsignInventoryReleaseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// TaobaoWlbWmsConsignInventoryReleaseResponse ERP释放预占用库存
type TaobaoWlbWmsConsignInventoryReleaseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object接口返回model
    Result Result `json:"result";xml:"result"`
    
}

// TaobaoWlbWmsConsignInventoryOccupancyRequest ERP发货库存预占用
type TaobaoWlbWmsConsignInventoryOccupancyRequest struct {
    
    // Content optional库存占用
    Content Content `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsConsignInventoryOccupancyRequest) GetAPIName() string {
	return "taobao.wlb.wms.consign.inventory.occupancy"
}

func (req *TaobaoWlbWmsConsignInventoryOccupancyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// TaobaoWlbWmsConsignInventoryOccupancyResponse ERP发货库存预占用
type TaobaoWlbWmsConsignInventoryOccupancyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsRetry Basic返回失败时，是否需求重试，为true时，建议系统自动重试
    IsRetry bool `json:"is_retry";xml:"is_retry"`
    
    // ItemInventoryList Object Array库存占用明细列表
    ItemInventoryList Iteminventorylist `json:"item_inventory_list";xml:"item_inventory_list"`
    
    // OrderCode BasicERP订单编码
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // WlErrorCode Basic错误码
    WlErrorCode string `json:"wl_error_code";xml:"wl_error_code"`
    
    // WlErrorMsg Basic错误信息
    WlErrorMsg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    // WlSuccess Basic是否成功
    WlSuccess string `json:"wl_success";xml:"wl_success"`
    
}

// TaobaoLogisticsAddressSearchRequest 通过此接口查询卖家地址库，
type TaobaoLogisticsAddressSearchRequest struct {
    
    // Rdef optional可选，参数列表如下：<br><font color='red'>no_def:查询非默认地址<br>get_def:查询默认取货地址，也即对应卖家后台地址库中发货地址（send_def暂不起作用）<br>cancel_def:查询默认退货地址<br>缺省此参数时，查询所有当前用户的地址库</font>
    Rdef string `json:"rdef";xml:"rdef"`
    
}

func (req *TaobaoLogisticsAddressSearchRequest) GetAPIName() string {
	return "taobao.logistics.address.search"
}

func (req *TaobaoLogisticsAddressSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Rdef"] = req.Rdef
    
    return params
}

// TaobaoLogisticsAddressSearchResponse 通过此接口查询卖家地址库，
type TaobaoLogisticsAddressSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Addresses Object Array一组地址库数据，
    Addresses AddressResult `json:"addresses";xml:"addresses"`
    
}

// TaobaoLogisticsAddressAddRequest 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
type TaobaoLogisticsAddressAddRequest struct {
    
    // Addr required详细街道地址，不需要重复填写省/市/区
    Addr string `json:"addr";xml:"addr"`
    
    // CancelDef optional默认退货地址。<br>
    // <font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
    CancelDef bool `json:"cancel_def";xml:"cancel_def"`
    
    // City required所在市
    City string `json:"city";xml:"city"`
    
    // ContactName required联系人姓名 <font color='red'>长度不可超过20个字节</font>
    ContactName string `json:"contact_name";xml:"contact_name"`
    
    // Country optional区、县
    // <br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    Country string `json:"country";xml:"country"`
    
    // GetDef optional默认取货地址。<br>
    // <font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
    GetDef bool `json:"get_def";xml:"get_def"`
    
    // Memo optional备注,<br><font color='red'>备注不能超过256字节</font>
    Memo string `json:"memo";xml:"memo"`
    
    // MobilePhone special手机号码，手机与电话必需有一个
    // <br><font color='red'>手机号码不能超过20位</font>
    MobilePhone string `json:"mobile_phone";xml:"mobile_phone"`
    
    // Phone special电话号码,手机与电话必需有一个
    Phone string `json:"phone";xml:"phone"`
    
    // Province required所在省
    Province string `json:"province";xml:"province"`
    
    // SellerCompany optional公司名称,<br><font color="red">公司名称长能不能超过20字节</font>
    SellerCompany string `json:"seller_company";xml:"seller_company"`
    
    // SendDef optional默认发货地址。<br>
    // <font color='red'>选择此项(true)，将当前地址设为默认发货地址，撤消原默认发货地址</font>
    SendDef bool `json:"send_def";xml:"send_def"`
    
    // ZipCode optional地区邮政编码
    // <br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    ZipCode string `json:"zip_code";xml:"zip_code"`
    
}

func (req *TaobaoLogisticsAddressAddRequest) GetAPIName() string {
	return "taobao.logistics.address.add"
}

func (req *TaobaoLogisticsAddressAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 13)
    
    params["Addr"] = req.Addr
    
    params["CancelDef"] = req.CancelDef
    
    params["City"] = req.City
    
    params["ContactName"] = req.ContactName
    
    params["Country"] = req.Country
    
    params["GetDef"] = req.GetDef
    
    params["Memo"] = req.Memo
    
    params["MobilePhone"] = req.MobilePhone
    
    params["Phone"] = req.Phone
    
    params["Province"] = req.Province
    
    params["SellerCompany"] = req.SellerCompany
    
    params["SendDef"] = req.SendDef
    
    params["ZipCode"] = req.ZipCode
    
    return params
}

// TaobaoLogisticsAddressAddResponse 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
type TaobaoLogisticsAddressAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddressResult Object只返回修改日期modify_date
    AddressResult AddressResult `json:"address_result";xml:"address_result"`
    
}

// TaobaoLogisticsOnlineSendRequest 用户调用该接口可实现在线订单发货（支持货到付款）
//调用该接口实现在线下单发货，有两种情况：<br>
//<font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br>
//如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font>
type TaobaoLogisticsOnlineSendRequest struct {
    
    // CancelId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
    CancelId int64 `json:"cancel_id";xml:"cancel_id"`
    
    // CompanyCode required物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
    CompanyCode string `json:"company_code";xml:"company_code"`
    
    // Feature optionalfeature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
    Feature string `json:"feature";xml:"feature"`
    
    // IsSplit optional表明是否是拆单，默认值0，1表示拆单
    IsSplit int64 `json:"is_split";xml:"is_split"`
    
    // OutSid optional运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    OutSid string `json:"out_sid";xml:"out_sid"`
    
    // SellerIp optional商家的IP地址
    SellerIp string `json:"seller_ip";xml:"seller_ip"`
    
    // SenderId optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    SenderId int64 `json:"sender_id";xml:"sender_id"`
    
    // SubTid optional需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
    SubTid int64 `json:"sub_tid";xml:"sub_tid"`
    
    // Tid required淘宝交易ID
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoLogisticsOnlineSendRequest) GetAPIName() string {
	return "taobao.logistics.online.send"
}

func (req *TaobaoLogisticsOnlineSendRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["CancelId"] = req.CancelId
    
    params["CompanyCode"] = req.CompanyCode
    
    params["Feature"] = req.Feature
    
    params["IsSplit"] = req.IsSplit
    
    params["OutSid"] = req.OutSid
    
    params["SellerIp"] = req.SellerIp
    
    params["SenderId"] = req.SenderId
    
    params["SubTid"] = req.SubTid
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoLogisticsOnlineSendResponse 用户调用该接口可实现在线订单发货（支持货到付款）
//调用该接口实现在线下单发货，有两种情况：<br>
//<font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br>
//如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font>
type TaobaoLogisticsOnlineSendResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Shipping Objectde
    Shipping Shipping `json:"shipping";xml:"shipping"`
    
}

// TaobaoLogisticsAddressRemoveRequest 用此接口删除卖家地址库
type TaobaoLogisticsAddressRemoveRequest struct {
    
    // ContactId required地址库ID
    ContactId int64 `json:"contact_id";xml:"contact_id"`
    
}

func (req *TaobaoLogisticsAddressRemoveRequest) GetAPIName() string {
	return "taobao.logistics.address.remove"
}

func (req *TaobaoLogisticsAddressRemoveRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ContactId"] = req.ContactId
    
    return params
}

// TaobaoLogisticsAddressRemoveResponse 用此接口删除卖家地址库
type TaobaoLogisticsAddressRemoveResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddressResult Object只返回修改日期modify_date
    AddressResult AddressResult `json:"address_result";xml:"address_result"`
    
}

// TaobaoLogisticsAddressModifyRequest 卖家地址库修改
type TaobaoLogisticsAddressModifyRequest struct {
    
    // Addr required详细街道地址，不需要重复填写省/市/区
    Addr string `json:"addr";xml:"addr"`
    
    // CancelDef optional默认退货地址。<br>
    // <font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
    CancelDef bool `json:"cancel_def";xml:"cancel_def"`
    
    // City required所在市
    City string `json:"city";xml:"city"`
    
    // ContactId required地址库ID
    ContactId int64 `json:"contact_id";xml:"contact_id"`
    
    // ContactName required联系人姓名
    // <font color='red'>长度不可超过20个字节</font>
    ContactName string `json:"contact_name";xml:"contact_name"`
    
    // Country optional区、县
    // <br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    Country string `json:"country";xml:"country"`
    
    // GetDef optional默认取货地址。<br>
    // <font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
    GetDef bool `json:"get_def";xml:"get_def"`
    
    // Memo optional备注,<br><font color='red'>备注不能超过256字节</font>
    Memo string `json:"memo";xml:"memo"`
    
    // MobilePhone special手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font>
    MobilePhone string `json:"mobile_phone";xml:"mobile_phone"`
    
    // Phone special电话号码,手机与电话必需有一个
    Phone string `json:"phone";xml:"phone"`
    
    // Province required所在省
    Province string `json:"province";xml:"province"`
    
    // SellerCompany optional公司名称,
    // <br><font color='red'>公司名称长能不能超过20字节</font>
    SellerCompany string `json:"seller_company";xml:"seller_company"`
    
    // SendDef optional默认发货地址。<br>
    // <font color='red'>选择此项(true)，将当前地址设为默认发货地址，撤消原默认发货地址</font>
    SendDef bool `json:"send_def";xml:"send_def"`
    
    // ZipCode optional地区邮政编码
    // <br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    ZipCode string `json:"zip_code";xml:"zip_code"`
    
}

func (req *TaobaoLogisticsAddressModifyRequest) GetAPIName() string {
	return "taobao.logistics.address.modify"
}

func (req *TaobaoLogisticsAddressModifyRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 14)
    
    params["Addr"] = req.Addr
    
    params["CancelDef"] = req.CancelDef
    
    params["City"] = req.City
    
    params["ContactId"] = req.ContactId
    
    params["ContactName"] = req.ContactName
    
    params["Country"] = req.Country
    
    params["GetDef"] = req.GetDef
    
    params["Memo"] = req.Memo
    
    params["MobilePhone"] = req.MobilePhone
    
    params["Phone"] = req.Phone
    
    params["Province"] = req.Province
    
    params["SellerCompany"] = req.SellerCompany
    
    params["SendDef"] = req.SendDef
    
    params["ZipCode"] = req.ZipCode
    
    return params
}

// TaobaoLogisticsAddressModifyResponse 卖家地址库修改
type TaobaoLogisticsAddressModifyResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddressResult Object只返回修改时间modify_date
    AddressResult AddressResult `json:"address_result";xml:"address_result"`
    
}

// TaobaoWlbWmsUnknownPackageUploadRequest 有货无单销退入库单消息回传
type TaobaoWlbWmsUnknownPackageUploadRequest struct {
    
    // Content requiredWlbWmsUnknownPackageUpload
    Content WlbWmsUnknownPackageUpload `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsUnknownPackageUploadRequest) GetAPIName() string {
	return "taobao.wlb.wms.unknown.package.upload"
}

func (req *TaobaoWlbWmsUnknownPackageUploadRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Content"] = req.Content
    
    return params
}

// TaobaoWlbWmsUnknownPackageUploadResponse 有货无单销退入库单消息回传
type TaobaoWlbWmsUnknownPackageUploadResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response ObjectWlbWmsUnknownPackageUploadResp
    Response WlbWmsUnknownPackageUploadResp `json:"response";xml:"response"`
    
}

// TaobaoTradeReceivetimeDelayRequest 延长交易收货时间
type TaobaoTradeReceivetimeDelayRequest struct {
    
    // Days required延长收货的天数，可选值为：3, 5, 7, 10。
    Days int64 `json:"days";xml:"days"`
    
    // Tid required主订单号
    Tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeReceivetimeDelayRequest) GetAPIName() string {
	return "taobao.trade.receivetime.delay"
}

func (req *TaobaoTradeReceivetimeDelayRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Days"] = req.Days
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoTradeReceivetimeDelayResponse 延长交易收货时间
type TaobaoTradeReceivetimeDelayResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Trade Object更新后的交易数据，只包括tid和modified两个字段。
    Trade Trade `json:"trade";xml:"trade"`
    
}

// TaobaoFenxiaoCooperationUpdateRequest 供应商更新合作的分销商等级
type TaobaoFenxiaoCooperationUpdateRequest struct {
    
    // DistributorId required分销商ID
    DistributorId int64 `json:"distributor_id";xml:"distributor_id"`
    
    // GradeId required等级ID(0代表取消)
    GradeId int64 `json:"grade_id";xml:"grade_id"`
    
    // TradeType optional分销方式(新增)： AGENT(代销)、DEALER(经销)(默认为代销)
    TradeType string `json:"trade_type";xml:"trade_type"`
    
}

func (req *TaobaoFenxiaoCooperationUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.cooperation.update"
}

func (req *TaobaoFenxiaoCooperationUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["DistributorId"] = req.DistributorId
    
    params["GradeId"] = req.GradeId
    
    params["TradeType"] = req.TradeType
    
    return params
}

// TaobaoFenxiaoCooperationUpdateResponse 供应商更新合作的分销商等级
type TaobaoFenxiaoCooperationUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic更新结果成功失败
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoFenxiaoGradesGetRequest 根据供应商ID，查询他的分销商等级信息
type TaobaoFenxiaoGradesGetRequest struct {
    
}

func (req *TaobaoFenxiaoGradesGetRequest) GetAPIName() string {
	return "taobao.fenxiao.grades.get"
}

func (req *TaobaoFenxiaoGradesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoFenxiaoGradesGetResponse 根据供应商ID，查询他的分销商等级信息
type TaobaoFenxiaoGradesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // FenxiaoGrades Object Array分销商等级信息
    FenxiaoGrades FenxiaoGrade `json:"fenxiao_grades";xml:"fenxiao_grades"`
    
}

// TaobaoFenxiaoDiscountsGetRequest 查询折扣信息
type TaobaoFenxiaoDiscountsGetRequest struct {
    
    // DiscountId optional折扣ID
    DiscountId int64 `json:"discount_id";xml:"discount_id"`
    
    // ExtFields optional指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
    ExtFields string `json:"ext_fields";xml:"ext_fields"`
    
}

func (req *TaobaoFenxiaoDiscountsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.discounts.get"
}

func (req *TaobaoFenxiaoDiscountsGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["DiscountId"] = req.DiscountId
    
    params["ExtFields"] = req.ExtFields
    
    return params
}

// TaobaoFenxiaoDiscountsGetResponse 查询折扣信息
type TaobaoFenxiaoDiscountsGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Discounts Object Array折扣数据结构
    Discounts Discount `json:"discounts";xml:"discounts"`
    
    // TotalResults Basic记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// QimenTaobaoAlphaxOpenJxtInvoiceRequest isv 开发票接口api
type QimenTaobaoAlphaxOpenJxtInvoiceRequest struct {
    
    // CompanyTitle optional公司抬头
    CompanyTitle string `json:"company_title";xml:"company_title"`
    
    // ExtendArg optionalapi新增字段，主要用于扩展参数，例如增值税扩展字段（registered_address 注册地址、registered_phone 注册电话、bank 开户行、账户 ）
    ExtendArg string `json:"extend_arg";xml:"extend_arg"`
    
    // InvoiceAttr required发票属性(0:公司；1：个人)
    InvoiceAttr int64 `json:"invoice_attr";xml:"invoice_attr"`
    
    // InvoiceKind required发票形态 （1:电子发票;   2：纸质发票)
    InvoiceKind int64 `json:"invoice_kind";xml:"invoice_kind"`
    
    // InvoiceType required发票类型（1:普通发票；2：增值税专用发票）
    InvoiceType int64 `json:"invoice_type";xml:"invoice_type"`
    
    // OrderId required订单id
    OrderId string `json:"order_id";xml:"order_id"`
    
    // SellerId required卖家主账号id
    SellerId string `json:"seller_id";xml:"seller_id"`
    
    // SellerNick required卖家名称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
    // TaxNo optional税号
    TaxNo string `json:"tax_no";xml:"tax_no"`
    
}

func (req *QimenTaobaoAlphaxOpenJxtInvoiceRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.jxt.invoice"
}

func (req *QimenTaobaoAlphaxOpenJxtInvoiceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["CompanyTitle"] = req.CompanyTitle
    
    params["ExtendArg"] = req.ExtendArg
    
    params["InvoiceAttr"] = req.InvoiceAttr
    
    params["InvoiceKind"] = req.InvoiceKind
    
    params["InvoiceType"] = req.InvoiceType
    
    params["OrderId"] = req.OrderId
    
    params["SellerId"] = req.SellerId
    
    params["SellerNick"] = req.SellerNick
    
    params["TaxNo"] = req.TaxNo
    
    return params
}

// QimenTaobaoAlphaxOpenJxtInvoiceResponse isv 开发票接口api
type QimenTaobaoAlphaxOpenJxtInvoiceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回值
    Result Data `json:"result";xml:"result"`
    
}

// TaobaoFuwuScoresGetRequest 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaoFuwuScoresGetRequest struct {
    
    // CurrentPage required当前页
    CurrentPage int64 `json:"current_page";xml:"current_page"`
    
    // Date required评价日期，查询某一天的评价
    Date time.Time `json:"date";xml:"date"`
    
    // PageSize optional每页获取条数。默认值40，最小值1，最大值100。
    PageSize int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoFuwuScoresGetRequest) GetAPIName() string {
	return "taobao.fuwu.scores.get"
}

func (req *TaobaoFuwuScoresGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CurrentPage"] = req.CurrentPage
    
    params["Date"] = req.Date
    
    params["PageSize"] = req.PageSize
    
    return params
}

// TaobaoFuwuScoresGetResponse 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaoFuwuScoresGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ScoreResult Object Array评价流水记录
    ScoreResult ScoreResult `json:"score_result";xml:"score_result"`
    
}

// TaobaoItemSchemaIncrementUpdateRequest 根据schema规则增量修改宝贝信息
type TaobaoItemSchemaIncrementUpdateRequest struct {
    
    // CategoryId optional商品类目id
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // Parameters required修改字段
    Parameters string `json:"parameters";xml:"parameters"`
    
}

func (req *TaobaoItemSchemaIncrementUpdateRequest) GetAPIName() string {
	return "taobao.item.schema.increment.update"
}

func (req *TaobaoItemSchemaIncrementUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CategoryId"] = req.CategoryId
    
    params["ItemId"] = req.ItemId
    
    params["Parameters"] = req.Parameters
    
    return params
}

// TaobaoItemSchemaIncrementUpdateResponse 根据schema规则增量修改宝贝信息
type TaobaoItemSchemaIncrementUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemId Basic商品id
    ItemId string `json:"item_id";xml:"item_id"`
    
}

// TaobaoVasSubscSearchRequest 用于ISV查询自己名下的应用及收费项目的订购记录
type TaobaoVasSubscSearchRequest struct {
    
    // ArticleCode required应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    ArticleCode string `json:"article_code";xml:"article_code"`
    
    // Autosub optional是否自动续费，true=自动续费 false=非自动续费 空=全部
    Autosub bool `json:"autosub";xml:"autosub"`
    
    // EndDeadline optional到期时间结束值
    EndDeadline time.Time `json:"end_deadline";xml:"end_deadline"`
    
    // ExpireNotice optional是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
    ExpireNotice bool `json:"expire_notice";xml:"expire_notice"`
    
    // ItemCode optional收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    ItemCode string `json:"item_code";xml:"item_code"`
    
    // Nick optional淘宝会员名
    Nick string `json:"nick";xml:"nick"`
    
    // PageNo optional页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional一页包含的记录数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartDeadline optional到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
    StartDeadline time.Time `json:"start_deadline";xml:"start_deadline"`
    
    // Status optional订购记录状态，1=有效 2=过期 空=全部
    Status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoVasSubscSearchRequest) GetAPIName() string {
	return "taobao.vas.subsc.search"
}

func (req *TaobaoVasSubscSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["ArticleCode"] = req.ArticleCode
    
    params["Autosub"] = req.Autosub
    
    params["EndDeadline"] = req.EndDeadline
    
    params["ExpireNotice"] = req.ExpireNotice
    
    params["ItemCode"] = req.ItemCode
    
    params["Nick"] = req.Nick
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartDeadline"] = req.StartDeadline
    
    params["Status"] = req.Status
    
    return params
}

// TaobaoVasSubscSearchResponse 用于ISV查询自己名下的应用及收费项目的订购记录
type TaobaoVasSubscSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ArticleSubs Object Array订购关系对象
    ArticleSubs ArticleSub `json:"article_subs";xml:"article_subs"`
    
    // TotalItem Basic总记录数
    TotalItem int64 `json:"total_item";xml:"total_item"`
    
}

// TaobaoItemQuantityUpdateRequest 提供按照全量或增量形式修改宝贝/SKU库存的功能
type TaobaoItemQuantityUpdateRequest struct {
    
    // NumIid required商品数字ID，必填参数
    NumIid int64 `json:"num_iid";xml:"num_iid"`
    
    // OuterId optionalSKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU
    OuterId string `json:"outer_id";xml:"outer_id"`
    
    // Quantity required库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
    Quantity int64 `json:"quantity";xml:"quantity"`
    
    // SkuId optional要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存
    SkuId int64 `json:"sku_id";xml:"sku_id"`
    
    // Type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
    Type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoItemQuantityUpdateRequest) GetAPIName() string {
	return "taobao.item.quantity.update"
}

func (req *TaobaoItemQuantityUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["NumIid"] = req.NumIid
    
    params["OuterId"] = req.OuterId
    
    params["Quantity"] = req.Quantity
    
    params["SkuId"] = req.SkuId
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoItemQuantityUpdateResponse 提供按照全量或增量形式修改宝贝/SKU库存的功能
type TaobaoItemQuantityUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Item Objectiid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    Item Item `json:"item";xml:"item"`
    
}

// TaobaoVasOrderSearchRequest 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。
//建议用于查询前一日的历史记录，不适合用作实时数据查询。
//现在只能查询90天以内的数据
//该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaoVasOrderSearchRequest struct {
    
    // ArticleCode required应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    ArticleCode string `json:"article_code";xml:"article_code"`
    
    // BizOrderId optional订单号
    BizOrderId int64 `json:"biz_order_id";xml:"biz_order_id"`
    
    // BizType optional订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
    BizType int64 `json:"biz_type";xml:"biz_type"`
    
    // EndCreated optional订单创建时间（订购时间）结束值
    EndCreated time.Time `json:"end_created";xml:"end_created"`
    
    // ItemCode optional收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    ItemCode string `json:"item_code";xml:"item_code"`
    
    // Nick optional淘宝会员名
    Nick string `json:"nick";xml:"nick"`
    
    // OrderId optional子订单号
    OrderId int64 `json:"order_id";xml:"order_id"`
    
    // PageNo optional页码
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional一页包含的记录数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartCreated optional订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
    StartCreated time.Time `json:"start_created";xml:"start_created"`
    
}

func (req *TaobaoVasOrderSearchRequest) GetAPIName() string {
	return "taobao.vas.order.search"
}

func (req *TaobaoVasOrderSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 10)
    
    params["ArticleCode"] = req.ArticleCode
    
    params["BizOrderId"] = req.BizOrderId
    
    params["BizType"] = req.BizType
    
    params["EndCreated"] = req.EndCreated
    
    params["ItemCode"] = req.ItemCode
    
    params["Nick"] = req.Nick
    
    params["OrderId"] = req.OrderId
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartCreated"] = req.StartCreated
    
    return params
}

// TaobaoVasOrderSearchResponse 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。
//建议用于查询前一日的历史记录，不适合用作实时数据查询。
//现在只能查询90天以内的数据
//该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaoVasOrderSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ArticleBizOrders Object Array商品订单对象
    ArticleBizOrders ArticleBizOrder `json:"article_biz_orders";xml:"article_biz_orders"`
    
    // TotalItem Basic总记录数
    TotalItem int64 `json:"total_item";xml:"total_item"`
    
}

// TaobaoVasSubscribeGetRequest 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
type TaobaoVasSubscribeGetRequest struct {
    
    // ArticleCode required商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码
    ArticleCode string `json:"article_code";xml:"article_code"`
    
    // Nick required淘宝会员名
    Nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoVasSubscribeGetRequest) GetAPIName() string {
	return "taobao.vas.subscribe.get"
}

func (req *TaobaoVasSubscribeGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ArticleCode"] = req.ArticleCode
    
    params["Nick"] = req.Nick
    
    return params
}

// TaobaoVasSubscribeGetResponse 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
type TaobaoVasSubscribeGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ArticleUserSubscribes Object Array用户订购信息
    ArticleUserSubscribes ArticleUserSubscribe `json:"article_user_subscribes";xml:"article_user_subscribes"`
    
}

// CainiaoWaybillIiQueryByWaybillcodeRequest 通过面单号查看面单号的当前状态，如签收、发货、失效等。
type CainiaoWaybillIiQueryByWaybillcodeRequest struct {
    
    // ParamList optional系统自动生成
    ParamList WaybillDetailQueryByWaybillCodeRequest `json:"param_list";xml:"param_list"`
    
}

func (req *CainiaoWaybillIiQueryByWaybillcodeRequest) GetAPIName() string {
	return "cainiao.waybill.ii.query.by.waybillcode"
}

func (req *CainiaoWaybillIiQueryByWaybillcodeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamList"] = req.ParamList
    
    return params
}

// CainiaoWaybillIiQueryByWaybillcodeResponse 通过面单号查看面单号的当前状态，如签收、发货、失效等。
type CainiaoWaybillIiQueryByWaybillcodeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Modules Object Array查询返回值
    Modules WaybillCloudPrintWithResultDescResponse `json:"modules";xml:"modules"`
    
}

// CainiaoWaybillIiQueryByTradecodeRequest 通过订单号查看面单的信息
type CainiaoWaybillIiQueryByTradecodeRequest struct {
    
    // ParamList optional订单号列表
    ParamList WaybillDetailQueryByBizSubCodeRequest `json:"param_list";xml:"param_list"`
    
}

func (req *CainiaoWaybillIiQueryByTradecodeRequest) GetAPIName() string {
	return "cainiao.waybill.ii.query.by.tradecode"
}

func (req *CainiaoWaybillIiQueryByTradecodeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ParamList"] = req.ParamList
    
    return params
}

// CainiaoWaybillIiQueryByTradecodeResponse 通过订单号查看面单的信息
type CainiaoWaybillIiQueryByTradecodeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Modules Object Array查询返回值
    Modules WaybillCloudPrintWithResultDescResponse `json:"modules";xml:"modules"`
    
}

// TaobaoWlbOrderConsignRequest 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaoWlbOrderConsignRequest struct {
    
    // WlbOrderCode required物流宝订单编号
    WlbOrderCode string `json:"wlb_order_code";xml:"wlb_order_code"`
    
}

func (req *TaobaoWlbOrderConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.consign"
}

func (req *TaobaoWlbOrderConsignRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["WlbOrderCode"] = req.WlbOrderCode
    
    return params
}

// TaobaoWlbOrderConsignResponse 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaoWlbOrderConsignResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ModifyTime Basic修改时间
    ModifyTime time.Time `json:"modify_time";xml:"modify_time"`
    
}

// CainiaoCloudprintSingleCustomareaGetRequest 商家所有快递公司模板只有一个自定义区
type CainiaoCloudprintSingleCustomareaGetRequest struct {
    
    // SellerId optional这是商家用户id
    SellerId int64 `json:"seller_id";xml:"seller_id"`
    
}

func (req *CainiaoCloudprintSingleCustomareaGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.single.customarea.get"
}

func (req *CainiaoCloudprintSingleCustomareaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["SellerId"] = req.SellerId
    
    return params
}

// CainiaoCloudprintSingleCustomareaGetResponse 商家所有快递公司模板只有一个自定义区
type CainiaoCloudprintSingleCustomareaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// TaobaoQimenShopSynchronizeRequest 店铺同步接口描述
type TaobaoQimenShopSynchronizeRequest struct {
    
    // Request optional请求
    Request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenShopSynchronizeRequest) GetAPIName() string {
	return "taobao.qimen.shop.synchronize"
}

func (req *TaobaoQimenShopSynchronizeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Request"] = req.Request
    
    return params
}

// TaobaoQimenShopSynchronizeResponse 店铺同步接口描述
type TaobaoQimenShopSynchronizeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response ObjectResponse
    Response Response `json:"response";xml:"response"`
    
}

// CainiaoCloudprintIsvtemplatesGetRequest 获取商家使用的标准模板
type CainiaoCloudprintIsvtemplatesGetRequest struct {
    
}

func (req *CainiaoCloudprintIsvtemplatesGetRequest) GetAPIName() string {
	return "cainiao.cloudprint.isvtemplates.get"
}

func (req *CainiaoCloudprintIsvtemplatesGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// CainiaoCloudprintIsvtemplatesGetResponse 获取商家使用的标准模板
type CainiaoCloudprintIsvtemplatesGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// AlibabaEinvoiceApplyGetRequest ERP获取开票申请数据
type AlibabaEinvoiceApplyGetRequest struct {
    
    // ApplyId optional开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息
    ApplyId string `json:"apply_id";xml:"apply_id"`
    
    // PlatformTid required平台订单号
    PlatformTid string `json:"platform_tid";xml:"platform_tid"`
    
}

func (req *AlibabaEinvoiceApplyGetRequest) GetAPIName() string {
	return "alibaba.einvoice.apply.get"
}

func (req *AlibabaEinvoiceApplyGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ApplyId"] = req.ApplyId
    
    params["PlatformTid"] = req.PlatformTid
    
    return params
}

// AlibabaEinvoiceApplyGetResponse ERP获取开票申请数据
type AlibabaEinvoiceApplyGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ApplyList Object Array开票明细
    ApplyList Apply `json:"apply_list";xml:"apply_list"`
    
    // IsSuccess Basicsuccess
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoWlbNotifyMessagePageGetRequest 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaoWlbNotifyMessagePageGetRequest struct {
    
    // EndDate optional记录截至时间
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // MsgCode optional通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功
    MsgCode string `json:"msg_code";xml:"msg_code"`
    
    // PageNo optional分页查询页数
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页查询的每页页数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartDate optional记录开始时间
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // Status optional消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM
    Status string `json:"status";xml:"status"`
    
}

func (req *TaobaoWlbNotifyMessagePageGetRequest) GetAPIName() string {
	return "taobao.wlb.notify.message.page.get"
}

func (req *TaobaoWlbNotifyMessagePageGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["EndDate"] = req.EndDate
    
    params["MsgCode"] = req.MsgCode
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartDate"] = req.StartDate
    
    params["Status"] = req.Status
    
    return params
}

// TaobaoWlbNotifyMessagePageGetResponse 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaoWlbNotifyMessagePageGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TotalCount Basic2000-01-01 00:00:00
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
    // WlbMessages Object Array通道消息
    WlbMessages WlbMessage `json:"wlb_messages";xml:"wlb_messages"`
    
}

// CainiaoCloudprintClientinfoPutRequest 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutRequest struct {
    
    // JsonData required客户端上传json数据
    JsonData string `json:"json_data";xml:"json_data"`
    
}

func (req *CainiaoCloudprintClientinfoPutRequest) GetAPIName() string {
	return "cainiao.cloudprint.clientinfo.put"
}

func (req *CainiaoCloudprintClientinfoPutRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["JsonData"] = req.JsonData
    
    return params
}

// CainiaoCloudprintClientinfoPutResponse 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result CloudPrintBaseResult `json:"result";xml:"result"`
    
}

// TaobaoWlbNotifyMessageConfirmRequest 确认物流宝可执行消息
type TaobaoWlbNotifyMessageConfirmRequest struct {
    
    // MessageId required物流宝通知消息的id，通过taobao.wlb.notify.message.page.get接口得到的WlbMessage数据结构中的id字段
    MessageId string `json:"message_id";xml:"message_id"`
    
}

func (req *TaobaoWlbNotifyMessageConfirmRequest) GetAPIName() string {
	return "taobao.wlb.notify.message.confirm"
}

func (req *TaobaoWlbNotifyMessageConfirmRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["MessageId"] = req.MessageId
    
    return params
}

// TaobaoWlbNotifyMessageConfirmResponse 确认物流宝可执行消息
type TaobaoWlbNotifyMessageConfirmResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic物流宝消息确认时间
    GmtModified time.Time `json:"gmt_modified";xml:"gmt_modified"`
    
}

// TaobaoWlbItemCombinationGetRequest 根据商品id查询商品组合关系
type TaobaoWlbItemCombinationGetRequest struct {
    
    // ItemId required要查询的组合商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemCombinationGetRequest) GetAPIName() string {
	return "taobao.wlb.item.combination.get"
}

func (req *TaobaoWlbItemCombinationGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TaobaoWlbItemCombinationGetResponse 根据商品id查询商品组合关系
type TaobaoWlbItemCombinationGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ItemIdList Basic Array组合子商品id列表
    ItemIdList int64 `json:"item_id_list";xml:"item_id_list"`
    
}

// TaobaoWlbItemMapGetRequest 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetRequest struct {
    
    // ItemId required要查询映射关系的物流宝商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoWlbItemMapGetRequest) GetAPIName() string {
	return "taobao.wlb.item.map.get"
}

func (req *TaobaoWlbItemMapGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TaobaoWlbItemMapGetResponse 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // OutEntityItemList Object Array外部商品实体
    OutEntityItemList OutEntityItem `json:"out_entity_item_list";xml:"out_entity_item_list"`
    
}

// TaobaoWlbItemMapGetByExtentityRequest 根据外部实体类型和实体id查询映射的物流宝商品id
type TaobaoWlbItemMapGetByExtentityRequest struct {
    
    // ExtEntityId required外部实体类型对应的商品id
    ExtEntityId int64 `json:"ext_entity_id";xml:"ext_entity_id"`
    
    // ExtEntityType required外部实体类型： IC_ITEM--ic商品 IC_SKU--ic销售单元
    ExtEntityType string `json:"ext_entity_type";xml:"ext_entity_type"`
    
}

func (req *TaobaoWlbItemMapGetByExtentityRequest) GetAPIName() string {
	return "taobao.wlb.item.map.get.by.extentity"
}

func (req *TaobaoWlbItemMapGetByExtentityRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["ExtEntityId"] = req.ExtEntityId
    
    params["ExtEntityType"] = req.ExtEntityType
    
    return params
}

// TaobaoWlbItemMapGetByExtentityResponse 根据外部实体类型和实体id查询映射的物流宝商品id
type TaobaoWlbItemMapGetByExtentityResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // ItemId Basic物流宝商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

// TaobaoWlbTmsorderQueryRequest 通过物流订单编号分页查询物流信息
type TaobaoWlbTmsorderQueryRequest struct {
    
    // OrderCode required物流订单编号
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // PageNo optional当前页
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    PageSize int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoWlbTmsorderQueryRequest) GetAPIName() string {
	return "taobao.wlb.tmsorder.query"
}

func (req *TaobaoWlbTmsorderQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["OrderCode"] = req.OrderCode
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    return params
}

// TaobaoWlbTmsorderQueryResponse 通过物流订单编号分页查询物流信息
type TaobaoWlbTmsorderQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TmsOrderList Object Array物流订单运单信息列表
    TmsOrderList WlbTmsOrder `json:"tms_order_list";xml:"tms_order_list"`
    
    // TotalCount Basic结果总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoItemIncrementUpdateSchemaGetRequest 获取增量更新规则，目前开发的字段有title, subTitle, description, wl_description
type TaobaoItemIncrementUpdateSchemaGetRequest struct {
    
    // CategoryId optional宝贝类目id
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // ItemId required宝贝id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // UpdateFields optional修改字段
    UpdateFields string `json:"update_fields";xml:"update_fields"`
    
}

func (req *TaobaoItemIncrementUpdateSchemaGetRequest) GetAPIName() string {
	return "taobao.item.increment.update.schema.get"
}

func (req *TaobaoItemIncrementUpdateSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CategoryId"] = req.CategoryId
    
    params["ItemId"] = req.ItemId
    
    params["UpdateFields"] = req.UpdateFields
    
    return params
}

// TaobaoItemIncrementUpdateSchemaGetResponse 获取增量更新规则，目前开发的字段有title, subTitle, description, wl_description
type TaobaoItemIncrementUpdateSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateRules Basic返回的结果集
    UpdateRules string `json:"update_rules";xml:"update_rules"`
    
}

// TaobaoWlbOrderPageGetRequest 分页查询物流宝订单
type TaobaoWlbOrderPageGetRequest struct {
    
    // EndTime optional查询截止时间
    EndTime time.Time `json:"end_time";xml:"end_time"`
    
    // OrderCode optional物流订单编号
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // OrderStatus optionalTMS拒签：-100 接收方拒签：-200
    OrderStatus int64 `json:"order_status";xml:"order_status"`
    
    // OrderSubType optional订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
    OrderSubType string `json:"order_sub_type";xml:"order_sub_type"`
    
    // OrderType optional订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
    OrderType string `json:"order_type";xml:"order_type"`
    
    // PageNo optional分页的第几页
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页多少条
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartTime optional查询开始时间
    StartTime time.Time `json:"start_time";xml:"start_time"`
    
}

func (req *TaobaoWlbOrderPageGetRequest) GetAPIName() string {
	return "taobao.wlb.order.page.get"
}

func (req *TaobaoWlbOrderPageGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 8)
    
    params["EndTime"] = req.EndTime
    
    params["OrderCode"] = req.OrderCode
    
    params["OrderStatus"] = req.OrderStatus
    
    params["OrderSubType"] = req.OrderSubType
    
    params["OrderType"] = req.OrderType
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartTime"] = req.StartTime
    
    return params
}

// TaobaoWlbOrderPageGetResponse 分页查询物流宝订单
type TaobaoWlbOrderPageGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // OrderList Object Array物流宝订单对象
    OrderList WlbOrder `json:"order_list";xml:"order_list"`
    
    // TotalCount Basic总条数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoWlbSubscriptionQueryRequest 查询商家定购的所有服务,可通过入参状态来筛选
type TaobaoWlbSubscriptionQueryRequest struct {
    
    // PageNo optional当前页
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // Status optional状态 
    // AUDITING 1-待审核; 
    // CANCEL 2-撤销 ;
    // CHECKED 3-审核通过 ;
    // FAILED 4-审核未通过 ;
    // SYNCHRONIZING 5-同步中;
    // 只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。
    Status string `json:"status";xml:"status"`
    
}

func (req *TaobaoWlbSubscriptionQueryRequest) GetAPIName() string {
	return "taobao.wlb.subscription.query"
}

func (req *TaobaoWlbSubscriptionQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["Status"] = req.Status
    
    return params
}

// TaobaoWlbSubscriptionQueryResponse 查询商家定购的所有服务,可通过入参状态来筛选
type TaobaoWlbSubscriptionQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // SellerSubscriptionList Object Array卖家定购的服务列表
    SellerSubscriptionList WlbSellerSubscription `json:"seller_subscription_list";xml:"seller_subscription_list"`
    
    // TotalCount Basic结果总数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoWlbItemUpdateRequest 修改物流宝商品信息
type TaobaoWlbItemUpdateRequest struct {
    
    // Color optional商品颜色
    Color string `json:"color";xml:"color"`
    
    // DeletePropertyKeyList optional需要删除的商品属性key列表
    DeletePropertyKeyList string `json:"delete_property_key_list";xml:"delete_property_key_list"`
    
    // GoodsCat optional商品货类
    GoodsCat string `json:"goods_cat";xml:"goods_cat"`
    
    // Height optional商品高度，单位厘米
    Height int64 `json:"height";xml:"height"`
    
    // Id required要修改的商品id
    Id int64 `json:"id";xml:"id"`
    
    // IsDangerous optional是否危险品
    IsDangerous bool `json:"is_dangerous";xml:"is_dangerous"`
    
    // IsFriable optional是否易碎品
    IsFriable bool `json:"is_friable";xml:"is_friable"`
    
    // Length optional商品长度，单位厘米
    Length int64 `json:"length";xml:"length"`
    
    // Name optional要修改的商品名称
    Name string `json:"name";xml:"name"`
    
    // PackageMaterial optional商品包装材料类型
    PackageMaterial string `json:"package_material";xml:"package_material"`
    
    // PricingCat optional商品计价货类
    PricingCat string `json:"pricing_cat";xml:"pricing_cat"`
    
    // Remark optional要修改的商品备注
    Remark string `json:"remark";xml:"remark"`
    
    // Title optional要修改的商品标题
    Title string `json:"title";xml:"title"`
    
    // UpdatePropertyKeyList optional需要修改的商品属性值的列表，如果属性不存在，则新增属性
    UpdatePropertyKeyList string `json:"update_property_key_list";xml:"update_property_key_list"`
    
    // UpdatePropertyValueList optional需要修改的属性值的列表
    UpdatePropertyValueList string `json:"update_property_value_list";xml:"update_property_value_list"`
    
    // Volume optional商品体积，单位立方厘米
    Volume int64 `json:"volume";xml:"volume"`
    
    // Weight optional商品重量，单位G
    Weight int64 `json:"weight";xml:"weight"`
    
    // Width optional商品宽度，单位厘米
    Width int64 `json:"width";xml:"width"`
    
}

func (req *TaobaoWlbItemUpdateRequest) GetAPIName() string {
	return "taobao.wlb.item.update"
}

func (req *TaobaoWlbItemUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 18)
    
    params["Color"] = req.Color
    
    params["DeletePropertyKeyList"] = req.DeletePropertyKeyList
    
    params["GoodsCat"] = req.GoodsCat
    
    params["Height"] = req.Height
    
    params["Id"] = req.Id
    
    params["IsDangerous"] = req.IsDangerous
    
    params["IsFriable"] = req.IsFriable
    
    params["Length"] = req.Length
    
    params["Name"] = req.Name
    
    params["PackageMaterial"] = req.PackageMaterial
    
    params["PricingCat"] = req.PricingCat
    
    params["Remark"] = req.Remark
    
    params["Title"] = req.Title
    
    params["UpdatePropertyKeyList"] = req.UpdatePropertyKeyList
    
    params["UpdatePropertyValueList"] = req.UpdatePropertyValueList
    
    params["Volume"] = req.Volume
    
    params["Weight"] = req.Weight
    
    params["Width"] = req.Width
    
    return params
}

// TaobaoWlbItemUpdateResponse 修改物流宝商品信息
type TaobaoWlbItemUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // GmtModified Basic修改时间
    GmtModified bool `json:"gmt_modified";xml:"gmt_modified"`
    
}

// TaobaoWlbInventorylogQueryRequest 通过商品ID等几个条件来分页查询库存变更记录
type TaobaoWlbInventorylogQueryRequest struct {
    
    // GmtEnd optional结束修改时间,小于等于该时间
    GmtEnd time.Time `json:"gmt_end";xml:"gmt_end"`
    
    // GmtStart optional起始修改时间,大于等于该时间
    GmtStart time.Time `json:"gmt_start";xml:"gmt_start"`
    
    // ItemId optional商品ID
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // OpType optional库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
    OpType string `json:"op_type";xml:"op_type"`
    
    // OpUserId optional可指定授权的用户来查询
    OpUserId int64 `json:"op_user_id";xml:"op_user_id"`
    
    // OrderCode optional单号
    OrderCode string `json:"order_code";xml:"order_code"`
    
    // PageNo optional当前页
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional分页记录个数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StoreCode optional仓库编码
    StoreCode string `json:"store_code";xml:"store_code"`
    
}

func (req *TaobaoWlbInventorylogQueryRequest) GetAPIName() string {
	return "taobao.wlb.inventorylog.query"
}

func (req *TaobaoWlbInventorylogQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 9)
    
    params["GmtEnd"] = req.GmtEnd
    
    params["GmtStart"] = req.GmtStart
    
    params["ItemId"] = req.ItemId
    
    params["OpType"] = req.OpType
    
    params["OpUserId"] = req.OpUserId
    
    params["OrderCode"] = req.OrderCode
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StoreCode"] = req.StoreCode
    
    return params
}

// TaobaoWlbInventorylogQueryResponse 通过商品ID等几个条件来分页查询库存变更记录
type TaobaoWlbInventorylogQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // InventoryLogList Object Array库存变更记录
    InventoryLogList WlbItemInventoryLog `json:"inventory_log_list";xml:"inventory_log_list"`
    
    // TotalCount Basic记录数
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
}

// TaobaoTopatsJushitaJdpDatadeleteRequest 异步删除rds库数据推送表的数据
type TaobaoTopatsJushitaJdpDatadeleteRequest struct {
    
    // EndDate required删除数据时间段的结束修改时间，格式为：yyyy-MM-dd HH:mm:ss，结束时间必须为前天的23:59:59秒以前，根据是业务的modified时间
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // RdsName optional数据库实例名，当删除用户的绑定关系已经不在推送里时，此参数必填
    RdsName string `json:"rds_name";xml:"rds_name"`
    
    // StartDate required删除数据时间段的起始修改时间，格式为：yyyy-MM-dd HH:mm:ss,根据是业务的modified时间
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
    // SyncTypes required推送的数据类型，可选值为：tb_trade(淘宝交易)，tb_item(淘宝商品)，tb_refund(淘宝退款)，fx_trade(分销订单)，同时删除多种类型以分号分隔，如："tb_trade;tb_item;tb_refund;fx_trade"
    SyncTypes string `json:"sync_types";xml:"sync_types"`
    
    // UserNick required用户昵称，必填
    UserNick string `json:"user_nick";xml:"user_nick"`
    
}

func (req *TaobaoTopatsJushitaJdpDatadeleteRequest) GetAPIName() string {
	return "taobao.topats.jushita.jdp.datadelete"
}

func (req *TaobaoTopatsJushitaJdpDatadeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 5)
    
    params["EndDate"] = req.EndDate
    
    params["RdsName"] = req.RdsName
    
    params["StartDate"] = req.StartDate
    
    params["SyncTypes"] = req.SyncTypes
    
    params["UserNick"] = req.UserNick
    
    return params
}

// TaobaoTopatsJushitaJdpDatadeleteResponse 异步删除rds库数据推送表的数据
type TaobaoTopatsJushitaJdpDatadeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result ResultDo `json:"result";xml:"result"`
    
}

// CainiaoWaybillIiSearchRequest 获取发货地&CP开通状态&账户的使用情况
type CainiaoWaybillIiSearchRequest struct {
    
    // CpCode optional物流公司code
    CpCode string `json:"cp_code";xml:"cp_code"`
    
}

func (req *CainiaoWaybillIiSearchRequest) GetAPIName() string {
	return "cainiao.waybill.ii.search"
}

func (req *CainiaoWaybillIiSearchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["CpCode"] = req.CpCode
    
    return params
}

// CainiaoWaybillIiSearchResponse 获取发货地&CP开通状态&账户的使用情况
type CainiaoWaybillIiSearchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // WaybillApplySubscriptionCols Object ArrayCP网点信息及对应的商家的发货信息
    WaybillApplySubscriptionCols WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_cols";xml:"waybill_apply_subscription_cols"`
    
}

// TaobaoJushitaJdpUsersGetRequest 获取开通的订单同步服务的用户，含有rds的路由关系
type TaobaoJushitaJdpUsersGetRequest struct {
    
    // EndModified optional此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
    EndModified time.Time `json:"end_modified";xml:"end_modified"`
    
    // PageNo optional当前页数
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页记录数，默认200
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartModified optional此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
    StartModified time.Time `json:"start_modified";xml:"start_modified"`
    
    // TargetAppkey optional普通isv不能传入此参数
    TargetAppkey string `json:"target_appkey";xml:"target_appkey"`
    
    // UserId optional如果传了user_id表示单条查询
    UserId int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoJushitaJdpUsersGetRequest) GetAPIName() string {
	return "taobao.jushita.jdp.users.get"
}

func (req *TaobaoJushitaJdpUsersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["EndModified"] = req.EndModified
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartModified"] = req.StartModified
    
    params["TargetAppkey"] = req.TargetAppkey
    
    params["UserId"] = req.UserId
    
    return params
}

// TaobaoJushitaJdpUsersGetResponse 获取开通的订单同步服务的用户，含有rds的路由关系
type TaobaoJushitaJdpUsersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TotalResults Basic总记录数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
    // Users Object Array用户列表
    Users JdpUser `json:"users";xml:"users"`
    
}

// TaobaoShopGetbytitleRequest 根据店铺名称获取店铺信息
type TaobaoShopGetbytitleRequest struct {
    
    // Fields optional代表需要获取的店铺信息：sid,cid,title,nick,desc,bulletin,pic_path,created,modified,shop_score
    Fields string `json:"fields";xml:"fields"`
    
    // Title required店铺名称，必须严格匹配（不支持模糊查询）
    Title string `json:"title";xml:"title"`
    
}

func (req *TaobaoShopGetbytitleRequest) GetAPIName() string {
	return "taobao.shop.getbytitle"
}

func (req *TaobaoShopGetbytitleRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["Fields"] = req.Fields
    
    params["Title"] = req.Title
    
    return params
}

// TaobaoShopGetbytitleResponse 根据店铺名称获取店铺信息
type TaobaoShopGetbytitleResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // ErrCode Basic错误码
    ErrCode string `json:"err_code";xml:"err_code"`
    
    // ErrMsg Basic错误信息
    ErrMsg string `json:"err_msg";xml:"err_msg"`
    
    // IsError Basic有无错误
    IsError bool `json:"is_error";xml:"is_error"`
    
    // ResultShop Basictest
    ResultShop map[string]interface{} `json:"result_shop";xml:"result_shop"`
    
}

// TmallProductSchemaAddRequest Schema体系发布一个产品
type TmallProductSchemaAddRequest struct {
    
    // BrandId special品牌ID
    BrandId int64 `json:"brand_id";xml:"brand_id"`
    
    // CategoryId required商品发布的目标类目，必须是叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // XmlData required根据tmall.product.add.schema.get生成的产品发布规则入参数据
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallProductSchemaAddRequest) GetAPIName() string {
	return "tmall.product.schema.add"
}

func (req *TmallProductSchemaAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["BrandId"] = req.BrandId
    
    params["CategoryId"] = req.CategoryId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TmallProductSchemaAddResponse Schema体系发布一个产品
type TmallProductSchemaAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddProductResult Basic新发产品结果
    AddProductResult string `json:"add_product_result";xml:"add_product_result"`
    
}

// TmallProductAddSchemaGetRequest 获取用户发布产品的规则
type TmallProductAddSchemaGetRequest struct {
    
    // BrandId special品牌ID
    BrandId int64 `json:"brand_id";xml:"brand_id"`
    
    // CategoryId required商品发布的目标类目，必须是叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TmallProductAddSchemaGetRequest) GetAPIName() string {
	return "tmall.product.add.schema.get"
}

func (req *TmallProductAddSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["BrandId"] = req.BrandId
    
    params["CategoryId"] = req.CategoryId
    
    return params
}

// TmallProductAddSchemaGetResponse 获取用户发布产品的规则
type TmallProductAddSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddProductRule Basic返回发布产品的规则文档
    AddProductRule string `json:"add_product_rule";xml:"add_product_rule"`
    
}

// TaobaoQimenReturnapplyReportRequest 退货异常包裹单通知接口
type TaobaoQimenReturnapplyReportRequest struct {
    
    // Request optional
    Request Request `json:"request";xml:"request"`
    
}

func (req *TaobaoQimenReturnapplyReportRequest) GetAPIName() string {
	return "taobao.qimen.returnapply.report"
}

func (req *TaobaoQimenReturnapplyReportRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Request"] = req.Request
    
    return params
}

// TaobaoQimenReturnapplyReportResponse 退货异常包裹单通知接口
type TaobaoQimenReturnapplyReportResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Response Object
    Response Response `json:"response";xml:"response"`
    
}

// QimenTaobaoAlphaxOpenGetExpressRequest 调用接口 从isv获取  商家配置物流公司名单
type QimenTaobaoAlphaxOpenGetExpressRequest struct {
    
    // SellerId required卖家id
    SellerId string `json:"seller_id";xml:"seller_id"`
    
    // SellerNick required卖家名称
    SellerNick string `json:"seller_nick";xml:"seller_nick"`
    
}

func (req *QimenTaobaoAlphaxOpenGetExpressRequest) GetAPIName() string {
	return "qimen.taobao.alphax.open.get.express"
}

func (req *QimenTaobaoAlphaxOpenGetExpressRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["SellerId"] = req.SellerId
    
    params["SellerNick"] = req.SellerNick
    
    return params
}

// QimenTaobaoAlphaxOpenGetExpressResponse 调用接口 从isv获取  商家配置物流公司名单
type QimenTaobaoAlphaxOpenGetExpressResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Object返回值
    Result ResultDO `json:"result";xml:"result"`
    
}

// TaobaoFenxiaoDealerRequisitionorderCloseRequest 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseRequest struct {
    
    // DealerOrderId required经销采购单编号
    DealerOrderId int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
    // Reason required关闭原因：
    // 1：长时间无法联系到分销商，取消交易。
    // 2：分销商错误提交申请，取消交易。
    // 3：缺货，近段时间都无法发货。
    // 4：分销商恶意提交申请单。
    // 5：其他原因。
    Reason int64 `json:"reason";xml:"reason"`
    
    // ReasonDetail required关闭详细原因，字数5-200字
    ReasonDetail string `json:"reason_detail";xml:"reason_detail"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}

func (req *TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["DealerOrderId"] = req.DealerOrderId
    
    params["Reason"] = req.Reason
    
    params["ReasonDetail"] = req.ReasonDetail
    
    return params
}

// TaobaoFenxiaoDealerRequisitionorderCloseResponse 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功。true：成功；false：失败。
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TmallItemAddSchemaGetRequest 通过类目以及productId获取商品发布规则；
type TmallItemAddSchemaGetRequest struct {
    
    // CategoryId required商品发布的目标类目，必须是叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // IsvInit optional正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
    IsvInit bool `json:"isv_init";xml:"isv_init"`
    
    // ProductId required商品发布的目标product_id
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // Type optional发布商品类型，一口价填“b”，拍卖填"a"
    Type string `json:"type";xml:"type"`
    
}

func (req *TmallItemAddSchemaGetRequest) GetAPIName() string {
	return "tmall.item.add.schema.get"
}

func (req *TmallItemAddSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["CategoryId"] = req.CategoryId
    
    params["IsvInit"] = req.IsvInit
    
    params["ProductId"] = req.ProductId
    
    params["Type"] = req.Type
    
    return params
}

// TmallItemAddSchemaGetResponse 通过类目以及productId获取商品发布规则；
type TmallItemAddSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddItemResult Basic返回发布商品的规则文档
    AddItemResult string `json:"add_item_result";xml:"add_item_result"`
    
}

// TaobaoFenxiaoDealerRequisitionorderAgreeRequest 供应商或分销商通过采购申请/经销采购单审核
type TaobaoFenxiaoDealerRequisitionorderAgreeRequest struct {
    
    // DealerOrderId required采购申请/经销采购单编号
    DealerOrderId int64 `json:"dealer_order_id";xml:"dealer_order_id"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}

func (req *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["DealerOrderId"] = req.DealerOrderId
    
    return params
}

// TaobaoFenxiaoDealerRequisitionorderAgreeResponse 供应商或分销商通过采购申请/经销采购单审核
type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic操作是否成功。true：成功；false：失败。
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TmallProductSchemaMatchRequest 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallProductSchemaMatchRequest struct {
    
    // CategoryId required商品发布的目标类目，必须是叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // Propvalues required根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
    Propvalues string `json:"propvalues";xml:"propvalues"`
    
}

func (req *TmallProductSchemaMatchRequest) GetAPIName() string {
	return "tmall.product.schema.match"
}

func (req *TmallProductSchemaMatchRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CategoryId"] = req.CategoryId
    
    params["Propvalues"] = req.Propvalues
    
    return params
}

// TmallProductSchemaMatchResponse 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallProductSchemaMatchResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // MatchResult Basic返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。
    MatchResult string `json:"match_result";xml:"match_result"`
    
}

// TmallProductMatchSchemaGetRequest ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallProductMatchSchemaGetRequest struct {
    
    // CategoryId required商品发布的目标类目，必须是叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TmallProductMatchSchemaGetRequest) GetAPIName() string {
	return "tmall.product.match.schema.get"
}

func (req *TmallProductMatchSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["CategoryId"] = req.CategoryId
    
    return params
}

// TmallProductMatchSchemaGetResponse ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallProductMatchSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // MatchResult Basic返回匹配product的规则文档
    MatchResult string `json:"match_result";xml:"match_result"`
    
}

// TmallItemSchemaAddRequest 天猫TopSchema发布商品。
type TmallItemSchemaAddRequest struct {
    
    // CategoryId required商品发布的目标类目，必须是叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // ProductId required发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId
    ProductId int64 `json:"product_id";xml:"product_id"`
    
    // XmlData required根据tmall.item.add.schema.get生成的商品发布规则入参数据
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TmallItemSchemaAddRequest) GetAPIName() string {
	return "tmall.item.schema.add"
}

func (req *TmallItemSchemaAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 3)
    
    params["CategoryId"] = req.CategoryId
    
    params["ProductId"] = req.ProductId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TmallItemSchemaAddResponse 天猫TopSchema发布商品。
type TmallItemSchemaAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddItemResult Basic返回商品发布结果
    AddItemResult string `json:"add_item_result";xml:"add_item_result"`
    
    // GmtCreate Basic发布商品操作成功时间
    GmtCreate time.Time `json:"gmt_create";xml:"gmt_create"`
    
}

// TaobaoFenxiaoDealerRequisitionorderGetRequest 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
type TaobaoFenxiaoDealerRequisitionorderGetRequest struct {
    
    // EndDate required采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天
    EndDate time.Time `json:"end_date";xml:"end_date"`
    
    // Fields optional多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
    Fields string `json:"fields";xml:"fields"`
    
    // Identity required查询者自己在所要查询的采购申请/经销采购单中的身份。
    // 1：供应商。
    // 2：分销商。
    // 注：填写其他值当做错误处理。
    Identity int64 `json:"identity";xml:"identity"`
    
    // OrderStatus optional采购申请/经销采购单状态。
    // 0：全部状态。
    // 1：分销商提交申请，待供应商审核。
    // 2：供应商驳回申请，待分销商确认。
    // 3：供应商修改后，待分销商确认。
    // 4：分销商拒绝修改，待供应商再审核。
    // 5：审核通过下单成功，待分销商付款。
    // 7：付款成功，待供应商发货。
    // 8：供应商发货，待分销商收货。
    // 9：分销商收货，交易成功。
    // 10：采购申请/经销采购单关闭。
    // 
    // 注：无值按默认值0计，超出状态范围返回错误信息。
    OrderStatus int64 `json:"order_status";xml:"order_status"`
    
    // PageNo optional页码（大于0的整数。无值或小于1的值按默认值1计）
    PageNo int64 `json:"page_no";xml:"page_no"`
    
    // PageSize optional每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
    PageSize int64 `json:"page_size";xml:"page_size"`
    
    // StartDate required采购申请/经销采购单最早修改时间
    StartDate time.Time `json:"start_date";xml:"start_date"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderGetRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.get"
}

func (req *TaobaoFenxiaoDealerRequisitionorderGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["EndDate"] = req.EndDate
    
    params["Fields"] = req.Fields
    
    params["Identity"] = req.Identity
    
    params["OrderStatus"] = req.OrderStatus
    
    params["PageNo"] = req.PageNo
    
    params["PageSize"] = req.PageSize
    
    params["StartDate"] = req.StartDate
    
    return params
}

// TaobaoFenxiaoDealerRequisitionorderGetResponse 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
type TaobaoFenxiaoDealerRequisitionorderGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DealerOrders Object Array采购申请/经销采购单结果列表
    DealerOrders DealerOrder `json:"dealer_orders";xml:"dealer_orders"`
    
    // TotalResults Basic按查询条件查到的记录总数
    TotalResults int64 `json:"total_results";xml:"total_results"`
    
}

// TaobaoQimenOrderstatusUpdateRequest 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
type TaobaoQimenOrderstatusUpdateRequest struct {
    
    // ActionTime optional事件发生时间
    ActionTime string `json:"action_time";xml:"action_time"`
    
    // AllocationCode required星盘派单号
    AllocationCode string `json:"allocation_code";xml:"allocation_code"`
    
    // Operator required操作人
    Operator string `json:"operator";xml:"operator"`
    
    // OrderCodes required淘系子订单号
    OrderCodes int64 `json:"order_codes";xml:"order_codes"`
    
    // Status required订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY
    Status string `json:"status";xml:"status"`
    
    // StoreId required目标门店的商户中心门店编码
    StoreId int64 `json:"store_id";xml:"store_id"`
    
    // Type required业务类型，（枚举值：FAHUO、ZITI）
    Type string `json:"type";xml:"type"`
    
}

func (req *TaobaoQimenOrderstatusUpdateRequest) GetAPIName() string {
	return "taobao.qimen.orderstatus.update"
}

func (req *TaobaoQimenOrderstatusUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 7)
    
    params["ActionTime"] = req.ActionTime
    
    params["AllocationCode"] = req.AllocationCode
    
    params["Operator"] = req.Operator
    
    params["OrderCodes"] = req.OrderCodes
    
    params["Status"] = req.Status
    
    params["StoreId"] = req.StoreId
    
    params["Type"] = req.Type
    
    return params
}

// TaobaoQimenOrderstatusUpdateResponse 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
type TaobaoQimenOrderstatusUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basicsuccess
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
    // Message Basicmessage
    Message string `json:"message";xml:"message"`
    
    // ResultCode BasicresultCode
    ResultCode string `json:"result_code";xml:"result_code"`
    
}

// TaobaoFenxiaoDealerRequisitionorderQueryRequest 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
type TaobaoFenxiaoDealerRequisitionorderQueryRequest struct {
    
    // DealerOrderIds required经销采购单编号。
    // 多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。
    DealerOrderIds int64 `json:"dealer_order_ids";xml:"dealer_order_ids"`
    
    // Fields optional多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
    Fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetAPIName() string {
	return "taobao.fenxiao.dealer.requisitionorder.query"
}

func (req *TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["DealerOrderIds"] = req.DealerOrderIds
    
    params["Fields"] = req.Fields
    
    return params
}

// TaobaoFenxiaoDealerRequisitionorderQueryResponse 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
type TaobaoFenxiaoDealerRequisitionorderQueryResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // DealerOrders Object Array经销采购单结果列表
    DealerOrders DealerOrder `json:"dealer_orders";xml:"dealer_orders"`
    
}

// TaobaoQimenEventProduceRequest 当订单被处理时，用于通知奇门系统。
type TaobaoQimenEventProduceRequest struct {
    
    // Create optional订单创建时间,数字
    Create int64 `json:"create";xml:"create"`
    
    // Ext optionalJSON格式扩展字段
    Ext string `json:"ext";xml:"ext"`
    
    // Nick optional外部商家名称。必须同时填写platform
    Nick string `json:"nick";xml:"nick"`
    
    // Platform optional商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
    Platform string `json:"platform";xml:"platform"`
    
    // Status required事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
    Status string `json:"status";xml:"status"`
    
    // Tid required淘宝订单号
    Tid string `json:"tid";xml:"tid"`
    
}

func (req *TaobaoQimenEventProduceRequest) GetAPIName() string {
	return "taobao.qimen.event.produce"
}

func (req *TaobaoQimenEventProduceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 6)
    
    params["Create"] = req.Create
    
    params["Ext"] = req.Ext
    
    params["Nick"] = req.Nick
    
    params["Platform"] = req.Platform
    
    params["Status"] = req.Status
    
    params["Tid"] = req.Tid
    
    return params
}

// TaobaoQimenEventProduceResponse 当订单被处理时，用于通知奇门系统。
type TaobaoQimenEventProduceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsSuccess Basic是否成功
    IsSuccess bool `json:"is_success";xml:"is_success"`
    
}

// TaobaoQimenEventsProduceRequest 批量发送消息
type TaobaoQimenEventsProduceRequest struct {
    
    // Messages required奇门事件列表, 最多50条
    Messages QimenEvent `json:"messages";xml:"messages"`
    
}

func (req *TaobaoQimenEventsProduceRequest) GetAPIName() string {
	return "taobao.qimen.events.produce"
}

func (req *TaobaoQimenEventsProduceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Messages"] = req.Messages
    
    return params
}

// TaobaoQimenEventsProduceResponse 批量发送消息
type TaobaoQimenEventsProduceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // IsAllSuccess Basic是否全部成功
    IsAllSuccess bool `json:"is_all_success";xml:"is_all_success"`
    
    // Results Object Array发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
    Results QimenResult `json:"results";xml:"results"`
    
}

// TaobaoFenxiaoTradePrepayOfflineReduceRequest 渠道分销供应商上传线下流水预存款（减少）
type TaobaoFenxiaoTradePrepayOfflineReduceRequest struct {
    
    // OfflineReducePrepayParam required减少流水
    OfflineReducePrepayParam TopOfflineReducePrepayDto `json:"offline_reduce_prepay_param";xml:"offline_reduce_prepay_param"`
    
}

func (req *TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetAPIName() string {
	return "taobao.fenxiao.trade.prepay.offline.reduce"
}

func (req *TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["OfflineReducePrepayParam"] = req.OfflineReducePrepayParam
    
    return params
}

// TaobaoFenxiaoTradePrepayOfflineReduceResponse 渠道分销供应商上传线下流水预存款（减少）
type TaobaoFenxiaoTradePrepayOfflineReduceResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result ResultTopDo `json:"result";xml:"result"`
    
}

// TaobaoQimenTradeUserAddRequest 添加奇门订单链路用户
type TaobaoQimenTradeUserAddRequest struct {
    
    // Memo optional商家备注
    Memo string `json:"memo";xml:"memo"`
    
}

func (req *TaobaoQimenTradeUserAddRequest) GetAPIName() string {
	return "taobao.qimen.trade.user.add"
}

func (req *TaobaoQimenTradeUserAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["Memo"] = req.Memo
    
    return params
}

// TaobaoQimenTradeUserAddResponse 添加奇门订单链路用户
type TaobaoQimenTradeUserAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Appkey Basicappkey
    Appkey string `json:"appkey";xml:"appkey"`
    
    // GmtCreate Basic创建时间
    GmtCreate time.Time `json:"gmt_create";xml:"gmt_create"`
    
    // Memo Basic卖家备注
    Memo string `json:"memo";xml:"memo"`
    
}

// TaobaoFenxiaoTradePrepayOfflineAddRequest 渠道分销供应商上传线下流水预存款（增加）
type TaobaoFenxiaoTradePrepayOfflineAddRequest struct {
    
    // OfflineAddPrepayParam required增加流水
    OfflineAddPrepayParam TopOfflineAddPrepayDto `json:"offline_add_prepay_param";xml:"offline_add_prepay_param"`
    
}

func (req *TaobaoFenxiaoTradePrepayOfflineAddRequest) GetAPIName() string {
	return "taobao.fenxiao.trade.prepay.offline.add"
}

func (req *TaobaoFenxiaoTradePrepayOfflineAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["OfflineAddPrepayParam"] = req.OfflineAddPrepayParam
    
    return params
}

// TaobaoFenxiaoTradePrepayOfflineAddResponse 渠道分销供应商上传线下流水预存款（增加）
type TaobaoFenxiaoTradePrepayOfflineAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Result Objectresult
    Result ResultTopDo `json:"result";xml:"result"`
    
}

// TaobaoQimenTradeUserDeleteRequest 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteRequest struct {
    
}

func (req *TaobaoQimenTradeUserDeleteRequest) GetAPIName() string {
	return "taobao.qimen.trade.user.delete"
}

func (req *TaobaoQimenTradeUserDeleteRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
    
    return params
}

// TaobaoQimenTradeUserDeleteResponse 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // Modal Basicmodal
    Modal bool `json:"modal";xml:"modal"`
    
}

// TaobaoQimenTradeUsersGetRequest 获取已开通奇门订单服务的用户列表
type TaobaoQimenTradeUsersGetRequest struct {
    
    // PageIndex required每页的数量
    PageIndex int64 `json:"page_index";xml:"page_index"`
    
    // PageSize required页数
    PageSize int64 `json:"page_size";xml:"page_size"`
    
}

func (req *TaobaoQimenTradeUsersGetRequest) GetAPIName() string {
	return "taobao.qimen.trade.users.get"
}

func (req *TaobaoQimenTradeUsersGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["PageIndex"] = req.PageIndex
    
    params["PageSize"] = req.PageSize
    
    return params
}

// TaobaoQimenTradeUsersGetResponse 获取已开通奇门订单服务的用户列表
type TaobaoQimenTradeUsersGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // TotalCount BasictotalCount
    TotalCount int64 `json:"total_count";xml:"total_count"`
    
    // Users Object Arraymodal
    Users QimenUser `json:"users";xml:"users"`
    
}

// TaobaoItemSchemaUpdateRequest 新模式下的商品编辑接口，在调用该接口的时候，需要先调用taobao.item.update.rules.get接口获取编辑规则和数据集。然后按照约定的参数传入规则，调用该接口进行编辑商品
type TaobaoItemSchemaUpdateRequest struct {
    
    // CategoryId optional如果重新选择类目后，传入重新选择的叶子类目id
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // Incremental optional是否增量更新。为true只改xml_data里传入的值。为false时，将根据xml_data的数据对原始宝贝数据进行全量覆盖,这个参数暂时不支持
    Incremental bool `json:"incremental";xml:"incremental"`
    
    // ItemId required编辑商品的商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
    // XmlData required编辑商品时候的修改内容
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TaobaoItemSchemaUpdateRequest) GetAPIName() string {
	return "taobao.item.schema.update"
}

func (req *TaobaoItemSchemaUpdateRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 4)
    
    params["CategoryId"] = req.CategoryId
    
    params["Incremental"] = req.Incremental
    
    params["ItemId"] = req.ItemId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TaobaoItemSchemaUpdateResponse 新模式下的商品编辑接口，在调用该接口的时候，需要先调用taobao.item.update.rules.get接口获取编辑规则和数据集。然后按照约定的参数传入规则，调用该接口进行编辑商品
type TaobaoItemSchemaUpdateResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateResult Basic编辑商品返回的结果信息，暂时只返回 itemId
    UpdateResult string `json:"update_result";xml:"update_result"`
    
}

// TaobaoItemAddSchemaGetRequest 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能
type TaobaoItemAddSchemaGetRequest struct {
    
    // CategoryId required发布宝贝的叶子类目id
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
}

func (req *TaobaoItemAddSchemaGetRequest) GetAPIName() string {
	return "taobao.item.add.schema.get"
}

func (req *TaobaoItemAddSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 1)
    
    params["CategoryId"] = req.CategoryId
    
    return params
}

// TaobaoItemAddSchemaGetResponse 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能
type TaobaoItemAddSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddRules Basic返回结果的集合
    AddRules string `json:"add_rules";xml:"add_rules"`
    
}

// TaobaoItemSchemaAddRequest isv将宝贝信息，组装成特定格式的xml数据作为参数，进行发布商品
type TaobaoItemSchemaAddRequest struct {
    
    // CategoryId required当前发布的叶子类目
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // XmlData required将需要发布的商品数据组装成的xml格式数据
    XmlData string `json:"xml_data";xml:"xml_data"`
    
}

func (req *TaobaoItemSchemaAddRequest) GetAPIName() string {
	return "taobao.item.schema.add"
}

func (req *TaobaoItemSchemaAddRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CategoryId"] = req.CategoryId
    
    params["XmlData"] = req.XmlData
    
    return params
}

// TaobaoItemSchemaAddResponse isv将宝贝信息，组装成特定格式的xml数据作为参数，进行发布商品
type TaobaoItemSchemaAddResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // AddResult Basic返回的结果
    AddResult string `json:"add_result";xml:"add_result"`
    
}

// TaobaoItemUpdateSchemaGetRequest 在新的编辑模式下，isv需要先获取一份编辑商品的规则和数据，然后根据规则传入数据。该接口提供编辑规则查询功能
type TaobaoItemUpdateSchemaGetRequest struct {
    
    // CategoryId optional商品类目id
    CategoryId int64 `json:"category_id";xml:"category_id"`
    
    // ItemId required商品id
    ItemId int64 `json:"item_id";xml:"item_id"`
    
}

func (req *TaobaoItemUpdateSchemaGetRequest) GetAPIName() string {
	return "taobao.item.update.schema.get"
}

func (req *TaobaoItemUpdateSchemaGetRequest) GetParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
    
    params["CategoryId"] = req.CategoryId
    
    params["ItemId"] = req.ItemId
    
    return params
}

// TaobaoItemUpdateSchemaGetResponse 在新的编辑模式下，isv需要先获取一份编辑商品的规则和数据，然后根据规则传入数据。该接口提供编辑规则查询功能
type TaobaoItemUpdateSchemaGetResponse struct {
    ErrorResponse ErrorResponse `json:"error_response",xml:"error_response"`
    
    // UpdateRules Basic返回的结果集
    UpdateRules string `json:"update_rules";xml:"update_rules"`
    
}

