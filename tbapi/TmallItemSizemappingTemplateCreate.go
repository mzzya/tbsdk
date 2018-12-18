package tbsdk

// TmallItemSizemappingTemplateCreateRequest 新增天猫商品尺码表模板

男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：
脚长（cm） 必选

内衣-文胸类目，尺码表维度为：
上胸围（cm） 必选
下胸围（cm） 必选
上下胸围差（cm） 必选
身高（cm）
体重（公斤）

内衣-内裤类目，尺码表维度为：
腰围（cm） 必选
臀围（cm） 必选
身高（cm）
体重（公斤）
裤长（cm）
裆部（cm）
脚口（cm）
腿围（cm）

内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
腰围（cm）
肩宽（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
领口（cm）

内衣-睡裤/保暖裤类目，尺码维度为：
身高（cm） 必选
腰围（cm） 必选
体重（公斤）
臀围（cm）
裆部（cm）
裤长（cm）
脚口（cm）
腿围（cm）
裤侧长（cm）

内衣-睡裙类目，尺码维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
裙长（cm）
腰围（cm）
袖长（cm）
肩宽（cm）
背宽（cm）
腿围（cm）
臀围（cm）
底摆（cm）

男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
肩宽（cm）
胸围（cm）
腰围（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
摆围（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
中腰（cm）
领深（cm）
领高（cm）
领宽（cm）
领围（cm）
圆摆后中长（cm）
平摆衣长（cm）
圆摆衣长（cm）

男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
腰围（cm）
臀围（cm）
裤长（cm）
裙长（cm）
裙摆长（cm）
腿围（cm）
膝围（cm）
小脚围（cm）
拉伸腰围（cm）
坐围（cm）
拉伸坐围（cm）
脚口（cm）
前浪（cm）
后浪（cm）
横档（cm）

如果上述维度满足，可以自定义最多5个维度。

模板格式为：
尺码值:维度名称:数值
如：M:身高（cm）:160,L:身高（cm）:170
type TmallItemSizemappingTemplateCreateRequest struct {
    
    /* template_content required尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。 */
    template_content string `json:"template_content";xml:"template_content"`
    
    /* template_name required尺码表模板名称 */
    template_name string `json:"template_name";xml:"template_name"`
    
}

func (req *TmallItemSizemappingTemplateCreateRequest) GetAPIName() string {
	return "tmall.item.sizemapping.template.create"
}

// TmallItemSizemappingTemplateCreateResponse 新增天猫商品尺码表模板

男鞋、女鞋、运动鞋、户外鞋类目，尺码表维度为：
脚长（cm） 必选

内衣-文胸类目，尺码表维度为：
上胸围（cm） 必选
下胸围（cm） 必选
上下胸围差（cm） 必选
身高（cm）
体重（公斤）

内衣-内裤类目，尺码表维度为：
腰围（cm） 必选
臀围（cm） 必选
身高（cm）
体重（公斤）
裤长（cm）
裆部（cm）
脚口（cm）
腿围（cm）

内衣-睡衣上衣/保暖上衣/睡袍类目，尺码表维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
腰围（cm）
肩宽（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
领口（cm）

内衣-睡裤/保暖裤类目，尺码维度为：
身高（cm） 必选
腰围（cm） 必选
体重（公斤）
臀围（cm）
裆部（cm）
裤长（cm）
脚口（cm）
腿围（cm）
裤侧长（cm）

内衣-睡裙类目，尺码维度为：
身高（cm） 必选
胸围（cm） 必选
体重（公斤）
裙长（cm）
腰围（cm）
袖长（cm）
肩宽（cm）
背宽（cm）
腿围（cm）
臀围（cm）
底摆（cm）

男装、女装、运动服、户外服等上装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
肩宽（cm）
胸围（cm）
腰围（cm）
袖长（cm）
衣长（cm）
背宽（cm）
前长（cm）
摆围（cm）
下摆围（cm）
袖口（cm）
袖肥（cm）
中腰（cm）
领深（cm）
领高（cm）
领宽（cm）
领围（cm）
圆摆后中长（cm）
平摆衣长（cm）
圆摆衣长（cm）

男装、女装、运动服、户外服等下装类目，尺码维度为（至少两项必选）：
身高（cm）
体重（公斤）
腰围（cm）
臀围（cm）
裤长（cm）
裙长（cm）
裙摆长（cm）
腿围（cm）
膝围（cm）
小脚围（cm）
拉伸腰围（cm）
坐围（cm）
拉伸坐围（cm）
脚口（cm）
前浪（cm）
后浪（cm）
横档（cm）

如果上述维度满足，可以自定义最多5个维度。

模板格式为：
尺码值:维度名称:数值
如：M:身高（cm）:160,L:身高（cm）:170
type TmallItemSizemappingTemplateCreateResponse struct {
    
    /* size_mapping_template Object尺码表模板 */
    size_mapping_template SizeMappingTemplateDo `json:"size_mapping_template";xml:"size_mapping_template"`
    
}
