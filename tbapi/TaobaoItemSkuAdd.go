package tbsdk

// TaobaoItemSkuAddRequest 新增一个sku到num_iid指定的商品中 
传入的iid所对应的商品必须属于当前会话的用户
type TaobaoItemSkuAddRequest struct {
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* item_price optionalsku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功 */
    item_price Price `json:"item_price";xml:"item_price"`
    
    /* lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
    lang string `json:"lang";xml:"lang"`
    
    /* num_iid requiredSku所属商品数字id，可通过 taobao.item.get 获取。必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* outer_id optionalSku的商家外部id */
    outer_id string `json:"outer_id";xml:"outer_id"`
    
    /* price requiredSku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
    price Price `json:"price";xml:"price"`
    
    /* properties requiredSku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity requiredSku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数 */
    quantity int64 `json:"quantity";xml:"quantity"`
    
}

func (req *TaobaoItemSkuAddRequest) GetAPIName() string {
	return "taobao.item.sku.add"
}

// TaobaoItemSkuAddResponse 新增一个sku到num_iid指定的商品中 
传入的iid所对应的商品必须属于当前会话的用户
type TaobaoItemSkuAddResponse struct {
    
    /* sku Objectsku */
    sku Sku `json:"sku";xml:"sku"`
    
}
