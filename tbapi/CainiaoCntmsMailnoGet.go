package tbsdk

// CainiaoCntmsMailnoGetRequest 打印面单时，通过此接口获取面单号及打打印信息
type CainiaoCntmsMailnoGetRequest struct {
    
    /* content optional获取菜鸟配送电子面单请求参数 */
    content CnTmsMailnoGetContent `json:"content";xml:"content"`
    
}

func (req *CainiaoCntmsMailnoGetRequest) GetAPIName() string {
	return "cainiao.cntms.mailno.get"
}

// CainiaoCntmsMailnoGetResponse 打印面单时，通过此接口获取面单号及打打印信息
type CainiaoCntmsMailnoGetResponse struct {
    
    /* allocator_code Basic揽货商（分拨中心）编码 */
    allocator_code string `json:"allocator_code";xml:"allocator_code"`
    
    /* allocator_name Basic揽货商（分拨中心）名称 */
    allocator_name string `json:"allocator_name";xml:"allocator_name"`
    
    /* logistics_code Basic物流商公司编码
（ERP在调用菜鸟发货接口时此字段赋值到tms_code, 调用淘宝“自己联系物流（线下物流）发货”时，做为company_code传入） */
    logistics_code string `json:"logistics_code";xml:"logistics_code"`
    
    /* logistics_name Basic物流商名称 */
    logistics_name string `json:"logistics_name";xml:"logistics_name"`
    
    /* mailno Basic面单号 */
    mailno string `json:"mailno";xml:"mailno"`
    
    /* package_center_code Basic集包地代码 */
    package_center_code string `json:"package_center_code";xml:"package_center_code"`
    
    /* package_center_name Basic集包地名称 */
    package_center_name string `json:"package_center_name";xml:"package_center_name"`
    
    /* sec_distribution Basic二级流向信息 */
    sec_distribution string `json:"sec_distribution";xml:"sec_distribution"`
    
    /* short_address Basic一级流向信息，大头笔 */
    short_address string `json:"short_address";xml:"short_address"`
    
    /* tms_code Basic二级配送公司编码 */
    tms_code string `json:"tms_code";xml:"tms_code"`
    
    /* tms_name Basic二级配送公司名称 */
    tms_name string `json:"tms_name";xml:"tms_name"`
    
}
