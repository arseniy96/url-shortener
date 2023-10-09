// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: src/proto/shortener.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CreateLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserSession string `protobuf:"bytes,2,opt,name=user_session,json=userSession,proto3" json:"user_session,omitempty"`
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLinkRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateLinkRequest) GetUserSession() string {
	if x != nil {
		return x.UserSession
	}
	return ""
}

type CreateLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateLinkResponse) Reset() {
	*x = CreateLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkResponse) ProtoMessage() {}

func (x *CreateLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateLinkResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLinkResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CreateLinksBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls        []*CreateLinksBatchRequestNested `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	UserSession string                           `protobuf:"bytes,2,opt,name=user_session,json=userSession,proto3" json:"user_session,omitempty"`
}

func (x *CreateLinksBatchRequest) Reset() {
	*x = CreateLinksBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinksBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinksBatchRequest) ProtoMessage() {}

func (x *CreateLinksBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinksBatchRequest.ProtoReflect.Descriptor instead.
func (*CreateLinksBatchRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLinksBatchRequest) GetUrls() []*CreateLinksBatchRequestNested {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *CreateLinksBatchRequest) GetUserSession() string {
	if x != nil {
		return x.UserSession
	}
	return ""
}

type CreateLinksBatchRequestNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CreateLinksBatchRequestNested) Reset() {
	*x = CreateLinksBatchRequestNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinksBatchRequestNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinksBatchRequestNested) ProtoMessage() {}

func (x *CreateLinksBatchRequestNested) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinksBatchRequestNested.ProtoReflect.Descriptor instead.
func (*CreateLinksBatchRequestNested) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *CreateLinksBatchRequestNested) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateLinksBatchRequestNested) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateLinksBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*CreateLinksBatchResponseNested `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *CreateLinksBatchResponse) Reset() {
	*x = CreateLinksBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinksBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinksBatchResponse) ProtoMessage() {}

func (x *CreateLinksBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinksBatchResponse.ProtoReflect.Descriptor instead.
func (*CreateLinksBatchResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *CreateLinksBatchResponse) GetUrls() []*CreateLinksBatchResponseNested {
	if x != nil {
		return x.Urls
	}
	return nil
}

type CreateLinksBatchResponseNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CreateLinksBatchResponseNested) Reset() {
	*x = CreateLinksBatchResponseNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinksBatchResponseNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinksBatchResponseNested) ProtoMessage() {}

func (x *CreateLinksBatchResponseNested) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinksBatchResponseNested.ProtoReflect.Descriptor instead.
func (*CreateLinksBatchResponseNested) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *CreateLinksBatchResponseNested) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateLinksBatchResponseNested) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ResolveLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ResolveLinkRequest) Reset() {
	*x = ResolveLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveLinkRequest) ProtoMessage() {}

