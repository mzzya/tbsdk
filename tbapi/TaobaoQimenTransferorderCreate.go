package tbsdk

// TaobaoQimenTransferorderCreateRequest 调拨单创建
type TaobaoQimenTransferorderCreateRequest struct {
    
    /* attributes optional扩展属性,HZ1234,string(500),, */
    attributes string `json:"attributes";xml:"attributes"`
    
    /* erpOrderCode optional外部ERP订单号,HZ1234,string(50),必填, */
    erpOrderCode string `json:"erpOrderCode";xml:"erpOrderCode"`
    
    /* expectStartTime optional期望调拨开始时间,Item1234,string(50),, */
    expectStartTime string `json:"expectStartTime";xml:"expectStartTime"`
    
    /* fromStoreCode optional出库仓编码,Item1234,string(50),必填, */
    fromStoreCode string `json:"fromStoreCode";xml:"fromStoreCode"`
    
    /* ownerCode required111 */
    ownerCode string `json:"ownerCode";xml:"ownerCode"`
    
    /* toStoreCode optional入库仓编码,HZ1234,string(50),必填, */
    toStoreCode string `json:"toStoreCode";xml:"toStoreCode"`
    
    /* transferItems optional项目集 */
    transferItems TransferItems `json:"transferItems";xml:"transferItems"`
    
}

func (req *TaobaoQimenTransferorderCreateRequest) GetAPIName() string {
	return "taobao.qimen.transferorder.create"
}

// TaobaoQimenTransferorderCreateResponse 调拨单创建
type TaobaoQimenTransferorderCreateResponse struct {
    
    /* code Basic响应码,0,string(50),, */
    code string `json:"code";xml:"code"`
    
    /* flag Basic响应结果:success|failure,success,string(10),必填, */
    flag string `json:"flag";xml:"flag"`
    
    /* message Basic响应信息,invalid appkey,string(100),, */
    message string `json:"message";xml:"message"`
    
    /* transferExecuteInfo Object调拨单信息 */
    transferExecuteInfo TransferExecuteInfo `json:"transferExecuteInfo";xml:"transferExecuteInfo"`
    
}
