package core

import (
	"fmt"
)

const (
	OpcodeEOF    = 0
	OpcodeSInt   = 1
	OpcodeUInt   = 2
	OpcodeFloat  = 3
	OpcodeDouble = 4
	OpcodeString = 5
)

type ModuleTypeHandler interface {
	ReadByte() (byte, error)
	ReadFull(buf []byte) error
	ReadString() ([]byte, error)
	ReadLength() (uint64, bool, error)
}

type moduleTypeHandlerImpl struct {
	dec *Decoder
}

func (m moduleTypeHandlerImpl) ReadByte() (byte, error) {
	return m.dec.readByte()
}

func (m moduleTypeHandlerImpl) ReadFull(buf []byte) error {
	return m.dec.readFull(buf)
}

func (m moduleTypeHandlerImpl) ReadString() ([]byte, error) {
	return m.dec.readString()
}

func (m moduleTypeHandlerImpl) ReadLength() (uint64, bool, error) {
	return m.dec.readLength()
}

type ModuleTypeHandleFunc func(handler ModuleTypeHandler, encVersion int) (interface{}, error)

func (dec *Decoder) readModuleType() (string, interface{}, error) {
	moduleId, _, err := dec.readLength()
	if err != nil {
		return "", nil, err
	}
	return dec.handleModuleType(moduleId)
}

func (dec *Decoder) handleModuleType(moduleId uint64) (string, interface{}, error) {
	moduleType := moduleTypeNameByID(moduleId)
	handler, found := dec.withSpecialTypes[moduleType]
	if !found {
		return moduleType, nil, fmt.Errorf("unknown module type: %s", moduleType)
	}
	encVersion := moduleId & 1023
	val, err := handler(moduleTypeHandlerImpl{dec: dec}, int(encVersion))
	return moduleType, val, err
}

func moduleTypeNameByID(moduleid uint64) string {

	cset := ModuleTypeNameCharSet

	name := make([]byte, 9)

	moduleid >>= 10
	for j := 0; j < 9; j++ {
		name[8-j] = cset[moduleid&63]
		moduleid >>= 6
	}

	return string(name)
}

const ModuleTypeNameCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
