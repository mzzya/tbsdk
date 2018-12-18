package tbsdk

// TaobaoTmcMessageProduceRequest 发布单条消息
type TaobaoTmcMessageProduceRequest struct {
    
    /* content required消息内容的JSON表述，必须按照topic的定义来填充 */
    content string `json:"content";xml:"content"`
    
    /* ex_content optional消息的扩增属性，用json格式表示 */
    ex_content string `json:"ex_content";xml:"ex_content"`
    
    /* media_content optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。 */
    media_content byte[] `json:"media_content";xml:"media_content"`
    
    /* media_content2 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content2 byte[] `json:"media_content2";xml:"media_content2"`
    
    /* media_content3 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content3 byte[] `json:"media_content3";xml:"media_content3"`
    
    /* media_content4 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content4 byte[] `json:"media_content4";xml:"media_content4"`
    
    /* media_content5 optional回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。 */
    media_content5 byte[] `json:"media_content5";xml:"media_content5"`
    
    /* target_appkey optional直发消息需要传入目标appkey */
    target_appkey string `json:"target_appkey";xml:"target_appkey"`
    
    /* target_group optional目标分组，一般为default */
    target_group string `json:"target_group";xml:"target_group"`
    
    /* topic required消息类型 */
    topic string `json:"topic";xml:"topic"`
    
}

func (req *TaobaoTmcMessageProduceRequest) GetAPIName() string {
	return "taobao.tmc.message.produce"
}

// TaobaoTmcMessageProduceResponse 发布单条消息
type TaobaoTmcMessageProduceResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
    /* msg_ids Basic Array消息ID */
    msg_ids string `json:"msg_ids";xml:"msg_ids"`
    
    /* total Basic投递目标数 */
    total int64 `json:"total";xml:"total"`
    
}
