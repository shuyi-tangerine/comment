// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package base

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

type RPCRequest struct {
}

func NewRPCRequest() *RPCRequest {
	return &RPCRequest{}
}

func (p *RPCRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(ctx, fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RPCRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RPCRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RPCRequest) Equals(other *RPCRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	return true
}

func (p *RPCRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RPCRequest(%+v)", *p)
}

// Attributes:
//   - Code
//   - Message
type RPCResponse struct {
	Code    int64  `thrift:"code,1,required" db:"code" json:"code"`
	Message string `thrift:"message,2,required" db:"message" json:"message"`
}

func NewRPCResponse() *RPCResponse {
	return &RPCResponse{}
}

func (p *RPCResponse) GetCode() int64 {
	return p.Code
}

func (p *RPCResponse) GetMessage() string {
	return p.Message
}
func (p *RPCResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetCode bool = false
	var issetMessage bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
				issetCode = true
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
				issetMessage = true
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetCode {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Code is not set"))
	}
	if !issetMessage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"))
	}
	return nil
}

func (p *RPCResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Code = v
	}
	return nil
}

func (p *RPCResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *RPCResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RPCResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RPCResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "code", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err)
	}
	return err
}

func (p *RPCResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err)
	}
	return err
}

func (p *RPCResponse) Equals(other *RPCResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Code != other.Code {
		return false
	}
	if p.Message != other.Message {
		return false
	}
	return true
}

func (p *RPCResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RPCResponse(%+v)", *p)
}
