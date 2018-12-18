package tbsdk

// TaobaoItemImgDeleteRequest 删除商品图片
type TaobaoItemImgDeleteRequest struct {
    
    /* id required商品图片ID；如果是竖图，请将id的值设置为1 */
    id int64 `json:"id";xml:"id"`
    
    /* is_sixth_pic optional标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下 */
    is_sixth_pic bool `json:"is_sixth_pic";xml:"is_sixth_pic"`
    
    /* num_iid required商品数字ID */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
}

func (req *TaobaoItemImgDeleteRequest) GetAPIName() string {
	return "taobao.item.img.delete"
}

// TaobaoItemImgDeleteResponse 删除商品图片
type TaobaoItemImgDeleteResponse struct {
    
    /* item_img Object商品图片结构 */
    item_img ItemImg `json:"item_img";xml:"item_img"`
    
}
