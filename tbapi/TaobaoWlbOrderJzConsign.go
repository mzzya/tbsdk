package tbsdk

// TaobaoWlbOrderJzConsignRequest 家装类订单使用该接口发货
type TaobaoWlbOrderJzConsignRequest struct {
    
    /* ins_receiver_to optional安装收货人信息,如果为空,则取默认收货人信息 */
    ins_receiver_to JzReceiverTo `json:"ins_receiver_to";xml:"ins_receiver_to"`
    
    /* ins_tp_dto optional安装公司信息,需要安装时,才填写 */
    ins_tp_dto Tpdto `json:"ins_tp_dto";xml:"ins_tp_dto"`
    
    /* jz_receiver_to optional家装收货人信息,如果为空,则取默认收货信息 */
    jz_receiver_to JzReceiverTo `json:"jz_receiver_to";xml:"jz_receiver_to"`
    
    /* jz_top_args optional发货参数 */
    jz_top_args JzTopArgs `json:"jz_top_args";xml:"jz_top_args"`
    
    /* lg_tp_dto required物流公司信息 */
    lg_tp_dto Tpdto `json:"lg_tp_dto";xml:"lg_tp_dto"`
    
    /* sender_id optional卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址 */
    sender_id int64 `json:"sender_id";xml:"sender_id"`
    
    /* tid required交易号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoWlbOrderJzConsignRequest) GetAPIName() string {
	return "taobao.wlb.order.jz.consign"
}

// TaobaoWlbOrderJzConsignResponse 家装类订单使用该接口发货
type TaobaoWlbOrderJzConsignResponse struct {
    
    /* result_error_code Basic错误码 */
    result_error_code string `json:"result_error_code";xml:"result_error_code"`
    
    /* result_error_msg Basic错误信息描述 */
    result_error_msg string `json:"result_error_msg";xml:"result_error_msg"`
    
    /* result_success Basic是否成功 */
    result_success bool `json:"result_success";xml:"result_success"`
    
}
