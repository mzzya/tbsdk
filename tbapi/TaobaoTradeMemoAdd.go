package tbsdk

// TaobaoTradeMemoAddRequest 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
type TaobaoTradeMemoAddRequest struct {
    
    /* flag optional交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0 */
    flag int64 `json:"flag";xml:"flag"`
    
    /* memo required交易备注。最大长度: 1000个字节 */
    memo string `json:"memo";xml:"memo"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeMemoAddRequest) GetAPIName() string {
	return "taobao.trade.memo.add"
}

// TaobaoTradeMemoAddResponse 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
type TaobaoTradeMemoAddResponse struct {
    
    /* trade Object对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created */
    trade Trade `json:"trade";xml:"trade"`
    
}
