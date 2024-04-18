package models

//聊天信息类型
const (
	MsgTypeUnKnown = iota
	MsgTypeVoice
	MsgTypeText
	MsgTypePic
)

//报警信息类型
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

//设备状态类型
const (
	StatusNone = iota
	StatusLowBattery
	StatusOutFence
	StatusInfence
	StatusTakeOff
	StatusStop
)

//通话类型
const (
	CallTypeAudio = iota  //音频
	CallTypeVideo //视频聊天
)

const (
	CallDevRespTypeUnknown  = iota //未获取到状态
	CallevRespTypeSuccess   //正常状态
	CallDevRespTypeHot  //手表过热
	VCallDevRespTypeInCall   //手表正在通话中
)
