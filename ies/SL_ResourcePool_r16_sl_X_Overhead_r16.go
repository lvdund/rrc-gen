package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ResourcePool_r16_sl_X_Overhead_r16_Enum_n0 uper.Enumerated = 0
	SL_ResourcePool_r16_sl_X_Overhead_r16_Enum_n3 uper.Enumerated = 1
	SL_ResourcePool_r16_sl_X_Overhead_r16_Enum_n6 uper.Enumerated = 2
	SL_ResourcePool_r16_sl_X_Overhead_r16_Enum_n9 uper.Enumerated = 3
)

type SL_ResourcePool_r16_sl_X_Overhead_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SL_ResourcePool_r16_sl_X_Overhead_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SL_ResourcePool_r16_sl_X_Overhead_r16", err)
	}
	return nil
}

func (ie *SL_ResourcePool_r16_sl_X_Overhead_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SL_ResourcePool_r16_sl_X_Overhead_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