func (x *ResolveLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveLinkRequest.ProtoReflect.Descriptor instead.
func (*ResolveLinkRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *ResolveLinkRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ResolveLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *ResolveLinkResponse) Reset() {
	*x = ResolveLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveLinkResponse) ProtoMessage() {}

func (x *ResolveLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveLinkResponse.ProtoReflect.Descriptor instead.
func (*ResolveLinkResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *ResolveLinkResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type UserUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSession string `protobuf:"bytes,1,opt,name=user_session,json=userSession,proto3" json:"user_session,omitempty"`
}

func (x *UserUrlsRequest) Reset() {
	*x = UserUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsRequest) ProtoMessage() {}

func (x *UserUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsRequest.ProtoReflect.Descriptor instead.
func (*UserUrlsRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *UserUrlsRequest) GetUserSession() string {
	if x != nil {
		return x.UserSession
	}
	return ""
}

type UserUrlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*UserUrlsResponseNested `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *UserUrlsResponse) Reset() {
	*x = UserUrlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsResponse) ProtoMessage() {}

func (x *UserUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsResponse.ProtoReflect.Descriptor instead.
func (*UserUrlsResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *UserUrlsResponse) GetUrls() []*UserUrlsResponseNested {
	if x != nil {
		return x.Urls
	}
	return nil
}

type UserUrlsResponseNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *UserUrlsResponseNested) Reset() {
	*x = UserUrlsResponseNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_shortener_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUrlsResponseNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsResponseNested) ProtoMessage() {}

func (x *UserUrlsResponseNested) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_shortener_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsResponseNested.ProtoReflect.Descriptor instead.
func (*UserUrlsResponseNested) Descriptor() ([]byte, []int) {
	return file_src_proto_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *UserUrlsResponseNested) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *UserUrlsResponseNested) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_src_proto_shortener_proto protoreflect.FileDescriptor

var file_src_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x48, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x7f, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x1d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x5e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x64, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x31, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22,
	0x38, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x4e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22,
	0x58, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x32, 0xb6, 0x03, 0x0a, 0x0e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x41, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x21, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x27, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x22, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x12,
	0x1f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_shortener_proto_rawDescOnce sync.Once
	file_src_proto_shortener_proto_rawDescData = file_src_proto_shortener_proto_rawDesc
)

func file_src_proto_shortener_proto_rawDescGZIP() []byte {
	file_src_proto_shortener_proto_rawDescOnce.Do(func() {
		file_src_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_shortener_proto_rawDescData)
	})
	return file_src_proto_shortener_proto_rawDescData
}

var file_src_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_src_proto_shortener_proto_goTypes = []interface{}{
	(*PingRequest)(nil),                    // 0: shortenerproto.PingRequest
	(*PingResponse)(nil),                   // 1: shortenerproto.PingResponse
	(*CreateLinkRequest)(nil),              // 2: shortenerproto.CreateLinkRequest
	(*CreateLinkResponse)(nil),             // 3: shortenerproto.CreateLinkResponse
	(*CreateLinksBatchRequest)(nil),        // 4: shortenerproto.CreateLinksBatchRequest
	(*CreateLinksBatchRequestNested)(nil),  // 5: shortenerproto.CreateLinksBatchRequestNested
	(*CreateLinksBatchResponse)(nil),       // 6: shortenerproto.CreateLinksBatchResponse
	(*CreateLinksBatchResponseNested)(nil), // 7: shortenerproto.CreateLinksBatchResponseNested
	(*ResolveLinkRequest)(nil),             // 8: shortenerproto.ResolveLinkRequest
	(*ResolveLinkResponse)(nil),            // 9: shortenerproto.ResolveLinkResponse
	(*UserUrlsRequest)(nil),                // 10: shortenerproto.UserUrlsRequest
	(*UserUrlsResponse)(nil),               // 11: shortenerproto.UserUrlsResponse
	(*UserUrlsResponseNested)(nil),         // 12: shortenerproto.UserUrlsResponseNested
}
var file_src_proto_shortener_proto_depIdxs = []int32{
	5,  // 0: shortenerproto.CreateLinksBatchRequest.urls:type_name -> shortenerproto.CreateLinksBatchRequestNested
	7,  // 1: shortenerproto.CreateLinksBatchResponse.urls:type_name -> shortenerproto.CreateLinksBatchResponseNested
	12, // 2: shortenerproto.UserUrlsResponse.urls:type_name -> shortenerproto.UserUrlsResponseNested
	0,  // 3: shortenerproto.ShortenerProto.Ping:input_type -> shortenerproto.PingRequest
	2,  // 4: shortenerproto.ShortenerProto.CreateLink:input_type -> shortenerproto.CreateLinkRequest
	4,  // 5: shortenerproto.ShortenerProto.CreateLinksBatch:input_type -> shortenerproto.CreateLinksBatchRequest
	8,  // 6: shortenerproto.ShortenerProto.ResolveLink:input_type -> shortenerproto.ResolveLinkRequest
	10, // 7: shortenerproto.ShortenerProto.UserUrls:input_type -> shortenerproto.UserUrlsRequest
	1,  // 8: shortenerproto.ShortenerProto.Ping:output_type -> shortenerproto.PingResponse
	3,  // 9: shortenerproto.ShortenerProto.CreateLink:output_type -> shortenerproto.CreateLinkResponse
	6,  // 10: shortenerproto.ShortenerProto.CreateLinksBatch:output_type -> shortenerproto.CreateLinksBatchResponse
	9,  // 11: shortenerproto.ShortenerProto.ResolveLink:output_type -> shortenerproto.ResolveLinkResponse
	11, // 12: shortenerproto.ShortenerProto.UserUrls:output_type -> shortenerproto.UserUrlsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_src_proto_shortener_proto_init() }
func file_src_proto_shortener_proto_init() {
	if File_src_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinksBatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinksBatchRequestNested); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinksBatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinksBatchResponseNested); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveLinkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveLinkResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUrlsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUrlsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_shortener_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUrlsResponseNested); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_shortener_proto_goTypes,
		DependencyIndexes: file_src_proto_shortener_proto_depIdxs,
		MessageInfos:      file_src_proto_shortener_proto_msgTypes,
	}.Build()
	File_src_proto_shortener_proto = out.File
	file_src_proto_shortener_proto_rawDesc = nil
	file_src_proto_shortener_proto_goTypes = nil
	file_src_proto_shortener_proto_depIdxs = nil
}
