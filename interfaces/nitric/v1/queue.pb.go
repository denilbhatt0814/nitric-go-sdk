// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: queue/v1/queue.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to push a single event to a queue
type QueueSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Nitric name for the queue
	// this will automatically be resolved to the provider specific queue identifier.
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// The task to push to the queue
	Task *NitricTask `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *QueueSendRequest) Reset() {
	*x = QueueSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueSendRequest) ProtoMessage() {}

func (x *QueueSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueSendRequest.ProtoReflect.Descriptor instead.
func (*QueueSendRequest) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{0}
}

func (x *QueueSendRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *QueueSendRequest) GetTask() *NitricTask {
	if x != nil {
		return x.Task
	}
	return nil
}

// Result of pushing a single task to a queue
type QueueSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueueSendResponse) Reset() {
	*x = QueueSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueSendResponse) ProtoMessage() {}

func (x *QueueSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueSendResponse.ProtoReflect.Descriptor instead.
func (*QueueSendResponse) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{1}
}

type QueueSendBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Nitric name for the queue
	// this will automatically be resolved to the provider specific queue identifier.
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// Array of tasks to push to the queue
	Tasks []*NitricTask `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *QueueSendBatchRequest) Reset() {
	*x = QueueSendBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueSendBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueSendBatchRequest) ProtoMessage() {}

func (x *QueueSendBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueSendBatchRequest.ProtoReflect.Descriptor instead.
func (*QueueSendBatchRequest) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{2}
}

func (x *QueueSendBatchRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *QueueSendBatchRequest) GetTasks() []*NitricTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// Response for sending a collection of tasks
type QueueSendBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of tasks that failed to be queued
	FailedTasks []*FailedTask `protobuf:"bytes,1,rep,name=failedTasks,proto3" json:"failedTasks,omitempty"`
}

func (x *QueueSendBatchResponse) Reset() {
	*x = QueueSendBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueSendBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueSendBatchResponse) ProtoMessage() {}

func (x *QueueSendBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueSendBatchResponse.ProtoReflect.Descriptor instead.
func (*QueueSendBatchResponse) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{3}
}

func (x *QueueSendBatchResponse) GetFailedTasks() []*FailedTask {
	if x != nil {
		return x.FailedTasks
	}
	return nil
}

type QueueReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The nitric name for the queue
	// this will automatically be resolved to the provider specific queue identifier.
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// The max number of items to pop off the queue, may be capped by provider specific limitations
	Depth int32 `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
}

func (x *QueueReceiveRequest) Reset() {
	*x = QueueReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueReceiveRequest) ProtoMessage() {}

func (x *QueueReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueReceiveRequest.ProtoReflect.Descriptor instead.
func (*QueueReceiveRequest) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{4}
}

func (x *QueueReceiveRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *QueueReceiveRequest) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

type QueueReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of tasks popped off the queue
	Tasks []*NitricTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *QueueReceiveResponse) Reset() {
	*x = QueueReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueReceiveResponse) ProtoMessage() {}

func (x *QueueReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueReceiveResponse.ProtoReflect.Descriptor instead.
func (*QueueReceiveResponse) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{5}
}

func (x *QueueReceiveResponse) GetTasks() []*NitricTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type QueueCompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The nitric name for the queue
	//  this will automatically be resolved to the provider specific queue identifier.
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// Lease id of the task to be completed
	LeaseId string `protobuf:"bytes,2,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
}

func (x *QueueCompleteRequest) Reset() {
	*x = QueueCompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueCompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueCompleteRequest) ProtoMessage() {}

func (x *QueueCompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueCompleteRequest.ProtoReflect.Descriptor instead.
func (*QueueCompleteRequest) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{6}
}

func (x *QueueCompleteRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *QueueCompleteRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

type QueueCompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueueCompleteResponse) Reset() {
	*x = QueueCompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueCompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueCompleteResponse) ProtoMessage() {}

func (x *QueueCompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueCompleteResponse.ProtoReflect.Descriptor instead.
func (*QueueCompleteResponse) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{7}
}

type FailedTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The task that failed to be pushed
	Task *NitricTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	// A message describing the failure
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FailedTask) Reset() {
	*x = FailedTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailedTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailedTask) ProtoMessage() {}

func (x *FailedTask) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailedTask.ProtoReflect.Descriptor instead.
func (*FailedTask) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{8}
}

