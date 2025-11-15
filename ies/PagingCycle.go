package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PagingCycle_Enum_rf32  uper.Enumerated = 0
	PagingCycle_Enum_rf64  uper.Enumerated = 1
	PagingCycle_Enum_rf128 uper.Enumerated = 2
	PagingCycle_Enum_rf256 uper.Enumerated = 3
)

type PagingCycle struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PagingCycle) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PagingCycle", err)
	}
	return nil
}

func (ie *PagingCycle) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PagingCycle", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
