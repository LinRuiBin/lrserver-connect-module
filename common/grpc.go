package common


import (

"gitee.com/linuibin/lrserver-connect-module/grpcs/HTServerGrpc"
	"gitee.com/linuibin/lrserver-connect-module/models/Common"
)

func RespWithOffline(request *HTServerGrpc.HTRequest) (reply *HTServerGrpc.HTReply, err error) {


	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
	}, Common.ErrDeviceOffline
}

func RespWithTimeout(request *HTServerGrpc.HTRequest) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
	}, Common.ErrHandleTimeout
}

func GrpcSuccessResp(request *HTServerGrpc.HTRequest) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
	}, nil
}

func GrpcSuccessRespWithJsonData(request *HTServerGrpc.HTRequest, data []byte) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
		Data: data,
	}, nil
}

func GrpcRespWithError(request *HTServerGrpc.HTRequest, inErr error) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
	}, inErr
}
