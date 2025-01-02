// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: openmeeting/meeting/meeting.proto

package meeting

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MeetingService_BookMeeting_FullMethodName                      = "/openmeeting.meeting.MeetingService/BookMeeting"
	MeetingService_CreateImmediateMeeting_FullMethodName           = "/openmeeting.meeting.MeetingService/CreateImmediateMeeting"
	MeetingService_JoinMeeting_FullMethodName                      = "/openmeeting.meeting.MeetingService/JoinMeeting"
	MeetingService_GetMeetingToken_FullMethodName                  = "/openmeeting.meeting.MeetingService/GetMeetingToken"
	MeetingService_LeaveMeeting_FullMethodName                     = "/openmeeting.meeting.MeetingService/LeaveMeeting"
	MeetingService_EndMeeting_FullMethodName                       = "/openmeeting.meeting.MeetingService/EndMeeting"
	MeetingService_GetMeetings_FullMethodName                      = "/openmeeting.meeting.MeetingService/GetMeetings"
	MeetingService_GetMeeting_FullMethodName                       = "/openmeeting.meeting.MeetingService/GetMeeting"
	MeetingService_UpdateMeeting_FullMethodName                    = "/openmeeting.meeting.MeetingService/UpdateMeeting"
	MeetingService_GetPersonalMeetingSettings_FullMethodName       = "/openmeeting.meeting.MeetingService/GetPersonalMeetingSettings"
	MeetingService_SetPersonalMeetingSettings_FullMethodName       = "/openmeeting.meeting.MeetingService/SetPersonalMeetingSettings"
	MeetingService_OperateRoomAllStream_FullMethodName             = "/openmeeting.meeting.MeetingService/OperateRoomAllStream"
	MeetingService_ModifyMeetingParticipantNickName_FullMethodName = "/openmeeting.meeting.MeetingService/ModifyMeetingParticipantNickName"
	MeetingService_RemoveParticipants_FullMethodName               = "/openmeeting.meeting.MeetingService/RemoveParticipants"
	MeetingService_SetMeetingHostInfo_FullMethodName               = "/openmeeting.meeting.MeetingService/SetMeetingHostInfo"
	MeetingService_CleanPreviousMeetings_FullMethodName            = "/openmeeting.meeting.MeetingService/CleanPreviousMeetings"
	MeetingService_ToggleRecordMeeting_FullMethodName              = "/openmeeting.meeting.MeetingService/ToggleRecordMeeting"
)

// MeetingServiceClient is the client API for MeetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Defines services related to meeting management.
type MeetingServiceClient interface {
	// Books a future meeting.
	BookMeeting(ctx context.Context, in *BookMeetingReq, opts ...grpc.CallOption) (*BookMeetingResp, error)
	// Creates an immediate meeting.
	CreateImmediateMeeting(ctx context.Context, in *CreateImmediateMeetingReq, opts ...grpc.CallOption) (*CreateImmediateMeetingResp, error)
	// Joins a meeting.
	JoinMeeting(ctx context.Context, in *JoinMeetingReq, opts ...grpc.CallOption) (*JoinMeetingResp, error)
	// get a specific meeting token
	GetMeetingToken(ctx context.Context, in *GetMeetingTokenReq, opts ...grpc.CallOption) (*GetMeetingTokenResp, error)
	// Leaves a meeting.
	LeaveMeeting(ctx context.Context, in *LeaveMeetingReq, opts ...grpc.CallOption) (*LeaveMeetingResp, error)
	// Ends a meeting.
	EndMeeting(ctx context.Context, in *EndMeetingReq, opts ...grpc.CallOption) (*EndMeetingResp, error)
	// Retrieves a list of meetings that the user has created or joined, filtered by status.
	GetMeetings(ctx context.Context, in *GetMeetingsReq, opts ...grpc.CallOption) (*GetMeetingsResp, error)
	// Gets detailed information about a specific meeting.
	GetMeeting(ctx context.Context, in *GetMeetingReq, opts ...grpc.CallOption) (*GetMeetingResp, error)
	// Updates specific fields of a meeting.
	UpdateMeeting(ctx context.Context, in *UpdateMeetingRequest, opts ...grpc.CallOption) (*UpdateMeetingResp, error)
	// Gets personal meeting settings.
	GetPersonalMeetingSettings(ctx context.Context, in *GetPersonalMeetingSettingsReq, opts ...grpc.CallOption) (*GetPersonalMeetingSettingsResp, error)
	// Sets personal meeting settings.
	SetPersonalMeetingSettings(ctx context.Context, in *SetPersonalMeetingSettingsReq, opts ...grpc.CallOption) (*SetPersonalMeetingSettingsResp, error)
	// operate room all stream.
	OperateRoomAllStream(ctx context.Context, in *OperateRoomAllStreamReq, opts ...grpc.CallOption) (*OperateRoomAllStreamResp, error)
	// modify meeting participant nickname
	ModifyMeetingParticipantNickName(ctx context.Context, in *ModifyMeetingParticipantNickNameReq, opts ...grpc.CallOption) (*ModifyMeetingParticipantNickNameResp, error)
	// batch remove participant out of the meeting room
	RemoveParticipants(ctx context.Context, in *RemoveMeetingParticipantsReq, opts ...grpc.CallOption) (*RemoveMeetingParticipantsResp, error)
	// modify host or co host of the meeting room
	SetMeetingHostInfo(ctx context.Context, in *SetMeetingHostInfoReq, opts ...grpc.CallOption) (*SetMeetingHostInfoResp, error)
	// clean previous meetings in rtc when login
	CleanPreviousMeetings(ctx context.Context, in *CleanPreviousMeetingsReq, opts ...grpc.CallOption) (*CleanPreviousMeetingsResp, error)
	// toggle record meeting
	ToggleRecordMeeting(ctx context.Context, in *ToggleRecordMeetingReq, opts ...grpc.CallOption) (*ToggleRecordMeetingResp, error)
}

type meetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingServiceClient(cc grpc.ClientConnInterface) MeetingServiceClient {
	return &meetingServiceClient{cc}
}

func (c *meetingServiceClient) BookMeeting(ctx context.Context, in *BookMeetingReq, opts ...grpc.CallOption) (*BookMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_BookMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) CreateImmediateMeeting(ctx context.Context, in *CreateImmediateMeetingReq, opts ...grpc.CallOption) (*CreateImmediateMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateImmediateMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_CreateImmediateMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) JoinMeeting(ctx context.Context, in *JoinMeetingReq, opts ...grpc.CallOption) (*JoinMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_JoinMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetMeetingToken(ctx context.Context, in *GetMeetingTokenReq, opts ...grpc.CallOption) (*GetMeetingTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMeetingTokenResp)
	err := c.cc.Invoke(ctx, MeetingService_GetMeetingToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) LeaveMeeting(ctx context.Context, in *LeaveMeetingReq, opts ...grpc.CallOption) (*LeaveMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_LeaveMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) EndMeeting(ctx context.Context, in *EndMeetingReq, opts ...grpc.CallOption) (*EndMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EndMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_EndMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetMeetings(ctx context.Context, in *GetMeetingsReq, opts ...grpc.CallOption) (*GetMeetingsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMeetingsResp)
	err := c.cc.Invoke(ctx, MeetingService_GetMeetings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetMeeting(ctx context.Context, in *GetMeetingReq, opts ...grpc.CallOption) (*GetMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_GetMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) UpdateMeeting(ctx context.Context, in *UpdateMeetingRequest, opts ...grpc.CallOption) (*UpdateMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_UpdateMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetPersonalMeetingSettings(ctx context.Context, in *GetPersonalMeetingSettingsReq, opts ...grpc.CallOption) (*GetPersonalMeetingSettingsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPersonalMeetingSettingsResp)
	err := c.cc.Invoke(ctx, MeetingService_GetPersonalMeetingSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) SetPersonalMeetingSettings(ctx context.Context, in *SetPersonalMeetingSettingsReq, opts ...grpc.CallOption) (*SetPersonalMeetingSettingsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetPersonalMeetingSettingsResp)
	err := c.cc.Invoke(ctx, MeetingService_SetPersonalMeetingSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) OperateRoomAllStream(ctx context.Context, in *OperateRoomAllStreamReq, opts ...grpc.CallOption) (*OperateRoomAllStreamResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperateRoomAllStreamResp)
	err := c.cc.Invoke(ctx, MeetingService_OperateRoomAllStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) ModifyMeetingParticipantNickName(ctx context.Context, in *ModifyMeetingParticipantNickNameReq, opts ...grpc.CallOption) (*ModifyMeetingParticipantNickNameResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyMeetingParticipantNickNameResp)
	err := c.cc.Invoke(ctx, MeetingService_ModifyMeetingParticipantNickName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) RemoveParticipants(ctx context.Context, in *RemoveMeetingParticipantsReq, opts ...grpc.CallOption) (*RemoveMeetingParticipantsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveMeetingParticipantsResp)
	err := c.cc.Invoke(ctx, MeetingService_RemoveParticipants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) SetMeetingHostInfo(ctx context.Context, in *SetMeetingHostInfoReq, opts ...grpc.CallOption) (*SetMeetingHostInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetMeetingHostInfoResp)
	err := c.cc.Invoke(ctx, MeetingService_SetMeetingHostInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) CleanPreviousMeetings(ctx context.Context, in *CleanPreviousMeetingsReq, opts ...grpc.CallOption) (*CleanPreviousMeetingsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CleanPreviousMeetingsResp)
	err := c.cc.Invoke(ctx, MeetingService_CleanPreviousMeetings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) ToggleRecordMeeting(ctx context.Context, in *ToggleRecordMeetingReq, opts ...grpc.CallOption) (*ToggleRecordMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ToggleRecordMeetingResp)
	err := c.cc.Invoke(ctx, MeetingService_ToggleRecordMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingServiceServer is the server API for MeetingService service.
// All implementations must embed UnimplementedMeetingServiceServer
// for forward compatibility.
//
// Defines services related to meeting management.
type MeetingServiceServer interface {
	// Books a future meeting.
	BookMeeting(context.Context, *BookMeetingReq) (*BookMeetingResp, error)
	// Creates an immediate meeting.
	CreateImmediateMeeting(context.Context, *CreateImmediateMeetingReq) (*CreateImmediateMeetingResp, error)
	// Joins a meeting.
	JoinMeeting(context.Context, *JoinMeetingReq) (*JoinMeetingResp, error)
	// get a specific meeting token
	GetMeetingToken(context.Context, *GetMeetingTokenReq) (*GetMeetingTokenResp, error)
	// Leaves a meeting.
	LeaveMeeting(context.Context, *LeaveMeetingReq) (*LeaveMeetingResp, error)
	// Ends a meeting.
	EndMeeting(context.Context, *EndMeetingReq) (*EndMeetingResp, error)
	// Retrieves a list of meetings that the user has created or joined, filtered by status.
	GetMeetings(context.Context, *GetMeetingsReq) (*GetMeetingsResp, error)
	// Gets detailed information about a specific meeting.
	GetMeeting(context.Context, *GetMeetingReq) (*GetMeetingResp, error)
	// Updates specific fields of a meeting.
	UpdateMeeting(context.Context, *UpdateMeetingRequest) (*UpdateMeetingResp, error)
	// Gets personal meeting settings.
	GetPersonalMeetingSettings(context.Context, *GetPersonalMeetingSettingsReq) (*GetPersonalMeetingSettingsResp, error)
	// Sets personal meeting settings.
	SetPersonalMeetingSettings(context.Context, *SetPersonalMeetingSettingsReq) (*SetPersonalMeetingSettingsResp, error)
	// operate room all stream.
	OperateRoomAllStream(context.Context, *OperateRoomAllStreamReq) (*OperateRoomAllStreamResp, error)
	// modify meeting participant nickname
	ModifyMeetingParticipantNickName(context.Context, *ModifyMeetingParticipantNickNameReq) (*ModifyMeetingParticipantNickNameResp, error)
	// batch remove participant out of the meeting room
	RemoveParticipants(context.Context, *RemoveMeetingParticipantsReq) (*RemoveMeetingParticipantsResp, error)
	// modify host or co host of the meeting room
	SetMeetingHostInfo(context.Context, *SetMeetingHostInfoReq) (*SetMeetingHostInfoResp, error)
	// clean previous meetings in rtc when login
	CleanPreviousMeetings(context.Context, *CleanPreviousMeetingsReq) (*CleanPreviousMeetingsResp, error)
	// toggle record meeting
	ToggleRecordMeeting(context.Context, *ToggleRecordMeetingReq) (*ToggleRecordMeetingResp, error)
	mustEmbedUnimplementedMeetingServiceServer()
}

// UnimplementedMeetingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMeetingServiceServer struct{}

func (UnimplementedMeetingServiceServer) BookMeeting(context.Context, *BookMeetingReq) (*BookMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) CreateImmediateMeeting(context.Context, *CreateImmediateMeetingReq) (*CreateImmediateMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImmediateMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) JoinMeeting(context.Context, *JoinMeetingReq) (*JoinMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) GetMeetingToken(context.Context, *GetMeetingTokenReq) (*GetMeetingTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingToken not implemented")
}
func (UnimplementedMeetingServiceServer) LeaveMeeting(context.Context, *LeaveMeetingReq) (*LeaveMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) EndMeeting(context.Context, *EndMeetingReq) (*EndMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) GetMeetings(context.Context, *GetMeetingsReq) (*GetMeetingsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetings not implemented")
}
func (UnimplementedMeetingServiceServer) GetMeeting(context.Context, *GetMeetingReq) (*GetMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) UpdateMeeting(context.Context, *UpdateMeetingRequest) (*UpdateMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) GetPersonalMeetingSettings(context.Context, *GetPersonalMeetingSettingsReq) (*GetPersonalMeetingSettingsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalMeetingSettings not implemented")
}
func (UnimplementedMeetingServiceServer) SetPersonalMeetingSettings(context.Context, *SetPersonalMeetingSettingsReq) (*SetPersonalMeetingSettingsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPersonalMeetingSettings not implemented")
}
func (UnimplementedMeetingServiceServer) OperateRoomAllStream(context.Context, *OperateRoomAllStreamReq) (*OperateRoomAllStreamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateRoomAllStream not implemented")
}
func (UnimplementedMeetingServiceServer) ModifyMeetingParticipantNickName(context.Context, *ModifyMeetingParticipantNickNameReq) (*ModifyMeetingParticipantNickNameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMeetingParticipantNickName not implemented")
}
func (UnimplementedMeetingServiceServer) RemoveParticipants(context.Context, *RemoveMeetingParticipantsReq) (*RemoveMeetingParticipantsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveParticipants not implemented")
}
func (UnimplementedMeetingServiceServer) SetMeetingHostInfo(context.Context, *SetMeetingHostInfoReq) (*SetMeetingHostInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMeetingHostInfo not implemented")
}
func (UnimplementedMeetingServiceServer) CleanPreviousMeetings(context.Context, *CleanPreviousMeetingsReq) (*CleanPreviousMeetingsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanPreviousMeetings not implemented")
}
func (UnimplementedMeetingServiceServer) ToggleRecordMeeting(context.Context, *ToggleRecordMeetingReq) (*ToggleRecordMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleRecordMeeting not implemented")
}
func (UnimplementedMeetingServiceServer) mustEmbedUnimplementedMeetingServiceServer() {}
func (UnimplementedMeetingServiceServer) testEmbeddedByValue()                        {}

// UnsafeMeetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingServiceServer will
// result in compilation errors.
type UnsafeMeetingServiceServer interface {
	mustEmbedUnimplementedMeetingServiceServer()
}

func RegisterMeetingServiceServer(s grpc.ServiceRegistrar, srv MeetingServiceServer) {
	// If the following call pancis, it indicates UnimplementedMeetingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MeetingService_ServiceDesc, srv)
}

func _MeetingService_BookMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).BookMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_BookMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).BookMeeting(ctx, req.(*BookMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_CreateImmediateMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImmediateMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).CreateImmediateMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_CreateImmediateMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).CreateImmediateMeeting(ctx, req.(*CreateImmediateMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_JoinMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).JoinMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_JoinMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).JoinMeeting(ctx, req.(*JoinMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetMeetingToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetMeetingToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_GetMeetingToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetMeetingToken(ctx, req.(*GetMeetingTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_LeaveMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).LeaveMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_LeaveMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).LeaveMeeting(ctx, req.(*LeaveMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_EndMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).EndMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_EndMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).EndMeeting(ctx, req.(*EndMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetMeetings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetMeetings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_GetMeetings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetMeetings(ctx, req.(*GetMeetingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_GetMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetMeeting(ctx, req.(*GetMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_UpdateMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).UpdateMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_UpdateMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).UpdateMeeting(ctx, req.(*UpdateMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetPersonalMeetingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonalMeetingSettingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetPersonalMeetingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_GetPersonalMeetingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetPersonalMeetingSettings(ctx, req.(*GetPersonalMeetingSettingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_SetPersonalMeetingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPersonalMeetingSettingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).SetPersonalMeetingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_SetPersonalMeetingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).SetPersonalMeetingSettings(ctx, req.(*SetPersonalMeetingSettingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_OperateRoomAllStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRoomAllStreamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).OperateRoomAllStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_OperateRoomAllStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).OperateRoomAllStream(ctx, req.(*OperateRoomAllStreamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_ModifyMeetingParticipantNickName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMeetingParticipantNickNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).ModifyMeetingParticipantNickName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_ModifyMeetingParticipantNickName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).ModifyMeetingParticipantNickName(ctx, req.(*ModifyMeetingParticipantNickNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_RemoveParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMeetingParticipantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).RemoveParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_RemoveParticipants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).RemoveParticipants(ctx, req.(*RemoveMeetingParticipantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_SetMeetingHostInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMeetingHostInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).SetMeetingHostInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_SetMeetingHostInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).SetMeetingHostInfo(ctx, req.(*SetMeetingHostInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_CleanPreviousMeetings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanPreviousMeetingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).CleanPreviousMeetings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_CleanPreviousMeetings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).CleanPreviousMeetings(ctx, req.(*CleanPreviousMeetingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_ToggleRecordMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleRecordMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).ToggleRecordMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeetingService_ToggleRecordMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).ToggleRecordMeeting(ctx, req.(*ToggleRecordMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MeetingService_ServiceDesc is the grpc.ServiceDesc for MeetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openmeeting.meeting.MeetingService",
	HandlerType: (*MeetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookMeeting",
			Handler:    _MeetingService_BookMeeting_Handler,
		},
		{
			MethodName: "CreateImmediateMeeting",
			Handler:    _MeetingService_CreateImmediateMeeting_Handler,
		},
		{
			MethodName: "JoinMeeting",
			Handler:    _MeetingService_JoinMeeting_Handler,
		},
		{
			MethodName: "GetMeetingToken",
			Handler:    _MeetingService_GetMeetingToken_Handler,
		},
		{
			MethodName: "LeaveMeeting",
			Handler:    _MeetingService_LeaveMeeting_Handler,
		},
		{
			MethodName: "EndMeeting",
			Handler:    _MeetingService_EndMeeting_Handler,
		},
		{
			MethodName: "GetMeetings",
			Handler:    _MeetingService_GetMeetings_Handler,
		},
		{
			MethodName: "GetMeeting",
			Handler:    _MeetingService_GetMeeting_Handler,
		},
		{
			MethodName: "UpdateMeeting",
			Handler:    _MeetingService_UpdateMeeting_Handler,
		},
		{
			MethodName: "GetPersonalMeetingSettings",
			Handler:    _MeetingService_GetPersonalMeetingSettings_Handler,
		},
		{
			MethodName: "SetPersonalMeetingSettings",
			Handler:    _MeetingService_SetPersonalMeetingSettings_Handler,
		},
		{
			MethodName: "OperateRoomAllStream",
			Handler:    _MeetingService_OperateRoomAllStream_Handler,
		},
		{
			MethodName: "ModifyMeetingParticipantNickName",
			Handler:    _MeetingService_ModifyMeetingParticipantNickName_Handler,
		},
		{
			MethodName: "RemoveParticipants",
			Handler:    _MeetingService_RemoveParticipants_Handler,
		},
		{
			MethodName: "SetMeetingHostInfo",
			Handler:    _MeetingService_SetMeetingHostInfo_Handler,
		},
		{
			MethodName: "CleanPreviousMeetings",
			Handler:    _MeetingService_CleanPreviousMeetings_Handler,
		},
		{
			MethodName: "ToggleRecordMeeting",
			Handler:    _MeetingService_ToggleRecordMeeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openmeeting/meeting/meeting.proto",
}