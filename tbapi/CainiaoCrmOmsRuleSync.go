package tbsdk

// CainiaoCrmOmsRuleSyncRequest 将商家ERP订单处理规则同步到菜鸟CRM系统
type CainiaoCrmOmsRuleSyncRequest struct {
    
    /* check_rule_msg optional人工审单规则描述 */
    check_rule_msg string `json:"check_rule_msg";xml:"check_rule_msg"`
    
    /* is_auto_check required是否系统智能处理订单（无人工介入） */
    is_auto_check bool `json:"is_auto_check";xml:"is_auto_check"`
    
    /* is_open_cnauto required是否开启菜鸟自动流转规则 */
    is_open_cnauto bool `json:"is_open_cnauto";xml:"is_open_cnauto"`
    
    /* is_sys_merge_order optional是否开启了订单合单 */
    is_sys_merge_order bool `json:"is_sys_merge_order";xml:"is_sys_merge_order"`
    
    /* merge_order_cycle optional订单合单周期（单位：分钟） */
    merge_order_cycle int64 `json:"merge_order_cycle";xml:"merge_order_cycle"`
    
    /* other_rule optional其他未定义订单处理规则，格式｛name;stauts;cycle;｝ */
    other_rule string `json:"other_rule";xml:"other_rule"`
    
    /* shop_code required店铺nick */
    shop_code string `json:"shop_code";xml:"shop_code"`
    
}

func (req *CainiaoCrmOmsRuleSyncRequest) GetAPIName() string {
	return "cainiao.crm.oms.rule.sync"
}

// CainiaoCrmOmsRuleSyncResponse 将商家ERP订单处理规则同步到菜鸟CRM系统
type CainiaoCrmOmsRuleSyncResponse struct {
    
    /* wl_error_code BasicerrorCode */
    wl_error_code string `json:"wl_error_code";xml:"wl_error_code"`
    
    /* wl_error_msg BasicerrorMsg */
    wl_error_msg string `json:"wl_error_msg";xml:"wl_error_msg"`
    
    /* wl_success Basicsuccess */
    wl_success bool `json:"wl_success";xml:"wl_success"`
    
}
