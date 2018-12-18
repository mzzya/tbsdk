package tbsdk

// TaobaoItemUpdateDelistingRequest * 单个商品下架
    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingRequest struct {
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemUpdateDelistingRequest) GetAPIName() string {
	return "taobao.item.update.delisting"
}

// TaobaoItemUpdateDelistingResponse * 单个商品下架
    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingResponse struct {
    
    /* item Object返回商品更新信息：返回的结果是:num_iid和modified */
    item Item `json:"item";xml:"item"`
    
}
