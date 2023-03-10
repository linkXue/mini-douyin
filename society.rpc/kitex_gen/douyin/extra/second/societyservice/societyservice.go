// Code generated by Kitex v0.4.4. DO NOT EDIT.

package societyservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	second "society.rpc/kitex_gen/douyin/extra/second"
)

func serviceInfo() *kitex.ServiceInfo {
	return societyServiceServiceInfo
}

var societyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SocietyService"
	handlerType := (*second.SocietyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ConcernAction":       kitex.NewMethodInfo(concernActionHandler, newConcernActionArgs, newConcernActionResult, false),
		"CancelConcernAction": kitex.NewMethodInfo(cancelConcernActionHandler, newCancelConcernActionArgs, newCancelConcernActionResult, false),
		"FollowList":          kitex.NewMethodInfo(followListHandler, newFollowListArgs, newFollowListResult, false),
		"FollowerList":        kitex.NewMethodInfo(followerListHandler, newFollowerListArgs, newFollowerListResult, false),
		"FriendList":          kitex.NewMethodInfo(friendListHandler, newFriendListArgs, newFriendListResult, false),
		"MessageChat":         kitex.NewMethodInfo(messageChatHandler, newMessageChatArgs, newMessageChatResult, false),
		"MessageSend":         kitex.NewMethodInfo(messageSendHandler, newMessageSendArgs, newMessageSendResult, false),
		"SocietyInfo":         kitex.NewMethodInfo(societyInfoHandler, newSocietyInfoArgs, newSocietyInfoResult, false),
		"IsFriend":            kitex.NewMethodInfo(isFriendHandler, newIsFriendArgs, newIsFriendResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin.extra.second",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func concernActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.ConcernActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).ConcernAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ConcernActionArgs:
		success, err := handler.(second.SocietyService).ConcernAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ConcernActionResult)
		realResult.Success = success
	}
	return nil
}
func newConcernActionArgs() interface{} {
	return &ConcernActionArgs{}
}

func newConcernActionResult() interface{} {
	return &ConcernActionResult{}
}

type ConcernActionArgs struct {
	Req *second.ConcernActionRequest
}

