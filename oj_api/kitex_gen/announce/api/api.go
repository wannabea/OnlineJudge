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
		"InsertAnnouncement":  kitex.NewMethodInfo(insertAnnouncementHandler, newApiInsertAnnouncementArgs, newApiInsertAnnouncementResult, false),
		"UpdateAnnouncement":  kitex.NewMethodInfo(updateAnnouncementHandler, newApiUpdateAnnouncementArgs, newApiUpdateAnnouncementResult, false),
		"DeleteAnnouncement":  kitex.NewMethodInfo(deleteAnnouncementHandler, newApiDeleteAnnouncementArgs, newApiDeleteAnnouncementResult, false),
		"HideAnnouncement":    kitex.NewMethodInfo(hideAnnouncementHandler, newApiHideAnnouncementArgs, newApiHideAnnouncementResult, false),
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

func insertAnnouncementHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*announce.ApiInsertAnnouncementArgs)
	realResult := result.(*announce.ApiInsertAnnouncementResult)
	success, err := handler.(announce.Api).InsertAnnouncement(ctx, realArg.Info, realArg.Idt)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newApiInsertAnnouncementArgs() interface{} {
	return announce.NewApiInsertAnnouncementArgs()
}

func newApiInsertAnnouncementResult() interface{} {
	return announce.NewApiInsertAnnouncementResult()
}

func updateAnnouncementHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*announce.ApiUpdateAnnouncementArgs)
	realResult := result.(*announce.ApiUpdateAnnouncementResult)
	success, err := handler.(announce.Api).UpdateAnnouncement(ctx, realArg.AnnounceId, realArg.Info, realArg.Idt)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newApiUpdateAnnouncementArgs() interface{} {
	return announce.NewApiUpdateAnnouncementArgs()
}

func newApiUpdateAnnouncementResult() interface{} {
	return announce.NewApiUpdateAnnouncementResult()
}

func deleteAnnouncementHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*announce.ApiDeleteAnnouncementArgs)
	realResult := result.(*announce.ApiDeleteAnnouncementResult)
	success, err := handler.(announce.Api).DeleteAnnouncement(ctx, realArg.Id, realArg.Idt)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newApiDeleteAnnouncementArgs() interface{} {
	return announce.NewApiDeleteAnnouncementArgs()
}

func newApiDeleteAnnouncementResult() interface{} {
	return announce.NewApiDeleteAnnouncementResult()
}

func hideAnnouncementHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*announce.ApiHideAnnouncementArgs)
	realResult := result.(*announce.ApiHideAnnouncementResult)
	success, err := handler.(announce.Api).HideAnnouncement(ctx, realArg.Op, realArg.Idt)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newApiHideAnnouncementArgs() interface{} {
	return announce.NewApiHideAnnouncementArgs()
}

func newApiHideAnnouncementResult() interface{} {
	return announce.NewApiHideAnnouncementResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetAnnouncementById(ctx context.Context, req *announce.AnnounceRequest) (r *announce.AnnounceInfo, err error) {
	var _args announce.ApiGetAnnouncementByIdArgs
	_args.Req = req
	var _result announce.ApiGetAnnouncementByIdResult
	if err = p.c.Call(ctx, "GetAnnouncementById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllAnnouncements(ctx context.Context) (r []*announce.AnnounceInfo, err error) {
	var _args announce.ApiGetAllAnnouncementsArgs
	var _result announce.ApiGetAllAnnouncementsResult
	if err = p.c.Call(ctx, "GetAllAnnouncements", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) InsertAnnouncement(ctx context.Context, info *announce.UpdateAnouncementRequest, idt *announce.Identity) (r int32, err error) {
	var _args announce.ApiInsertAnnouncementArgs
	_args.Info = info
	_args.Idt = idt
	var _result announce.ApiInsertAnnouncementResult
	if err = p.c.Call(ctx, "InsertAnnouncement", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAnnouncement(ctx context.Context, announceId int32, info *announce.UpdateAnouncementRequest, idt *announce.Identity) (r int32, err error) {
	var _args announce.ApiUpdateAnnouncementArgs
	_args.AnnounceId = announceId
	_args.Info = info
	_args.Idt = idt
	var _result announce.ApiUpdateAnnouncementResult
	if err = p.c.Call(ctx, "UpdateAnnouncement", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAnnouncement(ctx context.Context, id int32, idt *announce.Identity) (r int32, err error) {
	var _args announce.ApiDeleteAnnouncementArgs
	_args.Id = id
	_args.Idt = idt
	var _result announce.ApiDeleteAnnouncementResult
	if err = p.c.Call(ctx, "DeleteAnnouncement", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) HideAnnouncement(ctx context.Context, op int32, idt *announce.Identity) (r int32, err error) {
	var _args announce.ApiHideAnnouncementArgs
	_args.Op = op
	_args.Idt = idt
	var _result announce.ApiHideAnnouncementResult
	if err = p.c.Call(ctx, "HideAnnouncement", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
