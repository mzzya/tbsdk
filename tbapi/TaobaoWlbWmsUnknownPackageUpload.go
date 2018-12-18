package tbsdk

// TaobaoWlbWmsUnknownPackageUploadRequest 有货无单销退入库单消息回传
type TaobaoWlbWmsUnknownPackageUploadRequest struct {
    
    /* content requiredWlbWmsUnknownPackageUpload */
    content WlbWmsUnknownPackageUpload `json:"content";xml:"content"`
    
}

func (req *TaobaoWlbWmsUnknownPackageUploadRequest) GetAPIName() string {
	return "taobao.wlb.wms.unknown.package.upload"
}

// TaobaoWlbWmsUnknownPackageUploadResponse 有货无单销退入库单消息回传
type TaobaoWlbWmsUnknownPackageUploadResponse struct {
    
    /* response ObjectWlbWmsUnknownPackageUploadResp */
    response WlbWmsUnknownPackageUploadResp `json:"response";xml:"response"`
    
}
