package tbsdk

// TaobaoItemDeleteRequest 删除单条商品
type TaobaoItemDeleteRequest struct {
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemDeleteRequest) GetAPIName() string {
	return "taobao.item.delete"
}

// TaobaoItemDeleteResponse 删除单条商品
type TaobaoItemDeleteResponse struct {
    
    /* item Object被删除商品的相关信息 */
    item Item `json:"item";xml:"item"`
    
}
