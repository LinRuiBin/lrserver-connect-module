package models

// 查找在线设备数
type DevOnlineCountData struct {
	Count uint `json:"count"` //心率
}

//视频通话参数
type DevVideoResp struct {
	Code int `json:"code"`  //CallDevRespType
}


//视频通话参数
type ChatMsgResp struct {
	IsContinue int `json:"isContinue"`  //CallDevRespType
}

