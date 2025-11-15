package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms100  uper.Enumerated = 0
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms200  uper.Enumerated = 1
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms300  uper.Enumerated = 2
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms400  uper.Enumerated = 3
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms500  uper.Enumerated = 4
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms600  uper.Enumerated = 5
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms700  uper.Enumerated = 6
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms800  uper.Enumerated = 7
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms900  uper.Enumerated = 8
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_ms1000 uper.Enumerated = 9
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_spare6 uper.Enumerated = 10
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_spare5 uper.Enumerated = 11
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_spare4 uper.Enumerated = 12
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_spare3 uper.Enumerated = 13
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_spare2 uper.Enumerated = 14
	SL_PeriodCG_r16_sl_PeriodCG1_r16_Enum_spare1 uper.Enumerated = 15
)

type SL_PeriodCG_r16_sl_PeriodCG1_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SL_PeriodCG_r16_sl_PeriodCG1_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SL_PeriodCG_r16_sl_PeriodCG1_r16", err)
	}
	return nil
}

func (ie *SL_PeriodCG_r16_sl_PeriodCG1_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SL_PeriodCG_r16_sl_PeriodCG1_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
