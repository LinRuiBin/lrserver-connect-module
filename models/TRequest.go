package models

// http - tcp 电话本参数
type DevPhoneBook struct {
	name   string
	number string
}

// 查询设备是否在线
type DevOnlineReq struct {
	Imeis []string `json:"imeis"` //传入imei 数组 返回设备在线状态
}

//视频通话参数
type DevVideoReq struct {
	TicketId string `json:"ticketId"`
	RoomId string `json:"roomId"`
	Type int `json:"type"`   //CallType 呼叫类型 1为视频
}

