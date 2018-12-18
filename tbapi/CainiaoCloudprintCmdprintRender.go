package tbsdk

// CainiaoCloudprintCmdprintRenderRequest isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
type CainiaoCloudprintCmdprintRenderRequest struct {
    
    /* params required参数对象 */
    params CmdRenderParams `json:"params";xml:"params"`
    
}

func (req *CainiaoCloudprintCmdprintRenderRequest) GetAPIName() string {
	return "cainiao.cloudprint.cmdprint.render"
}

// CainiaoCloudprintCmdprintRenderResponse isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
type CainiaoCloudprintCmdprintRenderResponse struct {
    
    /* cmd_content Basic指令集内容串 */
    cmd_content string `json:"cmd_content";xml:"cmd_content"`
    
    /* cmd_encoding Basic指令集编码方式：origin-原串 gzip-采用gzip压缩并base64编码 */
    cmd_encoding string `json:"cmd_encoding";xml:"cmd_encoding"`
    
    /* ret_code Basic0成功，非0失败 */
    ret_code string `json:"ret_code";xml:"ret_code"`
    
    /* ret_msg Basic错误信息 */
    ret_msg string `json:"ret_msg";xml:"ret_msg"`
    
}
