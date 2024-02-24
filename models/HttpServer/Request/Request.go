package Request

// tcp - http 设备同步健康数据给服务器
type DevHealthData struct {
	Hr        uint   `json:"hr"` //心率
	Sbp       uint   `json:"sbp"` //舒张压
	Dbp       uint   `json:"dbp"` //收缩压
	Temp      float32 `json:"temp"`//体温
	TimeStamp uint    `json:"timeStamp"`//时间戳
}
