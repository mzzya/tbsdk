package tbsdk

// TaobaoTopSecretRegisterRequest 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TaobaoTopSecretRegisterRequest struct {
    
    /* user_id optional用户id，保证唯一 */
    user_id int64 `json:"user_id";xml:"user_id"`
    
}

func (req *TaobaoTopSecretRegisterRequest) GetAPIName() string {
	return "taobao.top.secret.register"
}

// TaobaoTopSecretRegisterResponse 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TaobaoTopSecretRegisterResponse struct {
    
    /* result Basic返回操作是否成功 */
    result bool `json:"result";xml:"result"`
    
}
