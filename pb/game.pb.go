// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: game.proto

package pb

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

type Vector2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             float64                `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y             float64                `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vector2) Reset() {
	*x = Vector2{}
	mi := &file_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vector2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector2) ProtoMessage() {}

func (x *Vector2) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector2.ProtoReflect.Descriptor instead.
func (*Vector2) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Vector2) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector2) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Controls struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Speed         float64                `protobuf:"fixed64,1,opt,name=speed,proto3" json:"speed,omitempty"`                          // Horizontal movement speed
	JumpForce     float64                `protobuf:"fixed64,2,opt,name=jump_force,json=jumpForce,proto3" json:"jump_force,omitempty"` // Force applied when jumping
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Controls) Reset() {
	*x = Controls{}
	mi := &file_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Controls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controls) ProtoMessage() {}

func (x *Controls) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controls.ProtoReflect.Descriptor instead.
func (*Controls) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Controls) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Controls) GetJumpForce() float64 {
	if x != nil {
		return x.JumpForce
	}
	return 0
}

type InputState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MoveLeft      bool                   `protobuf:"varint,1,opt,name=move_left,json=moveLeft,proto3" json:"move_left,omitempty"`    // Indicates if the player is moving left
	MoveRight     bool                   `protobuf:"varint,2,opt,name=move_right,json=moveRight,proto3" json:"move_right,omitempty"` // Indicates if the player is moving right
	Jump          bool                   `protobuf:"varint,3,opt,name=jump,proto3" json:"jump,omitempty"`                            // Indicates if the player is jumping
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputState) Reset() {
	*x = InputState{}
	mi := &file_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputState) ProtoMessage() {}

func (x *InputState) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputState.ProtoReflect.Descriptor instead.
func (*InputState) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *InputState) GetMoveLeft() bool {
	if x != nil {
		return x.MoveLeft
	}
	return false
}

func (x *InputState) GetMoveRight() bool {
	if x != nil {
		return x.MoveRight
	}
	return false
}

func (x *InputState) GetJump() bool {
	if x != nil {
		return x.Jump
	}
	return false
}

type GameObject struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                     // Unique identifier for the object
	Position      *Vector2               `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`         // Current position
	Velocity      *Vector2               `protobuf:"bytes,3,opt,name=velocity,proto3" json:"velocity,omitempty"`         // Current velocity
	Acceleration  *Vector2               `protobuf:"bytes,4,opt,name=acceleration,proto3" json:"acceleration,omitempty"` // Current acceleration
	Mass          float64                `protobuf:"fixed64,5,opt,name=mass,proto3" json:"mass,omitempty"`               // Mass of the object
	Elasticity    float64                `protobuf:"fixed64,6,opt,name=elasticity,proto3" json:"elasticity,omitempty"`   // Elasticity (bounciness)
	Friction      float64                `protobuf:"fixed64,7,opt,name=friction,proto3" json:"friction,omitempty"`       // Friction coefficient
	Controls      *Controls              `protobuf:"bytes,8,opt,name=controls,proto3" json:"controls,omitempty"`         // Optional, player-specific controls
	Grounded      bool                   `protobuf:"varint,9,opt,name=Grounded,proto3" json:"Grounded,omitempty"`        // Indicates if the player is on the ground
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameObject) Reset() {
	*x = GameObject{}
	mi := &file_game_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameObject) ProtoMessage() {}

func (x *GameObject) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameObject.ProtoReflect.Descriptor instead.
func (*GameObject) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *GameObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GameObject) GetPosition() *Vector2 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *GameObject) GetVelocity() *Vector2 {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *GameObject) GetAcceleration() *Vector2 {
	if x != nil {
		return x.Acceleration
	}
	return nil
}

func (x *GameObject) GetMass() float64 {
	if x != nil {
		return x.Mass
	}
	return 0
}

func (x *GameObject) GetElasticity() float64 {
	if x != nil {
		return x.Elasticity
	}
	return 0
}

func (x *GameObject) GetFriction() float64 {
	if x != nil {
		return x.Friction
	}
	return 0
}

func (x *GameObject) GetControls() *Controls {
	if x != nil {
		return x.Controls
	}
	return nil
}

func (x *GameObject) GetGrounded() bool {
	if x != nil {
		return x.Grounded
	}
	return false
}

type PhysicsEngineState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Objects       []*GameObject          `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"` // List of all game objects
	Gravity       *Vector2               `protobuf:"bytes,2,opt,name=gravity,proto3" json:"gravity,omitempty"` // Gravity vector applied to all objects
	Input         *InputState            `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`     // Current input state (optional for debugging or sync)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PhysicsEngineState) Reset() {
	*x = PhysicsEngineState{}
	mi := &file_game_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PhysicsEngineState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicsEngineState) ProtoMessage() {}

func (x *PhysicsEngineState) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicsEngineState.ProtoReflect.Descriptor instead.
func (*PhysicsEngineState) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *PhysicsEngineState) GetObjects() []*GameObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *PhysicsEngineState) GetGravity() *Vector2 {
	if x != nil {
		return x.Gravity
	}
	return nil
}

func (x *PhysicsEngineState) GetInput() *InputState {
	if x != nil {
		return x.Input
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0x25, 0x0a, 0x07, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x3f, 0x0a, 0x08, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6a,
	0x75, 0x6d, 0x70, 0x5f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6a, 0x75, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x5c, 0x0a, 0x0a, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x76, 0x65,
	0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x6f, 0x76,
	0x65, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x75, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x6a, 0x75, 0x6d, 0x70, 0x22, 0xbd, 0x02, 0x0a, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x32, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x32, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6d, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x50, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x67,
	0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x07, 0x67, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x17, 0x5a, 0x15,
	0x62, 0x6f, 0x75, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x6c, 0x6c, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_game_proto_goTypes = []any{
	(*Vector2)(nil),            // 0: game.Vector2
	(*Controls)(nil),           // 1: game.Controls
	(*InputState)(nil),         // 2: game.InputState
	(*GameObject)(nil),         // 3: game.GameObject
	(*PhysicsEngineState)(nil), // 4: game.PhysicsEngineState
}
var file_game_proto_depIdxs = []int32{
	0, // 0: game.GameObject.position:type_name -> game.Vector2
	0, // 1: game.GameObject.velocity:type_name -> game.Vector2
	0, // 2: game.GameObject.acceleration:type_name -> game.Vector2
	1, // 3: game.GameObject.controls:type_name -> game.Controls
	3, // 4: game.PhysicsEngineState.objects:type_name -> game.GameObject
	0, // 5: game.PhysicsEngineState.gravity:type_name -> game.Vector2
	2, // 6: game.PhysicsEngineState.input:type_name -> game.InputState
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
