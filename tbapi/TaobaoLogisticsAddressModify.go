package tbsdk

// TaobaoLogisticsAddressModifyRequest 卖家地址库修改
type TaobaoLogisticsAddressModifyRequest struct {
    
    /* addr required详细街道地址，不需要重复填写省/市/区 */
    addr string `json:"addr";xml:"addr"`
    
    /* cancel_def optional默认退货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
    cancel_def bool `json:"cancel_def";xml:"cancel_def"`
    
    /* city required所在市 */
    city string `json:"city";xml:"city"`
    
    /* contact_id required地址库ID */
    contact_id int64 `json:"contact_id";xml:"contact_id"`
    
    /* contact_name required联系人姓名
<font color='red'>长度不可超过20个字节</font> */
    contact_name string `json:"contact_name";xml:"contact_name"`
    
    /* country optional区、县
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
    country string `json:"country";xml:"country"`
    
    /* get_def optional默认取货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
    get_def bool `json:"get_def";xml:"get_def"`
    
    /* memo optional备注,<br><font color='red'>备注不能超过256字节</font> */
    memo string `json:"memo";xml:"memo"`
    
    /* mobile_phone special手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font> */
    mobile_phone string `json:"mobile_phone";xml:"mobile_phone"`
    
    /* phone special电话号码,手机与电话必需有一个 */
    phone string `json:"phone";xml:"phone"`
    
    /* province required所在省 */
    province string `json:"province";xml:"province"`
    
    /* seller_company optional公司名称,
<br><font color='red'>公司名称长能不能超过20字节</font> */
    seller_company string `json:"seller_company";xml:"seller_company"`
    
    /* send_def optional默认发货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认发货地址，撤消原默认发货地址</font> */
    send_def bool `json:"send_def";xml:"send_def"`
    
    /* zip_code optional地区邮政编码
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
    zip_code string `json:"zip_code";xml:"zip_code"`
    
}

func (req *TaobaoLogisticsAddressModifyRequest) GetAPIName() string {
	return "taobao.logistics.address.modify"
}

// TaobaoLogisticsAddressModifyResponse 卖家地址库修改
type TaobaoLogisticsAddressModifyResponse struct {
    
    /* address_result Object只返回修改时间modify_date */
    address_result AddressResult `json:"address_result";xml:"address_result"`
    
}
