package tbsdk

// TaobaoWlbItemUpdateRequest 修改物流宝商品信息
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

// TaobaoWlbItemUpdateResponse 修改物流宝商品信息
type TaobaoWlbItemUpdateResponse struct {
    
    /* gmt_modified Basic修改时间 */
    gmt_modified bool `json:"gmt_modified";xml:"gmt_modified"`
    
}
