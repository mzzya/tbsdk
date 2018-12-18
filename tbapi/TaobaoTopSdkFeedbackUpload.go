package tbsdk

// TaobaoTopSdkFeedbackUploadRequest sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaoTopSdkFeedbackUploadRequest struct {
    
    /* content optional具体内容，json形式 */
    content string `json:"content";xml:"content"`
    
    /* _type required1、回传加密信息 */
    _type string `json:"type";xml:"type"`
    
}

func (req *TaobaoTopSdkFeedbackUploadRequest) GetAPIName() string {
	return "taobao.top.sdk.feedback.upload"
}

// TaobaoTopSdkFeedbackUploadResponse sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaoTopSdkFeedbackUploadResponse struct {
    
    /* upload_interval Basic控制回传间隔（单位：秒） */
    upload_interval int64 `json:"upload_interval";xml:"upload_interval"`
    
}
