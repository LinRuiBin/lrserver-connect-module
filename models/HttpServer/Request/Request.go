package Request

// tcp - http 设备同步健康数据给服务器
type DevHealthData struct {
	Hr        uint   `json:"hr"` //心率
	Sbp       uint   `json:"sbp"` //舒张压
	Dbp       uint   `json:"dbp"` //收缩压
	Temp      float32 `json:"temp"`//体温
	TimeStamp uint    `json:"timeStamp"`//时间戳
}

// tcp - http 设备发送语音聊天数据给服务器
type DevchatVoiceData struct {
	Url        string   `json:"url"` //文件地址
	UserID       string   `json:"userID"` // userID 单聊的时候为用户id 群聊时无效
	Type       uint   `json:"type"` //0 单聊  1群聊
	Duration       uint   `json:"duration"` //时长
}
