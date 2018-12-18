package tbsdk

// Gfq1qp4287ShopexMatrxiOfflineSendRequest 商派矩阵线下发货接口
type Gfq1qp4287ShopexMatrxiOfflineSendRequest struct {
    
    /* params optional参数 */
    params string `json:"params";xml:"params"`
    
}

func (req *Gfq1qp4287ShopexMatrxiOfflineSendRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.matrxi.offline.send"
}

// Gfq1qp4287ShopexMatrxiOfflineSendResponse 商派矩阵线下发货接口
type Gfq1qp4287ShopexMatrxiOfflineSendResponse struct {
    
    /* msg_id Basic任务id */
    msg_id string `json:"msg_id";xml:"msg_id"`
    
}
