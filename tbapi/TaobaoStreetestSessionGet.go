package tbsdk

// TaobaoStreetestSessionGetRequest 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaoStreetestSessionGetRequest struct {
    
}

func (req *TaobaoStreetestSessionGetRequest) GetAPIName() string {
	return "taobao.streetest.session.get"
}

// TaobaoStreetestSessionGetResponse 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaoStreetestSessionGetResponse struct {
    
    /* stree_test_session_key Basic压测账号对应的sessionKey */
    stree_test_session_key string `json:"stree_test_session_key";xml:"stree_test_session_key"`
    
}
