package tbsdk

// CainiaoWaybillprintClientupdateGetconfigRequest 获取客户端更新配置信息
type CainiaoWaybillprintClientupdateGetconfigRequest struct {
    
    /* lan_ip required服务发起机器局域网ip */
    lan_ip string `json:"lan_ip";xml:"lan_ip"`
    
    /* mac required服务发起机器的物理地址 */
    mac string `json:"mac";xml:"mac"`
    
    /* msgid required当前消息版本 */
    msgid int64 `json:"msgid";xml:"msgid"`
    
    /* version required当前打印客户端版本 */
    version string `json:"version";xml:"version"`
    
}

func (req *CainiaoWaybillprintClientupdateGetconfigRequest) GetAPIName() string {
	return "cainiao.waybillprint.clientupdate.getconfig"
}

// CainiaoWaybillprintClientupdateGetconfigResponse 获取客户端更新配置信息
type CainiaoWaybillprintClientupdateGetconfigResponse struct {
    
    /* result Objectresult */
    result UpdateConfigTopResult `json:"result";xml:"result"`
    
}
