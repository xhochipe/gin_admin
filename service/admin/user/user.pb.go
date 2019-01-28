// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//创建和登录用户参数结构体
type UserParams struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	GroupId              string   `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserParams) Reset()         { *m = UserParams{} }
func (m *UserParams) String() string { return proto.CompactTextString(m) }
func (*UserParams) ProtoMessage()    {}
func (*UserParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *UserParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserParams.Unmarshal(m, b)
}
func (m *UserParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserParams.Marshal(b, m, deterministic)
}
func (m *UserParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserParams.Merge(m, src)
}
func (m *UserParams) XXX_Size() int {
	return xxx_messageInfo_UserParams.Size(m)
}
func (m *UserParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UserParams.DiscardUnknown(m)
}

var xxx_messageInfo_UserParams proto.InternalMessageInfo

func (m *UserParams) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserParams) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *UserParams) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

//修改用户和删除用户
//其实如果生成uid的时候把groupId包含进去，这时候只需要一个uid即可
type UpdateParams struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	GroupId              string   `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateParams) Reset()         { *m = UpdateParams{} }
func (m *UpdateParams) String() string { return proto.CompactTextString(m) }
func (*UpdateParams) ProtoMessage()    {}
func (*UpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *UpdateParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateParams.Unmarshal(m, b)
}
func (m *UpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateParams.Marshal(b, m, deterministic)
}
func (m *UpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateParams.Merge(m, src)
}
func (m *UpdateParams) XXX_Size() int {
	return xxx_messageInfo_UpdateParams.Size(m)
}
func (m *UpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateParams proto.InternalMessageInfo

func (m *UpdateParams) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateParams) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

//查询用户 可以根据groupId,账号,手机号码查询
type GetParams struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	GroupId              string   `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Ps                   int64    `protobuf:"varint,5,opt,name=ps,proto3" json:"ps,omitempty"`
	Pn                   int64    `protobuf:"varint,6,opt,name=pn,proto3" json:"pn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParams) Reset()         { *m = GetParams{} }
func (m *GetParams) String() string { return proto.CompactTextString(m) }
func (*GetParams) ProtoMessage()    {}
func (*GetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *GetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParams.Unmarshal(m, b)
}
func (m *GetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParams.Marshal(b, m, deterministic)
}
func (m *GetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParams.Merge(m, src)
}
func (m *GetParams) XXX_Size() int {
	return xxx_messageInfo_GetParams.Size(m)
}
func (m *GetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetParams proto.InternalMessageInfo

func (m *GetParams) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GetParams) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *GetParams) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *GetParams) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *GetParams) GetPs() int64 {
	if m != nil {
		return m.Ps
	}
	return 0
}

func (m *GetParams) GetPn() int64 {
	if m != nil {
		return m.Pn
	}
	return 0
}

type AuthUser struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUser) Reset()         { *m = AuthUser{} }
func (m *AuthUser) String() string { return proto.CompactTextString(m) }
func (*AuthUser) ProtoMessage()    {}
func (*AuthUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
}

func (m *AuthUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthUser.Unmarshal(m, b)
}
func (m *AuthUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthUser.Marshal(b, m, deterministic)
}
func (m *AuthUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUser.Merge(m, src)
}
func (m *AuthUser) XXX_Size() int {
	return xxx_messageInfo_AuthUser.Size(m)
}
func (m *AuthUser) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUser.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUser proto.InternalMessageInfo

func (m *AuthUser) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AuthUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//用户信息
type User struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Pwd                  string   `protobuf:"bytes,4,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Img                  string   `protobuf:"bytes,8,opt,name=img,proto3" json:"img,omitempty"`
	IsValid              int64    `protobuf:"varint,9,opt,name=isValid,proto3" json:"isValid,omitempty"`
	IsDel                int64    `protobuf:"varint,10,opt,name=isDel,proto3" json:"isDel,omitempty"`
	Create               int64    `protobuf:"varint,6,opt,name=create,proto3" json:"create,omitempty"`
	Update               int64    `protobuf:"varint,7,opt,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *User) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *User) GetIsValid() int64 {
	if m != nil {
		return m.IsValid
	}
	return 0
}

func (m *User) GetIsDel() int64 {
	if m != nil {
		return m.IsDel
	}
	return 0
}

