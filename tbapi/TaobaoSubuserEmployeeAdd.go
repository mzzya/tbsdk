package tbsdk

// TaobaoSubuserEmployeeAddRequest 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息）
type TaobaoSubuserEmployeeAddRequest struct {
    
    /* department_id required当前员工所属部门ID */
    department_id int64 `json:"department_id";xml:"department_id"`
    
    /* duty_id optional当前员工担任职务ID */
    duty_id int64 `json:"duty_id";xml:"duty_id"`
    
    /* employee_name required员工姓名 */
    employee_name string `json:"employee_name";xml:"employee_name"`
    
    /* employee_nickname optional员工花名 */
    employee_nickname string `json:"employee_nickname";xml:"employee_nickname"`
    
    /* employee_num optional员工工号(可以用大小写英文字母和数字) */
    employee_num string `json:"employee_num";xml:"employee_num"`
    
    /* entry_date optional员工入职时间 */
    entry_date Date `json:"entry_date";xml:"entry_date"`
    
    /* id_card_num optional员工身份证号码 */
    id_card_num string `json:"id_card_num";xml:"id_card_num"`
    
    /* leader_id optional直接上级的员工ID */
    leader_id int64 `json:"leader_id";xml:"leader_id"`
    
    /* office_phone optional办公电话 */
    office_phone string `json:"office_phone";xml:"office_phone"`
    
    /* personal_email optional员工私人邮箱 */
    personal_email string `json:"personal_email";xml:"personal_email"`
    
    /* personal_mobile optional员工手机号码 */
    personal_mobile string `json:"personal_mobile";xml:"personal_mobile"`
    
    /* sex required员工性别 1：男; 2:女 */
    sex int64 `json:"sex";xml:"sex"`
    
    /* sub_id required子账号ID */
    sub_id int64 `json:"sub_id";xml:"sub_id"`
    
    /* work_location optional工作地点 */
    work_location string `json:"work_location";xml:"work_location"`
    
}

func (req *TaobaoSubuserEmployeeAddRequest) GetAPIName() string {
	return "taobao.subuser.employee.add"
}

// TaobaoSubuserEmployeeAddResponse 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息）
type TaobaoSubuserEmployeeAddResponse struct {
    
    /* is_success Basic操作是否成功 true:操作成功; false:操作失败 */
    is_success bool `json:"is_success";xml:"is_success"`
    
}
