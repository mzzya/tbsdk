package tbsdk

// AlipayPointOrderAddRequest 向用户发送集分宝，发放集分宝之前，操作流程如下：
1、申请支付宝增值包。
2、申请支付宝应用上线时选择集分宝API。
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
7、调用发放API（alipay.point.order.add）发放集分宝。
type AlipayPointOrderAddRequest struct {
    
    /* auth_token special支付宝用户给应用发放集分宝的授权。 */
    auth_token string `json:"auth_token";xml:"auth_token"`
    
    /* memo required向用户展示集分宝发放备注 */
    memo string `json:"memo";xml:"memo"`
    
    /* merchant_order_no requiredisv提供的发放号订单号，由数字和组成，最大长度为32为，需要保证每笔发放的唯一性，支付宝会对该参数做唯一性控制。如果使用同样的订单号，支付宝将返回订单号已经存在的错误 */
    merchant_order_no string `json:"merchant_order_no";xml:"merchant_order_no"`
    
    /* order_time required发放集分宝时间 */
    order_time Date `json:"order_time";xml:"order_time"`
    
    /* point_count required发放集分宝的数量 */
    point_count int64 `json:"point_count";xml:"point_count"`
    
    /* user_symbol required用户标识符，用于指定集分宝发放的用户，和user_symbol_type一起使用，确定一个唯一的支付宝用户 */
    user_symbol string `json:"user_symbol";xml:"user_symbol"`
    
    /* user_symbol_type required用户标识符类型，现在支持ALIPAY_USER_ID:表示支付宝用户ID,ALIPAY_LOGON_ID:表示支付宝登陆号 */
    user_symbol_type string `json:"user_symbol_type";xml:"user_symbol_type"`
    
}

func (req *AlipayPointOrderAddRequest) GetAPIName() string {
	return "alipay.point.order.add"
}

// AlipayPointOrderAddResponse 向用户发送集分宝，发放集分宝之前，操作流程如下：
1、申请支付宝增值包。
2、申请支付宝应用上线时选择集分宝API。
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。
7、调用发放API（alipay.point.order.add）发放集分宝。
type AlipayPointOrderAddResponse struct {
    
    /* alipay_order_no Basic支付宝集分宝发放流水号 */
    alipay_order_no string `json:"alipay_order_no";xml:"alipay_order_no"`
    
    /* result_code Basic充值结果：SUCCESS表示成功，其他表示失败 */
    result_code bool `json:"result_code";xml:"result_code"`
    
}
