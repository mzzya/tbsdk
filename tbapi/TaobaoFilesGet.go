package tbsdk

// TaobaoFilesGetRequest 获取业务方暂存给ISV的文件列表
type TaobaoFilesGetRequest struct {
    
    /* end_date required搜索结束时间 */
    end_date Date `json:"end_date";xml:"end_date"`
    
    /* start_date required搜索开始时间 */
    start_date Date `json:"start_date";xml:"start_date"`
    
    /* status optional下载链接状态。1:未下载。2:已下载 */
    status int64 `json:"status";xml:"status"`
    
}

func (req *TaobaoFilesGetRequest) GetAPIName() string {
	return "taobao.files.get"
}

// TaobaoFilesGetResponse 获取业务方暂存给ISV的文件列表
type TaobaoFilesGetResponse struct {
    
    /* results Object Arrayresults */
    results TopDownloadRecordDo `json:"results";xml:"results"`
    
}
