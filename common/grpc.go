package common

import (
	"github.com/linuibin/lrserver-connect-module/grpcs/HTServerGrpc"
	"github.com/linuibin/lrserver-connect-module/models"
)

func GrpcRespWithData(request *HTServerGrpc.HTRequest, data []byte, inErr error) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
		Data: data,
	}, inErr
}

func RespWithOffline(request *HTServerGrpc.HTRequest) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
	}, models.ErrDeviceOffline
}

func RespWithTimeout(request *HTServerGrpc.HTRequest) (reply *HTServerGrpc.HTReply, err error) {

	return &HTServerGrpc.HTReply{
		Imei: request.Imei,
	}, models.ErrHandleTimeout
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
