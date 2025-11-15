package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph60  uper.Enumerated = 0
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph80  uper.Enumerated = 1
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph100 uper.Enumerated = 2
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph120 uper.Enumerated = 3
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph140 uper.Enumerated = 4
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph160 uper.Enumerated = 5
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph180 uper.Enumerated = 6
	SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16_Enum_kmph200 uper.Enumerated = 7
)

type SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16", err)
	}
	return nil
}

func (ie *SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
