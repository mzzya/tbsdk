package tbsdk

// TaobaoItemPropimgUploadRequest 添加一张商品属性图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 
商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoItemPropimgUploadRequest struct {
    
    /* id optional属性图片ID。如果是新增不需要填写 */
    id int64 `json:"id";xml:"id"`
    
    /* image optional属性图片内容。类型:JPG,GIF;图片大小不超过:3M */
    image byte[] `json:"image";xml:"image"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* position optional图片位置 */
    position int64 `json:"position";xml:"position"`
    
    /* properties required属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemPropimgUploadRequest) GetAPIName() string {
	return "taobao.item.propimg.upload"
}

// TaobaoItemPropimgUploadResponse 添加一张商品属性图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 
商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoItemPropimgUploadResponse struct {
    
    /* prop_img ObjectPropImg属性图片结构 */
    prop_img PropImg `json:"prop_img";xml:"prop_img"`
    
}
