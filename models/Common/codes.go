package Common

import (
	"errors"
)

var (

	ErrNoHanlder = errors.New("Router not found")
)

const (
	Common_Success = 0  //成功
	Common_Fail    = -1 //失败
	Common_NoRouter = -2 //未定义路由
)

// Http Grpc Opration Code
const (
	HCode_DevLogin = "H0001"
	HCode_HealData = "H0002"
)

// TCP Grpc Opration Code
const (
	TCode_FindDev = "T0001"
)
