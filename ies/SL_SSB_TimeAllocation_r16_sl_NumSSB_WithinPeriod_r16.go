package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n1  uper.Enumerated = 0
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n2  uper.Enumerated = 1
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n4  uper.Enumerated = 2
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n8  uper.Enumerated = 3
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n16 uper.Enumerated = 4
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n32 uper.Enumerated = 5
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n64 uper.Enumerated = 6
)

type SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16", err)
	}
	return nil
}

func (ie *SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
