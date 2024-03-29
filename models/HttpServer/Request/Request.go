package Request

const (
	MsgTypeUnKnown = iota
	MsgTypeVoice
	MsgTypeText
	MsgTypePic
)


// tcp - http 设备同步健康数据给服务器
type DevHealthData struct {
	Hr        uint   `json:"hr"` //心率
	Sbp       uint   `json:"sbp"` //舒张压
	Dbp       uint   `json:"dbp"` //收缩压
	Temp      float32 `json:"temp"`//体温
	TimeStamp uint    `json:"timeStamp"`//时间戳
}

// tcp <-> http 设备发送语音聊天数据给服务器
type DevChatData struct {

	IsGroup    uint    `json:"isGroup" form:"isGroup" `      //是否为群聊
	UserID    int    `json:"userId" form:"userId" `      //用户id
	Type      uint    `json:"type" form:"type" `          //聊天类型
	Content   string `json:"content" form:"content" `    //数据内容
	Duration  uint    `json:"duration" form:"duration" `  //录音时长
}
