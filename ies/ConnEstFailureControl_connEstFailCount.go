package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConnEstFailureControl_connEstFailCount_Enum_n1 uper.Enumerated = 0
	ConnEstFailureControl_connEstFailCount_Enum_n2 uper.Enumerated = 1
	ConnEstFailureControl_connEstFailCount_Enum_n3 uper.Enumerated = 2
	ConnEstFailureControl_connEstFailCount_Enum_n4 uper.Enumerated = 3
)

type ConnEstFailureControl_connEstFailCount struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ConnEstFailureControl_connEstFailCount) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ConnEstFailureControl_connEstFailCount", err)
	}
	return nil
}

func (ie *ConnEstFailureControl_connEstFailCount) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ConnEstFailureControl_connEstFailCount", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
