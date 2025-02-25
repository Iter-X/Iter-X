// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Environment int32

const (
	Environment_UNKNOWN Environment = 0
	Environment_DEV     Environment = 1
	Environment_PROD    Environment = 2
	Environment_TEST    Environment = 3
)

// Enum value maps for Environment.
var (
	Environment_name = map[int32]string{
		0: "UNKNOWN",
		1: "DEV",
		2: "PROD",
		3: "TEST",
	}
	Environment_value = map[string]int32{
		"UNKNOWN": 0,
		"DEV":     1,
		"PROD":    2,
		"TEST":    3,
	}
)

func (x Environment) Enum() *Environment {
	p := new(Environment)
	*p = x
	return p
}

func (x Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_conf_proto_enumTypes[0].Descriptor()
}

func (Environment) Type() protoreflect.EnumType {
	return &file_conf_proto_enumTypes[0]
}

func (x Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Environment.Descriptor instead.
func (Environment) EnumDescriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

type I18NFormat int32

const (
	I18NFormat_TOML I18NFormat = 0
	I18NFormat_JSON I18NFormat = 1
)

// Enum value maps for I18NFormat.
var (
	I18NFormat_name = map[int32]string{
		0: "TOML",
		1: "JSON",
	}
	I18NFormat_value = map[string]int32{
		"TOML": 0,
		"JSON": 1,
	}
)

func (x I18NFormat) Enum() *I18NFormat {
	p := new(I18NFormat)
	*p = x
	return p
}

func (x I18NFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (I18NFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_conf_proto_enumTypes[1].Descriptor()
}

func (I18NFormat) Type() protoreflect.EnumType {
	return &file_conf_proto_enumTypes[1]
}

func (x I18NFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use I18NFormat.Descriptor instead.
func (I18NFormat) EnumDescriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Environment   Environment            `protobuf:"varint,1,opt,name=environment,proto3,enum=conf.Environment" json:"environment,omitempty"`
	Log           *Log                   `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	Server        *Server                `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Auth          *Auth                  `protobuf:"bytes,5,opt,name=auth,proto3" json:"auth,omitempty"`
	I18N          *I18N                  `protobuf:"bytes,6,opt,name=i18n,proto3" json:"i18n,omitempty"`
	Agent         *Agent                 `protobuf:"bytes,7,opt,name=agent,proto3" json:"agent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetEnvironment() Environment {
	if x != nil {
		return x.Environment
	}
	return Environment_UNKNOWN
}

func (x *Bootstrap) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Bootstrap) GetI18N() *I18N {
	if x != nil {
		return x.I18N
	}
	return nil
}

func (x *Bootstrap) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type Log struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Level             string                 `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Format            string                 `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	Output            string                 `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	DisableCaller     bool                   `protobuf:"varint,4,opt,name=disable_caller,json=disableCaller,proto3" json:"disable_caller,omitempty"`
	DisableStacktrace bool                   `protobuf:"varint,5,opt,name=disable_stacktrace,json=disableStacktrace,proto3" json:"disable_stacktrace,omitempty"`
	EnableColor       bool                   `protobuf:"varint,6,opt,name=enable_color,json=enableColor,proto3" json:"enable_color,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Log) Reset() {
	*x = Log{}
	mi := &file_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Log) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Log) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Log) GetDisableCaller() bool {
	if x != nil {
		return x.DisableCaller
	}
	return false
}

func (x *Log) GetDisableStacktrace() bool {
	if x != nil {
		return x.DisableStacktrace
	}
	return false
}

func (x *Log) GetEnableColor() bool {
	if x != nil {
		return x.EnableColor
	}
	return false
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type Jwt struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Secret            string                 `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Expiration        *durationpb.Duration   `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Issuer            string                 `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	RefreshExpiration *durationpb.Duration   `protobuf:"bytes,4,opt,name=refresh_expiration,json=refreshExpiration,proto3" json:"refresh_expiration,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Jwt) Reset() {
	*x = Jwt{}
	mi := &file_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Jwt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jwt) ProtoMessage() {}

func (x *Jwt) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jwt.ProtoReflect.Descriptor instead.
func (*Jwt) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Jwt) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Jwt) GetExpiration() *durationpb.Duration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *Jwt) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Jwt) GetRefreshExpiration() *durationpb.Duration {
	if x != nil {
		return x.RefreshExpiration
	}
	return nil
}

type Auth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jwt           *Jwt                   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Oauth         *Auth_OAuth            `protobuf:"bytes,2,opt,name=oauth,proto3" json:"oauth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{5}
}

func (x *Auth) GetJwt() *Jwt {
	if x != nil {
		return x.Jwt
	}
	return nil
}

func (x *Auth) GetOauth() *Auth_OAuth {
	if x != nil {
		return x.Oauth
	}
	return nil
}

type I18N struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dir           string                 `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	Format        I18NFormat             `protobuf:"varint,2,opt,name=format,proto3,enum=conf.I18NFormat" json:"format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *I18N) Reset() {
	*x = I18N{}
	mi := &file_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *I18N) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I18N) ProtoMessage() {}

