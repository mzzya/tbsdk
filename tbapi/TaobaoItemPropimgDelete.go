package tbsdk

// TaobaoItemPropimgDeleteRequest 删除propimg_id 所指定的商品属性图片 
传入的num_iid所对应的商品必须属于当前会话的用户 
propimg_id对应的属性图片需要属于num_iid对应的商品
type TaobaoItemPropimgDeleteRequest struct {
    
    /* id required商品属性图片ID */
    id int64 `json:"id";xml:"id"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemPropimgDeleteRequest) GetAPIName() string {
	return "taobao.item.propimg.delete"
}

// TaobaoItemPropimgDeleteResponse 删除propimg_id 所指定的商品属性图片 
传入的num_iid所对应的商品必须属于当前会话的用户 
propimg_id对应的属性图片需要属于num_iid对应的商品
type TaobaoItemPropimgDeleteResponse struct {
    
    /* prop_img Object属性图片结构 */
    prop_img PropImg `json:"prop_img";xml:"prop_img"`
    
}
