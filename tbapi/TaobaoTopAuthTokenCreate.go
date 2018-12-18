package tbsdk

// TaobaoTopAuthTokenCreateRequest 用户通过code换获取access_token，https only
type TaobaoTopAuthTokenCreateRequest struct {
    
    /* code required授权code，grantType==authorization_code 时需要 */
    code string `json:"code";xml:"code"`
    
    /* uuid optional与生成code的uuid配对 */
    uuid string `json:"uuid";xml:"uuid"`
    
}

func (req *TaobaoTopAuthTokenCreateRequest) GetAPIName() string {
	return "taobao.top.auth.token.create"
}

// TaobaoTopAuthTokenCreateResponse 用户通过code换获取access_token，https only
type TaobaoTopAuthTokenCreateResponse struct {
    
    /* token_result Basic返回的是json信息，和之前调用https://oauth.taobao.com/tac/token https://oauth.alibaba.com/token 换token返回的字段信息一致 */
    token_result map[string]interface{} `json:"token_result";xml:"token_result"`
    
}
