package tbsdk

// TaobaoWlbOrderCreateRequest 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
type TaobaoWlbOrderCreateRequest struct {
    
    /* alipay_no optional支付宝交易号 */
    alipay_no string `json:"alipay_no";xml:"alipay_no"`
    
    /* attributes optional该字段暂时保留 */
    attributes string `json:"attributes";xml:"attributes"`
    
    /* buyer_nick optional买家呢称 */
    buyer_nick string `json:"buyer_nick";xml:"buyer_nick"`
    
    /* expect_end_time optional期望结束时间，在入库单会使用到 */
    expect_end_time Date `json:"expect_end_time";xml:"expect_end_time"`
    
    /* expect_start_time optional计划开始送达时间  在入库单中可能会使用 */
    expect_start_time Date `json:"expect_start_time";xml:"expect_start_time"`
    
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

// TaobaoWlbOrderCreateResponse 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
type TaobaoWlbOrderCreateResponse struct {
    
    /* create_time Basic订单创建时间 */
    create_time Date `json:"create_time";xml:"create_time"`
    
    /* order_code Basic物流宝订单创建成功后，返回物流宝的订单编号；如果订单创建失败，该字段为空。 */
    order_code string `json:"order_code";xml:"order_code"`
    
}
