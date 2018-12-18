package tbsdk

// TmallItemShiptimeUpdateRequest 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
1.
    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 0 ---更新SKU
    },

    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

2.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 0 --更新SKU
    },
    按照指定SKU删除指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

3.

    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 1 ---更新商品
    },

    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间

4.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 1 --更新商品
    },
    删除商品级的发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
type TmallItemShiptimeUpdateRequest struct {
    
    /* item_id required商品ID */
    item_id int64 `json:"item_id";xml:"item_id"`
    
    /* option required批量更新商品/SKU发货时间的备选项 */
    option UpdateItemShipTimeOption `json:"option";xml:"option"`
    
    /* ship_time optional被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。 */
    ship_time string `json:"ship_time";xml:"ship_time"`
    
    /* sku_ship_times optional被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。 */
    sku_ship_times UpdateSkuShipTime `json:"sku_ship_times";xml:"sku_ship_times"`
    
}

func (req *TmallItemShiptimeUpdateRequest) GetAPIName() string {
	return "tmall.item.shiptime.update"
}

// TmallItemShiptimeUpdateResponse 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
1.
    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 0 ---更新SKU
    },

    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

2.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 0 --更新SKU
    },
    按照指定SKU删除指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

3.

    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 1 ---更新商品
    },

    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间

4.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 1 --更新商品
    },
    删除商品级的发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
type TmallItemShiptimeUpdateResponse struct {
    
    /* shiptime_update_result Basic被更新商品ID */
    shiptime_update_result string `json:"shiptime_update_result";xml:"shiptime_update_result"`
    
}
