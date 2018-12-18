package tbsdk

// TaobaoWlbWmsSkuUpdateRequest 商品信息的更新
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

// TaobaoWlbWmsSkuUpdateResponse 商品信息的更新
type TaobaoWlbWmsSkuUpdateResponse struct {
    
    /* wl_error_code Basic错误编码 */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg Basic错误信息 */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basic是否成功 */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
