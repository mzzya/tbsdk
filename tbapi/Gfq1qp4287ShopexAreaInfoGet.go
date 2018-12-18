package tbsdk

// Gfq1qp4287ShopexAreaInfoGetRequest 地区编码同步
type Gfq1qp4287ShopexAreaInfoGetRequest struct {
    
    /* qimen_body optional请求参数 */
    qimen_body string `json:"qimen_body";xml:"qimen_body"`
    
}

func (req *Gfq1qp4287ShopexAreaInfoGetRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.area.info.get"
}

// Gfq1qp4287ShopexAreaInfoGetResponse 地区编码同步
type Gfq1qp4287ShopexAreaInfoGetResponse struct {
    
    /* data Basic返回内容 */
    data string `json:"data";xml:"data"`
    
    /* msg_id Basic请求任务id */
    msg_id string `json:"msg_id";xml:"msg_id"`
    
    /* res Basic描述 */
    res string `json:"res";xml:"res"`
    
    /* rsp Basic成功失败标识 */
    rsp string `json:"rsp";xml:"rsp"`
    
}
