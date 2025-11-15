package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ExtendedPagingCycle_r17_Enum_rf256  uper.Enumerated = 0
	ExtendedPagingCycle_r17_Enum_rf512  uper.Enumerated = 1
	ExtendedPagingCycle_r17_Enum_rf1024 uper.Enumerated = 2
	ExtendedPagingCycle_r17_Enum_spare1 uper.Enumerated = 3
)

type ExtendedPagingCycle_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ExtendedPagingCycle_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ExtendedPagingCycle_r17", err)
	}
	return nil
}

func (ie *ExtendedPagingCycle_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ExtendedPagingCycle_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
