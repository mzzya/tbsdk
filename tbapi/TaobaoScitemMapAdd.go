package tbsdk

// TaobaoScitemMapAddRequest 创建IC商品或分销商品与后端商品的映射关系
type TaobaoScitemMapAddRequest struct {
    
    /* item_id required前台ic商品id */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* need_check optional默认值为false
true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联
false:不进行高级校验 */
    need_check bool `json:"need_check";xml:"need_check"`
    
    /* outer_code optionalsc_item_id和outer_code 其中一个不能为空 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
    /* sc_item_id optionalsc_item_id和outer_code 其中一个不能为空 */
    sc_item_id int64 `json:"sc_item_id";xml:"sc_item_id"`
    
    /* sku_id optional前台ic商品skuid */
    sku_id int64 `json:"sku_id";xml:"sku_id"`
    
}

func (req *TaobaoScitemMapAddRequest) GetAPIName() string {
	return "taobao.scitem.map.add"
}

// TaobaoScitemMapAddResponse 创建IC商品或分销商品与后端商品的映射关系
type TaobaoScitemMapAddResponse struct {
    
    /* outer_code Basic接口调用返回结果信息：商家编码 */
    outer_code string `json:"outer_code";xml:"outer_code"`
    
}
