package tbsdk

// TaobaoOmniorderStoreGetconsignmailcodeRequest 用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
type TaobaoOmniorderStoreGetconsignmailcodeRequest struct {
    
    /* channel required淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS) */
    channel string `json:"channel";xml:"channel"`
    
    /* receiver optional收件人信息 */
    receiver ReceiverDto `json:"receiver";xml:"receiver"`
    
    /* sdt_extend_info_d_t_o optional扩展信息 */
    sdt_extend_info_d_t_o SdtExtendInfoDto `json:"sdt_extend_info_d_t_o";xml:"sdt_extend_info_d_t_o"`
    
    /* sender_contact optional发件人联系电话，如空则表示使用门店信息中的电话号码 */
    sender_contact string `json:"sender_contact";xml:"sender_contact"`
    
    /* store_id required门店ID */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* trades required订单信息，目前一次请求只支持一个主订单 */
    trades TradeOrderInfoDto `json:"trades";xml:"trades"`
    
}

func (req *TaobaoOmniorderStoreGetconsignmailcodeRequest) GetAPIName() string {
	return "taobao.omniorder.store.getconsignmailcode"
}

// TaobaoOmniorderStoreGetconsignmailcodeResponse 用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
type TaobaoOmniorderStoreGetconsignmailcodeResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
