// Code generated by protoc-gen-go.
// source: customer.proto
// DO NOT EDIT!

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	customer.proto

It has these top-level messages:
	CustomerRequest
	CustomerResponse
*/
package customer

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type CustomerRequest struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email     string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Iphone    string                     `protobuf:"bytes,4,opt,name=iphone" json:"iphone,omitempty"`
	Addresses []*CustomerRequest_Address `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *CustomerRequest) Reset()         { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()    {}

func (m *CustomerRequest) GetAddresses() []*CustomerRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type CustomerRequest_Address struct {
	Street            string `protobuf:"bytes,1,opt,name=street" json:"street,omitempty"`
	City              string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	State             string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Zip               string `protobuf:"bytes,4,opt,name=zip" json:"zip,omitempty"`
	IsShippingAddress bool   `protobuf:"varint,5,opt,name=isShippingAddress" json:"isShippingAddress,omitempty"`
}

func (m *CustomerRequest_Address) Reset()         { *m = CustomerRequest_Address{} }
func (m *CustomerRequest_Address) String() string { return proto.CompactTextString(m) }
func (*CustomerRequest_Address) ProtoMessage()    {}

type CustomerResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *CustomerResponse) Reset()         { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()    {}

func init() {
}
