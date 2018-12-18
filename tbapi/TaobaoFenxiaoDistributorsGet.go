package tbsdk

// TaobaoFenxiaoDistributorsGetRequest 查询和当前登录供应商有合作关系的分销商的信息
type TaobaoFenxiaoDistributorsGetRequest struct {
    
    /* nicks required分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。 */
    nicks string `json:"nicks";xml:"nicks"`
    
}

func (req *TaobaoFenxiaoDistributorsGetRequest) GetAPIName() string {
	return "taobao.fenxiao.distributors.get"
}

// TaobaoFenxiaoDistributorsGetResponse 查询和当前登录供应商有合作关系的分销商的信息
type TaobaoFenxiaoDistributorsGetResponse struct {
    
    /* distributors Object Array分销商详细信息 */
    distributors Distributor `json:"distributors";xml:"distributors"`
    
}
