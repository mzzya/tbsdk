package tbsdk

// CainiaoCloudprintCustomareaUpdateRequest 自定义区内容更新
type CainiaoCloudprintCustomareaUpdateRequest struct {
    
    /* custom_area_content required自定义区内容（可修改） */
    custom_area_content string `json:"custom_area_content";xml:"custom_area_content"`
    
    /* custom_area_id required自定义区id（不可修改） */
    custom_area_id int64 `json:"custom_area_id";xml:"custom_area_id"`
    
    /* custom_area_name required自定义区名称（可修改） */
    custom_area_name string `json:"custom_area_name";xml:"custom_area_name"`
    
}

func (req *CainiaoCloudprintCustomareaUpdateRequest) GetAPIName() string {
	return "cainiao.cloudprint.customarea.update"
}

// CainiaoCloudprintCustomareaUpdateResponse 自定义区内容更新
type CainiaoCloudprintCustomareaUpdateResponse struct {
    
    /* result Objectresult */
    result CloudPrintBaseResult `json:"result";xml:"result"`
    
}
