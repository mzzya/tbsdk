package tbsdk

// TaobaoAppstoreSubscribeGetRequest 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
type TaobaoAppstoreSubscribeGetRequest struct {
    
    /* lease_id optional插件实例ID */
    lease_id int64 `json:"lease_id";xml:"lease_id"`
    
    /* nick required用户昵称 */
    nick string `json:"nick";xml:"nick"`
    
}

func (req *TaobaoAppstoreSubscribeGetRequest) GetAPIName() string {
	return "taobao.appstore.subscribe.get"
}

// TaobaoAppstoreSubscribeGetResponse 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
type TaobaoAppstoreSubscribeGetResponse struct {
    
    /* user_subscribe Object用户订购信息 */
    user_subscribe UserSubscribe `json:"user_subscribe";xml:"user_subscribe"`
    
}
