package tbsdk

// TaobaoScitemUpdateRequest 根据商品ID或商家编码修改后端商品
type TaobaoScitemUpdateRequest struct {
    
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
    
    /* is_dangerous optional是否危险 0：不是  0：是 */
    is_dangerous int64 `json:"is_dangerous";xml:"is_dangerous"`
    
    /* is_friable optional是否易碎 0：不是  1：是 */
    is_friable int64 `json:"is_friable";xml:"is_friable"`
    
    /* is_warranty optional是否保质期：0:不是 1：是 */
    is_warranty int64 `json:"is_warranty";xml:"is_warranty"`
    
    /* item_id optional后端商品ID，跟outer_code必须指定一个 */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* item_name optional商品名称 */
    item_name string `json:"item_name";xml:"item_name"`
    
    /* item_type optional0.普通供应链商品 1.供应链组合主商品 */
    item_type int64 `json:"item_type";xml:"item_type"`
    
    /* length optional长度 单位：mm */
    length int64 `json:"length";xml:"length"`
    
    /* matter_status optional0:液体，1：粉体，2：固体 */
    matter_status int64 `json:"matter_status";xml:"matter_status"`
    
    /* outer_code special商家编码，跟item_id必须指定一个 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* price optionalprice */
    price int64 `json:"price";xml:"price"`
    
    /* remark optionalremark */
    remark string `json:"remark";xml:"remark"`
    
    /* remove_properties optional移除商品属性P列表,P由系统分配：p1；p2 */
    remove_properties string `json:"remove_properties";xml:"remove_properties"`
    
    /* spu_id optional淘宝SKU产品级编码CSPU ID */
    spu_id int64 `json:"spu_id";xml:"spu_id"`
    
    /* update_properties optional需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3 */
    update_properties string `json:"update_properties";xml:"update_properties"`
    
    /* volume optional体积：立方厘米 */
    volume int64 `json:"volume";xml:"volume"`
    
    /* weight optionalweight */
    weight int64 `json:"weight";xml:"weight"`
    
    /* width optional宽 单位：mm */
    width int64 `json:"width";xml:"width"`
    
    /* wms_code optional仓储商编码 */
    wms_code string `json:"wms_code";xml:"wms_code"`
    
}

func (req *TaobaoScitemUpdateRequest) GetAPIName() string {
	return "taobao.scitem.update"
}

// TaobaoScitemUpdateResponse 根据商品ID或商家编码修改后端商品
type TaobaoScitemUpdateResponse struct {
    
    /* update_rows Basic更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据 */
    update_rows int64 `json:"update_rows";xml:"update_rows"`
    
}
