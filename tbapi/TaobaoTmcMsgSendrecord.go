package tbsdk

// TaobaoTmcMsgSendrecordRequest 查询单条消息发送记录，只返回返回条数和时间。
type TaobaoTmcMsgSendrecordRequest struct {
    
    /* data_id required消息主键ID */
    data_id string `json:"data_id";xml:"data_id"`
    
    /* group_name required消息分组名 */
    group_name string `json:"group_name";xml:"group_name"`
    
    /* topic_name requiredTOPIC名称 */
    topic_name string `json:"topic_name";xml:"topic_name"`
    
}

func (req *TaobaoTmcMsgSendrecordRequest) GetAPIName() string {
	return "taobao.tmc.msg.sendrecord"
}

// TaobaoTmcMsgSendrecordResponse 查询单条消息发送记录，只返回返回条数和时间。
type TaobaoTmcMsgSendrecordResponse struct {
    
    /* tb_send_list Basic淘宝发送时间 */
    tb_send_list string `json:"tb_send_list";xml:"tb_send_list"`
    
    /* tb_send_times Basic淘宝发送测试 */
    tb_send_times int64 `json:"tb_send_times";xml:"tb_send_times"`
    
    /* tmc_send_list BasicTMC发送时间 */
    tmc_send_list string `json:"tmc_send_list";xml:"tmc_send_list"`
    
    /* tmc_send_times Basictmc发送次数 */
    tmc_send_times int64 `json:"tmc_send_times";xml:"tmc_send_times"`
    
}
