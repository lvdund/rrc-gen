package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s30  uper.Enumerated = 0
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s60  uper.Enumerated = 1
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s120 uper.Enumerated = 2
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s240 uper.Enumerated = 3
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s300 uper.Enumerated = 4
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s420 uper.Enumerated = 5
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s600 uper.Enumerated = 6
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s900 uper.Enumerated = 7
)

type ConnEstFailureControl_connEstFailOffsetValidity struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ConnEstFailureControl_connEstFailOffsetValidity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ConnEstFailureControl_connEstFailOffsetValidity", err)
	}
	return nil
}

func (ie *ConnEstFailureControl_connEstFailOffsetValidity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ConnEstFailureControl_connEstFailOffsetValidity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
