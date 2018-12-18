package tbsdk

// TaobaoItemJointImgRequest * 关联一张商品图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
type TaobaoItemJointImgRequest struct {
    
    /* id optional商品图片id(如果是更新图片，则需要传该参数) */
    id int64 `json:"id";xml:"id"`
    
    /* is_major optional上传的图片是否关联为商品主图 */
    is_major bool `json:"is_major";xml:"is_major"`
    
    /* is_rectangle optional是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图 */
    is_rectangle bool `json:"is_rectangle";xml:"is_rectangle"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* pic_path required图片URL,图片空间图片的相对地址 */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* position optional图片序号 */
    position int64 `json:"position";xml:"position"`
    
}

func (req *TaobaoItemJointImgRequest) GetAPIName() string {
	return "taobao.item.joint.img"
}

// TaobaoItemJointImgResponse * 关联一张商品图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
type TaobaoItemJointImgResponse struct {
    
    /* item_img Object商品图片信息 */
    item_img ItemImg `json:"item_img";xml:"item_img"`
    
}
