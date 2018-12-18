package tbsdk

// TaobaoWlbItemAddRequest 添加物流宝商品，支持物流宝子商品和属性添加
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

// TaobaoWlbItemAddResponse 添加物流宝商品，支持物流宝子商品和属性添加
type TaobaoWlbItemAddResponse struct {
    
    /* item_id Basic新增的商品 */
    item_id map[string]interface{} `json:"item_id";xml:"item_id"`
    
}
