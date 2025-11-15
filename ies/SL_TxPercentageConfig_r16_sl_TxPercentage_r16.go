package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxPercentageConfig_r16_sl_TxPercentage_r16_Enum_p20 uper.Enumerated = 0
	SL_TxPercentageConfig_r16_sl_TxPercentage_r16_Enum_p35 uper.Enumerated = 1
	SL_TxPercentageConfig_r16_sl_TxPercentage_r16_Enum_p50 uper.Enumerated = 2
)

type SL_TxPercentageConfig_r16_sl_TxPercentage_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SL_TxPercentageConfig_r16_sl_TxPercentage_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SL_TxPercentageConfig_r16_sl_TxPercentage_r16", err)
	}
	return nil
}

func (ie *SL_TxPercentageConfig_r16_sl_TxPercentage_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SL_TxPercentageConfig_r16_sl_TxPercentage_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
