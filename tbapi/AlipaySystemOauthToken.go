package tbsdk

// AlipaySystemOauthTokenRequest OAuth2.0 授权的第二步，换取访问令牌。
type AlipaySystemOauthTokenRequest struct {
    
    /* code optional授权码，用户对应用授权后得到。 */
    code string `json:"code";xml:"code"`
    
    /* grant_type required获取访问令牌的类型，authorization_code表示用授权码换，refresh_token表示用刷新令牌来换。 */
    grant_type string `json:"grant_type";xml:"grant_type"`
    
    /* refresh_token optional刷新令牌，上次换取访问令牌是得到。 */
    refresh_token string `json:"refresh_token";xml:"refresh_token"`
    
}

func (req *AlipaySystemOauthTokenRequest) GetAPIName() string {
	return "alipay.system.oauth.token"
}

// AlipaySystemOauthTokenResponse OAuth2.0 授权的第二步，换取访问令牌。
type AlipaySystemOauthTokenResponse struct {
    
    /* access_token Basic访问令牌 */
    access_token string `json:"access_token";xml:"access_token"`
    
    /* alipay_user_id Basic支付宝用户的id。 */
    alipay_user_id string `json:"alipay_user_id";xml:"alipay_user_id"`
    
    /* expires_in Basic访问令牌的有效时间，单位是秒。 */
    expires_in string `json:"expires_in";xml:"expires_in"`
    
    /* re_expires_in Basic刷新令牌的有效时间，单位是秒。 */
    re_expires_in string `json:"re_expires_in";xml:"re_expires_in"`
    
    /* refresh_token Basic刷新令牌 */
    refresh_token string `json:"refresh_token";xml:"refresh_token"`
    
}
