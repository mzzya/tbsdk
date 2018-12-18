package tbsdk

// TaobaoQimenEventProduceRequest 当订单被处理时，用于通知奇门系统。
type TaobaoQimenEventProduceRequest struct {
    
    /* create optional订单创建时间,数字 */
    create int64 `json:"create";xml:"create"`
    
    /* ext optionalJSON格式扩展字段 */
    ext string `json:"ext";xml:"ext"`
    
    /* nick optional外部商家名称。必须同时填写platform */
    nick string `json:"nick";xml:"nick"`
    
    /* platform optional商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他 */
    platform string `json:"platform";xml:"platform"`
    
    /* status required事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK */
    status string `json:"status";xml:"status"`
    
    /* tid required淘宝订单号 */
    tid string `json:"tid";xml:"tid"`
    
}

func (req *TaobaoQimenEventProduceRequest) GetAPIName() string {
	return "taobao.qimen.event.produce"
}

// TaobaoQimenEventProduceResponse 当订单被处理时，用于通知奇门系统。
type TaobaoQimenEventProduceResponse struct {
    
    /* is_success Basic是否成功 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
