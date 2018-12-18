package tbsdk

// TaobaoItemJointPropimgRequest * 关联一张商品属性图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
type TaobaoItemJointPropimgRequest struct {
    
    /* id optional属性图片ID。如果是新增不需要填写 */
    id int64 `json:"id";xml:"id"`
    
    /* num_iid required商品数字ID，必选 */
    num_iid int64 `json:"num_iid";xml:"num_iid"`
    
    /* pic_path required图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded ) */
    pic_path string `json:"pic_path";xml:"pic_path"`
    
    /* position optional图片序号 */
    position int64 `json:"position";xml:"position"`
    
    /* properties required属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。 */
    properties string `json:"properties";xml:"properties"`
    
}

func (req *TaobaoItemJointPropimgRequest) GetAPIName() string {
	return "taobao.item.joint.propimg"
}

// TaobaoItemJointPropimgResponse * 关联一张商品属性图片到num_iid指定的商品中
    * 传入的num_iid所对应的商品必须属于当前会话的用户
    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
type TaobaoItemJointPropimgResponse struct {
    
    /* prop_img Object属性图片对象信息 */
    prop_img PropImg `json:"prop_img";xml:"prop_img"`
    
}
