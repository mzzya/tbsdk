package tbsdk

// TaobaoOcTradetagAttachRequest 对订单添加标签和更新标签。标签分为官方标签和自定义标签。
官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731
自定义标签有2个通用属性:
    `show_str:给消费者显示的字符串（如果可以显示的话）
    `pic_urls:图片url,地址必须是图片空间的url,最多5张
type TaobaoOcTradetagAttachRequest struct {
    
    /* tag_name required标签名称 */
    tag_name string `json:"tag_name";xml:"tag_name"`
    
    /* tag_type optional标签类型       1：官方标签      2：自定义标签 */
    tag_type int64 `json:"tag_type";xml:"tag_type"`
    
    /* tag_value required标签值，json格式 */
    tag_value string `json:"tag_value";xml:"tag_value"`
    
    /* tid required订单id */
    tid int64 `json:"tid";xml:"tid"`
    
    /* visible optional该标签在消费者端是否显示,0:不显示,1：显示 */
    visible int64 `json:"visible";xml:"visible"`
    
}

func (req *TaobaoOcTradetagAttachRequest) GetAPIName() string {
	return "taobao.oc.tradetag.attach"
}

// TaobaoOcTradetagAttachResponse 对订单添加标签和更新标签。标签分为官方标签和自定义标签。
官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731
自定义标签有2个通用属性:
    `show_str:给消费者显示的字符串（如果可以显示的话）
    `pic_urls:图片url,地址必须是图片空间的url,最多5张
type TaobaoOcTradetagAttachResponse struct {
    
    /* result Basic操作成功或者操作失败 */
    result bool `json:"result";xml:"result"`
    
}
