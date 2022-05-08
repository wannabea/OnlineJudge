// Code generated by Kitex v0.3.0. DO NOT EDIT.

package api

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/wannabea/OnlineJudge/oj_api/kitex_gen/announce"
)

func serviceInfo() *kitex.ServiceInfo {
	return apiServiceInfo
}

var apiServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Api"
	handlerType := (*announce.Api)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetAnnouncementById": kitex.NewMethodInfo(getAnnouncementByIdHandler, newApiGetAnnouncementByIdArgs, newApiGetAnnouncementByIdResult, false),
		"GetAllAnnouncements": kitex.NewMethodInfo(getAllAnnouncementsHandler, newApiGetAllAnnouncementsArgs, newApiGetAllAnnouncementsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "announce",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.0",
		Extra:           extra,
	}
	return svcInfo
}

func getAnnouncementByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*announce.ApiGetAnnouncementByIdArgs)
	realResult := result.(*announce.ApiGetAnnouncementByIdResult)
	success, err := handler.(announce.Api).GetAnnouncementById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiGetAnnouncementByIdArgs() interface{} {
	return announce.NewApiGetAnnouncementByIdArgs()
}

func newApiGetAnnouncementByIdResult() interface{} {
	return announce.NewApiGetAnnouncementByIdResult()
}

func getAllAnnouncementsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*announce.ApiGetAllAnnouncementsResult)
	success, err := handler.(announce.Api).GetAllAnnouncements(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiGetAllAnnouncementsArgs() interface{} {
	return announce.NewApiGetAllAnnouncementsArgs()
}

func newApiGetAllAnnouncementsResult() interface{} {
	return announce.NewApiGetAllAnnouncementsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetAnnouncementById(ctx context.Context, req *announce.AnnounceRequest) (r *announce.AnnounceResponse, err error) {
	var _args announce.ApiGetAnnouncementByIdArgs
	_args.Req = req
	var _result announce.ApiGetAnnouncementByIdResult
	if err = p.c.Call(ctx, "GetAnnouncementById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllAnnouncements(ctx context.Context) (r []*announce.AnnounceResponse, err error) {
	var _args announce.ApiGetAllAnnouncementsArgs
	var _result announce.ApiGetAllAnnouncementsResult
	if err = p.c.Call(ctx, "GetAllAnnouncements", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}