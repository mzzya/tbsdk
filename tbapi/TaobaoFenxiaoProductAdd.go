package tbsdk

// TaobaoFenxiaoProductAddRequest 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。

    * 产品图片默认为空
    * 产品发布后默认为下架状态
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
    image byte[] `json:"image";xml:"image"`
    
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

// TaobaoFenxiaoProductAddResponse 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。

    * 产品图片默认为空
    * 产品发布后默认为下架状态
type TaobaoFenxiaoProductAddResponse struct {
    
    /* created Basic产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss */
    created Date `json:"created";xml:"created"`
    
    /* pid Basic产品ID */
    pid int64 `json:"pid";xml:"pid"`
    
}
