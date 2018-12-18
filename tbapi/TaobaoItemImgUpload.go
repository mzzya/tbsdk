package tbsdk

// TaobaoItemImgUploadRequest 添加一张商品图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
type TaobaoItemImgUploadRequest struct {
    
    /* id optional商品图片id(如果是更新图片，则需要传该参数) */
    id int64 `json:"id";xml:"id"`
    
    /* image optional商品图片内容类型:JPG,GIF;最大:3M 。支持的文件类型：gif,jpg,jpeg,png */
    image byte[] `json:"image";xml:"image"`
    
    /* is_major optional是否将该图片设为主图,可选值:true,false;默认值:false(非主图) */
    is_major bool `json:"is_major";xml:"is_major"`
    
    /* is_rectangle optional是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图 */
    is_rectangle bool `json:"is_rectangle";xml:"is_rectangle"`
    
    /* num_iid required商品数字ID，该参数必须 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* position optional图片序号 */
    position int64 `json:"position";xml:"position"`
    
}

func (req *TaobaoItemImgUploadRequest) GetAPIName() string {
	return "taobao.item.img.upload"
}

// TaobaoItemImgUploadResponse 添加一张商品图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
type TaobaoItemImgUploadResponse struct {
    
    /* item_img Object商品图片结构 */
    item_img ItemImg `json:"item_img";xml:"item_img"`
    
}
