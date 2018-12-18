package tbsdk

// TaobaoItemUpdateListingRequest * 单个商品上架
* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingRequest struct {
    
    /* num required需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num */
    num int64 `json:"num";xml:"num"`
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemUpdateListingRequest) GetAPIName() string {
	return "taobao.item.update.listing"
}

// TaobaoItemUpdateListingResponse * 单个商品上架
* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingResponse struct {
    
    /* item Object上架后返回的商品信息：返回的结果就是:num_iid和modified */
    item Item `json:"item";xml:"item"`
    
}
