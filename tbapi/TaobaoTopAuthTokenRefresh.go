package tbsdk

// TaobaoTopAuthTokenRefreshRequest 根据refresh_token重新生成token
type TaobaoTopAuthTokenRefreshRequest struct {
    
    /* refresh_token requiredgrantType==refresh_token 时需要 */
    refresh_token string `json:"refresh_token";xml:"refresh_token"`
    
}

func (req *TaobaoTopAuthTokenRefreshRequest) GetAPIName() string {
	return "taobao.top.auth.token.refresh"
}

// TaobaoTopAuthTokenRefreshResponse 根据refresh_token重新生成token
type TaobaoTopAuthTokenRefreshResponse struct {
    
    /* token_result Basic返回的是json信息 */
    token_result map[string]interface{} `json:"token_result";xml:"token_result"`
    
}
