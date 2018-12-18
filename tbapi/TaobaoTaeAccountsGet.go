package tbsdk

// TaobaoTaeAccountsGetRequest tae查询费用科目信息
type TaobaoTaeAccountsGetRequest struct {
    
    /* aids optional需要获取的科目ID */
    aids int64 `json:"aids";xml:"aids"`
    
    /* fields required需要返回的字段 */
    fields string `json:"fields";xml:"fields"`
    
}

func (req *TaobaoTaeAccountsGetRequest) GetAPIName() string {
	return "taobao.tae.accounts.get"
}

// TaobaoTaeAccountsGetResponse tae查询费用科目信息
type TaobaoTaeAccountsGetResponse struct {
    
    /* accounts Object Array返回的科目信息 */
    accounts TopAccountDto `json:"accounts";xml:"accounts"`
    
    /* total_results Basic返回记录行数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
