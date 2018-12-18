package tbsdk

// TaobaoScitemAddRequest 发布后端商品
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

// TaobaoScitemAddResponse 发布后端商品
type TaobaoScitemAddResponse struct {
    
    /* sc_item Object后台商品信息 */
    sc_item ScItem `json:"sc_item";xml:"sc_item"`
    
}