func (x *FailedTask) GetTask() *NitricTask {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *FailedTask) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A task to be sent or received from a queue.
type NitricTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique id for the task
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The lease id unique to the pop request, this must be used to complete, extend the lease or release the task.
	LeaseId string `protobuf:"bytes,2,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	// A content hint for the tasks payload
	PayloadType string `protobuf:"bytes,3,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
	// The payload of the task
	Payload *structpb.Struct `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *NitricTask) Reset() {
	*x = NitricTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_v1_queue_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NitricTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NitricTask) ProtoMessage() {}

func (x *NitricTask) ProtoReflect() protoreflect.Message {
	mi := &file_queue_v1_queue_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NitricTask.ProtoReflect.Descriptor instead.
func (*NitricTask) Descriptor() ([]byte, []int) {
	return file_queue_v1_queue_proto_rawDescGZIP(), []int{9}
}

func (x *NitricTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NitricTask) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *NitricTask) GetPayloadType() string {
	if x != nil {
		return x.PayloadType
	}
	return ""
}

func (x *NitricTask) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_queue_v1_queue_proto protoreflect.FileDescriptor

var file_queue_v1_queue_proto_rawDesc = []byte{
	0x0a, 0x14, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x22, 0x13, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x57, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x41, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65,
	0x70, 0x74, 0x68, 0x22, 0x49, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x69, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x69, 0x74,
	0x72, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x47,
	0x0a, 0x14, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x57, 0x0a, 0x0a, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2f,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e,
	0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x69, 0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x4e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xe7, 0x02, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x21, 0x2e, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x26, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x62, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x06, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x0c, 0x6e, 0x69, 0x74, 0x72, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xaa, 0x02, 0x15, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0xca,
	0x02, 0x15, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queue_v1_queue_proto_rawDescOnce sync.Once
	file_queue_v1_queue_proto_rawDescData = file_queue_v1_queue_proto_rawDesc
)

func file_queue_v1_queue_proto_rawDescGZIP() []byte {
	file_queue_v1_queue_proto_rawDescOnce.Do(func() {
		file_queue_v1_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_v1_queue_proto_rawDescData)
	})
	return file_queue_v1_queue_proto_rawDescData
}

var file_queue_v1_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_queue_v1_queue_proto_goTypes = []interface{}{
	(*QueueSendRequest)(nil),       // 0: nitric.queue.v1.QueueSendRequest
	(*QueueSendResponse)(nil),      // 1: nitric.queue.v1.QueueSendResponse
	(*QueueSendBatchRequest)(nil),  // 2: nitric.queue.v1.QueueSendBatchRequest
	(*QueueSendBatchResponse)(nil), // 3: nitric.queue.v1.QueueSendBatchResponse
	(*QueueReceiveRequest)(nil),    // 4: nitric.queue.v1.QueueReceiveRequest
	(*QueueReceiveResponse)(nil),   // 5: nitric.queue.v1.QueueReceiveResponse
	(*QueueCompleteRequest)(nil),   // 6: nitric.queue.v1.QueueCompleteRequest
	(*QueueCompleteResponse)(nil),  // 7: nitric.queue.v1.QueueCompleteResponse
	(*FailedTask)(nil),             // 8: nitric.queue.v1.FailedTask
	(*NitricTask)(nil),             // 9: nitric.queue.v1.NitricTask
	(*structpb.Struct)(nil),        // 10: google.protobuf.Struct
}
var file_queue_v1_queue_proto_depIdxs = []int32{
	9,  // 0: nitric.queue.v1.QueueSendRequest.task:type_name -> nitric.queue.v1.NitricTask
	9,  // 1: nitric.queue.v1.QueueSendBatchRequest.tasks:type_name -> nitric.queue.v1.NitricTask
	8,  // 2: nitric.queue.v1.QueueSendBatchResponse.failedTasks:type_name -> nitric.queue.v1.FailedTask
	9,  // 3: nitric.queue.v1.QueueReceiveResponse.tasks:type_name -> nitric.queue.v1.NitricTask
	9,  // 4: nitric.queue.v1.FailedTask.task:type_name -> nitric.queue.v1.NitricTask
	10, // 5: nitric.queue.v1.NitricTask.payload:type_name -> google.protobuf.Struct
	0,  // 6: nitric.queue.v1.Queue.Send:input_type -> nitric.queue.v1.QueueSendRequest
	2,  // 7: nitric.queue.v1.Queue.SendBatch:input_type -> nitric.queue.v1.QueueSendBatchRequest
	4,  // 8: nitric.queue.v1.Queue.Receive:input_type -> nitric.queue.v1.QueueReceiveRequest
	6,  // 9: nitric.queue.v1.Queue.Complete:input_type -> nitric.queue.v1.QueueCompleteRequest
	1,  // 10: nitric.queue.v1.Queue.Send:output_type -> nitric.queue.v1.QueueSendResponse
	3,  // 11: nitric.queue.v1.Queue.SendBatch:output_type -> nitric.queue.v1.QueueSendBatchResponse
	5,  // 12: nitric.queue.v1.Queue.Receive:output_type -> nitric.queue.v1.QueueReceiveResponse
	7,  // 13: nitric.queue.v1.Queue.Complete:output_type -> nitric.queue.v1.QueueCompleteResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_queue_v1_queue_proto_init() }
func file_queue_v1_queue_proto_init() {
	if File_queue_v1_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queue_v1_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueSendRequest); i {
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
		file_queue_v1_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueSendResponse); i {
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
		file_queue_v1_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueSendBatchRequest); i {
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
		file_queue_v1_queue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueSendBatchResponse); i {
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
		file_queue_v1_queue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueReceiveRequest); i {
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
		file_queue_v1_queue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueReceiveResponse); i {
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
		file_queue_v1_queue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueCompleteRequest); i {
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
		file_queue_v1_queue_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueCompleteResponse); i {
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
		file_queue_v1_queue_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailedTask); i {
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
		file_queue_v1_queue_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NitricTask); i {
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
			RawDescriptor: file_queue_v1_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queue_v1_queue_proto_goTypes,
		DependencyIndexes: file_queue_v1_queue_proto_depIdxs,
		MessageInfos:      file_queue_v1_queue_proto_msgTypes,
	}.Build()
	File_queue_v1_queue_proto = out.File
	file_queue_v1_queue_proto_rawDesc = nil
	file_queue_v1_queue_proto_goTypes = nil
	file_queue_v1_queue_proto_depIdxs = nil
}
