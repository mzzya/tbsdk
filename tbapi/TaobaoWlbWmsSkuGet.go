package tbsdk

// TaobaoWlbWmsSkuGetRequest 商品信息查询
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

// TaobaoWlbWmsSkuGetResponse 商品信息查询
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
