package tbsdk

// TaobaoFenxiaoProductUpdateRequest 更新分销平台产品数据，不传更新数据返回失败<br>
1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
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
    image byte[] `json:"image";xml:"image"`
    
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

// TaobaoFenxiaoProductUpdateResponse 更新分销平台产品数据，不传更新数据返回失败<br>
1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
type TaobaoFenxiaoProductUpdateResponse struct {
    
    /* modified Basic更新时间，时间格式：yyyy-MM-dd HH:mm:ss */
    modified Date `json:"modified";xml:"modified"`
    
    /* pid Basic产品ID */
    pid int64 `json:"pid";xml:"pid"`
    
}
