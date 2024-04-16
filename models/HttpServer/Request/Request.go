package Request

const (
	MsgTypeUnKnown = iota
	MsgTypeVoice
	MsgTypeText
	MsgTypePic
)

const (
	AlarmTypeNone = iota
	AlarmTypeSOS
	AlarmTypeLowBattery
	AlarmTypeOutFence
	AlarmTypeInFence
	AlarmTypeTakeOff
	AlarmTypeFallDown
	AlarmTypeFallWater
)

const (
	StatusNone = iota
	StatusLowBattery
	StatusOutFence
	StatusInfence
	StatusTakeOff
	StatusStop
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

// 定位数组命令解析结果字符串
type LocateCmd struct {
	Timestamp      uint      `json:"timestamp" form:"timestamp" `         //时间转为时间戳
	IsGpsValid     bool        `json:"isGpsValid" form:"isGpsValid" `      //gps标志位 A有效 V无效
	Latitude       float64     `json:"latitude" form:"latitude" `      //纬度
	Longitude      float64     `json:"longitude" form:"longitude" `      //经度
	Speed          float64     `json:"speed" form:"speed" `      //速度
	Direction      float64     `json:"direction" form:"direction" `      //方向
	Altitude       float64     `json:"altitude" form:"altitude" `      //海拔
	SatelliteCount int          `json:"satelliteCount" form:"satelliteCount" `     //卫星个数
	GsmSignal      int          `json:"gsmSignal" form:"gsmSignal" `     //*3/100-110后的值
	//Battery        int           `json:"battery" form:"battery" `    //电量
	//Step           int           `json:"step" form:"step" `    //步数 定位包不处理 时区有问题
	//RollS          int           `json:"rollS" form:"rollS" `    // 翻滚次数 定位包不处理 时区有问题
	//Alarms         string        `json:"alarms" form:"alarms" `    //设备状态和报警状态位 转为 DevStatus AlarmType
	Mnc            string        `json:"mnc" form:"mnc" `    //Mobile Network Code，移动网络号码（中国移动为00，中国联通为01）；
	Mcc            string       `json:"mcc" form:"mcc" `     //Mobile Country Code，移动国家代码（中国的为460）；
	LbsArr         []LBSCmdData `json:"lbsArr" form:"lbsArr" `   //基站数据数组
	WifiArr        []WifiCmdData `json:"wifiArr" form:"wifiArr" `  //wifi数据数组
	Accuracy       float64        `json:"accuracy" form:"accuracy" `   //精确度
	DevStatus      int            `json:"devStatus" form:"devStatus" `   //设备状态
	AlarmType      int             `json:"alarmType" form:"alarmType" `  //报警状态

}

type LBSCmdData struct {
	Lac    string `json:"lac" form:"lac" `  //Location Area Code，位置区域码；
	Cid    string `json:"cid" form:"cid" `  // Cell Identity，基站编号，是个16位的数据（范围是0到65535）。
	Signal int    `json:"signal" form:"signal" `  // -220 后的值
}

type WifiCmdData struct {
	Name   string `json:"name" form:"name" `  //ssid wifi名称
	Mac    string `json:"mac" form:"mac" `  //mac地址
	Signal int    `json:"signal" form:"signal" `  //信号强度
}

