package models

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
	HCode_DevLogin      = "H_Online" //设备上线
	HCode_DevOffLine    = "H_Offline" //设备离线
	HCode_CheckDevValid = "H_Valid" //查询设备是否入库
	HCode_CheckDevBind  = "H_IsBInd" //查询设备是否被绑定
	HCode_Members       = "H_BindInfo" //获取绑定人列表
	HCode_HealData      = "H_Health" //健康数据
	HCode_Chat          = "H_Chat" //聊天消息
	HCode_Locate        = "H_Locate" //定位数据
	HCode_Iccid         = "H_Iccid" //设备上传ICCID
	HCode_FuncConfig    = "H_Config" //设备上传功能配置

	HCode_VCallReq = "H_CallReq" //设备请求视频通话
	HCode_VCallRefuce = "H_CallRefuce" //设备拒接通话
	HCode_VCallCancel = "H_CallCancel" //设备取消通话
	HCode_VCallHandup = "H_CallHandup" //设备挂断通话
)

// TCP Grpc Opration Code
const (
	TCode_TcpCmd         = "T_Tcp" //指令透传 原数据发送给手表
	TCode_FindDev        = "T_FInd" //找设备
	TCode_GetOnlineCount = "T_Onlines" //获取在线设备数
	TCode_CheckOnlineDev = "T_IsOnline" //检查指定imei的设备是否在线
	TCode_Chat           = "T_Chat" //App发送聊天到设备

	TCode_VCallReq = "T_CallReq" //App请求视频通话
	TCode_VCallHandup = "T_CallHandup" //挂断
	TCode_VCallRefuce = "T_CallRefuce" //拒接
	TCode_VCallCancel = "T_CallCancel" //取消
	TCode_VCallBusy = "T_CallBusy" //忙线
)