func (p *ConcernActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.ConcernActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ConcernActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ConcernActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ConcernActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ConcernActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ConcernActionArgs) Unmarshal(in []byte) error {
	msg := new(second.ConcernActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ConcernActionArgs_Req_DEFAULT *second.ConcernActionRequest

func (p *ConcernActionArgs) GetReq() *second.ConcernActionRequest {
	if !p.IsSetReq() {
		return ConcernActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ConcernActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type ConcernActionResult struct {
	Success *second.ConcernActionResponse
}

var ConcernActionResult_Success_DEFAULT *second.ConcernActionResponse

func (p *ConcernActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.ConcernActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ConcernActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ConcernActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ConcernActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ConcernActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ConcernActionResult) Unmarshal(in []byte) error {
	msg := new(second.ConcernActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ConcernActionResult) GetSuccess() *second.ConcernActionResponse {
	if !p.IsSetSuccess() {
		return ConcernActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ConcernActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.ConcernActionResponse)
}

func (p *ConcernActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func cancelConcernActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.CancelConcernActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).CancelConcernAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CancelConcernActionArgs:
		success, err := handler.(second.SocietyService).CancelConcernAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CancelConcernActionResult)
		realResult.Success = success
	}
	return nil
}
func newCancelConcernActionArgs() interface{} {
	return &CancelConcernActionArgs{}
}

func newCancelConcernActionResult() interface{} {
	return &CancelConcernActionResult{}
}

type CancelConcernActionArgs struct {
	Req *second.CancelConcernActionRequest
}

func (p *CancelConcernActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.CancelConcernActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CancelConcernActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CancelConcernActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CancelConcernActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CancelConcernActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CancelConcernActionArgs) Unmarshal(in []byte) error {
	msg := new(second.CancelConcernActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CancelConcernActionArgs_Req_DEFAULT *second.CancelConcernActionRequest

func (p *CancelConcernActionArgs) GetReq() *second.CancelConcernActionRequest {
	if !p.IsSetReq() {
		return CancelConcernActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CancelConcernActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type CancelConcernActionResult struct {
	Success *second.CancelConcernActionResponse
}

var CancelConcernActionResult_Success_DEFAULT *second.CancelConcernActionResponse

func (p *CancelConcernActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.CancelConcernActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CancelConcernActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CancelConcernActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CancelConcernActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CancelConcernActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CancelConcernActionResult) Unmarshal(in []byte) error {
	msg := new(second.CancelConcernActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CancelConcernActionResult) GetSuccess() *second.CancelConcernActionResponse {
	if !p.IsSetSuccess() {
		return CancelConcernActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CancelConcernActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.CancelConcernActionResponse)
}

func (p *CancelConcernActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.FollowListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).FollowList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowListArgs:
		success, err := handler.(second.SocietyService).FollowList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowListArgs() interface{} {
	return &FollowListArgs{}
}

func newFollowListResult() interface{} {
	return &FollowListResult{}
}

type FollowListArgs struct {
	Req *second.FollowListRequest
}

func (p *FollowListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.FollowListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowListArgs) Unmarshal(in []byte) error {
	msg := new(second.FollowListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowListArgs_Req_DEFAULT *second.FollowListRequest

func (p *FollowListArgs) GetReq() *second.FollowListRequest {
	if !p.IsSetReq() {
		return FollowListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowListResult struct {
	Success *second.FollowListResponse
}

var FollowListResult_Success_DEFAULT *second.FollowListResponse

func (p *FollowListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.FollowListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowListResult) Unmarshal(in []byte) error {
	msg := new(second.FollowListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowListResult) GetSuccess() *second.FollowListResponse {
	if !p.IsSetSuccess() {
		return FollowListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowListResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.FollowListResponse)
}

func (p *FollowListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.FollowerListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).FollowerList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowerListArgs:
		success, err := handler.(second.SocietyService).FollowerList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowerListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowerListArgs() interface{} {
	return &FollowerListArgs{}
}

func newFollowerListResult() interface{} {
	return &FollowerListResult{}
}

type FollowerListArgs struct {
	Req *second.FollowerListRequest
}

func (p *FollowerListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.FollowerListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowerListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowerListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowerListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowerListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowerListArgs) Unmarshal(in []byte) error {
	msg := new(second.FollowerListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowerListArgs_Req_DEFAULT *second.FollowerListRequest

func (p *FollowerListArgs) GetReq() *second.FollowerListRequest {
	if !p.IsSetReq() {
		return FollowerListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowerListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowerListResult struct {
	Success *second.FollowerListResponse
}

var FollowerListResult_Success_DEFAULT *second.FollowerListResponse

func (p *FollowerListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.FollowerListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowerListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowerListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowerListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowerListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowerListResult) Unmarshal(in []byte) error {
	msg := new(second.FollowerListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowerListResult) GetSuccess() *second.FollowerListResponse {
	if !p.IsSetSuccess() {
		return FollowerListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowerListResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.FollowerListResponse)
}

func (p *FollowerListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.FriendListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).FriendList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FriendListArgs:
		success, err := handler.(second.SocietyService).FriendList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FriendListResult)
		realResult.Success = success
	}
	return nil
}
func newFriendListArgs() interface{} {
	return &FriendListArgs{}
}

func newFriendListResult() interface{} {
	return &FriendListResult{}
}

type FriendListArgs struct {
	Req *second.FriendListRequest
}

func (p *FriendListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.FriendListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FriendListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FriendListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FriendListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FriendListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FriendListArgs) Unmarshal(in []byte) error {
	msg := new(second.FriendListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FriendListArgs_Req_DEFAULT *second.FriendListRequest

func (p *FriendListArgs) GetReq() *second.FriendListRequest {
	if !p.IsSetReq() {
		return FriendListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FriendListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FriendListResult struct {
	Success *second.FriendListResponse
}

var FriendListResult_Success_DEFAULT *second.FriendListResponse

func (p *FriendListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.FriendListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FriendListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FriendListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FriendListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FriendListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FriendListResult) Unmarshal(in []byte) error {
	msg := new(second.FriendListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FriendListResult) GetSuccess() *second.FriendListResponse {
	if !p.IsSetSuccess() {
		return FriendListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FriendListResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.FriendListResponse)
}

func (p *FriendListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func messageChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.MessageChatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).MessageChat(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MessageChatArgs:
		success, err := handler.(second.SocietyService).MessageChat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MessageChatResult)
		realResult.Success = success
	}
	return nil
}
func newMessageChatArgs() interface{} {
	return &MessageChatArgs{}
}

func newMessageChatResult() interface{} {
	return &MessageChatResult{}
}

type MessageChatArgs struct {
	Req *second.MessageChatRequest
}

func (p *MessageChatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.MessageChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MessageChatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MessageChatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MessageChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessageChatArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MessageChatArgs) Unmarshal(in []byte) error {
	msg := new(second.MessageChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessageChatArgs_Req_DEFAULT *second.MessageChatRequest

func (p *MessageChatArgs) GetReq() *second.MessageChatRequest {
	if !p.IsSetReq() {
		return MessageChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessageChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessageChatResult struct {
	Success *second.MessageChatResponse
}

var MessageChatResult_Success_DEFAULT *second.MessageChatResponse

func (p *MessageChatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.MessageChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MessageChatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MessageChatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MessageChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessageChatResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MessageChatResult) Unmarshal(in []byte) error {
	msg := new(second.MessageChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessageChatResult) GetSuccess() *second.MessageChatResponse {
	if !p.IsSetSuccess() {
		return MessageChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessageChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.MessageChatResponse)
}

func (p *MessageChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func messageSendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.MessageSendRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).MessageSend(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MessageSendArgs:
		success, err := handler.(second.SocietyService).MessageSend(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MessageSendResult)
		realResult.Success = success
	}
	return nil
}
func newMessageSendArgs() interface{} {
	return &MessageSendArgs{}
}

func newMessageSendResult() interface{} {
	return &MessageSendResult{}
}

type MessageSendArgs struct {
	Req *second.MessageSendRequest
}

func (p *MessageSendArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.MessageSendRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MessageSendArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MessageSendArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MessageSendArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessageSendArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MessageSendArgs) Unmarshal(in []byte) error {
	msg := new(second.MessageSendRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessageSendArgs_Req_DEFAULT *second.MessageSendRequest

func (p *MessageSendArgs) GetReq() *second.MessageSendRequest {
	if !p.IsSetReq() {
		return MessageSendArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessageSendArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessageSendResult struct {
	Success *second.MessageSendResponse
}

var MessageSendResult_Success_DEFAULT *second.MessageSendResponse

func (p *MessageSendResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.MessageSendResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MessageSendResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MessageSendResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MessageSendResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessageSendResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MessageSendResult) Unmarshal(in []byte) error {
	msg := new(second.MessageSendResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessageSendResult) GetSuccess() *second.MessageSendResponse {
	if !p.IsSetSuccess() {
		return MessageSendResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessageSendResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.MessageSendResponse)
}

func (p *MessageSendResult) IsSetSuccess() bool {
	return p.Success != nil
}

func societyInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.SocietyInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).SocietyInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SocietyInfoArgs:
		success, err := handler.(second.SocietyService).SocietyInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SocietyInfoResult)
		realResult.Success = success
	}
	return nil
}
func newSocietyInfoArgs() interface{} {
	return &SocietyInfoArgs{}
}

func newSocietyInfoResult() interface{} {
	return &SocietyInfoResult{}
}

type SocietyInfoArgs struct {
	Req *second.SocietyInfoRequest
}

func (p *SocietyInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.SocietyInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SocietyInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SocietyInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SocietyInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SocietyInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SocietyInfoArgs) Unmarshal(in []byte) error {
	msg := new(second.SocietyInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SocietyInfoArgs_Req_DEFAULT *second.SocietyInfoRequest

func (p *SocietyInfoArgs) GetReq() *second.SocietyInfoRequest {
	if !p.IsSetReq() {
		return SocietyInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SocietyInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type SocietyInfoResult struct {
	Success *second.SocietyInfoResponse
}

var SocietyInfoResult_Success_DEFAULT *second.SocietyInfoResponse

func (p *SocietyInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.SocietyInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SocietyInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SocietyInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SocietyInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SocietyInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SocietyInfoResult) Unmarshal(in []byte) error {
	msg := new(second.SocietyInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SocietyInfoResult) GetSuccess() *second.SocietyInfoResponse {
	if !p.IsSetSuccess() {
		return SocietyInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SocietyInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.SocietyInfoResponse)
}

func (p *SocietyInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func isFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(second.IsFriendRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(second.SocietyService).IsFriend(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IsFriendArgs:
		success, err := handler.(second.SocietyService).IsFriend(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsFriendResult)
		realResult.Success = success
	}
	return nil
}
func newIsFriendArgs() interface{} {
	return &IsFriendArgs{}
}

func newIsFriendResult() interface{} {
	return &IsFriendResult{}
}

type IsFriendArgs struct {
	Req *second.IsFriendRequest
}

func (p *IsFriendArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(second.IsFriendRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IsFriendArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IsFriendArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IsFriendArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IsFriendArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *IsFriendArgs) Unmarshal(in []byte) error {
	msg := new(second.IsFriendRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsFriendArgs_Req_DEFAULT *second.IsFriendRequest

func (p *IsFriendArgs) GetReq() *second.IsFriendRequest {
	if !p.IsSetReq() {
		return IsFriendArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsFriendArgs) IsSetReq() bool {
	return p.Req != nil
}

type IsFriendResult struct {
	Success *second.IsFriendResponse
}

var IsFriendResult_Success_DEFAULT *second.IsFriendResponse

func (p *IsFriendResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(second.IsFriendResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IsFriendResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IsFriendResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IsFriendResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IsFriendResult")
	}
	return proto.Marshal(p.Success)
}

func (p *IsFriendResult) Unmarshal(in []byte) error {
	msg := new(second.IsFriendResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsFriendResult) GetSuccess() *second.IsFriendResponse {
	if !p.IsSetSuccess() {
		return IsFriendResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsFriendResult) SetSuccess(x interface{}) {
	p.Success = x.(*second.IsFriendResponse)
}

func (p *IsFriendResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ConcernAction(ctx context.Context, Req *second.ConcernActionRequest) (r *second.ConcernActionResponse, err error) {
	var _args ConcernActionArgs
	_args.Req = Req
	var _result ConcernActionResult
	if err = p.c.Call(ctx, "ConcernAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelConcernAction(ctx context.Context, Req *second.CancelConcernActionRequest) (r *second.CancelConcernActionResponse, err error) {
	var _args CancelConcernActionArgs
	_args.Req = Req
	var _result CancelConcernActionResult
	if err = p.c.Call(ctx, "CancelConcernAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, Req *second.FollowListRequest) (r *second.FollowListResponse, err error) {
	var _args FollowListArgs
	_args.Req = Req
	var _result FollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, Req *second.FollowerListRequest) (r *second.FollowerListResponse, err error) {
	var _args FollowerListArgs
	_args.Req = Req
	var _result FollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, Req *second.FriendListRequest) (r *second.FriendListResponse, err error) {
	var _args FriendListArgs
	_args.Req = Req
	var _result FriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageChat(ctx context.Context, Req *second.MessageChatRequest) (r *second.MessageChatResponse, err error) {
	var _args MessageChatArgs
	_args.Req = Req
	var _result MessageChatResult
	if err = p.c.Call(ctx, "MessageChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageSend(ctx context.Context, Req *second.MessageSendRequest) (r *second.MessageSendResponse, err error) {
	var _args MessageSendArgs
	_args.Req = Req
	var _result MessageSendResult
	if err = p.c.Call(ctx, "MessageSend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SocietyInfo(ctx context.Context, Req *second.SocietyInfoRequest) (r *second.SocietyInfoResponse, err error) {
	var _args SocietyInfoArgs
	_args.Req = Req
	var _result SocietyInfoResult
	if err = p.c.Call(ctx, "SocietyInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFriend(ctx context.Context, Req *second.IsFriendRequest) (r *second.IsFriendResponse, err error) {
	var _args IsFriendArgs
	_args.Req = Req
	var _result IsFriendResult
	if err = p.c.Call(ctx, "IsFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