func (x *I18N) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I18N.ProtoReflect.Descriptor instead.
func (*I18N) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{6}
}

func (x *I18N) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *I18N) GetFormat() I18NFormat {
	if x != nil {
		return x.Format
	}
	return I18NFormat_TOML
}

type Agent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jwt           *Jwt                   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Agent) Reset() {
	*x = Agent{}
	mi := &file_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{7}
}

func (x *Agent) GetJwt() *Jwt {
	if x != nil {
		return x.Jwt
	}
	return nil
}

type Server_HTTP struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Network          string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr             string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout          *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	EnableCors       bool                   `protobuf:"varint,4,opt,name=enable_cors,json=enableCors,proto3" json:"enable_cors,omitempty"`
	AllowedOrigins   string                 `protobuf:"bytes,5,opt,name=allowed_origins,json=allowedOrigins,proto3" json:"allowed_origins,omitempty"`
	AllowedMethods   []string               `protobuf:"bytes,6,rep,name=allowed_methods,json=allowedMethods,proto3" json:"allowed_methods,omitempty"`
	AllowedHeaders   []string               `protobuf:"bytes,7,rep,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	ExposedHeaders   []string               `protobuf:"bytes,8,rep,name=exposed_headers,json=exposedHeaders,proto3" json:"exposed_headers,omitempty"`
	AllowCredentials bool                   `protobuf:"varint,9,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
	GrpcAddr         string                 `protobuf:"bytes,10,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetEnableCors() bool {
	if x != nil {
		return x.EnableCors
	}
	return false
}

func (x *Server_HTTP) GetAllowedOrigins() string {
	if x != nil {
		return x.AllowedOrigins
	}
	return ""
}

func (x *Server_HTTP) GetAllowedMethods() []string {
	if x != nil {
		return x.AllowedMethods
	}
	return nil
}

func (x *Server_HTTP) GetAllowedHeaders() []string {
	if x != nil {
		return x.AllowedHeaders
	}
	return nil
}

func (x *Server_HTTP) GetExposedHeaders() []string {
	if x != nil {
		return x.ExposedHeaders
	}
	return nil
}

func (x *Server_HTTP) GetAllowCredentials() bool {
	if x != nil {
		return x.AllowCredentials
	}
	return false
}

func (x *Server_HTTP) GetGrpcAddr() string {
	if x != nil {
		return x.GrpcAddr
	}
	return ""
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_conf_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Auth_OAuth struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	GoogleClientId     string                 `protobuf:"bytes,1,opt,name=google_client_id,json=googleClientId,proto3" json:"google_client_id,omitempty"`
	GoogleClientSecret string                 `protobuf:"bytes,2,opt,name=google_client_secret,json=googleClientSecret,proto3" json:"google_client_secret,omitempty"`
	GoogleRedirectUrl  string                 `protobuf:"bytes,3,opt,name=google_redirect_url,json=googleRedirectUrl,proto3" json:"google_redirect_url,omitempty"`
	GithubClientId     string                 `protobuf:"bytes,4,opt,name=github_client_id,json=githubClientId,proto3" json:"github_client_id,omitempty"`
	GithubClientSecret string                 `protobuf:"bytes,5,opt,name=github_client_secret,json=githubClientSecret,proto3" json:"github_client_secret,omitempty"`
	GithubRedirectUrl  string                 `protobuf:"bytes,6,opt,name=github_redirect_url,json=githubRedirectUrl,proto3" json:"github_redirect_url,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Auth_OAuth) Reset() {
	*x = Auth_OAuth{}
	mi := &file_conf_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_OAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_OAuth) ProtoMessage() {}

func (x *Auth_OAuth) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_OAuth.ProtoReflect.Descriptor instead.
func (*Auth_OAuth) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Auth_OAuth) GetGoogleClientId() string {
	if x != nil {
		return x.GoogleClientId
	}
	return ""
}

