package tbsdk

// TaobaoTradeMemoUpdateRequest 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
type TaobaoTradeMemoUpdateRequest struct {
    
    /* flag optional卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0 */
    flag int64 `json:"flag";xml:"flag"`
    
    /* memo optional卖家交易备注。最大长度: 1000个字节 */
    memo string `json:"memo";xml:"memo"`
    
    /* reset optional是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值 */
    reset bool `json:"reset";xml:"reset"`
    
    /* tid required交易编号 */
    tid int64 `json:"tid";xml:"tid"`
    
}

func (req *TaobaoTradeMemoUpdateRequest) GetAPIName() string {
	return "taobao.trade.memo.update"
}

// TaobaoTradeMemoUpdateResponse 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
type TaobaoTradeMemoUpdateResponse struct {
    
    /* trade Object更新交易的备注信息后返回的Trade，其中可用字段为tid和modified */
    trade Trade `json:"trade";xml:"trade"`
    
}
