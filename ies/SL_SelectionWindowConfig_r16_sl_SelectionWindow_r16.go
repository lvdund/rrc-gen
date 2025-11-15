package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n1  uper.Enumerated = 0
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n5  uper.Enumerated = 1
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n10 uper.Enumerated = 2
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n20 uper.Enumerated = 3
)

type SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16", err)
	}
	return nil
}

func (ie *SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
