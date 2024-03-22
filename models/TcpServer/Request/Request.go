package Request

// http - tcp 电话本参数
type DevPhoneBook struct {
	name   string
	number string
}

// 查询设备是否在线
type DevOnlineReq struct {
	Imeis []string `json:"imeis"` //传入imei 数组 返回设备在线状态
}

