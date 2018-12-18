package tbsdk

// CainiaoEndpointLockerTopStationAddorupdateRequest 新增或者修改代收点相关信息
type CainiaoEndpointLockerTopStationAddorupdateRequest struct {
    
    /* station_info optional站点信息 */
    station_info StationInfo `json:"station_info";xml:"station_info"`
    
}

func (req *CainiaoEndpointLockerTopStationAddorupdateRequest) GetAPIName() string {
	return "cainiao.endpoint.locker.top.station.addorupdate"
}

// CainiaoEndpointLockerTopStationAddorupdateResponse 新增或者修改代收点相关信息
type CainiaoEndpointLockerTopStationAddorupdateResponse struct {
    
    /* result Objectresult */
    result SingleResult `json:"result";xml:"result"`
    
}
