// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20240319094008-0393e58bdf10
// source: authzed/api/v1/debug.proto

package v1

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	durationpb1 "github.com/planetscale/vtprotobuf/types/known/durationpb"
	structpb1 "github.com/planetscale/vtprotobuf/types/known/structpb"
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *DebugInformation) CloneVT() *DebugInformation {
	if m == nil {
		return (*DebugInformation)(nil)
	}
	r := new(DebugInformation)
	r.Check = m.Check.CloneVT()
	r.SchemaUsed = m.SchemaUsed
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DebugInformation) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CheckDebugTrace_SubProblems) CloneVT() *CheckDebugTrace_SubProblems {
	if m == nil {
		return (*CheckDebugTrace_SubProblems)(nil)
	}
	r := new(CheckDebugTrace_SubProblems)
	if rhs := m.Traces; rhs != nil {
		tmpContainer := make([]*CheckDebugTrace, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Traces = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CheckDebugTrace_SubProblems) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CheckDebugTrace) CloneVT() *CheckDebugTrace {
	if m == nil {
		return (*CheckDebugTrace)(nil)
	}
	r := new(CheckDebugTrace)
	r.Resource = m.Resource.CloneVT()
	r.Permission = m.Permission
	r.PermissionType = m.PermissionType
	r.Subject = m.Subject.CloneVT()
	r.Result = m.Result
	r.CaveatEvaluationInfo = m.CaveatEvaluationInfo.CloneVT()
	r.Duration = (*durationpb.Duration)((*durationpb1.Duration)(m.Duration).CloneVT())
	r.OptionalExpiresAt = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.OptionalExpiresAt).CloneVT())
	if m.Resolution != nil {
		r.Resolution = m.Resolution.(interface {
			CloneVT() isCheckDebugTrace_Resolution
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CheckDebugTrace) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *CheckDebugTrace_WasCachedResult) CloneVT() isCheckDebugTrace_Resolution {
	if m == nil {
		return (*CheckDebugTrace_WasCachedResult)(nil)
	}
	r := new(CheckDebugTrace_WasCachedResult)
	r.WasCachedResult = m.WasCachedResult
	return r
}

func (m *CheckDebugTrace_SubProblems_) CloneVT() isCheckDebugTrace_Resolution {
	if m == nil {
		return (*CheckDebugTrace_SubProblems_)(nil)
	}
	r := new(CheckDebugTrace_SubProblems_)
	r.SubProblems = m.SubProblems.CloneVT()
	return r
}

func (m *CaveatEvalInfo) CloneVT() *CaveatEvalInfo {
	if m == nil {
		return (*CaveatEvalInfo)(nil)
	}
	r := new(CaveatEvalInfo)
	r.Expression = m.Expression
	r.Result = m.Result
	r.Context = (*structpb.Struct)((*structpb1.Struct)(m.Context).CloneVT())
	r.PartialCaveatInfo = m.PartialCaveatInfo.CloneVT()
	r.CaveatName = m.CaveatName
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CaveatEvalInfo) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *DebugInformation) EqualVT(that *DebugInformation) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Check.EqualVT(that.Check) {
		return false
	}
	if this.SchemaUsed != that.SchemaUsed {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DebugInformation) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DebugInformation)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CheckDebugTrace_SubProblems) EqualVT(that *CheckDebugTrace_SubProblems) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Traces) != len(that.Traces) {
		return false
	}
	for i, vx := range this.Traces {
		vy := that.Traces[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &CheckDebugTrace{}
			}
			if q == nil {
				q = &CheckDebugTrace{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CheckDebugTrace_SubProblems) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CheckDebugTrace_SubProblems)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CheckDebugTrace) EqualVT(that *CheckDebugTrace) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Resolution == nil && that.Resolution != nil {
		return false
	} else if this.Resolution != nil {
		if that.Resolution == nil {
			return false
		}
		if !this.Resolution.(interface {
			EqualVT(isCheckDebugTrace_Resolution) bool
		}).EqualVT(that.Resolution) {
			return false
		}
	}
	if !this.Resource.EqualVT(that.Resource) {
		return false
	}
	if this.Permission != that.Permission {
		return false
	}
	if this.PermissionType != that.PermissionType {
		return false
	}
	if !this.Subject.EqualVT(that.Subject) {
		return false
	}
	if this.Result != that.Result {
		return false
	}
	if !this.CaveatEvaluationInfo.EqualVT(that.CaveatEvaluationInfo) {
		return false
	}
	if !(*durationpb1.Duration)(this.Duration).EqualVT((*durationpb1.Duration)(that.Duration)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.OptionalExpiresAt).EqualVT((*timestamppb1.Timestamp)(that.OptionalExpiresAt)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CheckDebugTrace) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CheckDebugTrace)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CheckDebugTrace_WasCachedResult) EqualVT(thatIface isCheckDebugTrace_Resolution) bool {
	that, ok := thatIface.(*CheckDebugTrace_WasCachedResult)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.WasCachedResult != that.WasCachedResult {
		return false
	}
	return true
}

