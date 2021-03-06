// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package primitives

import (
	"fmt"
)

type ProtocolVersion uint32

func (x ProtocolVersion) String() string {
  return fmt.Sprintf("%x", uint32(x))
}

func (x ProtocolVersion) Equal(y ProtocolVersion) bool {
  return x == y
}

func (x ProtocolVersion) KeyForMap() uint32 {
  return uint32(x)
}

type VirtualChainId uint32

func (x VirtualChainId) String() string {
  return fmt.Sprintf("%x", uint32(x))
}

func (x VirtualChainId) Equal(y VirtualChainId) bool {
  return x == y
}

func (x VirtualChainId) KeyForMap() uint32 {
  return uint32(x)
}

type BlockHeight uint64

func (x BlockHeight) String() string {
  return fmt.Sprintf("%x", uint64(x))
}

func (x BlockHeight) Equal(y BlockHeight) bool {
  return x == y
}

func (x BlockHeight) KeyForMap() uint64 {
  return uint64(x)
}

type TimestampNano uint64

func (x TimestampNano) String() string {
  return fmt.Sprintf("%x", uint64(x))
}

func (x TimestampNano) Equal(y TimestampNano) bool {
  return x == y
}

func (x TimestampNano) KeyForMap() uint64 {
  return uint64(x)
}

type ContractName string

func (x ContractName) String() string {
  return fmt.Sprintf(string(x))
}

func (x ContractName) Equal(y ContractName) bool {
  return x == y
}

func (x ContractName) KeyForMap() string {
  return string(x)
}

type MethodName string

func (x MethodName) String() string {
  return fmt.Sprintf(string(x))
}

func (x MethodName) Equal(y MethodName) bool {
  return x == y
}

func (x MethodName) KeyForMap() string {
  return string(x)
}

type ExecutionContextId uint32

func (x ExecutionContextId) String() string {
  return fmt.Sprintf("%x", uint32(x))
}

func (x ExecutionContextId) Equal(y ExecutionContextId) bool {
  return x == y
}

func (x ExecutionContextId) KeyForMap() uint32 {
  return uint32(x)
}


