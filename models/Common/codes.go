package Common

import (
	"errors"
)

var (
	ErrDataErr        = errors.New("Data error")
	ErrNoHanlder      = errors.New("Router not found")
	ErrDeviceOffline  = errors.New("Device Offline")
	ErrDeviceHandling = errors.New("Request is handling")
	ErrHandleTimeout  = errors.New("Request time out")
)

const (
	Common_Success = 0  //成功
	Common_Fail    = -1 //失败
	//Common_NoRouter = -2 //未定义路由
)

// Http Grpc Opration Code
const (
	HCode_DevLogin        = "H0001" //设备上线
	HCode_DevOffLine      = "H0002" //设备离线
	HCode_CheckDevValid   = "H0003" //查询设备是否入库
	HCode_CheckDevBind    = "H0004" //查询设备是否被绑定
	HCode_GetDevBindUsers = "H0005" //获取绑定人列表
	HCode_HealData        = "H0006" //健康数据
	HCode_Chat            = "H0007" //聊天消息
	HCode_Locate            = "H0008" //定位数据
)

// TCP Grpc Opration Code
const (
	TCode_TcpCmd         = "T0000" //指令透传 原数据发送给手表
	TCode_FindDev        = "T0001" //找设备
	TCode_GetOnlineCount = "T0002" //获取在线设备数
	TCode_CheckOnlineDev = "T0003" //检查指定imei的设备是否在线
	TCode_Chat           = "T0004" //App发送聊天到设备
)
