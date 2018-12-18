package tbsdk

// TaobaoWlbOrderJzQueryRequest 家装业务查询物流公司api
type TaobaoWlbOrderJzQueryRequest struct {
    
    /* ins_jz_receiver_t_o optional家装安装服务收货人信息 */
    ins_jz_receiver_t_o JzReceiverTo `json:"ins_jz_receiver_t_o";xml:"ins_jz_receiver_t_o"`
    
    /* jz_receiver_to optional家装收货人信息 */
    jz_receiver_to JzReceiverTo `json:"jz_receiver_to";xml:"jz_receiver_to"`
    
    /* sender_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址 */
    sender_id int64 `json:"sender_id";xml:"sender_id"`
    
    /* tid optional交易id */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoWlbOrderJzQueryRequest) GetAPIName() string {
	return "taobao.wlb.order.jz.query"
}

// TaobaoWlbOrderJzQueryResponse 家装业务查询物流公司api
type TaobaoWlbOrderJzQueryResponse struct {
    
    /* result Object结果信息 */
    result JzTopDto `json:"result";xml:"result"`
    
    /* result_error_code Basic错误编码 */
    result_error_code string `json:"result_error_code";xml:"result_error_code"`
    
    /* result_error_msg Basic错误信息 */
    result_error_msg string `json:"result_error_msg";xml:"result_error_msg"`
    
    /* result_success Basic是否成功 */
    result_success bool `json:"result_success";xml:"result_success"`
    
}