func (m *User) GetCreate() int64 {
	if m != nil {
		return m.Create
	}
	return 0
}

func (m *User) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

//这里返回值有点问题
type Response struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*UserParams)(nil), "UserParams")
	proto.RegisterType((*UpdateParams)(nil), "UpdateParams")
	proto.RegisterType((*GetParams)(nil), "GetParams")
	proto.RegisterType((*AuthUser)(nil), "AuthUser")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x8a, 0xd4, 0x40,
	0x14, 0x35, 0xaf, 0x4e, 0xe7, 0x8e, 0x8f, 0xa6, 0x18, 0xa4, 0x70, 0x21, 0x63, 0x74, 0xd1, 0xab,
	0x28, 0xe3, 0xce, 0xdd, 0xc0, 0xc0, 0x20, 0xb8, 0x90, 0xc8, 0xb8, 0x8f, 0xa9, 0x4b, 0xa6, 0x30,
	0x49, 0x15, 0x55, 0x15, 0xfd, 0x04, 0x3f, 0xc0, 0x3f, 0xf4, 0x4b, 0xa4, 0x1e, 0x99, 0x4e, 0x68,
	0x5b, 0x74, 0x13, 0xee, 0xb9, 0x39, 0xf7, 0xe4, 0xdc, 0x47, 0xe0, 0xc9, 0xa4, 0x51, 0xbd, 0xb6,
	0x8f, 0x4a, 0x2a, 0x61, 0x44, 0x59, 0x03, 0xdc, 0x6a, 0x54, 0x1f, 0x1b, 0xd5, 0x0c, 0x9a, 0x50,
	0xc8, 0x9b, 0xb6, 0x15, 0xd3, 0x68, 0x68, 0x74, 0x11, 0xed, 0x8b, 0x7a, 0x86, 0x64, 0x07, 0x89,
	0xfc, 0xce, 0x68, 0xec, 0xb2, 0x36, 0xb4, 0xdc, 0x4e, 0x89, 0x49, 0xbe, 0x67, 0x34, 0xf1, 0xdc,
	0x00, 0xcb, 0x77, 0xf0, 0xf0, 0x56, 0xb2, 0xc6, 0x60, 0x50, 0xdd, 0x41, 0x32, 0x71, 0x16, 0x14,
	0x6d, 0xb8, 0xac, 0x8d, 0xd7, 0xb5, 0x3f, 0x22, 0x28, 0x6e, 0xd0, 0xfc, 0x7f, 0xe5, 0xd2, 0x7b,
	0xb2, 0xf6, 0x7e, 0x0e, 0x99, 0xbc, 0x13, 0x23, 0xd2, 0xd4, 0xe5, 0x3d, 0x20, 0x8f, 0x21, 0x96,
	0x9a, 0x66, 0x17, 0xd1, 0x3e, 0xa9, 0x63, 0xa9, 0x1d, 0x1e, 0xe9, 0x26, 0xe0, 0xb1, 0xbc, 0x84,
	0xed, 0xd5, 0x64, 0xee, 0xec, 0x74, 0xfe, 0xe0, 0xe3, 0x1c, 0x32, 0x23, 0xbe, 0xe2, 0x18, 0x5c,
	0x78, 0x50, 0xfe, 0x8a, 0x20, 0x3d, 0x51, 0xb0, 0xb0, 0x17, 0x1f, 0xd9, 0xc3, 0xa1, 0xe1, 0x7d,
	0xb0, 0xed, 0xc1, 0x3c, 0xf0, 0xf4, 0x30, 0xf0, 0xfb, 0x36, 0xb2, 0x65, 0x1b, 0x3b, 0x48, 0xf8,
	0xd0, 0xd1, 0xad, 0xe7, 0xf1, 0xa1, 0xb3, 0x5f, 0xe2, 0xfa, 0x73, 0xd3, 0x73, 0x46, 0x0b, 0xd7,
	0xcd, 0x0c, 0xad, 0x02, 0xd7, 0xd7, 0xd8, 0x53, 0x70, 0x79, 0x0f, 0xc8, 0x53, 0xd8, 0xb4, 0x0a,
	0x1b, 0x83, 0xa1, 0xf9, 0x80, 0x6c, 0x7e, 0x72, 0x6b, 0xa4, 0xb9, 0xcf, 0x7b, 0x54, 0xbe, 0x81,
	0x6d, 0x8d, 0x5a, 0x8a, 0x51, 0x23, 0x21, 0x90, 0xb6, 0x82, 0x61, 0x68, 0xd4, 0xc5, 0xd6, 0xd1,
	0xa0, 0xbb, 0xf9, 0x54, 0x06, 0xdd, 0x5d, 0xfe, 0x8c, 0xe1, 0xcc, 0x8e, 0xe5, 0x13, 0xaa, 0x6f,
	0xbc, 0x45, 0xf2, 0xca, 0x2a, 0x74, 0x5c, 0x1b, 0x54, 0xe4, 0xac, 0x3a, 0xdc, 0xdf, 0xb3, 0xa2,
	0x9a, 0x95, 0xcb, 0x07, 0xe4, 0x05, 0x64, 0x1f, 0x44, 0xc7, 0xc7, 0xbf, 0x50, 0x9e, 0x43, 0x6a,
	0x77, 0x44, 0x8a, 0x6a, 0x5e, 0xd5, 0xfa, 0xfd, 0x4b, 0xc8, 0xaf, 0x18, 0x73, 0x1b, 0x39, 0x2d,
	0xb2, 0x07, 0xf0, 0xe7, 0xea, 0x78, 0x8f, 0xaa, 0xe5, 0xed, 0x1e, 0x31, 0xaf, 0xb1, 0xc7, 0x7f,
	0x60, 0x96, 0x90, 0xdf, 0xa0, 0x71, 0x34, 0xa8, 0xee, 0xef, 0x79, 0xc5, 0xf9, 0xb2, 0x71, 0x7f,
	0xe0, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x31, 0x0f, 0xa8, 0x94, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	//用户注册
	Register(ctx context.Context, in *UserParams, opts ...client.CallOption) (*Response, error)
	//用户登录
	Login(ctx context.Context, in *UserParams, opts ...client.CallOption) (*Response, error)
	//用户认证(有可能单独api需要鉴权访问)
	Auth(ctx context.Context, in *AuthUser, opts ...client.CallOption) (*Response, error)
	//添加用户
	AddUser(ctx context.Context, in *UserParams, opts ...client.CallOption) (*Response, error)
	//修改用户信息
	UpdateUser(ctx context.Context, in *UpdateParams, opts ...client.CallOption) (*Response, error)
	//删除用户(逻辑删除)
	DeleteUser(ctx context.Context, in *UpdateParams, opts ...client.CallOption) (*Response, error)
	//查询用户(可以查询所有，也可以根据分组查询)
	GetUser(ctx context.Context, in *GetParams, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "userservice"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Register(ctx context.Context, in *UserParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Register", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *UserParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Login", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *AuthUser, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUser(ctx context.Context, in *UserParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.AddUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.UpdateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UpdateParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.DeleteUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//用户注册
	Register(context.Context, *UserParams, *Response) error
	//用户登录
	Login(context.Context, *UserParams, *Response) error
	//用户认证(有可能单独api需要鉴权访问)
	Auth(context.Context, *AuthUser, *Response) error
	//添加用户
	AddUser(context.Context, *UserParams, *Response) error
	//修改用户信息
	UpdateUser(context.Context, *UpdateParams, *Response) error
	//删除用户(逻辑删除)
	DeleteUser(context.Context, *UpdateParams, *Response) error
	//查询用户(可以查询所有，也可以根据分组查询)
	GetUser(context.Context, *GetParams, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Register(ctx context.Context, in *UserParams, out *Response) error {
	return h.UserServiceHandler.Register(ctx, in, out)
}

func (h *UserService) Login(ctx context.Context, in *UserParams, out *Response) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *AuthUser, out *Response) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) AddUser(ctx context.Context, in *UserParams, out *Response) error {
	return h.UserServiceHandler.AddUser(ctx, in, out)
}

func (h *UserService) UpdateUser(ctx context.Context, in *UpdateParams, out *Response) error {
	return h.UserServiceHandler.UpdateUser(ctx, in, out)
}

func (h *UserService) DeleteUser(ctx context.Context, in *UpdateParams, out *Response) error {
	return h.UserServiceHandler.DeleteUser(ctx, in, out)
}

func (h *UserService) GetUser(ctx context.Context, in *GetParams, out *Response) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}