func (x *Auth_OAuth) GetGoogleClientSecret() string {
	if x != nil {
		return x.GoogleClientSecret
	}
	return ""
}

func (x *Auth_OAuth) GetGoogleRedirectUrl() string {
	if x != nil {
		return x.GoogleRedirectUrl
	}
	return ""
}

func (x *Auth_OAuth) GetGithubClientId() string {
	if x != nil {
		return x.GithubClientId
	}
	return ""
}

func (x *Auth_OAuth) GetGithubClientSecret() string {
	if x != nil {
		return x.GithubClientSecret
	}
	return ""
}

func (x *Auth_OAuth) GetGithubRedirectUrl() string {
	if x != nil {
		return x.GithubRedirectUrl
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

var file_conf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x12, 0x33, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c,
	0x6f, 0x67, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x31, 0x38, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x31,
	0x38, 0x6e, 0x52, 0x04, 0x69, 0x31, 0x38, 0x6e, 0x12, 0x21, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0xbc, 0x04, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x25, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04,
	0x68, 0x74, 0x74, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0xf8, 0x02, 0x0a, 0x04,
	0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x63, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78,
	0x70, 0x6f, 0x73, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x1a, 0x69, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0x73, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x12, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xed, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x03,
	0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x4a, 0x77, 0x74, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x05, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x1a, 0x9f, 0x02, 0x0a, 0x05, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x22, 0x42, 0x0a, 0x04, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x28, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x49, 0x31, 0x38, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x24, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4a, 0x77, 0x74, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x2a, 0x37, 0x0a,
	0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x56,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x52, 0x4f, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x45, 0x53, 0x54, 0x10, 0x03, 0x2a, 0x20, 0x0a, 0x0a, 0x49, 0x31, 0x38, 0x6e, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x4d, 0x4c, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x65, 0x72, 0x2d, 0x78, 0x2f, 0x69, 0x74,
	0x65, 0x72, 0x2d, 0x78, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData = file_conf_proto_rawDesc
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_proto_rawDescData)
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_conf_proto_goTypes = []any{
	(Environment)(0),            // 0: conf.Environment
	(I18NFormat)(0),             // 1: conf.I18nFormat
	(*Bootstrap)(nil),           // 2: conf.Bootstrap
	(*Log)(nil),                 // 3: conf.Log
	(*Server)(nil),              // 4: conf.Server
	(*Data)(nil),                // 5: conf.Data
	(*Jwt)(nil),                 // 6: conf.Jwt
	(*Auth)(nil),                // 7: conf.Auth
	(*I18N)(nil),                // 8: conf.I18n
	(*Agent)(nil),               // 9: conf.Agent
	(*Server_HTTP)(nil),         // 10: conf.Server.HTTP
	(*Server_GRPC)(nil),         // 11: conf.Server.GRPC
	(*Data_Database)(nil),       // 12: conf.Data.Database
	(*Auth_OAuth)(nil),          // 13: conf.Auth.OAuth
	(*durationpb.Duration)(nil), // 14: google.protobuf.Duration
}
var file_conf_proto_depIdxs = []int32{
	0,  // 0: conf.Bootstrap.environment:type_name -> conf.Environment
	3,  // 1: conf.Bootstrap.log:type_name -> conf.Log
	4,  // 2: conf.Bootstrap.server:type_name -> conf.Server
	5,  // 3: conf.Bootstrap.data:type_name -> conf.Data
	7,  // 4: conf.Bootstrap.auth:type_name -> conf.Auth
	8,  // 5: conf.Bootstrap.i18n:type_name -> conf.I18n
	9,  // 6: conf.Bootstrap.agent:type_name -> conf.Agent
	10, // 7: conf.Server.http:type_name -> conf.Server.HTTP
	11, // 8: conf.Server.grpc:type_name -> conf.Server.GRPC
	12, // 9: conf.Data.database:type_name -> conf.Data.Database
	14, // 10: conf.Jwt.expiration:type_name -> google.protobuf.Duration
	14, // 11: conf.Jwt.refresh_expiration:type_name -> google.protobuf.Duration
	6,  // 12: conf.Auth.jwt:type_name -> conf.Jwt
	13, // 13: conf.Auth.oauth:type_name -> conf.Auth.OAuth
	1,  // 14: conf.I18n.format:type_name -> conf.I18nFormat
	6,  // 15: conf.Agent.jwt:type_name -> conf.Jwt
	14, // 16: conf.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	14, // 17: conf.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		EnumInfos:         file_conf_proto_enumTypes,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_rawDesc = nil
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
