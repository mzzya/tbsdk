package tbsdk

// Gfq1qp4287ShopexMatrixDatapushRequest 订单转发
type Gfq1qp4287ShopexMatrixDatapushRequest struct {
    
    /* p optional数据内容 */
    p string `json:"p";xml:"p"`
    
}

func (req *Gfq1qp4287ShopexMatrixDatapushRequest) GetAPIName() string {
	return "gfq1qp4287.shopex.matrix.datapush"
}

// Gfq1qp4287ShopexMatrixDatapushResponse 订单转发
type Gfq1qp4287ShopexMatrixDatapushResponse struct {
    
    /* data Basicdata */
    data string `json:"data";xml:"data"`
    
    /* res Basicres */
    res string `json:"res";xml:"res"`
    
}