func (this *CheckDebugTrace_SubProblems_) EqualVT(thatIface isCheckDebugTrace_Resolution) bool {
	that, ok := thatIface.(*CheckDebugTrace_SubProblems_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.SubProblems, that.SubProblems; p != q {
		if p == nil {
			p = &CheckDebugTrace_SubProblems{}
		}
		if q == nil {
			q = &CheckDebugTrace_SubProblems{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *CaveatEvalInfo) EqualVT(that *CaveatEvalInfo) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Expression != that.Expression {
		return false
	}
	if this.Result != that.Result {
		return false
	}
	if !(*structpb1.Struct)(this.Context).EqualVT((*structpb1.Struct)(that.Context)) {
		return false
	}
	if !this.PartialCaveatInfo.EqualVT(that.PartialCaveatInfo) {
		return false
	}
	if this.CaveatName != that.CaveatName {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CaveatEvalInfo) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*CaveatEvalInfo)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *DebugInformation) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugInformation) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *DebugInformation) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.SchemaUsed) > 0 {
		i -= len(m.SchemaUsed)
		copy(dAtA[i:], m.SchemaUsed)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SchemaUsed)))
		i--
		dAtA[i] = 0x12
	}
	if m.Check != nil {
		size, err := m.Check.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CheckDebugTrace_SubProblems) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckDebugTrace_SubProblems) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CheckDebugTrace_SubProblems) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Traces) > 0 {
		for iNdEx := len(m.Traces) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Traces[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CheckDebugTrace) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckDebugTrace) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CheckDebugTrace) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Resolution.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.OptionalExpiresAt != nil {
		size, err := (*timestamppb1.Timestamp)(m.OptionalExpiresAt).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x52
	}
	if m.Duration != nil {
		size, err := (*durationpb1.Duration)(m.Duration).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x4a
	}
	if m.CaveatEvaluationInfo != nil {
		size, err := m.CaveatEvaluationInfo.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x42
	}
	if m.Result != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x28
	}
	if m.Subject != nil {
		size, err := m.Subject.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.PermissionType != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.PermissionType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Permission) > 0 {
		i -= len(m.Permission)
		copy(dAtA[i:], m.Permission)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Permission)))
		i--
		dAtA[i] = 0x12
	}
	if m.Resource != nil {
		size, err := m.Resource.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CheckDebugTrace_WasCachedResult) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CheckDebugTrace_WasCachedResult) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.WasCachedResult {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *CheckDebugTrace_SubProblems_) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CheckDebugTrace_SubProblems_) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SubProblems != nil {
		size, err := m.SubProblems.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *CaveatEvalInfo) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CaveatEvalInfo) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CaveatEvalInfo) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.CaveatName) > 0 {
		i -= len(m.CaveatName)
		copy(dAtA[i:], m.CaveatName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.CaveatName)))
		i--
		dAtA[i] = 0x2a
	}
	if m.PartialCaveatInfo != nil {
		size, err := m.PartialCaveatInfo.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.Context != nil {
		size, err := (*structpb1.Struct)(m.Context).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Result != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Expression) > 0 {
		i -= len(m.Expression)
		copy(dAtA[i:], m.Expression)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Expression)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DebugInformation) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Check != nil {
		l = m.Check.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.SchemaUsed)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *CheckDebugTrace_SubProblems) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Traces) > 0 {
		for _, e := range m.Traces {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *CheckDebugTrace) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resource != nil {
		l = m.Resource.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Permission)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.PermissionType != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.PermissionType))
	}
	if m.Subject != nil {
		l = m.Subject.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Result))
	}
	if vtmsg, ok := m.Resolution.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.CaveatEvaluationInfo != nil {
		l = m.CaveatEvaluationInfo.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Duration != nil {
		l = (*durationpb1.Duration)(m.Duration).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.OptionalExpiresAt != nil {
		l = (*timestamppb1.Timestamp)(m.OptionalExpiresAt).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *CheckDebugTrace_WasCachedResult) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *CheckDebugTrace_SubProblems_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubProblems != nil {
		l = m.SubProblems.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *CaveatEvalInfo) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Expression)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Result))
	}
	if m.Context != nil {
		l = (*structpb1.Struct)(m.Context).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.PartialCaveatInfo != nil {
		l = m.PartialCaveatInfo.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.CaveatName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *DebugInformation) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DebugInformation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugInformation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Check", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Check == nil {
				m.Check = &CheckDebugTrace{}
			}
			if err := m.Check.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckDebugTrace_SubProblems) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckDebugTrace_SubProblems: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckDebugTrace_SubProblems: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Traces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Traces = append(m.Traces, &CheckDebugTrace{})
			if err := m.Traces[len(m.Traces)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckDebugTrace) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckDebugTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckDebugTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &ObjectReference{}
			}
			if err := m.Resource.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permission = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermissionType", wireType)
			}
			m.PermissionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PermissionType |= CheckDebugTrace_PermissionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subject == nil {
				m.Subject = &SubjectReference{}
			}
			if err := m.Subject.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= CheckDebugTrace_Permissionship(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasCachedResult", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Resolution = &CheckDebugTrace_WasCachedResult{WasCachedResult: b}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubProblems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Resolution.(*CheckDebugTrace_SubProblems_); ok {
				if err := oneof.SubProblems.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &CheckDebugTrace_SubProblems{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Resolution = &CheckDebugTrace_SubProblems_{SubProblems: v}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaveatEvaluationInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CaveatEvaluationInfo == nil {
				m.CaveatEvaluationInfo = &CaveatEvalInfo{}
			}
			if err := m.CaveatEvaluationInfo.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Duration == nil {
				m.Duration = &durationpb.Duration{}
			}
			if err := (*durationpb1.Duration)(m.Duration).UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalExpiresAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OptionalExpiresAt == nil {
				m.OptionalExpiresAt = &timestamppb.Timestamp{}
			}
			if err := (*timestamppb1.Timestamp)(m.OptionalExpiresAt).UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CaveatEvalInfo) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CaveatEvalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CaveatEvalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expression", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expression = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= CaveatEvalInfo_Result(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &structpb.Struct{}
			}
			if err := (*structpb1.Struct)(m.Context).UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartialCaveatInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PartialCaveatInfo == nil {
				m.PartialCaveatInfo = &PartialCaveatInfo{}
			}
			if err := m.PartialCaveatInfo.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaveatName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CaveatName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
