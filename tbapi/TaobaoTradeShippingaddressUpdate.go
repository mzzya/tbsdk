package tbsdk

// TaobaoTradeShippingaddressUpdateRequest 只能更新一笔交易里面的买家收货地址 
只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 
更新后的发货地址可以通过taobao.trade.fullinfo.get查到 
参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaoTradeShippingaddressUpdateRequest struct {
    
    /* receiver_address optional收货地址。最大长度为228个字节。 */
    receiver_address string `json:"receiver_address";xml:"receiver_address"`
    
    /* receiver_city optional城市。最大长度为32个字节。如：杭州 */
    receiver_city string `json:"receiver_city";xml:"receiver_city"`
    
    /* receiver_district optional区/县。最大长度为32个字节。如：西湖区 */
    receiver_district string `json:"receiver_district";xml:"receiver_district"`
    
    /* receiver_mobile optional移动电话。最大长度为11个字节。 */
    receiver_mobile string `json:"receiver_mobile";xml:"receiver_mobile"`
    
    /* receiver_name optional收货人全名。最大长度为50个字节。 */
    receiver_name string `json:"receiver_name";xml:"receiver_name"`
    
    /* receiver_phone optional固定电话。最大长度为30个字节。 */
    receiver_phone string `json:"receiver_phone";xml:"receiver_phone"`
    
    /* receiver_state optional省份。最大长度为32个字节。如：浙江 */
    receiver_state string `json:"receiver_state";xml:"receiver_state"`
    
    /* receiver_zip optional邮政编码。必须由6个数字组成。 */
    receiver_zip string `json:"receiver_zip";xml:"receiver_zip"`
    
    /* tid required交易编号。 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeShippingaddressUpdateRequest) GetAPIName() string {
	return "taobao.trade.shippingaddress.update"
}

// TaobaoTradeShippingaddressUpdateResponse 只能更新一笔交易里面的买家收货地址 
只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 
更新后的发货地址可以通过taobao.trade.fullinfo.get查到 
参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaoTradeShippingaddressUpdateResponse struct {
    
    /* trade Object交易结构 */
    trade Trade `json:"trade";xml:"trade"`
    
}
