package tbsdk

// TaobaoItemAddRequest 此接口用于新增一个商品 
商品所属的卖家是当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
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
    
    /* delivery_time.delivery_time optional商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11 */
    delivery_time.delivery_time string `json:"delivery_time.delivery_time";xml:"delivery_time.delivery_time"`
    
    /* delivery_time.delivery_time_type optional发货时间类型：绝对发货时间或者相对发货时间 */
    delivery_time.delivery_time_type string `json:"delivery_time.delivery_time_type";xml:"delivery_time.delivery_time_type"`
    
    /* delivery_time.need_delivery_time optional设置是否使用发货时间，商品级别，sku级别 */
    delivery_time.need_delivery_time string `json:"delivery_time.need_delivery_time";xml:"delivery_time.need_delivery_time"`
    
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
    
    /* food_security.contact optional厂家联系方式 */
    food_security.contact string `json:"food_security.contact";xml:"food_security.contact"`
    
    /* food_security.design_code optional产品标准号 */
    food_security.design_code string `json:"food_security.design_code";xml:"food_security.design_code"`
    
    /* food_security.factory optional厂名 */
    food_security.factory string `json:"food_security.factory";xml:"food_security.factory"`
    
    /* food_security.factory_site optional厂址 */
    food_security.factory_site string `json:"food_security.factory_site";xml:"food_security.factory_site"`
    
    /* food_security.food_additive optional食品添加剂 */
    food_security.food_additive string `json:"food_security.food_additive";xml:"food_security.food_additive"`
    
    /* food_security.health_product_no optional健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
    food_security.health_product_no string `json:"food_security.health_product_no";xml:"food_security.health_product_no"`
    
    /* food_security.mix optional配料表 */
    food_security.mix string `json:"food_security.mix";xml:"food_security.mix"`
    
    /* food_security.period optional保质期 */
    food_security.period string `json:"food_security.period";xml:"food_security.period"`
    
    /* food_security.plan_storage optional储藏方法 */
    food_security.plan_storage string `json:"food_security.plan_storage";xml:"food_security.plan_storage"`
    
    /* food_security.prd_license_no optional生产许可证号 */
    food_security.prd_license_no string `json:"food_security.prd_license_no";xml:"food_security.prd_license_no"`
    
    /* food_security.product_date_end optional生产结束日期,格式必须为yyyy-MM-dd */
    food_security.product_date_end string `json:"food_security.product_date_end";xml:"food_security.product_date_end"`
    
    /* food_security.product_date_start optional生产开始日期，格式必须为yyyy-MM-dd */
    food_security.product_date_start string `json:"food_security.product_date_start";xml:"food_security.product_date_start"`
    
    /* food_security.stock_date_end optional进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
    food_security.stock_date_end string `json:"food_security.stock_date_end";xml:"food_security.stock_date_end"`
    
    /* food_security.stock_date_start optional进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
    food_security.stock_date_start string `json:"food_security.stock_date_start";xml:"food_security.stock_date_start"`
    
    /* food_security.supplier optional供货商 */
    food_security.supplier string `json:"food_security.supplier";xml:"food_security.supplier"`
    
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
    image byte[] `json:"image";xml:"image"`
    
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
    list_time Date `json:"list_time";xml:"list_time"`
    
    /* locality_life.choose_logis optional发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
    locality_life.choose_logis string `json:"locality_life.choose_logis";xml:"locality_life.choose_logis"`
    
    /* locality_life.eticket optional电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4 */
    locality_life.eticket string `json:"locality_life.eticket";xml:"locality_life.eticket"`
    
    /* locality_life.expirydate optional本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15 */
    locality_life.expirydate string `json:"locality_life.expirydate";xml:"locality_life.expirydate"`
    
    /* locality_life.merchant optional码商信息，格式为 码商id:nick */
    locality_life.merchant string `json:"locality_life.merchant";xml:"locality_life.merchant"`
    
    /* locality_life.network_id optional网点ID */
    locality_life.network_id string `json:"locality_life.network_id";xml:"locality_life.network_id"`
    
    /* locality_life.obs optional预约门店是否支持门店自提,1:是 */
    locality_life.obs string `json:"locality_life.obs";xml:"locality_life.obs"`
    
    /* locality_life.onsale_auto_refund_ratio optional电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
    locality_life.onsale_auto_refund_ratio int64 `json:"locality_life.onsale_auto_refund_ratio";xml:"locality_life.onsale_auto_refund_ratio"`
    
    /* locality_life.packageid optional新版电子凭证包 id */
    locality_life.packageid string `json:"locality_life.packageid";xml:"locality_life.packageid"`
    
    /* locality_life.refund_ratio optional退款比例，百分比%前的数字,1-100的正整数值 */
    locality_life.refund_ratio int64 `json:"locality_life.refund_ratio";xml:"locality_life.refund_ratio"`
    
    /* locality_life.refundmafee optional退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
    locality_life.refundmafee string `json:"locality_life.refundmafee";xml:"locality_life.refundmafee"`
    
    /* locality_life.verification optional核销打款 1代表核销打款 0代表非核销打款 */
    locality_life.verification string `json:"locality_life.verification";xml:"locality_life.verification"`
    
    /* locality_life.version optional新版电子凭证字段 */
    locality_life.version string `json:"locality_life.version";xml:"locality_life.version"`
    
    /* location.city required所在地城市。如杭州 。 */
    location.city string `json:"location.city";xml:"location.city"`
    
    /* location.state required所在地省份。如浙江 */
    location.state string `json:"location.state";xml:"location.state"`
    
    /* ms_payment.price optional订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm */
    ms_payment.price string `json:"ms_payment.price";xml:"ms_payment.price"`
    
    /* ms_payment.reference_price optional参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm */
    ms_payment.reference_price string `json:"ms_payment.reference_price";xml:"ms_payment.reference_price"`
    
    /* ms_payment.voucher_price optional尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm */
    ms_payment.voucher_price string `json:"ms_payment.voucher_price";xml:"ms_payment.voucher_price"`
    
    /* newprepay optional该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货； */
    newprepay string `json:"newprepay";xml:"newprepay"`
    
    /* num required商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。 */
    num int64 `json:"num";xml:"num"`
    
    /* o2o_bind_service optional汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。 */
    o2o_bind_service bool `json:"o2o_bind_service";xml:"o2o_bind_service"`
    
    /* outer_id optional商品外部编码，该字段的最大长度是64个字节 */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* paimai_info.deposit optional拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。 */
    paimai_info.deposit int64 `json:"paimai_info.deposit";xml:"paimai_info.deposit"`
    
    /* paimai_info.interval optional降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。 */
    paimai_info.interval int64 `json:"paimai_info.interval";xml:"paimai_info.interval"`
    
    /* paimai_info.mode optional拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。 */
    paimai_info.mode int64 `json:"paimai_info.mode";xml:"paimai_info.mode"`
    
    /* paimai_info.reserve optional降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
    paimai_info.reserve Price `json:"paimai_info.reserve";xml:"paimai_info.reserve"`
    
    /* paimai_info.valid_hour optional自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
    paimai_info.valid_hour int64 `json:"paimai_info.valid_hour";xml:"paimai_info.valid_hour"`
    
    /* paimai_info.valid_minute optional自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
    paimai_info.valid_minute int64 `json:"paimai_info.valid_minute";xml:"paimai_info.valid_minute"`
    
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

// TaobaoItemAddResponse 此接口用于新增一个商品 
商品所属的卖家是当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
type TaobaoItemAddResponse struct {
    
    /* item Object商品结构,仅有numIid和created返回 */
    item Item `json:"item";xml:"item"`
    
}
