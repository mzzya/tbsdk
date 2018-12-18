package tbsdk

// TaobaoQianniuCloudkefuStatuslogGetRequest 查询客服账号的操作记录。如：登录，下线，挂起等。接口即将下线，请ISV切换到taobao.qianniu.cloudkefu.onlinestatuslog.get接口上
type TaobaoQianniuCloudkefuStatuslogGetRequest struct {
    
    /* account_ids required子帐号列表，最多10个 */
    account_ids int64 `json:"account_ids";xml:"account_ids"`
    
    /* domian optional域，淘宝:cntaobao，1688:cnalichn */
    domian string `json:"domian";xml:"domian"`
    
    /* end_time optional查询结束时间，默认当天24点 */
    end_time Date `json:"end_time";xml:"end_time"`
    
    /* main_account_id required主帐号ID */
    main_account_id int64 `json:"main_account_id";xml:"main_account_id"`
    
    /* page_num optional页码，每页50个 */
    page_num int64 `json:"page_num";xml:"page_num"`
    
    /* page_size optional分页 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* start_time optional查询开始时间，默认当天0点 */
    start_time Date `json:"start_time";xml:"start_time"`
    
    /* _type required记录类型，PC上下线：8；移动上下线：4；挂起类型：5 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoQianniuCloudkefuStatuslogGetRequest) GetAPIName() string {
	return "taobao.qianniu.cloudkefu.statuslog.get"
}

// TaobaoQianniuCloudkefuStatuslogGetResponse 查询客服账号的操作记录。如：登录，下线，挂起等。接口即将下线，请ISV切换到taobao.qianniu.cloudkefu.onlinestatuslog.get接口上
type TaobaoQianniuCloudkefuStatuslogGetResponse struct {
    
    /* result Objectresult */
    result ResultTo `json:"result";xml:"result"`
    
}
