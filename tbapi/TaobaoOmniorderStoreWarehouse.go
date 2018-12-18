package tbsdk

// TaobaoOmniorderStoreWarehouseRequest 为了能够支持逆向到门店对应的区域仓，商家需要配置门店与区域仓的映射关系，这个接口可以给商家提供门店与区域仓映射关系的增删改查功能。
type TaobaoOmniorderStoreWarehouseRequest struct {
    
    /* operation required操作，1表示增加或者更新，2表示删除，3表示查询 */
    operation int64 `json:"operation";xml:"operation"`
    
    /* store_id optional门店id */
    store_id int64 `json:"store_id";xml:"store_id"`
    
    /* warehouse_address optional仓联系地址 */
    warehouse_address string `json:"warehouse_address";xml:"warehouse_address"`
    
    /* warehouse_code optional仓编码 */
    warehouse_code string `json:"warehouse_code";xml:"warehouse_code"`
    
    /* warehouse_contact optional仓联系人 */
    warehouse_contact string `json:"warehouse_contact";xml:"warehouse_contact"`
    
    /* warehouse_mobile optional仓联系电话 */
    warehouse_mobile string `json:"warehouse_mobile";xml:"warehouse_mobile"`
    
}

func (req *TaobaoOmniorderStoreWarehouseRequest) GetAPIName() string {
	return "taobao.omniorder.store.warehouse"
}

// TaobaoOmniorderStoreWarehouseResponse 为了能够支持逆向到门店对应的区域仓，商家需要配置门店与区域仓的映射关系，这个接口可以给商家提供门店与区域仓映射关系的增删改查功能。
type TaobaoOmniorderStoreWarehouseResponse struct {
    
    /* data Basic成功增加或者更新一条门店与区域仓的关联 */
    data string `json:"data";xml:"data"`
    
    /* fail_code Basiccode */
    fail_code string `json:"fail_code";xml:"fail_code"`
    
    /* fail_message Basicmessage */
    fail_message string `json:"fail_message";xml:"fail_message"`
    
}
