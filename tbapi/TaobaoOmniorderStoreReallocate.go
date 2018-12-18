package tbsdk

// TaobaoOmniorderStoreReallocateRequest 门店发货提供改派接口
type TaobaoOmniorderStoreReallocateRequest struct {
    
    /* main_order_id required主订单号 */
    main_order_id int64 `json:"main_order_id";xml:"main_order_id"`
    
    /* store_id optional门店Id */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* sub_order_ids required子订单号 */
    sub_order_ids int64 `json:"sub_order_ids";xml:"sub_order_ids"`
    
    /* warehouse_code optional电商仓code */
    warehouse_code string `json:"warehouse_code";xml:"warehouse_code"`
    
}

func (req *TaobaoOmniorderStoreReallocateRequest) GetAPIName() string {
	return "taobao.omniorder.store.reallocate"
}

// TaobaoOmniorderStoreReallocateResponse 门店发货提供改派接口
type TaobaoOmniorderStoreReallocateResponse struct {
    
    /* result Objectresult */
    result Result `json:"result";xml:"result"`
    
}
