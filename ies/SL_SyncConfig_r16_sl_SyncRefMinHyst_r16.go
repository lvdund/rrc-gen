package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SyncConfig_r16_sl_SyncRefMinHyst_r16_Enum_dB0  uper.Enumerated = 0
	SL_SyncConfig_r16_sl_SyncRefMinHyst_r16_Enum_dB3  uper.Enumerated = 1
	SL_SyncConfig_r16_sl_SyncRefMinHyst_r16_Enum_dB6  uper.Enumerated = 2
	SL_SyncConfig_r16_sl_SyncRefMinHyst_r16_Enum_dB9  uper.Enumerated = 3
	SL_SyncConfig_r16_sl_SyncRefMinHyst_r16_Enum_dB12 uper.Enumerated = 4
)

type SL_SyncConfig_r16_sl_SyncRefMinHyst_r16 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SL_SyncConfig_r16_sl_SyncRefMinHyst_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SL_SyncConfig_r16_sl_SyncRefMinHyst_r16", err)
	}
	return nil
}

func (ie *SL_SyncConfig_r16_sl_SyncRefMinHyst_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SL_SyncConfig_r16_sl_SyncRefMinHyst_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
