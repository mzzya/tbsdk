package tbsdk

// TaobaoItemSkuDeleteRequest 删除一个sku的数据
需要删除的sku通过属性properties进行匹配查找
type TaobaoItemSkuDeleteRequest struct {
    
    /* ignorewarning optional忽略警告提示. */
    ignorewarning string `json:"ignorewarning";xml:"ignorewarning"`
    
    /* item_num optionalsku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败 */
    item_num int64 `json:"item_num";xml:"item_num"`
    
    /* item_price optionalsku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功 */
    item_price Price `json:"item_price";xml:"item_price"`
    
    /* lang optionalSku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
    lang string `json:"lang";xml:"lang"`
    
    /* num_iid requiredSku所属商品数字id，可通过 taobao.item.get 获取。必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* properties requiredSku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemSkuDeleteRequest) GetAPIName() string {
	return "taobao.item.sku.delete"
}

// TaobaoItemSkuDeleteResponse 删除一个sku的数据
需要删除的sku通过属性properties进行匹配查找
type TaobaoItemSkuDeleteResponse struct {
    
    /* sku ObjectSku结构 */
    sku Sku `json:"sku";xml:"sku"`
    
}
