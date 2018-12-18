package tbsdk

// TaobaoTopSecretGetRequest top sdk通过api获取对应解密秘钥
type TaobaoTopSecretGetRequest struct {
    
    /* customer_user_id optional自定义用户id */
    customer_user_id int64 `json:"customer_user_id";xml:"customer_user_id"`
    
    /* random_num required伪随机数 */
    random_num string `json:"random_num";xml:"random_num"`
    
    /* secret_version optional秘钥版本号 */
    secret_version int64 `json:"secret_version";xml:"secret_version"`
    
}

func (req *TaobaoTopSecretGetRequest) GetAPIName() string {
	return "taobao.top.secret.get"
}

// TaobaoTopSecretGetResponse top sdk通过api获取对应解密秘钥
type TaobaoTopSecretGetResponse struct {
    
    /* app_config Basicapp配置信息 */
    app_config string `json:"app_config";xml:"app_config"`
    
    /* interval Basic下次更新秘钥间隔，单位（秒） */
    interval int64 `json:"interval";xml:"interval"`
    
    /* max_interval Basic最长有效期，容灾使用，单位（秒） */
    max_interval int64 `json:"max_interval";xml:"max_interval"`
    
    /* secret Basic秘钥值 */
    secret string `json:"secret";xml:"secret"`
    
    /* secret_version Basic秘钥版本号 */
    secret_version int64 `json:"secret_version";xml:"secret_version"`
    
}
